# Running scans

You can use Snyk to scan your code:

* [Manually](running-scans.md#run-tests-manually): using the Snyk CLI, the Snyk Web UI, and the Snyk API.
* [Automatically](running-scans.md#run-tests-automatically): after Project import or using the `snyk monitor` CLI command or using PR Checks to scan new PRs.

To start using Snyk scanning capabilities for open-source libraries, container images, and application code, see [Start scanning](../start-scanning-using-the-cli-web-ui-or-api.md).

{% hint style="info" %}
Scans may be limited on your account, depending on your[ Pricing Plan](../../more-info/plans.md). See [What counts as a scan?](what-counts-as-a-test.md) for more information.
{% endhint %}

## Scan manually

### Using the CLI

You can use the following Snyk [CLI](../../snyk-cli/cli-commands-and-options-summary.md) commands:

* Scan open-source code with `snyk test`.
* Scan application code with `snyk code test`.
* Scan container images with `snyk container test`.
* Scan Infrastructure as Code (IaC) files with `snyk iac test`.

See [Getting started with the CLI](../../snyk-cli/getting-started-with-the-snyk-cli.md) for details.

### Using the Snyk Web UI

A scan runs when you import a Snyk Project (see [Import a Project](../../getting-started/quickstart/import-a-project.md)) or click the **Retest now** button on a Project.

See [Exploring the Snyk Web UI](../../getting-started/explore-snyk-through-the-web-ui.md) for details.

### Using the API

Scans are counted when calls are made to the **https://snyk.io/api/v1/test** endpoint.

See the [API documentation](https://snyk.docs.apiary.io/#reference/test) for details.

## Scan automatically

### Using the CLI (snyk monitor)

Use the `snyk monitor` CLI command to create a snapshot of a project on the Snyk website that will be continuously monitored for new vulnerabilities.

{% hint style="info" %}
Projects are scanned at the frequency you select in your settings; the default is daily. After using `snyk monitor`, you will have recurring scans running on monitored Projects.
{% endhint %}

See [Monitor your projects at regular intervals](../../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md) for details.

### Using the Snyk Web UI

After you [import a Project](../../getting-started/quickstart/import-a-project.md), Snyk automatically runs periodic scans on that Project, to see if your code is affected by newly disclosed vulnerabilities.

{% hint style="info" %}
The default scan frequency and available frequencies vary depending on the type of Project: Open Source, Code analysis, Container, or IaC. For more information, see [Usage page details](../../snyk-admin/manage-settings/usage-settings.md). You can also set frequency in the Project **Settings** (see [View project settings](../../snyk-admin/introduction-to-snyk-projects/view-and-edit-project-settings.md)) or use the Snyk REST API: see [Updates project by project ID](https://apidocs.snyk.io/?version=2023-02-15#patch-/orgs/-org\_id-/projects/-project\_id-).
{% endhint %}

### Using PR Checks

Snyk can scan every new Pull Request (PR) submitted on your monitored repositories to help prevent new vulnerabilities from being added to your codebase.

See [Run PR Checks](../run-pr-checks/) for details.
