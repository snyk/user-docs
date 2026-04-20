# CLI support for uv

{% hint style="info" %}
**Release status**

CLI support for uv is in Early Access and available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://docs.snyk.io/snyk-admin/manage-settings/snyk-preview).
{% endhint %}

## Prerequites

Ensure you have `uv` version 0.9.29 or later installed.

Snyk uses the `uv.lock` file to build the dependency graph for a  `uv` application. This file must be present for Snyk to identify the Project, scan dependencies, and identify issues. Ensure you check this file into your repository and that it reflects your current `pyproject.toml`. That is, ensure you run `uv lock` after adding, removing, or updating dependencies.

You do not need to run `uv sync` before you run the Snyk CLI.

Test results are not specific to the platform where you run the test. For example, if your `uv` Project has a Windows-only dependency with a vulnerability, this vulnerability also appears in the results if you run your test on Linux. Contact Snyk to provide feedback on this behavior or if you want to narrow down results to the platform where you run the test, either by default or using a flag.

## Supported commands and options <a href="#set-the-python-version-in-the-cli" id="set-the-python-version-in-the-cli"></a>

Snyk supports the `snyk test`, `snyk monitor`, and `snyk sbom` commands.

### `snyk test`

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

### `snyk monitor`

Snyk does not render error messages from the `uv` plugin in the traditional CLI error style.

### `snyk sbom`

For SBOMs generated for `uv` Projects:

* When you generate an SBOM in a monorepo containing both `uv` and Maven Projects using the `--all-projects` flag, Snyk prunes some Maven dependencies in the resulting SBOM. All dependencies appear in the graph, but there are fewer instances than expected. When there are no `uv` Projects present, all dependencies are present as expected for Maven Projects.
