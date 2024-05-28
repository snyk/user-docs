package changelog

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	md "github.com/nao1215/markdown"
)

type testConfig struct {
	name            string
	baseURL         string // URL of currently stored file
	nextURL         string // URL of last GA version as returned by /openapi
	latestGAVersion string // latest ga version as returned by /openapi
	lastSyncVersion string // last version that was processed by the changelog tool
	want            string
}

var tests = []*testConfig{
	{
		name:            "01_new_version_added",
		baseURL:         "../testdata/changelog/01_new_version_added/2024-04-25.yaml",
		nextURL:         "../testdata/changelog/01_new_version_added/2024-05-23.yaml",
		latestGAVersion: "2024-05-23",
		lastSyncVersion: "2024-04-25",
		want:            "## 2024-05-23",
	},
	{
		name:            "02_no_changes_detected",
		baseURL:         "../testdata/changelog/02_no_changes_detected/2024-04-25.yaml",
		nextURL:         "../testdata/changelog/02_no_changes_detected/2024-05-23.yaml",
		latestGAVersion: "2024-04-25",
		lastSyncVersion: "2024-04-25",
		want:            "",
	},
	{
		name:            "03_current_version_updated_in_place",
		baseURL:         "../testdata/changelog/03_current_version_updated_in_place/2024-05-23.yaml",
		nextURL:         "../testdata/changelog/03_current_version_updated_in_place/2024-05-23_update.yaml",
		latestGAVersion: "2024-05-23",
		lastSyncVersion: "2024-05-23",
		want:            fmt.Sprintf("## %s - Updated %s", "2024-05-23", time.Now().Format("2006-01-02")),
	},
}

func testFn(config *testConfig) (string, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	groupedChanges, err := getChangeLog(config.nextURL, config.baseURL, loader)
	if err != nil {
		return "", err
	}

	groupedChanges = filterChanges(groupedChanges)

	if len(groupedChanges) == 0 {
		return "", err
	}
	builder := new(strings.Builder)
	markdown := md.NewMarkdown(builder)

	err = WriteToChangeLog(markdown, groupedChanges, config.latestGAVersion, config.nextURL, config.lastSyncVersion)
	if err != nil {
		return "", err
	}

	err = markdown.Build()
	if err != nil {
		return "", err
	}

	return builder.String(), err
}

func Test_delta(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotErr := testFn(tt)
			if gotErr != nil {
				t.Errorf("Expected not to fail %v", gotErr)
			}
			if !strings.Contains(gotRes, tt.want) {
				t.Errorf("Expected markdown to contain %v", tt.want)
			}
		})
	}
}
