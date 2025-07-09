# Manage users in Organizations

In the **Organization** where you want to manage users, select the **Members** menu option.

{% hint style="info" %}
You must have the permissions required to perform these tasks. For details, see [Pre-defined user roles](../../user-roles/pre-defined-roles.md) for a list of permissions.
{% endhint %}

## Add users

To add new users to your Organization, click **Add members**:

<div align="left"><figure><img src="../../../.gitbook/assets/Screen Shot 2022-02-24 at 12.51.45 PM.png" alt=""><figcaption><p>Add members to your Organization</p></figcaption></figure></div>

You can do the following on the **Add members** screen:

* Select **Invite new members** to send an email invitation to a new user. Enter the email addresses of users to invite, separated by commas, and click **Send invite**.
* Select **Add existing members** to add existing members of your Group to the Organization. Select the users when prompted and click **Invite members.**
* For Free plan users only:\
  Select **Invite by link** to send a link; click **Copy link** and send the link yourself. Note that an invite link expires after 48 hours.
* Use the **New members join as** dropdown to define the default role of a user when joining, such as **Org admin**. For details, see [Manage permissions](../../user-roles/pre-defined-roles.md).

For a demonstration of adding users, view this video:

{% embed url="https://thoughtindustries-1.wistia.com/medias/qqmkcaequj" %}
Inviting members to an Organization
{% endembed %}

## Revoke pending invites

Follow these steps to cancel pending invites:

1.  On the **Members** page, click the **Revoke pending invites link**, which appears when there is at least one pending invite.\


    <figure><img src="../../../.gitbook/assets/Revoke.png" alt=""><figcaption><p>Revoke pending invites</p></figcaption></figure>
2. In the **Pending invites in \_your organization's name**\_ modal that appears, click the trash icon next to the name of the user to cancel the invite.

## Change member roles

To change the role of a user, click on the **Role** entry for the member and use the dropdown to select the new role:

<figure><img src="../../../.gitbook/assets/org-member-change-role.png" alt=""><figcaption><p>Select new role</p></figcaption></figure>

{% hint style="warning" %}
For Enterprise plan customers who create custom roles, Snyk prevents users from assigning roles to other users who have more privileges. If you try to update a role, invite a new user, or add an existing user with a role that has more privileges than you have, you will see the error **Cannot assign higher privilege role**.
{% endhint %}

## Delete Organization members

Follow these steps to delete a member from the Organization:

1. Click the trash icon next to the user.
2. Click **Delete member from** \[**your Organization’s name]** when prompted.

## Filter and sort views of Organization members

### Filter views

Click the filter icon to expand the filter sidebar and then choose to filter the members displayed by role or authentication method:

<div align="left"><figure><img src="../../../.gitbook/assets/org-member-filters.png" alt=""><figcaption><p>Filter members by role or by authentication method</p></figcaption></figure></div>

### Sort views

Click on the column heading to sort user views:

<figure><img src="../../../.gitbook/assets/org-member-column-sort.png" alt=""><figcaption><p>Sort user views</p></figcaption></figure>

You can sort by name, authentication method, role, and date joined.
