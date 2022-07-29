# Visual Studio extension

The Visual Studio extension ([Snyk Security - Code and Open Source Dependencies](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs)) helps you find and fix security vulnerabilities in your projects. Within a few seconds, the extension provides a list of all the different types of security vulnerabilities identified together with actionable fix advice. The extension combines the power of two Snyk products: Snyk Open Source and Snyk Code.

1. Snyk Open Source finds known vulnerabilities in both the direct and indirect (transitive) open source dependencies you are pulling into the project.
2. Snyk Code finds known security vulnerabilities and code quality issues at blazing speed looking at the code you and your team wrote.

## Software requirements

* Operating system - Windows
* Supported versions of Visual Studio: 2015, 2017, 2019, 2022. Compatible with Community, Professional, and Enterprise

## Supported languages, package managers, and frameworks

* For Snyk Open Source: the Visual Studio extension supports all the languages and package managers supported by Snyk Open Source and the CLI. See the full list [in the docs](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code: the Visual Studio extension supports all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine).

## Install the extension

You can install the Snyk extension directly from the IDE; open **Extensions > Manage Extensions**.

![Manage extensions menu](../../.gitbook/assets/readme\_image\_2\_1\_1.png)

Search for _Snyk_ and select **Download** to download the Snyk Security - Code and Open Source Dependencies extension.

Once installed, use Snyk via the **Extensions > Snyk** menu (on Visual Studio versions older than 2019, Snyk will be part of the top menu bar).

![](<../../.gitbook/assets/image (76) (1) (1) (3).png>)

You can also open the Snyk tool window using **View > Other Windows > Snyk**_._

Once the tool window opens, wait while the Snyk extension downloads the latest Snyk CLI version.

![Snyk tool window, CLI downloading](../../.gitbook/assets/readme\_image\_2\_3.png)

After you install the extension and the CLI you must authenticate. You can use the **Connect Visual Studio to Snyk** link. For more information and additional ways to authenticate see [Authentication](./#authentication).

## Useful links

* This plugin works with projects written in .NET, Java, JavaScript, and many more languages. [See the full list of languages and package managers Snyk supports](https://support.snyk.io/hc/en-us/sections/360001087857-Language-package-manager-support)
* [Bug tracker](https://github.com/snyk/snyk-visual-studio-plugin/issues)
* [Github repository](https://github.com/snyk/snyk-visual-studio-plugin)

## Support and contact information

{% hint style="info" %}
Need more help? Submit a request to [Snyk support](https://snyk.zendesk.com/agent/dashboard).
{% endhint %}

**Share your experience.**

Snyk continuously strives to improve the Snyk plugins experience. Would you like to share with us your feedback about the Snyk Visual Studio extension? [Schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
