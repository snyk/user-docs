# Managing App details

## List Apps

To view a list of Snyk Apps owned by your Snyk organization, send a **GET** request to the `apps` endpoint (for details, see the [API documentation](https://apidocs.snyk.io/#get-/orgs/-org\_id-/apps)).

```
https://api.snyk.io/rest/orgs/{orgId}/apps?version={version}
```

## Rotate App clientSecret

You cannot view the **clientSecret** after the App is created. If you have misplaced it, you can rotate your **clientSecret** and receive a new one.

All secret management requests are performed by sending a POST request to the endpoint `/apps/{clientId}/secrets`. For more details refer to the [API documentation](https://apidocs.snyk.io/#post-/orgs/-org\_id-/apps/-client\_id-/secrets). The clientId can be found using the [List Apps endpoint](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/apps).

```
https://api.snyk.io/rest/orgs/{orgId}/apps/{clientId}/secrets?version={version}
```

There are currently three operations that can be performed which are indicated by the body of your POST request:

* create `{"mode": "create"}`
* delete `{"mode": "delete", "secret": "{clientSecret}"}`
* replace `{"mode": "replace"}`

Snyk recommends you adopt the following procedure when rotating your secrets:

1. Create new secret using `{"mode": "create"}`
2. Update your services with the newly generated secret
3. Remove the old secret using `{"mode": "delete", "secret": "{secret}"}`

#### Create a clientSecret

In normal operation it is recommended that you periodically rotate your client secrets. To start the process, send the request body `{"mode": "create"}` to the endpoint which will create a new secret. The returned value of this call will be your app with the new generated secret. Both the new secret and any existing secrets will be valid until they are manually replaced or deleted. You can also immediately replace a client secret.

An App can have a maximum of two active secrets at any time. This endpoint fails if you try to call create when you already have the maximum number of secrets active.

#### Delete a clientSecret

To clean up any unused secrets call the endpoint with `{"mode": "delete", "secret": "{clientSecret}"}` where `{clientSecret}` is your client secret that you want to delete. This action invalidates the secret immediately so it can no longer be used.

An App must have at least one active secret; calling delete with your last secret will fail.

#### Replace a clientSecret

In the event that your Apps clientSecret is leaked, you can generate a new one by using `{"mode": "replace"}`.

When you replace your clientSecret, your current secret is immediately invalid. Your App will not be able to connect to Snyk until you update the App's configuration with the new secret.

### Update App details

You can update your App's name, or the list of redirect URIs you have set.

To update an App, send a **PATCH** request to the `apps/{clientId}` endpoint (for details, see the [API documentation](https://apidocs.snyk.io/#patch-/orgs/-org\_id-/apps/-client\_id-)). The clientId can be found using the [List Apps endpoint](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/apps).

```
https://api.snyk.io/rest/orgs/{orgId}/apps/{clientId}?version={version}
```

### Delete an App

To delete an App from your Snyk Organization, send a **DELETE** request to the `apps` endpoint (for details, see the [API documentation](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#delete-/orgs/-org\_id-/apps/-client\_id-)):

```
https://api.snyk.io/rest/orgs/{orgId}/apps?version={version}
```

{% hint style="danger" %}
Deleting an App will revoke your App credentials and remove all of your App's installations. If you have active users, they will no longer be able to connect to Snyk through the App.
{% endhint %}
