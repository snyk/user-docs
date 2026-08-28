---
nav_context: agnostic
---

# AI asset visibility

The **Discovery agent** maps AI assets in your code to provide a complete view of your AI usage. It runs an AI Bill of Materials (AI-BOM) scan to detect AI components in your source code.

## Import repositories

Import repositories into Evo through the **Workspaces** integration or with a command-line interface (CLI) upload. After you enable Evo, Snyk scans the repositories you import within 24 hours. Scan duration depends on the number of repositories and the volume of code. Scan status appears on the **Scans and jobs** page.

## Supported languages and files

Evo scans Go, Java, JavaScript, Python, and TypeScript codebases. It also inspects non-code files to identify referenced models.

Evo scans files with these extensions: `.conf`, `.containerfile`, `.dockerfile`, `.env`, `.ini`, `.json`, `.properties`, `.tf`, `.tfvars`, `.toml`, `.xml`, `.yaml`, and `.yml`. It also scans files with these names: `build.gradle`, `build.gradle.kts`, `cargo.toml`, `composer.json`, `Containerfile`, `Dockerfile`, `environment.yml`, `Gemfile`, `go.mod`, `Justfile`, `Makefile`, `package.json`, `Pipfile`, `pom.xml`, `pubspec.yaml`, `pyproject.toml`, and `requirements.txt`.

## Asset types

An AI-BOM scan detects these asset types:

* Models: Foundation and open-source models. Evo provides model card metadata, license information, and documentation links for these models.
* Packages: AI libraries and software development kits (SDKs), detected from the imports and calls in your code.
* Agents: Agents built on common agent libraries and orchestration frameworks.
* Tools: Tools that use common tool-calling patterns and registration hooks, and tools that a Model Context Protocol (MCP) server exposes.
* MCP servers: MCP servers built with the official MCP SDK and other known server patterns.
* MCP clients: Code that connects to an MCP server as a client.
* MCP resources: Resources that an MCP server exposes.
* Datasets: Training datasets associated with a model in your inventory.

Evo identifies models by scanning configuration and manifest files. It identifies other asset types from the source code.

If the Snyk Risk Database contains a record for a model Evo discovers, Evo also adds the base models it derives from and the training datasets associated with it. These assets come from Snyk research rather than from your repository.

Asset details appear on the **Inventory** page and update after each scan.

## Custom discovery

Custom discovery finds AI components unique to your codebase. This includes internal model wrappers, custom naming conventions, and proprietary agent frameworks that generic scanners do not detect.

During a regular scan, Evo identifies candidates that resemble AI components. It validates each candidate using a large language model (LLM) inference and displays the result on the **Inventory** page with a confidence level. Review each candidate and accept or reject it. After you accept a candidate, it becomes a custom detection rule. This rule is scoped to your tenant and applies to all subsequent scans.

{% hint style="info" %}
Evo scopes custom patterns to your tenant and does not share data across tenants. It uses LLM inference to validate candidates, not to train models.
{% endhint %}

### Gate deployments

To block a deployment in your continuous integration/continuous delivery (CI/CD) pipeline based on your AI-BOM results, run the `aibom test` command in the Snyk CLI.
