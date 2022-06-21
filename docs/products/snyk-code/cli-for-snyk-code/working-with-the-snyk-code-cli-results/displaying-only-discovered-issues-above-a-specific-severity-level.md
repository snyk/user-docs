# Displaying only discovered issues above a specific severity level

You can filter the test results that are shown in the CLI terminal, and display only issues with a specific severity level and higher.

**To display only issues above a specific severity level:**

* In the terminal, enter:

```
snyk code test <path/to/folder> --severity-threshold=<low|medium|high|critical>
```

The results will include only issues with the specified severity level and issues with a higher severity level.

For example:

In the **snyk-goof-master** folder, 8 issues were found - 4 with a High severity level and 4 with Medium:

![](<../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - before (1).png>)

To display only issues with a High severity level and above, we enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --severity-threshold=high
```

The results show only 4 issues, all with a High severity level. Issues with a lower severity level are not displayed:

![](<../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - after.png>)
