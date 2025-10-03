# JetBrains AI assistant

You can add the Snyk MCP server to JetBrains AI Assistant to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](jetbrains-ai-assistant.md#install-jetbrains-ai-assistant)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](jetbrains-ai-assistant.md#install-the-snyk-mcp-server-in-jetbrains-ai-assistant)

### Install JetBrains AI Assistant

Add the JetBrains AI Assistant plugin to your JetBrains IDE (IntelliJ IDEA, WebStorm, PyCharm, and so on). For more details, see the official [JetBrains AI Assistant getting started page](https://www.jetbrains.com/help/ai-assistant/getting-started-with-ai-assistant.html).

Jetbrians AI Assistent tends to work more stably with the Snyk MCP service when using the [pre-installed CLI option](https://docs.google.com/document/d/1IfJ6dj6cEEvASoMKxg321IT3y2VkluP0juI_hHuDlJU/edit?tab=t.0#heading=h.nu8pk2tv8p88).

### Install the Snyk MCP Server in JetBrains AI Assistant

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install with pre-installed Snyk CLI

If you have the Snyk CLI installed, open **Settings/Preferences** > **Tools** > **AI Assistant** > **Model Context Protocol (MCP)**. Edit the MCP configuration and add the following instructions:

{% hint style="info" %}
* You might need to specify the full path to the Snyk executable CLI.
* JetBrains AI Assistant may have issues with the snyk\_trust function, needed before a dependency check (SCA scan). To fix this, add --disable-trust to the MCP command.
{% endhint %}

If you have the Snyk CLI installed and accessible on your system path add the following command.

<figure><img src="../../../.gitbook/assets/image2 (1).png" alt=""><figcaption></figcaption></figure>

If the snyk command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

Reload the IDE window for the changes to apply.

#### Install with Node.js and `npx`

Open Settings/Preferences ▸ Tools ▸ AI Assistant ▸ Model Context Protocol (MCP).

If you have the Node.js npx executable installed in your environment, add the following command:

<figure><img src="../../../.gitbook/assets/image1 (1).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
JetBrains AI Assistant may have issues with the snyk\_trust function, needed before a dependency check (SCA scan). To fix this, add --disable-trust to the MCP command.
{% endhint %}

Reload the IDE window for the changes to apply.

## Setting up the Snyk MCP Server

Verify that the Snyk MCP Server is available by checking the Status column in the MCP configuration and requesting it to list the tools it has available.

<figure><img src="../../../.gitbook/assets/image3 (2).png" alt=""><figcaption></figcaption></figure>

\
As a one-time setup, authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to work with the Snyk MCP Server.

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask JetBrains AI Assistant to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```
Scan this project for code security & dependency vulnerabilities and security issues with snyk
```
{% endcode %}

The AI Assistant indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

