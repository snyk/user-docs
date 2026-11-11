---
description: >-
  How to set up Snyk Studio and the Snyk MCP server in Gemini Code Assist
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Gemini Code Assist to Snyk

{% hint style="warning" %}
Google has ended Gemini CLI and Gemini Code Assist access for some account types and is retiring it for others. Visit [Migrate from Gemini tooling](antigravity-guide.md#migrate-from-gemini-tooling) for the schedule and Antigravity setup. The Studio installer's `gemini` target configures Gemini CLI, not Antigravity CLI.
{% endhint %}

Connect Gemini Code Assist to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* Install the code assistant extension
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Install the Gemini Code Assist extension in VS Code or a supported JetBrains IDE. Agent mode is in Preview. Visit the [Gemini Code Assist getting started page](https://codeassist.google/).

## Install the Snyk MCP server in Gemini Code Assist

Install the Snyk MCP server using the method that best suits your operating system and local development environment.

{% tabs %}
{% tab title="Node.js and npx" %}
For VS Code, create or edit `~/.gemini/settings.json`. Google documents a separate `mcp.json` location for IntelliJ. Follow the [IDE-specific MCP setup](https://docs.cloud.google.com/gemini/docs/codeassist/use-agentic-chat-pair-programmer#configure-mcp-servers) for JetBrains.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

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
{% endtab %}

{% tab title="Pre-installed Snyk CLI" %}
For VS Code, create or edit `~/.gemini/settings.json`. Google documents a separate `mcp.json` location for IntelliJ. Follow the [IDE-specific MCP setup](https://docs.cloud.google.com/gemini/docs/codeassist/use-agentic-chat-pair-programmer#configure-mcp-servers) for JetBrains.

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. Specify the absolute path to the Snyk CLI executable:

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

If the `snyk` command is not available, add it by following the instructions on the [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli) page.
{% endtab %}
{% endtabs %}

Reload the IDE window to apply the changes. In VS Code, open the **Command Palette** and select **Developer: Reload Window**.

Next, you can verify that the Snyk MCP server is available by enabling Agent mode on Gemini Code Assist and requesting it to list the tools it has available:

<figure><img src="../../.gitbook/assets/image (28).png" alt="Snyk MCP server available in Gemini Code Assist agent mode"><figcaption></figcaption></figure>

## Set up the Snyk MCP server

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to working with the Snyk MCP server.

## Gemini Code Assist security rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

Add the generic rule text as project context for Gemini Code Assist. For VS Code, use `GEMINI.md`. For IntelliJ, use `AGENT.md` or IDE instructions scoped to the project. Visit [Google agent context guidance](https://docs.cloud.google.com/gemini/docs/codeassist/use-agentic-chat-pair-programmer#create-a-context-file).

## Examples

In the chat, you can tag specific files or functions and ask Gemini to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Check the tool calls to confirm that Gemini invoked the requested Snyk MCP scans.

<figure><img src="../../.gitbook/assets/image (29).png" alt="Gemini calling the Snyk MCP server for a vulnerability scan"><figcaption></figcaption></figure>

If the connection fails, visit [Troubleshooting](../troubleshooting.md).
