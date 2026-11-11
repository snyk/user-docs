---
description: How to set up Snyk Studio and the Snyk MCP server in Goose CLI
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Goose CLI to Snyk

<a id="install-goosecli"></a>

Connect Goose CLI to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* [Install the Goose CLI](goose-cli-guide.md#install-goosecli)
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Install the Goose CLI. For more details, visit the official [Goose CLI installation instructions.](https://goose-docs.ai/docs/quickstart/)

## Install the Snyk MCP server in Goose CLI using npx command

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

To install the Snyk MCP server using the Goose CLI interactive session, proceed with the following instructions:

*   Start an interactive Goose configuration wizard by running the `goose configure` command:

    <figure><img src="../../.gitbook/assets/image (288).png" alt="Goose CLI interactive session used to add the Snyk MCP server"><figcaption></figcaption></figure>
* From the configuration menu, chose **Add Extension.**
* Choose **Command-line extension (run a local command or script).**
* Name this extension **Snyk Security.**
* Run the command by typing in: `npx -y snyk@latest mcp -t stdio`
* Continue with the defaults (timeout of 300 seconds), and optionally provide the Snyk API token if your Organization policy requires one.

Use the following Goose configuration example:

<figure><img src="../../.gitbook/assets/image (289).png" alt="Completed Goose configuration walkthrough"><figcaption></figcaption></figure>

Start or resume your Goose session with the Snyk MCP server enabled.

## Install the Snyk MCP server in Goose CLI using Snyk CLI

This installation instruction assumes you have installed the Snyk CLI and the `snyk` executable is available in your system path.

Install the Snyk MCP server using the Goose CLI interactive session by following these instructions:

* Start an interactive Goose configuration wizard by running the `goose configure` command.
* Choose **Add Extension f**rom the configuration menu.
* Select **Command-line extension (run a local command or script).**
* Name this extension **Snyk Security.**
* Run the command by typing in: `/path/to/snyk mcp -t stdio` - replacing `/path/to/snyk` with the fully qualified absolute path to your Snyk CLI.
* Continue with the defaults (timeout of 300 seconds), and optionally provide the Snyk API token if your Organization policy requires one.

{% hint style="info" %}
If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the absolute path to the `snyk` executable. For example: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

Start or resume your Goose session with the Snyk MCP server enabled.

## Set up the Snyk MCP server

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

The assistant invokes the authentication and trust workflows. Complete any browser confirmation request.

You can explicitly prompt to authenticate your account as follows:

```text
Authenticate my Snyk account
```

<figure><img src="../../.gitbook/assets/image (291).png" alt="Prompting Goose CLI to authenticate the Snyk account"><figcaption></figcaption></figure>

If you are not authenticated, this prompt opens a browser for you to authenticate to the Snyk Web UI, and then you can continue with the agent.

<figure><img src="../../.gitbook/assets/image (10).png" alt="Browser opened to authenticate to the Snyk Web UI" width="375"><figcaption></figcaption></figure>

## Secure at inception rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

Use the generic rule text with [Goose CLI recipes](https://goose-docs.ai/docs/guides/recipes/session-recipes/) or with the [Goose hints file](https://goose-docs.ai/docs/guides/context-engineering/using-goosehints/). Add it to the `.goosehints` file in the root directory of the project.

## Examples

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Goose CLI invokes Snyk scanning, providing output, and insights based on the results that Snyk returned:

<figure><img src="../../.gitbook/assets/image (292).png" alt="Goose CLI reporting the results of a Snyk scan"><figcaption></figcaption></figure>

If the connection fails, visit [Troubleshooting](../troubleshooting.md).
