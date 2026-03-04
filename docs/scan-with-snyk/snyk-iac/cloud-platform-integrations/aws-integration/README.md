# AWS integration

Snyk integrates with your [Amazon Web Services (AWS)](https://aws.amazon.com/) account to find issues in your cloud configurations and to generate cloud context to help you prioritize your vulnerabilities.

You can onboard an AWS account to Snyk using the following methods:

* [Snyk Web UI](aws-integration-web-ui/)
* [Snyk API](aws-integration-api/)

To set up an AWS integration, you need the following:

* A Snyk Enterprise [plan](https://snyk.io/plans/)
* A new Snyk Organization, with appropriate feature flags assigned by your Snyk contact
* A Snyk Group Administrator or Organization Administrator [role](../../../../snyk-platform-administration/user-roles/pre-defined-roles.md)
* Access to an [AWS](https://aws.amazon.com/) account and associated credentials with permissions to create a read-only IAM role
* Access to the [Terraform CLI](https://www.terraform.io/downloads), [AWS CLI](https://aws.amazon.com/cli/), or [AWS Management Console](https://console.aws.amazon.com) to create the IAM role for Snyk via Terraform or AWS CloudFormation
* If you are using Terraform or the AWS CLI, ensure you configure it with your AWS credentials. See the instructions for [Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration) or the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)
* API only: An Organization-level [service account](../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role to use the Snyk API
* API only: An API client such as [curl](https://curl.se/), [HTTPie](https://httpie.io/), or [Postman](https://www.postman.com/)
