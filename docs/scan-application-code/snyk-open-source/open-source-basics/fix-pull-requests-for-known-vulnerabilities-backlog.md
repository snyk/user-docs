# Configure Automatic fix PRs

{% hint style="info" %}
This feature is supported in GitHub, GitHub Enterprise, and Bitbucket Cloud integrations.



The Autofix PR settings may vary depending on the integration.



The fix strategy feature for getting dependency-oriented fixes is in [beta](../../../more-info/snyk-feature-release-process.md). We'll be happy to get your feedback.
{% endhint %}

**Known vulnerabilities** retrieve vulnerabilities from the Project's backlog. These are the previously declared vulnerabilities.

The following rules are applied to automatic PR creation for vulnerabilities:

* If a scan is manually run (you clicked **Retest now** for the Project), the 24-hour window is marked as having been run and no automatic PR is created until the next automated scan runs.
* One pull request is created per Project ([priority score](../../../manage-issues/prioritizing-and-managing-issues/priority-score.md) of **700 and above only**).
* Pull requests are created based on the **Test & Automated Pull Request Frequency** settings. To update the **Test & Automated Pull Request Frequency**, go to **Projects**, select your open source and licensing Project, then go to **Settings** (see screenshot below).

<figure><img src="../../../.gitbook/assets/Project testing and PR Checks frequency.png" alt="Project testing and PR Checks frequency."><figcaption><p>Project testing and PR Checks frequency</p></figcaption></figure>

To know when your last 24-hour window was kicked off, check the Project page for **Snapshot taken by recurring test**.&#x20;

<figure><img src="../../../.gitbook/assets/Test information with a focus on the latest snapshot taken.png" alt="Test information with focus on the latest snapshot taken."><figcaption><p>Test information with a focus on the latest snapshot taken</p></figcaption></figure>

For specific scan results, you can also check your inbox for an email titled **\[snyk] Vulnerability alert**.

## Configure Automatic fix PR at the integration level

Configure Automatic fix PR on a specific Git repository you have already integrated with Snyk, such as GitHub.

The configuration settings apply to all Projects in that Organization. You can also extend the configuration to Projects with custom settings.

1. Open Snyk Web UI and go to **Settings** <img src="../../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> > **Integrations**.
2. Select a Git repository integration (SCM). For this example, GitHub is configured.
3. Under **Automatic fix PRs** enable **Known vulnerabilities (backlog)**.

<figure><img src="../../../.gitbook/assets/Automatic fix PRs settings for Git integration.png" alt="Automatic fix PRs settings for Git integration."><figcaption><p>Automatic fix PRs settings for Git integration</p></figcaption></figure>

4. Select the **Fix Strategy** for your Backlog PRs.

* By default, the fix strategy will be a single PR at the vulnerability level. Snyk opens a PR per day for issues in your backlog, fixing the top vulnerability it finds as described above.
* Another option is checking the **Fix all vulnerabilities for the same dependency in a single PR** checkbox. This will pick the vulnerability with the highest priority, and suggest such a bump to solve it, as well as other vulnerabilities in the same dependency.

5. **Save** changes.
6. (Optional) Select **Save changes and apply to all overridden Projects** to extend the current configuration to Projects with custom settings (see Configure [Automatic fix PR at the integration level](fix-pull-requests-for-known-vulnerabilities-backlog.md#configure-automatic-fix-pr-at-the-project-level)). Use this option to apply the same configuration to all Projects.

{% hint style="warning" %}
Enabling Automatic fix PRs can result in larger version jumps.
{% endhint %}

<figure><img src="../../../.gitbook/assets/Fix all vulnerabilities for the same dependency in a single PR.png" alt="Fix all vulnerabilities for the same dependency in a single PR."><figcaption><p>Fix all vulnerabilities for the same dependency in a single PR</p></figcaption></figure>

## Configure Automatic fix PR at the Project level

You can configure Automatic fix PR to work only for specific Projects rather than inheriting the settings from the global integration. In this example, GitHub integration is used.

1. Go to **Projects**, then expand the [target](../../../manage-issues/snyk-projects/#target) containing your open source Project.
2. Go to **Settings >** **GitHub integration.**
3. Under the **Automatic fix pull requests** section:
   * Select **Customize for only this project**
   * Enable **Known vulnerabilities (backlog)**
4. Select the **Fix Strategy** for your Backlog PRs as described in the [Fix strategy step](fix-pull-requests-for-known-vulnerabilities-backlog.md#configure-automatic-fix-pr-at-the-integration-level).
5. **Save changes**.

<figure><img src="../../../.gitbook/assets/Automatic fix PRs settings at the Project level.png" alt="Automatic fix PRs settings at the Project level."><figcaption><p>Automatic fix PRs settings at the Project level</p></figcaption></figure>

