# Event Forwarding

Snyk event forwarding integrations allow you to push Snyk platform events directly to certain products on other platforms, enabling you to set up custom alerting, build your own reporting, trigger automation, and more.&#x20;

## Event types

Snyk supports sending two different types of events:

1. **Snyk issue events** - these events are sent when new issues are discovered in a Snyk Project, or when an issue is updated. Each event contains information about the vulnerability or other problem found, including whether a remediation is available.&#x20;
2. **Snyk platform audit events** - these events are sent every time a Snyk user performs an action within the Snyk platform. For more information, see [Audit logs](../../snyk-platform-administration/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md).&#x20;

{% hint style="info" %}
The **Snyk issue** event type does not include Snyk Cloud issues.
{% endhint %}

{% hint style="info" %}
The **Snyk platform audit** event type is available with Snyk Enterprise plans. See [Pricing plans](../../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md) for details.
{% endhint %}

## Supported integrations

Event forwarding integrations are available with the following products:

* [Amazon EventBridge](amazon-eventbridge.md) - **Issue events** and **Audit events**
* [AWS CloudTrail Lake](aws-cloudtrail-lake.md) - **Audit events**
* [AWS SecurityHub](aws-security-hub.md) - **Issue events**
* [Google Security Command Center](google-security-command-center.md) **- Issue events**
