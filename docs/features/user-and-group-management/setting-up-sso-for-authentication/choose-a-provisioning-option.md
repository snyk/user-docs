# Choose a provisioning option

Determine how new users in your organization get access to Snyk.

## Open to all

With the Open option, all users have automatic access to Snyk by using SSO to log in. They will have access to all organizations in your selected group. Their accounts will all be provisioned with the same role, with two options:

* The **Org** **administrator** role allows all new users to manage other Org admins and collaborators, view group reports, and work with organizations within the group.
* The **Org** **collaborator** role can access the organizations.

Let Snyk Support know whether new users will have the **administrator** role or the **collaborator** role for the organization. The selected role applies for all users.

## Invitation required

With the invitation required or **Group Member** option, admins can invite users or users can request access to an organization.

There are two ways to invite users to organizations. Invite members (see [manage-users-in-your-organizations.md](../managing-users-and-permissions/manage-users-in-your-organizations.md "mention") or automate the process using our [API endpoint.](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

If users who have not been invited use SSO to log in, they will gain access to Snyk, but they will not see any projects until an admin invites them or manually adds them to the organization(s). You can show a list of organizations with the appropriate contact person so that new users can [request access](https://docs.snyk.io/user-and-group-management/managing-users-and-permissions/organization-access-requests).

## Custom Mapping

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

You can provision user accounts with customized rules.

You can configure SSO differently for each of your different Snyk groups. You can also map users to a specific organization and role assignment based on information from the identity provider.

Work with your Customer Success Manager and Snyk Technical Services to prepare for implementing this SSO option.

Speak to your sales rep to learn more about the Custom Mapping option.
