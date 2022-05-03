# Outputting the test results to a JSON or SARIF format in the terminal

You can output the CLI Code test results to a JSON or SARIF format in the terminal, instead of displaying the results in the standard CLI format.

**Notes**:

* You can also [export the test results to a JSON or SARIF format file](exporting-the-test-results-to-a-json-or-sarif-file.md).
* SARIF is an open standard for the output of static analysis tools. For more information on SARIT, see [the official SARIF site](https://sarifweb.azurewebsites.net).

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
