# Audit logs

{% hint style="info" %}
**Feature availability**\
Audit logs are available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Introduction

Snyk has [Audit logs endpoints](https://snyk.docs.apiary.io/#reference/audit-logs) that enable you to retrieve information from your audit logs using the [Snyk API v1](../../snyk-api-info/).

For example, you can help retroactively triage any unexpected activity, find out when a new user is added, or monitor changes in a user's role to get early warning of any unusual behavior.

Use the Group level audit logs endpoint to get information about your Snyk Group, and the Organization level audit logs to get information about your Snyk Organization.

## Examples

Examples of events returned include:

* A Snyk Group or Organization was added or removed or a setting was changed
* Users were invited, added, removed, or a user's role was changed
* The license rule or policy was modified
* A service account was created, modified, or deleted

For a list of the events returned, refer to the request body JSON schema for the endpoint you are using.

{% hint style="info" %}
Logon and logout events for users are not returned, because the endpoints return information for the Group or Organization specified in the request.
{% endhint %}

## Limitations of audit logs

If you need audit log information for an earlier time, or information such as login and logoff events for a user, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new). The support team can retrieve information from storage on your behalf.
