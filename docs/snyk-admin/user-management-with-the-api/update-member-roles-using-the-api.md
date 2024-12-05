# Update member roles using the API

To migrate members of existing organizations to new roles, you must use the API. Follow the steps in each section of this document.

Snyk recommends running these updates with bounded concurrency in batches so as not to trip any of the rate limiters. Optimally, perform the updates in batches of **ten** concurrent requests at a time.

## Step 1: Get a list of roles in your Group

**Request**: `GET https://api.snyk.io/v1/group/{groupId}/roles`

**Endpoint:** [List all roles in a group](../../snyk-api/reference/groups-v1.md#group-groupid-roles)

This call returns an array of objects, each describing custom and non-custom (for example, default) roles. Default roles are indicated by the `customRole: false` property. Save the`publicId` of each role you want to assign to a user.

## Step 2: Get a list of Organization members

**Request**: `GET https://api.snyk.io/v1/org/{orgId}/members`

**Endpoint**: [List members](../../snyk-api/reference/organizations-v1.md#org-orgid-members)

This call returns an array of all non-admin members of the Organization. Save the `id` of each user who should have a new role.

Service accounts are not returned by the List members endpoint. You must get the `publicID` of each service account from the Service Account Settings page:

* In your **Service Accounts Settings**, select the **name of the service account** for which you want to get the `publicID`.
* When the **Edit account name** window opens, copy the string at the end of the URL; this is the `publicID` of the service account.

## Step 3: Update the role of users

**Request**: `PUT https://api.snyk.io/v1/org/{orgId}/members/update/{userId}`

**Endpoint**: [Update a member's role in the organization](../../snyk-api/reference/organizations-v1.md#org-orgid-members-update-userid)

For each user, call the endpoint Update a member's role in the organization to set the member's new role, using the user id and role id you collected previously.

You pass a `rolePublicId` in the JSON-formatted body of the request. This is the role `publicId` you saved in the first step.

For a successful request, the response is `200 OK`.

You can verify the change on the Organization members page (for humans) or the Service Account Settings page (for robots).
