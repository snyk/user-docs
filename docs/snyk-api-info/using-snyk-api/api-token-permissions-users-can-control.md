# Snyk API token permissions users can control

To set an API token to be read-only and unable to write to the platform, use a service account and set it to Group Viewer. **Note:** The [Get group audit logs endpoint](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs) requires Group Admin permissions.

Service accounts at the org level have only org admin and org collaborator permissions. Thus to set a service account to view-only you must use a group level service account.

For more information see [Service accounts](../../enterprise-setup/service-accounts/).
