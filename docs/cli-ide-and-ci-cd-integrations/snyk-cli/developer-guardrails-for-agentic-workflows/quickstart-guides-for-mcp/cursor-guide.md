# Cursor guide

You can add the Snyk MCP server to Cursor to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](cursor-guide.md#install-cursor)
* [Install the Snyk CLI](../../install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](cursor-guide.md#install-the-snyk-mcp-server-in-cursor)

### Install Cursor

* Click [this link](cursor://anysphere.cursor-deeplink/mcp/install?name=snyk\&config=eyJjb21tYW5kIjoibnB4IC15IHNueWtAbGF0ZXN0IG1jcCAtdCBzdGRpbyJ9) to directly add the server to Cursor.
* Confirm installation in the Cursor settings by clicking on `Install`

<figure><img src="../../../../.gitbook/assets/image (104).png" alt=""><figcaption></figcaption></figure>

### Install the Snyk MCP Server in Cursor

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install using the Snyk extension

Navigate to the list of MCP servers on the Cursor website and search for Snyk. Then install by clicking the `Add Snyk to Cursor` button.

<div data-full-width="false"><figure><img src="../../../../.gitbook/assets/image (43).png" alt="" width="375"><figcaption></figcaption></figure></div>

Confirm the installation by clicking `Install` in the Cursor settings.

#### Install manually in Cursor

* Open Cursor settings
* Navigate to **Tools & Integrations**
* Select **Add Custom MCP**&#x20;
* Add the following json snippet to the file in use. For this example, the assumption is that  Snyk is on the system path. If this is not applicable, then use the absolute path to the Snyk CLI.

```json5
{
  "mcpServers": {
    "Snyk": {
      "command": "npx -y snyk@latest mcp -t stdio",
      "env": {}
    }
  }
}
```

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog.&#x20;

If you need to authenticate and trust the current directory, then proceed and complete the process.

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Windsurf to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Cursor indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

Cursor is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found.

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Cursor rules](https://docs.cursor.com/en/context/rules):

{% code title="Cursor rule" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Append the following contents to the rules file at `.cursor/rules`:

{% code title=".cursor/rules" overflow="wrap" %}
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
