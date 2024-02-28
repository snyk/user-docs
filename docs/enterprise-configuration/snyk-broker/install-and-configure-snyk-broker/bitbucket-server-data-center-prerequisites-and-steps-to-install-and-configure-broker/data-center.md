# Bitbucket Server/Data Center - install and configure using Docker

Before installing, review the prerequisites and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise Bitbucket deployment.

## Configure Broker to be used with Bitbucket

The following explains how to configure Snyk Broker to use the Broker Client with a Bitbucket Server deployment.

To use the Snyk Broker Client with BitBucket, **run** `docker pull snyk/broker:bitbucket-server`. Refer to [BitBucket Server/Data Center - environment variables](bitbucket-server-data-center-environment-variables-for-snyk-broker.md) for Snyk Broker for definitions of the environment variables.

**If necessary,** go to the [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and **make any configuration changes** needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the Bitbucket instance is using a private certificate, and setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for Bitbucket

**Copy the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files (with the Code Agent), and Snyk AppRisk information. Enable [Snyk AppRisk](../../../../manage-risk/snyk-apprisk/) to identify your application assets, monitor them, and prioritize the risks.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BITBUCKET_USERNAME=<username> \
           -e BITBUCKET_PASSWORD=<password> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:bitbucket-server
```

{% hint style="info" %}
Snyk AppRisk is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the BitBucket Server/Data Center integration.

## Start the Broker Client container and verify the connection with Bitbucket

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the Bitbucket Integrations page shows the connection to Bitbucket and you can `Add Projects`

## Basic troubleshooting for Broker with BitBucket

* Run `docker logs <container id>` to look for any errors, where container id is the Bitbucket Broker container ID.
* Ensure relevant ports are exposed to Bitbucket.
