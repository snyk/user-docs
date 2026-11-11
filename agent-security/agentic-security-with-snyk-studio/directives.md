---
description: >-
  How directives govern the way AI coding assistants follow your security
  policy and standards
nav_context: classic
last_reviewed: 2026-07-21
last_reviewed_source: bulk-sync
---

# Directives

Directives guide an AI coding assistant's security workflow. Snyk Studio installs the hooks, commands, and skills supported by each assistant. For other clients, you can configure the MCP server and add rules manually.

| Mechanism | When it runs | Purpose |
| --- | --- | --- |
| Hook | At configured assistant lifecycle events | Runs programmatic logic, including background scans and evaluation of results. |
| Rule or instruction | When the assistant loads and follows the instruction | Tells the model when to invoke Snyk tools and how to respond to findings. |
| Command | When a developer or agent invokes the workflow | Starts a named procedure, such as vulnerability remediation. |
| Skill | When the assistant selects or invokes the packaged instructions | Supplies a workflow and supporting knowledge for a task. |

Guardrail directives are mainly hooks and rules that guide security behavior during development. Command directives are commands and skills that start named workflows on demand.

Visit [Getting started with Snyk Studio](getting-started-with-snyk-studio.md) to install the supported components. For enterprise deployment, use [Distribution at scale](distribution-at-scale.md).

## Guardrail directives

Guardrail directives guide security behavior during development. Hooks run programmatic logic when their configured event fires. Rules rely on the assistant following natural-language instructions, so the agent can skip a requested scan.

Installing a hook does not confirm that a scan completed. Authentication, tool permissions, product support, and scan failures still affect the result. Use [Studio logs](troubleshooting.md#verbose-logging-and-debugging) to diagnose activity.

## Secure at inception directives

For assistants with [hook support](getting-started-with-snyk-studio.md#coding-assistant-support), Studio tracks code and dependency changes and starts scans in the background. At the end of an agent turn, Studio evaluates the results and supplies remediation context for newly introduced vulnerabilities.

This workflow targets vulnerabilities introduced during the agent's work. Use the [remediation workflows](#remediation-directives) when you want to address an existing backlog.

The hooks run separately from direct MCP tool calls. You do not need a rule that tells the agent to invoke an MCP scan after every edit to start the hook-based workflow.

## Rules for manual MCP integrations

For a client without Studio hooks, add instructions using that client's rules mechanism. For example:

```text
Use snyk_code_scan to scan new or modified first-party code in supported languages.
Use snyk_sca_scan when adding or changing dependencies.
Review the scan results and address newly introduced vulnerabilities.
Rescan after changes to check the fixes.
Report scans that fail or cannot run instead of treating them as clean results.
```

These instructions guide the model. They do not install hooks or guarantee tool execution. This is the generic rule text. The quickstart guides link here and cover client-specific configuration, such as where each client reads rules. Visit the [quickstart guides](quickstart-guides/README.md).

<a id="secure-at-inception-package-health-check"></a>
<a id="secure-at-inception-package-health-check-experimental"></a>

## Package health checks

The secure dependency health check helps the assistant evaluate packages before recommending or adding them. It uses the `snyk_package_health_check` MCP tool to compare security, maintenance, popularity, and community signals.

The supported ecosystems are npm, PyPI, Maven, NuGet, and Go. The Studio installer includes the package-health workflow for the assistants in the [capability matrix](getting-started-with-snyk-studio.md#coding-assistant-support).

For example, ask your assistant:

```text
Compare the security and maintenance health of these packages before adding a dependency. Explain the recommendation and any trade-offs.
```

A package-health skill guides the assistant's decision. It does not establish a universal installation block. The separate [Cursor package-enforcement hook example](https://github.com/snyk/studio-recipes/tree/main/guardrail_directives/package_enforcement/cursor/hooks) implements a different workflow.

## Secrets at commit time

Secrets scanning at commit time checks staged files before a Git commit, including changes written by people and agents. It is separate from the code and dependency scans triggered during an agent turn.

The `snyk_secret_scan` tool is in the Full and Experimental MCP profiles. Its presence in the client's tool list does not confirm that a Git pre-commit hook is installed. Verify the commit-hook configuration and an actual commit-time scan separately.

For supported secrets scanning methods and prerequisites, visit [Snyk Secrets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-secrets). For a scan from your terminal, use the [Secrets test](https://docs.snyk.io/developer-tools/snyk-cli/commands/secrets-test) command.

## Remediation directives

Command directives start named workflows when a developer or agent invokes them. The installed form varies by assistant: some use slash commands or workflows, and others use skills.

Studio includes workflows to scan for existing vulnerabilities, apply fixes, and rescan the result:

| Workflow | Purpose |
| --- | --- |
| Snyk Fix | Investigates and fixes a selected vulnerability or vulnerability type, then validates the changes. |
| Snyk Batch Fix | Addresses multiple findings through a batch remediation workflow. |

For assistants with installed slash commands, invoke `/snyk-fix` or `/snyk-batch-fix`. Codex and GitHub Copilot CLI receive the Snyk Fix skill. Kiro uses steering content. The legacy `windsurf` installer target writes Cascade workflows. Devin Local does not support those workflows. Use [Connect Devin Desktop to Snyk](quickstart-guides/windsurf-guide.md) for manual MCP setup.

Review the changes and the validation results before accepting a fix. A proposed pull request is a review mechanism, not evidence that the application is free of vulnerabilities.

The maintained examples are in the [Studio recipes repository](https://github.com/snyk/studio-recipes). For the broader remediation setup and supported entry points, visit [Remediation Agent](https://docs.snyk.io/scan-fix-and-prevent/fix/remediation-agent).
