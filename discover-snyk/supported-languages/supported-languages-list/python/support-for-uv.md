---
description: Snyk support for the uv package manager with Python, available in Early Access on Enterprise plans through the CLI and SCM
---

# Support for uv

{% hint style="info" %}
**Release status**

CLI and SCM support for uv is in Early Access and available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

## CLI support for uv

### Prerequites

Ensure you have `uv` version 0.9.29 or later installed.

Snyk uses the `uv.lock` file to build the dependency graph for a `uv` application. This file must be present for Snyk to identify the Project, scan dependencies, and identify issues. Ensure you check this file into your repository and that it reflects your current `pyproject.toml`. That is, ensure you run `uv lock` after adding, removing, or updating dependencies.

You do not need to run `uv sync` before you run the Snyk CLI.

Test results are not specific to the platform where you run the test. For example, if your `uv` Project has a Windows-only dependency with a vulnerability, this vulnerability also appears in the results if you run your test on Linux. Contact Snyk to provide feedback on this behavior or if you want to narrow down results to the platform where you run the test, either by default or using a flag.

### Supported commands and options <a href="#set-the-python-version-in-the-cli" id="set-the-python-version-in-the-cli"></a>

Snyk supports the `snyk test`, `snyk monitor`, and `snyk sbom` commands.

#### `snyk test`

The following flags, options, and arguments are not available:

* `--detection-depth`: There is no maximum depth for directory traversal. In later releases, Snyk scans only `uv` Projects in directories up to the given depth.
* `--print-deps`: The test incorrectly fails with a `Could not detect supported target files` error. In later releases, Snyk outputs dependency graphs for all found Projects without running a test.
* Input directory positional argument: When you run `snyk test <directory>`, Snyk ignores the directory. In later releases, Snyk scans only the specified directory.

The following flags and options are supported with limitations:

* `--file`: Snyk respects this flag only for `uv.lock` files. In later releases, Snyk supports `pyproject.toml` manifest files (both root and workspace packages).
* `--exclude`: This flag works only to exclude `uv.lock` files. In later releases, this option supports workspace packages. For example, `--exclude=foo` excludes all dependencies from `packages/foo/pyproject.toml`.

{% hint style="info" %}
For `uv`, the `--package-manager` option is not supported because `uv` does not permit lockfiles with names other than `uv.lock`.
{% endhint %}

#### `snyk monitor`

Snyk does not render error messages from the `uv` plugin in the traditional CLI error style.

#### `snyk sbom`

For SBOMs generated for `uv` Projects:

* When you generate an SBOM in a monorepo containing both `uv` and Maven Projects using the `--all-projects` flag, Snyk prunes some Maven dependencies in the resulting SBOM. All dependencies appear in the graph, but there are fewer instances than expected. When there are no `uv` Projects present, all dependencies are present as expected for Maven Projects.

## SCM support for uv

Snyk imports and tests  `uv` Projects directly from your connected Git repositories.

Snyk resolves the following:

* Direct and transitive dependencies from `uv.lock`.
* Git-sourced dependencies (`[tool.uv.sources]` with a git entry) in a standard, non-workspace project, and their transitive dependencies.
* For `uv` workspaces, Snyk imports the root `pyproject.toml`. Snyk includes workspace members that the root depends on as path dependencies in the root Snyk Project's graph.

### Prerequisites

The repository must contain a `uv` project: a `pyproject.toml` and a `uv.lock` file in the same directory. Both files are required.

You must have a connected SCM integration on an Organization, on the Enterprise plan, with the uv SCM Preview feature enabled.

### How Snyk discovers uv Projects

When you import a repository, Snyk scans every directory for a `pyproject.toml` and `uv.lock` pair and creates one Snyk Project per pair.

* Both files must be in the same directory. If a directory contains only a `pyproject.toml` (no lock file) or only a `uv.lock` (no manifest file), Snyk does not create a Project.
* When both `uv.lock` and `poetry.lock` exist for the same `pyproject.toml` ,  `uv.lock` takes precedence, and Snyk identifies the Project as a `uv` Project.
* Snyk identifies a `requirements.txt` in the same repository as a separate pip Project, independent of the `uv` Project.
* Snyk identifies Projects by path. Two directories with byte-identical manifests, therefore, become two separate Projects.

Snyk does not support the following dependencies:

* Dev-dependency groups (PEP 735 `[dependency-groups]`)
* Mixed git/URL/path sources in workspace members
* Optional extras (`[project.optional-dependencies]`)
