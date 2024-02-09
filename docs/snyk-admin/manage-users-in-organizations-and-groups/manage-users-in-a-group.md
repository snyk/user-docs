# Manage users in a Group

{% hint style="warning" %}
**Release status**&#x20;

Groups available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Select the **Group** where you want to manage users and the **Members** menu option to manage your Group members:

<figure><img src="../../.gitbook/assets/snyk-group-member.png" alt="Manage Group members interface"><figcaption><p>Group membes page</p></figcaption></figure>

As a Group Admin you can do the following:

* [View Group and Organization members](manage-users-in-a-group.md#view-group-and-organization-members)
* [View individual members](manage-users-in-a-group.md#view-individual-members)
* [Promote a Group member to a Group admin](manage-users-in-a-group.md#promote-a-group-member-to-a-group-admin)
* [Delete Group members](manage-users-in-a-group.md#delete-group-members)
* [Filter and sort views](manage-users-in-a-group.md#filter-and-sort-views-of-group-members)

{% hint style="warning" %}
You cannot add external users directly to Groups; you must first add them to an Organization, and then to a Group. For more information, see [Manage users in Organizations](manage-users-in-organizations.md) for details.
{% endhint %}

## View Group and Organization members

On the Group members page, you can see all the members associated with your Group, their roles and authentication type, the number of Organizations they are members of, and the date they joined.

There are two standard roles available at the Group level, **Group Member** and **Group Admin**.

{% hint style="info" %}
Group Admins have all Snyk permissions; see [Pre-defined user roles](../manage-permissions-and-roles/pre-defined-user-roles.md). However, being a Group Member does not directly grant the user any rights. To have rights, users must be added as Organization members or promoted to Group Admin.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (250) (1).png" alt="View Group members"><figcaption><p>View Group members</p></figcaption></figure>

## View individual members

Click the line for each member to view the Group member details for that member.

If the user is a **Group Member**, you can see the user's role in each Organization where that user is a member. You can also see when the user was added and the user's authentication method.

A Group Member can have different roles in different Organizations. You can filter by role. You can also remove a user from a Group or Organization by using the available delete buttons.

<figure><img src="../../.gitbook/assets/image (19) (2) (1).png" alt="Group member details"><figcaption><p>Group member details</p></figcaption></figure>

Members with the Group Admin role have access to all Organizations in that Group, with the same access level as the Organization Admin role in these Organizations. You cannot change the role of a Group admin in a specific Organization, or delete a Group admin from one or more Organizations. However, you can remove a Group admin from the Group using the **Remove from group** option

## Promote a Group Member to a Group Admin

You can promote a Group Member to a Group Admin by selecting the **role** dropdown next to the user's name and choosing the Group Admin role.

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-09 at 12.40.00.png" alt="Promote to Group Admin"><figcaption><p>Promote to Group Admin</p></figcaption></figure>

{% hint style="warning" %}
If the user you want to promote to Group Admin is not already a part of your Group, you must first add that user as a member of at least one Organization; see [Add Members](manage-users-in-organizations.md#add-members) on the Manage users in Organizations page. The user then appears on the Group members page with the role of Group Member. You can then promote the user to Group Admin.
{% endhint %}

## Delete Group members

To delete a member from the Group:

1. Click the trash icon next to the user.
2. Click **Delete member** from the Group you are managing.

## Filter and sort views of Group members

### Filter views

On the Group members page, click the filter icon to expand the filter sidebar so you can filter the members displayed by role or authentication method:

<figure><img src="../../.gitbook/assets/Screenshot 2022-04-26 at 06.33.04.png" alt="Filter by role or authentication method"><figcaption><p>Filter by role or authentication method</p></figcaption></figure>

### Sort views

You can sort by name, authentication method, role, and date joined.

You can sort user views by clicking on the column heading:

<figure><img src="../../.gitbook/assets/Screenshot 2022-03-11 at 09.01.07.png" alt="Group members column headings"><figcaption><p>Group members column headings</p></figcaption></figure>

##

##
