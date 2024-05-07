# Revoke compromised refresh tokens

If you believe that a refresh token has been compromised then it is recommended that you revoke that token as soon as possible. This is to prevent misuse of the token to gain access to Snyk systems.&#x20;

To do this issue a POST request to the revoke endpoint:

```
https://api.snyk.io/oauth2/revoke
```

with the following properties in a x-www-form-urlencoded formatted request body:

```
token=(refresh token to be revoked)
&client_id=(clientId from the app creation)
&client_secret=(clientSecret from the app creation)
```

This endpoint will return an empty response on success and the given refresh token will immediately be revoked.
