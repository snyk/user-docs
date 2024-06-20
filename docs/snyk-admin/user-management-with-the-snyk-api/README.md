# User management with the Snyk API

{% hint style="info" %}
**Availability**\
Snyk API is available only for Enterprise plans.\
For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Users can be managed using the Snyk [API v1](https://snyk.docs.apiary.io) and [REST](https://apidocs.snyk.io/?version=2024-01-04#overview) API and the [Snyk Web UI](../manage-permissions-and-roles/user-role-management.md). You can also [manage service accounts](../../enterprise-configuration/service-accounts/manage-service-accounts-using-the-snyk-api.md) using the Snyk REST API.

[User provisioning ](provision-users-to-orgs-using-the-snyk-api-v1.md)uses the provisioning endpoints to organize and grant permissions under a specified role for SSO users before initial log-on.

Users can be programmatically [removed from Groups and Organizations](remove-members-from-groups-and-orgs-using-the-snyk-rest-and-v1-api.md) using the member endpoint.

[Audit logs](retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) can be retrieved for the past 90 days using the Group-level audit logs endpoint to monitor user activity.
