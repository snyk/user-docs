# Getting started with Snyk Cloud: AWS

Snyk Cloud scans the infrastructure configuration in your [Amazon Web Services (AWS)](https://aws.amazon.com/) provider account and detects misconfigurations that can lead to vulnerabilities.

You can can onboard an AWS account to Snyk using the following methods:

* [Snyk Web UI](snyk-cloud-for-aws-web-ui/)
* [Snyk API](snyk-cloud-for-aws-api/)

## Prerequisites

To start using Snyk Cloud, you need the following:

* A Snyk Business or Enterprise [plan](https://snyk.io/plans/)
* A new organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions)
* An organization-level [service account](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account) with an Org Admin role
* Access to an [AWS](https://aws.amazon.com/) account and associated credentials with permissions to create a read-only IAM role
* Access to the [Terraform CLI](https://www.terraform.io/downloads), [AWS CLI](https://aws.amazon.com/cli/), or [AWS Management Console](https://console.aws.amazon.com) to create the IAM role for Snyk via Terraform or AWS CloudFormation
  * If using Terraform or the AWS CLI, ensure you configure it with your AWS credentials. See instructions for [Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration) or the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)
* **API only:** An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
