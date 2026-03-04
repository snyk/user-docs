# Install and configure Broker using Docker

{% hint style="info" %}
[Broker version 4.205.1](https://github.com/snyk/broker/blob/cb4f89e05eb42605f076321b952cdb7e57bf4111/config.default.json#L8) has been [released](https://updates.snyk.io). In this version, all `ACCEPT` rule flags will be enabled by default. This reduces needed configuration. If you do not want a specific `ACCEPT` rule flag to be enabled, you can opt out of the default `ACCEPT` all behavior by adding `ACCEPT_=false` to your Broker client configuration.
{% endhint %}

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

Before starting installation, review the [Prerequisites](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/#prerequisites-for-snyk-broker) and other information on the page [Prepare Snyk Broker for deployment](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

If you are using Kubernetes, Snyk recommends that you install Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm). For details, see[ Install and configure Broker using Helm.](install-and-configure-broker-using-helm.md)

For all other environments, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. The pages listed here explain how to set up the Snyk Broker Client integrations using Docker.

## Install using Docker

The following pages explain how to install these special integrations.

* [GitHub](github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md)
* [GitHub Enterprise](github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)
* [Bitbucket Server/Data Centre](bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-docker.md)
* [Gitlab](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-docker.md)
* [Azure Repos](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-using-docker.md)
* [JFrog Artifactory Repository](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/artifactory-repository-install-and-configure-broker/artifactory-repository-install-and-configure-using-docker.md)
* [Nexus Repository Manager](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-docker.md)
* [Jira](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/jira-prerequisites-and-steps-to-install-and-configure-broker/jira-install-and-configure-using-docker.md)
* [Snyk Broker - Container Registry Agent](../../../../implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/) (needed to connect to Container Registries)

You can customize the configuration using the environment variables in the Docker images. For this reason, install separate, multiple instances of the Broker Client for different integration types to ensure proper configuration as well as redundancy.

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Advanced configuration using Docker

When you install using Docker, follow the instructions on the pages about [Advanced configuration for Snyk Broker Docker installation](../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/) as needed.
