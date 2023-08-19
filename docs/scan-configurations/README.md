# Scan cloud configurations

With Snyk, you can secure cloud infrastructure configurations before and after it is deployed. [Snyk Infrastructure as Code (IaC)](snyk-infrastructure-as-code/) enables you to develop secure cloud infrastructure, and [Snyk Cloud](snyk-iac+/) helps you keep your cloud environment secure.

{% hint style="info" %}
**Feature availability**\
Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. Integrated IaC is currently in closed beta. Please reach out to your account team with any questions.
{% endhint %}

* Scan [IaC](snyk-infrastructure-as-code/) and [cloud](snyk-iac+/) resources for misconfigurations using a comprehensive set of security rules.
* [Fix Cloud issues](snyk-iac+/fix-cloud-issues-in-integrated-iac.md) directly in the IaC source code that was used to deploy the misconfigured Cloud resources, by linking a Cloud issue to the underlying IaC template via a SCM source code link.
* Receive [fix advice](snyk-infrastructure-as-code/getting-started-snyk-iac.md) so you can make changes directly to code, before applications reach production.
* Suppress false positives in IaC tests by applying [context from deployed infrastructure](snyk-iac+/adding-cloud-context-to-your-iac-tests.md).
* [Detect drift](snyk-infrastructure-as-code/detect-drift-and-manually-created-resources/) and manually created resources in your cloud.
* Inspect every [cloud resource's configuration](snyk-iac+/cloud-and-integrated-iac-issues/view-cloud-and-integragted-iac-issues-in-the-snyk-web-ui.md) at a given moment in time, and see the attributes that caused an issue.
* [Filter issues](snyk-iac+/cloud-and-integrated-iac-issues/) to target the most mission-critical resources.
* Review a [report of issues](../manage-issues/reports/next-gen-reporting/available-snyk-reports.md#cloud-compliance-issues-report) for an entire organization, organized by compliance standard.

For a list of supported IaC environments and cloud providers, see [Supported providers - IaC and Cloud](supported-iac-and-cloud-providers.md).
