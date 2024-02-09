# Automatic and manual PRs with Snyk Open Source

{% hint style="info" %}
For the basic steps in fixing vulnerabilities, see  [Fix your vulnerabilities](../manage-vulnerabilities/fix-your-vulnerabilities.md). To be sure your language is supported, see [Languages supported for Fix Pull Requests or Merge Requests](../manage-vulnerabilities/troubleshoot-fixing-open-source-vulnerabilities.md#languages-supporting-fix-pull-requests-or-merge-requests).
{% endhint %}

Snyk provides actionable fix advice for vulnerabilities in your Open Source libraries through the following:

* [Automatic pull and merge requests (PRs and MRs)](./#automatic-pull-merge-requests)
* [Manual pull and merge requests](./#manual-pull-merge-requests-for-project-code)

## **Automatic pull and merge requests**

For Projects imported through an SCM integration, Snyk offers the following types of automated pull and merge requests:

* [Fix pull requests for new vulnerabilities](automated-pull-request-creation-for-new-fixes.md)
* [Fix pull requests to clear the backlog of vulnerabilities in priority order](automated-fix-pull-requests-for-backlog-issues-and-known-vulnerabilities.md)
* [Dependency upgrade pull requests](upgrade-dependencies-with-automatic-prs.md)

## Manual pull and merge requests for Project code

Follow these steps to generate a PR or MR directly from your Project in the Snyk Web UI:

1. Navigate to your Project from the Project list
2. Select the Project.
3. Select **Open a Fix PR/MR** or **Fix this vulnerability**\
   A preview screen appears, showing you what fixes will be applied.
4. Click **Open a Fix PR** on this screen to generate the pull request.

<figure><img src="../../../.gitbook/assets/image18.png" alt="Open Fix PR to fix Open Source Project vulnerabilities"><figcaption><p>Open Fix PR to fix Open Source Project vulnerabilities</p></figcaption></figure>

{% hint style="info" %}
PRs use a branch naming convention based on the issues that they fix. When a PR already exists for a specific change, Snyk does not create a new one, even if the original PR is closed. If you try to create a Fix PR that duplicates an existing one, an error may be displayed. If this happens, check to see if the branch already exists and re-open it.
{% endhint %}

## Snyk SCM webhooks

To track pull request events, Snyk adds webhooks to your imported repositories. For more information, see the [GitHub and Git repository integrations](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

Snyk uses these webhooks to:

* Track the state of Snyk pull requests: when PRs are triggered, created, updated, merged, and so on.
* Send push events to trigger PR checks
