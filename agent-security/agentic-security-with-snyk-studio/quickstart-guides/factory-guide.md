---
description: How to configure the local Snyk MCP server for the Factory App and Droid
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Factory to Snyk

Connect Factory to the local Snyk MCP server to scan code and dependencies from your coding assistant. The Factory App and Droid use the user-level MCP configuration at `~/.factory/mcp.json`.

{% hint style="info" %}
This is a manual MCP integration. It does not install Snyk Studio Secure at inception hooks or Agentic Development Security (ADS) Agent Behavior Governance hooks for Factory. Rules guide the model when invoking scans and do not guarantee tool execution.
{% endhint %}

## Install Factory

Install the [Factory App](https://docs.factory.com/factory-app/quickstart), or use [Droid in your terminal or IDE](factory-terminal-ide-guide.md). Install and authenticate the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli), and enable the Snyk products you want to scan with.

```bash
snyk auth
```

## Install the Snyk MCP server in Factory

Create or edit `~/.factory/mcp.json`. Merge the Snyk entry into the existing `mcpServers` object and preserve your other server entries:

```json
{
  "mcpServers": {
    "Snyk": {
      "type": "stdio",
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"]
    }
  }
}
```

Replace the `command` value with the absolute path to your Snyk CLI executable. For configuration scope and server management, visit the [Factory MCP documentation](https://docs.factory.com/harness/mcp).

## Set up the Snyk MCP server

Open the repository you want to scan. Check that Factory loads the Snyk MCP server, then complete any Snyk authentication or repository-trust request.

In Droid, use `/mcp` to inspect the server connection and available tools. For scan or authentication failures, visit [Troubleshooting](../troubleshooting.md).

## Secure at inception rules

To guide scans during development, add the generic rule text to your repository's `AGENTS.md`, for example under a `## Security scanning` heading. For the instruction file's location and behavior, visit the [Factory AGENTS.md documentation](https://docs.factory.com/harness/agents-md). For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

These are model instructions, not programmatic hooks. Visit [Directives](../directives.md) for the differences between rules, hooks, commands, and skills.

## Examples

Ask Factory to invoke Snyk scanning tools:

```text
Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
```

Review the tool calls and scan results. A successful direct MCP scan confirms that Factory can invoke Snyk tools. It does not verify automatic Studio hook scanning.
