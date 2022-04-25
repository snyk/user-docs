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

![Manage extensions menu](../../../.gitbook/assets/readme\_image\_2\_1\_1.png)

Search for _Snyk_ and select **Download** to download the Snyk Security - Code and Open Source Dependencies extension.

Once installed, use Snyk via the **Extensions > Snyk** menu (on Visual Studio versions older than  2019, Snyk will be part of the top menu bar).

__![](<../../../.gitbook/assets/image (76).png>)__

You can also open the Snyk tool window using **View > Other Windows > Snyk**_._

Once the tool window opens, wait while the Snyk extension downloads the latest Snyk CLI version.

![Snyk tool window, CLI downloading](../../../.gitbook/assets/readme\_image\_2\_3.png)

After you install the extension and the CLI you must authenticate. You can use the **Connect Visual Studio to Snyk** link. For more information see [Authentication](visual-studio-extension.md#authentication).

## Configuration

To analyze projects the plugin uses the Snyk CLI, which requires environment variables:

* `PATH`: specify the path to needed binaries, for example, to Maven
* `JAVA_HOME`: specify the path to the JDK you want to use for analysis of Java dependencies
* `http_proxy` and `https_proxy`: set if you are behind a proxy server, using the value in the format `http://username:password@proxyhost:proxyport`\
  **Note:** the leading `http://` in the value does not change to `https://` for `https_proxy`

You can set the variables using the GUI or on the command line using the `setx` tool.

## **Authentication**

Authenticate using **Connect Visual Studio to Snyk** link on Overview page.

![Connect Visual Studio to Snyk](../../../.gitbook/assets/readme\_image\_2\_4.png)

You can also authenticate using Options. Open Visual Studio **Options** and go to the **General Settings** of the Snyk extension or use the **Settings** button in the toolbar.

![Options and settings button](../../../.gitbook/assets/readme\_image\_2\_5.png)

If the automated method does not work, you can trigger authentication by pressing the **Authenticate** button or enter the user API token manually. You can also submit a request to [Snyk support](https://snyk.zendesk.com/agent/dashboard).

![Token field and Authenticate button](../../../.gitbook/assets/readme\_image\_2\_6.png)

![Click the  Authenticate button or enter your API token](../../../.gitbook/assets/install-5-a.png)

On the Snyk website, verify your identity and connect to the IDE extension. Click the **Authenticate** button.

![](../../../.gitbook/assets/install-6.png)

Once the authentication has been confirmed, close the browser and go back to the IDE extension. The Token field has been populated with the authentication token and authentication is complete.

![Token filed populated with the authentication token](../../../.gitbook/assets/readme\_image\_2\_8.png)

## Run analysis

Open your solution and run Snyk scan. Depending on the size of your solution and the time needed to build a dependency graph, it takes less than a minute to a couple of minutes to get the vulnerabilities.

The extension provides the user with two kinds of results:

* Open Source vulnerabilities
* Snyk Code issues

### Open Source vulnerabilities

* Note that your solution will have to be built successfully in order to allow the CLI to pick up the dependencies and find the vulnerabilities.
* If you see only npm vulnerabilities or vulnerabilities that are not related to your C#/.NET projects, that can mean your project was not built successfully and was not detected by the CLI. If you have difficulty or questions, submit a request to [Snyk support](https://snyk.zendesk.com/agent/dashboard).

![Run scan](../../../.gitbook/assets/readme\_image\_3\_1\_1.png)

![Open Source vulnerabilities](../../../.gitbook/assets/readme\_image\_3\_1\_2.png)

### Snyk Code issues

Snyk Code analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue and examine the Snyk suggestion information in the panel.

![Snyk suggestion panel](../../../.gitbook/assets/readme\_image\_3\_1\_3.png)

The Snyk suggestion panel shows the argumentation of the Snyk engine using, for example, variable names of your code and the line numbers in red. You can also see:

* Links to external resources to explain the bug pattern in more detail.
* Tags that were assigned by Snyk, such as Security (the issue found is a security issue), Database (the issue is related to database interaction), or In Test (the issue is within the test code).
* Code from open source repositories that can be of help to see how others have fixed the issue.

## View analysis results

You can filter vulnerabilities by name or by severity.

Filter by name by typing the name of the vulnerability in the search bar.

![Filter by name](../../../.gitbook/assets/readme\_image\_3\_2\_1.png)

Filter by severity by selecting one or more of the severities when you open the search bar filter.

![Filter by severity](../../../.gitbook/assets/readme\_image\_3\_2\_2.png)

Users can configure Snyk extension by **Project settings**.

Note that the “Scan all projects” option is enabled by default. It adds the `--all-projects` option for Snyk CLI. This option scans all projects by default.

![Scan all projects enabled](../../../.gitbook/assets/readme\_image\_3\_3.png)

## Extension configuration

After the plugin is installed, you can set the following configurations for the extension:

* **Token**: Enter the token the extension uses to connect to Snyk. You can manually replace it, if you need to switch to another account.
* **Custom endpoint**: Specify the custom Snyk API endpoint for your organization.
* **Ignore unknown CA**: Ignore unknown certificate authorities.
* **Organization**: Specify the ORG\_NAME to run Snyk commands tied to a specific organization.\
  This setting allows you to specify an organization slug name to run tests for that organization. The value must match the URL slug as displayed in the URL of your org in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`. If not specified, preferred organization as defined in your [web account settings](https://app.snyk.io/account) is used to run tests.
* **Send usage analytics**: To help Snyk improve the extension, let your Visual Studio send Snyk information about how the extension is working.
* **Project settings**: Specify any additional Snyk CLI parameters.
* **Scan all projects**: Auto-detect all projects in the working directory, enabled by default.

In the settings, you can also choose which results you want to receive:

* Open Source vulnerabilities
* Snyk Code Security vulnerabilities
* Snyk Code Quality issues



## Known issue

**Could not detect supported target files**

**Solution** Open Visual Studio Options to go to the Project Settings of the Snyk extension and check Scan all projects.

![](../../../.gitbook/assets/readme\_image\_4\_1.png)

## How tos

### How to find the log files

Logs can be found in the user AppData directory:

```
%HOMEPATH%\AppData\Local\Snyk\snyk-extension.log
```

### Build process

Clone this repository locally:

```
git clone https://github.com/snyk/snyk-visual-studio-plugin.git
```

Restore Nuget packages:

```
nuget restore
```

Run build:

```
msbuild -t:Build
```

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
