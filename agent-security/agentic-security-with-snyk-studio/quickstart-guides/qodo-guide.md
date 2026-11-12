---
description: How to set up Snyk Studio and the Snyk MCP server in Qodo
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Qodo to Snyk

{% hint style="warning" %}
Qodo IDE plugin support ends October 30, 2026. Qodo also deprecates its autocomplete and chat-based generation features. Treat the procedures on this page as legacy setup. Visit [Qodo IDE best practices](https://docs.qodo.ai/qodo-ide/code-intelligence/best-practices) for the current product notice and migration guidance.
{% endhint %}

Connect Qodo to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* Install the code assistant extension
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Add the Qodo extension to the VS Code or JetBrains IDE. Visit the official [Qodo.ai](https://qodo.ai/) page.

## Install the Snyk MCP server in Qodo

Install the Snyk MCP server using the method that best suits your operating system and local development environment.

In the Qodo chatbox, click the **Add MCP** **tools** > **Add new MCP**. Provide the required details for the Snyk MCP server:

* The MCP Server name
* `command`: Either `npx` or the absolute path to the Snyk CLI.
* `args`: The Snyk MCP command to start the MCP server.

`npx` option:

```json
{
    "Snyk": {
        "command": "npx",
        "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"]
    }
}

```

absolute path option:

```json
{
    "Snyk": {
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"]
  }
}
```

For the absolute path option, if the `snyk` command is not available, add it by following the instructions on the [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli) page.

The following example shows a Snyk MCP server that was successfully configured.

<figure><img src="../../.gitbook/assets/image (41).png" alt="Snyk MCP server configured in Qodo"><figcaption></figcaption></figure>

For additional MCP configuration options on VS Code and troubleshooting, consult the official [VS Code MCP server documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

## Set up the Snyk MCP server

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

Ask the assistant to authenticate or trust the repository when needed. Confirm that it invokes the Snyk MCP tools and review the returned results.

<figure><img src="../../.gitbook/assets/image (35).png" alt="Snyk MCP server running a workflow in Qodo"><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

<figure><img src="../../.gitbook/assets/image (36).png" alt="Completing the authenticate and trust prompt"><figcaption></figcaption></figure>

## Secure at inception rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

The following applies to the legacy Qodo IDE plugin. Add the generic rule text to `best_practices.md` in your project root, after this front matter. Review Qodo's retirement notice before using this workflow.

{% code title="best_practices.md" overflow="wrap" %}
```markdown
---
applyTo: "**"
---
```
{% endcode %}

## Qodo workflows

Create a custom Snyk workflow to scan and remediate security vulnerabilities. The following is a suggestion for [Qodo workflows](https://docs.qodo.ai/qodo-documentation/qodo-gen/agent/workflows). Import this legacy workflow from a TOML file in `~/.qodo/agents/` or through the workflows interface. Review Qodo's retirement notice before using it.

```toml
# Version of the agent configuration standard
version = "1.0"

[commands.remediate]
available_tools = [
  "Snyk",
  "Code Navigation",
  "filesystem",
  "git",
  "Terminal",
  "Web Search"
]
description = "Remediate application security issues with Snyk Studio"
instructions = """
You are an expert at developing secure code. Snyk flagged several SAST and SCA issues in this application.

Please do the following:

1. Use Snyk to get context on the issues.
2. Fix all of the issues.
3. Use Snyk again to confirm the issues are fixed. If the issues are not fixed, go back to the previous step.
4. Build the app and run unit tests to ensure you didn’t break it. I’ll run it and do additional validation when you’re done.

Thanks!"""
mcpServers = """
{
  "Snyk": {
    "command": "npx",
    "type": "CUSTOM",
    "env": {},
    "args": [
      "-y",
      "snyk@latest",
      "mcp",
      "-t",
      "stdio"
    ]
  }
}"""
```

## Examples

In the chat, you can tag specific files or functions and ask Qodo to scan your code before deployment:

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Review the assistant tool calls to confirm that it invoked the requested Snyk MCP scans.

Review the scan results and any fixes that the assistant proposes.
