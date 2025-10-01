# Claude Code guide

You can add the Snyk MCP server to Claude Code to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install Claude Code](claude-code-guide.md#install-claude-code)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](claude-code-guide.md#install-the-snyk-mcp-server-in-claude-code)

### Install Claude Code

Install Claude Code. For more details, see the official [Claude Code - Quickstart](https://docs.anthropic.com/en/docs/claude-code/quickstart).

### Install the Snyk MCP Server in Claude Code

Install the Snyk MCP Server using the method that best suits your operating system and local development environment. To add a MCP Server in Claude Code, see the official document for [installing MCP servers](https://docs.anthropic.com/en/docs/claude-code/mcp#installing-mcp-servers).

#### Install with Node.js and `npx`

Create or edit the MCP configuration file `~/.claude.json`.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

<pre><code><strong>{
</strong>  "servers": {
    "Snyk": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
</code></pre>

#### Install with pre-installed Snyk CLI

Create or edit the MCP configuration file `~/.claude.json` .

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. You might need to specify the full path to the Snyk executable CLI:

```
{
  "servers": {
    "Snyk": {
      "type": "stdio",
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.&#x20;

The following examples shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (512).png" alt="" width="563"><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (513).png" alt="" width="563"><figcaption></figcaption></figure>

For additional MCP configuration options on Claude Code and troubleshooting, consult the official [Claude Code MCP documentation](https://docs.anthropic.com/en/docs/claude-code/mcp).

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory.&#x20;

<figure><img src="../../../.gitbook/assets/image (514).png" alt=""><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Claude Code to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Claude Code indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (516).png" alt=""><figcaption></figcaption></figure>

Claude Code is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found:

<figure><img src="../../../.gitbook/assets/image (517).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for creating project rules for use with [Claude Code md files](https://www.anthropic.com/engineering/claude-code-best-practices):

{% code title="Claude Code rules" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Append the following contents to the rules file at `CLAUDE.md`:

<figure><img src="../../../.gitbook/assets/image (518).png" alt=""><figcaption></figcaption></figure>

{% code title="CLAUDE.md" overflow="wrap" %}
```
- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
