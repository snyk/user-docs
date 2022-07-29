# Visual Studio Code extension

{% hint style="warning" %}
The Snyk Visual Studio Code extension is available for installation on the [Visual Studio code marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner).
{% endhint %}

## **Capabilities**

Snyk Visual Studio Code extension supports comprehensive security for proprietary code and open-source dependencies, providing a Java vulnerability scanner and a custom code vulnerability scanner or an open-source security scanner.

Snyk can scan for vulnerabilities in seconds from within your integrated development system, showing context on the issue, the impact, and fix guidance. Snyk returns security analysis of your code and open source dependencies with the results categorized by issue type and severity. For open source projects Snyk provides algorithm-based fix suggestions for both direct and transitive dependencies.

## Supported languages, package managers, and frameworks

* For Snyk Open Source, the VS Code extension supports all the languages and package managers supported by Snyk Open Source and the CLI. See the full [list](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code, the VS Code extension supports all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine).

## Install the extension

You can find the Snyk extension in the [Visual Studio Code Marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner). To install:

* Open the settings or preferences in your IDE.
* Navigate to the [Snyk Extension on the Visual Studio Code Marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner) and click **Install**.\
  For more information see the [installation instructions](https://code.visualstudio.com/docs/editor/extension-marketplace#\_install-an-extension).
* Configure the Snyk CLI; see [Configuration](./#configuration).
* Authenticate with Snyk; see [Authentication](./#authentication).
* Navigate back to the IDE; the first scan starts automatically.



## Run analysis

In the IDE note that the extension is already picking up the files and uploading them for analysis.

Snyk Open Source requires the Snyk CLI, so it downloads in the background.

Snyk Code analysis runs quickly without the CLI, so results may already be available. Otherwise, you see the following screen while Snyk scans your workspace for vulnerabilities:

![Snyk Code scan](<../../.gitbook/assets/image (231) (1).png>)

Snyk analysis runs automatically when you open a folder or workspace.

* Snyk Code performs scans automatically on file saves.
* Snyk Open Source does not automatically run on save by default, but you can enable it in settings:

![Snyk Open Source settings](<../../.gitbook/assets/image (176) (1) (1).png>)

**Tip**: if you do not like to manually save while working, enable [AutoSave](https://code.visualstudio.com/docs/editor/codebasics#\_save-auto-save).

## Rescan

To manually trigger a scan, either Save or manually rescan using the rescan icon:

![Rescan icon](<../../.gitbook/assets/image (61) (1) (1).png>)

If you only need the Code Quality, Code Security, or Open Source Security portion of the findings, you can easily disable the feature with the results you do not want to see or collapse the view:

![Configure Features](../../.gitbook/assets/configure-features.png)

## Snyk Code advanced mode

Snyk Code has "Advanced" mode that allows you to control how scan is performed.

To manually perform the analysis, in the configuration of the extension you can enable Advanced Mode which enables you to control the scanning process:

![Advanced Mode](../../.gitbook/assets/run-analysis\_advanced-mode.png)

## View analysis results

Snyk analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue. Snyk suggestion information for the issue selected appears in a panel on the right side:

![Snyk suggestion information](<../../.gitbook/assets/image (243) (1) (1).png>)

### Snyk analysis panel

The Snyk analysis panel on the left of the preceding code screen shows how much time the analysis took plus a list of issues with the suggestions found for them.

The icons have the following meaning:

| ![](<../../.gitbook/assets/image (201) (1).png>) Critical severity | May allow attackers to access sensitive data and run code on your application.                                                               |
| ------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| ![](<../../.gitbook/assets/image (203) (1) (3).png>) High severity | May allow attackers to access sensitive data on your application.                                                                            |
| ![](<../../.gitbook/assets/image (180) (2).png>) Medium severity   | May allow attackers under some conditions to access sensitive data on your application.                                                      |
| ![](<../../.gitbook/assets/image (114) (1).png>) Low severity      | The application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application. |

You can filter the issues by setting the severities you want to see using the `snyk.severity` setting. For example, set `"snyk.severity": { "critical": true, "high": true, "medium": true, "low": false }` to hide low severity issues. You can also apply the setting in the Settings UI.

![Severity settings](<../../.gitbook/assets/image (65) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>)

### Snyk Code editor window

The editor window in the middle of the results screen shows the code that is inspected. This ensures that when you are inspecting a Snyk issue, you always have the code context close to the issue.

### Snyk Code vulnerability window

![Snyk Suggestion panel](<../../.gitbook/assets/image (243) (1) (6).png>)

The Snyk Suggestion panel on the right of the results screen shows the recommendation of the Snyk engine using, for example, variable names of your code and the line numbers in red. You can also see the following:

* Links to external resources to explain the bug pattern in more detail (the **More info** link).
* Tags that were assigned by Snyk, such as **Security** (the issue found is a security issue), **Database** (the issue is related to database interaction), or **In Test** (the issue is within the test code).
* Code from open source repositories that might be of help to see how others fixed the issue.
* Two buttons on the lower end of the panel which you can use to add ignore comments that would make Snyk ignore this particular suggestion, or all of these suggestions for the whole file. .

Snyk also includes a feedback mechanism to report false positives so others do not see the same issue.

### Snyk Open Source editor window

The editor window shows security vulnerabilities in open source modules while you code in JavaScript, TypeScript, and HTML. Receive feedback in-line with your code, such as how manyvulnerabilities a module contains that you are importing. The editor exposes only top-level dependency vulnerabilities; for the full list of vulnerabilities refer to the side panel.

You can find security vulnerabilities in the npm packages you import and see the number of known vulnerabilities in your imported npm packages as soon as you require them:

![Vulnerabilities in npm package](../../.gitbook/assets/oss-editor-vulnerability-count.png)

Code inline vulnerability counts are also shown in your `package.json` file:

![package.json file](../../.gitbook/assets/oss-editor-pjson.png)

Find security vulnerabilities in your JavaScript packages from well-known CDNs. The extension scans any HTML files in your projects and displays vulnerability information about the modules you include from your favorite CDN.

* Currently supported CDNs are:
  * unpkg.com
  * ajax.googleapis.com
  * cdn.jsdelivr.net
  * cdnjs.cloudflare.com
  * code.jquery.com
  * maxcdn.bootstrapcdn.com
  * yastatic.net
  * ajax.aspnetcdn.com

![Vulnerability from a CDN](../../.gitbook/assets/oss-editor-html.png)

You can navigate to the most severe vulnerability by triggering the provided code actions. This opens a vulnerability window to show more details:

![Code action](<../../.gitbook/assets/oss-editor-show-vulnerability (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

### Snyk Open Source vulnerability window

The Open Source Security (OSS) vulnerability window shows information about the vulnerable module.

* Links to external resources (CVE, CWE, Snyk Vulnerability DB) to explain the vulnerability in more detail.
* Displays CVSS score and exploit maturity.
* Provides detailed path on how vulnerability is introduced to the system.
* Shows summary of the vulnerability together with the remediation advice to fix it.

##

##

## Troubleshooting

To obtain extension logs, navigate to **Snyk:  Show Output Channel** in the VS Code command palette.&#x20;

Additional logs can be obtained by navigating to **Help** > **Toggle Developer Tools**, and opening the **Console** tab for more verbose output.

## Support and contact information

{% hint style="info" %}
Need more help? [Submit a request to Snyk support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

**Share your experience.**

Snyk continuously strives to improve the Snyk plugins experience. Would you like to share your feedback about the Snyk Visual Studio Code extension? [Schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
