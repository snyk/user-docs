---
description: How to set up Snyk Studio and the Snyk MCP server in Cline
nav_context: classic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Connect Cline to Snyk

Connect Cline to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* Install Cline
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Consult the official Cline documentation for [Cline installation instructions](https://docs.cline.bot/getting-started/installing-cline).

## Install the Snyk MCP server in Cline using npx

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

Open the Cline MCP configuration from **MCP Servers** in the top toolbar. Select **Configure**, then **Configure MCP Servers**. Visit the [Cline MCP documentation](https://docs.cline.bot/mcp/mcp-overview).

<figure><img src="../../.gitbook/assets/image (352).png" alt="Manage MCP Servers icon in the Cline bottom panel"><figcaption></figcaption></figure>

Click the **Configure MCP Servers** button from the **MCP Servers** tab:

<figure><img src="../../.gitbook/assets/image (353).png" alt="Configure MCP Servers button on the MCP Servers tab"><figcaption></figcaption></figure>

To install the Snyk MCP server, add the following `mcpServers` configuration block to the Cline MCP Servers configuration file:

```json5
{
  "mcpServers": {
    "Snyk Security": {
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"]
    }
  }
}
```

## Install the Snyk MCP server in Cline using Snyk CLI

Ensure that the Snyk CLI is installed and that the `snyk` executable is available in your system path.

To install the Snyk MCP server with the Snyk CLI, follow these instructions:

*   Add the following `mcpServers` configuration block to the Cline MCP Servers configuration file:

    ```json5
    {
      "mcpServers": {
        "Snyk Security": {
          "command": "/absolute/path/to/snyk",
          "args": ["mcp", "-t", "stdio"]
        }
      }
    }
    ```

{% hint style="info" %}
If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the absolute path to the `snyk` executable. For example: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

After saving the MCP Servers configuration, confirm the Snyk MCP server is successfully installed and listed in the Cline UI.

<figure><img src="../../.gitbook/assets/image (354).png" alt="Snyk MCP server listed as installed in Cline"><figcaption></figcaption></figure>

## Set up the Snyk MCP server

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

The assistant invokes the authentication and trust workflows. Complete any browser confirmation request.

<figure><img src="../../.gitbook/assets/image (9).png" alt="Cline running a Snyk workflow through the MCP server"><figcaption></figcaption></figure>

You can explicitly prompt to authenticate your account as follows:

```text
Authenticate my Snyk account
```

<figure><img src="../../.gitbook/assets/image (355).png" alt="Asking Cline to authenticate the Snyk account"><figcaption></figcaption></figure>

After enabling Act mode and approving the request, the agent opens a browser for you to authenticate to the Snyk Web UI, then continues the session.

<figure><img src="../../.gitbook/assets/image (10).png" alt="Browser opened by the agent to authenticate to the Snyk Web UI" width="375"><figcaption></figcaption></figure>

## Secure at inception rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

The following is a suggested instruction for [Cline rules](https://docs.cline.bot/features/cline-rules). Open the Rules and Workflows menu using the bottom chat navigation from Cline.

<figure><img src="../../.gitbook/assets/image (12).png" alt="Snyk rules added to the Cline rules configuration"><figcaption></figcaption></figure>

Add the generic rule text as a global rule.

## Examples

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Cline invokes Snyk scanning tools and summarizes the results:

<figure><img src="../../.gitbook/assets/image (11).png" alt="Cline reporting the results of a Snyk scan"><figcaption></figcaption></figure>

If the connection fails, visit [Troubleshooting](../troubleshooting.md).
