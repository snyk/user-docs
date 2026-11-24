---
description: Find which Snyk role grants a permission in a Snyk-managed account
nav_context: agnostic
last_reviewed: 2026-11-19
last_reviewed_source: reviewed
---

# Map Snyk roles to Snyk API & Web permissions

{% hint style="info" %}
This page applies to Snyk-managed accounts. For details, visit [Understand Snyk-managed accounts](snyk-managed-accounts.md).
{% endhint %}

In a Snyk-managed account, you do not assign roles in Snyk API & Web. Snyk API & Web reads each user's Snyk roles and grants the permissions listed on this page. Use this page to choose the Snyk role a user needs.

## Where each role applies

| Snyk role level | Applies to                                                           |
| --------------- | -------------------------------------------------------------------- |
| Tenant          | Your account, such as account settings, the audit log, and Discovery |
| Group           | Every Organization in the Group                                      |
| Organization    | That Organization only                                               |

If a user has both a Group role and an Organization role for the same Organization, the role that grants more permissions applies. A Tenant role does not grant permissions in Organizations.

## Permissions granted by built-in roles

A ✓ means the role grants the permission. For what each permission allows, visit [Understanding permissions](understanding-permissions.md).

| Permission | Tenant Admin | Tenant Viewer | Tenant Member | Group Admin | Group Viewer | Group Member | Organization Admin | Organization Collaborator |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Account Settings | ✓ | | | | | | | |
| Audit Log | ✓ | ✓ | ✓ | | | | | |
| Discovery | ✓ | | | | | | | |
| Discovery Read-Only | ✓ | ✓ | ✓ | | | | | |
| Scanning Agent Management | ✓ | | | | | | | |
| View Target | | | | ✓ | ✓ | | ✓ | ✓ |
| Create Target | | | | ✓ | | | ✓ | ✓ |
| Delete Target | | | | ✓ | | | ✓ | ✓ |
| Change Finding | | | | ✓ | ✓ | | ✓ | ✓ |
| Change Finding State | | | | ✓ | | | ✓ | ✓ |
| Change risk | | | | ✓ | | | ✓ | ✓ |
| Change Target Settings | | | | ✓ | | | ✓ | |
| Manage Credentials | | | | ✓ | | | ✓ | |
| Correlation Admin | | | | ✓ | | | ✓ | ✓ |
| Correlation Viewer | | | | ✓ | ✓ | | ✓ | ✓ |
| Start Scan | | | | ✓ | | | ✓ | ✓ |
| Start re-test | | | | ✓ | | | ✓ | ✓ |

The Group Member role lets a user see the Organization but grants no permissions in Snyk API & Web. Tenant Admin can also manage domains, and Tenant Admin and Group Admin can move targets between Organizations.

## Custom roles

Snyk API & Web supports the custom roles you create in Snyk. A custom role grants the permissions that match what it allows in Snyk. For example, a custom role that can create Projects in Snyk grants Create Target. If a user holds a built-in role and a custom role for the same Organization, the one that grants more applies.

## When a role change takes effect

A change to a user's Tenant role applies the next time they log in. A change to an Organization or Group role can take up to 5 minutes to refresh.

## Related information

* [Understand Snyk-managed accounts](snyk-managed-accounts.md)
* [Understanding permissions](understanding-permissions.md)
* [Select a Group and an Organization](select-a-group-and-an-organization.md)
* [Snyk pre-defined roles](https://docs.snyk.io/platform-administration/user-management/pre-defined-roles)
