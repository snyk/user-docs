# Configure Google provider

## Authentication for Google provider

To use `iac describe`, set up credentials to make authenticated requests to your GCP project.

Because the `iac describe` command uses the Cloud Asset API, you must use a service account.

For information on setting up a service account, see the [GoogleCloud documentation](https://cloud.google.com/docs/authentication/production).

```
GOOGLE_APPLICATION_CREDENTIALS=your-creds.json \
  CLOUDSDK_CORE_PROJECT=my-project \
  snyk iac describe --to="gcp+tf"
```

You can use any `env var` from the [GoogleCloud sdk environment variables](https://cloud.google.com/sdk/docs/properties#setting_properties_via_environment_variables).

## Least privilege policy​ <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The `iac describe` command uses the [Google Asset API](https://console.cloud.google.com/apis/api/cloudasset.googleapis.com/overview) to enumerate resources on your account and the [Cloud Resource Manager API](https://console.cloud.google.com/marketplace/product/google/cloudresourcemanager.googleapis.com) to enumerate project IAM resources. Be sure to enable these APIs for the GCP project you are using as shown in the following screenshot.

<figure><img src="https://docs.driftctl.com/assets/images/enable_api-dffb8e57a0ce1c667527ede14b2728df.png" alt="Enable Cloud Asset API"><figcaption><p>Enable Cloud Asset API</p></figcaption></figure>

To enumerate resources, you need at least the role [Cloud Asset Viewer](https://cloud.google.com/iam/docs/understanding-roles#cloud-asset-roles).

## **Required roles​**

To use `iac describe` with deep mode, you need access to retrieve the details of a resource, and the **Cloud Asset Viewer** role is not enough. To be able to get the details, set up the basic role of [**Viewer**](https://cloud.google.com/iam/docs/understanding-roles#basic-definitions) on your project. To read your IAM policies you also need the role [iam.securityReviewer](https://cloud.google.com/iam/docs/understanding-roles#iam-roles) on your project.

```
# Mandatory role to allow describe to enumerate resources
roles/cloudasset.viewer

# Required for deep mode only
roles/viewer
```
