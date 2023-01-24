# Remove a Snyk Cloud Environment

When you remove a Snyk Cloud Environment, Snyk removes all associated scans, issues, and records of resources. Your actual resources in the cloud provider are not affected.

## Web UI

In the Snyk Web UI, you can remove a cloud environment by navigating to your Organization's **Settings (cog icon) > Cloud environments**. See [View Snyk Cloud Environments](view-snyk-cloud-environments.md#remove-an-environment).

## API

To remove an environment using the Snyk API, send a request to the [`/cloud/environments`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#delete-/orgs/-org\_id-/cloud/environments/-environment\_id-) endpoint in the below format. You can find the environment ID using the method shown in [Find the environment ID](scan-a-snyk-cloud-environment.md#find-the-environment-id).

```
curl -X DELETE \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments/YOUR-ENVIRONMENT-ID?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-SERVICE-ACCOUNT-TOKEN'
```

There is no output if the command is successful.

## Remove the Snyk AWS IAM role

Removing an environment does not remove the Snyk Identity & Access Management (IAM) role. To fully remove Snyk's access to your Amazon Web Services (AWS) account, delete the IAM role using the same infrastructure as code tool that created it:

* **Terraform:** delete the role using the [terraform destroy](https://www.terraform.io/cli/commands/destroy) command.
* **CloudFormation:** delete the CloudFormation stack using the [AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-console-delete-stack.html) or the AWS CLI's [delete-stack command](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudformation/delete-stack.html).

## Remove the Google service account

Removing an environment does not remove the Google service account. To fully remove Snyk's access to your Google project, delete the Google service account using the [Google Cloud console](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-console) or [CLI](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-gcloud).
