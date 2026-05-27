# Pull Request experience

{% hint style="info" %}
**Release status**

As part of the Pull Request experience, Snyk Agent fix in the PR is in [Early Access](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/snyk-release-process#early-access-features).
{% endhint %}

The Pull Request experience builds on the foundational [Pull Request Checks](./) functionality. It streamlines the security review process by reducing the need to switch between different developer tools and providing contextually relevant feedback within your Source Code Manager (SCM) interface.

The Pull Request Experience consists of the following features:

* [Issue Summary Comment](pull-request-experience.md#issue-summary-comment)
* [Inline Comments](pull-request-experience.md#inline-comments)
* [Snyk Agent Fix in the PR](pull-request-experience.md#snyk-agent-fix-in-the-pr)

Issue summary comment provides a collated view of the last PR Check results, categorizing findings by severity and type directly within the pull request.

Inline comments give a granular view of the pull request with information on severity, the data flow of vulnerabilities, and more. This allows you to make quick decisions on issue prioritization and remediation.

Snyk Agent fix in the PR enables action to be taken based on the previous features in the Pull Request Experience and the recommendations given.

## Prerequisites

### User role requirement

To configure and manage the pull request experience, the user must be a [Group Admin](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/user-management/pre-defined-roles#group-level-permissions). This is to ensure access to all integrations for setup as the pull request experience is configured at the Organization level.

### Pull request checks enablement

{% hint style="info" %}
The Pull request experience is set up individually for each integration. If you have multiple integrations, you need to configure the pull request experience separately for each one.
{% endhint %}

Ensure the following is complete:

* [Configure PR checks](configure-pull-request-checks.md).
* Set up at least one SCM integration with Snyk and import a minimum of one Project from your repository.
* Snyk Code is enabled for your Snyk Organization. Contact your Account Representative or [Snyk Support](https://support.snyk.io/s/) if you do not have this enabled.

### Pull request experience feature requirements

In addition to the general SCM and PR Checks prerequisites, certain features in the pull request experience have their own requirements:

* To configure the inline comments feature, enable the **Code analysis** PR checks setting. This is located on the Organization level under **Settings** > **Integrations**.
* If you are using the GitHub Integration, specify a dedicated GitHub account by supplying a **GitHub Personal Access Token (PAT)** in the integration settings. This is required to be able to add inline comments or Agent Fix

### SCM permission and access scope requirements

The pull request experience integrates with various SCM platforms, each with specific requirements for a successful configuration with Snyk. Your existing SCM integration setup will work with the Pull Request experience out of the box, except for GitHub (OAuth) which requires an additional Fix PR token. For additional information, see [User permissions and access scopes](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/user-permissions-and-access-scopes).

{% hint style="info" %}
For information on which SCM integrations are supported in each pull request experience feature, see the individual feature sections: [Issue Summary Comment](pull-request-experience.md#issue-summary-comment), [Inline Comments](pull-request-experience.md#inline-comments), and [Snyk Agent Fix in the PR](pull-request-experience.md#snyk-agent-fix-in-the-pr).
{% endhint %}

## Configure pull request experience at the integration level

Configure the pull request experience for one or more integrations in your Snyk Organization that also have [PR checks](configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level) enabled.

1. Select the Snyk Organization for which you want to enable the pull request experience. Navigate to **Settings** > **Integrations** > **Source control** > **Edit settings** to open the settings configuration.
2. Configure and save the following changes:
   1. Enable [issue summary comment](pull-request-experience.md#issue-summary-comment): Enable this option to create an issue summary comment on each pull request, which aggregates the PR check results. If it is disabled, the entire pull request experience is disabled.
   2. Create comments for success cases: By default, an issue summary comment is created even if no vulnerabilities are detected by the PR check. Disable this option to stop creating issue summary comments for non-failing PR checks.
   3. Enable[ inline comments](pull-request-experience.md#inline-comments)**:** Enable inline comments to add a comment for each issue found by Snyk Code PR check.
   4. Enable [Snyk Agent fix in the PR](pull-request-experience.md#snyk-agent-fix-in-the-pr): Enable requesting Snyk Agent fix suggestions for issues found by Snyk Code PR check, and applying the fixes to the PR branch.

## Issue summary comment

The issue summary comment feature adds a comment to each pull request, summarizing the latest PR check results. The summary includes the type of checks performed and a breakdown of the findings by severity. Select **View Details** to access the PR check details in the Snyk Web UI.

## Inline comments

The inline comments feature adds a detailed comment for each issue identified by the Snyk Code pull request check. Each comment includes the severity level, the name, and a short description of the issue, helpful links for further information, and, if applicable, the data flow. For best results, Snyk recommends generating and applying fixes for a single inline comment at a time to avoid situations where applying a fix causes conflicts with another previously generated fix.

{% hint style="info" %}
This feature is limited to 10 inline comments at pull request level. The summary comment will display a message if the cap is surpassed.
{% endhint %}

For GitLab and Azure Repos, consider the following conditions:

* The Data Flow section in the inline comments is not available.
* Inline comments for vulnerabilities introduced outside modified lines are not available in GitLab.
* Inline comments for Code Analysis generated by the Snyk Local Code Engine are unavailable.

For Brokered integrations, the Data Flow section in the inline comments is available only for GitHub, GitHub Cloud App, Bitbucket Cloud, and Bitbucket Connect App.

## Snyk Agent fix in the PR

{% hint style="info" %}
**Release status**

Snyk Agent fix in the PR is in [Early Access](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/snyk-release-process#early-access-features).\
Snyk Agent fix in the PR will work only on inline comments created after the feature is enabled.
{% endhint %}

The Snyk Agent fix in the PR feature enables the user to request and apply fixes for vulnerabilities identified by the Snyk Code pull request check and posted as inline comments. By enabling this feature, the user is able to interact with inline comments in the following way:

* Request an initial fix by replying to an inline comment using the `@snyk /fix` command.

<figure><img src="../../../.gitbook/assets/image (299).png" alt=""><figcaption><p>Inline Comments with Snyk Agent Fix enabled</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (300).png" alt=""><figcaption><p>Request a fix by replying to the inline comment</p></figcaption></figure>

* Request a different suggestion by replying with the `@snyk /fix` command to a previously generated fix. Snyk Agent fix can generate up to five potential fixes, depending on the issue type.
* Apply a specific fix by using the `@snyk /apply #` command, where # is the number of the suggestion the user wishes to apply. A commit is created by Snyk on the PR branch, containing the selected fix.

### Exceptions

* The `@snyk /fix` command can be used only for automatically fixable vulnerabilities, identified in the inline comments with a zap icon and command description. See [fix-code-vulnerabilities-automatically.md](../../snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md "mention") for supported languages and limitations.
* Fixes expire after the time displayed in each suggestion, in accordance with the [How Snyk handles your data #Cache retention period related to vulnerability source data](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-snyk-handles-your-data#cache-retention-period-related-to-vulnerability-source-data "mention"). After expiration, a new fix can be requested by using the `@snyk /fix` command.
* Snyk Agent fix in the PR is not supported for [Snyk Code Local Engine](../../snyk-code/snyk-code-local-engine.md).
* The `@snyk /fix` and `@snyk /apply #` commands can be used only as replies to the Inline Comments created by Snyk, commands created on other comment threads will not be processed.

## Troubleshooting

{% hint style="info" %}
For detailed descriptions on troubleshooting scenarios and their relevant resolutions, see [Troubleshoot PR Checks](troubleshoot-pr-checks.md).
{% endhint %}

### Block pull requests based on Snyk findings

A PR test is configured to be optional or blocking within your source control management platform, such as GitHub’s branch protection rules. To learn more about issue prevention, visit [Enable and configure Snyk on PRs](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/implementation-guides/team-implementation-guide/phase-5-rolling-out-the-prevention-stage/enable-and-configure-snyk-on-prs).

### Handle false positives

If a finding is a false positive, you can manually resolve the comment in the SCM. It will remain resolved even if the issue is detected again in a subsequent PR check.

### Understand comment behavior when a new commit is pushed

#### Issue summary comments

When you push a new commit to a pull request, Snyk updates the existing summary comment with the latest scan results. This refreshes the issue count to reflect the most recent analysis. Snyk does not add a new issue summary comment, ensuring you get a concise view of potential vulnerabilities without receiving excessive notifications.

#### Inline comments

For every pushed commit, Snyk creates a review if the Snyk Code PR Check detects new issues. Snyk adds each issue as an inline comment. If you fix an issue from a previous commit, Snyk marks the corresponding comment as resolved. Unresolved issues remain as inline comments.

### PR comments do not appear in the repository

To ensure your repository receives PR comments, review the [Prerequisites for the Snyk Pull Request Experience](pull-request-experience.md#prerequisites) section. Ensure all required conditions are met and properly fulfilled.

Snyk Code pull request (PR) checks provide inline comments only for detected issues. To ensure issues appear as inline comments, verify these requirements:

* Enable Snyk Code for your Organization in **Organization settings** > **Snyk Code**.
* Enable **Code Analysis** in the **Pull request status checks** section of your integration settings. Snyk adds inline comments only for issues it detects during a PR check. If an issue has a lower severity than your fail condition, the check does not detect it.
* Enable **Inline Comments** in the **Pull Request Experience** section of your integration settings.

### Certain findings are not appearing as inline comments

Inline comments are only available for issues detected by Snyk Code PR checks. Their behavior may vary depending on the SCM platform. On most platforms, an inline comment is added for each issue found. However, on GitLab, inline comments are only added for issues found within the files that were modified in the pull request.

### Check Snyk Broker support for PR comments

Snyk supports brokered integrations for issue summary comments and inline comments. You can use both Classic Broker and Universal Broker.

The following features require specific Snyk Broker versions:

* **Issue summary and inline comments:** Version 4.194 or higher.
* **Snyk Agent fix in pull requests:** Version 4.219 or higher.

### The Snyk Agent fix in the PR is not working after enabling the setting

When you enable the Snyk Agent Fix in the PR setting, a background process is initiated to upgrade the [Snyk webhooks](../snyk-pull-or-merge-requests/#snyk-scm-webhooks) in the user's repository to include the pull request comment event subscription required for the feature. This process can take a couple of minutes to complete. During this time, the feature will not be fully active or available, and Snyk will not be able to react to commands in the PR.

To troubleshoot the webhook upgrade process, go to the repository settings page of your SCM, and confirm that the Snyk webhook is subscribed to the 'Pull request review comments' event.

### Snyk Agent fix returns "command execution is already in progress" when I submit a single command

This error occurs if you import a repository into multiple Snyk organizations and define multiple Snyk webhooks for it.

Ensure only one Snyk webhook subscribes to pull request comment events. You can edit the Snyk webhook event subscriptions on the repository settings page of your source control manager (SCM).

### A comment failed with a message about replying to the first comment in a Bitbucket thread

To resolve this error, reply to the first comment in the thread with your command.

Commands in comment replies fail in long or complex Bitbucket threads because the thread length or structure prevents Snyk from processing the reply. Replying to the first comment ensures Snyk processes the command.

### Applying Snyk Agent Fix does not work on some pull requests from forks in Bitbucket Cloud

To enable Snyk AI Agent Fix for cross-workspace pull requests in Bitbucket Cloud, switch to the Bitbucket Cloud PAT integration.

The Bitbucket Cloud Connect app restricts Snyk AI Agent Fix when the integration only has access to the base workspace. Because of this limitation, Snyk AI Agent Fix supports pull requests from repositories within the same workspace (for example, `base_workspace/fork_repo` to `base_workspace/base_repo`). It does not support pull requests from a fork located in a different workspace (for example, `fork_workspace/fork_repo` to `base_workspace/base_repo`).

### Troubleshoot the Snyk Agent Fix feature on Bitbucket Cloud

To troubleshoot agent-based fixes on Bitbucket Cloud, inspect the webhook history in your Bitbucket repository settings.

Snyk Agent Fix relies on webhooks to communicate between Bitbucket Cloud and Snyk. Check the webhook history to verify successful webhook delivery and inspect service responses. This helps diagnose network issues, configuration problems, or service errors.

To access the webhook history in Bitbucket Cloud:

1. Navigate to your repository on Bitbucket Cloud.
2. Click **Repository settings**.
3. Under **Workflow**, select **Webhooks**.
4. Locate the webhook used for the integration.
5. Click **View requests** next to the webhook URL.
6. Click **Enable History**.
7. Post a new Snyk Agent Fix command to view the webhook deliveries.
