# Exporting the test results to a JSON or SARIF file

You can export the CLI Code results to a JSON or SARIF format file. When you export the results, you need to provide a name for the new file.

**Note**: You can also [output the test results to a JSON or SARIF format in the terminal](outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal.md).

There are two methods to perform the export:

**Note**: The instructions below use a JSON file, but you can use a SARIF file as well.

* Export the results to a new file, without a result display in the terminal:

```
snyk code test --json > <path/to/new_file>
```

* Export the results to a new file, with a standard result display in the terminal:

```
snyk code test --json-file-output=<path/to/new_file>
```

**Note**: The `snyk code test --json-file-output=<path/to/new_file>` command is available from Snyk CLI v. 1.910.0 and higher. To update your Snyk CLI version, enter:

```
npm update snyk -g
```

&#x20;

&#x20;**To export the test results to a JSON file:**

* In the terminal, enter:

```
snyk code test --json-file-output=<path/to/new_json_file>
```

The test results appear in the terminal in the standard format, and a JSON file is created in the path you specified.

**Note**: If you want to export the results to a JSON file WITOUT displaying the results in the terminal, enter:

```
snyk code test --json > <path/to/new_json_file>
```

For example:

To export the test results of the **snyk-goof-master** folder to a JSON file called “**json**”, we changed the directory to the root folder of the repository, and then we enter:

```
snyk code test --json-file-output=json
```

In the terminal, the Code test results appear in the standard format:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results .png>)

In the repository folder, a JSON file is created:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - JSON file.png>)



**To export the test results to a SARIF file:**

* In the terminal, enter:

```
snyk code test --sarif-file-output=<path/to/new_sarif_file>
```

The test results appear in the terminal in the standard format, and a SARIF file is created in the path you specified.

**Note**: If you want to export the results to a SARIF file WITOUT displaying the results in the terminal, enter:

```
snyk code test --sarif > <path/to/new_sarif_file>
```

For example:

To export the test results of the **snyk-goof-master** folder to a SARIF file called “**sarif**”, we changed the directory to the root folder of the repository, and then we enter:

```
snyk code test --sarif-file-output=sarif
```

In the terminal, the test results appear in the standard format:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - export to SARIF - with terminal results.png>)

In the repository folder, a SARIF file is created:

![](<../../../../.gitbook/assets/snyk Code - CLI - results - export to SARIF - with terminal results - SARIF file.png>)

&#x20;
