# Managing permissions

Click settings ![](../../../.gitbook/assets/cog\_icon.png) > **Members** to invite new team members.

1. The Free plan only has administrators.
2. Enterprise plans have administrators and collaborators.

{% hint style="info" %}
**Feature availability**\
Our free subscription plan allows you to send up to 200 pending invitations every 7 days. If you have already reached that limit and try to invite additional folks to join you, we'll display an error message for you. Don't worry though! You can send more the next week, or whenever some of the currently pending ones have been accepted. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

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
| Create an organization                         | x                   |                            |              |
| Create group level Service accounts\*          | x                   |                            |              |
| Set a License policy\*                         | x                   |                            |              |
| Set a Security policy\*\*                      | x                   |                            |              |
| Set global notifications preferences           | x                   |                            |              |
| Access to the account overall reporting        | x                   |                            |              |

(\*) Only in paid accounts \
(\*\*) Only in Enterprise Plan

Groups are allowed multiple administrators. Group administrators assign additional users to the role of "group administrator" by:

1. Select "Settings" for group (Example link: https://app.snyk.io/org/your-org-name/manage/members)
2. Click on member to update
3. Update role for member

****
