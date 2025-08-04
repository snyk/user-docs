# Snyk Pull or Merge Requests

In addition to fixing advice, Snyk can automatically create pull requests (PRs) on your behalf to upgrade your dependencies based on the scan results. To create PRs automatically in implementations with Snyk Broker, your administrator should first upgrade to v4.55.0 or later.

{% hint style="info" %}
For the basic steps in fixing vulnerabilities, see  [Fix your vulnerabilities](../../snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md). To ensure your language is supported, see [Languages supported for Fix Pull Requests or Merge Requests](../../snyk-open-source/manage-vulnerabilities/troubleshoot-fixing-vulnerabilities-with-snyk-open-source.md#languages-supported-for-fix-pull-requests-or-merge-requests) and [Supported browsers](../../../getting-started/#supported-browsers) pages.
{% endhint %}

{% hint style="info" %}
Administrators and account owners can manage the settings for Snyk upgrade pull requests from the Snyk Web UI at both the Organization and Project levels. You can configure whether the feature is enabled (which is the default) and specify the conditions under which Snyk should submit upgrade pull requests, if at all.
{% endhint %}

Snyk provides actionable fix advice for vulnerabilities in your Open Source libraries through the following:

* [#automated-snyk-prs](./#automated-snyk-prs "mention"): automatic pull and merge requests
* [#manual-snyk-prs](./#manual-snyk-prs "mention"): manual pull and merge requests

## **Automated Snyk PRs**

For Projects imported through an SCM integration, Snyk offers the following types of automated pull and merge requests:

* [Fix pull requests for new vulnerabilities](create-automatic-prs-for-new-fixes-fix-prs.md) (Fix PRs)
* [Fix pull requests to clear the backlog of vulnerabilities in priority order](create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md) (Backlog PRs)
* [Dependency upgrade pull requests](upgrade-dependencies-with-automatic-prs-upgrade-prs/) (Upgrade PRs)

## Manual Snyk PRs

Follow these steps to generate a PR or MR directly from your Project in the Snyk Web UI:

1. Navigate to your Project from the Project list.
2. Select the Project.
3. Select **Open a Fix PR/MR** or **Fix this vulnerability.** A preview screen appears, showing you what fixes will be applied.
4. Click **Open a Fix PR** on this screen to generate the pull request.

{% hint style="info" %}
PRs use a branch naming convention based on the issues that they fix. When a PR already exists for a specific change, Snyk does not create a new one, even if the original PR is closed. If you try to create a fix PR that duplicates an existing one, an error may be displayed. If this happens, check to see if the branch already exists and re-open it.
{% endhint %}

## Reviewing the Snyk PRs

After Snyk submits a pull request on your behalf, you can view the pull request and all related details directly from the relevant repository.

To quickly review the pull request, hover over it. You can see the recommended upgrade and other pull request summary details:

<figure><img src="../../../.gitbook/assets/open-a-fix-pr-github.png" alt="Recommended upgrade"><figcaption><p>Recommended upgrade</p></figcaption></figure>

Open the pull request to view in-depth details, including package release notes and vulnerabilities included in the recommended upgrade.

<figure><img src="../../../.gitbook/assets/github-fix-pr-details.png" alt="Pull request details"><figcaption><p>Pull request details</p></figcaption></figure>

Click the Issue link from the table to view all details for the specified vulnerability directly from the Snyk database.

After you have reviewed the pull request, you can approve the merge.

## Generated Pull Requests report

Snyk provides a report which for Enterprise plan customers that gives an overview of how [Fix](create-automatic-prs-for-new-fixes-fix-prs.md), [Backlog](create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md), and [Upgrade PRs](upgrade-dependencies-with-automatic-prs-upgrade-prs/) are used and highlights the efficiency of PR merges. For more information, see [Snyk Generated Pull Requests report](../../../manage-issues/reporting/available-snyk-reports.md#snyk-generated-pull-requests).

## Snyk SCM webhooks

To track pull request events, Snyk adds webhooks to your imported repositories. For more information, see the [GitHub and Git repository integrations](../../../scm-integrations/organization-level-integrations/).

Snyk uses these webhooks to:

* Track the state of Snyk pull requests: when PRs are triggered, created, updated, merged, and so on.
* Send push events to trigger PR checks.
