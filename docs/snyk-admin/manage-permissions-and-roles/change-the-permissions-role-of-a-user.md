# Change the permissions (role) of a user

{% hint style="info" %}
**Feature availability**\
Enterprise plans have pre-defined Snyk roles: administrators, viewers, and collaborators, as well as custom roles. Roles are defined by permissions. Manging permissions is available on Snyk Enterprise plans. See [Plans and pricing](https://snyk.io/plans/) for details.
{% endhint %}

To change the permissions of a user,  you must change the role of that user.

## Prerequisites for changing the role of a user

You must have specific permissions in order to change the role of a user.

For example, only Group admins can assign the role of Group admin to other users. If you want to add a Group admin, and your company currently does not have any Group admins, ask Snyk [support](https://support.snyk.io/hc/en-us/requests/new) to promote at least one user.

## How to change the role of a user

Follow these steps to change the role of a Group or Organization user:

1. Log in to your Snyk account and navigate to the Group or Organization where you want to change the role of a user.
2. Select the **Members** option.
3. Find the member to update.
4. Update the role for that member using the dropdown next to the user.

Alternatively, use the [Manage member's roles in the organization API v1 calls](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) to update the role of members in your Organization. This API call does not support updating the role of Group members.

{% hint style="info" %}
If the user is not already a part of your Group, you must first add that user as a member of at least one Organization; see [Add members on the Manage users in Organizations page](../manage-users-in-organizations-and-groups/manage-users-in-organizations.md#add-members). The user then appears on the **Group > Members** page with the role **Group Member**, allowing you to change the role to Group Admin or Group Viewer.
{% endhint %}

