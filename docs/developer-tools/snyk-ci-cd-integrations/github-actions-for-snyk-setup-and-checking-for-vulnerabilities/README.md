# GitHub actions for Snyk setup and checking for vulnerabilities

## Overview of GitHub Actions Integration

Snyk offers a [set of GitHub actions](https://github.com/snyk/actions) for using Snyk to check for vulnerabilities in your GitHub projects. These actions are based on the [Snyk CLI](../../snyk-cli/), and you can use [all of its options and capabilities](../../snyk-cli/cli-commands-and-options-summary.md) with the `args` in the [properties](./#snyk-github-action-properties-for-open-source-languages-and-package-managers) of the action.

There is also a [Snyk Setup Action](snyk-setup-action.md).

For more information, see the [GitHub Actions feature](https://github.com/features/actions) page and the [GitHub custom actions](https://docs.github.com/en/actions/creating-actions/about-actions) documentation.

You must use a different action depending on the language or process you are using. This page provides detailed information that applies to all GitHub Actions for open-source languages and package managers. For Snyk Open Source examples, see the pages listed in the next section, [GitHub Actions for Open Source languages and package managers](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

For detailed information about the Docker and IaC GitHub Actions and examples, see the pages listed in the subsequent section, [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

For detailed information about the Setup Action and examples, see [Snyk Setup Action](snyk-setup-action.md).

### GitHub Actions for Open Source languages and package managers

* [Snyk CocoaPods Action](snyk-cocoapods-action.md)
* [Snyk dotNET Action](snyk-dotnet-action.md)
* [Snyk Golang Action](snyk-golang-action.md)
* [Snyk Gradle Action](snyk-gradle-action.md)
* [Snyk Gradle-jdk11 Action](snyk-gradle-jdk11-action.md)
* [Snyk Gradle-jdk12 Action](snyk-gradle-jdk12-action.md)
* [Snyk Gradle-jdk14 Action](snyk-gradle-jdk14-action.md)
* [Snyk Gradle-jdk16 Action](snyk-gradle-jdk16-action.md)
* [Snyk Gradle-jdk17 Action](snyk-gradle-jdk17-action.md)
* [Snyk Maven Action](snyk-maven-action.md)
* [Snyk Maven-3-jdk-11 Action](snyk-maven-3-jdk-11-action.md)
* [Snyk Node Action](snyk-node-action.md)
* [Snyk PHP Action](snyk-php-action.md)
* [Snyk Python Action](snyk-python-action.md)
* [Snyk Python-3.6 Action](snyk-python-3.6-action.md)
* [Snyk Python-3.7 Action](snyk-python-3.7-action.md)
* [Snyk Python-3.8 Action](snyk-python-3.8-action.md)
* [Snyk Ruby Action](snyk-ruby-action.md)
* [Snyk Scala Action](snyk-scala-action.md)

### GitHub Actions for Snyk Container and Snyk Infrastructure as Code

* [Snyk Docker Action](snyk-docker-action.md)
* [Snyk Infrastructure as Code Action](snyk-infrastructure-as-code-action.md)

## Snyk GitHub Action properties for Open Source languages and package managers

The Snyk GitHub Action for open-source languages and package managers has properties that are passed to the underlying image using `with`.

| Property | Default | Description                                                                                                                                                       |
| -------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| args     |         | Override the default arguments to the Snyk image. See the [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md) for all options. |
| command  | test    | Specify which command to run, for instance test or monitor.                                                                                                       |
| json     | false   | In addition to the stdout, save the results as snyk.json.                                                                                                         |

For the properties associated with the other Snyk GitHub Actions, see [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code) and [Snyk Setup Action](snyk-setup-action.md).

## Examples of using a Snyk GitHub Action

Examples follow of using a Snyk GitHub Action to test and monitor a Snyk Open Source Project. For information on using `snyk test` versus `snyk monitor`, see [What are the differences among snyk test, monitor, and protect?](https://support.snyk.io/s/article/What-are-the-differences-among-snyk-test-monitor-and-protect)

You can find examples specific to each language, package manager, and process on the pages listed in [GitHub Actions for Open Source languages and package managers](./#github-actions-for-open-source-languages-and-package-managers) and [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

Note that GitHub Actions will not pass on secrets set in the repository to forks being used in pull requests; thus, the Snyk Actions that require the token will fail to run from a forked repository.

### Snyk test example

An example follows of using a Snyk GitHub Action to test a Node.js Project:

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

### Snyk monitor example

If you want to send data to Snyk and be alerted when new vulnerabilities are discovered, run `snyk monitor` as follows:

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          command: monitor
```

### Example of reporting only on high severity vulnerabilities

By using the `args` property of the action, you can use [all of the options and capabilities of the Snyk CLI.](../../snyk-cli/cli-commands-and-options-summary.md) This example shows use of the option `--severity-threshold=high`.

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## GitHub Code Scanning support

Snyk GitHub Actions support integration with GitHub Code Scanning to show vulnerability information on the GitHub Security tab. The following applies to Snyk GitHub Actions for Open Source languages and package manager&#x73;**.** For information on specific languages, package managers, and processes see the pages listed in [GitHub Actions for Open Source languages and package managers](./#github-actions-for-open-source-languages-and-package-managers) and GitHub Actions for Snyk Container and Snyk Infrastructure as Code.

Using `--sarif-file-output` [Snyk CLI ](../../snyk-cli/cli-commands-and-options-summary.md)option and the [GitHub SARIF upload action](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github), you can upload Snyk scan results to GitHub Code Scanning as shown in the example that follows.

The Snyk Action fails when vulnerabilities are found. This would prevent the SARIF upload action from running. Thus, you must use a [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) option, as shown in the example that follows.

{% hint style="info" %}
To use this option for private repositories, you must have GitHub Advanced Security.

If you see the error `Advanced Security must be enabled for this repository to use code scanning`, check that GitHub Advanced Security is enabled. For more information, see [Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository).
{% endhint %}

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        continue-on-error: true # To make sure that SARIF upload gets called
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --sarif-file-output=snyk.sarif
      - name: Upload result to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v3
        with:
          sarif_file: snyk.sarif
```

After you upload to GitHub Code Scanning support, you will see vulnerabilities on the GitHub Security tab as shown in the following screen image.

<figure><img src="../../../.gitbook/assets/GitHub-showing-uploaded-vulnerabilty.png" alt="GitHub Security tab showing uploaded vulnerability"><figcaption><p>GitHub Security tab showing uploaded vulnerability</p></figcaption></figure>

## Use your own development environment

The Snyk GitHub Actions for each language automatically install all the required development tools for Snyk to determine the correct dependencies and hence vulnerabilities from different language environments. If you have a workflow where you already have the development tools installed, you can use the `snyk/actions/setup` action instead to install only the [Snyk CLI](https://github.com/snyk/snyk). An example follows:

```yaml
name: Snyk example
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - uses: snyk/actions/setup@master
      - uses: actions/setup-go@v1
        with:
          go-version: '1.19'
      - name: Snyk test
        run: snyk test
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

The example here uses `actions/setup-go`. You must select the right action to install the relevant development requirements for your project. If you are already using the same pipeline to build and test your application, you are likely already installing the relevant development requirements.

## Getting your Snyk token

The Snyk GitHub Actions examples on this page refer to a Snyk API token:

```yaml
env:
  SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

Every Snyk account has this token. After you have created an account with Snyk, you can find the API token in one of two ways:

1. In the UI, go to your Snyk account [settings page](https://app.snyk.io/account) and retrieve the API token, as explained on the page [Revoke and regenerate a Snyk API token](../../../snyk-api/authentication-for-api/revoke-and-regenerate-a-snyk-api-token.md).
2. If you are using the [Snyk CLI](../../snyk-cli/getting-started-with-the-snyk-cli.md) locally, you can retrieve the API token by running `snyk config get api`.
