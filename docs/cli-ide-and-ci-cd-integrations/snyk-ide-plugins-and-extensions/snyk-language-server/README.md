# Snyk Language Server

Snyk offers IDE integrations that allow you to use the functionality of Snyk in your Integrated Development Environment or Editor. This page describes the Snyk Language Server that can provide diagnostics for any IDE or Editor that supports the [Language Server Protocol](https://microsoft.github.io/language-server-protocol/). For information about all of the IDE plugins and extensions and their use, see [Snyk IDE plugins and extensions](../).

The Snyk Language Server scans for vulnerabilities, open-source license issues, and infrastructure misconfigurations and returns results with security issues categorized by issue type and severity.

For open source, you receive automated, algorithm-based fix suggestions for both direct and transitive dependencies.

Snyk Language Server scans for the following types of issues:

* [**Open Source Security**](https://snyk.io/product/open-source-security-management/) - security vulnerabilities and license issues in both the direct and indirect (transitive) open-source dependencies pulled into the Snyk Project. See also the [Open Source documentation](../../../scan-with-snyk/snyk-open-source/).
* [**Code Security**](https://snyk.io/product/snyk-code/) - security vulnerabilities in your code. See also the [Snyk Code documentation](../../../scan-with-snyk/snyk-code/).
* [**Infrastructure as Code (IaC) Security**](https://snyk.io/product/infrastructure-as-code-security/) - configuration issues in your IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager. See also the [Snyk Infrastructure as Code documentation](../../../scan-with-snyk/snyk-iac/scan-your-iac-source-code/).

After you have installed and configured the Language Server, every time you run it, open a file, or save, Snyk scans the manifest files, proprietary code, and configuration files in your Project. Snyk delivers actionable vulnerability, license, or misconfiguration issue details and displays the results natively within the LSP supporting Editor or IDE.

This page explains supported environments, support, and giving feedback, and provides installation instructions.

## Supported operating systems and architecture

{% hint style="warning" %}
Snyk plugins are not supported on any operating system that has reached End Of Life (EOL) with the distributor.
{% endhint %}

You can use the Language Server in the following environments:

* Linux: AMD64 and ARM64
* Linux Alpine: 386 and AMD64
* Windows: 386, AMD64, ARM through 386 compatibility
* MacOS: AMD64 and ARM64

## Where you can download the Language Server (Snyk CLI)

Snyk Language Server is nowadays included in the Snyk CLI. The CLI is automatically downloaded only when you use the Snyk IDE plugins.&#x20;

Please refer to [snyk-cli](../../snyk-cli/ "mention") for installation and manual download instructions.

## Usage of Snyk Language Server

### Starting the Snyk CLI in language server mode

```
// we assume the CLI is in the system path and named `snyk`
snyk language-server <flags>

// debug logging
snyk language-server -d

// trace logging
SNYK_LOG_LEVEL=trace snyk language-server
```

### Snyk LSP command line flags

`-d`  output debug level logs

`-c <FILE>` allows specifying a config file to load before all others.

`-l <LOGLEVEL>` allows specifying the log level (`trace`, `debug`, `info`, `warn`, `error`, `fatal`). The default log level is `info.`

`-o <FORMAT>` allows specifying the output format (`md` or `html`) for issues.

`-f <FILE>` allows specifying a log file instead of logging to the console.

`-licenses` displays the [licenses](https://github.com/snyk/snyk-ls/tree/main/licenses) used by Language Server.

### **LSP initialization options**

As part of the [Initialize message](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize) within `initializationOptions?: LSPAny;` Snyk supports the following settings:

```json
{
  "activateSnykOpenSource": "true", // Enables Snyk Open Source - defaults to true
  "activateSnykCode": "false", // Enables Snyk Code, if enabled for your organization - defaults to false
  "activateSnykIac":  "true", // Enables Infrastructure as Code - defaults to true
  "insecure": "false", // Allows custom CAs (Certification Authorities)
  "endpoint":  "https://example.com", // Snyk API Endpoint required for non-default multi-tenant and single-tenant setups
  "additionalParams": "--all-projects", // Any extra params for the Snyk CLI, separated by spaces
  "additionalEnv":  "MAVEN_OPTS=-Djava.awt.headless=true;FOO=BAR", // Additional environment variables, separated by semicolons
  "path": "/usr/local/bin", // Adds to the system path used by the CLI
  "organization": "a string", // The name of your organization, e.g. the output of: curl -H "Authorization: token $(snyk config get api)"  https://api.snyk.io/v1/cli-config/settings/sast | jq .org
  "token":  "secret-token", // The Snyk token, e.g.: snyk config get api
  "automaticAuthentication": "true", // Whether or not LS will automatically authenticate on scan start (default: true)
  "authenticationMethod": "oauth", // the authentication method (token, oauth, pat)
  "enableTrustedFoldersFeature": "true", // Whether or not LS will prompt to trust a folder (default: true)
  "trustedFolders": ["/a/trusted/path", "/another/trusted/path"], // An array of folder that should be trusted
}
```

For all .NET Projects, Snyk recommends adding the `--all-projects` additional parameter.

## **Authentication for Snyk Language Server**

When Snyk Language Server starts, it checks for a token in the initializationOption `token`. If a token is not there, Snyk Language Server tries to retrieve and authenticate.. If the CLI is not authenticated either, Snyk Language Server opens a browser window to authenticate. After successful authentication in the web browser, Snyk Language Server, in case of OAuth2 or API token authentication methods, automatically retrieves the Snyk authentication token from the CLI, but only for the session.

## **Environment variables for Snyk Language Server**

Snyk Language Server and Snyk CLI support and need certain environment variables to function:

1. `HTTP_PROXY`, `HTTPS_PROXY` and `NO_PROXY` to define the http proxy to be used
2. `JAVA_HOME` to analyze Java JVM-based projects via Snyk CLI
3. `PATH` to find `maven` when analyzing Maven projects, to find `python` and so on
4. `SNYK_LOG_LEVEL`  force a log-level (trace, debug, info, warn, error), default is info level

## **Auto-configuration of environment variables for Snyk Language Server**

To automatically add these variables to the environment, Snyk Language Server searches for the following files, with the order determining precedence. If the executable is not called from an already configured environment (for example, via `zsh -i -c 'snyk-ls'`), you can also specify the config file with the `-c` command line flag for setting the required variables. Snyk Language Server reads the following files in the given precedence and order, not overwriting the already loaded variables.

```
given config file via -c flag
<working-dir>/.snyk.env
$HOME/.snyk.env
```

Any lines that contain an environment variable in the format `VARIABLENAME=VARIABLEVALUE` are added automatically to the environment if not already there. This adheres to the `dotenv` format. In the case of `.profile`, `.zshrc`and so on, if a variable is directly exported, for example, via `export VARIABLENAME=VARIABLEVALUE`, it is not loaded. The export would need to be split of and be in its own line, for example

```bash
VARIABLENAME=VARIABLEVALUE
export VARIABLENAME
```

The PATH variable is treated differently frrom all other variables, as it is an aggregate of all PATH variables found in the files and in the environment. Also, the current working directory `.` is automatically added to the path, so a download of the Snyk CLI into the current working directory by an LSP client would yield a found Snyk CLI for the Language Server.

In addition to configuring variables via config files, Snyk Language Server adds the following directories to the path on Linux and macOS:

* /bin
* $HOME/bin
* /usr/local/bin
* $JAVA\_HOME/bin

If no JAVA\_HOME is set, Snyk Language Server automatically searches for a java executable first in `path`, then in the following directories, and adds the parent directory of its parent as JAVA\_HOME. The following directories are recursively searched:

* /usr/lib
* /usr/java
* /opt
* /Library
* $HOME/.sdkman
* C:\Program Files
* C:\Program Files (x86)

The same directories are searched for a Maven executable, and the parent directory is added to the path.

## Folder trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager (for example, pip, Gradle, Maven, Yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard against using the Language Server on untrusted folders, the Snyk Language Server asks for folder trust before running scans against these folders. When in doubt, do not grant trust.

The trust feature is enabled by default. When a folder is trusted, all sub-folders are also trusted. After a folder is trusted, Snyk Language Server notifies the Language Server Client with the custom `$/snyk.addTrustedFolders` notification, which contains a list of trusted folder paths. Based on this, a client can then implement logic to intercept this notification and persist the decision and trust in the IDE or Editor storage mechanism.

Trust dialogs can be disabled by setting `enableTrustedFoldersFeature` to `false` in the initialization options. This disables all trust prompts and checks.

An initial set of trusted folders can be provided by setting `trustedFolders` to an array of paths in the `initializationOptions`. These folders will be trusted on startup and will not prompt the user to trust them.

## Telemetry

Snyk collects telemetry from IDE plugins and CLI. For details, see [IDE and CLI usage telemetry](ide-and-cli-usage-telemetry.md).

## Support policy for Snyk Language Server

Snyk supports the latest 12 months of LS versions, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or troubleshooting.

Snyk only provides fixes in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io) to Snyk Support.
