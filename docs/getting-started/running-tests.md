# Running tests

You can use Snyk to test your code in different ways:

* [Run tests manually](running-tests.md#run-tests-manually): using the Snyk CLI, the Snyk Web UI, and the Snyk API.
* [Run tests automatically](running-tests.md#run-tests-automatically): on a recurring basis, or when a relevant repo change is made.

{% hint style="info" %}
Tests may be limited on your account; see [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-) for more information.
{% endhint %}

## Run tests manually

### Run tests with the CLI

With the Snyk [CLI](../snyk-cli/cli-reference.md) you can use the following commands:

* Scan open-source code with `snyk test`.
* Scan application code with `snyk code test`.
* Scan container images with `snyk container test`.
* Scan Infrastructure as Code (IaC) files with `snyk iac test`.

See [Getting started with the CLI](../snyk-cli/getting-started-with-the-cli.md) for details.

### Run tests with the Snyk Web UI

A test is run when you add a new project, or click the **Retest now** button on a Project in the Overview tab.

See [Snyk Web UI](../snyk-web-ui/) for details.

### Run tests with the API

Tests are counted when calls are made to the [**https://snyk.io/api/v1/test**](https://snyk.io/api/v1/test) endpoint.

See [API documentation](https://snyk.docs.apiary.io) for details.

## Run tests automatically

Snyk provides automatic scanning functions with recurring tests, allowing you to catch new vulnerabilities automatically. After you import a project, Snyk automatically runs periodic checks to see if your code is affected by newly disclosed vulnerabilities.

{% hint style="info" %}
Test frequency is set to daily by default. To change frequency, go to either the **Usage** page (see [Usage page details](https://docs.snyk.io/user-and-group-management/managing-settings/usage-page-details)) or the project **Settings** page (see [View project settings](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/view-project-settings)).
{% endhint %}

### Snyk monitor

Use the `snyk monitor` CLI command to create a snapshot of a project on the Snyk website that will be continuously monitored for new vulnerabilities.

See [Monitor your projects at regular intervals](https://docs.snyk.io/snyk-cli/secure-your-projects-in-the-long-term/monitor-your-projects-at-regular-intervals) for details.

### PR Checks

Snyk can be enabled to scan every Pull Request (PR) submitted on your monitored repositories for vulnerabilities, to help prevent new vulnerabilities being added to your codebase.
