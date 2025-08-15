# Install the Snyk Controller

## Prerequisites for installing the Snyk Controller

{% hint style="info" %}
**Feature availability**\
Snyk Controller is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Before you install the Snyk Controller:

* You must have an administrator account for your Snyk Organization.
* You must have a minimum of 50 GB of storage available in the cluster as an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir).
* You must have a minimum of the RAM requirements indicated in the code that follows to run on the cluster.

```
requests: cpu: "250m" memory: "400Mi"
limits: cpu: "1" memory: "2Gi"
```

* Your Kubernetes cluster must have a Kubernetes `linux/amd64` worker node.
* Your Kubernetes cluster must be able to communicate with Snyk outbound over HTTPS.
* [Enable the Kubernetes Integration ](../overview-of-kubernetes-integration/enable-the-kubernetes-integration.md)to get your **Integration ID.**
* Create a **Group** or **Organization** **service account token**. For more information, see [Service accounts](../../../../implementation-and-setup/enterprise-setup/service-accounts/). Different roles allow the integration to publish data:
  * Group Admin
  * Org Admin
  * Org custom role with the permission **Publish Kubernetes Resources**.
* Install [Helm](https://helm.sh/docs/intro/install/) locally.

{% hint style="info" %}
The Snyk Controller Kubernetes integration is designed to be deployed into your cluster. It has not been tested on Serverless FaaS deployment options such as Fargate, Google Cloud Run, and Azure Container Instances. Deployment into serverless platforms may be unreliable and is not recommended or supported.
{% endhint %}
