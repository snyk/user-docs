# User-defined custom images for CLI

## Context for user-defined custom images for CLI

Following Snyk announcements regarding [Snyk CLI Images](https://headwayapp.co/snyk-io-updates/deprecation-notice-for-snyk-cli-images-292562) and [Snyk Images](https://updates.snyk.io/deprecation-notice-for-obsolete-snyk-images-292563), Snyk is providing instructions that customers will find useful in building their own custom images. You can visit the [Snyk Images build tool chain](https://github.com/snyk/snyk-images) on GitHub for a list of [currently supported images](https://github.com/snyk/snyk-images?tab=readme-ov-file#current-images).&#x20;

{% hint style="info" %}
Snyk does not build or maintain images that contain end-of-life software by an upstream vendor. The list of unsupported images is available on the [GitHub repo](https://github.com/snyk/snyk-images?tab=readme-ov-file#vendor-unsupported-base-images).
{% endhint %}

Using custom images will allow you to extend environment support to any [environment](../snyk-cli/install-or-update-the-snyk-cli/) supported by the Snyk CLI.

## Requirements for user-defined custom images for CLI

For your custom image to be supported, it must:

* Use an environment [supported](../snyk-cli/install-or-update-the-snyk-cli/) by the Snyk CLI.
* Use a language and framework [supported](../../supported-languages/supported-languages-package-managers-and-frameworks.md) by Snyk.
* Have Snyk CLI installed on the image; see [Install or update the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/) for steps to install the CLI.
* Be publicly accessible; the integration will pull the image.

## Use of user-defined custom images for CLI

Providing a custom image gives you more control over your environment. For example, unless you use custom images, you cannot use an environment with Node LTS.

### Example: How to create a custom image using Dockerfile for Node LTS support

Given the base requirements, you can create a custom image to use Node LTS with the following Dockerfile:

{% code title="Dockerfile" %}
```docker
FROM alpine:3.18

# Install curl
RUN apk add --no-cache curl

# Install Node LTS
RUN apk add --no-cache nodejs

# Install & setup Snyk CLI
RUN curl -o ./snyk-alpine https://downloads.snyk.io/cli/stable/snyk-alpine && \
    curl -o ./snyk-alpine.sha256 https://downloads.snyk.io/cli/stable/snyk-alpine.sha256 && \
    sha256sum -c snyk-alpine.sha256 && \
    mv snyk-alpine /usr/local/bin/snyk && \
    chmod +x /usr/local/bin/snyk
```
{% endcode %}

The base image uses Alpine to keep things lightweight. You have installed Node and the Snyk CLI; this satisfies three-quarters of the requirements.

After the Dockerfile is defined, you can build and tag the image using [docker build](https://docs.docker.com/engine/reference/commandline/build/) and push the image using [docker push](https://docs.docker.com/engine/reference/commandline/push/):

```sh
# bulid image
docker build <PATH-TO-DOCKERFILE> --tag foobar/snyk:node-lts

# push image
docker push foobar/snyk:node-lts
```

### Example: how to use a custom image in a BitBucket pipeline&#x20;

Compatibility in the BitBucket Pipeline integration is limited to environments supported by the Docker container the integration runs in. Following the Snyk announcement [Decoupling Snyk Scan from Snyk CLI Docker Images](https://updates.snyk.io/decoupling-snyk-scan-from-snyk-cli-docker-images-277502), before v1.0.0, the experience was limited to environments supported by the Snyk CLI Docker Images.

With the release of v1.0.0, users can define custom images. If the list of environments provided by the [LANGUAGE](bitbucket-pipelines-integration-using-a-snyk-pipe/snyk-pipe-parameters-and-values-bitbucket-cloud.md#snyk-pipe-variables) variable does not support your particular build environment, you can define your own build environment in the form of a custom Docker image.

Ensure that the [Prerequisites for Bitbucket Pipelines integration](bitbucket-pipelines-integration-using-a-snyk-pipe/prerequisites-for-bitbucket-pipelines-integration.md) are met.

As long as the pushed image is publicly accessible, you can use the `SNYK_BASE_IMAGE` and `LANGUAGE` variables to reference your custom image and tag, respectively, in your Bitbucket pipeline:

{% code title="bitbucket-pipelines.yml" %}
```yaml
script:
  - npm install
  - npm test

  - pipe: snyk/snyk-scan:1.0.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "node-lts"
      SNYK_BASE_IMAGE: "foobar/snyk"

# rest of script
```
{% endcode %}
