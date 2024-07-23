# Visual Studio extension configuration

## Environment variables

To analyze projects the plugin uses the Snyk CLI, which requires environment variables:

* `PATH`: specify the path to needed binaries, for example, to Maven
* `JAVA_HOME`: specify the path to the JDK you want to use for analysis of Java dependencies
* `http_proxy` and `https_proxy`: set if you are behind a proxy server, using the value in the format `http://username:password@proxyhost:proxyport`\
  **Note:** the leading `http://` in the value does not change to `https://` for `https_proxy`

You can set the variables using the Web UI or on the command line using the `setx` tool.

## Configuration of the Visual Studio extension

After the plugin is installed, you can set the following configurations for the extension:

* **Token**: Enter the token the extension uses to connect to Snyk. You can manually replace it, if you need to switch to another account.
* **Custom endpoint**: Specify the custom Snyk API endpoint for your organization. Use this field for Single Tenant setups as well instead of https://app.snyk.io. See [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls).
* **Ignore unknown CA**: Ignore unknown certificate authorities.
* **Organization**: Specify the ORG\_ID to run Snyk commands tied to a specific organization. Snyk recommends using the ORG\_ID. If you specify the ORG\_NAME, that is, the organization slug name, the value must match the URL slug as displayed in the URL of your org in the Snyk UI: https://app.snyk.io/org/\[orgslugname]. If not specified, the Preferred Organization (as defined in your [account settings](https://app.snyk.io/account)) is used to run tests.
* **Send usage analytics**: To help Snyk improve the extension, let your Visual Studio send Snyk information about how the extension is working.
* **Project settings**: Specify any additional Snyk CLI parameters.\
  For all .NET Projects, Snyk recommends adding the `--all-projects` additional parameter.
* **Scan all projects**: Auto-detect all projects in the working directory, enabled by default.
* **Executable settings**: You can opt out of downloading the CLI through the plugin and thus use your own installation of the CLI.
  * When **Automatically manage needed binaries** is checked, the plugin automatically downloads the CLI and keeps the CLI updated.
  *   When **Automatically manage needed binaries** is unchecked, you must provide a valid path to the CLI. Use this option if downloading the CLI is not possible due to your network configuration (for example, due to firewall rules) and you need to obtain the CLI through other means. Snyk recommends recommends always using the most recent version of the CLI.

      <figure><img src="../../../.gitbook/assets/image (10) (1) (1) (1).png" alt=""><figcaption><p>VS Code extension executable settings</p></figcaption></figure>
* **Solution Settings**: Set additional `snyk test` [CLI options](../../../snyk-cli/commands/test.md) for the Open Source scanning. For **unmanaged** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **scanning**, use the CLI option `--unmanaged` to find vulnerabilities in open-source packages. This requires **Scan all projects** to be disabled. The `--unmanaged` option works only for unmanaged C/C++ scanning; do not use this option for other languages. Additional parameters do not apply to Snyk Code or IaC.

<figure><img src="../../../.gitbook/assets/VS_Oprions_Unmagaed.jpg" alt="VS Extension Solution Settings with --unmanaged"><figcaption><p>VS Extension Solution Settings with --unmanaged</p></figcaption></figure>

## Product selection for Visual Studio extension results

In the settings, you can also choose which results you want to receive:

* Open Source vulnerabilities
* Snyk Code Security vulnerabilities
* Snyk Code Quality issues
