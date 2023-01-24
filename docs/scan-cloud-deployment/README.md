# Scan cloud deployments

With Snyk, you can secure cloud infrastructure before and after it is deployed. [Snyk Infrastructure as Code (IaC)](snyk-infrastructure-as-code/) enables you to develop secure cloud infrastructure, and [Snyk Cloud](snyk-cloud/) helps you keep your cloud environment secure.

* Scan [IaC](snyk-infrastructure-as-code/) and [cloud](snyk-cloud/) resources for misconfigurations using a comprehensive set of security rules.
* Receive [fix advice](snyk-infrastructure-as-code/getting-started-snyk-iac.md) so you can make changes directly to code, before applications reach production.
* Suppress false positives in IaC tests by applying [context from deployed infrastructure](snyk-infrastructure-as-code/integrated-infrastructure-as-code/adding-cloud-context-to-your-iac-test.md).
* [Detect drift](snyk-infrastructure-as-code/detect-drift-and-manually-created-resources/) and manually created resources in your cloud.
* Inspect every [cloud resource's configuration](snyk-cloud/snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md) at a given moment in time, and see the attributes that caused an issue.
* [Filter issues](snyk-cloud/snyk-cloud-issues/) to target the most mission-critical resources.
* Review a [report of issues](../manage-issues/snyk-reports/reporting-beta-2022/available-snyk-reports.md#cloud-compliance-issues-report) for an entire organization, organized by compliance standard.

Snyk IaC supports the following configurations:

* [HashiCorp Terraform](snyk-infrastructure-as-code/scan-terraform-files/)
* [AWS CloudFormation](snyk-infrastructure-as-code/scan-cloudformation-files/)
* [Kubernetes](snyk-infrastructure-as-code/scan-kubernetes-configuration-files/)
* [Azure Resource Manager](snyk-infrastructure-as-code/scan-arm-configuration-files.md)

Snyk Cloud supports the following providers:

* [Amazon Web Services](snyk-cloud/getting-started-with-snyk-cloud-aws/)
* [Google](snyk-cloud/getting-started-with-snyk-cloud-google/)
