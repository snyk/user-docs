# Test your AWS CDK files with Snyk CLI

With Snyk Infrastructure as Code, you can test your configuration files with the CLI. You can scan the [Amazon Cloud Development Kit (CDK)](https://aws.amazon.com/cdk/) with the Snyk CLI by generating a CloudFormation file using the CDK CLI.

Follow these steps to scan a CDK:

**Navigate** to the directory that contains the application stack for which you want to generate the CloudFormation file.

**Generate** the CloudFormation file:

```
cdk synth
```

This is printed to your terminal as YAML output and a JSON file is created in the `cdk.out` directory

**Scan** the JSON file using the Snyk IaC CLI, replacing `cdk.out/.json` with the name of the application that you want to scan.

```
snyk iac test cdk.out/.json
```
