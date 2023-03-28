# Set up Snyk Broker with Bitbucket Server/Data Center

Follow the instructions on this page to set up Bitbucket Server/Data Center with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise Bitbucket deployment.

{% hint style="info" %}
**Prerequisites**\
****Ask your Snyk account team to provide you with a Broker token.

You need Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

## Configure Broker to be used with Bitbucket

The following explains how to configure Snyk Broker to use the Broker Client with a Bitbucket Server deployment.

To use the Snyk Broker Client with BitBucket, **run** `docker pull snyk/broker:bitbucket-server.`The following environment variables are required to configure the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Bitbucket Server integration settings view (app.snyk.io).
* `BITBUCKET_USERNAME` - the Bitbucket Server username.
* `BITBUCKET_PASSWORD` - the Bitbucket Server password.
* `BITBUCKET` - the hostname of your Bitbucket Server deployment, such as `your.bitbucket-server.domain.com`.
* `BITBUCKET_API` - the API endpoint of your Bitbucket Server deployment. Should be `$BITBUCKET/rest/api/1.0`.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Bitbucket Server for webhooks, such as `http://broker.url.example:8000.`This URL is required to access features such as PR Fixes or PR Checks.
* `PORT` - the local port at which the Broker client accepts connections. Default is 800.
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`

**If necessary,** go to the [Advanced Configuration page](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client/advanced-configuration-for-snyk-broker-docker-installation.md) and **make any configuration changes** needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the Bitbucket instance is using a private certificate, and setting up [proxy support](https://docs.snyk.io/integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client#proxy-support).

## Docker run command to set up a Broker Client for Bitbucket

**Use the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, and Code files (with the Code Agent).

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
       snyk/broker:bitbucket-server
```

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the BitBucket Server/Data Center integration.

## Start the Broker Client container and verify the connection with Bitbucket

Paste the Broker Client configuration to start the broker Client container.

Once the container is up, the Bitbucket Integrations page shows the connection to Bitbucket and you can `Add Projects`

## Custom allowlist through ACCEPT parameter for Bitbucket

In addition, a fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for Bitbucket is attached to use if you want to configure a custom allowlist. You cannot run the `ACCEPT_IAC` and `ACCEPT_CODE` arguments at the same time as the `ACCEPT` argument:

{% file src="../../../.gitbook/assets/accept (2).json" %}

## Basic troubleshooting for Broker with BitBucket

* Run `docker logs <container id>` to look for any errors, where container id is the Bitbucket Broker container ID.
* Ensure relevant ports are exposed to Bitbucket.
