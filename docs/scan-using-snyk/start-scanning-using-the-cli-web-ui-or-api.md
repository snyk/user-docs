# Start scanning using the CLI, Web UI, or API

{% hint style="info" %}
Ensure that the code in your repositories is in a supported language and platform. For more information, see [Supported languages and frameworks.](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview)
{% endhint %}

## Prerequisites for scanning applications

Regardless of how you use Snyk, ensure you have completed the steps in the [Quickstart](../getting-started/quickstart/).&#x20;

In addition, if you are using Snyk Code or scanning container images, complete the following:

* [Activate Snyk Code](start-scanning-using-the-cli-web-ui-or-api/scan-code/activate-snyk-code-using-the-web-ui.md) to scan your application code.
* [Set up integration with a supported container registry](../getting-started/quickstart/set-up-an-integration.md) to scan container images.

For general information about Snyk scans, see [Running scans](working-with-snyk-in-your-environment/running-scans.md).&#x20;

## Overview of ways to scan your Projects

Snyk provides multiple ways to scan Snyk Open Source, Snyk Code, and Snyk Container Projects. You can choose to run your scans from the Web UI, the CLI, the API, or with PR Checks.

| Features                                                                           | Snyk Web UI          | Snyk CLI             | Snyk API             | PR Checks            |
| ---------------------------------------------------------------------------------- | -------------------- | -------------------- | -------------------- | -------------------- |
| Auto scanning                                                                      | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: |
| Manual scanning                                                                    | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_minus\_sign: |
| Local scans                                                                        | :heavy\_minus\_sign: | :heavy\_check\_mark: | :heavy\_minus\_sign: | :heavy\_minus\_sign: |
| Incorporate into the CI/CD pipelines                                               | :heavy\_minus\_sign: | :heavy\_check\_mark: | :heavy\_minus\_sign: | :heavy\_minus\_sign: |
| Obtain results precisely reflecting the Project vulnerabilities and configurations | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: | :heavy\_check\_mark: |

If you choose to use the CLI for scanning your Snyk Open Source, Snyk Code, and Snyk Container Projects, you can run the scans locally on your machine or incorporate them into your CI/CD pipelines, providing more control and flexibility over the scanning process.&#x20;

In addition, using the CLI enables you to scan your code, open-source packages, and container images in their specific development environments, ensuring results that precisely reflect the dependencies and configurations of your Projects.

## Scan using the CLI

Before initiating a scan using the CLI, ensure you follow all the installation, authentication, and getting started steps:

* Follow the instructions on the [Install or update the Snyk CLI](../snyk-cli/install-or-update-the-snyk-cli/) page and choose the installation method that best suits your needs.&#x20;
* After the CLI is installed, you must authenticate your machine by following the instructions on the [Authentication](../snyk-cli/authenticate-the-cli-with-your-account.md) page.
* You can test your installation by running `snyk --help` in the CLI.&#x20;
* After the CLI setup in complete, you can start scanning. See [Getting started with the CLI ](../snyk-cli/getting-started-with-the-snyk-cli.md)for more details.

### Initiate a scan

When you scan your Projects using the CLI, you can use a variety of commands and options, some that apply only to Snyk Open Source, Snyk Code, Snyk Container, or Snyk IaC, and others that apply to scanning regardless of the method. For a complete list of Snyk CLI commands and options, see the [CLI commands and options summary](../snyk-cli/cli-commands-and-options-summary.md).

Each Snyk product has specific commands and options. Ensure you are using the correct options and commands for your needs. For information on using the product-specific CLI commands, see the CLI help:

* [Code](../snyk-cli/commands/code.md)
* [Container](../snyk-cli/commands/container.md)
* [IaC](../snyk-cli/commands/iac.md)
* [Test ](../snyk-cli/commands/test.md)for scanning with Snyk Open Source
* [Monito](../snyk-cli/commands/monitor.md)r for monitoring to detect new issues with Snyk Open Source

For more information, see the following pages:

* [Use Snyk Open Source from the CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-open-source-from-the-cli/)
* [Use Snyk Code from the CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/)
* [Use Snyk Container from the CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-container-from-the-cli/)
* [Snyk CLI for IaC](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/)

### Scan results

After you finish running a Snyk CLI scan, the results are displayed in the terminal. The following pages explain the results for Snyk Open Source, Snyk Code, and Snyk Container scans:

* [Review the Snyk Open Source CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-open-source-from-the-cli/review-the-snyk-open-source-cli-results.md)
* [Working with the Snyk Code CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/working-with-the-snyk-code-cli-results.md)
* [Understand Snyk Container CLI results](../snyk-cli/scan-and-maintain-projects-using-the-cli/use-snyk-container-from-the-cli/understanding-snyk-container-cli-results.md)

## Scan using the Web UI

Using a Snyk integration, you can scan your Projects from the Web UI> Running PR Checks also scans your code. See [Configure PR Checks](run-pr-checks/configure-pr-checks.md) to learn how to enable this functionality. See [Run an analysis with Visual Studio Code extension](../integrations/ide-tools/visual-studio-code-extension/run-an-analysis-with-visual-studio-code-extension.md) for an example of scanning using an integration.. For more information, see [Use Snyk in your IDE](../integrations/ide-tools/) and [Snyk CI/CD integrations](../integrations/snyk-ci-cd-integrations/).

## Scan using the API

The Snyk API v1 offers a set of endpoints to test your code. See the [API v1 Test docs](https://snyk.docs.apiary.io/#reference/test) for more information.







