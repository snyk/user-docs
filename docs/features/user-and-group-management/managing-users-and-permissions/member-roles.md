# Member Roles

{% hint style="info" %}
**Feature availability**\
This feature is available for Enterprise customers. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

### Introduction

Access Management is an integral part of the security strategy of an enterprise. While Snyk offers out-of-the-box roles like Group Admin, Group Member, Org Admin, and Org Collaborator, these roles are pre-defined and respective accesses cannot be modified. With **Member Roles**, we are introducing role-based access control (RBAC) at Snyk that allows you to create and enforce advanced access by assigning a set of permissions to a role that will be granted to users.

This is a self-serve capability for Group admins to grant members the permissions they need to do their jobs across the Snyk platform - and only those permissions - by creating customized roles and assigning specific permissions to them. This ensures the right people have the right access to the right resources at the right time, while maximizing transparency and reducing organizational risk.

{% hint style="info" %}
This version allows you to create custom roles for organizations in your group, and assign those roles to organization members and service accounts.\
Future versions will add group-level access (for example, policies, group member management, group reporting, and so on).
{% endhint %}

### Create a Role

**Group Admins** can find the option under Select Group > Settings > Member Roles.

You can find the default roles - Org Admin and Org Collaborator. When you click each of these roles you can view the associated permissions, but cannot modify the default roles.

![Create new role option under Group Settings](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.21.58.png>)

Click the **Create new Role** button and enter a role name and description. Role names should be unique, and can contain alphanumeric characters plus spaces.

![Create a new role](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.28.12.png>)

Click the **Create role** button. You will see basic details about the role in the top section.

![Role basic info](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.29.53.png>)

The bottom section lists all the permissions available at the organization level that you use to define the role.

![Organization level permissions](<../../../.gitbook/assets/Screenshot 2022-05-17 at 05.56.02.png>)

Choose the required permissions and click **Update Role Permissions**.

![Update Role Permissions](<../../../.gitbook/assets/Screenshot 2022-05-17 at 06.01.15.png>)

When creating the role is complete, you will see the confirmation message at the top.

![](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.32.32.png>)

### Edit a Role

**Group Admins** can select a role (except for the default roles that are marked as locked) from the Member Roles list page and update the name, description and permissions at any time . You can view how the default roles are set up and duplicate those roles, but you cannot edit them.

![](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.37.06.png>)

![Update Role Permissions](<../../../.gitbook/assets/Screenshot 2022-05-17 at 06.10.00.png>)

When updating the role is complete, you will see the confirmation message at the top.

![](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.38.52.png>)

### Duplicate a Role

Group Admins can create a copy of an existing role by using Duplicate role functionality. The system copies only the permissions associated with the role that you are duplicating and role memberships are not copied over.

You can use the Duplicate button next to each role in the Member Roles list page. Or, select a role from the Member Roles list page. On the Role details page, click the **Duplicate Role** button.

![Member Roles List page with Duplicate Role buttons](<../../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.09.png>)

![Role details page with Duplicate button](<../../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.40.png>)

Enter a unique name and description and click the **Duplicate Role** button. Group Admin can then edit this role to assign new permissions to it or rescind any permissions already assigned.

![](<../../../.gitbook/assets/image (6) (2).png>)

### Delete a Role

Group Admins can delete a role if it is no longer needed by opening the role from the Member Roles List page and clicking the **Delete** button.

![Delate role](<../../../.gitbook/assets/Screenshot 2022-05-17 at 06.22.22.png>)

If the role is assigned to one or more users including Service Accounts, select another role for them in order to delete the current role. This is to avoid having the Group Admin accidentally delete a role leaving members with no access to Snyk.

When the current role is deleted, all its existing members including Service Accounts are reassigned to the new role selected.

![Prompt to reassign members and delete a role](<../../../.gitbook/assets/Screenshot 2022-05-17 at 09.59.27.png>)

### Assigning roles to users

Users who hold the permissions to manage members can assign the roles to members across all organizations in the Group.

Using [this API call](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) users can update the role of the members in their organizations.

In the UI, select an **Org** > **Members**.

For any member (Name) except a Group Admin, the user can select the dropdown next to the current role and choose any role to assign that role to the member.

![](<../../../.gitbook/assets/image (11) (1).png>)

Click the **Add members** button > **Invite new members**.\
\
You can invite new members to the org by assigning them a specified role.

![](<../../../.gitbook/assets/image (8) (2).png>)

Choose **Add members** button > **Add existing members** to promote current Group Members to an org-specific role.

{% hint style="warning" %}
Snyk prevents users from assigning roles to others that have more privileges than what they already have. You would encounter the following error while trying to update the role of a member, invite a new member, or add an existing member with a role that has more privileges than the logged in user.
{% endhint %}

![User cannot assign more privileged role to another user](<../../../.gitbook/assets/Screenshot 2022-08-01 at 15.51.05 (1).png>)

### Assigning roles to Service Accounts

Users who have permission (Create Service Account) can set up new service accounts for their organization by choosing a role.

Select an **Org** > **Settings** > **Service Accounts** >\
Provide a name, choose a role, and click **Create**.

![Select a Role while creating Org Service Account](<../../../.gitbook/assets/Screenshot 2022-05-31 at 09.42.43.png>)

When you open a role that is assigned to Service Accounts, system would display a warning message. While updating the permissions associated with the role or deleting the role that would lead to reassigning the Service Accounts and users to a new role, be mindful of the potential impact.

![](<../../../.gitbook/assets/Screenshot 2022-06-23 at 15.49.49.png>)

{% hint style="warning" %}
Snyk prevents users from creating organisation service accounts with a role that has more privileges than what they already have. You would encounter the below error, while trying to create a service account with a role that has more privileges than the logged in user.
{% endhint %}

![](<../../../.gitbook/assets/Screenshot 2022-08-01 at 15.59.52.png>)

### Roles in Custom SSO

Member roles are supported as part of a Customized SSO authentication flow. All new and existing customers who have customized SSO will be able to use new roles they create in their IDP SAML assertions to provision users in their orgs with those roles.

New member role SAML assertions follow Snyk's existing pattern for declaring org memberships in IDP payloads: `{snyk-prefix}-{org-name}-{normalized-role-name},` for example: `snyk-goof-developer_readonly`

* snyk-prefix: `snyk`
* org-name: `goof`
* role-name: `developer_readonly`

![](<../../../.gitbook/assets/Screenshot 2022-06-23 at 16.31.11.png>)

### Sample Roles

#### Org Collaborator who cannot ignore issues

Create a new role similar to Org Collaborator but which blocks the ability to ignore issues.

Permissions:

* `Add Project`
* `Create Jira issues`
* `Create Pull Requests`
* `Edit Project`
* `Edit Project Tags`
* `Project Status`
* `Remove Project`
* `Remove Targets`
* `Test Packages`
* `Test Project`
* `User Leave`
* `View Audit Logs`
* `View Entitlements`
* `View Integrations`
* `View Jira issues`
* `View Organization`
* `View Organization Reports`
* `View Preview Features`
* `View Project`
* `View Project Ignores`
* `View Targets`
* `View Users`

#### Dashboard and Report Reviewer

Create a new role with permissions only to review dashboards and reporting for their management and executive teams.

Permissions:

* `View Organization`
* `View Organization Reports`
* `View Project`

For additional operations on the Dashboard add:

* `Add Project`
* `Create Pull Requests`

#### Read-only CLI Tester

Create a new role that blocks use of Snyk Monitor.

Permissions:

* `View Organization`
* `Test Packages`
* `Test Project`
* `View Preview Features`

**Full Access CLI Tester**

Create a new role that can use Snyk Test and Snyk Monitor.

Permissions:

* `View Organization`
* `Test Packages`
* `Add Project`
* `Test Project`
* `View Preview Features`

### Things to remember

* Permissions granted to users via Roles enable the same capabilities across all Snyk environments: Web UI, API, CLI, and IDE.
* If the Role is expected to view organisation or project-related data or both along with other operations - `View Organization and View Project` permissions should be added to the role.
* Snyk prevents role privileges from escalating so that users cannot assign a higher privileged role to others or cannot create service accounts with a higher privileged role.
* It is advisable to use the Duplicate Role functionality and create a copy of a standard role and then amend the permissions instead of building a role from scratch if you are unsure about the permissions.
