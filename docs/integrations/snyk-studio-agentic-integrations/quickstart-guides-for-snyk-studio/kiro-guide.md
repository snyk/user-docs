# Kiro guide

You can access Snyk Studio, including Snyk's MCP server, in Kiro to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisite

Install Kiro on your machine. For more details, visit the official [Kiro](https://kiro.dev/) page.

## Recommended installation method

Snyk Studio provides Amazon Kiro users with a unique, customized power to bypass manual installation. The power will:

* Install the latest version of the Snyk CLI (and keep it up to date).
* Install the Snyk Studio MCP server along with the custom ruleset to scan code immediately after generation.

### Add Snyk Studio using Kiro Powers

Navigate to the Powers panel in your Kiro agentic development environment.

<figure><img src="../../../.gitbook/assets/image (1).png" alt="View of the Powers panel in Kiro"><figcaption><p>View of the Powers panel in Kiro</p></figcaption></figure>

Click **Add Custom Power**. When the pop-up appears, select **Import Power from GitHub**.

<figure><img src="../../../.gitbook/assets/image (1) (3).png" alt=""><figcaption><p>Add a Custom Power in Kiro by importing from GitHub</p></figcaption></figure>

When prompted, copy and paste the following URL path into the window:

```
https://github.com/snyk/kiro-power-snyk/tree/main/snyk-power
```

You should now see the Snyk Studio custom **Snyk Secure at Inception** power in the list of installed powers on the left hand side of your window.

<figure><img src="../../../.gitbook/assets/image (3).png" alt="View of the Secure at Inception Power in the Installed Powers list"><figcaption><p>View of the Secure at Inception Power in the Installed Powers list</p></figcaption></figure>

To verify that the Snyk Studio MCP server is running, click the ghost icon above the Kiro powers icon. You should see the Snyk MCP server in the bottom left list of MCP servers with a checkmark, indicating it is properly installed and functioning.

<figure><img src="../../../.gitbook/assets/image (4).png" alt="Verification that the Secure at Inception Power is installed and functioning"><figcaption><p>Verification that the Secure at Inception Power is installed and functioning</p></figcaption></figure>

## Alternative manual installation

Manually Install Snyk Studio's MCP Server using the method suited to your OS and environment.

To do this, you need to create or edit the MCP configuration file `~/.kiro/settings/mcp.json` .

### Install with Node.js and `npx`

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

### Install with the Snyk CLI

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

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

The following example shows a Snyk MCP Server that was successfully configured and started:

<figure><img src="../../../.gitbook/assets/image (397).png" alt=""><figcaption></figcaption></figure>

## Authenticate and trust the Project directory

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

### Secure at inception directive

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows. If you installed Snyk Studio via the Kiro power, this step was already done for you.

The following is a suggested instruction for [Kiro Agent Steering](https://kiro.dev/docs/steering/). Append the following contents to an Agent Steering file:

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
