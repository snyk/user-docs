# Supported IaC languages and cloud providers

{% hint style="info" %}
For environments supported with other Snyk products, see: [Snyk Open Source - supported languages and package managers](../../scan-applications/snyk-open-source/snyk-open-source-supported-languages-and-package-managers.md), [Snyk Code - Supported languages and frameworks](../../scan-applications/snyk-code/snyk-code-language-and-framework-support.md), and [Snyk Container - Supported operating system distributions](../../scan-applications/snyk-container/how-snyk-container-works/supported-operating-system-distributions.md).
{% endhint %}

Use Snyk IaC to find, view, and fix issues in infrastructure configuration files for Terraform, Kubernetes Manifests, AWS CloudFormation, Azure Resource Manager (ARM), and Helm Charts.

Language support for Current IaC vs. IaC+ is documented below:

<table data-header-hidden><thead><tr><th width="271"></th><th width="218.33333333333331"></th><th width="224"></th></tr></thead><tbody><tr><td></td><td><strong>Current IaC support</strong></td><td><strong>IaC+ support</strong></td></tr><tr><td><strong>Terraform (single file)</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Terraform (modules)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>Terraform (variables file)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>AWS CloudFormation</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>Yes (CLI only)</td><td>Yes (all workflows)</td></tr><tr><td><strong>Kubernetes Manifests</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Helm Charts</strong></td><td>Yes (SCM only)</td><td>Coming soon</td></tr></tbody></table>

Snyk IaC supports scanning the following cloud providers:

* [Amazon Web Services](../../integrations/cloud-platforms-integrations/aws-integration/)
* [Azure](../../integrations/cloud-platforms-integrations/azure-integration-for-cloud-configurations/)
* [Google Cloud](../../integrations/cloud-platforms-integrations/google-cloud-integration/)
