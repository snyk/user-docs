# Bitbucket Data Center/Server integration

Snyk's Bitbucket Data Center / Server integration allows you to continuously perform security scanning across all the integrated repositories, detect vulnerabilities in your open source components, and use automated fixing. This integration supports Bitbucket Data Center / Server versions 4.0 and above.

> **Feature availability**\
> This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.

_Need a little help from your friends?_ Check out the [Snyk and Bitbucket best practices cheat sheet](https://snyk.io/blog/snyk-bitbucket-best-practices-cheat-sheet/) in our blog!

## Setting up a Bitbucket DC/Server Integration

> **Important**\
> Make sure the newly created user has **Admin** permissions to all the repositories you need to monitor with Snyk.

1. To give Snyk access to your Bitbucket DC/Server account, set up up a dedicated service account in Bitbucket DC/Server, with admin permissions. Visit [Bitbucket Server documentation ](https://confluence.atlassian.com/bitbucketserver/users-and-groups-776640439.html#Usersandgroups-Creatingauser)to learn more about creating users.
2. In Snyk, go to the **Integrations** page and click on **Bitbucket Server** card.
3. Enter your Bitbucket DC/Server URL, and the username and password for the service account you created. Alternatively, you can create a [personal access token](https://confluence.atlassian.com/bitbucketserver075/personal-access-tokens-1018784848.html) and use it instead of a password.
4. **Save** your changes.\
   Snyk connects to your Bitbucket DC/Server instance. When the connection succeeds, the following indications appear:

![](../../../.gitbook/assets/bitbucket\_server-18july2022.png)

To select the repositories for Snyk to monitor:

1. Click **Add your Bitbucket Server repositories to Snyk** to start importing repositories to Snyk.
2. When prompted, select the repositories to import to Snyk and click **Add selected repositories**.

After you add them, Snyk scans the selected repositories for dependency files in the entire directory tree, (that is, `package.json`, `pom.xml`, and so on) and imports them to Snyk as projects.\
\
The imported projects appear in your **Projects** page and are continuously checked for vulnerabilities.

![](../../../.gitbook/assets/bitbucketserver\_add-repos\_18july2022.png)

## Bitbucket DC/Server Integration Features

Once the integration is in place, you'll be able to use capabilities such as:

* [Project level security reports](bitbucket-data-center-server-integration.md#project-level-security-reports)
* [Project monitoring and automatic fix pull requests](bitbucket-data-center-server-integration.md#projects-monitoring-and-automatic-fix-pull-requests)
* [Pull request testing](bitbucket-data-center-server-integration.md#pull-request-testing)

### **Project level security reports**

Snyk produces advanced [security reports](https://docs.snyk.io/features/reports/reports-overview) that let you explore the vulnerabilities found in your repositories, and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

The example below presents a project level security report.

![](../../../.gitbook/assets/project\_lvl\_security\_rpt-18july2022.png)

### **Project monitoring and automatic fix Pull Requests**

Snyk scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and by opening automated Snyk Pull Requests with fixes for your repositories.

The example below presents a fix Pull Request opened by Snyk.

![](../../../.gitbook/assets/666.png)

To review and adjust the automatic fix pull request settings:

1. In Snyk, go to <img src="../../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Server**, and click **Edit Settings**.
2. Scroll to the **Automatic fix pull requests** section and configure the relevant options.

![](../../../.gitbook/assets/bitbucket\_server-autofixprs\_18july2022.png)

{% hint style="info" %}
Unlike manual pull requests opened from the Bitbucket interface, Snyk pull requests are not automatically assigned to the default reviewer set in your Bitbucket Cloud or Bitbucket Server account. More info on [Snyk automated pull requests](https://docs.snyk.io/products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities).
{% endhint %}

### **Pull request tests**

Snyk tests any newly created pull request in your repositories for security vulnerabilities, and sends a build check to Bitbucket DC/Server. You can to see whether the pull request introduces new security issues, directly from Bitbucket DC/Server.

The example below presents a Snyk pull request build check on the Bitbucket DC/Server **Pull Request** page\*\*.\*\*

![](../../../.gitbook/assets/888.png)

To review and adjust the pull request tests settings:

1. In Snyk, go to <img src="../../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Server** , and click **Edit Settings**.
2. Scroll to **Default Snyk test for pull requests > Open Source Security & Licenses**, and configure the relevant options.

![](../../../.gitbook/assets/bbs\_default-pr-test\_18july2022.png)

## Required permissions scope for the Bitbucket DC/Server integration

Snyk performs all the operations in Bitbucket DC/Server on behalf of the integrated service account.

For Snyk to perform the required operations on monitored repositories (such as reading manifest files on a frequent basis and opening fix or upgrade PRs), the integrated Bitbucket DC/Server service account needs **Admin** permissions on the imported repositories:

| **Action**                                          | **Purpose**                                                                                                                                                                                      | **Required permissions on the repository** |
| --------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------ |
| Daily / weekly tests                                | Used to read manifest files in private repositories.                                                                                                                                             | **Write** or above                         |
| Snyk tests on pull requests                         | Used to send pull request status checks when a new PR is created, or an existing PR is updated.                                                                                                  | **Write** or above                         |
| Opening fix and upgrade pull requests               | Used to create fix PRs in monitored repositories.                                                                                                                                                | **Write** or above                         |
| Snyk tests on pull requests - initial configuration | Used to add [Snyk webhooks ](https://docs.snyk.io/integrations/snyk-webhooks)to the imported repos, to notify Snyk when pull requests are created or updated, and enable Snyk to trigger a scan. | **Admin**                                  |

## **Disabling the Bitbucket Data Center/Server integration**

To disable this integration, in <img src="../../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Integrations:**

1. In your list of integrations, select the Bitbucket integration you want deactivate and click **Edit settings** to open a page with the current status of your integration.\
   \
   The page includes sections that are specific to each integration, where you can manage your credentials, API key, Service Principal, or connection details.
2. Scroll to the relevant section and click **Disconnect.**

![](../../../.gitbook/assets/101010.png)

{% hint style="warning" %}
**WARNING**\
\*\*\*\*When you disconnect Snyk from your repository projects, your credentials are removed from Snyk and any integration-specific projects that Snyk is monitoring are deactivated in Snyk.\
If you choose to re-enable this integration later, you'll need to re-enter your credentials and activate your projects.
{% endhint %}
