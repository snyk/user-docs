---
nav_context: agnostic
description: >-
  How Agent Behavior Governance secures what AI agents do inside their execution
  loop
---

# Agent Behavior Governance

Agent Behavior Governance secures what agents do as they run. It works within the agent's execution loop, evaluating each action against the policy before it runs. Because it sees the full session, including the sequence of actions, the tools in use, and the context behind them, it acts on patterns and intent, not on individual commands alone.

## Supported agents

Claude Code, Codex, Cursor, and GitHub Copilot.

{% hint style="warning" %}
**Codex in user mode requires each developer to accept a prompt**

After installation in user mode, Codex prompts the developer to allow hook configuration outside the managed path. If the developer does not accept, the hooks stay installed but do not run, and Agent Behavior Governance does not govern Codex on that machine. Managed mode does not show the prompt. To choose a mode, visit [Installation mode](../activation-and-deployment/README.md#installation-mode).
{% endhint %}

## Supported use cases

* [MCP Governance ](mcp-governance.md)

### Limits

* Rate limit: Each entitlement is limited to 5k behavioral guardrail requests (hook calls) per calendar day per Machine.
* Data retention limi&#x74;_:_ Data is retained for seven days. The data retention setting is fixed at seven days and cannot be changed.
