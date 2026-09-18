#!/usr/bin/env bash
#
# Approve and merge the newest qualifying IDE compatibility matrix PR.
#
# Every decision here is a deterministic check with exactly one right answer:
# string equality on the author, string equality on the changed path, a line
# prefix test on the diff, and check-run states. No judgement, no AI.
#
# Exit codes:
#   0  merged one PR, or found nothing to do (both are successful outcomes)
#   1  a PR from the automation exists but failed the envelope; needs a human

set -euo pipefail

REPO='snyk/user-docs'
AUTHOR='team-ide-user'
MATRIX_PATH='developer-tools/snyk-ide-plugins-and-extensions/compatibility-matrix.md'

# The matrix holds a rolling 12 months of plugin releases. A regeneration that
# drops most of the table is a bug in the upstream job, not a legitimate
# update, so require the result to stay near the established size.
MIN_TABLE_ROWS=20

DRY_RUN="${DRY_RUN:-false}"

log() { printf '%s\n' "$*"; }

# Emit a line to the workflow run summary so a skip or failure is legible in
# the Actions UI without opening the raw log.
summary() {
  if [[ -n "${GITHUB_STEP_SUMMARY:-}" ]]; then
    printf '%s\n' "$*" >>"$GITHUB_STEP_SUMMARY"
  fi
}

# ---------------------------------------------------------------------------
# 1. Find open PRs from the automation, newest last.
# ---------------------------------------------------------------------------

# `gh pr list --author` matches the PR author. Restricting to this login is
# what makes the rest of the envelope safe to trust.
#
# Kept as a newline-delimited string rather than a bash array: array append
# inside a `while read` loop fed by process substitution misbehaves on
# bash 3.2, and this script should not depend on the runner's bash version.
PR_NUMBERS="$(
  gh pr list --repo "$REPO" --author "$AUTHOR" --state open \
    --json number,createdAt \
    --jq 'sort_by(.createdAt) | .[].number'
)"

if [[ -z "$PR_NUMBERS" ]]; then
  log "No open PRs from $AUTHOR. Nothing to do."
  summary "No open PRs from \`$AUTHOR\`. Nothing to do."
  exit 0
fi

log "Open PRs from $AUTHOR (oldest first):"
printf '%s\n' "$PR_NUMBERS" | sed 's/^/  #/'

# Newest is the merge candidate. Everything older is a stale snapshot.
CANDIDATE="$(printf '%s\n' "$PR_NUMBERS" | tail -n 1)"
log "Merge candidate (newest): #$CANDIDATE"

# ---------------------------------------------------------------------------
# 2. Only the matrix file may be touched.
# ---------------------------------------------------------------------------

# Done entirely in jq so the "exactly one file, and it is this path" test is a
# single unambiguous expression rather than shell array handling.
# Piped to jq rather than using `gh --jq`, which does not support --arg.
PATHS_OK="$(
  gh pr view "$CANDIDATE" --repo "$REPO" --json files \
    | jq -r --arg want "$MATRIX_PATH" '
        if ([.files[].path] == [$want]) then "yes" else "no" end
      '
)"

if [[ "$PATHS_OK" != 'yes' ]]; then
  CHANGED_LIST="$(gh pr view "$CANDIDATE" --repo "$REPO" --json files --jq '[.files[].path] | join(", ")')"
  log "FAIL: #$CANDIDATE does not touch exactly the matrix file."
  log "Changed files: $CHANGED_LIST"
  summary "PR #$CANDIDATE **failed the envelope**: expected only \`$MATRIX_PATH\`, got: $CHANGED_LIST"
  exit 1
fi

log "OK: only $MATRIX_PATH changed."

# ---------------------------------------------------------------------------
# 3. The diff must be table rows only.
# ---------------------------------------------------------------------------

DIFF="$(gh pr diff "$CANDIDATE" --repo "$REPO")"

# Added and removed content lines, excluding the ---/+++ file headers. Every
# survivor must be a markdown table row, i.e. start with '|' after the diff
# marker. A changed heading, link, list item, blank line or frontmatter key
# all land here and fail the check.
OFFENDING="$(
  printf '%s\n' "$DIFF" \
    | grep -E '^[+-]' \
    | grep -vE '^(\+\+\+|---)' \
    | grep -vE '^[+-]\|' \
    || true
)"

if [[ -n "$OFFENDING" ]]; then
  log 'FAIL: diff contains non-table-row changes:'
  printf '%s\n' "$OFFENDING"
  summary "PR #$CANDIDATE **failed the envelope**: diff contains non-table-row changes."
  summary '```'
  summary "$OFFENDING"
  summary '```'
  exit 1
fi

log 'OK: diff is table rows only.'

# Guard against a regeneration that empties or truncates the table. Count the
# rows the PR would leave in place: existing rows, minus removed, plus added.
BASE_ROWS="$(
  gh api "repos/$REPO/contents/$MATRIX_PATH" --jq '.content' \
    | base64 --decode \
    | grep -cE '^\|' \
    || true
)"
REMOVED_ROWS="$(printf '%s\n' "$DIFF" | grep -cE '^-\|' || true)"
ADDED_ROWS="$(printf '%s\n' "$DIFF" | grep -cE '^\+\|' || true)"
RESULT_ROWS=$(( BASE_ROWS - REMOVED_ROWS + ADDED_ROWS ))

log "Table rows: base=$BASE_ROWS removed=$REMOVED_ROWS added=$ADDED_ROWS result=$RESULT_ROWS"

if [[ "$RESULT_ROWS" -lt "$MIN_TABLE_ROWS" ]]; then
  log "FAIL: resulting table would have $RESULT_ROWS rows, below the floor of $MIN_TABLE_ROWS."
  summary "PR #$CANDIDATE **failed the envelope**: resulting table would have $RESULT_ROWS rows (floor $MIN_TABLE_ROWS)."
  exit 1
fi

log "OK: resulting table size $RESULT_ROWS is at or above the floor."

# ---------------------------------------------------------------------------
# 4. Every check must be green.
# ---------------------------------------------------------------------------

# Deliberately not passing --required: only the CircleCI secrets scan is
# required by the ruleset, but the GitBook preview checks are what catch a
# broken table before it reaches docs.snyk.io. All checks must pass.
CHECK_BUCKETS="$(
  gh pr checks "$CANDIDATE" --repo "$REPO" --json bucket --jq '[.[].bucket] | unique | join(",")' \
    || true
)"

log "Check buckets: ${CHECK_BUCKETS:-<none>}"

if [[ -z "$CHECK_BUCKETS" ]]; then
  log "SKIP: no checks reported on #$CANDIDATE yet. Leaving it for the next run."
  summary "PR #$CANDIDATE has no checks reported yet. Skipped; will retry on the next run."
  exit 0
fi

# Pending is a skip, not a failure: the run fired early and the PR is still
# perfectly mergeable tomorrow.
if [[ ",$CHECK_BUCKETS," == *",pending,"* ]]; then
  log "SKIP: #$CANDIDATE still has pending checks. Leaving it for the next run."
  summary "PR #$CANDIDATE still has pending checks. Skipped; will retry on the next run."
  exit 0
fi

# Anything other than pass or skipping means a check is unhappy.
if [[ "$CHECK_BUCKETS" != 'pass' && "$CHECK_BUCKETS" != 'skipping' \
   && "$CHECK_BUCKETS" != 'pass,skipping' && "$CHECK_BUCKETS" != 'skipping,pass' ]]; then
  log "FAIL: #$CANDIDATE has checks that are not green (buckets: $CHECK_BUCKETS)."
  gh pr checks "$CANDIDATE" --repo "$REPO" || true
  summary "PR #$CANDIDATE **failed the envelope**: checks not green (buckets: $CHECK_BUCKETS)."
  exit 1
fi

log 'OK: all checks green.'

# Refuse to merge a PR that is not cleanly mergeable, e.g. a conflict against
# main. MERGEABLE is the only state safe to act on; UNKNOWN means GitHub has
# not finished computing it.
MERGE_STATE="$(gh pr view "$CANDIDATE" --repo "$REPO" --json mergeable --jq '.mergeable')"
log "Mergeable: $MERGE_STATE"

if [[ "$MERGE_STATE" == 'UNKNOWN' ]]; then
  log "SKIP: mergeability not yet computed for #$CANDIDATE. Leaving it for the next run."
  summary "PR #$CANDIDATE mergeability not yet computed. Skipped; will retry on the next run."
  exit 0
fi

if [[ "$MERGE_STATE" != 'MERGEABLE' ]]; then
  log "FAIL: #$CANDIDATE is not mergeable (state: $MERGE_STATE)."
  summary "PR #$CANDIDATE **failed the envelope**: not mergeable (state: $MERGE_STATE)."
  exit 1
fi

# ---------------------------------------------------------------------------
# 5. Approve, merge the newest, close the rest as superseded.
# ---------------------------------------------------------------------------

if [[ "$DRY_RUN" == 'true' ]]; then
  log "DRY RUN: would approve and merge #$CANDIDATE."
  for pr in $PR_NUMBERS; do
    [[ "$pr" == "$CANDIDATE" ]] && continue
    log "DRY RUN: would close #$pr as superseded."
  done
  summary "**Dry run.** PR #$CANDIDATE passed the full envelope and would have been merged."
  exit 0
fi

# The approval satisfies the CODEOWNERS review requirement on both the
# repo-level and org-level rulesets. It is not self-approval: the PR author is
# the automation machine user, not the token owner.
gh pr review "$CANDIDATE" --repo "$REPO" --approve \
  --body 'Automated approval: IDE compatibility matrix regeneration. Verified only `compatibility-matrix.md` changed, diff is table rows only, and all checks are green.'

log "Approved #$CANDIDATE."

# --merge, not --squash: squash and rebase are both disabled on this repo, so
# a merge commit is the only permitted method.
# --delete-branch is explicit because delete_branch_on_merge is off.
gh pr merge "$CANDIDATE" --repo "$REPO" --merge --delete-branch

log "Merged #$CANDIDATE."
summary "Merged PR #$CANDIDATE (IDE compatibility matrix)."

# Older PRs are superseded snapshots. Closing them keeps the queue from
# growing and stops a stale one from ever being merged over newer data.
for pr in $PR_NUMBERS; do
  [[ "$pr" == "$CANDIDATE" ]] && continue
  gh pr close "$pr" --repo "$REPO" --delete-branch \
    --comment "Superseded by #$CANDIDATE. Each run regenerates the whole matrix, so only the newest PR is merged."
  log "Closed #$pr as superseded."
  summary "Closed PR #$pr as superseded by #$CANDIDATE."
done
