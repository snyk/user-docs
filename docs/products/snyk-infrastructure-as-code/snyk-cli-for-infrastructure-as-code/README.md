# Snyk CLI for Infrastructure as Code

## Overview

To use the CLI you must first [install](../../../snyk-cli/install-the-snyk-cli.md) it and then [authenticate](../../../snyk-cli/commands/auth.md).

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI. See the following pages for details:

* [Test your configuration files](test-your-configuration-files.md)
* [IaC ignores using the `.snyk` policy file](iac-ignores-using-the-.snyk-policy-file.md)
* [Test your Terraform files with Snyk CLI](test-your-terraform-files-with-the-cli-tool.md)
* [Test your CloudFormation files with Snyk CLI](test-your-cloudformation-files-with-cli-tool.md)
* [Test your AWS CDK files with Snyk CLI](test-your-aws-cdk-files-with-our-cli-tool.md)
* [Test your Kubernetes files with Snyk CLI](test-your-kubernetes-files-with-our-cli-tool.md)
* [Test your ARM files with Snyk CLI](test-your-arm-files-with-the-cli-tool.md)

You can also test the following types of files:

* [Kustomize files](test-your-kustomize-files-with-our-cli-tool.md)
* [Helm charts](test-your-helm-charts-with-our-cli-tool.md)

See [Understanding the CLI output](understanding-the-cli-test-output/) for information about using the reports.

## Regularly testing IaC files

Currently, there is no equivalent command to `snyk monitor` for Snyk IaC because the CLI does not send IaC source files back to the platform for periodic testing.

For IaC CLI results to appear in the Snyk Web UI, you can use the [`snyk iac test --report`](https://docs.snyk.io/products/snyk-infrastructure-as-code/share-cli-results-with-the-snyk-web-ui) command to capture a one-time snapshot. Optionally, run the command on a recurring schedule to regularly test your IaC.

Alternatively, you can add an [SCM integration](https://docs.snyk.io/integrations/git-repository-scm-integrations) and Snyk will monitor and test a given git repository on a recurring basis.
