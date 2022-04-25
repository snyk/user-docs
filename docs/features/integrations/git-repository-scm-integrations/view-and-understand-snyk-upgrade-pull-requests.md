# View and understand Snyk upgrade pull requests

In addition to fix advice, Snyk can automatically create pull requests (PRs) on your behalf in order to upgrade your dependencies based on the scan results. Snyk currently supports this feature for npm, Yarn, Maven-Central, NuGet (C#), Pip and PyPI (Python), and Bundler (Ruby) projects through GitHub, GitHub Enterprise Server, BitBucket Cloud, GitLab, and Azure repos.&#x20;

To use with the Broker, your administrator should first upgrade to v4.55.0 or later.

{% hint style="info" %}
**NOTE**\
Administrators and account owners manage settings for Snyk upgrade pull requests from the Snyk Web UI on both the organization and the project levels, by configuring whether the feature is on (enabled by default) and under what conditions Snyk should submit upgrade pull requests, if at all.
{% endhint %}

## Viewing pull request details before merging

Once Snyk submits an upgrade pull request on your behalf, you can view the pull request and all related details directly from the relevant repository.

To quickly review the pull request, hover over it and you can see the recommended upgrade and other pull request summary details:

![](../../../.gitbook/assets/uuid-3683a529-6856-d15d-c49c-ca7ed318500d-en.png)

Open the pull request to view in-depth details, including package release notes, and vulnerabilities included in the recommended upgrade.:

![](../../../.gitbook/assets/uuid-508983f5-8844-c19f-a43e-5a65e4ffdae9-en.png)

Click the Issue link from the table to view all details for the specified vulnerability, directly from our database.

Once you've reviewed the pull request, you can approve it for merge.
