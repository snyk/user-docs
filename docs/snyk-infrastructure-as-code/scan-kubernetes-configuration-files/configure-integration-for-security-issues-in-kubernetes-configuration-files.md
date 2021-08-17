# Configure integration for security issues in Kubernetes configuration files

Snyk tests and monitors Kubernetes configurations stored in your source code repositories and provides information, tips, and tricks to better secure a Kubernetes environment--catching misconfigurations before they are pushed to production as well as providing fixes for vulnerabilities.

## Supported Git repositories and file formats

Snyk currently scans your Kubernetes configuration files in JSON and YAML format when imported from your integrated Git repository.

## Configure Snyk to scan Kubernetes configuration files

**Configure Snyk**

1. 1. Navigate to the relevant group and organization that you want to manage

   ![AddProjectMenu.gif](https://support.snyk.io/hc/article_attachments/360006777178/uuid-da316a4a-c823-cf03-f37f-5305446dc970-en.gif)

   **Note**

   Integrations are managed per organization.

2. Enable Snyk to detect Kubernetes configuration files by enabling the flag in the settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Infrastructure as code** page:

   ![Screenshot\_2020-08-18\_at\_17.29.49.png](https://support.snyk.io/hc/article_attachments/4402311177489/Screenshot_2020-08-18_at_17.29.49.png)

3. If needed, review and adjust settings in the **Infrastructure as code** settings:

   ![Configure-Policies.png](https://support.snyk.io/hc/article_attachments/360006701757/uuid-34af73f5-ffde-39bb-ffa4-364884089b2e-en.png)

The number of tests per product are based on account plans. For information on plans and test limits, [view our plans](https://snyk.io/plans/).

