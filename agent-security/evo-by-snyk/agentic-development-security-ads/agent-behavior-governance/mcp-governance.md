---
last_reviewed: 2026-10-30T00:00:00.000Z
last_reviewed_source: focused-edit
description: >-
  Establishing an MCP server policy in Evo, evaluating real usage against your
  environment, and safely enabling runtime enforcement.
---

# MCP Governance

Evo inspects the Model Context Protocol (MCP) servers declared in your code and installed on developer machines to build an inventory of the servers in use and can block server execution on the developer's machine at invocation.

## Prerequisites

To activate MCP Governance, ensure the following prerequisites are met:

* The Agentic Development Security (ADS) installer, deployed fleet-wide with both Agent Supply Chain Security and Agent Behavior Governance enabled, providing machine-level visibility into installed MCP servers and runtime enforcement of your policy.
* AI-SPM (SCM integration), connected to provide code-level visibility into MCP servers declared in your repositories.

For more information, visit [ADS Activation and deployment](../activation-and-deployment/), and [AI-SPM](../../ai-spm/README.md).

## How to use

MCP Governance takes three steps:

1. Understand your MCP usage: see which MCP servers are in use and by whom.
2. Define your rules and preview the impact: classify servers with an allowlist, set risk-based conditions, or both. Then return to the inventory to see which servers fall out of policy and how far they reach.
3. Enforce blocking (optional): turn on runtime blocking once your rules reflect production reality.

#### Understand your MCP usage

Go to the MCP Servers page to see all MCP servers detected across code repositories and developer machines, including which servers agents invoked.

* Exposure: see where each server appears (Code, Machines, Runtime) and how widely it is used.
* Risk: review each server's risk dimensions, including whether it comes from an Official vendor, alongside the other [risk dimensions scanned by Evo](../agent-supply-chain-security.md#mcp-server-risk-indexes).

#### Define your rules and preview the impact

Choose how you want to define what's allowed. You can use either approach or both.

**Option 1: Classify with an allowlist**

Classify the servers you trust. Everything else shows as Not Allowed, so you can focus on the servers outside your list.

1. Collect the MCP server names your organization allows, or use Evo's MCP risk profile to build the list.
2. Use Evo Chat to classify them or do it manually: go to Policies > Asset Classification > Allowed MCP Servers and enter the server names.

Best for: teams that want a defined set of approved servers and a clear status for every server, with the flexibility to avoid blocking developers over low-risk tools.

> Note: The allowlist ships empty and disabled by default.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-10-21 at 10.23.15.png" alt="Policy screen showing the enforcement toggle off, an approved servers list, and an Evo chat panel"><figcaption><p>MCP Servers classification</p></figcaption></figure>

**Option 2: Define risk-based conditions (no list required)**

Build a policy around Evo risk dimensions, tailored to your business rules. For example, flag servers with prompt injection in a tool, destructive capabilities, or unofficial publisher status.\
Go to Policies > Create policy > Monitor MCP servers that are not allowed.

Best for: teams that want governance to scale without managing a list.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-10-21 at 10.33.46.png" alt="MCP Governance policy"><figcaption><p>MCP Servers Risk based policy</p></figcaption></figure>

**Preview the impact**

With your rules active and blocking still off, check how your rules apply to your fleet. Nothing is blocked at this stage, so repeat as often as needed.

* Allowlist: return to the MCP Servers page. Each server shows its authorization status, either Allowed or Not Allowed. Use this to see which unapproved servers are in use, by whom, and how risky they are.
* Risk-based conditions: While building your policy, review the Block preview section before saving to assess the potential impact. Once saved, servers that meet your conditions will trigger policy issues (rather than receiving an Allowed status). Review these issues to pinpoint matching servers, user activity, and blast radius.

From either view, decide whether to add legitimate tools to your allowlist, adjust your conditions, or remediate unauthorized usage.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-10-21 at 10.39.47.png" alt="Table listing detected MCP servers with approval status, source, and issue severity"><figcaption><p>MCP servers inventory with approval status and issue counts</p></figcaption></figure>

#### Enforce blocking (optional)

When your rules reflect production reality, turn on blocking. Evo will then block non-compliant agent requests directly on the developer's machine at invocation.

You can enforce blocking by enabling the Block action on an existing policy, or by creating a dedicated policy through Policies > Create policy > Block MCP servers that are not allowed.

> Setup Requirement: Your policy must include at least the Type is MCP Server condition.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-10-23 at 09.52.50.png" alt=""><figcaption></figcaption></figure>

If you set an allowlist, it will act as a policy condition alongside Evo's risk dimensions, allowing you to tailor your enforcement scope:

* Block everything outside your allowlist: Use allowlist status as your sole condition. Any server marked as "Not Allowed" is immediately blocked.
* Block high-risk cases only: Combine allowlist status with specific risk conditions. For example, block unapproved servers that present high overall risk, contain prompt injection threats in a tool, or possess destructive capabilities. Lower-risk unapproved servers remain unblocked and visible for review.

You can also block based on risk alone without using the Allowlist attribute.

> Note: Enforcement can be toggled off at any time to immediately disable blocking across your fleet.

## Developer experience

When an agent attempts to call an unauthorized MCP server, Evo halts execution inline on the developer's workstation. The developer receives an explicit notification detailing:

* the specific policy enforced
* the blocked tool call
* the targeted MCP server name

## Use the Evo MCP server

The [Evo MCP server](../../platform-surfaces/evo-mcp-server/) lets you interact with your MCP Governance data in a conversational way, rather than only through the MCP Servers page. This is useful for ad hoc investigation, cross-referencing data you keep outside Evo, and building a policy without leaving the chat.

### Common use cases

#### See what is in use

* List your fleet's MCP servers: "Give me the list of MCP servers my fleet is using."
* Look up one server: "What's the risk of Slack MCP?"

#### Find problems

* Out-of-policy usage: "Show me the machines using an MCP server outside my allowlist."
* Risky servers in use: "Which high-risk servers are actually being called?" or "Which unapproved servers score above X?"

#### Take action

Creating a policy or updating the allowlist needs a Tenant role with Evo write access and a client that supports MCP elicitation. For details, visit [Available tools](../../platform-surfaces/evo-mcp-server/#available-tools).

* Create a policy: "Create a policy to flag or block servers above 600."
* Build your allowlist from an existing list: if your agent can access a list you already maintain, it can build the allowlist in Evo from it.&#x20;

#### Report and integrate

* Build a dashboard from this data, using the signals you need.
* Cross-reference the data against your own records.
* Route violations to Jira tickets. Visit [Evo MCP server common use cases](../../platform-surfaces/evo-mcp-server/common-use-cases.md#turn-a-violation-into-a-ticket).

For the full set of tools available through the Evo MCP server, see the [Evo MCP server](../../platform-surfaces/evo-mcp-server/) page.
