# AWS CloudTrail Lake

{% hint style="info" %}
**Feature availability**\
\*\*\*\*The AWS CloudTrail Lake integration feature is available with Snyk Enterprise plans. See [Pricing plans](../../more-info/plans.md) for details.

This feature is currently available in the Snyk MT-US environment. The feature will be rolled out to other Snyk Multi-tenant environments soon. If you are a Single-tenant customer, please contact your Snyk representative to enable the feature.
{% endhint %}

The AWS CloudTrail Lake integration allows you to forward [Snyk audit logs](https://docs.snyk.io/user-and-group-management/managing-users-and-permissions/audit-logs) to AWS CloudTrail Lake, which lets you run SQL-based queries on your logs and retain them for up to seven (7) years.

In order to enable the AWS CloudTrail Lake integration you must do the following:

1. Add a Snyk integration in AWS CloudTrail Lake.
2. Configure the integration in Snyk.

{% hint style="info" %}
This integration sends logs beginning at the time you enable it. Logs generated before enabling the integration are not sent, but may be available from Snyk API v1 [Get organization level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/organization-level-audit-logs/get-organization-level-audit-logs).
{% endhint %}

## Add a Snyk integration in AWS CloudTrail Lake

Follow the setup [instructions](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-integration.html) in the AWS CloudTrail Lake documentation, choosing Snyk as the integration type.

<figure><img src="../../.gitbook/assets/aws-ctl-1 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Choose Snyk Add integration for Snyk"><figcaption><p>Choose Snyk Add integration for Snyk</p></figcaption></figure>

During the setup, supply an **External ID** for the integration. Set this field to your Snyk **Organization ID**, which you can find by going to the Snyk [settings page](https://app.snyk.io/manage/settings).

<figure><img src="../../.gitbook/assets/aws-ctl-2 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Organization ID on Snyk settings page"><figcaption><p>Organization ID on Snyk settings page</p></figcaption></figure>

Copy your Snyk Organization ID to the **External ID** field in the AWS CloudTrail Lake integration setup and continue following the instructions in the AWS CloudTrail Lake documentation.

When you are done creating the Snyk integration in AWS CloudTrail Lake, copy the **Channel ARN** that is displayed on the integration page. You will need this for the next step.

## Configure the integration in Snyk

After creating the integration in **AWS CloudTrail Lake**, configure the integration on the Snyk side so that Snyk can start sending logs.

To do this, go to [the Snyk integrations page](https://app.snyk.io/integrations), navigate to **Cloud events**, and click the **AWS CloudTrail Lake** tile:

<figure><img src="../../.gitbook/assets/aws-ctl-3 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="CloudTrail Lake tile on Snyk integrations page"><figcaption><p>CloudTrail Lake tile on Snyk integrations page</p></figcaption></figure>

Enter a **name** for this integration, your **AWS Account ID**, and the **Channel ARN** from the previous step.

<figure><img src="../../.gitbook/assets/aws-ctl-4 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Integration name, AWS Account ID, Channel ARN"><figcaption><p>Integration name, AWS Account ID, Channel ARN</p></figcaption></figure>

Once this step is complete, Snyk will begin forwarding audit logs to AWS CloudTrail Lake immediately. You can click View settings or go to the [AWS CloudTrail Lake settings](https://app.snyk.io/manage/integrations/aws-cloudtrail) page to view and manage the integration.

## Remove an AWS CloudTrail Lake integration

\
Go to the [AWS CloudTrail Lake settings](https://app.snyk.io/manage/integrations/aws-cloudtrail) page and select the name of the integration you want to remove.

<figure><img src="../../.gitbook/assets/aws-ctl-5 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Select AWL CloudTrail Lake integration to remove"><figcaption><p>Select AWL CloudTrail Lake integration to remove</p></figcaption></figure>

Select **Remove integration** and confirm that you want to remove the integration.

<figure><img src="../../.gitbook/assets/aws-ctl-6 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Remove integration button"><figcaption><p>Remove integration button</p></figcaption></figure>

This action removes Snyk’s configuration for this integration, which will **prevent any further audit logs from being sent to AWS CloudTrail Lake**. This does not remove the Snyk integration in AWS CloudTrail Lake. To do this, go to AWS CloudTrail Lake and delete the Snyk integration from the **Integration** list.

## Query Snyk audit logs in AWS CloudTrail Lake

Once your Snyk audit logs are being forwarded to AWS CloudTrail Lake, you can access them with the AWS CloudTrail Lake **Query** functionality. You can use this example query to get started:\\

```sql
select 
    eventtime, 
    eventdata.useridentity, 
    eventdata.eventname,
    eventdata.additionaleventdata
from <EVENT-DATA-STORE-ID>
order by eventTime desc
limit 10
```

Replace `<EVENT-DATA-STORE-ID>` with the ID of the event data store that is associated with the Snyk integration in AWS CloudTrail Lake.

## Understanding the log data

There are three (3) key fields to note when using the Snyk audit log data in AWS CloudTrail Lake.

`eventdata.useridentity`

The event `useridentity` contains a field called `principalid`, which represents the Snyk user ID for the user associated with the audit event. You can use Snyk API v1 [Get organization level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/organization-level-audit-logs/get-organization-level-audit-logs) to match the Snyk user ID with a user in your organization.

`eventdata.eventname`

This represents the type of audit event (for example, `api.access` or `org.cloud_config.settings.edit`) and can be used to group or filter events.

`eventdata.additionaleventdata`

This field contains a raw JSON payload with more detailed information about the audit event. The content of the payload depends on the type of the event. For example, an API access event will include the accessed URL, while a settings change event will include before and after values for the changed setting.
