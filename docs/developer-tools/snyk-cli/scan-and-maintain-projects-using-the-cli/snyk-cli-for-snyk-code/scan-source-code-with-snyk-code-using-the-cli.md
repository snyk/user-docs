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

## Publishing Snyk Code CLI results

{% hint style="info" %}
The minimum supported CLI version is v1.1297.0. For best results, Snyk recommends using version v1.1299.1+ or later.
{% endhint %}

You can publish Snyk Code results to a Snyk Project with or without using an integration. Here’s how it works and what you need to know.

### How Snyk organizes CLI scans

Snyk uses the `--project-name` or `--target-name` you provide in the command to identify which Project to update.

You don't need to connect to a Git repository such as GitHub or GitLab for this to work. It works directly with the results you upload from Snyk CLI.&#x20;

* Sync with Snyk Web UI: After creating a Project from the CLI, you can manage it in the Snyk web interface. If you mark an issue as **Ignored** in the UI, future CLI scans for that Project are not being reported again.&#x20;
* If the Project is new, Snyk automatically creates a new CLI Project for you with the value provided in the `--project-name` option.
* If the Project already exists, Snyk adds your latest scan as a new snapshot to that same Project. This allows you to track its security history over time.

### Testing a specific commit using Snyk CLI

You can use the Snyk CLI to test a specific commit in a Git repository that's already connected to Snyk.&#x20;

When you run the `snyk code test` command, you can specify which SCM Project to update and the exact commit ID you want to test. Snyk then triggers a scan on that specific commit within your Git repository without scanning your local code. The results are saved as a new snapshot in your existing SCM Project. Any issues you have already marked as **Ignored** in Snyk Web UI for that Project will also be ignored in the CLI scan results (see [Ignore issues in the CLI](../../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/#ignore-issues-in-the-cli)).

See the CLI help for more information about the `snyk code test` command, including [exit codes](../../commands/code-test.md#exit-codes) and [available options](../../commands/code-test.md#options).

## **Publishing CLI results to a Snyk Code Project** <a href="#publish-cli-results-to-a-snyk-code-project" id="publish-cli-results-to-a-snyk-code-project"></a>

Using Snyk Code through the CLI allows you to publish test results of local code to a Snyk Project in Snyk Web UI. Future CLI tests of this Project will respect issues that were ignored in the Web UI.

This allows using Snyk Code as a blocking CI/CD gate to test and block builds at the main branch level and then have developers review the results in the Web UI, fix any newly introduced vulnerabilities, or ignore irrelevant ones.

In the terminal, enter the following command:

```
snyk code test --report --project-name="<PROJECT_NAME>"
```

* `project-name` must be in double quotation marks. Single quotes or missing quotes will result in an error.
* `project-name` must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).

After using this option, log in to Snyk and view your Projects to see the snapshot.

### Commands to publish Snyk Code CLI results <a href="#commands-to-publish-snyk-code-cli-results" id="commands-to-publish-snyk-code-cli-results"></a>

Running the `snyk code test` command with the `--report` option, as shown, returns the results to the terminal window, along with a URL to the Snyk Code Project where the results have been published. Refer to the following screenshot.

<figure><img src="../../../../.gitbook/assets/image (2) (6).png" alt=""><figcaption><p>Snyk code test results with --report option</p></figcaption></figure>

If a Snyk Code Project created using the CLI does not yet exist for the value provided in the `--project-name` option, the Snyk CLI creates a new Project. If a Project created by using the CLI exists, a new snapshot is created under the same Project.

To make the Project easier to interpret in Snyk Web UI, you can use additional commands to specify a target name and also target references, such as Git branches. The following command will create or upload an existing Project named `<PROJECT_NAME>` under a target named `<TARGET_NAME>`.

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>"
```

* Both `<PROJECT_NAME>` and `<TARGET_NAME>` must be in double quotation marks. Single quotes or missing quotes will result in an error.
* The Project name and Target name each must contain only alphanumeric characters, forward slashes (/), dashes (-), underscores (\_), and square brackets (\[]).

The following command creates or uploads an existing Project named `<PROJECT_NAME>` under a target named `<TARGET_NAME>` and grouped by the "`$(git branch --show-current)"` branch name.

```
snyk code test --report --project-name="<PROJECT_NAME>" --target-name="<TARGET_NAME>" --target-reference="$(git branch --show-current)"
```

<figure><img src="../../../../.gitbook/assets/image (4) (4).png" alt=""><figcaption><p>Code analysis Projects grouped by branch</p></figcaption></figure>

### **Troubleshooting published Snyk Code CLI results** <a href="#troubleshooting-publication-of-snyk-code-cli-results" id="troubleshooting-publication-of-snyk-code-cli-results"></a>

You may see this error: `There was a problem running Code analysis. The findings for this project may exceed the allowed size limit.`

This error indicates that the contents of the scanned Project exceed the limit. To complete the scan, consider the following troubleshooting steps:

* Partition the Project repository by scanning sub-directories instead of the whole repository, for example:
  * Create two Projects for your frontend and backend directories, and scan them separately.
  * Create and scan Projects for each MicroService.
* Exclude unnecessary files from the scanning process using the [.snyk exclude option](exclude-directories-and-files-from-snyk-code-cli-tests.md#exclude-directories-and-files-from-the-cli-test). For example, you can exclude test files from the scan.
* Set a severity threshold using the [`--severity-threshold=high`](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli#combining-security-policies-with-severity-threshold) to focus on more critical issues and gain visibility into urgent matters.

## Use the CLI to trigger a test for a specific commit to an existing Snyk Code Git repository Project <a href="#use-the-cli-to-trigger-a-test-for-a-specific-commit-to-an-existing-snyk-code-git-repository-project" id="use-the-cli-to-trigger-a-test-for-a-specific-commit-to-an-existing-snyk-code-git-repository-project"></a>

When you are using Snyk Code in the CI/CD, you can specify an existing Git repository Code Analysis Project along with the commit ID to trigger a test for which the results will be saved in the Web UI. The Web UI will also display code snippets that provide [Snyk Code Data flow](../../../../scan-with-snyk/snyk-code/manage-code-vulnerabilities/breakdown-of-code-analysis.md#data-flow). Ignores that have been added to the Web UI Project will be respected in this test.

#### Prerequisites

* A Git repository [Integration](../../../scm-integrations/organization-level-integrations/) is required, as the CLI command triggers a test using a Git repository Project.
* The Project ID is required. You can find the Project ID in the URL of the Project in the Snyk Web UI: `https://app.snyk.io/org/org_name/project/PROJECT_UUID.`
* After you run the CLI command, results are saved in the Git repository Project that you have specified. Given that scheduled tests for Git repository Projects scan the default branch, consider using this CLI command with the Project ID for the Project where the default branch for the specified repository is being monitored.
* In the Code Analysis Project, three snapshots are stored at a time. Consider outputting a static [JSON](../../commands/code-test.md#json) or[ SARIF ](../../commands/code-test.md#sarif)file to save scan results for future reference.

Use the following command to publish to a specific Project:

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```

## **Ignore CLI results for Snyk Code** <a href="#ignore-cli-results-for-snyk-code" id="ignore-cli-results-for-snyk-code"></a>

You can ignore issues in Snyk Web UI. The ignores will be used both to [publish CLI results to a Snyk Code Project](scan-source-code-with-snyk-code-using-the-cli.md#publish-cli-results-to-a-snyk-code-project) and in [using the CLI to trigger a test for a specific commit to an existing Snyk Code Git repository Project](scan-source-code-with-snyk-code-using-the-cli.md#testing-a-specific-commit-using-snyk-cli).

<figure><img src="../../../../.gitbook/assets/image (1) (7) (1).png" alt=""><figcaption><p>Ignoring issues in the Web UI</p></figcaption></figure>

[snyk-to-html](https://github.com/snyk/snyk-to-html) currently does not respect the ignored issues. Anything that is ignored in Snyk Web UI is not ignored in the report that `snyk-to-html` generates.

For [publishing workflows](scan-source-code-with-snyk-code-using-the-cli.md#publish-cli-results-to-a-snyk-code-project), after the CLI results are published to a Snyk Code Project, issues that were ignored in the Web UI will be ignored in CLI tests when you use the following command:

```
snyk code test --report --project-name="PROJECT_NAME"
```

* Assuming that the --project-name provided matches what is in the Web UI
* Ignores that have been applied to the Project with a `PROJECT_NAME` suppress the issue the next time that the CLI runs for the same `PROJECT_NAME`.

For the [trigger using the CLI workflow](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/publish-snyk-code-cli-results-and-ignore-issues#use-the-cli-to-trigger-a-snyk-code-git-repository-project-to-test-a-specific-commit-in-a-git-reposit), ignores added to the Git repository monitored Project will also be ignored the next time an SCM scan is triggered from the CLI for a given `<PROJECT_UUID>` and `<COMMIT_ID>.`

An example follows:

```
snyk code test --report --project-id="<PROJECT_UUID>" --commit-id="<COMMIT_ID>"
```
