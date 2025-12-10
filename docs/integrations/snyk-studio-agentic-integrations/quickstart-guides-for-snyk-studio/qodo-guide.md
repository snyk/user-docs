# Qodo guide

You can access Snyk Studio, including Snyk's MCP server, in Qodo to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](qodo-guide.md#install-qodo)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](qodo-guide.md#install-the-snyk-mcp-server-in-qodo)

### Install Qodo

Add the Qodo extension to the VS Code or JetBrains IDE. For more details, see the official [Qodo.ai](https://qodo.ai/) page.

### Install the Snyk MCP Server in Qodo

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

In the Qodo chatbox, click the **Add MCP** **tools** > **Add new MCP**. Provide the required details for the Snyk MCP server:

* The MCP Server name
* `command`: Either `npx` or the full path to the Snyk CLI.
* `args`: The Snyk MCP command to [start the MCP server](./#starting-the-snyk-mcp-server).

`npx` option:

```json
{
    "Snyk": {
        "command": "npx",
        "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"]
    }
}

```

full path option:

```json
{
    "Snyk": {
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"],
  }
}
```

For the full path option, if the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

The following example shows a Snyk MCP Server that was successfully configured.

<figure><img src="../../../.gitbook/assets/image (408).png" alt=""><figcaption></figcaption></figure>

For additional MCP configuration options on VS Code and troubleshooting, consult the official [VS Code MCP server documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (334).png" alt=""><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

<figure><img src="../../../.gitbook/assets/image (335).png" alt=""><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Qodo to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Qodo indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

Qodo is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found.

### "Secure at inception" rules

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Qodo best practices](https://docs.qodo.ai/qodo-documentation/qodo-gen/chat/chat-modes/standard-mode/commands/generate-best-practices). Append the following contents to the rules file at `best_practices.md`:

{% code title="best_practices.md" overflow="wrap" %}
```
---
applyTo: "**"
---

# Project security best practices

- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}

#### Qodo Workflows

Create a custom Snyk workflow to scan and remediate security vulnerabilities. The following is a suggestion for [Qodo workflows](https://docs.qodo.ai/qodo-documentation/qodo-gen/agent/workflows). Append the following contents to a Qodo agent file `remediate-snyk-studio.toml`:

```toml
# Version of the agent configuration standard
version = "1.0"

[commands.remediate]
available_tools = [
  "Snyk",
  "Code Navigation",
  "filesystem",
  "git",
  "Terminal",
  "Web Search"
]
description = "Remediate application security issues with Snyk Studio"
instructions = """
You are an expert at developing secure code. Snyk flagged several SAST and SCA issues in this application.

Please do the following:

1. Use Snyk to get context on the issues.
2. Fix all of the issues.
3. Use Snyk again to confirm the issues are fixed. If the issues are not fixed, go back to the previous step.
4. Build the app and run unit tests to ensure you didn’t break it. I’ll run it and do additional validation when you’re done.

Thanks!"""
mcpServers = """
{
  "Snyk": {
    "command": "npx",
    "type": "CUSTOM",
    "env": {},
    "args": [
      "-y",
      "snyk@latest",
      "mcp",
      "-t",
      "stdio"
    ]
  }
}"""
```
