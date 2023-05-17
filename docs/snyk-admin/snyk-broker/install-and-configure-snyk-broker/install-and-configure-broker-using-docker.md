# Install and configure Broker using Docker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access to your code by [snyk.io](http://snyk.io/) to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the [Snyk Broker user documentation](https://docs.snyk.io/snyk-admin/snyk-broker).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see [Install and configure Broker using Helm](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm).

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. The pages listed here explain how to set up the Snyk Broker Client integrations using Docker.

**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).

* [GitHub](broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](data-center.md)
* [Gitlab](setup-broker-with-gitlab.md)
* [Azure Repos](setup-broker-with-azure-repos.md)
* [JFrog Artifactory Repository](set-up-snyk-broker-with-artifactory-repository.md)
* [Nexus Repository Manager](set-up-snyk-broker-with-nexus-repository-manager.md)
* [Jira](setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../snyk-broker-code-agent/) (needed to enable SAST analysis)
* [Derived Docker images for Broker Client integrations and Container Registry Agent](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)

{% hint style="info" %}
You can customize the configuration using the environment variables in the Docker images. For this reason, install separate, multiple instances of the Broker Client for different integration types to ensure proper configuration as well as redundancy.
{% endhint %}

Once the Broker is running, visit the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a message like the following and start importing Projects:

<figure><img src="../../../.gitbook/assets/image (60) (1).png" alt="Brokered integration settings showing connected status"><figcaption><p>Brokered integration settings showing connected status</p></figcaption></figure>
