package changelog

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	md "github.com/nao1215/markdown"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/load"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
	"github.com/snyk/user-docs/tools/api-docs-generator/versions"
)

type Endpoint struct {
	Path      string
	Operation string
}

type ChangesByEndpoint struct {
	Endpoint
	checker.Changes
}

func UpdateChangelog(ctx context.Context, cfg *config.Config, syncStateCfg config.SyncStateConfig, docsDirectory string) error {
	changeLogFile := path.Join(docsDirectory, cfg.Changelog.ChangelogFile)
	syncStateFile := path.Join(docsDirectory, cfg.Changelog.SyncStateFile)

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	allVersions, err := versions.Find(ctx, cfg)
	if err != nil {
		return err
	}

	latestGAVersion := allVersions.LatestGA()

	nextURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, latestGAVersion)
	baseURL := path.Join(docsDirectory, cfg.Fetcher.Destination)

	groupedChanges, err := getChangeLog(nextURL, baseURL, loader)
	if err != nil {
		return err
	}

	if len(groupedChanges) == 0 {
		// skip all logic if no actual changes detected
		return nil
	}

	// changes detected
	historicalChangelog, err := os.ReadFile(changeLogFile)
	if err != nil {
		return err
	}

	writer, err := os.Create(changeLogFile)
	if err != nil {
		return err
	}

	defer func(writer *os.File) {
		writeErr := writer.Close()
		if writeErr != nil {
			fmt.Printf("Error closing writer for %s\n", changeLogFile)
		}
	}(writer)

	markdown := md.NewMarkdown(writer)
	err = writeToChangeLog(markdown, groupedChanges, latestGAVersion, nextURL, syncStateCfg.LastSyncedVersion)
	if err != nil {
		return err
	}

	err = markdown.Build()
	if err != nil {
		return err
	}

	_, err = writer.Write(historicalChangelog)
	if err != nil {
		return err
	}

	syncStateCfg.LastSyncedVersion = latestGAVersion
	err = config.UpdateSyncState(syncStateFile, syncStateCfg)

	return err
}

func GenerateHistorical(ctx context.Context, cfg *config.Config, docsDirectory string) error {
	allVersions, err := versions.Find(ctx, cfg)
	if err != nil {
		return err
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	gaVersions := allVersions.FilterGA()
	endVersionPos := sort.SearchStrings(gaVersions, cfg.Changelog.HistoricalVersionCutoff)
	gaVersions = gaVersions[:endVersionPos]
	slices.Reverse(gaVersions)
	nextVersion := gaVersions[0]

	writer, err := os.Create(path.Join(docsDirectory, cfg.Changelog.ChangelogFile))
	if err != nil {
		return err
	}

	for _, baseVersion := range gaVersions[1:] {
		nextURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, nextVersion)
		baseURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, baseVersion)

		groupedChanges, err := getChangeLog(nextURL, baseURL, loader)
		if err != nil {
			return err
		}

		groupedChanges = filterChanges(groupedChanges)

		if len(groupedChanges) != 0 {
			markdown := md.NewMarkdown(writer)

			err = writeToChangeLog(markdown, groupedChanges, nextVersion, nextURL, "")
			if err != nil {
				return err
			}

			err = markdown.Build()
			if err != nil {
				return err
			}
		}
		nextVersion = baseVersion
	}

	return writer.Close()
}
func filterChanges(changes []ChangesByEndpoint) []ChangesByEndpoint {
	filteredChanges := make([]ChangesByEndpoint, 0, len(changes))
	for _, endpointChange := range changes {
		if strings.HasPrefix(endpointChange.Path, "/openapi") {
			continue
		}
		filteredChanges = append(filteredChanges, endpointChange)
	}
	return filteredChanges
}

func writeToChangeLog(markdown *md.Markdown, groupedChanges []ChangesByEndpoint, baseVersion, nextVersionURL, lastSyncVersion string) error {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	localizer := checker.NewLocalizer("en")

	if baseVersion == lastSyncVersion {
		markdown.H2(fmt.Sprintf("%s - Updated %s\n", baseVersion, time.Now().Format("2006-01-02")))
	} else {
		markdown.H2(fmt.Sprintf("%s\n", baseVersion))
	}

	groupedChanges = filterChanges(groupedChanges)

	// sort for stability
	sort.Slice(groupedChanges, func(i, j int) bool {
		return groupedChanges[i].Path+groupedChanges[i].Operation > groupedChanges[j].Path+groupedChanges[j].Operation
	})

	for _, changeGroup := range groupedChanges {
		if len(changeGroup.Changes) == 1 && changeGroup.Changes[0].GetUncolorizedText(localizer) == "endpoint added" {
			markdown.H3f("%s - `%s` - Added", changeGroup.Operation, changeGroup.Path)
			parsedURL, err := url.Parse(nextVersionURL)
			if err != nil {
				return err
			}
			document, err := loader.LoadFromURI(parsedURL)
			if err != nil {
				return err
			}
			markdown.BulletList(normalizeQuotes(document.Paths.Find(changeGroup.Path).Operations()[changeGroup.Operation].Description))
		} else {
			markdown.H3f("%s - `%s` - Updated", changeGroup.Operation, changeGroup.Path)
			for index := range changeGroup.Changes {
				writeOperationChangeDetails(markdown, changeGroup, index, localizer)
			}
		}
		markdown.PlainText("\n")
	}
	return nil
}

func writeOperationChangeDetails(markdown *md.Markdown, changeGroup ChangesByEndpoint, index int, localizer checker.Localizer) {
	if changeGroup.Changes[index].IsBreaking() {
		markdown.BulletList(normalizeQuotes(changeGroup.Changes[index].GetUncolorizedText(localizer)))
		markdown.YellowBadge("Breaking")
	} else {
		markdown.BulletList(fmt.Sprintf("%s\n", normalizeQuotes(changeGroup.Changes[index].GetUncolorizedText(localizer))))
	}
}

func groupChanges(changes checker.Changes) []ChangesByEndpoint {
	apiChanges := map[Endpoint]checker.Changes{}

	for _, change := range changes {
		if change, ok := change.(checker.ApiChange); ok {
			ep := Endpoint{Path: change.GetPath(), Operation: change.GetOperation()}
			if changesForEndpoint, ok := apiChanges[ep]; ok {
				apiChanges[ep] = append(changesForEndpoint, change)
			} else {
				apiChanges[ep] = checker.Changes{change}
			}
		}
	}

	apiChangesSlice := make([]ChangesByEndpoint, 0, len(apiChanges))
	for endpoint, changes := range apiChanges {
		apiChangesSlice = append(apiChangesSlice, ChangesByEndpoint{Endpoint: endpoint, Changes: changes})
	}

	return apiChangesSlice
}

func getChangeLog(nextVersionURI, baseVersionURI string, loader *openapi3.Loader) ([]ChangesByEndpoint, error) {
	baseVersionURL, err := url.Parse(baseVersionURI)
	if err != nil {
		return nil, err
	}

	nextVersionURL, err := url.Parse(nextVersionURI)
	if err != nil {
		return nil, err
	}

	flattenAllOf := load.WithFlattenAllOf()
	flattenParams := load.WithFlattenParams()
	lowerHeaderNames := load.WithLowercaseHeaders()

	s1, err := load.NewSpecInfo(loader, load.NewSource(baseVersionURL.String()), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return nil, err
	}
	s2, err := load.NewSpecInfo(loader, load.NewSource(nextVersionURL.String()), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return nil, err
	}

	diffReport, sourcesMap, err := diff.GetWithOperationsSourcesMap(diff.NewConfig(), s1, s2)
	changes := checker.CheckBackwardCompatibilityUntilLevel(checker.GetDefaultChecks(), diffReport, sourcesMap, checker.INFO)

	groupedChanges := groupChanges(changes)
	if err != nil {
		return nil, err
	}

	return groupedChanges, nil
}

func normalizeQuotes(description string) string {
	return strings.ReplaceAll(description, "'", "`")
}
