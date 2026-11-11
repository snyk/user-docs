---
description: How to set up Snyk Studio and the Snyk MCP server in Continue
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Continue to Snyk

Connect Continue to the local Snyk MCP server to scan code and dependencies from your coding assistant.

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}


## Prerequisites

* Install the Continue extension
* [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli)

Add the Continue extension to your IDE. Visit the official [Continue installation page](https://docs.continue.dev/ide-extensions/install).

## Install the Snyk MCP server in Continue

Install the Snyk MCP server using the method that best suits your operating system and local development environment.

{% tabs %}
{% tab title="Pre-installed Snyk CLI" %}
Create or edit the MCP configuration file `~/.continue/mcpServers/new-mcp-server.yaml` .

```yaml
name: MCP Server
version: 0.0.1
schema: v1
mcpServers:
  - name: Snyk
    command: /absolute/path/to/snyk
    args:
      - "mcp"
      - "-t"
      - "stdio"      
    env: {}
```

If the `snyk` command is not available, add it by following the instructions on the [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli) page.
{% endtab %}

{% tab title="Node.js and npx" %}
Create or edit the MCP configuration file `~/.continue/mcpServers/new-mcp-server.yaml` .

```yaml
name: MCP Server
version: 0.0.1
schema: v1
mcpServers:
  - name: Snyk
    command: npx
    args:
      - "-y"
      - "snyk@latest"
      - "mcp"
      - "-t"
      - "stdio"       
    env: {}
```
{% endtab %}
{% endtabs %}

## Set up the Snyk MCP server <a href="#setting-up-the-snyk-mcp-server" id="setting-up-the-snyk-mcp-server"></a>

The MCP server attempts to start automatically. The following example shows a Snyk MCP server that was successfully configured and started.

<figure><img src="../../.gitbook/assets/image (15).png" alt="Snyk MCP server configured and started in Continue"><figcaption></figcaption></figure>

Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to work with the Snyk MCP server.

## Use Agent mode

Continue only connects MCP servers in Agent mode. Select **Agent** before asking Continue to run a Snyk scan. Visit the [Continue MCP documentation](https://docs.continue.dev/customize/deep-dives/mcp).

## Secure at inception rules <a href="#secure-at-inception-rules" id="secure-at-inception-rules"></a>

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

Create project rules for [Continue](https://docs.continue.dev/customize/deep-dives/rules) in the rules file at `.continue/rules/new-rule.md`. Start the file with the following front matter, then add the generic rule text:

```markdown
---
alwaysApply: true
---
```

## Examples

In the chat, you can tag specific files or functions and ask Continue to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Continue indicates that this request is related to security vulnerability scanning and calls the Snyk MCP server for various scans.

<figure><img src="../../.gitbook/assets/image (16).png" alt="Continue calling the Snyk MCP server for a vulnerability scan"><figcaption></figcaption></figure>

Continue combines the code and dependency scan results with Snyk recommendations to suggest fixes for some or all of the findings.

<figure><img src="../../.gitbook/assets/image (370).png" alt="Continue reporting combined code and dependency scan results"><figcaption></figcaption></figure>
