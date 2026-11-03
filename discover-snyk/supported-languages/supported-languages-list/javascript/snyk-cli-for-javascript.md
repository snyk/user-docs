---
nav_context: agnostic
description: >-
  How to test JavaScript Projects with the Snyk CLI, including report
  generation, output formats, and dependency filtering
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# CLI support for JavaScript

To help generate reports locally or at build time, visit the [snyk-to-html plugin](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html).

Use the `--json` and `--sarif` options to generate output that can be programmatically accessed.

For advanced filtering options, visit [snyk-filter](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter).

## Open Source libraries

The `snyk test` command tests the first manifest it can find and performs a test on that singular entry point. To have Snyk analyze all manifests in the directory, use the following options:

* `--all-projects`: This option detects and scans all Yarn and other Projects in this directory.
* `--yarn-workspaces`: Scans only Yarn Workspaces Projects. To test and monitor Yarn Workspaces together with packages from other package managers, use the `--all-projects` option instead.

{% hint style="info" %}
If you use a package manager that requires options, Snyk recommends targeting each manifest individually with `--file=`.
{% endhint %}

## Source code

* Framework support: visit [Supported languages, frameworks, and feature availability overview](../../supported-languages-package-managers-and-frameworks.md).
* Use the `snyk code test` command from the root of the Project to perform source code analysis.

## Containers

* As part of a container scan, Snyk automatically looks for application (open source) vulnerabilities. Snyk recommends integrating through the CLI earlier in the pipeline and using this as an additional signal of and insight into what is in production.
* If you ship your Node.js application in a container, be aware that the container can also bundle insecure packages (Linux, open source) alongside your application, in addition to the packages from the container base image. The Snyk Container CLI can help you identify a base image that minimizes the attack surface of your application.
* For more information on how you can filter to the layer you want to work on, such as identifying a secure base image to build on, the layers you are responsible for, or application (OS) vulnerabilities, visit [Snyk CLI for container security](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container).

## Infrastructure as code

Visit [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/).

## Resources

Visit the [CLI commands and options summary](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/cli-commands-and-options-summary) and the [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/). Use the `--help` option in the CLI for details of Snyk CLI commands.
