# Devin guide

You can add the Snyk MCP server to Devin to secure code generated with agentic workflows through an LLM. You can achieve this in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Get Started with Devin](devin-guide.md#get-started-with-devin)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](devin-guide.md#install-the-snyk-mcp-server-in-cursor)

### Get started with Devin

Navigate to [Devin.ai](https://devin.ai/) to get started using the AI software engineer.&#x20;

### Install the Snyk MCP Server in Devin

Install the Snyk MCP Server by navigating to **Settings** > **Organization settings** > **MCP marketplace** > **Add your own**.&#x20;

<figure><img src="../../../.gitbook/assets/image (535).png" alt=""><figcaption></figcaption></figure>

Devin uses [snapshots of virtual machines](https://docs.devin.ai/onboard-devin/repo-setup) to complete its tasks. Install the Snyk CLI on a virtual machine to enable the Snyk MCP Server to function properly.&#x20;

<figure><img src="../../../.gitbook/assets/image (541).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, it is possible that you need to authenticate and trust the current Project directory. If required, the agentic workflow will manage this automatically.

These workflows are automatically executed by the underlying model and the agentic code assistant, and you must approve them using a browser confirmation dialog.&#x20;

If you need to authenticate and trust the current directory, proceed and complete the process.

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Devin to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities security issues
```
{% endcode %}

Devin indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

Devin is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found.

<figure><img src="../../../.gitbook/assets/image (537).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (538).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows. The rules can be added to [Devin Knowledge](https://app.gitbook.com/u/vmSHyzn12dg748vWJFUXr9guzoj1).

The following is a suggested instruction for Devin rules:

{% code title="Devin rule" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

<figure><img src="../../../.gitbook/assets/image (539).png" alt=""><figcaption></figcaption></figure>
