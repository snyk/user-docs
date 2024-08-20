# Event Forwarding

{% hint style="info" %}
**Transition to Snyk Apps**

Snyk is transitioning event forwarding integrations to use the Snyk Apps platform. This change will enable new features and enhanced security across current and future Cloud Events integrations.&#x20;

During the transition, existing integrations will continue to function normally, and customers will have the opportunity to authorize the integrations to ensure they continue working once they become Snyk Apps. You can complete authorization for existing integrations by following these steps:

1. Go to the **Settings** page for your Organization
2. Go to the settings section for the integration you want to authorize, for example, Amazon EventBridge, AWS CloudTrail Lake, AWS Security Hub.
3. Click the **Authorize app** button and complete the App authorization flow

At the end of the transition window, **integrations that have not been authorized will no longer be able to forward events and will cease functioning.**
{% endhint %}

Snyk event forwarding integrations allow you to push Snyk platform events directly to certain products on other platforms, enabling you to set up custom alerting, build your own reporting, trigger automation, and more.&#x20;

## Event types

Snyk supports sending two different types of events:

1. **Snyk issue events** - these events are sent when new issues are discovered in a Snyk Project, or when an issue is updated. Each event contains information about the vulnerability or other problem found, including whether a remediation is available.&#x20;
2. **Snyk platform audit events** - these events are sent every time a Snyk user performs an action within the Snyk platform. For more information, see [Audit logs](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md).&#x20;

{% hint style="info" %}
The **Snyk issue** event type does not include Snyk Cloud issues.
{% endhint %}

{% hint style="info" %}
The **Snyk platform audit** event type is available with Snyk Enterprise plans. See [Pricing plans](../../implement-snyk/enterprise-implementation-guide/trial-limitations.md) for details.
{% endhint %}

## Supported integrations

Event forwarding integrations are available with the following products:

* [Amazon EventBridge](amazon-eventbridge.md) - **Issue events** and **Audit events**
* [AWS CloudTrail Lake](aws-cloudtrail-lake.md) - **Audit events**
* [AWS SecurityHub](aws-security-hub.md) - **Issue events**

Each integration is built on the [Snyk Apps](../../snyk-api-info/snyk-apps/) platform.
