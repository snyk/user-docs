# Quickstart guides for Snyk Studio

{% hint style="info" %}
**Feature availability**

Snyk does not offer a hosted, remote version of its MCP server.

The Snyk MCP Server is designed as a local MCP server, running on your system using the Snyk CLI to ensure local file access
{% endhint %}

To add an MCP server to an AI IDE or agent, consult the documentation for the AI system where you plan to integrate Snyk and review the specific MCP instructions.

This section includes quickstart guides for use with common AI coding tools.

For manual installation and use of Snyk Studio, download and install the Snyk CLI [v1.1298.0](https://github.com/snyk/cli/releases/tag/v1.1298.0) or later following the steps on the [installation page](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/). No other dependencies are needed. Snyk recommends always using the latest version of the CLI.

## Environment variables for MCP

You can set CLI environment variables for the MCP server in the following ways:

* In the MCP server configuration file, whether in your IDE or on the MCP host
* Directly on your system

For a full list of supported CLI environment variables, see [Environment variables for Snyk CLI](../../../developer-tools/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli.md).

## Starting the Snyk MCP server

Check our Quickstart guides for MCP. If you don’t find a guide for your specific AI agent or IDE, you can use the following sample configuration snippet:

```
{
  "mcpServers": {
    "Snyk": {
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```
