# Install the Snyk Controller

## Prerequisites for installing the Snyk Controller

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. For more details, see [pricing plans](https://snyk.io/plans/).
{% endhint %}

Before you install the Snyk Controller:

* You must have an administrator account for your Snyk Organization.
* You must have a minimum of 50 GB of storage available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster.
* You must have a minimum of the indicated RAM requirements to run on the cluster.

```
requests: cpu: "250m" memory: "400Mi"
limits: cpu: "1" memory: "2Gi"
```

* Your Kubernetes cluster must be able to communicate with Snyk outbound over HTTPS.
* [Enable the Kubernetes Integration ](../overview-of-the-kubernetes-integration/viewing-your-kubernetes-integration-settings.md)to get your **Integration ID.**
* Create a **Group** or **Organization** **service account token**. For more information, see [Service account](../../../../enterprise-setup/service-accounts/)s. There are different roles that allow the integration to publish data:
  * Group Admin
  * Org Admin
  * Org custom role with the permission **Publish Kubernetes Resources**
* Install [Helm](https://helm.sh/docs/intro/install/) locally.

The Snyk Controller defaults to using the US data center. If you are using Snyk deployed in [an alternative data center](../../../../more-info/data-residency-at-snyk.md), you must change the upstream endpoint `integrationApi` URL through an environment variable for your specific deployment:

* AU: [https://api.au.snyk.io/v2/kubernetes-upstream](https://api.au.snyk.io/v1/kubernetes-upstream)
* EU: [https://api.eu.snyk.io/v2/kubernetes-upstream](https://api.eu.snyk.io/v1/kubernetes-upstream)
* Helm example:

```
--set integrationApi=https://api.au.snyk.io/v2/kubernetes-upstream
```
