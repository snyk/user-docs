# Determine Member Roles

## Default or customized roles?

A key consideration when setting up Snyk is determining if the default member roles align with your needs, or if you need to customize member roles for your business requirements.&#x20;

{% hint style="info" %}
The default member roles in Snyk use a fixed set of permissions that cannot be changed. One of the biggest advantages to customizing member roles is that you can adjust the member role permissions to align with your desired level of control.&#x20;
{% endhint %}

## Default roles

The following are the default member roles:

* **Group Member**: a non-functional member role in Snyk. The member has access to the Group, but they require Organization-level permissions to use Snyk.&#x20;
* **Group Admin**: provides a full set of permissions at the Group and Organization level. No other roles need to be assigned. This role is typically used for members who manage Snyk.&#x20;
* **Org Admin**: Typically assigned to team leads. Users with this role can add/delete Projects, override Snyk checks, and provision Group members with an Organization-level role.&#x20;
* **Org Collaborator**: This is the default role in Snyk used for developers. This role is ideal for small teams or for a very developer-first organizational approach.&#x20;

{% hint style="info" %}
Large enterprises typically limit access to actions such as who can add or delete Projects, and so reduce permissions early on. For this purpose, you should use Custom Roles, for example, creating a **Team Lead** role to provide a middle ground between **Org Collaborator** and **Org Admin** roles.
{% endhint %}

As with all default roles, the permissions within the **Org Collaborator** role cannot be changed, so you cannot grant the developers the ability to "Mark SCM Check as Successful" - this role can be helpful when introducing developers to Snyk PR Checks. This is where Custom Member Roles is useful.&#x20;

## Customised roles

The [Manage Member Roles](../../../snyk-admin/manage-permissions-and-roles/manage-member-roles.md) feature provides Role-Based Access Control (RBAC) for Snyk, allowing you to create and enforce customized access by assigning a set of permissions to a role that will be granted to users. The permissions within these roles can be updated and changed even after assigning them to users.&#x20;

The most common types of custom roles include:

### Team Lead&#x20;

* This often starts as an equivalent to the Org Admin default role, but would provide the ability to customize or change the permissions within the role if needed. You can duplicate the default org admin role as a starting point.&#x20;
* Sample permissions might include the ability to Mark as Successful in SCM for pull request/merge request checks., but disabling the ignore capability.

### Developer&#x20;

* This is a custom role that at minimum allows developers to  view access to orgs, projects, etc and also give them the ability to test projects.&#x20;
* Often when deploying Snyk, developers might have the ability to override Snyk PR checks, but this permission might get revoked once they are comfortable using the Snyk IDE extensions and they start fixing issues earlier in the SDLC. Similarly you may start by allowing them to add projects and then limit it to a team lead.

{% hint style="info" %}
Don't forget to give developers the ability to run "test projects" and "test packages" when creating the developer role to ensure they can initiate tests!
{% endhint %}
