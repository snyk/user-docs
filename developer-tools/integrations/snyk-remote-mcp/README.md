---
description: Use Snyk Remote MCP to query existing Snyk data from AI assistants and agentic workflows
---

# Snyk Remote MCP

Snyk Remote MCP is a hosted [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server. It connects compatible AI assistants to read-only Snyk tools for exploring Organizations, Projects, issues, dependencies, software bills of materials (SBOMs), and other data that is already available in Snyk.

Snyk Remote MCP uses Snyk Apps OAuth 2.0 with Proof Key for Code Exchange (PKCE). You authorize access in the Snyk Web UI instead of copying a Snyk API token into the MCP client.

{% hint style="info" %}
Snyk Remote MCP does not modify Snyk data, onboard Projects, or run scans against source code on your local system. All advertised tools are read-only, non-destructive, and idempotent.
{% endhint %}

## Snyk Remote MCP and the local Snyk MCP Server

Snyk Remote MCP complements the local Snyk MCP Server provided by Snyk Studio.

|                             | Snyk Remote MCP                                    | Snyk MCP Server                                                       |
| --------------------------- | -------------------------------------------------- | --------------------------------------------------------------------- |
| Where it runs               | Hosted by Snyk and accessed over Streamable HTTP   | Locally through the Snyk CLI                                          |
| Authentication              | Snyk Apps OAuth 2.0 with PKCE                      | Snyk CLI authentication                                               |
| Snyk CLI required           | No                                                 | Yes                                                                   |
| Access to local source code | No                                                 | Yes                                                                   |
| Primary purpose             | Query and report on data already available in Snyk | Scan local code, open-source dependencies, containers, IaC, and SBOMs |

Use Snyk Remote MCP to ask questions such as:

* What are the highest-priority issues in an Organization or Project?
* Which issues have a concrete remediation?
* How does issue-instance risk compare across Organizations in a Group?
* Which Projects, targets, collections, or container images are available?
* What dependencies or SBOM components were captured in the latest monitored Project snapshot?
* Is a specific package version affected by known vulnerabilities?

Use the [Snyk MCP Server and Snyk Studio](https://app.gitbook.com/s/N5N885PkllOWeBmgm3Bp/agentic-security-with-snyk-studio/agentic-security-with-snyk-studio) when you want an AI assistant to scan files in a local workspace. You can configure both MCP servers in the same client.

## Prerequisites

Before connecting, ensure that:

* You have a Snyk account and access to the Snyk Organizations or Groups you want to query.
* Your MCP client supports remote Streamable HTTP servers, OAuth discovery, dynamic client registration, and PKCE.
* You know the [Snyk region](https://docs.snyk.io/snyk-data-and-governance/regional-hosting-and-data-residency#regional-urls) for your account.

The data returned by a tool depends on the access of the authorizing identity, the installed Snyk App, your Snyk plan, and product or feature availability. Authorizing Snyk Remote MCP does not grant access to Snyk data that the identity cannot otherwise view.

## Connect an MCP client

Select the endpoint for the Snyk region where your account is hosted:

| Region     | Snyk Remote MCP endpoint                 |
| ---------- | ---------------------------------------- |
| SNYK-US-01 | `https://api.snyk.io/mcp-server/mcp/`    |
| SNYK-US-02 | `https://api.us.snyk.io/mcp-server/mcp/` |
| SNYK-EU-01 | `https://api.eu.snyk.io/mcp-server/mcp/` |
| SNYK-AU-01 | `https://api.au.snyk.io/mcp-server/mcp/` |

For a single-tenant deployment, use the API hostname supplied for your deployment and append `/mcp-server/mcp/`.

### Connect from Cursor

Add Snyk Remote MCP to your Cursor MCP configuration:

```json
{
  "mcpServers": {
    "Snyk Remote MCP": {
      "url": "https://api.snyk.io/mcp-server/mcp/"
    }
  }
}
```

Replace the URL with the endpoint for your Snyk region.

### Connect from Claude Code

Run the following command:

```bash
claude mcp add -s user --transport http "Snyk Remote MCP" https://api.snyk.io/mcp-server/mcp/
```

Replace the URL with the endpoint for your Snyk region.

## Authorize access

On first use, the MCP client opens an authorization page in your browser.

1. Sign in to the Snyk Web UI for the same region as the configured MCP endpoint.
2. Review the requested Snyk App permissions.
3. Approve the app.
4. Return to the MCP client and use a Snyk Remote MCP tool.

The service requests the following read scopes:

| Scope                       | Used for                                                |
| --------------------------- | ------------------------------------------------------- |
| `org.read`                  | Identity and Organization discovery                     |
| `org.project.read`          | Projects and targets                                    |
| `org.project.snapshot.read` | Issues, SBOMs, dependency graphs, and issue aggregation |
| `org.package.test`          | Package vulnerability lookup and remediation enrichment |
| `org.report.read`           | Reporting data and Group issue views where available    |
| `org.collection.read`       | Project collections                                     |
| `org.container_image.read`  | Container image inventory                               |

The service validates the Snyk access token before every MCP request and uses that token for calls to the Snyk API. A read scope does not override Snyk roles, plan entitlements, or Early Access requirements. For example, Group data, audit logs, and batch package testing may not be available to every authorizing identity.

To review or revoke access, navigate to your personal **Account Settings** and select **Authorized Snyk Apps**. For more information, see [Managing Snyk Apps from the UI](../../snyk-api/using-specific-snyk-apis/snyk-apps-apis/about-snyk-apps.md#managing-snyk-apps-from-the-ui).

## Available tools

Snyk Remote MCP advertises 25 tools. These include discovery and raw-data tools, as well as reporting tools that combine Snyk API data into prioritized Markdown and structured results.

For the complete catalog and current limits, see [Snyk Remote MCP tools](available-tools.md).

## Work with tool results

List tools return one bounded page and include `has_more` and `next_cursor` when more data is available. The default page size is 50. A custom page size must be between 10 and 100 and a multiple of 10. This page-size constraint does not limit the total number of issues that can be retrieved.

List tools use `view: "summary"` by default to reduce the amount of data sent to the AI assistant. Use `view: "full"` only when you need the complete Snyk API resource.

Each tool returns stable MCP `structuredContent` and a JSON text representation for compatibility with different clients. Reporting tools also return a concise `markdown` field. When pagination, result limits, permissions, or failed enrichment can affect a conclusion, results include fields such as `coverage`, `truncated`, `has_more`, and `warnings`.

{% hint style="warning" %}
Issue reports count issue instances. One vulnerability or rule can produce multiple issue instances when it affects multiple Projects or scan items.
{% endhint %}

## Troubleshooting

### The authorization page does not open

Verify that your MCP client supports OAuth for remote Streamable HTTP servers. Check that the configured endpoint includes `/mcp-server/mcp/` and uses the API hostname for your Snyk region.

### A request returns `authentication` or HTTP 401

The access token is missing, expired, or invalid. Disconnect and reconnect Snyk Remote MCP in your MCP client to start the authorization flow again.

### A request returns `forbidden` or HTTP 403

A 403 response can indicate a Snyk role restriction, Snyk App access, plan or product entitlement, or an Early Access requirement. It does not always indicate a missing OAuth scope. Use the recovery hint in the tool result and verify that the authorizing identity can view the same resource in Snyk.

### A request is rate limited

Wait for the number of seconds in `retry_after_seconds`, then retry the request. Consider using a reporting tool instead of repeatedly retrieving and joining raw list results.

### A result is empty or incomplete

Check `warnings`, `coverage`, `truncated`, `has_more`, and `next_cursor` before treating an empty or partial result as complete. An empty SBOM or dependency graph can mean that the latest monitored Project snapshot does not contain the expected data.

If you need help, submit a [request](https://support.snyk.io) to Snyk Support.
