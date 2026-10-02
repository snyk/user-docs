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

	err = clearDir(path.Join(docsBasePath, cfg.Output.APIReferencePath))
	if err != nil {
		return err
	}

	summaryPath := path.Join(docsBasePath, cfg.Output.SummaryPath)
	referenceDir, err := filepath.Rel(path.Dir(summaryPath), path.Join(docsBasePath, cfg.Output.APIReferencePath))
	if err != nil {
		return err
	}

	pages := groupPagesByFileName(aggregatedDocs)
	entries := make([]summaryEntry, 0, len(pages))
	for fileName, page := range pages {
		destinationPath := path.Join(docsBasePath, cfg.Output.APIReferencePath, fileName)
		entries = append(entries, summaryEntry{label: page.label, link: path.Join(referenceDir, fileName)})

		err = renderReferenceDocsPage(destinationPath, page.label, docsBasePath, page.operations, cfg.CategoryContext)
		if err != nil {
			return err
		}
	}

	changes, err := syncSummary(summaryPath, referenceDir, entries)
	if err != nil {
		return err
	}
	if !changes.empty() {
		fmt.Printf("updated %s:\n", cfg.Output.SummaryPath)
		for _, label := range changes.added {
			fmt.Printf("+ %s\n", label)
		}
		for _, label := range changes.removed {
			fmt.Printf("- %s\n", label)
		}
	}

	return nil
}

type referencePage struct {
	label      string
	operations []operationPath
}

// groupPagesByFileName merges labels that render to the same file, such as the
// "OpenSourceSettings" and "OpensourceSettings" tags. Rendering them
// separately would make the second overwrite the first, dropping its endpoints
// from the docs depending on map iteration order.
//
// The page takes the label that sorts first, so the choice is stable across runs.
func groupPagesByFileName(aggregatedDocs map[string][]operationPath) map[string]referencePage {
	labels := make([]string, 0, len(aggregatedDocs))
	for label := range aggregatedDocs {
		labels = append(labels, label)
	}
	sort.Strings(labels)

	pages := make(map[string]referencePage)
	for _, label := range labels {
		fileName := labelToFileName(label)
		page, found := pages[fileName]
		if !found {
			page.label = label
		}
		page.operations = append(page.operations, aggregatedDocs[label]...)
		pages[fileName] = page
	}
	return pages
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

	_, err = fmt.Fprintf(docsFile, `---
description: %s
---

# %s

{%% hint style="info" %%}
%s
{%% endhint %%}
`, referencePageDescription(label), label, operation[0].docsHint)
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

// referencePageDescription builds the frontmatter description for a reference
// page. GitBook uses this as the page's meta description, and it is what search
// results and AI assistants quote when they surface the page.
//
// The text is derived from the label so that every generated page carries one
// without a maintainer having to write 60 of them by hand.
func referencePageDescription(label string) string {
	return fmt.Sprintf("Snyk API reference for the %s endpoints, including request parameters and response schemas", label)
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
