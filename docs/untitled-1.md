# Fix pull requests for known vulnerabilities \(backlog\)

* [ Fixing vulnerabilities](/hc/en-us/articles/360011484018-Fixing-vulnerabilities)
* [ Fix pull requests for new vulnerabilities](/hc/en-us/articles/360017186498-Fix-pull-requests-for-new-vulnerabilities)
* [ Fix pull requests for known vulnerabilities \(backlog\)](/hc/en-us/articles/360017186958-Fix-pull-requests-for-known-vulnerabilities-backlog-)
* [ Integrate Snyk into your workflow](/hc/en-us/articles/360011377157-Integrate-Snyk-into-your-workflow)

##  Fix pull requests for known vulnerabilities \(backlog\)

K**nown vulnerabilities** retrieve vulnerabilities from the project's backlog. These are the previously declared vulnerabilities. 

The following rules are applied to automatic PR creation for vulnerabilities:

* Pull requests are created based on the **Test & Automated Pull Request Frequency** \(see screenshot below\) setting
* If a scan is manually run \(you clicked **Retest now** for the project\), the 24-hour window is marked as having been run and no automatic PR is created until the next automated scan runs
* One pull request is created per project \(priority score of **700 and above only**\)

To know when your last 24-hour window was kicked off, check the project page for **Snapshot taken by recurring test**--also check your email for **\[snyk\] Vulnerability alert** for specific scan results:

### Enable or disable pull requests for integrations

To enable at the global integration level:

1. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations**.
2. Select an SCM integration \(for example, GitHub\).
3. Enable **Known vulnerabilities \(backlog\)**

**Apply changes to all overridden projects** will update all of the individual project settings for "Automatic fix pull requests". If a project previously had its own settings for this, clicking on this button will override it with the global setting.

### Enable or disable pull requests for a single project

Enabling/disabling at a project level will override this single project rather than inheriting it from the global integration setting.

1. Navigate to the project and select **Settings**
2. Select **GitHub integration**
3. Under the **Automatic fix pull requests** section:
   * Select to **Customize for only this project**
   * Enable **Known vulnerabilities \(backlog\)**

