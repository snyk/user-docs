# Set up to authorize users

When a user connects their Snyk account to your App, they must authorize access to their chosen organization or group and approve the requested scopes. This process is started when you direct users to the Snyk Apps authorization page and pass the appropriate parameters:

```
https://app.snyk.io/oauth2/authorize?response_type=code&client_id={clientId}&redirect_uri={redirectURI}&scope={scopes}&nonce={nonce}&state={state}&version={version}
```

{% hint style="warning" %}
Note that this is a webpage link and not an API endpoint.
{% endhint %}

The current **version** can be found in the Snyk [API documentation](https://snykoauth2.docs.apiary.io/#reference/apps/app-authorization/authorize-an-app).

The **scopes** and the **redirect\_uri** must match what was defined when the App was created.

The **state** value is used to carry any App-specific state from this `/authorize` call to the callback on the **redirect\_uri** (such as a user’s id). It must be verified in your callback to [prevent CSRF attacks](https://datatracker.ietf.org/doc/html/rfc6749#section-10.12).

The **nonce** value is a highly randomized string stored alongside a timestamp on the app side before calling `/authorize`, then verified on the returned access token. For more information see [The OAuth 2.0 Authorization Framework Access Token Types](https://datatracker.ietf.org/doc/html/rfc6749#section-7.1).

![](<../../../.gitbook/assets/image (118) (1) (1).png>)

After the connection is complete, the user is redirected to the specified redirect URI with query string parameters **code** and **state** added on, which are necessary for the next step.
