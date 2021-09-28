# Securing ACR

### Welcome

This workshop builds upon the **Securing AKS** module which provided you with sample patterns and reference architectures for securing [Microsoft Azure](https://azure.microsoft.com/en-us/) workloads running on [Microsoft AKS](https://azure.microsoft.com/en-us/services/kubernetes-service/) with [Snyk](https://snyk.io/). Now, we will extend that learning and cover scanning of container images in [Microsoft Container Registry \(ACR\)](https://azure.microsoft.com/en-us/services/container-registry/). We will provide you with step-by-step examples and sample code that will walk you through deployment of supporting infrastructure, sample applications, and configuration of the various Snyk integrations to Microsoft Azure.

### Prerequisites

In order to complete the exercises in this workshop, you will need both a [Microsoft Azure](https://azure.microsoft.com/) & [Snyk](https://snyk.io/) account as well as successfully completed the **Securing AKS** module.

* [Create](https://azure.microsoft.com/en-us/free) a free Microsoft Azure account.
* [Create](https://snyk.io/login) a free Snyk account.
* Complete **Securing AKS**, or have a functional AKS cluster running.

### Architecture

At the conclusion of this module, you will have a functional instance of ACR storing a sample container image. The following reference architecture is indicative of what you will build:

![](../../../.gitbook/assets/snyk-acr.jpg)

