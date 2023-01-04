# View analysis results from Visual Studio Code extension

Snyk analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue. Snyk suggestion information for the issue selected appears in a panel on the right side:

![Snyk suggestion information](<../../.gitbook/assets/image (76) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

## Snyk analysis panel

The Snyk analysis panel on the left of the preceding code screen shows how much time the analysis took plus a list of issues with the suggestions found for them.

The icons have the following meaning:

| ![](<../../.gitbook/assets/image (75) (1) (1) (1).png>) Critical severity                                                                                             | May allow attackers to access sensitive data and run code on your application.                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| ![](<../../.gitbook/assets/image (10) (1) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (1).png>) High severity | May allow attackers to access sensitive data on your application.                                                                            |
| ![](<../../.gitbook/assets/image (116) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (6) (2).png>) Medium severity                  | May allow attackers under some conditions to access sensitive data on your application.                                                      |
| ![](<../../.gitbook/assets/image (376).png>) Low severity                                                                                                             | The application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application. |

You can filter the issues by setting the severities you want to see using the `snyk.severity` setting. For example, set `"snyk.severity": { "critical": true, "high": true, "medium": true, "low": false }` to hide low severity issues. You can also apply the setting in the Settings UI.

![Severity settings](<../../.gitbook/assets/image (234) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>)

## Snyk Code editor window

The editor window in the middle of the results screen shows the code that is inspected. This ensures that when you are inspecting a Snyk issue, you always have the code context close to the issue.

## Snyk Code vulnerability window

![Snyk Suggestion panel](<../../.gitbook/assets/image (76) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

The Snyk Suggestion panel on the right of the results screen shows the recommendation of the Snyk engine using, for example, variable names of your code and the line numbers in red. You can also see the following:

* Links to external resources to explain the bug pattern in more detail (the **More info** link).
* Tags that were assigned by Snyk, such as **Security** (the issue found is a security issue), **Database** (the issue is related to database interaction), or **In Test** (the issue is within the test code).
* Code from open source repositories that might be of help to see how others fixed the issue.
* Two buttons on the lower end of the panel which you can use to add ignore comments that would make Snyk ignore this particular suggestion, or all of these suggestions for the whole file. .

Snyk also includes a feedback mechanism to report false positives so others do not see the same issue.

## Snyk Open Source editor window

The editor window shows security vulnerabilities in open source modules while you code in JavaScript, TypeScript, and HTML. Receive feedback in-line with your code, such as how manyvulnerabilities a module that you are importing contains. The editor exposes only top-level dependency vulnerabilities; for the full list of vulnerabilities refer to the side panel.

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

![Code action](<../../.gitbook/assets/oss-editor-show-vulnerability (1) (1) (1) (2).png>)

## Snyk Open Source vulnerability window

The Open Source Security (OSS) vulnerability window shows information about the vulnerable module.

* Provides links to external resources (CVE, CWE, Snyk Vulnerability DB) to explain the vulnerability in more detail.
* Displays CVSS score and exploit maturity.
* Provides detailed path on how vulnerability is introduced to the system.
* Shows summary of the vulnerability together with the remediation advice to fix it.
