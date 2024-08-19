# User management with the Snyk API

{% hint style="info" %}
**Availability**\
The Snyk API is available only for Enterprise plans.\
For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

You can manage users through the Snyk Web UI, and also the Snyk API [V1](https://snyk.docs.apiary.io) and [REST](https://apidocs.snyk.io/?version=2024-01-04#overview). In addition, you can [manage service accounts](../../enterprise-configuration/service-accounts/manage-service-accounts-using-the-snyk-api.md) using the Snyk REST API.

* You can use the [provisioning endpoints ](provision-users-to-organizations-using-the-snyk-api.md)to organize and grant permissions under a specified role for SSO users before initial log-on.
* You must use the API to [update member roles](update-member-roles-using-the-v1-api.md).
* You can [remove users from Groups and Organizations](remove-members-from-groups-and-orgs-using-the-api.md) programmatically using the member endpoints.
* You can retrieve [audit logs](retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) for the past 90 days using the Group-level audit logs endpoint to monitor user activity.
