package versions

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

func GetCurrentVersions(ctx context.Context, cfg *config.Config) ([]string, error) {
	// #nosec G107 // cfg.Fetcher.Source is a URL from config and does not contain user input
	resp, err := Get(ctx, cfg.Fetcher.Source)
	if err != nil {
		return nil, err
	}

	var versions []string

	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return versions, nil
}

func Get(ctx context.Context, urlToGet string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlToGet, http.NoBody)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func ExtractGAVersions(versions []string) []string {
	var gaVersions []string
	for _, version := range versions {
		if !strings.Contains(version, "~") {
			gaVersions = append(gaVersions, version)
		}
	}
	return gaVersions
}

func GetLatestGAVersion(versions []string) string {
	gaVersions := []string{}
	for _, version := range versions {
		if !strings.Contains(version, "~") {
			gaVersions = append(gaVersions, version)
		}
	}
	sort.Strings(gaVersions)
	return gaVersions[len(gaVersions)-1]
}
