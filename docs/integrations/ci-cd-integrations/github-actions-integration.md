# GitHub Actions integration

## Overview of GitHub Actions Integration

Snyk offers a [set of GitHub actions](https://github.com/snyk/actions) for using Snyk to check for vulnerabilities in your GitHub projects. These actions are are based on the [Snyk CLI](../../snyk-cli/cli-reference.md).

For more information about GitHub actions see the [GitHub site ](https://github.com/features/actions)and the [GitHub custom actions](https://docs.github.com/en/actions/creating-actions/about-actions) documentation.

You must use a different action depending on the language or build tool you are using. Snyk currently offers the following actions actions:

* [Snyk CocoaPods Action](https://github.com/snyk/actions/tree/master/cocoapods)
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
* [Snyk Docker Action](https://github.com/snyk/actions/tree/master/docker)
* [Snyk Infrastructure as Code Action](https://github.com/snyk/actions/tree/master/iac)
* Snyk Setup Action

## Example of using a Snyk GitHub Action

An example follows of using an Action to test a Node.js project:

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

For information on using `snyk test` versus `snyk monitor` refer to [What are the differences among snyk test, monitor, and protect?](https://support.snyk.io/hc/en-us/articles/360000920818-What-is-the-difference-between-snyk-test-protect-and-monitor-)

For information on using the Action for each language, refer to the repository for each language. Links are provided in the [overview](github-actions-integration.md#overview-of-github-actions-integration) on this page.

Note that GitHub Actions will not pass on secrets set in the repository to forks being used in pull requests, and so the Snyk actions that require the token will fail to run.

## Use your own development environment

The Actions for each language automatically install all the required development tools for Snyk to determine the correct dependencies and hence vulnerabilities from different language environments. If you have a workflow where you already have the development tools installed, you can instead use the `snyk/actions/setup` Action to just install [Snyk CLI](https://github.com/snyk/snyk):

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
          go-version: '1.13'
      - name: Snyk monitor
        run: snyk test
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

The example here uses `actions/setup-go`. You would need to select the right actions to install the relevant development requirements for your project. If you are already using the same pipeline to build and test your application you are likely already doing so.

## Getting your Snyk token

The Actions examples on this page refer to a Snyk API token:

```yaml
env:
  SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

Every Snyk account has this token. Once you [create an account](https://docs.snyk.io/getting-started/create-a-snyk-account) you can find the API token in one of two ways:

1. In the UI, go to your Snyk account [settings page](https://app.snyk.io/account) and retrieve the API token, as explained in  [Revoking and regenerating Snyk API tokens](../../snyk-api-info/revoking-and-regenerating-snyk-api-tokens.md).
2. If you're using the [Snyk CLI](https://docs.snyk.io/snyk-cli) locally you can retrieve the API token by running `snyk config get api`.

## GitHub Code Scanning support

All Snyk GitHub Actions support integration with GitHub Code Scanning to show vulnerability information in the GitHub Security tab. Refer to the repository for the Action for each language. Links are provided in the [overview](github-actions-integration.md#overview-of-github-actions-integration) on this page.

## Continuing on error

**When you use the Actions as shown in the examples, the workflow fails when issues are found.** To continue the Action even if Snyk finds vulnerabilities, use [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error)..

```yaml
name: Example workflow using Snyk with continue on error
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        continue-on-error: true
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```
