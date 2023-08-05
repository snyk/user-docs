# Exporting the test results to a JSON or SARIF file

You can export the `snyk code test` results to a JSON or SARIF format file. When you export the results, you must provide a name for the new file.

You can also [output the test results to JSON or SARIF format in the terminal](outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md).

The severity levels of the issues discovered by running `snyk code test` are displayed differently in JSON and SARIF files.  For more information, see [Understanding the severity levels in the JSON and SARIF files](exporting-the-test-results-to-a-json-or-sarif-file.md#understanding-the-severity-levels-in-the-json-and-sarif-files).

## Export methods

There are two methods for exporting the results to a JSON or SARIF file. The instructions that follow show a JSON file, but you can also export a SARIF file.

* Export the results to a new file, without a result display in the terminal:

```
snyk code test --json > <path/to/new_file>
```

* Export the results to a new file, with a standard result display in the terminal:

```
snyk code test --json-file-output=<path/to/new_file>
```

The `snyk code test --json-file-output=<path/to/new_file>` command is available in Snyk CLI v. 1.910.0 and higher. To update your Snyk CLI version, see [Install or update the Snyk CLI](../../../../snyk-cli/install-the-snyk-cli/).

## Exporting the test results to a JSON file

To export the test results to a JSON file, enter the following:

```
snyk code test --json-file-output=<path/to/new_json_file>
```

The test results appear in the terminal in the standard format, and a JSON file is created in the path you specified.

If you want to export the results to a JSON file without displaying the results in the terminal, enter:

```
snyk code test --json > <path/to/new_json_file>
```

For example, to export the test results of the `snyk-goof-master` folder to a JSON file called `json`, change the directory to the root folder of the repository, and enter the following:

```
snyk code test --json-file-output=json
```

In the terminal, the Code test results appear in the standard format:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - 2 .png" alt="snyk code test results in the terminal"><figcaption><p><code>snyk code tes</code>t results in the terminal</p></figcaption></figure>

In the repository folder, a JSON file is created:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - JSON file.png" alt="JSON file in repository"><figcaption><p>JSON file in repository</p></figcaption></figure>

## Exporting the test results to a SARIF file

**T**o export the test results to a SARIF file, enter the following:

```
snyk code test --sarif-file-output=<path/to/new_sarif_file>
```

The test results appear in the terminal in the standard format, and a SARIF file is created in the path you specified.

If you want to export the results to a SARIF file without displaying the results in the terminal, enter:

```
snyk code test --sarif > <path/to/new_sarif_file>
```

For example, to export the test results of the `snyk-goof-master` folder to a SARIF file called `sarif`, change the directory to the root folder of the repository, and  enter the following:

```
snyk code test --sarif-file-output=sarif
```

In the terminal, the test results appear in the standard format:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to SARIF - with terminal results - 2.png" alt="snyk code test results in the terminal"><figcaption><p><code>snyk code tes</code>t results in the terminal</p></figcaption></figure>

In the repository folder, a SARIF file is created:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - export to SARIF - with terminal results - SARIF file.png" alt="SARIF file in repository"><figcaption><p>SARIF file in repository</p></figcaption></figure>

## Understanding the severity levels in the JSON and SARIF files

The severity levels of the issues discovered by running `snyk code test` are displayed differently in JSON and SARIF files. The severity levels in the JSON and SARIF results are as follows:

* High = **error**
* Medium = **warning**
* Low = **note/info**

An example follows:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results.png" alt="Low severity level in JSON or SARIF file"><figcaption><p>Low severity level in JSON or SARIF file</p></figcaption></figure>
