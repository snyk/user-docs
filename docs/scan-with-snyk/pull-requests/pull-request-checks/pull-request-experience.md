# Pull Request Experience

{% hint style="info" %}
**Feature availability**

Pull Request Experience supports the following SCM integrations: GitHub, GitHub Enterprise, GitHub Cloud App, GitHub Server App, Bitbucket Cloud, and Bitbucket Cloud App.
{% endhint %}

The Snyk Pull Request Experience reduces context switching by displaying additional information about the PR Check scan results in the pull request.

## Prerequisites for the Snyk Pull Request Experience

* Snyk PR Checks must be enabled. For more information, see [Configure Pull Request Checks](configure-pull-request-checks.md).
* For the GitHub integration, a specific GitHub account needs to be set to open fix and upgrade PRs. The Personal Access Token (PAT) configured in this way is required by the Pull Request Experience to provide a consistent comment experience. For more information, see[ ](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)[Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).

## Issue Summary Comment

{% hint style="info" %}
**Release status**

Issue Summary Comment for Snyk PR Checks is in Early Access.
{% endhint %}

The Issue Summary Comment feature adds a comment to each pull request, summarizing the latest PR Check results. The summary includes the type of checks performed and a breakdown of the findings by severity. Select **View Details** to access the PR Check details in the Snyk Web UI.

<figure><img src="../../../.gitbook/assets/image (586).png" alt=""><figcaption><p>Issue Summary Comment feature for Pull Request Experience</p></figcaption></figure>

## Configure Pull Request Experience

You can configure the Pull Request Experience [at the Integration level](pull-request-experience.md#configure-pr-checks-at-the-integration-level) for your Snyk Organization.

{% hint style="warning" %}
This feature works only for integrations where the Pull Request Experience is configured. If you have multiple integrations you must configure the Pull Request Experience for each one.
{% endhint %}

### Configure Pull Request Experience at the Integration level

Configure the Pull Request Experience for one or more integrations in your Snyk Organization, for which [PR Checks](configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level) are also enabled.

1. In the Snyk Organization you wish to activate and configure the Pull Request Experience for, navigate to **Settings >** **Integrations** and select your connected source code manager to open the settings configuration.
2. Configure and save the following changes:
   1. **Enable issue summary comment**: Enable this option to create an Issue Summary Comment on each pull request, which aggregates the PR Check results. If it is disabled, the entire Pull Request Experience is disabled.
   2. **Create comments for success cases**: By default, an Issue Summary Comment is created even if no vulnerabilities are detected by the PR Check. Disable this option to stop creating Issue Summary Comments for non-failing PR Checks.

<figure><img src="../../../.gitbook/assets/image (1).png" alt=""><figcaption><p>Configuration details for the Pull request experience under your Integration settings</p></figcaption></figure>

For the GitHub integration, configure opening fix and upgrade pull requests from a fixed GitHub account, by providing a Personal Access Token (PAT), which has `write` permissions or above to the repos monitored by Snyk. For more information, see[  ](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)[Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).

<figure><img src="../../../.gitbook/assets/image.png" alt=""><figcaption><p>Configuration details for the GitHub integration</p></figcaption></figure>
