# Retrieve audit logs of user-initiated activity by API for an Org or Group

{% hint style="info" %}
**Feature availability**\
Audit logs are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Snyk has [audit logs](../../snyk-api/reference/audit-logs.md) endpoints that enable you to retrieve your audit logs of user-initiated activity that took place in the past 90 days. Use the endpoints [Search Group audit logs](../../snyk-api/reference/audit-logs.md#groups-group_id-audit_logs-search) to get information about your Snyk Group and [Search Organization audit logs](../../snyk-api/reference/audit-logs.md#orgs-org_id-audit_logs-search) to get information about your Snyk Organization.

You may want to retrieve audit logs for purposes such as finding out when a new user was added, analyzing unexpected activity and retroactively triaging the issues to address, or monitoring changes in a user's role to see any unusual behavior.

Examples of events returned include the following:

* A Snyk Group or Organization was added or removed, or a setting was changed
* Users were invited, added, removed, or a user's role was changed
* The license rule or policy was modified
* A service account was created, modified, or deleted

For a list of the events returned, refer to the [API Audit logs documentation](../../snyk-api/reference/audit-logs.md). Login and logout events for users are not returned because the endpoints return information for the Group or Organization specified in the request.
