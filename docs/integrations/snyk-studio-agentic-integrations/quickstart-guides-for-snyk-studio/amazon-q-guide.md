# Amazon Q guide

You can access Snyk Studio, including Snyk's MCP server, in Amazon Q to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install the code assistant extension](amazon-q-guide.md#install-amazon-q)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](amazon-q-guide.md#install-the-snyk-mcp-server-in-the-amazon-q-ide-extension)

### Install Amazon Q

Add the Amazon Q extension to your IDE. For more details, see the official [Installing the Amazon Q Developer extension or plugin in your IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-setup.html).

### Install the Snyk MCP Server in the Amazon Q IDE extension

You can configure the Snyk MCP Server in VS Code and JetBrains for Amazon Q.

Using the AmazonQ extension UI in your chosen IDE:

1. Add a new MCP Server
2. Provide the following values in the specified fields:
   1. Command: `npx`
   2. Arguments: `-y snyk@latest mcp -t stdio -o=ostemp`
   3. Timeout: `0`

The `-o` option instructs the MCP server to write the scan results to a file. To direct the results to a specific folder, provide an absolute path instead of using `ostemp`. For example, `-o=/absolute/path/to/folder`

{% hint style="info" %}
For additional MCP configuration options on Amazon Q and troubleshooting, visit the official [Amazon Q MCP server documentation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/qdev-mcp.html).
{% endhint %}

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

These workflows are mostly automatically executed by the underlying model and the agentic code assistant, and you will need to approve them using a browser confirmation dialog. The process should look similar to this:

<figure><img src="../../../.gitbook/assets/image (341).png" alt=""><figcaption></figcaption></figure>

If you need to authenticate and trust the current directory, then proceed and complete the process.

<figure><img src="../../../.gitbook/assets/image (342).png" alt=""><figcaption></figcaption></figure>

## Examples

### Scanning for security vulnerabilities

In the chat, you can tag specific files or functions and ask Amazon Q to scan your code for safely deploying to production:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Amazon Q indicates that this request is related to security vulnerability scanning and calls the Snyk MCP Server for various scans.

<figure><img src="../../../.gitbook/assets/image (344).png" alt=""><figcaption></figcaption></figure>

Amazon Q is then able to gather all the security scanning results, both code (SAST) and dependencies (SCA), and based on recommendations and contextual information provided by the Snyk MCP Server, it is now able to provide code changes that address some or all of the security vulnerabilities that were found:

<figure><img src="../../../.gitbook/assets/image (345).png" alt=""><figcaption></figcaption></figure>

### "Secure at inception" rules

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for creating project rules for use with [Amazon Q Developer chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-project-rules.html). Append the following contents to the rules file at `.amazonq/rules/amazonq-rules.md`:

<figure><img src="../../../.gitbook/assets/image (346) (1).png" alt=""><figcaption></figcaption></figure>

{% code title=".amazonq/rules/amazonq-rules.md" overflow="wrap" %}
```
- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}
