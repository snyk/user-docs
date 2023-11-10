---
description: >-
  When upgrading from Bitbucket Pipelines < 1.0.0 to  v1.0.0+, there are some
  changes to note
---

# Migrating to Bitbucket Pipelines v1.0.0

## Changes required

When upgrading from < 1.0.0 to 1.0.0+, the following changes should be made to your configuration:

* Please use tags supported by [Snyk Images](https://hub.docker.com/r/snyk/snyk/tags) for the `LANGUAGE` variable instead of tags supported by [Snyk CLI Docker images](https://hub.docker.com/r/snyk/snyk-cli/tags)
* Alternately, users can provide their own [custom image](user-defined-custom-images.md) using the new `SNYK_BASE_IMAGE` variable

### Example

#### Basic app dependencies scan example <a href="#markdown-header-basic-app-dependencies-scan-example" id="markdown-header-basic-app-dependencies-scan-example"></a>

Uses Snyk to scan a Node.js application and break the build if any vulnerabilities found.

{% code title="snyk/snyk-scan:1.0.0" %}
```yaml
script:
  - npm install

  - npm test

  - pipe: snyk/snyk-scan:1.0.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "node" # language tag is "node"

  - npm publish
```
{% endcode %}

{% code title="snyk/snyk-scan:0.7.0" %}
```yaml
script:
  - npm install

  - npm test

  - pipe: snyk/snyk-scan:0.7.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "npm" # language tag is "npm"

  - npm publish
```
{% endcode %}

## Equivalent Snyk Images

The table below lists the Snyk CLI Docker images used in Bitbucket Pipelines < 1.0.0 and the equivalent Snyk Images that can be used in Bitbucket Pipelines > 1.0.0.

{% hint style="info" %}
NodeJS 14 is installed in all Snyk CLI Docker images for the purpose of installing the CLI. In comparison, NodeJS is only installed Snyk Images containing the `node` tag
{% endhint %}

{% hint style="warning" %}
Note that the images will not be exactly like-for-like. Properties like the base image, the installed Snyk CLI version, etc will be different.

But for the purpose of Bitbucket Pipelines, they are functionally equivalent.
{% endhint %}

{% hint style="danger" %}
Where the supported language/framework has reached EOL, it is **highly recommended** to use your own [user defined custom images](user-defined-custom-images.md) with newer versions of the language/framework that is still supported by the vendor.
{% endhint %}

| **Snyk CLI Docker images tag**                                                                                                                                                                                      | **Snyk Images tag**                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [snyk/snyk-cli:composer](https://hub.docker.com/layers/snyk/snyk-cli/composer/images/sha256-5dabf21ff787a2533d8d53f74f80e690ab69112325181fa8c5b87570b381e8e6?context=explore)                                       | [snyk/snyk:node-14](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=node-14)                                                                                 |
| [snyk/snyk-cli:docker](https://hub.docker.com/layers/snyk/snyk-cli/docker/images/sha256-affb32b0be2f9d4dbf62abd2f63a2060d81f724fa0b14a755945385c08993ae4?context=explore)                                           | [snyk/snyk:docker](https://hub.docker.com/layers/snyk/snyk/docker/images/sha256-5031561bd647169eaa169d4ad8e7a9f61ede7ca4ed9d6a448894d499309da4de?context=explore) |
| [snyk/snyk-cli:gradle-2.8](https://hub.docker.com/layers/snyk/snyk-cli/gradle-2.8/images/sha256-791d44e3cd23813139eda15862a4013b44d8d2f6e82a43d38db85e61ca8f9811?context=explore)                                   | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:gradle-4.4](https://hub.docker.com/layers/snyk/snyk-cli/gradle-4.4/images/sha256-d230db03d1c9e959485323e00edb255957c4e115316cfe194e55aa69cdc2743a?context=explore)                                   | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:gradle-5.4](https://hub.docker.com/layers/snyk/snyk-cli/gradle-5.4/images/sha256-7d6741a7724093c77a7351bd3d84f81c69760c149d050c8e7714156fc4307aa8?context=explore)                                   | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:gradle-5.4\_java11](https://hub.docker.com/layers/snyk/snyk-cli/gradle-5.4\_java11/images/sha256-e46ec043d496232331eecf352d6e4e8af41958c6ec195696eeb8ece69f8bf9f7?context=explore)                   | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:npm](https://hub.docker.com/layers/snyk/snyk-cli/npm/images/sha256-fcb80d3dc0f6d837327b19d2ccfe35461a22779f897ee929dd99dc008fc1ff2a?context=explore)                                                 | [snyk/snyk:node-14](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=node-14)                                                                                 |
| [snyk/snyk-cli:nuget](https://hub.docker.com/layers/snyk/snyk-cli/nuget/images/sha256-c9163ae6deb39b9951bfa2e7cf99b539398fbc0ea9134d0059e58bbd00a95c55?context=explore)                                             | [snyk/snyk:node-14](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=node-14)                                                                                 |
| [snyk/snyk-cli:python-2](https://hub.docker.com/layers/snyk/snyk-cli/python-2/images/sha256-f7b2e87412242171d2f58133d8efd3e99c30e603ce940d4420058fea6411c2d0?context=explore)                                       | [snyk/snyk:python-2.7](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=python-2.7)                                                                           |
| [snyk/snyk-cli:python-3.6](https://hub.docker.com/layers/snyk/snyk-cli/python-3.6/images/sha256-085d514e77fd535eb490777b6a69b0495297e7db5eaf80a9a35ba870e69f2f2c?context=explore)                                   | [snyk/snyk:pyhon-3.6](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=python-3.6)                                                                            |
| [snyk/snyk-cli:rubygems](https://hub.docker.com/layers/snyk/snyk-cli/rubygems/images/sha256-50d3ccb5a2547db4f0c29332330e19b8118dca6f4fd4c7afa8cad82854ec3e46?context=explore)                                       | [snyk/snyk:node-14](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=node-14)                                                                                 |
| [snyk/snyk-cli:1.1228.0-composer](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-composer/images/sha256-b2c8555c0503108d64c1784d441c4336ec96f80a7a305b2889d64679fba6353f?context=explore)                     | [snyk/snyk:node-14](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=node-14)                                                                                 |
| [snyk/snyk-cli:1.1228.0-docker](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-docker/images/sha256-07cc3c26dc78cf183f09c98f681551d81f3db224fa55fc464694a0d37cb6da5e?context=explore)                         | [snyk/snyk:docker](https://hub.docker.com/layers/snyk/snyk/docker/images/sha256-5031561bd647169eaa169d4ad8e7a9f61ede7ca4ed9d6a448894d499309da4de?context=explore) |
| [snyk/snyk-cli:1.1228.0-gradle-2.8](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-gradle-2.8/images/sha256-58985a5c9599008f7c30defd497298fb624f990359aaba5f042b7539b42a62f6?context=explore)                 | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:1.1228.0-gradle-4.4](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-gradle-4.4/images/sha256-f09d66d19f8cc38fa1523b833d7f4262f3500e2338dfaa80292887877d559cf8?context=explore)                 | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:1.1228.0-gradle-5.4](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-gradle-5.4/images/sha256-2bc490a3b49398117282f39d0619213c9e7ba4b20dc8ed695bb121201dab7b38?context=explore)                 | _no direct equivalent\*_                                                                                                                                          |
| [snyk/snyk-cli:1.1228.0-gradle-5.4\_java11](https://hub.docker.com/layers/snyk/snyk-cli/1.1228.0-gradle-5.4\_java11/images/sha256-9a05884de9da4de8b7098039bb28d634e3ca6ae0f3e47d4c5df43ff697e48c20?context=explore) | _no direct equivalent\*_                                                                                                                                          |

&#x20;_\* There are a_ [_selection_](https://hub.docker.com/r/snyk/snyk/tags?page=1\&name=gradle) _of gradle Snyk Images which may suit your needs_
