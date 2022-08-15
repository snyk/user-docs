# Remove a Snyk Cloud Environment

When you remove a Snyk Cloud Environment, Snyk removes all associated scans, issues, and records of resources. Your actual resources in the cloud provider are not affected.

To remove an environment, send a request to the Snyk API's `/cloud/environments` endpoint in the below format. You can find the environment ID using the method shown in [Find the environment ID](scan-a-snyk-cloud-environment.md#find-the-environment-id).

```
curl -X DELETE \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments/YOUR-ENVIRONMENT-ID?version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-SERVICE-ACCOUNT-TOKEN'
```

There is no output if the command is successful.

## Remove the Snyk IAM role

Removing an environment does not remove the Snyk IAM role. To fully remove Snyk's access to your cloud provider account, delete the Amazon Web Services (AWS) Identity & Access Management (IAM) role using the same infrastructure as code tool that created it:

* **Terraform:** delete the role using the [terraform destroy](https://www.terraform.io/cli/commands/destroy) command.
* **CloudFormation:** delete the CloudFormation stack using the [AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-console-delete-stack.html) or the AWS CLI's [delete-stack command](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudformation/delete-stack.html).
