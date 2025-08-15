# Choose a provisioning option

{% hint style="warning" %}
Contact your Snyk account team or Snyk Support to turn on custom mapping once you have completed the setup steps.
{% endhint %}

Determine how new users in your company will get access to Snyk:

* [Open to all](choose-a-provisioning-option.md#open-to-all)
* [Invitation required](choose-a-provisioning-option.md#invitation-required)
* [Custom mapping](choose-a-provisioning-option.md#custom-mapping)

## Open to all

With the Open option, all users have automatic access to Snyk by using SSO to log in. They will have access to all Organizations in your selected Group. Their accounts will all be provisioned with the same role, with two options:

* The **Org** **administrator** role allows all new users to manage other Org admins and collaborators, view Group reports, and work with Organizations within the Group as well as perform non-administrative tasks in the Organization.
* The **Org** **collaborator** role can perform non-administrative tasks in the Organization.

Let Snyk Support know whether new users will have the administrator role or the collaborator role for the Organization. The selected role applies for all users. For details, see [Pre-defined roles](../../../snyk-platform-administration/user-roles/pre-defined-roles.md).

## Invitation required

With the invitation required or Group Member option, admins can invite users or users can request access to an Organization.

There are two ways to invite users to Organizations. Invite members (see [Manage users in Organizations](../../../snyk-platform-administration/groups-and-organizations/organizations/manage-users-in-organizations.md)) or automate the process using the Snyk [API Invite users endpoint](../../../snyk-api/reference/organizations-v1.md#org-orgid-invite).

If users who have not been invited use SSO to log in, they will gain access to Snyk, but they will not see any Projects until an admin invites them or manually adds them to the applicable Organizations. You can provide a list of Organizations with the appropriate contact person so that new users can [request access](../../../snyk-platform-administration/groups-and-organizations/organizations/requests-for-access-to-an-organization.md).

## Custom mapping

{% hint style="info" %}
**Feature availability**\
Custom Mapping is available only for Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).&#x20;
{% endhint %}

You can provision user accounts with customized rules using the [Custom Mapping Option](custom-mapping/).

You can configure SSO differently for each of your different Groups. You can also map users to a specific Organization and role assignment based on information from the identity provider.
