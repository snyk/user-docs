# Goose CLI guide

You can add the Snyk MCP server to Goose CLI to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the Goose CLI](gemini-cli-guide-1.md#install-goosecli)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](gemini-cli-guide-1.md#install-the-snyk-mcp-server-in-gemini-cli-using-npx-and-settings.json)

### Install GooseCLI

Install the GooseCLI. For mode details, see the official [Goose CLI installation instructions.](https://block.github.io/goose/docs/quickstart)

### Install the Snyk MCP Server in Goose CLI using npx command

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

To install the Snyk MCP Server using the Goose CLI interactive session, proceed with the following instructions:

*   Start an interactive Goose configuration wizard by running the `goose configure` command:\


    <figure><img src="../../../.gitbook/assets/image (396) (1).png" alt=""><figcaption></figcaption></figure>
* From the configuration menu, chose **Add Extension.**
* Choose **Command-line extension (run a local command or script).**
* Name this extension **Snyk Security.**
* Run the command by typing in: `npx -y snyk@latest mcp -t stdio`
* Continue with the defaults (timeout of 300 seconds), and optionally provide the Snyk API token if your Organization policy requires one.&#x20;

A complete Goose configuration walkthrough should look as follows:

<figure><img src="../../../.gitbook/assets/image (397).png" alt=""><figcaption></figcaption></figure>

Start or resume your Goose session with the Snyk MCP server enabled.

### Install the Snyk MCP Server in Goose CLI using Snyk CLI

This installation instruction assumes you have installed the Snyk CLI and the `snyk` executable is available in your system's path.

Install the Snyk MCP Server using the Goose CLI interactive session by following these instructions:

* Start an interactive Goose configuration wizard by running the `goose configure` command.
* Choose **Add Extension f**rom the configuration menu.
* Select **Command-line extension (run a local command or script).**
* Name this extension **Snyk Security.**
* Run the command by typing in: `/path/to/snyk mcp -t stdio` - replacing `/path/to/snyk` with the fully qualified absolute path to your Snyk CLI.
* Continue with the defaults (timeout of 300 seconds), and optionally provide the Snyk API token if your Organization policy requires one.&#x20;

{% hint style="info" %}
&#x20;If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the full path to the `snyk` executable. The path may look like this: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

Start or resume your Goose session with the Snyk MCP server enabled.

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog.&#x20;

You can explicitly prompt to authenticate your account as follows:

```
Authenticate my Snyk account
```

<figure><img src="../../../.gitbook/assets/image (398).png" alt=""><figcaption></figcaption></figure>

If you are not authenticated, this prompt opens a browser for you to authenticate to the Snyk Web UI, and then you can continue with the agent.

<figure><img src="../../../.gitbook/assets/image (486).png" alt="" width="375"><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Goose CLI invokes Snyk scanning, providing output, and insights based on the results that Snyk returned:

<figure><img src="../../../.gitbook/assets/image (399).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction to use with [Goose CLI recipes](https://block.github.io/goose/docs/guides/recipes/session-recipes) or with [Goose hints file](https://block.github.io/goose/docs/guides/using-goosehints)

{% code title="Goose CLI rules" overflow="wrap" %}
```
Always run Snyk Code scanning tool for new first party code generated.
Always run Snyk SCA scanning tool for new dependencies or dependency updates.
If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
Repeat this process until no issues are found.
```
{% endcode %}

Add the following contents to the `.goosehints` file in the root directory of the Project:

{% code title=".goosehints" overflow="wrap" %}
```
# Project security best practices

- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
