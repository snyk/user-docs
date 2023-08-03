# Manage permissions

{% hint style="info" %}
**Feature availability**\
The Snyk free subscription plan allows you to send up to 200 pending invitations every seven days and has only administrator roles. Enterprise plans have administrators, viewers and collaborators, and custom roles.\
See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Prerequisites for editing permissions

To edit permissions, you must have the required permissions. For example, only Group admins can assign other users as Group admins.

If your company currently does not have any Group admins, ask Snyk [support](https://support.snyk.io/hc/en-us/requests/new) to promote at least one user.

## Changing the role of a user

Follow these steps to change the role of a Group or Organization user:

1. Log in to your Snyk account and navigate to the Group or Organization where you want to change a user's role.
2. Select the **Members** option.
3. Find the member to update.
4. Update the role for that member, using the dropdown next to the role.

Alternatively, use the [Update a member's role in the Organization API call](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) to update the role of members in your Organization. This API call does not support updating the role of Group members.

{% hint style="info" %}
If the user is not already a part of your Group, you must first add that user as a member of at least one Organization; see [Add members](manage-users-in-your-organizations.md#add-members). The user then appears on the **Group > Members** page with the role **Group Member**, allowing you to change the role to Group Admin or Group Viewer.
{% endhint %}

## Permissions per role

The following list identifies the available permissions, what is allowed for each permission, and the roles that have each permission.

### Project Management permissions

#### Add/Delete Projects

_What is allowed_: Users with this permission can add and remove Snyk Projects for the Organization.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### Update Project with New Snapshot

_What is allowed_: Users with this permission can update the Project with a new snapshot.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### View Project Collections

_What is allowed_: Users with this permission can view Project collections.

_Roles with this permission_: Group Admin, Group Viewer, Org Admin, Org Collaborator

#### Create and Configure Project Collections

_What is allowed_: Users with this permission can create and configure Project collections.

_Roles with this permission_: Group Admin, Org Admin

#### Open Fix PRs

_What is allowed_: Users with this permission can open pull requests to fix issues in the Project.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### Delete Snapshot from Project History

_What is allowed_: Users with this permission can remove a snapshot from the Project history.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### Edit Project Attributes

_What is allowed_: Users with this permission can edit Project attributes.

_Roles with this permission_: Group Admin, Org Admin

### User Management permissions

#### View Team Members

_What is allowed_: Users with this permission can view the team members in the Organization.

_Roles with this permission_: Group Admin, Group Viewer, Org Admin, Org Collaborator

{% hint style="info" %}
The **Group Member** role allows the member to view team members. For additional permissions, a Group Member must be granted an Organization role or promoted to Group Viewer or Group Admin.
{% endhint %}

#### Invite/Remove Team Members

_What is allowed_: Users with this permission invite new members to a team or remove existing members.

_Roles with this permission_: Group Admin, Org Admin

#### Change Team Members' Roles

_What is allowed_: This permission allows users to change the roles of team members.

_Roles with this permission_: Group Admin, Org Admin

#### Leave Organization

_What is allowed:_ Users with this permission can choose to leave the Organization.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

### Organization Management permissions

#### Manage Plans and Billing for This Organization

_What is allowed_: Users with this permission can manage plans and billing for the Organization.

_Roles with this permission_: Group Admin, Org Admin

#### Delete Organization

_What is allowed_: Users with this permission can delete the Organization.

_Roles with this permission_: Group Admin, Org Admin

#### Create an Organization

_What is allowed_: Users with this permission can create a new Organization.

_Roles with this permission_: Group Admin

### Reports Management permissions

#### View Organization Reporting

_What is allowed_: Users with this permission can access and view reports at the Organization level.

_Roles with this permission_: Group Admin, Group Viewer, Org Admin, Org Collaborator

#### Access to the Account Overall Reporting (Group Level Reports)

_What is allowed_: Users with this permission can access and view reports at the group level.

_Roles with this permission_: Group Admin, Group Viewer

### Integration Management permissions

#### View Organization Integrations

_What is allowed_: Users with this permission can view integrations at the Organization level.

_Roles with this permission_: Group Admin, Group Viewer, Org Admin, Org Collaborator

#### Edit Organization Integrations

_What is allowed_: Users with this permission can edit integrations at the Organization level.

_Roles with this permission_: Group Admin, Org Admin

### Service Account Management permissions

#### Create an Org Level Service Account

_What is allowed_: Users with this permission can create a service account at the Organization level. This permission is available only for users with paid accounts.

_Roles with this permission_: Group Admin, Org Admin

#### Create Group Level Service Accounts

_What is allowed_: Users with this permission can create a service account at the group level. This permission is available only for users with paid accounts.

_Roles with this permission_: Group Admin

### Policy Setting permissions

#### Set a Security Policy

_What is allowed_: Users with this permission can set a security policy. This is available only for users with Enterprise plans.

_Roles with this permission_: Group Admin

#### Set a License Policy

_What is allowed_: Users with this permission can set a license policy. This is available only for users with paid accounts.

_Roles with this permission_: Group Admin

### Project Ignore Management permissions

#### View Project Ignores

_What is allowed_: Users with this permission can view Project ignores.

_Roles with this permission_: Group Admin, Group Viewer, Org Admin, Org Collaborator

#### Create Project Ignores

_What is allowed_: Users with this permission can create Project ignores.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### Edit Project Ignores

_What is allowed_: Users with this permission can edit Project ignores.

_Roles with this permission_: Group Admin, Org Admin, Org Collaborator

#### Remove Project Ignores

_What is allowed_: Users with this permission can remove Project ignores.

Roles with this permission: Group Admin, Org Admin, Org Collaborator

## Editing Project attributes from the Snyk CLI

The Organization Collaborator role lacks permission to edit Project attributes, including using the CLI [`snyk monitor`](../../snyk-cli/commands/monitor.md) command with arguments such as `--project-environment`. Attempting to use these arguments without the proper permission causes the `snyk monitor` command to fail.

Users or service accounts requiring the ability to edit Project attributes must be an Organization Admin or use a custom role with `org.project.attributes.edit` assigned. Alternatively, remove any arguments that edit the Project attributes from the `monitor` command.
