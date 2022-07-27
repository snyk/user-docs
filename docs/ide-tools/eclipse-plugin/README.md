---
description: Use this documentation to get started with the Eclipse plugin.
---

# Eclipse plugin

Install the **Snyk Security - Code,​ Open Source,​ IaC Configurations** extension in your Eclipse workflow to find and fix security vulnerabilities, license issues, infrastructure misconfigurations, and code quality issues. Find and fix these early in the development lifecycle to help you ace your security reviews and avoid a costly fix later down the line. If you’re an individual developer, open-source contributor, or maintainer at a large organization, Snyk helps you ship secure code, faster.

Snyk scans for issue types around:

* [**Open Source Security**](https://snyk.io/product/open-source-security-management/) - security vulnerabilities in both the direct and in-direct (transitive) open-source dependencies you are pulling into the project.
* [**Code Security**](https://snyk.io/product/snyk-code/) - security vulnerabilities identified in your own code.
* [**Infrastructure as Code (IaC) Security**](https://snyk.io/product/infrastructure-as-code-security/) - configuration issues in your IaC templates (Terraform, Kubernetes, CloudFormation, and Azure Resource Manager)
* [**Code Quality**](https://snyk.io/product/snyk-code/) - code quality issues in your own code

After you have installed and configured the Eclipse plugin, every time you run it, open a file, or autosave, Snyk scans the manifest files, proprietary code, and configuration files in your project. Snyk delivers actionable vulnerability, license, code quality, or misconfiguration issue details and displays the results natively within the Eclipse UI.

{% hint style="info" %}
The Snyk Eclipse plugin is available for installation on the [Eclipse Marketplace](https://marketplace.eclipse.org/content/snyk-security-scanner).
{% endhint %}

## Supported languages, package managers, and frameworks

* For Snyk Open Source, the Eclipse plugin supports the languages and package managers supported by Snyk Open Source and the CLI except C/C++. See the full [list](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code, the Eclipse plugin supports all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine).
* For Snyk IaC, the Eclipse plugin supports the following IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager.

## Installing the Snyk Eclipse plugin

Navigate to the Marketplace from within your running Eclipse instance. Search for Snyk and click **Install**.

![Eclipse Marketplace search showing Snyk plugin and Install button](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.29.29.png>)

When prompted accept the license agreement add the **Snyk Security** certificate to complete the installation (this happens only once).

![Add Snyk Security certificate](<../../.gitbook/assets/Screenshot 2022-05-13 at 09.08.52 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>)

Restart the Eclipse instance:

![Restart Eclipse](<../../.gitbook/assets/Screenshot 2022-05-13 at 09.16.37.png>)

Once Eclipse is restarted, navigate to **Eclipse Preferences** to ensure that **Snyk** now appears in the list:

![Eclipse preferences showing Snyk.](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

## Use the Snyk plugin to secure your Eclipse projects

Once the language server is downloaded and the authentication is done, the plugin will successfully start the workspace scan. You might notice a confirmation that a workspace scan is starting. Snyk shows such a notification when there is no workspace scan available.

![Starting workspace scan](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.55.41.png>)

All of the issues found by Snyk are now natively integrated with Eclipse's flows. Issues are shown in the Problems tab (see the following screenshot). There is a squiggly line indicating the issue while you code plus the gutter icons to indicate where the issue is.

![Problems tab](<../../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png>)

## Download urls

* Eclipse Marketplace (recommended): [https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations](https://marketplace.eclipse.org/content/snyk-security-code%E2%80%8B-open-source%E2%80%8B-iac-configurations)
* Preview update site (CI/CD, on commit): [https://storage.googleapis.com/snyk-eclipse-plugin-test/preview/repository/](https://storage.googleapis.com/snyk-eclipse-plugin-test/preview/repository/)
* Update site (weekly): [https://storage.googleapis.com/snyk-eclipse-plugin/weekly/repository/](https://storage.googleapis.com/snyk-eclipse-plugin/weekly/repository/)
* Manual downloads: [https://github.com/snyk/snyk-eclipse-plugin/releases](https://github.com/snyk/snyk-eclipse-plugin/releases)

#### Signing Information for Jars

If you want to verify the correct provenance of your download, please verify the signing details from within the Eclipse dialogue with this data.

```bash
Creation date: May 9, 2022
Owner: CN=Snyk Ltd, OU=Road Runner, O=Snyk Ltd, L=London, ST=London, C=GB
Issuer: CN=Snyk Ltd, OU=Road Runner, O=Snyk Ltd, L=London, ST=London, C=GB
Serial number: 740679006
Valid from: Mon May 09 21:13:17 CEST 2022 until: Wed May 08 21:13:17 CEST 2024
```

### Supported Operating Systems

* MacOSX
* Linux
* Windows 10

### Supported Eclipse Versions

* 2022-03
* 2021-12
* 2021-09
* 2021-06
* 2021-03

## Support / Contact

{% hint style="info" %}
If you need help, submit a [request](https://support.snyk.io/hc/en-us/requests/new) to Snyk Support.
{% endhint %}

## Share your experience

Snyk continuously strives to improve the Snyk plugins experience. If you would you like to share your feedback about the Snyk Eclipse plugin, [schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
