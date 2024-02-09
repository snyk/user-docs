# Import Projects

Depending on the integrations you have configured, and the language / package managers in your tech stack, you can import Projects into Snyk using:&#x20;

* A source control integration with your Git repositories
* The Snyk CLI with CI/CD.

The best import route varies based on the languages and package managers in your tech stack.&#x20;

Here are some key points to determine the best starting point. For details, see [Git repositories and CI/CD comparisons](../../../integrate-with-snyk/git-repository-and-ci-cd-integrations-comparisons.md).

## Getting started with Snyk

{% hint style="info" %}
For details, see [Getting started](../../../getting-started/) and [Start scanning](../../../scan-with-snyk/start-scanning-using-the-cli-web-ui-or-api.md).
{% endhint %}

Depending on your needs, Snyk offers various integration methods:

### Git Integration

For details, see [Git repositories (SCMs) integrations with Snyk](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

Connect your repositories for automatic scanning.

For a small number of applications, typically under a hundred:

1. From the Snyk Web UI, connect to Git code repositories from the **Settings-Integrations** page.
2. In the integration settings:
   1. Disable the automatic fixes and PR/Merge checks when first onboarding Projects.
   2. Enable them once a steady state is reached and blocking is desired.
3. From the **Projects** page, add the Projects.
4. Monitor results in Git code repositories.

For hundreds or thousands of repositories:

* At scale, Snyk recommends using the API. APIs are available with the Snyk Enterprise plan.
  * Use the [Snyk API](../../../snyk-api/) to import your Projects. This leverages an existing source control integration and can be used to automate processes.
  * The  [snyk-api-import](../../../snyk-api-info/other-tools/tool-snyk-api-import/) tool uses the API to manage onboarding at scale for large enterprises and is the suggested tool to use at scale. The source control structure will need to be mirrored.

## Snyk CLI

For details, see [Snyk CLI](../../../snyk-cli/).

The CLI allows granular scanning of individual Projects.&#x20;

{% hint style="info" %}
A command must be formulated for each type of test to perform (open source, code, infrastructure as code, and container).
{% endhint %}

To use the Snyk CLI:

1. [Install CLI](https://docs.snyk.io/snyk-cli/install-or-update-the-snyk-cli) using one of the appropriate methods as part of the build script.
2. In the script, navigate to the Project folder.
3. Run the appropriate `snyk test` or `snyk monitor` commands and options for the type of scan you want to run. \
   \
   Where to implement testing in your scripts is generally flexible but most commonly prior to deployment. Use the monitor command alone for Snyk Open Source and Snyk Container to passively report. When using gating via the test command, the purpose is to break the build if issues are found that meet particular criteria like `--severity-threshold` or any number of options in the CLI or snyk-filter plugin. \
   \
   In general, Snyk Open source is typically run in test and/or monitor after the dependencies are installed on the build system.\
   \
   A typical command can look as follows:
   * Code: `snyk code test --org=[org-id]`
   * Open source:&#x20;
     * `snyk test --all-projects --org=[org-id]`
     * `snyk monitor --all-projects --org=[org-id]`\
       Replace `[org-id]` with the ID of your Organization.
   * For Container and Infrastructure as Code scans, see [Container](../../../scan-with-snyk/snyk-container/scan-container-images.md) and [Infrastructure as Code](../../../scan-with-snyk/scan-infrastructure/), as this will vary based on the type being scanned.
4. Review results either locally when running `snyk test`, or via the Snyk Web UI when using monitor or report.

For demonstrations of various pipeline integrations, see [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples).

