# Snyk PHP action

This page provides examples of using the Snyk GitHub action for [PHP](https://github.com/snyk/actions/tree/master/php). For instructions on using the action and further information, see [GitHub Actions for Snyk setup and checking for vulnerabilities](./).

## Using the Snyk PHP action to check for vulnerabilities

You can use the Snyk PHP Action to check for vulnerabilities as follows:

```yaml
name: Example workflow for PHP using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/php@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

You can use the Snyk PHP action to check for only high severity vulnerabilities as follows:

```yaml
name: Example workflow for PHP using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/php@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## Using the Snyk PHP Action to run snyk monitor

For an example of running `snyk monitor`, see this [Snyk monitor example](./#snyk-monitor-example).

## Uploading Snyk scan results to GitHub Code Scanning using the Snyk PHP Action

Using `--sarif-file-output` [Snyk CLI option](../../snyk-cli/cli-commands-and-options-summary.md) and the [GitHub SARIF upload action](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github), you can upload Snyk scan results to GitHub Code Scanning as shown in the example that follows.

The Snyk action fails when vulnerabilities are found. This would prevent the SARIF upload action from running. Thus, you must use a [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) option as shown in this example:

```yaml
name: Example workflow for PHP using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/php@master
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

{% hint style="info" %}
To use the upload-sarif option for private repositories, you must have GitHub Advanced Security.

If you see the error `Advanced Security must be enabled for this repository to use code scanning`, check that GitHub Advanced Security is enabled. For more information, see [Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository).
{% endhint %}
