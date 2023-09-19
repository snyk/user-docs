# Service accounts

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

You can set up a **service account** to be used for continuous integration (CI) and other automation purposes without using a Snyk user’s token.

Service accounts are a special type of system user. Creating a service account generates an API token that is the only token associated with the service account and takes the place of standard user credentials.

Use this token to provide credentials to authenticate with your Snyk account when you are setting up integration with your development tools and when you are working with the Snyk [CLI](../snyk-cli/) and [API](../snyk-api/). Snyk needs authentication in order to initiate Snyk processes.

You can generate single or multiple tokens on the Organization or Group levels to manage your integrations. Use Group-level tokens to use Group API endpoints, Organization API endpoints, and the CLI for all Organizations in the Group.

{% hint style="info" %}
Each service account has a unique name to make it easier to recognize. This name cannot be reused.
{% endhint %}

## Using service accounts

By using a service account token, you can:

* Create multiple tokens for different uses or integrations to manage each separately.
* Ensure seamless integrations, for example, when employees change roles or close their Snyk accounts.

{% hint style="info" %}
Group roles are only for service accounts on the Group level and are only for paid accounts.
{% endhint %}

{% hint style="info" %}
Service accounts can be used for GitHub Enterprise integrations. If your team needs to set up a service account in GitHub, the service account must be set up as a GitHub Enterprise integration. GHE is available only with Snyk Enterprise accounts.
{% endhint %}

## Set up a Group or Organization level service account

Generate single or multiple tokens on the Group or Organization levels to manage your integrations.

### Prerequisites to set up a service account

{% hint style="warning" %}
Group viewers are not able to create service accounts, regardless of their Org role.
{% endhint %}

To create a **Group service account,** you must be a Group admin. To create an **Organization service account,** you must be either a Group member and Org Admin, or a Group admin.

This process describes all options. Repeat the steps to create multiple tokens for the same or any other Group or Organization.

### How to set up a Group or Organization service account

* Log in to your account and navigate to the relevant Group and Organization that you want to manage.
* Click on **Settings** > **Service accounts** to view existing service accounts and their details.
* Click **Create a service account** to create a new one.\
  The screen that loads varies depending on whether you chose a **Group** or an **Organization.**

Note that while creating a **Group service account**, you can choose a Group level role.

<figure><img src="../.gitbook/assets/Screenshot 2022-07-06 at 12.01.28.png" alt="Group settings"><figcaption><p>Group settings</p></figcaption></figure>

In contrast, while creating an **Organization service account,** you can choose Org level roles, including custom [member roles](../snyk-admin/manage-permissions-and-roles/manage-member-roles.md) that you have set up for your Organizations.

<figure><img src="../.gitbook/assets/Screenshot 2022-07-06 at 12.06.35.png" alt="Organization settings"><figcaption><p>Organization settings</p></figcaption></figure>

#### Enter a service account name

In the **Service Account** name field, enter a unique name for this token. Remember, this name can be used only once for tokens in the same area, either an **Organization** or a **Group**.

<figure><img src="../.gitbook/assets/uuid-01c4cc98-23c9-3cb1-4972-1aa4f83ad98e-en.png" alt="Service account name and role"><figcaption><p>Service account name and role</p></figcaption></figure>

#### Select a role

From the **Role** dropdown list, select an appropriate role.

<figure><img src="../.gitbook/assets/image (1) (4).png" alt="Roles"><figcaption><p>Roles</p></figcaption></figure>

For Group service accounts, choose from the following list of roles to configure the scope of the token; Snyk recommends selecting Viewer or Admin.

* **Group Viewer** enables read-only access. Note that to set an API token to be read-only and unable to write to the platform, you must use a service account and set it to Group Viewer. See [Snyk API token permissions users can control](../snyk-api-info/using-snyk-api/api-token-permissions-users-can-control.md).
* **Group Admin** enables full administrator access.
* **Group Member** associates a service account with a group but does not grant any specific access.

For **Organization service accounts**, choose from the standard roles, **Org Admin** or **Org** **Collaborator**, or a custom role if you have set up any custom roles. See [Managing permissions](../snyk-admin/manage-permissions-and-roles/permissions-associated-with-each-pre-defined-role.md) for the scope of the Org Admin and Org Collaborator roles.

### Create the service account

Click **Create**.

The token is generated and displayed.

Make sure you copy this token, as you will not see it again. You can click **Close and Hide** once you have copied the token; whether you do or not, when you navigate away from this page, the token will no longer be visible. This is a standard security practice to keep your tokens safe.

#### How the token is associated with a Group and Organizations

The new token is also added to your **Existing service accounts** list, like the list in this example:

<figure><img src="../.gitbook/assets/uuid-799b88fc-d1d7-72c9-5ceb-30fb2a8d572e-en (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (14).png" alt="Existing service accounts for a Group"><figcaption><p>Existing service accounts for a Group</p></figcaption></figure>

In addition, if you created the token for the entire Group with a **Group Admin** role, the token also appears in the **Existing service accounts** list for each of its Organizations, though it can only be edited at the **Group** level.

<figure><img src="../.gitbook/assets/uuid-1110723e-74e7-3090-3e69-da65f93acfcc-en.png" alt="Existing accounts for an organization"><figcaption><p>Existing accounts for an Organization</p></figcaption></figure>

If you created the token from an Organization that is part of a Group, the token now also appears in the **Existing service account** list on the Group level. From that list, the Group Admin can also change the token name or delete it.

<figure><img src="../.gitbook/assets/uuid-50563edb-6a75-9f37-2040-cd814fdf9ead-en.png" alt="Group service accounts with Organization accounts listed"><figcaption><p>Group service accounts with Organization accounts listed</p></figcaption></figure>

## Update the name for a service account token

Click any of the links to update the name for a service account token:

* For **Group-level tokens**, from the **Group** level only
* For **Organization-level tokens**, from the relevant **Organization** and also from the **Group** level:

<figure><img src="../.gitbook/assets/uuid-b34e3d10-bb0c-b608-bc08-12f2bf0a4fc0-en.png" alt="Update a service account name"><figcaption><p>Update a service account name</p></figcaption></figure>

## Edit and delete a service account

Administrators can change token names and delete tokens.

### What happens to the service account token for a deleted account

When you delete a service account, the API token associated with it is invalidated immediately.

When an account is managed with Groups, the Organization and the Group admins can delete tokens for the Organization; only Group admins can view and manage tokens on the Group level.

Deleting a service account is the same as revoking the API token.

### How to edit or delete a service account

*   Log in to your account and navigate to the Group and Organization that you want to manage.

    For **Group tokens**, navigate to the Group level.\
    For **Organization tokens**, group admins can delete from either the Group or the relevant Organization; Organization admins should navigate to the relevant Organization.
* Click on **Settings** > **Service accounts**.
* Scroll to find the list of existing service accounts:

<figure><img src="../.gitbook/assets/uuid-799b88fc-d1d7-72c9-5ceb-30fb2a8d572e-en (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (14).png" alt="Existing service accounts for a Group"><figcaption><p>Existing service accounts for a Group</p></figcaption></figure>

* From the list of existing tokens:
  * Click the token name to navigate to **change the token name** and click **Save**.
  * Click **Delete** to **delete a token and invalidate it immediately**. When prompted, click **OK**. Remember that you cannot re-generate the same token.
