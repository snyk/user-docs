# Set up the authorization code exchange

After you receive an authorization **code**, you must exchange it for an access token.

To request an access token, make a POST request to the token endpoint (more details in the [API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-tokens/token-exchange-&-refresh)):

```
https://api.snyk.io/oauth2/token
```

with the following properties in a x-www-form-urlencoded formatted request body:

```
grant_type=authorization_code
&code=(code from the previous step)
&redirect_uri=(redirect URI from the previous step)
&client_id=(clientId from the app creation),
&client_secret=(clientSecret from the app creation)
```

The response includes details necessary for your app to communicate with the Snyk APIs: **access\_token** and **refresh\_token**, as well as their expiry.

**Both tokens must be stored in a secured data store. It is highly recommended to encrypt the values before storing them.**

The **access\_token** will be used for future API calls on behalf of the user and organization. The access-token has a much shorter lifespan than the **refresh\_token**.

The **refresh\_token** can be used a single time to get a new **access\_token** when it expires. In other words, the **refresh\_token** will no longer be usable if its own expiry time passes or if it is used to refresh the **access\_token**.
