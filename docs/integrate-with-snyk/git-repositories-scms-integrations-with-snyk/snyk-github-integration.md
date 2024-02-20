# Snyk GitHub integration

{% hint style="info" %}
**Feature availability**

The Snyk GitHub integration is available for all Snyk customers regardless of plan level. See the [Plans and pricing page](https://snyk.io/plans/) for more details.
{% endhint %}

## Prerequisites for Snyk GitHub integration

* Internet-accessible repositories.\
  If your repositories are not internet-accessible, you must use [Snyk Broker](../../enterprise-configuration/snyk-broker/).
* A public or private GitHub project.

## Known limitations of the Snyk GitHub integration

You cannot use the Snyk GitHub integration with a Snyk [Service Account](../../enterprise-configuration/service-accounts/), as the GitHub integration is associated with your user account, not with the Snyk Organization.

Use the [GitHub Enterprise integration](snyk-github-enterprise-integration.md) to import public and private Projects using the API with a Snyk Service Account.

## Snyk GitHub integration features

The Snyk GitHub integration allows you to:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes and upgrades

## Per user, not per Organization

The GitHub integration is set up for your user account, not for a Snyk Organization. GitHub integration settings apply to all Organizations associated with your user account but do not automatically apply to other user accounts in an Organization.

When you import a Snyk Project using your GitHub integration with the Snyk PR functionality enabled, Snyk PRs are created for that Project. However, if another user imports Projects with their GitHub integration after disabling the Snyk PR functionality, Snyk PRs are not created for the Projects they import.

## How to set up the Snyk GitHub integration

To connect your GitHub repositories to Snyk for scanning, you need to set up the integration and then import Projects.&#x20;

See [Set up an integration](../../getting-started/quickstart/set-up-an-integration.md) and [Import a Project](../../getting-started/quickstart/import-a-project.md) for details of this process.

## Snyk GitHub integration settings

To see all settings for your GitHub integration, go to the Snyk GitHub Integration settings page, then navigate to Organization **Settings**, and select **GitHub** in the **Integrations** section:

<figure><img src="../../.gitbook/assets/Github-integrations-intro.png" alt="GitHub integration settings"><figcaption><p>GitHub integration settings</p></figcaption></figure>

You can then scroll down to the section required, and set the options accordingly:

* [General settings](snyk-github-integration.md#github-integration-general-settings)
* Pull requests:
  * [Open Snyk automatic PRs from a fixed GitHub account](snyk-github-integration.md#setting-an-account-to-open-snyk-prs)
  * [Automatic fix PRs](snyk-github-integration.md#project-monitoring-and-automatic-fix-pull-requests)
  * [Automatic dependency upgrade PRs](../../scan-with-snyk/snyk-open-source/automatic-and-manual-prs-with-snyk-open-source/upgrade-dependencies-with-automatic-prs.md)
  * [Pull request assignees](snyk-github-integration.md#pr-assignment)
  * [Snyk vulnerability patches](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-patches-to-fix-vulnerabilities.md)
* Pull request status checks ([PR Checks](../../scan-with-snyk/run-pr-checks/))
  * [Open Source security and licenses](../../scan-with-snyk/run-pr-checks/configure-pr-checks.md#configure-pr-checks-at-the-integration-level)
  * [Code analysis](../../scan-with-snyk/run-pr-checks/configure-pr-checks.md#configure-pr-checks-at-the-integration-level)
* Dockerfiles
  * [Scan you Dockerfile](../../scan-with-snyk/snyk-container/scan-your-dockerfile/)
  * [Fix vulnerable base images in your Dockerfile](../../scan-with-snyk/snyk-container/scan-your-dockerfile/fix-vulnerable-base-images-in-your-dockerfile.md)

## General Snyk GitHub integration settings

Select **General** to view general settings:

<div align="left">

<figure><img src="../../.gitbook/assets/Github-integrations-general.png" alt="GitHub general settings" width="563"><figcaption><p>GitHub general settings</p></figcaption></figure>

</div>

* **Integration ID**: The unique ID for this integration, needed if you use the [Snyk API](../../snyk-api/).
* **Repository access**: Whether Snyk can access private repos (in addition to public repos). Changing this setting affects existing Projects.

## Snyk GitHub integration features

After you have connected GitHub to Snyk, you can use:

* [Project-level security reports](snyk-github-integration.md#project-level-security-reports)
* [Project monitoring and automatic fix pull requests](snyk-github-integration.md#project-monitoring-and-automatic-fix-pull-requests)
* [Commit signing](snyk-github-integration.md#commit-signing)
* [Pull request testing](snyk-github-integration.md#pull-request-testing)

### **Project-level security reports**

{% hint style="info" %}
**Feature availability**\
Reports are available with Enterprise plans. See the [plans and pricing](https://snyk.io/plans/) page for details.
{% endhint %}

Snyk produces advanced [security reports](../../manage-issues/reporting/legacy-reports/legacy-reports-overview.md) that let you explore the vulnerabilities found in your repositories and fix them right away by opening a fix pull request directly in your repository, with the required upgrades or patches.

This example shows a Project-level security report.

<figure><img src="../../.gitbook/assets/project_lvl_security_rpt-18july2022.png" alt="Project-level security report"><figcaption><p>Project-level security report</p></figcaption></figure>

### **Project monitoring and automatic fix pull requests**

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you via email and opens automated pull requests with fixes for your repositories.

The example that follows shows a fix pull request opened by Snyk.

<figure><img src="../../.gitbook/assets/github_fix_pr_cropped-14july2022 (1) (1).png" alt="Fix pull request opened by Snyk"><figcaption><p>Fix pull request opened by Snyk</p></figcaption></figure>

To review and adjust the automatic fix pull request settings in the Snyk GitHub Integration settings page, navigate to Organization **Settings** **>** **Integrations > Source control > GitHub**.

Scroll down to the **Automatic fix PRs** section and set the options accordingly:

<div align="left">

<figure><img src="../../.gitbook/assets/Github-integrations-auto-fix-PRs.png" alt="Automatic fix pull request settings" width="563"><figcaption><p>Automatic fix pull request settings</p></figcaption></figure>

</div>

### Commit signing

{% hint style="info" %}
**Feature availability**\
For availability with Snyk Broker, see the [Commit signing](../../enterprise-configuration/snyk-broker/snyk-broker-commit-signing.md) page in the Broker docs.
{% endhint %}

All the commits in Snyk's pull requests are done by `snyk-bot@snyk.io` (a verified user on GitHub), and signed with a PGP key. All Snyk pull requests appear as verified on GitHub, thus providing your developers with the confidence that the fix and upgrade pull requests are generated by a trusted source.

### **Pull request status checks**

The Snyk [PR Checks](../../scan-with-snyk/run-pr-checks/) feature allows Snyk to test any new PR in your repositories for security vulnerabilities and sends a status check to GitHub. This lets you see, directly in GitHub, whether or not the pull request introduces new security issues.

This example shows how Snyk PR checks appear on the GitHub pull request page.

<figure><img src="../../.gitbook/assets/pr_testing-14july2022.png" alt="Snyk pull request checks on GitHub pull request page"><figcaption><p>Snyk pull request checks on GitHub pull request page</p></figcaption></figure>

You can review and adjust the pull request test settings using the Snyk GitHub Integration settings page in Organization **Settings** **>** **Integrations > Source control > GitHub**.

<div align="left">

<figure><img src="../../.gitbook/assets/Github-integrations-PR-request-status-checks.png" alt="Pull request status checks settings" width="563"><figcaption><p>Pull request status checks settings</p></figcaption></figure>

</div>

## Required permissions scope for the Snyk GitHub integration

The table that follows provides a summary of the required access scopes for GitHub integration. For information about the token in a brokered integration, see [GitHub - install and configure using Docker](../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md). For details about permissions in a non-brokered integration, refer to the information that follows this table.

| **Action**                                              | **Purpose**                                                                                                                                                                                                                                           | **Required permissions in GitHub** |
| ------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| Daily / weekly tests                                    | Used to read manifest files in private repositories.                                                                                                                                                                                                  | _repo (all)_                       |
| Manual fix pull requests (triggered by the user)        | Used to create fix PRs in the monitored repositories.                                                                                                                                                                                                 | _repo (all)_                       |
| Automatic fix and upgrade pull requests                 | Used to create fix or upgrade PRs in the monitored repositories.                                                                                                                                                                                      | _repo (all)_                       |
| Snyk tests on pull requests                             | Used to send pull request status checks whenever a new PR is created or an existing PR is updated.                                                                                                                                                    | _repo (all)_                       |
| Importing new projects to Snyk                          | Used to present a list of all the available repos in the GitHub org in the **Add Projects** screen (import popup).                                                                                                                                    | _admin:read:org, repo (all)_       |
| Snyk tests on pull requests - **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests when PRs are created, updated triggered, merged, and so on.</li><li>Send push events to trigger PR checks.</li></ul> | _admin:repo\_hooks (read & write)_ |

In non-brokered GitHub integrations, operations that are triggered via the Snyk Web UI, for example, opening a Fix PR or re-testing a Project, are performed on behalf of the acting user.&#x20;

Therefore, a user who wants to perform this operation on GitHub via the Snyk UI must connect their GitHub account to Snyk with the required permission scope for the repositories where they want to perform these operations. See the [Required permissions scope for repositories](snyk-github-integration.md#h\_01eefvj14p8b3depeffvyvdwzj) section for details.

Operations that are not triggered via the Snyk Web UI, such as daily and weekly tests and automatic PRs (fix and upgrade), are performed on behalf of random Snyk Organization members who have connected their GitHub accounts to Snyk and have the required permission scope for the repository.

For public repositories that are non-brokered, some operations, such as creating the PR, may occasionally be performed by `snyk-bot@snyk.io`.

{% hint style="info" %}
A Snyk Organization administrator can [designate a specific GitHub account to use for opening fix and upgrade PRs](introduction-to-git-repository-integrations/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).

Note that Snyk will continue to use a random Snyk Organization member's GitHub account to perform all the other operations. Therefore using this feature does not eliminate the need to connect users' GitHub accounts to Snyk.
{% endhint %}

## Required permission scope for repositories <a href="#h_01eefvj14p8b3depeffvyvdwzj" id="h_01eefvj14p8b3depeffvyvdwzj"></a>

For Snyk to perform the required operation on monitored repositories, that is, reading manifest files on a frequent basis and opening fix or upgrade PRs, the accounts that are connected to Snyk, either directly or via Snyk Broker, must have the following access to the repositories:

| **Action**                                              | **Purpose**                                                                                                                                                                                                                                             | **Required permissions on the repository** |
| ------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| Daily / weekly tests                                    | Used to read manifest files in private repos.                                                                                                                                                                                                           | _Read_ or higher                           |
| Snyk tests on pull requests                             | Used to send pull request status checks whenever a new PR is created or an existing PR is updated.                                                                                                                                                      | _Write_ or higher                          |
| Opening fix and upgrade pull requests                   | Used to create fix and upgrade PRs in the monitored repos.                                                                                                                                                                                              | _Write_ or higher                          |
| Snyk tests on pull requests - **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests (when PRs are created, updated triggered, merged, and so on).</li><li>Send push events to trigger PR checks.</li></ul> | _Admin_                                    |

## **How to set up a GitHub account to open Snyk PRs**

Snyk lets you designate a specific GitHub account to open fix and upgrade pull requests.

{% hint style="info" %}
The configured account is only used for opening PRs. All other operations are still performed on behalf of a randomly selected Snyk Organization member who has connected their GitHub accounts to Snyk.
{% endhint %}

To use this feature, follow these steps:

1. Go to the GitHub Integrations settings page in the Snyk Web UI using Organization **Settings** **>** **Integrations > Source control > GitHub.**
2. In the **Open Snyk automatic PRs from a fixed GitHub account** section, enter your GitHub personal access token.\
   You can [generate this from your GitHub account](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).
3. Click **Save** to enable this feature.

<div align="left">

<figure><img src="../../.gitbook/assets/image (5) (1) (3).png" alt="Set an account to open Snyk PRs" width="563"><figcaption><p>Set an account to open Snyk PRs</p></figcaption></figure>

</div>

{% hint style="info" %}
Ensure that the GitHub account that you designate to open Snyk PRs has **write-level** permissions or higher for the repos you want to monitor with Snyk.

See [repository permission levels on GitHub](snyk-github-integration.md#required-permissions-scope-for-the-github-integration) for more information.
{% endhint %}

## **How to assign pull requests to users** <a href="#pr-assignment" id="pr-assignment"></a>

{% hint style="info" %}
**Feature availability**

The Auto-assign PRs feature is supported only for private repositories.
{% endhint %}

Snyk can automatically assign the pull requests it creates to ensure that they are handled by the right team members.

Auto-assign for PRs can be enabled for the GitHub integration and all Projects imported via GitHub, or on a per-Project basis.

Users can either be manually specified, and all will be assigned, or automatically selected based on the last commit user account.

### **Enable Auto-assign for all Projects in the GitHub integration**

To configure the Auto-assign settings for all the Projects from an imported private repository, go to the Github integration settings using Organization **Settings** **>** **Integrations > Source control > GitHub** and select **Enable pull request assignees**.&#x20;

You can then choose to assign PRs to the last user to change the manifest file or specified contributors.

<div align="left">

<figure><img src="../../.gitbook/assets/image (7) (2) (1).png" alt="Auto-assign PRs in private repos"><figcaption><p>Auto-assign PRs in private repos</p></figcaption></figure>

</div>

### **Enable Auto-assign for a single Project**

To configure the Auto-assign settings for a specific Project from an imported private repository, follow these steps:

1. In the **Projects** tab for your Organization, select and expand the relevant private repository, select a Target, and click the **Settings** cog.\
   <img src="../../.gitbook/assets/image (56) (2) (3) (3) (3) (3) (4) (5) (5) (5) (5) (3) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (2).png" alt="Settings cog for target settings" data-size="original">\
   The Project page opens.
2. On the Project page, apply unique settings for that specific Project.\
   Select the **Settings** tab in the upper right and the **Github integration** \_\_ option in the left sidebar.
3. Go to the **Pull request assignees for private repos** section at the bottom of the page and choose to **Inherit from integration settings** or **Customize only for this Project**.
4. Ensure **Auto-assign PRs for this private Project** is enabled.
5. Choose to assign PRs to the last user to change the manifest file or named contributors.

<div align="left">

<figure><img src="../../.gitbook/assets/image (5) (2) (1).png" alt="Auto-assign PRs for this private Project" width="375"><figcaption><p>Auto-assign PRs for this private Project</p></figcaption></figure>

</div>

## How to disable the Snyk GitHub integration

The Snyk GitHub SCM integration leverages the OAuth app integration. If you integrated GitHub without using Snyk Broker, you can disconnect it by following these steps:

1. In GitHub, log in to the GitHub account that you used to create the integration.
2. Go to your GitHub account settings and select the **Applications** option in the left sidebar.
3. Select the **Authorized OAuth Apps** tab.\
   You can also reach the [Authorized OAuth Apps tab directly](https://github.com/settings/applications).
4. Find the **Snyk** entry, click the three (3) dots on the right, and select **Revoke**.

<figure><img src="../../.gitbook/assets/disconnect_github_3-july-2022.png" alt="Revoke OAuth authorization"><figcaption><p>Revoke OAuth authorization</p></figcaption></figure>

Revoking this access effectively disconnects Snyk’s access to that GitHub account.

* Existing imported snapshots will persist in Snyk and continue to be re-scanned based on the existing snapshots until deleted.
* Snyk will no longer be able to import new Projects from the GitHub integration and will no longer re-scan on new code merges.

In addition, you must confirm that Snyk is not enabled on any existing **Branch protection rules.**

{% hint style="info" %}
Note that branch protection is active only after a PR has been raised.
{% endhint %}

1. From the main page of your GitHub repository, go to **Settings > Branches > Branch protection rules.**
2. Ensure there are no **Status checks found in the last week for this repository.**

{% hint style="info" %}
A disconnected GitHub integration will still appear as configured in the Integrations menu of the Snyk UI. However, clicking on the integration settings will show that it is not connected. In this case, the "configured" integration can safely be ignored.&#x20;
{% endhint %}
