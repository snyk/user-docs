# Deployment recommendations

If you try to implement all the SCM integration features at the same time, you risk causing friction in your software development life cycle ([SDLC](https://snyk.io/learn/secure-sdlc/)), which in turn leads to a poor developer experience.

To ensure a smooth rollout of Snyk across your organization, Snyk provides a suggested deployment timeline consisting of deployment stages, configuration steps, and the desired outcome for each stage.

| Deployment stages                                                                                                                                                                                                                                                                               | Configurations and tests                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Outcome                                                                                                                                                                                                                                                                                                                                                          |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><a href="deployment-recommendations.md#stage-1-set-up-your-organization-and-group-level-scm-integrations"><strong>Stage 1:</strong></a><br>Set up an Organization-level SCM integration in your first Snyk Organization.<br><br>Set up a Group-level SCM integration in Snyk Essentials.</p> | <p>See the <a href="organization-level-integrations/#organization-level-snyk-scm-integrations">setup documentation</a> for the applicable SCM integration, and read the <a href="user-permissions-and-access-scopes.md">User permissions and access scope requirements</a> section to ensure accessibility for the right users in your Snyk hierarchy.<br><br>See the <a href="group-level-integrations/">setup documentation</a> for the applicable Essentials SCM integration.</p> | <p>A configured integration, ready to import Projects.<br><br>Essentials will detect your Git repositories so you can track your progress in rolling out Snyk scanning.</p>                                                                                                                                                                                      |
| <p><a href="deployment-recommendations.md#stage-2-import-projects"><strong>Stage 2:</strong></a><br>Import all of the relevant projects from your integrated SCM.</p>                                                                                                                           | <p>• Test both public and private repos</p><p>• Disable all user notifications</p><p>• Disable Snyk PR status checks</p><p>• Disable Automatic fix pull requests</p><p>• Disable Automatic dependency upgrade pull requests</p><p>• Disable Update Dockerfile base images</p>                                                                                                                                                                                                        | <p>Snyk will scan your imported repositories to detect all supported files and create them as Projects where you can view the vulnerabilities detected.<br><br>Disabling these features removes potential blockers for your initial setup and use of Snyk. This way, you have scheduled scans and high visibility of reporting without blocking deployments.</p> |
| <p><a href="deployment-recommendations.md#stage-3-enable-snyk-test-on-prs"><strong>Stage 3:</strong></a><br>Build a fix plan using Snyk prioritization reporting capabilities.</p>                                                                                                              | Learn more about [Snyk Priority Score](../../manage-risk/prioritize-issues-for-fixing/priority-score.md) and [Snyk Risk Score](../../manage-risk/prioritize-issues-for-fixing/risk-score.md) to understand how you can leverage these features to help prioritize which issues to fix.                                                                                                                                                                                               | Alignment between Developers and Security on what should be fixed and when to streamline the fix process.                                                                                                                                                                                                                                                        |
| <p><a href="deployment-recommendations.md#stage-4-enable-blocking-prs"><strong>Stage 4:</strong></a><br>Configure Snyk PR Checks and auto-fix PRs to alert developers to issues in real-time, with available fixes.</p>                                                                         | Enable [Snyk PR status checks](../../scan-with-snyk/pull-requests/) so developers know whether their PR contains high-severity issues with available fixes.                                                                                                                                                                                                                                                                                                                          | Make it easy for your developers to see new issues before they are merged into the repository, decreasing the number of new vulnerabilities added and enabling them to make decisions on blocking a merge.                                                                                                                                                       |
| <p><a href="deployment-recommendations.md#stage-5-automatic-fix-prs"><strong>Stage 5:</strong></a><br>Prevent developers from introducing any new vulnerabilities.</p>                                                                                                                          | Enable branch protection in your SCM tool so that PRs with failing Snyk PR status checks cannot be merged.                                                                                                                                                                                                                                                                                                                                                                           | This change will prevent PRs from being merged into your repository if Snyk detects new vulnerabilities. Over time, you can adjust the sensitivity of the check to include all high or critical severity issues, regardless of whether a fix is available.                                                                                                       |
| <p><a href="deployment-recommendations.md#stage-6-dependency-upgrade-prs"><strong>Stage 6:</strong></a><br>Enable Snyk automated features to fix issues and keep dependencies up-to-date to reduce technical security debt.</p>                                                                 | <p>Enable <a href="../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/#automated-snyk-prs">Automatic fix pull requests</a>.<br>Enable <a href="../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/">Automatic dependency upgrades</a>.</p>                                                                                                                                                               | Snyk will automatically raise PRs to fix issues with the highest severity. Keeping your dependencies up to date helps to reduce future design headaches and hot fixes, which can be time consuming to research and address.                                                                                                                                      |

## Stage 1: Set up your Organization and Group-level SCM integrations

Snyk has Organization-level SCM integrations available, including GitHub, GitHub Enterprise, Bitbucket Cloud, and more.

For details, see [Set up a Snyk integration](../../discover-snyk/getting-started/#set-up-a-snyk-integration).

Snyk has Group-level SCM integrations available for Snyk Essentials, including GitHub, GitLab, Azure DevOps, and Bitbucket.

See [Group level SCM integrations](group-level-integrations/) for details.

### **SCM permissions on repositories**

Operations triggered using the Snyk UI, such as opening a Fix PR or retesting a Project, are performed for the acting user. Thus to perform these operations, you must connect your own SCM user or service account. This gives Snyk the required permissions for the repositories where you want to perform these operations.

For details on these permissions, see [User permissions and access scope requirements](user-permissions-and-access-scopes.md) for your chosen SCM integration.

### **Change notification settings**

By default, Snyk emails every Org user when a new issue or fix in a Project’s dependencies is found, and provides each user with a weekly update of the security status across the Organization. If you plan to import many Projects to an Organization, to avoid sending multiple email notifications sent to users, consider disabling all the notifications for that Organization.

To customize the emails your Org users receive, go to **Settings** > **Notifications > Notification Setting**. The changes you make here will affect all of your Organization’s members, although Org users can override these default settings in their user-level account settings.

To disable notifications for all the users in an Org ahead of your import, deselect the appropriate notification boxes. See [Manage notifications](../../snyk-platform-administration/manage-notifications.md) for more details.

## Stage 2: Import Projects

Navigate to the **Projects** page in the Snyk UI, select **Add projects**, select the repos to import to Snyk, and click **Add selected repositories**.

* Snyk starts scanning the selected repos for dependency files (for example, **package.json**) in the entire directory tree and imports these files as Projects.
* Snyk evaluates root folders and any custom file locations defined. If no manifest or configuration files are found, Snyk alerts you that no files can be imported.
* Snyk detects the manifest files (Projects), tests them, then displays the results.\
  Imported Projects appear underneath the repository name.\
  After a Project is imported, it is continuously checked for vulnerabilities.

{% hint style="info" %}
To confirm that a Project was imported, go to the **Add project** import page for the integration. Imported Projects are indicated by a ✔ next to the repository name.
{% endhint %}

For details, see [Import a Project](../../discover-snyk/getting-started/#import-a-project-to-scan-and-identify-issues).

## Stage 3: Enable Snyk test on PRs

### **PR test settings and workflows**

By default, Snyk scans every pull request submitted on your monitored repositories and displays the results and recommendations grouped together in a single security check and a single license check.

### **Status details**

Click the **Details** link to display the status of the Snyk check. The status options are:

* **Success**: no issues were identified and all checks passed
* **Processing**: this status is displayed until the Snyk test is completed
* **Failure**: identified issues that must be fixed for the check to pass
* **Error**: indicates that one of the following issues may have occurred:
  * Your manifest file may be out of sync
  * Snyk could not read the manifest file
  * Snyk could not find the manifest file

### Manage PR Check settings

An administrator can manage Snyk [PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/) settings for each SCM integration at the Organization level and then apply these settings, either to all the Projects for that integration or to selected specific Projects. You can configure whether this feature is on (enabled by default) and set fail conditions to define when Snyk should fail your PR checks.

See [Configure PR Checks at the Integration level](../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level) and [Configure PR Checks at the Project level](../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-pr-checks-at-the-project-level) for details of this process.&#x20;

{% hint style="info" %}
Use Snyk License policies to ensure that your Snyk PRs are free of licensing issues.\
See [Licenses](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md) for more information.
{% endhint %}

### **Initial step: get visibility and set fail conditions**

When you first roll out your SCM integration, Snyk recommends that you start with Snyk set up to test your PRs _without_ failing them, so that your developers can get used to seeing the Snyk commit check.

1. Decide to apply PR testing for your integrations at the Organization level, or for specific Projects.
2. Set the Fail conditions as described in the section [Manage PR Check settings](deployment-recommendations.md#manage-pr-check-settings):
   * Select the fail condition **Only fail when the PR is adding a dependency with issues**.
   * Check both **Only Fail for high or critical severity issues** and **Only** **fail when the issues found have a fix available**.

## Stage 4: Enable Blocking PRs

After you have embedded Snyk into your software development life cycle (SDLC), and have built good developer awareness, you can start to apply stricter policies to improve your overall security posture, for example:

* **Low priority Projects**: you can fail the PR only for new high-severity issues that are fixable.
* **Medium priority Projects**: fail the PR only for high-severity issues.
* **High priority Projects (PCI/GDPR compliance)**: fail the PR for any issue.

{% hint style="info" %}
To align vulnerability severity with your internal policy, use security policies to change the severity of issues and attach them to relevant Project attributes. See [Security policies](../../manage-risk/policies/security-policies/) for more details.
{% endhint %}

## Stage 5: Automatic Fix PRs

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk emails you and opens automated PRs with fixes for the repositories.

To configure the automatic fix PR settings for all Projects in an Organization, select the relevant Organization and navigate to Organization **Settings** > **Integrations > Edit Settings.**

<div align="center"><figure><img src="../../.gitbook/assets/image (101).png" alt="Automatic fix pull request settings" width="563"><figcaption><p>Automatic fix pull request settings</p></figcaption></figure></div>

{% hint style="info" %}
Settings can also be configured per Project by selecting the specific Project in your Organization and navigating to the **Settings** tab. For more information, see [View and edit Project settings](../../snyk-platform-administration/snyk-projects/view-and-edit-project-settings.md).
{% endhint %}

Snyk recommends that you exclude patches from the auto fix PRs if your developers are not familiar with how to use and execute them.

Ask your developers to consider the merge advice label that appears on the auto fix PRs.

{% hint style="info" %}
Snyk auto fix PRs are generated only for new issues.
{% endhint %}

If your SCM is GitHub and you are not using Snyk Broker, then by default, Snyk rotates every Org user's credentials to open the auto fix PRs. You can change this if needed and set the user credentials to open the auto fix PRs. See [Opening fix and upgrade pull requests from a fixed GitHub account](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md) for details.

## Stage 6 - Dependency Upgrade PRs

When your Group is ready to start tackling security technical debt, you can configure Snyk to automatically create pull requests (PRs) on your behalf to upgrade your dependencies.

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

Snyk supports automatic upgrade PRs for npm, Yarn, and Maven-Central Projects through GitHub, GitHub Enterprise Server, and BitBucket Cloud. Use with Snyk Broker is supported. For use with the Broker, ask your admin first to upgrade to v4.55.0 or later.

### **Enable automatic dependency upgrade PRs for a specific Project**

To set PR Settings on the Project level and override the PR settings configured at the Organization level:

1. Open the Organization for which you would like to enable automatic upgrade PRs and go to the **Projects** tab.
2. Select and expand the relevant Project, select the relevant Target, and using the cog icon, open the **Settings.**
3. In the Settings area, in the left panel menu, click the Integration settings to apply unique settings for that specific Project.
4. In settings that load, scroll to the **Automatic dependency upgrade pull requests** and click **Disabled**.
5. In the options that appear:
   1. Snyk creates PRs up to a maximum of ten open simultaneously per repo.\
      To limit this number further, select the maximum number of PRs from the dropdown list.\
      For more information, see [Upgrading dependencies with automatic PRs](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/).
   2. In the **Dependencies to ignore** field, enter the exact name of any dependencies that should _not_ be handled as part of the automatic functionality.\
      This field accepts only lowercase letters.
   3. After you click **Upgrade dependency settings**, each time Snyk scans this Project, Snyk automatically submits upgrade PRs based on the scan results.\
      If a newer version is released for an existing Snyk upgrade PR or for an existing fix PR, before Snyk can raise a new PR, the existing PR must be closed or merged.
