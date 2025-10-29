# Gemini CLI guide

You can access Snyk Studio, including Snyk's MCP server, in Gemini CLI to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

Install [the Gemini CLI](https://github.com/google-gemini/gemini-cli?tab=readme-ov-file) to your local machine.&#x20;

### Install the Snyk Extension or MCP Server in Gemini CLI

Get started with Snyk and Gemini CLI by using the Gemini extension, which installs the Snyk MCP Server. Run in your terminal:

```sh
gemini extensions install https://github.com/snyk/agentic-integration-wrappers
```

<figure><img src="../../../.gitbook/assets/image (425).png" alt="Gemini extensions install in the CLI"><figcaption></figcaption></figure>

To confirm the extension was installed, run the following command in the Gemini CLI:

```
/extensions list
```

<figure><img src="../../../.gitbook/assets/image (382).png" alt=""><figcaption></figcaption></figure>

To install the Snyk MCP Server manually, use the method that best suits your operating system and local development environment.

#### Install with Node.js and `npx`

Create or edit the file `~/.gemini/settings.json` in the root directory of your Project.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

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

#### Install with pre-installed Snyk CLI

Create or edit the file `~/.gemini/settings.json` in the root directory of your Project.

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. Specify the full path to the Snyk executable CLI:

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

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

Run `/mcp list` to confirm the Snyk MCP Server is connected.

<figure><img src="../../../.gitbook/assets/image (383).png" alt=""><figcaption></figcaption></figure>

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

If you are already authenticated with the Snyk CLI installed on your system path, or through a prior interaction, then proceed to working with the Snyk MCP Server.

## Examples

### Scanning code and dependencies for security vulnerabilities

In the chat, you can tag specific files or functions and ask Gemini CLI to scan your code for safe deployment to production.

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Gemini CLI indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (384).png" alt=""><figcaption></figcaption></figure>

### Gemini CLI Rules

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

Append the following contents to the context file at `~/.gemini/GEMINI.md`

{% code title="~/.gemini/GEMINI.md" overflow="wrap" %}
```
# Project security best practices

- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}

{% code title="~/.gemini/GEMINI.md" overflow="wrap" %}
```
# Project security best practices

- Always run Snyk Code scanning tool for new first party code generated.
- Always run Snyk SCA scanning tool for new dependencies or dependency updates.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no issues are found.
```
{% endcode %}
