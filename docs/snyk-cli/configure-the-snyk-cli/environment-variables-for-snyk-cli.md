# Environment variables for Snyk CLI

You can set the following environment variables to change CLI settings.

`SNYK_TOKEN`

Allows you to override the token that may be available in your Snyk configuration settings (`~/.config/configstore/snyk.json`). Use `SNYK_TOKEN` in a CI/CD environment. After setting `SNYK_TOKEN` you can [get started](../getting-started-with-the-cli.md) using the CLI..

For information on how to get your account token see [Authenticate the CLI with your account](../authenticate-the-cli-with-your-account.md). You can also use a service account to authenticate; for more information see [Service accounts](../../features/user-and-group-management/structure-account-for-high-application-performance/service-accounts.md). For additional information, see [Authentication for third-party tools](../../features/user-and-group-management/authentication/authentication-for-third-party-tools.md).

`SNYK_CFG_<KEY>`

Allows you to override any key that is also available as a `snyk config` option.

For example, `SNYK_CFG_ORG=myorg` overrides the default org option in `config` with "myorg".

`SNYK_REGISTRY_USERNAME`

For the `snyk container` commands, specify a username to use when connecting to a container registry. Note that using the `--username` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.

`SNYK_REGISTRY_PASSWORD`

For the `snyk container` commands, specify a password to use when connecting to a container registry. Note that using the `--password` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.
