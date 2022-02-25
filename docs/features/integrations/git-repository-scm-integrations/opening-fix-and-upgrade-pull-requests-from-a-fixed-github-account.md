# Opening fix and upgrade pull requests from a fixed GitHub account

You can set a specific GitHub account to open fix and upgrade PRs, rather than use a GitHub user account picked randomly by Snyk.

{% hint style="info" %}
Other operations not performed by the UI, such as daily / weekly tests, are still performed by Snyk organization members who have connected their GitHub accounts to Snyk, picked randomly.
{% endhint %}

To use this feature:

1. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **Integrations**.
2. For the GitHub entry, click **Edit Settings**.
3. Enable the toggle button under the **Open fix and upgrade pull requests from a fixed GitHub account** setting
4. Follow the in-page instructions for creating a personal access token in GitHub.
5. Provide the newly generated token to Snyk so it can be used to perform operations in GitHub (such as opening Fix PRs).

{% hint style="info" %}
Ensure that the GitHub account for which the token is provided, has **write** level permissions or above, to the repos you'd like to monitor with Snyk.
{% endhint %}

[Read more](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration) about repository permission levels in GitHub.
