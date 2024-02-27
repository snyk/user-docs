# Install and configure Broker using Docker

Before starting installation, review the [Prerequisites](../prepare-snyk-broker-for-deployment.md#prerequisites-for-snyk-broker) and other information on the page [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment.md).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see[ Install and configure Broker using Helm.](install-and-configure-broker-using-helm.md)

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. The pages listed here explain how to set up the Snyk Broker Client integrations using Docker.

## Install using Docker

The following pages explain how to install these special integrations.

* [GitHub](github-prerequisites-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](github-enterprise-prerequisites-to-install-and-configure-broker/setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](bitbucket-server-data-center-prerequisites-to-install-and-configure-broker/data-center.md)
* [Gitlab](gitlab-prerequisites-to-install-and-configure-broker/setup-broker-with-gitlab.md)
* [Azure Repos](azure-repos-prerequisites-to-install-and-configure-broker/setup-broker-with-azure-repos.md)
* [JFrog Artifactory Repository](artifactory-repository-install-and-configure-broker/set-up-snyk-broker-with-artifactory-repository.md)
* [Nexus Repository Manager](nexus-repository-prerequisites-to-install-and-configure-broker/set-up-snyk-broker-with-nexus-repository-manager.md)
* [Jira](jira-prerequisites-to-install-and-configure-broker/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../snyk-broker-code-agent/) (needed to enable SAST analysis)
* [Derived Docker images for Broker Client integrations and Container Registry Agent](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)

You can customize the configuration using the environment variables in the Docker images. For this reason, install separate, multiple instances of the Broker Client for different integration types to ensure proper configuration as well as redundancy.

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Advanced configuration using Docker

When you install using Docker, follow the instructions in the pages on [Advanced configuration for Snyk Broker Docker installation](advanced-configuration-for-snyk-broker-docker-installation/) as needed.

##

