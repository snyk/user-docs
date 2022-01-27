# Managing permissions

You can manage permissions to change roles for users - for example, to change a user from a collaborator to an organization administrator.

{% hint style="info" %}
To edit permissions, you will need relevant permissions yourself. For example, only group admins can assign other users as group admins.
{% endhint %}

{% hint style="info" %}
**Feature availability**\
Our free subscription plan allows you to send up to 200 pending invitations every 7 days, and only has administrator roles. Enterprise plans have administrators and collaborators.\
&#x20;See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

### Changing roles

To change roles for a user:

1. Go to the **Members** page (Example link: **https://app.snyk.io/org/your-org-name/manage/members**)
2. Find the the member to update.
3. Update the role for that member, using the drop down next to their current role.

### Permissions per role

| User permissions                               | Group Administrator | Organization Administrator | Collaborator |
| ---------------------------------------------- | ------------------- | -------------------------- | ------------ |
| Add/delete projects                            | x                   | x                          | x            |
| Update project with new snapshot               | x                   | x                          | x            |
| Open fix PR's                                  | x                   | x                          |              |
| Delete snapshot from project history           | x                   | x                          | x            |
| Invite/remove team members                     | x                   | x                          |              |
| Change team members’ roles                     | x                   | x                          |              |
| Create an org level service account\*          | x                   | x                          |              |
| Manage plans and billing for this organization | x                   | x                          |              |
| Leave organization                             | x                   | x                          | x            |
| Delete organization                            | x                   | x                          |              |
| View organization reporting                    | x                   | x                          | x            |
| Create an organization                         | x                   |                            |              |
| Create group level Service accounts\*          | x                   |                            |              |
| Set a License policy\*                         | x                   |                            |              |
| Set a Security policy\*\*                      | x                   |                            |              |
| Set global notifications preferences           | x                   |                            |              |
| Access to the account overall reporting        | x                   |                            |              |

(\*) Only in paid accounts \
(\*\*) Only in Enterprise Plan
