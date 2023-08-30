# Test your IaC files

With Snyk Infrastructure as Code, you can test your configuration files with the CLI. This page provides detailed information on how to use certain options for the `snyk iac test command`.  For information about all of the options see the `snyk iac test` command [help](../../../snyk-cli/commands/iac-test.md). For details about testing the various configuration files, see the following pages:

* [Test your Terraform files with Snyk CLI](test-your-terraform-files-with-snyk-cli.md)
* [Test your CloudFormation files with Snyk CLI](test-your-cloudformation-files-with-snyk-cli.md)
* [Test your AWS CDK files with Snyk CLI](test-your-aws-cdk-files-with-snyk-cli.md)
* [Test your Kubernetes files with Snyk CLI](test-your-kubernetes-files-with-snyk-cli.md)
* [Test your ARM files with Snyk CLI](test-your-arm-files-with-snyk-cli.md)

You can also test the following types of files:

* [Kustomize files](test-your-kustomize-files-with-snyk-cli.md)
* [Helm charts](test-your-helm-charts-with-snyk-cli.md)
* [Serverless files](test-your-serverless-files-with-snyk-cli.md)

{% hint style="info" %}
Beginning with CLI version 1.594.0, all configuration files are processed locally, ensuring that they do not leave your machine. Earlier versions by default send the configuration files to Snyk to be processed. Snyk recommends that you upgrade to the latest version of the CLI.
{% endhint %}

In the examples that follow, you can replace the sample file names with the names of your own files, like `deployment.yaml`.

## Test for an issue in specified files

When you provide no arguments, the `snyk iac test` command recursively traverses the current working directory and scans every file it finds:

```
snyk iac test
```

You can scan specific files under the current working directory. If you provide one or more file paths, the command scans only those files:

```
snyk iac test file-1.tf dir/file-2.tf
```

The command returns an error if you provide file paths outside the current working directory. For example, this is **not a valid invocation of the command**:

```
snyk iac test ../main.tf
```

## Test for an issue in a directory of files

When you provide no arguments, the command recursively traverses the current working directory and scans every file it finds:

```
snyk iac test
```

You can restrict the scan to a specific directory:

```
snyk iac test my-folder
```

You can limit the depth of the directories that are traversed. The current working directory has a depth of one; directories under the current working directory have a depth of two, and so on. For example, if you want to restrict the search to the current working directory and two more levels of directories, you can invoke the command like this:

```
snyk iac test --detection-depth=3
```

The command returns an error if you provide directory paths outside the current working directory. For example, this is **not a valid invocation of the command**:

```
snyk iac test ../my-folder
```

## Output the test format as JSON

Use the following command to receive output in the JSON file format:

```
snyk iac test  --json
```

This can be helpful if you want to store a snapshot of the results locally or process the results in another tool for reporting and further analysis.

Example:

```
snyk iac test main.tf --json
```

## Output the test format as SARIF

SARIF is an open standard for the output of static analysis tools. You can view and save the results of your tests as a SARIF file for analysis in another tool.

Use the following command to receive output in the SARIF file format

```
snyk iac test main.tf --sarif
```

To save this to a file output, you can run the following command:

```
snyk iac test main.tf --sarif-file-output=snyk.sarif
```

## Display issues only above a specific severity level

Use the following command to limit results displayed to issues above a specified severity.

```
snyk iac test  --severity-threshold=medium
```

Example:

```
snyk iac test main.tf --severity-threshold=medium
```

This displays only results that have a severity value of medium or higher.

## Target a specific Snyk Organization

You can control the severity settings of your security rules at the Organization level in the Snyk UI. By targeting a specific Organization in your CLI tests, you can determine which rules should be run and their severity.

Use the following command to specify the Organization:

```
snyk iac test  --org=infrastructure
```

Example:

```
snyk iac test main.tf --org=infrastructure
```

You can also set the `org` flag in `snyk config`, so you do not need to use the `--org` option each time you want to specify the Organization.

```
snyk config set org=infrastructure
```
