# Automated pull requests for known vulnerabilities (backlog)

{% hint style="info" %}
This feature is supported in GitHub , GitHub Enterprise, and Bitbucket Cloud integrations.

The Autofix PR settings may vary depending on the integration.
{% endhint %}

**Known vulnerabilities** retrieve vulnerabilities from the Project's backlog. These are the previously declared vulnerabilities.

The following rules are applied to automatic PR creation for vulnerabilities:

* Pull requests are created based on the **Test & Automated Pull Request Frequency** (see screenshot below) setting
* If a scan is manually run (you clicked **Retest now** for the project), the 24-hour window is marked as having been run and no automatic PR is created until the next automated scan runs
* One pull request is created per Project (priority score of **700 and above only**)

<figure><img src="../../../.gitbook/assets/os1.png" alt="Project testing and PR Checks frequency."><figcaption><p>Project testing and PR Checks frequency</p></figcaption></figure>

To know when your last 24-hour window was kicked off, check the Project page for **Snapshot taken by recurring test**--also check your email for **\[snyk] Vulnerability alert** for specific scan results:

<figure><img src="../../../.gitbook/assets/os2.png" alt="Test information with focus on the latest snapshot taken."><figcaption><p>Test information with focus on the latest snapshot taken</p></figcaption></figure>

## Enable or disable pull requests for integrations

To enable at the global integration level:

1. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Integrations**
2. Select an SCM integration (for example, GitHub)
3. Enable **Known vulnerabilities (backlog)**

**Apply changes to all overridden projects** will update all of the individual Project settings for "Automatic fix pull requests". If a Project previously had its settings for this, clicking on this button will override it with the global setting.

<figure><img src="../../../.gitbook/assets/screen_shot_2021-05-24_at_12.23.38_pm.png" alt="Automatic fix pull request settings at the integration level."><figcaption><p>Automatic fix pull request settings at the integration level</p></figcaption></figure>

## Enable or disable pull requests for a single project

Enabling/disabling at a project level will override this single project rather than inheriting it from the global integration setting.

1. Navigate to the project and select **Settings**
2. Select **GitHub integration**
3. Under the **Automatic fix pull requests** section:
   * Select **Customize for only this project**
   * Enable **Known vulnerabilities (backlog)**

<figure><img src="../../../.gitbook/assets/os3.png" alt="Automatic fix pull request settings at the project level,"><figcaption><p>Automatic fix pull request settings at the project level</p></figcaption></figure>
