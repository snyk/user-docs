# Cline guide

You can add the Snyk MCP server to Cline to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install Cline](cline-guide.md#install-cline)
* [Install the Snyk CLI](../../../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](cline-guide.md#install-the-snyk-mcp-server-in-cline-using-npx)

### Install Cline

Consult the official Cline documentation for [Cline installation instructions](https://docs.cline.bot/getting-started/installing-cline).

### Install the Snyk MCP Server in Cline using npx

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

Open the MCP Server configuration in Cline. Click the **Manage MCP Servers** icon at the bottom panel, then click the setup wheel icon at the top right.

<figure><img src="../../../.gitbook/assets/image (95).png" alt=""><figcaption></figcaption></figure>

Click the **Configure MCP Servers** button from the **MCP Servers** tab:

<figure><img src="../../../.gitbook/assets/image (96).png" alt=""><figcaption></figcaption></figure>

To install the Snyk MCP Server, add the following `mcpServers` configuration block to the Cline MCP Servers configuration file:

```
{
  "mcpServers": {
    "Snyk Security": {
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
    }
  }
}
```

### Install the Snyk MCP Server in Cline using Snyk CLI&#x20;

Ensure that the Snyk CLI is installed and that the `snyk` executable is available in your system's path.

To install the Snyk MCP Server with the Snyk CLI, follow these instructions:

*   Add the following `mcpServers` configuration block to the Cline MCP Servers configuration file:\


    ```
    {
      "mcpServers": {
        "Snyk Security": {
          "command": "/absolute/path/to/snyk",
          "args": ["mcp", "-t", "stdio"],
        }
      }
    }
    ```

{% hint style="info" %}
&#x20;If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the full path to the `snyk` executable. The path may look like this: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

After saving the MCP Servers configuration, confirm the Snyk MCP Server is successfully installed and listed in the Cline UI.

<figure><img src="../../../.gitbook/assets/image (97).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (89).png" alt=""><figcaption></figcaption></figure>

You can explicitly prompt to authenticate your account as follows:

```
Authenticate my Snyk account
```

<figure><img src="../../../.gitbook/assets/image (98).png" alt=""><figcaption></figcaption></figure>

After enabling Act mode and approving the request, the agent opens a browser for you to authenticate to the Snyk Web UI, then continues the session.

<figure><img src="../../../.gitbook/assets/image (91).png" alt="" width="375"><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Cline will invoke Snyk scanning, provide output and insights based on the results that Snyk returned:

<figure><img src="../../../.gitbook/assets/image (99).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Cline rules](https://docs.cline.bot/features/cline-rules):

{% code title="Cline rules" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Open the Rules and Workflows menu using the bottom chat navigation from Cline.

<figure><img src="../../../.gitbook/assets/image (100).png" alt=""><figcaption></figcaption></figure>

Add the following contents as a global rule:

{% code title="" overflow="wrap" %}
```
# Project security best practices

- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
