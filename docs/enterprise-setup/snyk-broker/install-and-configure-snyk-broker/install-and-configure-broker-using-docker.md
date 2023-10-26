# Install and configure Broker using Docker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access by [snyk.io](http://snyk.io/) to your code to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the Snyk Broker user documentation.

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see[ Install and configure Broker using Helm.](install-and-configure-broker-using-helm.md)

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. The pages listed here explain how to set up the Snyk Broker Client integrations using Docker.

## Installation using Docker

**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).

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

When you install using Docker, follow these instructions to configure Broker as needed:

* [Changing the auth method with Docker](advanced-configuration-for-snyk-broker-docker-installation/changing-the-auth-method-with-docker.md)
* [Credential pooling with Docker](advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md)
* [HTTPS for Broker Client with Docker](advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker.md)
* [Backend requests with an internal certificate for Docker](advanced-configuration-for-snyk-broker-docker-installation/backend-requests-with-an-internal-certificate-for-docker.md)
* [Proxy support with Docker](advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md)
* [Disable certificate verification with Docker](advanced-configuration-for-snyk-broker-docker-installation/disable-certificate-verification-with-docker.md)
* [Adding custom allowlist for Docker installation](broken-reference)
* [Custom approved-listing filter with Docker](advanced-configuration-for-snyk-broker-docker-installation/custom-approved-listing-filter-with-docker.md)
* [Mounting secrets with Docker](advanced-configuration-for-snyk-broker-docker-installation/mounting-secrets-with-docker.md)
* [Snyk Code - Clone capability with Broker for Docker](advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md)
* [High availability mode](../high-availability-mode.md)
* [Large manifest files rule addition](advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup.md)
* [Insecure downstream mode](advanced-configuration-for-snyk-broker-docker-installation/insecure-downstream-mode.md)

## Common questions about Snyk Broker

**How often is Snyk Broker updated?**\
Snyk Broker is updated each time new features become available and when there are fixes.

**How often is Snyk Broker checked for vulnerabilities?**\
The Snyk Broker application and images are tested daily for vulnerabilities.

**What is the SLA to fix vulnerabilities?**\
There is a 14-day SLA for fixing high vulnerabilities and a five-day SLA for fixing critical vulnerabilities in public images.

## Additional information for developers

If you need to upgrade, see [Upgrade the Snyk Broker Client](https://docs.snyk.io/snyk-admin/snyk-broker/upgrade-the-snyk-broker-client).

Troubleshooting information is provided on the [Troubleshooting Broker](https://docs.snyk.io/snyk-admin/snyk-broker/troubleshooting-broker) page.

You can view the [license, Apache License, Version 2.0](https://github.com/snyk/broker/blob/master/LICENSE).

To submit pull requests, see [Contributing](https://github.com/snyk/broker/blob/master/.github/CONTRIBUTING.md).

See [Security](https://github.com/snyk/broker/blob/master/SECURITY.md) for specific information about Broker.

