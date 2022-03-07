# Google

## Authentication

To use describe, we need credentials to make authenticated requests to your GCP project.

**SERVICE ACCOUNT**

Since snyk iac describe use Cloud Asset API, using a service account is **mandatory**.

Please refer to [official documentation](https://cloud.google.com/docs/authentication/production) to setup a proper service account.

```
GOOGLE_APPLICATION_CREDENTIALS=your-creds.json\
  CLOUDSDK_CORE_PROJECT=my-project\
  snyk iac describe --to=gcp+tf
```

You can use any env var from [google cloud sdk environment variable](https://cloud.google.com/sdk/docs/properties#setting\_properties\_via\_environment\_variables).

### Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/google/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

Describe uses [Google Asset API](https://console.cloud.google.com/apis/api/cloudasset.googleapis.com/overview) to enumerate efficiently resources on your account. It also uses [Cloud Resource Manager API](https://console.cloud.google.com/marketplace/product/google/cloudresourcemanager.googleapis.com) to enumerate project IAM Resources. Be sure to have enabled these APIs for the GCP project you are using.

![enable\_api](https://docs.driftctl.com/assets/images/enable\_api-dffb8e57a0ce1c667527ede14b2728df.png)

To be able to enumerate resources, you need at least the role [Cloud Asset Viewer](https://cloud.google.com/iam/docs/understanding-roles#cloud-asset-roles).

**DEEP MODE**

If you want to use describe with deep mode, it will also need to retrieve resource's details and the **Cloud Asset Viewer** will not be enough. If you want to be able to get the details you should set up the basic role [Viewer](https://cloud.google.com/iam/docs/understanding-roles#basic-definitions) on your project. To read your IAM policies you will also need role [iam.securityReviewer](https://cloud.google.com/iam/docs/understanding-roles#iam-roles) on your project.

#### Required roles[​](https://docs.driftctl.com/0.22.0/providers/google/authentication#required-roles) <a href="#required-roles" id="required-roles"></a>

```
# Mandatory role to allow describe to enumerate resources
roles/cloudasset.viewer

# Required for deep mode only
roles/viewer
```
