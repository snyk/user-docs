# Artifactory Repository - install and configure using Docker

{% hint style="info" %}
**Feature availability**

Integration with Artifactory Repository is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise Artifactory Repository deployment.

For information about non-brokered integration with Artifactory Repository see [Artifactory Repository setup](../../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/). For information about brokered integration with Artifactory Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).

## Configure Broker to be used for Artifactory Registry

To use the Broker client with an Artifactory Registry deployment, **run** `docker pull snyk/broker:artifactory`. Refer to [Artifactory Repository - environment variables for Snyk Broker](artifactory-repository-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

## Docker run commands to set up a Broker Client for Artifactory Repository

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`.This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](https://snyksec.atlassian.net/o/-M4tdxG8qotLgGZnLpFR/s/-MdwVZ6HOZriajCf5nXH/~/changes/8951/working-with-snyk/regional-hosting-and-data-residency#broker-urls).
{% endhint %}

**Copy the following command** to set up a fully configured Broker Client to use with Artifactory Registry. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e ARTIFACTORY_URL=<yourdomain>.artifactory.com/artifactory \
       snyk/broker:artifactory
```

For an **npm or Yarn integration**, use the following **command**.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e ARTIFACTORY_URL=acme.com/artifactory \
  -e RES_BODY_URL_SUB=http://acme.com/artifactory \ 
  snyk/broker:artifactory
```

As an **alternative to using the Docker run command**, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Artifactory integration.

## Start the Broker Client container and verify the connection with Artifactory Repository

Paste the Broker Client configuration to start the Broker Client container.

You can check the status of the connection by refreshing the Artifactory Integration Settings page. When the connection is set up correctly, there is no connection error.
