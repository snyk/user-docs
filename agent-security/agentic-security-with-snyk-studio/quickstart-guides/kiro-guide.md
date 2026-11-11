---
description: How to set up Snyk Studio in Kiro with the Snyk CLI
nav_context: classic
last_reviewed: 2026-09-17
last_reviewed_source: bulk-sync
---

# Connect Kiro to Snyk

Set up Snyk Studio in Kiro to invoke Snyk scanning tools, evaluate package health, and use remediation workflows.

This integration includes MCP access and remediation and package-health components. The Studio installer does not add Secure at inception hooks for Kiro. For the differences between assistants, visit the [capability matrix](../getting-started-with-snyk-studio.md#coding-assistant-support).

For enterprise deployment, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).

## Install Snyk Studio

Install Kiro and the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli). Enable the Snyk products you want to scan with. Complete the account and product setup in [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md).

1. Authenticate with Snyk:

   ```bash
   snyk auth
   ```

2. Run the installer as the developer who uses Kiro:

   ```bash
   snyk studio install --experimental --ade=kiro
   ```

3. Start a new session in Kiro and open the repository you want to work in. Complete any MCP authentication or folder-trust request.

## Verify installation

Run:

```bash
snyk studio install --experimental --verify --ade=kiro
```

This checks the installed assets. For merged configuration files, the verifier checks only file existence. It does not confirm that a scan completed.

Ask Kiro to scan the repository:

```text
Scan this repository for code and dependency vulnerabilities. Summarize the findings before making changes.
```

## Invoke scans and remediation

Ask the assistant to invoke Snyk tools when you want a scan. If you add rules about when to scan, they rely on the model following those instructions and do not provide programmatic hook execution.

Use the installed remediation and package-health workflows for existing vulnerabilities and package selection. Visit [Directives](../directives.md) for their behavior and limitations.

## Update or remove Snyk Studio

After updating the Snyk CLI, run the installation command again to refresh the components, then start a new assistant session.

To remove the local setup for Kiro, run:

```bash
snyk studio install --experimental --uninstall --ade=kiro
```

If your company deployed Studio through ADS, use the [ADS removal procedure](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#uninstall).

## Manual MCP configuration

For an MCP-only connection, use the [manual configuration instructions](../getting-started-with-snyk-studio.md#configure-the-snyk-mcp-server) with your client's supported configuration format. Manual MCP setup does not install the full Studio components.
