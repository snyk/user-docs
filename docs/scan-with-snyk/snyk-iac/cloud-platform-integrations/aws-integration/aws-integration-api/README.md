# AWS Integration: API

Before you can onboard an AWS account to Snyk via the API, you need access to the AWS account and associated credentials with permissions to create a read-only Identity and Access Management (IAM) role. See the prerequisites on the [AWS integration](../) page.

The steps follow to onboard an AWS account via the API:

1. [Download an infrastructure as code (IaC) template](step-1-download-iam-role-iac-template-api.md): to give Snyk permissions to scan your account.
2. [Create an AWS IAM role](step-2-create-the-snyk-iam-role-api.md): using the template you downloaded.
3. [Create and scan a Cloud Environment.](step-3-create-and-scan-a-cloud-environment-api.md)

When you have completed the steps, you will be able to do the following:

* View the cloud configuration issues Snyk finds. See [Manage cloud issues](../../../getting-started-with-cloud-scans/manage-cloud-issues/).
* Prioritize your vulnerabilities with cloud context.
