# GitHub actions migration

This page explains how to transition away from affected GitHub actions

Snyk recommends that you update the affected workflows to use a newer action that is not slated for removal.

## Alternative (a) switch to an alternate supported version of the software

### Follow these steps for Python-3.6/Python-3.7

1. Identify and find the Actions:
   * Determine which action in your workflow needs to be updated.
   * In this case, you are looking to replace the `python-3.6` action with an action that is available in the snyk build tool chain, such as [python-3.10](https://github.com/snyk/actions/tree/master/python-3.10)
2. Update the workflow file:
   * Open the workflow file where the current action is defined.
   * Locate the section that specifies the current action, such as `python:3.6`.
   * Replace the current action with the newer action.
3. Save your changes: Save the updated workflow file with the new action version.
4. Test the workflow: Run a test on the updated workflow to ensure that the new action functions as expected.

<mark style="color:red;">Example before:</mark>

```yaml
name: Example workflow for Python-3.6 using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.6@master // <- Using python 3.6
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

<mark style="color:green;">Example after:</mark>

```yaml
name: Example workflow for Python using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.10@master // <- Using python 3.10
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

### Follow these steps for scala/sbt <a href="#a.2-please-follow-these-steps-for-scala-sbt" id="a.2-please-follow-these-steps-for-scala-sbt"></a>

1. Identify and find the Actions:
   * Determine which action in your workflow needs to be updated.
   * In this case, you are looking to replace the `scala` action with an action that is available in the snyk build tool chain, such as [https://github.com/snyk/actions/tree/master/sbt1.10.0-scala3.4.2](https://github.com/snyk/actions/tree/master/sbt1.10.0-scala3.4.2)
2. Update the workflow file:
   * Open the workflow file where the current action is defined.
   * Locate the section that specifies the current action, such as `scala`.
   * Replace the current action with the newer action, `sbt1.10.0-scala3.4.2@master`.
3. Save your changes: Save the updated workflow file with the new action version.
4. Test the workflow: Run a test on the updated workflow to ensure that the new action functions as expected.

<mark style="color:red;">Example before:</mark>

```yaml
name: Example workflow for Scala using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/scala@master // <- Using old scala action
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

<mark style="color:green;">Example after:</mark>

```yaml
name: Example workflow for Scala using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/sbt1.10.0-scala3.4.2@master // <- Using new scala action
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      
```

## Alternative (b) create your own custom actions <a href="#b.-you-can-roll-your-own-custom-actions" id="b.-you-can-roll-your-own-custom-actions"></a>

Snyk customers who prefer to move away from pre-built actions provided by Snyk can create custom actions tailored to their specific needs. This approach allows for greater customization and control over the actions used in their workflows.

By creating your own actions, you can avoid the effects of future cleanup and removal events when images and actions lose vendor support.

### Leveraging the [Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup) <a href="#b.1-leveraging-the-snyk-setup-action" id="b.1-leveraging-the-snyk-setup-action"></a>

This [action](../github-actions-for-snyk-setup-and-checking-for-vulnerabilities/snyk-setup-action.md) offers a versatile method of incorporating Snyk into your workflows effectively.

Consider using this action when:

* You have a workflow where you already have the development tools installed
* You do not want to depend on a predefined Snyk action for a specific environment, but still want a robust way to set up the Snyk CLI for your workflows
* You are unable to find an action built for your specific environment

### Direct CLI Installation <a href="#b.2-direct-cli-installation" id="b.2-direct-cli-installation"></a>

Another option is to install and use the Snyk CLI directly in your GitHub Actions workflow. This method allows you to skip the requirement of dedicated GitHub Actions integration.

```yaml
name: Example workflow using Snyk
on: push
jobs:
  snyk_scan:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Install Snyk CLI
      run: |
        curl https://downloads.snyk.io/cli/stable/snyk-linux -o snyk-linux
        curl https://downloads.snyk.io/cli/stable/snyk-linux.sha256 -o snyk.sha256
        sha256sum -c snyk.sha256
        chmod +x snyk-linux
        sudo mv snyk-linux /usr/local/bin/snyk
    - name: Run Snyk to test project dependencies
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      run: |
        snyk test
```
