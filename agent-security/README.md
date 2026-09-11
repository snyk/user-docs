---
description: >-
  How Snyk Studio embeds security directives into AI-assisted development
  workflows
nav_context: agnostic
description: >-
  Secure AI-assisted development and the AI systems your organization runs,
  across coding assistants, cloud AI platforms, and organization-wide AI risk
---

# Overview

Snyk secures AI-assisted development and the AI systems your organization runs, across three areas.

## Secure coding assistants with Snyk Studio

Agentic workflows speed up development, but AI-generated code can introduce vulnerabilities or insecure practices. [Snyk Studio](agentic-security-with-snyk-studio/) embeds security directives directly into your AI-assisted workflows, connecting the Snyk platform, your development environment, and your AI tools.

Use Snyk Studio when you want to:

* Guide coding assistants such as Claude Code, Codex CLI, Cursor, and Gemini CLI toward secure output as they generate code.
* Run Snyk scans (Snyk Code, Snyk Open Source, Snyk IaC) directly from an agent through the local Snyk MCP Server.
* Roll out consistent security directives across your organization's AI tooling.

To get started, see [Getting started with Snyk Studio](agentic-security-with-snyk-studio/getting-started-with-snyk-studio.md) or jump to a [quickstart guide](agentic-security-with-snyk-studio/quickstart-guides/) for your specific tool.

## Govern cloud AI platforms

Beyond an individual developer's coding assistant, Snyk extends visibility to the [cloud AI platforms](cloud-ai-platforms/overview.md) your organization runs, such as Anthropic's Claude Enterprise. These integrations surface AI models, approved MCP servers, and tool-level permissions as assets in your Evo inventory.

## Manage AI risk and posture with Evo by Snyk

[Evo by Snyk](evo-by-snyk/overview.md) is the platform layer that ties the above together: it surfaces risk context for AI assets discovered by AI-SPM and Agent Supply Chain Security, lets you apply that context in policies, and raises issues when an asset violates one. Use Evo when your focus is organization-wide AI risk visibility and policy enforcement rather than a single tool or platform.

## Choosing where to start

* Securing what an individual developer generates with an AI coding assistant → start with Snyk Studio.
* Governing a cloud AI platform your organization has adopted → start with Cloud AI platforms.
* Getting a unified view of AI risk and enforcing policy across both → start with Evo by Snyk.
