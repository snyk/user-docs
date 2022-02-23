# Introduction

### Overview

Snyk Code Local Engine is a fully contained version of the Snyk Code Engine, allowing users to avoid uploading their code to the internet.

This high-level architecture diagram shows the different components and their interactions.

![Snyk Code Local Engine Architecture](<../../../../.gitbook/assets/Screen Shot 2021-11-11 at 2.36.41 PM.png>)

{% hint style="info" %}
Only the scanning is done locally. Your scan results are uploaded to Snyk, for you to view in the regular Snyk hosted UI.
{% endhint %}

### System Requirements

The core requirements to deploy the Snyk Code Local Engine are:

* **Kubernetes**
  * Version 1.16.0 - 1.21.5
* **Helm**
  * 3.5.0
* **3 Nodes each with**
  * **Disk:** 500 GB (>300GB Ephemeral Storage)
  * **CPU and RAM:**
    * AWS instance type m5.8xlarge OR
    * GCP instance type e2-standard-32 OR
    * 32 core CPU with 120GB RAM
