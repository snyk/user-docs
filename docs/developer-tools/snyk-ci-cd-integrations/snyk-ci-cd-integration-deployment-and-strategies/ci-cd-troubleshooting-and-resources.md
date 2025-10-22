# CI/CD troubleshooting and resources

This page provides a few tips to help troubleshoot or scale CI/CD integrations.

### Step 1: Try to replicate with the Snyk CLI

If the CLI and the pipeline are running the same engine, try to clone the Project and scan with the CLI.

Try various CLI options. Use the Snyk CLI to find and fix known vulnerabilities as you run it in the pipeline. For more information, see the [CLI documentation](../../snyk-cli/).

### Step 2: Get logs

If you can replicate the issue using the Command Line Interface (CLI) and the problem still exists, consult the [Debugging the Snyk CLI ](../../snyk-cli/debugging-the-snyk-cli.md) troubleshooting guidelines for capturing logs in a debug mode.

### Step 3: Use the CLI instead of the plugin

Try to replace the native plugin with the CLI by installing the CLI. See [Install the Snyk CLI ](../../snyk-cli/install-or-update-the-snyk-cli/)for instructions.

{% hint style="info" %}
The following repository provides some examples of binary and npm integrations for various CI/CD tools: [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).

To learn more about CI/CD, see the article [What is CI/CD? CI/CD Pipeline and Tools Explained](https://snyk.io/learn/what-is-ci-cd-pipeline-and-tools-explained/).
{% endhint %}
