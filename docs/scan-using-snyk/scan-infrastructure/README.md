# Scan infrastructure

{% hint style="info" %}
**Feature availability**\
IaC+ is a new version of Snyk IaC that includes more accurate results, an expanded security ruleset, and code to cloud capabilities. IaC+ is now in early access via [Snyk Preview](../../snyk-admin/manage-settings/snyk-preview.md).
{% endhint %}

With Snyk Infrastructure as Code (IaC), you can secure cloud infrastructure configurations before and after deployment. There are two version of Snyk IaC available today:

* **Current IaC**: The generally available version of Snyk IaC
* **IaC+**: A new version of Snyk IaC that is currently in early access

With both versions of Snyk IaC, you can:

* Write secure configurations for [HashiCorp Terraform](scan-your-iac-source-code/scan-terraform-files/), [AWS CloudFormation](scan-your-iac-source-code/scan-cloudformation-files/), [Kubernetes](scan-your-iac-source-code/scan-kubernetes-configuration-files/), and [Azure Resource Manager (ARM)](scan-your-iac-source-code/scan-arm-configuration-files.md) - for IDE, SCM, CLI, and Terraform Cloud/Enterprise workflows.
* View issues and receive [fix advice](getting-started-with-current-iac.md) so you can make changes directly to code, before applications reach production.
* [Detect drift](iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/) and manually created resources in your cloud.
* **Onboard, scan, and test deployed cloud environments** for misconfigurations for AWS, Azure, and Google Cloud environments.

IaC+ is built on a new engine and ruleset that also powers Snyk IaC’s cloud scanning capabilities. IaC+ enables the following improvements vs. Current IaC:

* Includes consistent support for languages - such as Azure Resource Manager - across all IaC workflows.
* Adds multi-file analysis for Terraform (support for modules and variables files).
* Utilizes an expanded security ruleset that is mapped to more than a dozen compliance standards (CIS Benchmarks, PCI, SOC 2, and more).
* Supports custom rules with Rego that are managed in the Snyk platform, and work consistently across all IaC workflows.
* Introduces projects (for SCM) that capture issues for an entire repository, instead of only for a single IaC file - in alignment with Snyk Code.
* Supports recurring (daily or weekly) scans for IaC+ SCM projects.
* Utilizes a new organization-wide Cloud Issues page for IaC+ and cloud issues that enables users to group issues by rule or resource, filter and inspect the configuration of relevant resources for a given issue, and take action on issues.

IaC+ also adds support for “code to cloud” use cases that work with Snyk IaC’s ability to onboard, scan, and test deployed cloud environments:

* [Fix Cloud issues](iac+-code-to-cloud-capabilities/fix-cloud-issues-in-iac.md) directly in the IaC source code that was used to deploy the misconfigured cloud resources by linking a cloud issue to the underlying IaC template with an SCM source code link.
* Suppress false positives in IaC tests by applying [context from deployed infrastructure](iac+-code-to-cloud-capabilities/add-cloud-context-to-your-iac-tests.md).
* For Terraform - the same custom rule applies across the entire SDLC for all workflows (IaC to cloud).
* View an inventory of IaC and cloud resources generated from your IaC files via the [resources API](https://apidocs.snyk.io/?version=2023-09-20%7Ebeta#get-/orgs/-org\_id-/cloud/resources).

For a list of supported IaC languages and cloud providers, see [Supported IaC and cloud providers](supported-iac-languages-cloud-providers-and-cloud-resources/).
