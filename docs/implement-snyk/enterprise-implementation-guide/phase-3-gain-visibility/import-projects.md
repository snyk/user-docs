# Import Projects

Depending on the integrations you have configured and the language and package managers in your tech stack, you can import Projects into Snyk using the following:&#x20;

* A source control integration with your Git repositories
* The Snyk CLI with CI/CD

The best import route varies based on the languages and package managers in your tech stack.&#x20;

Here are some key points to determine the best starting point. See [Git repositories and CI/CD comparisons](broken-reference) for more details.

## Ways to get started with Snyk

{% hint style="info" %}
See the [Getting started](../../../getting-started/) section and [Start scanning using the CLI, Web UI, or AP](../../../scan-with-snyk/start-scanning.md) for more details.
{% endhint %}

Snyk offers various integration methods to meet your needs, as described here.

### Git integration

You can connect your repositories for automatic scanning. See [Snyk SCM integrations](../../../developer-tools/scms/organization-level-integrations/) for more details.

For a small number of applications, typically under a hundred, follow these steps.

1. Navigate to **Settings** and then to **Integrations**, and connect to Git code repositories using a tile on the Integrations page.
2. In the settings for the integration:
   * Disable the automatic fixes and PR/Merge checks when first onboarding Projects
   * Enable after a steady state is reached and blocking is desired.
3. From the Projects listing, add Projects using the Web UI.
4. Monitor results in Git code repositories.

For hundreds or thousands of repositories, you can use the API endpoint [Import targets ](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)to import your Projects. This takes advantage of an existing source control integration and can be used to automate processes.

The [API-import](../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) tool uses the API to manage onboarding at scale for large enterprises and is the suggested tool to use at scale. You must mirror the source control structure when using the snyk-api-import tool.

### Snyk CLI

The [Snyk CLI ](../../../cli-ide-and-ci-cd-integrations/snyk-cli/)allows granular scanning of individual Projects.&#x20;

{% hint style="info" %}
You must formulate a command for each type of test to perform: open source, code, infrastructure as code, or container tests.
{% endhint %}

Follow these steps to use the CLI:

1. [Install the CLI ](../../../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/)using one of the appropriate methods as part of the build script.
2. In the script, navigate to the Project folder.
3. Run the appropriate `snyk test` or `snyk monitor` commands with the appropriate options for the type of scan being run. \
   Where to implement testing in your scripts is generally flexible, but most commonly, testing is done prior to deployment. Use the `monitor` command alone for Snyk Open Source and Snyk Container to report vulnerabilities passively. In using gating with the `test` command, the idea is to break the build if issues are found that meet particular criteria like `--severity-threshold` or any number of options in the CLI or snyk-filter plugin. \
   \
   In general, the `test` or `monitor` commands or both are typically run for open source after the dependencies are installed on the build system.\
   \
   A typical command might look like one of the following:
   * Code: `snyk code test --org=[org-id]`
   * Open source:&#x20;
     * `snyk test --all-projects --org=[org-id]`
     * `snyk monitor --all-projects --org=[org-id]`
   * Refer to the documentation for [container](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) and [infrastructure as code](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/) scans for information about how to scan those types of Projects.
4. Review results either locally when running `snyk test`, or in the Web UI when using monitor or report functions.

{% hint style="info" %}
Demonstrations of various pipeline integrations can be found on [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)
{% endhint %}

## Snyk API

You can scan using the [Snyk API](../../../snyk-api/overview.md). This enables large-scale automated scanning.

The process is:

1. Navigate to **Settings** and **Service Accounts** and generate an API token.
2. Call the Snyk API in pipelines.
3. Handle the results programmatically.
4. Change the code if needed.

Scanning using the API is useful to accomplish the following:

* Trigger scans through pipelines.
* Scale across your Project portfolio.
* Identify new issues in real-time.



