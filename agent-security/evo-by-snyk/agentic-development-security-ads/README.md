---
nav_context: agnostic
---

# Agentic Development Security (ADS)

{% hint style="info" %}
**Feature availability**

Agent Behavior Governance is in open preview. As such it is not connected to the platform surfaces (Inventory, Policies & issues, Reports, and Evo chat) and comes with limitations described [here](agent-behavior-governance.md#open-preview-limits).
{% endhint %}

Agentic Development Security (ADS) secures the AI agents that build software. AI coding assistants and development agents pull in external tools, take actions across systems, and generate code. ADS secures three parts of that activity: what agents use, what they do, and what they generate.

## What ADS secures

* **Agent Supply Chain Security** secures what agents use. It evaluates the components your agents connect to - the Model Context Protocol (MCP) servers, skills, and external tools they use - before those components enter your development workflow. It discovers the components in use on end users' machines across your environment and assesses each one for security issues such as prompt injection, tool poisoning, and hardcoded secrets.
* **Agent Behavior Governance** secures what agents do. It works inside the agent execution loop, evaluating an agent's actions against policy before they run, with awareness of the full session: the sequence of actions, the tools in use, and the intent behind them.
* **Trusted Output Assurance**, also known as Snyk Studio, secures agent-generated output at creation. It runs inside the coding assistant. Hooks scan and fix AI-generated code as developers write it. This catches vulnerabilities and insecure dependencies before the code reaches a repository. Learn more about [Agentic security with Snyk Studio](https://app.gitbook.com/s/N5N885PkllOWeBmgm3Bp/).
