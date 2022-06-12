# Outputting the test results to a JSON or SARIF format in the terminal

You can output the CLI Code test results to a [JSON](outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal.md#outputting-the-test-results-to-a-json-format) or [SARIF](outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal.md#outputting-the-test-results-to-a-json-format-1) format in the terminal, instead of displaying the results in the standard CLI format.

**Notes**:

* You can also [export the test results to a JSON or SARIF format file](exporting-the-test-results-to-a-json-or-sarif-file.md).
* SARIF is an open standard for the output of static analysis tools. For more information on SARIT, see [the official SARIF site](https://sarifweb.azurewebsites.net/).
* The severity levels of the issues discovered in the Snyk Code test are displayed differently in the JSON and SARIF outputs. The severity levels in the JSON and SARIF results are as follows:
  * High = **error**
  * Medium = **warning**
  * Low = **note/info**

For example:

![](<../../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results - in the Terminal.png>)

### Outputting the test results to a JSON format

**To output the test results to a JSON format:**

* In the terminal, enter:

```
snyk code test <path/to/folder> --json
```

The test results appear in the terminal in a JSON format.

For example:

To output the test results of the **snyk-goof-master** folder in a JSON format in the terminal, we enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --json
```

The test results appear in the terminal in a JSON format:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - JSON output in the terminal.png>)

### Outputting the test results to a JSON format

**To output the test results to a SARIF format:**

* In the terminal, enter:

```
snyk code test <path/to/folder> --sarif
```

The test results appear in the terminal in a SARIF format.

For example:

To output the test results of the **snyk-goof-master** folder in a SARIF format, we enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --sarif
```

The test results appear in the terminal in a SARIF format:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - SARIF output in the terminal.png>)

&#x20;
