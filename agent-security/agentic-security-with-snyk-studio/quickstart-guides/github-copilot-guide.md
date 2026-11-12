---
description: How to set up Snyk Studio in GitHub Copilot with the Snyk CLI
nav_context: classic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Connect GitHub Copilot to Snyk

Set up Snyk Studio in GitHub Copilot to scan code and dependencies, evaluate package health, and fix vulnerabilities from your coding assistant.

For enterprise deployment across your development teams, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md). The following procedure is for local developer setup.

## Install Snyk Studio with hooks support

Install GitHub Copilot and the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli). Enable the Snyk products you want to scan with. Complete the account and product setup in [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md).

1. Authenticate with Snyk:

   ```bash
   snyk auth
   ```

2. Run the installer as the developer who uses GitHub Copilot:

   ```bash
   snyk studio install --experimental --ade=copilot-cli,copilot-vscode
   ```

3. Start a new session in GitHub Copilot and open the repository you want to work in. Complete any MCP authentication or folder-trust request.

The example configures both Copilot CLI and Copilot in VS Code. Use `--ade=copilot-cli` or `--ade=copilot-vscode` to configure only one of them. This setup does not require the legacy Studio opt-in in the Snyk IDE extension.

## Verify installation

Run:

```bash
snyk studio install --experimental --verify --ade=copilot-cli,copilot-vscode
```

This checks the installed assets. For merged configuration files, the verifier checks only file existence. It does not confirm that a scan completed.

Ask GitHub Copilot to scan the repository:

```text
Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
```

## Check background hook scanning

Run a coding task in a new assistant session. Studio hooks trigger background scans for supported code and dependency changes. Inspect the activity with:

```bash
snyk studio logs --experimental --level=debug
```

Check for hook events and scan results. A successful direct MCP scan does not verify that the hooks ran. For scan failures, authentication, or missing activity, visit [Troubleshooting](../troubleshooting.md).

## Use remediation

Codex and GitHub Copilot CLI receive the Snyk Fix skill. Ask the assistant to use that skill for existing vulnerabilities. Other installed assistant components use their command or workflow mechanism. Visit [Remediation directives](../directives.md#remediation-directives) for the differences.

## Update or remove Snyk Studio

After updating the Snyk CLI, run the installation command again to refresh the components, then start a new assistant session.

To remove the local setup for GitHub Copilot, run:

```bash
snyk studio install --experimental --uninstall --ade=copilot-cli,copilot-vscode
```

If your company deployed Studio through ADS, use the [ADS removal procedure](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#uninstall).

## Manual MCP configuration

For an MCP-only connection, use the [manual configuration instructions](../getting-started-with-snyk-studio.md#configure-the-snyk-mcp-server) with your client's supported configuration format. Manual MCP setup does not install the full Studio components.
