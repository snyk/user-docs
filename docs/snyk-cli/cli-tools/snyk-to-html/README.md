# snyk-to-html

Use the the `snyk-to-html` plugin when you want to do either of the following:

* Create a readable build artifact as part of the output of a scan
* Create a readable report of the full context of a vulnerability

The CLI provides a direct or automated way to fail the build and, by default, provides only summary information unless you use the `--json` or `--sarif` format to direct the output to a file. Those formats are not human-readable.

For a full report that you can read, convert the `--json` output to HTML by using the s`nyk-to-html` plugin and view the HTML report in a browser.

For example, when you run the `snyk code test` command, summary results are displayed that do not include the data flow and fix analysis information you can see in the Web UI for each discovered issue:

![](<../../../.gitbook/assets/Snyk-to-HTML - Results in the CLI Terminal - 2.png>)

When you use `snyk-to-html` to convert test results to an HTML file, you can view the vulnerability findings with their data flow and fix analysis information as an HTML report in a browser:

![](<../../../.gitbook/assets/Snyk-to-HTML - HTML Report - 2.png>)

The steps to use `snyk-to-htm`l are as follows:

* [Download and install ](install-snyk-to-html.md)the `snyk-to-html` plugin.
* [Run the `snyk-to-htm`l command ](running-the-snyk-to-html-command.md)to export the test results to a JSON or SARIF file and convert the information in the JSON or SARIF file to an HTML file, either as a readable build artifact or a file you can open and read in a browser.
* [View the HTML report in a browser](viewing-the-html-results.md).
