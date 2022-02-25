# Advanced Snyk Container CLI usage

## Testing archives

In addition to testing images from a local Docker daemon or remote registry, Snyk can directly test or monitor a Docker or OCI archive.

```
snyk container test docker-archive:archive.tar
snyk container test oci-archive:archive.tar
```

## Testing multi-platform images

Some repositories represent multi-manifests, pointing to several different images depending on the operating system and architecture required. The Snyk CLI `container` command can be used to explicitly test an image for a specific platform:

```
snyk container test --platform=linux/arm64 debian
```

The `--platform` flag should contain one of:

* linux/amd64
* linux/arm64
* linux/riscv64
* linux/ppc64le
* linux/s390x
* linux/386
* linux/arm/v7
* linux/arm/v

## Authenticating to a remote container registry

When Docker is installed, the Snyk CLI `container` command uses any pre-configured registry authentication. If you are not using Docker you can pass the credentials on the command line:

* Either use the following environment variables: `SNYK_REGISTRY_USERNAME` and `SNYK_REGISTRY_PASSWORD`
* Or  pass the username and password as follows:

```
snyk container test <repository>:<tag> --username= --password=
```

Note that the options take precedence over the environment variables in the case both are passed.

## Common additional options

Some useful CLI options include the following:

| Option                       | Description                                                                                                                                                                                                  |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--json`                     | Print results in JSON format, useful for integrating with other tools                                                                                                                                        |
| `--sarif`                    | Return results in [SARIF](https://www.oasis-open.org/committees/tc\_home.php?wg\_abbrev=sarif) format, useful for integrating with other tools. Note this requires the test to be run with `--file` as well. |
| `--exclude-base-image-vulns` | Do not  show vulnerabilities introduced only by the base image. Available when using `snyk container test` only.                                                                                             |
| `--severity-threshold`       | Report only vulnerabilities at the specified level or higher.                                                                                                                                                |
| `--app-vulns`                | Snyk allows detection of vulnerabilities in your application dependencies from container images, as well as from the operating system, all in one single scan.                                               |
| `--nested-jars-depth`        | When using `--app-vulns` use the `--nested-jars-depth=n` option to set how many levels of nested jars Snyk is to unpack.                                                                                     |

For more options, see the Snyk CLI `container` [help](../../../features/snyk-cli/commands/container.md) or display the help by running:

```
snyk container --help
```

## More information

* [Test images with the Snyk Container CLI](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security)
* [Understand Snyk Container CLI results](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/understanding-snyk-container-cli-results)
