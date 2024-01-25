# Determine Member Roles

## Use default or customized roles?

A key consideration when setting up Snyk is determining if the default member roles align with your needs, or if you need to customize member roles for your business requirements.&#x20;

{% hint style="info" %}
The default member roles in Snyk have a fixed set of permissions that cannot be changed. One of the biggest advantages of customizing member roles is that you can adjust the member role permissions to align with your desired level of control.&#x20;
{% endhint %}

## Default roles

The following are the default member roles:

* **Group Member**: a non-functional member role in Snyk. The member has access to the Group but requires Organization-level permissions to use Snyk.&#x20;
* **Group Admin**: provides a full set of permissions at the Group and Organization level. No other roles need to be assigned. This role is typically used for members who manage Snyk.&#x20;
* **Org Admin**: Typically assigned to team leads. Users with this role can add and delete Projects, override Snyk checks, and provision Group members with an Organization-level role.&#x20;
* **Org Collaborator**: This is the default role in Snyk used for developers. This role is ideal for small teams or a very developer-first organizational approach.&#x20;

{% hint style="info" %}
Large enterprises typically limit access to actions such as who can add or delete Projects, and so reduce permissions early on. For this purpose, you should use Custom Roles, for example, creating a **Team Lead** role to provide a middle ground between the **Org Collaborator** and **Org Admin** roles.
{% endhint %}

As with all default roles, the permissions within the **Org Collaborator** role cannot be changed, so you cannot grant the developers the ability to **Mark SCM Check as Successful**. This role can be helpful when introducing developers to Snyk PR Checks. This is where Custom Member Roles is useful.&#x20;

## Customized roles

The [Manage Member Roles](../../../snyk-admin/manage-permissions-and-roles/manage-user-roles.md) feature provides Role-Based Access Control (RBAC) for Snyk, allowing you to create and enforce customized access by assigning a set of permissions to a role that will be granted to users. The permissions within these roles can be updated and changed even after assigning them to users.&#x20;

The most common types of custom roles include Team Lead and Developer.

### Team Lead&#x20;

* This often starts as an equivalent to the Org Admin default role but would provide the ability to customize or change the permissions within the role if needed. You can duplicate the default Org Admin role as a starting point.&#x20;
* Sample permissions might include the ability to **Mark as Successful** in SCM for pull request/merge request checks., with the ignore capability disabled.

### Developer&#x20;

* This is a custom role that, at minimum, allows developers to view Organizations and Projects and also gives them the ability to test Projects.&#x20;
* Often, when a company is deploying Snyk, developers might have the ability to override Snyk PR checks, but this permission might get revoked after developers are comfortable using the Snyk IDE extensions and start fixing issues earlier in the SDLC. Similarly, you may start by allowing them to add Projects and then limit that permission to a team lead.

{% hint style="info" %}
Remember to give developers the ability to `test` Projects and `test` packages when you create the developer role to ensure developers can initiate tests.
{% endhint %}
