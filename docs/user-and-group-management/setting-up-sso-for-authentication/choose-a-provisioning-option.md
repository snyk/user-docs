# Choose a provisioning option

## Introduction to provisioning options

Determine how new users in your company get access to Snyk: [Open to all](choose-a-provisioning-option.md#open-to-all), [Invitation required](choose-a-provisioning-option.md#invitation-required), or [Custom mapping](choose-a-provisioning-option.md#custom-mapping).

### Open to all

With the Open option, all users have automatic access to Snyk by using SSO to log in. They will have access to all Organizations in your selected Group. Their accounts will all be provisioned with the same role, with two options:

* The **Org** **administrator** role allows all new users to manage other Org admins and collaborators, view Group reports, and work with Organizations within the gFoup.
* The **Org** **collaborator** role can access the Organizations.

Let Snyk Support know whether new users will have the **administrator** role or the **collaborator** role for the organization. The selected role applies for all users.

### Invitation required

With the invitation required or **Group Member** option, admins can invite users or users can request access to an organization.

There are two ways to invite users to organizations. Invite members (see [Manage users in your Organizations](../managing-users-and-permissions/manage-users-in-your-organizations.md)) or automate the process using the Snyk [API Invite users endpoint.](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

If users who have not been invited use SSO to log in, they will gain access to Snyk, but they will not see any projects until an admin invites them or manually adds them to the Organization(s). You can show a list of Organizations with the appropriate contact person so that new users can [request access](https://docs.snyk.io/user-and-group-management/managing-users-and-permissions/organization-access-requests).

### Custom mapping

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

You can provision user accounts with customized rules, using the [Custom Mapping Option](custom-mapping-option/).

You can configure SSO differently for each of your different Groups. You can also map users to a specific Organization and role assignment based on information from the identity provider.

Work with your Snyk account team to prepare for implementing this SSO option.
