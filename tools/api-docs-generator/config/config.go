package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Fetcher struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

type Spec struct {
	Path     string `yaml:"path"`
	Suffix   string `yaml:"suffix,omitempty"`
	DocsHint string `yaml:"docsHint,omitempty"`
}

type Output struct {
	SummaryPath      string `yaml:"summaryPath"`
	APIReferencePath string `yaml:"apiReferencePath"`
}

type Config struct {
	Fetcher         Fetcher          `yaml:"fetcher"`
	Specs           []Spec           `yaml:"specs"`
	Output          Output           `yaml:"output"`
	CategoryContext CategoryContexts `yaml:"categoryContext"`
}

type CategoryContexts []CategoryContext

func (contexts CategoryContexts) ToMap() map[string]string {
	m := make(map[string]string)
	for i := range contexts {
		m[contexts[i].Name] = contexts[i].Hint
	}
	return m
}

type CategoryContext struct {
	Name string `yaml:"name"`
	Hint string `yaml:"hint"`
}

func Parse(filename string) (*Config, error) {
	cfg := Config{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
