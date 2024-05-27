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

	if cfg.Changelog.HistoricalDate == "" {
		log.Panic("Missing historical changelog date")
	}

	//
	// err = changelog.GenerateHistorical(ctx, &cfg, "docs/snyk-api/changelog.md", cfg.Changelog.HistoricalDate)
	// if err != nil {
	//	log.Panic(err)
	// }

	updatedToVersion, err := changelog.UpdateChangelog(ctx, cfg, "docs/snyk-api/CHANGELOG.md")
	if err != nil {
		log.Panic(err)
	}

	if updatedToVersion != "" {
		cfg.Changelog.LastSyncDate = updatedToVersion
		err := config.Update(configFile, cfg)
		if err != nil {
			log.Panic(err)
		}
	}
}
