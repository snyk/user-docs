package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
	"github.com/snyk/user-docs/tools/api-docs-generator/fetcher"
	"github.com/snyk/user-docs/tools/api-docs-generator/generator"
)

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()
	if len(os.Args) != 3 {
		log.Panicf("usage: api-docs <config-file> <docs-dir>")
	}
	cfg, err := config.Parse(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	docsDirectory := os.Args[2]

	err = fetcher.FetchSpec(ctx, cfg, docsDirectory)
	if err != nil {
		log.Panic(err)
	}

	err = generator.GenerateReferenceDocs(cfg, docsDirectory)
	if err != nil {
		log.Panic(err)
	}
}
