# AWS CDK files

With Snyk Infrastructure as Code, you can test your configuration files with the CLI. You can scan the [Amazon Web Services Cloud Development Kit (AWS CDK)](https://aws.amazon.com/cdk/) with the Snyk CLI by generating a CloudFormation file using the CDK CLI.

Follow these steps to scan a CDK application:

**Navigate** to the directory that contains the application stack for which you want to generate the CloudFormation.

**Generate** the CloudFormation file.

```
cdk synth
```

This is displayed on your terminal as YAML output, and a JSON file is created in the `cdk.out` directory

**Scan** the JSON file using the following Snyk IaC CLI command, replacing `cdk.out/*.json` with the name of the application that you want to scan.

```
snyk iac test cdk.out/*.json
```
