# Snyk GitHub Enterprise integration

{% hint style="info" %}
**Feature availability**\
GitHub Enterprise integration is available to Snyk Enterprise plan customers. If you have a Legacy Business plan, contact [Snyk support](https://support.snyk.io/hc/en-us) for access. See the  [Plans and pricing](https://snyk.io/plans/) page for details.
{% endhint %}

{% hint style="info" %}
If you are a Snyk Enterprise plan customer, Snyk recommends that you use the GitHub Enterprise integration. If you use the self-hosted GitHub Enterprise product, you must use the Snyk GitHub Enterprise integration. See [Using GItHub or GItHub Enterprise integration](using-github-or-github-enterprise-integration.md) for details.
{% endhint %}

Snyk GitHub Enterprise integration lets you:

* Continuously perform security scanning across all the integrated repositories.
* Detect vulnerabilities in your Open Source components.
* Provide automated fixes and upgrades.

## How to set up a Snyk GitHub Enterprise integration

{% hint style="info" %}
If your repositories are not internet-accessible, you must use [Snyk Broker](../../enterprise-setup/snyk-broker/). This requires creating a startup script. For the script and instructions, see [GitHub Enterprise - install and configure using Docker](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/setup-broker-with-github-enterprise.md).
{% endhint %}

Follow these steps to connect Snyk with your GitHub repositories:

1. If you have not already done so, create a dedicated GitHub service account with `write-level` or higher permissions for the repositories to monitor.\
   For more information, see [Types of GitHub accounts](https://docs.github.com/en/get-started/learning-about-github/types-of-github-accounts). For information about the necessary permissions, see [Required permissions scope for Snyk GitHub Enterprise integration](snyk-github-enterprise-integration.md#required-permissions-scope-for-the-github-integration).
2. Generate a personal access token for the GitHub service account with these permissions:
   * `repo (all)`
   * `admin:read:org`
   * `admin:repo_hooks` (read and write)
3. If you are using fine-grained personal access tokens, the following scope is required. For details, see the [GitHub Enterprise documentation](https://docs.github.com/en/enterprise-server@2.22/github/authenticating-to-github/creating-a-personal-access-token).
   * Repository permissions
     * Administration: `Read-only`&#x20;
     * Commit Status: `Read and write`&#x20;
     * Contents: `Read and write`&#x20;
     * Metadata: `Read-only`&#x20;
     * Pull requests: `Read and write`&#x20;
     * Webhooks: `Read and write`
   * Organization permissions - Members: `Read-only`
4. Authorize your personal access token and enable SSO:
   1. In Snyk, navigate to the **Integrations** page and click the **GitHub Enterprise** tile.
   2. Enter your GitHub URL and the personal access token (PAT) for the GitHub service account you created, and **Save** your changes.\
      Snyk connects to your GitHub instance. When the connection succeeds, the list of available repositories is displayed.\
      **Note**: To use this integration with your GitHub Enterprise Cloud instance, provide the following URL: [https://api.github.com](https://api.github.com).
   3. If your Snyk Organization enforces SAML/SSO, select **Configure SSO** next to the PAT for the GitHub service account.\
      **Note:** Occasionally, SSO is enforced in your Organizations after a PAT and Integration are configured. If this happens, any Projects that have already been imported are displayed in Snyk but retests, PR Checks, and so on will not be performed.\
      If this happens, check the **Configure SSO** settings to ensure that the Snyk Organization is authorized for the GitHub Enterprise integration.\
      On occasion, an Organization shows as authorized, but the retests and PR checks do not work. If this happens, de-authorizing the Organization and then re-authorizing it may help.
5. Select the repositories to import to Snyk and click **Add selected repositories**.

Snyk starts scanning the selected repositories for dependency files, such as `package.json`, in the entire directory tree, and imports the repositories to Snyk as Projects.

The imported Projects appear on your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/github_integration-fix_15dec2022 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (16) (26).jpeg" alt="Imported Projects on the Projects page in Snyk"><figcaption><p>Imported Projects on the Projects page in Snyk</p></figcaption></figure>

## Uses of Snyk GitHub Enterprise integration

After the integration is set up, you can use the capabilities explained in the sections that follow.

### **Obtain Project-level security reports**

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories and fix them by opening a fix pull request directly to your repository with the required upgrades or patches.

The example that follows shows a Project-level security report.

<figure><img src="../../.gitbook/assets/project_lvl_security_rpt-18july2022.png" alt="Project-level security report"><figcaption><p>Project-level security report</p></figcaption></figure>

### **Monitor Projects and generate automatic fix pull requests**

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and opens an automated pull request with fixes for your repositories.

The example that follows shows a fix pull request opened by Snyk.

<figure><img src="../../.gitbook/assets/github_fix_pr_cropped-14july2022 (1) (1).png" alt="Fix pull request created by Snyk"><figcaption><p>Fix pull request created by Snyk</p></figcaption></figure>

To review and update the automatic fix pull request settings:

1. In Snyk, navigate to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> **Settings >** **Integrations > Source control > GitHub Enterprise**, and select **Edit Settings**.
2. Scroll to the **Automatic fix pull requests** section, then select options as required:

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.41.56.png" alt="Automatic pull requests settings"><figcaption><p>Automatic pull requests settings</p></figcaption></figure>

### **Test new pull requests**

The [PR Checks](../../scan-applications/run-pr-checks/) feature enables Snyk to test any newly-created pull requests in your repositories for security vulnerabilities and sends a status check to GitHub. This allows you to see, directly from GitHub, whether the pull request introduces new security issues.

The following example shows how Snyk pull request checks appear on the pull requests page in GitHub.

<figure><img src="../../.gitbook/assets/pr_testing-14july2022.png" alt="Pull request checks shown in GitHub Enterprise"><figcaption><p>Pull request checks shown in GitHub Enterprise</p></figcaption></figure>

To review and adjust the pull request tests settings: In Snyk, navigate to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> Organization **Settings** > **Integrations > Source control > GitHub Enterprise**, and select **Edit Settings**.

1. Scroll to **Snyk PR status checks**; see [Configure PR Checks](../../scan-applications/run-pr-checks/configure-pr-checks.md) for details.

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.43.34.png" alt="Default Snyk test for pull requests setting enabled"><figcaption><p>Default Snyk test for pull requests setting enabled</p></figcaption></figure>

## Required permissions scope for Snyk GitHub Enterprise integration

All the operations, whether triggered manually or automatically, are performed for a GitHub service account that has its token configured on the integrations settings page. This shows the required access scopes for the configured token:

| **Action**                                              | **Purpose**                                                                                                                                                                                                                                                     | **Required permissions in GitHub** |
| ------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| Daily / weekly tests                                    | Used to read manifest files in private repos.                                                                                                                                                                                                                   | `repo (all)`                       |
| Manual fix pull requests (triggered by the user)        | Used to create fix PRs in the monitored repos.                                                                                                                                                                                                                  | `repo (all)`                       |
| Automatic fix and upgrade pull requests                 | Used to create fix or upgrade PRs in the monitored repos.                                                                                                                                                                                                       | `repo (all)`                       |
| Snyk tests on pull requests                             | Used to send pull request status checks whenever a new PR is created or an existing PR is updated.                                                                                                                                                              | `repo (all)`                       |
| Importing new Projects to Snyk                          | Used to present a list of all the available repos in the GitHub org in the **Add Projects** screen (import popup).                                                                                                                                              | `admin:read:org, repo (all)`       |
| Snyk tests on pull requests : **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests, that is, when PRs are created, updated triggered, merged, and so on.</li><li>Send push events to trigger PR checks.</li></ul> | `admin:repo_hooks (read & write)`  |
