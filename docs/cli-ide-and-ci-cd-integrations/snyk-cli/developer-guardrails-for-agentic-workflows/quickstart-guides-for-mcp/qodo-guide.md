# Qodo guide

{% include "../../../../.gitbook/includes/ai-assisted-coding-survey.md" %}

You can add the Snyk MCP server to Qodo to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](qodo-guide.md#install-qodo)
* [Install the Snyk CLI](../../install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](qodo-guide.md#install-the-snyk-mcp-server-in-qodo)

### Install Qodo

Add the Qodo extension to the VS Code or JetBrains IDE. For more details, see the official [Qodo.ai](https://qodo.ai/) page.

### Install the Snyk MCP Server in Qodo

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

In the Qodo chatbox, click the **Add MCP** **tools** > **Add new MCP**. Provide the required details for the Snyk MCP server:

* The MCP Server name
* The path to the Snyk CLI
* The Snyk MCP command to [start the MCP server](../snyk-mcp-early-access/snyk-mcp-installation-configuration-and-startup.md#starting-the-snyk-mcp-server).&#x20;
* You can see a list of all MCP Servers and their tool. Locate Snyk from the list and enable all of its tools.

#### Install with Node.js and `npx`

Create or edit the MCP configuration file `.vscode/mcp.json` in the root directory of your Project.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

```json5
{
  "servers": {
    "Snyk": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```

#### Install with pre-installed Snyk CLI

Create or edit the MCP configuration file `.vscode/mcp.json` in the root directory of your Project.

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. You might need to specify the full path to the Snyk executable CLI:

```json5
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

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../install-or-update-the-snyk-cli/) page.&#x20;

The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../../.gitbook/assets/image (73).png" alt=""><figcaption></figcaption></figure>

For additional MCP configuration options on VS Code and troubleshooting, consult the official [VS Code MCP server documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../../.gitbook/assets/image (74).png" alt=""><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

<figure><img src="../../../../.gitbook/assets/image (75).png" alt=""><figcaption></figcaption></figure>

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

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Qodo best practices](https://docs.qodo.ai/qodo-documentation/qodo-gen/chat/chat-modes/standard-mode/commands/generate-best-practices):

{% code title="Qodo rule" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Append the following contents to the rules file at `best_practices.md`:

{% code title="best_practices.md" overflow="wrap" %}
```
---
applyTo: "**"
---

# Project security best practices

- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
