package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"path"
	"sort"
	"strings"
)

type operationPath struct {
	operation *openapi3.Operation
	pathItem  *openapi3.PathItem
	pathUrl   string
	specPath  string
	method    string
	docsHint  string
}

func generateReferenceDocs(config config, docsPath string) error {
	aggregatedDocs := map[string][]operationPath{}

	assetPathBase := path.Join(docsPath, "docs")

	for _, spec := range config.Specs {
		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromFile(path.Join(assetPathBase, spec.Path))
		if err != nil {
			return err
		}
		for pathUrl, pathItem := range doc.Paths.Map() {
			for method, operation := range pathItem.Operations() {
				for _, tag := range operation.Tags {
					tag = tag + spec.Suffix
					aggregatedDocs[tag] = append(aggregatedDocs[tag], operationPath{
						operation: operation,
						pathItem:  pathItem,
						pathUrl:   pathUrl,
						specPath:  spec.Path,
						method:    method,
						docsHint:  spec.DocsHint,
					})
				}
			}
		}
	}

	var summary []string
	for label, operation := range aggregatedDocs {
		if label == "OpenAPI" {
			continue
		}
		filePath := path.Join(docsPath, "docs/", config.Output.APIReferencePath, labelToFileName(label))
		docsFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		summary = append(summary, fmt.Sprintf("* [%s](%s)\n", label, path.Join(config.Output.APIReferencePath, labelToFileName(label))))

		fmt.Fprintf(docsFile, `# %s

{%% hint style="info" %%}
%s
{%% endhint %%}`, label, operation[0].docsHint)

		// sort for stability
		sort.Slice(operation, func(i, j int) bool {
			return operation[i].pathUrl+operation[i].method > operation[j].pathUrl+operation[j].method
		})

		for _, op := range operation {
			_, err = fmt.Fprintf(docsFile,
				`
{%% swagger src="../../%s" path="%s" method="%s" %%}
[spec.yaml](../../%s)
{%% endswagger %%}
`,
				op.specPath,
				op.pathUrl,
				strings.ToLower(op.method),
				op.specPath,
			)
			if err != nil {
				return err
			}
		}
	}
	sort.Strings(summary)
	fmt.Printf("generated menu for summary:\n")
	fmt.Printf("%s", strings.Join(summary, ""))

	return nil
}

func labelToFileName(label string) string {
	replacements := []string{"(", ")"}
	for _, replacement := range replacements {
		label = strings.ReplaceAll(label, replacement, "")
	}
	return strings.ToLower(strings.ReplaceAll(label, " ", "-")) + ".md"
}
