# Displaying the CLI results in an HTML format using the Snyk-to-HTML feature

When using Snyk Code via the CLI, the test results are displayed in a concise text format in the terminal, and they do not include the Data Flow and Fix Analysis information that is available for each discovered issue in the Web UI. The Snyk-to-HTML feature enables you to convert these text results into an HTML file, and to display the vulnerability findings along with their Data Flow and Fix Analysis information as an HTML Report in a browser.

For example:

When using Snyk Code via the CLI for testing a certain repository, the vulnerability findings are displayed in the terminal in the following way:

![](<../../../.gitbook/assets/Snyk-to-HTML - Results in the CLI Terminal - 2.png>)

When converting the CLI textual findings into HTML using the Snyk-to-HTML feature, the following **HTML Report** is generated:

![](<../../../.gitbook/assets/Snyk-to-HTML - HTML Report - 2.png>)

The Snyk-to-HTML feature uses the Snyk-to-HTML tool, which you need to download and install before running the snyk-to-html command. Then, you should test your source code with Snyk Code via the CLI, and export the results to a JSON or SARIF file. The snyk-to-html command takes the information on the test results from the JSON/SARIF file, and convert it to an HTML file that can be displayed in a browser.

**The workflow of using the Snyk-to-HTML feature is as follows:**

1\. [Installing the Snyk-to-HTML tool via the CLI](installing-the-snyk-to-html-tool.md).

2\. [Running the Snyk Code test of your repository in the CLI, and exporting the results to a JSON/SARIF file](running-the-snyk-to-html-command.md#running-the-snyk-to-html-feature-in-several-steps).

3\. [Running the snyk-to-html command in the CLI, and generating an HTML file with the test results](running-the-snyk-to-html-command.md#running-the-snyk-to-html-feature-in-several-steps).

**Note**: [Steps 2-3 above can be done in one command line](running-the-snyk-to-html-command.md#running-the-snyk-to-html-command-in-one-step).

4\. [Opening the generated HTML file in a browser to view the discovered vulnerabilities](viewing-the-html-results.md).
