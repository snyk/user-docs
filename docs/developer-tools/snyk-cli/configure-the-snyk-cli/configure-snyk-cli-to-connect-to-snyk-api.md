# Configure Snyk CLI to connect to Snyk API

## **Configure the Snyk platform URL**

By default, the Snyk CLI connects to `https://api.snyk.io/`. You can use the following variables to configure your connection.

`SNYK_API`

Specifying this variable sets the API host that will be used for Snyk requests. This is useful for [regional hosting](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#cli-and-ci-pipeline-urls), on-premise instances, or when you are using a proxy server. If this variable is set with the `http` protocol, the CLI upgrades the requests to `https` unless `SNYK_HTTP_PROTOCOL_UPGRADE` is set to `0`.

`SNYK_HTTP_PROTOCOL_UPGRADE=0`

If you set this variable to the value of `0`, API and CLI requests aimed at `http` URLs are not upgraded to `https`. If the protocol is not set, the default behavior is to upgrade these requests from `http` to `https`. Setting this variable is useful, for example, for reverse proxies.

## **Configure CLI Analytics**

`SNYK_DISABLE_ANALYTICS=1`

Specifying this variable disables all Snyk CLI analytics.

## **Configure the authentication method used to log in to the Snyk platform**

`SNYK_TOKEN`

Specifying this variable allows you to override the token that may be available in your Snyk configuration settings (`~/.config/configstore/snyk.json`). Use `SNYK_TOKEN` in a CI/CD environment. After setting `SNYK_TOKEN` you can [get started](../getting-started-with-the-snyk-cli.md) using the CLI.

For information on how to get your account token see [Authenticate the CLI with your account](../authenticate-to-use-the-cli.md). You can also use a service account to authenticate; for more information see [Service accounts](../../../implementation-and-setup/enterprise-setup/service-accounts/). For additional information, see [Authentication for third-party tools](../../../implementation-and-setup/enterprise-setup/authentication-for-third-party-tools.md).

`SNYK_OAUTH_TOKEN=<OAuth token>`

Using this variable specifies the OAuth token if required for verification.
