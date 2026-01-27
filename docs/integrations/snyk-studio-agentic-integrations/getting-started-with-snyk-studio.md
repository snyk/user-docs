# Getting started with Snyk Studio

{% hint style="info" %}
**Feature availability**

The Snyk MCP Server is designed as a local MCP server, running on your system using the Snyk CLI to ensure local file access.

Snyk does not offer a hosted remote MCP server.
{% endhint %}

## Prerequisites

Snyk recommends that you have:

* `npx`
* `Node.js` installed on your machine

If these are not available to you, the Snyk CLI can be manually introduced in the MCP server configuration. For instructions on how to do so, refer to your coding assistant's official documentation.

## Configure the Snyk MCP server

Navigate to the coding assistant MCP server configuration setup. This often involves modifying a file called `mcp.json` or similar. The following snippet can be used to configure the Snyk MCP server:

{% code overflow="wrap" expandable="true" %}
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
{% endcode %}

To validate the MCP server configuration, prompt your coding agent with natural language, for example, "scan my directory for security issues".

## Configure Secure at inception directives

The Secure at inception directives are used to guide agents to test for security issues at the time of code generation. To learn how to configure these directives with code examples, visit [Secure at inception directives](directives.md#secure-at-inception-directives).

## Configure Remediation directives

Remediation directives are commands that can be used to accelerate the remediation of pre-existing security issues. To learn how to configure these directives with code examples, visit [Remediation directives](directives.md#remediation-directives).

{% hint style="info" %}
Snyk maintains quickstart guides for some of the more common coding assistants. For detailed configuration steps of a specific coding assistant, visit [Quickstart guides](quickstart-guides-for-snyk-studio/).
{% endhint %}
