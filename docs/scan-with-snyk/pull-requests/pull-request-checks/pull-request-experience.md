# Pull Request Experience

{% hint style="info" %}
**Feature availability**

Pull Request Experience supports the following SCM integrations: GitHub, GitHub Enterprise, GitHub Cloud App, GitHub Server App, Bitbucket Cloud, and Bitbucket Cloud App.
{% endhint %}

The Snyk Pull Request Experience reduces context switching by displaying additional information about the PR Check scan results in the pull request.

The following features are part of the Pull Request Experience:

* [Issue Summary Comment](pull-request-experience.md#issue-summary-comment)
* [Inline Comments](pull-request-experience.md#inline-comments)

## Prerequisites for the Snyk Pull Request Experience

* **Snyk PR Checks** must be enabled. For more information, see [Configure Pull Request Checks](configure-pull-request-checks.md).
* For the **GitHub integration**, a specific **GitHub account** needs to be set to open fix and upgrade PRs. The Personal Access Token (PAT) configured in this way is required by the Pull Request Experience to provide a consistent comment experience. See[ ](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)[Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md) for more information.
* For inline comments, the **Code analysis PR checks** setting needs to be enabled at the [integration level](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience#configure-pull-request-experience-at-the-integration-level). In the Snyk Web UI, under Settings > Integrations > Edit settings,  verify that the option is enabled under **Pull request status checks > Code analysis**. If "Code analysis" is not visible, please reach out to your Snyk account team to enable Snyk Code for your account. If it is already enabled, check under **Settings** > **Snyk Code** to ensure it is turned on.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-11-27 at 14.14.35.png" alt=""><figcaption><p>Code analysis feature for Pull Request status checks</p></figcaption></figure>

## Issue Summary Comment

{% hint style="info" %}
**Release status**

Issue Summary Comment for Snyk PR Checks is in [Early Access](../../../getting-started/snyk-release-process.md#early-access).
{% endhint %}

The Issue Summary Comment feature adds a comment to each pull request, summarizing the latest PR Check results. The summary includes the type of checks performed and a breakdown of the findings by severity. Select **View Details** to access the PR Check details in the Snyk Web UI.

<figure><img src="../../../.gitbook/assets/image (586).png" alt=""><figcaption><p>Issue Summary Comment feature for Pull Request Experience</p></figcaption></figure>

## Inline Comments

{% hint style="info" %}
**Release status**

Inline Comments for Snyk PR Checks is in [Early Access](../../../getting-started/snyk-release-process.md#early-access) and available only for **Snyk Code PR Checks**.
{% endhint %}

The Inline Comments feature adds a detailed comment for each issue identified by the Snyk Code Pull Request Check. Each comment includes the severity level, the name and a short description of the issue, helpful links for further information, and, if applicable, the data flow.&#x20;

This feature is limited to 10 inline comments at Pull Request level. The Summary Comment will display a message if the cap is surpassed.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-11-27 at 14.03.12.png" alt=""><figcaption><p>Inline Comment feature for Pull Request Experience</p></figcaption></figure>

## Configure Pull Request Experience

You can configure the Pull Request Experience [at the Integration level](pull-request-experience.md#configure-pr-checks-at-the-integration-level) for your Snyk Organization.

{% hint style="warning" %}
The Pull Request Experience is set up individually for each integration. If you have multiple integrations, you need to configure the Pull Request Experience separately for each one.
{% endhint %}

### Configure Pull Request Experience at the Integration level

Configure the Pull Request Experience for one or more integrations in your Snyk Organization, for which [PR Checks](configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level) are also enabled.

1. In the Snyk Organization you wish to activate and configure the Pull Request Experience for, navigate to **Settings,** **Integrations** and select **Edit settings** your connected source code manager to open the settings configuration.
2. Configure and save the following changes:
   1. **Enable issue summary comment**: Enable this option to create an Issue Summary Comment on each pull request, which aggregates the PR Check results. If it is disabled, the entire Pull Request Experience is disabled.
   2. **Create comments for success cases**: By default, an Issue Summary Comment is created even if no vulnerabilities are detected by the PR Check. Disable this option to stop creating Issue Summary Comments for non-failing PR Checks.
   3. **Enable inline comments**: Enable inline comments to add a comment for each issue found by Snyk Code PR Check.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-11-27 at 11.44.50.png" alt=""><figcaption><p>Configuration details for the GitHub integration</p></figcaption></figure>

For the **GitHub integration**, configure opening fix and upgrade pull requests from a fixed **GitHub account**, by providing a Personal Access Token (PAT), which has `write` permissions or above to the repos monitored by Snyk. See[  ](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)[Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md) for more information.

<figure><img src="../../../.gitbook/assets/image (1).png" alt=""><figcaption><p>Configuration details for the GitHub integration</p></figcaption></figure>
