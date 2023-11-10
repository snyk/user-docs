# GitLab - install and configure using Docker

Follow the instructions on this page to set up GitLab with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

{% hint style="info" %}
**Prerequisites**

Ask your Snyk account team to provide you with a Broker token.

Ensure that the GItLab permissions are correct for integration with Snyk. For details, see [GitLab integration](../../../../integrate-with-snyk/git-repository-scm-integrations/snyk-gitlab-integration.md).

You need Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

## Configure Broker to be used for GitLab

To use the Broker client with GitLab.com or an on-prem GitLab deployment, **run** `docker pull snyk/broker:gitlab`. Refer to [GitLab - environment variables for Snyk Broker](gitlab-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

**If necessary,** go to the  [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and **make any configuration changes** needed, such as providing the CA (Certificate Authority to the Broker Client configuration if the GitlLab instance is using a private certificate, or setting up [proxy support](https://docs.snyk.io/integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client#proxy-support). See also [Adding custom accept.json for Docker installation](broken-reference).

## Docker run command to set up a Broker Client for GitLab

**Copy the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, and Code files (with the Code Agent).

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
       snyk/broker:gitlab
```

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the GitLab integration.

## Start the Broker Client container and verify the connection with GitLab

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the GitLab Integrations page shows the connection to GitLab and you can `Add Projects`.

## Basic troubleshooting for Broker with GitLab

* Run `docker logs <container id>` to look for any errors, where container id is the GitLab Broker container ID.
* Ensure relevant ports are exposed to GitLab.
