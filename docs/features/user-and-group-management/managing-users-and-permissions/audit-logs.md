# Audit logs

Snyk API v1 has [Audit logs endpoints](https://snyk.docs.apiary.io/#reference/audit-logs) that enable you to retrieve information from your audit logs using the Snyk API. Use the Group level audit logs endpoint to get information about your Snyk Group, and the Organization level audit logs to get information about your Snyk Organization.  Examples of events returned include:

* A Snyk Group or Organization was added or removed or a setting was changed
* Users were invited, added, removed, or a user's role was changed
* The license rule or policy was modified
* A service account was created, modified, or deleted

For a list of the events returned refer to the request body JSON schema for the endpoint you are using. Note that logon and logout events for users are not returned because the endpoints return information for the Group or Organization specified in the request.

You can retrieve user-initiated activity for the past 90 days. If you need audit log information for an earlier time, or information such as login and logoff events for a user, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new). The support team can retrieve information from storage on your behalf.

Use the Audit logs endpoints to help retroactively triage any unexpected activity, to find out whenever a new user is added so you can assist with onboarding, or to monitor changes in a user's role in order to get early warning of any unusual behavior.report.

{% hint style="info" %}
**Feature availability**\
The Snyk API is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}
