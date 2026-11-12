---
description: How to set up Snyk Studio and the Snyk MCP server in Augment Code
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Augment Code to Snyk

Connect Augment Code to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* [Install the code assistant extension](https://www.augmentcode.com/)
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Visit the [Augment Code](https://www.augmentcode.com/) website to download the correct version of the IDE plugin.

## Install the Snyk MCP server in Augment Code

Install the Snyk MCP server using the method that best suits your operating system and local development environment.

{% tabs %}
{% tab title="Node.js and npx" %}
* Name: Snyk
* Command: npx -y snyk@latest mcp -t stdio
{% endtab %}

{% tab title="Pre-installed Snyk CLI" %}
* Name: Snyk
* Command: /absolute/path/to/snyk mcp -t stdio

If the `snyk` command is not available, add it by following the instructions on the [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli) page.
{% endtab %}
{% endtabs %}

## Set up the Snyk MCP server

The MCP server starts automatically. The following example shows a Snyk MCP server that was successfully configured and started.

<figure><img src="../../.gitbook/assets/image (3).png" alt="Snyk MCP server configured and started in Augment Code"><figcaption></figcaption></figure>

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to work with the Snyk MCP server.

## Secure at inception rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

Add the generic rule text to your Augment Code user guidelines.

<figure><img src="../../.gitbook/assets/image (5).png" alt="Secure at inception rules added to the Augment Code configuration"><figcaption></figcaption></figure>

## Examples <a href="#examples" id="examples"></a>

In the chat, you can tag specific files or functions and ask Augment Code to scan your code for safe deployment to production.

```text
Scan this directory for code security & dependency vulnerabilities and security issues
```

Augment Code indicates that this request is related to security vulnerability scanning and calls the Snyk MCP server for various scans.

<figure><img src="../../.gitbook/assets/image (4).png" alt="Augment Code calling the Snyk MCP server for a vulnerability scan"><figcaption></figcaption></figure>
