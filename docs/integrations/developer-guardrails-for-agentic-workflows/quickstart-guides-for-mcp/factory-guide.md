# Factory guide

You can add the Snyk MCP server to Factory to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install Factory](factory-guide.md#install-factory)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](factory-guide.md#install-the-snyk-mcp-server-in-amazon-q)

### Install Factory

Add the Factory extension to your IDE. For more details, see the [Factory official documentation](https://docs.factory.ai/user-guides/factory-bridge/installation-and-usage).&#x20;

### Install the Snyk MCP Server in Factory

Install the Snyk MCP Server using the method suited to your OS and environment.

#### Install with pre-installed Snyk CLI

Create or edit the MCP configuration file `~/Library/Application Support/Factory Bridge/mcp.json` .

```
{
  "mcpServers": { 
    "Snyk": {
    "command": "/absolute/path/to/snyk",
    "args": [
      "mcp",
      "-t",
      "stdio"
    ],
    "env": {}
  }
}
}
```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

#### Install with Node.js and `npx`

```
{
  "mcpServers": { 
    "Snyk": {
    "command": "npx",
    "args": [
      "-y",
      "snyk@latest",
      "mcp",
      "-t",
      "stdio"
    ],
    "env": {}
  }
}
}
```

The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image.png" alt=""><figcaption></figcaption></figure>

This json file can also be found by clicking on the Factory icon on the top toolbar and then selecting "Open MCP Config File".&#x20;

### Setting up the Snyk MCP Server <a href="#setting-up-the-snyk-mcp-server" id="setting-up-the-snyk-mcp-server"></a>

The MCP Server attempts to start automatically. The following example shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

As a one-time setup, authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to work with the Snyk MCP Server.

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Factory to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Factory indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

Factory is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found:

<figure><img src="../../../.gitbook/assets/image (215).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules <a href="#secure-at-inception-rules" id="secure-at-inception-rules"></a>

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [creating memories with Factory](https://docs.factory.ai/user-guides/memory/understanding-memory#adding-new-facts):

```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
