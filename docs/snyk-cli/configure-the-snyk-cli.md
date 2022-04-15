# Configure the Snyk CLI

You can use [environment variables](configure-the-snyk-cli.md#environment-variables) to configure the Snyk CLI. You can also set variables to configure the Snyk CLI to [connect to the Snyk API](configure-the-snyk-cli.md#configuration-to-connect-to-the-snyk-api).

## Environment variables

You can set the following environment variables to change CLI settings.

`SNYK_TOKEN`

Allows you to override the token that may be available in your Snyk configuration settings (`~/.config/configstore/snyk.json`). Use `SNYK_TOKEN` in a CI/CD environment. After setting `SNYK_TOKEN` you can [get started](getting-started-with-the-cli/) using the CLI..

For information on how to get your account token see [Authenticate the CLI with your account](authenticate-the-cli-with-your-account/). You can also use a service account to authenticate; for more information see [Service accounts](../features/integrations/managing-integrations/service-accounts.md). For additional information, see [Authentication for third-party tools](../features/user-and-group-management/authentication/authentication-for-third-party-tools.md).

`SNYK_CFG_<KEY>`

Allows you to override any key that is also available as a `snyk config` option.

For example, `SNYK_CFG_ORG=myorg` overrides the default org option in `config` with "myorg".

`SNYK_REGISTRY_USERNAME`

For the container command, specify a username to use when connecting to a container registry. Note that using the `--username` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.

`SNYK_REGISTRY_PASSWORD`

For the container command, specify a password to use when connecting to a container registry. Note that using the `--password` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.

## Configuration to connect to the Snyk API

By default the Snyk CLI connects to `https://snyk.io/api/v1`.

`SNYK_API`

Sets the API host to use for Snyk requests. Useful for on-premise instances or when using a proxy server. If set with the `http` protocol the CLI upgrades the requests to `https`, unless `SNYK_HTTP_PROTOCOL_UPGRADE` is set to `0`.

`SNYK_HTTP_PROTOCOL_UPGRADE=0`

If set to the value of `0`, API (and CLI) requests aimed at `http` URLs are not upgraded to `https`. If the protocol is not set, the default behavior is to upgrade these requests from `http` to `https`. This is useful, for example, for reverse proxies.

`HTTPS_PROXY` and `HTTP_PROXY`

Allows you to specify a proxy to use for `https` and `http` calls. The `https` in the `HTTPS_PROXY` means that _requests using `https` protocol_ use this proxy. The proxy itself doesn't need to use `https`.

`SNYK_DISABLE_ANALYTICS=1`

Disables all Snyk CLI analytics.

`SNYK_OAUTH_TOKEN=<OAuth token>`

Specifies the OAuth token if required for verification.

For more information see [How can I use Snyk behind a proxy?](https://support.snyk.io/hc/en-us/articles/360000925358-How-can-I-use-Snyk-behind-a-proxy-)
