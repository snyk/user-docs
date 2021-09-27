# Test your CloudFormation files with CLI tool

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

Snyk Infrastructure as Code for CloudFormation supports scanning yaml & json formats.

{% hint style="info" %}
You can also scan AWS CDK applications. See [Test your AWS CDK files with our CLI tool](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-aws-cdk-files-with-our-cli-tool).
{% endhint %}

You can use the CLI as follows:

## To test for an issue on specified files:

```text
snyk iac test
```

For example, from the CLI enter the following:

```text
snyk iac test deploy.yaml
```

You can also specify multiple files by appending the file names after each other, such as:

```text
snyk iac test file-1.yaml file-2.yaml
```

