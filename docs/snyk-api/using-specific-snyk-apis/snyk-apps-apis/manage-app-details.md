# Manage App details

## List Apps created by an Organization

To view a list of Snyk Apps owned by your Snyk Organization, send a `GET` request to the `apps/creations` endpoint:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations?version={version}`

For details, see the endpoint [Get a list of apps created by an Organization](../../reference/apps.md#orgs-org_id-apps-1).

## Update App details

You can update the name of your App or the list of redirect URIs you have set.

To update an App, send a `PATCH` request to the `apps/creations{app_id}` endpoint:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations{app_id}?version={version}`

The `app_id` path parameter is the `id` in the response to a [`GET` request to the `apps/creations` endpoint](manage-app-details.md#list-apps-created-by-an-organization).

For details, see the endpoint[Update app creation attributes such as name, redirect URLs, and access token time to live using the App ID](../../reference/apps.md#orgs-org_id-apps-creations-app_id).

## Delete an App

To delete an App from your Snyk Organization, send a DELETE request to the endpoint `apps/creations{app_id}`:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations/{app_id}?version={version}`

The `app_id` path parameter is the `id` in the response to a [`GET` request to the `apps/creations` endpoint](manage-app-details.md#list-apps-created-by-an-organization).

For details, see the endpoint [Delete an app by its App ID](../../reference/apps.md#orgs-org_id-apps-client_id-2).

Deleting an App revokes your App credentials and removes all of your App's installations. If you have active users, they will no longer be able to connect to Snyk through the App.

## Rotate App clientSecret

You cannot view the `clientSecret` after the App is created. If you have misplaced it, you can rotate your `clientSecret` and receive a new one.

Perform secret management requests for apps you have created by sending a `POST` request to the endpoint `apps/creations{app_id}/secrets`:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations/{app_id}/secrets?version={version}`

The `app_id` path parameter is the `id` in the response to a [`GET` request to the `apps/creations` endpoint](manage-app-details.md#list-apps-created-by-an-organization).

For details, see the endpoint [Manage client secret for the Snyk App](../../reference/apps.md#orgs-org_id-apps-creations-app_id-secrets).

{% hint style="info" %}
For client credentials apps that you have installed, see the endpoint [Manage client secret for non-interactive Snyk App installations](../../reference/apps.md#groups-group_id-apps-installs-install_id-secrets).
{% endhint %}

You can perform three operations that are indicated by the body of your POST request:

* create `{"mode": "create"}`
* delete `{"mode": "delete", "secret": "{clientSecret}"}`
* replace `{"mode": "replace"}`

Snyk recommends you adopt the following procedure when rotating your secrets:

1. Create a new secret using `{"mode": "create"}`
2. Update your services with the newly generated secret
3. Remove the old secret using `{"mode": "delete", "secret": "{secret}"}`

### Create a clientSecret

It is recommended that in normal operation you periodically rotate your client secrets. To start the process, send the request body `{"mode": "create"}` to the endpoint which will create a new secret. The returned value of this call will be your app with the newly generated secret. Both the new secret and any existing secrets will be valid until they are manually replaced or deleted. You can also immediately replace a client secret.

An App can have a maximum of two active secrets at any time. This endpoint fails with the error  `cannot update secrets` if you try to call `create` when you already have the maximum number of secrets active.

### Replace a clientSecret

In the event that your App's `clientSecret` is leaked, you can generate a new one by using `{"mode": "replace"}`.

When you replace your `clientSecret`, your current secret is immediately invalid. Your App will not be able to connect to Snyk until you update the App's configuration with the new secret.

### Delete a clientSecret

To clean up any unused secrets, call the endpoint with `{"mode": "delete", "secret": "{clientSecret}"}` where `{clientSecret}` is your client secret that you want to delete. This action invalidates the secret immediately so it can no longer be used.

An App must have at least one active secret; calling delete with your last secret will fail.
