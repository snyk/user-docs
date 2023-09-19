# Snyk CLI for Infrastructure as Code

## Overview

To use the CLI, you must first [install](../../../snyk-cli/install-or-update-the-snyk-cli/) it and then [authenticate](../../../snyk-cli/commands/auth.md).

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI. See the following pages for details:

* [Test your configuration files](test-your-iac-files/)
* [Share CLI results with the Snyk Web UI](share-cli-results-with-the-snyk-web-ui.md)
* [IaC ignores using the `.snyk` policy file](../../../scan-infrastructure/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file.md)
* [Understanding the IaC CLI test results](understanding-the-cli-test-output/) (has information about reports)
* [IaC exclusions using the command line](iac-exclusions-using-the-command-line.md)
* [Test your Terraform files with Snyk CLI](test-your-iac-files/test-your-terraform-files-with-snyk-cli.md)
* [Test your CloudFormation files with Snyk CLI](test-your-iac-files/test-your-cloudformation-files-with-snyk-cli.md)
* [Test your AWS CDK files with Snyk CLI](test-your-iac-files/test-your-aws-cdk-files-with-snyk-cli.md)
* [Test your Kubernetes files with Snyk CLI](test-your-iac-files/test-your-kubernetes-files-with-snyk-cli.md)
* [Test your ARM files with Snyk CLI](test-your-iac-files/test-your-arm-files-with-snyk-cli.md)
* [Test your Serverless files with Snyk CLI](test-your-iac-files/test-your-serverless-files-with-snyk-cli.md)

You can also test the following types of files:

* [Kustomize files](test-your-iac-files/test-your-kustomize-files-with-snyk-cli.md)
* [Helm charts](test-your-iac-files/test-your-helm-charts-with-snyk-cli.md)
* [Serverless files](test-your-iac-files/test-your-serverless-files-with-snyk-cli.md)

## Regularly testing IaC files

Snyk Infrastructure as Code has no equivalent command to `snyk monitor` because the CLI does not send IaC source files back to the platform for periodic testing.

For IaC CLI results to appear in the Snyk Web UI, use [`snyk iac test --report`](https://docs.snyk.io/products/snyk-infrastructure-as-code/share-cli-results-with-the-snyk-web-ui) to capture a one-time snapshot. Optionally, run the command on a recurring schedule to regularly test your IaC files.

Alternatively, you can add an [SCM integration](../../../integrations/git-repository-scm-integrations/), and Snyk will monitor and test a given Git repository on a recurring basis.

## Using Snyk behind a proxy

If you are using a proxy, see [Proxy configuration for Snyk CLI](../../../snyk-cli/configure-the-snyk-cli/proxy-configuration-for-snyk-cli.md).

For IaC scans specifically, you must also whitelist the \*.snyk.io address, as explained[ ](https://support.snyk.io/hc/en-us/articles/360002153077-How-can-we-whitelist-Snyk-IP-addresses-)on the page [How can we whitelist Snyk IP addresses?](https://support.snyk.io/hc/en-us/articles/360002153077-How-can-we-whitelist-Snyk-IP-addresses-)
