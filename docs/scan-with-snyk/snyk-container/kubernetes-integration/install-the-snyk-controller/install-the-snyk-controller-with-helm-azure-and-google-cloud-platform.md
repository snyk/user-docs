# Install the Snyk Controller with Helm (Azure and Google Cloud Platform)

{% hint style="info" %}
Ensure you have reviewed the [prerequisites for installing the Snyk Controller](./#prerequisites-for-installing-the-snyk-controller).
{% endhint %}

To receive vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk Controller onto your cluster. The Snyk Controller is published in [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor).

The installation steps cover:

* Snyk integration for most Kubernetes platforms
* Additional configuration steps for integration when using Amazon Elastic Container Registry (ECR) with your Amazon Elastic Kubernetes Service (EKS) clusters.

## Installing the Snyk Controller with Helm

To install the Snyk Controller with Helm, follow these mandatory steps:

1. Access your Kubernetes environment. Run the following command to add the Snyk Charts repository to Helm:

<pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
</strong></code></pre>

2. After the repository is added, create a unique namespace for the Snyk Controller:

```
kubectl create namespace snyk-monitor
```

{% hint style="info" %}
As a good practice for Kubernetes applications, use a unique namespace to isolate the controller resources easily.&#x20;

Ensure you remember the namespace `snyk-monitor`. You will use it when configuring other resources.
{% endhint %}

Snyk monitor requires:

* Snyk Integration ID
* Service Account Token
* dockercfg.json file.

### Installing the Snyk Controller to scan images from a public container registry

To install the Snyk Controller to scan images from a public container registry, you must create a Kubernetes secret called `snyk-monitor` containing the Snyk Integration ID and the service account token.

To do this, run the following command:

```
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

{% hint style="info" %}
For a successful integration, the secret must be called `snyk-monitor`.
{% endhint %}

### Install the Snyk Helm chart

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster"
```

If you are running your own instance of Snyk, you must specify the API endpoint when installing the controller.

In the following command, provide the full hostname of your Snyk instance.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set integrationApi=https://<server>/kubernetes-upstream
```

* Replace `"Production cluster"` with a name based on the cluster you are monitoring. You can use this label to find workloads in Snyk later.
* Using `/` (slash) in the cluster name is not allowed. Any `/`in the cluster names are removed.
* To avoid renaming the cluster on every update, you can use the existing option from Helm `--reuse-values`. When upgrading, Helm reuses the values from the last release and merges them in any overrides from the command line using `--set` and `-f`

### Integrate AKS with ACR using Managed Identities

To do this:

1. When you use AKS with user-managed identities to authorize access to ACR, and there are multiple identities that assign the `AcrPull` role to the VM scale set, you must also specify the Client ID of the desired user-managed identity to be used. This value must be set as an override, in `.Values.azureEnvVars`:

```yaml
azureEnvVars:
  - name: AZURE_CLIENT_ID
    value: "abcd1234-abcd-1234-abcd-1234abcd1234"
```

2. With the YAML above saved in `override.yaml`, run the following command:

```bash
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  -f override.yaml
```

By default, this value is set to an empty string, and it is not used as such.

{% hint style="info" %}
When using the system-managed identity with the `AcrPull` role assigned, setting this variable is not necessary.&#x20;
{% endhint %}

## Update an existing installation

If you are an existing customer and are updating your Snyk Controller:

1. Create a service account token. For more information, see [Prerequisites for installing the Snyk Controller](./#prerequisites-for-installing-the-snyk-controller). This token is stored in the `snyk-monitor` secret.
2. Delete your existing `snyk-monitor` secret:

```shell
kubectl delete secret snyk-monitor -n snyk-monitor
```

3. Follow the [mandatory installation steps](install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md#mandatory-installation-steps-for-the-snyk-controller-with-helm). To get the latest Helm chart version, run the following command:

<pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
</strong></code></pre>
