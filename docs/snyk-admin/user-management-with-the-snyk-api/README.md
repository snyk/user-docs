# User management with the Snyk API

{% hint style="info" %}
**Feature availability**

The Snyk API is available to Enterprise customers. See [Plans and pricing](https://snyk.io/plans/) for details.
{% endhint %}

{% hint style="warning" %}
You can manage service accounts using the [Snyk REST API](https://apidocs.snyk.io/?version=2024-01-04#tag--ServiceAccounts).
{% endhint %}

Users can be managed using the Snyk [API v1](https://snyk.docs.apiary.io) and [REST](https://apidocs.snyk.io/?version=2024-01-04#overview) API and the [Snyk Web UI](../manage-permissions-and-roles/manage-user-roles.md).

[User provisioning ](provision-users-to-orgs-using-the-snyk-api-v1.md)uses the provisioning endpoints to organize and grant permissions under a specified role for SSO users before initial log-on.

Users can be programmatically [removed from Groups and Organizations](remove-members-from-groups-and-orgs-using-the-snyk-rest-and-v1-api.md) using the member endpoint.

[Audit logs](retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) can be retrieved for the past 90 days using the Group-level audit logs endpoint to monitor user activity.
