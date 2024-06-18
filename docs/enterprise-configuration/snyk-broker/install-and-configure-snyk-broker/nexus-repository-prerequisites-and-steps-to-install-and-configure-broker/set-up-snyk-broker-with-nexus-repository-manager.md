# Nexus Repository - install and configure using Docker

{% hint style="info" %}
**Feature availability**

Integration with Nexus Repository Manager is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise Nexus Repository Manager deployment.

For information about non-brokered integration with Nexus Repository Manager including supported environments and versions and user permissions, see [Nexus Repository Manager setup](../../../../scan-using-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/). For information about brokered integration with Nexus Container Registry see [Snyk Broker -Container Registry Agent](../../snyk-broker-container-registry-agent/).

## Configure Broker to be used for Nexus plugins

### Docker pull for Nexus 3 and Nexus 2 configuration

To use the Broker client with a Nexus 3 deployment, **run** `docker pull snyk/broker:nexus`.

To use the Broker client with a Nexus 2 deployment, **run** `docker pull snyk/broker:nexus2`.

For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

### Docker run commands to set up Broker Client for Nexus 3 and Nexus 2 integrations

**Copy the following command** to set up a fully configured Broker Client to use with Nexus 3. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 7341:7341 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BASE_NEXUS_URL=https://[<user>:<pass>@]<your.nexus.hostname> \
           -e BROKER_CLIENT_VALIDATION_URL=https://<your.nexus.hostname>/service/rest/v1/status[/check] \
           -e RES_BODY_URL_SUB=https://<your.nexus.hostname>/repository \
       snyk/broker:nexus
```

**Copy the following command** to set up a fully configured Broker Client to use with Nexus 2. You can run the Docker container by providing the relevant configuration:

```
docker run --restart=always \
  -p 7341:7341 \
  -e BROKER_TOKEN=<secret-broker-token> \
  -e BASE_NEXUS_URL=https://[username:password]@acme.com \
  -e RES_BODY_URL_SUB=https://acme.com/nexus/content/(groups|repositories) \ 
  snyk/broker:nexus2
```

As an **alternative to using the Docker run command**, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Nexus3 integration.

## Start the Broker Client container and verify the connection with Nexus Repository Manager

Paste the Broker Client configuration to start the Broker Client container.

Check connection status by making a request to the Broker Client `/systemcheck` endpoint.

Example: `curl http://172.17.0.2:7341/systemcheck`

You see success output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"brokerClientValidationUrlStatusCode":200,"ok":true}`

Or failure output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"ETIMEDOUT"}`
