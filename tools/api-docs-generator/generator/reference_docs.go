package generator

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
)

type operationPath struct {
	operation *openapi3.Operation
	pathItem  *openapi3.PathItem
	pathURL   string
	specPath  string
	method    string
	docsHint  string
}

func GenerateReferenceDocs(cfg config.Config, docsPath string) error {
	aggregatedDocs := map[string][]operationPath{}

	assetPathBase := path.Join(docsPath, "docs")

	for _, spec := range cfg.Specs {
		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromFile(path.Join(assetPathBase, spec.Path))
		if err != nil {
			return err
		}
		for pathURL, pathItem := range doc.Paths.Map() {
			for method, operation := range pathItem.Operations() {
				for _, tag := range operation.Tags {
					tag += spec.Suffix
					aggregatedDocs[tag] = append(aggregatedDocs[tag], operationPath{
						operation: operation,
						pathItem:  pathItem,
						pathURL:   pathURL,
						specPath:  spec.Path,
						method:    method,
						docsHint:  spec.DocsHint,
					})
				}
			}
		}
	}

	summary := make([]string, len(aggregatedDocs))
	for label, operation := range aggregatedDocs {
		filePath := path.Join(docsPath, "docs/", cfg.Output.APIReferencePath, labelToFileName(label))
		docsFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		summary = append(summary, fmt.Sprintf("* [%s](%s)\n", label, path.Join(cfg.Output.APIReferencePath, labelToFileName(label))))

		fmt.Fprintf(docsFile, `# %s

{%% hint style="info" %%}
%s
{%% endhint %%}`, label, operation[0].docsHint)

		// sort for stability
		sort.Slice(operation, func(i, j int) bool {
			return operation[i].pathURL+operation[i].method > operation[j].pathURL+operation[j].method
		})

		for _, op := range operation {
			_, err = fmt.Fprintf(docsFile,
				`
{%% swagger src="../../%s" path="%s" method="%s" %%}
[spec.yaml](../../%s)
{%% endswagger %%}
`,
				op.specPath,
				op.pathURL,
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
