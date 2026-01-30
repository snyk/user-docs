# Eclipse plugin

## **Scan early, fix as you develop: elevate your security posture**

Integrating security checks early in your development lifecycle helps you pass security reviews seamlessly and avoid expensive fixes down the line.

The Snyk Eclipse extension allows you to analyze your code, open-source dependencies, and Infrastructure as Code (IaC) configurations. With actionable insights directly in your IDE, you can address issues as they arise.

Key features:

* In-line issue highlighting: Security issues are flagged directly within your code, categorized by type and severity for quick identification and resolution.
* Comprehensive scanning: The extension scans for a wide range of security issues, including:
  * [Open Source Security](https://snyk.io/product/open-source-security-management/): Detects vulnerabilities and license issues in both direct and transitive open-source dependencies. Automated fix suggestions simplify remediation. Explore more in the [Snyk Open Source documentation](https://docs.snyk.io/scan-using-snyk/snyk-open-source).
  * [Code Security](https://snyk.io/product/snyk-code/): Identifies security vulnerabilities in your custom code. Explore more in the [Snyk Code documentation](https://docs.snyk.io/scan-using-snyk/snyk-code).
  * [IaC Security](https://snyk.io/product/infrastructure-as-code-security/): Uncovers configuration issues in your Infrastructure as Code templates (Terraform, Kubernetes, CloudFormation, Azure Resource Manager). Explore more in the [IaC documentation](https://docs.snyk.io/scan-using-snyk/snyk-iac).
* Broad language and framework support: Snyk Open Source and Snyk Code cover a wide array of package managers, programming languages, and frameworks, with ongoing updates to support the latest technologies. For the most up-to-date information on supported languages, package managers, and frameworks, see the [supported language technologies pages](https://docs.snyk.io/supported-languages-package-managers-and-frameworks).

## How to install and set up the extension

{% hint style="info" %}
For information about the versions of Eclipse supported by the Eclipse plugin, see [Snyk IDE plugins and extensions](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions). Snyk recommends always using the latest version of the Eclipse plugin.
{% endhint %}

You can use the Eclipse plugin in the following environments:

* Linux: AMD64 and ARM64
* Windows: 386, AMD64, and ARM64
* MacOS: AMD64 and ARM64

Install the plugin at any time free of charge from the [Eclipse marketplace](https://marketplace.eclipse.org/content/snyk-security) and use it with any Snyk account, including the Free plan.

When you are prompted, accept the license agreement and add the **Snyk Security** certificate to complete the installation (this happens only once).

Continue by following the instructions in the other Eclipse extension docs:

* [Download the CLI with the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/download-the-cli-and-language-server-with-the-eclipse-plugin)
* [Authentication for the Eclipse plugin](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/authentication-for-the-eclipse-plugin)
* [Eclipse plugin folder trust](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/folder-trust)
* [Configuration of the Eclipse plugin](https://docs.snyk.io/ide-tools/eclipse-plugin/configuration-of-the-eclipse-plugin)
* [Environment variables for the Eclipse plugin](environment-variables-for-the-eclipse-plugin.md)
* [Use the Snyk plugin to secure your Eclipse projects](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects)
