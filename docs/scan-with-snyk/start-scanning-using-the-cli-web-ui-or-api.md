# Start scanning using the CLI, Web UI, or API

You can use Snyk to scan your code manually and automatically, using the [Snyk CLI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-cli), the [Snyk Web UI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-web-ui), the [Snyk API](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-api), and by running [PR Checks](start-scanning-using-the-cli-web-ui-or-api.md#using-pr-checks).

{% hint style="info" %}
Scans may be limited on your account, depending on your[ Pricing Plan](../implement-snyk/enterprise-implementation-guide/trial-limitations.md). See [What counts as a scan?](scanning-overview/what-counts-as-a-test.md) for more information.
{% endhint %}

## Prerequisites

1. Ensure that the code in your repositories is in a supported language and platform. For more information, see [Supported languages and frameworks.](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview)
2. Ensure you have completed the steps in the [Quickstart](../getting-started/quickstart/).&#x20;
3. If you are using Snyk Code, [Enable Snyk Code](snyk-code/configure-snyk-code.md#enable-snyk-code-in-snyk-web-ui).
4. If you are scanning container images, [Set up integration with a supported container registry](../getting-started/quickstart/set-up-an-integration.md).&#x20;

## Overview of ways to scan your Projects

You can run your scans from the [Web UI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-web-ui), the [CLI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-cli), the [API](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-api), or with [PR Checks](run-pr-checks/).

| Features                                                                           | Snyk Web UI          | Snyk CLI             | Snyk API             | PR Checks            |
| ---------------------------------------------------------------------------------- | -------------------- | -------------------- | -------------------- | -------------------- |
| Auto scanning                                                                      | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: |
| Manual scanning                                                                    | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_minus\_sign: |
| Local scans                                                                        | :heavy\_minus\_sign: | :heavy\_check\_mark: | :heavy\_minus\_sign: | :heavy\_minus\_sign: |
| Incorporate into the CI/CD pipelines                                               | :heavy\_minus\_sign: | :heavy\_check\_mark: | :heavy\_minus\_sign: | :heavy\_minus\_sign: |
| Obtain results precisely reflecting the Project vulnerabilities and configurations | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: |

## Scan using the CLI

{% hint style="info" %}
See [Getting started with the CLI](../snyk-cli/getting-started-with-the-snyk-cli.md) for more details.
{% endhint %}

### Overview

If you use the [Snyk CLI](../snyk-cli/) for scanning, you can run scans locally, or you can incorporate them into your CI/CD pipelines, providing more control and flexibility over the scanning process.&#x20;

In addition, using the CLI enables you to scan your code, open-source packages, and container images in their development environments, ensuring results that precisely reflect the dependencies and configurations of your Projects.

### CLI scanning prerequisites

Before initiating a scan using the CLI, ensure you follow all the installation, authentication, and getting started steps:

* Follow the instructions on the [Install or update the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/) page and choose the installation method that best suits your needs.&#x20;
* After the CLI is installed, you must authenticate your machine by following the instructions on the [Authentication](../snyk-cli/authenticate-the-cli-with-your-account.md) page.
* You can test your installation by running `snyk --help` in the CLI.&#x20;
* After the CLI setup in complete, you can start scanning. See [Getting started with the CLI ](../snyk-cli/getting-started-with-the-snyk-cli.md)for more details.

### Run a CLI scan

Each Snyk product has specific commands and options; some apply to specific scanning methods such as Snyk Code, and some apply across all methods. For a complete list of Snyk CLI commands and options, see the [CLI commands and options summary](../snyk-cli/cli-commands-and-options-summary.md).

You can use the following Snyk [CLI](../snyk-cli/cli-commands-and-options-summary.md) commands for specific scanning methods:

<table><thead><tr><th width="190">Command</th><th width="236">Function</th><th>More details</th></tr></thead><tbody><tr><td><a href="../snyk-cli/commands/test.md">snyk test</a></td><td>Scan open-source code</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/">Use Snyk Open Source from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/code.md">snyk code test</a></td><td>Scan application code</td><td> <a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/">Use Snyk Code from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/container.md">snyk container test</a></td><td>Scan container images</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/">Use Snyk Container from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/iac.md">snyk iac test</a></td><td>Scan infrastructure as code (IaC) files</td><td> <a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">Snyk CLI for IaC</a></td></tr></tbody></table>

### View CLI scan results

After you run a Snyk CLI scan, the results are displayed in the terminal. The following pages explain the results for Snyk Open Source, Snyk Code, and Snyk Container scans:

* [Review the Snyk Open Source CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/review-the-snyk-open-source-cli-results.md)
* [Working with the Snyk Code CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md)
* [Understand Snyk Container CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/understand-snyk-container-cli-results.md)

### Scan continuously using the CLI (snyk monitor)

Use the [snyk monitor](../snyk-cli/commands/monitor.md) CLI command to create a snapshot of a Project on the Snyk website that will be continuously monitored for new vulnerabilities.

{% hint style="info" %}
Projects are scanned at the frequency you select in your settings; the default is daily. After using `snyk monitor`, you will have recurring scans running on monitored Projects.
{% endhint %}

See [Monitor your projects at regular intervals](../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md) for detail

## Scan using the Web UI

A scan runs when you import a Snyk Project (see [Import a Project](../getting-started/quickstart/import-a-project.md)) or click the **Retest now** button on a Project. Snyk then automatically runs periodic scans on that imported Project, to see if your code is affected by newly disclosed vulnerabilities.

See [Explore Snyk through the Web UI](../getting-started/explore-snyk-through-the-web-ui.md).

{% hint style="info" %}
The default scan frequency and available frequencies vary depending on the type of Project: for more information, see [Usage page details](../snyk-admin/manage-settings/usage-settings.md). You can also set the frequency in the Project **Settings** (see [View project settings](../snyk-admin/snyk-projects/view-and-edit-project-settings.md)) or use the Snyk REST API: see [Updates project by project ID](https://apidocs.snyk.io/?version=2023-02-15#patch-/orgs/-org\_id-/projects/-project\_id-).
{% endhint %}

## Scan using the API

The Snyk API v1 offers a set of endpoints to test your code. Scans are counted when calls are made to the **https://snyk.io/api/v1/test** endpoint.

See the [API documentation](https://snyk.docs.apiary.io/#reference/test) and the [API v1 Test docs](https://snyk.docs.apiary.io/#reference/test) for more information.

## Using PR Checks

Snyk can scan every new Pull Request (PR) submitted on your monitored repositories to help prevent new vulnerabilities from being added to your codebase.

See [Run PR Checks](run-pr-checks/) for details.
