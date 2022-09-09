# Provision users to Orgs via API

These API endpoints allow you to organise and grant permissions to your SSO users, before the users log into the Snyk platform.&#x20;

The provisioned users will not need to accept invites; when they first log into Snyk, they will automatically have all permissions ready for them. This provides an alternative to add users to orgs at scale ahead of the first login.

### Pre-requisites

* The user being provisioned must not already exist in our system.
* Inviting user must call the API with their personal token.
* The Snyk Group to which orgs belong should have [Single Sign On (SSO) configured](../setting-up-sso-for-authentication/).
* Both inviting user and provisioned user must log in via SSO.
* The API does not support service accounts to be used as the inviting user or provisioned user.
*   The inviting user should have the permission `Provision Users` to invoke these calls. All Group and Org Admins by default have this permission.\


    <figure><img src="../../../.gitbook/assets/Screenshot 2022-09-09 at 09.57.17.png" alt=""><figcaption></figcaption></figure>

### Working with provision user API



{% embed url="https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization" %}
Click to access API Documentation
{% endembed %}

#### 1. Provision a user to the organization

Provision user to specified org with a role. When a user first logs into Snyk, they will automatically be assigned with the permissions as defined in the role.

**`POST`** `https://api.snyk.io/api/v1/org/orgId/provision`

**Request model:**&#x20;

`{`&#x20;

`"email": "test@example.com",`&#x20;

`"rolePublicId": "",`&#x20;

&#x20;`"role": "ADMIN"`&#x20;

`}`

**Response model:**&#x20;

`{`

`"email": "test@example.com",`&#x20;

`"rolePublicId": "",`

`"role": "ADMIN",`&#x20;

`"created": Date`&#x20;

`}`

{% hint style="info" %}
For Business plan users, the role must be one of `ADMIN`, `COLLABORATOR`, or `RESTRICTED_COLLABORATOR`.&#x20;

Enterprise plan users can define their own customized [member role](member-roles.md) and can use `rolePublicId` for assignment.\
\
You can use either `role` or `rolePublicId` but not both in the same call.
{% endhint %}

#### 2. List pending user provisions

The following endpoint will return pending provisioned users in their response.

**`GET`** `https://api.snyk.io/api/v1/org/orgId/provision`

**Response model:**&#x20;

`[`

`....`

`{`

`"email": "test@example.com",`&#x20;

`"rolePublicId": "",`

`"role": "ADMIN",`&#x20;

`"created": Date`&#x20;

`},`

`....`

`]`

#### 3. Delete pending user provision

Remove pending provision request.

**`DELETE`** `https://api.snyk.io/api/v1/org/orgId/provision`

Query parameters

* email (string) - The email of the user.

**Response model:**&#x20;

`{`

`"ok": true`

`}`
