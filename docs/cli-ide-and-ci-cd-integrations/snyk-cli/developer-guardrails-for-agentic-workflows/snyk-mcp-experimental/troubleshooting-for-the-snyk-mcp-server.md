# Troubleshooting for the Snyk MCP server

If you encounter issues with the Snyk MCP server or its integration, try the troubleshooting steps provided here.

## Ensure your Snyk CLI version is compatible

* After downloading or updating the CLI, run `snyk version`.
* The version must be ≥ v1.1296.2. Snyk recommends using the latest version.

## Verify Snyk CLI path and permissions

* If you have specified a direct path to the `snyk` executable in your `mcpconfig.json` , double-check that this path is correct.
* Ensure the Snyk CLI binary has execute permissions.

## Basic repository scanning (crucial diagnostic)

* This is a key step for many issues. Before suspecting complex MCP integration problems, confirm that the Snyk CLI you are using to run the MCP server can scan your repository directly from your terminal.
* Navigate to the root directory of your Project and run:
  * `/path/to/your/snykCli test` (for open-source vulnerabilities)
  * `/path/to/your/snykCli code test` (for code issues)
* If these direct scans fail, resolve those issues first, for example, authentication, Organization settings, and Snyk Code enablement for your Organization.

## Authentication issues

* Some MCP hosts (the client application integrating the Snyk MCP server) might restrict MCP server processes, which can interfere with the Snyk authentication flow, for example, browser-based login.
* Mitigation strategies
  * Try starting the Snyk MCP server in `sse` transport mode instead of `stdio`: `snyk mcp -t sse --experimental` and set the URL in your `mcpconfig.json` file.
  * Provide a Snyk authentication token directly using the `SNYK_TOKEN` environment variable.

## Snyk Organization configuration

If your Snyk account is part of multiple Organizations, or if scans are not appearing in the expected place, ensure the correct Snyk Organization is configured. You can set this using:

* The command `snyk config set org=<YOUR_ORG_ID>`
* The environment variable `SNYK_CFG_ORG=<YOUR_ORG_ID>`

## Environment variable propagation

Verify that the necessary environment variables, for example, `SNYK_TOKEN`, `SNYK_CFG_ORG`, and proxy settings, are correctly propagated to the Snyk MCP server process.

## Proxy configuration

If you are behind a corporate proxy, ensure that the `http_proxy` and `https_proxy` environment variables are correctly set and accessible to the Snyk CLI and MCP server process.

## SSE Transport specifics

If you are using SSE  (if using `snyk mcp -t sse`):

* Check to see if the local firewall might be blocking incoming connections to the port used by the Snyk MCP SSE server.
* Remember that SSE transport supports running the MCP server locally only.

## Folder Trust

if you are experiencing issues related to folder trust, you can disable this feature by using the&#x20;

`--disable-trust` CLI flag.

## Verbose logging and debugging

Use these suggestions to improve and expand on your Snyk CLI debug output to troubleshoot MCP-related issues:

* For more detailed Snyk CLI logs, which are useful whether you are starting the `snyk mcp` server or performing direct test scans (see [Basic repository scanning](troubleshooting-for-the-snyk-mcp-server.md#basic-repository-scanning-crucial-diagnostic)), you can add verbosity parameters to your Snyk commands.\
  \
  These include using the `-d` or `--debug` flag for debug level output, for example:
  * `snyk mcp -t sse --experimental -d`
  * `snyk test -d`
  * `snyk code test -d`
* For even more granular, trace-level logging, you can use the `--log-level=trace` option or set the `SNYK_LOG_LEVEL=trace` environment variable:
  * `snyk mcp -t sse --experimental -d --log-level=trace`&#x20;
  * `SNYK_LOG_LEVEL=trace snyk mcp -t sse --experimental -d`
* Inspect the MCP client and host logs from your AI tool, IDE, or MCP client application. These logs might contain errors related to connecting to or communicating with the Snyk MCP server.
