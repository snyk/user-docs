# Using Snyk Code via the CLI

The [Snyk Command Line Interface](../../../snyk-cli/) (CLI) for Snyk Code enables you to bring the functionality of Snyk Code into your development workflow. By using the Snyk CLI, you can run Snyk Code tests locally, or incorporate them with your CI/CD pipeline to scan your source code for security vulnerabilities.

### Prerequisites for the Snyk Code CLI

Before you start using the Snyk CLI for testing your source code with Snyk Code, verify you have the following prerequisites:

* A Snyk Account.\
  **Note**: For more information on creating a Snyk Account, see [Setting Up your Snyk Account](https://docs.snyk.io/tutorials/getting-started/snyk-integrations/snyk-account).
* Repositories with code in a supported language and framework.\
  **Note**: For more information, see [Snyk Code - Supported languages and frameworks](../snyk-code-ai-engine-web-ui-supported-integrations-languages-frameworks/snyk-code-supported-languages-and-frameworks.md).
* [The **Snyk Code** option enabled in your Snyk Org Setting](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-1-enabling-the-snyk-code-option)
* An installed and authenticated Snyk CLI.\
  **Notes**:
  * For instructions see [Install or update the Snyk CLI](../../../snyk-cli/install-the-snyk-cli/) and [Authenticate the Snyk CLI](https://docs.snyk.io/snyk-cli/authenticate-the-cli-with-your-account).
  * The minimum Snyk CLI version for Snyk Code is 1.716.0. Snyk recommends using the latest version of the CLI. See [Install or update the Snyk CLI ](../../../snyk-cli/install-the-snyk-cli/)for instructions.

**Note**: For more information about Snyk Code, see [Getting started with Snyk Code](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code).

### Using the Snyk CLI for Snyk Code tests

To test your repository code via the Snyk Code CLI, use the `snyk code test` command.\
For more information, see [Testing your source code via the CLI](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/testing-your-source-code-via-the-cli).

Before the test, you may want to do the following:

* [Update your Snyk CLI version](../../../snyk-cli/install-the-snyk-cli/).
* [Set the organization for the CLI tests](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/before-you-start-set-the-organization-for-the-cli-tests).
* [Exclude directories and files from the Snyk Code test](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/excluding-directories-and-files-from-the-snyk-code-cli-test).

After the test, you can do the following:

* [Explore the Snyk Code CLI results](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/snyk-code-cli-results).
* [Filter the CLI results, and display only issues with a specific severity level and higher](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/displaying-only-discovered-issues-above-a-specific-severity-level).
* [Output the CLI results to a JSON or SARIF format in the terminal](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal).
* [Export the CLI results to a JSON or SARIF format file](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file).
* [Display the CLI results in an HTML format using the Snyk-to-HTML feature](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature).
