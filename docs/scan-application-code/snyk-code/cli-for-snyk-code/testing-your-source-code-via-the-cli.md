# Testing your source code via the CLI

Snyk Code enables you to test the source code of your repositories via the Snyk CLI.

**Notes**:

* When testing a folder, all its sub-folders and files will be tested as well.
* To exclude certain directories or files from the Snyk Code CLI test, you can use the following Exclusion methods:
  * Using the CLI `snyk ignore --file-path` command – see [Excluding directories and files from the Snyk Code test](excluding-directories-and-files-from-the-snyk-code-cli-test.md).
  * Manually creating a `.snyk` file in the tested folder – see [Excluding directories and files from the import process](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process).

When testing your repository code via the CLI, you can:

* Test the repository directly from its root folder (see below).
* [Test the repository from another location](testing-your-source-code-via-the-cli.md#testing-a-repository-from-a-different-location).

### **Testing a repository from its root folder**

**To test the current repository folder:**

* In the terminal, enter:

```
snyk code test
```

**Note**: No additional parameters are required for running Snyk Code to test a repository from its root folder.

Snyk Code tests the current folder, and displays the [test results](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/snyk-code-cli-results) in the terminal.

For example:

To test the **snyk-goof** repository from its root folder, we first changed the directory to the root folder of the repository. Then, we enter:

```
snyk code test
```

Snyk Code tests the **snyk-goof** repository, and displays the vulnerability issue that were discovered in the terminal:

![](<../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - 1 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4).png>)

### **Testing a repository from a different location**

**To test a repository from another folder:**

* In the terminal, enter:

```
snyk code test <path/to/folder>
```

Where:

* `path/to/folder` – the full path of the repository you want to test via Snyk Code.

For example:

To test the **snyk-goof** repository from another directory, we enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof
```

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Any folder - 2 (1).png>)

* To explore the test results, see [Understanding the Snyk Code CLI results](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/snyk-code-cli-results).
* To work with the test results, see:
  * [Displaying only discovered issues above a specific severity level](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/displaying-only-discovered-issues-above-a-specific-severity-level).
  * [Outputting the test results to a JSON or SARIF format in the terminal](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-a-json-or-sarif-format-in-the-terminal).
  * [Exporting the test results to a JSON or SARIF file](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file).
  * [Displaying the CLI results in an HTML format using the Snyk-to-HTML feature](https://docs.snyk.io/products/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature).
