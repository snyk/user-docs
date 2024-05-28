package main

import (
	"context"
	"log"
	"time"

	"github.com/snyk/user-docs/tools/api-docs-generator/changelog"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

const configFile = "tools/api-docs-generator/config.yml"

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancelCtx()

	cfg, err := config.Parse(configFile)
	if err != nil {
		log.Panic(err)
	}

	if cfg.Changelog.HistoricalVersionCutoff == "" {
		log.Panic("Missing historical version cutoff")
	}

	err = changelog.GenerateHistorical(ctx, cfg, "docs/snyk-api/changelog.md", cfg.Changelog.HistoricalVersionCutoff)
	if err != nil {
		log.Panic(err)
	}
}
