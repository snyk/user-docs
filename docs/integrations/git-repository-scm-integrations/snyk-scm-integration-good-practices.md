# Snyk Git repository integration: deployment recommendations

You can integrate Snyk with your Git repository, to quickly and easily gain visibility across all the [Snyk Projects](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects) that you add to the **Projects** menu in Snyk.

Snyk's [Git repository integrations](https://docs.snyk.io/integrations/git-repository-scm-integrations) allow you to:

* Continuously perform security scanning across all integrated repositories
* Detect vulnerabilities in your open source components
* Provide automated fixes.

## Recommended deployment order

To ensure a smooth deployment for your team, follow the stages outlined below for each integration.

If you try to implement all the SCM integration features at the same time, you risk causing friction in your software development life cycle ([SDLC](https://snyk.io/learn/secure-sdlc/)), which in turn leads to a poor developer experience.

| **Deployment Stages**                                                                                                                                                                                                       | **Configurations**                                                                                                                                                                                                             | **Outcome**                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------- |
| <p><a href="snyk-scm-integration-good-practices.md#stage-1-set-up-your-scm-integration"><strong>Stage 1</strong></a></p><p>Set up your Git repository integration</p>                                                       | See the configuration documentation for each [Git repository integration](./)                                                                                                                                                  | A configured integration, ready to import projects                                                      |
| <p><a href="snyk-scm-integration-good-practices.md#stage-2-import-projects"><strong>Stage 2</strong></a></p><p>Import all of the projects from your Git repository.</p>                                                     | <p>• Test both public and private repos</p><p>• Disable all user notifications</p><p>• Disable Snyk PR checks</p><p>• Disable Auto-fix PRs</p><p>• Disable Failing PR checks</p><p>• Disable Automatic dependency upgrades</p> | A complete Software Bill of Materials (SBOM) that enables you to assess risk across your applications.  |
| <p><a href="snyk-scm-integration-good-practices.md#stage-3-enable-snyk-test-on-prs"><strong>Stage 3</strong></a></p><p>Developer and Security teams use Snyk prioritization reporting capabilities to build a fix plan.</p> | Watch this [video ](https://www.youtube.com/watch?v=\_kAY94JwQHY)to learn more about Snyk prioritization reporting.                                                                                                            | Alignment between Developers and Security on what should be fixed and when streamlines the fix process. |
| <p><a href="snyk-scm-integration-good-practices.md#stage-4-enable-blocking-prs"><strong>Stage 4</strong></a></p><p>Alert developers to issues in real-time, and educate them on available fixes.</p>                        | <p>• Enable Snyk PR checks and fail PRs if they contain High severity issues with available fixes</p><p>• Enable Auto-fix PRs</p>                                                                                              | Improve your organization’s mean-time-to-fix (MTTF).                                                    |
| <p><a href="snyk-scm-integration-good-practices.md#stage-5-automatic-fix-prs"><strong>Stage 5</strong></a></p><p>Prevent developers from introducing any new vulnerabilities</p>                                            | • Enable Failing PR checks, for ANY High severity issues (fix or no fix)                                                                                                                                                       | ‘Secure by Design’ methodology achieved.                                                                |
| <p><a href="snyk-scm-integration-good-practices.md#stage-6-dependency-upgrade-prs"><strong>Stage 6</strong></a></p><p>Reduce technical security debt</p>                                                                    | • Enable [Automatic dependency upgrades](https://docs.snyk.io/products/snyk-open-source/dependency-management/upgrading-dependencies-with-automatic-prs)                                                                       | Reduce future design headaches, and hot fixes, which can be time-consuming to research & address.       |

### Stage 1: Set up your Git integration

Snyk has pre-built integrations for Git repositories, including GitHub, GitHub Enterprise, Bitbucket Cloud and others.

See [Set up an integration](../../getting-started/quickstart/set-up-an-integration.md) for details.

#### **SCM permissions on repositories**

Operations triggered using the Snyk UI (such as opening a Fix PR or retesting a project) are performed for the acting user. So to perform these operation, you must connect your own SCM user or service account. This gives Snyk the required permissions for the repositories for where you would like to perform these operations.

For example, in GitHub, the accounts connected to Snyk need the following access on the target repositories:

| **Action**                                              | **Purpose**                                                                                                                                  | **Required permissions on the repository** |
| ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| **Daily / weekly tests**                                | Used to read manifest files in private repos                                                                                                 | _**Write**_ or higher                      |
| **Snyk tests on pull requests**                         | Used to send pull request status checks when a new PR is created or an existing PR is updated                                                |                                            |
| **Opening fix and upgrade pull requests**               | Used to send fix or upgrade PRs in the monitored repos                                                                                       |                                            |
| **Snyk tests on pull requests - initial configuration** | Used to add Snyk webhooks to the imported repos, to notify Snyk when pull requests are created or updated, and enable Snyk to trigger a scan | _**Admin**_                                |

#### **Change notification settings**

By default, Snyk emails every Org user when a new issue or fix in a project’s dependencies is found, and provides them with a weekly update of the security status across your organization. If you plan to import many projects to an Org, to avoid sending multiple email notifications sent to users, consider disabling all the notifications for that Org.

To customize the emails your Org users receive, go to settings (<img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line">) > **Notifications > Notification Setting**. The changes you make here will affect all of your organization’s members, although Org users can override these default settings in their user-level account settings.

To disable notifications for all the users in an Org ahead of your import, deselect the appropriate notification boxes. See [Notification management](https://docs.snyk.io/features/user-and-group-management/notifications) for more details.

### Stage 2: Import Projects

Go to the **Projects** page in the Snyk UI, select **Add projects**, select the repos to import to Snyk, and click **Add selected repositories**.

![](<../../.gitbook/assets/image (2) (4) (1) (1) (1) (1).png>)

* Snyk starts scanning the selected repos for dependency files (for example, **package.json**) in the entire directory tree and imports these files as projects.
* Snyk evaluates root folders and any custom file locations defined. If no manifest or configuration files are found, Snyk alerts you that no files can be imported.
* Snyk detects the manifest files (projects), tests them, then displays the results.\
  Imported projects appear underneath the repository name.\
  After a project is imported, it is continuously checked for vulnerabilities.

{% hint style="info" %}
To confirm that a project was imported, go to the **Add project** import page for the integration. Imported projects are indicated by a ✔ next to the repo name: <img src="../../.gitbook/assets/image (256) (1) (1) (1) (1) (1) (1) (1).png" alt="" data-size="line">
{% endhint %}

See [Import a Project](../../getting-started/running-tests/import-a-project.md) for more details.

### Stage 3: Enable Snyk test on PRs

#### **PR Test Settings & Workflows**

By default, Snyk scans every pull request submitted on your monitored repositories and displays the results and recommendations grouped together in a single security check and a single license check.

![](../../.gitbook/assets/checks-passed.png)

#### **Status details**

Click the **Details** link to display the status of the Snyk check. The status options are:

* **Success**: no issues were identified and all checks passed
* **Processing**: this status is displayed until the Snyk test is completed
* **Failure**: identified issues that must be fixed for the check to pass
* **Error**: indicates that one of the following issues may have occurred:
  * Your manifest file may be out of sync
  * Snyk couldn't read the manifest file
  * Snyk couldn't find the manifest file

![](../../.gitbook/assets/security-check.png)

#### **Manage PR test settings**

An administrator can manage settings for Snyk PR tests for each SCM integration at the organization level and apply them to all the projects for that integration, or select specific projects to apply the PR tests. You can configure whether this feature is on (enabled by default) and set fail conditions to define when Snyk should fail your PR checks.

To configure the PR test settings for your Organization:

1. For the relevant organization, go to (Organization Settings) <img src="../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> **>** **Integrations**, select the relevant integration, and click **Edit Settings**.
2. Set the the toggle to **Enabled** and set the **Fail conditions** as needed:
3. Click **Update settings**.

![](../../.gitbook/assets/image13.png)

To manage the PR test settings per project

1. To configure the pull request test settings for a specific project, go to **Projects** > **Projects Settings > Edit Settings.**\
   \*\*\*\*The project page opens.
2. In the project page, select the **Settings** tab in the upper right, and the SCM integration \_\_ option in the left sidebar.
3. Select and confirm the relevant options in each of the sections.

![](../../.gitbook/assets/main.png)

{% hint style="info" %}
Use Snyk [License policies](https://docs.snyk.io/products/snyk-open-source/license-policies) to ensure that your Snyk PRs are free of licensing issues.\
See [Licenses](https://docs.snyk.io/products/snyk-open-source/licenses) for more information.
{% endhint %}

#### **Initial step: get visibility and set fail conditions**

When you first roll out your SCM integration, we recommend that you start with Snyk set up to test your PRs _without_ failing them, so that your developers can get used to seeing the Snyk commit check.

1. Decide to apply PR testing for your integrations at the organization level, or for specific projects.
2. Set the conditions (as described above):
   * **Fail only** for **Only fail when the PR is adding a dependency with issues**.
   * Check both **Only Fail for high severity issues** and Only **fail when the issues found have a fix available**

### Stage 4: Enable Blocking PRs

After you’ve embedded Snyk into your software development life cycle (SDLC), and have built good developer awareness, you can start to apply stricter policies to improve your overall security posture. For example:

* **Low priority projects**: you can fail the PR only for new high severity that are fixable.
* **Medium priority projects**: fail the PR only for high severity issues.
* **High priority projects (PCI/GDPR compliance)**: fail the PR for any issue.

{% hint style="info" %}
To align vulnerability severity with your internal policy, use security policies to change the severity of issues and attach them to relevant projects attributes. See [Security policies](https://docs.snyk.io/features/fixing-and-prioritizing-issues/security-policies) for more details.
{% endhint %}

### Stage 5: Automatic Fix PRs

Snyk scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk emails you and opens automated PRs with fixes for the repositories.

The example below presents a fix pull request opened by Snyk:

![](<../../.gitbook/assets/mceclip0 (1).png>)

To configure the PR test settings for specific projects, select the relevant organization and go to (Organization settings) <img src="../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> > **Integrations > Edit Settings.**

![](../../.gitbook/assets/automatic.png)

We recommend that you exclude patches from the auto fix PRs, if your developers are not familiar with how to use and execute them.

You should ask your developers to consider the merge advice label that appears on the auto fix PRs:

![](<../../.gitbook/assets/merge-advice-review-recommended (2) (2) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1 (18).png>)

![](<../../.gitbook/assets/advice-green (1) (2) (2) (4) (3) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) ( (19).png>)

![](<../../.gitbook/assets/merge-advice (2) (2) (4) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (  (9).png>)

{% hint style="info" %}
Snyk auto fix PRs are only generated for new issues.
{% endhint %}

If your SCM is Github and you are not using Snyk Broker, then by default, Snyk rotates every Org user's credentials to open the auto fix PRs. You can change this if needed and set the user credentials to open the auto fix PRs. See [Opening fix and upgrade pull requests from a fixed GitHub account](https://docs.snyk.io/integrations/git-repository-scm-integrations/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account) for details.

### Stage 6 - Dependency Upgrade PRs

When your group is ready to start tackling security technical debt, you can configure Snyk to automatically create pull requests (PRs) on your behalf to upgrade your dependencies.

![](../../.gitbook/assets/upgrade-node-uuid.png)

#### **How Automatic upgrade PRs work**

1. Integration is configured and users enable automatic upgrade PRs.
2. Snyk scans your projects as you import them and continues to monitor your projects, scanning on a regular basis.
3. For each scan, when Snyk identifies new versions for your dependencies:
   * Snyk creates automatic upgrade PRs (frequency based on Snyk project settings)
   * Snyk does not open a new upgrade PR for a dependency that is already changed (upgraded or patched) in another open Snyk PR.
   * Snyk opens separate PRs for each dependency.
   * Snyk does not create upgrade PRs for a repo that has 5 or more Snyk PRs open - if the limit of open PRs is reached, no new ones are created. This number can set to between 1-10 from the Settings. This limit only applies when creating upgrade PRs, but does count fix PRs. Fix PRs are not limited in this way.
   * By default, Snyk recommends only patch and minor upgrades, but major version upgrade can be enabled in the settings where the feature is enabled.
   * If the latest eligible version contains vulnerabilities not already found in your project, Snyk does not recommend an upgrade.
   * Snyk does not recommend upgrades to versions that are less than 21 days old. This is to avoid versions that introduce functional bugs and subsequently get unpublished, or versions that are released from a compromised account (where the account owner has lost control to someone with malicious intent).

#### **Supported languages and repos**

Snyk currently supports this feature for npm, Yarn, and Maven-Central projects through GitHub, GitHub Enterprise Server, and BitBucket Cloud, including use of the Snyk Broker. For use with the Broker, your admin should first upgrade to v4.55.0 or later.

#### **Enable automatic dependency upgrade PRs for a specific project**

To set PR Settings on the project level and override the PR settings configured at organization level:

1. Open the organization for which you would like to enable automatic upgrade PRs and go to the **Projects** tab.
2.  Select and expand the relevant project, select the relevant target, and click the **Settings** cog:

    <img src="../../.gitbook/assets/image (56) (2) (3) (3) (3) (3) (4) (5) (5) (5) (5) (3) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (6).png" alt="" data-size="original">
3. In the Settings area, in the left panel menu, click the Integration settings to apply unique settings for that specific project.
4. In settings that load, scroll to the **Automatic dependency upgrade pull requests** and click **Disabled**.
5. In the options that appear:
   1. Snyk creates PRs up to a maximum of 10 open simultaneously - per repo.\
      To limit this number further, select the maximum number of PRs from the dropdown list.\
      For more information, see [Upgrading dependencies with automatic PRs](https://docs.snyk.io/snyk-open-source/dependency-management/upgrading-dependencies-with-automatic-prs).
   2. In the **Dependencies to ignore** field, enter the exact name of any dependencies that should _not_ be handled as part of the automatic functionality.\
      This field accepts only lower case letters.
   3. After you click **Upgrade dependency settings**, each time Snyk scans this project, it will automatically submit upgrade PRs based on the scan results.\
      If a newer version is released for an existing Snyk upgrade PR or for an existing fix PR, before Snyk can raise a new PR, the existing PR must be closed or merged.

![](../../.gitbook/assets/general-github\_integration.png)
