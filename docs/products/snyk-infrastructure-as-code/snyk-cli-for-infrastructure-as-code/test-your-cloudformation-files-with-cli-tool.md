# Test your CloudFormation files with Snyk CLI

With Snyk Infrastructure as Code, you can test your CloudFormation files using the CLI. You can test files in both YAML and JSON formats. You can also scan AWS CDK applications. See [Test your AWS CDK files with Snyk CLI](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-aws-cdk-files-with-our-cli-tool).

Test for an issue on specified files using the following command:

```
snyk iac test
```

Example:

```
snyk iac test deploy.yaml
```

You can also specify multiple files by appending the file names after each other, for example:

```
snyk iac test file-1.yaml file-2.yaml
```
