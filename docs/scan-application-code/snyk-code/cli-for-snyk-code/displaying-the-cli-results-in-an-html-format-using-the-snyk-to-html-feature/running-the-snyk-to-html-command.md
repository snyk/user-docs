# Running the snyk-to-html command

You can run the Snyk-to-HTML feature in one command (below), or you can first [test and export the results to a JSON/SARIF file, and then pass the JSON/SARIF file to the Snyk-to-HTML feature](running-the-snyk-to-html-command.md#running-the-snyk-to-html-feature-in-several-steps). When running the snyk-to-html command, you can customize it using its [command options](running-the-snyk-to-html-command.md#snyk-to-html-command-options). After you generated the results in an HTML file, open it in a browser to [view the discovered vulnerabilities with their Data Flow and Fix Analysis information](viewing-the-html-results.md).

**Notes**:

* To display the Help information for the Snyk-to-HTML feature in your terminal, use either the `-h` or `--help` options with the `snyk-to-html` command:

```
snyk-to-html --help
```

* The `snyk-to-html` command currently does not generate the standard exit codes. For an example of a workaround for integrating the `snyk-to-html` command with a CI/CD pipeline, see: [https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/AzurePipelines/AzurePipelines-npm-generic-html.yml](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/AzurePipelines/AzurePipelines-npm-generic-html.yml)

### **Running the Snyk-to-HTML feature in one step**

**Note**: The instructions below use a JSON file, but you can use a SARIF file as well.

**To run the Snyk-to-HTML feature in one step:**

1\. On the terminal, change the directory to the root folder of the repository you want to test.

2\. In order to test the repository, export the results to a JSON file, and convert it to an HTML file called “**results**”. Perform this by entering one of the following commands:

```
snyk code test --json | snyk-to-html > results.html
```

\- or -

```
snyk code test --json | snyk-to-html -o results.html
```

Where:

`-o` – The name of the output file of the generated HTML results.

**Note**: To customize the command, use the [command options](running-the-snyk-to-html-command.md#snyk-to-html-command-options).

A **results.html** file is created in your repository folder, allowing you to [view your test results in an HTML format](viewing-the-html-results.md).

### **Running the Snyk-to-HTML feature in several steps**

You can run the Snyk-to-HTML feature in several steps. First, you need to test the repository and export the test results to a JSON or SARIF file. Then, you should pass the results in the JSON/SARIF file to the Snyk-to-HTML feature, in order to generate the results in an HTML format.

**Note**: The instructions below use a JSON file, but you can also use a SARIF file.

**To run the Snyk-to-HTML feature in several steps:**

1\. On the terminal, change the directory to the root folder of the repository you want to test.

2\. Test your repository with Snyk Code, and export the results to a JSON file called “**results**”:

**Note**: You can change the name of the JSON file, but make sure you use this new file name for the generation of the HTML results, as described in the following step.

```
snyk code test --json > results.json
```

3\. Pass the JSON file to the Snyk-to-HTML feature:

```
snyk-to-html -i results.json -o results.html
```

Where:

`-i` – The input path of the JSON file.

`-o` – The name of the output file of the generated HTML results.

### **snyk-to-html command options**

You can add the following optional commands to the snyk-to-html command:

| **Short** | **Long**   | **Description**                                                                                        | **Default** |
| --------- | ---------- | ------------------------------------------------------------------------------------------------------ | ----------- |
| `-i`      | `--input`  | The input path of the JSON or SARIF file that contains the test results.                               | Stdin       |
| `-o`      | `--output` | <p>The name of the output file of the HTML results.</p><p>Example:<br><code>-o results.html</code></p> | Stdout      |
| `-d`      | `--debug`  | Run the CLI in debug mode.                                                                             |             |
