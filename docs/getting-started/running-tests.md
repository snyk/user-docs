# Running tests

You can use Snyk to test your code in different ways:

* [Run tests manually](running-tests.md#run-tests-manually): by using the Snyk CLI, the Snyk Web UI, and the Snyk API
* [Run tests automatically](running-tests.md#run-tests-automatically): after Project import or using the `snyk monitor` CLI command or using PR Checks to scan new PRs

{% hint style="info" %}
Tests may be limited on your account; see [What counts as a test?](what-counts-as-a-test.md) for more information.
{% endhint %}

## Run tests manually

### Run tests manually with the CLI

You can use the following commands to run tests with the Snyk [CLI](../snyk-cli/cli-commands-and-options-summary.md):

* Scan open-source code with `snyk test`.
* Scan application code with `snyk code test`.
* Scan container images with `snyk container test`.
* Scan Infrastructure as Code (IaC) files with `snyk iac test`.

See [Getting started with the CLI](../snyk-cli/getting-started-with-the-cli.md) for details.

### Run tests manually with the Snyk Web UI

A test is run when you import a Snyk Project (see [Import a Project](quickstart/import-a-project.md)) or click the **Retest now** button on a Project in the Overview tab. During retests, Snyk fetches the assets that need to be scanned for Snyk to find any new issues.

See [Exploring the Snyk Web UI](exploring-the-snyk-web-ui.md) for details.

### Run tests manually with the API

Tests are counted when calls are made to the **https://snyk.io/api/v1/test** endpoint.

See the [API documentation](https://snyk.docs.apiary.io/#reference/test) for details.

## Run tests automatically

### Run tests automatically with the Snyk Web UI

After you [import a Project](quickstart/import-a-project.md), Snyk automatically runs periodic scans on that Project, to see if your code is affected by newly disclosed vulnerabilities.

{% hint style="info" %}
The default test frequency and available test frequencies vary depending on the type of Project: Open Source, Code analysis, Container, or IaC. For more information, see [Usage page details](../snyk-admin/manage-settings/usage-settings.md) (Set test frequency). You can also set test frequency in the Project **Settings** (see [View project settings](../manage-risk/snyk-projects/view-and-edit-project-settings.md)) or use theSnyk REST API: [Updates project by project ID](https://apidocs.snyk.io/?version=2023-02-15#patch-/orgs/-org\_id-/projects/-project\_id-).
{% endhint %}

### Run tests automatically with snyk monitor

Use the `snyk monitor` CLI command to create a snapshot of a project on the Snyk website that will be continuously monitored for new vulnerabilities.

{% hint style="info" %}
Projects are tested at the test frequency you select in your settings; the default is daily. After using `snyk monitor`, you will have recurring tests.  Recurring tests are scheduled retests that run on monitored projects.
{% endhint %}

See [Monitor your projects at regular intervals](../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md) for details.

### Run tests automatically using PR Checks

Snyk can scan every new Pull Request (PR) submitted on your monitored repositories to help prevent new vulnerabilities from being added to your codebase.

See [Run PR Checks](../scan-applications/run-pr-checks/) for details.
