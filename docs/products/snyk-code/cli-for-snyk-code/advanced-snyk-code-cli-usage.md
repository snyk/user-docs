# Filtering and exporting Snyk Code CLI results

The following are examples of specific uses of the `snyk code` CLI command.

## Display only issues above a specific severity level

Usage:

```
snyk code test <folder> --severity-threshold=<low|medium|high|critical>
```

Example: display only results with a severity value of `medium` or higher:

```
snyk code test ./my-proj --severity-threshold=medium
```

## Output the test format as JSON

Usage:

```
snyk code test <folder> --json
```

This can be helpful if you want to store a snapshot of the results locally, or process the results in another tool for reporting and further analysis.

Example: output the results for a test of `./my-proj` as JSON:

```
snyk code test ./my-proj --json
```

## Output the test format as SARIF

{% hint style="info" %}
SARIF is an open standard for the output of static analysis tools. See the [official SARIF page](https://sarifweb.azurewebsites.net) for details.
{% endhint %}

Example: view the results of your tests as a SARIF file for analysis in another tool:

```
snyk code test ./my-proj --sarif
```

Example: save the results to an output SARIF file:

```
snyk code test ./my-proj --sarif-file-output=snyk.sarif
```
