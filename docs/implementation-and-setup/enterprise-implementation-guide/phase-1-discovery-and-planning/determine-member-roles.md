# Determine user roles

## Use default or custom roles?

A key consideration when setting up Snyk is determining if the [pre-defined user roles](../../../snyk-platform-administration/user-roles/pre-defined-roles.md) align with your needs or if you need to customize roles for your business requirements.

{% hint style="info" %}
The pre-defined user roles in Snyk have a fixed set of permissions that cannot be changed. One of the biggest advantages of customizing roles is that you can adjust the role permissions to align with your desired level of control.&#x20;
{% endhint %}

## Pre-defined roles

The following are the standard pre-defined user roles. See [Pre-defined roles](../../../snyk-platform-administration/user-roles/pre-defined-roles.md) for more details.

* **Group Member**: a non-functional role in Snyk. The user can access the Group but requires Organization-level permissions to use Snyk.
* **Group Viewer**: a user that can access the Group level but requires Organization-level permissions to take actions in Snyk. This is normally used as a starting point during onboarding with Snyk to understand functions tied to Group permissions and design a custom Group role for post-deployment use.
* **Group Admin**: provides a full set of permissions at the Group and Organization level. No other roles need to be assigned. This role is typically used for users who manage Snyk.
* **Organization Admin**: Typically assigned to Team Leads. Users with this role can add and delete Projects, override Snyk checks, and provision Group users with an Organization-level role.
* **Organization Collaborator**: This is the standard role in Snyk used for developers. This role is ideal for small teams or a very developer-first organizational approach.

{% hint style="info" %}
Large enterprises typically limit access to actions such as who can add or delete Projects, and so reduce permissions early on. For this purpose, you should use Custom Roles, for example, creating a **Team Lead** role to provide a middle ground between the **Organization Collaborator** and **Organization Admin** roles.
{% endhint %}

## Custom roles

The [User role management](../../../snyk-platform-administration/user-roles/user-role-management.md) feature provides Role-Based Access Control (RBAC) for Snyk, allowing you to create and enforce customized access by assigning a set of permissions to a role that will be granted to users. The permissions within these roles can be updated and changed even after assigning them to users.

As with all pre-defined roles, the permissions within the **Organization Collaborator** role cannot be changed, so you cannot grant the developers the ability to **Mark SCM Check as Successful**. This pre-defined role can be helpful when introducing developers to Snyk PR Checks but cannot be expanded when a developer is fully functioning in a Snyk environment. This is where Custom Roles are useful.

The most common types of custom roles include [Team Lead](../../../snyk-platform-administration/user-roles/custom-role-templates/team-lead-role-template.md) and [Developer](../../../snyk-platform-administration/user-roles/custom-role-templates/developer-role-template.md), but your company can have custom roles that suit your business structure. Snyk has provided multiple templates for you to use in structuring your Enterprise in Snyk; for more custom role suggestions and the listed permissions, see [Custom role templates](../../../snyk-platform-administration/user-roles/custom-role-templates/).
