# User permissions and access scopes

Snyk SCM integrations may require different permission requirements based on the connection method.

See the following for detailed permission requirements:

* GitHub and GitHub Enterprise
* GitHub Cloud App
* GitHub Server App
* GitHub Server App for Universal Broker
* GitLab
* Bitbucket
* Azure Repositories (TFS)

### GitHub and GitHub Enterprise permissions requirements&#x20;

{% hint style="info" %}
For information about token permissions in a brokered integration, see [GitHub - prerequisites and steps to install and configure Broker](../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/) and [Integrated SCM tokens for Snyk Broker](scm-integrations-and-snyk-broker.md#integrated-scm-tokens-for-classic-broker).
{% endhint %}

The Snyk GitHub Enterprise integration is bound to a single user, preferably a GitHub service account. The level of access for the integration is defined by the combination of the user's permissions in GitHub and the access defined for the Personal Access Token (PAT) on that user's account. If the PAT is defined with more permission than the user's GitHub account, the integration will not be able to use that permission.

The following table details the access scopes required in GitHub and GitHub Enterprise for Personal Access Tokens (PAT) and the scopes required for Snyk to perform the required operations on monitored repositories, such as reading manifest files on a frequent basis and opening fix or upgrade PRs. GitHub custom roles are not supported.

<table><thead><tr><th width="259">Action and purpose</th><th align="center">PAT scopes</th><th align="center">Repository scopes</th></tr></thead><tbody><tr><td><strong>Daily/weekly tests:</strong><br>Read manifest files in private repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>read</code></td></tr><tr><td><strong>Manual fix pull requests:</strong><br>Create fix PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center"></td></tr><tr><td><strong>Automatic fix and upgrade pull requests:</strong><br>Create fix or upgrade PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Snyk tests on pull requests:</strong><br>Send PR status checks whenever a new PR is created, or an existing PR is updated.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Initial configuration of Snyk tests on pull requests:</strong><br>Used to add SCM webhooks to the imported repo</td><td align="center"><code>admin:repo_hooks (read &#x26; write)</code></td><td align="center"><code>admin</code></td></tr><tr><td><strong>Import new Projects to Snyk:</strong><br>Present a list of all the available repos in the GitHub org in the <strong>Add Projects</strong> screen.</td><td align="center"><code>admin:read:org</code><br><code>repo (all)</code></td><td align="center"></td></tr></tbody></table>

A fine-grained PAT requires additional repository access scopes:

* `Administration: Read-only`
* `Commit Status: Read and write`
* `Content: Read and write`
* `Metadata: Read-only`
* `Pull requests: Read and write`
* `Webhooks: Read and write`
* `Members access: Read-only (Organization access scope)`

The `Administration: Read-only` permission on the PAT is crucial for Snyk to identify and list the user's accessible GitHub organizations, a prerequisite for importing a new Project.

Snyk uses PRs to tell [GitHub Enterprise](organization-level-integrations/github-enterprise.md) that a merge is to occur. To do this, change content is pushed into a branch, which requires the `content: write` scope. A separate call is then made to create the fix PR, which requires the `pull request: write` scope. GitHub Enterprise is then instructed to create a PR, merging the change branch into the default branch.

Snyk uses SCM webhooks to:

* Track the state of Snyk pull requests when PRs are created, updated triggered, merged, and so on.
* Send push events to trigger PR checks.

### GitHub Cloud App permission requirements

The [Snyk GitHub Cloud App](organization-level-integrations/github-cloud-app.md) integration uses role-based access control, meaning access control is not dependent on individual users or their role, it is instead tied to the app entity.

To set up the GitHub Cloud app integration you must be a:

* Snyk Organization Admin.
* GitHub Organization Admin.
* GitHub Repository Admin (if installing through the GitHub UI).

{% hint style="info" %}
While some permissions may be optional from GitHub’s perspective, they are necessary to support Snyk functions. These permissions **cannot be customized** for your individual needs because the app is registered under the Snyk Organization.&#x20;
{% endhint %}

The following table states the required GitHub App permissions and scopes:

<table><thead><tr><th width="191">Action and scope</th><th width="254">Scope</th><th>Level</th><th>Permission</th></tr></thead><tbody><tr><td>Determine if the GitHub user has admin role on the GitHub org, to restrict app installation reuse to only admin users</td><td>Members</td><td>Organization</td><td>Read</td></tr><tr><td>Search repositories, and access repository metadata.</td><td>Metadata</td><td>Repository</td><td>Read</td></tr><tr><td>Interact with the GitHub Checks tab</td><td>Checks</td><td>Repository</td><td>Read and write</td></tr><tr><td>Create commits and branches</td><td>Contents</td><td>Repository</td><td>Read and write</td></tr><tr><td>Send PR check results as commit statuses</td><td>Commit status</td><td>Repository</td><td>Read and write</td></tr><tr><td>Get pull requests details, post related comments (next gen PR experience)</td><td>Pull request</td><td>Repository</td><td>Read and write</td></tr><tr><td>Manage webhooks which trigger the PR checks</td><td>Repository hooks</td><td>Repository</td><td>Read and write</td></tr></tbody></table>

### GitHub Server App permission requirements

To set up the GitHub Server App you must be a:

* Snyk Organization Admin.
* GitHub Organization Admin.

You must have a public or private GitHub repository.

### GitHub Server App for Universal Broker permission requirements

To set up the GitHub Server App for Universal Broker you must be a:

* Snyk Organization Admin.
* GitHub Organization Admin.

You must have a self-hosted instance of GitHub.

### GitLab permission requirements

The [Snyk GitLab integration](organization-level-integrations/gitlab.md#gitlab-access-tokens) uses either a personal access token (PAT) or group access token (GAT), depending on the GitLab account tier you are on.

To set up the Snyk GitLab integration you must be a:

* Snyk Group or Organization Admin.
* GitLab Owner or Maintainer

A PAT is used for managing personal GitLab projects and requires the `api` scope. For Snyk Essentials to show all repositories from GitLab, the user generating the PAT should be part of the GitLab group where their GitLab permissions can be a minimum of Guest.

A GAT is used for managing multiple GitLab projects in a GitLab group and requires the `api` scope and maintainer role selected from the dropdown. You must be a GitLab Premium or Ultimate account tier holder to create a GAT.

### Bitbucket permission requirements

The Snyk Bitbucket integrations use different access control mechanisms to connect with Snyk:

* [Snyk Bitbucket Cloud](user-permissions-and-access-scopes.md#bitbucket-cloud-and-bitbucket-data-center-server-scopes) requires the creation of an [app password](organization-level-integrations/bitbucket-cloud.md#how-to-set-up-the-bitbucket-cloud-integration).
* [Snyk Bitbucket Cloud App](user-permissions-and-access-scopes.md#bitbucket-cloud-app-scopes) requires [Bitbucket workspace authorization](organization-level-integrations/bitbucket-cloud-app.md#setting-up-a-bitbucket-cloud-app) and related permissions.
* [Snyk Bitbucket Data Center/Server](user-permissions-and-access-scopes.md#bitbucket-cloud-and-bitbucket-data-center-server-scopes) requires a [dedicated username and password](organization-level-integrations/bitbucket-data-center-server.md#how-to-set-up-a-bitbucket-dc-server-integration) or a personal access token (PAT).

{% hint style="warning" %}
To set up any Snyk Bitbucket integration, you must be a Bitbucket Workspace Admin.
{% endhint %}

#### Bitbucket Cloud and Bitbucket Data Center/Server scopes

The following table details the required permission scopes in Bitbucket Cloud and Bitbucket Data Center/Serve&#x72;**:**

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

The [Snyk Azure Repositories (TFS) integration](organization-level-integrations/azure-repositories-tfs.md) uses an Azure DevOps personal access token (PAT). This token is configured with the specific permissions Snyk needs to access your Azure repositories.

To set up the Snyk Azure Repositories (TFS) integration you must be:

* A [Snyk Organization Admin](../snyk-admin/user-roles/pre-defined-roles.md).
* A member of the [Project Administrators group](https://learn.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops\&tabs=preview-page) in Azure. This ensures the PAT has the `edit subscriptions permissions` required to enable webhooks.

In Azure, the PAT requires the following permissions for Snyk access:

* **Expiry**: Custom defined. Snyk recommends choosing a token expiration date that is far in the future.&#x20;
* **Scopes**: Custom defined. `Read & write` permissions are needed for the **Code** scope.
