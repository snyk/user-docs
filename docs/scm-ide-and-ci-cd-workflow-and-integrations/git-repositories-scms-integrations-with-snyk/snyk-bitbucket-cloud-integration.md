# Snyk Bitbucket Cloud integration

{% hint style="info" %}
**Feature availability**\
This feature is available for all plans. See [pricing plans](https://snyk.io/plans/) for more details.

Snyk recommends installing or [migrating](migrate-a-bitbucket-cloud-personal-access-token.md) to the [Bitbucket Cloud Application](snyk-bitbucket-cloud-app-integration.md) for smoother integration and to ensure long-term support.
{% endhint %}

The Snyk Bitbucket Cloud (PAT) integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes and upgrades

## How to set up the Snyk Bitbucket Cloud Integration

{% hint style="info" %}
Admin permissions are required; however, Snyk's access is ultimately limited by the [permissions assigned to the App Password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).
{% endhint %}

1. To give Snyk access to your Bitbucket account, set up a dedicated service account in Bitbucket with admin permissions. See the [Bitbucket documentation ](https://support.atlassian.com/bitbucket-cloud/docs/grant-access-to-a-workspace/)to learn more about adding users to a workspace.\
   The newly created user must have **Admin** permissions to all the repositories you need to monitor with Snyk.
2. In Snyk, go to the **Integrations** page, open the **Bitbucket Cloud** card, and configure the **Account credentials**.
3. In the **Account credentials >** **Creating an app password** section in Snyk, use the link <img src="../../.gitbook/assets/image (365) (1) (1) (1) (1) (1) (1) (2).png" alt="Create an app password" data-size="line"> (**Create an App password**) to jump to your Bitbucket Cloud account.
4.  Follow the Bitbucket procedure to set up an account with the following permissions:

    * **Account: Email & Read**
    * **Workspace membership: Read**
    * **Projects: Read**
    * **Repositories: Read & Write**
    * **Pull requests: Read & Write**
    * **Webhooks: Read & Write**

    See the [Bitbucket documentation](https://confluence.atlassian.com/bitbucket/app-passwords-828781300.html) for more details about the procedure.
5. Enter the username and the [App Password for the Bitbucket account](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/) you created and **Save** your changes.\
   You can find your username under the Bitbucket **Personal settings.**\
   Snyk connects to your Bitbucket Cloud account. When the connection succeeds, the following confirmation appears:\
   \
   <img src="../../.gitbook/assets/settings (1).png" alt="Bitbucket Cloud settings successfully updated." data-size="original">

## How to add Bitbucket repositories to Snyk

After you connect Snyk to your Bitbucket Cloud account, you can select repositories for Snyk to monitor.

1. In Snyk, go to **Integrations** > **Bitbucket Cloud** card, and click **Add your Bitbucket Cloud repositories to Snyk** to start importing repositories to Snyk.
2. Choose the repositories you want to import to Snyk and click **Add selected repositories**.

After you add the selected repositories, Snyk scans them for dependency files in the entire directory tree, that is, `package.json`, `pom.xml`, and so on, and imports them to Snyk as Projects.

The imported projects appear on your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/444 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (12) (2).png" alt="The Imported projects on your Projects page"><figcaption><p>The Imported projects on your <strong>Projects</strong> page</p></figcaption></figure>

## Bitbucket integration features

After the integration is in place, you will be able to use capabilities such as:

* [Project-level security reports](snyk-bitbucket-cloud-integration.md#project-level-security-reports)
* [Project monitoring and automatic fix pull requests](snyk-bitbucket-cloud-integration.md#project-monitoring-and-automatic-fix-pull-requests)
* [Pull request testing](snyk-bitbucket-cloud-integration.md#pull-request-tests)

### Project-level security reports

Snyk produces advanced [security reports](https://docs.snyk.io/features/reports/reports-overview) that let you explore the vulnerabilities found in your repositories and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

The example that follows shows a Project-level security report.

<figure><img src="../../.gitbook/assets/bbc_project-sec-rpt_21sept2022.png" alt="An example of a Project-level security report"><figcaption><p>An example of a Project-level security report</p></figcaption></figure>

### Project monitoring and automatic fix Pull Requests

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and by opening [automated pull requests](../../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/create-automatic-prs-for-new-fixes.md) with fixes for your repositories.

The example that follows shows a fix Pull Request opened by Snyk.

<figure><img src="../../.gitbook/assets/666.png" alt="Example of an automatic fix Pull Request opened by Snyk"><figcaption><p>Example of an automatic fix Pull Request opened by Snyk</p></figcaption></figure>

To review and adjust the automatic fix pull request settings:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud**, and click **Edit Settings**.
2. Scroll to the **Automatic fix PRs** section and configure the relevant options.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-03 at 14.49.59.png" alt="Configure Automatic fix PRs"><figcaption><p>Configure Automatic fix PRs</p></figcaption></figure>

{% hint style="info" %}
Unlike manual pull requests opened from the Bitbucket interface, Snyk pull requests are _not_ automatically assigned to the default reviewer set in your Bitbucket Cloud account.

For more information, see [Snyk automated pull requests](../../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/create-automatic-prs-for-new-fixes.md).
{% endhint %}

### Pull request tests

Snyk tests any newly-created pull request in your repositories for security vulnerabilities and sends a build check to Bitbucket Cloud. You can see directly from Bitbucket Cloud whether or not the pull request introduces new security issues.

The example that follows shows a Snyk pull request build check on the Bitbucket Cloud **Pull Request** page.

<figure><img src="../../.gitbook/assets/888.png" alt="Example of a Snyk pull request build check on the Bitbucket Cloud Pull Request page"><figcaption><p>Example of a Snyk pull request build check on the Bitbucket Cloud <strong>Pull Request</strong> page</p></figcaption></figure>

To review and adjust the pull request tests settings:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud**, and click **Edit Settings**.
2. Scroll to **Default Snyk test for pull requests > Open Source Security & Licenses**, and configure the relevant options.

<figure><img src="../../.gitbook/assets/Screenshot 2022-03-16 at 10.07.50.png" alt="Configuring the options for pull request Open Source Security &#x26; Licenses"><figcaption><p>Configuring the options for pull request Open Source Security &#x26; Licenses</p></figcaption></figure>

## Required permission scope for the Bitbucket Cloud integration

All the operations, whether triggered manually or automatically, are performed for a Bitbucket Cloud [service account](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/service-accounts) that has its token (App Password) configured in the **Integration settings**.

For Snyk to perform the required operations on monitored repositories, such as reading manifest files on a frequent basis and opening fix or upgrade PRs, the integrated Bitbucket Cloud service account needs **Admin** permissions on the imported repositories.

For detailed information on the permission scopes required, see [Bitbucket permission requirements](./#bitbucket-permission-requirements).

## How to disconnect Snyk from Bitbucket Cloud

{% hint style="warning" %}
When you disconnect Snyk from your repository Projects, your credentials are removed from Snyk, and any integration-specific Projects that Snyk is monitoring are deactivated in Snyk.\
If you choose to re-enable this integration, you must re-enter your credentials and activate your Projects.
{% endhint %}

To disconnect this integration, in **Organization settings** > **Integrations:**

1. In your list of integrations, select the Bitbucket integration you want to deactivate and click **Edit settings** to open a page with the current status of your integration.\
   \
   The page includes sections that are specific to each integration, where you can manage your credentials, API key, Service Principal, or connection details.
2. Scroll to the relevant section and click **Disconnect.**

<figure><img src="../../.gitbook/assets/mceclip2-4-.png" alt="Disconnect button at the bottom left of the Disconnect from Bitbucket Cloud section"><figcaption><p>Disconnect button at the bottom left of the Disconnect from Bitbucket Cloud section</p></figcaption></figure>

## Migrate to the Snyk Bitbucket Cloud App

This section describes how to migrate your existing [Bitbucket Cloud Personal Access Token (PAT) integration](snyk-bitbucket-cloud-integration.md), displayed in Snyk as Bitbucket Cloud, to the [**Bitbucket Cloud App**](snyk-bitbucket-cloud-app-integration.md) integration.

To migrate to the new app integration, you must remove all the previously imported Projects from Snyk, delete the PAT integration and its Projects, set up the new app integration, and reimport your Projects to Snyk from the new integration.

{% hint style="info" %}
Before going through the migration process, you should note that the following Project-level information will not persist:

* Historic Project-related data, including trend numbers for fixing vulnerabilities
* Project-related metadata: ignores and tags
{% endhint %}

### Migration process

The migration process includes the following steps:

1. [Deleting the existing Projects](snyk-bitbucket-cloud-integration.md#delete-existing-projects) that are connected to the Bitbucket Cloud PAT integration in Snyk.
2. [Disconnecting the PAT integration](snyk-bitbucket-cloud-integration.md#disconnect-the-pat-integration) in Snyk.
3. [Removing the first-party extension](snyk-bitbucket-cloud-integration.md#remove-the-snyk-tab-for-the-pat-integration-in-bitbucket-cloud-optional) for the PAT integration in Bitbucket (optional).
4. [Connecting the Bitbucket Cloud App](snyk-bitbucket-cloud-integration.md#set-up-the-bitbucket-cloud-app-integration) and importing Projects.

#### Delete existing Projects

Delete all the existing Projects in Snyk that were previously imported from the Legacy integration. To use the bulk delete action on the Projects page, change the grouping filter to **Group by none**. You can now select multiple Projects in the list individually or by selecting the checkbox at the top to **Select all visible projects**. To delete a Project, select the trash icon, **Delete selected projects**.

<figure><img src="../../.gitbook/assets/2023-11-20_14-29-35.png" alt="Change the Projects filter to Group by none"><figcaption><p>Change the Projects filter to <strong>Group by none</strong></p></figcaption></figure>

<figure><img src="../../.gitbook/assets/2023-11-20_14-41-16.png" alt="Bulk delete the selected Projects" width="375"><figcaption><p>Bulk delete the selected Projects</p></figcaption></figure>

#### Disconnect the PAT integration

To disconnect the Bitbucket Cloud PAT integration, navigate to the settings page of Bitbucket Cloud integration, scroll to the relevant section, and click **Disconnect.**

<figure><img src="../../.gitbook/assets/image (524).png" alt="Disconnect the Bitbucket Cloud PAT (Legacy) integration"><figcaption><p>Disconnect the Bitbucket Cloud PAT (Legacy) integration</p></figcaption></figure>

#### Remove the Snyk tab for the PAT integration in Bitbucket Cloud (optional)

The Bitbucket Cloud integration has an optional first-party interface app.

This app can be installed on your Bitbucket Cloud workspace to enrich the PAT integration with a first-party interface as the "_**Snyk**_" tab)

If you have used this app, before setting up the Snyk Bitbucket Cloud App in the next step, remove the previous interface app in Bitbucket Cloud.\
This functionality is supported out-of-the-box in the Snyk App integration.\
\
Go to your **Workspace settings** page in **Bitbucket.org > Manage installed apps**, expand the **Snyk Security for Bitbucket Cloud** app, and click **Remove.**

<figure><img src="../../.gitbook/assets/remove_snyk-security-bbc_11oct2022.png" alt="Remove the first-party Snyk Legacy interface app in Bitbucket"><figcaption><p>Remove the first-party Snyk Legacy interface app in Bitbucket</p></figcaption></figure>

#### Set up the Bitbucket Cloud App integration

See the [Bitbucket Cloud App integration](snyk-bitbucket-cloud-app-integration.md) topic for instructions.

### Migration demo

In less than five minutes, Marco Morales, a Partner Solutions Architect at Snyk, talks about the Snyk Bitbucket Cloud App and goes through the process of migrating an existing Bitbucket Cloud integration to the Snyk Bitbucket Cloud App.

_Go to timestamp 2:34 to jump right into the demo._

{% embed url="https://thoughtindustries-1.wistia.com/medias/32rgw3hkdk" %}
How to migrate to the new Snyk Bitbucket Cloud App integration
{% endembed %}
