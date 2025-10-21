# Kiro guide

You can add the Snyk MCP server to Kiro to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](kiro-guide.md#install-kiro)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](kiro-guide.md#install-the-snyk-mcp-server-in-kiro)

### Install Kiro

Install Kiro on your machine. For more details, visit the official [Kiro](https://kiro.dev/) page.

### Install the Snyk MCP Server in Kiro

Install the Snyk MCP Server using the method suited to your OS and environment.

Create or edit the MCP configuration file `~/.kiro/settings/mcp.json` .

#### Install with Node.js and `npx`

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

```json5
{
  "mcpServers": {
    "Snyk": {
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```

#### Install with pre-installed Snyk CLI

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. You might need to specify the full path to the Snyk executable CLI:

```json5
{
  "mcpServers": {
    "Snyk": {
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.&#x20;

The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (397).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (398).png" alt=""><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

## Examples

### Scanning for security vulnerabilities

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Kiro indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

Kiro is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found.

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Kiro Agent Steering](https://kiro.dev/docs/steering/):

{% code title="Kiro Agent Steering" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Append the following contents to an Agent Steering file

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
