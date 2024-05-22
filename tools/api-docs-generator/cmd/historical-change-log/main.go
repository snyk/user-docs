package main

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/formatters"
	"github.com/tufin/oasdiff/load"
	"github.com/tufin/oasdiff/report"
	"log"
	"net/http"
	"net/url"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	// 1. Fetch existing GA specs, less than a config date
	// 2. oas diff to generate log
	// 3. Write the changelog file

	err := Generate("docs/snyk-api/CHANGELOG.md", "2024-04-29")
	if err != nil {
		log.Panic(err)
	}

}

func Generate(changeLogFileName string, endVersion string) error {
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
	baseVersion := gaVersions[0]
	writer, err := os.Create(changeLogFileName)
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, "# Change Log\n\n")

	for _, nextVersion := range gaVersions[1:5] {
		baseVersionURI := fmt.Sprintf("https://api.snyk.io/rest/openapi/%s", baseVersion)
		nextVersionURI := fmt.Sprintf("https://api.snyk.io/rest/openapi/%s", nextVersion)
		baseVersion = nextVersion
		diffReportText, err := getChangeLog(baseVersionURI, nextVersionURI, loader)

		fmt.Fprintf(writer, "## Version %s\n\n", nextVersion)
		fmt.Printf("## Version %s\n", nextVersion)

		_, err = writer.Write(diffReportText)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "\n\n")
		if err != nil {
			return err
		}
	}
	return writer.Close()
}

func getChangeLog(nextVersionURI, baseVersionURI string, loader *openapi3.Loader) ([]byte, error) {
	baseVersionUrl, err := url.Parse(baseVersionURI)
	if err != nil {
		return nil, err
	}

	nextVersionUrl, err := url.Parse(nextVersionURI)
	if err != nil {
		return nil, err
	}

	//baseOpenAPISpec, err := loader.LoadFromURI(baseVersionUrl)
	//if err != nil {
	//	return nil, err
	//}
	//
	//nextOpenAPISpec, err := loader.LoadFromURI(nextVersionUrl)
	//if err != nil {
	//	return nil, err
	//}

	//diffReport, err := diff.Get(diff.NewConfig(), baseOpenAPISpec, nextOpenAPISpec)
	//if err != nil {
	//	return nil, err
	//}
	formatter, err := formatters.Lookup(string(formatters.FormatText), formatters.DefaultFormatterOpts())
	if err != nil {
		return nil, err
	}

	flattenAllOf := load.WithFlattenAllOf()
	flattenParams := load.WithFlattenParams()
	lowerHeaderNames := load.WithLowercaseHeaders()

	s1, err := load.NewSpecInfo(loader, load.NewSource(baseVersionUrl.String()), flattenAllOf, flattenParams, lowerHeaderNames)

	s2, err := load.NewSpecInfo(loader, load.NewSource(nextVersionUrl.String()), flattenAllOf, flattenParams, lowerHeaderNames)

	pair := load.NewSpecInfoPair(s1, s2)
	diffReport, sourcesMap, err := diff.GetWithOperationsSourcesMap(diff.NewConfig(), s1, s2)
	changes := checker.CheckBackwardCompatibilityUntilLevel(checker.GetDefaultChecks(), diffReport, sourcesMap, checker.INFO)
	_, err = formatter.RenderChangelog(changes, formatters.NewRenderOpts(), pair)

	diffReportMarkdown := report.GetTextReportAsBytes(diffReport)
	if err != nil {
		return nil, err
	}
	return diffReportMarkdown, nil
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
	if err != nil {
		return nil, err
	}
	var versions []string
	err = json.NewDecoder(versionsResponse.Body).Decode(&versions)
	return versions, err
}
