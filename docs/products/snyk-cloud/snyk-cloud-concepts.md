# Snyk Cloud concepts

Snyk Cloud has a number of unique concepts, separate from [Snyk core concepts](broken-reference/).

## Environments

A Snyk Cloud **Environment** is an organizing concept that equates to an Amazon Web Services (AWS) account, Azure subscription, or Google Cloud project.

Unlike a Snyk [Project](broken-reference/), an environment contains scannable entities known as [resources](snyk-cloud-concepts.md#resources). Resources can be interrelated; one resource can be a child or sibling resource of another. Resources also have attributes that can be tested, and these attributes can be misconfigured, which generates Snyk Cloud issues. This makes environments and their resources different from Projects.

A Snyk Cloud Environment also includes integration settings for a cloud provider. For example, each environment can represent an integration with a different AWS account.

Use the [`/cloud/environments`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/cloud/environments) Snyk API endpoint to retrieve a list of all environments and optionally filter by attribute, such as name and scan status.

Currently supported providers:

* [Amazon Web Services](https://aws.amazon.com/)
* [Google Cloud](https://cloud.google.com/)

## Resources

A **resource** is a cloud infrastructure entity such as an AWS S3 bucket, Identity & Access Management (IAM) role, or Virtual Private Cloud (VPC) flow log.

On each scan, Snyk records the configuration attributes of each resource in an environment.

You can use the [`/cloud/resources`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/cloud/resources) Snyk API endpoint to retrieve a list of all resources for an organization and optionally filter by an attribute such as environment ID, resource ID, or resource type.

For a list of supported resource types, see the following:

* [Supported AWS resources](supported-aws-resources-for-snyk-cloud.md)
* [Supported Google resources](supported-google-resources-for-snyk-cloud.md)

## Rules

A security **rule** checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems. Snyk has a set of [predefined rules](https://snyk.io/security-rules/cloud) that can be applied to both Snyk Cloud Environments and Snyk IaC Projects.

An example rule is “S3 bucket does not have all block public access options enabled.” Snyk can scan the configuration of an AWS S3 bucket to see if it fails the rule, and so is vulnerable to a data breach.

## Issues

A cloud **issue** represents a misconfiguration that can lead to a security problem. It is associated with a resource and a rule. For instance, an AWS S3 bucket can be tested against the rule “S3 bucket does not have all block public access options enabled.” If the bucket fails the rule, Snyk opens a cloud issue.

After Snyk creates an issue, Snyk keeps it open until the misconfiguration is fixed, at which point the issue is closed.

You can view your Organization's issues in the Snyk Web UI. See [View cloud issues in the Snyk Web UI](snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md).

## Compliance standard <a href="#docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e" id="docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e"></a>

A compliance standard is a framework that establishes guidelines and controls for organizations to secure their IT systems and infrastructure. Compliance standards are “versioned,” with versions being released at various cadences. Examples: NIST 800-53 (vRev5), CIS AWS Foundations Benchmark (v1.4.0). Snyk provides a [Cloud Compliance Issues report](../../manage-issues/snyk-reports/reporting-beta-2022/available-snyk-reports.md#cloud-compliance-issues-report).

See [supported compliance standards](cloud-compliance.md#supported-compliance-standards).

## Compliance control <a href="#docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b" id="docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b"></a>

A compliance control is a specific recommendation or guideline from a compliance standard that prescribes how an organization should secure systems or infrastructure. Example: control 2.1.5 of CIS AWS Foundations Benchmark (v1.4.0) is “Ensure that S3 Buckets are configured with ‘Block public access (bucket settings)’”. To be compliant with this control, an organization would enable the “block public access” settings for all of their S3 buckets.

## Compliance mapping

Snyk “maps” security [rules](snyk-cloud-concepts.md#rules) to compliance controls, which means each rule is associated with one or more controls, and each control is associated with one or more rules.

For example, control 2.1.5 of CIS AWS Foundations Benchmark (v1.4.0) is “Ensure that S3 Buckets are configured with ‘Block public access (bucket settings)’” and it maps to the security rule[ SNYK-CC-00195](https://snyk.io/security-rules/cloud/SNYK-CC-00195/s3-bucket-does-not-have-all-block-public-access-options-enabled/), which is “S3 bucket does not have all block public access options enabled.”
