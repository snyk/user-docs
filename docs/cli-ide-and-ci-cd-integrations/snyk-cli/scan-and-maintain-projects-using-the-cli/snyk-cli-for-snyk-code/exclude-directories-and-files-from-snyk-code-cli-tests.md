# Exclude directories and files from Snyk Code CLI tests

When you test a Snyk Code repository using the CLI, you can exclude certain directories and files from the CLI test by using the `snyk ignore --file-path` command. When you run this command, the `.snyk` file is created automatically in your repository, containing the name of the directory or file you specified for exclusion.

{% hint style="info" %}
* You can also create the `.snyk` file manually in your repository, and use it to exclude directories and files from the CLI test. For more information about the manual creation of the `.snyk` file, see [Exclude directories and files from Project import](../../../../scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import.md).
* The `snyk ignore --file-path` command does not ignore specific vulnerability issues. It excludes only directories and files from the CLI test.
* Consider excluding directories and files only if you do not publish or compile them into production. If a trace goes through an excluded file or directory with existing vulnerabilities, Snyk might miss potential issues.
{% endhint %}

## **Exclude directories and files from the CLI test**

Follow these steps to exclude Snyk Code directories and files from the CLI test:

1\. In the terminal, change the directory to the folder you want to test.

The command `snyk ignore --file-path` applies only to the folder from which you are running it and the sub-folders and files of that folder.

2\. In the terminal, enter the following:

```
snyk ignore --file-path=<directory_or_file>
```

where `directory_or_file` is the name of the directory or file you want to exclude from the test, for example, `db.js`.

The `.snyk` file is created in the root folder, containing the directory or file that was specified for exclusion.

The `.snyk` file is created as a hidden file. If you do not see it in your root folder, use the **View hidden files** option on your machine.

3\. Optionally, to specify several directories or files for exclusion enter:

```
snyk ignore --file-path=<directory1_or_file1> && snyk ignore --file-path=<directory2_or_file2> && snyk ignore --file-path=<directory3_or_file3>
```

From now on, when you run the `snyk code test` command from the selected folder, the specified directories or files will be excluded from the test.

## Re-include excluded files in the CLI test&#x20;

To re-include in the test directories or files that were excluded from it, manually edit or delete the `.snyk` file.

1\. In the `snyk-goof-master` folder, 12 issues were found in three different files: `app.js`, `db.js`, and `routes/index.js`:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - before -2.png" alt="Issues found by a CLI test"><figcaption><p>Issues found by a CLI test</p></figcaption></figure>

2\. To exclude the `app.js` and `db.js` files, and display only issues that are discovered in the `routes/index.js` file, enter:

```
snyk ignore --file-path=app.js && snyk ignore --file-path=db.js
```

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example command.png" alt="snyk ignore command in the terminal"><figcaption><p><code>snyk ignore</code> command in the terminal</p></figcaption></figure>

3\. When you enter the command `snyk ignore`, the `.snyk` file is created automatically in the `snyk-goof-master` folder:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file.png" alt=".snyk file listed in the folder"><figcaption><p><code>.snyk</code>file listed in the folder</p></figcaption></figure>

This `.snyk` file contains the files  specified for exclusion:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file - content.png" alt="Contents of .snyk file"><figcaption><p>Contents of .snyk file</p></figcaption></figure>

4\. When the test runs again, the `app.js` and `db.js` files are excluded from the test, and the results show only the issues that were found in the `routes/index.js` file:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - after - 2.png" alt="Issues found after using the ignore command"><figcaption><p>Issues found after using the <code>ignore</code> command</p></figcaption></figure>
