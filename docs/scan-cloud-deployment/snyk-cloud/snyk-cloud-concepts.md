# Snyk Cloud concepts

Snyk Cloud has a number of unique concepts, separate from [Snyk core concepts](../../products/snyk-cloud/broken-reference/).

## Environments

A Snyk Cloud **Environment** is an organizing concept that equates to an Amazon Web Services (AWS) account, Azure subscription, or Google Cloud project.

Unlike a Snyk [Project](../../getting-started/introducing-snyk/key-concepts/projects.md), an environment contains scannable entities known as [resources](snyk-cloud-concepts.md#resources). Resources can be interrelated; one resource can be a child or sibling resource of another. Resources also have attributes that can be tested, and these attributes can be misconfigured, which generates Snyk Cloud issues. This makes environments and their resources different from Projects.

A Snyk Cloud Environment also includes integration settings for a cloud provider. For example, each environment can represent an integration with a different AWS account.

Use the `/cloud/environments` Snyk API endpoint to retrieve a list of all environments and optionally filter by attribute, such as name and scan status.

Currently supported providers:

* [Amazon Web Services](https://aws.amazon.com/)

## Resources

A **resource** is a cloud infrastructure entity such as an AWS S3 bucket, Identity & Access Management (IAM) role, or Virtual Private Cloud (VPC) flow log.

On each scan, Snyk records the configuration attributes of each resource in an environment.

You can use the `/cloud/resources` Snyk API endpoint to retrieve a list of all resources for an organization and optionally filter by an attribute such as environment ID, resource tag, or resource type.

For a list of supported resource types, see [Supported AWS resources for Snyk Cloud](supported-aws-resources-for-snyk-cloud.md).

## Rules

A security **rule** checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems. Snyk has a set of predefined rules that can be applied to both Snyk Cloud Environments and Snyk IaC Projects.

An example rule is “S3 bucket does not have all block public access options enabled.” Snyk can scan the configuration of an AWS S3 bucket to see if it fails the rule, and so is vulnerable to a data breach.

## Issues

A cloud **issue** represents a misconfiguration that can lead to a security problem. It is associated with a resource and a rule. For instance, an AWS S3 bucket can be tested against the rule “S3 bucket does not have all block public access options enabled.” If the bucket fails the rule, Snyk opens a cloud issue.

After Snyk creates an issue, Snyk keeps it open until the misconfiguration is fixed, at which point the issue is closed.

You can view your Organization's issues in the Snyk Web UI or through the `/cloud/issues` Snyk API endpoint. See [View cloud issues in the Snyk Web UI](snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md) and [View cloud issues in the API](snyk-cloud-issues/view-cloud-issues-in-the-api.md).
