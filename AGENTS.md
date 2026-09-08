# AGENTS.md

Instructions for AI coding agents working in `snyk/user-docs`.

This file covers what you cannot infer from the file tree. For writing style, use the Snyk documentation writing rules — do not infer style from surrounding pages, and do not treat the conventions in this file as style guidance.

## This repository is one half of a two-way sync

GitBook renders [docs.snyk.io](https://docs.snyk.io) from `main`, and Git Sync runs **bidirectionally**. Writers edit in GitBook, and GitBook commits straight to `main` as `gitbook-bot` with messages like `GITBOOK-64: vc-update title`. Those commits do not go through a pull request.

Consequences you must respect:

- **`main` moves without warning.** Rebase or re-pull before you push, and expect conflicts on content files.
- **Never rewrite published history** (`push --force`, amend on `main`, squash-merge older commits). It desynchronizes GitBook from Git, and recovery is manual.
- **There is no build and no local preview.** You cannot compile the site, run a link checker against it, or verify rendering locally. Verification happens in the GitBook preview on the pull request.
- If a file you edited changes under you with a `GITBOOK-` commit, a human is editing that page right now. Stop and leave it alone.

## Structure: SUMMARY.md is the only source of truth

Each top-level directory containing a `SUMMARY.md` is one GitBook space, and the set of spaces matches the site structure of docs.snyk.io 1:1. There is no root `SUMMARY.md` and no `.gitbook.yaml`.

- **A page that is not listed in its space's `SUMMARY.md` does not exist on the site.** Adding, renaming, or removing a page means editing `SUMMARY.md` in the same commit.
- **A file's location does not determine where it appears in the navigation.** `SUMMARY.md` freely points across subdirectories within its space, so the on-disk tree and the rendered tree legitimately disagree. Do not move, rename, or reorganize files to make them match. Structural changes are a Docs team decision, not cleanup.
- Keep changes inside one space. Do not touch another space's `SUMMARY.md`.
- `.gitbook/includes/` holds reusable blocks (for example `new-navigation-banner.md`) that are transcluded into many pages, sometimes across spaces. Editing one changes every page that includes it.
- Assets belong in the owning space's `.gitbook/assets/`.

## Links: relative, never absolute

Use relative Markdown paths between pages (`../v1-api.md`), never absolute `https://docs.snyk.io/...` URLs. Absolute links do not survive restructuring and rot into dead links and soft 404s; a recent cleanup had to repoint over 120 of them. Absolute URLs are correct only for genuinely external destinations.

## Generated files — do not hand-edit

Edits to these paths are silently overwritten. Fix the generator or the upstream source instead.

| Path | Produced by |
| --- | --- |
| `developer-tools/snyk-api/reference/` | `tools/api-docs-generator`, from the OpenAPI specs |
| `developer-tools/snyk-api/changelog.md` | `tools/api-docs-generator` |
| `developer-tools/.gitbook/assets/rest-spec.json` | fetched from `https://api.snyk.io/rest/openapi` |
| `developer-tools/.gitbook/assets/v1-api-spec.yaml` | source of truth for the v1 API; owned by `@snyk/platformeng_api` |
| `tools/api-docs-generator/sync-state.yml` | generator changelog state |
| `developer-tools/snyk-ide-plugins-and-extensions/compatibility-matrix.md` | a workflow in [`snyk/snyk-ls`](https://github.com/snyk/snyk-ls) |
| `scan-fix-and-prevent/scan-with-snyk/error-catalog.md` | synced from [`snyk/error-catalog`](https://github.com/snyk/error-catalog) |
| `error-catalog/` | transient checkout during the sync workflow; gitignored, never commit it |

`developer-tools/SUMMARY.md` is partly machine-written: the generator rewrites its API reference entries. Edit the rest of that file normally, but do not hand-curate the reference section.

Automation that opens these pull requests lives in `.github/workflows/`: `sync-api-docs.yml` (hourly, Mon–Fri) and `sync-error-catalog.yml`.

## The API docs generator

`tools/api-docs-generator` is a Go module and the only code in this repository. Run from that directory:

```bash
make test    # go mod tidy, golangci-lint, go test ./...
make lint
make format  # goimports -local=github.com/snyk/user-docs/tools/api-docs-generator
make dry-run # generate without writing
make run     # regenerate into the repo — writes real files under developer-tools/
```

`make test` is the only test suite here; there are no tests for Markdown content. Prefer `make dry-run` unless you intend to commit generated output. This directory is owned by `@snyk/platformeng_api`, not the Docs team.

## Pull requests

- **Work on a branch in this repository, not a fork.** GitBook previews do not build for forks, so a fork cannot be reviewed properly.
- Sign your commits.
- One space and one concern per pull request. Content and generator changes do not belong together.
- CI runs a gitleaks secrets scan (CircleCI `prodsec/secrets-scan`); changes under `tools/` also run the generator tests. `pre-commit install` catches secrets locally.
- `CODEOWNERS` routes all content to `@snyk/design-content_docs`. Review is where writing style is enforced.
- The `/ship-it` Slack workflow is the intake path for **internal Snyk contributors only**, and a human runs it — it is not a step you perform. If you are working for an internal contributor, opening the pull request is not the last step and they still need to submit it. External contributions end at the pull request. See [README.md](README.md).

## Do not touch

- Any generated path in the table above.
- `.gitleaksignore` — audited allowlist, not noise to clean up.
- `.github/workflows/`, `.circleci/`, `catalog-info.yaml`, `CODEOWNERS` — owning teams review these.
- Directory structure and page locations, absent an explicit request from the Docs team.
