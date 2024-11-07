# Language support and CLI, CI/CD, and SCM integrations

## Steps to start using the CLI

You must have a Snyk account to use the Snyk CLI.

* [Install Snyk CLI and authenticate your machine](../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
* [Set the default Organization for the `snyk test` or `snyk code test` commands ](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md)

## CLI for Snyk Code

To start testing your code using Snyk Code through the CLI, open your repository in a terminal and run `snyk code test`.

For information about customizing test options, running other commands, excluding directories and files, and viewing and exploring the results in different formats, see the following:

* [Available commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from Snyk Code CLI tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Output test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Export test results](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)
* [snyk-to-html](../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)

After you have run `snyk code test`, you can:

* [Open a Fix PR ](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/)
* [Configure PR Checks](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

## CLI for Snyk Open Source

Ensure you have installed the relevant package manager and you have included the relevant manifest files supported by Snyk before testing.

To test your Open Source Project for vulnerabilities, run the `snyk test` command.

## Steps to start using SCM integrations

* [Set up an integration](../getting-started/#set-up-a-snyk-integration).
* For details, see [Snyk SCM integrations](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/).
* For language-specific information, see [Git repositories with Maven and Gradle](java-and-kotlin/git-repositories-with-maven-and-gradle.md), [Git repositories and JavaScript](javascript/git-repositories-and-javascript.md), and [Git repositories and Python](python/git-repositories-and-python.md).
