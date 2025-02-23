# Environment variables for Snyk CLI

`SNYK_TOKEN`

Allows you to override the token that may be available in your Snyk configuration settings (`~/.config/configstore/snyk.json`). Use `SNYK_TOKEN` in a CI/CD environment. After setting `SNYK_TOKEN` you can [get started](../getting-started-with-the-snyk-cli.md) using the CLI..

For information on how to get your account token see [Authenticate the CLI with your account](../authenticate-to-use-the-cli.md). You can also use a service account to authenticate; for more information see [Service accounts](../../enterprise-setup/service-accounts/). For additional information, see [Authentication for third-party tools](../../enterprise-setup/authentication-for-third-party-tools.md).

This page identifies environment variables that you can use to configure specific settings for the CLI.

## **Configure the CLI Cache folder**

`SNYK_CACHE_PATH`

By specifying the environment variable `SNYK_CACHE_PATH`, you can define the path to the folder where the CLI will cache files. You must meet the [access requirements](../security-concept-of-operations-for-snyk/access-requirements.md) for the folder specified.

## **Configure a timeout for CLI operations**

`SNYK_TIMEOUT_SECS`&#x20;

When you specify an environment variable such as `SNYK_TIMEOUT_SECS=60`, the CLI will cancel all operations after 60 seconds and exit with exit code 69.

## **Configure authentication for container registries**

`SNYK_REGISTRY_USERNAME`

For the `snyk container` commands, specify a `USERNAME` to use when connecting to a container registry. Note that using the `--username` option overrides this value. This value is ignored in favor of local Docker binary credentials when Docker is present.

`SNYK_REGISTRY_PASSWORD`

For the `snyk container` commands, specify a `PASSWORD` to use when connecting to a container registry. Note that using the `--password` option overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.

## **Configure CLI settings**

You can set the following environment variable to change CLI settings.

`SNYK_CFG_<KEY>`

By specifying this variable, you can override any key that is also available as a `snyk config` option.

For example, `SNYK_CFG_ORG=myorg` overrides the default `ORG` in `config` with `myorg`.
