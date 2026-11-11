---
description: How to connect Devin Desktop to the local Snyk MCP server
nav_context: classic
last_reviewed: 2026-08-20
last_reviewed_source: focused-edit
---

# Connect Devin Desktop to Snyk

Connect the Devin Local agent in Devin Desktop to Snyk tools to scan code and dependencies.

[Devin Desktop is the new name for Windsurf](https://devin.ai/blog/windsurf-is-now-devin-desktop). Cognition removed the legacy Cascade agent from Devin Desktop on September 8, 2026. Visit the [Devin Desktop changelog](https://docs.devin.ai/desktop/changelog).

{% hint style="info" %}
This procedure configures agent-invoked MCP access. It does not install Studio Secure at inception hooks, remediation commands, or package-health skills.

The Studio installer still accepts `--ade=windsurf` and writes legacy configuration under `.codeium/windsurf`. It does not configure the Devin Local MCP files. Devin Local uses skills rather than legacy Cascade workflows.
{% endhint %}

## Configure the local Snyk MCP server

Install [Devin Desktop](https://devin.ai/desktop) and complete the account and product setup in [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md).

1. Authenticate the Snyk CLI:

   ```bash
   snyk auth
   ```

2. Add the Snyk MCP server to `.devin/mcp_config.local.json` in your repository. Preserve existing entries in `mcpServers`:

   ```json
   {
     "mcpServers": {
       "Snyk": {
         "command": "/absolute/path/to/snyk",
         "args": ["mcp", "-t", "stdio"]
       }
     }
   }
   ```

   Replace the `command` value with the absolute path to your Snyk CLI executable. This local configuration is not shared with your team. Visit [Devin MCP configuration](https://docs.devin.ai/cli/extensibility/mcp/configuration) for other scopes.

3. Start a new Devin Local session in the repository and ask the assistant to scan it:

   ```text
   Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
   ```

4. Approve the Snyk MCP tool calls and complete Snyk authentication or repository trust requests when prompted. Inspect the tool calls and scan results to verify the connection.

## Remove the MCP connection

Remove the `Snyk` entry from the configuration file where you added it, then start a new assistant session. Preserve other MCP server entries.

If the connection fails, visit [Troubleshooting](../troubleshooting.md).

If the connection fails, visit [Troubleshooting](../troubleshooting.md).
