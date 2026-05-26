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
