# Advanced Snyk Code CLI usage

### To output the test format as JSON

Enter:

```text
snyk code test <folder> --json
```

This can be helpful if you want to store a snapshot of the results locally, or process the results in another tool for reporting and further analysis.

For example, from the CLI enter the following:

```text
snyk code test ./my-proj --json
```

### To output the test format as SARIF

{% hint style="info" %}
SARIF is an open-standard for the output of static analysis tools. See the [official SARIF page](https://sarifweb.azurewebsites.net/) for details.
{% endhint %}

You can view and save the results of your tests as a SARIF file for analysis in another tool:

```text
snyk code test ./my-proj --sarif
```

Or to save this to a file output, you can run:

```text
snyk code test ./my-proj --sarif-file-output=snyk.sarif
```

### To only display issues above a specific severity level

Enter:

```text
snyk code test <folder> --severity-threshold=<low|medium|high>
```

For example, to only display results with a severity value of **medium** or higher:

```text
snyk code test ./my-proj --severity-threshold=medium
```

