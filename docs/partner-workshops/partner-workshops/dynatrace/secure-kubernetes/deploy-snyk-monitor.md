# Deploy Snyk Monitor

{% hint style="warning" %}
Please review the **prerequisites** section of [product documentation page](https://support.snyk.io/hc/en-us/articles/360003916158-Install-the-Snyk-controller-with-Helm) before proceeding.
{% endhint %}

### Enable the integration

From the Snyk web console, navigate to `Integrations`. Search and select `Kubernetes`. Click `Connect` and copy the `Integration ID` to your clipboard. The `Integration ID` will be a UUID with a format similar to `abcd1234-abcd-1234-abcd-1234abcd1234`.

![](../../../.gitbook/assets/snyk_integrations_01.png)

### Install the Snyk controller

#### Helm chart

From the terminal, ensure that you have [helm installed ](https://helm.sh/docs/intro/install/)then, add the Snyk Charts repository to Helm with the following command:

```bash
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor/
```

If successful, you will see output similar to the following:

```text
"snyk-charts" has been added to your repositories
```

#### Namespace

Once added, we will need to create a unique namespace for the Snyk controller. Run the following command:

```bash
kubectl create namespace snyk-monitor
```

If successful, you will see output similar to the following:

```text
namespace/snyk-monitor created
```

#### Secret

The Snyk monitor runs by using your Snyk `Integration ID`, and using a `dockercfg` file. If you are not using any private registries, create a Kubernetes secret called `snyk-monitor` containing the Snyk `Integration ID` from the previous step and run the following command:

```bash
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=$IntegrationId
```

If successful, you will see output similar to the following:

```text
secret/snyk-monitor created
```

#### Deploy

Now, install the Snyk Helm chart to your AKS cluster:

```bash
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="myK8sCluster"
```

If successful, you will see output similar to the following:

```text
Release "snyk-monitor" does not exist. Installing it now.
NAME: snyk-monitor
NAMESPACE: snyk-monitor
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

#### Test

We can also validate our pod is running with the following command:

```bash
kubectl get pods --namespace snyk-monitor
```

You will want to see the `STATUS` display `Running` as in the following example output:

```text
NAME                            READY   STATUS    RESTARTS   AGE
snyk-monitor-544ff7ccd9-qkwj8   1/1     Running   0          4m47s
```

Note that Snyk Monitor will require outbound internet access.

### Import Projects

From the Kubernetes integration menu, select the **Add your Kubernetes workloads to Snyk** button.

![](../../../.gitbook/assets/snyk-k8s-config-01.png)

Then select the desired workloads and click **Add selected workloads**.

![](../../../.gitbook/assets/snyk-k8s-import-01.png)

Your project will be scanned and results available in the **Projects** menu.

