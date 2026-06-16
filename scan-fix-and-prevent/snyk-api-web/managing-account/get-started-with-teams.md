# Get started with teams

Plan and manage target scans between different teams in your organization.

Teams provide a way for you to group Users, API keys, and Targets in one place so that managing them becomes easier and time-efficient. You can create multiple teams for one account. Teams act as independent units and help you separate and assign targets for each team, limit the number of targets, set scopes, and user roles.

To understand more about Teams, visit [What are teams and how can they help me?](https://probely.com/blog/what-are-teams-and-how-can-they-help-me).

This article covers the following Team-related key aspects for you to get started:

* Adding a team with or without a target quota.
* Managing a team.
* Managing Users, API Keys, and Targets.
* Managing Scanning Agents.

The following sections detail each one of these aspects.

## Add a team with or without a target quota

Follow these steps to add a Team:

1. Log in to your account, open the **Settings** dropdown menu on the bottom-left corner of the navigation bar, and choose **Teams**.

<figure><img src="../../../.gitbook/assets/managing-account-get-started-with-teams-settings-teams.png" alt="Settings menu with Teams option"></figure>

1. Click **Add team** and give it a name. This is required to identify the team.
1. Optionally, you can limit the number of targets each team can create by assigning a target quota. This value depends on how many targets can still be added to your Snyk API & Web account.
1. If you assign a target quota (a limit of targets), you can reserve target slots for the current team. Doing so prevents other teams from using those slots beforehand. Reserving that quota to the team is optional.

<figure><img src="../../../.gitbook/assets/managing-account-get-started-with-teams-add-team.png" alt="Add team dialog with name and quota options"></figure>

## Manage a team

When you access a team's details, you have three tabs where you can manage that team: Users, API Keys, and Targets. The Users tab is selected by default.

To manage a team, go to its details and:

* From the **Users** tab, add users in bulk and assign them a role in the context of that team.
* Click on the **API Keys** tab, and manage API Keys precisely as you did for users.
* Go to the **Targets** tab, and add targets to the team by selecting the desired ones from the list of existing targets.

Once you have your team created and set up, users who were added to that team will be able to access the respective targets with the role you defined. Note that one user can be added to different teams and assigned roles that do not affect one another, such as an Admin or Developer. Also, bear in mind that a user with the Admin role within a team will be able to list all of the account's users, so that they are able to manage the users from their team independently.

## Manage Users, API Keys, and Targets

Similar to how you can manage every component of the Team (Users, API Keys, and Targets) from the team's details, you can also manage these components from their respective sections.

For instance, through the Users list, you can add new or edit existing users to accommodate distinct roles for different access scopes:

* With a **Global (account)** scope, the user can use the assigned role on the entire account.
* With a **specific Team** scope, the user can only use the assigned role in the context of that team.
* With a **specific Target** scope, the user can only use the assigned role in the context of that target.

To do this, select the role you want to assign and choose the respective scope using the Teams or Targets toggle and selecting the intended scope from the respective dropdown.

<figure><img src="../../../.gitbook/assets/managing-account-get-started-with-teams-user-role-scope.png" alt="User role assignment with scope selection"></figure>

You can do precisely the same for API Keys by accessing the new menu entry API Keys and either adding a new API Key or editing an existing one.

Just like for Users and API Keys, you can also define a specific target's scope by adding a new target or editing an existing one. If you assign a target to a team, members of that team can access it. Depending on the role they have in the context of that team, what they can do on the target may vary.

If the target is bound to the account or global level, only users with a Global role can list, view, change, or otherwise interact with that target. Depending on their role in a global or account context, what they can do on the target may vary.

* With the **Global (account)** scope, they will have access to all targets.
* With a specific **team** scope, they will only have access to targets that belong to that team.

<figure><img src="../../../.gitbook/assets/managing-account-get-started-with-teams-target-scope.png" alt="Target scope selection dialog"></figure>

Note that a target can only belong to one team at a time.

If a specific target belongs to a team, and you need to allow a user to access that target but do not want to give them access to every target from that team, you can always assign a target-specific role to that user by going to the Users list, as explained above.

## Manage Scanning Agents

Finally, you can manage scanning agents based on your organization's requirements.

* You can set a scanning agent to be used freely for the account, regardless of the scope of the target. Therefore, the scanning agent's scope is not restricted.
* Alternatively, you can restrict a scanning agent to one or more teams. Therefore, only targets that belong to those teams can use the scanning agent.

<figure><img src="../../../.gitbook/assets/managing-account-get-started-with-teams-scanning-agents.png" alt="Scanning agent team restrictions"></figure>
