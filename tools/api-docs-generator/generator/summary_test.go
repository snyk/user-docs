package generator

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSummary = `# Table of contents

## Snyk API

* [Changelog](snyk-api/changelog.md)
* [Reference](snyk-api/reference/README.md)
  * [AccessRequests](snyk-api/reference/accessrequests.md)
  * [AIBOM](snyk-api/reference/aibom.md)
  * [Groups (v1)](snyk-api/reference/groups-v1.md)
  * [Groups](snyk-api/reference/groups.md)
  * [Monitor (v1)](snyk-api/reference/monitor-v1.md)
* [API endpoints index and tips](snyk-api/api-endpoints-index-and-tips/README.md)
`

func writeSummary(t *testing.T, contents string) string {
	t.Helper()
	summaryPath := path.Join(t.TempDir(), "SUMMARY.md")
	require.NoError(t, os.WriteFile(summaryPath, []byte(contents), 0o600))
	return summaryPath
}

func readSummary(t *testing.T, summaryPath string) string {
	t.Helper()
	contents, err := os.ReadFile(summaryPath)
	require.NoError(t, err)
	return string(contents)
}

func entriesFor(labels ...string) []summaryEntry {
	entries := make([]summaryEntry, len(labels))
	for i, label := range labels {
		entries[i] = summaryEntry{label: label, link: path.Join("snyk-api/reference", labelToFileName(label))}
	}
	return entries
}

func Test_syncSummary_noChanges(t *testing.T) {
	summaryPath := writeSummary(t, testSummary)

	changes, err := syncSummary(summaryPath, "snyk-api/reference",
		entriesFor("Monitor (v1)", "AccessRequests", "Groups", "AiBom", "Groups (v1)"))
	require.NoError(t, err)

	assert.True(t, changes.empty())
	assert.Equal(t, testSummary, readSummary(t, summaryPath))
}

func Test_syncSummary_addsNewPageInOrder(t *testing.T) {
	summaryPath := writeSummary(t, testSummary)

	changes, err := syncSummary(summaryPath, "snyk-api/reference",
		entriesFor("AccessRequests", "AiBom", "Groups (v1)", "Groups", "Model", "Monitor (v1)"))
	require.NoError(t, err)

	assert.Equal(t, []string{"Model"}, changes.added)
	assert.Empty(t, changes.removed)
	assert.Contains(t, readSummary(t, summaryPath),
		"  * [Groups](snyk-api/reference/groups.md)\n"+
			"  * [Model](snyk-api/reference/model.md)\n"+
			"  * [Monitor (v1)](snyk-api/reference/monitor-v1.md)\n")
}

func Test_syncSummary_removesPageThatWasNotGenerated(t *testing.T) {
	summaryPath := writeSummary(t, testSummary)

	changes, err := syncSummary(summaryPath, "snyk-api/reference",
		entriesFor("AccessRequests", "AiBom", "Groups", "Monitor (v1)"))
	require.NoError(t, err)

	assert.Empty(t, changes.added)
	assert.Equal(t, []string{"Groups (v1)"}, changes.removed)
	assert.NotContains(t, readSummary(t, summaryPath), "groups-v1.md")
}

func Test_syncSummary_keepsHandEditedLabels(t *testing.T) {
	summaryPath := writeSummary(t, testSummary)

	_, err := syncSummary(summaryPath, "snyk-api/reference",
		entriesFor("AccessRequests", "AiBom", "Groups (v1)", "Groups", "Model", "Monitor (v1)"))
	require.NoError(t, err)

	summary := readSummary(t, summaryPath)
	assert.Contains(t, summary, "[AIBOM](snyk-api/reference/aibom.md)")
	assert.NotContains(t, summary, "[AiBom]")
}

func Test_syncSummary_leavesTheRestOfTheFileAlone(t *testing.T) {
	summaryPath := writeSummary(t, testSummary)

	_, err := syncSummary(summaryPath, "snyk-api/reference", entriesFor("AccessRequests"))
	require.NoError(t, err)

	assert.Equal(t, `# Table of contents

## Snyk API

* [Changelog](snyk-api/changelog.md)
* [Reference](snyk-api/reference/README.md)
  * [AccessRequests](snyk-api/reference/accessrequests.md)
* [API endpoints index and tips](snyk-api/api-endpoints-index-and-tips/README.md)
`, readSummary(t, summaryPath))
}

func Test_syncSummary_errorsWithoutReferenceEntry(t *testing.T) {
	summaryPath := writeSummary(t, "# Table of contents\n\n* [Overview](README.md)\n")

	_, err := syncSummary(summaryPath, "snyk-api/reference", entriesFor("Apps"))
	assert.ErrorContains(t, err, "snyk-api/reference/README.md")
}

func Test_syncSummary_refusesToFlattenNestedEntries(t *testing.T) {
	nested := `* [Reference](snyk-api/reference/README.md)
  * [Apps](snyk-api/reference/apps.md)
    * [Hand-written child](snyk-api/reference/apps/child.md)
`
	summaryPath := writeSummary(t, nested)

	_, err := syncSummary(summaryPath, "snyk-api/reference", entriesFor("Apps", "Model"))
	assert.ErrorContains(t, err, "nested entry")
	assert.Equal(t, nested, readSummary(t, summaryPath), "file is not modified on error")
}

func Test_groupPagesByFileName_mergesLabelsThatShareAFile(t *testing.T) {
	pages := groupPagesByFileName(map[string][]operationPath{
		"OpensourceSettings": {{pathURL: "/orgs/{org_id}/settings/opensource/broker", method: "GET"}},
		"OpenSourceSettings": {{pathURL: "/orgs/{org_id}/settings/opensource", method: "GET"}},
		"Apps":               {{pathURL: "/apps", method: "GET"}},
	})

	require.Len(t, pages, 2)
	page := pages["opensourcesettings.md"]
	assert.Equal(t, "OpenSourceSettings", page.label, "label that sorts first wins")
	assert.Len(t, page.operations, 2, "operations from both labels are kept")
}
