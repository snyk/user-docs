# Eclipse plugin

The Snyk Eclipse plugin provides analysis of your code, containers, and infrastructure as code configurations.

Snyk scans for the following types of issues:

* [**Open Source Security**](https://snyk.io/product/open-source-security-management/) - security vulnerabilities and license issues in both the direct and indirect (transitive) open-source dependencies pulled into the Snyk Project. See also the [`Open Source docs`](https://docs.snyk.io/products/snyk-open-source).
* [**Code Security**](https://snyk.io/product/snyk-code/) - security vulnerabilities in your code. See also the [Snyk Code docs](https://docs.snyk.io/products/snyk-code).
* [**Infrastructure as Code (IaC) Security**](https://snyk.io/product/infrastructure-as-code-security/) - configuration issues in your IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager. See also the [Snyk Infrastructure as Code docs](https://docs.snyk.io/products/snyk-infrastructure-as-code).

The Eclipse plugin provides automated, algorithm-based fix suggestions for both direct and transitive dependencies. This single plugin provides a Java vulnerability scanner and an open-source security scanner.

After you have installed and configured the Eclipse plugin, every time you run it, open a file, or autosave, Snyk scans the manifest files, proprietary code, and configuration files in your project. Snyk delivers actionable vulnerability, license, or misconfiguration issue details and displays the results natively within the Eclipse UI.

This page explains supported environments, support, and giving feedback and provides installation instructions. **After you complete the steps on this page**, you will continue by following the instructions in the other Eclipse plugins docs:

* [Download the CLI and language server with the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/download-the-cli-and-language-server-with-the-eclipse-plugin)
* [Authentication for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/authentication-for-the-eclipse-plugin)
* [Configuration of the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/configuration-of-the-eclipse-plugin)
* [Environment variables for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/environment-variables-for-the-eclipse-plugin)
* [Use the Snyk plugin to secure your Eclipse projects](https://docs.snyk.io/ide-tools/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects)
* [SAST scanning results (SAST, Snyk Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/sast-scanning-results-sast-snyk-code)
* [Misconfiguration scanning results (Snyk Infrastructure as Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/misconfiguration-scanning-results-snyk-infrastructure-as-code)
* [Third party dependency scanning (SCA, Snyk Open Source)](https://docs.snyk.io/ide-tools/eclipse-plugin/third-party-dependency-scanning-sca-snyk-open-source)
* [Troubleshooting for the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/troubleshooting-for-the-eclipse-plugin)

## Where you can download the Eclipse plugin

* **Eclipse Marketplace (recommended)**: [https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations)
* Preview update site (CI/CD, on commit): [https://storage.googleapis.com/snyk-eclipse-plugin-test/preview-2.1/repository/](https://storage.googleapis.com/snyk-eclipse-plugin-test/preview-2.1/repository/)
* Update site (weekly): [https://storage.googleapis.com/snyk-eclipse-plugin/weekly-2.1/repository/](https://storage.googleapis.com/snyk-eclipse-plugin-test/weekly-2.1/repository/)
* Manual downloads: [https://github.com/snyk/snyk-eclipse-plugin/releases](https://github.com/snyk/snyk-eclipse-plugin/releases)

**Signing Information for Jars**

If you want to verify the correct provenance of your download, verify the signing details from the Eclipse dialog with this data.

<figure><img src="../../../.gitbook/assets/image (134) (2) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="The signing key details to verify the integrity and origin of the download plugin"><figcaption><p>The signing key details to verify the integrity and origin of the download plugin</p></figcaption></figure>

The plugin runs on

* macOS
* Linux
* Windows

## Supported Eclipse versions

* 2023-03
* 2022-12
* 2022-09
* 2022-06
* 2022-03
* 2021-12
* 2021-09
* 2021-06
* 2021-03

## Supported languages, package managers, and frameworks

* For Snyk Open Source, the Eclipse plugin supports the languages and package managers supported by Snyk Open Source and the CLI except C/C++. For more information, see [Supported languages, frameworks, and feature availability overview, Open Source section](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source).
* For Snyk Code, the Eclipse plugin supports the languages and frameworks supported by Snyk Code. For more information, For more information, see [Supported languages, frameworks, and feature availability overview, Snyk Code section](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code).
* For Snyk IaC, the Eclipse plugin supports the following IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager.

## Supported operating systems and architecture

You can use the Eclipse plugin in the following environments:

* Linux: AMD64 and ARM64
* Windows: 386 and AMD64
* MacOS: AMD64 and ARM64

## How to install the Snyk Eclipse plugin

Navigate to the Marketplace from your running Eclipse instance. Search for Snyk and click **Install**.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-17 at 16.29.29.png" alt="Eclipse Marketplace search showing Snyk plugin and Install button"><figcaption><p>Eclipse Marketplace search showing Snyk plugin and Install button</p></figcaption></figure>

When you are prompted, accept the license agreement and add the **Snyk Security** certificate to complete the installation (this happens only once).

Restart the Eclipse instance:

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 09.16.37.png" alt="Restart Eclipse"><figcaption><p>Restart Eclipse</p></figcaption></figure>

Once Eclipse is restarted, the Snyk Wizard should run; this will set up your Snyk API endpoint and authentication token.

Once the Snyk configuration wizard runs; follow the instructions to set up your Snyk API:

<figure><img src="../../../.gitbook/assets/eclipseSnykWizard (1).png" alt="Snyk configuration wizard"><figcaption><p>Snyk configuration wizard</p></figcaption></figure>

Once the Snyk is configured, navigate to **Eclipse Preferences** to ensure that **Snyk** now appears in the list:

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png" alt="Eclipse preferences showing Snyk."><figcaption><p>Eclipse preferences showing Snyk.</p></figcaption></figure>

When you open the preferences, you can opt out of downloading the CLI through the plugin and thus use your own installation of the CLI.

Continue with the steps to [Download the CLI and language server with the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/download-the-cli-and-language-server-with-the-eclipse-plugin).

## Support

If you need help, submit a [request](https://support.snyk.io/hc/en-us/requests/new) to Snyk Support.
