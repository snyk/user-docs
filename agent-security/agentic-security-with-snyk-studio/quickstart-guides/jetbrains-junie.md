---
description: How to set up Snyk Studio and the Snyk MCP server in JetBrains Junie
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect JetBrains Junie to Snyk

Connect JetBrains Junie to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* Install the code assistant extension
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Add the JetBrains Junie plugin to your JetBrains IDE (IntelliJ IDEA, WebStorm, PyCharm, and so on). Visit the official [JetBrains Junie getting started page](https://www.jetbrains.com/help/junie/get-started-with-junie.html).

## Install the Snyk MCP server in JetBrains Junie

Install the Snyk MCP server using the method that best suits your operating system and local development environment.

{% tabs %}
{% tab title="Node.js and npx" %}
If you have the Node.js npx executable installed, open **Settings/Preferences** > **Tools** > **Junie** > **MCP Settings**. Edit the MCP configuration and add the following instructions:

```json5
{
  "mcpServers": {
    "snyk": {
      "command": "npx",
      "args": [
        "-y",
        "snyk@latest",
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}
```

<figure><img src="../../.gitbook/assets/image3 (1).png" alt="Snyk MCP server command added to the Junie configuration"><figcaption></figcaption></figure>
{% endtab %}

{% tab title="Pre-installed Snyk CLI" %}
If you have the Snyk CLI installed, open **Settings/Preferences** > **Tools** > **Junie** > **MCP Settings**. Edit the MCP configuration and add the following instructions:

```json5
{
  "mcpServers": {
    "snyk": {
      "command": "/absolute/path/to/snyk",
      "args": [
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}
```

If the `snyk` command is not available, add it by following the instructions on the Install the Snyk CLI page.

Reload the IDE window for the changes to apply.
{% endtab %}
{% endtabs %}

## Set up the Snyk MCP server

Verify that the Snyk MCP server is available by checking the **Status** column in the MCP configuration and requesting it to list the tools it has available.

<figure><img src="../../.gitbook/assets/image1 (1).png" alt="Status column confirming the Snyk MCP server is available"><figcaption></figcaption></figure>

\
Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to working with the Snyk MCP server.

## Examples

In the chat, you can ask Junie to scan your code before deployment:

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Junie indicates in the plan that this request is related to security vulnerability scanning and calls the Snyk MCP server for various scans.

<figure><img src="../../.gitbook/assets/image2 (1).png" alt="Junie plan showing a call to the Snyk MCP server"><figcaption></figcaption></figure>
