# Publish Snyk Code CLI results

Using Snyk Code, you can publish results to a Snyk Project with or without using an integration. The sections of this page explain how:

* [Publish CLI results to a Snyk Code Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#publish-cli-results-to-a-snyk-code-project) created using the CLI
  * Run Snyk Code in the CI/CD and append results to Snyk Projects for reports.&#x20;
  * Issues that are ignored in the WebUI Project will apply to the CLI scans.&#x20;
  * This does not require an [SCM Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).&#x20;
* [Test and publish CLI results to an existing Snyk Code SCM Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#test-and-publish-cli-results-to-an-existing-snyk-code-scm-project)
  * Run Snyk Code in the CI/CD, and push the results to an existing SCM Code Analysis Project to get Code Snippets.&#x20;
  * Issues that are ignored in the WebUI Project will apply to the CLI scans.&#x20;
  * The command will scan the specific code in the repository, not locally, as you do when you publish Snyk CLI results to a Code Project.&#x20;
  * The results are appended to the existing SCM Project, which requires an [SCM Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

If a CLI-based Snyk Code Project does not yet exist for the value provided in the `--project-name` option, Snyk creates a new CLI-based Project.&#x20;

A new snapshot is created under the same Project if a CLI-based Project already exists.

The identity is based on `--project-name` or `--target-name`.

This page also explains how to [ignore CLI results](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#ignore-cli-results).

See the CLI help for more information about `snyk code test`, such as [exit codes](../../commands/code-test.md#exit-codes) and [available options](../../commands/code-test.md#options).

## **Publish CLI results to a Snyk Code Project**

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI results is in [Closed Beta](../../../more-info/snyk-feature-release-process.md#closed-beta) and available only by invitation from Snyk.

The minimum supported CLI version is v1.1194.0.
{% endhint %}

Using Snyk Code through the CLI supports publishing test results to a Snyk Project in the Web UI and respecting issues that were ignored in a Snyk Project in the Web UI so you can filter them from the analysis results.

This allows you to use Snyk Code as a blocking CI/CD gate to test and block builds at the main branch level and then have developers review the results in the Web UI, fix any newly introduced vulnerabilities, or ignore irrelevant ones.

In the terminal, enter the following command:

```
snyk code test --report --project-name="<PROJECT_NAME>"
```

* PROJECT\_NAME must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The Project name must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).
* There is a temporary limit of 6MB for the resulting payload. If the SARIF output is larger than 6MB, the process will not complete.

Running the `snyk code test` command with the `--report` option as shown returns the results to the terminal window, along with a URL to the Snyk Code Project where the results have been published. Refer to the following screenshot.

<figure><img src="../../../.gitbook/assets/image (2) (6).png" alt="Snyk code test results with --report option"><figcaption><p>Snyk code test results with --report option</p></figcaption></figure>

If a Snyk Code Project created using the CLI does not yet exist for the value provided in the `--project-name` option, the Snyk CLI creates a new Project. If a Project created using the CLI exists, a new snapshot is created under the same Project.

The following command will create or upload an existing Project named \<PROJECT\_NAME> under a target named \<TARGET\_NAME>

* Both PROJECT\_NAME and TARGET\_NAME must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The project name and target name must each contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>"
```

The following command will create or upload an existing Project named \<PROJECT-NAME> under a target named \<TARGET-NAME> and grouped by the "$(git branch --show-current)" branch name.

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>" --target-reference="$(git branch --show-current)"
```

<figure><img src="../../../.gitbook/assets/image (4) (4).png" alt="Code analysis Projects grouped by branch"><figcaption><p>Code analysis Projects grouped by branch</p></figcaption></figure>

### **Troubleshooting**

#### **There was a problem running Code analysis. The findings for this project may exceed the allowed size limit.**

If you encounter this error after running the Code Analysis, it indicates that the contents of the scanned Project exceed the 6MB limit. To resolve:

To complete the scan, consider the following troubleshooting steps:

* Partition the Project repository by scanning sub-directories instead of the whole repository. For example:
  * Create two Projects for your frontend and backend directories, and scan them separately.
  * Create and scan Projects for each MicroService.
* Exclude unnecessary files from the scanning process using the [.snyk exclude option](../../../scan-using-snyk/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process.md#exclusion-syntax-of-the-.snyk-file). For example, you can exclude test files from the scan.
* Set a severity threshold using the [`--severity-threshold=high`](../failing-of-builds-in-snyk-cli.md#combining-security-policies-with-severity-threshold). You can focus on more critical issues and gain visibility into urgent matters.

## **Test and publish CLI results to an existing Snyk Code SCM Project**

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI issues in SCM Projects is in [Closed Beta](../../../more-info/snyk-feature-release-process.md#closed-beta) and available only by invitation from Snyk.

The minimum supported CLI version is 1.1194.0.
{% endhint %}

Using Snyk Code via the CLI supports publishing test results to an existing Snyk Code SCM Project in the Web UI and respecting issues ignored in a Snyk Project in the Web UI so you can filter them from the analysis results. This provides [Snyk Code Data flow](broken-reference) for the CLI results snapshot.

Be sure you have the prerequisites and publish the results as you intend:

* An [SCM Integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/) is required, as the CLI results are appended there. &#x20;
* The Project ID is required. You can find it as follows:
  * Find the Project ID in the URL of the Project in the Snyk Web UI: https://app.snyk.io/org/org\_name/project/PROJECT\_UUID
  * Or use the GET Projects method through the [Snyk API.](https://apidocs.snyk.io/?version=2023-05-29#get-/orgs/-org\_id-/projects)&#x20;
* The CLI results are pushed to the SCM Project that is monitoring your **default branch**. Consider using the command on your default branch.
* In the SCM integration, three snapshots are stored at a time. To save scan results, consider outputting a static [JSON or SARIF file](broken-reference) for future reference.

Use the following command to publish to a specific Project:

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```

## &#x20;**Ignore CLI results**

You can use the Ignore button to ignore issues from CLI results in the Web UI both to [publish CLI results to a Snyk Code Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#publish-cli-results-to-a-snyk-code-project) and to [test and publish CLI results to an existing Snyk Code SCM Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#test-and-publish-cli-results-to-an-existing-snyk-code-scm-project).

<figure><img src="../../../.gitbook/assets/image (1) (7) (1).png" alt="Ignoring issues in the Web UI"><figcaption><p>Ignoring issues in the Web UI</p></figcaption></figure>

{% hint style="info" %}
&#x20;[snyk-to-html](https://github.com/snyk/snyk-to-html) currently does not honor the ignored issues. Anything that is ignored in the Web UI is NOT ignored in the report that `snyk-to-html` generates.
{% endhint %}

Issues that are ignored in the Web UI are ignored in CLI tests when you use the following command:&#x20;

```
snyk code test --report --project-name="PROJECT_NAME"
```

Ignores that have been applied to the Project with a `PROJECT_NAME` suppress the issue the next time that the CLI runs for the same `PROJECT_NAME`.

To ignore the CLI issues appended in the SCM Project, you can use the following command :

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```
