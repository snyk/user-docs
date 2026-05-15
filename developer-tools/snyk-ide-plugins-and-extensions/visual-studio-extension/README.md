# Visual Studio extension

{% hint style="info" %}
Visual Studio 2026 not supported.
{% endhint %}

## **Scan early, fix as you develop: elevate your security posture**

Integrating security checks early in your development lifecycle helps you pass security reviews seamlessly and avoid expensive fixes down the line.

The Snyk Visual Studio extension allows you to analyze your code, infrastructure configuration, and open-source dependencies. With actionable insights directly in your IDE, you can address issues as they arise.

**Key features:**

* **Comprehensive scanning:** The extension scans for a wide range of security issues, including:
  * [**Open Source Security**](https://snyk.io/product/open-source-security-management/)**:** Detects vulnerabilities and license issues in both direct and transitive open-source dependencies. Automated fix suggestions simplify remediation. Explore more in the [Snyk Open Source documentation](https://docs.snyk.io/scan-using-snyk/snyk-open-source).
  * [**Code Security**](https://snyk.io/product/snyk-code/)**:** Identifies security vulnerabilities in your custom code. Explore more in the [Snyk Code documentation](https://docs.snyk.io/scan-using-snyk/snyk-code).
  * [**IaC Security**](https://snyk.io/product/infrastructure-as-code-security/)**:** Uncovers configuration issues in your Infrastructure as Code templates (Terraform, Kubernetes, CloudFormation, Azure Resource Manager). Explore more in the [IaC documentation](https://docs.snyk.io/scan-using-snyk/snyk-iac).
* **Broad language and framework support:** Snyk Open Source and Snyk Code cover a wide array of package managers, programming languages, and frameworks, with ongoing updates to support the latest technologies. For the most up-to-date information on supported languages, package managers, and frameworks, see the [supported language technologies pages](https://docs.snyk.io/supported-languages-package-managers-and-frameworks).

## How to install and set up the extension

**Note:** For information about the versions of Visual Studio supported by the Visual Studio extension, see [Snyk IDE plugins and extensions](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions). Snyk recommends always using the latest version of the Visual Studio extension.

You can use the Snyk Visual Studio extension in the following environments:

* Windows: 386, AMD64, and ARM64
* MacOS: Visual Studio Windows plugin in a Windows virtual machine inside a Mac with an ARM64 processor

Install the plugin at any time free of charge from the [Visual Studio marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs-2022) and use it with any Snyk account, including the Free plan. For more information, see the [VS extension installation guide](https://learn.microsoft.com/en-us/visualstudio/ide/finding-and-using-visual-studio-extensions?view=vs-2022#find-and-install-extensions).

After the extension is installed, use Snyk through the **Extensions > Snyk** menu.

<figure><img src="../../../.gitbook/assets/image (24).png" alt=""><figcaption><p>Snyk extensions menu</p></figcaption></figure>

You can also open the Snyk tool window using **View > Other Windows > Snyk**_._

When the extension is installed, it automatically downloads the [Snyk CLI,](https://docs.snyk.io/snyk-cli) which includes the [Language Server](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/snyk-language-server).

Continue by following the instructions in the other Visual Studio extension docs:

* [Visual Studio extension configuration, environment variables, and proxy](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-configuration)
* [Authentication for Visual Studio extension](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-authentication)
* [Visual Studio Workspace trust](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/workspace-trust)
* [Run an analysis with Visual Studio extension](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/run-an-analysis-with-visual-studio-extension)
* [View analysis results from Visual Studio extension](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/view-analysis-results-from-visual-studio-extension)

## Support

For troubleshooting and known issues, see [Troubleshooting and known issues with Visual Studio extension](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/troubleshooting-and-known-issues-with-visual-studio-extension).

If you need help, submit a request to [Snyk Support](https://support.snyk.io).
