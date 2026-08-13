---
nav_context: classic
---

# Remediation Agent

The Remediation Agent is an LLM-driven tool that automates vulnerability remediation in your projects. It scans your project with Snyk, generates a fix plan enriched with Snyk's security intelligence, applies fixes, and verifies the result — all within your coding assistant or the Snyk CLI.

The Remediation Agent uses Snyk's breakability signals to predict whether a fix will break your build before applying it. When a dependency upgrade carries risk, the agent provides specific guidance on what to watch for. When the upgrade is safe, the agent proceeds with confidence.

## How the Remediation Agent works

The Remediation Agent has two entry points:

* **Agentic IDE (ADE)** — Run `/snyk-fix` in your coding assistant to scan your project and apply a fix for the top vulnerability. Use `/snyk-batch-fix` to address multiple issues at once.
* **Snyk CLI** — Run `snyk fix --agentic` from your terminal to trigger the same remediation flow without an IDE.

Both entry points follow the same flow: scan, enrich with Snyk intelligence, fix, and verify.

## Snyk intelligence injected by the Remediation Agent

Before the agent modifies a dependency, Snyk performs a breaking change assessment. The agent receives guidance on whether the version upgrade is likely to break the build and what to watch for if it does. When there is no breakability risk, the agent proceeds without additional friction.

{% hint style="info" %}
The breaking change assessment is currently in preview. When it is unavailable, the agent falls back to a local heuristic assessment.
{% endhint %}

## Prerequisites

* A Snyk account with Snyk Open Source or Snyk Code enabled.
* A supported coding assistant: Cursor, Claude, Gemini, Kiro, Codex, Windsurf, or GitHub Copilot.
* An LLM API key from one of the following providers: Anthropic, OpenAI, Vertex AI, LiteLLM, or Ollama.
* The Snyk CLI. See [Install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/).

## Set up the Remediation Agent

### Install with the Snyk Studio installer

The [Snyk Studio installer](https://github.com/snyk/studio-recipes/tree/main/installer#install) installs and configures everything you need: the Snyk CLI, the Snyk MCP server, the `/snyk-fix` and `/snyk-batch-fix` skills, and the secure at inception hooks.

**macOS and Linux**

```bash
curl -fsSL 'https://raw.githubusercontent.com/snyk/studio-recipes/main/installer/dist/snyk-studio-install.sh' -o snyk-studio-install.sh
bash ./snyk-studio-install.sh
```

**Windows (PowerShell)**

```powershell
powershell -Command "Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/snyk/studio-recipes/main/installer/dist/snyk-studio-install.ps1' -OutFile snyk-studio-install.ps1"
powershell -ExecutionPolicy Bypass -File .\snyk-studio-install.ps1
```

By default, the installer applies to every coding assistant it detects. To target a specific assistant, pass `--ade [agent]`:

```bash
bash ./snyk-studio-install.sh --ade claude
```

Accepted values: `cursor`, `claude`, `gemini`, `kiro`, `codex`, `windsurf`, `copilot-cli`, `copilot-vscode`.

### Enable the experimental MCP profile

The breaking change assessment requires the Snyk MCP server to run in experimental mode. After running the installer, add `SNYK_MCP_PROFILE: experimental` to your MCP server configuration.

**Cursor** — Navigate to **Tools & MCPs**, find the **Snyk** MCP entry, and click the edit icon. Add `"SNYK_MCP_PROFILE": "experimental"` to the `env` block.

**Claude** — Open `~/.claude.json`, find the `mcpServers.Snyk` entry, and add `"SNYK_MCP_PROFILE": "experimental"` to the `env` block:

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

**Codex** — Add the profile to `~/.codex/config.toml`:

```toml
[mcp_servers.snyk-security]
command = "npx"
args = ["-y", "snyk@latest", "mcp", "-t", "stdio"]

[mcp_servers.snyk-security.env]
SNYK_MCP_PROFILE = "experimental"
```

### Restart your coding assistant

Restart your coding assistant for the updated MCP configuration to take effect. After restarting, Snyk appears as a connected MCP server. Type `/snyk-fix` in the prompt to run the Remediation Agent.

## Use `snyk fix --agentic`

The Snyk Studio installer installs the CLI automatically. If you skipped the installer, see [Install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) for installation options.

The agentic CLI flow requires an LLM provider API key. Set one of the following environment variables before running the command:

* **Anthropic (default):** `export ANTHROPIC_API_KEY=sk-ant-...`
* **OpenAI:** `export OPENAI_API_KEY=sk-...`
* **Vertex AI:** authenticate with `gcloud auth application-default login` and set `GOOGLE_CLOUD_PROJECT` and `GOOGLE_CLOUD_LOCATION`
* **LiteLLM:** set `LITELLM_BASE_URL` and `LITELLM_API_KEY`
* **Ollama:** start Ollama with your model, then pass `--provider=ollama --model=<model>` to the command

Run the command with `--sca` for dependency vulnerabilities or `--sast` for source-code issues:

```bash
snyk fix --agentic --experimental --sca [path]
snyk fix --agentic --experimental --sast [path]
```

For the full command reference including all flags, see [Fix](../../developer-tools/snyk-cli/commands/fix.md).

## Next steps

* [Getting started with Snyk Studio](getting-started-with-snyk-studio.md)
* [Directives](directives.md)
* [Distribution at scale](distribution-at-scale.md)
