# Publishing Snyk Code CLI results to a Snyk Project and ignoring CLI results (beta)

Snyk Code CLI supports publishing the results to a Snyk Project in the Web UI and respecting issues that were ignored in a Snyk Project in the Web UI so you can filter them from the analysis results.

## Before you begin

The two sections support different use cases:

1. [Publishing CLI results to a Snyk Code **CLI Project**](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#publishing-cli-results-to-a-snyk-code-cli-project)**:**&#x20;
   * Run Snyk Code in the CI/CD and append results to Snyk Projects for reports.&#x20;
   * Issues that are ignored in the WebUI Project will apply to the CLI scans.&#x20;
   * This does not require an [SCM Integration](../../../integrations/git-repository-scm-integrations/).&#x20;
2. [Test and publish CLI results to an existing Snyk Code **SCM Project**](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#test-and-publish-cli-results-to-an-existing-snyk-code-scm-project)**:**&#x20;
   * Run Snyk Code in the CI/CD, and push the results to an existing SCM Code Analysis Project to get Code Snippets.&#x20;
   * Issues that are ignored in the WebUI Project will apply to the CLI scans.&#x20;
   * The command will scan the specific code in the Repository, not locally, as you would do when you publish Snyk CLI results to a Code Project.&#x20;
   * This will be appended to the existing SCM Project, which requires an [SCM Integration](../../../integrations/git-repository-scm-integrations/).

## General considerations

* If a CLI-based Snyk Code Project does not yet exist for the value provided in the `--project-name` option, Snyk creates a new CLI-based Project.&#x20;
* A new snapshot is created under the same Project if a CLI-based Project already exists.
* The identity is based on `--project-name` or `--target-name`.

## **Publishing CLI results to a Snyk Code CLI Project**

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI results is in [Closed Beta](../../../more-info/snyk-feature-release-process.md#closed-beta) and only available by invitation from Snyk.

Minimum supported CLI version: v1.1194.0
{% endhint %}

Snyk Code CLI supports publishing the results to a Snyk Project in the Web UI and respecting issues that were ignored in a Snyk Project in the Web UI so you can filter them from the analysis results.

This enables Code to be used as a blocking CI/CD gate to test and block builds at the main branch level and then have developers review the results in the Web UI, fix any newly introduced vulnerabilities, or ignore irrelevant ones.

The following commands are supported.

In the terminal, enter:

```
snyk code test --report --project-name="<PROJECT_NAME>"
```

* PROJECT\_NAME must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The Project name must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).
* There is a (temporary) limit of 3MB for the resulting payload, meaning that for large Projects with many issues, the process will not complete.

Running the `snyk code test` command with the `--report` option, as shown, returns the results to the terminal window, along with a URL to the Snyk Code Project where the results have been published. Refer to the following screenshot.

<figure><img src="../../../.gitbook/assets/image (2) (6).png" alt="Snyk code test results with --report option"><figcaption><p>Snyk code test results with --report option</p></figcaption></figure>

If a CLI-based Snyk Code Project does not yet exist for the value provided in the `--project-name` option, Snyk creates a new CLI-based Project. A new snapshot is created under the same Project if a CLI-based Project already exists.

The following command will create or upload an existing Project named \<PROJECT-NAME> under a target named \<TARGET-NAME>

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET-NAME>"
```

The following command will create or upload an existing Project named \<PROJECT-NAME> under a target named \<TARGET-NAME> and grouped by the "$(git branch --show-current)" branch name.

```
snyk code test --report --project-name="<PROJECT-NAME>" --target-name="<TARGET-NAME>" --target-reference="$(git branch --show-current)"
```

<figure><img src="../../../.gitbook/assets/image (4).png" alt=""><figcaption></figcaption></figure>

## **Test and publish CLI results to an existing Snyk Code SCM Project**

{% hint style="info" %}
**Feature availability**\
Publishing and ignoring CLI issues in SCM Projects is in [Closed Beta](../../../more-info/snyk-feature-release-process.md#closed-beta) and only available by invitation from Snyk.

Minimum supported CLI version: 1.1194.0
{% endhint %}

Snyk Code CLI supports publishing the results to an existing Snyk Code SCM Project in the Web UI and respecting issues ignored in a Snyk Project in the Web UI so you can filter them from the analysis results. This will provide [Snyk Code Data flow](https://docs.snyk.io/scan-application-code/snyk-code/exploring-and-working-with-the-snyk-code-results/exploring-the-vulnerability-issues-discovered-by-snyk-code/exploring-the-data-flow-and-fix-analysis-pages-of-an-issue/exploring-the-data-flow-page) for the CLI results snapshot.

### Considerations in publishing SLI results

* An [SCM Integration](../../../integrations/git-repository-scm-integrations/) is required, as the CLI results are appended there. &#x20;
* The Project ID is required.
  * Project URL: https://app.snyk.io/org/org\_name/project/**projectID**
  * Use the GET Projects method through the [API](https://apidocs.snyk.io/?version=2023-05-29#get-/orgs/-org\_id-/projects).&#x20;
* The CLI results are pushed to the SCM Project monitoring your **default branch**. Consider using the command on your default branch.
* In the SCM integration, three snapshots are stored at a time. To save scan results, consider outputting a static [json or SARIF](https://docs.snyk.io/scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file) file for future reference.

### Command to publish to a specific Project

A Project ID can be retrieved from the Snyk WebUI: https://app.snyk.io/org/org\_name/project/PROJECT\_UUID

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```

## &#x20;**Ignoring CLI results**

You can ignore issues from CLI results in the Web UI both for [Publishing CLI results to a Snyk Code Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#publishing-cli-results-to-a-snyk-code-cli-project) and to [Test and publish CLI results to an existing Snyk Code SCM Project](publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md#test-and-publish-cli-results-to-an-existing-snyk-code-scm-project), using the ignore button.

<figure><img src="../../../.gitbook/assets/image (1) (7) (1).png" alt="Ignoring issues in the Web UI"><figcaption><p>Ignoring issues in the Web UI</p></figcaption></figure>

{% hint style="info" %}
&#x20;[snyk-to-html](https://github.com/snyk/snyk-to-html) currently does not honor the ignored issues. It does not show ignored results when using snyk-to-html.
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
