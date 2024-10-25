# Configuration for the Snyk JetBrains plugin and IDE proxy

You can set the following configurations for the plugin, using **Preferences → Tools → Snyk**.

## **Snyk account**

* **Authentication method:** Specifies whether to authenticate with OAuth2 or with an API token. `OAuth2` by default
* **Token**: Sets the token that should be used for authentication with Snyk. For details, see [Authentication for the JetBrains plugins](authentication-for-the-jetbrains-plugins.md).
* **Custom endpoint**: Specifies the Snyk API endpoint for custom multi-tenant or single-tenant setup. The default is `https://api.snyk.io`. For details, see [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).
* **Ignore unknown CA**: Ignores the SSL cert, if needed
*   **Organization**: Sets the Organization to run `snyk test` against (similar to the `--org=` option in the CLI). Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`.

    If this is not specified, the preferred Organization, as defined in your [web account settings](https://app.snyk.io/account), is used to run tests.

## Scan configuration

* **Snyk Open Source**: Enables a scanner for open source dependencies; enabled by default.
* **Snyk Infrastructure as Code**: Enables a scanner for insecure configurations in Terraform and Kubernetes code; enabled by default.
* **Snyk Container vulnerabilities**: Enables a scanner for container vulnerabilities in container images and Kubernetes applications; enabled by default.
* **Snyk Code Security issues**: Enables a scanner for security vulnerabilities in your application code; enabled by default
* **Snyk Code Quality issues**: Enables a scanner for code quality issues in your application code; disabled by default.
* **Severity selection**: Filters issues by their severity, from Low to Critical.
* **All Issues vs Net New Issues**: Specifies whether to see all issues or only net new issues. The latter requires a Git repository, where it compares findings with those in the base branch.
*   **Additional parameters**: Set additional `snyk test` [CLI options](../../../snyk-cli/cli-commands-and-options-summary.md) for Open Source scanning.

    For **unmanaged** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **scanning**, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This option works only for unmanaged C/C++ scanning; do not use this option for other languages. Additional parameters do not apply to Snyk Code or IaC.

## **Executable settings**

These options allow to customize the handling of the plugin dependency on Snyk CLI.&#x20;

* **Base URL to download the CLI:** Allows you to specify an alternative download location of the CLI. This location must conform to the same file and folder layout as https://downloads.snyk.io. For example, FIPS-supported CLIs would use the base URL https://static.snyk.io/fips.
* **Path to Snyk CLI:** Allows changing a file path of the Snyk CLI.
* When **Automatically manage needed binaries** is checked, the plugin will download the CLI and update it regularly to the defined CLI Path. Uncheck this option if downloading the CLI is not possible due to your network configuration, for example, due to firewall rules, and you need to obtain the CLI through other means.
* **CLI release channel:** Allows you to specify a release channel (**preview**, **rc**, **stable**) for the CLI. You can also pin the CLI to a version here, specifying the version, for example,`v1.1293.1`.

## **User experience**

* **Scan automatically on start-up and save:** If enabled, activates automatic scans when saving files and when opening a project.

## Proxy for the JetBrains plugins

If you need to use a proxy server to connect to the internet, configure it using the [JetBrains IDE settings](https://www.jetbrains.com/help/idea/settings-http-proxy.html). The Snyk plugin will use IDE settings automatically.
