# CI/CD troubleshooting and resources

## CI/CD troubleshooting

This section provides a few tips to help troubleshoot or scale CI/CD integrations.

### Step 1: Try to replicate with Snyk CLI

If CLI and pipeline are running the same engine, try to clone the project and scan with CLI.

Try various CLI options. Use the Snyk CLI tool to find and fix known vulnerabilities as you run it in the pipeline. For more information see the [CLI docs](../../../snyk-cli/).

### Step 2: Get logs

If you are able to replicate with CLI and the problem still exists, ask the CLI to output the debug logging using the following command: `DEBUG=*` or use the `-d` option to capture logs:

```
snyk test -d
```

or

```
DEBUG=* snyk test
```

### Step 3: Use the CLI instead of the plugin

Try to replace the native plugin with the CLI by installing the CLI. See [Install the Snyk CLI ](../../../snyk-cli/install-or-update-the-snyk-cli/)for instructions.

## **Useful resources for CI/CD**

The following repo shares some examples of binary and npm integrations for various CI/CD tools: [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).

To earn more about CI/CD see [What is CI/CD? CI/CD Pipeline and Tools Explained](https://snyk.io/learn/what-is-ci-cd-pipeline-and-tools-explained/).
