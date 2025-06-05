# BitBucket Pipelines migration

## For users of `snyk/snyk-scan` < v1.0.0 <a href="#users-using-snyk-snyk-scan-less-than-v1.0.0" id="users-using-snyk-snyk-scan-less-than-v1.0.0"></a>

`snyk/snyk-scan` \<v1.0.0 uses Snyk CLI Images. As all Snyk CLI Images will be removed, `snyk/snyk-scan` < v1.0.0 will stop working also.

See the [upgrade guide in Snyk Docs](../bitbucket-pipelines-integration-using-a-snyk-pipe/migrating-to-bitbucket-pipelines-v1.0.0.md) to upgrade to `snyk/snyk-scan` >= v1.0.0

## For users of `snyk/snyk-scan` >= v1.0.0 <a href="#users-using-snyk-snyk-scan-greater-than-v1.0.0" id="users-using-snyk-snyk-scan-greater-than-v1.0.0"></a>

### Create your own custom image <a href="#create-your-own-custom-image" id="create-your-own-custom-image"></a>

Users can create their own custom images to use. This option is available for `snyk/snyk-scan` >= v1.0.0 only. For details, see [User-defined custom images for CLI.](../user-defined-custom-images-for-cli.md)

Creating a custom image should guarantee compatibility with your system. However, there are alternative images to which you can upgrade if creating a custom image is not possible.

### Upgrade to a supported Snyk Image <a href="#upgrade-to-a-supported-snyk-image" id="upgrade-to-a-supported-snyk-image"></a>

After you have validated that you are using a Snyk Image that will be removed, as outlined [for users of `snyk/snyk-scan` < v1.0.0](bitbucket-pipelines-migration.md#users-using-snyk-snyk-scan-less-than-v1.0.0), refer to the [Snyk images migration](snyk-images-migration.md) guidelines to view upgrade paths for your configuration.

{% hint style="info" %}
Remember to use pinned versions where available for better stability. for example, `snyk/snyk:dotnet-8.0` is preferable to `snyk/snyk:dotnet`
{% endhint %}

An example follows of upgrading to a supported Snyk Image.

In the example `bitbucket-pipeline.yml` configuration that follows, a Snyk image is configured that will be removed on 12 Aug 2024:

```yaml
#  Example bitbucket-pipelines.yml using `snyk/snyk:node-16` Snyk Image
#  Template NodeJS build

#  This template allows you to validate your NodeJS code.
#  The workflow allows running tests and code linting on the default branch.

image: atlassian/default-image:latest

pipelines:
  default:
    - parallel:
        - step:
            name: Build
            caches:
              - node
            script:
              - npm install
        - step:
            name: Snyk scan
            script:
              - pipe: snyk/snyk-scan:1.0.1
                variables:
                  SNYK_TOKEN: $SNYK_TOKEN
                  LANGUAGE: "node-16" # <------ Using the `snyk/snyk:node-16` Snyk Image
                  EXTRA_ARGS: "--all-projects" # Optional
                  DEBUG: "true" # Optional
```

Following the [Snyk images migration](snyk-images-migration.md) guidelines, you can upgrade to a supported Snyk Image as shown here:

```yaml
#  Upgrading to supported Snyk Image `snyk/snyk:node-22`
#  Template NodeJS build

#  This template allows you to validate your NodeJS code.
#  The workflow allows running tests and code linting on the default branch.

image: atlassian/default-image:latest

pipelines:
  default:
    - parallel:
        - step:
            name: Build
            caches:
              - node
            script:
              - npm install
        - step:
            name: Snyk scan
            script:
              - pipe: snyk/snyk-scan:1.0.1
                variables:
                  SNYK_TOKEN: $SNYK_TOKEN
                  LANGUAGE: "node-22" # <------ Upgrade to the `snyk/snyk:node-22` Snyk Image
                  EXTRA_ARGS: "--all-projects" # Optional
                  DEBUG: "true" # Optional
```

## Download and install Snyk CLI directly <a href="#download-and-install-snyk-cli-directly" id="download-and-install-snyk-cli-directly"></a>

If you do not want to use the Bitbucket `snyk/snyk-scan` integration, you have the option to install and use the Snyk CLI directly.

{% hint style="info" %}
If you use this option, you will be unable to use integration features such as Code Insight Results
{% endhint %}

The following example shows using the CLI directly.

In the example `bitbucket-pipeline.yml` configuration that follows, a pipeline is configured that does the following:

* Downloads the CLI
* Validates the CLI with a SHASUM check
* Runs the CLI to test the code

```yaml
image: node:18

pipelines:
  default:
    - step:
        name: Build
        caches:
          - node
        script:
          - npm install
    - step:
        name: Snyk scan
        script:
          # Download Snyk Linux CLI
          - curl https://downloads.snyk.io/cli/latest/snyk-linux -o snyk-linux
          # Download Snyk Linux CLI SHASUM
          - curl https://downloads.snyk.io/cli/latest/snyk-linux.sha256 -o snyk.sha256
          # Validate binary using SHASUM
          - sha256sum -c snyk.sha256
          # Configure CLI for executtion
          - chmod +x snyk-linux
          - mv snyk-linux /usr/local/bin/snyk
          # Run Snyk CLI
          - snyk test --all-projects -d
```

&#x20;

\
