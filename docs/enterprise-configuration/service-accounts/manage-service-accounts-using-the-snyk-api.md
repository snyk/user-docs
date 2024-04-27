# Manage service accounts using the Snyk API

You can manage service accounts using the [Snyk REST API](https://apidocs.snyk.io/#get-/groups/-group\_id-/service\_accounts).

{% hint style="info" %}
Specific permissions are required to perform all these tasks; see [Service accounts-select a role](./#select-a-role).
{% endhint %}

## Service account attributes

`id` - The ID of the service account.

`name` - A human-friendly name for the service account.

`auth_type` - Authentication strategy for the service account. The following options are available:

* `api_key` - The service account uses a regular Snyk API key.
* `oauth_client_secret` - The service account uses an [OAuth 2.0 access token](./#service-accounts-using-oauth-2.0), which is retrieved with a client secret.
* `oauth_private_key_jwt` - The service account uses an [OAuth 2.0 access token](./#service-accounts-using-oauth-2.0), which is retrieved with a JWT signed with a private key.

`role_id` - The role of the service account, which defines the permissions it has. Available roles can be found using the [List all roles in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/list-all-roles-in-a-group) Snyk API v1 endpoint.

`jwks_url` - A JWKs URL hosting your public keys used to verify signed JWT requests; this must be `https`. Required only when `auth_type` is `oauth_private_key_jwt`.

`access_token_ttl_seconds` - The time, in seconds, that a generated access token will be valid for. Defaults to 1 hour if unset. Required only when `auth_type` is `oauth_client_secret` or `oauth_private_key_jwt`.

## Manage Group-level service accounts

### Get a list of service accounts in your Group

**Request**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API documentation:** [https://apidocs.snyk.io/#get-/groups/-group\_id-/service\_accounts](https://apidocs.snyk.io/?version=2023-09-07#get-/groups/-group\_id-/service\_accounts)

This [paginated](../../snyk-api/make-calls-to-the-snyk-api/links-for-pagination-in-snyk-rest-api.md) call returns an array of objects, each describing a service account.

### Create a service account for your Group

**Request**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API documentation:** [https://apidocs.snyk.io/#post-/groups/-group\_id-/service\_accounts](https://apidocs.snyk.io/?version=2023-09-07#post-/groups/-group\_id-/service\_accounts)

This call creates a new service account. You pass a `role_id` in the JSON-formatted body of the request, which defines the permissions a service account can use. This role id can be found using the [List all roles in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/list-all-roles-in-a-group) Snyk API v1 endpoint. Roles can be re-used for multiple service accounts.

### Get a service account from your Group

**Request**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/#get-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#get-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

This call returns details describing a specific service account.

### Update a service account in your Group

**Request**: `PATCH https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/#patch-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#patch-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

This call updates the details of a specific service account, at this time, the name of the service account.

### Delete a service account from your Group

**Request**: `DELETE https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/#delete-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#delete-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

This call permanently deletes the specified service account and revokes its credentials.

### Manage a service account client secret for your Group

**Request**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}/secrets`

**API documentation:** [https://apidocs.snyk.io/#post-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-/secrets](https://apidocs.snyk.io/?version=2023-09-07#post-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-/secrets)

This call allows you to manage the client secret for `oauth_client_secret` service accounts. You can perform the following operations:

* `create` - generate a new client secret. A service account can have a maximum of two active secrets at a time.
* `delete` - delete an existing client secret. This requires putting `client_secret` in the request body. Deleting an existing client secret would render it invalid. A service account must have at least one active secret; calling delete with your last secret will fail.
* `replace` - simultaneously delete the existing client secret and generate a new secret. This option is recommended if your `client_secret` is compromised.

## Manage Organization-level service accounts

### Get a list of service accounts in your Organization

**Request**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API documentation:** [https://apidocs.snyk.io/#get-/orgs/-org\_id-/service\_accounts](https://apidocs.snyk.io/?version=2023-09-07#get-/orgs/-org\_id-/service\_accounts)

This [paginated](../../snyk-api/make-calls-to-the-snyk-api/links-for-pagination-in-snyk-rest-api.md) call returns an array of objects, each describing a service account.

### Create a service account for your Organization

**Request**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API documentation:** [https://apidocs.snyk.io/#post-/orgs/-org\_id-/service\_accounts](https://apidocs.snyk.io/?version=2023-09-07#post-/orgs/-org\_id-/service\_accounts)

This call creates a new service account. You pass a `role_id` in the JSON-formatted body of the request, which defines the permissions a service account can use. This `role id` can be found using the [List all roles in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/list-all-roles-in-a-group) Snyk API v1 endpoint. Roles can be re-used for multiple service accounts.

### Get a service account from your Organization

**Request**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/#get-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#get-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

This call returns details describing a specific service account.

### Update a service account in your Organization

**Request**: `PATCH https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/?version=2023-09-07#patch-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#patch-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

This call updates the details of a specific service account, at this time, the name of the service account.

### Delete a service account from your Organization

**Request**: `DELETE https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API documentation:** [https://apidocs.snyk.io/#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2023-09-07#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

This call permanently deletes the specified service account.

### Manage a service account client secret for your Organization

**Request**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}/secrets`

**API documentation:** [https://apidocs.snyk.io/#post-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-/secrets](https://apidocs.snyk.io/?version=2023-09-07#post-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-/secrets)

This call allows you to manage the client secret for `oauth_client_secret` service accounts. You can perform the following operations:

* `create` - generate a new client secret. A service account can have a maximum of two active secrets at a time.
* `delete` - delete an existing client secret. This requires putting `client_secret` in the request body. Deleting an existing client secret would render it invalid. A service account must have at least one active secret; calling delete with your last secret will fail.
* `replace` - simultaneously delete the existing client secret and generate a new secret. This option is recommended if your `client_secret` is compromised.
