# Gemini Code Assist guide

You can access Snyk Studio, including Snyk's MCP server, in Gemini Code Assist to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](gemini-code-assist-guide.md#install-gemini-code-assist)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](gemini-code-assist-guide.md#install-the-snyk-mcp-server-in-gemini-code-assist)

### Install Gemini Code Assist

Add the Gemini Code Assist extension to the VS Code or JetBrains IDE. For more details, see the official [Gemini Code Assist getting started page](https://codeassist.google/).

{% hint style="info" %}
MCP support is available only when using the agent preview mode. To enable the agent preview mode on VS Code, you have to open the settings.json VS Code user file or global settings file and enable the Insiders channel by adding the `"geminicodeassist.updateChannel": "Insiders",` line.
{% endhint %}

### Install the Snyk MCP Server in Gemini Code Assist

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install with Node.js and `npx`

Create or edit the file `.gemini/settings.json` in the root directory of your Project.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

```json5
{
  "mcpServers": {
    "snyk": {
      "command": "npx",
      "args": [
        "-y",
        "snyk@latest",
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}
```

#### Install with pre-installed Snyk CLI

Create or edit the file `.gemini/settings.json` in the root directory of your Project.

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. You might need to specify the full path to the Snyk executable CLI:

```json5
{
  "mcpServers": {
    "snyk": {
      "command": "/absolute/path/to/snyk",
      "args": [
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}

```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

You might need to reload the IDE window for the changes to apply. In VS Code, do this by opening the Command Palette and choosing Developer: Reload Window.

Next, you can verify that the Snyk MCP Server is available by enabling Agent mode on Gemini Code Assist and requesting it to list the tools it has available:

<figure><img src="../../../.gitbook/assets/image (359).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to working with the Snyk MCP Server.

## Examples

### Scanning code and dependencies for security vulnerabilities

In the chat, you can tag specific files or functions and ask Gemini to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Gemini indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (360).png" alt=""><figcaption></figcaption></figure>

### Gemini Code Assist security rules

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for Gemini Code Assist rules. Following the [documentation proposal for Gemini Code Assist style guide](https://developers.google.com/gemini-code-assist/docs/customize-gemini-behavior-github), append the following contents to the rules file at `.gemini/styleguide.md`:

{% code title=".gemini/styleguide.md" overflow="wrap" %}
```

# Project security best practices

- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}

#### Add Gemini Code Assist security rules using the Settings UI

You can add these rules using VS Code **UI Settings** and locate the **Geminicodeassist: Rules** configuration:

<figure><img src="../../../.gitbook/assets/image (361).png" alt=""><figcaption></figcaption></figure>
