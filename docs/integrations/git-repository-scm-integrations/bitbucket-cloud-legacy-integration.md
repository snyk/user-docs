# Bitbucket Cloud legacy integration

{% hint style="info" %}
Support for the Bitbucket Cloud (Legacy) Personal Access Token (PAT) integration will continue for Snyk users who configured it prior to mid October 2022.

If you're new to Snyk, the Legacy integration tile may not be available on the Integrations page.

Snyk recommends installing or [migrating](migrate-a-bitbucket-cloud-legacy-integration.md) to the [Bitbucket Cloud Application](bitbucket-cloud-app-integration.md) for smoother integration and to ensure long-term support.
{% endhint %}

Snyk's Bitbucket Cloud PAT integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open source components
* Provide automated fixes and upgrades

> **Feature availability**\
> This feature is available for all plans. See [pricing plans](https://snyk.io/plans/) for more details.

## Setting up a Bitbucket Cloud Integration

> The newly created user must have **Admin** permissions to all the repositories you need to monitor with Snyk.
>
> Admin permissions are required; however, Snyk's access is ultimately limited by the [permissions assigned to the App Password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).

1. To give Snyk access to your Bitbucket account, set up a dedicated service account in Bitbucket, with admin permissions. See the [Bitbucket documentation ](https://support.atlassian.com/bitbucket-cloud/docs/grant-access-to-a-workspace/)to learn more about adding users to a workspace.
2. In Snyk, go to the **Integrations** page, open the **Bitbucket Cloud** card and configure the **Account credentials**.
3. In the **Account credentials >** **Creating an app password** section in Snyk, use the link <img src="../../.gitbook/assets/image (114) (1) (1) (1) (1) (1) (2) (1) (1) (2).png" alt="" data-size="line"> (**Create an App password**) to jump to your Bitbucket Cloud account.
4.  Follow the Bitbucket procedure to set up an account with the following permissions:

    * **Account: Email & Read**
    * **Workspace membership: Read**
    * **Projects: Read**
    * **Repositories: Read & Write**
    * **Pull requests: Read & Write**
    * **Webhooks: Read & Write**

    See the [Bitbucket documentation](https://confluence.atlassian.com/bitbucket/app-passwords-828781300.html) for more procedure details.
5. Enter the username and the [App Password for the Bitbucket account](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/) you created, and **Save** your changes.\
   You can find your username under the Bitbucket **Personal settings.**\
   Snyk connects to your Bitbucket Cloud account. When the connection succeeds, the following confirmation appears:\
   \
   <img src="../../.gitbook/assets/settings (1).png" alt="Confirmation that Snyk has connected to your Bitbucket account" data-size="original">

## Adding Bitbucket repositories to Snyk

After you connect Snyk to your Bitbucket Cloud account, you can select repositories for Snyk to monitor.

1. In Snyk, go to **Integrations** > **Bitbucket Cloud** card, and click **Add your Bitbucket Cloud repositories to Snyk** to start importing repositories to Snyk.
2. Choose the repositories you want to import to Snyk and click **Add selected repositories**.

After you add them, Snyk scans the selected repositories for dependency files in the entire directory tree, (that is, `package.json`, `pom.xml`, and so on) and imports them to Snyk as projects.

The imported projects appear in your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/444 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (16).png" alt="The Imported projects in your Projects page"><figcaption><p>The Imported projects in your <strong>Projects</strong> page</p></figcaption></figure>

## Bitbucket integration features

Once the integration is in place, you'll be able to use capabilities such as:

* [Project level security reports](bitbucket-cloud-legacy-integration.md#project-level-security-reports)
* [Project monitoring and automatic fix pull requests](bitbucket-cloud-legacy-integration.md#project-monitoring-and-automatic-fix-pull-requests)
* [Pull request testing](bitbucket-cloud-legacy-integration.md#pull-request-tests)

### Project-level security reports

Snyk produces advanced [security reports](https://docs.snyk.io/features/reports/reports-overview) that let you explore the vulnerabilities found in your repositories, and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

The example below presents a project-level security report.

<figure><img src="../../.gitbook/assets/bbc_project-sec-rpt_21sept2022 (1).png" alt="An example of a project-level security report"><figcaption><p>An example of a project-level security report</p></figcaption></figure>

### Project monitoring and automatic fix Pull Requests

Snyk scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and by opening [automated pull requests](https://docs.snyk.io/products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities) with fixes for your repositories.

The example below presents a fix Pull Request opened by Snyk.

<figure><img src="../../.gitbook/assets/666.png" alt="Example of an automatic fix Pull Request opened by Snyk"><figcaption><p>Example of an automatic fix Pull Request opened by Snyk</p></figcaption></figure>

To review and adjust the automatic fix pull request settings:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud**, and click **Edit Settings**.
2. Scroll to the **Automatic fix PRs** section and configure the relevant options.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-03 at 14.49.59.png" alt="Configure Automatic fix PRs"><figcaption><p>Configure Automatic fix PRs</p></figcaption></figure>

{% hint style="info" %}
Unlike manual pull requests opened from the Bitbucket interface, Snyk pull requests are _not_ automatically assigned to the default reviewer set in your Bitbucket Cloud account.

Click for more info on [Snyk automated pull requests](https://docs.snyk.io/products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities).
{% endhint %}

### Pull request tests

Snyk tests any newly created pull request in your repositories for security vulnerabilities and sends a build check to Bitbucket Cloud. You can to see whether the pull request introduces new security issues, directly from Bitbucket Cloud.

The example below presents a Snyk pull request build check on the Bitbucket Cloud **Pull Request** page.

<figure><img src="../../.gitbook/assets/888.png" alt="Example of a Snyk pull request build check on the Bitbucket Cloud Pull Request page"><figcaption><p>Example of a Snyk pull request build check on the Bitbucket Cloud <strong>Pull Request</strong> page</p></figcaption></figure>

To review and adjust the pull request tests settings,

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud**, and click **Edit Settings**.
2. Scroll to **Default Snyk test for pull requests > Open Source Security & Licenses**, and configure the relevant options.

<figure><img src="../../.gitbook/assets/Screenshot 2022-03-16 at 10.07.50.png" alt="Configuring the options for pull request Open Source Security &#x26; Licenses"><figcaption><p>Configuring the options for pull request Open Source Security &#x26; Licenses</p></figcaption></figure>

## Required permissions scope for the Bitbucket Cloud integration

All the operations, whether triggered manually or automatically, are performed for a Bitbucket Cloud [service account](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/service-accounts) that has its token (App Password) configured in the **Integration settings**.

The table below presents the required access scopes for the configured token:

| Action                                                  | Purpose                                                                                                                                                                                                                                               | Required permissions in Bitbucket                                                                                                     |
| ------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Daily / weekly tests                                    | Used to read manifest files in private repos                                                                                                                                                                                                          | Repositories: **Read**                                                                                                                |
| Manual fix pull requests (triggered by the user)        | Used to create fix PRs in the monitored repos                                                                                                                                                                                                         | <p>Repositories: <strong>Read</strong>, <strong>Write</strong></p><p>Pull requests: <strong>Read</strong>, <strong>Write</strong></p> |
| Automatic fix and upgrade pull requests                 | Used to create fix / upgrade PRs in the monitored repos                                                                                                                                                                                               | Repositories: **Read**, **Write** Pull requests: **Read**, **Write**                                                                  |
| Snyk tests on pull requests                             | Used to send pull request status checks whenever a new PR is created / an existing PR is updated                                                                                                                                                      | Repositories: **Read**, **Write** Pull requests: **Read**, **Write**                                                                  |
| Importing new projects to Snyk                          | Used to present a list of all the available repos in the Bitbucket in the "Add Projects" screen (import popup)                                                                                                                                        | <p>Account: <strong>Read</strong></p><p>Workspace membership: <strong>Read</strong></p><p>Projects: <strong>Read</strong></p>         |
| Snyk tests on pull requests - **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests (when PRs are created, updated triggered, merged, and so on)</li><li>Send push events to trigger PR checks</li></ul> | Webhooks: **Read**, **Write**                                                                                                         |

## Required permissions scope for repositories <a href="#h_01eefvj14p8b3depeffvyvdwzj" id="h_01eefvj14p8b3depeffvyvdwzj"></a>

For Snyk to perform the required operations on monitored repositories (such as reading manifest files on a frequent basis and opening fix or upgrade PRs), the integrated Bitbucket Cloud service account needs **Admin** permissions on the imported repositories:

| Action                                                  | Purpose                                                                                                                                                                                                                                               | Required permissions on the repository |
| ------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------- |
| Daily / weekly tests                                    | Used to read manifest files in private repositories.                                                                                                                                                                                                  | **Write** or above                     |
| Snyk tests on pull requests                             | Used to send pull request status checks when a new PR is created, or an existing PR is updated.                                                                                                                                                       | **Write** or above                     |
| Opening fix and upgrade pull requests                   | Used to create fix PRs in monitored repositories.                                                                                                                                                                                                     | **Write** or above                     |
| Snyk tests on pull requests - **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests (when PRs are created, updated triggered, merged, and so on)</li><li>Send push events to trigger PR checks</li></ul> | **Admin**                              |

## Disabling the Bitbucket Cloud integration

To disable this integration, in <img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings)> **Integrations:**

1. In your list of integrations, select the Bitbucket integration you want to deactivate and click **Edit settings** to open a page with the current status of your integration.\
   \
   The page includes sections that are specific to each integration, where you can manage your credentials, API key, Service Principal, or connection details.
2. Scroll to the relevant section and click **Disconnect.**

![Disconnect button at the bottom left of the Disconnect from Bitbucket Cloud section](../../.gitbook/assets/mceclip2-4-.png)

{% hint style="warning" %}
**WARNING**\
When you disconnect Snyk from your repository projects, your credentials are removed from Snyk and any integration-specific projects that Snyk is monitoring are deactivated in Snyk.\
If you choose to re-enable this integration later, you'll need to re-enter your credentials and activate your projects.
{% endhint %}
