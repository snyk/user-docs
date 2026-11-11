---
description: How to connect Droid in a terminal or IDE to the local Snyk MCP server
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Factory Terminal/IDE to Snyk

Connect Droid in your terminal or IDE to the local Snyk MCP server. Use the same user-level MCP configuration as the Factory App: `~/.factory/mcp.json`.

<a id="install-factory-terminalide"></a>
<a id="install-snyk-studio"></a>
<a id="authenticate"></a>
<a id="run-snyk-studio"></a>

## Connect Droid to Snyk

1. Follow the [Factory IDE integration instructions](https://docs.factory.com/ide-integrations) for your environment.
2. Follow [Install the Snyk MCP server in Factory](factory-guide.md#install-the-snyk-mcp-server-in-factory) to add the Snyk entry to `~/.factory/mcp.json`. Use this shared configuration for the Factory App and Droid.
3. Run `snyk auth` as the developer who runs Droid. Open the repository and complete any Snyk repository-trust request.
4. In Droid, use `/mcp` to inspect the Snyk connection and tool list. Ask the assistant to scan the repository:

   ```text
   Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
   ```

## Set up the Snyk MCP server

For the configuration example and troubleshooting handoff, visit [Connect Factory to Snyk](factory-guide.md). This setup provides agent-invoked MCP tools. It does not install Studio Secure at inception hooks or Agentic Development Security (ADS) Agent Behavior Governance hooks for Factory.

## Secure at inception rules

Use the [Factory project instructions](factory-guide.md#secure-at-inception-rules) to request scans during development. Rules rely on the model following the instructions and do not guarantee scan execution. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

## Examples

Use the [Factory scanning example](factory-guide.md#examples) and check the actual tool calls and scan results.
