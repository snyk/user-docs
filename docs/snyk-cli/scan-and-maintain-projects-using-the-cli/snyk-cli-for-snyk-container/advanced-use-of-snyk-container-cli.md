# Advanced use of Snyk Container CLI

## Scan archives

In addition to scanning images from a local Docker daemon or remote registry, Snyk can directly scan or monitor a Docker or OCI or Kaniko archive by running `snyk container test docker-archive:<filename>.tar` or `snyk container test oci-archive:<filename>.tar.` (coming soon `snyk container test kaniko-archive:<filename>.tar.)`

For example:

```
snyk container test docker-archive:archive.tar
snyk container test oci-archive:archive.tar
(coming soon) snyk container test kaniko-archive:archive.tar
```

{% hint style="info" %}
For`crane`, Snyk supports only the formats `--format=oci` and`--format=legacy`.
{% endhint %}

## Test multi-platform images

Some repositories represent multi-manifests, pointing to several different images depending on the operating system and the architecture required. To explicitly scan an image for a specific platform, you can use the Snyk CLI `container test` command.

For example:

```
snyk container test --platform=linux/arm64 debian
```

The `--platform` option must contain one of the following:

* linux/amd64
* linux/arm64
* linux/riscv64
* linux/ppc64le
* linux/s390x
* linux/386
* linux/arm/v7
* linux/arm/v

## Authenticate to a remote container registry

When Docker is installed, the Snyk CLI `container` commands use any pre-configured registry authentication. If you are not using Docker, you can pass the credentials on the command line in one of the following ways:

* Use the following environment variables: `SNYK_REGISTRY_USERNAME` and `SNYK_REGISTRY_PASSWORD`
* Pass the username and password:

```
snyk container test <repository>:<tag> --username= --password=
```

{% hint style="info" %}
When both are passed, the options take precedence over the environment variables.
{% endhint %}

## Other commonly used CLI options

Frequently used CLI options include:

* `--json` - useful for integrating with other tools
* `--sarif` - useful for integrating with other tools. The option is only available with `container test`. See also [OASIS Static Analysis Results Interchange Format (SARIF)](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif).
* `--exclude-base-image-vulns` - only available with `container test`
* `--severity-threshold` - only available with `container test`
* `--exclude-app-vulns`
* `--nested-jars-depth`
* `--fail-on` - only available with `container test`

For more details and CLI options, see the [Snyk CLI container](../../commands/container.md) help or display the help by running:

```
snyk container --help
```
