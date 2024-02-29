# Publish Snyk Code CLI results and ignore issues

This page explains how you can publish Snyk Code results to a Snyk Project with or without using an integration.

You can [publish CLI results to a Snyk Code Project](publish-snyk-code-cli-results-and-ignore-issues.md#publish-cli-results-to-a-snyk-code-project) created using the CLI. Run Snyk Code in the CI/CD and upload results to a Snyk Project for reports. When you do this:

* After you have created a Snyk Code Project from the CLI, any issues that are ignored in the Web UI Project will apply to future CLI scans.
* A Git repository[ Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/) is not required.
* If a CLI-based Snyk Code Project does not yet exist for the value provided in the `--project-name` option, Snyk creates a new CLI-based Project.&#x20;
* If a CLI-based Snyk Code Project does exist, a new snapshot is created under the same Project.
* The identity is based on `--project-name` or `--target-name`.

You can [use the CLI to trigger a Snyk Code Git repository Project to test a specific commit in a Git repository](publish-snyk-code-cli-results-and-ignore-issues.md#use-the-cli-to-trigger-a-snyk-code-git-repository-project-to-test-a-specific-commit-in-a-git-reposit). Run Snyk Code in the CI/CD and specify an existing SCM Code Analysis Project along with the commit ID you wish to test, which will then also display code snippets. When you do this:

* Issues that are ignored in the Web UI Project will apply to scans triggered through this method.&#x20;
* The command will trigger a scan of the specified code in the repository, not local code, based on the commit ID you specified.&#x20;
* The results are saved as a new snapshot of the existing SCM Project, which requires a Git repository[ Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

This page also explains how to [ignore CLI results](publish-snyk-code-cli-results-and-ignore-issues.md#ignore-cli-results).

See the CLI help for more information about the `snyk code test` command, including [exit codes](../../commands/code-test.md#exit-codes) and [available options](../../commands/code-test.md#options).

## **Publish CLI results to a Snyk Code Project**

{% hint style="warning" %}
**Release status and feature availability**

Publishing and ignoring CLI results is in [Closed Beta](../../../getting-started/snyk-release-process.md#closed-beta) and available only for Enterprise plans.

For more information, see [Plans and pricing.](https://snyk.io/plans)

The minimum supported CLI version is v1.1194.0.
{% endhint %}

Using Snyk Code through the CLI allows you to publish test results of local code to a Snyk Project in the Web UI. Future CLI tests of this Project will respect issues that were ignored in the Web UI.

This allows using Snyk Code as a blocking CI/CD gate to test and block builds at the main branch level and then have developers review the results in the Web UI, fix any newly introduced vulnerabilities, or ignore irrelevant ones.

In the terminal, enter the following command:

```
snyk code test --report --project-name="<PROJECT_NAME>"
```

* `project-name` must be in double quotation marks. Single quotes or missing quotes will result in an error.
* `project-name` must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).
* There is a temporary limit of 6MB for the resulting payload. If the SARIF output is larger than 6MB, the process will not complete.

### Commands to publish Snyk Code CLI results

Running the `snyk code test` command with the `--report` option, as shown, returns the results to the terminal window, along with a URL to the Snyk Code Project where the results have been published. Refer to the following screenshot.

<figure><img src="../../../.gitbook/assets/image (2) (6).png" alt="Snyk code test results with --report option"><figcaption><p>Snyk code test results with --report option</p></figcaption></figure>

If a Snyk Code Project created using the CLI does not yet exist for the value provided in the `--project-name` option, the Snyk CLI creates a new Project. If a Project created by using the CLI exists, a new snapshot is created under the same Project.

To make the Project easier to interpret in the Web UI, you can use additional commands to specify a target name and also target references, such as Git branches. The following command will create or upload an existing Project named `<PROJECT_NAME>` under a target named `<TARGET_NAME>`.

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>"
```

* Both `<PROJECT_NAME>` and `<TARGET_NAME>` must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The Project name and Target name each must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).

The following command will create or upload an existing Project named `<PROJECT_NAME>` under a target named `<TARGET_NAME>` and grouped by the "`$(git branch --show-current)"` branch name.

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>" --target-reference="$(git branch --show-current)"
```

<figure><img src="../../../.gitbook/assets/image (4) (4).png" alt="Code analysis Projects grouped by branch"><figcaption><p>Code analysis Projects grouped by branch</p></figcaption></figure>

### **Troubleshooting publication of Snyk Code CLI results**

You may see this error:\
`There was a problem running Code analysis. The findings for this project may exceed the allowed size limit.`

This error indicates that the contents of the scanned Project exceed the 6MB limit. To complete the scan, consider the following troubleshooting steps:

* Partition the Project repository by scanning sub-directories instead of the whole repository, for example:
  * Create two Projects for your frontend and backend directories, and scan them separately.
  * Create and scan Projects for each MicroService.
* Exclude unnecessary files from the scanning process using the [.snyk exclude option](../../../scan-with-snyk/snyk-code/import-repository-to-snyk/excluding-directories-and-files-from-the-import-process.md#exclusion-syntax-of-the-.snyk-file). For example, you can exclude test files from the scan.
* Set a severity threshold using the [`--severity-threshold=high`](../failing-of-builds-in-snyk-cli.md#combining-security-policies-with-severity-threshold). You can focus on more critical issues and gain visibility into urgent matters.

## Use the CLI to trigger a test for a specific commit to an existing Snyk Code Git repository Project

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI issues in SCM Projects is in [Closed Beta](../../../getting-started/snyk-release-process.md#closed-beta) and available only by invitation from Snyk.

The minimum supported CLI version is 1.1194.0.
{% endhint %}

When you are using Snyk Code in the CI/CD, you can specify an existing Git repository Code Analysis Project along with the commit ID to trigger a test for which the results will be saved in the Web UI. The Web UI will also display Code Snippets that provide [Snyk Code Data flow](broken-reference). Ignores that have been added to the Web UI Project will be respected in this test.

Be sure you have the prerequisites and publish the results as you intend:

* A Git repository[ Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/) is required, as the CLI command triggers a test using a Git repository Project. &#x20;
* The Project ID is required. You can find the Project ID in the URL of the Project in the Snyk Web UI: `https://app.snyk.io/org/org_name/project/PROJECT_UUID.`
* After you run the CLI command, results are saved in the Git repository Project that you have specified. Given that scheduled tests for Git repository Projects scan the default branch, Snyk recommends using this CLI command with the Project ID for the Project where the default branch for the specified repository is being monitored.
* In the Code Analysis Project, three snapshots are stored at a time. Consider outputting a static [JSON](https://docs.snyk.io/snyk-cli/commands/code-test#json) or[ SARIF ](https://docs.snyk.io/snyk-cli/commands/code-test#sarif)file to save scan results for future reference.

Use the following command to publish to a specific Project:

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```

## &#x20;**Ignore CLI results for Snyk Code**

You can ignore issues in the Web UI. The ignores will be used both to [publish CLI results to a Snyk Code Project](publish-snyk-code-cli-results-and-ignore-issues.md#publish-cli-results-to-a-snyk-code-project) and in [using the CLI to trigger a test for a specific commit to an existing Snyk Code Git repository Project](publish-snyk-code-cli-results-and-ignore-issues.md#use-the-cli-to-trigger-a-test-for-a-specific-commit-to-an-existing-snyk-code-git-repository-project).

<figure><img src="../../../.gitbook/assets/image (1) (7) (1).png" alt="Ignoring issues in the Web UI"><figcaption><p>Ignoring issues in the Web UI</p></figcaption></figure>

{% hint style="info" %}
&#x20;[snyk-to-html](https://github.com/snyk/snyk-to-html) currently does not honor the ignored issues. Anything that is ignored in the Web UI is NOT ignored in the report that `snyk-to-html` generates.
{% endhint %}

For [publishing workflows](publish-snyk-code-cli-results-and-ignore-issues.md#publish-cli-results-to-a-snyk-code-project), after the CLI results are published to a Snyk Code Project, issues that were ignored in the Web UI will be ignored in CLI tests when you use the following command:

```
snyk code test --report --project-name="PROJECT_NAME"
```

* Assuming that the --project-name provided matches what is in the Web UI
* Ignores that have been applied to the Project with a `PROJECT_NAME` suppress the issue the next time that the CLI runs for the same `PROJECT_NAME`.

For the [trigger using the CLI workflow](publish-snyk-code-cli-results-and-ignore-issues.md#use-the-cli-to-trigger-a-snyk-code-git-repository-project-to-test-a-specific-commit-in-a-git-reposit),  ignores added to the Git repository monitored Project will also be ignored the next time an SCM scan is triggered from the CLI for a given `<PROJECT_UUID>` and `<COMMIT_ID>.`

An example follows:

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```
