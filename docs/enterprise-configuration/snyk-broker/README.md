# Snyk Broker

{% hint style="info" %}
**Feature availability**

Snyk Broker is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, allowing for access by [snyk.io](http://snyk.io/) to your code to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For more information, see [How Snyk Broker works](./#how-snyk-broker-works).

## How to download and install Snyk Broker

Snyk Broker is hosted on [GitHub](https://github.com/snyk/broker) and published as a set of Docker images for specific integrations. Snyk provides a [Helm Chart](https://github.com/snyk/snyk-broker-helm) to deploy Snyk Broker if you are using Kubernetes. To deploy Broker, you must install and configure an integration.

You can install and configure using [Helm](install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md) or [Docker](install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md). You can install using Docker to run the Snyk Broker Client or run npm `install snyk-broker`. Snyk recommends using Helm as the simplest way to deploy Snyk Broker.&#x20;

## **Integrations with Snyk Broker**

Install each type of integration and configure using environment variables, as explained for [Docker](install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md) and [Helm](install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md).

Types of integrations supported with Broker are:

* Your Source Code Management (SCM) system ([GitHub](install-and-configure-snyk-broker/github-prerequisites-to-install-and-configure-broker/), [GitHub Enterprise](install-and-configure-snyk-broker/github-enterprise-prerequisites-to-install-and-configure-broker/), [BitBucket Server/Data Center](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-to-install-and-configure-broker/), [GitLab](install-and-configure-snyk-broker/gitlab-prerequisites-to-install-and-configure-broker/), [Azure Repos](install-and-configure-snyk-broker/azure-repos-prerequisites-to-install-and-configure-broker/))
  * SCM that is not internet-reachable
  * Publicly-accessible SCM, allowing you to view and control Snyk activity for increased data security
* Your on-premise [Jira](install-and-configure-snyk-broker/jira-prerequisites-to-install-and-configure-broker/), [JFrog Artifactory](install-and-configure-snyk-broker/artifactory-repository-install-and-configure-broker/), or [Nexus](install-and-configure-snyk-broker/nexus-repository-prerequisites-to-install-and-configure-broker/) installation
* Network-restricted [container registries](snyk-broker-container-registry-agent/)
* [Infrastructure as code (IaC) configuration files](snyk-broker-infrastructure-as-code-detection/) on private Git-based repositories

You can also use [derived Docker images](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/snyk-broker-set-up-examples/derived-docker-images-for-broker-client-integrations-and-container-registry-agent) for each integration and the Container Registry Agent.

For information about advanced configuration as needed for your installation, see[ Advanced configuration for Snyk Broker Docker installation ](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation)and [Advanced setup for Helm Chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation).

## **Using Snyk Broker to scan your code**

To use **Snyk Open Source** with Snyk Broker, you need only the Broker Server and Broker Client components. The Broker Client is published as a set of Docker images, each configured for a specific Git service. Configure each type of integration using environment variables following the links in the section [Integrations with Snyk Broker](./#integrations-with-snyk-broker).

To scan other types of code with Snyk Broker, you must add a component or configurations and add parameters to the Broker Client setup:

* **Snyk Code** – add the [Code Agent](snyk-broker-code-agent/) component to enable Snyk Code analysis of repositories in SCMs that are integrated through Snyk Broker. You can also grant Broker access to perform a [Git clone of your repository](install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md) by adding an environment variable: `ACCEPT_CODE=true.`
* **Snyk Container** – add the [Container Registry Agent](snyk-broker-container-registry-agent/) to enable the connection to network-restricted container registries and the analysis of container images. There are instructions for [installing with Docker](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent#configuring-and-running-the-broker-client-for-container-registry-agent) and [installing with Helm](snyk-broker-container-registry-agent/install-broker-for-container-registry-agent-using-helm.md).
* **Snyk Infrastructure as Code** – configure the [`accept.json` file with additional parameters](snyk-broker-infrastructure-as-code-detection/) to detect and analyze Terraform, CloudFormation, and Kubernetes configuration files through Snyk Broker.

## How Snyk Broker works

Snyk Broker is designed to connect Snyk products to self-hosted integrations that are not publicly accessible from the internet. Snyk Broker also allows you to do the following:

* Control Snyk access to your network by limiting the files to which Snyk has access and the actions that Snyk can perform.
* Manage a fixed private IP for your integration, targeting the Broker.

Snyk Broker includes a Server and a Client, basic components that are the same across all integrations. The Broker Server runs on the Snyk SaaS backend and is provided by Snyk; no installation is required. The Broker Client is a [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure. For more information, see [Components of Snyk Broker](components-of-snyk-broker.md) and [Connections with Snyk Broker](connections-with-snyk-broker.md).

See [Prepare Snyk Broker for deployment](prepare-snyk-broker-for-deployment.md) for information about prerequisites, choosing components, network configuration, and credentials.
