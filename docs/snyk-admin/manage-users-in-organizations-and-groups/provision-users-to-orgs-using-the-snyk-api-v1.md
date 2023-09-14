# Provision users to Orgs using the Snyk API v1

The [Provision user endpoints](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization) in the Snyk API v1 allow you to organize and grant permissions to your single sign-on users before the users log in to the Snyk platform.

Provisioned users do not need to accept invites. When provisioned users first log in to Snyk, they will have all their permissions. You can use the [Provision a user to the organization](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization) endpoint to add users to Organizations at scale before their first login.

## Pre-requisites for provisioning users via API

{% hint style="warning" %}
The API does not support using service accounts as the inviting user or provisioned user.
{% endhint %}

* The user being provisioned must not already exist in the Snyk system.
* The inviting user must call the API using a personal token.
* The Snyk Group to which the Organizations belong should have [Single Sign On (SSO) configured](../../enterprise-setup/using-single-sign-on-sso-for-authentication/).
* Both the inviting user and the provisioned user must log in using SSO.
* The inviting user should have the permission `Provision Users` to invoke these calls. All Group and Org Admins by default have this permission.

<figure><img src="../../.gitbook/assets/Screenshot 2022-09-09 at 09.57.17.png" alt="Enable Provision Users permission"><figcaption><p>Enable Provision Users permission</p></figcaption></figure>

## How to use the Provision user API

The following explains how to use the Provision user endpoints. For more information, see the [API documentation](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization).

### Provision a user to the Organization

You provision a user to the specified Organization with a role. When a user first logs into Snyk, the user is automatically assigned the permissions as defined in the role.

**`POST`** `https://api.snyk.io/api/v1/org/orgId/provision`

**Request model:**

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN"`

`}`

**Response model:**

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN",`

`"created": Date`

`}`

{% hint style="info" %}
Enterprise plan users can define their own customized [member roles](../manage-permissions-and-roles/manage-member-roles.md) and can use `rolePublicId` for assignment.\
\
You can use either `role` or `rolePublicId` but not both in the same call.
{% endhint %}

### List pending user provisions

The List pending user permissions endpoint returns pending provisioned users in the response.

**`GET`** `https://api.snyk.io/api/v1/org/orgId/provision`

**Response model:**

`[`

`....`

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN",`

`"created": Date`

`},`

`....`

`]`

### Delete pending user provision

Use the Delete pending use providion endpoint to remove a pending provision request.

**`DELETE`** `https://api.snyk.io/api/v1/org/orgId/provision`

Query parameters

* email (string) - The email of the user.

**Response model:**

`{`

`"ok": true`

`}`
