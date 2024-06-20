# Remove members from Groups and Orgs using the Snyk REST and v1 API

To remove members from Groups and Organizations programmatically, you can use the Snyk API v1 and the Snyk REST API, as explained in the steps that follow.

You cannot use the following API calls to remove Service Accounts. The API calls are for use to remove human user accounts only.

## Remove Organization memberships

### Step 1: Get a list of Organization members

**Request**: `GET https://snyk.io/api/v1/org/{orgId}/members`

**API documentation**: [https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members](https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members)

This call returns an array of all non-admin members of the Organization. Save the `id` of each user who needs to be removed from the Organization.

### Step 2: Remove a member from an Organization

**Request**: `DELETE https://snyk.io/api/v1/org/{orgId}/members/update/{userId}`

**API documentation**: [https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/remove-a-member-from-the-organization](https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/remove-a-member-from-the-organization)

For each user, call the endpoint to remove that member from the Organization using the user id you saved previously.

For a successful request, the response is `200 OK`.

Look at the Organization members page to verify that the member has been removed.

{% hint style="info" %}
When a member is removed from an Organization, if the Organization is a part of a Group, the user continues to exist in the Group as a Group Member. To completely remove the user from the Group, follow the steps in the next section.
{% endhint %}

## Remove Group memberships

### Step 1: Get a list of Group members

**Request**: `GET https://api.snyk.io/api/v1/group/groupId/members`

**API documentation**: [https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group/list-all-members-in-a-group](https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group/list-all-members-in-a-group)

This call returns an array of all members of the Group. Save the `id` of each user who needs to be removed from the Group.

### Step 2: Remove a member from a Group

**Request**: PATCH https://api.snyk.io/rest/groups/{group\_id}/users/{id}?version=2023-01-30\~beta

**Body:**

```postman_json
{
    "data": {
        "attributes": {
            "membership": null
        },
        "id": "<user-id>",
        "type": "user"
    }
}
```

**API documentation**: [https://apidocs.snyk.io/?version=2023-01-30%7Ebeta#patch-/groups/-group\_id-/users/-id-](https://apidocs.snyk.io/?version=2023-01-30%7Ebeta#patch-/groups/-group\_id-/users/-id-)

For each user, call the endpoint to remove that member from the Group using the user id you saved previously.

For a successful request, the response is `200 OK`.

Look at the Group members page to verify the user has been removed.

{% hint style="info" %}
When a member is removed from a Group, the user continues to exist in Snyk. To completely delete all data associated with the user, follow the step in the next section.
{% endhint %}

## Delete Group members

When an SSO connection is associated with only one Group, the following call can completely delete a Group member from the system. This delete action also complies with the GDPR (General Data Protection Regulation) requirements.

**Request**: DELETE `https://api.snyk.io/rest/groups/{group_id}/sso_connections/{sso_id}/users/{user_id}?version=2023-01-30~beta`

**API documentation**: [https://apidocs.snyk.io/?version=2023-01-30%7Ebeta#delete-/groups/-group\_id-/sso\_connections/-sso\_id-/users/-user\_id-](https://apidocs.snyk.io/?version=2023-01-30%7Ebeta#delete-/groups/-group\_id-/sso\_connections/-sso\_id-/users/-user\_id-)

You can find the `{sso_id}` on the Snyk Web UI; navigate to **Group** >**Settings** >**SSO** >**Step 3**. If you need help, reach out to your Account team.

<figure><img src="../../.gitbook/assets/Screenshot 2023-02-22 at 10.27.19.png" alt="Self Serve SSO screen, Step 3, sso_id highlighted"><figcaption><p>Self Serve SSO screen, Step 3, sso_id highlighted</p></figcaption></figure>

For a successful request, the response is `200 OK`.

Use the following request to verify the member has been deleted.\
`GET https://api.snyk.io/api/v1/user/userId`
