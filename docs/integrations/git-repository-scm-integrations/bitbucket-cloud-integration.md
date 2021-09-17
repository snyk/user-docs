# Bitbucket Cloud integration

Snyk's Bitbucket Cloud integration allows you to continuously perform security scanning across all the integrated repositories, detect vulnerabilities in your open source components, and use automated remediation.

> **Feature availability**  
> This feature is available for all plans. See [pricing plans](https://snyk.io/plans/) for more details.

## Setting up a Bitbucket Cloud Integration

> Make sure the newly created user has **Admin** permissions to all the repositories you need to monitor with Snyk.

1. To give Snyk access to your Bitbucket account, set up up a dedicated service account in Bitbucket, with admin permissions. Visit [Bitbucket documentation ](https://support.atlassian.com/bitbucket-cloud/docs/grant-access-to-a-workspace/)to learn more about creating users.
2. In Snyk, go to the **Integrations** page and click on **Bitbucket Cloud** card.
3. Access your Bitbucket Cloud account and create a unique App Password for Snyk with the following permissions: **Account: Email & Read** **Workspace membership: Read Projects: Read Repositories: Read & Write Pull requests: Read & Write Webhooks: Read & Write** Follow [Bitbucket documentation](https://confluence.atlassian.com/bitbucket/app-passwords-828781300.html/) to learn how to do it. 
4. Enter the username and the App Password for the service account you created:
5. Click **Save**.  
   Snyk connects to your Bitbucket Cloud account. When the connection succeeds, the following indications appear:

   ![](../../.gitbook/assets/settings%20%281%29.png/)  
   You can now select the repositories for Snyk to monitor.

6. Click **Add your Bitbucket Cloud repositories to Snyk** to start importing repositories to Snyk.
7. Select the repositories to import to Snyk when prompted, then click **Add selected repositories**.
8. Snyk scans the selected repositories for dependency files \(such as package.json and pom.xml\) in the entire directory tree, and import them to Snyk as projects:
9. The imported projects appear in your **Projects** page and are continuously checked for vulnerabilities.

![](../../.gitbook/assets/444%20%282%29%20%284%29%20%284%29%20%284%29%20%285%29%20%284%29%20%282%29.png/)

## Bitbucket Cloud Integration Features

After the integration is done, you can use the following capabilities:

## Project level security reports

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories, and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

This is an example of a project level security report:

![](../../.gitbook/assets/mceclip0-22-%20%282%29%20%285%29%20%286%29%20%281%29%20%287%29.png/)

## Projects monitoring and automatic fix pull requests

Snyk frequently scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, it notifies you by email and by opening an automated pull requests with fixes to repositories.

Here is an example of a fix pull request opened by Snyk:

![](../../.gitbook/assets/666.png/)

To review and adjust the automatic fix pull request settings:

Click on settings![cog\_icon.png](../../.gitbook/assets/cog_icon.png/) &gt; **Integrations**. 2. Select **Edit Settings** for Bitbucket Cloud. 3. Navigate to **Automatic fix pull requests**:

![](../../.gitbook/assets/mceclip4%20%281%29%20%282%29%20%286%29%20%287%29%20%283%29%20%2817%29.png/)

## Pull request tests

Snyk tests any newly created pull request in your repositories for security vulnerabilities, and sends a build check to Bitbucket Cloud. You can to see whether the pull request introduces new security issues, directly from Bitbucket Cloud.

This is how Snyk pull request build check appears in the **Pull Request** page in Bitbucket Cloud:

![](../../.gitbook/assets/888.png/)

To review and adjust the pull request tests settings,

1. Click on settings![cog\_icon.png](../../.gitbook/assets/cog_icon.png/) &gt; **Integrations**.
2. Select **Edit Settings** for Bitbucket Cloud. 
3. Navigate to **Default Snyk test for pull requests**

![](../../.gitbook/assets/999.png/)

## Required permissions scope for the Bitbucket Cloud integration

All the operations, triggered manually or automatically, are performed for a Bitbucket Cloud service account that has its token \(App Password\) is configured in the integrations settings. This shows the required access scopes for the configured token:

| **Action** | **Why?** | **Required permissions in Bitbucket** |
| :--- | :--- | :--- |
| Daily / weekly tests | For reading manifest files in private repos | _Repositories read_ |
| Manual fix pull requests \(triggered by the user\) | For creating fix PRs in the monitored repos | _Repositories \(read & write\)_ _pull requests \(read & write\)_ |
| Automatic fix and upgrade pull requests | For creating fix / upgrade PRs in the monitored repos | _Repositories \(read & write\)_ _pull requests \(read & write\)_ |
| Snyk tests on pull requests | For sending pull request status checks whenever a new PR is created / an existing PR is updated | _Repositories \(read & write\)_ _pull requests \(read & write\)_ |
| Importing new projects to Snyk | For presenting a list of all the available repos in the Bitbucket in the "Add Projects" screen \(import popup\) | _Account \(read\)_ _Workspace membership \(read\)_ _Projects \(read\)_ |
| Snyk tests on pull requests - initial configuration | For adding Snyk's webhooks to the imported repos, so Snyk will be informed whenever pull requests are created or updated and be able to trigger scans. | _webhooks \(read & write\)_ |

## Required permissions scope for repositories <a id="h_01EEFVJ14P8B3DEPEFFVYVDWZJ"></a>

For Snyk to perform the required operations on monitored repositories \(such as reading manifest files on a frequent basis and opening fix or upgrade PRs\), the integrated Bitbucket Cloud service account needs **Admin** permissions on the imported repositories:

| **Action** | **Why?** | **Required permissions on the repository** |
| :--- | :--- | :--- |
| Daily / weekly tests | To read manifest files in private repositories. | **Write** or above |
| Snyk tests on pull requests | To send pull request status checks when a new PR is created, or an existing PR is updated. |  |
| Opening fix and upgrade pull requests | To create fix PRs in monitored repositories. |  |
| Snyk tests on pull requests - initial configuration | To add Snyk's webhooks to the imported repos, so Snyk is informed when pull requests are created or updated, and can trigger scans. | **Admin** |

## 1st Party Integration \(Connect App\/)

The Connect App integration is a layer on top of the regular App Password integration, which allows your developers to consume Snyk from the Bitbucket interface.

> The Connect App currently supports [Snyk Open Source](../../snyk-open-source/) and [Snyk Container](../../snyk-container/) products only.

By adding the Connect App to your Bitbucket workspaces, your workspaces members can import repos and see the security data in a dedicated Snyk tab in Bitbucket Cloud:

![](../../.gitbook/assets/mceclip1-13-.png/)

## Installing the Connect App

To install the app, navigate to the **Security** tab in one of your workspace's repos in Bitbucket Cloud, then click **Try now**:

![](../../.gitbook/assets/mceclip2-3-.png/)

## Uninstalling the Connect App

To remove the Connect App from your workspace in Bitbucket Cloud:

1. Navigate to the **workspace settings** page &gt; **Installed apps**.
2. Find **Snyk Security for Bitbucket Cloud** in the installed applications list.
3. Click **remove**. 

## Associating the Connect App to a different Snyk account / organization

The Connect App is associated to a specific Snyk account and organization, as defined during the app onboarding process.

To change these settings later, navigate to the workspace settings and select **Security for Bitbucket Cloud Integration Settings**:

![](../../.gitbook/assets/mceclip0-23-.png/)

## Disabling the Bitbucket Cloud integration

To disable this integration:

1. Click on settings![cog\_icon.png](../../.gitbook/assets/cog_icon.png/) &gt; **Integrations** in Snyk.
2. Find the specific integration to deactivate in your list of integrations, and click Edit settings. 
3. A page appears showing the current status of your integration and a place to update your credentials, specific to each integration \(credentials, API key, Service Principal, or connection details\):
4. Click **Disconnect**.

![](../../.gitbook/assets/mceclip2-4-.png/)

> Your credentials are removed from Snyk and any integration-specific projects Snyk is monitoring are deactivated on Snyk.  
> If you then choose to re-enable this integration at any time, you will need to re-enter your credentials and activate your projects.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

