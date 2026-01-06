# Visual Studio extension configuration, environment variables, and proxy

After the plugin is installed, you can set the following configurations for the extension.

## Account <a href="#snyk-account" id="snyk-account"></a>

* **Authentication method:** Specify whether to authenticate with OAuth2 or with an API token. `OAuth2` is the default.
* **Custom endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#ides-urls).\
  Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Ignore unknown CA**: Ignore unknown certificate authorities.
* **Organization**: Specify the `ORG_ID` to run Snyk commands tied to a specific Organization. Snyk recommends using the `ORG_ID`. If you specify the `ORG_NAME`, that is, the Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk Web UI: `https://app.snyk.io/org/[orgslugname]`.\
  Organization selection follows this order:
  1. [Project-specific settings](visual-studio-extension-configuration-environment-variables-and-proxy.md#solution-settings) (if configured)
  2. [General setting](visual-studio-extension-configuration-environment-variables-and-proxy.md#snyk-account) (if the project-specific setting is empty)
  3. Your [web account's](https://app.snyk.io/account) preferred Organization (if both general and Project-specific settings are empty).

<figure><img src="../../../.gitbook/assets/Screenshot 2025-11-24 at 17.24.39.png" alt=""><figcaption><p>VS Extension Account settings</p></figcaption></figure>

## Scan configuration <a href="#scan-configuration" id="scan-configuration"></a>

* **Open Source**: Enable a scanner for open source dependencies. Enabled by default.
* **Snyk Code security issues**: Enable a scanner for security vulnerabilities in your application code. Enabled by default.
* **Snyk IaC**: Enable a scanner for insecure configurations in Terraform and Kubernetes code. Enabled by default.
* **All issues vs Net new issues**: Specify whether to see all issues or only net new issues. The latter requires a Git repository, where the extension compares findings with those in the base branch.

## User experience

**Scanning mode:** The auto option activates automatic scans when saving files and when opening a Project. This works with Snyk Code and Snyk IaC.

**Snyk Window Recommendation:** The Snyk window should be docked either to the bottom or to the side of the IDE, to enable smoother navigation. It should not be used in full-screen mode.

## Experimental

This section contains experimental features that may change at any time.

These settings are not part of the stable functionality and are not officially supported.

## CLI settings

You may opt to either use a CLI instance downloaded and managed by the extension or to use your own installation of the CLI.

* **Base URL to download the CLI**: The URL that the extension will use to download the CLI.
* **CLI path:** Allow changing the file path to the CLI (optional field).
* **Update and install Snyk dependencies automatically**: Specify whether or not the extension should download and use its own CLI instance. If this is disabled, you must provide a valid path to your own CLI instance.
* **CLI release channel:** For CLI instances managed by the extension, choose whether to use the stable, rc (release candidate), or preview versions of the CLI. For details on the CLI release channels, see [Releases and channels for the Snyk CLI](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md).

## Solution settings

For example, to enable unmanaged [C/C++](../../../supported-languages/supported-languages-list/c-c++/) scanning to find vulnerabilities in open-source packages, use the CLI option `--unmanaged`. Note that `--unmanaged` works only for unmanaged C/C++ scanning; do not use this option for other languages.

{% hint style="info" %}
**Additional parameters** do not apply to Snyk Code or Snyk Infrastructure as Code.
{% endhint %}

Settings on this page are scoped to the active solution

* **Additional parameters:** Passed to `snyk test` [CLI options](../../snyk-cli/commands/test.md) for **Snyk Open Source** scanning.\
  For example, to enable unmanaged [C/C++](../../../supported-languages/supported-languages-list/c-c++/) scanning to find vulnerabilities in open-source packages, use the CLI option `--unmanaged`. Note that `--unmanaged` works only for unmanaged C/C++ scanning.
* **Auto-select Organization:** When enabled, Snyk will automatically select the most appropriate Organization for your Project using context found in your repository and your authentication. If an Organization is configured manually, this feature is overridden. If an appropriate Organization cannot be identified automatically, the preferred Organization defined in your [web account settings](https://app.snyk.io/account) is used as a fallback.
* **Organization**: Specify the Organization (ID or name) for Snyk to run scans against for this specific IDE project. Retrieve the Organization ID from the Organization settings in the Snyk Web UI: `https://app.snyk.io/org/[ORG_NAME]/manage/settings` and copy the ID from the Organization ID section. If the Organization is provided manually, automatic Organization selection is overridden. If the Organization value is blank or invalid, the value from the global Organization field is used.

<div data-full-width="false"><figure><img src="../../../.gitbook/assets/Screenshot 2025-11-24 at 17.58.07.png" alt=""><figcaption><p>VS Extension Solution settings</p></figcaption></figure></div>

## Environment variables

To analyze Projects, the extension uses the Snyk CLI, which requires the following environment variables:

* `PATH`: the path to needed binaries, for example, Maven.
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies.
* `http_proxy` and `https_proxy`: Set if you are behind a proxy server, using the value in the format `http://username:password@proxyhost:proxyport`.

{% hint style="info" %}
The leading `http://` in the value does not change to `https://` for `https_proxy.`
{% endhint %}

You can set environment variables using the Windows GUI or on the command line using the `setx` tool.

## Proxy <a href="#proxy" id="proxy"></a>

If you are behind a proxy, set the proxy settings using the `http_proxy` and `https_proxy` environment variables. You can set environment variables using the Windows GUI or on the command line using the `setx` tool.

For example, the commonly used setting `Proxy Strict SSL` specifies that the proxy server certificate should be verified against the list of supplied CAs specific to Snyk Code.
