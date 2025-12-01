# Configuration for the Snyk JetBrains plugin and IDE proxy

Navigate to **Preferences** > **Tools** > **Snyk** to set the following configurations for the plugin:

<figure><img src="../../../.gitbook/assets/Screenshot 2025-11-24 at 16.31.39.png" alt=""><figcaption></figcaption></figure>

## General settings

* **Authentication method:** Specify whether to authenticate with OAuth2 or with an API token. `OAuth2` is the default
* **Token**: Set the token to use for authentication with Snyk. For details, see [Authentication for the JetBrains plugins](authentication-for-the-jetbrains-plugins.md).
* **Custom endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).\
  Multi-tenant users who do not belong to the `SNYK-US-01` region are automatically redirected to the domain for the email with which the user authenticated. The redirect does not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Ignore unknown CA**: Ignore the SSL cert, if needed.
* **Organization**: Set the Organization to run `snyk test` against (similar to the `--org=` option in the CLI). Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk Web UI: `https://app.snyk.io/org/[orgslugname]`.\
  Organization selection follows this order:
  1. [Project-specific settings](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#project-settings) (if configured)
  2. [General setting](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#general-settings) (if the project-specific setting is empty)
  3. Your [web account's](https://app.snyk.io/account) preferred Organization (if both General and Project-specific settings are empty).
* **Snyk Open Source**: Enable a scanner for open-source dependencies. Enabled by default.
* **Snyk infrastructure as code**: Enable a scanner for insecure configurations in Terraform and Kubernetes code. Enabled by default.
* **Snyk Code security issues**: Enable a scanner for security vulnerabilities in your application code. Enabled by default.
* **Severity selection**: Filter issues by their severity, from Low to Critical.
* **All Issues vs Net New Issues**: Specify whether to see all issues or only net new issues. The latter requires an SCM integration, where it compares findings with those in the base branch.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-11-24 at 16.32.44.png" alt=""><figcaption></figcaption></figure>

## Project settings

* **Additional parameters**: Set additional `snyk test` [CLI options](../../snyk-cli/cli-commands-and-options-summary.md) for Open Source scanning. For **unmanaged** [**C/C++**](../../../supported-languages/supported-languages-list/c-c++/) **scanning**, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This option works only for unmanaged C/C++ scanning. Additional parameters do not apply to Snyk Code or IaC.
* **Auto-select Organization:** When enabled, Snyk automatically selects the most appropriate Organization for your Project using context found in your repository and your authentication. If an Organization is configured manually, this feature is overridden. If an appropriate Organization cannot be identified automatically, the preferred Organization defined in your [web account settings](https://app.snyk.io/account) is used as a fallback.
* **Preferred Organization**: Specify the Organization (ID or name) for Snyk to run scans against for this specific IDE Project. Retrieve the Organization ID from the Organization settings in the Snyk Web UI: `https://app.snyk.io/org/[ORG_NAME]/manage/settings` and copy the ID from the Organization ID section. If the Organization value is blank or invalid, the value from the global Organization field is used.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-11-24 at 16.34.07.png" alt=""><figcaption></figcaption></figure>

## Executable settings

These options allow you to customize the handling of the plugin dependency on Snyk CLI.

* **Base URL to download the CLI:** Allow specifying an alternative download location of the CLI. This location must conform to the same file and folder layout as `https://downloads.snyk.io`. For example, FIPS-supported CLIs would use the base URL `https://static.snyk.io/fips`.
* **Path to Snyk CLI:** Allow changing the file path of the Snyk CLI.
* **Automatically manage needed binaries:** Check to have the plugin download the CLI and update it regularly to the defined CLI path. Uncheck this option if downloading the CLI is not possible due to your network configuration, for example, due to firewall rules, and you need to obtain the CLI through other means.
* **CLI release channel:** Allow specifying a release channel (**preview**, **rc**, **stable**) for the CLI. You can also pin the CLI to a version here, specifying the version, for example,`v1.1293.1`.

## User experience

**Scan automatically on start-up and save:** Enable to activate automatic scans when saving files and when opening a Project.

## Proxy for the JetBrains plugins

If you need to use a proxy server to connect to the internet, configure it using the [JetBrains IDE settings](https://www.jetbrains.com/help/idea/settings-http-proxy.html). The Snyk plugin automatically uses the IDE settings.
