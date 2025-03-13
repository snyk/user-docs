# Visual Studio extension configuration

This page explains [Visual Studio extension settings](visual-studio-extension-configuration.md#visual-studio-extension-settings) and [environment variables](visual-studio-extension-configuration.md#environment-variables) needed for the CLI.

## Visual Studio extension settings

After the Visual Studio extension is installed, you can set the following options.

### Account settings

* **Authentication method**: Specify whether to authenticate with OAuth2 or with  a Snyk API token.  Snyk recommends using OAuth2, becuase it is more secure.
* **Token**: Set the authentication token for Snyk.
* **Custom endpoint**: Specify the Snyk API endpoint for a custom multi-tenant or single-tenant setup. If you are using `https://api.snyk.io`, no configuration is required. For details, see the list of [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).\
  Multi-tenant users who do not belong to the default region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.
* **Ignore unknown CA**: Ignore unknown certificate authorities.
* **Organization**: Specify the ORG\_ID to run Snyk commands tied to a specific Organization. Snyk recommends using the ORG\_ID. If you specify the ORG\_NAME, that is, the Organization slug name, the value must match the URL slug as displayed in the URL of your Organization in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If not specified, the Preferred Organization (as defined in your [account settings](https://app.snyk.io/account)) is used to run tests.

### CLI settings

You may opt to either use a CLI instance downloaded and managed by the extension, or to use your own installation of the CLI.

* **Base URL to download the CLI**: The URL that the extension will use to download the CLI.
* **Snyk CLI Path**: Browse to a custom installation of the CLI or Reset to default to use the installation managed by the extension.
* **Update and install Snyk dependencies automatically**: Specify whether or not the extension should download and use its own CLI instance. If this is disabled, you must provide a valid path to your own CLI instance.
* **CLI release channel:** For CLI instances managed by the extension, choose whether to use the stable, rc (release candidate) or preview versions of the CLI. For details on the CLI release channels, see [Releases and channels for the Snyk CLI](../../../snyk-cli/releases-and-channels-for-the-snyk-cli.md).

### Experimental

This section contains experimental features that may change suddenly.&#x20;

These settings are not part of the stable functionality and are not officially supported. &#x20;

### Scan configuration

* **Snyk Open Source**: Enable a scanner for open source dependencies; enabled by default.
* **Snyk Infrastructure as Code**: Enable a scanner for insecure configurations in Terraform and Kubernetes code; enabled by default.
* **Snyk Code Security**: Enable a scanner for security vulnerabilities in your application code; enabled by default
* **Snyk Code Quality**: Enable a scanner for code quality issues in your application code; disabled by default.

{% hint style="info" %}
Effective beginning June 24, 2025, Snyk Code Quality issues will no longer be provided.
{% endhint %}

* **All issues versus net new issues**: Specify whether to see all issues or only net new issues. The latter requires a Git repository, where the extension compares findings with those in the base branch.

### Solution settings

Use **Solution settings** to set **Additional Parameters** specifying `snyk test` [CLI options](../../../snyk-cli/commands/test.md) for Snyk Open Source scanning.&#x20;

For example, to enable unmanaged [C/C++](../../../supported-languages-package-managers-and-frameworks/c-c++/) scanning to find vulnerabilities in open-source packages, use the CLI option `--unmanaged`. Note that `--unmanaged`should not be used for other languages.&#x20;

**Additional Parameters** do not apply to Snyk Code or Snyk Infrastructure as Code.&#x20;

<div data-full-width="false"><figure><img src="../../../.gitbook/assets/Screenshot 2025-01-08 164652.png" alt=""><figcaption><p>VS Extension Solution settings with --unmanaged</p></figcaption></figure></div>

### User Experience

**Scan automatically on stat-up and save**: Specify whether to enable automatic scanning. Regardless of this setting, you may run a scan manually at any time. For details, see [Run an analysis with Visual Studio extension](run-an-analysis-with-visual-studio-extension.md).

## Environment variables

To analyze Projects the plugin uses the Snyk CLI, which requires the following environment variables:

* `PATH`: Include the paths to any needed binaries (for example, to Maven).
* `JAVA_HOME`: Specify the path to the JDK you want to use for analysis of Java dependencies
* `http_proxy` and `https_proxy`: Set if you are behind a proxy server, using the value in the format `http://username:password@proxyhost:proxyport`\
  Note that the leading `http://` in the value does not change to `https://` for `https_proxy.`

You can set environment variables using the Windows GUI or on the command line using the `setx` tool.
