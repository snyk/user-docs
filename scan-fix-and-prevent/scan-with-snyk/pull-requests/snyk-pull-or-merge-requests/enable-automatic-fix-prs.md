---
description: How to enable automatic Snyk Fix PRs
---

# Enable automatic Fix PRs

{% hint style="info" %}
**Feature availability**

The automatic Fix PRs feature is supported for the following SCM integrations: GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab, and Azure Repos.

The automatic Fix PR settings can vary depending on the integration.
{% endhint %}

The following rules are applied to the creation of automatic PRs for vulnerabilities:

* Pull requests are created based on the **Test & Automated Pull Request Frequency** setting, which you can configure from the Project details page, by navigating to **Settings**.
* If you click **Retest now** for a Project, a manual scan runs. The 24-hour window is marked as having had the scan run, and no automatic PR is created until the next automated scan.
* One pull request is created per Project.
* A new vulnerability is a vulnerability in the current recurring scan that was not present in the previous scan of the same Project.
* If either the vulnerability is new and has a fix available or the fix is new and is not ignored, a Fix PR can be created.
* Fixing a vulnerability by upgrading a package may sometimes introduce a new vulnerability. Snyk will only automatically create a pull request if the fixed vulnerabilities are of a higher severity than any new ones introduced.

For known vulnerabilities, visit [Configure Automatic Backlog PRs.](enable-automatic-backlog-prs-for-previously-known-vulnerabilities.md)

To determine when your last 24-hour window began, check the Project issue card for **Snapshot taken by recurring test**, and check your email for **\[snyk] Vulnerability alert** for specific scan results.

Pull requests for new vulnerabilities are enabled by default for new integrations.

Visit the [Git repository SCM integrations](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations) pages for details about supported integrations.

## Enable or disable pull requests for an integration

Follow these steps to enable pull requests at the global integration level in an Organization:

1. From the Project details page, navigate to the **Settings** tab > **Integrations**.
2. Select an SCM integration, for example, GitHub.
3. In the **Automatic Fix PRs** section, enable **New vulnerabilities**.
4. Click **Save**.

Applying changes to all overridden Projects updates all individual Project settings for Automatic fix PRs. Clicking the button overrides individual Project settings with the global setting, even if a Project previously had its own configuration for automatic fix pull requests.

If you check **Fix all vulnerabilities for the same dependency in a single PR**, this adds any PR opened to address an issue with any potential PR for upgrades which also fix the same issue. Fixing an instance of the vulnerability often fixes other vulnerabilities as well.

## Set creation thresholds for score and severity

For every new actionable vulnerability found on each recurring test, Snyk raises a Fix PR. This may not be ideal depending on the velocity you are looking for in your Organization, so setting up specific criteria to match your needs can be achieved by setting thresholds.

To decide which automatic Fix PRs are visible to you, you can set a custom threshold for **Score** or **Severity**. Depending on your Organization's configuration, the dropdown contains either the **Risk Score** or the **Priority Score**.

### Score threshold

Snyk creates Fix PRs only above the threshold you set in the **Score** field. The selected score appears in the dropdown. This threshold ranges from 0 to 1,000.

The set defaults for score are as follows:

* Organizations created before December 5, 2024, have a default score of 0.
* Organizations created after December 5, 2024, have a default score of 700.
* After June 5, 2025, all Organizations have a default score of 700 unless you have configured a score in the settings.

### Severity threshold

Snyk creates Fix PRs for the severity levels selected, for example, **Critical** and **High**.

## Enable or disable pull requests for a single Project

Enabling or disabling at the Project level overrides the pull request setting for this single Project, so it will not inherit from the global integration setting.

1. From the Project details page, navigate to the **Settings** tab.
2. Select an SCM integration, for example, GitHub.
3. In the **Automatic fix pull requests** section:
   * Select **Customize for only this project**.
   * Enable **New vulnerabilities**.
   * Select **Save changes**.

## Automatic closure of obsolete Fix PRs

Snyk automatically closes Fix PRs when all vulnerabilities they target are resolved. This behavior keeps your PR backlog focused on active issues.

{% hint style="info" %}
Auto close of obsolete Fix and backlog PRs is enabled by default. You do not need extra configuration.
{% endhint %}

**Auto Close Pull Requests** is available on SCM providers that support Snyk Fix PRs, including GitHub, GitLab, Bitbucket, and Azure Repos.

### How Auto Close Pull Requests works

During each recurring test, Snyk evaluates open Fix PRs against the latest Project state:

1. Snyk compares the issue IDs targeted by each open Fix PR with the issues in the latest Project scan.
2. If all targeted issues are already resolved, Snyk marks the PR as obsolete. This can happen after a manual fix, a transitive dependency update, a dependency removal, or a vulnerability data change.
3. Snyk closes the obsolete PR in your SCM and then posts a comment that lists the resolved issues.

### Fix PR closure comments

When Snyk closes an obsolete Fix PR, it adds a comment to the closed PR thread. The comment:

* confirms that Snyk closed the PR because it no longer detects the targeted issues
* explains how the issues were resolved, such as a manual update, a transitive dependency update, or a vulnerability database change
* lists the resolved Snyk issue IDs and explains that you can reopen the PR if needed

### Safeguards

Snyk includes safeguards to avoid disruptive behavior:

| Safeguard                | Description                                                                                                           |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------- |
| Limit per test execution | Snyk closes up to 10 obsolete PRs during each recurring test execution.                                               |
| Comment after closure    | Snyk posts the closure comment only after it closes the PR successfully. This prevents orphaned comments on open PRs. |

### Turn off Auto Close Pull Requests

1. Navigate to **Settings** > **Integrations**.
2. Turn off Automatic closure of obsolete Fix PRs.
3. Snyk stops evaluating and closing obsolete Fix PRs in the next recurring test cycle.

### How to enable the automatic upgrade PRs option for an Organization

Project Settings inherit Organization Settings until they are explicitly changed on the Project Settings page. To force push Project Setting overrides at the Group level, they must be re-saved with the option "Apply changes to all overridden projects."

To enable Automatic Upgrade PRs across an entire Organization, follow these steps:

1. In the Snyk web UI, open the desired Organization.
2. Navigate to **Settings** > **Organization settings**. Find and click your configured SCM in the **Integrations** sidebar.
3. On the **Settings** page of the selected integration, navigate to the **Automatic dependency upgrade PRs** section.
4. In this section, perform the following actions:
   * Slider - change to **Enable**.
   * **Maximum number of open upgrade PRs allowed** – define how many open Snyk PRs (fix, upgrade, and backlog) a Project can have to also receive a dependency upgrade PR; the maximum is ten. When the limit of the open PRs is reached, no new upgrade PRs are created.
   * **Include major version in upgrade recommendation** – select whether to include major version upgrades in the recommendations. By default, only patches and minor versions are included in the upgrade recommendations.
   * **Dependencies to ignore** – enter the exact name of the dependencies that should not be included in the **Automatic upgrade** operation. You can only enter lowercase letters.
5. To save and apply your changes, click one of the following from the **Save** dropdown:
   * **Save** – your changes are saved and will be applied to all the Projects in the Organization that are configured to inherit these Settings from the Organization. Projects that have Custom Settings will not be influenced by this change.
   * **Save changes and apply to all overridden Projects** – your changes are saved and will be applied to all the Projects in the Organization. Projects that have Custom Settings will inherit these Organization Settings, and their Custom Settings will be overridden.

From now on, every time Snyk scans any Project in the Organization, it automatically submits Upgrade PRs if the scan discovers that an upgrade is available.

If a newer version is released for an existing Snyk Upgrade PR, or for an existing Fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.

### How to enable the automatic upgrade PRs option for a Project

The Settings on the Project level typically override the Settings on the Organization level. If an Organization owner forces settings overrides, Project users need to reconfigure their settings. We recommend alignment between Project and Organization owners to avoid situations like this.

Follow these steps to configure automatic upgrade PRs for a specific Project:

1. From the Snyk web UI, open the Organization that includes the Project you want to configure.
2. In the list of Projects, locate and expand the **Project** for which you want to enable automatic upgrade PRs.
3. Click the **Project settings** at the end of the Project row.
4. On the **Project** **Settings** page, select the integration you are using.
5. On the **Integration** page, scroll to the **Automatic dependency upgrade pull requests** section and choose one of the following:
   * **Inherit from Integration settings** – apply the Integration Settings of the Organization to the selected Project.\
     If the **Automatic dependency upgrade PRs option is disabled for the Organization**, this option will also be disabled for the Project.
   * **Customize for only this Project** – apply specific settings of the **Automatic dependency upgrade PRs** option on the Project. If you click this option:
     * Change the slider to **Enabled**.
     * In **Include major version in upgrade recommendation,** click one of the available options to define whether major version upgrades will be included in the recommendations.\
       By default, only patches and minor versions are included in the upgrade recommendations.
     * In **Limit Snyk to this many dependency upgrade PRs open simultaneously,** define how many open Snyk PRs a Project can have to also receive a dependency upgrade PR. You can set a number between 1 and 10.\
       When the limit of the open PRs is reached, no new upgrade PRs are created.\
       By default, to _also_ receive a dependency upgrade PR, a Project can have _up to four_ open PRs.
     * In **Dependencies to ignore**, enter the exact name of the dependencies to _exclude_ from the **Automatic upgrade** operation.\
       You can only enter lowercase letters.
     * Click **Update dependency upgrade settings** to save your changes.

After you have completed these steps, Snyk scans the Project and automatically submits Upgrade PRs if the scan discovers that an upgrade is available. If a newer version is released for an existing Snyk Upgrade PR or an existing Fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.

## Troubleshooting

### What Snyk considers resolved

Snyk considers an issue resolved when it no longer appears in the latest Project scan. This can happen when:

* you manually update the vulnerable dependency
* a transitive dependency update resolves the vulnerability
* you remove the vulnerable dependency from the Project
* Snyk updates the vulnerability data, for example, by withdrawing or reclassifying the issue

### Can you reopen an automatically closed PR?

Yes. If Snyk closes a PR in error, you can reopen it directly in your SCM. After you reopen the PR, Snyk does not close it again.

### Will you receive notifications when Snyk closes a PR?

Snyk posts a closure comment on each closed PR that lists the resolved issues. Depending on your SCM notification settings, you may also receive email or webhook notifications from your SCM provider.
