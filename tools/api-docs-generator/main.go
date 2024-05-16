package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: api-docs <config-file> <docs-dir>")
	}
	cfg, err := parseConfigFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	docsDirectory := os.Args[2]

	err = fetchSpec(cfg, docsDirectory)
	if err != nil {
		log.Fatal(err)
	}

	err = generateReferenceDocs(cfg, docsDirectory)
	if err != nil {
		log.Fatal(err)
	}

}
