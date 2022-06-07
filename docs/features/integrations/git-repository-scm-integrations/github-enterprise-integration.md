# GitHub Enterprise integration

{% hint style="info" %}
Self-managed source code, like GitHub Enterprise, is available with our Enterprise plan, [our plans and pricing page](https://snyk.io/plans) for more info.
{% endhint %}

Snyk's GitHub Enterprise integration allows you to:

* Continuously perform security scanning across all the integrated repositories.
* Detect vulnerabilities in your open source components.
* Provide automated fixes and upgrades.

## Setting up a GitHub Enterprise Integration

1. Create a dedicated service account in GitHub Enterprise, with _**write**_ level or above permissions, to the repos you want to monitor with Snyk permissions. See [Required permissions scope for the GitHub integration](github-enterprise-integration.md#required-permissions-scope-for-the-github-integration) for details.
2. Generate a personal access token for that account, with **repo (all)**, **admin:read:org**, and **admin:repo\_hooks (read & write)** permissions scope. See [GitHub Enterprise documentation ](https://docs.github.com/en/enterprise-server@2.22/github/authenticating-to-github/creating-a-personal-access-token)for details.
3. **Authorize** your personal access token and Enable SSO:
4. Go to the **Integrations** page in Snyk and click on **GitHub Enterprise**:
5. Enter your Github Enterprise URL, and the personal access token for the service account you created: **Note**: You can use this integration to integrate to your GitHub Enterprise Cloud, by providing the following URL [https://api.github.com](https://api.github.com)
6. Click **Save**. Snyk connects to your GitHub Enterprise instance. When the connection succeeds, the following indications appear:
7. Select the repos to import to Snyk, then click **Add selected repositories**.
8. Snyk starts scanning the selected repos for dependency files (such as package.json) in the entire directory tree and imports them to Snyk as projects:
9. The imported projects appear on your **Projects** page and are continuously checked for vulnerabilities.

![](<../../../.gitbook/assets/which\_repos (3) (5) (9) (7) (18) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (27) (1) (1) (1) (1) (1) (30).jpg>)

## GitHub Enterprise Broker startup script

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e GITHUB=your.ghe.domain.com \
           -e GITHUB_API=your.ghe.domain.com/api/v3 \
           -e GITHUB_GRAPHQL=your.ghe.domain.com/api \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
       snyk/broker:github-enterprise
```

## GitHub Enterprise Integration Features

After the integration is set up, you can use the following capabilities:

**Project-level security reports**

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories and fix them by opening a fix pull request directly to your repository, with the required upgrades or patches.

This is an example of a project-level security report:

![](<../../../.gitbook/assets/mceclip0-22- (2) (5) (6) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (37).png>)

**Project monitoring and automatic fix pull requests**

Snyk frequently scans your projects on either a daily or a weekly basis. When new vulnerabilities are found, it notifies you by email and by opening an automated pull request with fixes to repositories.

This is an example of a fix pull request opened by Snyk:

![](../../../.gitbook/assets/uuid-6cfdaf0b-c349-468d-fe65-4f80bad110ea-en.png)

To review and adjust the automatic fix pull request settings:

Click on settings![cog\_icon.png](../../../.gitbook/assets/cog\_icon.png) > **Integrations**. 2. Select **Edit Settings** for GitHub Enterprise. 3. Navigate to **Automatic fix pull requests**:

![](<../../../.gitbook/assets/mceclip4 (1) (2) (6) (7) (3) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (38).png>)

**Pull request testing**

Snyk tests any newly created pull requests in your repositories for security vulnerabilities and sends a status check to GitHub Enterprise. This allows you to see whether the pull request introduces new security issues, directly from GitHub Enterprise.

This is how Snyk pull request checks appear in the Pull Request page in GitHub Enterprise:

![](<../../../.gitbook/assets/uuid-87113833-be79-dbe2-8860-a3f224d654c4-en (2) (2) (6) (5) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (23).png>)

To review and adjust the pull request tests settings:

1. Click on settings![cog\_icon.png](../../../.gitbook/assets/cog\_icon.png) > **Integrations**.
2. Select **Edit Settings** for GitHub Enterprise.
3. Navigate to **Default Snyk test for pull requests**:

![](<../../../.gitbook/assets/mceclip5 (1) (1).png>)

## Required permissions scope for the GitHub integration

All the operations, triggered manually or automatically, are performed for a GitHub service account that has its token is configured in the integrations settings. This shows the required access scopes for the configured token:

| **Action**                                          | **Why?**                                                                                                                                              | **Required permissions in GitHub** |
| --------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| Daily / weekly tests                                | For reading manifest files in private repos                                                                                                           | _repo (all)_                       |
| Manual fix pull requests (triggered by the user)    | For creating fix PRs in the monitored repos                                                                                                           | _repo (all)_                       |
| Automatic fix and upgrade pull requests             | For creating fix/upgrade PRs in the monitored repos                                                                                                   | _repo (all)_                       |
| Snyk tests on pull requests                         | For sending pull request status checks whenever a new PR is created / an existing PR is updated                                                       | _repo (all)_                       |
| Importing new projects to Snyk                      | For presenting a list of all the available repos in the GitHub org in the "Add Projects" screen (import popup)                                        | _admin:read:org, repo (all)_       |
| Snyk tests on pull requests - initial configuration | For adding Snyk's webhooks to the imported repos, so Snyk will be informed whenever pull requests are created or updated and be able to trigger scans | _admin:repo\_hooks (read & write)_ |

**Required permissions scope for repositories**

For Snyk to perform the required operation on monitor repositories, such as reading manifest files on a frequent basis, the accounts connected to Snyk (either directly or using Snyk Broker) need the following access on the repositories:

| **Action**                                          | **Why?**                                                                                                                                              | **Required permissions on the repository** |
| --------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| Daily / weekly tests                                | For reading manifest files in private repos                                                                                                           | _Write_ or above                           |
| Snyk tests on pull requests                         | For sending pull request status checks whenever a new PR is created / an existing PR is updated                                                       | _Write_ or above                           |
| Opening fix and upgrade pull requests               | For creating fix/upgrade PRs in the monitored repos                                                                                                   | _Write_ or above                           |
| Snyk tests on pull requests - initial configuration | For adding Snyk's webhooks to the imported repos, so Snyk will be informed whenever pull requests are created or updated and be able to trigger scans | _Admin_                                    |
