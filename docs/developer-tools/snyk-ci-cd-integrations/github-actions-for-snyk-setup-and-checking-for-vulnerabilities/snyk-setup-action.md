# Snyk Setup action

The [Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup) provides a way to install the Snyk CLI to check for vulnerabilities. For information about when to use this action, see [Use your own development environment](./#use-your-own-development-environment).

You can use the Snyk Setup Action to install Snyk to check for vulnerabilities as follows:

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
        go-version: "1.13"
    - name: Snyk test
      run: snyk test
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

When using the Setup Action, you are responsible for setting up the development environment required to run Snyk. In this case, this is a Go project, so `actions/setup-go` was used, but this would be specific to your Project. The [GitHub language and frameworks guides](https://docs.github.com/en/actions/language-and-framework-guides) are a good starting point.

The Snyk Setup Action has a property that is passed to the underlying image using `with`.

| Property     | Default | Description                         |
| ------------ | ------- | ----------------------------------- |
| snyk-version | latest  | Install a specific version of Snyk. |

The Action also has output:

| Property | Default | Description                                 |
| -------- | ------- | ------------------------------------------- |
| version  |         | The full version of the Snyk CLI installed. |

For example, you can choose to install a specific version of Snyk and grab the installed version from the output:

```yaml
name: Snyk example
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - uses: snyk/actions/setup@master
      id: snyk
      with:
        snyk-version: v1.391.0
    - uses: actions/setup-go@v1
      with:
        go-version: "1.13"
    - name: Snyk version
      run: echo "${{ steps.snyk.outputs.version }}"
    - name: Snyk monitor 
      run: snyk monitor
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```
