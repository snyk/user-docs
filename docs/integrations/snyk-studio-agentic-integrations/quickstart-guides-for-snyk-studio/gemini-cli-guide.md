# Gemini CLI guide

You can access Snyk Studio, including Snyk's MCP server, in Gemini CLI to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, Snyk Studio may ask for trust and trigger authentication.

## Prerequisite

Install Gemini CLI. For more details, visit the [official Gemini CLI readme file.](https://github.com/google-gemini/gemini-cli?tab=readme-ov-file)

## Install Snyk Studio

Install Snyk Studio using the method that best suits your operating system and local development environment. Snyk recommends leveraging the ['single command install' using npx](gemini-cli-guide.md#install-with-npx). For other ways to install MCP servers in Gemini CLI, see Google's [official documentation](https://geminicli.com/docs/tools/mcp-server/).

### Install with `npx`

Open up a terminal window and paste the following command:

`npx -y snyk@latest mcp configure --tool=gemini-cli`&#x20;

This command:

* Downloads the latest version of Snyk's CLI.
* Set sup Snyk Studio within Gemini CLI.
* Configures Snyk Studio's Secure at inception directives within Gemini CLI's global rules file.

To verify installation, use the `/mcp list` command within Gemini CLI.

The list of tools installed as a part of Snyk Studio are listed below. These tools can be found listed underneath `Snyk` on the `/mcp list` screen.

| Tool                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_aibom`          | Generates an AI Bill of Materials (AI-BOM) for Python software projects in CycloneDX v1.6 JSON format. This experimental feature analyzes local Python projects to identify AI models, datasets, tools, and other AI-related components. Requires an active internet connection and access to the experimental feature (available on request). The command must be run within a Python project directory and requires the CLI from the preview release channel. |
| `snyk_auth`           | Authenticates the user with Snyk.                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `snyk_code_scan`      | <p>Performs Static Application Security Testing (SAST) directly from the Snyk MCP. This command analyzes an application's source code with a SAST scan to identify security vulnerabilities and weaknesses without executing the code. </p><p></p><p>Snyk supports: </p><p>Apex, C/C++, Dart and Flutter, Elixir, Go, Groovy, Java and Kotlin, Javascript, .NET, PHP, Python, Ruby, Rust, Scala, Swift and Objective-C, Typescript, VB.NET   </p>               |
| `snyk_container_scan` | Scans container images for known vulnerabilities in OS packages and application dependencies.                                                                                                                                                                                                                                                                                                                                                                   |
| `snyk_iac_scan`       | <p>Analyzes Infrastructure as Code (IaC) files for security misconfigurations. </p><p></p><p>Supports Terraform (.tf, .tf.json, plan files), Kubernetes (YAML, JSON), AWS CloudFormation (YAML, JSON), Azure Resource Manager (ARM JSON), and Serverless Framework.  </p>                                                                                                                                                                                       |
| `snyk_logout`         | Logs the Snyk MCP out of the current Snyk account by clearing the locally stored authentication token.                                                                                                                                                                                                                                                                                                                                                          |
| `snyk_sbom_scan`      | Experimental command. Analyzes an existing SBOM file for known vulnerabilities in its open-source components. Requires components in SBOM to be identified using PackageURLs (purls).                                                                                                                                                                                                                                                                           |
| `snyk_sca_scan`       | Analyzes projects for open-source vulnerabilities and license compliance issues by inspecting manifest files (e.g., package.json, pom.xml, requirements.txt) to understand dependencies and then queries the Snyk vulnerability database.                                                                                                                                                                                                                       |
| `snyk_send_feedback`  | Can be used to send feedback to Snyk if needed.                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `snyk_trust`          | Trust a given folder to allow Snyk to scan it.                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `snyk_version`        | Displays the installed Snyk MCP version.                                                                                                                                                                                                                                                                                                                                                                                                                        |

### Optional: Adjust scan frequency

Snyk recommends you use Snyk Studio with the Secure at inception directives, but also provides a smart scan option that allows the LLM to determine when to call Snyk Studio. This option results in lower overall token usage and faster iterating, but it increases the risk of insecure code being added to your codebase. Expand the options below for instructions on adjusting directives at installation or after installation.

<details>

<summary>Adjust at install</summary>

To utilize smart scan from install, add the following argument to the npx install command:

`npx -y snyk@latest mcp configure --tool=gemini-cli --rule-type=smart-apply`

</details>

<details>

<summary>Adjust after install</summary>

The default ruleset frequency can be adjusted by editing Gemini CLI's global `Gemini.md` file. For reference, the following are the smart apply rules Snyk places in Gemini CLI's global rules file when prompted:

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

### Alternative installation methods

Expand the relevant method below for installation instructions.

<details>

<summary>Install with Gemini CLI</summary>

Use the Gemini extension to install Snyk Studio. In your terminal, run:

```sh
gemini extensions install https://github.com/snyk/agentic-integration-wrappers
```

<figure><img src="../../../.gitbook/assets/image (425).png" alt="Gemini extensions install in the CLI"><figcaption><p>Example of using the Gemini install command in your terminal</p></figcaption></figure>

To confirm the extension was installed, run the following command in the Gemini CLI:

```
/extensions list
```

<figure><img src="../../../.gitbook/assets/image (382).png" alt=""><figcaption><p>Result of the /extensions list terminal command</p></figcaption></figure>

To install Snyk Studio manually, use the method that best suits your operating system and local development environment.

</details>

<details>

<summary>Install with Node.js and <code>npx</code></summary>

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

</details>

<details>

<summary>Install with Snyk CLI</summary>

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

</details>

#### Set up the Snyk MCP Server for alternate installations

As a one-time setup, you need to authenticate and trust the current Project directory. If required, the agentic workflow will likely manage this automatically.

You will not be prompted to authenticate if you already authenticated with the Snyk CLI installed on your system path, or through a prior interaction.

#### Set Gemini CLI Rules for alternate installations

To practice [Secure at inception](../../../discover-snyk/getting-started/glossary.md#secure-at-inception), Snyk recommends adopting directives that align the LLM for secure code generation and workflows.

Append one of the following contents to the context file at `~/.gemini/GEMINI.md`  &#x20;

{% code title="~/.gemini/GEMINI.md" overflow="wrap" %}
```
# Project security best practices

- Always run **snyk_code_scan** tool for new first party code that is generated in a Snyk-supported language.
- If any security issues are found based on newly introduced or modified code or dependencies, attempt to fix the issues using the results context from **Snyk**.
- Rescan the code after fixing the issues to ensure that the issues were fixed and that there are no newly introduced issues.
- Repeat this process until no new issues are found.
```
{% endcode %}

## Uninstall Snyk Studio

To remove Snyk Studio and the associated rulesets from your environment, use:

```
snyk mcp configure --rm --tool=gemini-cli
```
