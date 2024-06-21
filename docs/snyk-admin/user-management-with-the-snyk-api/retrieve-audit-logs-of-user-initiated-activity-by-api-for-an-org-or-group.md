# Retrieve audit logs of user-initiated activity by API for an Org or Group

{% hint style="info" %}
**Availability**\
Audit logs are available only for Enterprise plans.\
For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Snyk has [audit logs endpoints](https://apidocs.snyk.io/?#tag--Audit-Logs) that enable you to retrieve your audit logs of user-initiated activity that took place in the past 90 days. use [Search Group audit logs](https://apidocs.snyk.io/?#get-/groups/-group\_id-/audit\_logs/search) to get information about your Snyk Group and [Search Organization audit logs](https://apidocs.snyk.io/?version=2024-05-08%7Ebeta#get-/orgs/-org\_id-/audit\_logs/search) to get information about your Snyk Organization.

You may want to retrieve audit logs for purposes such as finding out when a new user was added, analyzing unexpected activity, and retroactively triaging the issues to address, or monitoring changes in a user's role to see any unusual behavior.

Examples of events returned include the following:

* A Snyk Group or Organization was added or removed, or a setting was changed
* Users were invited, added, removed, or a user's role was changed.
* The license rule or policy was modified.
* A service account was created, modified, or deleted.

For a list of the events returned, refer to the [API documentation](https://apidocs.snyk.io/?version=2023-08-24%7Ebeta#tag--Audit-Logs). Login and logout events for users are not returned because the endpoints return information for the Group or Organization specified in the request.
