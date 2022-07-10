# Excluding directories and files from the Snyk Code CLI test

{% hint style="info" %}
This Exclusion command is applicable to Snyk Code only.
{% endhint %}

When you test a repository via the CLI, you can exclude certain directories and files from the Snyk Code test by using the `snyk ignore --file-path` command. When running this command, the `.snyk` file is created automatically in your repository, containing the name of the directory or file you specified for exclusion.

**Notes**:

* You can also manually create the `.snyk` file in your repository, and use it to exclude directories and files from the CLI test. For more information about the manual creation of the `.snyk` file, see [Excluding directories and files from the import process.](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process)
* The `snyk ignore --file-path` command does NOT ignore specific vulnerability issues. It ONLY excludes directories and files from the Snyk Code test.

****

**To exclude directories and files from the Snyk Code test:**

1\.  In the terminal, change the directory to the folder you want to test.

**Note**: The `snyk ignore --file-path` command applies only to the folder (and its sub-folders and files) from which you are running it.

2\.  In the terminal, enter the following:

```
snyk ignore --file-path=<directory_or_file>
```

Where:

* `directory_or_file` - the name of the directory or file you want to exclude from the test. For example, `db.js`.

The `.snyk` file is created in the root folder, containing the directory/file that was specified for exclusion.

**Note**: The `.snyk` file is created as a hidden file. If you do not see it in your root folder, use the **View hidden files** option on your machine.

3\.  \[Optional] To specify several directories or files for exclusion, enter:

```
snyk ignore --file-path=<directory1_or_file1> && snyk ignore --file-path=<directory2_or_file2> && snyk ignore --file-path=<directory3_or_file3>
```

From now on, when you will run the `snyk code test` command from the selected folder, the specified directories or files will be excluded from the test.

**Note**: To re-include in the test directories/files that were excluded from it, manually edit or delete the `.snyk` file.

&#x20;

For example:

1\.  In the **snyk-goof-master** folder, 12 issues were found in 3 different files: **app.js**, **db.js**, and **routes/index.js**:

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - before -2.png>)

2\.  To exclude the **app.js** and **db.js** files, and display only issues that are discovered in the **routes/index.js** file, we enter:

```
snyk ignore --file-path=app.js && snyk ignore --file-path=db.js
```

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example command.png>)

3\.  When the `snyk ignore` command is entered, the `.snyk` file is created automatically in the **snyk-goof-master** folder:

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file.png>)

This `.snyk` file contains the files we specified for exclusion:

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file - content.png>)

4\.  When we run the test again, the **app.js** and **db.js** files are excluded from the test, and we receive only the issues that were found in the **routes/index.js** file:

![](<../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - after - 2.png>)

&#x20;
