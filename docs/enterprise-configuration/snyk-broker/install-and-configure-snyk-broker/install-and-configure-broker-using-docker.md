# Install and configure Broker using Docker

Before starting installation, review the [Prerequisites](../prepare-snyk-broker-for-deployment.md#prerequisites-for-snyk-broker) and other information on the page [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment.md).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see[ Install and configure Broker using Helm.](install-and-configure-broker-using-helm.md)

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. The pages listed here explain how to set up the Snyk Broker Client integrations using Docker.

## Installation using Docker

The following pages explain how to install these special integrations.

* [GitHub](github-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](github-enterprise-install-and-configure-broker/setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](bitbucket-server-data-center-install-and-configure-broker/data-center.md)
* [Gitlab](gitlab-install-and-configure-broker/setup-broker-with-gitlab.md)
* [Azure Repos](azure-repos-install-and-configure-broker/setup-broker-with-azure-repos.md)
* [JFrog Artifactory Repository](artifactory-repository-install-and-configure-broker/set-up-snyk-broker-with-artifactory-repository.md)
* [Nexus Repository Manager](nexus-repository-install-and-configure-broker/set-up-snyk-broker-with-nexus-repository-manager.md)
* [Jira](jira-install-and-configure-broker/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../snyk-broker-code-agent/) (needed to enable SAST analysis)
* [Derived Docker images for Broker Client integrations and Container Registry Agent](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)

{% hint style="info" %}
You can customize the configuration using the environment variables in the Docker images. For this reason, install separate, multiple instances of the Broker Client for different integration types to ensure proper configuration as well as redundancy.
{% endhint %}

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Advanced configuration using Docker

When you install using Docker, follow the instructions in the pages on [Advanced configuration for Snyk Broker Docker installation](advanced-configuration-for-snyk-broker-docker-installation/) as needed.

## Common questions about Snyk Broker

**How often is Snyk Broker updated?**\
Snyk Broker is updated each time new features become available and when there are fixes.

**How often is Snyk Broker checked for vulnerabilities?**\
The Snyk Broker application and images are tested daily for vulnerabilities.

**What is the SLA to fix vulnerabilities?**\
There is a 14-day SLA for fixing high vulnerabilities and a five-day SLA for fixing critical vulnerabilities in public images.

## Additional information for developers

If you need to upgrade, see [Upgrade the Snyk Broker Client](../upgrade-the-snyk-broker-client.md).

Troubleshooting information is provided on the [Troubleshooting Broker](../troubleshooting-broker.md) page.

You can view the [license, Apache License, Version 2.0](https://github.com/snyk/broker/blob/master/LICENSE).

To submit pull requests, see [Contributing](https://github.com/snyk/broker/blob/master/.github/CONTRIBUTING.md).

See [Security](https://github.com/snyk/broker/blob/master/SECURITY.md) for specific information about Broker.

