# View Snyk Code CLI results

The Snyk CLI enables you to perform the following actions on the results of the `snyk code test` command:

* [Analyze Snyk Code CLI results](view-snyk-code-cli-results.md#analyze-snyk-code-cli-results): View test results and analyze vulnerabilities.
* [Filter results by severity level](view-snyk-code-cli-results.md#filter-results-by-severity-level): Filter the `snyk code test` results shown in the terminal to display only issues with a specific severity level and higher.
* [Output test results](view-snyk-code-cli-results.md#output-test-results): Output the `snyk code test` results to a JSON or SARIF format in the terminal instead of displaying the results in the standard CLI format.
* [Export test results](view-snyk-code-cli-results.md#export-test-results): Export the CLI Code results to a JSON or SARIF format file.

{% hint style="info" %}
For `snyk code test`, the JSON and SARIF formats are the same; thus, examples are shown in only one format.

You can also [display the CLI results in HTML format using `snyk-to-html`](../cli-tools/snyk-to-html.md).
{% endhint %}

## Analyze Snyk Code CLI results

After you run the `snyk code test` command in the CLI, the results of the test are displayed:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results Details - 2.png" alt="Snyk Code test restuls from the CLI"><figcaption><p>Snyk Code test restuls from the CLI</p></figcaption></figure>

Note that if you ignored issues on the Snyk Web UI, these issues would still appear in the CLI results. Each section on this page explains one section of the displayed results.

### List of vulnerability issues detected by Snyk Code

The list of issues discovered in the Snyk Code test is organized by the severity level of the issues, from low to high.

For each detected issue, the following information is provided:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Issue summary - 2.png" alt="CLI test information for each Snyk Code issue"><figcaption><p>CLI test information for each Snyk Code issue</p></figcaption></figure>

* Header: The severity level and vulnerability type of the issue.
* Path: The file name and the line in the file where the issue was found. These location details refer to the sink of the issue, meaning where the vulnerability may be executed in the tested repository.
* Info: A description of the data flow of the issue.

The message that appears in the `Info` section is the same as the one in the **Data flow** section on the Web UI:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Issue summary - In the UI - 2.png" alt="CLI test Info for Snyk Code issue in the Data flow section"><figcaption><p>CLI test Info for Snyk Code issue in the Data flow section</p></figcaption></figure>

### General information about the test results

The general information about the test results includes the following details:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Test summary - 2.png" alt="Details in general information about the CLI test results"><figcaption><p>Detals in general information about the CLI test results</p></figcaption></figure>

* Test success: Whether the test was completed or not.
* Organization: The Snyk ID or internal name of the Organization under which the test run. For more information, see [Set the Snyk Organization for the CLI tests](set-the-snyk-organization-for-the-cli-tests.md).
* Test type: The type of test command that generated the results. For Snyk Code, it is always `Static code analysis`**.**
* Project path: The path of the tested repository.

### Summary of the test findings

The summary of the test findings includes the following details:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Summary - 2.png" alt="Summary ot CLI test findings for Snyk Code issues"><figcaption><p>Summary ot CLI test findings for Snyk Code issues</p></figcaption></figure>

* The number of vulnerability issues that Snyk Code discovered in the tested repository.
* The number of discovered issues at each severity level.

{% hint style="info" %}
The `snyk code test` command has exit codes. See the help for [definitions of these codes](../../commands/code-test.md#exit-codes). To see the exit code, run `snyk code test -d`.

For a summary of exit codes for all CLI commands, see the [CLI commands and options summary](../../cli-commands-and-options-summary.md).
{% endhint %}

## Filter results by severity level

You can filter the test results that are shown in the CLI terminal and display only issues with a specific severity level and higher.

To display only issues above a specific severity level, enter the following:

```
snyk code test <path/to/folder> --severity-threshold=<low|medium|high>
```

The results will include only issues with the specified severity level and issues with a higher severity level.

For example, in the `snyk-goof-master` folder, eight issues were found, four with a High severity level and four with Medium:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - before - 2.png" alt="CLI test results for Snyk Code with High and Medium severity"><figcaption><p>CLI test results for Snyk Code with High and Medium severity</p></figcaption></figure>

To display only issues with a High severity level and above, enter the following:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --severity-threshold=high
```

The results show four issues, all with a High severity level. Issues with a lower severity level are not displayed:

![CLI test results for Snyk Code with High severity](<../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - after - 2.png>)

## Severity levels in JSON and SARIF files

The severity levels of the issues discovered by running `snyk code test` are displayed differently in JSON and SARIF files. The severity levels in the JSON and SARIF results are as follows:

* High = **error**
* Medium = **warning**
* Low = **note/info**

The designation Critical is not used in Snyk Code.

An example of medium-level severity displayed in the terminal follows:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results - in the Terminal.png" alt="Medium severity in JSON or SARIF output"><figcaption><p>Medium severity in JSON or SARIF output</p></figcaption></figure>

The following shows examples of high-level and low-level severity in a file:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results.png" alt="High and low severity levels in JSON or SARIF file"><figcaption><p>High and low severity levels in JSON or SARIF file</p></figcaption></figure>

## Output test results

You can output the `snyk code test` results to JSON or SARIF format in the terminal instead of displaying the results in the Snyk CLI format.

You can also [export the test results to a JSON or SARIF format file](view-snyk-code-cli-results.md#export-test-results). SARIF is an open standard for the output of static analysis tools. For more information, see the [SARIF site](https://sarifweb.azurewebsites.net/).&#x20;

The severity levels of the issues discovered by running `snyk code test` and reported in JSON and SARIF files are displayed differently from the results in the terminal.  For more information, see [Severity levels in the JSON and SARIF files](view-snyk-code-cli-results.md#severity-levels-in-json-and-sarif-files).

To output the test results to JSON format, enter the following:

```
snyk code test <path/to/folder> --json
```

To output the test results to SARIF format, enter the following:

```
snyk code test <path/to/folder> --sarif
```

The test results appear in the terminal in JSON or SARIF format.

Because JSON and SARIF are the same for `snyk code test`, only a JSON example is shown here. The example shows how to output the test results of the `snyk-goof-master` folder in JSON format in the terminal by using the following command:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --json
```

The test results appear in the terminal in JSON format:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - JSON output in the terminal.png" alt="snyk code test results in JSON format"><figcaption><p><code>snyk code test</code> results in JSON format</p></figcaption></figure>

## Export test results

You can export the `snyk code test` results to a JSON or SARIF format file. When you export the results, you must provide a name for the new file.

You can also [output the test results to JSON or SARIF format in the terminal](view-snyk-code-cli-results.md#output-test-results).

The severity levels of the issues discovered by running `snyk code test` and reported in JSON and SARIF files are displayed differently from the results in the terminal.  For more information, see [Severity levels in the JSON and SARIF files](view-snyk-code-cli-results.md#severity-levels-in-json-and-sarif-files).

You can use two methods to export the results to either a JSON or SARIF file. The following instructions show a JSON file, but you can also export a SARIF file.

### Export test results to a new file with a standard display of results in the terminal

The `snyk code test --json-file-output=<path/to/new_file>` command is available in the Snyk CLI v. 1.910.0 and higher. To update your Snyk CLI version, see [Install or update the Snyk CLI](../../install-or-update-the-snyk-cli/).

To export the results to a new JSON file, with a standard  display of results in the terminal, use the following command:

```
snyk code test --json-file-output=<path/to/new_json_file>
```

To export the test results to a new SARIF file, use the following command:

```
snyk code test --sarif-file-output=<path/to/new_sarif_file>
```

The test results appear in the terminal in the standard format, and a JSON or SARIF file is created in the path you specified.

Because JSON and SARIF are the same for snyk code test, only a JSON example is shown here. To export the test results of the `snyk-goof-master` folder to a JSON file called `json`, change the directory to the root folder of the repository, and enter the following:

```
snyk code test --json-file-output=json
```

In the terminal, the Code test results appear in the standard format:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - 2 .png" alt="snyk code test results in the terminal"><figcaption><p><code>snyk code tes</code>t results in the terminal</p></figcaption></figure>

In the repository folder, a JSON file is created:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - JSON file.png" alt="JSON file in repository"><figcaption><p>JSON file in repository</p></figcaption></figure>

### Export the results to a new file without a display of results in the terminal

To export the results to a new JSON file without displaying the results in the terminal, use the following command:

```
snyk code test --json > <path/to/new_json_file>
```

To export the results to a SARIF file without displaying the results in the terminal, use the following command:

```
snyk code test --sarif > <path/to/new_sarif_file>
```

