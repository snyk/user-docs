# Configure integration for Amazon Elastic Container Registry \(ECR\)

Enable integration between one Amazon ECR registry and a Snyk organization, and start managing your image security. To integrate with multiple registries, create a unique organization for each one.

**Automated Process:**

You have the option of establishing cross-account access to enable Snyk's Amazon ECR integration as a 1-click deployment. This options is available as an official [AWS Quick Start](https://github.com/aws-quickstart/quickstart-snyk-security) and eliminates the need for manual configuration. 

![](../../../.gitbook/assets/quickstart-snyk-security-ecr.png)


You will need your Snyk **Organization ID** and AWS IAM [role ARN](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns) to complete the integration. The role ARN will be provided for you in the AWS CloudFormation Console's Output tab.

[![cloudformation-launch-stack.png](https://support.snyk.io/hc/article_attachments/360010120798/cloudformation-launch-stack.png)

](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Snyk-Security-ECR&templateURL=https://aws-quickstart.s3.amazonaws.com/quickstart-snyk-security/templates/snyk-ecr.yaml)

**Manual Process:**

To enable integration, you must first create a read-only AWS Identity and Access Management \(IAM\) role. The role delegates read-only access to all repositories in your registry for Snyk per organization by indicating the list of permitted Snyk-assigned organization IDs.

Thereafter, when integrating additional organizations, you can simply add the additional organization IDs as necessary.

Follow these steps to set up your integration:

* [Enable permissions to access Amazon ECR for the first time](https://support.snyk.io/hc/articles/360003947017#UUID-a8420931-6623-3dd4-faac-b5245b325ebf)
* [Add additional organizations to your AWS IAM role for Snyk authentication](https://support.snyk.io/hc/articles/360003947037#UUID-4e0116d9-dbb0-0a25-24de-406e3658c6ae)
* [Amazon ECR: configure your integration with Snyk](https://support.snyk.io/hc/articles/360003947057#UUID-2f4b05ad-e072-0883-b6e8-453f8a2702ea)

