# Run the snyk-to-html command

You can run `snyk-to-html` as part of a `snyk test` command to create a readable build artifact as part of the output of a scan.

You can also run a `snyk test` command with the option to export the results to a JSON file and then convert the JSON  file to HTML using by using  s`nyk-to-html`. You can also export Snyk Code results to a SARIF file and convert that file to HTML.

When you run the `snyk-to-html` command, you can customize it with the [command options](run-the-snyk-to-html-command.md#snyk-to-html-command-options).&#x20;

After you generate the results in an HTML file, open the file in a browser to view a full report of the scan results. For example, for Snyk Code, the results include the discovered vulnerabilities with their data flow and fix analysis information.

{% hint style="info" %}
The `snyk-to-html` command does not generate the standard exit codes.

For an example of a workaround for integrating the `snyk-to-html` command in a CI/CD pipeline, see the [Snyk CI/CD integration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/AzurePipelines/AzurePipelines-npm-generic-html.yml).

To display the help for `snyk-to-html`, use the command `snyk-to-html --help`. You can also use `--h`.
{% endhint %}

## Create a readable build artifact as part of the output of a scan

Follow these steps to create a build artifact by running `snyk-to-html` as part of a `snyk test` command.

1\. Change the directory to the root folder of the repository you want to test.

2\. To test the repository, export the results to a JSON file, and convert that file to an HTML file called `results` appended with the snyk test you ran.

Example commands for a JSON file  for `snyk code tes`t follow:

```
snyk code test --json | snyk-to-html > results-code.html
snyk code test --json | snyk-to-HTML -o results-code.html
```

{% hint style="info" %}
The filenames for the other `snyk test` commands are `results-opensource`, `results-container`, and `results-iac`.
{% endhint %}

If you want to customize the command, use the [command options](run-the-snyk-to-html-command.md#snyk-to-html-command-options).

An HTML file with one of these names is created in your repository folder, allowing you to [view your test results in HTML format](viewing-the-html-results.md).

## Convert a JSON or SARIF file to HTML to view in a browser

To run `snyk test` and then convert the output file to HTML, follow these steps:

1\. Change the directory to the root folder of the repository you want to test.

2\. Run the `test` command, `snyk code test` in this example, and export the results to a JSON OR SARIF file called, for example,  `results-code`:

```
snyk code test --json > results.json
```

{% hint style="info" %}
The filenames for the other `snyk test` commands are `results-opensource`, `results-container`, and `results-iac`.
{% endhint %}

3\. Pass the JSON file to `snyk-to-html` to be converted to HTML. Ensure you use the name of the output file you generated.

```
snyk-to-html -i results.json -o results.html
```

where:

`-i` is the input path of the JSON file.

`-o` precedes the name of the output file of generated HTML results.

## **`snyk-to-html` command options**

You can use the following options with the `snyk-to-html` command:

| **Short** | **Long**   | **Description**                                                                                                 | **Default** |
| --------- | ---------- | --------------------------------------------------------------------------------------------------------------- | ----------- |
| `-i`      | `--input`  | The input path of the JSON or SARIF file that contains the test results.                                        | `stdin`     |
| `-o`      | `--output` | <p>Precedes the name of the output file of the HTML results.</p><p>Example:<br><code>-o results.html</code></p> | `stdout`    |
| `-d`      | `--debug`  | Run the command in debug mode.                                                                                  |             |
