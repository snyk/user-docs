# Opening fix and upgrade pull requests from a fixed GitHub account

You can set a specific GitHub account to open, fix, and upgrade PRs, rather than use a GitHub user account picked randomly by Snyk.

{% hint style="info" %}
Other operations not performed by the UI, such as daily and weekly tests, are still performed by Snyk Organization members who have connected their GitHub accounts to Snyk, picked randomly.
{% endhint %}

To use this feature:

1. In Snyk Web UI, navigate to **Settings** > **Integrations** > **Source control** > **GitHub**, and click **Edit Settings**.
2. Enable the toggle button under the **Open fix and upgrade pull requests from a fixed GitHub account** setting.
3. Follow the in-page instructions for creating a personal access token in GitHub.
4. Provide the newly generated token to Snyk so it can be used to perform operations in GitHub, such as opening Fix PRs.

{% hint style="info" %}
Ensure that the GitHub account for which the token is provided has write permissions or above to the repos you want to monitor with Snyk.
{% endhint %}

For more information about repository permission levels in GitHub, see [GitHub integration](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md).
