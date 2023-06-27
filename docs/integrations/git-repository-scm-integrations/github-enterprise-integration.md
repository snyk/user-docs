# GitHub Enterprise integration

{% hint style="info" %}
**Feature availability**\
Scanning of self-managed source code, like GitHub Enterprise, is available with the Enterprise plan. See[ the Snyk plans and pricing page](https://snyk.io/plans) for more information.
{% endhint %}

Snyk GitHub Enterprise integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes and upgrades

{% hint style="info" %}
Snyk recommends GitHub Enterprise integration for Enterprise customers, as it allows the use of a single Organization-wide Personal Access Token (PAT) rather than a user account-based PAT. See [Using GItHub or GItHub Enterprise integration](using-github-or-github-enterprise-integration.md) for details.
{% endhint %}

## Setting up a GitHub Enterprise integration

The process to connect Snyk with your GitHub Enterprise repositories includes the following steps. Note that using a service account is recommended but not required.

1. Create a dedicated service account in GitHub Enterprise with _**write**_ level or higher permissions for the repos you want to monitor with Snyk permissions.\
   See [Types of GitHub accounts](https://docs.github.com/en/get-started/learning-about-github/types-of-github-accounts) and [Required permissions scope for the GitHub integration](github-enterprise-integration.md#required-permissions-scope-for-the-github-integration) for details.
2. Generate a personal access token for that account with the following permissions:\
   \- repo (all)\
   \- admin:read:org\
   \- admin:repo\_hooks (read & write)\
   \
   If you are using fine-grained personal access tokens, the following scope is required:\
   \
   Repository Permissions\
   \- Administration: Read-only \
   \- Commit Status: Read and write \
   \- Contents: Read and write \
   \- Metadata: Read-only \
   \- Pull requests: Read and write \
   \- Webhooks: Read and write \
   \
   Organization Permissions\
   \- Members: Read-only\
   \
   See the [GitHub Enterprise documentation](https://docs.github.com/en/enterprise-server@2.22/github/authenticating-to-github/creating-a-personal-access-token) for details.
3. **Authorize** your personal access token and Enable SSO:
   1. In Snyk, go to the **Integrations** page and click the **GitHub Enterprise** card.
   2. Enter your Github Enterprise URL and the personal access token (PAT) for the service account you created, and **Save** your changes.\
      Snyk connects to your GitHub Enterprise instance. When the connection succeeds, the list of available repositories is displayed.\
      **Note**: To use this integration to integrate with your GitHub Enterprise Cloud instance, provide the following URL: [https://api.github.com](https://api.github.com).
   3. If your GitHub Enterprise organization enforces SAML/SSO, select **Configure SSO** next to the PAT in GitHub after the PAT has been created.\
      **Note:** Occasionally, SSO is enforced in your GitHub Enterprise organizations after a PAT and Integration are configured. If this happens, any Projects that have already been imported show in Snyk but retests, PR Checks, and so on will not be performed.\
      If this happens, check the **Configure SSO** settings here to ensure that the GitHub Enterprise organization is Authorized.\
      On occasion, an organization shows as Authorized, but the retests and PR checks do not work. If this happens, de-authorizing the organization and then re-authorizing it may help.
4. Select the repositories you want to import to Snyk and click **Add selected repositories**.

Snyk starts scanning the selected repositories for dependency files, such as package.json, in the entire directory tree and imports them to Snyk as Projects.

The imported Projects appear on your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/github_integration-fix_15dec2022 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (16).jpeg" alt="Imported Projects on the Projects page"><figcaption><p>Imported Projects on the Projects page</p></figcaption></figure>

## GitHub Enterprise Broker startup script

For the script and instructions, see [GitHub Enterprise - install and configure using Docker](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/setup-broker-with-github-enterprise.md).

## GitHub Enterprise integration features

After the integration is set up, you can use the capabilities explained in the sections that follow.

### **Project-level security reports**

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories and fix them by opening a fix pull request directly to your repository with the required upgrades or patches.

The example that follows shows a Project-level security report.

<figure><img src="../../.gitbook/assets/project_lvl_security_rpt-18july2022.png" alt="Project-level security report"><figcaption><p>Project-level security report</p></figcaption></figure>

### **Project monitoring and automatic fix pull requests**

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and opens an automated pull request with fixes for your repositories.

The example that follows shows a fix pull request opened by Snyk.

<figure><img src="../../.gitbook/assets/github_fix_pr_cropped-14july2022 (1).png" alt="Fix pull request created by Snyk"><figcaption><p>Fix pull request created by Snyk</p></figcaption></figure>

To review and update the automatic fix pull request settings:

1. In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> **Settings >** **Integrations > Source control > GitHub Enterprise**, and select **Edit Settings**.
2. Scroll to the **Automatic fix pull requests** section, then select options as required:

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.41.56.png" alt="Automatic pull requests settings"><figcaption><p>Automatic pull requests settings</p></figcaption></figure>

### **Pull request testing**

The [PR Checks](../../scan-application-code/run-pr-checks/) feature enables Snyk to test any newly-created pull requests in your repositories for security vulnerabilities and sends a status check to GitHub Enterprise. This allows you to see, directly from GitHub Enterprise, whether the pull request introduces new security issues.

The example that follows shows how Snyk pull request checks appear on the GitHub Enterprise Pull Request page.

<figure><img src="../../.gitbook/assets/pr_testing-14july2022.png" alt="Pull request checks shown in GitHub Enterprise"><figcaption><p>Pull request checks shown in GitHub Enterprise</p></figcaption></figure>

To review and adjust the pull request tests settings: In Snyk, go to <img src="../../.gitbook/assets/cog_icon.png" alt="Settings" data-size="line"> Organization **Settings** > **Integrations > Source control > GitHub Enterprise**, and select **Edit Settings**.

1. Scroll to **Snyk PR status checks**; see [Configure PR Checks](../../scan-application-code/run-pr-checks/configure-pr-checks.md) for details.

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.43.34.png" alt="Default Snyk test for pull requests setting enabled"><figcaption><p>Default Snyk test for pull requests setting enabled</p></figcaption></figure>

## Required permissions scope for the GitHub integration

All the operations, whether triggered manually or automatically, are performed for a GitHub service account that has its token configured in the integrations settings page. This shows the required access scopes for the configured token:

| **Action**                                              | **Purpose**                                                                                                                                                                                                                                                     | **Required permissions in GitHub** |
| ------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| Daily / weekly tests                                    | Used to read manifest files in private repos.                                                                                                                                                                                                                   | _repo (all)_                       |
| Manual fix pull requests (triggered by the user)        | Used to create fix PRs in the monitored repos.                                                                                                                                                                                                                  | _repo (all)_                       |
| Automatic fix and upgrade pull requests                 | Used to create fix or upgrade PRs in the monitored repos.                                                                                                                                                                                                       | _repo (all)_                       |
| Snyk tests on pull requests                             | Used to send pull request status checks whenever a new PR is created or an existing PR is updated.                                                                                                                                                              | _repo (all)_                       |
| Importing new Projects to Snyk                          | Used to present a list of all the available repos in the GitHub org in the **Add Projects** screen (import popup).                                                                                                                                              | _admin:read:org, repo (all)_       |
| Snyk tests on pull requests : **initial configuration** | <p>Used to add SCM webhooks to the imported repos. Snyk uses these webhooks to:</p><ul><li>Track the state of Snyk pull requests, that is, when PRs are created, updated triggered, merged, and so on.</li><li>Send push events to trigger PR checks.</li></ul> | _admin:repo\_hooks (read & write)_ |

