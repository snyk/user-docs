---
nav_context: agnostic
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

To activate and deploy Agent Supply Chain Security, visit [Activation and deployment](activation-and-deployment.md).

## What gets scanned

Agent Supply Chain Security discovers two types of assets that agents use: MCP servers and skills. It scans the standard configuration locations for each supported agent.

## Supported agents

The following table shows agent support by operating system. A check mark (✓) means supported. A cross (✗) means the agent supports this but no paths are detected yet. A dash (—) means not applicable for that operating system.

| Agent          | macOS MCP | macOS Skills | Linux MCP | Linux Skills | Windows MCP | Windows Skills |
| -------------- | --------- | ------------ | --------- | ------------ | ----------- | -------------- |
| Claude Code    | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Claude Desktop | ✓         | ✗            | —         | —            | ✓           | ✗              |
| Cursor         | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| VS Code        | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Windsurf       | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Kiro           | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Gemini CLI     | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Antigravity    | ✓         | ✓            | ✓         | ✓            | ✓           | ✓              |
| Codex          | ✓         | ✓            | ✓         | ✓            | —           | —              |
| Amp            | ✗         | ✓            | ✗         | ✓            | ✗           | ✓              |
| Amazon Q       | ✓         | ✗            | ✓         | ✗            | ✓           | ✗              |
| OpenClaw       | ✗         | ✓            | ✗         | ✓            | ✗           | ✓              |
| OpenCode       | ✗         | ✗            | ✗         | ✗            | ✗           | ✗              |

## Risk indexes

Agent Supply Chain Security scores each MCP server and skill it discovers across a set of risk indexes. Each index scores one category of risk from 0 to 1,000. The higher the score, the more severe the finding. A component's Risk profile shows only the indexes that scored non-zero. Default policies raise an issue for any index at High or Critical severity (a score of 600 or above). Lower-severity findings appear in the Risk profile but do not raise an issue by default.

### MCP server risk indexes

* **Dangerous words**: manipulative language in a tool description that tries to influence the agent's decisions.
* **Prompt injection in tool**: an agent processes hidden instructions within a tool description as commands
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
* **Missing SKILL.md**: the skill lacks the SKILL.md file needed to evaluate it.

{% hint style="info" %}
Evo stores and displays the contents of .md files. Evo assesses every file the skill depends on, including code files, but does not store the contents of those files.
{% endhint %}

## Findings

Agent Supply Chain Security reports its findings to Evo, where you review them across your fleet: the machines that have run a scan, the MCP servers and skills found across those machines, and the issues raised against them.

### Machines and components

Your fleet shows every machine that has reported a scan, including how many MCP servers and skills Evo discovered and how many it successfully scanned.

For each component — MCP server or skill — Evo surfaces its risk profile, the issues raised against it, and where it appears across the fleet. For MCP servers, Evo also shows the tools the server exposes and the instructions it provides to the agent.

A component's scan status is **Complete** or **Discovery only**. Discovery only means Evo found the component but could not complete its scan, so it carries no risk assessment. A scan ends in Discovery only when, for example, the component's configuration cannot be parsed, a remote server requires credentials the scanner does not have, the server cannot start, or the analysis times out. For stdio MCP servers, a scan also ends in Discovery only when the server is not yet in Snyk's catalog of regularly assessed servers.

### Issues

Each issue records the asset it was raised against, its severity, the policy that triggered it, and where the asset appears across the fleet. The issue includes the risk evidence from the assessment and, when the policy defines them, remediation steps.
