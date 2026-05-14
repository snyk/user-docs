# Authenticate to private container registries

If you are using private container registries, you must create a `dockercfg.json` file that contains the credentials to the registry. Then you must create a secret, which must be called `snyk-monitor`.

The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually, your credentials are in `$HOME/.docker/config.json`. However, the credentials must also be added to the `dockercfg.json` file. The Snyk Controller is not able to access these registries if the credentials are only stored in `$HOME/.docker/config.json`

The steps below explain how to authenticate to private container registries.

## Configure the dockercfg.json file

Create a file named `dockercfg.json`. Store your credentials in this file.

{% hint style="info" %}
Ensure the file containing your credentials is named `dockercfg.json`. This filename is required by the `snyk-monitor`.
{% endhint %}

{% hint style="info" %}
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

If you are using Nexus Repository\*\*,\*\* your `dockercfg.json` file must contain:

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

If you are using Artifactory Container Registry to host multiple private repositories\*\*,\*\* your `dockercfg.json` file must contain:

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
  	"auth": "<output from “echo -n "_json_key:$(cat <path to file containing raw key>)" | base64"}
  }
}

```

This method relies on creating a service account. See [Google Cloud service account key](https://cloud.google.com/artifact-registry/docs/docker/authentication#json-key). Ensure you have the raw key saved to a file.

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
In addition, for clusters running on AKS and using ACR, see [Entra ID Workload Identity service account](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html#service-account). It is possible that you are required to configure labels and annotations on the `snyk-monitor` ServiceAccount.
{% endhint %}

You can configure different credential helpers for different registries.

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
