package generator

import (
	"os"
	"path"
	"testing"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"

	"github.com/stretchr/testify/assert"
)

func Test_labelToFileName(t *testing.T) {
	tests := []struct {
		name  string
		label string
		want  string
	}{
		{
			label: "Apps",
			want:  "apps.md",
		},
		{
			label: "spaces to dashes",
			want:  "spaces-to-dashes.md",
		},
		{
			name:  "removes parentheses",
			label: "Apps (v1)",
			want:  "apps-v1.md",
		},
		{
			label: "Apps (test)",
			want:  "apps-test.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := labelToFileName(tt.label); got != tt.want {
				t.Errorf("labelToFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renderReferenceDocsPage(t *testing.T) {
	type args struct {
		filePath        string
		label           string
		docsPath        string
		operation       []operationPath
		categoryContext config.CategoryContexts
	}
	testDir := t.TempDir()
	tests := []struct {
		name    string
		args    args
		wantErr bool
		checker func(t *testing.T, arg args)
	}{
		{
			name: "renders reference docs page",
			args: args{
				filePath: createTempFile(t, testDir, "existing content"),
				label:    "Apps",
				docsPath: testDir,
				operation: []operationPath{
					{
						specPath: "foo/test/apps-spec.yaml",
						pathURL:  "/apps",
						method:   "GET",
						docsHint: "This is a hint",
					},
				},
				categoryContext: config.CategoryContexts{},
			},
			checker: func(t *testing.T, args args) {
				t.Helper()
				content, err := os.ReadFile(args.filePath)
				if err != nil {
					t.Fatal(err)
				}
				renderedFileContents := string(content)
				assert.Contains(t, renderedFileContents, "# Apps", "rendered file contains header")
				assert.Contains(t, renderedFileContents, "This is a hint", "rendered file does not contain hint")
				assert.Contains(t, renderedFileContents, `path="/apps"`, "rendered file does not contain path")
				assert.Contains(t, renderedFileContents, `method="get"`, "rendered file does not contain method")
				assert.Contains(t, renderedFileContents, `src="../foo/test/apps-spec.yaml"`, "renders relative path to spec file")
			},
		},
		{
			name: "renders reference docs page, with category context hint",

			args: args{
				filePath: createTempFile(t, testDir, "existing content"),
				label:    "Apps",
				docsPath: testDir,
				operation: []operationPath{
					{
						specPath: "foo/test/apps-spec.yaml",
						pathURL:  "/apps",
						method:   "GET",
						docsHint: "This is a hint",
					},
				},
				categoryContext: config.CategoryContexts{
					{
						Name: "apps",
						Hint: "This is a hint from category context",
					},
					{
						Name: "not-apps",
						Hint: "This is a hint from another context",
					},
				},
			},
			checker: func(t *testing.T, args args) {
				t.Helper()
				content, err := os.ReadFile(args.filePath)
				if err != nil {
					t.Fatal(err)
				}
				renderedFileContents := string(content)
				assert.Contains(t, renderedFileContents, "# Apps", "rendered file contains header")
				assert.Contains(t, renderedFileContents, "This is a hint", "rendered file does not contain hint")
				assert.Contains(t, renderedFileContents, `path="/apps"`, "rendered file does not contain path")
				assert.Contains(t, renderedFileContents, `method="get"`, "rendered file does not contain method")
				assert.Contains(t, renderedFileContents, `src="../foo/test/apps-spec.yaml"`, "renders relative path to spec file")
				assert.Contains(t, renderedFileContents, `This is a hint from category context`, "renders category context hint")
			},
		},
		{
			name: "renders reference docs page, without category context hint if no matches",

			args: args{
				filePath: createTempFile(t, testDir, "existing content"),
				label:    "Apps",
				docsPath: testDir,
				operation: []operationPath{
					{
						specPath: "foo/test/apps-spec.yaml",
						pathURL:  "/apps",
						method:   "GET",
						docsHint: "This is a hint",
					},
				},
				categoryContext: config.CategoryContexts{
					{
						Name: "not-apps",
						Hint: "This is a hint from category context",
					},
				},
			},
			checker: func(t *testing.T, args args) {
				t.Helper()
				content, err := os.ReadFile(args.filePath)
				if err != nil {
					t.Fatal(err)
				}
				renderedFileContents := string(content)
				assert.Contains(t, renderedFileContents, "# Apps", "rendered file contains header")
				assert.Contains(t, renderedFileContents, "This is a hint", "rendered file does not contain hint")
				assert.Contains(t, renderedFileContents, `path="/apps"`, "rendered file does not contain path")
				assert.Contains(t, renderedFileContents, `method="get"`, "rendered file does not contain method")
				assert.Contains(t, renderedFileContents, `src="../foo/test/apps-spec.yaml"`, "renders relative path to spec file")
				assert.NotContains(t, renderedFileContents, `This is a hint from category context`, "does not render category context hint")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := renderReferenceDocsPage(tt.args.filePath,
				tt.args.label,
				tt.args.docsPath,
				tt.args.operation,
				tt.args.categoryContext); (err != nil) != tt.wantErr {
				t.Errorf("renderReferenceDocsPage() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.checker(t, tt.args)
		})
	}
}

func createTempFile(t *testing.T, baseDir, content string) string {
	t.Helper()
	fileBaseDir := path.Join(baseDir, "somepath")
	err := os.MkdirAll(fileBaseDir, 0o755)
	assert.NoError(t, err)
	file, err := os.CreateTemp(fileBaseDir, "output")
	assert.NoError(t, err)
	_, err = file.WriteString(content)
	assert.NoError(t, err)
	return file.Name()
}
