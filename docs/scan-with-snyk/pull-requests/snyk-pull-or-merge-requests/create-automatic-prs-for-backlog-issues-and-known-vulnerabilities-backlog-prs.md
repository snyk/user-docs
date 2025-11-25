# Create automatic PRs for backlog issues and known vulnerabilities (Backlog PRs)

{% hint style="info" %}
**Feature availability**

* The Automatic Fix PRs for known vulnerabilities (Backlog PRs) feature is supported for the following SCM integrations: GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab, and Azure Repos.
* The Automatic Fix PR settings may vary depending on the integration.
{% endhint %}

When Snyk creates automatic PRs for vulnerabilities, the following rules are applied:

* If you select **Retest now** for the Project, a scan is run manually, and the 24-hour window is marked as having had a scan run. No automatic PR is created until the next automated scan runs.
* One pull request is created per Project with a [priority score](../../../manage-risk/prioritize-issues-for-fixing/priority-score.md) of 700 and above.
* Fixing a vulnerability by upgrading a package may sometimes introduce a new vulnerability. Snyk will only automatically create such a pull request if the fixed vulnerabilities are a higher severity than any new ones introduced.
* Pull requests are created based on the **Test & Automated Pull Request Frequency** settings.
  * To update the **Test & Automated Pull Request Frequency**, navigate to **Projects** and select your Open Source Project.
  * Navigate to **Settings** and select an option from the pulldown list.

<figure><img src="../../../.gitbook/assets/project-settings-test-pull-request-frequency (1).png" alt="Project test and automated PR Checks frequency setting"><figcaption><p>Project test and automated PR Checks frequency setting</p></figcaption></figure>

To determine when your last 24-hour window began, check the Project issue card for **Snapshot taken by recurring test**.

<figure><img src="../../../.gitbook/assets/project-snapshot-taken.png" alt="Snapshot taken by recurrint test 13 hours ago"><figcaption><p>Snapshot taken by recurring test 13 hours ago</p></figcaption></figure>

For specific scan results, you can also check your inbox for an email titled **\[snyk] Vulnerability alert**.

## Configure Automatic Fix PRs at the integration level

Follow these steps to configure Automatic fix PRs on a specific Git repository you have already integrated with Snyk, such as GitHub.

{% hint style="warning" %}
Enabling **Automatic Fix PRs** can result in larger version jumps.
{% endhint %}

The configuration settings apply to all Projects in that Organization. You can also extend the configuration to Projects with custom settings.

1. Open Snyk Web UI and navigate to **Settings** > **Integrations**.
2. Select a Git repository integration (SCM). For this example, GitHub is configured.
3. Under **Automatic Fix PRs,** enable **Known vulnerabilities (backlog)**.\
   This retrieves previously declared vulnerabilities from the Project's backlog.

<figure><img src="../../../.gitbook/assets/project-settings-github-automatic-fix-prs.png" alt="Automatic fix PRs settings for Git integration."><figcaption><p>Automatic Fix PRs settings for Git integration</p></figcaption></figure>

4. Select the **Fix Strategy** for your Backlog PRs.

* By default, the fix strategy will be a single PR at the vulnerability level. Snyk opens one PR each day for issues in your backlog, fixing the top vulnerability it finds.
* You can check **Fix all vulnerabilities for the same dependency in a single PR**. This selects the vulnerability with the highest priority and a fix to resolve it, as well as fixes for other vulnerabilities in the same dependency.

5. Click **Save**.
6. (Optional) Select **Save changes and apply to all overridden Projects** to extend the current configuration to Projects with custom settings.\
   Use this option to apply the same configuration to all Projects. Selecting this option updates all of the individual Project settings for Automatic fix PRs. For Projects that previously had their own settings for automatic fix full requests, selecting this option overrides the Project setting with the global setting.

## Configure Automatic Fix PRs at the Project level

You can configure Automatic fix PRs to work only for specific Projects rather than having Projects inherit the settings from the global integration.

1. Navigate to **Projects** and expand the [target](../../../snyk-platform-administration/snyk-projects/#target) containing your Open Source Project.
2. Navigate to **Settings** and select an integration, for example, GitHub.
3. In the **Automatic fix pull requests** section:
   * Select **Customize for only this project**
   * Enable **Known vulnerabilities (backlog)**
4. Select the **Fix Strategy** for your Backlog PRs as described in the [Fix strategy step of configuring for integrations](create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md#configure-automatic-fix-prs-at-the-integration-level).
5. Click **Save changes**.

<figure><img src="../../../.gitbook/assets/project-settings-github-integration-automatic-fix-pull-requests.png" alt="Automatic fix PRs settings at the Project level."><figcaption><p>Automatic Fix PRs settings at the Project level</p></figcaption></figure>
