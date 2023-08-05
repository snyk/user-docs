# Displaying only discovered issues above a specific severity level

You can filter the test results that are shown in the CLI terminal and display only issues with a specific severity level and higher.

To display only issues above a specific severity level, enter the following:

```
snyk code test <path/to/folder> --severity-threshold=<low|medium|high|critical>
```

The results will include only issues with the specified severity level and issues with a higher severity level.

For example, in the `snyk-goof-master` folder, eight issues were found, four with a High severity level and four with Medium:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - before - 2.png" alt="CLI test results for Snyk Code with High and Medium severity"><figcaption><p>CLI test results for Snyk Code with High and Medium severity</p></figcaption></figure>

To display only issues with a High severity level and above, enter the following:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --severity-threshold=high
```

The results show four issues, all with a High severity level. Issues with a lower severity level are not displayed:

![CLI test results for Snyk Code with High severity](<../../../../.gitbook/assets/Snyk Code - CLI - snyk code test - Results - Filter Severity - Example - after - 2.png>)
