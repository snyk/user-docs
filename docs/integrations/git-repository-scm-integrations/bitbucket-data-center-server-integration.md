# Bitbucket Data Center/Server integration

1. To give Snyk access to your Bitbucket DC/Server account, set up up a dedicated service account in Bitbucket DC/Server, with admin permissions. Visit [Bitbucket Server documentation ](https://confluence.atlassian.com/bitbucketserver/users-and-groups-776640439.html#Usersandgroups-Creatingauser)to learn more about creating users.

   **Important**

   Make sure the newly created user has **Admin** permissions to all the repositories you need to monitor with Snyk.

2. In Snyk, go to the **Integrations** page and click on **Bitbucket Server** card:![111.png](../../.gitbook/assets/111.png)
3. Enter your Bitbucket DC/Server URL, and the username and password for the service account you created: ![222.png](../../.gitbook/assets/222.png)
4. Click **Save**. Snyk connects to your Bitbucket DC/Server instance. When the connection succeeds, the following indications appear: ![333.png](../../.gitbook/assets/333.png) You can now select the repositories for Snyk to monitor.
5. Click **Add your Bitbucket Server repositories to Snyk** to start importing repositories to Snyk.
6. Select the repositories to import to Snyk when prompted, then click **Add selected repositories**.
7. Snyk scans the selected repositories for dependency files \(such as package.json and pom.xml\) in the entire directory tree, and import them to Snyk as projects: ![444.png](../../.gitbook/assets/444.png)
8. The imported projects appear in your **Projects** page and are continuously checked for vulnerabilities.

## Bitbucket DC/Server Integration Features

After the integration is done, you can use the following capabilities:

### **Project level security reports**

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories, and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

This is an example of a project level security report:

![](../../.gitbook/assets/555.png)

### **Projects monitoring and automatic fix pull requests**

Snyk frequently scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, it notifies you by email and by opening an automated pull requests with fixes to repositories.

Here is an example of a fix pull request opened by Snyk:![666.png](../../.gitbook/assets/666.png)

To review and adjust the automatic fix pull request settings:

### **Pull request tests**

Snyk tests any newly created pull request in your repositories for security vulnerabilities, and sends a build check to Bitbucket DC/Server. You can to see whether the pull request introduces new security issues, directly from Bitbucket DC/Server.

This is how Snyk pull request build check appears in the **Pull Request** page in Bitbucket DC/Server:![888.png](../../.gitbook/assets/888.png)

To review and adjust the pull request tests settings:

1. Click on settings ![cog\_icon.png](../../.gitbook/assets/cog_icon.png) &gt; **Integrations**.
2. Select **Edit Settings** for Bitbucket Server
3. Navigate to **Default Snyk test for pull requests**: ![999.png](../../.gitbook/assets/999.png)

