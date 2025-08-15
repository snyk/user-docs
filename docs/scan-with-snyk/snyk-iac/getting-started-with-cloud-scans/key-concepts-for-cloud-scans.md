# Key concepts for cloud scans

Cloud scans have a number of unique concepts that are different from Snyk core concepts, such as [Environments](key-concepts-for-cloud-scans.md#environments) and [Resources](key-concepts-for-cloud-scans.md#resources).

## Environments

A Snyk environment is an organizing concept that equates to an Amazon Web Services (AWS) account, Azure subscription, or Google Cloud project.

Unlike a Snyk [Project](../../../snyk-platform-administration/snyk-projects/#project), an environment contains scannable entities known as [resources](key-concepts-for-cloud-scans.md#resources). Resources can be interrelated; one resource can be a child or sibling resource of another. Resources also have attributes that can be tested, and these attributes can be misconfigured, which generates Issues. This makes environments and their resources different from Projects.

A Snyk environment also includes integration settings for a cloud provider. For example, each environment can represent an integration with a different AWS account.

Use the [`/cloud/environments`](https://apidocs.snyk.io/#get-/orgs/-org_id-/cloud/environments) Snyk REST API endpoint to retrieve a list of all environments and optionally filter by attribute, such as name and scan status.

The following cloud providers are supported:

* [Amazon Web Services](https://aws.amazon.com/)
* [Microsoft Azure](https://azure.microsoft.com/en-us/)
* [Google Cloud](https://cloud.google.com/)

## Resources

A **resource** is a cloud infrastructure entity such as an AWS S3 bucket; Identity and Access Management (IAM) role; or Virtual Private Cloud (VPC) flow log.

On each scan, Snyk records the configuration attributes of each resource in an environment.

You can use the [`/cloud/resources`](https://apidocs.snyk.io/?version=2023-05-29%7Ebeta#get-/orgs/-org_id-/cloud/resources) Snyk REST API endpoint to retrieve a list of all resources for an Organization and optionally filter by an attribute such as environment ID, resource ID, or resource type.

For a list of supported resource types for cloud environments, see the following:

* [Supported AWS resources](../detect-manually-created-resources/supported-resources/aws-resources.md)
* [Supported Azure resources](../detect-manually-created-resources/supported-resources/azure-resources.md)
* [Supported Google resources](../detect-manually-created-resources/supported-resources/google-resources.md)

## Rules

A security rule checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems.&#x20;

An example rule is "S3 bucket does not have all block public access options enabled". Snyk can scan the configuration of an AWS S3 bucket to see if it fails the rule and so is vulnerable to a data breach.

## Issues

An issue represents a misconfiguration that can lead to a security problem. It is associated with a resource and a rule. For instance, an AWS S3 bucket can be tested against the rule “S3 bucket does not have all block public access options enabled.” If the bucket fails the rule, Snyk opens a cloud issue.

After Snyk creates an issue, Snyk keeps it open until the misconfiguration is fixed, at which point the issue is closed.

You can view your Organization's issues in the Snyk Web UI. See [View cloud issues in the Snyk Web UI](manage-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md).

## Compliance standard <a href="#docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e" id="docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e"></a>

A compliance standard is a framework that establishes guidelines and controls for Organizations to secure their IT systems and infrastructure. Compliance standards are versioned, with versions being released at various cadences. Examples: NIST 800-53 (vRev5), CIS AWS Foundations Benchmark (v1.4.0). Snyk provides a [Cloud Compliance Issues report](../../../manage-risk/reporting/available-snyk-reports.md#cloud-compliance-issues-report).

For more information, see [supported compliance standards](../view-cloud-compliance-reporting.md#supported-compliance-standards).

## Compliance control <a href="#docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b" id="docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b"></a>

A compliance control is a specific recommendation or guideline from a compliance standard that prescribes how an Organization should secure systems or infrastructure. Example: control 2.1.5 of CIS AWS Foundations Benchmark (v1.4.0) is “Ensure that S3 Buckets are configured with ‘Block public access (bucket settings)’”. To be compliant with this control, an Organization would enable the “block public access” settings for all of their S3 buckets.

## Compliance mapping

Snyk “maps” security [rules](key-concepts-for-cloud-scans.md#rules) to compliance controls, which means each rule is associated with one or more controls and each control is associated with one or more rules.

For example, control 2.1.5 of CIS AWS Foundations Benchmark (v1.4.0) is “Ensure that S3 Buckets are configured with ‘Block public access (bucket settings’” and it maps to the security rule [SNYK-CC-00195](https://security.snyk.io/rules/cloud/SNYK-CC-00195), which is “S3 bucket does not have all block public access options enabled.”
