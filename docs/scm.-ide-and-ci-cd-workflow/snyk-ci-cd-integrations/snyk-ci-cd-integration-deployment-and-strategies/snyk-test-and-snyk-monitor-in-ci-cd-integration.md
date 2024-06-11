# Snyk test and snyk monitor in CI/CD integration

Depending on your approach and goals, you may need to use both the `snyk monitor` and `snyk test` commands in your pipeline. Examples and details follow.

## **CLI examples in a build pipeline**

Use `snyk monitor` to expose vulnerabilities and post to the Snyk UI for ongoing monitoring:

```
snyk monitor --all-projects --org=snyk-apps
```

Use `snyk test` to fail the build on high severity issues:

```
snyk test --all-projects --org=snyk-apps --severity-threshold=high
```

To see the full list of options in the CLI, run the `snyk test --help`, `snyk monitor --help`, and `snyk container --help` commands or see the [help docs](../../../snyk-cli/commands/).

## **Exit Codes**

The `snyk test` command is synchronous; it ends with an exit code. Your build system can use exit codes to either pass or fail the build based on the test results. See the [help docs](../../../snyk-cli/commands/) for the command you are using to find the meaning of the exit codes.

The `snyk monitor` command posts a snapshot of the dependency tree for your project to your Snyk account and monitors that snapshot for vulnerabilities. It is an asynchronous command that does not end with an exit code based on the vulnerability status. For `snyk monitor`, exit codes signify success or failure in creating the snapshot to monitor.

To silence Snyk CLI exit codes for the `snyk test` command to avoid failing the build step, use `|| true` at the end of the command. `|| true` sets the exit code of the scan to 0. This can be used to continue with a CI/CD pipeline even when there are vulnerabilities.

## Common CLI options in a CI/CD integration

Among the most common options used in a CI/CD integration are the following:

`-- all-projects`: Auto-detect all projects in working directory

`-p`: Prune dependency trees, removing duplicate sub-dependencies. Continues to find all vulnerabilities, but may not find all of the vulnerable paths.

\--org=\<ORG\_ID>: Specify the ORG\_ID to run Snyk commands for a specific organization. This influences where new projects are created after running the `monitor` command, some features availability, and private tests limits. If you have multiple organizations, you can set a default from the CLI using:

```
$ snyk config set org=<ORG_ID>
```

Set a default to ensure all newly tested or monitored projects are tested under your default organization. If you need to override the default, use the `--org=<ORG_ID>` option.

The default`<ORG_ID>` is the current preferred organization in your [Account settings](https://app.snyk.io/account).

## Configuring options for failing builds

You can add options to the `snyk test` command to refine parameters that can result in a failed build:

* `--severity-threshold=high`: Fail the build only for high severity issues
* `--fail-on=upgradable`: fail the build only for issues that are upgradable (can be fixed with Snyk fix advice)

You can also fail the build for any other parameter in the Snyk JSON output (such as CVSS score), using a wrapper like [snyk-filter](https://github.com/snyk-tech-services/snyk-filter), or using additional tooling like [snyk-delta](https://github.com/snyk-tech-services/snyk-delta) to fail the build only for issues found since the last build. For information on using snyk-filter and snyk-delta see [Failing of builds in the Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli.md).

## Ignoring issues

By default if issues are not ignored, or if you are not using [snyk-delta](https://github.com/snyk-tech-services/snyk-delta), a `snyk test` in your pipeline fails the build when issues are found. To allow builds to continue without resolving these issues, you can:

* [Ignore issues using a .snyk policy file](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)
* [Ignore issues from the Snyk UI](https://support.snyk.io/hc/en-us/articles/360000923498-How-can-I-ignore-a-vulnerability-)
* [Ignore issues from the Snyk API](https://snyk.docs.apiary.io/#reference/projects/project-ignores)
* Use the Snyk Python API for bulk ignores: see [https://github.com/snyk-labs/pysnyk](https://github.com/snyk-labs/pysnyk) and [https://github.com/snyk-labs/pysnyk/blob/master/examples/api-demo-9c-bulk-ignore-vulns-by-issueIdList.py](https://github.com/snyk-labs/pysnyk/blob/master/examples/api-demo-9c-bulk-ignore-vulns-by-issueIdList.py)

## Creating custom build artifacts

You can use JSON output from Snyk commands to create custom test reports as build artifacts, using the [snyk-to-html](https://github.com/snyk/snyk-to-html) utility or other custom processing you develop.

## Creating work items for new vulnerabilities

Snyk allows you to automatically create new work items in JIRA (see [Jira integration](https://docs.snyk.io/integrations/untitled-3/jira) documentation). You can customize this code for your specific requirements, or adapt it to work with other work management systems.

See [Jira tickets for new vulns](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns) to get started, or review the [API to create Jira tickets](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues).
