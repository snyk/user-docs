package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type SyncStateConfig struct {
	LastSyncedVersion string `yaml:"lastSyncedVersion"`
}

func LoadSyncState(filename string) (SyncStateConfig, error) {
	cfg := SyncStateConfig{}
	file, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	err = yaml.NewDecoder(file).Decode(&cfg)
	return cfg, err
}

func UpdateSyncState(filename string, syncState SyncStateConfig) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = yaml.NewEncoder(file).Encode(syncState)
	return err
}
