# Supported languages, package managers, and frameworks

**Note:** The language pages are being restructured and updated on June 22 and June 23. The open source and code analysis details have been moved to each language page.

## Code analysis (Snyk Code)

The pages for code analysis support provide details about fully supported frameworks and features for Snyk Code.

Snyk supports Apex, C++, Go, Java and Kotlin, JavaScript, .NET, PHP. Python, Ruby, Swift, TypeScript, and VB NET.

## Open source and licensing (Snyk Open Source)

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

The pages for open-source support provide details about fully supported package managers and features supported for Snyk Open Source.

Snyk supports Bazel, C.C++, Dart and Flutter, Elixir, Go, Java and Kotlin, Javascrpt. .NET, PHP, Python, Ruby, and Swift and Objective-C,

## Snyk CLI

To use the Snyki CLI:

* [Create a Snyk account](../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
* [Install Snyk CLI and authenticate your machine](../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
* [Set the default Organization for `snyk code test`](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md)

### CLI for code analysis

To start testing your code using Snyk Code, open your repository in a terminal and run `snyk code test`.

For information about customizing test options, running other commands, excluding directories and files, and viewing and exploring the results in different formats, see the following:

* [Available commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from Snyk Code CLI tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Output test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Export test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)

After you have run `snyk code test`, you can:

* [Open a Fix PR ](../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/)
* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk CLI for open source

To test your open-source Project for vulnerabilities, run the `snyk test` command.

## IDE and CI/CD

For integrated development environments, see [Snyk IDEs](../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).

If you use continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software.
