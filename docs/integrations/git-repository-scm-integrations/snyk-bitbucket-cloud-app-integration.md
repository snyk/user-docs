---
description: >-
  The Bitbucket Cloud App is positioned to be the default Bitbucket Cloud
  integration
---

# Snyk Bitbucket Cloud App integration

{% hint style="info" %}
the **Feature availability**\
This feature is available for all plans. See [pricing plans](https://snyk.io/plans/) for more details.

Snyk offers two Bitbucket Cloud integrations: the Personal Access Token (PAT) legacy integration and the new App integration. The App integration is available only on multi-tenant US. If you are a user of multi-tenant EU/AU or single tenant, use the PAT integration.
{% endhint %}

The Snyk Bitbucket Cloud App integration lets you connect your Snyk Organization to a Bitbucket Cloud Workspace and get all Snyk's core SCM integration features:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your Open Source components
* Provide automated fixes and upgrades
* Provides developer teams with first-party visibility for security issues directly in the Bitbucket interface

{% hint style="info" %}
Snyk recommends using the Bitbucket Cloud App integration for smoother integration and to ensure long-term support.

If you are using the [Bitbucket Cloud Personal Access Token (Legacy) integration](snyk-bitbucket-cloud-legacy-integration.md), see [Migrate a Bitbucket Cloud Legacy integration](migrate-a-bitbucket-cloud-personal-access-token.md) for more information.
{% endhint %}

## Setting up a Bitbucket Cloud App

To give Snyk access to your Bitbucket account, you need to install the Snyk App on your Bitbucket Cloud workspace.

{% hint style="info" %}
To install the Snyk App on your Bitbucket Cloud workspace, you must have **Admin** permissions for the Workspace in Bitbucket.
{% endhint %}

1. In Snyk, go to **Integrations (Source control) >** **Bitbucket Cloud App** tile, and click **Connect** to install the Snyk Bitbucket Cloud App on your Bitbucket Cloud workspace.
2.  In the new Bitbucket tab, select the relevant workspace to connect to your Snyk Organization from the list and [**Grant access** to let Snyk](snyk-bitbucket-cloud-app-integration.md#required-snyk-bitbucket-cloud-app-permissions-scope):

    * Read your account information 
    * Read and modify your repositories and their pull requests
    * Read and modify your repositories' webhooks

    <figure><img src="../../.gitbook/assets/image (108) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (4).png" alt="Allow access for Snyk to Bitbucket Cloud" width="563"><figcaption><p>Allow access for Snyk to Bitbucket Cloud</p></figcaption></figure>
3.  Grant access to your Snyk Organization when you're prompted.

    <figure><img src="../../.gitbook/assets/bitbucket-cloud-permissions_10nov2022.png" alt="Allow Bitbucket Cloud access to your Snyk Organization" width="365"><figcaption><p>Allow Bitbucket Cloud access to your Snyk Organization<br></p></figcaption></figure>

    After you allow access to the Snyk Organization, the Snyk **Organization Settings** page opens and confirms that the Bitbucket Cloud App is connected.\


After Snyk is integrated with Bitbucket Cloud, you can see the new Snyk security tab on the repository page and [import and explore the issues and vulnerabilities for your repository Projects](../../getting-started/quickstart/import-a-project.md) directly in Bitbucket.

<figure><img src="../../.gitbook/assets/bbcloud-app_snyk-security_6oct2022.png" alt="Bitbucket security insights with Snyk Bitbucket Cloud App"><figcaption><p>Bitbucket security insights with Snyk Bitbucket Cloud App</p></figcaption></figure>

Watch this short video to see how to set up **Snyk security** in Bitbucket Cloud.

{% embed url="https://thoughtindustries-1.wistia.com/medias/peusq0bkie" %}
Set up Snyk security in Bitbucket Cloud
{% endembed %}

## Installing the Snyk App from Bitbucket Cloud

If you need to, you can also install the Snyk Bitbucket Cloud App integration while you are in Bitbucket Cloud.

In one of your Bitbucket Cloud workspaces, go to the **Security** tab in one of your repositories, click **Try now**, and follow the procedure.

<figure><img src="../../.gitbook/assets/install-app-bbc_6oct-2022.png" alt="Install the Snyk Bitbucket Cloud App from Bitbucket"><figcaption><p>Install the Snyk Bitbucket Cloud App from Bitbucket</p></figcaption></figure>

## Adding Bitbucket repositories to Snyk

After you connect Snyk to your Bitbucket Cloud account, you can select repositories for Snyk to monitor.

1. In Snyk, go to **Integrations** > **Bitbucket Cloud App** card and click to start importing repositories to Snyk.
2. Choose the repositories you want to import to Snyk and click **Add selected repositories**.

After you add them, Snyk scans the selected repositories for dependency files in the entire directory tree, that is, `package.json`, `pom.xml`, and so on, and imports them to Snyk as Projects.

The imported Projects appear on your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/bbc-snyk_redo_21sept2022 (1).png" alt="Snyk Projects page"><figcaption><p>Snyk Projects page</p></figcaption></figure>

## Bitbucket integration features

After the integration is in place, you can use capabilities such as:

* [Project-level security reports](snyk-bitbucket-cloud-app-integration.md#project-level-security-reports)
* [Pull request testing](snyk-bitbucket-cloud-app-integration.md#pull-request-tests)
* [First-party interface in Bitbucket Cloud](snyk-bitbucket-cloud-app-integration.md#first-party-interface-in-bitbucket-cloud)

### Project-level security reports

Snyk produces advanced [security reports](../../manage-issues/reporting/legacy-reports/legacy-reports-overview.md) that let you explore the vulnerabilities found in your repositories, and fix them immediately by opening a fix pull request directly to your repository, with the required upgrades or patches.

The example that follows shows a Project-level security report.

<figure><img src="../../.gitbook/assets/bbc_project-sec-rpt_21sept2022.png" alt="Project-level security report"><figcaption><p>Project-level security report</p></figcaption></figure>

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and by opening [automated pull requests](../../scan-application-code/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) with fixes for your repositories.

The example that follows shows a fix pull request opened by Snyk.

<figure><img src="../../.gitbook/assets/bbc-app_pr_6oct2022.png" alt="Fix pull request opened by Snyk"><figcaption><p>Fix pull request opened by Snyk</p></figcaption></figure>

To review and adjust the automatic fix pull request settings:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud App**, and click **Edit Settings**.
2.  Scroll to the **Automatic fix PRs** section and configure the relevant options.

    <div align="left">

    <figure><img src="../../.gitbook/assets/Screenshot 2023-05-02 at 11.19.09.png" alt="Automatic fix PR settings "><figcaption><p>Automatic fix PR settings </p></figcaption></figure>

    </div>

{% hint style="info" %}
Unlike manual pull requests opened from the Bitbucket interface, Snyk pull requests are _not_ automatically assigned to the default reviewer set in your Bitbucket Cloud account.

For more information, see [Automated pull request creation for new fixes](../../scan-application-code/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md).
{% endhint %}

### Pull request tests

Snyk tests any newly created pull request in your repositories for security vulnerabilities and sends a build check to Bitbucket Cloud. You can see directly from Bitbucket Cloud whether or not the pull request introduces new security issues.

The example that follows shows a Snyk pull request build check on the Bitbucket Cloud **Pull Request** page.

<figure><img src="../../.gitbook/assets/888.png" alt="BitBucket Cloud pull request page showing Snyk pull request"><figcaption><p>BitBucket Cloud pull request page showing Snyk pull request</p></figcaption></figure>

To review and adjust the pull request test settings, follow these steps:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> (Organization settings) > **Integrations > Source control > Bitbucket Cloud App**, and click **Edit Settings**.
2. Scroll to **Default Snyk test for pull requests > Open Source Security & Licenses**, and configure the relevant options. See [Configure PR Checks](../../scan-application-code/run-pr-checks/configure-pr-checks.md) for more details.

### First-party interface in Bitbucket Cloud

When you install the Snyk Bitbucket Cloud App integration in your Bitbucket workspace, the members of your workspace can import repositories and see the security issues related to their repositories in a dedicated Snyk security tab in Bitbucket Cloud.

<figure><img src="../../.gitbook/assets/bbcloud-app_snyk-security_6oct2022.png" alt="Snyk Security tab in Bitbucket Cloud"><figcaption><p>Snyk Security tab in Bitbucket Cloud</p></figcaption></figure>

{% hint style="warning" %}
The first-party interface currently supports only the [Snyk Open Source](../../scan-application-code/snyk-open-source/) and [Snyk Container](../../scan-applications/snyk-container/) products. Issues from other Snyk products do not show up on this page.
{% endhint %}

You can **associate a first-party interface with a different Snyk account or Organization**.

During the first-time Bitbucket Cloud App onboarding process, the first-party interface is associated with a specific Snyk Organization. This is the Snyk Organization to which Bitbucket users will import repositories and for which they will view Snyk issues.

To change the Snyk Organization after onboarding, go to the workspace settings > **Security for Bitbucket Cloud > Integration Settings** and click **Connect via a different Snyk user/organization**.

<figure><img src="../../.gitbook/assets/bbc-app-1st-party_change-org_6oct2022.png" alt="Create integration settings for a different Organization"><figcaption><p>Create integration settings for a different Organization</p></figcaption></figure>

The installation process begins again, and you can choose the relevant Snyk Organization.

## Required permissions scope for the Snyk Bitbucket Cloud App integration

The table that follows shows the required access scopes for the Bitbucket Cloud App.

| Action                                              | Purpose                                                                                                                                       | Required Scope                                            |
| --------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| Daily / weekly tests                                | Used to read manifest files in private repositories.                                                                                          | Read and modify your repositories and their pull requests |
| Snyk tests on pull requests                         | Used to send pull request status checks when a new PR is created, or an existing PR is updated.                                               | Read and modify your repositories and their pull requests |
| Opening fix and upgrade pull requests               | Used to create fix PRs in monitored repositories.                                                                                             | Read and modify your repositories and their pull requests |
| Snyk tests on pull requests - initial configuration | Used to add Snyk webhooks to the imported repos, to notify Snyk when pull requests are created or updated, and enable Snyk to trigger a scan. | Read and modify your repositories' webhooks               |

## Disabling the Bitbucket Cloud App integration

{% hint style="warning" %}
When you disconnect Snyk from your repository Projects, your credentials are removed from Snyk, and any integration-specific Projects that Snyk is monitoring are deactivated in Snyk.

If you choose to re-enable this integration later, you must re-enter your credentials and activate your Projects.
{% endhint %}

To disable this integration, in <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> (Organization settings)> **Integrations > Source Control > Bitbucket Cloud App:**

1. In your list of integrations, select the Bitbucket Cloud App integration you want to deactivate and click **Edit settings** to open a page with the current status of your integration.\
   \
   The page includes sections that are specific to each integration, where you can manage your credentials, API key, Service Principal, or connection details.
2. Scroll to the **Disconnect** section and click **Remove** to remove the integration.

{% hint style="info" %}
Disconnecting the integration from the Snyk side does not uninstall the app from your workspace in Bitbucket Cloud. To uninstall the Bitbucket app, go to your workspace settings in Bitbucket.org --> Installed Apps and remove the _Snyk Security for Bitbucket Cloud_ app.
{% endhint %}
