# Manage users in your group

{% hint style="warning" %}
**This feature is currently a Private Beta.**&#x20;
{% endhint %}

{% hint style="info" %}
**Feature availability**\
Groups are available with Enterprise and Business plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Select the Group and **Members** top-level menu option to manage your group members:

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 05.38.32 (1).png>)

As a Group Admin (see [managing-permissions.md](managing-permissions.md "mention")), you can:

* [View group and org members](manage-users-in-your-organizations-1.md#view-group-and-org-members)
* [View individual members](manage-users-in-your-organizations-1.md#explore-individual-members)
* [Filter and sort views](manage-users-in-your-organizations-1.md#filter-and-sort-views)
* [Delete members](manage-users-in-your-organizations-1.md#delete-members)
* [Promote group member to a group admin](manage-users-in-your-organizations-1.md#undefined)

{% hint style="warning" %}
You cannot add external users directly to Groups; you must first add them to an Organization, then to a Group. See [manage-users-in-your-organizations.md](manage-users-in-your-organizations.md "mention") for details.
{% endhint %}

### View group and org members

In the group members page you can find all the members associated with your group, their respective roles and authentication type, and the number of orgs they are members of.

There are two standard roles available under group level - **Group Member** and **Group Admin**. Group Admins have all permissions at Snyk; see [managing-permissions.md](managing-permissions.md "mention"). However, being a Group Member does not directly grant the user any rights. They need to be added as org members or promoted as Group Admins.

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 05.52.57.png>)

### View individual members

Click on each member to view more details about their memberships.

If the user is a **Group Member**, you can see their role for each of the orgs they are a member of. You can filter by role since a Group Member can have different roles for different orgs. You can also remove the user from the group or orgs by invoking the respective delete buttons.

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 06.28.53.png>)

For a **Group Admin**, they are by default added as Org Admin across all Organizations in your Group. You cannot change a group admin's role for a specific org, or delete them from one or more orgs. However, you can remove a group admin from the group using the **Remove from group** option.

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 06.24.09.png>)

### Filter and sort views

#### Filter views

Click the filter icon (<img src="../../../.gitbook/assets/Screenshot 2022-03-11 at 08.47.59.png" alt="" data-size="line">) to expand the filter sidebar, to filter members displayed, by role or authentication method:

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 06.33.04.png>)

#### Sort views

You can sort by Name, Authentication method, Role, and Date joined.&#x20;

You can sort user views by clicking on the column heading:

![](<../../../.gitbook/assets/Screenshot 2022-03-11 at 09.01.07.png>)

### Delete members

To delete a member from the group:

1. Click the ![](<../../../.gitbook/assets/Screenshot 2022-03-11 at 08.05.56.png>) icon next to the user.
2. Click **Delete member from** _**your group's name**_ when prompted.

### Promote Group Member to a Group Admin

You can promote a Group Member to a Group Admin by selecting the role dropdown next to them and choosing the Group Admin role.

![](<../../../.gitbook/assets/Screenshot 2022-04-26 at 06.40.05.png>)

{% hint style="warning" %}
If the user is not already a part of your group, you must first add that user as a member of at least one org; see [Add Members](manage-users-in-your-organizations.md#add-members). The user then appears here with the role as Group Member, so you can then promote the user to Group Admin.
{% endhint %}
