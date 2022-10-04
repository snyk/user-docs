# Manually import of Kubernetes workload project

{% hint style="info" %}
With the same integration id, you can import **many** clusters to **one** Snyk organization by giving clusters a unique cluster name during installation.

You can also import **one** cluster to **multiple** Snyk organizations in manual import, this can be done via [cloning an integration across your Snyk orgs](../../../integrations/managing-integrations/clone-an-integration-across-your-snyk-orgs.md), so they have the same integration id.

It is important to note that one Snyk organization has a limit of 25,000 projects; please consult your SCM for more information.
{% endhint %}

## **Prerequisites**

You must have an account with Snyk and be onboarded to your organization by an administrator.

In addition, the integration must be configured between Snyk and your Kubernetes environment per organization. To verify the integration is configured, ensure you have an [**Integration ID**](../kubernetes-integration-overview/viewing-your-kubernetes-integration-settings.md)\*\*\*\*

1. Deploy Snyk Controller in your Kubernetes cluster.
2. Deploy Kubernetes workloads that you want to be imported into the Kubernetes cluster

## Manually add workloads

Go to the **Projects page**, click **Add project** and select the **Kubernetes** option.

![](../../../.gitbook/assets/uuid-619a153d-6c77-f7dc-854c-ff77b3173191-en.png)

The import screen loads, similar to the one below, displaying all namespaces from the Kubernetes environment on the left and relevant namespace workloads on the right:

![](../../../.gitbook/assets/uuid-3a8568e0-b5a4-34af-d612-83466b206882-en.png)

We purposely ignore scanning certain namespaces which we believe are internal to Kubernetes (any namespace starting with _**kube-\***_), therefore you won't see those namespaces in the list, as well as the workloads they contain.\
The full list of ignored namespaces can be found [here](https://github.com/snyk/kubernetes-monitor/blob/master/src/supervisor/watchers/internal-namespaces.ts). This can be configured by adding the following to helm when setting up the snyk-monitor:

```
      --set excludedNamespaces={kube-node-lease,local-path-storage,some_namespace}
```

* Select one or multiple namespaces from the left side and for each namespace, select one or multiple workloads to import from the right side.

![Select\_namespace.gif](../../../.gitbook/assets/uuid-27db0a60-f18d-5ab0-9215-5a81e467f013-en.gif)

* When ready, click **Add selected workloads** from the top right of the screen. When the import completes, the Projects page loads and all workloads that you’ve imported appear, with a unique Kubernetes icon:

![Kubernetes icon](../../../.gitbook/assets/uuid-24e0b69a-01c3-9434-9dac-9b44864bd269-en.png)

Each item is named according to its Kubernetes metadata as follows: **\<namespace>/\<kind>/\<name>**.

You can filter for Kubernetes projects only:

![](<../../../.gitbook/assets/image (5) (1) (1) (1).png>)
