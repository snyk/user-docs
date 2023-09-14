# Retrieve audit logs of user-initiated activity by API for an Org or Group

{% hint style="info" %}
**Feature availability**\
Audit logs are available with Enterprise plans. See [Plans and pricing](https://snyk.io/plans/) for details.
{% endhint %}

Snyk has [Audit logs endpoints in the Snyk REST API](https://apidocs.snyk.io/?version=2023-08-24%7Ebeta#tag--Audit-Logs) that enable you to retrieve your audit logs of user-initiated activity that took place in the past 90 days. Use the [Group level audit logs endpoint](https://apidocs.snyk.io/?version=2023-08-24%7Ebeta#get-/groups/-group\_id-/audit\_logs/search) to get information about your Snyk Group, and the [Organization level audit logs endpoint](https://apidocs.snyk.io/?version=2023-08-24%7Ebeta#get-/orgs/-org\_id-/audit\_logs/search) to get information about your Snyk Organization.

You may want to retrieve audit logs for purposes such as finding out when a new user was added, analyzing unexpected activity and retroactively triaging the issues to address, or monitoring changes in a user's role to see any unusual behavior.

Examples of events returned include the following:

* A Snyk Group or Organization was added or removed, or a setting was changed
* Users were invited, added, removed, or a user's role was changed.
* The license rule or policy was modified.
* A service account was created, modified, or deleted.

For a list of the events returned, refer to the [API documentation](https://apidocs.snyk.io/?version=2023-08-24%7Ebeta#tag--Audit-Logs).

{% hint style="info" %}
Login and logout events for users are not returned because the endpoints return information for the Group or Organization specified in the request.
{% endhint %}
