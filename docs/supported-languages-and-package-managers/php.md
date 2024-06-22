# PHP

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Getting started with Snyk for PHP across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md).
2. [Install Snyk CLI and authenticate your machine](../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine).
3. [Set the default Organization for all Snyk tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md) (code analysis).
4. Ensure you have installed the relevant package manager before you begin using the Snyk CLI (open source).
5. Ensure you have included the relevant manifest files supported by Snyk before testing.

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Exporting the test results to a JSON or SARIF file](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)

#### Open source and licensing

To build the dependency tree, Snyk analyzes the `composer.json` and `composer.lock` files to identify the dependencies and their versions.

There are no unique options for use when running Snyk for PHP.

For more information about Snyk CLI see [Getting started with the CLI](../snyk-cli/getting-started-with-the-snyk-cli.md).

### Snyk Web UI (Git repository integration)

PHP Projects can be imported from any of the Git services Snyk supports. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

Once you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import will not process the composer.json manifest.

By default, Snyk scans your production dependencies. From the Snyk Web UI you can configure whether to include your development dependencies (`require_dev`) in the scan for vulnerabilities.

**To update language preferences:**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing Projects when they are re-tested.

#### What's next?

* [Open a Fix PR](php.md#open-a-fix-pr)&#x20;
* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/) and [Snyk API](../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;

