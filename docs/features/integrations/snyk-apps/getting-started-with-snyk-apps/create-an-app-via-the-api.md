# Create an App

## Create using API

Now that you have an API token and **orgId**, you can use the Snyk API to create your Snyk App by sending a **POST** request to the `apps` endpoint:

```
https://api.snyk.io/rest/orgs/{orgId}/apps?version={version}
```

For more information see the [API documentation](https://apidocs.snyk.io/#post-/orgs/-org\_id-/apps).

The request body should contain the details for your new App, including the App's **name**, **redirectURIs**, and [**scopes**](create-an-app-via-the-api.md#requesting-scopes).

The response includes details necessary to complete the integration: **clientId** and **clientSecret**. You will use these values with our API endpoints within your App, so consider storing them as part of your App's configuration.

{% hint style="danger" %}
Never share the **clientSecret** publicly, as this is used to authenticate your App. This is also the only time you’ll be able to get the **clientSecret**, so keep it secure and private. If you lose it or if the secret is leaked, you can [rotate your App's **clientSecret**](managing-app-details.md#rotate-app-client-secret).
{% endhint %}

## Create using Snyk CLI

The [Snyk CLI](../../../../snyk-cli/) can be used to create the Snyk App. If you don't already have the Snyk CLI installed on your machine, you can easily do so by following our [getting started with Snyk CLI docs](../../../../snyk-cli/getting-started-with-the-cli/). See [Create a Snyk App using the CLI](../../../../snyk-cli/create-a-snyk-app-using-the-snyk-cli.md) for details.

{% hint style="warning" %}
All `apps` commands are only accessible behind the `--experimental` flag and the behaviour can change at any time without prior notice. You are advised to use all the commands with caution.

If you are already authenticated with Snyk CLI (`snyk auth`) you can use the `apps` sub-commands as is. Otherwise, first authenticate using `snyk auth` and then use the Snyk Apps commands.
{% endhint %}

Snyk Apps related sub-commands are under the `snyk apps` command.

Use the `create` sub-command to create the Snyk Apps. There are two ways to use the command.

THe first is the normal mode, for example:

```
snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='My Awesome App' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta
```

*   `--org-id=<ORG_ID>`

    Required for the `create` command. Specify the `<ORG_ID>` to create the Snyk App under.
*   `--name=<SNYK_APP_NAME>`

    Required for the `create` command. This name will be displayed to the end user later on, when the user authorizes your App.
*   `--redirect-uris=<REDIRECT_URIS>`

    Required for the `create` command. A comma-separated list of redirect URIs. This forms a list of allowed redirect URIs to call back after authentication.
*   `--scopes=<SCOPES>`

    Required for the `create` command. A comma-separated list of scopes that are required by your Snyk App. This will form a list of scopes that your app is allowed to request during authorization.

The second is the interactive mode, which prompts you to enter all the values in a similar fashion as the normal mode. The following is an example of the interactive mode:

```
snyk apps create --experimental --interactive

? Name of the Snyk App (visible to users when they install the Snyk App)? My Awesome Snyk App
? Your Snyk App's redirect URIs (comma-separated list.  Ex: https://example1.com,https://example2.com)?:  https://example1.com
? Your Snyk App's permission scopes (comma-separated list.  Ex: apps:beta)?:  apps:beta
? Please provide the org id under which you want to create your Snyk App:  48ebb069-472f-40f4-b5bf-d2d103bc02d4
```

## Requesting scopes

Scopes define the permissions your Snyk App has to perform actions in a user’s account. When a user authorizes your App to access their Snyk account, they see the list of scopes you are requesting and then decide whether or not they approve the connection.

When deciding which scopes your Snyk App will need, consider the actions your App will be performing. It may seem better to request every available scope, but users may refuse to install an App that asks for more permissions than required. Also, a user installing your App will not be able to complete the authorization process if they don’t have all the permissions matching the scopes you request.

The following lists the **available scopes**.

| Scope     | Description                                                                                                                                               |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| apps:beta | Allows the App to test and monitor your existing projects, as well as read information on your Snyk Organization, existing projects, issues, and reports. |

## Generally Available

{% hint style="warning" %}
As Snyk move towards more granular scopes, Snyk is allowing Snyk Apps to be created with more granular scopes. This is to promote the principle of least privilege. Refer to [Snyk Apps Scopes](../snyk-apps-scopes.md) for the list of scopes available behind a feature flag (`appsGranularPermissions`). Being behind a feature flag means they are still in active development. If you would like to test them out as they give a look into how Snyk Apps scopes will look in the future, contact the Snyk [support team](https://snyk.io/contact-us/).
{% endhint %}

{% hint style="info" %}
You cannot currently update scopes for a Snyk App after it has been created. If you change your mind about which scopes you need during the App development process, create a new Snyk App with a new list of scopes, and replace the clientId and clientSecret in your App’s configuration. If users have installed the App already, the users will need to authorize the new App with their Snyk account.
{% endhint %}
