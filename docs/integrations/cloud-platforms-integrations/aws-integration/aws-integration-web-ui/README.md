# AWS Integration: Web UI

Before you can onboard an AWS account via the Web UI, you need access to the AWS account and associated credentials with permissions to create a read-only Identity and Access Management (IAM) role. See the prerequisites on the [AWS integration](../) page.

The steps follow to onboard an AWS account via the Web UI:

1. [Download an infrastructure as code (IaC) template](step-1-download-iam-role-iac-template-web-ui.md) to give Snyk permissions to scan your account.
2. [Create an AWS IAM role](step-2-create-the-snyk-iam-role.md) using the template you downloaded.
3. [Create and scan a Cloud Environment.](step-3-create-and-scan-a-snyk-cloud-environment-web-ui.md)

When you have completed the steps, you will be able to do the following:

* View the cloud configuration issues Snyk finds. See [Cloud and IaC+ issues](../../../../scan-using-snyk/scan-infrastructure/introduction-to-iac+/cloud-and-integrated-iac-issues/).
* Prioritize your vulnerabilities with cloud context.
