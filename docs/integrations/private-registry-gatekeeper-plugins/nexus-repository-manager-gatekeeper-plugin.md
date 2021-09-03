# Nexus Repository Manager Gatekeeper plugin

### **Nexus Gatekeeper plugin: overview**

{% hint style="info" %}
**Feature availability**  
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Install Snyk plugin directly on the Nexus instance to track open source vulnerabilities and license details in your artifacts based on your configurations.

Once installed, Snyk runs in the background and whenever a download is requested from the developer's CLI, Snyk automatically:

* scans artifacts for licenses and for vulnerabilities and delivers remediation data for known vulnerabilities found in the artifact
* blocks developers from downloading vulnerable packages—based on Snyk results and your severity threshold configurations

By scanning artifacts as part of your work and then displaying those test results, the Snyk plugin enables developers transparency to the risks associated with their packages, and enables administrators to more quickly track and identify issues that risk your application’s security and avoid using those artifacts in your organization’s projects.

**Supported languages and package managers**

Snyk supports all languages that are supported by both Nexus and the Snyk API. See [our docs](https://snyk.docs.apiary.io/#reference/test) for more information about the package managers and languages we support.

**How it works**

Use the Snyk capability with your Nexus Repository Manager to test your artifacts for vulnerabilities and license issues at installation of this plugin and also every time developers request to download any one of them.

1. The admin installs and updates the Snyk configurations on the Nexus instance from the **Capabilities** section, including the authentication token and organization ID from Snyk.
2. Snyk authenticates the account configuration using the API token and Organization ID that the admin entered, and then runs in the background continuously.
3. A developer attempts to download an artifact from the Nexus instance to their local environment.
4. Based on the severity thresholds configured for the Snyk Security Configuration capability, the package is blocked.
5. The error appears in the developer’s CLI \(including a link to the error with full details\) and from the Nexus interface for administrators, detailing the number of known issues for vulnerabilities and licenses.

### Administrators’ guide for the Snyk Nexus Gatekeeper plugin

To set up and configure the plugin, start scanning and managing your organization's artifacts for vulnerabilities, check out this information:

#### Prerequisites:

* You must sign up for an Enterprise account with Snyk before you begin.
* You must be an administrator or owner of the Snyk account.
* You must have one of the following installed for your team first:
  * Nexus Repository Manager OSS
  * Nexus Repository Manager Pro v3.15.0 or greater

### Install the plugin on your Nexus server

1. Download the bundle from [our GitHub repository](https://github.com/snyk/nexus-snyk-security-plugin/releases).
2. Copy the `nexus-snyk-security-plugin.kar` file from the bundle onto your Nexus server at`/deploy`.
3. From the Nexus interface, enable the Snyk Security Configuration from the Capabilities area. For more information about this, see the [Sonatype documentation](https://help.sonatype.com/repomanager2/configuration/accessing-and-configuring-capabilities).

### **Configure the capability**

1. Go to your Snyk account to copy and save your personal API token or your service account token, and an **Organization ID**. Both a token and an organization ID are mandatory and must be configured in order for Snyk to authenticate your account. Because this plugin does not import any data to Snyk, you can use any of your organization IDs.
2. From your Nexus instance, navigate to the Capabilities section and select to edit the **Snyk Security Configuration** from the list.
3. Ensure **Enable this capability** is checked, and enter details for the remaining fields as follows:
   * **Snyk API URL** - enter the API endpoint if you’re setting up on-prem or Broker for Snyk
   * **Snyk API token** - paste the token value you saved from step 1
   * **Snyk Organization ID** - paste the token value you saved from step 1
   * **Vulnerability Threshold**—default is \*low\*. Valid values include low, medium, high. Manually update the configuration based on your needs.
   * **License Threshold**—default is \*low\*. Valid values include low, medium, high. Manually update the configuration based on your needs. 
4. Shut down the Nexus service instance and then restart it. 
5. Log in to your Nexus instance and double check that the Snyk bundle has been installed successfully.

![](../../.gitbook/assets/uuid-9745b82a-ed7e-bce0-75dd-0070514f274d-en.png)

### Track vulnerabilities in your team's artifacts

1. Once installed, every time a developer requests to download an artifact, the following occurs:
   * Based on the severity thresholds that you configured, the download is blocked.
   * Scan results are displayed for the developer with a link to full details for the error:
   * Results are stored in the Snyk Security part of the Attributes section from the Nexus interface for the artifact:

![](../../.gitbook/assets/uuid-a2c354a2-21ca-bdfb-7862-a2ef26eec59e-en.png)

![](../../.gitbook/assets/image%20%2833%29.png)

By reviewing the results, you can evaluate the issues found in your artifact and determine a course of action.

Work with Snyk properties from Nexus as follows:

| **Property** | **Description** |
| :--- | :--- |
| issues\_licenses | Regardless of the thresholds configured, this row displays license summary scan results. |
| issues\_url | This is the URL to our database and explanation of the vulnerability, including specific details about vulnerable versions, available upgrades and Snyk patches as well. |
| issues\_vulnerabilities | Regardless of the thresholds configured, this row displays vulnerability summary scan results. |

#### Troubleshooting

If your connection to Snyk is unsuccessful, try checking the following points or contact us at snyk@support.io \(`<`[`snyk@support.io`](mailto:snyk@support.io)`>`\):

* Check Nexus logs for any related errors.
* Ensure you’ve entered the API URL correctly for the configuration of the capability.
* For on-prem and Broker configurations, ensure the Snyk service is running.

### **Use Snyk in your build**

Snyk continuously runs in the background on the Nexus instance, and whenever any team member requests a download, Snyk automatically scans the artifact to evaluate vulnerabilities and license issues and blocks the request based on your administrator’s configurations.

When the scan ends, you get an error message if the download was blocked, with a link to the full details for the known vulnerabilities found in that artifact, similar to the following example:

![](../../.gitbook/assets/uuid-a2c354a2-21ca-bdfb-7862-a2ef26eec59e-en.png)

Click the link to view details, itemizing the number of vulnerabilities found in the artifact, and their severities.

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}