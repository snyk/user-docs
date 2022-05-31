# Prerequisite setting

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Before installing Snyk Controller, please make sure you have the prerequisites:

* You must have an administrator account for your Snyk organization.
* You must have a minimum of 50 GB of storage must be available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster.
* You must have a minimum of the indicated [RAM requirements](https://github.com/snyk/kubernetes-monitor/blob/staging/snyk-monitor/values.yaml#L82-L89) to run on the cluster.
* Your Kubernetes cluster must be able to communicate with Snyk outbound over HTTPS.
* [Enable the Kubernetes Integration ](../kubernetes-integration-overview/viewing-your-kubernetes-integration-settings.md)to get your **Integration ID**.
* Install [Helm](https://helm.sh/docs/intro/install/) locally&#x20;

![Copy the Integration ID](../../../../.gitbook/assets/CopyIntegration.gif)
