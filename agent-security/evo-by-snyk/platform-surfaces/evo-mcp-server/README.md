---
description: >-
  How to connect AI assistants to Evo through the hosted Evo MCP server, and
  which tools, permissions, and limits apply
nav_context: agnostic
---

# Evo MCP server

The Evo MCP server is a hosted [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server. It connects compatible AI assistants to your AI inventory, risk data, policies, and issues, so your own agent harness can answer questions and manage policy without leaving the tool you already work in.

{% hint style="info" %}
This is a remote MCP server. Nothing is installed on your machine, and the server has no access to your filesystem.
{% endhint %}

## Prerequisites

You need access to an Evo Tenant. To get set up, visit [Access and authentication](../../access-and-authentication.md).

What you can do through the MCP server matches what you can do in Evo. Your Tenant role determines which tools are available to you. For details, visit [Available tools](#available-tools).

{% hint style="info" %}
The MCP server supports only non-destructive actions. Delete actions are not available.
{% endhint %}

## Evo MCP endpoint

| Region              | URL                                                       |
| ------------------- | --------------------------------------------------------- |
| Default, SNYK-US-01 | [https://evo.snyk.io/mcp](https://evo.snyk.io/mcp)       |
| SNYK-US-02          | [https://evo.us.snyk.io/mcp](https://evo.us.snyk.io/mcp) |
| SNYK-AU-01          | [https://evo.au.snyk.io/mcp](https://evo.au.snyk.io/mcp) |
| SNYK-EU-01          | [https://evo.eu.snyk.io/mcp](https://evo.eu.snyk.io/mcp) |

For a single-Tenant deployment, add the `evo.` prefix to the hostname supplied for your deployment, then append `/mcp`: `https://evo.`_`your-deployment-hostname`_`/mcp`. Do not use the API hostname.

## Connect Evo MCP

Connect Evo MCP to Cursor, Claude Code, or Codex. Select your client, then add the regional endpoint for your deployment. On first use, the client opens your browser to authorize access.

{% tabs %}
{% tab title="Cursor" %}
Add Evo MCP to your Cursor MCP configuration:

```json
{
  "mcpServers": {
    "Evo MCP": {
      "url": "https://evo.snyk.io/mcp"
    }
  }
}
```
{% endtab %}

{% tab title="Claude Code" %}
Run the following command:

```bash
claude mcp add --scope user --transport http evo https://evo.snyk.io/mcp
```

To authorize up front, or to authorize again later, run:

```bash
claude mcp login evo
```
{% endtab %}

{% tab title="Codex" %}
Edit `~/.codex/config.toml`, or your Codex CLI configuration file, to add Evo MCP with the HTTP transport:

```toml
[mcp_servers.evo]
url = "https://evo.snyk.io/mcp"
```

Then log in:

```bash
codex mcp login evo
```

This opens your browser to complete the authorization flow. Codex stores the credentials until the token expires.
{% endtab %}
{% endtabs %}

## Authorize access

On first use, the MCP client opens an authorization page in your browser.

1. Log in to Snyk.
2. Review the requested Evo permissions and approve the app.
3. Return to the MCP client and use an Evo MCP tool.

The agent runs as you, with your identity and your permissions.

The service requests the `org.read` OAuth scope to identify you and discover your Organizations. This scope does not grant access to Evo data. Access to Evo data comes from your Tenant role, which Snyk checks on every request. The service also validates your Snyk access token before every MCP request and uses that token for all downstream calls.

## Available tools

Evo MCP advertises eight tools: six read tools and two write tools. Which tools you see depends on your Tenant role:

| Tenant role                               | Available tools                    |
| ----------------------------------------- | ---------------------------------- |
| Tenant Viewer                             | The six read tools                 |
| Tenant Admin, or a role with Evo access   | All eight tools                    |

The write tools are not listed at all for users who cannot use them. For the full list of roles, visit [Access and authentication](../../access-and-authentication.md#add-members).

### Schema discovery

| Tool               | Description                                                                                                           |
| ------------------ | --------------------------------------------------------------------------------------------------------------------- |
| `get_query_schema` | Returns the fields available for a given schema, with their meanings. Invoked automatically before a query tool runs. |

### Inventory and relationships

| Tool                   | Description                                                                                                                                                                                               |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `query_assets`         | Returns AI assets and their metadata: models, agents, MCP servers, skills, prompts, datasets, packages, repositories, applications, developer machines, and COS targets.                                  |
| `query_related_assets` | Returns relationships between assets, including indirect ones. Used for questions that map one asset type to another, such as the models a repository references or the MCP servers present on a machine. |

### Policies and issues

| Tool             | Description                                                                                                                      |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `query_policies` | Returns policy definitions and metadata, including each policy's conditions and whether it is a Snyk default or a custom policy. |
| `query_issues`   | Returns policy violations, with the violating asset and the location it was found in.                                            |

### Policy authoring

| Tool             | Description                                                                                                                                 |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `resolve_values` | Returns the canonical value for a policy clause from a readable name, such as a vendor or a license. Used when preparing a policy.          |
| `create_policy`  | Creates a new custom policy for your Tenant. Requires write access and explicit approval in your client.                                    |
| `update_policy`  | Updates an existing custom policy by UUID. Requires write access and explicit approval in your client. You can update only custom policies. |

The write tools ask for your approval through MCP elicitation, a client feature that prompts you to confirm an action. Your client must support elicitation to use `create_policy` and `update_policy`. If you do not respond within five minutes, the request times out and nothing is written.

## What can you ask

Some examples of what the tools support:

* Which models are we running, and from which vendors and countries?
* Which assets have the most critical policy violations?
* Show me the policies governing our estate, then add a condition to one of them.

For multi-step workflows that combine Evo MCP with other MCP servers, visit [Common use cases](common-use-cases.md).

## Supported features

The Evo MCP server supports the following:

* Querying AI assets in your inventory and the relationships between them
* Querying policies and the issues they raise
* Creating and updating custom policies

Preview features are not available through the MCP server.

## Rate limits

Snyk applies rate limits per Snyk user. The limits are shared across every client and connection you use in the same region.

| Request type    | Per second | Per minute | Per hour |
| --------------- | ---------- | ---------- | -------- |
| Authenticated   | 20         | 300        | 10,000   |
| Unauthenticated | 10         | 120        | 2,000    |

Snyk applies unauthenticated limits per source IP address.

When you exceed a limit, the server returns an HTTP 429 response with a `Retry-After` header. Wait the number of seconds in the header, then retry.

## Troubleshooting

### Authorization page does not open

Verify that your MCP client supports OAuth for remote Streamable HTTP servers. Check that the configured endpoint includes `/mcp`.

### Request returns HTTP 401

The access token is missing, expired, or invalid. Disconnect and reconnect Evo MCP in your MCP client to start the authorization flow again.

### Tool call returns permission_denied

Your Tenant role does not include Evo access. Ask your Tenant Admin to assign a role with Evo access, then reconnect Evo MCP in your MCP client.

### Write tools return unsupported_client

Your MCP client does not support elicitation, so it cannot ask you to approve policy changes. Use a client that supports MCP elicitation to create or update policies. The read tools work in any compatible client.

### Request returns HTTP 429

You exceeded a rate limit. Wait the number of seconds in the `Retry-After` header, then retry. For the limits, visit [Rate limits](#rate-limits).
