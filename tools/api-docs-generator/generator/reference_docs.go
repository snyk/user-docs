package generator

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

func GenerateReferenceDocs(cfg *config.Config, docsBasePath string) error {
	aggregatedDocs, err := aggregateSpecs(cfg, docsBasePath)
	if err != nil {
		return err
	}

	summary := make([]string, len(aggregatedDocs))
	err = clearDir(path.Join(docsBasePath, cfg.Output.APIReferencePath))
	if err != nil {
		return err
	}

	for label, operations := range aggregatedDocs {
		destinationPath := path.Join(docsBasePath, cfg.Output.APIReferencePath, labelToFileName(label))
		summary = append(summary, fmt.Sprintf("* [%s](%s)\n", label, path.Join(cfg.Output.APIReferencePath, labelToFileName(label))))

		err = renderReferenceDocsPage(destinationPath, label, docsBasePath, operations, cfg.CategoryContext)
		if err != nil {
			return err
		}
	}
	sort.Strings(summary)

	matches, err := matchCurrentSummary(path.Join(docsBasePath, cfg.Output.SummaryPath), summary)
	if err != nil {
		return err
	}

	if !matches {
		fmt.Printf("generated menu for summary:\n")
		fmt.Printf("%s", strings.Join(summary, ""))
	}

	return nil
}

func matchCurrentSummary(summaryPath string, summary []string) (bool, error) {
	contents, err := os.ReadFile(summaryPath)
	if err != nil {
		return false, fmt.Errorf("failed to read summary file: %w", err)
	}
	currentSummary := string(contents)
	for _, menuItem := range summary {
		if !strings.Contains(currentSummary, menuItem) {
			return false, nil
		}
	}
	return true, nil
}

func clearDir(dirName string) error {
	dir, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	for _, child := range dir {
		if strings.HasPrefix(child.Name(), "README") {
			continue
		}
		err = os.RemoveAll(path.Join(dirName, child.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func aggregateSpecs(cfg *config.Config, docsBasePath string) (map[string][]operationPath, error) {
	aggregatedDocs := make(map[string][]operationPath)

	for _, spec := range cfg.Specs {
		specDocs, err := processSpec(spec, docsBasePath)
		if err != nil {
			return nil, err
		}

		for tag, ops := range specDocs {
			aggregatedDocs[tag] = append(aggregatedDocs[tag], ops...)
		}
	}

	return aggregatedDocs, nil
}

func processSpec(spec config.Spec, docsBasePath string) (map[string][]operationPath, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(path.Join(docsBasePath, spec.Path))
	if err != nil {
		return nil, err
	}

	specDocs := make(map[string][]operationPath)
	for pathURL, pathItem := range doc.Paths.Map() {
		err := processPathItem(pathURL, pathItem, spec, specDocs)
		if err != nil {
			return nil, err
		}
	}
	return specDocs, nil
}

func processPathItem(pathURL string, pathItem *openapi3.PathItem, spec config.Spec, specDocs map[string][]operationPath) error {
	for method, operation := range pathItem.Operations() {
		err := processOperation(pathURL, pathItem, method, operation, spec, specDocs)
		if err != nil {
			return err
		}
	}
	return nil
}

func processOperation(pathURL string,
	pathItem *openapi3.PathItem,
	method string,
	operation *openapi3.Operation,
	spec config.Spec,
	specDocs map[string][]operationPath) error {
	for _, tag := range operation.Tags {
		if tag == "OpenAPI" {
			continue
		}

		if snykDocsExtension, ok := operation.Extensions["x-snyk-documentation"]; ok && snykDocsExtension != nil {
			var err error
			tag, err = extractCategoryNameFromExtension(snykDocsExtension)
			if err != nil {
				return err
			}
		}

		tag += spec.Suffix
		specDocs[tag] = append(specDocs[tag], operationPath{
			operation: operation,
			pathItem:  pathItem,
			pathURL:   pathURL,
			specPath:  spec.Path,
			method:    method,
			docsHint:  spec.DocsHint,
		})
	}
	return nil
}

func extractCategoryNameFromExtension(extension interface{}) (string, error) {
	extensionMap, worked := extension.(map[string]interface{})
	if !worked {
		return "", fmt.Errorf("failed to parse docs extension as an object")
	}
	categoryValue, worked := extensionMap["category"].(string)
	if !worked {
		return "", fmt.Errorf("x-snyk-documentation extension category field not a string")
	}
	return categoryValue, nil
}

func renderReferenceDocsPage(filePath, label, docsPath string, operation []operationPath, categoryContext config.CategoryContexts) error {
	docsFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(docsFile, `# %s

{%% hint style="info" %%}
%s
{%% endhint %%}
`, label, operation[0].docsHint)
	if err != nil {
		return err
	}
	if categoryContextHint, found := categoryContext.ToMap()[labelToKey(label)]; found {
		_, err = fmt.Fprintln(docsFile)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(docsFile, categoryContextHint)
		if err != nil {
			return err
		}
	}

	// sort for stability
	sort.Slice(operation, func(i, j int) bool {
		return operation[i].pathURL+operation[i].method > operation[j].pathURL+operation[j].method
	})
	for _, op := range operation {
		relativePathToSpec, err := filepath.Rel(path.Dir(filePath), path.Join(docsPath, op.specPath))
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(docsFile,
			`
{%% openapi src="%s" path="%s" method="%s" %%}
[%s](%s)
{%% endopenapi %%}
`,
			relativePathToSpec,
			op.pathURL,
			strings.ToLower(op.method),
			path.Base(relativePathToSpec),
			relativePathToSpec,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func labelToFileName(label string) string {
	return labelToKey(label) + ".md"
}

func labelToKey(label string) string {
	replacements := []string{"(", ")"}
	for _, replacement := range replacements {
		label = strings.ReplaceAll(label, replacement, "")
	}

	return strings.ToLower(strings.ReplaceAll(label, " ", "-"))
}
