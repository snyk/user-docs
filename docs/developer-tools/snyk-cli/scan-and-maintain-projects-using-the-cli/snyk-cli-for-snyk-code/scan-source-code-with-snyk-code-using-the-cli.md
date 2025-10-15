# Scan source code with Snyk Code using the CLI

Snyk Code enables you to test the source code of your repositories using the Snyk CLI.

When testing your repository code via the CLI, you can:

* [Test the repository directly from its root folder](scan-source-code-with-snyk-code-using-the-cli.md#testing-a-repository-from-its-root-folder).
* [Test the repository from another location](scan-source-code-with-snyk-code-using-the-cli.md#testing-a-repository-from-a-different-location).

When you test a folder, all its sub-folders and files are also tested. You can test a single file in the current folder, or a single file in another folder by specifying the absolute path to the file.

You can also test files with a relative path reference by prefixing the path with `$PWD`, for example,  `snyk code test $PWD/path/to/file`. This works with bash.

To exclude certain directories or files from the Snyk Code CLI test, you can use the following means:

* The  `snyk ignore --file-path` command. See [Excluding directories and files from the Snyk Code test](exclude-directories-and-files-from-snyk-code-cli-tests.md).
* Manually creating a `.snyk` file in the tested folder. See [Excluding directories and files from the import process](../../../../scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import.md).

## Testing a repository from its root folder

To test the current repository folder, in the terminal, enter the following:

```
snyk code test
```

No additional options are required for using the `snyk code test` command to test a repository from its root folder.

Snyk Code tests the current folder and displays the [test results](view-snyk-code-cli-results.md) in the terminal.

For example, to test the `snyk-goof` repository from its root folder, first change the directory to the root folder of the repository. Then enter:

```
snyk code test
```

Snyk Code tests the `snyk-goof` repository, and displays the vulnerability issues that were discovered:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - 1 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Example of Snyk Code CLI test results"><figcaption><p>Example of Snyk Code CLI test results</p></figcaption></figure>

## Testing a repository from a different location

To test a repository from another folder, in the terminal, enter the following:

```
snyk code test <path/to/folder>
```

The `path/to/folder` is the full path of the repository you want to test using Snyk Code via the CLI.

For example, to test the **snyk-goof** repository from another directory, enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof
```

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Any folder - 2.png" alt="Example of Snyk Code CLI test results"><figcaption><p>Example of Snyk Code CLI test results</p></figcaption></figure>

* To explore the test results, see [View Snyk Code CLI results](view-snyk-code-cli-results.md).
* To work with the test results, see [Displaying the CLI results in an HTML format using the Snyk-to-HTML feature](../cli-tools/snyk-to-html.md).
