# Step 1: Download service account IaC template (Web UI)

Before you can create a Cloud Environment, you must download an infrastructure as code (IaC) template declaring a tightly-scoped Google service account that gives Snyk permission to scan the configuration of resources in your Google Project.

The template also enables a set of [Google service APIs](https://cloud.google.com/service-usage/docs/enabled-service) for your Google Cloud Project. This ensures that Snyk can use the necessary APIs to scan your Project's resources.

You will use this IaC template to provision the role in [Step 2: Create the Google service account](step-2-create-the-google-service-account-web-ui.md).

## Download the IaC template

1. In the [Snyk Web UI](https://app.snyk.io/), navigate to **Integrations** > **Cloud platforms**.
2. Select **Google Cloud**.
3. In the **Add Google Cloud Environment** modal, select the **Terraform** button to download a `snyk-permissions-google.tf` file:

<figure><img src="../../../../../.gitbook/assets/Bildschirmfoto 2023-07-18 um 12.16.54 (1) (1).png" alt=""><figcaption><p>The Add Google Cloud Environment modal</p></figcaption></figure>

You can now proceed to [Step 2: Create the Google service account.](step-2-create-the-google-service-account-web-ui.md)

{% hint style="info" %}
You can also add a cloud environment from **Settings** > **Cloud environments**. See [View Environments, Add a Cloud environment](../../../getting-started-with-cloud-scans/manage-cloud-environments/view-add-and-remove-environments.md).
{% endhint %}
