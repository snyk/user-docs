# Snyk Code Local Engine

{% hint style="info" %}
Snyk Code Local Engine is currently in Private Beta. If you would like to set it up in your Organization, contact your Snyk account team.
{% endhint %}

Snyk Code Local Engine is a fully contained version of the Snyk Code Engine that allows users to avoid uploading their code to the internet.

This high-level architecture diagram shows the components and their interactions. Snyk scans a request for a Git repository through the Snyk Code Local Engine, which returns the results to Snyk.

<figure><img src="../../.gitbook/assets/Screen Shot 2021-11-11 at 2.36.41 PM.png" alt="Snyk Code Local Engine high-level architecture"><figcaption><p>Snyk Code Local Engine high-level architecture</p></figcaption></figure>

{% hint style="info" %}
When you use the Local Engine, only the scan is performed locally. Your scan results are uploaded to Snyk so you can view them on the Snyk Web UI.
{% endhint %}

## System requirements for Snyk Code Local Engine

The core requirements to deploy the Snyk Code Local Engine are:

* Kubernetes version 1.19.0 - 1.27.0:
  * _Recommended:_ a dedicated Kubernetes cluster
  *   Outbound HTTPS connections supporting WebSockets from the cluster to \*.snyk.io

      This connection is needed for all flows (CLI, IDE, SCM, and PR Checks)
  * Kubernetes – one of the following:
    * Managed public cloud Kubernetes service - EKS, AKS, GKE\
      \- or -
    * Unmanaged Kubernetes (a self-hosted cluster)
  * PR Checks and Snyk CLI support require network access to the Kubernetes cluster:
    * From your Kubernetes cluster to your Git and CI/CD tooling
    * From users running Snyk CLI to the Kubernetes cluster
* Helm version 3.8.0 or newer

## Resource requirements for Snyk Code Local Engine

The Snyk Code Local Engine is modular and can be customized to run specific integrations or everything at once. It can also scale as you prefer, with one or multiple nodes.

However**,** there is a minimum requirement for the core engine. Each of the Kubernetes nodes will require at least the following free resources to run it:

* 55 GB RAM
* 14 Core CPU
* 20GB Ephemeral Storage

The rest of the capabilities have different requirements, **i**ncluding the minimum single-node requirements for the core engine.

The minimum required resources for each flavor of the Snyk Code Local Engine are identified in the following list. Actual memory and storage consumption depend on the usage and the size of repositories scanned. The minimum total required resources can then be divided into multiple nodes.

* Local Engine CLI/IDE only
  * 165GB RAM
  * 60 Core CPU
  * 55GB Ephemeral Storage
* Local Engine SCM Only
  * 170GB RAM
  * 65 Core CPU
  * 65GB Ephemeral Storage
* Local Engine SCM and PR Check
  * 200GB RAM
  * 90 Core CPU
  * 160GB Ephemeral Storage
* Local Engine Full deployment (all features)
  * 220GB RAM
  * 100 Core CPU
  * 160GB Ephemeral Storage

Using the CLI and IDEs with the Snyk Code Local Engine

To use the Snyk CLI and IDEs with Snyk Code Local Engine, provide your Snyk account team with your Snyk Code Local Engine URL, the URL of the Snyk Code Local Engine running on your premises.\
\
After your CSM has configured the URL for your Organization, you can view it from **Settings** --> **Snyk Code:**

<figure><img src="../../.gitbook/assets/Snyk Code Local Engine settings showing Local Engine URL (1).png" alt="Snyk Code Local Engine settings Local Engine URL"><figcaption><p>Snyk Code Local Engine settings Local Engine URL</p></figcaption></figure>

## Configure and deploy the Local Engine

The instructions to configure and deploy the Local Engine in your environment are in the Readme file in the Local Engine installation package. For more information and assistance, contact your Snyk account team.

