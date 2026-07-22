# Rule Extensions permissions

Rule Extensions are available to Enterprise customers. Access is granted through a [custom role](https://docs.snyk.io/snyk-platform-administration/user-roles/user-role-management). Rule Extensions are **managed at the Group level** and **tested at the Organization level**, so a role combines Group-level and Organization-level permissions.

Grant only the permissions a role needs. The set differs slightly depending on whether the role is used through the **API** or the **in-product UI**.

## Permissions for API access

A service-account role used with the Rule Extensions REST API needs only the Rule Extensions permissions:

| Permission | Slug | Grants |
| --- | --- | --- |
| View Rule Extensions | `group.rule_extension.read` | View Rule Extensions in the Group. Required by Create, Edit, and Delete. |
| Create Rule Extensions | `group.rule_extension.create` | Create Rule Extensions in the Group. |
| Edit Rule Extensions | `group.rule_extension.edit` | Update Rule Extensions in the Group. |
| Delete Rule Extensions | `group.rule_extension.delete` | Permanently delete Rule Extensions from the Group. |
| Test Rule Extensions | `org.rule_extension.project.test` | Run an impact test on a Project in the Organization. |

To use the API, create a service account that holds this role and authenticate with its token.

## Permissions for UI access

Reaching the Rule Extensions screens in the Snyk Web UI (**Group settings → Snyk Code**) additionally requires permission to view the Group and its Organizations:

| Permission | Slug | Grants |
| --- | --- | --- |
| View Group | `group.read` | View details of the Group. |
| View Group Settings | `group.settings.read` | View the Group's settings. |
| View Organizations | `group.org.list` | View Organizations in the Group. |
| View Rule Extensions | `group.rule_extension.read` | View Rule Extensions in the Group. |
| Create Rule Extensions | `group.rule_extension.create` | Create Rule Extensions in the Group. |
| Edit Rule Extensions | `group.rule_extension.edit` | Update Rule Extensions in the Group. |
| Delete Rule Extensions | `group.rule_extension.delete` | Permanently delete Rule Extensions from the Group. |
| Test Rule Extensions | `org.rule_extension.project.test` | Run an impact test on a Project in the Organization. |

## Create the custom role

1. Go to **Group settings → Member Roles → Create new role** and choose the Group role type.
2. Add the Group-level and Organization-level permissions above that the role needs.
3. Assign the role to the member (for UI access) or service account (for API access) that manages Rule Extensions.

{% hint style="info" %}
Customers who managed access during the closed beta must add the updated Rule Extensions permissions to their custom roles — Snyk does not backfill them. Snyk removes the deprecated SAST Rule Extensions permissions from the UI after a 30-day grace period.
{% endhint %}
