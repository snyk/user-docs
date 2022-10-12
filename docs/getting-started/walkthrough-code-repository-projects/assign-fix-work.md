# Assign fix work

{% hint style="info" %}
**Recap**\
You have understood and fixed a vulnerability. Now we'll see how this process works in your team for your applications.
{% endhint %}

### Decide what to fix first

Your fix priority (deciding what to fix first) depends on the workflows and business processes of your teams. This decision-making process for fixing vulnerabilities may be for individual developers to decide on an ad-hoc basis, or may be for team leads to assign work as part of a Sprint.

Different teams approach fixes in different ways, depending on the tools they use, their own workflow maturity, and competing work priorities. Typically, smaller teams have less process, and Enterprise-level teams are more formal.

#### Example triage process

For example, your team could follow a triage-based process for each issue, driven largely by the severity of the issue:

![](<../../.gitbook/assets/image (110).png>)

### Example workflow: team-lead driven, using Jira

{% hint style="info" %}
**Feature availability**\
Jira integration is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Some teams base all their work around Jira tasks; we’ll look at this as an example.

Let’s assume that your development team assigns fix work based on Sprints, and decides that the next development Sprint will be dedicated to vulnerability fixing.

As part of this Sprint planning, the team leader can:

* Review the vulnerabilities in a project
* Decide which vulnerabilities to fix
* Create a Jira issue for each vulnerability.
* Assign these Jira issues as task to developers to fix these vulnerabilities,
* Track progress on these tasks during the Sprint.

Snyk [Jira integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira) allows you to run this process from the Snyk Web UI.

#### Assign a Jira issue

Navigate to the issue you have decided to fix, then click **Create a Jira issue**:

![](<../../.gitbook/assets/image (221) (1) (2) (1).png>)

You can then define the Jira task details for this fix:

![](<../../.gitbook/assets/image (297).png>)

You can assign this task to a developer in the team, following your team’s normal Sprint processes.

{% hint style="info" %}
You may want to create a Jira issue even if Snyk knows how to fix the change, and even if it’s a very minor upgrade; assigning issues allows your team to manage, justify and track code changes.
{% endhint %}

#### More information

* See [Jira issues](https://training.snyk.io/learn/course/introduction-to-the-snyk-ui/issue-fix-options/open-source-fix-advice?page=2) training for more details.
* See [Fixing and Prioritizing issues](https://docs.snyk.io/features/fixing-and-prioritizing-issues) docs for more general details.

### What's next?

Now, you can look at [using Snyk Reports to manage team work](manage-team-work-using-reports.md).
