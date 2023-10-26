# Import Projects

Depending on the integrations you have configured, and the language / package managers in your tech stack, you can import Projects into Snyk using:&#x20;

* A source control integration with your Git repositories.
* The Snyk CLI with CI/CD.

The best import route varies based on the languages and package managers in your tech stack.&#x20;

Here are some key points to determine the best starting point. See [this article](../../../integrations/git-repository-and-ci-cd-integrations-comparisons.md) for more details.

## Getting started with Snyk

{% hint style="info" %}
See the [Getting started](../../../getting-started/) and [Start scanning](../../../scan-applications/start-scanning/) sections for more details.
{% endhint %}

Snyk offers various integration methods to meet your needs:

### [Git Integration](../../../integrations/git-repository-scm-integrations/)

Connect your repositories for automatic scanning.

**For a small number of applications, typically under a hundred:**

1. Connect to Git code repositories from the Settings->Integrations screen from within Snyk
2. Within the settings for the integration:
   1. Disable the automatic fixes and PR/Merge checks when first onboarding projects
   2. Enable once a steady state is reached and blocking is desired.
3. From Projects, add the projects via UI
4. Monitor results in git code repositories

**For hundreds or thousands of repositories**

* Alternatively, use the [Snyk API](../../../snyk-api/) to import your Projects. This leverages an existing source control integration and can be used to automate processes.
* The  [snyk-api-import](../../../snyk-api-info/other-tools/tool-snyk-api-import/) tool uses the API to manage onboarding at scale for large enteprises and is the suggested tool to use at scale. The source control structure will need to be mirrored.

## [Snyk CLI](../../../snyk-cli/)

The CLI allows granular scanning of individual Projects.&#x20;

{% hint style="info" %}
A command must be formulated for each type of test to perform (open source, code, infrastructure as code, and container).
{% endhint %}

How to use:

1. [Install CLI](https://docs.snyk.io/snyk-cli/install-or-update-the-snyk-cli) using one of the appropriate methods as part of the build script
2. In script navigate to the project folder
3. Run the appropriate **snyk test** or **snyk monitor** commands and options for the type of scan being run. \
   Where to implement testing in your scripts is generally pretty flexible, but most commonly prior to deployment. Use the monitor command alone for Snyk Open Source and Snyk Container to passively report. When using gating via the test command, the idea is to break the build if issues are found that meet a particulat criteria like --severity-threshold or any number of options in the CLI or snyk-filter plugin. \
   \
   In general, Snyk Open source is typically ran in test and/or monitor after the dependencies are installed on the build system.\
   \
   A typical command might looks like:
   * code: **snyk code test --org=\[org-id]**
   * open source:&#x20;
     * **snyk test --all-projects --org=\[org-id]**
     * **snyk monitor --all-projects --org=\[org-id]**
   * For [container](https://docs.snyk.io/scan-applications/snyk-container/use-snyk-container-from-the-cli) and [infrastructure as code](https://docs.snyk.io/scan-infrastructure/snyk-cli-for-infrastructure-as-code) scans, please refer to the relevant documentation as this will vary based on type being scanned
4. Review results either locally when running snyk test, or via the web interface when using monitor or report.

{% hint style="info" %}
Demonstrations of various pipeline integrations can be found on [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)
{% endhint %}

## [Snyk API](../../../snyk-api/) (Advanced use case)

1. Generate API token under Settings->Service Accounts
2. Call API in pipelines
3. Handle results programmatically
4. May require code changes

**How to use:**

* Trigger scans through pipelines
* Scale across project portfolio
* Identify new issues in real-time

**Advantages:**

The API enables large-scale automated scanning.

