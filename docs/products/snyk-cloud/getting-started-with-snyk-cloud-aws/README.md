# Getting started with Snyk Cloud: AWS

Snyk Cloud scans the infrastructure configuration in your cloud provider account and detects misconfigurations that can lead to vulnerabilities.

Currently, Snyk Cloud supports [Amazon Web Services (AWS)](https://aws.amazon.com/) accounts.

Before you can onboard an AWS account to Snyk Cloud, you need access to the AWS account and associated credentials with permissions to create a read-only Identity & Access Management (IAM) role.

To onboard an AWS account to Snyk Cloud:

1. [Download an infrastructure as code (IaC) template giving Snyk Cloud permissions to scan your account.](step-1-download-iam-role-iac-template.md)
2. [Create an AWS IAM role using the Terraform or AWS CloudFormation template you downloaded.](step-2-create-the-snyk-iam-role.md)
3. [Create and scan a Snyk Cloud Environment.](step-3-create-and-scan-a-snyk-cloud-environment.md)

You can now view the issues Snyk finds. See [Snyk Cloud issues](../snyk-cloud-issues/).

## Prerequisites

To start using Snyk Cloud, you need the following:

* A Snyk Business or Enterprise [plan](https://snyk.io/plans/)
* A new organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions)
* An organization-level [service account](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account) with an Org Admin role
* Access to an [AWS](https://aws.amazon.com/) account and associated credentials with permissions to create a read-only IAM role
* A supported browser:
  * [Chrome](https://www.google.com/chrome/) >=87
  * [Firefox](https://www.mozilla.org/en-US/firefox/browsers/) >=78
  * [Safari](https://www.apple.com/safari/) >=13
  * [Edge](https://www.microsoft.com/en-us/edge) >=88
* Access to the [Terraform CLI](https://www.terraform.io/downloads), [AWS CLI](https://aws.amazon.com/cli/), or [AWS Management Console](https://console.aws.amazon.com) to create the IAM role for Snyk via Terraform or AWS CloudFormation
  * If using Terraform or the AWS CLI, ensure you configure it with your AWS credentials. See instructions for [Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration) or the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)
* An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
