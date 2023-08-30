# Private container registry authentication

If you are using any **private container registries**, you must create a `dockercfg.json` file containing credentials to the registry, and then create a secret, which must be called `snyk-monitor`.&#x20;

The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually, your credentials can be found in `$HOME/.docker/config.json`. However, the credentials must also be added to the `dockercfg.json` file.&#x20;

The Snyk Controller will not be able to access these registries if the credentials are only stored in `$HOME/.docker/config.json`

**The steps to authenticate for private container registries follow.**

## Configure the dockercfg.json file

Create a file named `dockercfg.json`.  Store your credentials in this file.

Be sure the file with your credentials is named `dockercfg.json`. This filename is required by the `snyk-monitor`.

Pay careful attention to formatting including new line characters and whitespace in the `dockercfg.json` file. Malformed files will result in authentication failures.

Where your cluster runs and where your registries run determine the combination of entries in your `dockercfg.json` file. The file can contain credentials for multiple registries.

If credentials already reside in `$HOME/.docker/config.json`, you can copy this information to the `dockercfg.json` file.

If the `auth` entry is empty in your `$HOME/.docker/config.json`, you can run this command and paste the output to `auth` entry in `dockercfg.json`:

```
echo -n 'username:password' | base64
```

## Examples of dockercfg.json file

If your cluster does not run on **GKE**, or it runs on **GKE** and pulls images from **other private registries**, your`dockercfg.json` file should look like this:

```
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

If you are using **Artifactory Container Registry to host multiple private repositories,** your `dockercfg.json` file should look like this:

```
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

If your cluster runs on **GKE** and you are using **GCR**, your`dockercfg.json` file should look like this:

```
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

If your cluster runs on **EKS** and you are using **ECR**, add the following:

```
{
  "credsStore": "ecr-login"
}
```

To use this credential helper for a specific **ECR** registry, create a credHelpers section with the URI of your ECR registry:

```
{
  "credHelpers": {
    "public.ecr.aws": "ecr-login",
    "<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "ecr-login"
  }
}
```

If your cluster runs on **AKS** and you're using **ACR**, add the following:

```
{
  "credHelpers": { 
    "myregistry.azurecr.io": "acr-env"
  }
}
```

{% hint style="info" %}
Additionally, for clusters running on AKS and using ACR, see [Azure AD Workload Identity service account](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html#service-account).

You may need to configure labels and annotations on the snyk-monitor ServiceAccount
{% endhint %}

You can configure different credential helpers for different registries.&#x20;

## Create the secret

Create the secret in Kubernetes by running the following command:

{% code overflow="wrap" %}
```sh
kubectl create secret generic snyk-monitor -n snyk-monitor \ 
        --from-file=./dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=aabb1212-abab-1212-dcba-4321abcd4321
```
{% endcode %}
