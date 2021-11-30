# Deploy Snyk

You can scan Kubernetes workloads by deploying the `snyk-monitor` into your cluster. This is published as an official [helm chart](https://artifacthub.io/packages/helm/snyk/snyk-monitor) and we will be basing these steps on that deployment option. To learn more about Snyk's Kubernetes integration, please visit our [documentation pages](https://docs.snyk.io/products/snyk-container/image-scanning-library/kubernetes-workload-and-image-scanning/kubernetes-integration-overview). For convenience, we will cover the steps at a high level here.&#x20;

#### Step 1

Create the namespace:

```bash
kubectl create namespace snyk-monitor
```

#### Step 2

Create the secret:

```bash
kubectl create secret generic snyk-monitor -n snyk-monitor --from-literal=dockercfg.json={} --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
```

{% hint style="info" %}
Locate your **Integration** ID from the [Snyk Integrations page ](https://app.snyk.io/org/YOUR-ORGANIZATION-NAME/manage/integrations/kubernetes)and copy it.
{% endhint %}

#### Step 3

Add the Helm repo:

```bash
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor/ --force-update
```

#### Step 4

Install the chart:

```bash
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor --namespace snyk-monitor --set clusterName="Production cluster"
```

{% hint style="info" %}
Replace "Production cluster" with the name of your cluster.
{% endhint %}

Now, we wait for the app to be ready and can check the status by running the following command:

```bash
kubectl get pods -n snyk-monitor
```

Like the previous example, we want a ready status to be displayed.
