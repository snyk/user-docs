# Apex

{% hint style="info" %}
Snyk for Apex is supported only for Snyk Code.
{% endhint %}

## Code analysis

Interfile is currently not supported. The data flow is monitored within a single file, not between multiple files.

`.trigger` files are supported.

## Getting started with Snyk for Apex across environments

### Snyk CLI&#x20;

#### Prerequisites

* [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/)
* [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
* [Set the default Organization for all Snyk tests](../snyk-code/using-snyk-code-from-the-cli/set-the-snyk-organization-for-the-cli-tests/setting-the-default-organization-for-all-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../snyk-code/using-snyk-code-from-the-cli/excluding-directories-and-files-from-the-snyk-code-cli-test.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../snyk-code/using-snyk-code-from-the-cli/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md)
* [Exporting the test results to a JSON or SARIF file](../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md)

#### What's next?

* [Open a Fix PR](apex.md#open-a-fix-pr)&#x20;
* [Configure PR Checks](../../scan-application-code/run-pr-checks/configure-pr-checks.md)

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../integrations/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
