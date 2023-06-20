# Automated pull request creation for new fixes

The following rules are applied to automatic PR creation for vulnerabilities:

* Pull requests are created based on the **Test & Automated Pull Request Frequency** (see screenshot below) setting
* If a scan is manually run (you clicked **Retest now** for the project), the 24-hour window is marked as having been run and no automatic PR is created until the next automated scan runs
* One pull request is created per project
* If **either** the vulnerability itself is new and has a fix available **or** if the fix is new and is not ignored
* For known vulnerabilities see [Automated pull request for known vulnerabilities (backlog)](fix-pull-requests-for-known-vulnerabilities-backlog.md)

![](../../../.gitbook/assets/os1.png)

To know when your last 24-hour window was kicked off, check the project page for **Snapshot taken by recurring test**--also check your email for **\[snyk] Vulnerability alert** for specific scan results:

![](../../../.gitbook/assets/os2.png)

Pull requests for new vulnerabilities are enabled by default for new integrations.

See the [Git repository SCM integrations](../../../integrations/git-repository-scm-integrations/) pages for details of supported integrations.

## Enable or disable pull requests for an integration

Enable at the global integration level:

1. Navigate to settings ![](../../../.gitbook/assets/cog\_icon.png) > **Integrations**.
2. Select an SCM integration (for example, GitHub).
3. Enable **New vulnerabilities**

**Apply changes to all overridden projects** will update all of the individual project settings for "Automatic fix pull requests". If a project previously had its own settings for this, clicking on this button will override it with the global setting.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-03 at 14.49.59.png" alt="Configure Automatic fix PRs"><figcaption><p>Configure Automatic fix PRs</p></figcaption></figure>

## Enable or disable pull requests for a single project

Enabling/disabling at a project level will override this single project rather than inheriting it from the global integration setting.

1. Under **Projects** select a project and select **Settings** (top right-hand corner).
2. Select **GitHub integration**.
3. Under the **Automatic fix pull requests** section:
   * Select **Customize for only this project**
   * Enable **New vulnerabilities**
   * Select **Save changes**

<figure><img src="../../../.gitbook/assets/os3.png" alt="Automatic fix pull requests settings at the project level"><figcaption><p>Automatic fix pull requests settings at the project level</p></figcaption></figure>
