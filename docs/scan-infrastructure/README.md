# Scan infrastructure

{% hint style="info" %}
**Feature availability**\
Snyk IaC+ is a new version of Snyk IaC that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. IaC+ is currently in closed beta. Reach out to your account team if you would like access.
{% endhint %}

With Snyk, you can secure cloud infrastructure configurations before and after deployment. [Snyk Infrastructure as Code (IaC)](snyk-infrastructure-as-code/) enables you to develop secure cloud infrastructure, and [Snyk IaC+](snyk-iac+/) helps you keep your cloud environment secure. Snyk supports the following actions:

* Snyk IaC helps developers write secure configurations for [HashiCorp Terraform](snyk-infrastructure-as-code/scan-terraform-files/), [AWS CloudFormation](snyk-infrastructure-as-code/scan-cloudformation-files/), [Kubernetes](snyk-infrastructure-as-code/scan-kubernetes-configuration-files/), and [Azure Resource Manager (ARM)](snyk-infrastructure-as-code/scan-arm-configuration-files.md).
* IaC+ is a new version of Snyk IaC that has cloud context to secure cloud configurations across the entire SDLC, from code to deployed AWS, Azure, and Google Cloud environments. IaC+ has the Snyk cloud context interfaces, workflows, policy engine, and data model.
* [Fix Cloud issues](snyk-iac+/fix-cloud-issues-in-integrated-iac.md) directly in the IaC source code that was used to deploy the misconfigured cloud resources by linking a cloud issue to the underlying IaC template with an SCM source code link.
* Receive [fix advice](getting-started-with-snyk-iac/) so you can make changes directly to code, before applications reach production.
* Suppress false positives in IaC tests by applying [context from deployed infrastructure](snyk-iac+/add-cloud-context-to-your-iac-tests.md).
* [Detect drift](snyk-iac+/detect-drift-and-manually-created-resources/) and manually created resources in your cloud.
* Inspect the configuration of every [cloud resource](key-concepts-in-iac+/cloud-and-integrated-iac-issues/view-cloud-and-integragted-iac-issues-in-the-snyk-web-ui.md) at a given moment in time and see the attributes that caused an issue.
* [Filter issues](key-concepts-in-iac+/cloud-and-integrated-iac-issues/) to target the most mission-critical resources.
* Review a [report of issues](../manage-issues/reports/next-gen-reporting/available-snyk-reports.md#cloud-compliance-issues-report) for an entire Organization, organized by compliance standard.

For a list of supported IaC environments and cloud providers. See [Supported IaC and cloud providers](snyk-infrastructure-as-code/supported-iac-and-cloud-providers.md).
