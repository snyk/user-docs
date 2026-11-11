---
description: >-
  How Snyk Studio embeds security directives into AI-assisted development
  workflows
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: focused-edit
---

# Overview

Snyk Studio brings Snyk security scanning and remediation into AI-assisted development. It connects your coding assistant to Snyk tools, scans generated code and dependency changes in supported assistants, and gives the agent context to fix vulnerabilities.

Studio supports three related workflows:

| Workflow | How it works |
| --- | --- |
| Secure at inception | Hooks trigger background code and dependency scans during the agent's work. Studio evaluates the results and directs the agent to address newly introduced vulnerabilities. |
| Package selection | Package-health tools and skills help the agent compare dependencies before adding them. |
| Remediation | Commands and skills guide the agent through scanning, fixing, and validating existing vulnerabilities. |

## How Studio works

The Snyk CLI runs locally. Supported assistant lifecycle events trigger Studio hooks, which start background scans and evaluate results. The assistant receives security context for remediation when Studio identifies new vulnerabilities.

The local Model Context Protocol (MCP) server provides another way for the assistant to invoke Snyk tools. A background hook scan and a direct MCP tool call are different execution paths. Installing only the MCP server does not configure the complete Studio workflow.

Scan coverage depends on the assistant integration, Snyk product support, repository configuration, and authentication. Hook execution does not guarantee that every scan completes or that the agent resolves every finding. Use [Troubleshooting](troubleshooting.md) to check installation and scan activity.

## Coding assistant support

For the assistants the installer configures and the components each one receives, visit [Getting started with Snyk Studio](getting-started-with-snyk-studio.md#coding-assistant-support).

Other clients that support a local MCP server can invoke Snyk tools through [manual integration](getting-started-with-snyk-studio.md#configure-the-snyk-mcp-server). Rules give an agent instructions about when to scan. They do not provide the same execution mechanism as hooks.

## Choose a deployment path

* For enterprise deployment across teams, use the [Agentic Development Security (ADS) installer](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md). Visit [Distribution at scale](distribution-at-scale.md) for the rollout sequence.
* For local developer setup, use [Getting started with Snyk Studio](getting-started-with-snyk-studio.md#install-on-your-machine).

## MCP server supported tools

The MCP server exposes Snyk scanning and supporting tools according to the configured [MCP profile](getting-started-with-snyk-studio.md#configure-the-snyk-mcp-profile). These include code, dependency, container, infrastructure as code, and software bill of materials scanning, package-health checks, and authentication and folder-trust tools.

Inspect the tool list in your MCP client to check which tools your installed CLI and selected profile expose. Tool availability does not mean a Studio hook automatically runs that tool.

{% hint style="info" %}
Dependency scans can run ecosystem tools, such as Gradle or Maven, on your machine to resolve the dependency tree.
{% endhint %}

## Resources

* [Getting started with Snyk Studio](getting-started-with-snyk-studio.md) for setup and the assistant capability matrix
* [Directives](directives.md) for hooks, rules, remediation, package health, and commit-time secrets scanning
* [Distribution at scale](distribution-at-scale.md) for enterprise rollout through ADS
* [Troubleshooting](troubleshooting.md) for installation and scan diagnostics
* [Usage analytics](usage-analytics.md) for adoption and scan reporting
