---
nav_context: agnostic
description: >-
  How Agentic Development Security secures the AI agents that build your
  software
last_reviewed: 2026-10-13
last_reviewed_source: focused-edit
---

# Agentic Development Security (ADS)

Agentic Development Security (ADS) secures the AI agents that build software. AI coding assistants and development agents pull in external tools, take actions across systems, and generate code. ADS secures three parts of that activity: what agents use, what they do, and what they generate.

## What ADS secures

* **Agent Supply Chain Security** secures what agents use. It evaluates the components your agents connect to - the Model Context Protocol (MCP) servers, skills, and external tools they use - before those components enter your development workflow. It discovers the components in use on end users' machines across your environment and assesses each one for security issues such as prompt injection, tool poisoning, and hardcoded secrets.
* **Agent Behavior Governance** secures what agents do. It works inside the agent execution loop, evaluating an agent's actions against policy before they run, with awareness of the full session: the sequence of actions, the tools in use, and the intent behind them.
* **Trusted Output Assurance**, also known as Snyk Studio, helps secure agent-generated output during development. In supported assistants, hooks trigger scans and evaluate results as code changes. Snyk surfaces findings and remediation guidance to the agent before code reaches a repository. Learn more about [Overview](../../agentic-security-with-snyk-studio/agentic-security-with-snyk-studio.md).
