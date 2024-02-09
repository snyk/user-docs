# Snyk Git repositories: deployment recommendations

You can integrate Snyk with your Git repository to quickly and easily gain visibility across all the [Snyk Projects](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects) that you add to the **Projects** list in Snyk.

Snyk [Git repository integrations](../) allow you to:

* Continuously perform security scanning across all integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes

## Recommended deployment order

To ensure a smooth deployment for your team, follow the stages outlined below for each integration.

If you try to implement all the SCM integration features at the same time, you risk causing friction in your software development life cycle ([SDLC](https://snyk.io/learn/secure-sdlc/)), which in turn leads to a poor developer experience.

| **Deployment Stages**                                                                                                                                                                                                       | **Configurations**                                                                                                                                                                                                             | **Outcome**                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------- |
| <p><a href="snyk-scm-integration-good-practices.md#stage-1-set-up-your-scm-integration"><strong>Stage 1</strong></a></p><p>Set up your Git repository integration</p>                                                       | See the configuration documentation for each [Git repository integration](../)                                                                                                                                                 | A configured integration, ready to import Projects                                                        |
| <p><a href="snyk-scm-integration-good-practices.md#stage-2-import-projects"><strong>Stage 2</strong></a></p><p>Import all of the Projects from your Git repository.</p>                                                     | <p>• Test both public and private repos</p><p>• Disable all user notifications</p><p>• Disable Snyk PR checks</p><p>• Disable Auto-fix PRs</p><p>• Disable Failing PR checks</p><p>• Disable Automatic dependency upgrades</p> | A complete Software Bill of Materials (SBOM) that enables you to assess risk across your applications.    |
| <p><a href="snyk-scm-integration-good-practices.md#stage-3-enable-snyk-test-on-prs"><strong>Stage 3</strong></a></p><p>Developer and Security teams use Snyk prioritization reporting capabilities to build a fix plan.</p> | Watch this [video ](https://www.youtube.com/watch?v=\_kAY94JwQHY)to learn more about Snyk prioritization reporting.                                                                                                            | Alignment between Developers and Security on what should be fixed and when to streamline the fix process. |
| <p><a href="snyk-scm-integration-good-practices.md#stage-4-enable-blocking-prs"><strong>Stage 4</strong></a></p><p>Alert developers to issues in real time and educate them on available fixes.</p>                         | <p>• Enable Snyk PR checks and fail PRs if they contain high-severity issues with available fixes</p><p>• Enable Auto-fix PRs</p>                                                                                              | Improve your organization’s mean-time-to-fix (MTTF).                                                      |
| <p><a href="snyk-scm-integration-good-practices.md#stage-5-automatic-fix-prs"><strong>Stage 5</strong></a></p><p>Prevent developers from introducing any new vulnerabilities</p>                                            | Enable Failing PR checks for ANY high-severity issues (fix or no fix)                                                                                                                                                          | ‘Secure by Design’ methodology achieved.                                                                  |
| <p><a href="snyk-scm-integration-good-practices.md#stage-6-dependency-upgrade-prs"><strong>Stage 6</strong></a></p><p>Reduce technical security debt</p>                                                                    | Enable [Automatic dependency upgrades](../../../scan-with-snyk/snyk-open-source/automatic-and-manual-prs-with-snyk-open-source/upgrade-dependencies-with-automatic-prs.md)                                                     | Reduce future design headaches and hot fixes, which can be time-consuming to research and address.        |

## Stage 1: Set up your Git integration

Snyk has pre-built integrations for Git repositories, including GitHub, GitHub Enterprise, Bitbucket Cloud, and others.

See [Set up an integration](../../../getting-started/quickstart/set-up-an-integration.md) for details.

### **SCM permissions on repositories**

Operations triggered using the Snyk UI, such as opening a Fix PR or retesting a Project, are performed for the acting user. Thus to perform these operations, you must connect your own SCM user or service account. This gives Snyk the required permissions for the repositories where you want to perform these operations.

For example, in GitHub, the accounts connected to Snyk need the following access on the target repositories:

| **Action**                                              | **Purpose**                                                                                                                                  | **Required permissions on the repository** |
| ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| **Daily / weekly tests**                                | Used to read manifest files in private repos                                                                                                 | _**Write**_ or higher                      |
| **Snyk tests on pull requests**                         | Used to send pull request status checks when a new PR is created or an existing PR is updated                                                |                                            |
| **Opening fix and upgrade pull requests**               | Used to send fix or upgrade PRs in the monitored repos                                                                                       |                                            |
| **Snyk tests on pull requests - initial configuration** | Used to add Snyk webhooks to the imported repos, to notify Snyk when pull requests are created or updated, and enable Snyk to trigger a scan | _**Admin**_                                |

### **Change notification settings**

By default, Snyk emails every Org user when a new issue or fix in a Project’s dependencies is found, and provides each user with a weekly update of the security status across the Organization. If you plan to import many Projects to an Organization, to avoid sending multiple email notifications sent to users, consider disabling all the notifications for that Org.

To customize the emails your Org users receive, go to **Settings** <img src="../../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> > **Notifications > Notification Setting**. The changes you make here will affect all of your Organization’s members, although Org users can override these default settings in their user-level account settings.

To disable notifications for all the users in an Org ahead of your import, deselect the appropriate notification boxes. See [Manage notifications](../../../snyk-admin/manage-notifications.md) for more details.

## Stage 2: Import Projects

Navigate to the **Projects** page in the Snyk UI, select **Add projects**, select the repos to import to Snyk, and click **Add selected repositories**.

<figure><img src="../../../.gitbook/assets/image (2) (4) (1) (1) (1) (1) (1) (1) (1).png" alt="Add projects to Snyk"><figcaption><p>Add projects to Snyk</p></figcaption></figure>

* Snyk starts scanning the selected repos for dependency files (for example, **package.json**) in the entire directory tree and imports these files as Projects.
* Snyk evaluates root folders and any custom file locations defined. If no manifest or configuration files are found, Snyk alerts you that no files can be imported.
* Snyk detects the manifest files (Projects), tests them, then displays the results.\
  Imported Projects appear underneath the repository name.\
  After a Project is imported, it is continuously checked for vulnerabilities.

{% hint style="info" %}
To confirm that a Project was imported, go to the **Add project** import page for the integration. Imported Projects are indicated by a ✔ next to the repository name: <img src="../../../.gitbook/assets/image (223) (2) (2).png" alt="Check next to name of repo" data-size="line">
{% endhint %}

See [Import a Project](../../../getting-started/quickstart/import-a-project.md) for more details.

## Stage 3: Enable Snyk test on PRs

### **PR test settings and workflows**

By default, Snyk scans every pull request submitted on your monitored repositories and displays the results and recommendations grouped together in a single security check and a single license check.

<figure><img src="../../../.gitbook/assets/checks-passed.png" alt="Security check"><figcaption><p>Security check</p></figcaption></figure>

### **Status details**

Click the **Details** link to display the status of the Snyk check. The status options are:

* **Success**: no issues were identified and all checks passed
* **Processing**: this status is displayed until the Snyk test is completed
* **Failure**: identified issues that must be fixed for the check to pass
* **Error**: indicates that one of the following issues may have occurred:
  * Your manifest file may be out of sync
  * Snyk could not read the manifest file
  * Snyk could not find the manifest file

<figure><img src="../../../.gitbook/assets/security-check (1).png" alt="Snyk security check failure"><figcaption><p>Snyk security check failure</p></figcaption></figure>

### **Manage PR Check settings**

An administrator can manage Snyk [PR Checks](../../../scan-with-snyk/run-pr-checks/) settings for each SCM integration at the Organization level and then apply these settings, either to all the Projects for that integration or to selected specific Projects. You can configure whether this feature is on (enabled by default) and set fail conditions to define when Snyk should fail your PR checks.

See [Configure PR Checks at the Integration level](../../../scan-with-snyk/run-pr-checks/configure-pr-checks.md#configure-pr-checks-at-the-integration-level) and [Configure PR Checks at the Project level](../../../scan-with-snyk/run-pr-checks/configure-pr-checks.md#configure-pr-checks-at-the-project-level) for details of this process.&#x20;

{% hint style="info" %}
Use Snyk License policies to ensure that your Snyk PRs are free of licensing issues.\
See [Licenses](../../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md) for more information.
{% endhint %}

### **Initial step: get visibility and set fail conditions**

When you first roll out your SCM integration, Snyk recommends that you start with Snyk set up to test your PRs _without_ failing them, so that your developers can get used to seeing the Snyk commit check.

1. Decide to apply PR testing for your integrations at the Organization level, or for specific Projects.
2. Set the Fail conditions as described in the section [Manage PR Check settings](snyk-scm-integration-good-practices.md#manage-pr-check-settings):
   * Select the fail condition **Only fail when the PR is adding a dependency with issues**.
   * Check both **Only Fail for high or critical severity issues** and **Only** **fail when the issues found have a fix available**.

## Stage 4: Enable Blocking PRs

After you have embedded Snyk into your software development life cycle (SDLC), and have built good developer awareness, you can start to apply stricter policies to improve your overall security posture, for example:

* **Low priority Projects**: you can fail the PR only for new high-severity issues that are fixable.
* **Medium priority Projects**: fail the PR only for high-severity issues.
* **High priority Projects (PCI/GDPR compliance)**: fail the PR for any issue.

{% hint style="info" %}
To align vulnerability severity with your internal policy, use security policies to change the severity of issues and attach them to relevant Project attributes. See [Security policies](../../../scan-with-snyk/policies/security-policies/) for more details.
{% endhint %}

## Stage 5: Automatic Fix PRs

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk emails you and opens automated PRs with fixes for the repositories.

The example that follows shows a fix pull request opened by Snyk:

<figure><img src="../../../.gitbook/assets/mceclip0 (1).png" alt="Fix pull request opened by Snyk"><figcaption><p>Fix pull request opened by Snyk</p></figcaption></figure>

To configure the PR test settings for specific Projects, select the relevant Organization and go to Organization **Settings** <img src="../../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> > **Integrations > Edit Settings.**

<div align="left">

<figure><img src="../../../.gitbook/assets/image (4) (1) (1).png" alt="Automatic fix pull request settings" width="563"><figcaption><p>Automatic fix pull request settings</p></figcaption></figure>

</div>

Snyk recommends that you exclude patches from the auto fix PRs if your developers are not familiar with how to use and execute them.

Ask your developers to consider the merge advice label that appears on the auto fix PRs:

<div align="left">

<figure><img src="../../../.gitbook/assets/merge-advice-review-recommended (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (17) (20).png" alt="Merge advice label - review recommended" width="324"><figcaption><p>Merge advice label - review recommended</p></figcaption></figure>

</div>

<div align="left">

<figure><img src="../../../.gitbook/assets/advice-green (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (13) (18).png" alt="Merge advice label - high chance of success" width="333"><figcaption><p>Merge advice label - high chance of success</p></figcaption></figure>

</div>

<div align="left">

<figure><img src="../../../.gitbook/assets/merge-advice (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (29) (15).png" alt="Merge advice label - not enough data yet" width="312"><figcaption><p>Merge advice label - not enough data yet</p></figcaption></figure>

</div>

{% hint style="info" %}
Snyk auto fix PRs are generated only for new issues.
{% endhint %}

If your SCM is GitHub and you are not using Snyk Broker, then by default, Snyk rotates every Org user's credentials to open the auto fix PRs. You can change this if needed and set the user credentials to open the auto fix PRs. See [Opening fix and upgrade pull requests from a fixed GitHub account](opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md) for details.

## Stage 6 - Dependency Upgrade PRs

When your Group is ready to start tackling security technical debt, you can configure Snyk to automatically create pull requests (PRs) on your behalf to upgrade your dependencies.

<figure><img src="../../../.gitbook/assets/upgrade-node-uuid.png" alt="Automatic upgrade PR"><figcaption><p>Automatic upgrade PR</p></figcaption></figure>

### **How Automatic upgrade PRs work**

1. ProjectIntegration is configured, and users enable automatic upgrade PRs.
2. Snyk scans your Projects as you import them and continues to monitor your Projects, scanning on a regular basis.
3. For each scan, when Snyk identifies new versions for your dependencies:
   * Snyk creates automatic upgrade PRs with the frequency based on Snyk Project settings.
   * Snyk does not open a new upgrade PR for a dependency that is already changed (upgraded or patched) in another open Snyk PR.
   * Snyk opens separate PRs for each dependency.
   * Snyk does not create upgrade PRs for a repo that has five or more Snyk PRs open. If the limit of open PRs is reached, no new ones are created. This number can be set to between one and ten from the settings. This limit only applies when creating upgrade PRs, but does count fix PRs. Fix PRs are not limited in this way.
   * By default, Snyk recommends only patch and minor upgrades, but major version upgrades can be enabled in the settings where the feature is enabled.
   * If the latest eligible version contains vulnerabilities not already found in your Project, Snyk does not recommend an upgrade.
   * Snyk does not recommend upgrades to versions that are less than 21 days old. This is to avoid versions that introduce functional bugs and subsequently get unpublished, or versions that are released from a compromised account (where the account owner has lost control to someone with malicious intent).

### **Supported languages and repos**

Snyk currently supports automatic upgrade PRs for npm, Yarn, and Maven-Central Projects through GitHub, GitHub Enterprise Server, and BitBucket Cloud. Use with Snyk Broker is supported. For use with the Broker, ask your admin first to upgrade to v4.55.0 or later.

### **Enable automatic dependency upgrade PRs for a specific Project**

To set PR Settings on the Project level and override the PR settings configured at the Organization level:

1. Open the Organization for which you would like to enable automatic upgrade PRs and go to the **Projects** tab.
2.  Select and expand the relevant Project, select the relevant Target, and using the cog icon, open the **Settings**:

    <img src="../../../.gitbook/assets/image (56) (2) (3) (3) (3) (3) (4) (5) (5) (5) (5) (3) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (2).png" alt="Cog icon for Project settings" data-size="original">
3. In the Settings area, in the left panel menu, click the Integration settings to apply unique settings for that specific Project.
4. In settings that load, scroll to the **Automatic dependency upgrade pull requests** and click **Disabled**.
5. In the options that appear:
   1. Snyk creates PRs up to a maximum of ten open simultaneously per repo.\
      To limit this number further, select the maximum number of PRs from the dropdown list.\
      For more information, see [Upgrading dependencies with automatic PRs](../../../scan-with-snyk/snyk-open-source/automatic-and-manual-prs-with-snyk-open-source/upgrade-dependencies-with-automatic-prs.md).
   2. In the **Dependencies to ignore** field, enter the exact name of any dependencies that should _not_ be handled as part of the automatic functionality.\
      This field accepts only lowercase letters.
   3. After you click **Upgrade dependency settings**, each time Snyk scans this Project, Snyk automatically submits upgrade PRs based on the scan results.\
      If a newer version is released for an existing Snyk upgrade PR or for an existing fix PR, before Snyk can raise a new PR, the existing PR must be closed or merged.

<figure><img src="../../../.gitbook/assets/general-github_integration.png" alt="Automatic dependency pull request settings"><figcaption><p>Automatic dependency pull request settings</p></figcaption></figure>
