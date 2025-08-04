# User management with the API

{% hint style="info" %}
**Feature availability**

The Snyk API is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

You can manage users through the Snyk Web UI and Snyk API. You can also [manage service accounts](../../implementation-and-setup/enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md) using the API.

* You can use the [provisioning endpoints ](provision-users-to-organizations-using-the-api.md)to organize and grant permissions under a specified role for SSO users before initial log-on.
* You must use the API to [update member roles](update-member-roles-using-the-api.md).
* You can [remove users from Groups and Organizations](remove-members-from-groups-and-orgs-using-the-api.md) programmatically using the member endpoints.
* You can retrieve [audit logs](retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) for the past 90 days using the Group-level audit logs endpoint to monitor user activity.
