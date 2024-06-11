# Snyk SCM Integrations

You can integrate Snyk with your Git repository to quickly and easily gain visibility across all the [Snyk Projects](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects) that you add to your **Projects** list.

Snyk Source Control Manager (SCM) integrations allow you to:

* Continuously perform security scanning across all integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes

Snyk can integrate with the following SCMs to help you track, monitor, and fix the issues and vulnerabilities in your code:

* [GitHub Cloud App](snyk-github-cloud-app.md)
* [GitHub Enterprise](snyk-github-enterprise-integration.md)
* [GitHub](snyk-github-integration.md)
* [GitHub Read-only Projects](snyk-github-read-only-projects.md)
* [GitLab](snyk-gitlab-integration.md)
* [Bitbucket Cloud](snyk-bitbucket-cloud-integration.md)
* [Bitbucket Cloud (Legacy)](migrate-a-bitbucket-cloud-personal-access-token.md)
* [Bitbucket Cloud App](snyk-bitbucket-cloud-app-integration.md)
* [Bitbucket Data Center/Server](snyk-bitbucket-data-center-server-integration.md)
* [Azure Repositories (TFS)](snyk-azure-repositories-tfs-integration.md)

## User permissions and access scopes for SCM integrations

This section details:

* [GitHub and GitHub Enterprise permission requirements](./#github-and-github-enterprise-permissions-requirements)

### GitHub and GitHub Enterprise permissions requirements

{% hint style="info" %}
For information about token permissions in a brokered integration, see [GitHub - install and configure using Docker](../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md).
{% endhint %}

The Snyk GitHub integration is bound to a single user, preferably a GitHub service account. The level of access for the integration is defined by the combination of the user's permissions in GitHub and the access defined for the PAT on that user's account. If the PAT is defined with more permission than the user's GitHub account, the integration will not be able to use that permission.

The following table details the access scopes required in GitHub for Personal Access Tokens (PAT) and the scopes required for Snyk to perform the required operations on monitored repositories, such as reading manifest files on a frequent basis and opening fix or upgrade PRs. GitHub custom roles are not supported.

<table><thead><tr><th width="259">Action and purpose</th><th align="center">PAT scopes</th><th align="center">Repository scopes</th></tr></thead><tbody><tr><td><strong>Daily/weekly tests:</strong><br>Read manifest files in private repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>read</code></td></tr><tr><td><strong>Manual fix pull requests:</strong><br>Create fix PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">N/A</td></tr><tr><td><strong>Automatic fix and upgrade pull requests:</strong><br>Create fix or upgrade PRs in monitored repositories.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Snyk tests on pull requests:</strong><br>Send PR status checks whenever a new PR is created, or an existing PR is updated.</td><td align="center"><code>repo (all)</code></td><td align="center">≥ <code>write</code></td></tr><tr><td><strong>Initial configuration of Snyk tests on pull requests:</strong><br>Used to add SCM webhooks to the imported repo</td><td align="center"><code>admin:repo_hooks (read &#x26; write)</code></td><td align="center"><code>admin</code></td></tr><tr><td><strong>Import new Projects to Snyk:</strong><br>Present a list of all the available repos in the GitHub org in the <strong>Add Projects</strong> screen.</td><td align="center"><code>admin:read:org</code><br><code>repo (all)</code></td><td align="center">N/A</td></tr></tbody></table>

Snyk uses PRs to tell GitHub Enterprise that a merge is to occur. To do this, change content is pushed into a branch, which requires the `content: write` scope. A separate call is then made to create the fix PR, which requires the `pull request: write` scope. GitHub Enterprise is then instructed to create a PR, merging the change branch into the default branch.

Snyk uses SCM webhooks to:

* Track the state of Snyk pull requests when PRs are created, updated triggered, merged, and so on.
* Send push events to trigger PR checks.
