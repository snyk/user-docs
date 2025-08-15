# Snyk Code Local Engine

{% include "../../.gitbook/includes/release-status-snyk-code-local-engine.md" %}

Snyk Code Local Engine (SCLE) is a fully contained version of the Snyk Code Engine that allows you to avoid uploading your code to the internet. When you use the Local Engine, only the scan is performed locally. Your scan results are uploaded to Snyk so you can view them on the Snyk Web UI.

This high-level architecture diagram shows the components and their interactions. Snyk scans a request for a Git repository through the Snyk Code Local Engine, which returns the results to Snyk.

<figure><img src="../../.gitbook/assets/Screen Shot 2021-11-11 at 2.36.41 PM.png" alt="Snyk Code Local Engine high-level architecture"><figcaption><p>Snyk Code Local Engine high-level architecture</p></figcaption></figure>

## System requirements for Snyk Code Local Engine

The core requirements to deploy the Snyk Code Local Engine are:

* Kubernetes version 1.21.0 - 1.28.0:
  * Recommended: a dedicated Kubernetes cluster
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
* Environment with x86 CPUs

### Resource requirements for Snyk Code Local Engine

The Snyk Code Local Engine is modular and can be customized to run specific integrations or everything at once. It can also scale as you prefer, with one or multiple nodes.

Each of the Kubernetes nodes will require at least the following free resources to run it:

* 55 GB RAM
* 14 Core CPU
* 50GB Ephemeral Storage

The size of the minimum node will depend on what your environment requires, but these are the minimum free resources required.&#x20;

The total required resources for each flavor of the Snyk Code Local Engine are identified in the following list. Actual memory and storage consumption depend on the usage and the size of repositories scanned. The minimum total required resources can then be divided into multiple nodes.

<table><thead><tr><th width="236">Deployment options</th><th width="263.3333333333333">Resources required</th><th>Use cases</th></tr></thead><tbody><tr><td>CLI </td><td><ul><li>165GB RAM</li><li>60 Core CPU</li><li>55GB Ephemeral Storage</li></ul></td><td><ul><li>Run the SCLE in the pipeline</li></ul></td></tr><tr><td>IDE </td><td><ul><li>165GB RAM</li><li>60 Core CPU</li><li>55GB Ephemeral Storage</li></ul></td><td><ul><li>Developers using it on their IDE</li></ul></td></tr><tr><td>SCM </td><td><ul><li>170GB RAM</li><li>65 Core CPU</li><li>65GB Ephemeral Storage</li></ul></td><td><ul><li>Import repositories for monitoring purposes</li></ul></td></tr><tr><td>SCM and PR Checks </td><td><ul><li>200GB RAM</li><li>90 Core CPU</li><li>160GB Ephemeral Storage</li></ul></td><td><ul><li>Import repositories for monitoring purposes</li><li>Scan every PR for new vulnerabilities</li></ul></td></tr><tr><td>SCM, PR Checks and CLI </td><td><ul><li>220GB RAM</li><li>100 Core CPU</li><li>160GB Ephemeral Storage</li></ul></td><td><ul><li>Run the SCLE in the pipeline</li><li>Import repositories for monitoring purposes</li><li>Scan every PR for new vulnerabilities</li></ul></td></tr><tr><td>Full deployment (all features)</td><td><ul><li>330GB RAM</li><li>140 Core CPU</li><li>220GB Ephemeral Storage</li></ul></td><td><ul><li>All of the above</li></ul></td></tr></tbody></table>

## Using the CLI and IDEs with the Snyk Code Local Engine

To use the Snyk CLI and IDEs with Snyk Code Local Engine, provide your Snyk account team with your Snyk Code Local Engine URL, the URL of the Snyk Code Local Engine running on your premises.\
\
After your CSM has configured the URL for your Organization, you can view it from **Settings** > **Snyk Code:**

<figure><img src="../../.gitbook/assets/Snyk Code Local Engine settings showing Local Engine URL.png" alt="Snyk Code Local Engine settings Local Engine URL"><figcaption><p>Snyk Code Local Engine settings Local Engine URL</p></figcaption></figure>

## Configure and deploy the Local Engine

The instructions to configure and deploy the Local Engine in your environment are in the Readme file in the Local Engine installation package. For more information and assistance, contact your Snyk account team.
