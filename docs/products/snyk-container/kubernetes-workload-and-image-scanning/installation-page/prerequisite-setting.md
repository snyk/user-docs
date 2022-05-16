# Prerequisite setting

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Before installing Snyk Controller, please make sure you are following this prerequisite

* An administrator account for your Snyk organization.
* A minimum of 50 GB of storage must be available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster.&#x20;
* Your Kubernetes cluster needs to be able to communicate with Snyk outbound over HTTPS.
* [Enable the Kubernetes Integration ](../kubernetes-integration-overview/viewing-your-kubernetes-integration-settings.md)to get your **Integration ID.**
* Install [Helm](https://helm.sh/docs/intro/install/) locally&#x20;
