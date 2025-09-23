# Configuration for the Snyk JetBrains plugin and IDE proxy

Navigate to **Preferences** > **Tools** > **Snyk** to set the following configurations for the plugin:

## **Snyk account**

* **Authentication method:** Specify whether to authenticate with OAuth2 or with an API token. `OAuth2` is the default
* **Token**: Set the token to use for authentication with Snyk. For details, see [Authentication for the JetBrains plugins](authentication-for-the-jetbrains-plugins.md).
* **Custom endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).\
  Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Ignore unknown CA**: Ignore the SSL cert, if needed
* **Organization**: Set the Organization to run `snyk test` against (similar to the `--org=` option in the CLI). Snyk recommends using the `ORG_ID`. If you specify an Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If this is not specified or incorrect, the preferred Organization, as defined in your [web account settings](https://app.snyk.io/account), is used to run tests.

## Scan configuration

* **Snyk Open Source**: Enable a scanner for open-source dependencies; enabled by default.
* **Snyk Infrastructure as Code**: Enable a scanner for insecure configurations in Terraform and Kubernetes code; enabled by default.
* **Snyk Code Security issues**: Enable a scanner for security vulnerabilities in your application code; enabled by default
* **Severity selection**: Filter issues by their severity, from Low to Critical.
* **All Issues vs Net New Issues**: Specify whether to see all issues or only net new issues. The latter requires a Git repository, where it compares findings with those in the base branch.
*   **Additional parameters**: Set additional `snyk test` [CLI options](../../snyk-cli/cli-commands-and-options-summary.md) for Open Source scanning.

    For **unmanaged** [**C/C++**](../../../supported-languages/supported-languages-list/c-c++/) **scanning**, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This option works only for unmanaged C/C++ scanning; do not use this option for other languages. Additional parameters do not apply to Snyk Code or IaC.

## **Executable settings**

These options allow you to customize the handling of the plugin dependency on Snyk CLI.&#x20;

* **Base URL to download the CLI:** Allow specifying an alternative download location of the CLI. This location must conform to the same file and folder layout as `https://downloads.snyk.io`. For example, FIPS-supported CLIs would use the base URL `https://static.snyk.io/fips`.
* **Path to Snyk CLI:** Allow changing the file path of the Snyk CLI.
* **Automatically manage needed binaries:** Check to have the plugin download the CLI and update it regularly to the defined CLI Path. Uncheck this option if downloading the CLI is not possible due to your network configuration, for example, due to firewall rules, and you need to obtain the CLI through other means.
* **CLI release channel:** Allow specifying a release channel (**preview**, **rc**, **stable**) for the CLI. You can also pin the CLI to a version here, specifying the version, for example,`v1.1293.1`.

## User experience

**Scan automatically on start-up and save:** Enable to activate automatic scans when saving files and when opening a Project.

## Proxy for the JetBrains plugins

If you need to use a proxy server to connect to the internet, configure it using the [JetBrains IDE settings](https://www.jetbrains.com/help/idea/settings-http-proxy.html). The Snyk plugin will use IDE settings automatically.
