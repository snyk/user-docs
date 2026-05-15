# Advanced use of Snyk Container CLI

## Scan archives

In addition to scanning images from a local Docker daemon or remote registry, Snyk can directly scan or monitor a Docker or Open Container Initiative (OCI) archive when you use the following commands:

* `snyk container test docker-archive:<filename>.tar`
* OR `snyk container test oci-archive:<filename>.tar.`&#x20;

Beginning with CLI version 1.1296.0, you can scan and monitor Kaniko image archives using the following commands:

* `snyk container test kaniko-archive:<filename>.tar`
* `snyk container monitor kaniko-archive:<filename>.tar`.

Beginning with CLI version 1.1297.0, you can scan and monitor image archives without specifying the archive type:

* `snyk container test <filename>.tar`
* `snyk container monitor <filename>.tar`.

This update maintains full support for the current CLI scanning features and ensures backward compatibility.

Examples:

```
snyk container test docker-archive:archive.tar
snyk container test oci-archive:archive.tar
snyk container test kaniko-archive:archive.tar
snyk container test archive.tar
```

{% hint style="info" %}
For`Crane`, Snyk supports only the formats `--format=oci` and`--format=legacy`.
{% endhint %}

## Test multi-platform images

Some repositories represent multi-manifests, pointing to several different images depending on the operating system and the architecture required. To explicitly scan an image for a specific platform, you can use the Snyk CLI `container test` command, for example:

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

## Use an alternate Docker context

The Snyk CLI `container` commands always use the default [Docker context](https://docs.docker.com/engine/manage-resources/contexts/). To force the Snyk CLI to use an alternate context's connection, set the `DOCKER_HOST` environment variable to the desired context URI.&#x20;

## Other commonly used CLI options

Frequently used CLI options include:

* `--json` - useful for integrating with other tools
* `--sarif` - useful for integrating with other tools. The option is available  only with `container test`. See also [OASIS Static Analysis Results Interchange Format (SARIF)](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif).
* `--exclude-base-image-vulns` - only available with `container test`
* `--severity-threshold` - available only with `container test`
* `--exclude-app-vulns`
* `--nested-jars-depth`
* `--fail-on` - available  only with `container test`

For more details and CLI options, see the [Snyk CLI container help ](../../commands/container.md)or display the help by running:

```
snyk container --help
```
