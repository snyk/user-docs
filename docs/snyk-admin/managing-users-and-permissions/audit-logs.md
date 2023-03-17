# Audit logs

{% hint style="info" %}
**Feature availability**\
Audit logs are available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Introduction

Snyk has [Audit logs endpoints in the Snyk API v1](https://snyk.docs.apiary.io/#reference/audit-logs) that enable you to retrieve your audit logs of user-initiated activity that took place in the past 90 days.

You may want to retrieve audit logs for purposes such as helping to retroactively triage any unexpected activity, finding out when a new user is added, or monitoring changes in a user's role to get early warning of any unusual behavior.

Use the [Group level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs) endpoint to get information about your Snyk Group, and the [Organization level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/organization-level-audit-logs) endpoint to get information about your Snyk Organization.

## Examples of events returned in audit logs

Examples of events returned include:

* A Snyk Group or Organization was added or removed or a setting was changed
* Users were invited, added, removed, or a user's role was changed
* The license rule or policy was modified
* A service account was created, modified, or deleted

For a list of the events returned, refer to the request body JSON schema for the endpoint you are using.

{% hint style="info" %}
Logon and logout events for users are not returned, because the endpoints return information for the Group or Organization specified in the request.
{% endhint %}
