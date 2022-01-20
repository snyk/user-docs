---
description: Use this documentation to get started with the Eclipse plugin.
---

# Eclipse plugin

Install our **Snyk Vuln Scanner** in your Eclipse workflow to catch vulnerabilities and license issues directly from within your IDE (integrated development environment), before they are merged into your codebase.

Once installed and configured, every time you run the plugin, Snyk scans your project’s manifest files and:

* analyzes and delivers actionable vulnerability and license issue details
* records results per package
* displays results directly from the Eclipse UI

The Snyk plugin enables you to track and identify issues that risk your application’s security and avoid ever adding those issues to your shared repo.

{% hint style="info" %}
Snyk's Eclipse plugin is available for install on the marketplace: [https://marketplace.eclipse.org/content/snyk-security-scanner](https://marketplace.eclipse.org/content/snyk-security-scanner).
{% endhint %}

## Supported languages and repos

Snyk supports all languages that are supported by both Eclipse and Snyk. Additionally, the Snyk plugin can also be implemented with our Broker and on-prem solutions.

## Installing the Eclipse Snyk plugin

1. Navigate to the Marketplace from within your running Eclipse instance.
2. Search for Snyk and click **Install**.
3. When prompted accept the license agreement and the **Snyk Security** certificate to complete the installation.
4. Restart the Eclipse instance and navigate to **Eclipse Preferences** to ensure **Snyk Vuln Scanner** now appears in the list:

![](../../../.gitbook/assets/uuid-01198b42-f020-2cc5-c20f-93817eeb44a4-en.png)

![](../../../.gitbook/assets/uuid-928012b7-8e49-fe6f-4965-77c5db026784-en.png)

## Use the Snyk plugin to secure your Eclipse projects

From the Snyk results click ![](../../../.gitbook/assets/uuid-aa090aa8-d4fe-eb5d-2505-54a0b1555be9-en.png) whenever you are ready to scan your projects. It shouldn’t take too long for the results to appear—but no worries! You can continue to work as usual in the meantime anyway.

If for any reason you need to stop the scan before the build ends, click: ![](../../../.gitbook/assets/uuid-29be01e6-6913-25f8-15ed-a8cf47230fa0-en.png) If you only want to scan a single project in your workspace, navigate to the Package Explorer panel, right-click the root of the project you want to test, and then choose **Snyk test**.

When the scan ends, results and any relevant error messages as well, are displayed from the **Snyk results**, grouped by project similar to the following:

![](../../../.gitbook/assets/uuid-e868f739-eb55-9bd6-be33-acbb230ec1fa-en.png)

Work with Snyk results from Eclipse as follows:

| **Column**         |                                                                                                                           | **Description**                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Context menu**   | Right-click menu                                                                                                          | <p>Options include:</p><p>Ignore issue—Hover over the specific issue that you want to ignore for the next 30 days and then access the context menu.</p><p>Snyk test—Run the Snyk test for the entire workspace.</p><p>Preferences—Access and update Snyk Vuln Scanner preferences directly from the right-click menu.</p>                                                                                                                                                                                                                                                                                               |
| **When collapsed** |                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| **Title**          |                                                                                                                           | The name of the project.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **Dependency**     |                                                                                                                           | A summary of vulnerabilities and the number of affected paths found per project.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| **When expanded**  |                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| **Title**          |                                                                                                                           | The full name of the vulnerability affecting your project, linked to a description and complete details of the vulnerability in our database, to assist you in resolving the issue.                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **Dependency**     |                                                                                                                           | The name of the direct dependency package in your project (the package you explicitly installed) that is affected by the vulnerability, either directly or indirectly.                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
|                    |                                                                                                                           | <p>All details appear on a single row and the Dependency (the name of the package explicitly used in the code) and Package (the name of the package that actually contains the vulnerability) columns both display the name of the same package:</p><p><img src="../../../.gitbook/assets/uuid-e7accdc1-7495-e7a5-7a64-2403b066cb03-en.png" alt="image13.png"></p>                                                                                                                                                                                                                                                      |
|                    | <p><strong>When your project is affected by an indirect vulnerability:</strong></p><p><strong>Collapsed mode</strong></p> | <p>An arrow appears on the row, grouping together all relevant details, similar to the following examples:</p><p><img src="../../../.gitbook/assets/uuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png" alt="image14.png"></p><p>For example:</p><p>Package X uses Package Y, which in turn uses Package Z.</p><p>Package Z contains a Cross-Site Scripting (XSS) vulnerability, indirectly affecting your project.</p><p>The Dependency (the name of the package explicitly used in the code) is Package X; the Package field displays Package Z (the name of the package that actually contains the vulnerability).</p> |
|                    | **Expanded mode**                                                                                                         | <p>Click the arrow on the row to expand and view the full path from the direct dependency to the actual vulnerable package.</p><p><img src="../../../.gitbook/assets/uuid-35658aaf-3359-80c2-c094-41a34c7863cc-en.png" alt="image15.png"></p><p>In the example above, the full path would appear as:</p><p>[Name of Package X]-->[Name of Package Y]-->[Name of Package Z]</p>                                                                                                                                                                                                                                          |
| **Package**        |                                                                                                                           | <p>The name of the package in your project that is directly affected by the vulnerability.</p><p>In the example above:</p><ul><li>the Dependency is indicated as Package X—this is the package the developer explicitly uses in the code</li><li>the Package field displays Package Z, which is the package that actually contains the vulnerability.</li></ul>                                                                                                                                                                                                                                                         |
| **Fix**            |                                                                                                                           | The name of the package, if such exists, and the version that it can be upgraded to in order to resolve the issue.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |

### Support / Contact

For support and help, visit [Snyk IDE Plugins Help](../snyk-ide-plugins/help.md#a693dbb5-063c-46d4-9b4e-d21e73b1e485).

### Share your experience

We continuously strive to improve our plugins experience. Would you like to share with us your feedback about Snyk's Eclipse Plugin: [schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
