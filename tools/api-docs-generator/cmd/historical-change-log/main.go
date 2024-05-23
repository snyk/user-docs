package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	md "github.com/nao1215/markdown"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/load"
)

func main() {
	// 1. Fetch existing GA specs, less than a config date
	// 2. oas diff to generate log
	// 3. Write the changelog file

	err := Generate("docs/snyk-api/changelog.md", "2024-04-29")
	if err != nil {
		log.Panic(err)
	}
}

// Change shadow declaration because I cant import Change from the original package formatters/changes.go
type Change struct {
	Id          string
	Text        string
	Comment     string
	Level       checker.Level
	Operation   string
	OperationId string
	Path        string
	Source      string
	Section     string
	IsBreaking  bool
}

type Changes []Change

func Generate(changeLogFileName, endVersion string) error {
	versions, err := fetchAllVersions()
	if err != nil {
		return err
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	gaVersions := gaVersions(versions)
	endVersionPos := sort.SearchStrings(gaVersions, endVersion)
	gaVersions = gaVersions[:endVersionPos]
	slices.Reverse(gaVersions)
	nextVersion := gaVersions[0]

	writer, err := os.Create(changeLogFileName)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "# API Change Log\n\n")
	if err != nil {
		return err
	}
	for _, baseVersion := range gaVersions[1:] {
		nextURL := fmt.Sprintf("https://api.snyk.io/rest/openapi/%s", nextVersion)
		baseURL := fmt.Sprintf("https://api.snyk.io/rest/openapi/%s", baseVersion)
		nextVersion = baseVersion
		groupedChanges, _, err := getChangeLog(nextURL, baseURL, loader)
		if err != nil {
			return err
		}

		if len(groupedChanges) == 0 {
			continue
		}
		markdown := md.NewMarkdown(writer)
		markdown.H2(fmt.Sprintf("%s\n", baseVersion))

		for op, changes := range groupedChanges {
			changesVal := *changes
			if len(changesVal) == 1 && changesVal[0].Text == "endpoint added" {
				markdown.H3f("%s - %s - Added", op.Operation, op.Path)
				parsedUrl, err := url.Parse(nextURL)
				if err != nil {
					return err
				}
				document, err := loader.LoadFromURI(parsedUrl)
				if err != nil {
					return err
				}
				markdown.BulletList(replaceWithCodeQuotes(document.Paths.Find(op.Path).Operations()[op.Operation].Description))
			} else {
				markdown.H3f("%s - %s - Updated", op.Operation, op.Path)
				for _, change := range changesVal {
					if change.IsBreaking {
						markdown.BulletList(replaceWithCodeQuotes(change.Text))
						markdown.YellowBadge("Breaking")
					} else {
						markdown.BulletList(fmt.Sprintf("%s\n", replaceWithCodeQuotes(change.Text)))
					}
				}
			}
			markdown.PlainText("\n")
		}
		err = markdown.Build()

		if err != nil {
			return err
		}
	}

	return writer.Close()
}

func replaceWithCodeQuotes(description string) string {
	return strings.ReplaceAll(description, "'", "`")
}

type Endpoint struct {
	Path      string
	Operation string
}

type ChangesByEndpoint map[Endpoint]*Changes

func GroupChanges(changes checker.Changes, l checker.Localizer) ChangesByEndpoint {
	apiChanges := ChangesByEndpoint{}

	for _, change := range changes {
		switch change.(type) {
		case checker.ApiChange:
			ep := Endpoint{Path: change.GetPath(), Operation: change.GetOperation()}
			if c, ok := apiChanges[ep]; ok {
				*c = append(*c, Change{
					IsBreaking: change.IsBreaking(),
					Text:       change.GetUncolorizedText(l),
				})
			} else {
				apiChanges[ep] = &Changes{Change{
					IsBreaking: change.IsBreaking(),
					Text:       change.GetUncolorizedText(l),
				}}
			}
		}
	}

	return apiChanges
}

func getChangeLog(nextVersionURI, baseVersionURI string, loader *openapi3.Loader) (ChangesByEndpoint, *diff.Diff, error) {
	baseVersionURL, err := url.Parse(baseVersionURI)
	if err != nil {
		return nil, nil, err
	}

	nextVersionURL, err := url.Parse(nextVersionURI)
	if err != nil {
		return nil, nil, err
	}

	flattenAllOf := load.WithFlattenAllOf()
	flattenParams := load.WithFlattenParams()
	lowerHeaderNames := load.WithLowercaseHeaders()

	s1, err := load.NewSpecInfo(loader, load.NewSource(baseVersionURL.String()), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return nil, nil, err
	}
	s2, err := load.NewSpecInfo(loader, load.NewSource(nextVersionURL.String()), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return nil, nil, err
	}

	diffReport, sourcesMap, err := diff.GetWithOperationsSourcesMap(diff.NewConfig(), s1, s2)
	changes := checker.CheckBackwardCompatibilityUntilLevel(checker.GetDefaultChecks(), diffReport, sourcesMap, checker.INFO)

	groupedChanges := GroupChanges(changes, checker.NewLocalizer("en"))
	if err != nil {
		return nil, nil, err
	}

	return groupedChanges, diffReport, nil
}

func gaVersions(versions []string) []string {
	var gaVersions []string
	for _, version := range versions {
		if !strings.Contains(version, "~") {
			gaVersions = append(gaVersions, version)
		}
	}
	return gaVersions
}

func fetchAllVersions() ([]string, error) {
	versionsResponse, err := http.Get("https://api.snyk.io/rest/openapi")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(versionsResponse.Body)

	if err != nil {
		return nil, err
	}

	var versions []string
	err = json.NewDecoder(versionsResponse.Body).Decode(&versions)
	return versions, err
}
