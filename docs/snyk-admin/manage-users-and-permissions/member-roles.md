# Member Roles

{% hint style="info" %}
**Feature availability**\
This feature is available for Enterprise customers. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk pre-defined roles (such as Group Admin) cannot be changed. Member Roles provides Role-Based Access Control (RBAC) for Snyk, allowing you to create and enforce advanced access by assigning a set of permissions to a role that will be granted to users.

You can [manage roles](member-roles.md#manage-roles), [assign roles](member-roles.md#assign-roles) to users or to service accounts, and [use roles with SSO](member-roles.md#use-roles-in-custom-sso).

## Manage roles

Member roles allows Group Admins to:

* [Create a Role](member-roles.md#create-a-role)
* [Edit a Role](member-roles.md#edit-a-role)
* [Duplicate a Role](member-roles.md#duplicate-a-role)
* [Delete a Role](member-roles.md#delete-a-role)

This allows you to manage custom roles, granting your users the precise permissions they need to do their jobs across the Snyk platform. So you can ensure the right people have the right access to the right resources at the right time, while maximizing transparency and reducing organizational risk.

### Create a Role

**Group Admins** can find the option under Select Group > Settings > Member Roles.

You can find the default roles - Org Admin and Org Collaborator. When you click each of these roles you can view the associated permissions, but cannot modify the default roles.

Click the **Create new Role** button and enter a role name and description. Role names should be unique and can contain alphanumeric characters plus spaces.

![Create a new role](<../../.gitbook/assets/Screenshot 2022-06-23 at 15.28.12.png>)

Click the **Create role** button. You will see basic details about the role in the top section.

![Role basic info](<../../.gitbook/assets/Screenshot 2022-06-23 at 15.29.53.png>)

The bottom section lists all the permissions available at the organization level that you use to define the role.

<figure><img src="../../.gitbook/assets/Screenshot 2023-02-22 at 17.07.18.png" alt=""><figcaption><p>Organization level permissions</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-30 at 17.03.22.png" alt=""><figcaption><p>Organization level permissions</p></figcaption></figure>

Choose the required permissions and click **Update Role Permissions**.

![Update Role Permissions](<../../.gitbook/assets/Screenshot 2022-05-17 at 06.01.15.png>)

When creating the role is complete, you will see the confirmation message at the top.

<figure><img src="../../.gitbook/assets/image (13).png" alt="Role creation confirmation message"><figcaption><p>Role creation confirmation message</p></figcaption></figure>

### Edit a Role

**Group Admins** can select a role (except for the default roles that are marked as locked) from the Member Roles list page and update the name, description and permissions at any time . You can view how the default roles are set up and duplicate those roles, but you cannot edit them.

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 15.37.06.png" alt="Edit role details"><figcaption><p>Edit role details</p></figcaption></figure>

![Update Role Permissions](<../../.gitbook/assets/Screenshot 2022-05-17 at 06.10.00.png>)

When updating the role is complete, you will see the confirmation message at the top.

<figure><img src="../../.gitbook/assets/image (272).png" alt="Role change confirmation message"><figcaption><p>Role change confirmation message</p></figcaption></figure>

### Duplicate a Role

Group Admins can create a copy of an existing role by using Duplicate role functionality. The system copies only the permissions associated with the role that you are duplicating and role memberships are not copied over.

You can use the Duplicate button next to each role in the Member Roles list page. Or, select a role from the Member Roles list page. On the Role details page, click the **Duplicate Role** button.

![Member Roles List page with Duplicate Role buttons](<../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.09.png>)

![Role details page with Duplicate button](<../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.40 (1).png>)

Enter a unique name and description and click the **Duplicate Role** button. Group Admin can then edit this role to assign new permissions to it or rescind any permissions already assigned.

<figure><img src="../../.gitbook/assets/image (394) (1) (1) (1) (1) (1) (1) (1).png" alt="Enter new role details"><figcaption><p>Enter new role details</p></figcaption></figure>

### Delete a Role

Group Admins can delete a role if it is no longer needed by opening the role from the Member Roles List page and clicking the **Delete** button.

![Delate role](<../../.gitbook/assets/Screenshot 2022-05-17 at 06.22.22.png>)

If the role is assigned to one or more users including Service Accounts, select another role for them in order to delete the current role. This is to avoid having the Group Admin accidentally delete a role leaving members with no access to Snyk.

When the current role is deleted, all its existing members including Service Accounts are reassigned to the new role selected.

![Prompt to reassign members and delete a role](<../../.gitbook/assets/Screenshot 2022-05-17 at 09.59.27.png>)

## Assign roles

### Assign roles to users

Users who hold the permissions to manage members can assign the roles to members across all Organizations in the Group.

Using [this API call](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) users can update the role of the members in their organizations.

In the Web UI, select an **Org** > **Members**.

For any member (Name) except a Group Admin, the user can select the dropdown next to the current role and choose any role to assign that role to the member.

<figure><img src="../../.gitbook/assets/image (104) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Select member role"><figcaption><p>Select member role</p></figcaption></figure>

Click the **Add members** button > **Invite new members**.\
\
You can invite new members to the org by assigning them a specified role.

<figure><img src="../../.gitbook/assets/image (76).png" alt="Invite new members"><figcaption><p>Invite new members</p></figcaption></figure>

Choose **Add members** button > **Add existing members** to promote current Group Members to an org-specific role.

{% hint style="warning" %}
Snyk prevents users from assigning roles to others that have more privileges than what they already have. You would encounter the following error while trying to update the role of a member, invite a new member, or add an existing member with a role that has more privileges than the logged in user.
{% endhint %}

![User cannot assign more privileged role to another user](<../../.gitbook/assets/Screenshot 2022-08-01 at 15.51.05.png>)

### Assign roles to Service Accounts

Users who have permission (Create Service Account) can set up new service accounts for their organization by choosing a role.

Select an **Org** > **Settings** > **Service Accounts** >\
Provide a name, choose a role, and click **Create**.

![Select a Role while creating Org Service Account](../../.gitbook/assets/snyk-service-accounts.png)

When you open a role that is assigned to Service Accounts, the system would display a warning message. While updating the permissions associated with the role or deleting the role that would lead to reassigning the Service Accounts and users to a new role, be mindful of the potential impact.

![](<../../.gitbook/assets/Screenshot 2022-06-23 at 15.49.49.png>)

{% hint style="warning" %}
Snyk prevents users from creating organization service accounts with a role that has more privileges than what they already have. You would encounter the below error while trying to create a service account with a role that has more privileges than the logged in user.
{% endhint %}

![](<../../.gitbook/assets/Screenshot 2022-08-01 at 15.59.52.png>)

## Use Roles in Custom SSO

Member roles are supported as part of a Customized SSO authentication flow. All new and existing customers who have customized SSO will be able to use new roles they create in their IDP SAML assertions to provision users in their orgs with those roles.

{% hint style="info" %}
If you are a customer who already has Custom SSO set up or you are planning to create Member Roles after setting up Custom SSO, you can use Member Roles _without_ any modification to the Custom SSO config at Snyk side, as long as you send normalized role name in your payload in the agreed format.
{% endhint %}

New member role SAML assertions follow Snyk's existing pattern for declaring org memberships in IDP payloads: `{snyk-prefix}-{org-slug}-{normalized-role-name},` for example: `snyk-goof-developer_readonly`

* snyk-prefix: `snyk`
* org-name: `goof`
* role-name: `developer_readonly`

![](<../../.gitbook/assets/Screenshot 2022-06-23 at 16.31.11.png>)

## Sample Roles

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
* `View Project History`
* `View Project Ignores`
* `View Targets`
* `View Users`

#### Dashboard and Report Reviewer

Create a new role with permissions only to review dashboards and reporting for their management and executive teams.

Permissions:

* `View Organization`
* `View Organization Reports`
* `View Project`
* `View Project History`

For additional operations on the Dashboard add:

* `Add Project`
* `Create Pull Requests`

#### Read-only CLI Tester

Create a new role that blocks use of Snyk Monitor.

Permissions:

* `View Organization`
* `View Project`
* `Test Packages`
* `Test Project`
* `View Preview Features`

**Full Access CLI Tester**

Create a new role that can use Snyk Test and Snyk Monitor.

Permissions:

* `View Organization`
* `View Project`
* `View Project History`
* `Test Packages`
* `Add Project`
* `Test Project`
* `View Preview Features`

## Things to remember

* Permissions granted to users via Roles enable the same capabilities across all Snyk environments: Web UI, API, CLI, and IDE.
* `View Organization` permission is needed by default for all organization level member roles.
* If the Role is expected to view project-related data of an organization along with other operations - `View Organization , View Project, and View Project History` permissions should be added to the role at a minimum.
* `View Preview Features` permission is required to run a Snyk Container test and Snyk IaC test.
* Snyk prevents role privileges from escalating so that users cannot assign a higher privileged role to others or cannot create service accounts with a higher privileged role.
* It is advisable to use the Duplicate Role functionality and create a copy of a standard role and then amend the permissions instead of building a role from scratch if you are unsure about the permissions.
