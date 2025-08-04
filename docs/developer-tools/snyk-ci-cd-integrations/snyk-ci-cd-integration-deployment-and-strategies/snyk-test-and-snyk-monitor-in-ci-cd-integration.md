# Snyk test and snyk monitor in CI/CD integration

Depending on your approach and goals for your Snyk Open Source Project, you may need to use both the `snyk monitor` and `snyk test` commands in your pipeline. Examples and details follow.

## **CLI examples in a build pipeline**

Use `snyk monitor` to expose vulnerabilities and post to the Snyk UI for ongoing monitoring:

```
snyk monitor --all-projects --org=snyk-apps
```

Use `snyk test` to fail the build on high-severity issues:

```
snyk test --all-projects --org=snyk-apps --severity-threshold=high
```

To see the full list of options in the CLI, run the `snyk test --help`, `snyk monitor --help`, and `snyk container --help` commands or see the [CLI help](../../snyk-cli/commands/).

## **Exit Codes**

The `snyk test` command is synchronous; it ends with an exit code. Your build system can use exit codes to either pass or fail the build based on the test results. See the [CLI help](../../snyk-cli/commands/) for the command you are using to find the meaning of the exit codes.

The `snyk monitor` command posts a snapshot of the dependency tree for your Project to your Snyk account and monitors that snapshot for vulnerabilities. It is an asynchronous command that does not end with an exit code based on the vulnerability status. For `snyk monitor`, exit codes signify success or failure in creating the snapshot to monitor.

To silence Snyk CLI exit codes for the `snyk test` command to avoid failing the build step, use `|| true` at the end of the command. `|| true` sets the exit code of the scan to 0. This can be used to continue with a CI/CD pipeline even when there are vulnerabilities.

## Common CLI options in a CI/CD integration

Among the most common options used in a CI/CD integration are the following:

`-- all-projects`: Auto-detect all Projects in the working directory

`-p`: Prune dependency trees, removing duplicate sub-dependencies. Continues to find all vulnerabilities, but may not find all of the vulnerable paths.

\--org=\<ORG\_ID>: Specify the `ORG_ID` to run Snyk commands for a specific Organization. This influences where new Projects are created after running the `monitor` command, some availabilty of features, and private test limits. If you have multiple Organizations, you can set a default from the CLI using:

```
$ snyk config set org=<ORG_ID>
```

Set a default to ensure all newly tested or monitored Projects are tested under your default Organization. If you need to override the default, use the `--org=<ORG_ID>` option.

The default`<ORG_ID>` is the current preferred Organization in your Snyk [Account settings](https://app.snyk.io/account).

## Configuring options for failing builds

You can add options to the `snyk test` command to refine parameters that can result in a failed build:

* `--severity-threshold=high`: Fail the build only for high-severity issues
* `--fail-on=upgradable`: fail the build only for issues that are upgradable (can be fixed with Snyk fix advice)

You can also fail the build for any other parameter in the Snyk JSON output (such as CVSS score), using a wrapper like [`snyk-filter`](https://github.com/snyk-tech-services/snyk-filter), or using additional tooling like [`snyk-delta`](https://github.com/snyk-tech-services/snyk-delta) to fail the build only for issues found since the last build. For information on using `snyk-filter` and `snyk-delta` see [Failing of builds in the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli.md).

## Ignoring issues

By default, if issues are not ignored, or if you are not using [`snyk-delta`](https://github.com/snyk-tech-services/snyk-delta), a `snyk test` in your pipeline fails the build when issues are found. To allow builds to continue without resolving these issues, you can:

* [Ignore issues using a .snyk policy file](../../snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md)
* [Ignore issues from the Snyk UI](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/#ignore-issues-in-the-snyk-web-ui)
* [Ignore issues using the Ignores API](../../../snyk-api/reference/ignores-v1.md)
* Use the Snyk Python API for bulk ignores: see [pysnyk](https://github.com/snyk-labs/pysnyk) and the demo [bulk-ignore-vulns-by-issueIdList](https://github.com/snyk-labs/pysnyk/blob/master/examples/api-demo-9c-bulk-ignore-vulns-by-issueIdList.py).

## Creating custom build artifacts

You can use JSON output from Snyk commands to create custom test reports as build artifacts, using the [`snyk-to-html`](https://github.com/snyk/snyk-to-html) utility or other custom processing you develop. For more information, see the [`snyk-to-html`](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) documentation.

## Creating work items for new vulnerabilities

Snyk allows you to automatically create new work items in JIRA; for details, see the [Jira integration ](../../../integrations/jira-and-slack-integrations/jira-integration.md)documentation. You can customize this code for your specific requirements or adapt it to work with other work management systems.

See [Jira tickets for new vulns](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns) to get started. You can also use the API endpoints [Create jira issue](../../../snyk-api/reference/jira-v1.md#org-orgid-project-projectid-issue-issueid-jira-issue) and [List all jira issues](../../../snyk-api/reference/jira-v1.md#org-orgid-project-projectid-jira-issues).
