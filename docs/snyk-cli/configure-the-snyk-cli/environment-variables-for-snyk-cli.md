# Environment variables for Snyk CLI

You can set the following environment variables to change CLI settings.

`SNYK_CFG_<KEY>`

Allows you to override any key that is also available as a `snyk config` option.

For example, `SNYK_CFG_ORG=myorg` overrides the default org option in `config` with "myorg".

`SNYK_REGISTRY_USERNAME`

For the `snyk container` commands, specify a username to use when connecting to a container registry. Note that using the `--username` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.

`SNYK_REGISTRY_PASSWORD`

For the `snyk container` commands, specify a password to use when connecting to a container registry. Note that using the `--password` flag overrides this value. This is ignored in favor of local Docker binary credentials when Docker is present.
