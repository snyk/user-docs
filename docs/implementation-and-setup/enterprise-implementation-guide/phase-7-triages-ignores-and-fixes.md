# Phase 7: Triages, ignores, and fixes

## Tips for deciding on prioritization

After you implement a strategy to prevent new issues from entering your repositories, whether blocking builds or running in a non-blocking, advisory mode, the next step is to prioritize and start fixing issues in your backlog.

* In [Phase 4: Create a Fix strategy](phase-4-create-a-fix-strategy.md), you created a plan for prioritizing your Projects and issues. To implement this, you can schedule regular meetings with development team leads to assist them with this process.
* In the [Snyk Tools](../../scan-with-snyk/snyk-tools/) section, there is a tool called [jira-tickets-for-new-vulns](../../scan-with-snyk/snyk-tools/tool-jira-tickets-for-new-vulns.md)**,** which can be scheduled to run on a regular basis to automatically create Jira tickets for vulnerabilities that meet your specified criteria. Whether this process is automated or not, creating tickets for your developers to review can be a great way to help make issues identified by Snyk more visible.
* If you use Jira Cloud, you can download and install the [Snyk Security in Jira Cloud](https://marketplace.atlassian.com/apps/1230482/snyk-security-in-jira-cloud) plugin from the Atlassian marketplace. This allows you to view information on your Snyk Vulnerabilities directly in Jira, and use Jira Automation to create new tickets when new vulnerabilities are identified.

## Use the ignore feature

### When should you ignore an issue?

When deciding your priority for fixing issues, you may see specific packages or vulnerabilities that you do not currently want to fix. This could be for a range of reasons, such as:

* The fix introduces breaking changes, and you do not have time to fix them.
* This vulnerability does not apply to you for an environmental or contextual reason.

In each case, you can use the ignore feature to stop these issues from appearing each time you run a test.

### Implement the ignore feature

{% hint style="info" %}
Confirm an ignore with a Group Admin or Organization Admin. They may need to complete this step themselves.
{% endhint %}

When adding an ignore:

* Ensure you add a detailed reason, so the ignore reason is clear to others who see this issue.
* Set an expiration date for the ignore, rather than having a permanent ignore. This is essential, as even if the issue may not be fixable or relevant today, it should be reviewed regularly, monthly or quarterly, to see if it is possible to implement a fix.

{% hint style="info" %}
It is common to use the General Settings to limit access for users to permission to ignore an issue and to require a reason.
{% endhint %}

By default, the **Organization Collaborator** role has permission to ignore issues, but this can be controlled for each Organization on the settings page, that is, restricted to **Organization admins** only.

See [Ignore issues](../../manage-risk/prioritize-issues-for-fixing/ignore-issues/) for more details.

## Review Snyk reports to track the adoption of Snyk across your teams

There are a number of different reports in Snyk to help you get an overview of your issues and vulnerabilities. For more information, see [Reports tab](../../manage-risk/analytics/reports-tab/).

* On the **Issues Summary** report, the **Risk Breakdown** section displays open, new, and resolved issues across different Organizations. Tracking the amount of activity in each Organization can help you identify which teams are adopting the tool most actively.
* Individual Organization administrators can also view reports focused on their Organization as a way to help identify which vulnerabilities are most common across their repositories, and also to track issues resolved in different Projects.
