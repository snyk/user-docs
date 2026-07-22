---
nav_context: agnostic
---

# Agent Behavior Governance

{% hint style="info" %}
**Feature availability**

Agent Behavior Governance is in open preview. It is not connected to the platform surfaces: Inventory, Policies and issues, Reports, and Evo chat. Refer to the full list of limitations in the [Open preview limits](agent-behavior-governance.md#open-preview-limits) section.
{% endhint %}

Agent Behavior Governance secures what agents do as they run. It works within the agent's execution loop, evaluating each action against the policy before it runs. Because it sees the full session — the sequence of actions, the tools in use, and the context behind them — it acts on patterns and intent, not on individual commands alone.

## Capabilities

* Observe agent actions across systems, including tool use, command execution, and data access.
* Enforce policies that control how agents interact with systems, tools, and data.
* Block high-risk or unauthorized actions before they complete.
* Modify actions in real time, such as steering an agent prompt before it leaves your environment.
* Block or redirect agent actions targeting specific systems, environments, or services based on context and policy.
* Detect and respond to threats such as prompt injection, agent hijacking, and data exfiltration.

## Supported agents

Claude Code, Codex, and Cursor.

## How it works

You view and manage Agent Behavior Governance in the **Observe** area. During open preview, Observe is separate from the platform surfaces.

Agent Behavior Governance captures the runtime activity of your AI agents — prompts, tools, Model Context Protocol (MCP) calls, shell commands, and file access — and runs a set of scanners over that activity. Each scanner raises findings, and a policy decides what happens when a finding type is detected.

### Observe

The Observe view is organized into the following tabs:

* **Overview**: a summary of security posture with three sub-tabs:
  * **Findings**: all detections raised by scanners.
  * **Issues**: policy enforcement outcomes — the subset of findings where a policy decided to log, steer, ask, or block.
  * **Active Scanners**: the list of active scanner policies with trend data.
* **Machines**: machines running Agentic Development Security hooks, with per-machine agents, activity, findings, and last active time.
* **Activity**: a timeline of every agent action across your fleet — prompts, tool calls, shell commands, and file operations — grouped by machine and session.
* **Policies**: policy management — create, edit, reorder, and configure guard policies.

### Scanners

Agent Behavior Governance includes eight scanners:

* **Indirect Prompt Injection**: hidden instructions in tool outputs that attempt to hijack agent behavior.
* **Secrets Detection**: exposed credentials, API keys, tokens, and private keys in agent messages and tool outputs.
* **PII Detection**: personally identifiable information (PII) — names, emails, phone numbers, government IDs, financial identifiers, and other sensitive data types across multiple jurisdictions — in prompts and in arguments to shell commands.
* **Sensitive Shell Commands**: destructive operations, requests to external endpoints, and reads of sensitive files.
* **Workspace Boundary**: file operations outside the session workspace or in protected directories such as `node_modules` and `.venv`. Supports custom allow and deny path globs per Tenant.
* **MCP Tool Observability**: MCP tool calls, logging tool invocations for observability and audit.
* **Web Search & HTTP**: external HTTP activity from shell commands such as `curl` and `wget`, and from web-search MCP tools.
* **Toxic Flows**: dangerous combinations across findings — for example, secrets sent to a public endpoint, or untrusted input followed by a destructive command — that individual scanners miss.

### Findings

A **finding** is a detection raised by a scanner. Each finding records the scanner that raised it, its severity, and the machine and client it came from. A finding's detail explains its potential impact, includes the model's reasoning for the detection, shows the evidence, and links to the related agent activity and the full session.

### Activity and sessions

Activity tracking records every agent action across your fleet — prompts, tool calls, shell commands, and file operations — including which machine, session, and policy it involved. A session provides a full timeline of every action the agent took.

### Machines

Machines appear automatically after you install Agentic Development Security hooks. Each machine shows its agents, activity, findings, and last active time. Selecting a machine shows its findings, activity, and related sessions.

### Open preview limits

Agent Behavior Governance is available only as an open preview and is not generally available (GA). The product will be limited in the following ways:

* Rate limit: Each entitlement in open preview will be limited to 10k behavioral guardrailing requests (hook calls) per calendar day, per single Tenant.
* Data retention limi&#x74;_:_ Data is retained for seven days. During the preview, the data retention setting is fixed at seven days and cannot be changed.
* Product support: This experimental preview is subject to change. It is a try-out experience and should not be used in production. On-call and service availability will be limited during this period. This means there is limited support for feedback and feature requests.
