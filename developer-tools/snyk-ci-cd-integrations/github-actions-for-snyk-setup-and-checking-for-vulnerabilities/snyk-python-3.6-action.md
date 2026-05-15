# Snyk Python-3.6 action

{% hint style="warning" %}
This image has been removed on 12 Aug 2024. It is highly recommended that users consider migrating to a newer action to ensure continued support and up-to-date functionality. If you are currently using this image, plan an upgrade as soon as possible to avoid any disruptions in your workflow post this date.
{% endhint %}

This page provides examples of using the Snyk GitHub action for [Python (3.6)](https://github.com/snyk/actions/tree/master/python-3.6). For instructions on using the action and further information, see [GitHub Actions for Snyk setup and checking for vulnerabilities](./).

## Using the Snyk Python (3.6) action to check for vulnerabilities

The examples that follow show how you can use a Snyk Python GitHub action.

Snyk requires that Python download the dependencies before running or triggering the Snyk checks.

The Python image checks and installs dependencies only if the manifest files are present in the current path, that is, the path from where the action is being triggered.

* If pip is present on the current path , and Snyk finds a `requirements.txt` file, then Snyk runs `pip install -r requirements.txt`.
* If pipenv is present on the current path, and Snyk finds a `Pipfile` without a `Pipfile.lock`, then Snyk runs `pipenv update`.
* If `pyproject.toml` is present in the current path and Snyk does not find `poetry.lock` then Snyk runs `pip install poetry`.

If manifest files are present under any location other root then they must be installed prior to running Snyk.

You can use the Snyk Python (3.6) action to check for vulnerabilities as follows:

```yaml
name: Example workflow for Python-3.6 using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.6@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

You can use the Snyk Python (3.6) action to check for only high severity vulnerabilities as follows:

```yaml
name: Example workflow for Python-3.6 using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.6@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## Using the Snyk Python (3.6) action to run snyk monitor

For an example of running `snyk monitor`, see this [Snyk monitor example](./#snyk-monitor-example).

## Uploading Snyk scan results to GitHub Code Scanning using the Snyk Python (3.6) action

Using `--sarif-file-output` [Snyk CLI option](../../snyk-cli/cli-commands-and-options-summary.md) and the [GitHub SARIF upload action](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github), you can upload Snyk scan results to GitHub Code Scanning as shown in the example that follows.

The Snyk action fails when vulnerabilities are found. This would prevent the SARIF upload action from running. Thus, you must use a [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) option as shown in this example:

```yaml
name: Example workflow for Python-3.6 using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.6@master
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
