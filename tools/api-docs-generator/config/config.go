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

type Changelog struct {
	HistoricalDate string `yaml:"historicalDate"`
	LastSyncDate   string `yaml:"lastSyncDate"`
}

type Output struct {
	APIReferencePath string `yaml:"apiReferencePath"`
}

type Config struct {
	Fetcher   Fetcher   `yaml:"fetcher"`
	Specs     []Spec    `yaml:"specs"`
	Output    Output    `yaml:"output"`
	Changelog Changelog `yaml:"changelog"`
}

func Parse(filename string) (*Config, error) {
	cfg := Config{}
	file, err := os.Open(filename)
	if err != nil {
		return &cfg, err
	}
	err = yaml.NewDecoder(file).Decode(&cfg)
	return &cfg, err
}

func Update(filename string, cfg *Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = yaml.NewEncoder(file).Encode(cfg)
	return err
}
