# Managing permissions

{% hint style="info" %}
**Feature availability**\
Our free subscription plan allows you to send up to 200 pending invitations every 7 days, and only has administrator roles. Enterprise plans have administrators and collaborators.\
See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

To edit permissions, you will need the relevant permissions yourself--for example, only group admins can assign other users as group admins.

If your company currently does not have any group admins you will need to ask Snyk [support](https://support.snyk.io/hc/en-us/requests/new) to promote at least one user.

### Changing roles

To change roles for a user:

1. Click on the **Members** tab in the Snyk Web UI (example link: **https://app.snyk.io/org/your-org-name/manage/members**)
2. Find the member to update
3. Update the role for that member, using the dropdown next to role

### Permissions per role

**Group Member** role does not directly provide the user any rights. They need to be granted an organization role or promoted as Group Admins.

| User permissions                                                                                    | Group Administrator | Organization Administrator | Organization Collaborator |
| --------------------------------------------------------------------------------------------------- | ------------------- | -------------------------- | ------------------------- |
| Add/delete projects                                                                                 | x                   | x                          | x                         |
| Update project with new snapshot                                                                    | x                   | x                          | x                         |
| Open fix PR's                                                                                       | x                   | x                          | x                         |
| Delete snapshot from project history                                                                | x                   | x                          | x                         |
| Edit [project attributes](../../../snyk-web-ui/introduction-to-snyk-projects/project-attributes.md) | x                   | x                          |                           |
| Invite/remove team members                                                                          | x                   | x                          |                           |
| Change team members’ roles                                                                          | x                   | x                          |                           |
| Create an org level service account\*                                                               | x                   | x                          |                           |
| Manage plans and billing for this organization                                                      | x                   | x                          |                           |
| Leave organization                                                                                  | x                   | x                          | x                         |
| Delete organization                                                                                 | x                   | x                          |                           |
| View organization reporting                                                                         | x                   | x                          | x                         |
| Create an organization                                                                              | x                   |                            |                           |
| Create group level Service accounts\*                                                               | x                   |                            |                           |
| Set a License policy\*                                                                              | x                   |                            |                           |
| Set a Security policy\*\*                                                                           | x                   |                            |                           |
| Set global notifications preferences                                                                | x                   |                            |                           |
| Access to the account overall reporting                                                             | x                   |                            |                           |

(\*) Only in paid accounts\
(\*\*) Only in Enterprise Plan

### Editing project attributes from the Snyk CLI

The Organization Collaborator role lacks permission to edit project attributes, including via the Snyk CLI's [`monitor` command](../../../snyk-cli/commands/monitor.md) using arguments such as `--project-environment`. This was adjusted on August 25, 2022.  Attempting to use these arguments without the proper permission causes the `monitor` command to fail.&#x20;

Users or service accounts requiring the ability to edit project attributes will need to be an Organization Admin or use a custom role with `org.project.attributes.edit` assigned. Alternatively, remove any arguments that edit the project attributes from the `monitor` call.
