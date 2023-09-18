# Advanced use of Snyk Container CLI

## Scanning archives

In addition to testing images from a local Docker daemon or remote registry, Snyk can directly scan or monitor a Docker or OCI archive by running, `snyk container test docker-archive:<filename>.tar` or `snyk container test oci-archive:<filename>.tar.`&#x20;

Examples:

```
snyk container test docker-archive:archive.tar
snyk container test oci-archive:archive.tar
```

## Testing multi-platform images

Some repositories represent multi-manifests, pointing to several different images depending on the operating system and architecture required. The Snyk CLI `container test` command can be used to explicitly test an image for a specific platform.

Example:

```
snyk container test --platform=linux/arm64 debian
```

The `--platform` option should contain one of the following:

* linux/amd64
* linux/arm64
* linux/riscv64
* linux/ppc64le
* linux/s390x
* linux/386
* linux/arm/v7
* linux/arm/v

## Authenticating to a remote container registry

When Docker is installed, the Snyk CLI `container` commands use any pre-configured registry authentication. If you are not using Docker, you can pass the credentials on the command line:

* Either use the following environment variables: `SNYK_REGISTRY_USERNAME` and `SNYK_REGISTRY_PASSWORD`
* Or pass the username and password as follows:

```
snyk container test <repository>:<tag> --username= --password=
```

Note that the options take precedence over the environment variables when both are passed.

## Additional commonly used CLI options

CLI options that are used frequently include the following:

* `--json` (Useful for integrating with other tools)
* `--sarif` (See [OASIS Static Analysis Results Interchange Format (SARIF)](https://www.oasis-open.org/committees/tc\_home.php?wg\_abbrev=sarif); useful for integrating with other tools; available with `container tes`t only)
* `--exclude-base-image-vulns` (Available with `container test` only)
* `--severity-threshold` (Available with `container test` only)
* `--exclude-app-vulns`
* `--nested-jars-depth`
* `--fail-on` (Available with container test only)

For details and more options, see the [Snyk CLI containe](../../../snyk-cli/commands/container.md)r help or display the help by running:

```
snyk container --help
```

## More information about Snyk Container CLI

* [Snyk CLI for container security](./)
* [Understanding Snyk Container CLI results](understanding-snyk-container-cli-results.md)
