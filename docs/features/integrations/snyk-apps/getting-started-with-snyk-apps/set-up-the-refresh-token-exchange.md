# Set up the refresh token exchange

As the **access\_token** will expire in a short time, the App will need to frequently request a new one using the **refresh\_token**. This must be done while the **refresh\_token** itself is still valid.

To exchange for a fresh **access\_token**, make a POST request to the token endpoint \(more details found in the [API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-tokens/token-exchange-&-refresh)\):

```text
https://api.snyk.io/oauth2/token
```

as the `refresh_token` grant type, with the following properties in the request body:

```text
{
    grant_type: refresh_token,
    client_id: clientId - from when the app was created,
    client_secret: clientSecret - from when the app was created,
    refresh_token: refresh_token - from the previous step
}
```

The response to this call provides a new **access\_token**, **refresh\_token**, and expiry for each.

