# GitHub Enterprise Integration

**Submit a Support ticket Sign in to Support Sign up for free**

#### [ ](untitled-259.md) <a id="category-name"></a>

**Git repository \(SCM\) integrations**

* [ GitHub integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032117-GitHub-integration/README.md)
* [ GitHub Enterprise Integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360015951318-GitHub-Enterprise-Integration/README.md)
* [ Bitbucket Cloud integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032097-Bitbucket-Cloud-integration/README.md)
* [ Bitbucket Data Center / Server integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002218-Bitbucket-Data-Center-Server-integration/README.md)
* [ GitLab integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002238-GitLab-integration/README.md)
* [ Azure Repos integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002198-Azure-Repos-integration/README.md)
* [ GitHub Read-Only Projects](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360010766557-GitHub-Read-Only-Projects/README.md)
* [ Opening fix and upgrade pull requests from a fixed GitHub account](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360010843797-Opening-fix-and-upgrade-pull-requests-from-a-fixed-GitHub-account-/README.md)
* [ Test your PRs for vulnerabilities before merging](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360006528057-Test-your-PRs-for-vulnerabilities-before-merging/README.md)
* [ Snyk checks on pull requests](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360006581938-Snyk-checks-on-pull-requests/README.md)

  [See more](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/sections/360001138098-Git-repository-SCM-integrations/README.md)

## GitHub Enterprise Integration

Snyk's GitHub Enterprise integration allows you to:

* Continuously perform security scanning across all the integrated repositories.
* Detect vulnerabilities in your open source components.
* Provide automated remediation and upgrade fixes.

**GHE compatibility and availability**  
On-premise GHE integrations are not currently supported. However, some on-premise SCM integrations, like Github, are available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.

### Setting up a GitHub Enterprise Integration

1. Create a dedicated service account in GitHub Enterprise, with _**write**_ level or above permissions, to the repos you want to monitor with Snyk permissions. See [Required permissions scope for the GitHub integration](untitled-259.md) for details.
2. Generate a personal access token for that account, with **repo \(all\)**, **admin:read:org**, and **admin:repo\_hooks \(read & write\)** permissions scope. See [GitHub Enterprise documentation ](https://docs.github.com/en/enterprise-server@2.22/github/authenticating-to-github/creating-a-personal-access-token)for details.
3. **Authorize** your personal access token and Enable SSO: 
4. Go to the **Integrations** page in Snyk and click on **GitHub Enterprise**: 
5. Enter your Github Enterprise URL, and the personal access token for the service account you created: **Note:** You can use this integration to integrate to your GitHub Enterprise Cloud, by providing the following URL [https://api.github.com](https://api.github.com). 
6. Click **Save**. Snyk connects to your GitHub Enterprise instance. When the connection succeeds, the following indications appear: 
7. Select the repos to import to Snyk, then click **Add selected repositories**.
8. Snyk starts scanning the selected repos for dependency files \(such as package.json\) in the entire directory tree and imports them to Snyk as projects:
9. The imported projects appear on your **Projects** page and are continuously checked for vulnerabilities.

### GitHub Enterprise Integration Features

After the integration is set up, you can use the following capabilities:

**Project-level security reports**

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories and fix them by opening a fix pull request directly to your repository, with the required upgrades or patches.

This is an example of a project-level security report:

**Projects monitoring and automatic fix pull requests**

Snyk frequently scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, it notifies you by email and by opening an automated pull request with fixes to repositories.

This is an example of a fix pull request opened by Snyk:

To review and adjust the automatic fix pull request settings:

1. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations**.
2. Select **Edit Settings** for GitHub Enterprise.
3. Navigate to **Automatic fix pull requests**: 

**Pull request testing**

Snyk tests any newly created pull request in your repositories for security vulnerabilities and sends a status check to GitHub Enterprise. This allows you to see whether the pull request introduces new security issues, directly from GitHub Enterprise.

This is how Snyk pull request checks appear in the Pull Request page in GitHub Enterprise:

To review and adjust the pull request tests settings:

1. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations**.
2. Select **Edit Settings** for GitHub Enterprise.
3. Navigate to **Default Snyk test for pull requests**: 

### Required permissions scope for the GitHub integration <a id="h_01ER1W3EZ4DXGHGKT12DWQEJV2"></a>

All the operations, triggered manually or automatically, are performed for a GitHub service account that has its token is configured in the integrations settings. This shows the required access scopes for the configured token:

| **Action** | **Why?** | **Required permissions in GitHub** |
| :--- | :--- | :--- |
| Daily / weekly tests | For reading manifest files in private repos | _repo \(all\)_ |
| Manual fix pull requests \(triggered by the user\) | For creating fix PRs in the monitored repos | _repo \(all\)_ |
| Automatic fix and upgrade pull requests | For creating fix/upgrade PRs in the monitored repos | _repo \(all\)_ |
| Snyk tests on pull requests | For sending pull request status checks whenever a new PR is created / an existing PR is updated | _repo \(all\)_ |
| Importing new projects to Snyk | For presenting a list of all the available repos in the GitHub org in the "Add Projects" screen \(import popup\) | _admin:read:org, repo \(all\)_ |
| Snyk tests on pull requests - initial configuration | For adding Snyk's webhooks to the imported repos, so Snyk will be informed whenever pull requests are created or updated and be able to trigger scans | _admin:repo\_hooks \(read & write\)_ |

**Required permissions scope for repositories**

For Snyk to perform the required operation on monitor repositories, such as reading manifest files on a frequent basis, the accounts connected to Snyk \(either directly or using Snyk Broker\) need the following access on the repositories:

| **Action** | **Why?** | **Required permissions on the repository** |
| :--- | :--- | :--- |
| Daily / weekly tests | For reading manifest files in private repos | _Write_ or above |
| Snyk tests on pull requests | For sending pull request status checks whenever a new PR is created / an existing PR is updated |  |
| Opening fix and upgrade pull requests | For creating fix/upgrade PRs in the monitored repos |  |
| Snyk tests on pull requests - initial configuration | For adding Snyk's webhooks to the imported repos, so Snyk will be informed whenever pull requests are created or updated and be able to trigger scans | _Admin_ |

