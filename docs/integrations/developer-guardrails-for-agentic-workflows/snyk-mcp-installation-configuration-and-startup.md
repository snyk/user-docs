# Snyk MCP installation, configuration, and startup

To use the Snyk MCP server, download and install the Snyk CLI [v1.1298.0](https://github.com/snyk/cli/releases/tag/v1.1298.0) or later following the steps on the [installation page](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/). No other dependencies are needed. Snyk recommends always using the latest version of the CLI.

## Environment variables for MCP

You can set CLI environment variables for the MCP server in the following ways:

* In the MCP server configuration file, whether in your IDE or on the MCP host
* Directly on your system

For a full list of supported CLI environment variables, see [Environment variables for Snyk CLI](../../developer-tools/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli.md).

## Starting the Snyk MCP server

To start the Snyk MCP server, use the `snyk mcp` command for a supported transport type, `stdio` or `sse` as follows:

`snyk mcp -t sse` - Start the Snyk MCP server using `sse`, HTTP Server-Sent Events Transport. The available endpoint is `/sse`.

`snyk mcp -t stdio` - Start the Snyk MCP server using `stdio`, Stdio (Standard IO) Transport.
