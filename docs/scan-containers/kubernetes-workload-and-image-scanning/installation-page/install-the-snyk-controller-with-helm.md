# Install the Snyk controller with Helm (Azure and Google Cloud Platform)

## Updating from V1 to V2 (existing installations only)

If you are an existing customer and are updating your Snyk controller to V2:

* Create a service account token as described in the [prerequisite setting page](prerequisite-setting.md). This token will be stored in the `snyk-monitor` secret.
* Delete your existing `snyk-monitor` secret:

```shell
kubectl delete secret snyk-monitor -n snyk-monitor
```

* Follow the instructions in the [Installation steps](install-the-snyk-controller-with-helm.md#installation-steps) section. In order to get the latest Helm chart version, make sure you run the following command:

<pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
</strong></code></pre>

## Prerequisites

{% hint style="danger" %}
Before following this installation page, review the [prerequisite setting page](prerequisite-setting.md).
{% endhint %}

To get vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk controller onto your cluster. The Snyk controller is published in [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor).

These pages cover:

* Snyk integration for most Kubernetes platforms
* Additional configuration steps for integration when using Amazon Elastic Container Registry (ECR) with your Amazon Elastic Kubernetes Service (EKS) clusters

## Installation steps

1.  Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

    <pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
    </strong></code></pre>
2.  Once the repository is added, create a unique namespace for the Snyk controller:

    ```
    kubectl create namespace snyk-monitor
    ```

    **Tip:** Use a unique namespace to isolate the controller resources easily. This is generally good practice for Kubernetes applications. Notice the namespace is called `snyk-monitor`; you will need this later when configuring other resources.
3. Snyk monitor requires your Snyk **Integration ID**, **Service Account Token** and **dockercfg.json** file.

### Public Container Registry installation

Create a Kubernetes secret called `snyk-monitor`(the secret must be called `snyk-monitor` for the integration to work) containing the Snyk **Integration ID** and **Service Account Token** by running the following command:

```
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

### Private Container Registry installation

Create a file named `dockercfg.json` and provide credentials to access those registries by creating a secret, which must be called `snyk-monitor`. The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually your credentials reside in `$HOME/.docker/config.json;` you can copy this information to the `dockercfg.json` file.\
\
If the `auth` entry is empty in your `$HOME/.docker/config.json` you can run this command and paste the output to `auth` entry in `dockercfg.json`:

```
echo -n 'username:password' | base64
```

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

\
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

\
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

\
Create a secret with the file added:

```
kubectl create secret generic snyk-monitor \        
        -n snyk-monitor --from-file=dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

Install the Snyk Helm chart:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster"
```

If you are running your own instance of Snyk you need to specify the API endpoint when installing the controller. In the following, provide the full hostname of your Snyk instance.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set integrationApi=https://<server>/kubernetes-upstream
```

{% hint style="info" %}
* Replace the name **Production cluster** with a name based on the cluster you are monitoring. You can use this label to find workloads in Snyk later.
* Please note that **`/`** in cluster name is **disallowed**. Any **`/`** in cluster names will be removed.
* To avoid renaming the cluster on every update, you can use the existing option from Helm **--reuse-values**. This means when upgrading, Helm will reuse the values from the last release and merge in any overrides from the command line via **--set** and **-f**
{% endhint %}

## Optional installation steps

For any additional, optional Snyk Controller steps that fit your environment refer to [Optional installation steps for Snyk Controller with Helm](optional-installation-steps-for-snyk-controller-with-helm.md)[.](optional-installation-steps-for-snyk-controller-with-helm.md)
