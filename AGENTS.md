# Working in this repository

This repository holds the source files for the [Snyk user documentation](https://docs.snyk.io), published through GitBook. The notes below cover the mechanics of making a change here. They do not cover writing style: follow the `snyk-docs-writing-rules` skill and the internal [writing user documentation](https://snyksec.atlassian.net/wiki/spaces/DRC/pages/1819541615/Writing+user+documentation) page for that.

## Repository layout

Each top-level directory is a separate GitBook space with its own navigation file:

| Directory | Navigation file |
| --- | --- |
| `agent-security/` | `agent-security/SUMMARY.md` |
| `developer-tools/` | `developer-tools/SUMMARY.md` |
| `discover-snyk/` | `discover-snyk/SUMMARY.md` |
| `platform-administration/` | `platform-administration/SUMMARY.md` |
| `scan-fix-and-prevent/` | `scan-fix-and-prevent/SUMMARY.md` |
| `snyk-data-and-governance/` | `snyk-data-and-governance/SUMMARY.md` |

When you add a page, add it to the `SUMMARY.md` of that space only. There is no repository-wide navigation file, so a new page that is not listed in its own space's `SUMMARY.md` does not appear in the published navigation.

Images and other assets belong in the `.gitbook/assets/` directory of the space that uses them. Link to other pages using repository-relative paths, matching the surrounding pages.

`tools/` is not documentation. It holds the API docs generator described below. `evo-by-snyk/` currently holds only assets and has no navigation file.

## Generated files: do not edit by hand

These paths are produced by automation and any manual edit is overwritten on the next run:

| Path | Produced by |
| --- | --- |
| `developer-tools/snyk-api/reference/` | `tools/api-docs-generator` via `.github/workflows/sync-api-docs.yml` |
| `developer-tools/snyk-api/changelog.md` | Same |
| `developer-tools/.gitbook/assets/rest-spec.json` | Fetched from the API spec source |
| `developer-tools/.gitbook/assets/v1-api-spec.yaml` | Same |
| `scan-fix-and-prevent/scan-with-snyk/error-catalog.md` | Generated from [snyk/error-catalog](https://github.com/snyk/error-catalog) via `.github/workflows/sync-error-catalog.yml` |

The API sync runs hourly on weekdays, so edits to those paths are usually reverted within the hour. The generator also writes the API reference entries in `developer-tools/SUMMARY.md`; when editing that file, change only the entries for hand-written pages.

To change generated API reference content, change the API spec or `tools/api-docs-generator` rather than the output.

## Making a change

- Work on a branch in this repository rather than a fork. GitBook builds previews only for branches, and reviewers rely on those previews.
- Sign your commits. This is required for both internal and external contributions.
- Keep pull requests scoped to one space where possible, since ownership is per-path.
- `.github/CODEOWNERS` assigns review. `@snyk/design-content_docs` owns the repository by default, and API spec assets and workflow files are co-owned by `@snyk/platformeng_api`.

## Checks that run

- `gitleaks` runs as a pre-commit hook, configured in `.pre-commit-config.yaml`. Install the hooks with `pre-commit install` before committing.
- CircleCI runs a Snyk ProdSec secrets scan on every pull request, defined in `.circleci/config.yml`.

Never commit tokens, API keys, or customer data. Examples in documentation must use clearly fake values.

## Requesting review and publication

Documentation and release-notification requests go to the Docs team through the `/ship-it` Slack workflow, which creates the tracking ticket. Open the pull request first and have its URL ready before running the workflow. See [README.md](README.md) for the prerequisites.
