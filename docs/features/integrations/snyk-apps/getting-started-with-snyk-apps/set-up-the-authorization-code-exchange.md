# Set up the authorization code exchange

After you receive an authorization **code**, you must exchange it for an access token.

To request an access token, make a POST request to the token endpoint (more details in the [API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-tokens/token-exchange-&-refresh)):

```
https://api.snyk.io/oauth2/token
```

as the `authorization_code` grant type, with the following properties in the request body:

```
{
    grant_type: authorization_code,
    code: code - from the previous step,
    redirect_uri: redirect URI - from the previous step,
    client_id: clientId - from when the app was created,
    client_secret: clientSecret - from when the app was created
}
```

The response includes details necessary for your app to communicate with the Snyk APIs: **access_token** and **refresh_token**, as well as their expiry.** **Both tokens must be stored in a secured data store. It is highly recommended to encrypt the values before storing them.

The **access_token** will be used for future API calls on the user & organization’s behalf. It has a much shorter lifespan as compared to the **refresh_token**.

The **refresh_token** can be used a single time to get a new **access_token** when it expires. In other words, the **refresh_token** will no longer be usable if its own expiry time passes or if it’s used to refresh the **access_token**.
