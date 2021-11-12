# Create an App via the API

Now that you have an API token and **orgId**, you can use our API to create your Snyk App by sending a **POST** request to the `apps` endpoint (more details found in the [API documentation](https://snykv3.docs.apiary.io/#reference/apps/app-management/create-an-app)):

```
https://api.snyk.io/v3/orgs/{orgId}/apps?version={version}
```

The request body should contain the details for your new App, including the App's **name**, **redirectURIs**, and [**scopes**](create-an-app-via-the-api.md#requesting-scopes).

The response includes details necessary to complete the integration: **clientId** and **clientSecret**. You will use these values with our API endpoints within your App, so consider storing them as part of your App's configuration.

{% hint style="danger" %}
Never share the **clientSecret** publicly, as this is used to authenticate your App. This is also the only time you’ll be able to get the **clientSecret**, so keep it secure and private. If you lose it or if the secret is leaked, you can [rotate your App's **clientSecret**](managing-app-details.md#rotate-app-client-secret).
{% endhint %}

## Requesting scopes

Scopes define the permissions your Snyk App has to perform actions in a user’s account. When a user authorizes your App to access their Snyk account, they see the list of scopes you are requesting, and then decide whether or not they approve the connection.&#x20;

When deciding which scopes your Snyk App will need, consider the actions your App will be performing. It may seem better to request every available scope, but users may refuse to install an App that asks for more permissions than required. Also, a user installing your App will not be able to complete the authorization process if they don’t have all the permissions matching the scopes you request.

#### Available Scopes

| Scope     | Description                                                                                                                                               |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| apps:beta | Allows the App to test and monitor your existing projects, as well as read information on your Snyk Organization, existing projects, issues, and reports. |

{% hint style="info" %}
You cannot currently update scopes for a Snyk App after it has been created. If you change your mind about which scopes you need during the App development process, create a new Snyk App with a new list of scopes, and replace the clientId and clientSecret in your App’s configuration. If any user has installed the App already, they will need to authorize the new App with their Snyk account.
{% endhint %}
