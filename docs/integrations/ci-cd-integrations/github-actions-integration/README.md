# GitHub Actions integration

{% hint style="info" %}
As of December 15, 2022, the GitHub Actions integration pages are being moved from the repository to the Snyk docs site. During this process the explanations will provide the same basic information but vary in presentation. If you need help contact [Snyk support](https://support.snyk.io/hc/en-us).
{% endhint %}

## Overview of GitHub Actions Integration

Snyk offers a [set of GitHub actions](https://github.com/snyk/actions) for using [Snyk](https://snyk.io/) to check for vulnerabilities in your GitHub projects. These actions are are based on the [Snyk CLI ](https://docs.snyk.io/snyk-cli)and you can use [all of its options and capabilities](https://docs.snyk.io/snyk-cli/cli-reference) with the `args` in the [properties](./#snyk-github-action-properties-for-open-source-languages-and-package-managers) of the action.

There is also a [Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup).

For additional information see the [GitHub Actions feature ](https://github.com/features/actions)page and the [GitHub custom actions](https://docs.github.com/en/actions/creating-actions/about-actions) documentation.

You must use a different action depending on the language or process you are using. This page provides detailed information that applies to **all GitHub Actions for Open Source languages and package managers**. For **Open Source examples**, see the pages listed in the next section, [GitHub Actions for Open Source languages and package managers](./#github-actions-for-open-source-languages-and-package-managers).

For detailed information about the **Docker and IaC GitHub Actions and examples** see the pages listed in the subsequent section, [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

For detailed information about the **Setup Action and examples**, see [Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup).

### GitHub Actions for Open Source languages and package managers

* [Snyk CocoaPods Action](snyk-cocoapods-action.md)
* [Snyk dotNET Action](https://github.com/snyk/actions/tree/master/dotnet)
* [Snyk Golang Action](https://github.com/snyk/actions/tree/master/golang)
* [Snyk Gradle Action](https://github.com/snyk/actions/tree/master/gradle)
* [Snyk Gradle-jdk11 Action](https://github.com/snyk/actions/tree/master/gradle-jdk11)
* [Snyk Gradle-jdk12 Action](https://github.com/snyk/actions/tree/master/gradle-jdk12)
* [Snyk Gradle-jdk14 Action](https://github.com/snyk/actions/tree/master/gradle-jdk14)
* [Snyk Gradle-jdk16 Action](https://github.com/snyk/actions/tree/master/gradle-jdk16)
* [Snyk Gradle-jdk17 Action](https://github.com/snyk/actions/tree/master/gradle-jdk17)
* [Snyk Maven Action](https://github.com/snyk/actions/tree/master/maven)
* [Snyk Maven-3-jdk-11 Action](https://github.com/snyk/actions/tree/master/maven-3-jdk-11)
* [Snyk Node Action](https://github.com/snyk/actions/tree/master/node)
* [Snyk PHP Action](https://github.com/snyk/actions/tree/master/php)
* [Snyk Python Action](https://github.com/snyk/actions/tree/master/python)
* [Snyk Python-3.6 Action](https://github.com/snyk/actions/tree/master/python-3.6)
* [Snyk Python-3.7 Action](https://github.com/snyk/actions/tree/master/python-3.7)
* [Snyk Python-3.8 Action](https://github.com/snyk/actions/tree/master/python-3.8)
* [Snyk Ruby Action](https://github.com/snyk/actions/tree/master/ruby)
* [Snyk Scala Action](https://github.com/snyk/actions/tree/master/scala)

### GitHub Actions for Snyk Container and Snyk Infrastructure as Code

* [Snyk Docker Action](https://github.com/snyk/actions/tree/master/docker)
* [Snyk Infrastructure as Code Action](https://github.com/snyk/actions/tree/master/iac)

## Snyk GitHub Action properties for Open Source languages and package managers

The Snyk GitHub Action for Open Source languages and package managers has properties which are passed to the underlying image using `with`.

| Property | Default | Description                                                                                                                                                |
| -------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| args     |         | Override the default arguments to the Snyk image. See [Snyk CLI commands and options summary for all options](https://docs.snyk.io/snyk-cli/cli-reference) |
| command  | test    | Specify which command to run, for instance test or monitor                                                                                                 |
| json     | false   | In addition to the stdout, save the results as snyk.json                                                                                                   |

For the properties associated with the other Snyk GitHub Actions, see the pages listed in the section [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code) and [Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup).

## Examples of using a Snyk GitHub Action

Examples follow of using  a Snyk GitHub Action to test and monitor an Open Source project. For information on using `snyk test` versus `snyk monitor` see [What are the differences among snyk test, monitor, and protect?](https://support.snyk.io/hc/en-us/articles/360000920818-What-is-the-difference-between-snyk-test-protect-and-monitor-)

You can find examples specific to each language, package manager, and process on the pages listed in [GitHub Actions for Open Source languages and package managers](./#github-actions-for-open-source-languages-and-package-managers) and [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

**Note:** GitHub Actions will not pass on secrets set in the repository to forks being used in pull requests, and so the Snyk Actions that require the token will fail to run from a forked repository.

### Snyk test example

An example follows of using a Snyk GItHub Action to test a Node.js project:

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

By using the `args` property of the action you can use [all of the options and capabilities of the Snyk CLI](https://docs.snyk.io/snyk-cli/cli-reference). This example shows use of the option `--severity-threshold=high`.

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/nodemaster
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## GitHub Code Scanning support

Snyk GitHub Actions support integration with GitHub Code Scanning to show vulnerability information in the GitHub Security tab. The following applies to **Snyk GitHub Actions for Open Source languages and package managers**. For information on **specific languages, package managers, and processes** see the pages listed in [GitHub Actions for Open Source languages and package managers](./#github-actions-for-open-source-languages-and-package-managers) and [GitHub Actions for Snyk Container and Snyk Infrastructure as Code](./#github-actions-for-snyk-container-and-snyk-infrastructure-as-code).

Using `--sarif-file-output` [Snyk CLI option](https://docs.snyk.io/snyk-cli/cli-reference) and the [GitHub SARIF upload action](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github), you can upload Snyk scan results to the GitHub Code Scanning as shown in the example that follows.

The Snyk Action fails when vulnerabilities are found. This would prevent the SARIF upload action from running. Thus you must use a [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) option as shown in this example:

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/nodemaster
        continue-on-error: true # To make sure that SARIF upload gets called
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --sarif-file-output=snyk.sarif
      - name: Upload result to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

<figure><img src="../../../.gitbook/assets/GitHub-showing-uploaded-vulnerabilty.png" alt="GitHub showing uploaded vulnerability"><figcaption><p>GitHub showing uploaded vulnerability</p></figcaption></figure>

## Use your own development environment

The Snyk GitHub Actions for each language automatically install all the required development tools for Snyk to determine the correct dependencies and hence vulnerabilities from different language environments. If you have a workflow where you already have the development tools installed, you can instead use the `snyk/actions/setup` Action to install only [Snyk CLI](https://github.com/snyk/snyk). An example follows:

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

Every Snyk account has this token. Once you [create an account](https://docs.snyk.io/getting-started/create-a-snyk-account) you can find the API token in one of two ways:

1. In the UI, go to your Snyk account [settings page](https://app.snyk.io/account) and retrieve the API token, as explained in  [Revoking and regenerating Snyk API tokens](https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens).
2. If you're using the [Snyk CLI](https://docs.snyk.io/snyk-cli/getting-started-with-the-cli) locally you can retrieve the API token by running `snyk config get api`.

