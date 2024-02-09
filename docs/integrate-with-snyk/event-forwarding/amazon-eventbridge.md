# Amazon EventBridge

{% hint style="info" %}
**Transition to Snyk Apps**

Snyk is currently transitioning event forwarding integrations to use the Snyk Apps platform. This change will enable new features and enhanced security across current and future Cloud Events integrations.&#x20;



During the transition, existing integrations will continue to function normally and customers will have the opportunity to authorize the integrations to ensure they continue working once they become Snyk Apps. You can complete authorization for existing integrations by following these steps:

1. Go to the **Settings** page for your organization
2. Go to the settings section for the integration you want to authorize (e.g. Amazon EventBridge, AWS CloudTrail Lake, AWS Security Hub)
3. Click the **Authorize app** button and complete the App authorization flow



At the end of the transition window, **integrations which have not been authorized will no longer be able to forward events and will cease functioning.**
{% endhint %}

The [Amazon EventBridge](https://aws.amazon.com/eventbridge/) integration sends Snyk platform events to EventBridge, allowing you to integrate Snyk events into your existing AWS environments. The integration can be configured to send two different types of events:

* **Snyk issue events** - these events are sent when new issues are discovered in a Snyk Project, or when an issue is updated. Each event contains information about the vulnerability or other problem found, including whether a remediation is available.
* **Snyk platform audit events** - these events are sent every time a Snyk user performs an action within the Snyk platform. For more information, see [Audit logs](https://docs.snyk.io/snyk-admin/manage-users-and-permissions/audit-logs). This event type is available with Snyk Enterprise plans. See [Pricing plans](../../implement-snyk/enterprise-implementation-guide/trial-limitations.md) for details

To set up the integration, there are two steps:

1. Configure an EventBridge integration in the Snyk dashboard. This will create a Snyk **Partner Event Source** in your AWS account, which you can see in the EventBridge dashboard.
2. Configure the Snyk integration in Amazon EventBridge. This step involves associating the Snyk event source created in step one with an EventBridge **Event Bus**.&#x20;

After you complete these steps, Snyk immediately starts sending events to the configured event bus.

## Configuring EventBridge in the Snyk dashboard

Navigate to [the Snyk integrations page](https://app.snyk.io/integrations) and search for **EventBridge** or navigate to the **Cloud events** section. Click on the **EventBridge** tile to start creating a new integration.

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-tile.png" alt=".Create new EventBridge integration"><figcaption><p>Create new EventBridge integration</p></figcaption></figure>

Enter a **name** for this integration, along with the **AWS Account ID** and **AWS Region** where you want to receive events. Then, select the **Event Type** you want to forward with this integration. To send more than one event type to the same account/region, create a separate integration for each event type.

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-dialog.png" alt="Enter integration details"><figcaption><p>Enter integration details</p></figcaption></figure>

When the form is completed, click **Add integration**. After this step is done, you must complete the integration set up in the Amazon EventBridge console.

### Snyk App Authorization

If this is the first time you have set up an Amazon EventBridge integration for your organization, you will be prompted to complete the Snyk App authorization flow.&#x20;

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-auth.png" alt="" width="375"><figcaption></figcaption></figure>

After completing the authorization flow you will be redirected to the settings page for the integration.&#x20;

## Configure the integration in Amazon EventBridge

After configuring the EventBridge integration on the Snyk side, you should see a new **Partner Event Source** in the EventBridge console. Go to the EventBridge console and navigate to the **Partner event sources** page under the **Integration** section.

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-eventsource.png" alt="Partner event sources"><figcaption><p>Partner event sources</p></figcaption></figure>

Snyk-generated event sources will have a naming pattern like this:

`aws.partner/snyk.io/org_<SNYK_ORG_ID>/<EVENT_TYPE>`\
\
Click on the name of the event source then click **Associate with event bus** and follow the prompts to associate the event source with an event bus. After the event source is associated with an event bus, Snyk will immediately be able to start sending events, which you can use for any actions supported by EventBridge.

## Managing and deleting an EventBridge integration

Go to the [EventBridge integration settings page](https://app.snyk.io/manage/integrations/aws-eventbridge) in the Snyk dashboard and click on the name of the integration you want to manage.

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-snyksettings.png" alt="Select Amazon EventBridge integration"><figcaption><p>Select Amazon EventBridge integration</p></figcaption></figure>

Clicking on the name of the integration opens the integration settings page, which displays configuration information for the integration.

{% hint style="info" %}
Because EventBridge integrations create an external resource that depends on the configured AWS Account ID, Region, and event type, it is not possible to edit these configuration fields. If you need to change one of these fields, delete the integration and create a new one. This deletes the existing **partner event source** in AWS and creates a new one, which you will need to associate with an **event bus** as described above.
{% endhint %}

To delete an integration, scroll to the bottom of the page and click the **Remove integration** button, then confirm the deletion.

<figure><img src="../../.gitbook/assets/integrations-eventforwarding-eventbridge-delete.png" alt="Remove integration"><figcaption><p>Remove integration</p></figcaption></figure>

This deletes the integration configuration on the Snyk side, and also deletes the **Partner Event Source** associated with this integration in AWS. You can verify that the event source has been deleted in the EventBridge console.

## Understanding event data

### Snyk issue events

This event type includes core data about Snyk issues, including:

* Vulnerability type and CVE identifiers
* Issue severity
* Whether a remediation is available

Events are JSON formatted using the [Open Cybersecurity Schema Framework _finding_](https://schema.ocsf.io/1.0.0-rc.2/classes/security\_finding?extensions=) schema.&#x20;

{% hint style="info" %}
Not all Snyk issue data is included in these events, though Snyk is continually working to provide more complete event data.
{% endhint %}

### Snyk audit events

{% hint style="info" %}
This event type is available with Snyk Enterprise plans. See [Pricing plans](../../implement-snyk/enterprise-implementation-guide/trial-limitations.md) for details.
{% endhint %}

This event type forwards Snyk platform audit logs with the same schema defined as part of the Snyk [v1 Audit Log API](https://snyk.docs.apiary.io/#reference/audit-logs). &#x20;
