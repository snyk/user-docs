# Install the Snyk Controller

## Prerequisites for installing the Snyk Controller

{% hint style="warning" %}
**Release status**

Snyk Controller is available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
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
* [Enable the Kubernetes Integration ](../overview-of-the-kubernetes-integration/enable-the-kubernetes-integration.md)to get your **Integration ID.**
* Create a **Group** or **Organization** **service account token**. For more information, see [Service account](../../../../enterprise-configuration/service-accounts/)s. There are different roles that allow the integration to publish data:
  * Group Admin
  * Org Admin
  * Org custom role with the permission **Publish Kubernetes Resources**
* Install [Helm](https://helm.sh/docs/intro/install/) locally.

{% hint style="info" %}
The Snyk-monitor Kubernetes integration is designed to be deployed into your cluster. It has not been tested on Serverless FaaS deployment options such as Fargate, Google Cloud Run, and Azure Container Instances. Deployment into serverless platforms may be unreliable and is not recommended or supported.
{% endhint %}

The Snyk Controller defaults to using the US data center. If you are using Snyk deployed in [an alternative data center](../../../../working-with-snyk/regional-hosting-and-data-residency.md), you must change the upstream endpoint `integrationApi` URL through an environment variable for your specific deployment:

* AU: [https://api.au.snyk.io/v2/kubernetes-upstream](https://api.au.snyk.io/v1/kubernetes-upstream)
* EU: [https://api.eu.snyk.io/v2/kubernetes-upstream](https://api.eu.snyk.io/v1/kubernetes-upstream)
* Helm example:

```
--set integrationApi=https://api.au.snyk.io/v2/kubernetes-upstream
```
