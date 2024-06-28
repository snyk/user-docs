# Snyk SCM Integrations

You can integrate Snyk with your Git repository to quickly and easily gain visibility across all the [Snyk Projects](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects) that you add to your **Projects** list.

Snyk Source Control Manager (SCM) integrations allow you to:

* Continuously perform security scanning across all integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes

Snyk can integrate with the following SCMs to help you track, monitor, and fix the issues and vulnerabilities in your code:

* [GitHub Cloud App](snyk-github-cloud-app.md)
* [GitHub Server App](snyk-github-server-app.md)
* [GitHub Enterprise](snyk-github-enterprise-integration.md)
* [GitHub](snyk-github-integration.md)
* [GitHub Read-only Projects](snyk-github-read-only-projects.md)
* [GitLab](snyk-gitlab-integration.md)
* [Bitbucket Cloud](snyk-bitbucket-cloud-integration.md)
* [Bitbucket Cloud App](snyk-bitbucket-cloud-app-integration.md)
* [Bitbucket Data Center/Server](snyk-bitbucket-data-center-server-integration.md)
* [Azure Repositories (TFS)](snyk-azure-repositories-tfs-integration.md)

## Deployment recommendations for SCM integrations

If you try to implement all the SCM integration features at the same time, you risk causing friction in your software development life cycle ([SDLC](https://snyk.io/learn/secure-sdlc/)), which in turn leads to a poor developer experience.

To ensure a smooth rollout of Snyk across your organization, Snyk provides a suggested deployment timeline consisting of deployment stages, configuration steps, and the desired outcome for each stage.

For detailed steps, see [Deployment recommendations for SCM integrations](./#deployment-recommendations-for-git-integrations).

## User permissions and access scope requirements

Snyk SCM integrations may require different permission requirements based on the connection method.

See the following for detailed permission requirements:

* [GitHub and GitHub Enterprise](./#github-and-github-enterprise-permissions-requirements)
* [GitHub Cloud App](./#github-cloud-app-permission-requirements)
* [GitHub Server App](./#github-server-app-permission-requirements)
* [GitLab](./#gitlab-permission-requirements)
* [Bitbucket](./#bitbucket-permission-requirements)
* [Azure Repositories (TFS)](./#azure-repositories-tfs-permission-requirements)

### GitHub and GitHub Enterprise permission requirements

{% hint style="info" %}
For information about token permissions in a brokered integration, see [GitHub - prerequisites and steps to install and configure Broker](https://docs.snyk.io/enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker).
{% endhint %}

The Snyk GitHub Enterprise integration is bound to a single user, preferably a GitHub service account. The level of access for the integration is defined by the combination of the user's permissions in GitHub and the access defined for the PAT on that user's account. If the PAT is defined with more permission than the user's GitHub account, the integration will not be able to use that permission.

The following table details the access scopes required in GitHub for Personal Access Tokens (PAT) and the scopes required for Snyk to perform the required operations on monitored repositories, such as reading manifest files on a frequent basis and opening fix or upgrade PRs. GitHub custom roles are not supported.

<table><thead><tr><th width="259">Action and purpose</th><th align="center">PAT scopes</th><th align="center">Repository scopes</th></tr></thead><tbody><tr><td><strong>Daily/weekly tests:</strong><br>Read manifest files in private repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>read</code></td></tr><tr><td><strong>Manual fix pull requests:</strong><br>Create fix PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center"></td></tr><tr><td><strong>Automatic fix and upgrade pull requests:</strong><br>Create fix or upgrade PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Snyk tests on pull requests:</strong><br>Send PR status checks whenever a new PR is created, or an existing PR is updated.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Initial configuration of Snyk tests on pull requests:</strong><br>Used to add SCM webhooks to the imported repo</td><td align="center"><code>admin:repo_hooks (read &#x26; write)</code></td><td align="center"><code>admin</code></td></tr><tr><td><strong>Import new Projects to Snyk:</strong><br>Present a list of all the available repos in the GitHub org in the <strong>Add Projects</strong> screen.</td><td align="center"><code>admin:read:org</code><br><code>repo (all)</code></td><td align="center"></td></tr></tbody></table>

A fine-grained PAT requires additional repository access scopes:

* `Administration: Read-only`
* `Commit Status: Read and write`
* `Content: Read and write`
* `Metadata: Read-only`
* `Pull requests: Read and write`
* `Webhooks: Read and write`
* `Members access: Read-only (Organization access scope)`

Snyk uses PRs to tell [GitHub Enterprise](snyk-github-enterprise-integration.md) that a merge is to occur. To do this, change content is pushed into a branch, which requires the `content: write` scope. A separate call is then made to create the fix PR, which requires the `pull request: write` scope. GitHub Enterprise is then instructed to create a PR, merging the change branch into the default branch.

Snyk uses SCM webhooks to:

* Track the state of Snyk pull requests when PRs are created, updated triggered, merged, and so on.
* Send push events to trigger PR checks.

### GitHub Cloud App permission requirements

The [Snyk GitHub Cloud App](snyk-github-cloud-app.md) integration uses role-based access control, meaning access control is not dependent on individual users or their role, it is instead tied to the app entity.

To set up the GitHub Cloud app integration you must be a:

* Snyk Organization Admin.
* GitHub Organization Admin.
* GitHub Repository Admin (if installing through the GitHub UI).

### GitHub Server App permission requirements

{% hint style="warning" %}
To utilize the Snyk GitHub Server App you must be using a self-hosted instance of GitHub.
{% endhint %}

The [Snyk GitHub Server App](snyk-github-server-app.md) uses role-based access control, meaning access control is not dependent on individual users or their role, it is instead tied to the app entity.

To set up the GitHub Cloud app integration you must be a:

* Snyk Organization Admin.
* GitHub Organization Admin.
* GitHub Repository Admin (if installing through the GitHub UI).

### GitLab permission requirements

The [Snyk GitLab integration](snyk-gitlab-integration.md#gitlab-access-tokens) uses either a personal access token (PAT) or group access token (GAT), depending on the GitLab account tier you are on.

To set up the Snyk GitLab integration you must be a:

* Snyk Group or Organization Admin.
* GitLab Owner or Maintainer

A PAT is used for managing personal GitLab projects and requires the `api` scope.

A GAT is used for managing multiple GitLab projects in a GitLab group and requires the `api` scope and maintainer role selected from the dropdown. You must be a GitLab Premium or Ultimate account tier holder to create a GAT.

### Bitbucket permission requirements

The Snyk Bitbucket integrations use different access control mechanisms to connect with Snyk:

* [Snyk Bitbucket Cloud](./#bitbucket-cloud-and-bitbucket-data-center-server-scopes) requires the creation of an [app password](snyk-bitbucket-cloud-integration.md#how-to-set-up-the-snyk-bitbucket-cloud-integration).
* [Snyk Bitbucket Cloud App](./#bitbucket-cloud-app-scopes) requires [Bitbucket workspace authorization](snyk-bitbucket-cloud-app-integration.md#setting-up-a-bitbucket-cloud-app) and related permissions.
* [Snyk Bitbucket Data Center/Server](./#bitbucket-cloud-and-bitbucket-data-center-server-scopes) requires a [dedicated username and password](snyk-bitbucket-data-center-server-integration.md#how-to-set-up-a-bitbucket-dc-server-integration) or a personal access token (PAT).

{% hint style="warning" %}
To set up any Snyk Bitbucket integration, you must be a Bitbucket Workspace Admin.
{% endhint %}

#### Bitbucket Cloud and Bitbucket Data Center/Server scopes

The following table details the required permission scopes in Bitbucket Cloud and Bitbucket Data Center/Server**:**

| Action and purpose                                                                                                                                      |                                                                       App password requirements                                                                      | Bitbucket permissions |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------: | :-------------------: |
| <p><strong>Daily / weekly tests:</strong><br>Read manifest files in private repos.</p>                                                                  |                                                                       **Repositories:** `read`                                                                       |       ≥ `write`       |
| <p><strong>Manual fix pull requests (triggered by the user):</strong><br>Create fix PRs in repos.</p>                                                   |                      <p><strong>Repositories</strong>: <code>read, write</code><br><strong>Pull Requests</strong>: <code>read, write</code></p>                      |                       |
| <p><strong>Automatic fix and upgrade pull requests:</strong><br>Create fix/upgrade PRs in repos.</p>                                                    |                      <p><strong>Repositories</strong>: <code>read, write</code><br><strong>Pull requests</strong>: <code>read, write</code></p>                      |       ≥ `write`       |
| <p><strong>Snyk tests on pull requests:</strong><br>Send PR status checks when a new PR is created or a PR is updated.</p>                              |                      <p><strong>Repositories</strong>: <code>read, write</code><br><strong>Pull requests</strong>: <code>read, write</code></p>                      |       ≥ `write`       |
| <p><strong>Snyk tests on pull requests (initial configuration):</strong><br>Add SCM webhooks to imported repos.</p>                                     |                                                                      **Webhooks**: `read, write`                                                                     |        `admin`        |
| <p><strong>Importing new projects to Snyk:</strong><br>Lists available repos in the Bitbucket instance in the <strong>Add Projects</strong> screen.</p> | <p><strong>Account</strong>: <code>read</code></p><p><strong>Workspace membership</strong>: <code>read</code></p><p><strong>Projects</strong>: <code>read</code></p> |                       |

Snyk uses SCM webhooks in Bitbucket to:

* Track the state of pull requests when PRs are created, updated triggered, merged, and so on.
* Send push events to trigger PR checks.

#### Bitbucket Cloud App scopes

The following table details the permissions required for the **Bitbucket Cloud App**:

| Action                                              | Purpose                                                                                                                                       | Required Scope                                            |
| --------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| Daily / weekly tests                                | Used to read manifest files in private repositories.                                                                                          | Read and modify your repositories and their pull requests |
| Snyk tests on pull requests                         | Used to send pull request status checks when a new PR is created, or an existing PR is updated.                                               | Read and modify your repositories and their pull requests |
| Opening fix and upgrade pull requests               | Used to create fix PRs in monitored repositories.                                                                                             | Read and modify your repositories and their pull requests |
| Snyk tests on pull requests - initial configuration | Used to add Snyk webhooks to the imported repos, to notify Snyk when pull requests are created or updated, and enable Snyk to trigger a scan. | Read and modify your repositories' webhooks               |

### Azure Repositories (TFS) permission requirements

The [Snyk Azure Repositories (TFS) integration](snyk-azure-repositories-tfs-integration.md) uses an Azure DevOps personal access token (PAT). This token is configured with the specific permissions Snyk needs to access your Azure repositories.

To set up the Snyk Azure Repositories (TFS) integration you must be:

* A [Snyk Organization Admin](../../snyk-admin/manage-permissions-and-roles/pre-defined-roles.md).
* A member of the [Project Administrators group](https://learn.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops\&tabs=preview-page) in Azure. This ensures the PAT has the `edit subscriptions permissions` required to enable webhooks.

In Azure, the PAT requires the following permissions for Snyk access:

* **Expiry**: Custom defined. Snyk recommends choosing a token expiration date that is far in the future.&#x20;
* **Scopes**: Custom defined. `Read & write` permissions are needed for the **Code** scope.
