# Visual Studio extension configuration, environment variables, and proxy

After the plugin is installed, you can set the following configurations for the extension.

## Snyk account <a href="#snyk-account" id="snyk-account"></a>

* **Authentication method:** Specify whether to authenticate with OAuth2 or with an API token. `OAuth2` is the default
* **Custom Endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of IDE URLs.\
  Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Organization**: Specify the `ORG_ID` to run Snyk commands tied to a specific Organization. Snyk recommends using the `ORG_ID`. If you specify the `ORG_NAME`, that is, the Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If the Organization is not specified, the preferred Organization (as defined in your [account settings](https://app.snyk.io/account)) is used to run tests.
* **Ignore unknown CA**: Ignore unknown certificate authorities.

## Scan configuration <a href="#scan-configuration" id="scan-configuration"></a>

* **Open Source**: Enable a scanner for open source dependencies; enabled by default.
* **Snyk Code Security issues**: Enable a scanner for security vulnerabilities in your application code; enabled by default
* **Infrastructure as Code**: Enable a scanner for insecure configurations in Terraform and Kubernetes code; enabled by default.
* **All Issues vs Net New Issues**: Specify whether to see all issues or only net new issues. The latter requires a Git repository, where the extension compares findings with those in the base branch.
*   **Additional parameters**: Set additional `snyk test` CLI options for Open Source scanning.

    For unmanaged [C/C++](../../../supported-languages/supported-languages-list/c-c++/) scanning, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This option works only for unmanaged C/C++ scanning; do not use this option for other languages. Additional parameters do not apply to Snyk Code or IaC.

## User experience

**Scanning mode:** The **auto** option activates automatic scans when saving files and when opening a Project; this works with Snyk Code and IaC.

## Experimental

This section contains experimental features that may change at any time.&#x20;

These settings are not part of the stable functionality and are not officially supported. &#x20;

## CLI

You may opt to either use a CLI instance downloaded and managed by the extension or to use your own installation of the CLI.

* **Base URL to download the CLI**: The URL that the extension will use to download the CLI.
* **CLI Path:** Allow changing the file path to the CLI (optional field).
* **Update and install Snyk dependencies automatically**: Specify whether or not the extension should download and use its own CLI instance. If this is disabled, you must provide a valid path to your own CLI instance.
* **CLI release channel:** For CLI instances managed by the extension, choose whether to use the stable, rc (release candidate), or preview versions of the CLI. For details on the CLI release channels, see [Releases and channels for the Snyk CLI](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md).

## Solution settings

Use **Solution settings** to set **Additional Parameters** specifying `snyk test` [CLI options](../../snyk-cli/commands/test.md) for Snyk Open Source scanning.&#x20;

For example, to enable unmanaged [C/C++](../../../supported-languages/supported-languages-list/c-c++/) scanning to find vulnerabilities in open-source packages, use the CLI option `--unmanaged`. Note that `--unmanaged` works only for unmanaged C/C++ scanning; do not use this option for other languages.&#x20;

**Additional Parameters** do not apply to Snyk Code or Snyk Infrastructure as Code.&#x20;

<div data-full-width="false"><figure><img src="../../../.gitbook/assets/Screenshot 2025-01-08 164652.png" alt=""><figcaption><p>VS Extension Solution settings with --unmanaged</p></figcaption></figure></div>

## Environment variables

To analyze Projects, the extension uses the Snyk CLI, which requires the following environment variables:

* `PATH`: the path to needed binaries, for example, Maven.
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies.
* `http_proxy` and `https_proxy`: Set if you are behind a proxy server, using the value in the format `http://username:password@proxyhost:proxyport`.\
  Note that the leading `http://` in the value does not change to `https://` for `https_proxy.`

You can set environment variables using the Windows GUI or on the command line using the `setx` tool.

## Proxy <a href="#proxy" id="proxy"></a>

If you are behind a proxy, set the proxy settings using the `http_proxy` and `https_proxy` environment variables. You can set environment variables using the Windows GUI or on the command line using the `setx` tool.

For example, the commonly used setting `Proxy Strict SSL` specifies that the proxy server certificate should be verified against the list of supplied CAs specific to Snyk Code.
