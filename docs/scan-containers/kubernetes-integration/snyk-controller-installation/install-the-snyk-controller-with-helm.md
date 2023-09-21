# Install the Snyk controller with Helm (Azure and Google Cloud Platform)

{% hint style="danger" %}
Before following this installation page, review the [prerequisites page](../../../scan-applications/snyk-container/kubernetes-integration/use-the-snyk-controller/prerequisites-for-installing-the-snyk-controller.md).
{% endhint %}

To get vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk Controller onto your cluster. The Snyk Controller is published in [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor).

These installation pages cover:

* Snyk integration for most Kubernetes platforms
* Additional configuration steps for integration when using Amazon Elastic Container Registry (ECR) with your Amazon Elastic Kubernetes Service (EKS) clusters

## Steps to install the Snyk Controller with Helm

1. Access your Kubernetes environment and run the following command to add the Snyk Charts repository to Helm:

<pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
</strong></code></pre>

2. After the repository is added, create a unique namespace for the Snyk controller:

```
kubectl create namespace snyk-monitor
```

{% hint style="info" %}
Use a unique namespace to isolate the controller resources easily. This is generally good practice for Kubernetes applications. Notice the namespace is called snyk-monitor; you will need this later when configuring other resources.
{% endhint %}

Snyk monitor requires your Snyk **Integration ID**, **Service Account Token** and **dockercfg.json** file.

### Public Container Registry installation

Create a Kubernetes secret called `snyk-monitor`containing the Snyk **Integration ID** and **Service Account Token** by running the following command:

```
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

{% hint style="info" %}
The secret must be called `snyk-monitor` for the integration to work.
{% endhint %}

### Private Container Registry installation

For additional setup for Private Registries, see [Private Container Registry authentication](private-container-registry-authentication.md).

### Install the Snyk Helm chart

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster"
```

If you are running your own instance of Snyk, you must specify the API endpoint when installing the controller. In the following, provide the full hostname of your Snyk instance.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set integrationApi=https://<server>/kubernetes-upstream
```

{% hint style="info" %}
* Replace "**Production cluster"** with a name based on the cluster you are monitoring. You can use this label to find workloads in Snyk later.
* Note that **`/`** in cluster name is **disallowed**. Any **`/`** in cluster names will be removed.
* To avoid renaming the cluster on every update, you can use the existing option from Helm **--reuse-values**. When upgrading, Helm will reuse the values from the last release and merge in any overrides from the command line via **--set** and **-f**
{% endhint %}

### Integrating AKS with ACR using Managed Identities

1. When you are using AKS with user-managed identities to authorise access to ACR, and there are multiple identities that assign the `AcrPull` role to the VM scale set, you must also specify the Client ID of the desired user-managed identity to be used. This value must be set as an override, in `.Values.azureEnvVars`:

```yaml
azureEnvVars:
  - name: AZURE_CLIENT_ID
    value: "abcd1234-abcd-1234-abcd-1234abcd1234"
```

2. With the YAML above saved in `override.yaml`, run the following:

```bash
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  -f override.yaml
```

By default, this value is set to an empty string, and it will not be used as such.

{% hint style="info" %}
When using the system-managed identity with the `AcrPull` role assigned, setting this variable is not necessary.&#x20;
{% endhint %}

## Optional installation steps

To add any additional, optional Snyk Controller steps that fit your environment, refer to [Optional installation steps for Snyk Controller with Helm](optional-installation-steps-for-snyk-controller-with-helm.md)[.](optional-installation-steps-for-snyk-controller-with-helm.md)

## Updating an existing installation

If you are an existing customer and are updating your Snyk Controller:

1. Create a service account token as described on the [prerequisites page](../../../scan-applications/snyk-container/kubernetes-integration/use-the-snyk-controller/prerequisites-for-installing-the-snyk-controller.md). This token will be stored in the `snyk-monitor` secret.
2. Delete your existing `snyk-monitor` secret:

```shell
kubectl delete secret snyk-monitor -n snyk-monitor
```

3. Follow the instructions in the [Installation steps](install-the-snyk-controller-with-helm.md#installation-steps) section. To get the latest Helm chart version, make sure you run the following command:

<pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
</strong></code></pre>
