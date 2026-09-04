---
description: >-
  How AI assistants and agents read Snyk documentation, including the MCP
  server, llms.txt, Markdown page URLs, and content negotiation
---

# AI and agent access

Snyk publishes its documentation in formats that AI assistants and agents can consume directly. You can connect an agent to the documentation MCP server, fetch any page as Markdown, or read a single index of the whole site.

Use this page when you are building an agent, configuring an AI assistant, or writing a tool that reads Snyk documentation programmatically.

## Access methods

| Method | Endpoint | Use it for |
|---|---|---|
| MCP server | `https://docs.snyk.io/~gitbook/mcp` | Search and read documentation from an MCP-compatible client |
| Documentation index | `https://docs.snyk.io/llms.txt` | A linked index of every page |
| Full-text corpus | `https://docs.snyk.io/llms-full.txt` | The documentation as a single text file |
| Markdown pages | Any page URL with `.md` appended | Fetching one page without HTML |
| Content negotiation | Any page URL with an `Accept` header | Fetching Markdown without changing the URL |

## MCP server

The documentation MCP server exposes Snyk documentation to any client that supports the Model Context Protocol (MCP), such as Claude, Cursor, or a custom agent.

The endpoint is:

```
https://docs.snyk.io/~gitbook/mcp
```

The server accepts JSON-RPC over HTTP POST and responds with a server-sent event stream. A `GET` request returns `405 Method Not Allowed`, which is expected.

### Available tools

The server exposes four tools:

| Tool | What it does |
|---|---|
| `searchDocumentation` | Searches the documentation and returns matching pages with links |
| `getPage` | Fetches the complete Markdown content of one page |
| `askQuestion` | Answers a natural-language question and cites the source pages |
| `sendFeedback` | Reports a documentation issue to the Snyk documentation team |

### Connect a client

Add the endpoint to the MCP configuration for your client. For a client that reads a JSON configuration file, the entry is:

```json
{
  "mcpServers": {
    "snyk-docs": {
      "url": "https://docs.snyk.io/~gitbook/mcp"
    }
  }
}
```

The server requires no authentication and no API token.

{% hint style="info" %}
This server provides the Snyk documentation. It does not scan code or return findings from your Snyk account. To give an agent access to scanning, visit [Snyk Studio](https://docs.snyk.io/agent-security/agentic-security-with-snyk-studio/getting-started-with-snyk-studio).
{% endhint %}

## Documentation index

`llms.txt` lists every documentation page as a Markdown link with a one-line description, grouped by site section:

```
https://docs.snyk.io/llms.txt
```

Use it to discover what exists before fetching individual pages.

`llms-full.txt` contains the full text of the documentation in a single file:

```
https://docs.snyk.io/llms-full.txt
```

This file is large. For most agents, fetching individual Markdown pages costs fewer tokens than loading the whole corpus.

## Markdown pages

Append `.md` to any documentation URL to get the Markdown source instead of the rendered HTML page:

| Page | Markdown |
|---|---|
| `https://docs.snyk.io/whats-snyk` | `https://docs.snyk.io/whats-snyk.md` |
| `https://docs.snyk.io/developer-tools/snyk-cli` | `https://docs.snyk.io/developer-tools/snyk-cli.md` |

Each Markdown page opens with a blockquote pointing to `llms.txt` and to the Markdown version of that page.

## Content negotiation

To request Markdown without changing the URL, send an `Accept` header:

```bash
curl -H "Accept: text/markdown" https://docs.snyk.io/whats-snyk
```

The response has a `Content-Type` of `text/markdown; charset=utf-8`.

## Crawling the documentation

`https://docs.snyk.io/robots.txt` allows all user agents and declares `Content-Signal: ai-train=yes, search=yes, ai-input=yes`.

The sitemap lists every page with a last-modified date:

```
https://docs.snyk.io/sitemap.xml
```

## Report a documentation problem

If an agent finds documentation that is wrong, missing, or out of date, use the `sendFeedback` tool on the MCP server. The report goes to the Snyk documentation team.
