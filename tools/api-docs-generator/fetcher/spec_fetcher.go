package fetcher

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

func FetchSpec(ctx context.Context, cfg config.Config, directory string) error {
	// #nosec G107 // cfg.Fetcher.Source is a URL from config and does not contain user input
	resp, err := get(ctx, cfg.Fetcher.Source)
	if err != nil {
		return err
	}

	var versions []string

	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}

	gaVersion := getLatestGAVersion(versions)

	specPath, err := url.JoinPath(cfg.Fetcher.Source, gaVersion)
	if err != nil {
		return err
	}

	// #nosec G107 // specPath is a URL from config and does not contain user input
	resp, err = get(ctx, specPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	jsonSpec, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	formattedSpec := bytes.NewBufferString("")
	err = json.Indent(formattedSpec, jsonSpec, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(directory, "docs", cfg.Fetcher.Destination), formattedSpec.Bytes(), 0o644)
}

func getLatestGAVersion(versions []string) string {
	gaVersions := []string{}
	for _, version := range versions {
		if !strings.Contains(version, "~") {
			gaVersions = append(gaVersions, version)
		}
	}
	sort.Strings(gaVersions)
	return gaVersions[len(gaVersions)-1]
}

func get(ctx context.Context, urlToGet string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlToGet, http.NoBody)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
