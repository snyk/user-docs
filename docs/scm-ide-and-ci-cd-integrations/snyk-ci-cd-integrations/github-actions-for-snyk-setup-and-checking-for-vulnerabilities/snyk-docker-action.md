# Snyk Docker action

This page provides instructions for and examples of using the Snyk GitHub Action for [Docker](https://github.com/snyk/actions/tree/master/docker). For general instructions and information, see [GitHub Actions for Snyk setup and checking for vulnerabilities](./).

In order to use the Docker Action, you must have a Snyk API token. See [Getting your Snyk token](./#getting-your-snyk-token), or you can [sign up for free](https://snyk.io/login).

## Using the Snyk Docker Action to check for vulnerabilities

You can use the Snyk Docker Action to check for vulnerabilities in your Docker images as follows:

```yaml
name: Example workflow for Docker using Snyk 
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - name: Run Snyk to check Docker image for vulnerabilities
      uses: snyk/actions/docker@master
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      with:
        image: your/image-to-test
```

## Snyk Docker Action properties

The Snyk Docker Action has properties that are passed to the underlying image. These are passed to the action using `with`.

| Property | Default | Description                                                 |
| -------- | ------- | ----------------------------------------------------------- |
| args     |         | Override the default arguments to the Snyk image.           |
| command  | test    | Specify which command to run, for instance test or monitor. |
| image    |         | The name of the image to test.                              |
| json     | false   | In addition to the stdout, save the results as snyk.json.   |
| sarif    | true    | In addition to the stdout, save the results as snyk.sarif.  |

For example, you can use the Snyk Docker Action to check for only high severity vulnerabilities as follows:

```yaml
name: Example workflow for Docker using Snyk 
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - name: Run Snyk to check Docker images for vulnerabilities
      uses: snyk/actions/docker@master
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      with:
        image: your/image-to-test
        args: --severity-threshold=high
```

## Uploading Snyk scan results to GitHub Code Scanning using the Snyk Docker Action

The Docker Action also supports integrating with GitHub Code Scanning and can show issues in the GitHub Security tab. As long as you reference a Dockerfile with `--file=Dockerfile` in the `args`, a `snyk.sarif` file will be generated, which can be uploaded to GitHub Code Scanning.

```yaml
name: Snyk Container
on: push
jobs:
  snyk:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Build a Docker image
      run: docker build -t your/image-to-test .
    - name: Run Snyk to check Docker image for vulnerabilities
      # Snyk can be used to break the build when it detects vulnerabilities.
      # In this case we want to upload the issues to GitHub Code Scanning
      continue-on-error: true
      uses: snyk/actions/docker@master
      env:
        # In order to use the Snyk Action you will need to have a Snyk API token.
        # See https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#getting-your-snyk-token
        # or you can sign up for free at https://snyk.io/login
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      with:
        image: your/image-to-test
        args: --file=Dockerfile
    - name: Upload result to GitHub Code Scanning
      uses: github/codeql-action/upload-sarif@v3
      with:
        sarif_file: snyk.sarif
```

{% hint style="info" %}
To use the upload-sarif option for private repositories, you must have GitHub Advanced Security.

If you see the error `Advanced Security must be enabled for this repository to use code scanning`, check that GitHub Advanced Security is enabled. For more information, see [Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository).
{% endhint %}
