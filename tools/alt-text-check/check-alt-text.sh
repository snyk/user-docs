#!/usr/bin/env bash
# Report images that have no alt text.
#
# Flags three patterns in Markdown files:
#   ![](path)           Markdown image with empty alt text
#   <img ... alt="">    HTML image with empty alt text
#   <img ...>           HTML image with no alt attribute
#
# Usage:
#   check-alt-text.sh [file ...]   Check the named files
#   check-alt-text.sh              Check every tracked Markdown file
#
# Generated pages are skipped; see is_excluded below.
#
# Exits 0 when no violations are found, 1 otherwise. The GitHub workflow that
# calls this script does not fail the build on exit 1; it reports the findings
# as warnings. See .github/workflows/check-alt-text.yml.

set -uo pipefail

if [ "$#" -gt 0 ]; then
  files=("$@")
else
  IFS=$'\n' read -r -d '' -a files < <(git ls-files '*.md' && printf '\0')
fi

[ "${#files[@]}" -gt 0 ] || { printf 'No Markdown files to check.\n'; exit 0; }

# Paths this check skips. These pages are not hand-authored prose:
# the API reference is generated from the OpenAPI spec by
# tools/api-docs-generator (see its config.yml apiReferencePath), and the CLI
# help pages mirror the output of snyk --help. A writer cannot act on an alt
# text finding in either, so reporting one is noise.
is_excluded() {
  case "$1" in
    node_modules/*|*/node_modules/*) return 0 ;;
    developer-tools/snyk-api/reference/*) return 0 ;;
    developer-tools/snyk-cli/commands/*) return 0 ;;
  esac
  return 1
}

existing=()
skipped=0
for file in "${files[@]}"; do
  [ -f "$file" ] || continue
  if is_excluded "$file"; then
    skipped=$((skipped + 1))
    continue
  fi
  existing+=("$file")
done

if [ "$skipped" -gt 0 ]; then
  printf 'Skipped %s generated file(s): API reference and CLI help.\n' "$skipped"
fi

[ "${#existing[@]}" -gt 0 ] || { printf 'No Markdown files to check.\n'; exit 0; }

# awk finds each offending image and prints one file:line: message per hit.
# Using awk rather than grep -P keeps this portable to BSD grep on macOS.
output=$(awk '
  function report(line) {
    text = line
    sub(/^[ \t]+/, "", text)
    if (length(text) > 120) text = substr(text, 1, 120)
    printf "%s:%d: image is missing alt text: %s\n", FILENAME, FNR, text
    found++
  }
  {
    if ($0 ~ /!\[\]\(/) { report($0); next }

    rest = $0
    while (match(rest, /<img[^>]*>/)) {
      tag = substr(rest, RSTART, RLENGTH)
      rest = substr(rest, RSTART + RLENGTH)
      if (tag ~ /alt[ \t]*=[ \t]*""/ || tag !~ /[ \t]alt[ \t]*=/) {
        report($0)
        next
      }
    }
  }
  END { exit(found > 0 ? 1 : 0) }
' "${existing[@]}")
status=$?

if [ -n "$output" ]; then
  printf '%s\n' "$output"
fi

if [ "$status" -ne 0 ]; then
  count=$(printf '%s\n' "$output" | grep -c 'image is missing alt text')
  printf '\n%s image(s) missing alt text.\n' "$count"
  printf 'Add a short description of what each image shows.\n'
  exit 1
fi

printf 'All images have alt text.\n'
