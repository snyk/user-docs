# Test your CloudFormation files with our CLI tool

## Test your CloudFormation files with our CLI tool

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

Snyk Infrastructure as Code for CloudFormation supports scanning yaml & json formats.

You can use the CLI as follows:

### To test for an issue on specified files:

```text
snyk iac test <my-cloudformation-filepath>
```

For example, from the CLI enter the following:

```text
snyk iac test deploy.yaml
```

You can also specify multiple files by appending the file names after each other, such as:

```text
snyk iac test file-1.yaml file-2.yaml
```

