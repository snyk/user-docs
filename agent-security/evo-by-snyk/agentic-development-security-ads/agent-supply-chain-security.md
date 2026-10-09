---
nav_context: agnostic
description: >-
  How Agent Supply Chain Security assesses the MCP servers, skills, and tools
  your agents use
---

# Agent Supply Chain Security

Agent Supply Chain Security secures the components your agents use. AI agents connect to Model Context Protocol (MCP) servers, run skills, and call external tools. These components enter your environment without the review that first-party code receives, and they can introduce security issues into your development workflows. Agent Supply Chain Security discovers these components and assesses each one for risk.

## How it works

Agent Supply Chain Security inspects the agents installed on an end user's machine to build an inventory of the components they use, then assesses each component for security issues and reports its findings to Evo.

How it assesses an MCP server depends on the server:

* Agent Supply Chain Security matches a local (stdio) MCP server against the Snyk catalog and assesses it from that signature. It does not start the server. Snyk regularly computes these risk assessments. Not every stdio server is in the catalog yet, and coverage expands as Snyk adds servers.
* Agent Supply Chain Security contacts a remote server directly to retrieve its current tools and capabilities.

It assesses skills from their files.

## Setup

To activate and deploy Agent Supply Chain Security, visit [Activation and deployment](activation-and-deployment/).

{% hint style="info" %}
To view Agent Supply Chain Security data, you must have a Tenant role with full Evo access. Snyk discovers these assets on end users' machines. Because these assets do not belong to a Snyk Organization, users restricted to specific Organizations cannot see this data. Visit [Access and authentication](../access-and-authentication.md) for more information.
{% endhint %}

## What gets scanned

Agent Supply Chain Security discovers two types of assets that agents use: MCP servers and skills. It scans the standard configuration locations for each supported agent.

## Supported agents

The following table shows agent support by operating system. A check mark (✓) means supported. A cross (✗) means the agent supports this, but no paths are detected yet. A dash (—) means not applicable for that operating system.

<table><thead><tr><th width="135.62109375">Agent</th><th>macOS MCP</th><th>macOS Skills</th><th>Linux/WSL MCP</th><th>Linux/WSL Skills</th><th>Windows MCP</th><th>Windows Skills</th></tr></thead><tbody><tr><td>Claude Code</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Claude Desktop</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✗</td></tr><tr><td>Cursor</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>VS Code</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>GitHub Copilot</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Windsurf</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Kiro</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Gemini CLI</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Antigravity</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td></tr><tr><td>Codex</td><td>✓</td><td>✓</td><td>✓</td><td>✓</td><td>—</td><td>—</td></tr><tr><td>Amp</td><td>✗</td><td>✓</td><td>✗</td><td>✓</td><td>✗</td><td>✓</td></tr><tr><td>Amazon Q</td><td>✓</td><td>✗</td><td>✓</td><td>✗</td><td>✓</td><td>✗</td></tr><tr><td>OpenClaw</td><td>✗</td><td>✓</td><td>✗</td><td>✓</td><td>✗</td><td>✓</td></tr><tr><td>OpenCode</td><td>✗</td><td>✗</td><td>✗</td><td>✗</td><td>✗</td><td>✗</td></tr></tbody></table>

#### Claude component detection coverage

The following table shows which Claude components Snyk detects and where they appear in Evo.

<table><thead><tr><th>Component</th><th></th><th>Supported</th><th>Appears on</th><th data-hidden>Detected when</th></tr></thead><tbody><tr><td rowspan="6">Claude Code</td><td>MCP server (local or remote)</td><td>Yes</td><td><strong>Machines</strong> tab and <strong>MCP servers</strong> page</td><td>Scan finds the server in a config file</td></tr><tr><td>MCP server from a Claude Code plugin</td><td>Yes</td><td><strong>Machines</strong> tab and <strong>MCP servers</strong> page</td><td></td></tr><tr><td>Skill from a Claude Code plugin</td><td>Yes</td><td><strong>Machines</strong> tab</td><td>Scan finds the server in <code>claude_desktop_config.json</code></td></tr><tr><td>Claude connector</td><td>Yes, only when used through Claude Code (terminal or the <strong>Code</strong> tab in Claude Desktop)</td><td><strong>MCP servers</strong> page only</td><td>Scan finds <code>&#x3C;skill>/SKILL.md</code> in a standard skills folder</td></tr><tr><td>Skill via Claude Skill Marketplace</td><td>No</td><td>—</td><td></td></tr><tr><td>Standalone skill (<code>SKILL.md</code>)</td><td>Yes</td><td><strong>Machines</strong> tab</td><td></td></tr><tr><td rowspan="5">Claude Desktop</td><td>MCP server from a Claude Desktop plugin</td><td>Yes (on macOS and Linux only)</td><td><strong>Machines</strong> tab and <strong>MCP servers</strong> page</td><td>Scan finds <code>&#x3C;skill>/SKILL.md</code> in the installed plugin</td></tr><tr><td>Local MCP server in the Claude Desktop configuration</td><td>Yes</td><td><strong>Machines</strong> tab and <strong>MCP servers</strong> page</td><td></td></tr><tr><td>Skill from a Claude Desktop plugin</td><td>Yes (on macOS and Linux only)</td><td><strong>Machines</strong> tab</td><td>Scan finds the server in the installed plugin</td></tr><tr><td>Skill via Claude Skill Marketplace</td><td>No</td><td>—</td><td>Scan finds <code>&#x3C;skill>/SKILL.md</code> in the installed plugin</td></tr><tr><td>Standalone skill (<code>SKILL.md</code>)</td><td>Yes</td><td><strong>Machines</strong> tab</td><td></td></tr></tbody></table>



## MCP server and skill risk indexes

{% hint style="info" %}
MCP server and skill risk indexes use their own severity ranges. They are not on the same scale as the Model Risk Score used for AI models, and the two are not comparable. For model risk, visit [Risk Intelligence](../ai-spm/risk-intelligence/).
{% endhint %}

Agent Supply Chain Security scores each MCP server and skill it discovers across a set of risk indexes. Each index scores one category of risk from 0 to 1,000. The higher the score, the more severe the finding.

A component's Risk profile shows only the indexes that scored non-zero. Default policies raise an issue for any index at High or Critical severity (a score of 600 or above). Lower-severity findings appear in the Risk profile but do not raise an issue by default.

### MCP server risk indexes

* **Dangerous words**: manipulative language in a tool description that tries to influence the agent's decisions.
* **Prompt injection in a tool**: an agent processes hidden instructions within a tool description as commands.
* **Untrusted content**: tools that pull in attacker-controllable content, such as inbound emails or issue trackers.
* **Private data**: tools that retrieve sensitive, non-public data, such as personal communications, financial records, or credentials.
* **Destructive capabilities**: tools that can modify shared infrastructure, run system commands, or move money.

### Skill risk indexes

* **Prompt injection**: hidden or deceptive instructions in a skill's content.
* **Suspicious download URL**: instructions to download or run files from untrusted or obscured URLs.
* **Malicious code**: code patterns that indicate exfiltration, backdoors, or obfuscated payloads.
* **Insecure credential handling**: instructions that pass secrets such as API keys, tokens, or passwords through the model.
* **Hardcoded secret**: live credentials embedded directly in the skill.
* **Direct money access**: tools that let the agent execute financial transactions on its own.
* **Third-party content exposure**: instructions to fetch and act on untrusted public content, such as web pages or social posts.
* **Unverifiable dependencies**: instructions to fetch external code or prompts from remote URLs at runtime.
* **Attempt to modify system services**: instructions to change the host's system files, accounts, or privileges.
* **Missing SKILL.md**: the skill lacks the `SKILL.md` file needed to evaluate it.

{% hint style="info" %}
Evo stores and displays the contents of .md files. Evo assesses every file the skill depends on, including code files, but does not store the contents of those files.
{% endhint %}

## Findings

Agent Supply Chain Security reports its findings to Evo, where you review them across your fleet: the machines that have run a scan, the MCP servers and skills found across those machines, and the issues raised against them.

### Machines and components

Your fleet shows every machine that has reported a scan, including how many MCP servers and skills Evo discovered and how many it successfully scanned.

For each component — MCP server or skill — Evo surfaces its risk profile, the issues raised against it, and its distribution across the fleet. For MCP servers, Evo also shows the tools the server exposes and the instructions it provides to the agent.

A component's scan status is **Complete** or **Discovery only**. Discovery only means Evo found the component but could not complete its scan, so it carries no risk assessment. A scan ends in Discovery only when, for example, the component's configuration cannot be parsed, a remote server requires credentials the scanner does not have, the server cannot start, or the analysis times out. For stdio MCP servers, a scan also ends in Discovery only when the server is not yet in Snyk's catalog of regularly assessed servers.

### Issues

Each issue records the asset it was raised against, its severity, the policy that triggered it, and where the asset appears across the fleet. The issue includes the risk evidence from the assessment and, when the policy defines them, remediation steps.
