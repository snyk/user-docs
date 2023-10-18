# Automatic and manual PRs with Snyk Open Source

{% hint style="info" %}
See [Fix your vulnerabilities](../../../scan-applications/snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md) for more details of general fix functions.
{% endhint %}

Snyk provides actionable fix advice for vulnerabilities in your Open Source libraries, using:

* [Automatic pull / merge requests (PRs / MRs)](./#automatic-pull-merge-requests).
* [Manual pull / merge requests](./#manual-pull-merge-requests-for-project-code).

{% hint style="info" %}
See [Language support for Fix Pull Request and Merge Requests](../starting-to-fix-vulnerabilities/troubleshooting-open-source-vulnerability-fixing.md#languages-supporting-fix-pull-requests-or-merge-requests).
{% endhint %}

## **Automatic pull / merge requests**

For Projects imported via an SCM (Source Code Manager), Snyk offers the following types of automated pull / merge requests:

* [Fix pull requests for new vulnerabilities](fix-pull-requests-for-new-vulnerabilities.md)
* [Fix pull requests to clear the backlog of vulnerabilities in priority order](fix-pull-requests-for-known-vulnerabilities-backlog.md)
* [Dependency upgrade pull requests](upgrading-dependencies-with-automatic-prs.md)

## Manual pull / merge requests for Project code

To generate a PR / MR directly from your Project using the Snyk UI:

1. Navigate to your Project from the Project list
2. Select the file
3. Select **Open a Fix PR/MR** or **Fix this vulnerability**
4. A preview screen appears, showing you what fixes will be applied
5. Click **Open a Fix PR** on this screen to generate the pull request

<figure><img src="../../../.gitbook/assets/image18.png" alt="Open Fix PR to fix Open Source Project vulnerabilities"><figcaption><p>Open Fix PR to fix Open Source Project vulnerabilities</p></figcaption></figure>

{% hint style="warning" %}
PRs use a branch naming convention based on the issues that they fix. When a PR already exists for a specific change, Snyk doesn't create a new one, even if that original PR is closed.\
If you try to create a Fix PR which duplicates an existing one, an error may be displayed. If this happens, check if the branch already exists and re-open it.
{% endhint %}

## Snyk SCM webhooks

To track pull request events, Snyk adds webhooks to your imported repositories (see [GitHub and Git repository integrations](../../../integrations/git-repository-scm-integrations/)).

Snyk uses these webhooks to:

* Track the state of Snyk pull requests: when PRs are created, updated, triggered, merged, and so on.
* Send push events to trigger PR checks
