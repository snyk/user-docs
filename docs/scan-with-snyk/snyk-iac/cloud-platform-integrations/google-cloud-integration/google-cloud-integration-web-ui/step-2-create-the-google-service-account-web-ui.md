# Step 2: Create the Google service account (Web UI)

{% hint style="info" %}
**Recap**\
You have downloaded the Terraform template declaring the [Google service account](https://cloud.google.com/iam/docs/service-accounts) for Snyk. Now you need to provision the infrastructure.
{% endhint %}

The process to create the Google service account is the same whether you are using the [Snyk Web UI](step-1-download-service-account-iac-template-web-ui.md) or [Snyk API](../google-cloud-integration-api/step-1-download-service-account-iac-template-api.md) to onboard your Google Project.

To scan a Google Cloud Project, Snyk takes the permissions of a tightly-scoped Google service account that allows Snyk to scan the configuration of your Project resources.

The service account you create is granted the following read-only Identity and Access Management (IAM) roles:

* [Security Reviewer](https://cloud.google.com/iam/docs/understanding-roles#iam.securityReviewer)
* [Viewer](https://cloud.google.com/iam/docs/understanding-roles)

Snyk Cloud's service account is granted the [Service Account Token Creator](https://cloud.google.com/iam/docs/understanding-roles#iam.serviceAccountTokenCreator) IAM role to enable it to generate short-lived credentials for your service account.

Additionally, Snyk has a mechanism in place to lock a service account to the Organization that onboards it. This is a security feature to ensure that nobody can guess a service account name and onboard it into a separate Organization to see those resources.

## Set Google Cloud Project ID

Snyk scans the Google Cloud Project specified by the `project_id` [variable](https://www.terraform.io/language/values/variables) in the Terraform template. You must set the variable's value using one of the following methods:

* Set the `project_id` variable directly in the Terraform template. On line 4 of the template, change the default value of the `project_id` variable to your Project ID:

```
default = "your-project-id"
```

* Set the `project_id` variable when you apply the Terraform. In the following section, [Apply  Terraform](step-2-create-the-google-service-account-web-ui.md#apply-terraform), you will apply Terraform to create the Google service account. At that time, you can use Terraform's [-var](https://www.terraform.io/language/values/variables#variables-on-the-command-line) option to set the `project_id` variable to your Project ID:

```
terraform apply -var="project_id=your-project-id"
```

* Use the `GOOGLE_PROJECT` environment variable. See Terraform's [documentation.](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#full-reference)

## Apply Terraform

{% hint style="info" %}
Before you use the [Terraform CLI](https://www.terraform.io/downloads), ensure you [configure it to use your Google Cloud credentials](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started).
{% endhint %}

To provision the Google service account using Terraform:

1. In your terminal, navigate to the directory containing your `.tf` file (named `snyk-permissions-google.tf` if it has been downloaded from the Web UI).
2. Using the Terraform CLI, initialize the Terraform Project:

```
terraform init
```

3\. Review and apply the Terraform plan:

```
terraform apply
```

4\. Enter `yes` when Terraform asks if you want to perform the actions.

Terraform then creates the Google service account. When it is finished, you will see the following output:

```
Apply complete! Resources: 22 added, 0 changed, 0 destroyed.

Outputs:

service_account_email = "snyk-cloud-mt-us-abcd1234@my-project.iam.gserviceaccount.com"
identity_provider = "https://iam.googleapis.com/projects/12345567/locations/global/workloadIdentityPools/workload-identity-123456/providers/identity-provider-123456"
```

Copy the service account email and identity provider for use in the next step.

## What's next?

The next step is to create and scan the Cloud Environment. See [Step 3: Create and scan a Cloud Environment for Google (Web UI)](step-3-create-and-scan-a-cloud-environment-for-google-web-ui.md) or [Step 3: Create and scan a Cloud Environment for Google (API)](../google-cloud-integration-api/step-3-create-and-scan-a-cloud-environment-for-google-api.md).
