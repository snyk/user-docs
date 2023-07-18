# Nexus Repository Manager Gatekeeper plugin

## **Nexus Gatekeeper plugin: overview**

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Install the Snyk plugin directly on the Nexus instance to track open source vulnerabilities and license details in your artifacts based on your configurations.

After the plugin in installed, Snyk runs in the background and whenever a download is requested from the developer's CLI, Snyk automatically does the following:

* Scans artifacts for license issues and for vulnerabilities and delivers fix data for known vulnerabilities found in the artifact
* Blocks developers from downloading vulnerable packages based on Snyk results and your severity threshold configurations

By scanning artifacts as part of your work and then displaying those test results, the Snyk plugin enables developers to see to the risks associated with their packages, and enables administrators to track and identify issues that risk your application security more quickly and avoid using those artifacts in your projects.

### **Supported languages and package managers**

Snyk Nexus Plugin supports Maven and npm dependency scanning.

### **How it works**

Use the Snyk capability with your Nexus Repository Manager to test your artifacts for vulnerabilities and license issues at installation of this plugin and also every time developers request to download any one of the artifacts.

1. The admin installs and updates the Snyk configurations on the Nexus instance from the **Capabilities** section, including the authentication token and organization ID from Snyk.
2. Snyk authenticates the account configuration using the API token and Organization ID that the admin entered, and then runs in the background continuously.
3. When developer attempts to download an artifact from the Nexus instance to the local environment, based on the severity thresholds configured for the Snyk Security Configuration capability, the package is blocked.
4. The error appears in the developer’s CLI (including a link to the error with full details) and from the Nexus interface for administrators, detailing the number of known issues for vulnerabilities and licenses.

## Administrators’ guide for the Snyk Nexus Gatekeeper plugin

To set up and configure the plugin and start scanning and managing your organization's artifacts for vulnerabilities, follow the instructions in this guide.

### Prerequisites

* You must sign up for an Enterprise account with Snyk before you begin.
* You must be an administrator or owner of the Snyk account.
* You must have one of the following installed for your team:
  * Nexus Repository Manager OSS
  * Nexus Repository Manager Pro v3.15.0 or greater

### Install the plugin on your Nexus server

1. Download the bundle from the [Snyk Nexus security plugin GitHub repository](https://github.com/snyk/nexus-snyk-security-plugin/releases).
2. Copy the `nexus-snyk-security-plugin.kar` file from the bundle onto your Nexus server at`/deploy`.
3. From the Nexus interface, enable the Snyk Security Configuration from the **Capabilities** area. For more information about this, see the [Sonatype documentation](https://help.sonatype.com/repomanager2/configuration/accessing-and-configuring-capabilities).

### **Configure the capability**

1. Go to your Snyk account to copy and save your personal API token or your service account token, and an **Organization ID**. Both a token and an organization ID must be configured in order for Snyk to authenticate your account. Because this plugin does not import any data to Snyk, you can use any of your organization IDs.
2. From your Nexus instance, navigate to the **Capabilities** section and select the **Snyk Security Configuration** in order to edit it.
3. Ensure **Enable this capability** is checked, and enter details for the remaining fields as follows:
   * **Snyk API URL** - enter the API endpoint to use for Snyk requests
   * **Snyk API token** - paste the token value you saved from step 1
   * **Snyk Organization ID** - paste the token value you saved from step 1
   * **Vulnerability Threshold** - default is \*low\*. Valid values include none (will not block artifacts download), low, medium, high. Manually update the configuration based on your needs.
   * **License Threshold** - default is \*low\*. Valid values include none (will not block artifacts download), low, medium, high. Manually update the configuration based on your needs.
4. Shut down the Nexus service instance and then restart it.
5. Log in to your Nexus instance and check that the Snyk bundle has been installed successfully.

![Capabilities to configure](../../.gitbook/assets/uuid-9745b82a-ed7e-bce0-75dd-0070514f274d-en.png)

### Track vulnerabilities in your team's artifacts

1. Once installed, every time a developer requests download of an artifact, the following occurs:
   * Based on the severity thresholds that you configured, the download is blocked.
   * Scan results are displayed for the developer with a link to full details for the error.
   * Results are stored in the Snyk Security part of the **Attributes** section of the Nexus interface for the artifact.

![Results of a scan](../../.gitbook/assets/uuid-a2c354a2-21ca-bdfb-7862-a2ef26eec59e-en.png)

![Attributes showing results of a scan](<../../.gitbook/assets/image (33) (2).png>)

By reviewing the results, you can evaluate the issues found in your artifact and determine a course of action.

The following explains the Snyk properties from Nexus.

| **Property**            | **Description**                                                                                                                                                        |
| ----------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| issues\_licenses        | Regardless of the thresholds configured, this row displays license summary scan results.                                                                               |
| issues\_url             | This is the URL to the Snyk database and explanation of the vulnerability, including specific details about vulnerable versions, available upgrades, and Snyk patches. |
| issues\_vulnerabilities | Regardless of the thresholds configured, this row displays vulnerability summary scan results.                                                                         |

### Troubleshooting

If your connection to Snyk is unsuccessful, try checking the following points or [submit a request to Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

* Check Nexus logs for any related errors.
* Ensure you have entered the API URL correctly for the configuration of the capability.
* For Broker configurations, ensure the Snyk service is running.

## **Use Snyk in your build**

Snyk continuously runs in the background on the Nexus instance, and whenever any team member requests a download, Snyk automatically scans the artifact to evaluate vulnerabilities and license issues and blocks the request based on your configurations.

When the scan ends, you get an error message if the download was blocked, with a link to the full details for the known vulnerabilities found in that artifact, similar to the following example:

![Results a scan](../../.gitbook/assets/uuid-a2c354a2-21ca-bdfb-7862-a2ef26eec59e-en.png)

Click the link to view the number of vulnerabilities found in the artifact and their severities.
