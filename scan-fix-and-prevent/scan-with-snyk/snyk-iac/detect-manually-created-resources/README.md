# Detect manually created resources

Snyk IaC can report unmanaged resources. Unmanaged resources are resources that are present on your cloud provider but not on your Terraform state. You can import these resources into Terraform or delete them from your IaaS account.

For information about detecting drift of managed resources, see [Command: plan](https://developer.hashicorp.com/terraform/cli/commands/plan) in the Terraform CLI documentation.

The information in this group of pages supports using the `snyk iac describe` command. Information is provided about the following:

* [Get started with Snyk IaC Describe on AWS](get-started-with-snyk-iac-describe-on-aws.md)
* [Configure cloud providers](configure-cloud-providers/)
* [Supported resources](supported-resources/)
* [IaC describe command examples](iac-describe-command-examples.md)
* [Filter rules](filter-rules.md)
* [Ignore resources for drift](ignore-unmanaged-resources.md)
* [IAC sources usage](iac-sources-usage.md)
