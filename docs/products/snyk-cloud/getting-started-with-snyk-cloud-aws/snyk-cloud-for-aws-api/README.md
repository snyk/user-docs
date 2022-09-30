# Snyk Cloud for AWS: API

Before you can onboard an AWS account to Snyk Cloud via the API, you need access to the AWS account and associated credentials with permissions to create a read-only Identity & Access Management (IAM) role. See all [Prerequisites](../#prerequisites).

To onboard an AWS account to Snyk Cloud via the API:

1. [Download an infrastructure as code (IaC) template](step-1-download-iam-role-iac-template.md): to give Snyk Cloud permissions to scan your account.
2. [Create an AWS IAM role](step-2-create-the-snyk-iam-role-api.md): using the template you downloaded.
3. [Create and scan a Snyk Cloud Environment.](step-3-create-and-scan-a-snyk-cloud-environment.md)

You can now view the issues Snyk finds. See [Snyk Cloud issues](../../snyk-cloud-issues/).
