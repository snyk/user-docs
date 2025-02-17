# GitLab - install and configure using Docker

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

## Configure Broker to be used for GitLab

To use the Broker client with GitLab.com or an on-prem GitLab deployment, **run** `docker pull snyk/broker:gitlab`. Refer to [GitLab - environment variables for Snyk Broker](gitlab-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

**If necessary,** go to the  [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and **make any configuration changes** needed, such as providing the CA (Certificate Authority to the Broker Client configuration if the GitlLab instance is using a private certificate, or setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for GitLab

**Copy the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files, and Snyk AppRisk information. Enable [Snyk AppRis](../../../../scan-with-snyk/snyk-apprisk.md)k to identify your application assets, monitor them, and prioritize the risks.

{% hint style="info" %}
**Multi-tenant settings for regions other than the default**\
When you set up Snyk Broker for use in regions other than the default, additional environment variables with specific URLs are required. For the URLs and examples, see [Regional hosting and data residency, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls).
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITLAB_TOKEN=<secret-gitlab-token> \
           -e GITLAB=<your.gitlab.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:gitlab
```

Use the environment variable `REMOVE_X_FORWARDED_HEADERS=true` to remove the `XFF` headers from the requests made by the Broker client to GitLab. This allows the Broker to work properly.

Snyk AppRisk is set by default to `false`. Enable it by setting the flag to `true`.

As an **alternative to using the Docker run command**, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the GitLab integration.

## Start the Broker Client container and verify the connection with GitLab

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the GitLab Integrations page shows the connection to GitLab and you can `Add Projects`.

## Basic troubleshooting for Broker with GitLab

* Run `docker logs <container id>` to look for any errors, where `container id` is the GitLab Broker container ID.
* Ensure relevant ports are exposed to GitLab.
