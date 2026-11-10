---
description: >-
  How to diagnose Snyk Studio and Snyk MCP server problems, including CLI
  version and authentication issues
nav_context: classic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Troubleshooting

Diagnose installation, background hook scans, and direct MCP tool calls separately. A connected MCP server does not confirm that Studio hooks are installed or that a background scan completed.

For an enterprise deployment, also use [Agentic Development Security (ADS) troubleshooting](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/troubleshooting.md).

| Symptom | Track | Go to |
| --- | --- | --- |
| The installer finds no assistant, or the installation fails | Installation | [No assistant detected](#no-assistant-detected), [Check Snyk CLI version compatibility](#check-snyk-cli-version-compatibility), [Verify Snyk CLI path and permissions](#verify-snyk-cli-path-and-permissions) |
| Installed files look wrong or are missing | Installation | [Verify a local Studio installation](#verify-a-local-studio-installation) |
| Background hook scans do not start or do not finish | Background hook scans | [Approve Codex hooks](#approve-codex-hooks), [Verbose logging and debugging](#verbose-logging-and-debugging), [Slow or incomplete scans](#slow-or-incomplete-scans) |
| The MCP server does not connect or cannot scan | Direct MCP tool calls | [Verify Snyk CLI path and permissions](#verify-snyk-cli-path-and-permissions), [Folder trust](#folder-trust), [SSE transport specifics](#sse-transport-specifics), [Environment variable propagation](#environment-variable-propagation) |
| Scans fail with authentication or Organization errors | All tracks | [Authentication issues](#authentication-issues), [Snyk Organization configuration](#snyk-organization-configuration) |
| Scans fail only behind a proxy | All tracks | [Proxy configuration](#proxy-configuration) |
| Scans fail in the CLI as well as in the assistant | All tracks | [Basic repository scanning](#basic-repository-scanning) |

## Check Snyk CLI version compatibility

Run:

```bash
snyk version
```

Studio commands require Snyk CLI version 1.1307.4 or later. Update the CLI using your [installation method](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli), then refresh a local Studio installation with `snyk studio install --experimental`. Start a new assistant session after updating.

The CLI version required for a Studio installation is separate from the version prerequisites on analytics reports and other Snyk features.

## Verify a local Studio installation

Run:

```bash
snyk studio install --experimental --verify
```

The verifier compares copied files with the assets embedded in your CLI. It reports changed or missing files. For merged configuration files, it checks only that the file exists, so a successful verification does not confirm that all configuration entries are correct or that a scan ran.

To inspect the intended changes before reinstalling, run:

```bash
snyk studio install --experimental --dry-run
```

For an ADS installation, use [Verify an installation](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/verify-an-installation.md).

## Approve Codex hooks

If Codex lists the Snyk hooks but background scans do not start, open `/hooks` in the Codex CLI, review the Snyk hooks, and trust them. Codex 0.129.0 and later skips non-managed hooks until you trust their current definitions. Review them again after an update changes a hook.

Start a new assistant session, run a coding task, and check for hook events and scan results in the [Studio logs](#verbose-logging-and-debugging). For the setup procedure, visit [Connect Codex to Snyk](quickstart-guides/codex-cli-guide.md#install-snyk-studio-with-hooks-support).

## Verbose logging and debugging

For local CLI-based Studio activity, read the persistent log:

```bash
snyk studio logs --experimental
snyk studio logs --experimental --level=debug
```

Use debug output to check whether hook events arrived, why a scan started or did not start, and whether it completed.

To follow new log entries:

```bash
snyk studio logs --experimental --level=debug --watch
```

To collect structured output:

```bash
snyk studio logs --experimental --format=jsonl
```

For direct CLI scans and the MCP server, add `-d` for debug output:

```bash
snyk test -d
snyk code test -d
snyk mcp -t stdio -d
```

Also inspect the coding assistant's MCP and hook logs. For an ADS deployment, collect the corresponding [ADS installation logs](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/troubleshooting.md).

## No assistant detected

The local installer detects assistants from configuration directories in the home directory of the account running the command. Run the installer as the developer who uses the assistant.

To select an assistant explicitly, use its [supported `--ade` value](getting-started-with-snyk-studio.md#coding-assistant-support). For example:

```bash
snyk studio install --experimental --ade=cursor
```

If the machine has a managed Studio installation, follow your administrator's setup instructions. A user installation and a managed installation use different scopes.

## Verify Snyk CLI path and permissions

The installed hooks and MCP configuration refer to an absolute Snyk CLI path. Ensure the file exists and the developer can run it.

If the CLI moved or a new installation uses a different path, run `snyk studio install --experimental` from the new binary to refresh a local installation. For managed deployments, ask your administrator to update the installed path.

For manual MCP setup, check the `command` in the client's configuration. Use the absolute path to the intended Snyk executable.

## Basic repository scanning

Use the same Snyk CLI binary as Studio to check whether scans run outside the assistant. From the repository root, run:

```bash
snyk test
snyk code test
```

The first command checks dependencies. The second checks first-party code. Resolve direct scanning failures, such as authentication, unsupported project configuration, or Snyk Code settings, before diagnosing the assistant integration.

## Authentication issues

Run `snyk auth` as the developer who uses the assistant and complete the browser authentication flow. The MCP client also requests repository trust when required.

For an ADS deployment, the push key associates deployment data with an Evo Tenant. Each developer still authenticates scanning with the Snyk CLI and MCP server individually.

If your environment supplies `SNYK_TOKEN`, ensure the CLI process launched by the assistant receives it. For token details, visit [Authentication for API](https://docs.snyk.io/developer-tools/snyk-api/authentication-for-api).

## Snyk Organization configuration

If you belong to multiple Organizations, configure the one to use for scanning:

```bash
snyk config set org=YOUR_ORG_ID
```

Replace *YOUR_ORG_ID* with your Snyk Organization ID. You can also use the `SNYK_CFG_ORG` environment variable. Ensure the account has access to that Organization and enable the Snyk products you need.

## Environment variable propagation

An assistant launched from a desktop application does not necessarily inherit the same environment as your terminal. Check that the process running Snyk receives the required authentication, Organization, region, and proxy settings.

For manual MCP integration, use the client's supported configuration mechanism to pass environment variables to the server process.

## Proxy configuration

If your network uses a proxy, ensure the `http_proxy` and `https_proxy` environment variables reach the Snyk CLI and MCP process. Check the proxy and certificate configuration for the installation route you use.

## Slow or incomplete scans

Open the repository you intend to scan and check the workspace path in the logs. In particular, check whether the assistant started in your home directory or in a larger workspace than intended.

Large repositories and dependency resolution can take longer to scan. Inspect the logs for completion, timeouts, and errors. A turn ending without a finding is not evidence that every scan completed successfully.

Use direct CLI scans to separate repository scanning problems from hook or MCP integration problems.

## SSE transport specifics

For a manual MCP integration using local server-sent events (SSE) transport, check the configured server URL and local firewall access to the server port. The Studio MCP server still runs locally in this mode.

The default Studio installer configures stdio transport. Use the configuration for your installation route when diagnosing a connection.

## Folder trust

Complete the MCP client's folder-trust request for the repository you want to scan. Trust applies to the directory the MCP server accesses. Check that the client opens the expected repository path.
