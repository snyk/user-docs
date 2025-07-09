# Opening fix and upgrade pull requests from a fixed GitHub account

You can set a specific GitHub account to open, fix, and upgrade PRs, rather than use a GitHub user account picked randomly by Snyk.

{% hint style="info" %}
Other operations not performed by the UI, such as daily and weekly tests, are still performed by Snyk Organization members who have connected their GitHub accounts to Snyk, picked randomly.
{% endhint %}

To use this feature:

1. In Snyk Web UI, navigate to **Settings** > **Integrations** > **Source control** > **GitHub**, and click **Edit Settings**.
2. Under **Open Snyk automatic PRs from a fixed GitHub account**, enter your GitHub personal access token. You can generate this from your GitHub account. See [Managing your personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens).
3. Click **Save**.

{% hint style="info" %}
Ensure that the GitHub account for which the token is provided has write permissions or above to the repos you want to monitor with Snyk.
{% endhint %}

For more information about repository permission levels in GitHub, see [GitHub integration](../../../scm-integrations/organization-level-integrations/github.md).
