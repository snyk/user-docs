# Environment variables for Snyk CLI

This page identifies environment variables that you can use to configure specific settings for the CLI.

## **Configure the CLI Cache folder**

`SNYK_CACHE_PATH`

By specifying the environment variable `SNYK_CACHE_PATH`, you can define the path to the folder where the CLI will cache files. You must meet the [access requirements](../security-concept-of-operations-for-snyk/access-requirements.md) for the folder specified.

## **Configure a timeout for CLI operations**

`SNYK_TIMEOUT_SECS`

When you specify an environment variable such as `SNYK_TIMEOUT_SECS=60`, the CLI will cancel all operations after 60 seconds and exit with exit code 69.

The timeout can range from 1 to 2.147.483.648 seconds, approximately 596.523 hours.&#x20;

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
