---
description: >-
  Establishing an MCP server policy in Evo, evaluating real usage against your
  environment, and safely enabling runtime enforcement.
---

# MCP Governance

Evo inspects the Model Context Protocol (MCP) servers declared in your code and installed on developer machines to build an inventory of the servers in use. Under the allowlist policy, Evo blocks unauthorized server execution.

## Prerequisites

To activate MCP Governance, ensure the following prerequisites are met:

* The Agentic Development Security (ADS) installer, deployed fleet-wide with both Agent Supply Chain Security and Agent Behavior Governance enabled, providing machine-level visibility into installed MCP servers and runtime enforcement of your policy.
* AI-SPM (SCM integration), connected to provide code-level visibility into MCP servers declared in your repositories.

For more information, visit [ADS Activation and deployment](../activation-and-deployment.md), and [AI-SPM](https://docs.snyk.io/agent-security/evo-by-snyk/ai-spm).

## How to use

MCP Governance rolls out in four phases, from establishing a baseline to enforcing your policy at runtime:

1. Baseline your inventory: review the unified MCP Servers Inventory Report to discover which servers are in use, how widely they are deployed, and their risks before you enforce rules.
2. Define your allowlist: import authorized servers into Evo based on your findings and company policies.
3. Analyze the impact: revisit the unified MCP Servers Inventory Report to evaluate the effect on your fleet and triage findings before you enforce rules.
4. Enforce blocking (optional): enable runtime blocking after your baseline and policy reflect production reality.

#### Baseline your inventory

* Navigate to the **MCP Servers** page to review active usage: all MCP servers detected across code repositories and developer machines, including which servers agents invoked.
* Review the exposure spread of each server (Code, Machines, Runtime) and associated risks to understand how your fleet uses it and its potential reach.

#### Define your allowlist

* Collect the list of MCP server names your company allows.
* Navigate to **Policies** > **Asset Classification** > **Allowed MCP Servers** to configure the policy.

{% hint style="info" %}
This policy ships empty and disabled by default.
{% endhint %}

* Enter allowed servers by name in the policy.

<figure><img src="https://2912686697-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FN5N885PkllOWeBmgm3Bp%2Fuploads%2FAiykJW6wEQYoePJWsadw%2Funknown.png?alt=media&#x26;token=2c1c5c58-49c0-4f10-b3c8-9f16d63464b7" alt="Policy screen showing the enforcement toggle off, an approved servers list, and an Evo chat panel" width="563"><figcaption><p>Allowed MCP Servers policy with enforcement disabled and an Evo chat session</p></figcaption></figure>

#### Impact analysis

Review the updated report, then return to the MCP Servers page with your defined policy active and blocking disabled. Each row displays its updated authorization status, either **Allowed** or **Out of Policy**.

Decide whether to add legitimate tools to the allowlist or remediate unauthorized usage.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-09-30 at 09.43.44.png" alt="Table listing detected MCP servers with approval status, source, and issue severity" width="563"><figcaption><p>MCP servers inventory with approval status and issue counts</p></figcaption></figure>

#### Enforce blocking

To enable blocking, return to the policy and set enforcement to **On**. After you enable blocking, Evo blocks any agent that attempts to execute a server designated as **Out of Policy**, directly on the developer's machine at invocation.

You can turn off enforcement at any time. Doing so disables blocking immediately.

## Developer experience

When an agent attempts to call an unauthorized MCP server, Evo halts execution inline on the developer's workstation. The developer receives an explicit notification detailing:

* the specific policy enforced
* the blocked tool call
* the targeted MCP server name

## Use the Evo MCP server

The [Evo MCP server](../../platform-surfaces/evo-mcp-server/) lets you interact with your MCP Governance data in a conversational way, rather than only through the MCP Servers page. This is useful for ad hoc investigation, cross-referencing data you keep outside Evo, and building a policy without leaving the chat.

Common use cases:

* See what your fleet is using, for example: "Give me the list of MCP servers my fleet is using."
* Build your allow list from a list you already maintain within your Organization, assuming your agent has access to it. This list is separate from the Allowed MCP Servers policy and does not affect blocking.
* Find where out-of-policy usage is happening, for example: "Show me the list of machines using an MCP server outside my allowlist."
* Build a dashboard from this data, based on your preference and the signals you need.
* Cross-reference against your own records.
* Route violations to Jira tickets (see Evo MCP server common use cases).

For the full set of tools available through Evo MCP, see the [Evo MCP server](../../platform-surfaces/evo-mcp-server/) page.
