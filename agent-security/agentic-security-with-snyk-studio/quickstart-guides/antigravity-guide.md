---
description: >-
  How to set up Snyk Studio and the Snyk MCP server in Google Antigravity
nav_context: classic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Connect Antigravity to Snyk

Choose the procedure for the Antigravity product you use:

| Product | Setup | What it provides |
| --- | --- | --- |
| Antigravity CLI | [Antigravity CLI](#antigravity-cli) | Local Snyk MCP tools. Does not install Studio hooks, remediation commands, or skills. |
| Antigravity IDE | [Install the Snyk Security plugin](#install-the-snyk-security-plugin) | Plugin setup and model-directed scan rules. Rules depend on the model and do not install Studio hooks. |

{% hint style="info" %}
This setup provides MCP access. It does not install Studio Secure at inception hooks. Rules instruct the agent when to scan and depend on the model following those instructions.

For clients supported by the Studio CLI installer, use [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md). For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).
{% endhint %}

## Antigravity CLI

Install and authenticate the Snyk CLI, then add the local Snyk server from [Configure the Snyk MCP server](../getting-started-with-snyk-studio.md#configure-the-snyk-mcp-server) to one of these files:

* `~/.gemini/config/mcp_config.json` for all projects
* `.agents/mcp_config.json` for the current workspace

Preserve existing entries in `mcpServers`. Open `/mcp` in Antigravity CLI to inspect the server connection, then ask the assistant to scan the repository. This configures agent-invoked MCP access without installing Studio hooks, remediation commands, or skills. Visit [Google's MCP configuration documentation](https://antigravity.google/docs/mcp).

## Antigravity IDE

### Install the Snyk Security plugin

Visit the [Google Antigravity](https://antigravity.google/) website to download the correct version of the IDE.

{% hint style="warning" %}
The following procedure uses the Snyk Security plugin in the Antigravity IDE. For Antigravity CLI, use the MCP configuration procedure on this page.
{% endhint %}

1. Open the [Snyk Security plugin in Google Antigravity](antigravity:extension/snyk-security.snyk-vulnerability-scanner).
2. Click **Install**.

If asked to trust the publisher, click **Trust Publisher & Install**.

<figure><img src="../../.gitbook/assets/Screenshot 2025-11-24 at 12.08.02 PM.png" alt="Popup in Snyk Security plugin asking for trust verification" width="410"></figure>

### Enable Secure at inception

After installation, the plugin prompts you to opt in to [Secure at inception](https://docs.snyk.io/glossary#secure-at-inception-directive). This action configures rules that instruct the model to scan new AI-generated code. The rules do not install Studio hooks or guarantee scan execution. Additional options are available on the **Settings** page for the plugin.

<figure><img src="../../.gitbook/assets/Screenshot 2025-11-24 at 12.08.20 PM.png" alt="Modal prompting you to opt in to Secure at inception" width="404"></figure>

### Authenticate

{% hint style="info" %}
Authenticate with Snyk and complete the client's trust request for the repository directory when prompted.
{% endhint %}

After selecting the Secure at inception option, authenticate with Snyk. You can authenticate at two points in this process:

* Immediately after plugin install
* Before your first Snyk code scan

Snyk opens a browser window. Create a Snyk account or log in with an existing account.

If you are creating an account, select a registration method and agree to the terms. After authentication, return to your IDE.

{% hint style="info" %}
To use Snyk Studio, specifically first-party code scanning, enable [Snyk Code](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code). Snyk Code analyzes your code for vulnerabilities. It temporarily clones the repository or uploads code and caches it according to the Snyk [data retention policy](https://docs.snyk.io/snyk-data-and-governance/how-snyk-handles-your-data).

With the Snyk Free plan, Snyk Code offers unlimited scans for open source projects, and limited tests for first-party code. For more details, visit [Plans and Pricing](https://snyk.io/plans/).
{% endhint %}

For existing users, select the login method associated with your account. If you do not have access to Snyk Code, enable it before your first scan. You can also [enable it directly in Snyk's Settings](https://docs.snyk.io/implementation-guides/enterprise-implementation-guide/create-a-template-organization#connect-your-development-tools).

{% hint style="info" %}
If you are enabling Snyk Code for the first time, you must import or re-import existing projects to properly scan them.
{% endhint %}

### Run Snyk Studio

After authentication, ask the assistant to scan your repository. If you use rules to request scans during code generation, check the tool calls to confirm that the model invoked a scan. Restart the client if it has not loaded the MCP configuration.

{% hint style="info" %}
Free users are limited to a set number of scans. If you reach the allotment, Snyk recommends [contacting sales](https://snyk.io/contact-us/) to unlock additional thresholds.
{% endhint %}

### Examples

In the chat, you can tag specific files or functions and ask Google Antigravity to scan your code before deployment:

{% code title="prompt" overflow="wrap" %}
```text
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Google Antigravity indicates that this request is related to security vulnerability scanning and calls the Snyk MCP server for various scans.

Google Antigravity combines the code and dependency scan results with Snyk recommendations to suggest fixes for some or all of the findings.

## Secure at inception rules

Add rules to guide the assistant when invoking Snyk scans. Rules do not guarantee that the model runs a scan. For the generic rule text, visit [Rules for manual MCP integrations](../directives.md#rules-for-manual-mcp-integrations).

To manage these instructions outside the Snyk setting, add the generic rule text to your environment's global rules file. For project-level rules, add it to each project.

## Migrate from Gemini tooling

Google directs affected Gemini CLI and Gemini Code Assist users to Antigravity. Visit the [consumer access notice](https://developers.google.com/gemini-code-assist/docs/deprecations/code-assist-individuals) and [Standard and Enterprise subscription schedule](https://docs.cloud.google.com/gemini/docs/codeassist/sunset) for the restrictions that apply to your account.

Google ended Gemini CLI and Gemini Code Assist IDE extension access for Gemini Code Assist for individuals, Google AI Pro, and Google AI Ultra on June 18, 2026.

Google stopped selling new Gemini Code Assist Standard and Enterprise subscriptions on October 9, 2026. Existing subscriptions remain supported through their contract end. Additional licenses and annual renewals end after January 31, 2027. Monthly renewals end after December 31, 2027.

Gemini CLI also remains accessible with paid Gemini API or Gemini Enterprise Agent Platform API keys, as described in [Google's CLI transition announcement](https://github.com/google-gemini/gemini-cli/discussions/27274).

The Studio CLI installer's `--ade=gemini` target configures Gemini CLI. It does not configure Antigravity CLI. Antigravity CLI uses different MCP configuration files. Visit [Google's migration guide](https://antigravity.google/docs/cli/gcli-migration/) when moving existing configuration.
