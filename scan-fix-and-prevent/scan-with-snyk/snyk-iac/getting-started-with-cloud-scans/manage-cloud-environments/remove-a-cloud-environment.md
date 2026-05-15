# Remove a cloud environment

When you remove an environment, Snyk removes all associated scans, issues, and records of resources. For cloud environments, your actual resources in the cloud provider are not affected.

## Web UI

In the Snyk Web UI, you can remove an environment by navigating to your Organization **Settings** > **Cloud environments**. See [Remove an IaC+ or cloud environment](view-add-and-remove-environments.md#remove-a-cloud-environment).

## API

To remove an environment using the Snyk API, send a request to the [Delete environment](https://apidocs.snyk.io/#delete-/orgs/-org_id-/cloud/environments/-environment_id-) endpoint in the following format. You can find the environment ID using the method shown on the page [Find an environment ID](find-an-environment-id.md).

```
curl -X DELETE \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments/YOUR-ENVIRONMENT-ID?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-SERVICE-ACCOUNT-TOKEN'
```

There is no output if the command is successful.

## Remove the Snyk AWS IAM role

Removing an environment does not remove the Snyk Identity and Access Management (IAM) role. To fully remove Snyk's access to your Amazon Web Services (AWS) account, delete the IAM role using the same infrastructure as code tool that created it:

* **Terraform:** delete the role using the [terraform destroy](https://www.terraform.io/cli/commands/destroy) command.
* **CloudFormation:** delete the CloudFormation stack using the [AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-console-delete-stack.html) or the AWS CLI [delete-stack command](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudformation/delete-stack.html).

## Remove the Google service account

Removing an environment does not remove the Google service account. To fully remove Snyk's access to your Google project, delete the Google service account using the [Google Cloud console](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-console) or [CLI](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-gcloud).

## Remove the Azure app registration, federated identity credential, and service principal

Removing an environment does not remove the Azure app registration, federated identity credential, or service principal. To fully remove Snyk's access to your Azure subscription, delete the infrastructure according to the method you used to create it:

* **Terraform:** use the [terraform destroy](https://www.terraform.io/cli/commands/destroy) command.
* **Azure CLI Bash script:** first, delete the Reader role assignment for the app registration (see docs: [Azure CLI](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-remove#azure-cli) or [Azure Portal](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-remove#azure-portal)), then delete the Azure app registration (see docs: [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/ad/app?view=azure-cli-latest#az-ad-app-delete) or [Azure Portal](https://learn.microsoft.com/en-us/azure/active-directory/develop/howto-remove-app#remove-an-application-authored-by-you-or-your-organization)).
