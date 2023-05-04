# Snyk Broker

{% hint style="info" %}
**Feature availability**\
Snyk Broker is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Introduction to Snyk Broker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, allowing [snyk.io](http://snyk.io/) access to your code for scanning and returning results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. Supported integrations with Broker include the following:

* Your Source Code Management (SCM) that is not internet reachable
* Your publicly-accessible SCM, which allows you to view and control Snyk activity for increased data security
* Network-restricted [container registries](snyk-broker-container-registry-agent/)
* [Infrastructure as code (IaC) configuration](snyk-broker-infrastructure-as-code-detection/) files using Snyk IaC located on private Git-based repositories

You can also use Snyk Broker to enable a secure connection with your on-premise Jira installation or JFrog Artifactory installation.

### How Snyk Broker works

Snyk Broker is designed to connect Snyk products to self-hosted integrations that are not publicly accessible from the internet. Snyk Broker also allows you to do the following:

* Control Snyk's access to your network by limiting the files to which Snyk has access and the actions that Snyk can perform.
* Manage a fixed private IP for your integration, targeting the Broker.

Snyk Broker includes a Server and a Client, basic components that are the same across all integrations. The Broker Server runs on the Snyk SaaS backend and is provided by Snyk; no installation is required. The Broker Client is a [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure. For more information, see [Components of Snyk Broker](components-of-snyk-broker.md) and [Connections with Snyk Broker](connections-with-snyk-broker.md).

### How to download and install Snyk Broker

The Snyk Broker project is hosted on [GitHub](https://github.com/snyk/broker) and published as a set of Docker images for specific integrations. Snyk provides a [Helm Chart](https://github.com/snyk/snyk-broker-helm) to deploy the Snyk Broker. To deploy Broker, you must install and configure an integration.

{% hint style="info" %}
Snyk recommends using Helm as the simplest way to deploy Snyk Broker. You can also use Docker to run the Snyk Broker Client or run npm `install snyk-broker`.

This documentation provides detailed instructions for using [Helm](install-and-configure-broker-using-helm/) and [Docker](install-and-configure-broker-using-docker/).
{% endhint %}

## **Integrations with Snyk Broker**

Install each type of integration and configure using environment variables as explained in these pages:

* [Install and configure Broker using Helm](install-and-configure-broker-using-helm/)
* [Install and configure Broker using Docker](install-and-configure-broker-using-docker/)

Snyk Broker currently integrates with the Git repository systems listed here, with links to the Docker instructions. Installing with Helm is explained for each repository on one page, [Install and configure Broker using Helm](install-and-configure-broker-using-helm/).

{% hint style="info" %}
Configuration information for the environment variables is found on the pages with Docker instructions.
{% endhint %}

* [GitHub](../../integrations/git-repository-scm-integrations/github-integration.md) and [GitHub Enterprise](../../integrations/git-repository-scm-integrations/github-enterprise-integration.md) (Cloud and On-prem)
* [GitLab](../../integrations/git-repository-scm-integrations/gitlab-integration.md) (Cloud and On-prem)
* [Bitbucket Server / Data Center](../../integrations/git-repository-scm-integrations/bitbucket-data-center-server-integration.md) (On-prem)
* [Azure Repos](../../integrations/git-repository-scm-integrations/azure-repositories-integration.md) (Cloud and On-prem)

In addition, Snyk Broker integrates with [Jira Server/Jira Data Center](../../integrations/notifications-ticketing-system-integrations/jira.md), [JFrog Artifactory](../../integrations/private-registry-integrations/artifactory-repository-setup.md), and [Nexus Repository Manager](../../integrations/private-registry-integrations/nexus-repo-manager-setup.md).

With the [Container Registry Agent](snyk-broker-container-registry-agent/), Snyk Broker also connects to all [Snyk-supported container registries](snyk-broker-container-registry-agent/).

You can also use [derived Docker images](install-and-configure-broker-using-docker/snyk-broker-set-up-examples/derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for each integration and the Container Registry Agent.

For information about advanced configuration as needed for your installation, see [Advanced configuration for Snyk Broker Docker installation ](install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/)and [Advanced setup for Helm Chart installation](install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/).

## **Using Snyk Broker to scan your code**

To use **Snyk Open Source** with Snyk Broker, you need only the Broker Server and  Broker Client components. The Broker Client is published as a set of Docker images, each configured for a specific Git service. Configure each type of integration using environment variables as explained in the [Snyk Broker-Client integration setups](install-and-configure-broker-using-docker/snyk-broker-set-up-examples/).

To scan other types of code with Snyk Broker, you must add a component or configurations and add parameters to the Broker Client setup:

* **Snyk Code** – add the [**Code Agent** ](snyk-broker-code-agent/)component to enable Snyk Code analysis of repositories in SCMs that are integrated through Snyk Broker.
* **Snyk Container** – add the [**Container Registry Agent**](snyk-broker-container-registry-agent/) to enable the connection to network-restricted container registries and the analysis of container images.
* **Snyk Infrastructure as Code** – configure the [**`accept.json`** file with additional parameters](snyk-broker-infrastructure-as-code-detection/) to detect and analyze Terraform, CloudFormation, and Kubernetes configuration files through Snyk Broker.

## Common questions about Snyk Broker

**How often is Snyk Broker updated?**\
Snyk Broker is updated each time new features become available and when there are fixes.

**How often is Snyk Broker checked for vulnerabilities?**\
The Snyk Broker application and images are being tested daily for vulnerabilities.

**What is the SLA to fix vulnerabilities?**\
There is a 14-day SLA for fixing high vulnerabilities and a five-day SLA for fixing critical vulnerabilities in public images.

## Additional information for users

You can view the [License: Apache License, Version 2.0](https://github.com/snyk/broker/blob/master/LICENSE)

To submit pull requests, see [Contributing](https://github.com/snyk/broker/blob/master/.github/CONTRIBUTING.md)

See [Security](https://github.com/snyk/broker/blob/master/SECURITY.md) for specific information about Broker.
