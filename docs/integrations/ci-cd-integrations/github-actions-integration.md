# GitHub Actions integration

Snyk offers a [set of GitHub actions](https://github.com/snyk/actions) for using Snyk to check for vulnerabilities in your GitHub projects. These actions are are based on the [Snyk CLI](https://docs.snyk.io/snyk-cli/guides-for-our-cli/cli-reference).

For more information about GitHub actions see the [GitHub site ](https://github.com/features/actions)and the [GitHub custom actions](https://docs.github.com/en/actions/creating-actions/about-actions) documentation.

A different action is required depending on which language or build tool you are using. Snyk currently offers the following actions (during development of these pages find the content on [GitHub](https://github.com/snyk/actions)):

* [Snyk CocoaPods Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-cocoapods-action)
* [Snyk dotNET Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-dotnet-action)
* [Snyk Golang Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-golang-action)
* [Snyk Gradle Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-action)
* [Snyk Gradle-jdk11 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-jdk11-action)
* [Snyk Gradle-jdk12 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-jdk12-action)
* [Snyk Gradle-jdk14 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-jdk14-action)
* [Snyk Gradle-jdk16 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-jdk16-action)
* [Snyk Gradle-jdk17 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-gradle-jdk17-action)
* [Snyk Maven Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-maven-action)
* [Snyk Maven-3-jdk-11 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-maven-3-jdk-11-action)
* [Snyk Node Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-node-action)
* [Snyk PHP Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-php-action)
* [Snyk Python Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-python-action)
* [Snyk Python-3.6 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-python-3.6-action)
* [Snyk Python-3.7 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-python-3.7-action)
* [Snyk Python-3.8 Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-python-3.8-action)
* [Snyk Ruby Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-ruby-action)
* [Snyk Scala Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-scala-action)
* [Snyk Docker Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-docker-action)
* [Snyk Infrastructure as Code Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-infrastructure-as-code-action)
* [Snyk Setup Action](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration/snyk-setup-action)

Here's an example of using one of the Actions, in this case to test a Node.js project:

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

If you want to send data to Snyk, and be alerted when new vulnerabilities are discovered, you can run [Snyk monitor](https://support.snyk.io/hc/en-us/articles/360000920818-What-is-the-difference-between-snyk-test-protect-and-monitor-) like so:

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

See the individual Actions linked above for per-language instructions.

Note that GitHub Actions will not pass on secrets set in the repository to forks being used in pull requests, and so the Snyk actions that require the token will fail to run.

## Bring your own development environment

The per-language Actions automatically install all the required development tools for Snyk to determine the correct dependencies and hence vulnerabilities from different language environments. If you have a workflow where you already have those installed then you can instead use the `snyk/actions/setup` Action to just install [Snyk CLI](https://github.com/snyk/snyk):

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

The example here uses `actions/setup-go` would you would need to select the right actions to install the relevant development requirements for your project. If you are already using the same pipeline to build and test your application you're likely already doing so.

## Getting your Snyk token

The Actions example above refer to a Snyk API token:

```yaml
env:
  SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

Every Snyk account has this token. Once you [create an account](https://snyk.co/SignUpGH) you can find it in one of two ways:

1. In the UI, go to your Snyk account's [settings page](https://app.snyk.io/account) and retrieve the API token, as shown in the following [Revoking and regenerating Snyk API tokens](https://support.snyk.io/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens).
2. If you're using the [Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI) locally you can retrieve it by running `snyk config get api`.

## GitHub Code Scanning support

All Snyk GitHub Actions support integration with GitHub Code Scanning to show vulnerability information in the GitHub Security tab. You can see full details on the individual action READMEs linked above.

## Continuing on error

The above examples will fail the workflow when issues are found. If you want to ensure the Action continues, even if Snyk finds vulnerabilities, then [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) can be used..

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
