package main

import (
	"context"
	"log"
	"time"

	"github.com/snyk/user-docs/tools/api-docs-generator/changelog"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancelCtx()

	cfg, err := config.Parse("tools/api-docs-generator/config.yml")
	if err != nil {
		log.Panic(err)
	}

	err = changelog.GenerateHistorical(ctx, cfg, "docs/snyk-api/changelog.md", "2024-05-08")
	if err != nil {
		log.Panic(err)
	}

	//err = changelog.UpdateChangelog(ctx, cfg, "docs/snyk-api/changelog.md", "2024-04-29")
	//if err != nil {
	//	log.Panic(err)
	//}
}
