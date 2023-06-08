# Introduction to Snyk Code Local Engine

### Overview

Snyk Code Local Engine is a fully contained version of the Snyk Code Engine, allowing users to avoid uploading their code to the internet.

This high-level architecture diagram shows the different components and their interactions.

![Snyk Code Local Engine Architecture](<../../../../.gitbook/assets/Screen Shot 2021-11-11 at 2.36.41 PM.png>)

{% hint style="info" %}
When you use the Local Engine, only the scan is performed locally. Your scan results are uploaded to Snyk, so you can view them in the Snyk Web UI.
{% endhint %}

### System Requirements

The core requirements to deploy the Snyk Code Local Engine are:

* **Kubernetes** – version 1.19 - 1.24:
  * _Recommended:_ a dedicated Kubernetes cluster
  * Outbound HTTPS connections supporting websockets from the cluster to \*.snyk.io
  * Kubernetes – one of the following:
    * Managed public cloud Kubernetes service - EKS, AKS, GKE\
      \- or -
    * Unmanaged Kubernetes (a self-hosted cluster)
  * PR Checks and Snyk CLI support requires network access:
    * From your Kubernetes cluster to your Git and CI/CD tooling
    * From users running Snyk CLI to the Kubernetes cluster
* **Helm** – version 3.8.0 or newer
* **3 Nodes** – each one with the following:
  * **Disk:** 500 GB (>300GB Ephemeral Storage)
  * **CPU and RAM:**
    * \[For AWS] instance type m5.8xlarge
    * \[For GCP] instance type e2-standard-32
    * 32 Core CPU with 120GB RAM
