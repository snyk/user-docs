# JetBrains Junie

You can access Snyk Studio, including Snyk's MCP server, in JetBrains Junie to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](jetbrains-junie.md#install-jetbrains-junie)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](jetbrains-junie.md#install-the-snyk-mcp-server-in-jetbrains-junie)

### Install JetBrains Junie

Add the JetBrains Junie plugin to your JetBrains IDE (IntelliJ IDEA, WebStorm, PyCharm, and so on). For more details, see the official [JetBrains Junie getting started page](https://www.jetbrains.com/help/junie/get-started-with-junie.html).

### Install the Snyk MCP Server in JetBrains Junie

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install with Node.js and `npx`

If you have the Node.js npx executable installed, open **Settings/Preferences** > **Tools** > **Junie** > **MCP Settings**. Edit the MCP configuration and add the following instructions:

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

<figure><img src="../../../.gitbook/assets/image3 (3).png" alt=""><figcaption></figcaption></figure>

#### Install with pre-installed Snyk CLI

If you have the Snyk CLI installed, open **Settings/Preferences** > **Tools** > **Junie** > **MCP Settings**. Edit the MCP configuration and add the following instructions:

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

If the snyk command is not available, add it by following the instructions on the Installing or updating the Snyk CLI page.

Reload the IDE window for the changes to apply.

## Setting up the Snyk MCP Server

Verify that the Snyk MCP Server is available by checking the Status column in the MCP configuration and requesting it to list the tools it has available.

<figure><img src="../../../.gitbook/assets/image1 (2).png" alt=""><figcaption></figcaption></figure>

\
As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to working with the Snyk MCP Server.

## Examples

### Scanning for security vulnerabilities

In the chat, you can ask Junie to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Junie indicates in the plan that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image2 (2).png" alt=""><figcaption></figcaption></figure>
