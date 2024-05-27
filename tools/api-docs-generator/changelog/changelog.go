package changelog

import (
	"context"
	"fmt"
	"net/url"
	"os"
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

type ChangesByEndpoint map[Endpoint]checker.Changes

func UpdateChangelog(ctx context.Context, cfg *config.Config, changeLogFileName string) (string, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	allVersions, err := versions.GetCurrentVersions(ctx, cfg)
	if err != nil {
		return "", err
	}

	latestGAVersion := versions.GetLatestGAVersion(allVersions)

	historicalChangelog, err := os.ReadFile(changeLogFileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	writer, err := os.Create(changeLogFileName)
	if err != nil {
		return "", err
	}

	defer func(writer *os.File) {
		writeErr := writer.Close()
		if writeErr != nil {
			fmt.Printf("Error closing writer for %s\n", historicalChangelog)
		}
	}(writer)

	nextURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, latestGAVersion)
	baseURL := cfg.Fetcher.Destination

	groupedChanges, err := getChangeLog(nextURL, baseURL, loader)
	if err != nil {
		return "", err
	}

	if len(groupedChanges) == 0 {
		return "", nil
	}

	markdown := md.NewMarkdown(writer)

	err = WriteToChangeLog(markdown, groupedChanges, latestGAVersion, nextURL, cfg.Changelog.LastSyncDate)
	if err != nil {
		return "", err
	}

	err = markdown.Build()
	if err != nil {
		return "", err
	}

	_, err = writer.Write(historicalChangelog)
	if err != nil {
		return "", err
	}

	return latestGAVersion, err
}

func GenerateHistorical(ctx context.Context, cfg *config.Config, changeLogFileName, endVersion string) error {
	allVersions, err := versions.GetCurrentVersions(ctx, cfg)
	if err != nil {
		return err
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	gaVersions := versions.ExtractGAVersions(allVersions)
	endVersionPos := sort.SearchStrings(gaVersions, endVersion)
	gaVersions = gaVersions[:endVersionPos]
	slices.Reverse(gaVersions)
	nextVersion := gaVersions[0]

	writer, err := os.Create(changeLogFileName)
	if err != nil {
		return err
	}

	for _, baseVersion := range gaVersions[1:] {
		nextURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, nextVersion)
		baseURL := fmt.Sprintf("%s/%s", cfg.Fetcher.Source, baseVersion)
		nextVersion = baseVersion
		groupedChanges, err := getChangeLog(nextURL, baseURL, loader)
		if err != nil {
			return err
		}

		groupedChanges = filterChanges(groupedChanges)

		if len(groupedChanges) == 0 {
			continue
		}

		markdown := md.NewMarkdown(writer)

		err = WriteToChangeLog(markdown, groupedChanges, baseVersion, nextURL, cfg.Changelog.LastSyncDate)
		if err != nil {
			return err
		}

		err = markdown.Build()
		if err != nil {
			return err
		}
	}

	return writer.Close()
}
func filterChanges(changes ChangesByEndpoint) ChangesByEndpoint {
	filteredChanges := ChangesByEndpoint{}
	for endpoint, change := range changes {
		if strings.HasPrefix(endpoint.Path, "/openapi") {
			continue
		}
		filteredChanges[endpoint] = change
	}
	return filteredChanges
}

func WriteToChangeLog(markdown *md.Markdown, groupedChanges ChangesByEndpoint, baseVersion, nextVersionURL, lastSyncVersion string) error {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	localizer := checker.NewLocalizer("en")

	if baseVersion == lastSyncVersion {
		markdown.H2(fmt.Sprintf("%s - Updated %s\n", baseVersion, time.Now().Format("2006-01-02")))
	} else {
		markdown.H2(fmt.Sprintf("%s\n", baseVersion))
	}

	groupedChanges = filterChanges(groupedChanges)

	for op, changes := range groupedChanges {
		if len(changes) == 1 && changes[0].GetUncolorizedText(localizer) == "endpoint added" {
			markdown.H3f("%s - `%s` - Added", op.Operation, op.Path)
			parsedURL, err := url.Parse(nextVersionURL)
			if err != nil {
				return err
			}
			document, err := loader.LoadFromURI(parsedURL)
			if err != nil {
				return err
			}
			markdown.BulletList(versions.ReplaceWithCodeQuotes(document.Paths.Find(op.Path).Operations()[op.Operation].Description))
			continue
		}
		markdown.H3f("%s - Updated %s", op.Operation, op.Path)
		for index := range changes {
			if changes[index].IsBreaking() {
				markdown.BulletList(versions.ReplaceWithCodeQuotes(changes[index].GetUncolorizedText(localizer)))
				markdown.YellowBadge("Breaking")
			} else {
				markdown.BulletList(fmt.Sprintf("%s\n", versions.ReplaceWithCodeQuotes(changes[index].GetUncolorizedText(localizer))))
			}
		}

		markdown.PlainText("\n")
	}
	return nil
}

func groupChanges(changes checker.Changes) ChangesByEndpoint {
	apiChanges := ChangesByEndpoint{}

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

	return apiChanges
}

func getChangeLog(nextVersionURI, baseVersionURI string, loader *openapi3.Loader) (ChangesByEndpoint, error) {
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
