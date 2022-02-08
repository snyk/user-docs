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

* Kubernetes version 1.16.0 - 1.21.5
* 3 nodes, **each** with following specs:
  * CPU 32 cores
  * RAM 120GB
  * Disk 500 GB (>300GB Ephemeral Storage)
* Helm 3.5.0
