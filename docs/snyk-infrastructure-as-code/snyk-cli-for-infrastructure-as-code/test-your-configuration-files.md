# Test your configuration files

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

You can scan both your Kubernetes and Terraform files using the CLI.

You can find more detailed information here

* [Terraform](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-terraform-files-with-the-cli-tool)
* [Kubernetes](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-kubernetes-files-with-our-cli-tool)

{% hint style="info" %}
As of version 1.594.0 all configuration files are processed locally, ensuring that they do not leave your machine.  
Earlier versions by default will send the configuration files to Snyk to be processed. We recommend you upgrade to the latest version of the CLI.
{% endhint %}

You can use the CLI as follows:

Where the examples show \`main.tf\` you can replace this for your filename e.g. \`deployment.yaml\`

## To test for an issue on specified files:

```text
snyk iac test
```

For example, from the CLI enter the following:

```text
snyk iac test main.tf
```

You can also specify multiple files by appending the file names after each other, such as:

```text
snyk iac test file-1.tf file-2.tf
```

## To test for an issue on a directory of files:

You can scan a directory of configuration files. This will scan recursively through all files & folders.

For example, to scan all directories relative to your current path

```text
snyk iac test
```

Or you can specify a specific folder to scan and recurse through

```text
snyk iac test my-folder
```

Alternatively, you can restrict the directory depth to be scanned as follows

```text
snyk iac test --detection-depth=3
```

This will limit search to provided directory \(or current directory if no PATH provided\) plus two levels of subdirectories.

## To output the test format as JSON:

```text
snyk iac test  --json
```

This can be helpful if you want to store a snapshot of the results locally, or process the results in another tool for reporting and further analysis.

For example, from the CLI enter the following:

```text
snyk iac test main.tf --json
```

## To output the test format as SARIF:

SARIF is an open-standard for the output of static analysis tools.

You can view and save the results of your tests as a SARIF file for analysis in another tool.

```text
snyk iac test main.tf --sarif
```

Or to save this to a file output, you can run

```text
snyk iac test main.tf --sarif-file-output=snyk.sarif
```

## To only display issues above a specific severity level:

```text
snyk iac test  --severity-threshold=
```

For example, from the CLI enter the following:

```text
snyk iac test main.tf --severity-threshold=medium
```

This will only display to the terminal results that have a severity value of medium or higher.

## To target a specific Snyk organisation:

You can control the severity settings of your security rules at the organisation level in the Snyk UI. By targeting a specific organisation in your CLI tests, you can determine which rules should be run and the severity of them.

```text
snyk iac test  --org=
```

For example, from the CLI enter the following:

```text
snyk iac test main.tf --org=infrastructure
```

You can also set the Org flag in Snyk Config, so you do not need to provide the flag each time

```text
snyk config set org=
```



