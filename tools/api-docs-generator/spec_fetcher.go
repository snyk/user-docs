package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
)

func fetchSpec(cfg config, directory string) error {
	resp, err := http.Get(cfg.Fetcher.Source)
	if err != nil {
		return err
	}

	var versions []string

	if err = json.NewDecoder(resp.Body).Decode(&versions); err != nil {
		return err
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}

	gaVersion := getLatestGAVersion(versions)

	specPath, err := url.JoinPath(cfg.Fetcher.Source, gaVersion)
	if err != nil {
		return err
	}

	resp, err = http.Get(specPath)
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

	return os.WriteFile(path.Join(directory, "docs", cfg.Fetcher.Destination), formattedSpec.Bytes(), 0644)
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
