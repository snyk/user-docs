---
description: Use this documentation to get started with the Eclipse plugin.
---

# Eclipse plugin

Snyk offers IDE integrations that allow you to use the functionality of Snyk in your Integrated Development Environment. This page describes the Snyk Eclipse plugin. For information about all of the IDE plugins and their use, see [Snyk for IDEs](https://docs.snyk.io/ide-tools) in the docs.

The Snyk Eclipse plugin provides analysis of your code, containers, and Infrastructure as Code configurations.

Snyk scans for vulnerabilities, open source license issues, code quality, and infrastructure misconfigurations and returns results with security issues categorized by issue type and severity.

For open source, you receive automated, algorithm-based fix suggestions for both direct and transitive dependencies. This single plugin provides a Java vulnerability scanner or an open-source security scanner.

Snyk scans for the following types of issues:

* [**Open Source Security**](https://snyk.io/product/open-source-security-management/) - security vulnerabilities and license issues in both the direct and in-direct (transitive) open-source dependencies pulled into the Snyk Project. See also the [`Open Source docs`](https://docs.snyk.io/products/snyk-open-source).
* [**Code Security**](https://snyk.io/product/snyk-code/) and [**Code Quality**](https://snyk.io/product/snyk-code/) - security vulnerabilities and quality issues in your code. See also the [Snyk Code docs](https://docs.snyk.io/products/snyk-code).
* [**Infrastructure as Code (IaC) Security**](https://snyk.io/product/infrastructure-as-code-security/) - configuration issues in your IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager. See also the [Snyk Infrastructure as Code docs](https://docs.snyk.io/products/snyk-infrastructure-as-code).

After you have installed and configured the Eclipse plugin, every time you run it, open a file, or autosave, Snyk scans the manifest files, proprietary code, and configuration files in your project. Snyk delivers actionable vulnerability, license, code quality, or misconfiguration issue details and displays the results natively within the Eclipse UI.

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
* Preview update site (CI/CD, on commit): [https://storage.googleapis.com/snyk-eclipse-plugin-test/preview/repository/](https://storage.googleapis.com/snyk-eclipse-plugin-test/preview/repository/)
* Update site (weekly): [https://storage.googleapis.com/snyk-eclipse-plugin/weekly/repository/](https://storage.googleapis.com/snyk-eclipse-plugin/weekly/repository/)
* Manual downloads: [https://github.com/snyk/snyk-eclipse-plugin/releases](https://github.com/snyk/snyk-eclipse-plugin/releases)

**Signing Information for Jars**

If you want to verify the correct provenance of your download, please verify the signing details from the Eclipse dialog with this data.

![The signing key details to verify the integrity and origin of the download plugin](<../../.gitbook/assets/image (134) (2) (1) (1) (1) (1).png>)

The plugin runs on

* macOS
* Linux
* Windows

## Supported Eclipse Versions

* 2022-06
* 2022-03
* 2021-12
* 2021-09
* 2021-06
* 2021-03

## Supported languages, package managers, and frameworks

* For Snyk Open Source, the Eclipse plugin supports the languages and package managers supported by Snyk Open Source and the CLI except C/C++. See [Open Source - Supported languages and package managers](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code, the Eclipse plugin supports all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine).
* For Snyk IaC, the Eclipse plugin supports the following IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager.

## How to install the Snyk Eclipse plugin

Navigate to the Marketplace from your running Eclipse instance. Search for Snyk and click **Install**.

![Eclipse Marketplace search showing Snyk plugin and Install button](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.29.29.png>)

When prompted accept the license agreement add the **Snyk Security** certificate to complete the installation (this happens only once).

![Add Snyk Security certificate](../../.gitbook/assets/Screenshot%202022-05-13%20at%2009.08.52)

Restart the Eclipse instance:

![Restart Eclipse](<../../.gitbook/assets/Screenshot 2022-05-13 at 09.16.37.png>)

Once Eclipse is restarted, the Snyk Wizard should run; this will setup your Snyk API endpoint and authentication token:

Once the Snyk configuration wizard runs; follow the instructions to set up your Snyk API:

<figure><img src="../../.gitbook/assets/eclipseSnykWizard.png" alt="Snyk configuration wizard"><figcaption><p>Snyk configuration wizard</p></figcaption></figure>

Once the Snyk is configured, navigate to **Eclipse Preferences** to ensure that **Snyk** now appears in the list:

![Eclipse preferences showing Snyk.](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

When you open the preferences you can opt out of downloading the CLI through the plugin and thus use your own installation of the CLI.

Continue with the steps to [Download the CLI and language server with the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/download-the-cli-and-language-server-with-the-eclipse-plugin).

## Support

If you need help, submit a [request](https://support.snyk.io/hc/en-us/requests/new) to Snyk Support.
