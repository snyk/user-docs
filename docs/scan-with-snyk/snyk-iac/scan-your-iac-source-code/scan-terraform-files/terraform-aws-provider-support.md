---
description: Support and limitations for the Terraform AWS provider.
---

# Terraform AWS Provider support

{% hint style="success" %}
Version 4.0.0 of the AWS Terraform Provider introduced changes in how S3 services are defined. With v4.0, the definition of S3 services is now spread across several resource blocks within Terraform. If you defined an instance of an S3 bucket across multiple files, this update is a breaking change and may have negatively impacted your security results from Snyk IaC.
{% endhint %}

{% hint style="info" %}
From the Terraform documentation: "To help distribute the management of S3 bucket settings via independent resources, various arguments and attributes in the `aws_s3_bucket` resource have become read-only. Configurations dependent on these arguments should be updated to use the corresponding `aws_s3_bucket_*` resource."
{% endhint %}

To migrate to Terraform v4.0.0, you must refactor and re-import your S3 service definitions. Depending on how you choose to do this, it may limit your coverage of security findings.

See the [Terraform V4 upgrade guide from Hashicorp](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/version-4-upgrade) for more details

## Example of Terraform AWS Provider support

Consider the following example of defining an S3 bucket with an ACL using Terraform v3.x.x in a file named `s3.tf`

{% code title="s3.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... other configuration ...
  acl = "public-read-write"
}
```
{% endcode %}

The definition of the S3 bucket is in one resource block. If you scanned this file using `snyk iac test s3.tf` you would get a security finding for the permissive ACL settings.

With v4.0.0 of the Provider, certain S3 bucket properties are now defined in their own resources. Continuing the previous example, the ACL property has moved to its own resource, so the refactored Terraform looks like this.

{% code title="s3.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... other configuration ...
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.example.id
  acl    = "public-read-write"
}
```
{% endcode %}

If you scan this file using `snyk iac test s3.tf` you will get the same results as before for the permissive ACL settings.

## Limitations of support for Terraform AWS Provider

You may choose to structure your code differently, for example, by having the S3 bucket definition and its associated properties in separate Terraform files:

{% code title="s3_bucket.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... other configuration ...
}
```
{% endcode %}

{% code title="s3_acl.tf" %}
```hcl
resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.example.id
  acl    = "public-read-write"
}
```
{% endcode %}

If you scan these files, you will either not receive a security issue or receive a false positive for the permissive ACL. This is because Snyk cannot currently link the two resources together.

Snyk is working on adding support for this additional use case to the product.

An interim workaround is to define the resources in a single Terraform file.
