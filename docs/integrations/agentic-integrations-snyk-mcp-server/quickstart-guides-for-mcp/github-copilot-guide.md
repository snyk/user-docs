# GitHub Copilot guide

{% include "../../../.gitbook/includes/ai-assisted-coding-survey.md" %}

You can add the Snyk MCP server to GitHub Copilot to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](github-copilot-guide.md#install-github-copilot)
* [Install the Snyk CLI](../../../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](github-copilot-guide.md#install-the-snyk-mcp-server-in-github-copilot)

### Install GitHub Copilot

Add the GitHub Copilot extension to VS Code. For mode details, see the official [Setup GitHub Copilot on VS Code guide](https://code.visualstudio.com/docs/copilot/setup).

### Install the Snyk MCP Server in GitHub Copilot

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install using the Snyk extension (preferred)

Installing the extension using one of the following methods:

* View the Snyk Security extension for VS Code from the [marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner) and install it.
*   Open the **Extensions: Install Extensions** side-pane, search for Snyk Security and install it.\


    <figure><img src="../../../.gitbook/assets/image (72).png" alt=""><figcaption></figcaption></figure>

VS Code automatically detects the Snyk MCP Server, but it needs to be explicitly enabled.

Enable the Snyk MCP Server using one of the following methods:

* Open the Command Palette by pressing **CMD + SHIFT + P** on macOS or **CTRL + SHIFT + P** on Windows, choose the **MCP: List Servers,** and then find Snyk MCP Server in the list and enable all tools.
* In the GitHub Copilot chat box, click the **Tools** icon.

![](<../../../.gitbook/assets/image (76).png>)\
\
You can see a list of all MCP Servers and their tool. Locate Snyk from the list and enable all of its tools:\
![](<../../../.gitbook/assets/image (77).png>)

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

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/) page.&#x20;

The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (73).png" alt=""><figcaption></figcaption></figure>

For additional MCP configuration options on VS Code and troubleshooting, consult the official [VS Code MCP server documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (74).png" alt=""><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

<figure><img src="../../../.gitbook/assets/image (75).png" alt=""><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Copilot to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Copilot indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (78).png" alt=""><figcaption></figcaption></figure>

Copilot is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found:

<figure><img src="../../../.gitbook/assets/image (79).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [GitHub Copilot custom instructions](https://docs.github.com/en/copilot/how-tos/configure-custom-instructions/add-repository-instructions):

{% code title="GitHub Copilot rule" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Append the following contents to the rules file at `.github/copilot-instructions.md`:

{% code title=".github/copilot-instructions.md" overflow="wrap" %}
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
