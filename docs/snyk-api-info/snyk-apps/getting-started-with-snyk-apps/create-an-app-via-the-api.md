# Create an App

## Create using API

Now that you have an API token and **orgId**, you can use the Snyk API to create your Snyk App by sending a **POST** request to the `apps` endpoint:

```
https://api.snyk.io/rest/orgs/{orgId}/apps?version={version}
```

For more information see the [API documentation](https://apidocs.snyk.io/#post-/orgs/-org\_id-/apps). See the example request at the end of this section.

The request body should contain the details for your new App, including the App's **name**, **redirectURIs**, and [**scopes**](create-an-app-via-the-api.md#requesting-scopes).

The response includes details necessary to complete the integration: **clientId** and **clientSecret**. You will use these values with Snyk API endpoints within your App, so consider storing them as part of your App's configuration.

{% hint style="danger" %}
Never share the **clientSecret** publicly, as this is used to authenticate your App. This is also the only time you’ll be able to get the **clientSecret**, so keep it secure and private. If you lose it or if the secret is leaked, you can [rotate your App's **clientSecret**](managing-app-details.md#rotate-app-client-secret).
{% endhint %}

Example CURL request to create a Snyk App:

```
curl -X POST -H "Content-Type: application/vnd.api+json" \                                 
-H "Authorization: token <REPLACE_WITH_API_TOKEN>" \
-d '{"name": "My Awesome Snyk App", "redirect_uris": ["https://example.com/callback"], "scopes": ["org.read"]}' \
https://api.snyk.io/rest/orgs/<REPLACE_WITH_YOUR_ORGID>/apps?version=2023-11-06
```

## Create using Snyk CLI

The [Snyk CLI](../../../snyk-cli/) can be used to create the Snyk App. If you do not already have the Snyk CLI installed on your machine, you can install it by following the steps in [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-snyk-cli.md). See [Create a Snyk App using the CLI](../create-a-snyk-app-using-the-snyk-cli.md) for details.

{% hint style="warning" %}
All `apps` commands are only accessible behind the `--experimental` flag and the behavior can change at any time without prior notice. You are advised to use all the commands with caution.

If you are already authenticated with Snyk CLI (`snyk auth`) you can use the `Snyk Apps` commands. Otherwise, first, authenticate using `snyk auth` and then use the Snyk Apps commands.
{% endhint %}

Snyk Apps commands are prefaced with snyk apps.

Use the `snyk apps create` command to create your Snyk Apps. There are two ways to use the command.

The first is the normal mode, for example:

```
snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='My Awesome App' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta
```

*   `--org-id=<ORG_ID>`

    Required for the `create` command. Specify the `<ORG_ID>` to create the Snyk App under.
*   `--name=<SNYK_APP_NAME>`

    Required for the `create` command. This name will be displayed to the end-user later on when the user authorizes your App.
*   `--redirect-uris=<REDIRECT_URIS>`

    Required for the `create` command. A comma-separated list of redirect URIs. This forms a list of allowed redirect URIs to call back after authentication.
*   `--scopes=<SCOPES>`

    Required for the `create` command. A comma-separated list of scopes that are required by your Snyk App. This will form a list of [scopes](create-an-app-via-the-api.md#requesting-scopes) that your app is allowed to request during authorization.

The second is the interactive mode, which prompts you to enter all the values in a similar fashion as the normal mode. The following is an example of the interactive mode:

```
snyk apps create --experimental --interactive

? Name of the Snyk App (visible to users when they install the Snyk App)? My Awesome Snyk App
? Your Snyk App's redirect URIs (comma-separated list.  Ex: https://example1.com,https://example2.com)?:  https://example1.com
? Your Snyk App's permission scopes (comma-separated list.  Ex: org.read)?:  apps:beta
? Please provide the org id under which you want to create your Snyk App:  48ebb069-472f-40f4-b5bf-d2d103bc02d4
```

## Requesting scopes

Scopes define the permissions your Snyk App has to perform actions in a user’s account. When a user authorizes your Snyk App to access their Snyk account, they see the list of scopes you are requesting and then decide whether or not they approve the connection.

When deciding which scopes your Snyk App will need, consider the actions your App will be performing. It may seem better to request every available scope, but users may refuse to install an App that asks for more permissions than required. Also, a user installing your App will not be able to complete the authorization process if they don’t have all the permissions matching the scopes you request.

The following lists the **available scopes**.

{% hint style="warning" %}
`org.read` is a mandatory scope and should always be included.
{% endhint %}

| Scope                         | Description                                                                 |
| ----------------------------- | --------------------------------------------------------------------------- |
| org.read                      | View organization information and settings                                  |
| org.edit                      | Edit organization information and settings                                  |
| org.report.read               | View reports in your organization                                           |
| org.project.create            | Add new projects                                                            |
| org.project.read              | View project information and settings and view organization targets         |
| org.project.edit              | Edit project information                                                    |
| org.project.delete            | Permanently remove projects and permanently remove organization targets     |
| org.project.status            | Activate and deactivate projects                                            |
| org.project.test              | Test projects                                                               |
| org.project.ignore.create     | Create new project ignores                                                  |
| org.project.ignore.read       | View project ignore information                                             |
| org.project.ignore.edit       | Configure project ignores                                                   |
| org.project.ignore.delete     | Permanently remove project ignores                                          |
| org.project.tag.edit          | Create, apply and remove project tags                                       |
| org.project.pr.create         | Create fix pull requests for projects                                       |
| org.project.pr.skip           | Skip failed security tests on pull requests by marking checks as successful |
| org.project.jira.issue.read   | View Jira issue information                                                 |
| org.project.jira.issue.create | Create new Jira issues                                                      |
| org.package.test              | Test packages in our supported ecosystems                                   |

{% hint style="info" %}
You cannot currently update scopes for a Snyk App after it has been created. If you change your mind about which scopes you need during the App development process, create a new Snyk App with a new list of scopes, and replace the clientId and clientSecret in your App’s configuration. If users have installed the Snyk App already, the users will need to authorize the new App with their Snyk account.
{% endhint %}
