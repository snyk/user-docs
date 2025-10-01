# GitLab - install and configure using Docker

Before installing, review the [prerequisites](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/) and the general instructions for installation using [Docker](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

## Configure Broker to be used for GitLab

To use the Broker client with GitLab.com or an on-prem GitLab deployment, run `docker pull snyk/broker:gitlab`. Refer to [GitLab - environment variables for Snyk Broker](gitlab-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

If necessary, go to the  [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and make any configuration changes needed, such as providing the CA (Certificate Authority to the Broker Client configuration if the GitLab instance is using a private certificate, or setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for GitLab

Copy the following command to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files, and Snyk Essentials information. Enable [Snyk Essentials ](../../../../../../scan-with-snyk/snyk-essentials.md)to identify your application assets, monitor them, and prioritize the risks.

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e GITLAB_TOKEN=<secret-gitlab-token> \
           -e GITLAB=<your.gitlab.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_ESSENTIALS=true \
       snyk/broker:gitlab
```

Use the environment variable `REMOVE_X_FORWARDED_HEADERS=true` to remove the `XFF` headers from the requests made by the Broker client to GitLab. This allows the Broker to work properly.

Snyk Essentials is set by default to `false`. Enable it by setting the flag to `true`.

## Start the Broker Client container and verify the connection with GitLab

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the GitLab Integrations page shows the connection to GitLab, and you can `Add Projects`.

## Basic troubleshooting for Broker with GitLab

* Run `docker logs <container id>` to look for any errors, where `container id` is the GitLab Broker container ID.
* Ensure relevant ports are exposed to GitLab.
