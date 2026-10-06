package generator

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// summaryEntry is one generated reference page as it should appear in SUMMARY.md.
type summaryEntry struct {
	label string
	// link is relative to the directory that holds SUMMARY.md, which is how
	// GitBook resolves the links in a table of contents.
	link string
}

// summaryChanges reports what syncSummary changed, for the PR description.
type summaryChanges struct {
	added   []string
	removed []string
}

func (c summaryChanges) empty() bool {
	return len(c.added) == 0 && len(c.removed) == 0
}

var summaryItemPattern = regexp.MustCompile(`^( *)\* \[(.*)\]\((.*)\)\s*$`)

// syncSummary rewrites the children of the reference README entry in
// SUMMARY.md so they list exactly the generated reference pages.
//
// GitBook only shows a page in the navigation if SUMMARY.md lists it, so a
// generated page that is missing from SUMMARY.md is unreachable on the docs
// site. Children for pages that were not generated are removed, because
// GenerateReferenceDocs has already deleted those files and the links would
// be broken.
//
// Labels already in SUMMARY.md are kept, so hand-edited capitalization such as
// "AIBOM" survives. New pages use the generated label.
func syncSummary(summaryPath, referenceDir string, entries []summaryEntry) (summaryChanges, error) {
	contents, err := os.ReadFile(summaryPath)
	if err != nil {
		return summaryChanges{}, fmt.Errorf("failed to read summary file: %w", err)
	}
	lines := strings.Split(string(contents), "\n")

	parentLink := referenceDir + "/README.md"
	parent := -1
	parentIndent := ""
	for i, line := range lines {
		if m := summaryItemPattern.FindStringSubmatch(line); m != nil && m[3] == parentLink {
			parent, parentIndent = i, m[1]
			break
		}
	}
	if parent == -1 {
		return summaryChanges{}, fmt.Errorf("no entry linking to %s in %s", parentLink, summaryPath)
	}

	childIndent := parentIndent + "  "
	var existing []summaryEntry
	end := parent + 1
	for ; end < len(lines); end++ {
		m := summaryItemPattern.FindStringSubmatch(lines[end])
		if m == nil || len(m[1]) <= len(parentIndent) {
			break
		}
		if m[1] != childIndent {
			return summaryChanges{}, fmt.Errorf(
				"%s:%d: nested entry under %s; generated reference pages must be direct children",
				summaryPath, end+1, parentLink)
		}
		existing = append(existing, summaryEntry{label: m[2], link: m[3]})
	}

	var changes summaryChanges
	wanted := make(map[string]bool, len(entries))
	for _, entry := range entries {
		wanted[entry.link] = true
	}

	// Keep existing entries in their current order, so a run that adds one page
	// changes one line instead of re-sorting the whole block.
	children := make([]summaryEntry, 0, len(entries))
	listed := make(map[string]bool, len(existing))
	for _, entry := range existing {
		if !wanted[entry.link] {
			changes.removed = append(changes.removed, entry.label)
			continue
		}
		listed[entry.link] = true
		children = append(children, entry)
	}

	var added []summaryEntry
	for _, entry := range entries {
		if !listed[entry.link] {
			added = append(added, entry)
		}
	}
	sort.Slice(added, func(i, j int) bool {
		return summarySortKey(added[i]) < summarySortKey(added[j])
	})
	for _, entry := range added {
		changes.added = append(changes.added, entry.label)
		children = insertSorted(children, entry)
	}
	if changes.empty() {
		return changes, nil
	}

	rendered := make([]string, len(children))
	for i, child := range children {
		rendered[i] = fmt.Sprintf("%s* [%s](%s)", childIndent, child.label, child.link)
	}

	updated := make([]string, 0, len(lines)-(end-parent-1)+len(rendered))
	updated = append(updated, lines[:parent+1]...)
	updated = append(updated, rendered...)
	updated = append(updated, lines[end:]...)

	err = os.WriteFile(summaryPath, []byte(strings.Join(updated, "\n")), 0o600)
	if err != nil {
		return summaryChanges{}, fmt.Errorf("failed to write summary file: %w", err)
	}
	return changes, nil
}

// insertSorted inserts entry before the first child that sorts after it. The
// existing children are not re-sorted.
func insertSorted(children []summaryEntry, entry summaryEntry) []summaryEntry {
	key := summarySortKey(entry)
	i := 0
	for i < len(children) && summarySortKey(children[i]) < key {
		i++
	}
	children = append(children, summaryEntry{})
	copy(children[i+1:], children[i:])
	children[i] = entry
	return children
}

// summarySortKey orders entries case-insensitively by label, with "Groups (v1)"
// before "Groups", which is the order SUMMARY.md already uses.
func summarySortKey(entry summaryEntry) string {
	return strings.ToLower("[" + entry.label + "](" + entry.link)
}
