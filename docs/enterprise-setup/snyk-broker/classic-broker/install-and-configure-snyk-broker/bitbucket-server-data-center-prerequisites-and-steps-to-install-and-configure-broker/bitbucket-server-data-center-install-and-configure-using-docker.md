# Bitbucket Server/Data Center - install and configure using Docker

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise Bitbucket deployment.

This page describes two distinct authentication schemes: [Basic Auth](bitbucket-server-data-center-install-and-configure-using-docker.md#configure-broker-to-be-used-with-bitbucket-using-basic-auth) and [Bearer (Personal Access Token)](bitbucket-server-data-center-install-and-configure-using-docker.md#configure-broker-to-be-used-with-bitbucket-using-a-personal-access-token-pat). Your Bitbucket Server settings might preclude Basic Auth usage, in which case Bearer Auth is preferred.

## Configure Broker to be used with Bitbucket using Basic Auth

The following explains how to configure Snyk Broker to use the Broker Client with a Bitbucket Server deployment.

To use the Snyk Broker Client with BitBucket, run `docker pull snyk/broker:bitbucket-server`. Refer to [BitBucket Server/Data Center - environment variables](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md) for Snyk Broker for definitions of the environment variables.

If necessary, navigate to the [Advanced configuration page](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/) and make any configuration changes needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the Bitbucket instance is using a private certificate, and setting up [proxy support](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for Bitbucket using Basic Auth

Copy the following command to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files, and Snyk Essentials information. Enable [Snyk Essentials](../../../../../scan-with-snyk/snyk-essentials.md) to identify your application assets, monitor them, and prioritize the risks.

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e BITBUCKET_USERNAME=<username> \
           -e BITBUCKET_PASSWORD=<password> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_Essentials=true \
       snyk/broker:bitbucket-server
```

{% hint style="info" %}
Snyk Essentials is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}

## Configure Broker to be used with Bitbucket using an API token

The following explains how to configure Snyk Broker to use the Broker Client with a Bitbucket Server deployment using an API token.

To use the Snyk Broker Client with BitBucket, **run** `docker pull snyk/broker:bitbucket-server-bearer-auth`. For definitions of the environment variables, refer to [Bitbucket Server/Data Center - environment variables for Snyk Broker Basic Auth](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md) and [Bitbucket Server/Data Center - environment variables for Snyk Broker Personal Access Token (PAT)](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-personal-access-token-pat.md).

If necessary, go to the [Advanced configuration page](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/) and make any configuration changes needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the Bitbucket instance is using a private certificate, and setting up [proxy support](../../../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for Bitbucket using a PAT

Copy the following command to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files, and Snyk Essentials information. Enable [Snyk Essentials](../../../../../scan-with-snyk/snyk-apprisk.md) to identify your application assets, monitor them, and prioritize the risks.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BITBUCKET_PAT=<personal-access-token> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_Essentials=true \
       snyk/broker:bitbucket-server-bearer-auth
```

{% hint style="info" %}
Snyk Essentials is set by default to `false`. Enable it by setting the flag to `true`.
{% endhint %}

## Start the Broker Client container and verify the connection with Bitbucket

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the Bitbucket Integrations page shows the connection to Bitbucket, and you can `Add Projects`

## Basic troubleshooting for Broker with BitBucket

* Run `docker logs <container id>` to look for any errors, where `container id` is the Bitbucket Broker container ID.
* Ensure relevant ports are exposed to Bitbucket.
