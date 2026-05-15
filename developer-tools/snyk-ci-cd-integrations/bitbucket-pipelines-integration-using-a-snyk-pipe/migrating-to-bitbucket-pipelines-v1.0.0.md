# Migrating to Bitbucket Pipelines v1.0.0

When you are upgrading from < 1.0.0 to 1.0.0+, make the following changes in your configuration:

* Use tags supported by [Snyk images](https://hub.docker.com/r/snyk/snyk/tags) for the `LANGUAGE` variable instead of tags supported by [Snyk CLI Docker images](https://hub.docker.com/r/snyk/snyk-cli/tags).
* Alternatively, you can provide your own [custom image](../user-defined-custom-images-for-cli.md) using the new `SNYK_BASE_IMAGE` variable

<figure><img src="../../../.gitbook/assets/Untitled.jpg" alt="Users can now define custom images in v1.0.0"><figcaption><p>Users can now define custom images in v1.0.0</p></figcaption></figure>

## Basic app dependencies scan example <a href="#markdown-header-basic-app-dependencies-scan-example" id="markdown-header-basic-app-dependencies-scan-example"></a>

This example uses Snyk to scan a Node.js application and break the build if any vulnerabilities are found.

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

## Equivalent Snyk images

The table that follows lists the Snyk CLI Docker images used in Bitbucket Pipelines < 1.0.0 and the equivalent Snyk images that can be used in Bitbucket Pipelines > 1.0.0.

{% hint style="info" %}
Node.js 14 is installed in all Snyk CLI Docker images for installing the CLI. In comparison, NodeJS is installed only in Snyk images containing the `node` tag.
{% endhint %}

Note that the images will not be exactly like-for-like. Properties, including the base image, the installed Snyk CLI version, and others, will be different.

However, for the purpose of Bitbucket Pipelines, they are functionally equivalent.

Where the supported language or framework has reached end of life, Snyk highly recommends that you use your own [user-defined custom images](../user-defined-custom-images-for-cli.md) with newer versions of the language or framework that are still supported by the vendor.

| **Language / Tooling** | **Snyk CLI Image**                                                       | **Snyk Image**                                                                                                                                                                                                                                                                                                                                                                         |
| ---------------------- | ------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| PHP / Composer         | <ul><li>1*.*-composer</li><li>composer</li></ul>                         | <ul><li>snyk/snyk:composer</li></ul>                                                                                                                                                                                                                                                                                                                                                   |
| Docker                 | <ul><li>1.*.*-docker</li><li>docker</li></ul>                            | <ul><li>snyk/snyk:docker-latest</li><li>snyk/snyk:docker</li></ul>                                                                                                                                                                                                                                                                                                                     |
| Java / Gradle          | <ul><li>1.*.*-gradle-*</li><li>gradle-*</li><li>gradle-*_java*</li></ul> | <ul><li>snyk/snyk:gradle</li><li>snyk/snyk:gradle-jdk11</li><li>snyk/snyk:gradle-jdk12</li><li>snyk/snyk:gradle-jdk13</li><li>snyk/snyk:gradle-jdk14</li><li>snyk/snyk:gradle-jdk16</li><li>snyk/snyk:gradle-jdk17</li><li>snyk/snyk:gradle-jdk18</li><li>snyk/snyk:gradle-jdk19</li><li>snyk/snyk:gradle-jdk20</li><li>snyk/snyk:gradle-jdk21</li><li>snyk/snyk:gradle-jdk8</li></ul> |
| Java / Maven           | <ul><li>1*.*-maven-*</li><li>maven-*</li><li>maven-*_java*</li></ul>     | <ul><li>snyk/snyk:maven</li><li>snyk/snyk:maven-3-jdk-11</li><li>snyk/snyk:maven-3-jdk-17</li><li>snyk/snyk:maven-3-jdk-20</li><li>snyk/snyk:maven-3-jdk-21</li><li>snyk/snyk:maven-3-jdk-22</li><li>snyk/snyk:maven-3-jdk-8</li></ul>                                                                                                                                                 |
| JS (Node) / NPM        | <ul><li>1*.*-npm</li><li>npm</li></ul>                                   | <ul><li>snyk/snyk:node</li><li>snyk/snyk:node-18</li><li>snyk/snyk:node-20</li><li>snyk/snyk:node-22</li></ul>                                                                                                                                                                                                                                                                         |
| .Net / Nuget           | <ul><li>1*.*-nuget</li><li>nuget</li></ul>                               | <ul><li>snyk/snyk:dotnet</li><li>snyk/snyk:dotnet-8.0</li></ul>                                                                                                                                                                                                                                                                                                                        |
| Python                 | <ul><li>1*.*-python-*</li><li>python-*</li></ul>                         | <ul><li>snyk/snyk:python</li><li>snyk/snyk:python-3.8</li><li>snyk/snyk:python-3.9</li><li>snyk/snyk:python-3.10</li><li>snyk/snyk:python-3.11</li><li>snyk/snyk:python-alpine</li></ul>                                                                                                                                                                                               |
| Ruby / Ruby Gems       | <ul><li>1*.*-rubygems</li><li>rubygems</li></ul>                         | <ul><li>snyk/snyk:ruby</li><li>snyk/snyk:ruby-3.3</li><li>snyk/snyk:ruby-alpine</li></ul>                                                                                                                                                                                                                                                                                              |
| Scala / SBT            | <ul><li>1*.*-sbt-*</li><li>sbt-*</li></ul>                               | <ul><li>snyk/snyk:sbt1.10.0-scala3.4.2</li></ul>                                                                                                                                                                                                                                                                                                                                       |
