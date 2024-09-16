# Service accounts using OAuth 2.0

{% hint style="info" %}
**Feature availability**\
Service accounts using OAuth 2.0 are available to Enterprise customers through the Snyk REST API. See [Manage service accounts using the Snyk API](manage-service-accounts-using-the-snyk-api.md).
{% endhint %}

You can create service accounts that authenticate with the [Auth 2.0 `client_credentials` grant flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) rather than with the Snyk API key. OAuth 2.0 access tokens can be used the same way a Snyk API key would be used, but they have the added security of a short time-to-live (TTL) and can be automatically refreshed.

## OAuth 2.0 with client secret

You can create an `oauth_client_secret` service account at either the Group or Organization level by calling the API endpoint [Create a service account for an organization](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts) or [Create a service account for a group](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts).

The response returns the `client_secret` and the `client_id`.  You cannot view the `client_secret` again after the service account is created. If you have misplaced it, you can [rotate](manage-service-accounts-using-the-snyk-api.md#manage-a-service-account-client-secret-for-your-group) your `client_secret` and receive a new one.

{% hint style="danger" %}
Never share the `client_secret` publicly, as this is used to authenticate your service account. Keep it secure and private.
{% endhint %}

### Retrieve an OAuth 2.0 access token

After the service account is created, you can retrieve an `access_token` through the OAuth 2.0 endpoint [Request an access token](../../snyk-api/oauth2-api.md#token) using the `client_secret`. The body format and `Content-Type` header must be form-urlencoded.

You can use an`access_token` the same way as you would use a Snyk API key, but with the `Authorization: bearer $access_token` header or the `SNYK_OAUTH_TOKEN` environment variable with the Snyk CLI.

The `access_token` has a short time to live and must be refreshed once it expires. There are many OAuth 2.0 libraries available that greatly simplify this process.

### Authenticate with the Snyk CLI

You can also use the `client_secret` and the `client_id` to authenticate with the Snyk CLI. For details, see the [`snyk auth`](../../snyk-cli/commands/auth.md) command help.

## OAuth 2.0 with Private Key JWT &#x20;

Service accounts can also use the `private_key_jwt` form of `client_assertion` for authentication, as defined in the [OIDC Core 1.0 spec](https://openid.net/specs/openid-connect-core-1\_0.html#ClientAuthentication)) rather than the more traditional `client_secret`.

This is an advanced feature for customers who have the infrastructure to host a publicly accessible [JWKS](https://datatracker.ietf.org/doc/html/rfc7517) endpoint and manage public and private signing keys. It creates an added layer of security for OAuth service accounts by requiring access token requests to be sent in the form of signed JWTs rather than with a simple client id and client secret pair in a request body.

### Prerequisites to set up the Private Key JWT service account

Before creating the service account, you must do the following:

* Generate public and private signing keys.
* Host a publicly accessible JWKS endpoint to access the public keys; this must be `https`.

Customer-managed signing keys can be rotated at will.

### Create the Private Key JWT service account

You can create an `oauth_private_key_jwt` service account at either the Group or Organization level by using the Snyk REST API. To create the service account, you must provide the URL from your publicly accessible JWKS endpoint and the role you wish to assign to the service account. For more information, see [Manage service accounts using the Snyk API](manage-service-accounts-using-the-snyk-api.md).

The response includes the `client_id`, which is needed for the next step.

### Prepare the JWT signed with a private key

Snyk recommends you prepare a tool or script to build a `private_key_jwt` with the proper claims and sign it with the private signing key you generated. as a prerequisite

The JWT should include the [claims](https://datatracker.ietf.org/doc/html/rfc7519#section-4) outlined in the [Snyk OAuth 2.0 token endpoint](../../snyk-api/oauth2-api.md#token) for the `client_assertion` property. Note that the `aud` claim may vary based on the Snyk instance, for example, `api.snyk.io`, or `api.eu.snyk.io`. For more information, see [Regional hosting and data residency](../../working-with-snyk/regional-hosting-and-data-residency.md).

### Retrieve a Private Key JWT access token

After the service account is created and signed JWT is prepared, you can retrieve an `access_token` using the [Snyk OAuth 2.0 token endpoint](../../snyk-api/oauth2-api.md#token). This access token can be used as a Snyk API key would be used. The request body should include the following:

* `grant_type: client_credentials`
* `client_assertion_type: private_key_jwt`
* `client_assertion:` `<signed JWT>`

The `access_token` has a short time-to-live and must be refreshed once it expires. There are many OAuth 2.0 libraries available that will greatly simplify this process.
