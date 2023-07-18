# Manage permissions

{% hint style="info" %}
**Feature availability**\
The Snyk free subscription plan allows you to send up to 200 pending invitations every seven days and has only administrator roles. Enterprise plans have administrators, collaborators, and custom roles.\
See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

To edit permissions, you need the relevant permissions yourself, for example, only Group admins can assign other users as Group admins.

If your company currently does not have any Group admins, ask Snyk [support](https://support.snyk.io/hc/en-us/requests/new) to promote at least one user.

## Changing roles

To change roles for a user:

1. Click on the **Members** tab in the Snyk Web UI (example link: **https://app.snyk.io/org/your-org-name/manage/members**).
2. Find the member to update.
3. Update the role for that member, using the dropdown next to role.

Alternatively, use the [Update a member's role in the organization API call](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) to update the role of the members in your Organizations. This API call does not support updating the role of Group members. For updating the role of a Group member, follow the preceding steps to make changes via the UI.

{% hint style="info" %}
If the user is not already a part of your Group, you must first add that user as a member of at least one Organization; see [Add members](manage-users-in-your-organizations.md#add-members). The user then appears on the **Group > Members** page with the role **Group Member**, allowing you to change the role to Group Admin or Group Viewer.
{% endhint %}

## Permissions per role

The **Group Member** role allows the member to view team members. For additional permissions, a Group Member must be granted an Organization role or promoted to Group Admin.

<table><thead><tr><th width="139">Permission</th><th width="134">Group Admin</th><th width="139">Group Viewer</th><th width="118">Org Admin</th><th width="168">Org Collaborator</th></tr></thead><tbody><tr><td>Add/delete projects</td><td>x</td><td></td><td>x</td><td>x</td></tr><tr><td>Update project with new snapshot</td><td>x</td><td></td><td>x</td><td>x</td></tr><tr><td>View Project collections</td><td>x</td><td>x</td><td>x</td><td></td></tr><tr><td>Create and configure Project collections</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Open fix PRs</td><td>x</td><td></td><td>x</td><td>x</td></tr><tr><td>Delete snapshot from project history</td><td>x</td><td></td><td>x</td><td>x</td></tr><tr><td>Edit <a href="../../manage-issues/introduction-to-snyk-projects/project-attributes.md">Project attributes</a></td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>View team members</td><td>x</td><td>x</td><td>x</td><td>x</td></tr><tr><td>Invite/remove team members</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Change team members’ roles</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Create an Org level service account*</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Manage plans and billing for this organization</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Leave organization</td><td>x</td><td></td><td>x</td><td>x</td></tr><tr><td>Delete organization</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>View organization reporting</td><td>x</td><td>x</td><td>x</td><td>x</td></tr><tr><td>View organizations</td><td>x</td><td>x</td><td>x</td><td>x</td></tr><tr><td>Create an organization</td><td>x</td><td></td><td></td><td></td></tr><tr><td>View organization integrations</td><td>x</td><td>x</td><td>x</td><td>x</td></tr><tr><td>Edit organization integrations</td><td>x</td><td></td><td>x</td><td></td></tr><tr><td>Create group level Service accounts*</td><td>x</td><td></td><td></td><td></td></tr><tr><td>Set a License policy*</td><td>x</td><td></td><td></td><td></td></tr><tr><td>Set a Security policy**</td><td>x</td><td></td><td></td><td></td></tr><tr><td>Set global notifications preferences</td><td>x</td><td></td><td></td><td></td></tr><tr><td>Access to the account overall reporting<br>(Group level reports)</td><td>x</td><td>x</td><td></td><td></td></tr></tbody></table>

(\*) Only in paid accounts\
(\*\*) Only in Enterprise Plan

## Editing Project attributes from the Snyk CLI

The Organization Collaborator role lacks permission to edit Project attributes, including using the CLI [`snyk monitor`](../../snyk-cli/commands/monitor.md) command with arguments such as `--project-environment`. Attempting to use these arguments without the proper permission causes the `snyk monitor` command to fail.

Users or service accounts requiring the ability to edit project attributes must be an Organization Admin or use a custom role with `org.project.attributes.edit` assigned. Alternatively, remove any arguments that edit the project attributes from the `monitor` call.
