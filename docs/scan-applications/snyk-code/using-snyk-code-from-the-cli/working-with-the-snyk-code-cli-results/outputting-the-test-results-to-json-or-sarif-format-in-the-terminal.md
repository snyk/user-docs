# Outputting the test results to JSON or SARIF format in the terminal

You can output the `snyk code test` results to [JSON](outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md#outputting-the-test-results-to-json-format) or [SARIF](outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md#outputting-the-test-results-to-sarif-format) format in the terminal instead of displaying the results in the Snyk CLI format.

You can also [export the test results to a JSON or SARIF format file](../../../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md).

SARIF is an open standard for the output of static analysis tools. For more information, see the [SARIF site](https://sarifweb.azurewebsites.net/).

The severity levels of the issues discovered by running `snyk code test` are displayed differently in JSON and SARIF outputs. The severity levels in the JSON and SARIF results are as follows:

* High = **error**
* Medium = **warning**
* Low = **note/info**

{% hint style="info" %}
The designation Critical is not used in Snyk Code.
{% endhint %}

An example for medium level severity follows:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results - in the Terminal.png" alt="Medium severity in JSON or SARIF output"><figcaption><p>Medium severity in JSON or SARIF output</p></figcaption></figure>

## Outputting the test results to JSON format

To output the test results to a JSON format, enter the following:

```
snyk code test <path/to/folder> --json
```

The test results appear in the terminal in JSON format.

For example, to output the test results of the `snyk-goof-master` folder in JSON format in the terminal, enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --json
```

The test results appear in the terminal in JSON format:

<figure><img src="../../../../.gitbook/assets/snyk Code - CLI - results - JSON output in the terminal.png" alt="snyk code test results in JSON format"><figcaption><p><code>snyk code test</code> results in JSON format</p></figcaption></figure>

## Outputting the test results to SARIF format

To output the test results to SARIF format, enter the following:

```
snyk code test <path/to/folder> --sarif
```

The test results appear in the terminal in SARIF format.

For example, to output the test results of the `snyk-goof-master` folder in SARIF format, enter:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --sarif
```

The test results appear in the terminal in SARIF format:

![ code test results in SARIF format](<../../../../.gitbook/assets/snyk Code - CLI - results - SARIF output in the terminal.png>)
