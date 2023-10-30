# Phase 7: Triages, Ignores and Fixes

## Tips for deciding on prioritization

After you implement a strategy to prevent new issues from entering your repositories, whether blocking builds or running in a non-blocking/advisory mode, the next step is to prioritize and start fixing issues in your backlog.

* In [Phase 4: Create a Fix strategy](phase-4-create-a-fix-strategy.md), you created a plan for prioritizing your Projects and issues. To implement this, you can schedule regular meetings with development team leads, to assist them with this process.&#x20;
* In the [Snyk Tools](../../snyk-api-info/other-tools/) section, there is a tool called **jira-tickets-for-new-vulns** which can be scheduled to run on a regular basis to automatically create Jira tickets for vulnerabilities that meet your specified criteria. Whether this process is automated or not, creating tickets for your developers to review can be a great way to help make issues identified by Snyk more visible.
* If you use Jira Cloud, you can download and install the [Snyk Security in Jira Cloud](https://marketplace.atlassian.com/apps/1230482/snyk-security-in-jira-cloud) plugin from the Atlassian marketplace. This allows you to view information on your Snyk Vulnerabilities directly in Jira, and use Jira Automation to create new tickets when new vulnerabilities are identified.

## When should we ignore an issue?

When deciding your priority for fixing issues, you may see specific packages or vulnerabilities that you do not currently want to fix. This could be for a range of reasons, such as:

* The fix introduces breaking changes, and you don't have time to fix them.
* This vulnerability doesn't apply to you for an environmental/contextual reason.

## Use the ignore feature

In each case, you can use the 'ignore' feature to stop these from appearing each time you run a test.&#x20;

{% hint style="info" %}
Confirm ignore with a Group Admin or Organization Admin (they may need to complete this step themselves).&#x20;
{% endhint %}

When adding the ignore:

* Ensure you add a detailed reason, so the ignore reason is clear to others who see this issue.
* Set an expiration date for the ignore, rather than having a permanent ignore. This is essential, as whilst the issue may not be fixable/relevant today, it should be reviewed regularly (monthly or quarterly) to see if it is possible to implement a fix.

{% hint style="info" %}
In **Settings -> General** it's common to limit access to who can ignore an issue and require a reason.
{% endhint %}

By default, the **Organization Collaborator** role has permission to ignore issues, but this can be controlled per-Organization in the settings page (that is, restricted to **Organization admins** only).&#x20;

See [Ignore issues](../../manage-risk/find-and-manage-priority-issues/ignore-issues.md) for more details.

## Review Snyk reports to track adoption of Snyk across your teams

[Reporting](../../manage-issues/reporting/): there are a number of different reports in Snyk to help to get an overview of your issues and vulnerabilities.&#x20;

* On the **Issues Summary** report, the **Risk Breakdown** section displays open/new/resolved issues across your different Organizations. By tracking the amount of activity in each Organization, it can help you to identify which teams are adopting the tool most actively.
* Individual Organization Administrators can also view reports focused on their Organization, as a way to help identify which vulnerabilities are most common across their repositories, and also track issues resolved in different Projects.
