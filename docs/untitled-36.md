# What languages do we support Fix Pull Requests or Merge Requests?

##  Test your AWS CDK files with our CLI tool

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

If you are using [Amazons Cloud Development Kit \("CDK"\)](https://aws.amazon.com/cdk/) you scan this using our CLI by generating a CloudFormation file using the CDK CLI. 

### Scan a CDK application:

First, navigate to the directory that contains your application stack that you want to generate the CloudFormation for

Next, generate the CloudFormation file

```text
cdk synth
```

This will be printed to your terminal as YAML output and a JSON file will be created in the **cdk.out** directory

Scan the created json file using the Snyk IaC CLI, replacing &lt;applicationstack&gt; with the name of your application that you want to scan. 

```text
snyk iac test cdk.out/<applicationstack>.json
```

