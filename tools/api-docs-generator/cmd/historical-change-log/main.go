package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/snyk/user-docs/tools/api-docs-generator/changelog"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancelCtx()
	if len(os.Args) != 3 {
		log.Panicf("usage: api-docs <config-file> <docs-dir>")
	}

	fmt.Println(os.Args)

	cfg, err := config.Parse(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	docsDirectory := os.Args[2]

	if cfg.Changelog.HistoricalVersionCutoff == "" {
		log.Panic("Missing historical version cutoff")
	}

	err = changelog.GenerateHistorical(ctx, cfg, docsDirectory)
	if err != nil {
		log.Panic(err)
	}
}
