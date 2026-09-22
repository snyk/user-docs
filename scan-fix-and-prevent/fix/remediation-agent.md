---
description: >-
  How the Remediation Agent generates, applies, and verifies vulnerability fixes
  in your Snyk Projects
---

# Remediation Agent

{% hint style="info" %}
The Remediation Agent is in Open Preview. To turn it on, visit [Snyk Preview](../../platform-administration/snyk-platform-administration/snyk-preview.md) or ask your Snyk account team.
{% endhint %}

The Remediation Agent is an LLM-driven tool that automates vulnerability remediation in your Snyk Projects. It scans a Project, generates a fix plan enriched with Snyk security intelligence, applies the fix, and verifies the result. It runs in your coding assistant or in the Snyk CLI.

{% hint style="warning" %}
The Remediation Agent is under rapid development. Some options on this page are available only in the latest preview release of the Snyk CLI. Run `snyk version` to check which version you have, and expect option names and defaults to change between releases.

The bring-your-own-model provider integrations are the newest part of the feature. Treat their configuration as unstable.
{% endhint %}

## How it works

Traditional vulnerability remediation requires a developer to review each finding, decide on a fix, apply it, and verify the outcome. The Remediation Agent automates that cycle.

The agent follows the same core flow at every entry point:

1. Scan: Snyk scans the Project for vulnerabilities with Snyk Open Source (SCA), Snyk Code (SAST), or Snyk Container.
2. Plan: The agent builds a fix plan enriched with Snyk security intelligence, including breakability signals for SCA fixes.
3. Fix: The agent applies the fix. For SCA, it bumps dependency versions or adds overrides. For SAST, it applies Snyk Agent Fix suggestions to the source code. For containers, it edits the Dockerfile.
4. Verify: The agent rescans to confirm the vulnerability is resolved, and runs the tests that exist in the application, such as unit tests. When a fix introduces a problem, the agent reports the outcome.

## Breaking change assessment

Before the agent changes an open-source dependency, Snyk assesses whether the version upgrade is likely to break your build. The agent receives that assessment along with what to watch for if the upgrade does break something. Breakability applies to open source dependencies only, so SAST and container fixes carry no breakability rating.

How much the agent does on its own depends on where you run it. In an agentic IDE, the agent proceeds automatically when there is no breakability risk. In the CLI, the agent prompts you to choose what to fix unless you run it with `--auto-approve`.

{% hint style="info" %}
The breaking change assessment is in preview. When it is unavailable, the agent falls back to a local heuristic assessment.
{% endhint %}

## Entry points

* Agentic IDE (ADE): run `/snyk-fix` in your coding assistant to scan the Project and fix the top vulnerability. Run `/snyk-batch-fix` to address several vulnerabilities in one pass. Install the `/snyk-fix` skill before you use it. The Snyk Studio installer installs it, or you can follow the [studio-recipes](https://github.com/snyk/studio-recipes/tree/main) setup instructions.
* Snyk CLI: run `snyk fix --agentic` in your terminal for an interactive, human-in-the-loop remediation flow without an IDE.

`/snyk-batch-fix` is an ADE skill rather than a CLI command. To fix several vulnerabilities from the CLI, use `--auto-approve`.

## Prerequisites

### Agentic IDE

* A Snyk account with Snyk Open Source or Snyk Code enabled.
* A supported coding assistant: Cursor, Claude, Gemini, Kiro, Codex, Windsurf, or GitHub Copilot.
* The `/snyk-fix` skill installed in the coding assistant, through the Snyk Studio installer or manually.

### Snyk CLI

* A Snyk account with Snyk Open Source, Snyk Code, or Snyk Container enabled.
* The Snyk CLI. Visit [Install the Snyk CLI](../../developer-tools/snyk-cli/install-the-snyk-cli/README.md).
* Access to a model provider. Visit [Configure a model provider](#configure-a-model-provider).

## Set up the Remediation Agent

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

The Snyk Studio installer installs the CLI for you. If you skipped the installer, visit [Install the Snyk CLI](../../developer-tools/snyk-cli/install-the-snyk-cli/README.md) for the other installation options.

### Configure a model provider

The agentic CLI flow runs against a model you supply. Pick a provider with `--provider` and give it credentials. How each provider authenticates differs, and only three of them take an API key:

| `--provider` | Authentication | Required configuration | Optional configuration |
|---|---|---|---|
| `anthropic` (default) | API key | `ANTHROPIC_API_KEY` | `ANTHROPIC_BASE_URL` to route through a gateway |
| `openai` | API key | `OPENAI_API_KEY` | — |
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

An interactive run prompts you for each fix. Narrow what the agent touches with these options:

| Option | Effect |
|---|---|
| `--auto-approve` | Applies fixes without prompting. It also trusts the folder, continues after test failures, skips advisories, and approves partial plans when the agent reaches a budget warning. |
| `--issue-ids=<ID>[,<ID>]` | Fixes only these vulnerabilities. Requires `--auto-approve`. |
| `--exclude-ids=<ID>[,<ID>]` | Fixes everything except these. Requires `--auto-approve`, and does not combine with `--issue-ids`. |
| `--severity-filter=<SEVERITY>` | Fixes only vulnerabilities at exactly these severities, from `low`, `medium`, `high`, and `critical`. Unlike `--severity-threshold`, this is an exact match, and it applies to interactive runs as well as auto-approved ones. |
| `--breakability-filter=<RATING>` | Fixes only SCA fixes at exactly these breakability ratings, from `low`, `medium`, and `high`. SCA only, so it skips the SAST leg. |
| `--no-breakability` | Skips the Snyk Breakability API and uses the local heuristic only. |
| `--dry-run` | Shows the fix plan without changing any files. |

For the full command reference, including all flags, visit [Fix](../../developer-tools/snyk-cli/commands/fix.md).

## Fix container vulnerabilities

`--container` runs the container remediation flow against a single Dockerfile. It addresses two kinds of finding:

* Base image upgrades. The agent updates the base image reference in the final stage's `FROM` instruction. When that reference is built from global `ARG` defaults, the agent edits the `ARG` default values rather than the structure of the `FROM` lines. Verification confirms that the upgraded image keeps a runtime user equivalent to the original, so a base image bump does not silently change the user your container runs as.
* OS package upgrades. The agent patches operating system packages in the image. The run presents the available package upgrades and you select which ones to apply.

Verification builds the image and rescans it, so a candidate Dockerfile that no longer builds is rejected rather than reported as a fix.

{% hint style="info" %}
The container flow selects fixes at the Dockerfile and package level rather than per vulnerability, so it does not accept `--issue-ids`, `--exclude-ids`, `--severity-threshold`, or `--severity-filter`. Passing any of them ends the run before the scan starts. `--container` also requires `--agentic`.
{% endhint %}

## Best practices

* Run the agent against a test repository or a small Project first, to get familiar with the output and the review process.
* Start with SCA fixes rated low breakability. These are the least likely to disrupt your build, and they give you the clearest signal on how the agent performs in your environment.
* Review every fix before you merge it. An interactive run shows you the plan, and an auto-approved run does not.

## Get help

Send feedback on outcomes to Snyk through your account manager or through [Snyk Support](https://support.snyk.io). Your input shapes the feature while it is in Open Preview.
