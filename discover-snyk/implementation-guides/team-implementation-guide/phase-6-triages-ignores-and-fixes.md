---
description: 'Phase 6 of the Snyk team implementation: triage, prioritize, ignore, and fix vulnerabilities in your backlog'
nav_context: agnostic
---

# Phase 6: Triages, ignores, and fixes

## Tips for deciding on prioritization

After you implement a strategy to prevent new issues from entering your repositories, whether blocking builds or running in a non-blocking/advisory mode, the next step is to prioritize and start fixing issues in your backlog.

* In [Phase 4: Create a fix strategy](phase-4-create-a-fix-strategy.md), you created a plan for prioritizing your Projects and issues. To implement this plan, you can schedule regular meetings with development team leads to assist them with this process.
* If you use Jira Cloud, you can download and install the [Snyk Security in Jira Cloud](https://marketplace.atlassian.com/apps/1230482/snyk-security-in-jira-cloud) plugin from the Atlassian Marketplace. You can then view information on your Snyk vulnerabilities directly in Jira and use Jira Automation to create new tickets when new vulnerabilities are identified.

## When should you ignore an issue?

When deciding your priority for fixing issues, you can find specific packages or vulnerabilities that you do not want to fix now. The reasons vary, for example:

* The fix introduces breaking changes, and you do not have time to fix them.
* This vulnerability does not apply to you for an environmental/contextual reason.

## Use the ignore feature

In these cases, use the ignore feature to stop these issues from appearing each time you run a test.

{% hint style="info" %}
Confirm the ignore with an Organization Admin. If only Organization Admins can ignore issues in your Organization, the Organization Admin must complete this step.
{% endhint %}

When adding the ignore:

* Add a detailed reason so that others who see this issue understand why it is ignored.
* Set an expiration date for the ignore rather than making it permanent. This is essential: even if the issue is not fixable or relevant today, review it regularly (monthly or quarterly) to check whether a fix is now possible.

{% hint style="info" %}
In **Settings** > **General**, it is common to limit who can ignore an issue and to require a reason.
{% endhint %}

By default, the **Organization Collaborator** role has permission to ignore issues. You can restrict this permission to **Organization Admins** only on the **Settings** page of each Organization.

For more details, visit [Ignore issues](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/ignore-issues).
