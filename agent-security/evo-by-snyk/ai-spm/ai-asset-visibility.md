# AI asset visibility

The **Discovery agent** maps the AI assets in your code so you have a complete view of your AI usage. It runs an AI Bill of Materials (AI-BOM) scan to detect AI components in your source code.

## Import repositories

Import repositories into Evo through the Workspaces integration or with a command-line interface (CLI) upload. After you enable Evo, Snyk scans the repositories you import within 24 hours. Scan duration depends on the number of repositories and the volume of code. Scan status appears on the **Scans and jobs** page.

## Supported languages and files

Evo scans Python and TypeScript codebases. It also inspects non-code files to identify models, datasets, and other AI assets, including these extensions:

Here is the full list of extensions that Evo is currently dissecting:  `.yaml`, `.yml`, `.xml`, `.json`, `.toml`, `.ini`, `.conf`, `.env`, `.dockerfile`, `.containerfile`, `Gemfile`, `Makefile`, `build.grade`, `go.mod`, and `requirements.txt`.

## Asset types

An AI-BOM scan detects the following asset types:

* Models: foundation and open-source models. When available, Evo provides model card metadata, license information, and documentation links.
* Agents: agents built on common agent libraries and orchestration frameworks.
* Tools: tools that use common tool-calling patterns and registration hooks.
* MCP servers: Model Context Protocol (MCP) servers built with the official MCP SDK and other known server patterns.

Asset details appear on the **Inventory** page and update after each scan.

## Custom discovery

Custom discovery finds AI components unique to your codebase - internal model wrappers, custom naming conventions, and proprietary agent frameworks - that generic scanners do not detect.

During a regular scan, Evo identifies candidates that resemble AI components, validates each one with large language model (LLM) inference, and surfaces it on the **Inventory** page with a confidence level. You review each candidate and accept or reject it. After you accept a candidate, it becomes a custom detection rule scoped to your tenant and applies to every later scan.

{% hint style="info" %}
Evo scopes custom patterns to your tenant and shares no data across tenants. It uses LLM inference to validate candidates, not to train models.
{% endhint %}

### Gate deployments

To block a deployment in your continuous integration/continuous delivery (CI/CD) pipeline based on your AI-BOM results, run the `aibom test` command in the Snyk CLI.
