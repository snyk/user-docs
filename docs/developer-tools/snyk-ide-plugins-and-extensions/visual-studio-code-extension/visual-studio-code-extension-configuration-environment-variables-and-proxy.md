# Visual Studio Code extension configuration, environment variables, and proxy

## Configuration of the Visual Studio Code extension

After the plugin is installed, you can set the following configurations for the extension.

### Snyk account

* **Authentication method:** Specify whether to authenticate with OAuth2 or with an API token. `OAuth2` is the default
* **Custom Endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).\
  Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect does not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Organization**: Sets the Organization to run `snyk test` against, similar to the `--org=` option in the CLI. Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If the Organization value is blank or invalid, the preferred Organization defined in your [web account settings](https://app.snyk.io/account) will be used. Setting the Organization at the Folder level is Experimental.
* **Auto-select Organization**: When enabled, Snyk automatically selects the most appropriate Organization for your Project using context found in your repository and your authentication. If an Organization is configured manually, this feature is overridden. If an appropriate Organization cannot be identified automatically, the preferred Organization defined in your [web account settings](https://app.snyk.io/account) is used as a fallback.
* **Send error reports to Snyk**: Help Snyk to improve the stability of the plugin by analyzing such reports.

### Scan configuration

* **Open Source**: Enable a scanner for open source dependencies; enabled by default.
* **Snyk Code Security issues**: Enable a scanner for security vulnerabilities in your application code. Enabled by default
* **Snyk IaC**: Enable a scanner for insecure configurations in Terraform and Kubernetes code. Enabled by default.
* **Severity selection**: Filter issues by their severity, from Low to Critical.
* **All issues vs Net new issues**: Specify whether to see all issues or only net new issues. The latter requires an SCM integration, where the extension compares findings with those in the base branch.
*   **Additional parameters**: Set additional `snyk test` CLI options for Open Source scanning.

    For unmanaged [C/C++](../../../supported-languages/supported-languages-list/c-c++/) scanning, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This option works only for unmanaged C/C++ scanning. Additional parameters do not apply to Snyk Code or IaC.

## User experience

* **Scanning mode:** The **auto** option activates automatic scans when saving files and when opening a Project. This works with Snyk Code and Snyk IaC.
* **Auto Scan Open Source Security:** Set to run Open Source Security analysis in automatic mode.
* **Snyk window recommendation**: Dock the Snyk window at the bottom or side of your screen for easier navigation. Avoid using it in full-screen mode.

## Experimental

This section contains experimental features that may change suddenly.

These settings are not part of the stable functionality and are not officially supported.

## Initialization

**Trusted Folders**: Link to the `settings.json` file, which has a list of folders that are marked as trusted. Use this setting only in advanced cases or when certain folders should be marked as not trusted.

## CLI and Language Server

* When **Automatic Dependency Management** is checked, the plugin will download the [CLI](../../snyk-cli/) and update it regularly to the defined CLI path, if defined. Uncheck this option if downloading the CLI is not possible due to your network configuration, for example, due to firewall rules, and you need to obtain this dependency through other means.
* **CLI Path:** Allow changing the file path of the CLI (optional field).

## Environment variables

To analyze Projects, the plugin uses the Snyk language server built into the Snyk CLI, which may require environment variables:

* `PATH`: the path to needed binaries, for example, to Maven.
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies.

Setting these variables only in a shell environment, for example, using `~/.bashrc`, is not sufficient if you do not start the IDE from the command line, or create a script file that starts the IDE using a shell environment.

* On Windows, you can set the variables using the GUI or on the command line using the `setx` tool.
* On macOS, the process `launchd` must know the environment variables to launch the IDE from the Finder directly. You can set environment variables for applications launched using the Finder by using the `launchctl setenv` command, for example, on start-up or through a script you launch at user login.

{% hint style="info" %}
The provision of environment variables to the `macOS` UI can change between operating system releases, so it may be easier to create a small shell script that launches the IDE to leverage the shell environment that can be defined using `~/.bashrc`.
{% endhint %}

* On Linux, updating the `file /etc/environment` can propagate environment variables to the Windows manager and UI.

## Proxy

If you are behind a proxy, configure the proxy settings using VS Code proxy settings `Application > Proxy` or set the proxy settings using the `http_proxy` and `https_proxy` environment variables.

For example, the commonly used setting **Proxy Strict SSL** specifies that the proxy server certificate should be verified against the list of supplied CAs specific to Snyk Code.
