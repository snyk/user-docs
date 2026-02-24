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

## Configure the Snyk MCP profile

Snyk supports three tool profiles that let you control which tools are available to your AI agent. Profiles help you balance token cost and capability based on your needs.

The available profiles are:

|  Profile type  | Description                                                                                                                  |
| :------------: | ---------------------------------------------------------------------------------------------------------------------------- |
|      Lite      | The minimum set of essential security scanning tools. Ideal for keeping token use and context size low.                      |
| Full (default) | All stable tools. Provides comprehensive security coverage across code, dependencies, containers, IaC, and SBOMs.            |
|  Experimental  | Everything in the full profile, plus new tools under evaluation. These may change or be promoted to full in future releases. |

These tools are available for the following profile types:

<table><thead><tr><th width="269.86328125" align="center">Tool</th><th data-type="checkbox">Lite profile</th><th data-type="checkbox">Full profile</th><th data-type="checkbox">Experimental profile</th></tr></thead><tbody><tr><td align="center"><code>snyk_auth</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_code_scan</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_sca_scan</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_version</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_logout</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_trust</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_send_feedback</code></td><td>true</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_container_scan</code></td><td>false</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_iac_scan</code></td><td>false</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_sbom_scan</code></td><td>false</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_aibom</code></td><td>false</td><td>true</td><td>true</td></tr><tr><td align="center"><code>snyk_package_health_check</code></td><td>false</td><td>false</td><td>true</td></tr></tbody></table>

To set a profile, use either the `--profile` CLI flag or the `SNYK_MCP_PROFILE` environment variable CLI flag:

```
npx -y snyk@latest mcp -t stdio --profile=lite
```

{% code title="mcp.json" %}
```
{
  "mcpServers": {
    "Snyk": {
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {
        "SNYK_MCP_PROFILE": "experimental"
      }
    }
  }
}

```
{% endcode %}

If a profile is not specified in setup, the profile defaults to full.

## Configure Secure at inception directives

The Secure at inception directives are used to guide agents to test for security issues at the time of code generation. To learn how to configure these directives with code examples, visit [Secure at inception directives](directives.md#secure-at-inception-directives).

## Configure Remediation directives

Remediation directives are commands that can be used to accelerate the remediation of pre-existing security issues. To learn how to configure these directives with code examples, visit [Remediation directives](directives.md#remediation-directives).

{% hint style="info" %}
Snyk maintains quickstart guides for some of the more common coding assistants. For detailed configuration steps of a specific coding assistant, visit [Quickstart guides](quickstart-guides-for-snyk-studio/).
{% endhint %}
