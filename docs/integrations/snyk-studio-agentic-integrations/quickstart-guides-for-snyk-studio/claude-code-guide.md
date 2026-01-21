# Claude Code guide

You can access Snyk Studio in Claude Code to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisite

Install Claude Code. For more details, visit the official [Claude Code - Quickstart](https://docs.anthropic.com/en/docs/claude-code/quickstart).

## Install Snyk Studio

Install Snyk Studio using the method that best suits your operating system and local development environment. Snyk recommends leveraging the ['single command install' using `npx`](claude-code-guide.md#install-with-npx) . For different ways to install MCP servers into Claude Code, see Anthropic's [official documentation](https://docs.anthropic.com/en/docs/claude-code/mcp#installing-mcp-servers).

### Install with `npx`

Open a terminal window and paste the following command:

`npx -y snyk@latest mcp configure --tool=claude-cli`&#x20;

This command:

* Downloads the latest version of Snyk's CLI.
* Sets up Snyk Studio within Claude Code.
* Configures Snyk Studio's Secure at inception directives within Claude Code's global rules file.

<figure><img src="../../../.gitbook/assets/Screenshot 2026-01-09 at 3.23.45 PM.png" alt=""><figcaption></figcaption></figure>

To verify installation, use the `/mcp` command within Claude:

<figure><img src="../../../.gitbook/assets/Screenshot 2026-01-09 at 3.26.38 PM.png" alt=""><figcaption></figcaption></figure>

Select **View Tools** to look at all of the commands and tooling Snyk utilizes as part of its execution The descriptions also include instructions specific for the LLM. These are capitalized to help you differentiate. These tools include:

| Tool                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_aibom`          | Generates an AI Bill of Materials (AI-BOM) for Python software projects in CycloneDX v1.6 JSON format. This experimental feature analyzes local Python projects to identify AI models, datasets, tools, and other AI-related components. Requires an active internet connection and access to the experimental feature (available on request). The command must be run from within a Python project directory and requires the CLI from the preview release channel. |
| `snyk_auth`           | Authenticate the user with Snyk.                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `snyk_code_scan`      | <p>Performs Static Application Security Testing (SAST) directly from the Snyk MCP. It analyzes an application's source code with a SAST scan to identify security vulnerabilities and weaknesses without executing the code. </p><p></p><p>Supported languages: </p><p>Apex, C/C++, Dart and Flutter, Elixir, Go, Groovy, Java and Kotlin, Javascript, .NET, PHP, Python, Ruby, Rust, Scala, Swift and Objective-C, Typescript, VB.NET   </p>                        |
| `snyk_container_scan` | Scans container images for known vulnerabilities in OS packages and application dependencies.                                                                                                                                                                                                                                                                                                                                                                        |
| `snyk_iac_scan`       | <p>Analyzes Infrastructure as Code (IaC) files for security misconfigurations. </p><p></p><p>Supports Terraform (.tf, .tf.json, plan files), Kubernetes (YAML, JSON), AWS CloudFormation (YAML, JSON), Azure Resource Manager (ARM JSON), and Serverless Framework.  </p>                                                                                                                                                                                            |
| `snyk_logout`         | Logs the Snyk MCP out of the current Snyk account by clearing the locally stored authentication token.                                                                                                                                                                                                                                                                                                                                                               |
| `snyk_sbom_scan`      | Experimental command. Analyzes an existing SBOM file for known vulnerabilities in its open-source components. Requires components in SBOM to be identified using PackageURLs (purls).                                                                                                                                                                                                                                                                                |
| `snyk_sca_scan`       | Analyzes Projects for open-source vulnerabilities and license compliance issues by inspecting manifest files (for example package.json, pom.xml, requirements.txt) to understand dependencies and then queries the Snyk vulnerability database.                                                                                                                                                                                                                      |
| `snyk_send_feedback`  | Can be used to send feedback to Snyk.                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `snyk_trust`          | Trusts a given folder to allow Snyk to scan it.                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `snyk_version`        | Displays the installed Snyk MCP version.                                                                                                                                                                                                                                                                                                                                                                                                                             |

### Optional: Adjust scan frequency

Snyk recommends you use Snyk Studio with the Secure at inception directives, but also provides a smart scan option that allows the LLM to determine when to call Snyk Studio. This option results in lower overall token usage and faster iterating, but it increases the risk of insecure code being added to your codebase. Expand the options below for instructions on adjusting directives at installation or after installation.

<details>

<summary>Adjust at install</summary>

To utilize smart-scan from install, add the following argument to the npx install command:

`npx -y snyk@latest mcp configure --tool=claude-cli --rule-type=smart-apply`

</details>

<details>

<summary>Adjust after install</summary>

The default ruleset frequency can be adjusted by editing the global `CLAUDE.md` file.

&#x20;For reference, the following are the smart apply rules Snyk places in Claude's global rules file when prompted:

{% code overflow="wrap" %}
```
BEFORE declaring task complete: Run snyk_code_scan tool when a significant change has been made in first party code.
- This should only apply for Snyk-supported coding language
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}

</details>

### Alternate installation methods

Expand the relevant method below for installation instructions.

<details>

<summary>Install with Node.js and <code>npx</code></summary>

Create or edit the MCP configuration file `~/.claude.json`.

If you have the Node.js `npx` executable installed in your environment, add the following JSON snippet to the file:

<pre><code><strong>{
</strong>  "mcpServers": {
    "Snyk": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "snyk@latest", "mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
</code></pre>

</details>

<details>

<summary>Install with Snyk CLI</summary>

Create or edit the MCP configuration file `~/.claude.json` .

If you have the Snyk CLI installed and accessible on your system path, include the following JSON snippet in the file. You might need to specify the full path to the Snyk executable CLI:

```
{
  "mcpServers": {
    "Snyk": {
      "type": "stdio",
      "command": "/absolute/path/to/snyk",
      "args": ["mcp", "-t", "stdio"],
      "env": {}
    }
  }
}
```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

</details>

<details>

<summary>Install with Claude Code CLI commands</summary>

Run the Snyk MCP Server in `sse` transport mode using the Snyk CLI:

```
snyk mcp -t sse 
```

Then run the Claude Code CLI for adding a new MCP server:

```
claude mcp add --transport sse snyk http://127.0.0.1:7695/sse
```

Watch this step-by-step [YouTube video](https://www.youtube.com/watch?v=qAcTPMXbfyU). showing how to configure the Snyk MCP Server with Claude Code CLI.

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

The following examples shows a Snyk MCP Server that was successfully configured and started.

<figure><img src="../../../.gitbook/assets/image (372).png" alt="" width="563"><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (373).png" alt="" width="563"><figcaption></figcaption></figure>

For additional MCP configuration options on Claude Code and troubleshooting, visit the official [Claude Code MCP documentation](https://docs.anthropic.com/en/docs/claude-code/mcp).

</details>

### Authenticate the Snyk MCP Server

As a one-time setup, you need to authenticate and trust the current Project directory.

<figure><img src="../../../.gitbook/assets/image (374).png" alt=""><figcaption></figcaption></figure>

## Uninstall Snyk Studio

To remove Snyk Studio and its associated ruleset from your environment, use:

```
snyk mcp configure --rm --tool=claude-cli
```

## Secure at inception directives

To practice [Secure at Inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting rules that align the LLM for secure code generation and workflows.

The following is a suggested instruction for creating project rules for use with [Claude Code md files](https://www.anthropic.com/engineering/claude-code-best-practices). Append the following contents to the rules file at `CLAUDE.md`:

<figure><img src="../../../.gitbook/assets/image (381) (1).png" alt=""><figcaption></figcaption></figure>

{% code title="CLAUDE.md" overflow="wrap" %}
```
- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}
