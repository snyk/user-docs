# Set up Snyk Broker with GitLab

Follow the instructions on this page to set up GitLab with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

{% hint style="info" %}
**Prereqisites**

Ask your Snyk account team to provide you with a Broker token.

You need Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

## Configure Broker to be used for GitLab

To use the Broker client with GitLab.com or an on-prem GitLab deployment, **run** `docker pull snyk/broker:gitlab`. The following environment variables are required to configure the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your GitLab integration settings view (app.snyk.io).
* `GITLAB_TOKEN` - a GitLab personal access token with `api` scope
* `GITLAB` - the hostname of your GitLab deployment, such as `your.gitlab.domain.com` or `GitLab.com`.
* `PORT` - the local port at which the Broker Client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker Client as needed to be reachable by either GitLab.com or on-prem GitLab deployment in order to establish webhook connectivity. This must be a full URL like `http://broker.url.example:8000`
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`

**If necessary,** go to the  [Advanced configuration page](../install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation.md) and **make any configuration changes** needed, such as providing the CA (Certificate Authority to the Broker Client configuration if the GitlLab instance is using a private certificate, or setting up [proxy support](https://docs.snyk.io/integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client#proxy-support). See also [Adding custom accept.json for Docker installation](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client/adding-custom-accept.json-for-docker-installation.md).

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

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the GitLab integration.

## Start the Broker Client container and verify the connection with GitLab

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the GitLab Integrations page shows the connection to GitLab and you can `Add Projects`.

## Basic troubleshooting for Broker with GitLab

* Run `docker logs <container id>` to look for any errors, where container id is the GitLab Broker container ID.
* Ensure relevant ports are exposed to GitLab.
