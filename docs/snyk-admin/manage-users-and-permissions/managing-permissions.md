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

| User permissions                                                                                   | Group Admin | Group Viewer | Org Admin | Org Collaborator |
| -------------------------------------------------------------------------------------------------- | ----------- | ------------ | --------- | ---------------- |
| Add/delete projects                                                                                | x           |              | x         | x                |
| Update project with new snapshot                                                                   | x           |              | x         | x                |
| Open fix PR's                                                                                      | x           |              | x         | x                |
| Delete snapshot from project history                                                               | x           |              | x         | x                |
| Edit [Project attributes](../../manage-issues/introduction-to-snyk-projects/project-attributes.md) | x           |              | x         |                  |
| View team members                                                                                  | x           | x            | x         | x                |
| Invite/remove team members                                                                         | x           |              | x         |                  |
| Change team members’ roles                                                                         | x           |              | x         |                  |
| Create an Org level service account\*                                                              | x           |              | x         |                  |
| Manage plans and billing for this organization                                                     | x           |              | x         |                  |
| Leave organization                                                                                 | x           |              | x         | x                |
| Delete organization                                                                                | x           |              | x         |                  |
| View organization reporting                                                                        | x           | x            | x         | x                |
| View organizations                                                                                 | x           | x            | x         | x                |
| Create an organization                                                                             | x           |              |           |                  |
| View organization integrations                                                                     | x           | x            | x         | x                |
| Edit organization integrations                                                                     | x           |              | x         |                  |
| Create group level Service accounts\*                                                              | x           |              |           |                  |
| Set a License policy\*                                                                             | x           |              |           |                  |
| Set a Security policy\*\*                                                                          | x           |              |           |                  |
| Set global notifications preferences                                                               | x           |              |           |                  |
| <p>Access to the account overall reporting<br>(Group level reports)</p>                            | x           | x            |           |                  |

(\*) Only in paid accounts\
(\*\*) Only in Enterprise Plan

## Editing Project attributes from the Snyk CLI

The Organization Collaborator role lacks permission to edit Project attributes, including using the CLI [`snyk monitor`](../../snyk-cli/commands/monitor.md) command with arguments such as `--project-environment`. Attempting to use these arguments without the proper permission causes the `snyk monitor` command to fail.

Users or service accounts requiring the ability to edit project attributes must be an Organization Admin or use a custom role with `org.project.attributes.edit` assigned. Alternatively, remove any arguments that edit the project attributes from the `monitor` call.
