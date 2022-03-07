# Google

## Authentication

To use `iac describe`, set up credentials to make authenticated requests to your Google Cloud Platform (GCP) project.

Since the `iac describe` command uses the cloud asset API, you **must use a service account**.

See the [Google Authenticating as a service account page](https://cloud.google.com/docs/authentication/production) for information on how to set up a service account.

```
GOOGLE_APPLICATION_CREDENTIALS=your-creds.json\
  CLOUDSDK_CORE_PROJECT=my-project\
  snyk iac describe --to=gcp+tf
```

You can use any environment variables from the [Google Cloud sdk environment variables](https://cloud.google.com/sdk/docs/properties#setting\_properties\_via\_environment\_variables).

## Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/google/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The `iac describe` command uses the [Google Asset API](https://console.cloud.google.com/apis/api/cloudasset.googleapis.com/overview) to efficiently enumerate resources on your account. It also uses the [Cloud Resource Manager API](https://console.cloud.google.com/marketplace/product/google/cloudresourcemanager.googleapis.com) to enumerate project IAM Resources. Be sure to enable these APIs for the GCP project you are using before you run the `iac describe` command.

![Enable the Cloud Asset APIi](https://docs.driftctl.com/assets/images/enable\_api-dffb8e57a0ce1c667527ede14b2728df.png)

To enumerate resources, you need at least the role [Cloud Asset Viewer](https://cloud.google.com/iam/docs/understanding-roles#cloud-asset-roles).

## **Deep mode**

If you want to use `snyk iac describe` with deep mode, the command also must be able to retrieve the details of a resource. The **Cloud Asset Viewer** is not enough. To get the details, set up the basic role [Viewer](https://cloud.google.com/iam/docs/understanding-roles#basic-definitions) on your project. To read your IAM policies you need the role [iam.securityReviewer](https://cloud.google.com/iam/docs/understanding-roles#iam-roles) on your project.

The following code sets up the **required roles**.[​](https://docs.driftctl.com/0.22.0/providers/google/authentication#required-roles)

```
# Mandatory role to allow describe to enumerate resources
roles/cloudasset.viewer

# Required for deep mode only
roles/viewer
```
