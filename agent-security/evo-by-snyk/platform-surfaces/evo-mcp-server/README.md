# Evo MCP Server

The Evo MCP server is a hosted [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server. It connects compatible AI assistants to your AI inventory, risk data, policies, and issues, so your own agent harness can answer questions and manage policy without leaving the tool you already work in.

{% hint style="info" %}
This is a remote MCP server. Nothing is installed on your machine, and the server has no access to your filesystem.
{% endhint %}

## Prerequisites

You need access to an Evo Tenant. To get set up, visit [Access and authentication](https://docs.snyk.io/agent-security/evo-by-snyk/access-and-authentication).

What you can do through the MCP server matches what you can do in Evo.

{% hint style="info" %}
The MCP server supports only non-destructive actions. Delete actions are not available.
{% endhint %}

## Evo MCP endpoint

| Region              | URL                                                  |
| ------------------- | ---------------------------------------------------- |
| Default, SNYK-US-01 | [https://evo.snyk.io/mcp](https://evo.snyk.io)       |
| SNYK-US-02          | [https://evo.us.snyk.io/mcp](https://evo.us.snyk.io) |
| SNYK-AU-01          | [https://evo.au.snyk.io/mcp](https://evo.au.snyk.io) |
| SNYK-EU-01          | [https://evo.eu.snyk.io/mcp](https://evo.eu.snyk.io) |

For a single-Tenant deployment, use the API hostname supplied for your deployment and append `/mcp`.

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
claude mcp add --scope user --transport http evo https://evo.snyk.io/mcp/
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

1. Sign in to the Evo Web UI.
2. Review the requested Evo permissions.
3. Approve the app.
4. Return to the MCP client and use an Evo MCP tool.

The agent runs as you, with your identity and your permissions.

The service requests the following read scope:

| Scope      | Used for                            |
| ---------- | ----------------------------------- |
| `org.read` | Identity and Organization discovery |

The service validates the Snyk access token before every MCP request, and uses that token for all downstream calls.

Available tools

Evo MCP advertises eight tools: six read tools and two write tools. Which tools you see depends on your permissions. Users with read access see the six read tools. Users with write access see all eight. The write tools are not listed at all for users who cannot use them.

### Schema Discovery

| Tool               | Description                                                                                                          |
| ------------------ | -------------------------------------------------------------------------------------------------------------------- |
| `get_query_schema` | Returns the fields available for a given schema, with their meanings. Invoked automatically before a query tool runs |

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

## What can you ask

Some examples of what the tools support:

* Which models are we running, and from which vendors and countries?
* Which assets have the most critical policy violations?
* Show me the policies governing our estate, then add a condition to one of them.

## Limits

Requests to the Evo MCP server are rate limited to 100 req/min per user. The limit applies to your Snyk user and is shared across every client you connect from. When a limit is exceeded, the server returns a 429 response with a `retry_after_seconds` value. Wait that many seconds, then retry.

## Troubleshooting

### Authorization page does not open

Verify that your MCP client supports OAuth for remote Streamable HTTP servers. Check that the configured endpoint includes `/mcp`.

### Request returns HTTP 401

The access token is missing, expired, or invalid. Disconnect and reconnect Evo MCP in your MCP client to start the authorization flow again.
