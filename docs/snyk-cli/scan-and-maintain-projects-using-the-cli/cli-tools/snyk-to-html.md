# snyk-to-html

The CLI provides a direct or automated way to fail the build and, by default, provides only summary information unless you use the `--json` or `--sarif` format to direct the output to a file. These formats are not human-readable.

You can use `snyk-to-html` (the Snyk JSON to HTML Mapper) to:

* Create a readable build artifact as part of the output of a scan
* Create an HTML report of the full context of a vulnerability that you can view in a browser

This page explains how to [install `snyk-to-html`](snyk-to-html.md#install-snyk-to-html), [run the `snyk-to-html` command](snyk-to-html.md#run-the-snyk-to-html-command) to export the test results to a JSON or SARIF file, and [view the test results in a browser](snyk-to-html.md#view-test-results-in-html-format).

## Install `snyk-to-html`

To install `snyk-to-html`, you must have the required permissions on your machine. If the installation fails, contact your IT administrator to request the required permissions.

To install the `snyk-to-html` plugin locally only, clone the [`snyk-to-html`](https://github.com/snyk/snyk-to-html) GitHub repository.

To install `snyk-to-html`, in the terminal, enter:

```
npm install snyk-to-html -g
```

## Run the `snyk-to-html` command

You can run `snyk-to-html` as part of a `snyk test` command to create a readable build artifact as part of the output of a test.

You can also run a `snyk test` command with the option to export the results to a JSON file and then convert the JSON  file to HTML using  `snyk-to-html`. You can also export Snyk Code results to a SARIF file and convert that file to HTML.

When you run the `snyk-to-html` command, you can customize it with the following command options:

<table data-header-hidden><thead><tr><th width="105"></th><th width="134"></th><th></th><th></th></tr></thead><tbody><tr><td><strong>Short</strong></td><td><strong>Long</strong></td><td><strong>Description</strong></td><td><strong>Default</strong></td></tr><tr><td><code>-i</code></td><td><code>--input</code></td><td>The input path of the JSON or SARIF file that contains the test results.</td><td><code>stdin</code></td></tr><tr><td><code>-o</code></td><td><code>--output</code></td><td><p>Precedes the name of the output file of the HTML results.</p><p>Example:<br><code>-o results.html</code></p></td><td><code>stdout</code></td></tr><tr><td><code>-d</code></td><td><code>--debug</code></td><td>Run the command in debug mode.</td><td></td></tr></tbody></table>

The `snyk-to-html` command does not generate the standard exit codes.

To display the help for `snyk-to-html`Use the command `snyk-to-html --help` or `--h`.

If you want to use the `snyk-to-html` command in a CI/CD pipeline, see the [Snyk CI/CD integration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/AzurePipelines/AzurePipelines-npm-generic-html.yml) for an example of a workaround for Azure Pipelines.

### Create a readable build artifact as part of the output of a test

Follow these steps to create a build artifact by running `snyk-to-html` as part of a `snyk test` command.

1. Change the directory to the root folder of the repository you want to test.
2. To test the repository, export the results to a JSON file, and convert that file to an HTML file called `results` appended with output of the `snyk test` you ran.

The output filenames for the `snyk test` commands are `results-opensource`, `results-container`, `results-code`, and `results-iac`.

Example commands for a JSON file  for `snyk code test` follow:

```
snyk code test --json | snyk-to-html > results-code.html
snyk code test --json | snyk-to-HTML -o results-code.html
```

An HTML file with one of these names is created in your repository folder, allowing you to [view your test results in HTML format](snyk-to-html.md#view-test-results-in-html-format).

### Convert a JSON or SARIF file to HTML to view in a browser

Follow these steps to run `snyk test` and then convert the output file to HTML:

1. Change the directory to the root folder of the repository you want to test.
2. Run the `test` command, as shown in the example that follows for `snyk code test`:\
   `snyk code test --json > results.json`\
   The output filenames for the `snyk test` commands are `results-opensource`, `results-container`, `results-code`, and `results-iac`.
3. Pass the JSON file to `snyk-to-html` to be converted to HTML. Ensure you use the name of the output file you generated:

```
snyk-to-html -i results.json -o results.html
```

where:

`-i` is the input path of the JSON file.

`-o` precedes the name of the output file of generated HTML results.

## View test results in HTML format

To view the HTML file, locate the output file in your repository and double-click it. The following example shows a `results.html` file.

<figure><img src="../../../.gitbook/assets/Snyk-to-HTML - Results file.png" alt="results.html file in the repository"><figcaption><p><code>results.html</code> file in the repository</p></figcaption></figure>

If you used a different name for your HTML file, locate and open that file.

The test results report opens in the browser. The following example shows `snyk code test` results.

<figure><img src="../../../.gitbook/assets/Snyk-to-HTML - HTML Report - 2.png" alt="Snyk Code Report"><figcaption><p>Snyk Code Report</p></figcaption></figure>

You can view the **Data Flow** and **Fix Analysis** information for the discovered vulnerabilities by clicking the corresponding buttons for each issue:

<figure><img src="../../../.gitbook/assets/Snyk-to-HTML - Example - HTML Report - Fix Analysis tab - 2.png" alt="Data Flow and Fix Analysis buttons for an issue"><figcaption><p>Data Flow and Fix Analysis buttons for an issue</p></figcaption></figure>

