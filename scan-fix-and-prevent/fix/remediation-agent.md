---
description: >-
  How the Remediation Agent generates, applies, and verifies vulnerability fixes
  in your Snyk Projects
---

# Remediation Agent

{% hint style="info" %}
The coding-assistant and Snyk CLI entry points are available for all Snyk plans and are in [Early Access](https://docs.snyk.io/discover-snyk/getting-started/snyk-release-process#early-access-features). To turn them on, visit [Snyk Preview](https://docs.snyk.io/platform-administration/snyk-platform-administration/snyk-preview) or ask your Snyk account team.

Remediation campaigns require Evo by Snyk. Snyk enables campaigns for each Tenant. To request access, contact your Snyk account team.
{% endhint %}

The Remediation Agent automates vulnerability remediation: it scans a Project, creates a fix plan enriched with Snyk security intelligence, applies changes, and verifies the result. You can run it in a coding assistant, from the Snyk CLI, or through remediation campaigns in Evo by Snyk.

## How the products relate

Use this comparison to choose the tools for your workflow:

| Product or feature | Role | When to use it |
|---|---|---|
| Snyk Open Source | Finds vulnerabilities in open-source dependencies through software composition analysis (SCA). | Identify dependency vulnerabilities for the agent to fix. |
| Snyk Code | Finds vulnerabilities in your source code through static application security testing (SAST). | Identify source code vulnerabilities for the agent to fix. |
| Snyk Container | Finds vulnerabilities in container images. | Identify container vulnerabilities for the CLI remediation flow. |
| Remediation Agent | Coordinates scanning, fix planning, changes, and verification. | Apply and verify fixes for vulnerabilities found by Snyk. Supported scan types depend on the entry point. |
| [Snyk Agent Fix](../scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md) | Generates fix suggestions for vulnerabilities detected by Snyk Code. The CLI SAST remediation flow uses these suggestions. | Generate a source code fix, or use the Remediation Agent to apply and verify it. |
| [Snyk Studio](https://docs.snyk.io/agent-security/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio) | Connects Snyk security tools to coding assistants through the MCP server and directives, including hooks and remediation skills. | Check code during AI-assisted development and run remediation from your coding assistant. |
| [Remediation campaigns in Evo by Snyk](remediation-agent/remediation-campaigns.md) | Runs the Remediation Agent in Snyk-hosted infrastructure across selected repositories, once or on a schedule. | Manage open-source remediation across repositories and review the pull requests the agent opens. |

Snyk Studio provides the coding-assistant integration. Evo by Snyk provides the hosted campaign experience. The Remediation Agent performs the remediation workflow. Snyk Agent Fix supplies source code fix suggestions for the CLI SAST flow.

For a comparison of where each workflow runs, model configuration, and review, visit [Entry points](#entry-points).

{% hint style="warning" %}
The Remediation Agent is under rapid development. Some options on this page are available only in the latest preview release of the Snyk CLI. Run `snyk version` to check which version you have, and expect option names and defaults to change between releases.

The bring-your-own-model provider integrations are the newest part of the feature. Treat their configuration as unstable.
{% endhint %}

## Prerequisites

### Agentic IDE

* A Snyk account with Snyk Open Source or Snyk Code enabled.
* A supported coding assistant: Cursor, Claude, Gemini, Kiro, Codex, Windsurf, or GitHub Copilot.
* The `/snyk-fix` skill installed in the coding assistant, through the Snyk Studio installer or manually.

### Snyk CLI

* A Snyk account with Snyk Open Source, Snyk Code, or Snyk Container enabled.
* The Snyk CLI. Visit [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli).
* Access to a model provider. Visit [Configure a model provider](#configure-a-model-provider).

## How it works

Traditional vulnerability remediation requires a developer to review each finding, decide on a fix, apply it, and verify the outcome. The Remediation Agent automates that cycle.

The agent follows the same core flow at every entry point:

1. Scan: Snyk scans the Project for vulnerabilities with Snyk Open Source (SCA), Snyk Code (SAST), or Snyk Container.
2. Plan: The agent builds a fix plan enriched with Snyk security intelligence, including breakability signals for SCA fixes.
3. Fix: The agent applies the fix. For SCA, it bumps dependency versions or adds overrides. For SAST, it applies Snyk Agent Fix suggestions to the source code. For containers, it edits the Dockerfile.
4. Verify: The agent rescans to confirm the vulnerability is resolved, and runs the tests that exist in the application, such as unit tests. When a fix introduces a problem, the agent reports the outcome.

## Entry points

Choose an entry point based on where you want to run fixes and who supplies the model.

| Entry point | How to start | Supported scan types | Where fixes run | Model configuration | Review |
|---|---|---|---|---|---|
| [Coding assistant with Snyk Studio](#set-up-in-an-agentic-ide) | Run `/snyk-fix` to fix the top vulnerability, or `/snyk-batch-fix` to fix several in one pass. | Snyk Open Source and Snyk Code | Your coding assistant's development environment | Runs as a skill in your coding assistant. No separate model-provider configuration is required for the skill. | Review the changes before merging. |
| [Snyk CLI](#fix-vulnerabilities-from-the-snyk-cli) | Run `snyk fix --agentic` for an interactive remediation flow without an IDE. | Snyk Open Source, Snyk Code, and Snyk Container | The environment where you run the CLI | You configure a [model provider](#configure-a-model-provider). | Prompts for each fix unless you use `--auto-approve`. Review the changes before merging. |
| [Remediation campaigns](remediation-agent/remediation-campaigns.md) | In Evo by Snyk, create a campaign to run once or on a schedule. | Snyk Open Source | Snyk-hosted infrastructure | Snyk supplies the model. | The agent opens pull requests for you to review and merge. |

Campaigns require Evo by Snyk with remediation campaigns enabled for your Tenant. The IDE and CLI prerequisites on this page apply to their respective entry points. For campaign setup and limitations, visit [Remediation campaigns](remediation-agent/remediation-campaigns.md).

`/snyk-batch-fix` is an ADE skill rather than a CLI command. To fix several vulnerabilities from the CLI, use `--auto-approve`. Review its additional effects in [Scope the run](#scope-the-run).

## Breaking change assessment

Before the agent changes an open-source dependency, Snyk assesses whether the version upgrade is likely to break your build. The agent receives that assessment along with what to watch for if the upgrade does break something. Breakability applies to open source dependencies only, so SAST and container fixes carry no breakability rating.

How much the agent does on its own depends on where you run it. In an agentic IDE, the agent proceeds automatically when there is no breakability risk. In the CLI, the agent prompts you to choose what to fix unless you run it with `--auto-approve`. Visit [Scope the run](#scope-the-run) for the options that control this.

{% hint style="info" %}
The breaking change assessment is in preview. When it is unavailable, the agent falls back to a local heuristic assessment.
{% endhint %}

## Set up in an agentic IDE

{% stepper %}
{% step %}
### Install with the Snyk Studio installer

The [Snyk Studio installer](https://github.com/snyk/studio-recipes/tree/main/installer#install) installs and configures everything you need: the Snyk CLI, the Snyk MCP server, the `/snyk-fix` and `/snyk-batch-fix` skills, and the secure at inception hooks.

#### macOS and Linux

```bash
curl -fsSL 'https://raw.githubusercontent.com/snyk/studio-recipes/main/installer/dist/snyk-studio-install.sh' -o snyk-studio-install.sh
bash ./snyk-studio-install.sh
```

#### Windows (PowerShell)

```powershell
powershell -Command "Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/snyk/studio-recipes/main/installer/dist/snyk-studio-install.ps1' -OutFile snyk-studio-install.ps1"
powershell -ExecutionPolicy Bypass -File .\snyk-studio-install.ps1
```

By default, the installer applies to every coding assistant it detects. To target one assistant, pass `--ade [agent]`:

```bash
bash ./snyk-studio-install.sh --ade claude
```

Accepted values: `cursor`, `claude`, `gemini`, `kiro`, `codex`, `windsurf`, `copilot-cli`, `copilot-vscode`.
{% endstep %}

{% step %}
### Turn on the experimental MCP profile

The breaking change assessment requires the Snyk MCP server to run in experimental mode. After you run the installer, add `SNYK_MCP_PROFILE: experimental` to your MCP server configuration.

{% tabs %}
{% tab title="Cursor" %}
Navigate to **Tools & MCPs**, find the **Snyk** MCP entry, and click the edit icon. Add `"SNYK_MCP_PROFILE": "experimental"` to the `env` block.
{% endtab %}

{% tab title="Claude" %}
Open `~/.claude.json`, find the `mcpServers.Snyk` entry, and add `"SNYK_MCP_PROFILE": "experimental"` to the `env` block:

```json
"Snyk": {
  "args": ["mcp", "-t", "stdio"],
  "command": "/usr/local/bin/snyk",
  "env": {
    "SNYK_API": "https://api.snyk.io",
    "SNYK_MCP_PROFILE": "experimental"
  }
}
```
{% endtab %}

{% tab title="Codex" %}
Add the profile to `~/.codex/config.toml`:

```toml
[mcp_servers.snyk-security]
command = "npx"
args = ["-y", "snyk@latest", "mcp", "-t", "stdio"]

[mcp_servers.snyk-security.env]
SNYK_MCP_PROFILE = "experimental"
```
{% endtab %}
{% endtabs %}
{% endstep %}

{% step %}
### Restart your coding assistant

Restart your coding assistant so the updated MCP configuration takes effect. After the restart, Snyk appears as a connected MCP server. Type `/snyk-fix` in the prompt to run the Remediation Agent.
{% endstep %}
{% endstepper %}

## Fix vulnerabilities from the Snyk CLI

The Snyk Studio installer installs the CLI for you. If you skipped the installer, visit [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli) for the other installation options.

### Configure a model provider

The agentic CLI flow runs against a model you supply. Pick a provider with `--provider` and give it credentials. How each provider authenticates differs, and only three of them take an API key:

| `--provider` | Authentication | Required configuration | Optional configuration |
|---|---|---|---|
| `anthropic` (default) | API key | `ANTHROPIC_API_KEY` | `ANTHROPIC_BASE_URL` to route through a gateway |
| `openai` | API key | `OPENAI_API_KEY` | None |
| `bedrock` | The standard AWS credential chain, so no API key | AWS credentials the SDK can resolve, such as a profile or an instance role | `AWS_REGION` to pin the region |
| `vertex` | Google Application Default Credentials, so no API key. Authenticate with `gcloud auth application-default login` | `GOOGLE_CLOUD_PROJECT` and `GOOGLE_CLOUD_LOCATION` | `VERTEX_AUTH_TOKEN` for a gateway that expects a bearer token, and `VERTEX_BASE_URL` |
| `litellm` | A LiteLLM virtual key. The proxy holds the real provider credentials | `LITELLM_BASE_URL`, which must use HTTPS | `LITELLM_API_KEY` |
| `ollama` | None, because the model runs locally | `--model`, which Ollama requires | `OLLAMA_HOST`, which defaults to `http://localhost:11434` |

Each provider has a default model, and `--model` overrides it. For Amazon Bedrock, model IDs depend on your account and region, and newer Claude models are often reachable only through a cross-region inference profile. Pass `--model` when the default is not enabled for your account.

Snyk identifies its Bedrock traffic with the application ID `SNYK_AGENTIC_FIX`, so you can attribute or authorize it downstream.

{% hint style="info" %}
To add custom headers to every provider request, for example an end-user identity header required by a gateway, set `REMY_EXTRA_HEADERS` to a comma-separated list of `key=value` pairs.
{% endhint %}

### Run the command

Pass exactly one of `--sca` for dependency vulnerabilities, `--sast` for source code vulnerabilities, or `--container` for container vulnerabilities. The `--experimental` flag is required alongside `--agentic`:

```bash
snyk fix --agentic --experimental --sca [path]
snyk fix --agentic --experimental --sast [path]
snyk fix --agentic --experimental --container [path]
```

### Scope the run

An interactive run prompts you for each fix. These options change that. For what each option does and how they combine, visit [Fix](https://docs.snyk.io/developer-tools/snyk-cli/commands/fix).

* Run without prompting: `--auto-approve`. It has side effects, described in the warning that follows this list.
* Choose which vulnerabilities to fix: `--issue-ids`, `--exclude-ids`, `--severity-threshold`, `--severity-filter`, and `--breakability-filter`.
* Preview or adjust the run: `--dry-run` and `--no-breakability`.

The `--issue-ids`, `--exclude-ids`, and `--severity-threshold` options require `--auto-approve`.

{% hint style="warning" %}
`--auto-approve` does more than approve fixes. It also trusts the folder, continues after test failures, skips advisories, and approves partial plans when the agent reaches a budget warning. Review every change before you merge it.
{% endhint %}

## Fix container vulnerabilities

`--container` runs the container remediation flow against a single Dockerfile. It addresses two kinds of finding:

* Base image upgrades. The agent updates the base image reference in the final stage's `FROM` instruction. When that reference is built from global `ARG` defaults, the agent edits the `ARG` default values rather than the structure of the `FROM` lines. Verification confirms that the upgraded image keeps a runtime user equivalent to the original, so a base image bump does not silently change the user your container runs as.
* OS package upgrades. The agent patches operating system packages in the image. The run presents the available package upgrades and you select which ones to apply.

Verification builds the image and rescans it, so the agent rejects a candidate Dockerfile that no longer builds instead of reporting it as a fix.

{% hint style="info" %}
The container flow selects fixes at the Dockerfile and package level rather than per vulnerability, so it does not accept the options that select individual vulnerabilities. Passing any of them ends the run before the scan starts. The [Fix](https://docs.snyk.io/developer-tools/snyk-cli/commands/fix) reference lists them. `--container` also requires `--agentic`.
{% endhint %}

## Best practices

* Run the agent against a test repository or a small Project first, to get familiar with the output and the review process.
* Start with SCA fixes rated low breakability. These are the least likely to disrupt your build, and they give you the clearest signal on how the agent performs in your environment.
* Review every fix before you merge it. An interactive run shows you the plan, and an auto-approved run does not.

## Get help

Send feedback on outcomes to Snyk through your account manager or through [Snyk Support](https://support.snyk.io). Your input shapes the feature while it is in Early Access.
