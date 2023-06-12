# Set up the Terraform Cloud integration for IaC

{% hint style="warning" %}
You must be an administrator of the Snyk organization to configure the Terraform Cloud integration.
{% endhint %}

Navigate to the dedicated Terraform Cloud integration settings page, under the **Integrations** page in the Snyk Web UI, then follow these steps to set up Terraform plan scanning:

In the Snyk Web UI, go to the Terraform Cloud integration settings page for your organization:\
`https://app.snyk.io/org/{YOUR-SNYK-ORG}/manage/integrations/terraform-cloud` to find the provided URL and HMAC Key.

![Snyk Integration settings for Terraform Cloud](../../../.gitbook/assets/terraform\_cloud\_02oct2022.png)

## Create the Snyk Run Task for Terraform Cloud

Now navigate to [Terraform Cloud](https://app.terraform.io) in the organization global settings:

![Terraform Cloud Settings](<../../../.gitbook/assets/image (126) (1) (2).png>)

1. Go to the run tasks settings for your TFC organization:\
   `https://app.terraform.io/app/{YOUR_TFC_ORG}/settings/tasks`
2. Create a new run task for Snyk with the URL and HMAC key values.\
   The HMAC key is mandatory for the Snyk integration to work, even though it is identified as optional on Terraform Cloud.

## Associate the Run Task to your Terraform Cloud Workspace

1. Navigate to your Terraform Cloud Workspace, enter the **Settings** \_\_ menu and chose **Run Tasks**_._
2. The run task you created is available under **Available Run Tasks**; click on **+** to associate it.
3. Choose the enforcement level (**Advisory** or **Mandatory**) and click **Create**.

Once your integration is set up, Snyk scans Terraform plans for each run triggered in your workspace.
