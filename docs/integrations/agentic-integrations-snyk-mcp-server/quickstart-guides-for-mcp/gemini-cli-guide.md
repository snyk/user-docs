# Gemini CLI guide

You can add the Snyk MCP server to Gemini CLI to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the Gemini CLI](gemini-cli-guide.md#install-gemini-cli)
* [Install the Snyk CLI](../../../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](gemini-cli-guide.md#install-the-snyk-mcp-server-in-gemini-cli-using-npx-and-settings.json)

### Install Gemini CLI

Consult the official Gemini documentation on GitHub for [Gemini CLI installation instructions](https://github.com/google-gemini/gemini-cli).

### Install the Snyk MCP Server in Gemini CLI using npx and settings.json

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

To install the Snyk MCP Server using the `.gemini/settings.json` Gemini file, proceed with the following instructions:

* Create or edit the `.gemini/settings.json` file (project-scoped configuration file for Gemini CLI)
*   Insert the following `mcpServers` configuration block to the file:\


    ```
    {
      "mcpServers": {
        "Snyk Security": {
          "command": "npx",
          "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
          "trust": true
        }
      }
    }
    ```
* Restart the Gemini CLI

### Install the Snyk MCP Server in Gemini CLI using Snyk CLI and settings.json

This installation instruction assumes you have installed the Snyk CLI and the `snyk` executable is available in your system's path.

To install the Snyk MCP Server using Gemini's own `.gemini/settings.json` file, proceed with the following instructions:

* Create or edit the `.gemini/settings.json` file (project-scoped configuration file for Gemini CLI)
*   Insert the following `mcpServers` configuration block to the file:\


    ```
    {
      "mcpServers": {
        "Snyk Security": {
          "command": "/absolute/path/to/snyk",
          "args": ["mcp", "-t", "stdio"],
          "trust": true
        }
      }
    }
    ```



{% hint style="info" %}
&#x20;If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the full path to the `snyk` executable. The path may look like this: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

* Restart the Gemini CLI

Use the Gemini CLI to list all MCP servers and verify that the Snyk MCP Server is installed and available by running the `mcp list` command.

<figure><img src="../../../.gitbook/assets/image (88).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (89).png" alt=""><figcaption></figcaption></figure>

You can explicitly prompt to authenticate your account as follows:

```
Authenticate my Snyk account
```

<figure><img src="../../../.gitbook/assets/image (90).png" alt=""><figcaption></figcaption></figure>

If you are not authenticated, this prompt opens a browser for you to authenticate to the Snyk Web UI, and then you can continue with the agent.

<figure><img src="../../../.gitbook/assets/image (91).png" alt=""><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Gemini will invoke Snyk scanning, provide output and insights based on the results that Snyk returned:

<figure><img src="../../../.gitbook/assets/image (92).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Gemini CLI's GEMINI.md context file](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/configuration.md#context-files-hierarchical-instructional-context):

{% code title="Gemini CLI rules" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Add the following contents to the `GEMINI.md` file in the root directory of the Project:

{% code title="GEMINI.md" overflow="wrap" %}
```
# Project security best practices

- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
