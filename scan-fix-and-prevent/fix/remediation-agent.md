# Remediation Agent

The Remediation Agent is an LLM-driven tool that automates vulnerability remediation in your projects. It scans your Project with Snyk, generates a fix plan enriched with Snyk security intelligence, applies the fixes, and verifies the result, all in your coding assistant or the Snyk CLI.

## How it works

Traditional vulnerability remediation requires a developer to review each finding, decide on a fix, apply it manually, and verify the outcome. The Remediation Agent automates this cycle. Its goal is to move from human-in-the-loop validation toward fully autonomous, mergeable pull requests.

The agent follows the same core flow regardless of the entry point:

1. Scan: Snyk scans the Project for vulnerabilities using Snyk Open Source (SCA) or Snyk Code (SAST).
2. Plan: The agent generates a fix plan enriched with Snyk security intelligence, including breakability signals for SCA fixes.
3. Fix: The agent applies the fix. For SCA, this means bumping dependency versions or adding overrides. For SAST, this means applying Snyk Agent Fix suggestions to the source code.
4. Verify: The agent rescans to confirm the vulnerability is resolved and runs any tests that exist in the application, such as unit tests. If a fix introduces a problem, the agent reports the outcome.

## Entry points

* **Agentic IDE (ADE)**: Run `/snyk-fix` in your coding assistant to scan the project and apply a fix for the top vulnerability. Use `/snyk-batch-fix` to address multiple issues at once. The `/snyk-fix` skill must be installed in the ADE before use. It installs automatically with the Snyk Studio one-line installer or manually by following the [studio-recipes](https://github.com/snyk/studio-recipes/tree/main) setup instructions.
* **Snyk CLI**: Run `snyk fix --agentic` from your terminal to trigger a human-in-the-loop interactive remediation experience without an IDE.

## Prerequisites

### For Agentic IDE (ADE)

* A Snyk account with Snyk Open Source or Snyk Code enabled.
* A supported coding assistant: Cursor, Claude, Gemini, Kiro, Codex, Windsurf, or GitHub Copilot.
* The `/snyk-fix` skill installed in the coding assistant (through the Snyk Studio installer or manually).

### For Snyk CLI

* A Snyk account with Snyk Open Source or Snyk Code enabled.
* The Snyk CLI. Visit [Install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/).
* An LLM API key from one of the following providers: Anthropic, OpenAI, Vertex AI, LiteLLM, or Ollama.

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

By default, the installer applies to every coding assistant it detects. To target a specific assistant, pass `--ade [agent]`:

```bash
bash ./snyk-studio-install.sh --ade claude
```

Accepted values: `cursor`, `claude`, `gemini`, `kiro`, `codex`, `windsurf`, `copilot-cli`, `copilot-vscode`.
{% endstep %}

{% step %}
### Enable the experimental MCP profile



The breaking change assessment requires the Snyk MCP server to run in experimental mode. After running the installer, add `SNYK_MCP_PROFILE: experimental` to your MCP server configuration.

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

Restart your coding assistant for the updated MCP configuration to take effect. After restarting, Snyk appears as a connected MCP server. Type `/snyk-fix` in the prompt to run the Remediation Agent.
{% endstep %}
{% endstepper %}

## Use `snyk fix --agentic`

The Snyk Studio installer installs the CLI automatically. If you skipped the installer, visit [Install the Snyk CLI](../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) for installation options.

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

For the full command reference, including all flags, visit [Fix](../../developer-tools/snyk-cli/commands/fix.md).

## Snyk intelligence injected by the Remediation Agent

Before the agent modifies an open-source dependency, Snyk performs a breaking change assessment. The agent receives guidance on whether the version upgrade is likely to break the build and what to watch for if it does. In the ADE, the agent proceeds automatically when there is no breakability risk. In the CLI, the agent still prompts you to choose what to fix, unless you run it with `--auto-approve`.

{% hint style="info" %}
The breaking change assessment is in preview. When it is unavailable, the agent falls back to a local heuristic assessment.
{% endhint %}

1. Try the Remediation Agent on a test or smaller repository first to get familiar with the output and the review process.
2. Start with SCA fixes rated as low breakability. These are the least likely to disrupt your build and give you the best signal on how the agent performs in your environment.
3. Send feedback about outcomes directly to Snyk through your account manager or through [Snyk Support](https://support.snyk.io). The feature is in public preview, and your input helps shape its development.

## Best practices

Before running the Remediation Agent on production code, consider the following:

1. Try the Remediation Agent on a test or smaller repository first to get familiar with the output and the review process.
2. Start with SCA fixes rated as low breakability. These are the least likely to disrupt your build and give you the best signal on how the agent performs in your environment.
3. Send feedback about outcomes directly to Snyk through your account manager or through [Snyk Support](https://support.snyk.io). The feature is in public preview, and your input helps shape its development.

* **Agentic IDE (ADE)**: Run `/snyk-fix` in your coding assistant to scan the project and apply a fix for the top vulnerability. Use `/snyk-batch-fix` to address multiple issues at once. The `/snyk-fix` skill must be installed in the ADE before use. It installs automatically with the Snyk Studio one-line installer or manually by following the [studio-recipes](https://github.com/snyk/studio-recipes/tree/main) setup instructions.
* **Snyk CLI**: Run `snyk fix --agentic` from your terminal to trigger a human-in-the-loop interactive remediation experience without an IDE.
