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

After you have installed and configured the Eclipse plugin, every time you run it, open a file, or  autosave, Snyk scans the manifest files, proprietary code, and configuration files in your project. Snyk delivers actionable vulnerability, license, code quality, or misconfiguration issue details and displays the results natively within the Eclipse UI.

{% hint style="info" %}
The Snyk Eclipse plugin is available for installation on the [Eclipse Marketplace](https://marketplace.eclipse.org/content/snyk-security-scanner).
{% endhint %}

## Supported languages, package managers, and frameworks

* For Snyk Open Source, the Eclipse plugin supports all the languages and package managers supported by Snyk Open Source and the CLI. See the full [list](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code, the Eclipse plugin supports all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine).
* For Snyk IaC, the Eclipse plugin supports the following IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager.

## Installing the Snyk Eclipse plugin

Navigate to the Marketplace from within your running Eclipse instance. Search for Snyk and click **Install**.

![Eclipse Marketplace search showing Snyk plugin and Install button](<../.gitbook/assets/Screenshot 2022-05-17 at 16.29.29.png>)

When prompted accept the license agreement add the **Snyk Security** certificate to complete the installation (this happens only once).

![Add Snyk Security certificate](<../.gitbook/assets/Screenshot 2022-05-13 at 09.08.52.png>)

Restart the Eclipse instance:

![Restart Eclipse](<../.gitbook/assets/Screenshot 2022-05-13 at 09.16.37.png>)

Once Eclipse is restarted, navigate to **Eclipse Preferences** to ensure that **Snyk** now appears in the list:

![Eclipse preferences showing Snyk](<../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

## First run after the installation

The Snyk Eclipse plugin now works with an underlying language server for optimal Eclipse experience. After restart, on opening a file that Snyk supports, the Eclipse plugin attempts to start a workspace scan:

![Eclipse plugin starting a scan](<../.gitbook/assets/Screenshot 2022-05-13 at 09.28.30 (1).png>)

Once the plugin is installed, all of the plugin's prerequisites are triggered on opening a file that Snyk supports (Snyk hooks into this action). The prerequisites include downloading the Snyk CLI, the language server, and asking you to authenticate. All of these steps are shown in the following sections, in the order they happen.

### Snyk CLI and language server download

The process downloads the Snyk CLI and the language server and uses the CLI for authentication.

![Download the Snyk CLI](<../.gitbook/assets/Screenshot 2022-05-13 at 11.27.00.png>)

### Authentication

Once the CLI is downloaded you will be redirected to the browser to authenticate and then connect the IDE with your account. Steps are as follows:

* You get a notification that a browser window will open.

![Notification, browser window opening for authentication](<../.gitbook/assets/Screenshot 2022-05-13 at 11.29.37.png>)

* Once you are redirected to the browser for authentication, click Authenticate.

![Authenticate](<../.gitbook/assets/Screenshot 2022-05-13 at 11.30.02.png>)

* You should see a successful message saying you've been authenticated.

![Confirmation of authentication](<../.gitbook/assets/Screenshot 2022-05-13 at 11.30.30 (1).png>)

* Going back to the IDE, you should see a confirmation that the IDE has been successfully connected and the API token has been securely stored.

![Confirmation of connection](<../.gitbook/assets/Screenshot 2022-05-13 at 11.30.54.png>)

To verify that the Eclipse plugin is ready to start scanning, be sure the Snyk preferences show the following after the downloads and the authentication are done:

* Language server path that the Eclipse plugin will use
* Snyk CLI path that the Eclipse plugin will use
* Snyk API token, securely stored through the Eclipse's secure storage mechanism

![Snyk preferences after downloads and authentication](<../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

### Configure the API token manually

You can provide the API token by copying it from your [account settings](https://app.snyk.io/account) and paste it into the Eclipse preferences Snyk API Token field. Click **Apply and Close** once the token has been set**.**

![Providing the API token manually](<../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

### Configuration

In the Snyk preferences, the following configuration options are available.

* `Snyk API Token`: The authentication token from Snyk.
* `Path`: Your additions to the path to find needed third party tools such as Gradle or Maven.
* `Custom Endpoint`: The custom endpoint for Single Tenant setups instead of https://app.snyk.io.
* `Allow unknown certificate authorities`: Disable certificate checks for SSL connections.
* `Custom Snyk LS Path`: Specify the location of your language server binaries. If set, no updates are downloaded and updates must be performed manually to synchronize Eclipse features and preferences with the Language Server.
* `Snyk Open Source enabled`: Enable/Disable Snyk Open Source Dependency Scans via Language Server. Default: `Enabled` during beta
* `Snyk Code enabled`: Enable/Disable Snyk Code Scans via Language Server. Default: `Disabled` during beta.
* `Snyk Infrastructure-as-Code enabled` : Enable/Disable Snyk IaC Scans via Language Server. Default: `Enabled` during beta.
* `Organization`: Specify the Snyk Organization to use for scanning.
* `Additional Parameters`: Specify additional parameters to pass to the CLI (for example, `--all-projects` or `-d.`
* `Additional Environment`: Add environment variables to Language Server.
* `Send error reports to Snyk`: Send errors from Language Server to Snyk to enable quicker bug fixing. Default: `Enabled` during beta.
* `Send usage statistics to Snyk`: Allow Snyk to get usage data to improve workflows. Default: `Disabled` during beta.

### Environment variables

To analyze projects, the plugin uses the Snyk CLI. The CLI needs the following environment variables:

* `PATH`: The path to needed binaries, for example, to Maven. The `PATH` variable can also be manually adjusted using the **`Path`** field in the settings dialog.
* `JAVA_HOME`: The path to the JDK you want to use to analyze Java dependencies
* `http_proxy` and `https_proxy`: Proxy set using the value in the format `http://username:password@proxyhost:proxyport`.\
  **Note:** the leading `http://` in the value does not change to `https://` for an `https_proxy`.

Setting these variables only in a shell environment (for example, using **\~/.bashrc**) is not enough, if you do not start Eclipse from the command line or create a script file that starts Eclipse using a shell environment.

* On **Windows**, set the variables using the GUI or on the command line using the [setx](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/setx) tool.
* On **macOS**, the process `launchd` needs to know the environment variables to launch Eclipse directly from the Finder. Set these environment variables using the `launchctl setenv` command (for example, on start up or using a script you launch at user login).\
  **Note:** Provision of environment variables to the macOS UI may change between operating system releases, so it can be easier to create a small shell script that launches the Eclipse app to leverage the shell environment. This can be defined via `~/.bashrc`.
* On **Linux**, updating the file `/etc/environment` can be used to propagate the environment variables to the windows manager and UI.

{% hint style="info" %}
If you need to use a **proxy server** to connect to the internet, configure it using the **environment variables**.
{% endhint %}

## Use the Snyk plugin to secure your Eclipse projects

Once the language server is downloaded and the authentication is done, the plugin will successfully start the workspace scan. You might notice a confirmation that a workspace scan is starting. Snyk shows such a notification when there is no workspace scan available.

![Starting workspace scan](<../.gitbook/assets/Screenshot 2022-05-13 at 11.55.41.png>)

All of the issues found by Snyk are now natively integrated with Eclipse's flows. Issues are shown in the Problems tab (see the following screenshot). There is a squiggly line indicating the issue while you code plus the gutter icons to indicate where the issue is.

![Problems tab](<../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png>)

## SAST scanning results (SAST, Snyk Code)

Starting version 2.0.0 and above, Snyk is introducing a deeper integration within Eclipse with the native flows of Eclipse (inline highlights, problems integrations, information about the issue on hover). The following shows all of these for a high severity security vulnerability found in a `js` file:&#x20;

1. The security vulnerability is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in your code. You will see the vulnerability ID, what the issue is, on hover. Snyk is updating the panel to include more information about the found Code issues, so you have even more details like data flow and example fixes.
2. You see the integration with the problems tab, which comes in handy if you use the problems tab to show only issues in the current file. Snyk also shows the line where the issue is.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer, etc.
{% endhint %}

![SAST scanning results](<../.gitbook/assets/Screenshot 2022-05-13 at 12.56.46.png>)

## Misconfiguration scanning results (Snyk Infrastructure as Code)

Starting version 2.0.0 and later, Snyk is introducing a deeper integration within Eclipse with the native flows of Eclipse (inline highlights, problems integrations, information about the issue on hover). The following shows all of these for a high severity misconfiguration found in a Terraform file:&#x20;

1. The misconfiguration is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in this file and the line number. You have all the information on hover; you can scroll, read, or click the links (when available) for even more information. Advice on how to resolve the misconfiguration is right there where the misconfiguration is.
2. You see the integration with the problems tab, which comes in handy if you use the problems tab to show only issues in the current file. Snyk also indicates the line where the issue is.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer, etc.
{% endhint %}

![Misconfiguration scanning results](<../.gitbook/assets/Screenshot 2022-05-13 at 12.59.40.png>)

## Third party dependency scanning (SCA, Snyk Open Source)

Starting version 2.0.0 and later, Snyk is introducing a deeper integration within Eclipse with the native flows of Eclipse (inline highlights, problems integrations, information about the issue on hover). The following shows all of these for a security vulnerability found in a third party dependency:&#x20;

1. The vulnerable package is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in this package. You have all the information on hover; you can scroll, read, or click the links for even more information. Advice on what action to take and how is right there where the vulnerability is.
2. You see the integration with the problems tab, which comes in handy if you use the problems tab to show only issues in the current file. Snyk also indicates the line where the issue is.
3. You can see the gutter icons on the left, as well as the file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer, etc.
{% endhint %}

![Third party dependency scanning results, problems tab](<../.gitbook/assets/Screenshot 2022-05-13 at 13.01.53.png>)

Third party dependency scanning results are **also available in the already existing Snyk Results panel**:

![Third party dependency scanning results, results panel](<../.gitbook/assets/Screenshot 2022-05-13 at 12.13.37.png>)

### **Context menu**

Right click menu options include:

**Ignore issue**—Hover over the specific issue that you want to ignore for the next 30 days and then access the context menu.

**Snyk test**—Run the Snyk test for the entire workspace.

**Preferences**—Access and update Snyk Vuln Scanner preferences directly from the right click menu.

### **Snyk View when collapsed**

**Title:** The name of the project.

**Dependency:** A summary of vulnerabilities and the number of affected paths found for each project.

### Snyk View when expanded

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

**Example:**

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

### My Snyk Results panel is not visible

If you close the Snyk Results panel by accident, or for some reason you don't see it, here is how to enable it:

Navigate to **Windows ->  Show View -> Other...**

![Show View, Other](<../.gitbook/assets/Screenshot 2022-05-13 at 12.04.07.png>)

Search for Snyk in the **Show View** dialog window.

![Show View dialog window](<../.gitbook/assets/Screenshot 2022-05-13 at 12.02.06 (2).png>)

You should now be able to see the Snyk Results panel:

![Snyk Results panel](<../.gitbook/assets/Screenshot 2022-05-13 at 12.02.18.png>)

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
