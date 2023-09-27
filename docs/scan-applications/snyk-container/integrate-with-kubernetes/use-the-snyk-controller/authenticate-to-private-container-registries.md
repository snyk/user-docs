# Authenticate to private container registries

If you are using private container registries, you must create a `dockercfg.json` file that contains the credentials to the registry. Then you must create a secret, which must be called `snyk-monitor`.&#x20;

The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually, your credentials are in `$HOME/.docker/config.json`. However, the credentials must also be added to the `dockercfg.json` file.  The Snyk Controller is not able to access these registries if the credentials are only stored in `$HOME/.docker/config.json`

The steps below explain how to authenticate to private container registries.

## Configure the dockercfg.json file

Create a file named `dockercfg.json`.  Store your credentials in this file.

{% hint style="info" %}
Ensure the file containing your credentials is named `dockercfg.json`. This filename is required by the `snyk-monitor`.
{% endhint %}

{% hint style="warning" %}
Ensure the formatting is correct, including new line characters and whitespace in the `dockercfg.json` file. Malformed files will result in authentication failures.
{% endhint %}

The locations where your cluster runs and where your registries run determine the combination of entries in your `dockercfg.json` file. The file can contain credentials for multiple registries.

If your credentials are already in `$HOME/.docker/config.json`, copy this information to the `dockercfg.json` file.

If the `auth` entry is empty in the `$HOME/.docker/config.json`, run the following command and paste the output to `auth` entry in `dockercfg.json`:

```
echo -n 'username:password' | base64
```

### Examples of dockercfg.json file configuration

#### For private registries other than Nexus

If your cluster does not run on `GKE`, or it runs on `GKE` and pulls images from other private registries, your`dockercfg.json` file must contain:

```json
{  
  "auths": {
    "gcr.io": {
      "auth": "BASE64-ENCODED-AUTH-DETAILS"
    },
    // Add other registries as necessary, for example:
    "<yourdomain>.azurecr.io": {
      "auth": "BASE64-ENCODED-AUTH-DETAILS"
    }
  }
}
```

#### For Nexus Repository

If you are using Nexus Repository**,** your `dockercfg.json` file must contain:

```json
{
  "auths": {
    "<registry>": {
        "auth": "BASE64-ENCODED-AUTH-DETAILS"
    },
  }
}
```

#### For Artifactory Container Registry

If you are using Artifactory Container Registry to host multiple private repositories**,** your `dockercfg.json` file must contain:

```json
{
  "auths": {
    "<registry1>": {
        "auth": "BASE64-ENCODED-AUTH-DETAILS"
       },
    "<registry2>": {
       "auth": "BASE64-ENCODED-AUTH-DETAILS"
    }
  }
}
```

#### For GKE using GCR

If your cluster runs on `GKE` and you are using `GCR`, your`dockercfg.json` file must contain:

```json
{
  "credHelpers": {
    "us.gcr.io": "gcloud",
    "asia.gcr.io": "gcloud",
    "marketplace.gcr.io": "gcloud",
    "gcr.io": "gcloud",
    "eu.gcr.io": "gcloud",
    "staging-k8s.gcr.io": "gcloud"
  }
}
```

#### For GKE using Google Artifact Registry (GAR)

If your cluster runs on `GKE` and you are using `GAR`, your`dockercfg.json` file must contain:

```json
{
  "auths": {
	"northamerica-northeast2-docker.pkg.dev": {
  	"auth": "<output from “echo -n _json_key_base64:BASE64-ENCODED-AUTH-DETAILS"}
  }
}

```

This method relies on creating a service account. See [Google Cloud service account key](https://cloud.google.com/artifact-registry/docs/docker/authentication#json-key). Ensure you follow the optional steps to base64 encode the file.

The `“auth”` line is generated with the following command, where the username is json\_key\_base64 and the password is the entire contents of the base64 json keyfile.

```sh
echo -n 'username:password' | base64 
```

For example, the output of this command is used in the `“auth”` line of the dockercfg.json

```sh
echo -n ‘json_key_base64:ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAic255ay1jeC1zZS1kZW1vIiwKICAicHJpdmF0ZV9rZXlfaWQiOiAiMDNhOTQyMWNhMDhkZTA0MzIyNTc4OTFjMDRiNGFjZmJlYTM0Y2U1ZCIsCiAgInByaXZhdGVfa2V5IjogIi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlJRXZRSUJBREFOQmdrcWhraUc5dzBCQVFFRkFBU0NCS2N3Z2dTakFnRUFBb0lCQVFEQkpxT2M0NzFIMjFIOVxuSjBKMi9zRTlJb2tmY1N0SjVKbEYyYTg3bGRZTHd6QkNmcW5UUk1lSGl6RGgrTXFCVWxjdE01ZVVFdFJXRmkwMFxud2NrVTN4dTNMVlRMekhnNWVMeGZ0MHVQZWtRRVBUZ3RQMW3lsk3tzVHEydms2T1VmTHI2ZHh0TXpUQzZtMWdzQVxuYS9rTm9mMmE1TkN6aHZ5d1NqU1E0dTdMWHBhWUhNY2ZUN0lTZmZJSEJVMi9OdDlWZUpYc3F0b3l3YXNsY3l5TVxub0dEV3JTL3BqblVWZzkxOXVvaUhBaE9MR0piMFczeWQwWWR2Z3ZPMGFrbzNzTWcxbkJiSTd2S1JEaE1aNDlETVxucWZPakNOZm56UUNKREdFOHB5akhUdWNxdWx0dkxxUjFBSUh5Sy9pNG53amNNdjB0MXR0d0xaOFBxSGpHd0ZPMlxuK3FsUnlHVTdBZ01CQUFFQ3lsv0RUFEV1RCQnJHU0lBZjQxUDJpY2I3enBtb2RLUlYvWTNYYkhRbGR5ZHQzaHNSSFxudGV5em1RZFZjTFE1dFFtNy9TQzVFOVRXaDNtUXlORnIzQk10L3VrRHNuMk51ajRZL2g4OXJNTjRsVi9zbElDc1xuUXhMM3o0OHhHMkdSTDBQcEUzTDUyOVg4TWp2dnBqa0VkVWlIY2ljUC8yd3JmcTkzR1VCa0NjSDZ2aFoyaWVDaVxuR1l1QzBRVGs1eXkwZTNnV05FMkpLdkk4WkI0Y2dwUHpKMnFhRHdWSXRHV3FOY3hwbWp6WjhNSS8wTXNaQS9IeFxuNTQ0MVpISkJKZXRzWVdxV2ZkWGV0cG1EN24vRVVzTFlQTmtRTnBodTE2Z2o5LzR3Wk9RaXhoTGpTWXJ5OVdqdFxuL2JtYlkyQ2xtRmVybFB5UjdMUElQeno1TXZqVGI4dHVFRjZNWG1nc0xRS0JnUURxMkNUcmVKZ05IUDB2YVprRVxuMXU0Tk5naTJxSkxZengrZHN2SENCVHgzREpvdWtXejd3Tzh2OXFjOWg2RkptalRzTFVUNUdZUVFDTjJ4V29zb1xubEJYcE1QY1g0VmNPTGlKNFRoSGJHYmV1RFhnZXZFYnVHOU90amtxVk1xdnQvaVBJL3pzRDhGL0ZxbmxSTkdxSlxuZjFmRjcwTWIxdGMwa0EvUWhTajNQK2lxdFFLQmdRRFNqUHZ4bEtXMVhjSm1KZEN1NWxHck1JTS8vdC9INTdpYVxuc0FLa0JaektBWVZhSHh1UUdxeUZ2RjRTSlRhelNEMUdRQ2RVcExZTUJhWGZpb25DTWVIYS9CWEsrQS9kVVd5OFxuUDFFSG9MNHJTTVZsMVl6WHRFWGh5b3FDcS9lSFRrcStQa2R4Z0pJSjBpcktvV0Y3ZDk3OTVqNmVtaStEeHYyeFxuN3VQOWNiK1dMd0tCZ0Y4WG9ITjhoRTBqQk40eTZ4UUxsNTdmMTAxbkd2Y1JmMkxTdDVQeG5OY3owaWF6R2ljaVxucTNlSGI1YTVtYlI4N1pzSWhabzhHNzZHYUlaTS9IWTA2RjVoUmx4MEVWVWJsemVSblNkVDFZMXp4TVRsUmU5YVxuY3k4ZW85S2dEd0F5WFBraGFCc2pOUlNMLzgzQzVMVENUSjlJVDZzeEpqa1JjR1hsMVgyd2NoelZBb0dCQUxIbFxubHhZRi80RGZLRnFBUnZNUC9SOEVUVkVyKzAzL1ZuVzBrM2FjbTEzK3JQcDVZQ09BdGhZRkV3S0gyTkRnRDQya1xudE5hS21KcE54MW01eHkyQ1VnOWhnTlJPaGJEOGxELzF5M1FEZDhwQW9UQ3FuMmE5bFhIeVhOZU5qd1lPdTQ1RVxuTnI4SzM5bFdidnRvSVdKZDVOWm56SzdiSFp4YzdJdURpYlRoZi92WEFvR0FWNkhWMmIwS0hFbjloQURQOXZITVxuZzg1Nm0yVnZUTGhjNmZaS3lDY28zWDYwZTNHaDJNLzUyZ3YxMTlWTGFvdWlFbzc1YWVYejQxNFlJUFcvL2VTZFxubVdibmtvWTZDVkFBdWFBOUh4dlp6Y2JQV1h4dElyK3lFcFcxMDNjMFN5bFErbW9PeSthRStyQm0yQld3VzBpYlxuWXU1SHBoWjFWczN6OThFaVlXWFk4aDQ9XG4tLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tXG4iLAogICJjbGllbnRfZW1haWwiOiAic3BlcmNpYmFsbGktc255a21vbml0b3JAc255ay1jeC1zZS1kZW1vLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwKICAiY2xpZW50X2lkIjogIjEwOTE4NDIxMDg1NTc0MTIzOTUxMyIsCiAgImF1dGhfdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwKICAidG9rZW5fdXJpIjogImh0dHBzOi8vb2F1dGgyLmdvb2dsZWFwaXMuY29tL3Rva2VuIiwKICAiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL29hdXRoMi92MS9jZXJ0cyIsCiAgImNsaWVudF94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3JvYm90L3YxL21ldGFkYXRhL3g1MDkvc3BlcmNpYmFsbGktc255a21vbml0b3IlNDBzbnlrLWN4LXNlLWRlbW8uaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJ1bml2ZXJzZV9kb21haW4iOiAiZ29vZ2xlYXBpcy5jb20iCn0K
```

#### For EKS using ECR

If your cluster runs on `EKS` and you are using `ECR`, add the following:

```json
{
  "credsStore": "ecr-login"
}
```

To use this credential helper for a specific `ECR` registry, create a credHelpers section with the URI of your ECR registry:

```json
{
  "credHelpers": {
    "public.ecr.aws": "ecr-login",
    "<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "ecr-login"
  }
} 
```

#### For AKS using ACR

If your cluster runs on `AKS` and you're using `ACR`, add the following:

```json
{
  "credHelpers": { 
    "myregistry.azurecr.io": "acr-env"
  }
}
```

{% hint style="info" %}
In addition, for clusters running on AKS and using ACR, see [Azure AD Workload Identity service account](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html#service-account). It is possible that you are required to configure labels and annotations on the `snyk-monito`r ServiceAccount.
{% endhint %}

You can configure different credential helpers for different registries.&#x20;

## Create the Kubernetes secret

Create the secret in Kubernetes by running the following command:

{% code overflow="wrap" %}
```sh
kubectl create secret generic snyk-monitor -n snyk-monitor \ 
        --from-file=./dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=aabb1212-abab-1212-dcba-4321abcd4321
```
{% endcode %}
