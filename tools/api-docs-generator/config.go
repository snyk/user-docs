package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Fetcher struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	} `yaml:"fetcher"`
	Specs []struct {
		Path     string `yaml:"path"`
		Suffix   string `yaml:"suffix,omitempty"`
		DocsHint string `yaml:"docsHint,omitempty"`
	} `yaml:"specs"`
	Output struct {
		APIReferencePath string `yaml:"apiReferencePath"`
	} `yaml:"output"`
}

func parseConfigFile(filename string) (config, error) {
	cfg := config{}
	file, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	return cfg, yaml.NewDecoder(file).Decode(&cfg)
}
