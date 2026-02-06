# Manage users in a Tenant

{% hint style="info" %}
**Feature availability**

Tenant functions are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Select the name of your Tenant and the **Members** menu option to manage members:

<figure><img src="../../../.gitbook/assets/tenant-nav-member.png" alt="Members page option in the Tenant menu" width="201"><figcaption><p>Members page option in the Tenant menu</p></figcaption></figure>

## View Tenant members

<figure><img src="../../../.gitbook/assets/Tenant-member-list.png" alt="Tenant member management list with assigned roles"><figcaption><p>Tenant member management list with assigned roles</p></figcaption></figure>

On the Tenant **Members** page, you can see all the users associated with your Tenant, their authentication type, and their Tenant role. The [pre-defined](../../user-roles/pre-defined-roles.md#role-types) Tenant roles are:

* **Tenant Admin**: can access all Tenant products and settings. Reserved for Snyk Admins only.
* **Tenant Viewer:** can see the list of all Tenant users, all the Groups, and all the Organizations of the Tenant.
* **Tenant Member**: the default role of all users of the Tenant with no access to any Tenant level option.

Users with the Tenant Admin or Tenant Viewer roles can navigate the list of users by:

* Reading through the pages of users
* Searching by full name
* Filtering the list by role

## Change Tenant roles

{% hint style="info" %}
Only Tenant Admins can change the roles of Tenant users.
{% endhint %}

You can promote a Tenant Member to a Tenant Viewer or Admin by selecting the role dropdown next to the user's name and choosing the appropriate option.

## Delete Tenant Members

{% hint style="info" %}
Only Tenant Admins can delete Tenant users.
{% endhint %}

To delete a member from the Tenant, click the trash icon next to the user.
