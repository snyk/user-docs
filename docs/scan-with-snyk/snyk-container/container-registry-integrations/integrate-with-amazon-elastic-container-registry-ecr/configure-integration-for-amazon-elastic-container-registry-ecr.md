# Configure integration for Amazon Elastic Container Registry (ECR)

{% hint style="warning" %}
When you connect to the ECR integration, ensure that the us-east-2 region is activated. This is required for the STS (Security Token Service) to work properly. For more information, see the [related support article](https://support.snyk.io/hc/en-us/articles/19883432587293-Connecting-to-ECR-Integration-gives-error-Could-not-connect-to-ECR-Please-ensure-your-credentials-are-correctly-configured).
{% endhint %}

This page explains how to enable integration between one Amazon ECR registry and a Snyk Organization and start managing your image security. To integrate with multiple registries, create a unique Organization for each one.

## **Automated integration process for ECR**

You can establish cross-account access to enable Snyk's Amazon ECR integration as a one-click deployment using an [AWS Quick Start](https://github.com/aws-quickstart/quickstart-snyk-security), This eliminates the need for manual configuration.

<figure><img src="../../../../.gitbook/assets/configure_integration_Amazon_ecr.png" alt=""><figcaption><p>AWS ECR and Snyk integration cross-account IAM role</p></figcaption></figure>

You must have your Snyk **Organization ID** and AWS IAM [role ARN](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference\_identifiers.html#identifiers-arns) to complete the integration. The role ARN is provided in the AWS CloudFormation Console's Output tab.

## **Manual integration process for ECR**

To enable integration, you must first create a read-only AWS Identity and Access Management (IAM) role. The role delegates read-only access to all repositories in your registry for Snyk per Organization by indicating the list of permitted Snyk-assigned Organization IDs.

After you create the IAM role, when integrating additional organizations, you can add the additional Organization IDs as needed.
