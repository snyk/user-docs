# User-defined custom images

## Context for user-defined custom images

Compatibility in the BitBucket Pipeline integration is limited to environments supported by the Docker container the integration runs in. Before v1.0.0, this was limited to environments supported by the Snyk CLI Docker Images.

With the release of v1.0.0, users can define custom images. If the list of environments provided by the [LANGUAGE](snyk-pipe-parameters-and-values-bitbucket-cloud.md#snyk-pipe-variables) variable does not support your particular build environment, you can define your own build environment in the form of a custom Docker image.

Using custom images allows you to:

* Extend environment support to any [environment](../../../snyk-cli/install-or-update-the-snyk-cli/) supported by the Snyk CLI
* Extend support to newer versions of languages and frameworks not available in Bitbucket Pipelines versions < 1.0.0

<figure><img src="../../../.gitbook/assets/Untitled.jpg" alt="Users can now define custom images in v1.0.0"><figcaption><p>Users can now define custom images in v1.0.0</p></figcaption></figure>

## Requirements for user-defined custom images

In order for your custom image to be supported, it must:

* Meet the [Prerequisites for Bitbucket Pipelines integration](prerequisites-for-bitbucket-pipelines-integration.md)
* Use an environment [supported](../../../snyk-cli/install-or-update-the-snyk-cli/) by the Snyk CLI
* Use a language and framework [supported](../../../scan-using-snyk/supported-languages-and-frameworks.md) by Snyk
* Have Snyk CLI installed on the image - see [Install or update the Snyk CLI](../../../snyk-cli/install-or-update-the-snyk-cli/) for steps to install the CLI
* Be publicly accessible; the integration will pull the image.

## Use of user-defined custom images

Providing a custom image gives you more control over your environments. For example, unless you use custom images, you cannot  able to use an environment with Node LTS.

### How to create a custom image using Dockerfile

Given the base requirements, you can create a custom image to use Node LTS with the following Dockerfile:

{% code title="Dockerfile" %}
```docker
FROM alpine:3.18

# Install curl
RUN apk add --no-cache curl

# Install Node LTS
RUN apk add --no-cache nodejs

# Install & setup Snyk CLI
RUN curl -o ./snyk-alpine https://static.snyk.io/cli/latest/snyk-alpine && \
    curl -o ./snyk-alpine.sha256 https://static.snyk.io/cli/latest/snyk-alpine.sha256 && \
    sha256sum -c snyk-alpine.sha256 && \
    mv snyk-alpine /usr/local/bin/snyk && \
    chmod +x /usr/local/bin/snyk
```
{% endcode %}

The base image uses Alpine to keep things lightweight. You have installed Node and the Snyk CLI; this satisfies three-quarters of the requirements.

### How to build and push the Docker image

Once the Dockerfile is defined, you can build and tag the image via [docker build](https://docs.docker.com/engine/reference/commandline/build/) and push the image via [docker push](https://docs.docker.com/engine/reference/commandline/push/). for example:

```sh
# bulid image
docker build --tag foobar/snyk:node-lts

# push image
docker push foobar/snyk:node-lts
```

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
