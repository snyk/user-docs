---
description: Set up Snyk Studio in your coding assistant to scan code and dependencies, check package health, and fix vulnerabilities
last_reviewed: 2026-07-21
last_reviewed_source: bulk-sync
---
# Getting started with Snyk Studio

Set up Snyk Studio to scan code and dependencies from your AI coding assistant, evaluate packages, and fix vulnerabilities. For local setup, you authenticate, install Studio, start a new assistant session, and verify a scan.

Choose the setup path that matches your deployment:

| Deployment | Setup path |
| --- | --- |
| Enterprise deployment across your development teams | Use the [Agentic Development Security (ADS) installer](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md). |
| Local setup on your own machine | [Install with the Snyk CLI](#install-on-your-machine). |
| Manual integration with another local MCP client | [Configure the Snyk MCP server](#configure-the-snyk-mcp-server). |

{% hint style="info" %}
For local setup, install [Snyk CLI version 1.1307.4 or later](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli). Studio commands were introduced in [CLI 1.1307.4](https://github.com/snyk/cli/releases/tag/v1.1307.4). Install your coding assistant and create a [Snyk account](https://app.snyk.io/). Enable the Snyk products you want to scan with, including [Snyk Code](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/configure-snyk-code) for first-party code.

If your administrator manages Studio on your machine, follow your company's deployment instructions.
{% endhint %}

## Coding assistant support

The default CLI installation configures the following components for each assistant:

| Coding assistant | `--ade` value | Secure at inception hooks | MCP server | Remediation and package health |
| --- | --- | --- | --- | --- |
| [Claude Code](quickstart-guides/claude-code-guide.md) | `claude` | Yes | Yes | Yes |
| [Cursor](quickstart-guides/cursor-guide.md) | `cursor` | Yes | Yes | Yes |
| [Codex](quickstart-guides/codex-cli-guide.md) | `codex` | Yes | Yes | Yes |
| [Gemini CLI](quickstart-guides/gemini-cli-guide.md) (access restrictions apply) | `gemini` | Yes | Yes | Yes |
| [GitHub Copilot CLI](quickstart-guides/github-copilot-guide.md) | `copilot-cli` | Yes | Yes | Yes |
| [GitHub Copilot in VS Code](quickstart-guides/github-copilot-guide.md) | `copilot-vscode` | Yes | Yes | Yes |
| [Kiro](quickstart-guides/kiro-guide.md) | `kiro` | No | Yes | Yes |

Kiro receives MCP configuration and remediation and package-health components without Secure at inception hooks. Other clients use [manual MCP integration](#configure-the-snyk-mcp-server).

For agent-invoked Snyk tools in Devin Desktop, follow [Connect Devin Desktop to Snyk](quickstart-guides/windsurf-guide.md). This manual setup does not install Studio hooks, remediation commands, or package-health skills. The installer still accepts the legacy `windsurf` target, which writes to the old Cascade configuration paths and does not configure Devin Local.

To import Snyk findings into OpenAI Daybreak Defense Factory, visit [Connect Daybreak Defense Factory to Snyk](quickstart-guides/daybreak-defense-factory-guide.md). The Studio installer does not configure that import.

{% hint style="warning" %}
Google has ended Gemini CLI and Gemini Code Assist access for some account types and is retiring it for others. Visit [Migrate from Gemini tooling](quickstart-guides/antigravity-guide.md#migrate-from-gemini-tooling) for the schedule and Antigravity setup. The Studio installer's `gemini` target configures Gemini CLI, not Antigravity CLI.
{% endhint %}

## What the installer configures

| Component | Purpose |
| --- | --- |
| Secure at inception hooks | Trigger background code and dependency scanning in assistants with hook support. |
| Local Snyk MCP server | Gives the assistant access to Snyk tools that it can invoke directly. |
| Remediation commands and skills | Provide workflows to scan, fix, and validate existing vulnerabilities. The form depends on the assistant. |
| Secure dependency health check | Guides package selection using security, maintenance, popularity, and community data. |

Visit [Directives](directives.md) for the differences between these workflows.

{% hint style="warning" %}
The hooks and MCP configuration use the absolute path of the Snyk CLI binary that ran the installation. If you move the binary or reinstall it at a different path, run `snyk studio install --experimental` again from the new binary. Otherwise, the installed configuration points to the old path.
{% endhint %}

<a id="install-snyk-studio-for-ades-with-hooks-support"></a>
<a id="install-commands"></a>

## Install on your machine

Run these commands as the developer who uses the coding assistant.

{% stepper %}
{% step %}
**Authenticate the Snyk CLI**

Run:

```bash
snyk auth
```

Complete authentication in your browser. If you belong to multiple Snyk Organizations, [configure the Organization for your scans](troubleshooting.md#snyk-organization-configuration).
{% endstep %}

{% step %}
**Install Studio**

Run:

```bash
snyk studio install --experimental
```

The installer detects assistants by looking for their configuration directories in your home directory. To select assistants explicitly, use `--ade`. For example:

```bash
snyk studio install --experimental --ade=claude,cursor
```

To preview the files and configuration changes before installation, run:

```bash
snyk studio install --experimental --dry-run
```

The installer merges Snyk entries into existing configuration files and installs the components supported by each assistant. Studio installation and log commands require the Studio installer flag `--experimental`. This required flag is separate from the installer recipe profile and the MCP tool profile.
{% endstep %}

{% step %}
**Start a new assistant session**

Start a new session so your assistant loads the configuration. Open the repository you want to work in.

If the MCP client requests authentication or trust for the repository directory, complete those steps. In Codex, approve the installed hooks when prompted or review them with `/hooks`. Non-interactive Codex sessions skip unapproved hooks.
{% endstep %}

{% step %}
**Verify installation and scan a repository**

Check the installed files:

```bash
snyk studio install --experimental --verify
```

The verifier reports **ok** for copied files that match the embedded assets, **DIFF** for changed files, and **MISS** for missing files. For merged configuration files, it checks file existence. This check does not confirm that a scan ran successfully.

Ask your assistant to scan the repository:

```text
Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
```

For assistants with hooks, continue with a coding task and inspect hook activity using:

```bash
snyk studio logs --experimental --level=debug
```

Use the logs to distinguish background hook activity from a scan you requested through MCP. For diagnosis, visit [Troubleshooting](troubleshooting.md).
{% endstep %}
{% endstepper %}

<a id="install-snyk-studio-for-ades-without-hooks-support"></a>

## Configure the Snyk MCP server

Use manual MCP configuration for a client outside the CLI installer support matrix, or when you need only agent-invoked Snyk tools.

{% hint style="info" %}
The Snyk MCP server runs locally through the Snyk CLI and needs access to your local files. Snyk Studio does not provide a hosted remote scanning MCP server.

Adding the MCP server alone does not install Studio hooks, commands, or skills.
{% endhint %}

Install and authenticate the Snyk CLI, then add a local server to your client's MCP configuration. For a client that uses `mcpServers`, the configuration is:

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

Replace the `command` value with the absolute path to your Snyk CLI executable. Clients use different configuration formats. Visit the [assistant-specific guides](quickstart-guides/README.md) for client instructions.

After the client loads the server, ask the assistant to scan your repository. A successful MCP scan verifies that the assistant can invoke a Snyk tool. It does not verify automatic hook scanning.

## Configure the Snyk MCP profile

The MCP profile controls which tools the server exposes to the assistant:

| Profile | Purpose |
| --- | --- |
| `lite` | Exposes essential scanning and supporting tools with a smaller tool set. |
| `full` | Exposes the stable tool set. This is the default. |
| `experimental` | Adds tools under evaluation to the full tool set. |

Set the MCP tool profile with the server's `--profile` argument or the `SNYK_MCP_PROFILE` environment variable. This profile selects which MCP tools the assistant can access. It does not control Studio installation. For example:

```bash
snyk mcp -t stdio --profile=lite
```

The MCP tool profile is separate from the Studio installer's recipe profile. Changing the MCP profile does not install or remove hooks.

| Tool | Purpose | Lite | Full | Experimental |
| --- | --- | --- | --- | --- |
| `snyk_auth` | Authenticate with Snyk | Yes | Yes | Yes |
| `snyk_sca_scan` | Scan dependencies | Yes | Yes | Yes |
| `snyk_code_scan` | Scan first-party code | Yes | Yes | Yes |
| `snyk_container_scan` | Scan container images | No | Yes | Yes |
| `snyk_iac_scan` | Scan infrastructure as code | No | Yes | Yes |
| `snyk_sbom_scan` | Scan a software bill of materials | No | Yes | Yes |
| `snyk_secret_scan` | Scan for hard-coded secrets | No | Yes | Yes |
| `snyk_version` | Check version information | Yes | Yes | Yes |
| `snyk_logout` | Log out | Yes | Yes | Yes |
| `snyk_aibom` | Generate an AI-BOM | No | Yes | Yes |
| `snyk_trust` | Trust a repository directory | Yes | Yes | Yes |
| `snyk_send_feedback` | Report remediation feedback | Yes | Yes | Yes |
| `snyk_package_health_check` | Evaluate package health | No | Yes | Yes |
| `snyk_breakability_check` | Assess a package upgrade for breaking changes | No | Yes | Yes |

Inspect your MCP client's tool list to check what your installed CLI exposes. A tool's profile membership does not mean that Studio hooks run it automatically.

## Update or remove a local installation

Update the Snyk CLI using your [CLI installation method](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli), then run `snyk studio install --experimental` again to refresh the installed components. Start a new assistant session after updating.

To remove a local installation:

```bash
snyk studio install --experimental --uninstall
```

To preview removal first:

```bash
snyk studio install --experimental --uninstall --dry-run
```

Add `--ade` to target specific assistants. Removal deletes files the installer copied and removes Snyk entries from configuration files it merged into. It preserves other configuration entries.

For an enterprise deployment, use the [ADS removal procedure](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#uninstall).


<a id="configure-secure-at-inception-directives"></a>
<a id="configure-remediation-directives"></a>

## Use Studio in your workflow

* For automatic code and dependency scanning, visit [Secure at inception directives](directives.md#secure-at-inception-directives).
* To evaluate a dependency before adding it, visit [Package health checks](directives.md#package-health-checks).
* To fix existing vulnerabilities, visit [Remediation directives](directives.md#remediation-directives).
* For enterprise rollout, visit [Distribution at scale](distribution-at-scale.md).
