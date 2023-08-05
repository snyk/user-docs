# Understanding the Snyk Code CLI results

After you run the `snyk code test` command in the CLI, the results of the test are displayed:

<figure><img src="../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results Details - 2.png" alt="Snyk Code test restuls from the CLI"><figcaption><p>Snyk Code test restuls from the CLI</p></figcaption></figure>

Note that if you ignored issues on the Snyk Web UI, these issues would still appear in the CLI results. Each section on this page explains one section of the results displayed.

## List of vulnerability issues detected by Snyk Code

The list of issues discovered in the Snyk Code test is organized by the severity level of the issues, from low to high.

For each detected issue, the following information is provided:

<figure><img src="../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Issue summary - 2.png" alt="CLI test information for each Snyk Code issue"><figcaption><p>CLI test information for each Snyk Code issue</p></figcaption></figure>

* Header: The severity level and vulnerability type of the issue.
* Path: The file name and the line in the file where the issue was found. These location details refer to the Sink of the issue, meaning where the vulnerability may be executed in the tested repository.
* Info: A description of the data flow of the issue.

The message that appears in the `Info` section is the same as the one in the **Data flow** section on the Web UI:

<figure><img src="../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Issue summary - In the UI - 2.png" alt="CLI test Info for Snyk Code issue in the Data flow section"><figcaption><p>CLI test Info for Snyk Code issue in the Data flow section</p></figcaption></figure>

## General information about the test results

The general information about the test results includes the following details:

<figure><img src="../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Test summary - 2.png" alt="Details in general information about the CLI test results"><figcaption><p>Detals in general information about the CLI test results</p></figcaption></figure>

* Test success: Whether the test was completed or not.
* Organization: The Snyk ID or internal name of the Organization under which the test run. For more information, see [Set the Snyk Organization for the CLI tests](set-the-snyk-organization-for-the-cli-tests/).
* Test type: The type of test command that generated the results. For Snyk Code, it is always `Static code analysis`**.**
* Project path: The path of the tested repository.

## Summary of the test findings

The summary of the test findings includes the following details:

<figure><img src="../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Summary - 2.png" alt="Summary ot CLI test findings for Snyk Code issues"><figcaption><p>Summary ot CLI test findings for Snyk Code issues</p></figcaption></figure>

* The number of vulnerability issues that Snyk Code discovered in the tested repository.
* The number of discovered issues at each severity level.

{% hint style="info" %}
The `snyk code test` command has exit codes. See the help for [definitions of these codes](../../../snyk-cli/commands/code-test.md#exit-codes). To see the exit code, run `snyk code test -d`.

For a summary of exit codes for all CLI commands, see the [CLI commands and options summary](../../../snyk-cli/cli-reference.md).
{% endhint %}
