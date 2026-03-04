# Manage service accounts using the Snyk API

You can manage service accounts using the [Snyk REST API](../../../snyk-api/reference/serviceaccounts.md).

{% hint style="info" %}
Specific permissions are required to perform all these tasks. For more information, visit [Service accounts-select a role](./#select-a-role).
{% endhint %}

## Service account attributes

`id` - The ID of the service account.

`name` - A human-friendly name for the service account.

`auth_type` - Authentication strategy for the service account. The following options are available:

* `access_token` - The service account uses a Snyk access token with expiry.
* `api_key` - The service account uses a legacy Snyk API key (no expiry).
* `oauth_client_secret` - The service account uses an [OAuth 2.0 access token](service-accounts-using-oauth-2.0.md#oauth-2.0-with-client-secret), retrieved with a client secret.
* `oauth_private_key_jwt` - The service account uses an [OAuth 2.0 access token](service-accounts-using-oauth-2.0.md#oauth-2.0-with-private-key-jwt), retrieved with a JWT signed with a private key.

`role_id` - The role of the service account, which defines the permissions it has. Available roles can be found using the endpoint [List all roles in a group](../../../snyk-api/reference/groups-v1.md#group-groupid-roles).

`access_token_expires_at` - The expiry date of the access token in ISO-8601 format, for example, `2025-08-16T00:00:00Z`. The maximum expiry is one year from creation time. Required only when `auth_type` is `access_token`.

`jwks_url` - A JWKs URL hosting your public keys used to verify signed JWT requests. This must be `https`. Required only when `auth_type` is `oauth_private_key_jwt`.

`access_token_ttl_seconds` - The time, in seconds, that a generated access token will be valid for. Defaults to one hour if unset. Required only when `auth_type` is `oauth_client_secret` or `oauth_private_key_jwt`.

## Manage Group-level service accounts

### Get a list of service accounts in your Group

**Request**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API endpoint:** [Get a list of group service accounts](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts-1)

This [paginated](../../../snyk-api/rest-api/about-the-rest-api.md#pagination) call returns an array of objects, each describing a service account.

### Create a service account for your Group

**Request**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API endpoint:** [Create a service account for a group](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts)

This call creates a new service account. You pass a `role_id` in the JSON-formatted body of the request, which defines the permissions a service account can use. This role id can be found using the endpoint [List all roles in a group](../../../snyk-api/reference/groups-v1.md#group-groupid-roles). Roles can be re-used for multiple service accounts.

### Get a service account from your Group

**Request**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Get a group service account](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts-serviceaccount_id-1)

This call returns details describing a specific service account.

### Update a service account in your Group

**Request**: `PATCH https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Update a group service account](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts-serviceaccount_id)

This call updates the details of a specific service account, at this time, the name of the service account.

### Delete a service account from your Group

**Request**: `DELETE https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Delete a group service account](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts-serviceaccount_id-secrets)

This call permanently deletes the specified service account and revokes its credentials.

### Manage a service account client secret for your Group

**Request**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}/secrets`

**API endpoint:** [Manage a group service account’s client secret](../../../snyk-api/reference/serviceaccounts.md#groups-group_id-service_accounts-serviceaccount_id-secrets)

This call allows you to manage the client secret for `oauth_client_secret` service accounts. You can perform the following operations:

* `create` - generate a new client secret. A service account can have a maximum of two active secrets at a time.
* `delete` - delete an existing client secret. This requires putting `client_secret` in the request body. Deleting an existing client secret would render it invalid. A service account must have at least one active secret; calling delete with your last secret will fail.
* `replace` - simultaneously delete the existing client secret and generate a new secret. This option is recommended if your `client_secret` is compromised.

## Manage Organization-level service accounts

### Get a list of service accounts in your Organization

**Request**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API endpoint:** [Get a list of organization service accounts](../../../snyk-api/reference/serviceaccounts.md#orgs-org_id-service_accounts-1)

This [paginated](../../../snyk-api/rest-api/about-the-rest-api.md#pagination) call returns an array of objects, each describing a service account.

### Create a service account for your Organization

**Request**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API endpoint:** [Create a service account for an organization](../../../snyk-api/reference/serviceaccounts.md#orgs-org_id-service_accounts)

This call creates a new service account. You pass a `role_id` in the JSON-formatted body of the request, which defines the permissions a service account can use. This `role id` can be found using the endpoint [List all roles in a group](../../../snyk-api/reference/groups-v1.md#group-groupid-roles). Roles can be re-used for multiple service accounts.

### Get a service account from your Organization

**Request**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Get an organization service account](../../../snyk-api/reference/serviceaccounts.md#orgs-org_id-service_accounts-serviceaccount_id-1)

This call returns details describing a specific service account.

### Update a service account in your Organization

**Request**: `PATCH https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Update an organization service account](../../../snyk-api/reference/serviceaccounts.md#patch-orgs-org_id-service_accounts-serviceaccount_id)

This call can update the details of a specific service account. The name of the service account is updated.

### Delete a service account from your Organization

**Request**: `DELETE https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API endpoint:** [Delete a service account in an organization](../../../snyk-api/reference/serviceaccounts.md#orgs-org_id-service_accounts-serviceaccount_id-2)

This call permanently deletes the specified service account.

### Manage a service account client secret for your Organization

**Request**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}/secrets`

**API endpoint:** [Manage an organization's service account's client secret](../../../snyk-api/reference/serviceaccounts.md#orgs-org_id-service_accounts-serviceaccount_id-secrets)

This call allows you to manage the client secret for `oauth_client_secret` service accounts. You can perform the following operations:

* `create` - generate a new client secret. A service account can have a maximum of two active secrets at a time.
* `delete` - delete an existing client secret. This requires putting `client_secret` in the request body. Deleting an existing client secret would render it invalid. A service account must have at least one active secret; calling delete with your last secret will fail.
* `replace` - simultaneously delete the existing client secret and generate a new secret. This option is recommended if your `client_secret` is compromised.
