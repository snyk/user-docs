# Nexus Repository - install and configure using Docker

{% hint style="info" %}
**Feature availability**

Integration with Nexus Repository Manager is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise Nexus Repository Manager deployment.

For information about non-brokered integration with Nexus Repository Manager including supported environments and versions and user permissions, see [Nexus Repository Manager setup](../../../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/). For information about brokered integration with Nexus Container Registry see [Snyk Broker -Container Registry Agent](../../../snyk-broker-container-registry-agent/).

## Configure Broker to be used for Nexus plugins

### Docker pull for Nexus 3 and Nexus 2 configuration

To use the Broker client with a Nexus 3 deployment, run `docker pull snyk/broker:nexus`.

To use the Broker client with a Nexus 2 deployment, run `docker pull snyk/broker:nexus2`.

For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

### Docker run commands to set up Broker Client for Nexus 3 and Nexus 2 integrations

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`.This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

Copy the following command to set up a fully configured Broker Client to use with Nexus 3. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 7341:7341 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e BASE_NEXUS_URL=https://[<user>:<pass>@]<your.nexus.hostname> \
           -e BROKER_CLIENT_VALIDATION_URL=https://<your.nexus.hostname>/service/rest/v1/status[/check] \
           -e RES_BODY_URL_SUB=https://<your.nexus.hostname>/repository \
       snyk/broker:nexus
```

Copy the following command to set up a fully configured Broker Client to use with Nexus 2. You can run the Docker container by providing the relevant configuration:

```
docker run --restart=always \
  -p 7341:7341 \
  -e BROKER_TOKEN=<secret-broker-token> \
  -e BROKER_SERVER_URL=<broker-region-url> \
  -e BASE_NEXUS_URL=https://[username:password]@acme.com \
  -e RES_BODY_URL_SUB=https://acme.com/nexus/content/(groups|repositories) \
  snyk/broker:nexus2
```

You can find your `BASE_NEXUS_URL` by visiting the Nexus UI and navigating to the server tab under **Administration**, then selecting the **Base URL** entry without a trailing slash. This will typically end with `/nexus`, but may vary with non-default deployments. If you have a custom base URL then you must also set the `NEXUS_URL` environment variable to point to the URL where your repositories live. By default this is configured as `/nexus/content` but should follow a similar format to your base URL.

## Start the Broker Client container and verify the connection with Nexus Repository Manager

Paste the Broker Client configuration to start the Broker Client container.

Check connection status by making a request to the Broker Client `/systemcheck` endpoint.

Example: `curl http://172.17.0.2:7341/systemcheck`

You see success output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"brokerClientValidationUrlStatusCode":200,"ok":true}`

Or failure output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"ETIMEDOUT"}`
