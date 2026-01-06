# Augment Code guide

Add Snyk Studio to Augment Code to secure code generated with agentic workflows through a Large Language Model (LLM). You can achieve this in several ways. When you use it for the first time, Snyk Studio asks for trust and, if necessary, trigger authentication.

## Prerequisites

* [Install the code assistant extension](https://www.augmentcode.com/)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](augment-code-guide.md#install-the-snyk-mcp-server-in-augment-code)

### Install Augment Code

Visit the [Augment Code](https://www.augmentcode.com/) website to download the correct version of the IDE plugin.

### Install the Snyk MCP Server in Augment Code

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install with Node.js and npx

* Name: Snyk
* Command: npx -y snyk@latest mcp -t stdio

#### Install with pre-installed Snyk CLI

* Name: Snyk
* Command: /absolute/path/to/snyk mcp -t stdio

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

## Setting up the Snyk MCP Server

The MCP Server starts automatically. The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (80).png" alt=""><figcaption></figcaption></figure>

As a one-time setup, authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to work with the Snyk MCP Server.

## Examples <a href="#examples" id="examples"></a>

### Scanning for security vulnerabilities <a href="#scanning-for-security-vulnerabilities" id="scanning-for-security-vulnerabilities"></a>

In the chat, you can tag specific files or functions and ask Augment Code to scan your code for safe deployment to production.

```
Scan this directory for code security & dependency vulnerabilities and security issues
```

Augment Code indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (84).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for Augment Code user guidelines:

```
- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```

<figure><img src="../../../.gitbook/assets/image (86).png" alt=""><figcaption></figcaption></figure>
