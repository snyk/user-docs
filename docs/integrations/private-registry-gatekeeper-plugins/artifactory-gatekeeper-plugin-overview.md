# Artifactory Gatekeeper plugin overview

**Feature availability**  
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.

With the Snyk plugin for Artifactory, you can track open source vulnerabilities and license details in your cached artifacts.

Once installed, the Snyk plugin runs in the background and automatically:

1. Blocks devs from downloading packages with vuln/license issues according to a predefined threshold that the admin sets
2. Adds vuln and license data from Snyk as properties in artifact

### Note

 This article refers to the Artifactory _Plugin_, an independent piece of software that is installed on the Artifactory machine and serves as a gatekeeper, blocking vulnerable packages from being downloaded from the Artifactory instance rather than the [Artifactory _Integration_](https://support.snyk.io/hc/en-us/articles/360005507418-Artifactory-Registry-for-Maven) - an internal integration in Snyk app, that allows configuring SCM / CLI scans to use custom package registries.

By scanning artifacts as part of your workflow and then displaying those test results directly from the Artifactory UI, the Snyk plugin enables you to more quickly track and identify issues that risk your application’s security and avoid using those artifacts in your projects.

#### Artifactory Gatekeeper plugin: supported languages and repos

Our Artifactory plugin is supported for the following:

* npm
* Maven \(.jar files\)
* Gradle \(.jar files\)
* sbt \(.jar files\)
* pip

#### Install the Artifactory Gatekeeper plugin

Install or upgrade the Snyk Security plugin with these steps. Once complete, Snyk automatically scans your artifacts every time you request to download them.

* To install our plugin, first, download the archived \(.zip\) distribution of the Snyk Security Artifactory plugin as described in the steps below.

  This archive contains the following structure, files, and folders:

  ```text
  *plugins
  snykSecurityPlugin.groovy—our plugin
  snykSecurityPlugin.properties—the configuration file for the plugin
  *lib—this is the folder that contains the dependencies for this plugin.
  artifactory-snyk-security-core.jar
  ```

* You must sign up for an Enterprise account with Snyk before you begin.
* Snyk Artifactory plugin is compatible with Artifactory version 7.4.3 and higher.

**How to install the Artifactory Snyk plugin**

1. Log in to your Snyk account.
2. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **General** to locate, copy and save the following on the side:

   * service account token or Organization API token:
   * the Organization ID for \(any\) one of your organizations

   ![APIToken\_OrgID.png](https://support.snyk.io/hc/article_attachments/360007064657/uuid-d2c9a8bf-7873-249d-1719-2aca68c0970e-en.png)

3. Go to [our repo in GitHub](https://github.com/snyk/artifactory-snyk-security-plugin) and navigate to the Releases.
4. From the most current release, open the Assets section to download the artifactory-snyk-security-plugin-.zip archive.
5. Extract the folders and files and move the contents of the plugins folder to /artifactory/etc/plugins
6. Right-click the snykSecurityPlugin.properties file to open and edit it with any text editor.
7. The file contents appear as the code snippet below:

   **Note**

   When a backup file is created for the .properties file, Artifactory cannot recognize the difference between the original and the new file. Therefore, disable any backup features configured for the editor you choose before editing the file.

8. The following properties can be configured in this file:
   * **snyk.api.url**—on-prem customers should update the URL of their Snyk API endpoint based on their Snyk deployment; other users need not configure a URL.
   * **snyk.api.token**—this property is mandatory and must be configured in order for Snyk to authenticate your account, before scanning your artifacts. This is the token you copied in step 1.
   * **snyk.api.organization**—this property is mandatory and must be configured in order for Snyk to authenticate your account, along with your API token. Because this plugin does not import any data to your Snyk account, you can use the ID from any of your organizations. This is the organization ID you copied in step 1.
   * **snyk.artifactory.scanner.vulnerability.threshold**—default is \*low\*. Valid values include **none, low, medium, high, and critical**. Manually update the configuration based on your needs.
   * **snyk.artifactory.scanner.license.threshold**—default is \*low\*. Valid values include **none, low, medium, high, and critical**. Manually update the configuration based on your needs.
   * **snyk.http.proxyHost** and **snyk.http.proxyPort**—fill in your proxy connection string. HTTP proxies are supported.
9. Paste the token and the organization ID in place of the sample values for each of the parameters.
10. Copy and paste the plugin into **${ARTIFACTORY\_HOME}/etc/plugins/**
11. Restart your Artifactory server.

    **Note**

    Refresh now or Reload is not sufficient. Artifactory must be restarted.

12. Log in to your Artifactory instance and navigate to the System Logs to double-check Snyk has been installed successfully.

    ![image2.png](https://support.snyk.io/hc/article_attachments/360007146298/uuid-018348b6-2153-20e7-27d8-71aca312f6eb-en.png)

#### Artifactory plugin: use Snyk for your artifacts

Snyk runs in the background and whenever a download is requested from the UI or from the CLI, Snyk automatically scans the artifact to evaluate vulnerabilities and license issues.

When the scan ends, results are displayed in the Artifactory UI, in the artifact details.

To view details about download status, open the System Logs:

When the scan fails, based on the configurations that were set during installation, the download request is blocked with a HTTP status code 403 Forbidden. By reviewing the results, you can evaluate the issues found in your artifact and determine a course of action, before ever using that artifact.

When your setup blocks downloads with issues, you can override this configuration at the artifact level—enabling downloads even when issues are identified, per artifact.

#### Artifactory: understanding the Snyk test results

From the UI, the Snyk properties are displayed similar to the following:

Work with Snyk properties from Artifactory as follows:

| **Property** | **Description** |
| :--- | :--- |
| **snyk.issue.url** | This is the URL to our database and explanation of the vulnerability, including specific details about vulnerable versions, available upgrades and Snyk patches as well. |
| **snyk.issue.vulnerabilities** | Regardless of the thresholds configured, this row displays vulnerability summary scan results. |
| **snyk.issue.vulnerabilities.forceDownload** | When "true", allows downloads for this artifact even when there are vulnerabilities. |
| **snyk.issue.vulnerabilities.forceDownload.info** | Use this field to provide additional information for why the forceDownload is enabled. |
| **snyk.issue.licenses** | Regardless of the thresholds configured, this row displays license summary scan results. |
| **snyk.issue.licenses.forceDownload** | When "true", allows downloads for this artifact even when there are license issues. |
| **snyk.issue.licenses.forceDownload.info** | Use this field to provide additional information for why the forceDownload is enabled. |

