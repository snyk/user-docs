# Configure integration for Amazon Elastic Container Registry (ECR)

{% hint style="warning" %}
When you connect to the ECR integration, ensure that the us-east-2 region is activated. This is required for the STS (Security Token Service) to work properly. For more information, see the [related support article](https://support.snyk.io/s/article/Connecting-to-ECR-Integration-gives-error-Could-not-connect-to-ECR-Please-ensure-your-credentials-are-correctly-configured).
{% endhint %}

This page explains how to enable integration between one Amazon ECR registry and a Snyk Organization and start managing your image security. To integrate with multiple registries, create a unique Organization for each one.

To enable integration, you must first create a read-only AWS Identity and Access Management (IAM) role. The role delegates read-only access to all repositories in your registry for Snyk per Organization by indicating the list of permitted Snyk-assigned Organization IDs.

After you create the IAM role, when integrating additional organizations, you can add the additional Organization IDs as needed.

Additionally, after you create the IAM role, allow a few minutes for AWS to update the role on their servers before continuing:

1. From AWS, copy the **Role ARN** key that appears at the top of the **Summary** section of the **Role** area.
2. Log in to your Snyk account.
3. Navigate to **Integrations** and click the Amazon ECR option.\
   The Amazon ECR configuration page in the Settings area loads.
4. Enter credentials as follows:
   1. **AWS Region**—use the format `region-part-#`, for example, `eu-west-3`.\
      You must enter the default region as configured for your AWS account for your repositories and images to be available for import.
   2. **Role ARN**—copy from your AWS account in the format `arn:aws:iam:::role/`.
5. Click **Save**

An example follows:

```
   arn:aws:iam::881001789406:role/TestSnykIntegration_role
```

Snyk tests the connection values, and the page reloads, now displaying Amazon ECR integration details as you entered them. A confirmation message that the details were saved also appears in green at the top of the screen.

If the connection to AWS fails, a notification appears under the **Connected to Amazon ECR** section.
