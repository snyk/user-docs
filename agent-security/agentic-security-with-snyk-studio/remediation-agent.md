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

If you used the Snyk Studio installer, the CLI is already installed. Otherwise, install it with npm or Homebrew:

```bash
npm install -g snyk@latest
```

```bash
brew install snyk-cli
```

For additional installation methods, see [Install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/).

### Set up an LLM provider

The agentic CLI flow requires one of the following LLM providers:

**Anthropic (default)**

```bash
export ANTHROPIC_API_KEY=sk-ant-...
```

**OpenAI**

```bash
export OPENAI_API_KEY=sk-...
```

**Vertex AI**

```bash
export GOOGLE_CLOUD_PROJECT=your_project
export GOOGLE_CLOUD_LOCATION=your_location
gcloud auth application-default login
```

**LiteLLM**

```bash
export LITELLM_BASE_URL=https://your-proxy
export LITELLM_API_KEY=your-virtual-key
```

**Ollama (local)**

Start Ollama with your model first, then pass `--provider` and `--model` to the `snyk fix` command.

### Run the command

**Snyk Open Source (dependency vulnerabilities)**

```bash
snyk fix --agentic --experimental --sca [path]
```

**Snyk Code (source-code issues)**

```bash
snyk fix --agentic --experimental --sast [path]
```

`[path]` is the project directory to scan. It defaults to the current directory.

You must pass exactly one of `--sca` or `--sast`. Running `--agentic` without one of these flags prints usage guidance and exits without scanning.

### CLI flags

| Flag | Description |
| ---- | ----------- |
| `--agentic` | Enable the LLM-driven fix flow. |
| `--experimental` | Required alongside `--agentic` (early-access acknowledgement). |
| `--sca` | Remediate Snyk Open Source (dependency) vulnerabilities. Mutually exclusive with `--sast`. |
| `--sast` | Remediate Snyk Code (source-code) issues. Mutually exclusive with `--sca`. |
| `--provider` | LLM provider: `anthropic` (default), `openai`, `vertex`, `litellm`, or `ollama`. |
| `--model` | Model ID — required for Ollama (for example, `llama3.1`), optional override for Anthropic and OpenAI. |
| `--dry-run` | Show the fix plan without applying changes. |
| `--auto-approve` | Automatically approve fixes for all discovered issues. Combine with `--issue-ids` or `--severity-threshold` to restrict scope. |
| `--issue-ids` | Comma-separated list of Snyk issue IDs to fix. |
| `--severity-threshold` | Fix issues at or above this severity: `CRITICAL`, `HIGH`, or `MEDIUM`. |
| `--no-breakability` | Skip the Snyk breakability assessment and use local heuristics only. |
| `-d` / `--debug` | Print debug information to stderr. |

## Next steps

* [Getting started with Snyk Studio](getting-started-with-snyk-studio.md)
* [Directives](directives.md)
* [Distribution at scale](distribution-at-scale.md)
