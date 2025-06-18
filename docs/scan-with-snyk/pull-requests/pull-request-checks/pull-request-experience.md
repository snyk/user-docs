# Pull Request Experience

The Snyk Pull Request Experience reduces context switching by displaying additional information about the PR Check scan results in the pull request.

The following features are part of the Pull Request Experience:

* [Issue Summary Comment](pull-request-experience.md#issue-summary-comment)
* [Inline Comments](pull-request-experience.md#inline-comments)
* [Snyk Agent Fix in the PR](pull-request-experience.md#snyk-agent-fix-in-the-pr)

## Prerequisites for the Snyk Pull Request Experience

* Snyk PR Checks must be enabled. For more information, see [Configure Pull Request Checks](configure-pull-request-checks.md).
* For the GitHub integration, you must specify a GitHub account  to provide a consistent comment experience. For more information, see [Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md).
* For inline comments, enable the **Code analysis PR checks** setting at the [integration level](pull-request-experience.md#configure-pull-request-experience-at-the-integration-level). In the Snyk Web UI, under **Settings** > **Integrations** > **Edit settings**,  verify that the option is enabled under **Pull request status checks** > **Code analysis**.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-11-27 at 14.14.35.png" alt=""><figcaption><p>Code analysis feature for Pull Request status checks</p></figcaption></figure>

## Issue Summary Comment

{% hint style="info" %}
**Release status**

Issue Summary Comment for Snyk PR Checks is in [Early Access](../../../getting-started/snyk-release-process.md#early-access-features) for Bitbucket Server, GitLab and Azure Repos, and available only for Snyk Code PR Checks.

Issue Summary Comment for Snyk PR Checks fully supports GitHub integrations and Bitbucket Cloud integrations.
{% endhint %}

The Issue Summary Comment feature adds a comment to each pull request, summarizing the latest PR Check results. The summary includes the type of checks performed and a breakdown of the findings by severity. Select **View Details** to access the PR Check details in the Snyk Web UI.

<figure><img src="../../../.gitbook/assets/image (586).png" alt=""><figcaption><p>Issue Summary Comment feature for Pull Request Experience</p></figcaption></figure>

## Inline Comments

{% hint style="info" %}
**Release status**

Inline Comments for Snyk PR Checks is in [Early Access](../../../getting-started/snyk-release-process.md#early-access-features) for Bitbucket Server, GitLab and Azure Repos, and available only for Snyk Code PR Checks.

Inline Comments for Snyk PR Checks fully supports GitHub integrations and Bitbucket Cloud integrations.
{% endhint %}

The Inline Comments feature adds a detailed comment for each issue identified by the Snyk Code Pull Request Check. Each comment includes the severity level, the name and a short description of the issue, helpful links for further information, and, if applicable, the data flow.&#x20;

This feature is limited to 10 inline comments at Pull Request level. The Summary Comment will display a message if the cap is surpassed.

For GitLab and Azure Repos, consider the following conditions:

* The Data Flow section in the inline comments is not available.
* Inline comments for vulnerabilities introduced outside modified lines are unavailable for GitLab.
* Inline comments for Code Analysis done using Snyk Local Code Engine are unavailable.

For Brokered integrations, the Data Flow section in the inline comments is available only for GitHub, GitHub Cloud App, Bitbucket Cloud, and Bitbucket Connect App.

<figure><img src="../../../.gitbook/assets/inline_comment_feature.png" alt=""><figcaption><p>Inline Comment feature for Pull Request Experience</p></figcaption></figure>

## Snyk Agent Fix in the PR

{% hint style="info" %}
**Release status**

Snyk Agent Fix in the PR is in [Early Access](https://docs.snyk.io/getting-started/snyk-release-process#early-access)  and available only for GitHub Integrations.
{% endhint %}

The Snyk Agent Fix in the PR feature enables the user to request and apply fixes for vulnerabilities identified by the Snyk Code Pull Request Check and posted as inline comments. By enabling this feature, the user is able to interact with inline comments in the following way:&#x20;

* Request an initial fix by replying to an inline comment using the `@snyk /fix` command.

<figure><img src="../../../.gitbook/assets/image (20).png" alt=""><figcaption><p>Inline Comments with Snyk Agent Fix enabled</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (22).png" alt=""><figcaption><p>Request a fix by replying to the inline comment</p></figcaption></figure>

* Request a different suggestion by replying with the `@snyk /fix` command to a previously generated fix. Snyk Agent Fix can generate up to five potential fixes, depending on the issue type.
* Apply a specific fix by using the `@snyk /apply #` command, where # is the number of the suggestion the user wishes to apply. A commit is created by Snyk on the PR branch, containing the selected fix.

#### Notes

* The `@snyk /fix` command can be used only for automatically fixable vulnerabilities, identified in the inline comments with a zap icon and command description. See  [fix-code-vulnerabilities-automatically.md](../../snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md "mention") for supported languages and limitations.
* Fixes expire after the time displayed in each suggestion, in accordance with the [#cache-retention-period-related-to-vulnerability-source-data](../../../working-with-snyk/how-snyk-handles-your-data.md#cache-retention-period-related-to-vulnerability-source-data "mention"). After expiration, a new fix can be requested by using the `@snyk /fix` command.
* Snyk Agent Fix in the PR is not supported for [Snyk Code Local Engine](../../snyk-code/snyk-code-local-engine.md).
* The `@snyk /fix` and `@snyk /apply #` commands can be used only as replies to the Inline Comments created by Snyk, commands created on other comment threads will not be processed.
* For best results, Snyk recommends generating and applying fixes for a single inline comment at a time, to avoid situations where applying a fix causes conflicts with another previously generate fix.

## Configure Pull Request Experience

You can configure the Pull Request Experience [at the Integration level](pull-request-experience.md#configure-pull-request-experience-at-the-integration-level) for your Snyk Organization.

{% hint style="warning" %}
The Pull Request Experience is set up individually for each integration. If you have multiple integrations, you need to configure the Pull Request Experience separately for each one.
{% endhint %}

### Configure Pull Request Experience at the Integration level

Configure the Pull Request Experience for one or more integrations in your Snyk Organization, for which [PR Checks](configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level) are also enabled.

1. In the Snyk Organization you wish to activate and configure the Pull Request Experience for, navigate to **Settings** > **Integrations** and select **Edit settings** your connected source code manager to open the settings configuration.
2. Configure and save the following changes:
   1. **Enable Issue Summary Comment:** Enable this option to create an Issue Summary Comment on each pull request, which aggregates the PR Check results. If it is disabled, the entire Pull Request Experience is disabled.
   2. **Create comments for success cases:** By default, an Issue Summary Comment is created even if no vulnerabilities are detected by the PR Check. Disable this option to stop creating Issue Summary Comments for non-failing PR Checks.
   3. **Enable Inline Comments:** Enable inline comments to add a comment for each issue found by Snyk Code PR Check.
   4. **Enable Snyk Agent Fix in the PR:** Enable requesting Snyk Agent Fix suggestions for issues found by Snyk Code PR Check, and applying the fixes to the PR branch.&#x20;

<figure><img src="../../../.gitbook/assets/image (23).png" alt=""><figcaption><p>Configuration details for the GitHub integration</p></figcaption></figure>

For the GitHub integration, configure opening fix and upgrade pull requests from a fixed GitHub account, by providing a Personal Access Token (PAT), which has `write` permissions or above to the repos monitored by Snyk. See[ ](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)[Opening fix and upgrade pull requests from a fixed GitHub account](../snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md) for more information.

<figure><img src="../../../.gitbook/assets/pat_config.png" alt=""><figcaption><p>Configuration details for the GitHub integration</p></figcaption></figure>

## Frequently Asked Questions (FAQs)

### How can I handle false positives?

If a finding is a false positive, you can manually resolve the comment in the SCM. It will remain resolved even if the issue is detected again in a subsequent PR check.

### What happens to comments when a new commit is pushed?

#### **Issue Summary Comment**

When a new commit is pushed to the pull request, the existing summary comment is updated with the latest Snyk scan results for that commit. This means the issues count is refreshed to reflect the most recent analysis. No new issue summary comment is added, ensuring developers get a concise view of potential vulnerabilities without being flooded with notifications.

#### **Inline Comments**

For every pushed commit, a review is created if Snyk Code PR Check detects new issues. Each issue is added as an inline comment. If an issue from a previous commit is fixed in the new commit, its comment is marked as resolved. Unresolved issues remain as inline comments.

### Why are PR comments not appearing in my repository?

To ensure your repository receives PR comments, review the [Prerequisites for the Snyk Pull Request Experience](pull-request-experience.md#prerequisites-for-the-snyk-pull-request-experience) section. Ensure all required conditions are met and properly fulfilled.

Inline comments are available only for issues detected by Snyk Code PR Checks. To ensure that issues appear as inline comments, verify the following requirements:

1. **Snyk Code** is enabled for your Organization.

<figure><img src="../../../.gitbook/assets/enable_snyk_code.png" alt=""><figcaption><p>Configuration details to enable Snyk Code</p></figcaption></figure>

2. **Code Analysis** is enabled in the Pull request status checks section for your integration.

<figure><img src="../../../.gitbook/assets/enable_code_analysis.png" alt=""><figcaption><p>Configuration details to enable Code Analysis</p></figcaption></figure>

3. **Inline Comments** are enabled in the **Pull Request Experience** section for your integration.

<figure><img src="../../../.gitbook/assets/image (24).png" alt=""><figcaption><p>Configuration details to enable Inline Comments</p></figcaption></figure>

### Why are some findings not appearing as inline comments?

Inline comments are only available for issues detected by Snyk Code PR Checks. Their behavior may vary depending on the SCM platform. On most platforms, an inline comment is added for each issue found. However, on GitLab, inline comments are only added for issues found within the files that were modified in the pull request.

### Is Snyk Broker supported?

Yes. Brokered integrations are supported for both Issue Summary Comment and Inline Comments, with compatibility for both Classic and Universal Broker.&#x20;

Supported Snyk Broker version 4.194 or higher.

Snyk Agent Fix in the PR is supported in Snyk Broker version 4.216 or higher.

### Why is Snyk Agent Fix in the PR not working after enabling the setting?

When you enable the Snyk Agent Fix in the PR setting, a background process is initiated to upgrade the [Snyk webhooks](../snyk-pull-or-merge-requests/#snyk-scm-webhooks) in the user's repository with the pull request comment event subscription necessary for the feature. This process can take a couple of minutes to complete. During this time, the feature will not be fully active or available, and Snyk will not be able to react to commands in the PR.&#x20;

Snyk Agent Fix in the PR will work only on inline comments created after the feature is enabled.

To troubleshoot the webhook upgrade process, go to the repository settings page of your SCM, and confirm that the Snyk webhook is subscribed to the 'Pull request review comments' event.

### Why is Snyk Agent Fix replying with command execution is already in progress, when I submitted a single command?

This can happen when a repository is imported into multiple Snyk Organisations and there are multiple Snyk webhooks defined on the repository. For best results ensure that only a single Snyk webhook is subscribed to pull request comment events. If needed, the Snyk webhook event subscriptions can be edited from the repository settings page of your SCM.
