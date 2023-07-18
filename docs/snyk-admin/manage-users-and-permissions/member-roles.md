# Member Roles

{% hint style="info" %}
**Feature availability**\
This feature is available for Enterprise customers. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk pre-defined roles, such as Group Admin, cannot be changed. The Member Roles feature provides Role-Based Access Control (RBAC) for Snyk, allowing you to create and enforce customized access by assigning a set of permissions to a role that will be granted to users.

You can [manage roles](member-roles.md#manage-roles), [assign roles](member-roles.md#assign-roles) to users or to service accounts, and [use roles with SSO](member-roles.md#use-roles-in-custom-sso).

## Manage roles

The Member Roles feature allows Group Admins to:

* [Create a Role](member-roles.md#create-a-role)
* [Edit a Role](member-roles.md#edit-a-role)
* [Duplicate a Role](member-roles.md#duplicate-a-role)
* [Delete a Role](member-roles.md#delete-a-role)

You can manage custom roles, granting your users exactly the permissions they need to do their jobs across the Snyk platform. Thus you can ensure the right people have the right access to the right resources at the right time, thus maximizing transparency, and reducing organizational risk.

### Create a Role

**Group Admins** can find this option under **Select Group > Settings > Member Roles**.

You will find the default roles, **Org Admin** and **Org Collaborator**. When you click each of these roles, you can view the associated permissions, but you cannot modify the default roles.

Click the **Create new Role** button and enter a role name and description. Role names should be unique and can contain alphanumeric characters plus spaces.

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 15.28.12.png" alt="Create a new role"><figcaption><p>Create a new role</p></figcaption></figure>

Click the **Create role** button. Basic details about the role appear in the top section.

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 15.29.53.png" alt="Basic information about a role"><figcaption><p>Basic information about a role</p></figcaption></figure>

The bottom section lists all the permissions available at the Organization level that you can use to define the role.

<figure><img src="../../.gitbook/assets/Screenshot 2023-02-22 at 17.07.18.png" alt="Organization level permissions"><figcaption><p>Organization level permissions</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-30 at 17.03.22.png" alt="More Organization level permissions"><figcaption><p>More Organization level permissions</p></figcaption></figure>

Choose the required permissions and click **Update Role Permissions**.

{% hint style="info" %}
If you specify Move project permissions for the role, you must include Add project permissions for the Organization to which the Project is being moved.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-17 at 06.01.15.png" alt="Update Role Permissions"><figcaption><p>Update Role Permissions</p></figcaption></figure>

When the process of creating a role completes, a confirmation message appears.

<figure><img src="../../.gitbook/assets/image (18) (4).png" alt="Role creation confirmation message"><figcaption><p>Role creation confirmation message</p></figcaption></figure>

### Edit a Role

**Group Admins** can select a role (except for the default roles that are marked as locked) from the Member Roles list page and update the name, description, and permissions at any time . You can view how the default roles are set up and duplicate those roles, but you cannot edit them.

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 15.37.06.png" alt="Update Role Details"><figcaption><p>Update Role Details</p></figcaption></figure>

Select the permissions from the list at the bottom and click **Update Role Permissions**.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-17 at 06.10.00.png" alt="Update Role Permissions"><figcaption><p>Update Role Permissions</p></figcaption></figure>

When the process of updating a role completes, a confirmation message appears.

<figure><img src="../../.gitbook/assets/image (196).png" alt="Role update confirmation message"><figcaption><p>Role update confirmation message</p></figcaption></figure>

### Duplicate a Role

Group Admins can create a copy of an existing role by using the Duplicate role functionality. The system copies only the permissions associated with the role that you are duplicating and role memberships are not copied over.

To copy a role, use the **Duplicate** button next to each role in the Member Roles list, or select a role from the Member Roles list page, and when the Role details page opens, click the **Duplicate Role** button.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.09.png" alt="Member Roles List with Duplicate button for each role"><figcaption><p>Member Roles List with Duplicate button for each role</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 09.38.40.png" alt="Role details page with Duplicate button"><figcaption><p>Role details page with Duplicate button</p></figcaption></figure>

Enter a unique name and description and click the **Duplicate Role** button. A Group Admin can then edit this role to assign new permissions to it or rescind any permissions already assigned.

<figure><img src="../../.gitbook/assets/image (102) (1).png" alt="Enter new role details"><figcaption><p>Enter new role details</p></figcaption></figure>

### Delete a Role

Group Admins can delete a role if it is no longer needed by opening the role from the Member Roles List, clicking the **Delete** button, and confirming the delete action.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-17 at 06.22.22.png" alt="Delete role"><figcaption><p>Delate role</p></figcaption></figure>

If the role is assigned to one or more users including Service Accounts, you must select another role for them in order to delete the current role. This restriction prevents the Group Admin from accidentally deleting a role and leaving members with no access to Snyk.

When the current role is deleted, all its existing members including Service Accounts are reassigned to the new role selected.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-17 at 09.59.27.png" alt="Prompt to delete a role and reassign members"><figcaption><p>Prompt to delete a role and reassign members</p></figcaption></figure>

## Assign roles

### Assign roles to users

Users who hold the permissions to manage members can assign roles to members across all Organizations in the Group.

Using the [Update a member's role in the organization API call](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-a-member's-role-in-the-organization) you can update the role of the members in their Organizations.

In the Web UI, select an **Org** > **Members**.

For any member (Name) except a Group Admin, you can select the dropdown next to the current role and choose any role to assign that role to the member.

<figure><img src="../../.gitbook/assets/image (388) (1).png" alt="Select member role"><figcaption><p>Select member role</p></figcaption></figure>

\
You can invite new members to the Organization by assigning them a specified role.

Click the **Add members** button > **Invite new members**.

<figure><img src="../../.gitbook/assets/image (531).png" alt="Invite new members"><figcaption><p>Invite new members</p></figcaption></figure>

Click the **Add members** button > **Add existing members** to promote current Group Members to an org-specific role.

{% hint style="warning" %}
Snyk prevents users from assigning roles to others that have more privileges than those the user who is assigning roles already has. If you tried to update the role of a member, invite a new member, or add an existing member with a role that has more privileges than you have, you would see the error **Cannot assign higher privilege role**.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-01 at 15.51.05 (1).png" alt="User cannot assign more privileged role to another user"><figcaption><p>User cannot assign more privileged role to another user</p></figcaption></figure>

### Assign roles to Service Accounts

Users who have permission (Create Service Account) can set up new service accounts for their organization by choosing a role.

Select an **Org** > **Settings** > **Service Accounts** >\
Provide a name, choose a role, and click **Create**.

<figure><img src="../../.gitbook/assets/snyk-service-accounts.png" alt="Select a Role while creating Org Service Account"><figcaption><p>Select a Role while creating Org Service Account</p></figcaption></figure>

When you open a role that is assigned to Service Accounts, the system displays a warning message. When you update the permissions associated with or delete a role that would lead to reassigning the Service Accounts and users to a new role, be mindful of the potential impact.

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 15.49.49.png" alt="Warning that you are about to change a role assigned to a servie account"><figcaption><p>Warning that you are about to change a role assigned to a servie account</p></figcaption></figure>

{% hint style="warning" %}
Snyk prevents users from creating Organization service accounts with a role that has more privileges than what the user creating the service account has. If you try to create a service account with a role that has more privileges than you have, you will see the error **Cannot create a service account with a higher privilege role than yours**.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2022-08-01 at 15.59.52.png" alt="User cannot assign a more privileged role to a service account"><figcaption><p>User cannot assign a more privileged role to a service account</p></figcaption></figure>

## Use roles in custom SSO

Member roles are supported as part of a customized SSO authentication flow. All new and existing customers who have customized SSO can use new roles they create in their IDP SAML assertions to provision users in their Organizations with those roles.

{% hint style="info" %}
If you are a customer who already has Custom SSO set up, or you are planning to create Member Roles after setting up Custom SSO, you can use Member Roles _without_ any modification to the Custom SSO config on the the Snyk side, as long as you send the normalized role name in your payload in the agreed format.
{% endhint %}

New member role SAML assertions follow Snyk's existing pattern for declaring Organization memberships in IDP payloads: `{snyk-prefix}-{org-slug}-{normalized-role-name},` for example: `snyk-goof-developer_readonly`

* snyk-prefix: `snyk`
* org-name: `goof`
* role-name: `developer_readonly`

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-23 at 16.31.11.png" alt="Noarmalized name for a member role displayed in the Web UI"><figcaption><p>Noarmalized name for a member role displayed in the Web UI</p></figcaption></figure>

## Sample roles

### Org Collaborator who cannot ignore issues

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
* `View Project Collections`
* `View Targets`
* `View Users`

### Dashboard and report reviewer

Create a new role with permissions only to review dashboards and reporting for their management and executive teams.

Permissions:

* `View Organization`
* `View Organization Reports`
* `View Project`
* `View Project History`

For additional operations on the Dashboard add:

* `Add Project`
* `Create Pull Requests`

### Read-only CLI Tester

Create a new role that blocks use of `snyk monitor`.

Permissions:

* `View Organization`
* `View Project`
* `Test Packages`
* `Test Project`
* `View Preview Features`

### **Full Access CLI Tester**

Create a new role that can use `snyk test` and `snyk monitor`.

Permissions:

* `View Organization`
* `View Project`
* `View Project History`
* `Test Packages`
* `Add Project`
* `Test Project`
* `View Preview Features`

## Things to remember in working with Member Roles

* Permissions granted to users via Roles enable the same capabilities across all Snyk environments: Web UI, API, CLI, and IDE.
* The permission `View Organization` is needed by default for all Organization-level member roles.
* If the role is expected to view project-related data of an organization along with other operations, the `View Organization`, `View Project`, and `View Project History` permissions should be added to the role at minimum.
* The permission `View Preview Features` is required to run `snyk container test` and `snyk iac test`.
* Snyk prevents role privileges from escalating so that users cannot assign a higher privileged role to others and cannot create service accounts with a higher privileged role.
* It is advisable to use the Duplicate Role functionality and create a copy of a standard role and then amend the permissions rather than build a role from scratch if you are unsure about the permissions.
