# Manage users in your organizations

You can use member functions to add, remove, and change member details for your organization.

Click on the **Members** tab in the app (example link: **https://app.snyk.io/org/your-org-name/manage/members**)

![](<../../../.gitbook/assets/Screen Shot 2022-02-24 at 1.05.40 PM.png>)

{% hint style="info" %}
Relevant permissions are required to perform all these tasks--see [managing-permissions.md](managing-permissions.md "mention").
{% endhint %}

### Add members

Click **Add members**:

![](<../../../.gitbook/assets/Screen Shot 2022-02-24 at 12.51.45 PM.png>)

From here:

* Select **Invite new members** to send an email invitation; enter the email addresses of users to invite (separated by commas), then click **Send invite**.
* Select **Add existing members** to add existing members of your group to this organization; select the members when prompted then click **Invite members.**

{% hint style="info" %}
Any users who have been invited are assigned a default role (such as **Collaborator**).To change this, use the **New members join as** drop-down. See also [managing-permissions.md](managing-permissions.md "mention").
{% endhint %}

#### Free Tier options:

![Snyk's Free Tier options](<../../../.gitbook/assets/Add members.png>)

The option to select **Invite by link** to send a link and click **Copy link** (in screenshot below) is available on the free tier so that you can send the link yourself.

![](<../../../.gitbook/assets/Members page.png>)

### Remove pending invites

Any pending invites appear on the members page in your organization with a number to state how many pending invites you have.

To remove these pending invites:

1. Select **pending invite** beneath the search field on the members page
2. Click the icon next to each user that has a pending invite to remove the invite.

### Remove users from the organization

1. Click the icon next to the user to remove them from the organization
2. Click **Delete member from** _**your organization’s name**_ in the modal

The user has been successfully removed from the organization.

### Change organization user roles

1. Roles are assigned via the drop-down list on the members page under the **Existing members**
2. Select the new role from the drop-down--the change is automatically saved
3. The user is then assigned the newly selected role.

### Additional actions

Filtering is available to users in the following fields:

* Role
* Authentication method

Sorting users is available with the following criteria:

* Name
* Authentication method
* Role
* Date joined
