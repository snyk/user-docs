# Configure Snyk CLI to connect to Snyk API

By default, the Snyk CLI connects to `https://api.snyk.io/`. You can use the following variables to configure your connection.

`SNYK_API`

Sets the API host to use for Snyk requests. Useful for [regional hosting](../../working-with-snyk/regional-hosting-and-data-residency.md#cli-and-ci-pipelines-urls), on-premise instances, or when using a proxy server. If set with the `http` protocol the CLI upgrades the requests to `https`, unless `SNYK_HTTP_PROTOCOL_UPGRADE` is set to `0`.

`SNYK_HTTP_PROTOCOL_UPGRADE=0`

If set to the value of `0`, API (and CLI) requests aimed at `http` URLs are not upgraded to `https`. If the protocol is not set, the default behavior is to upgrade these requests from `http` to `https`. This is useful, for example, for reverse proxies.



`SNYK_DISABLE_ANALYTICS=1`

Disables all Snyk CLI analytics.



`SNYK_OAUTH_TOKEN=<OAuth token>`

Specifies the OAuth token if required for verification.



`SNYK_TOKEN`

Allows you to override the token that may be available in your Snyk configuration settings (`~/.config/configstore/snyk.json`). Use `SNYK_TOKEN` in a CI/CD environment. After setting `SNYK_TOKEN` you can [get started](../getting-started-with-the-snyk-cli.md) using the CLI..

For information on how to get your account token see [Authenticate the CLI with your account](../authenticate-to-use-the-cli.md). You can also use a service account to authenticate; for more information see [Service accounts](../../enterprise-setup/service-accounts/). For additional information, see [Authentication for third-party tools](../../enterprise-setup/authentication-for-third-party-tools.md).
