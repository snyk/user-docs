# Configure Snyk CLI to connect to Snyk API

By default the Snyk CLI connects to `https://snyk.io/api/v1`. You can use the following variables to configure your connection.

`SNYK_API`

Sets the API host to use for Snyk requests. Useful for on-premise instances or when using a proxy server. If set with the `http` protocol the CLI upgrades the requests to `https`, unless `SNYK_HTTP_PROTOCOL_UPGRADE` is set to `0`.

`SNYK_HTTP_PROTOCOL_UPGRADE=0`

If set to the value of `0`, API (and CLI) requests aimed at `http` URLs are not upgraded to `https`. If the protocol is not set, the default behavior is to upgrade these requests from `http` to `https`. This is useful, for example, for reverse proxies.

`SNYK_DISABLE_ANALYTICS=1`

Disables all Snyk CLI analytics.

`SNYK_OAUTH_TOKEN=<OAuth token>`

Specifies the OAuth token if required for verification.
