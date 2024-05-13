# Prioritize your issues

Snyk has several features that help you determine which issues you discover are the most important for you to fix and the sequence in which to fix the issues.

{% hint style="info" %}
For information on how to ignore and exclude issues, see [The .snyk file](../policies/the-.snyk-file.md) and the [Policies ](../policies/)pages that explain how to create policies and assign them to Projects, as well as Security and License policies.
{% endhint %}

Some tools use only the single factor of severity to prioritize issues, but this can result in thousands of results, with no clear starting point for fixing these issues.

You can prioritize at the Project level when looking at a specific Project. Enterprise customers can prioritize across all Projects.

Snyk Priority Score and Risk Score rank the [severity](severity-levels.md) of an issue and the urgency of fixing it. For details, see [Priority Score vs Risk Score](priority-score-vs-risk-score.md), [Priority Score](../../scan-with-snyk/find-and-manage-priority-issues/priority-score.md), and [Risk Score](risk-score.md).

You can [ignore issues](ignore-issues/) and [triage issues](triage-for-issues.md) to establish your issue management strategy.

[View exploits](view-exploits.md) to see how vulnerabilities can be taken advantage of. You can then start evaluating and prioritizing vulnerabilities using guidance from the [Snyk Priority Score](../../scan-with-snyk/find-and-manage-priority-issues/priority-score.md) for each issue.

Consider [Malicious packages](malicious-packages.md) and how to address them in your Projects.

You can set up [reachable vulnerability analysis ](reachable-vulnerabilities.md)to identify vulnerabilities with a path to your code. This helps you asse are calculated as part of the priority score.

[Vulnerabilities with Social Trends](vulnerabilities-with-social-trends.md) are calculated as part of the Priority Score.

Based on your priorities, you can start [fixing vulnerabilities](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/).

See [Prioritize issues in the Snyk Web UI](https://learn.snyk.io/lesson/prioritize-issues-snyk) to learn about prioritization in action.

You can use many features of [Snyk Projects](../../snyk-admin/snyk-projects/) help you to focus on priority issues:

* [View Project information](../../snyk-admin/snyk-projects/project-information.md).
* Apply and remove [Project attributes ](../../snyk-admin/snyk-projects/project-attributes.md)and [Project tags](../../snyk-admin/introduction-to-snyk-projects/project-tags.md) to characterize Projects.
* Look at [Project collections groupings](../../snyk-admin/snyk-projects/project-collections-groupings/).
* [View Project issues, fixes, and dependencies](../../snyk-admin/snyk-projects/view-project-issues-fixes-and-dependencies.md).
* [View Project settings](../../snyk-admin/snyk-projects/view-and-edit-project-settings.md).



