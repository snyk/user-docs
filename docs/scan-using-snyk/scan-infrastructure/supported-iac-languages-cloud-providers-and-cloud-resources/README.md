# Supported IaC languages, cloud providers, and cloud resources

{% hint style="info" %}
For environments supported with other Snyk products, see [Supported languages and frameworks](../../supported-languages-and-frameworks/).
{% endhint %}

Use Snyk IaC to find, view, and fix issues in infrastructure configuration files for Terraform, Kubernetes Manifests, AWS CloudFormation, Azure Resource Manager (ARM), and Helm Charts.

## Supported IaC languages

Language support for Current IaC vs. IaC+ is documented below:

<table data-header-hidden><thead><tr><th width="271"></th><th width="218.33333333333331"></th><th width="224"></th></tr></thead><tbody><tr><td></td><td><strong>Current IaC support</strong></td><td><strong>IaC+ support</strong></td></tr><tr><td><strong>Terraform (single file)</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Terraform (modules)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>Terraform (variables file)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>AWS CloudFormation</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>Yes (CLI only)</td><td>Yes (all workflows)</td></tr><tr><td><strong>Kubernetes Manifests</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Helm Charts</strong></td><td>Yes (SCM only)</td><td>Coming soon</td></tr></tbody></table>

## Supported cloud providers

Snyk IaC supports scanning the following cloud providers:

* [Amazon Web Services](../../../integrate-with-snyk/cloud-platforms-integrations/aws-integration/)
* [Azure](../../../integrate-with-snyk/cloud-platforms-integrations/azure-integration-for-cloud-configurations/)
* [Google Cloud](../../../integrate-with-snyk/cloud-platforms-integrations/google-cloud-integration/)

## Supported resources for cloud scans

For a list of supported cloud resource types, see the documentation for each cloud provider:

* [AWS](supported-aws-resources-for-cloud-context.md)
* [Azure](supported-azure-resources-for-cloud-context.md)
* [Google](supported-google-resources-for-cloud-context.md)
