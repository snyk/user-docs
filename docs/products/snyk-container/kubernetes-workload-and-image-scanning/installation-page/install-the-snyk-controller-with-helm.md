# Install the Snyk controller with Helm (Azure and Google Cloud Platform)

{% hint style="danger" %}
Before following this installation page, review the [prerequisite setting page](prerequisite-setting.md).
{% endhint %}

To get vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk controller onto your cluster. The Snyk controller is published in [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor).

This section covers:

* Snyk integration for most Kubernetes platforms
* Additional configuration steps for integration when using Amazon Elastic Container Registry (ECR) with your Amazon Elastic Kubernetes Service (EKS) clusters

The **installation steps** follow.

1.  Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

    ```
    helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
    ```
2.  Once the repository is added, create a unique namespace for the Snyk controller:

    ```
    kubectl create namespace snyk-monitor
    ```

    **Tip:** Use a unique namespace to isolate the controller resources easily. This is generally good practice for Kubernetes applications. Notice the namespace is called `snyk-monitor`; you’ll need this later when configuring other resources.
3.  Snyk monitor runs by using your Snyk **Integration ID** and a `dockercfg` file. \
    \
    <mark style="color:purple;">**Public Container Registry**</mark>\ <mark style="color:purple;"></mark>Create a Kubernetes secret called `snyk-monitor`(the secret must be called `snyk-monitor` for the integration to work) containing the Snyk **Integration ID** from the previous step by running the following command:

    ```
    kubectl create secret generic snyk-monitor -n snyk-monitor \
            --from-literal=dockercfg.json={} \
            --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
    ```

    \
    <mark style="color:purple;">**Private Container Registry**</mark>\ <mark style="color:purple;"></mark>Create a file named `dockercfg.json` and provide credentials to access those registries by creating a secret (which must be called `snyk-monitor`). The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually your credentials reside in `$HOME/.docker/config.json,` you can copy this information to the `dockercfg.json` file.\
    \
    If the `auth` entry is empty in your `$HOME/.docker/config.json` you can run this command and paste the output to `auth` entry in `dockercfg.json`

    ```
    echo -n 'username:password' | base64
    ```

    \
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
    If your cluster runs on **GKE** and you are using **GCR**, your`dockercfg.json` file should look like this:&#x20;

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
    If you are using **Artifactory Container Registry to host multiple private repositories,** your `dockercfg.json` file should look like this:&#x20;

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
    Create a secret with the file added:&#x20;

    ```
    kubectl create secret generic snyk-monitor \        
            -n snyk-monitor --from-file=dockercfg.json \
            --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
    ```


4. Install the Snyk Helm chart:

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
* To avoid renaming the cluster on every update, you can use the existing option from Helm **--reuse-values**. This means when upgrading, it will reuse the value of the last release and merge in any overrides from the command line via **--set** and **-f**. If **--reset-values** is specified, this is ignored.
{% endhint %}

For any additional Snyk Controller optional steps that fit your environment refer to [Optional installation steps for Snyk Controller with Helm](optional-installation-steps-for-snyk-controller-with-helm.md)[.](optional-installation-steps-for-snyk-controller-with-helm.md)
