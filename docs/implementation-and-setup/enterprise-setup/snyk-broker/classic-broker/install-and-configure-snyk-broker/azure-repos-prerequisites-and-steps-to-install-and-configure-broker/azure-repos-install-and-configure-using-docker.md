# Azure Repos - install and configure using Docker

{% hint style="info" %}
**Feature availability**

Snyk Azure Repos are available only for Azure DevOps/TFS 2020 or above.
{% endhint %}

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise or cloud Azure Repos deployment.

## Configure Broker to be used with Azure Repos

To use the Broker Client with [Azure](https://azure.microsoft.com/en-us/services/devops/), run `docker pull snyk/broker:azure-repos`. Refer to [Azure Repos - environment variables for Snyk Broker](azure-repos-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

If necessary, go to the [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and make any configuration changes needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the Azure Repos instance is using a private certificate, and setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for Azure Repos

Copy the following command to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files, and Snyk Essentials information for one Azure organization. Enable [Snyk Essentials](../../../../../../scan-with-snyk/snyk-essentials.md) to identify your application assets, monitor them, and prioritize the risks.

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

Note that if you have more than one Azure organization, you must deploy a Broker for each one. Snyk Essentials is set by default to `false`. Enable it by setting the flag to `true`.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e AZURE_REPOS_TOKEN=<secret-azure-token> \
           -e AZURE_REPOS_ORG=<org-name> \
           -e AZURE_REPOS_HOST=<your.azure-server.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_ESSENTIALS=true \
       snyk/broker:azure-repos
```

## Start the Broker Client container and verify the connection with Azure Repos

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the Azure Repos Integrations page shows the connection to Azure Repos and you can `Add Projects.`

## Basic troubleshooting for Broker with Azure Repos

* Run `docker logs <container id>` to look for any errors, where `container id` is the Azure Repos Broker container ID.
* Ensure relevant ports are exposed to Azure Repos.
* Make sure that file permissions for the local path to the `accept.json` file, as well as the `accept.json` file itself, are correct and accessible.
