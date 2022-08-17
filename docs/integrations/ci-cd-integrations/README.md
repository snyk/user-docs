# Snyk CI/CD Integration

## CI/CD troubleshooting and advanced tips

This section provides a few tips to help troubleshoot or scale CI/CD integrations.

### Step 1: Try to replicate with Snyk CLI

If CLI and pipeline are running the same engine, try to clone the project and scan with CLI.

Play with the CLI options. Use the Snyk CLI tool to find and fix known vulnerabilities as you run it in the pipeline. For more information see the [CLI reference](../../snyk-cli/cli-reference.md).

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

Try to replace the native plugin with the CLI by installing the CLI. See [Install the Snyk CLI ](../../snyk-cli/install-the-snyk-cli.md)for instructions.

## Common CLI options in a CI/CD integration

Among the most common options used in a CI/CD integration are the following:

`-- all-projects`: Auto-detect all projects in working directory

`-p`: Prune dependency trees, removing duplicate sub-dependencies. Continues to find all vulnerabilities, but may not find all of the vulnerable paths.

\--org=\<ORG\_ID>: Specify the ORG\_ID to run Snyk commands tied to a specific organization. This influences where new projects are created after running the `monitor` command, some features availability, and private tests limits. If you have multiple organizations, you can set a default from the CLI using:

```
$ snyk config set org=<ORG_ID>
```

Set a default to ensure all newly tested or monitored projects are tested under your default organization. If you need to override the default, use the `--org=<ORG_ID>` option.

Default: `<ORG_ID>` that is the current preferred organization in your [Account settings](https://app.snyk.io/account).

## **Useful resources**

The following repo shares some examples of binary and npm integrations for various CI/CD tools: [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).

To earn more about CI/CD see [What is CI/CD? CI/CD Pipeline and Tools Explained](https://snyk.io/learn/what-is-ci-cd-pipeline-and-tools-explained/).

##
