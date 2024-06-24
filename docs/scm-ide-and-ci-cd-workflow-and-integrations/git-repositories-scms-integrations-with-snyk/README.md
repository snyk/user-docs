# Snyk SCM Integrations

You can integrate Snyk with your Git repository to quickly and easily gain visibility across all the [Snyk Projects](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects) that you add to your **Projects** list.

Snyk Source Control Manager (SCM) integrations allow you to:

* Continuously perform security scanning across all integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes

Snyk can integrate with the following SCMs to help you track, monitor, and fix the issues and vulnerabilities in your code:

* [GitHub Cloud App](../../snyk-scm-ide-and-ci-cd-integrations/git-repositories-scms-integrations-with-snyk/snyk-github-cloud-app.md)
* [GitHub Server App](snyk-github-server-app.md)
* [GitHub Enterprise](snyk-github-enterprise-integration.md)
* [GitHub](snyk-github-integration.md)
* [GitHub Read-only Projects](snyk-github-read-only-projects.md)
* [GitLab](snyk-gitlab-integration.md)
* [Bitbucket Cloud (legacy)](snyk-bitbucket-cloud-integration.md)
* [Bitbucket Cloud App](snyk-bitbucket-cloud-app-integration.md)
* [Bitbucket Data Center/Server](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/snyk-bitbucket-data-center-server-integration.md)
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

### GitHub and GitHub Enterprise permissions requirements

{% hint style="info" %}
For information about token permissions in a brokered integration, see [GitHub - install and configure using Docker](../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md).
{% endhint %}

The [Snyk GitHub integration](snyk-github-integration.md) is bound to a single user, preferably a GitHub service account. The level of access for the integration is defined by the combination of the user's permissions in GitHub and the access defined for the PAT on that user's account. If the PAT is defined with more permission than the user's GitHub account, the integration will not be able to use that permission.

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

The [Snyk GitHub Cloud App](../../snyk-scm-ide-and-ci-cd-integrations/git-repositories-scms-integrations-with-snyk/snyk-github-cloud-app.md) integration uses role-based access control, meaning access control is not dependent on individual users or their role, it is instead tied to the app entity.

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
