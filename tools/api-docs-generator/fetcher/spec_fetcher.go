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

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
	"github.com/snyk/user-docs/tools/api-docs-generator/versions"
)

func FetchSpec(ctx context.Context, cfg *config.Config, directory string) error {
	allVersions, err := versions.GetCurrentVersions(ctx, cfg)
	if err != nil {
		return err
	}

	gaVersion := versions.GetLatestGAVersion(allVersions)

	formattedSpec, err := getSpecByVersion(ctx, cfg, gaVersion)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(directory, cfg.Fetcher.Destination), formattedSpec.Bytes(), 0o644)
}

func get(ctx context.Context, urlToGet string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlToGet, http.NoBody)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func getSpecByVersion(ctx context.Context, cfg *config.Config, version string) (*bytes.Buffer, error) {
	specPath, err := url.JoinPath(cfg.Fetcher.Source, version)
	if err != nil {
		return nil, err
	}

	// #nosec G107 // specPath is a URL from config and does not contain user input
	resp, err := get(ctx, specPath)
	if err != nil {
		return nil, err
	}

	jsonSpec, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	formattedSpec := bytes.NewBufferString("")
	err = json.Indent(formattedSpec, jsonSpec, "", "  ")
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return formattedSpec, nil
}
