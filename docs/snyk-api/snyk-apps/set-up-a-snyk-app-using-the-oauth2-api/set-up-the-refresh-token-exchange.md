# Set up the refresh token exchange

As the `access_token` will expire in a short time, the App will need to frequently request a new one using the `refresh_token`. This must be done while the `refresh_token` itself is still valid.

To exchange for a fresh `access_token`, make a POST request to the token endpoint:

```
https://api.snyk.io/oauth2/token
```

with the following properties in a x-www-form-urlencoded formatted request body:

```
grant_type=refresh_token
&refresh_token=(refresh token from the previous step)
&client_id=(clientId from the app creation)
&client_secret=(clientSecret from the app creation)
```

The response to this call provides a new `access_token`, `refresh_token`, and expiry for each.

Be sure to store the new `refresh_token`, as the old one is now invalid. Be aware that concurrent calls to this endpoint may result in refresh tokens being marked invalid unexpectedly, make sure to only request a new refresh token pair once per given refresh token. If this process fails then a user will need to perform the authorization flow again.
