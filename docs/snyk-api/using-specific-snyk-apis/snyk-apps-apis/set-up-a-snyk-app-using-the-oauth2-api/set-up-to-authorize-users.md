# Set up to authorize users

When users connect their Snyk account to your App, they must authorize access to their chosen Organization or Group and approve the requested scopes. This process starts when you direct users to the Snyk Apps authorization web page and pass the appropriate parameters: `https://app.snyk.io/oauth2/authorize?response_type=code&client_id={clientId}&redirect_uri={redirectURI}&state={state}&code_challenge={codeChallenge}&code_challenge_method=S256`

If your Snyk App is configured with multiple redirect URIs, passing `redirect_uri` will select which redirect URI to use. The given value must match exactly one of the values that are defined for your Snyk App. If not specified the first value defined on the Snyk App is used.

The `state` value is used to carry any App-specific state from this `/authorize` call to the callback on the `redirect_uri` (such as a user’s id). It must be verified in your callback to [prevent CSRF attacks](https://datatracker.ietf.org/doc/html/rfc6749#section-10.12).

The `code_challenge` value is a URL-safe base64-encoded string of the SHA256 hash of a generated code verifier. This code verifier must be a highly randomised string generated on the app side before calling `/authorize`, the code verifier is sent when exchanging the returned authorization code for an access token. For more information see [Proof Key for Code Exchange by OAuth Public Clients](https://datatracker.ietf.org/doc/html/rfc7636). Note that Snyk Apps only supports `S256` for the code challenge method.

<figure><img src="../../../../.gitbook/assets/image (118).png" alt="Authorize access to Organization"><figcaption><p>Authorize access to Organization</p></figcaption></figure>

After the connection is complete, the user is redirected to the specified redirect URI with query string parameters `code` and `state` added on, which are necessary for the next step.
