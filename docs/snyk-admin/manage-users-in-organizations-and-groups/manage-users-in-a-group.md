# Manage users in a Group

{% hint style="info" %}
**Feature availability**\
Groups are available with the Snyk Enterprise plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Select the **Group** where you want to manage users and the **Members** menu option to manage your Group members:

<figure><img src="../../.gitbook/assets/snyk-group-member.png" alt="Manage Group members interface"><figcaption><p>Manage Group members interface</p></figcaption></figure>

As a Group Admin you can do the following:

* [View Group and Organization members](manage-users-in-a-group.md#view-group-and-organization-members)
* [View individual members](manage-users-in-a-group.md#view-individual-members)
* [Promote a Group member to a Group admin](manage-users-in-a-group.md#promote-a-group-member-to-a-group-admin)
* [Delete Group members](manage-users-in-a-group.md#delete-group-members)
* [Filter and sort views](manage-users-in-a-group.md#filter-and-sort-views-of-group-members)

{% hint style="warning" %}
You cannot add external users directly to Groups; you must first add them to an Organization, then to a Group. See [Manage users in your Organizations](manage-users-in-organizations.md) for details.
{% endhint %}

## View Group and Organization members

In the Group members page, you can find all the members associated with your Group, their respective roles and authentication type, and the number of Organizations they are members of.

There are two standard roles available under group level - **Group Member** and **Group Admin**. Group Admins have all permissions at Snyk; see [Managing permissions](../manage-permissions-and-roles/managing-permissions.md). However, being a Group Member does not directly grant the user any rights. They need to be added as org members or promoted as Group Admins.

<figure><img src="../../.gitbook/assets/image (250).png" alt="View Group members"><figcaption><p>View Group members</p></figcaption></figure>

## View individual members

Click on each member to view more details about their memberships.

If the user is a **Group Member**, you can see their role for each of the orgs they are a member of. You can filter by role since a Group Member can have different roles for different orgs. You can also remove the user from the group or orgs by invoking the respective delete buttons.

<figure><img src="../../.gitbook/assets/image (19) (2).png" alt="View member roles"><figcaption><p>View member roles</p></figcaption></figure>

Group Admin roles have access to all Organizations in that Group, with the same access level as an Organization Admin role in these Organizations.&#x20;

You cannot change a Group admin's role for a specific Organization, or delete them from one or more Organizations. However, you can remove a Group admin from the Group using the **Remove from group** option

## Promote a Group Member to a Group Admin

You can promote a Group Member to a Group Admin by selecting the role dropdown next to them and choosing the Group Admin role.

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-09 at 12.40.00.png" alt="Promote to Group Admin"><figcaption><p>Promote to Group Admin</p></figcaption></figure>

{% hint style="warning" %}
If the user is not already a part of your group, you must first add that user as a member of at least one Organization; see [Add Members](manage-users-in-organizations.md#add-members). The user then appears here with the role as Group Member, so you can then promote the user to Group Admin.
{% endhint %}

## Delete Group members

To delete a member from the group:

1. Click the trash icon next to the user.
2. Click **Delete member** from the Group you are managing.

## Filter and sort views of Group members

### Filter views

Click the filter icon (<img src="../../.gitbook/assets/Screenshot 2022-03-11 at 08.47.59.png" alt="" data-size="line">) to expand the filter sidebar, to filter members displayed, by role or authentication method:

<figure><img src="../../.gitbook/assets/Screenshot 2022-04-26 at 06.33.04.png" alt="Filter views"><figcaption><p>Filter views</p></figcaption></figure>

### Sort views

You can sort by Name, Authentication method, Role, and Date joined.

You can sort user views by clicking on the column heading:

<figure><img src="../../.gitbook/assets/Screenshot 2022-03-11 at 09.01.07.png" alt="Sort views"><figcaption><p>Sort views</p></figcaption></figure>

##

##
