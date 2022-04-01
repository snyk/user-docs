---
description: Use this documentation to get started with the Eclipse plugin.
---

# Eclipse plugin

Install the **Snyk Security - Code and Open Source Dependencies** extension in your Eclipse workflow to expose vulnerabilities and license issues directly from within your IDE (integrated development environment), before they are merged into your codebase.

After you have installed and configured the Eclipse plugin, every time you run it Snyk scans the manifest files in your project and takes the following action:

* analyzes and delivers actionable vulnerability and license issue details
* records results per package
* displays results directly from the Eclipse UI

The Snyk Eclipse plugin enables you to track and identify issues that are risks to the security of your application and avoid ever adding those issues to your shared repo.

{% hint style="info" %}
The Snyk Eclipse plugin is available for installation on the [Eclipse Marketplace](https://marketplace.eclipse.org/content/snyk-security-scanner).
{% endhint %}

## Supported languages and repos

Snyk supports all languages that are supported by both Eclipse and Snyk. Additionally, the Snyk plugin can be implemented with the Snyk Broker and on-prem solutions.

## Installing the Snyk Eclipse plugin

1. Navigate to the Marketplace from within your running Eclipse instance.
2. Search for Snyk and click **Install**.
3. When prompted accept the license agreement add the **Snyk Security** certificate to complete the installation.
4. Restart the Eclipse instance and navigate to **Eclipse Preferences** to ensure **Snyk Security - Code and Open Source Dependencies** now appears in the list:

![Eclipse Marketplace search showing Snyk plugin and Install button](../../../.gitbook/assets/uuid-01198b42-f020-2cc5-c20f-93817eeb44a4-en.png)

## Configuration

To use Snyk, you must supply the plugin with environment variables and the Snyk Token.

### API token

To provide your API token, copy it from your [account settings](https://app.snyk.io/account) and paste it into the Eclipse preferences Snyk API Token field. Click **Apply and Close** to start analyzing.

![Eclipse preferences with Snyk API Token filed and Apply and Close button](../../../.gitbook/assets/uuid-928012b7-8e49-fe6f-4965-77c5db026784-en.png)

### Environment variables

To analyze projects, the plugin uses the Snyk CLI. The CLI needs the following environment variables:

* `PATH`: the path to needed binaries, for example, to Maven. The `PATH` variable can also be manually adjusted using the `Path` field in the settings dialog
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies
* `http_proxy` and `https_proxy`: set using the value in the format `http://username:password@proxyhost:proxyport`.\
  **Note:** the leading `http://` in the value does not change to `https://` for `https_proxy`.

Setting these variables only in a shell environment (for example, using **\~/.bashrc**) is not enough, if you do not start Eclipse from the command line or create a script file that starts Eclipse using a shell environment.

* On **Windows**, set the variables using the GUI, or on the command line using the [setx](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/setx) tool.
* On **macOS**, the process `launchd` needs to know the environment variables to launch Eclipse directly from Finder. Set these environment variables using the `launchctl setenv` command (for example, on start up or using a script you launch at user login).\
  **Note:** Provision of environment variables to the macOS UI may change between operating system releases, so it can be easier to create a small shell script that launches the Eclipse app to leverage the shell environment, that can be defined via `~/.bashrc`.
* On **Linux**, updating the file `/etc/environment` can be used to propagate the environment variables to the windows manager and UI.

## Use the Snyk plugin to secure your Eclipse projects

From the Snyk results click the green arrow (![](../../../.gitbook/assets/uuid-aa090aa8-d4fe-eb5d-2505-54a0b1555be9-en.png)) whenever you are ready to scan your projects. Results appear in a short time, and your continue to work as usual in the meantime.

If for any reason you need to stop the scan before the build ends, click the red square(![](../../../.gitbook/assets/uuid-29be01e6-6913-25f8-15ed-a8cf47230fa0-en.png)). If you want to scan only a single project in your workspace, navigate to the Package Explorer panel, right click the root of the project you want to test, and then choose **Snyk test**.

When the scan ends, results and any relevant error messages are displayed from the **Snyk results**, grouped by project similar to the following:

![Results screen for scan with Eclipse plugin](../../../.gitbook/assets/uuid-e868f739-eb55-9bd6-be33-acbb230ec1fa-en.png)

## Work with Snyk results for Eclipse projects

### **Context menu**

Right click menu

Options include:

**Ignore issue**—Hover over the specific issue that you want to ignore for the next 30 days and then access the context menu.

**Snyk test**—Run the Snyk test for the entire workspace.

**Preferences**—Access and update Snyk Vuln Scanner preferences directly from the right click menu.

### **When collapsed**

**Title:** The name of the project.

**Dependency:** A summary of vulnerabilities and the number of affected paths found for each project.

### When expanded

**Title:** The full name of the vulnerability affecting your project, linked to a description and complete details of the vulnerability in the Snyk database, to assist you in resolving the issue.

**Dependency:** The name of the direct dependency package in your project (the package you explicitly installed) that is affected by the vulnerability, either directly or indirectly.

All details appear on a single row and the Dependency (the name of the package explicitly used in the code) and Package (the name of the package that actually contains the vulnerability) columns both display the name of the same package:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-4cdd086d6be47b598fc1a9a52c63023d59cff825%2Fuuid-e7accdc1-7495-e7a5-7a64-2403b066cb03-en.png?alt=media&token=e3bf024a-ba92-4b76-87be-b728d7edf092" %}
Eclipse results details
{% endembed %}

An arrow appears on the row, grouping together all relevant details, similar to the following examples:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
Eclipse results arrow on row grouping details
{% endembed %}

**Dependency when your project is affected by an indirect vulnerability, collapsed mode:**

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
Collapsed mode, indirect vulnerability
{% endembed %}

Example:

Package X uses Package Y, which in turn uses Package Z.

Package Z contains a Cross-Site Scripting (XSS) vulnerability, indirectly affecting your project.

The Dependency (the name of the package explicitly used in the code) is Package X; the Package field displays Package Z (the name of the package that actually contains the vulnerability).

**Dependency when your project is affected by an indirect vulnerability, expanded mode:**

Click the arrow on the row to expand and view the full path from the direct dependency to the vulnerable package.

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-992b169b89e7f3c45782fdeb47b205e3c0a95af8%2Fuuid-35658aaf-3359-80c2-c094-41a34c7863cc-en.png?alt=media&token=53c91ccc-f9bc-4ba7-a55f-8def3aa50d86" %}
Expanded mode, indierct vulnerability
{% endembed %}

On the preceding screen the full path would appear as:

\[Name of Package X]-->\[Name of Package Y]-->\[Name of Package Z]

**Package:** The name of the package in your project that is directly affected by the vulnerability. On the preceding screen:

* The Dependency is indicated as Package X—this is the package the developer explicitly uses in the code
* the Package field displays Package Z, which is the package that contains the vulnerability.

**Fix:** The name of the package if any and the version that it can be upgraded to in order to resolve the issue.

## Support / Contact

{% hint style="info" %}
If you need help, submit a [request](https://support.snyk.io/hc/en-us/requests/new) to Snyk Support.
{% endhint %}

## Share your experience

Snyk continuously strives to improve the Snyk plugins experience. If you would you like to share your feedback about the Snyk Eclipse plugin, [schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
