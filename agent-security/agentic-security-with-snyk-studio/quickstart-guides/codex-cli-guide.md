# Codex CLI guide

Access Snyk Studio in Codex CLI to ensure generated code is secure at inception (SAI). Snyk Studio supports a default hooks-based approach to SAI for Codex CLI.

## Install Snyk Studio with hooks support

For installation instructions, visit [Getting started with Snyk Studio](../getting-started-with-snyk-studio.md#install-snyk-studio-for-ades-with-hooks-support).

## Install Snyk Studio (legacy rules-based approach)

### Prerequisites

* [Install the Codex CLI](codex-cli-guide.md#install-codex-cli)
* [Install the Snyk CLI](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/install-the-snyk-cli)
* [Install the Snyk MCP](codex-cli-guide.md#install-the-snyk-mcp-server-in-codex-cli-using-npx)

### Install Codex CLI

Consult the official OpenAI Codex CLI documentation on GitHub for [Codex CLI installation instructions](https://github.com/openai/codex).

### Install the Snyk MCP Server in Codex CLI using npx

This installation instruction assumes you have a Node.js local development environment setup with the `npx` executable.

To install the Snyk MCP Server using Codex's own `.codex/config.toml` file, proceed with the following instructions:

* Create or edit the `.codex/config.toml` file
*   Insert the following `mcp_servers` configuration block to the file:\\

    ```
    [mcp_servers.snyk-security]
    command = "npx"
    args = ["-y", "snyk@latest", "mcp", "-t", "stdio"]
    ```
* Restart the Codex CLI

### Install the Snyk MCP Server in Codex CLI using Snyk CLI

This installation instruction assumes you have installed the Snyk CLI and the `snyk` executable is available in your system's path.

To install the Snyk MCP Server using the `.codex/config.toml` Codex file, follow these instructions:

* Create or edit the `.codex/config.toml` file
*   Insert the following `mcpServers` configuration block to the file:\\

    ```
    [mcp_servers.snyk-security]
    command = "/absolute/path/to/snyk"
    args = ["mcp", "-t", "stdio"]
    ```

{% hint style="info" %}
If you installed `snyk` using npm as a global module and you are using fnm to manage your versions, then you must specify the full path to the `snyk` executable. The path may look like this: `/Users/username/.local/state/fnm_multishells/8831_1753881185071/bin/snyk`.
{% endhint %}

* Restart the Codex CLI

## Set up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../.gitbook/assets/image (9).png" alt=""><figcaption></figcaption></figure>

You can explicitly prompt to authenticate your account as follows:

```
Authenticate my Snyk account
```

<figure><img src="../../.gitbook/assets/image (13).png" alt=""><figcaption></figcaption></figure>

If you are not authenticated, this prompt opens a browser for you to authenticate to the Snyk Web UI, and then you can continue with the agent.

<figure><img src="../../.gitbook/assets/image (10).png" alt="" width="375"><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask the agent to scan your code and third-party dependencies of the Project.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Codex CLI will invoke Snyk scanning, provide output and insights based on the results that Snyk returned:

<figure><img src="../../.gitbook/assets/image (14).png" alt=""><figcaption></figcaption></figure>

### Secure at inception rules

To practice [Secure at Inception](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/glossary#secure-at-inception-directive), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for [Codex CLI's AGENTS.md context file](https://github.com/openai/codex?tab=readme-ov-file#memory--project-docs). Add the following contents to the `AGENTS.md` file in your root directory of your Project:

{% code title="AGENTS.md" overflow="wrap" %}
```
# Project security best practices

- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}
