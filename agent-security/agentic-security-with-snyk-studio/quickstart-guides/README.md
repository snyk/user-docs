---
description: >-
  Quickstart guides for setting up Snyk Studio with common AI coding
  assistants
nav_context: agnostic
last_reviewed: 2026-07-21
last_reviewed_source: bulk-sync
---

# Quickstart guides

Choose the setup for your coding assistant. For local developer setup, the [Studio CLI installer](../getting-started-with-snyk-studio.md#install-on-your-machine) configures supported components together. For enterprise rollout, use the [Agentic Development Security (ADS) installer](../../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md).

## Assistants configured by the Studio installer

| Assistant | Setup guide |
| --- | --- |
| Claude Code | [Connect Claude Code to Snyk](claude-code-guide.md) |
| Cursor | [Connect Cursor to Snyk](cursor-guide.md) |
| Codex | [Connect Codex to Snyk](codex-cli-guide.md) |
| Gemini CLI (access restrictions apply) | [Connect Gemini CLI to Snyk](gemini-cli-guide.md) |
| GitHub Copilot CLI and VS Code | [Connect GitHub Copilot to Snyk](github-copilot-guide.md) |
| Kiro | [Connect Kiro to Snyk](kiro-guide.md) |

The installer configures MCP access and remediation and package-health components for all assistants in this table. Visit the [capability matrix](../getting-started-with-snyk-studio.md#coding-assistant-support) to see which assistants also receive Secure at inception hooks.

Google is restricting access to Gemini CLI and Gemini Code Assist. Visit [Migrate from Gemini tooling](antigravity-guide.md#migrate-from-gemini-tooling) for the schedule and the Antigravity path. The `gemini` installer target does not configure Antigravity CLI.

The installer retains a legacy `windsurf` target for old Cascade configuration. For agent-invoked Snyk tools in Devin Local, use [Connect Devin Desktop to Snyk](windsurf-guide.md). This manual integration does not install Studio hooks or skills.

## Import findings into Daybreak Defense Factory

Daybreak Defense Factory imports Snyk findings through the Codex Security CLI. The Studio installer does not configure it. For setup, visit [Connect Daybreak Defense Factory to Snyk](daybreak-defense-factory-guide.md).

## Manual MCP integrations

These guides explain how to connect a client to the local Snyk MCP server:

* [Connect Amazon Q to Snyk](amazon-q-guide.md)
* [Connect Antigravity to Snyk](antigravity-guide.md)
* [Connect Augment Code to Snyk](augment-code-guide.md)
* [Connect Cline to Snyk](cline-guide.md)
* [Connect Continue to Snyk](continue-guide.md)
* [Connect Devin to Snyk](devin-guide.md)
* [Connect Devin Desktop to Snyk](windsurf-guide.md)
* [Connect Factory to Snyk](factory-guide.md)
* [Connect Factory Terminal/IDE to Snyk](factory-terminal-ide-guide.md)
* [Connect Gemini Code Assist to Snyk](gemini-code-assist-guide.md)
* [Connect Goose CLI to Snyk](goose-cli-guide.md)
* [Connect JetBrains AI Assistant to Snyk](jetbrains-ai-assistant.md)
* [Connect JetBrains Junie to Snyk](jetbrains-junie.md)
* [Connect Qodo to Snyk](qodo-guide.md)

A manual MCP connection gives the assistant access to Snyk tools. It does not install the Studio Secure at inception hooks. Rules instruct the assistant to scan, but depend on the model following those instructions.

If the connection fails, visit [Troubleshooting](../troubleshooting.md).
