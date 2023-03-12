# Set up Snyk Broker with Azure Repos

Follow the instructions on this page to set up Snyk Azure Repos with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise or cloud Azure Repos deployment.

## Configure Broker to be used for Azure Repos

{% hint style="info" %}
Ask your Snyk account team to provide you with a Broker token.
{% endhint %}

{% hint style="info" %}
You need Docker or a way to run Docker containers.
{% endhint %}

{% hint style="info" %}
Snyk supports only Azure DevOps/TFS 2020 or above.
{% endhint %}

The following explains how to configure Snyk Broker to be used with Snyk Azure Repos.

**Obtain** your Azure Repos Broker token from Snyk.

To use the Broker client with [Azure](https://azure.microsoft.com/en-us/services/devops/), **run** `docker pull snyk/broker:azure-repos` tag. The following environment variables are required to configure the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Azure Repos integration settings view (app.snyk.io).
* `AZURE_REPOS_TOKEN` - an Azure Repos personal access token. Refer to this [Guide](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page) for how to get or create the token. Required scopes: ensure Custom defined is selected and under Code select _Read & write._
* `AZURE_REPOS_ORG` - organization name, which can be found in your Organization Overview page in Azure. For Azure Repos on-prem, the typical organization name is `tfs/Main`
* `AZURE_REPOS_HOST` - the hostname of your Azure Repos Server deployment, such as `your.azure-server.domain.com`.
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Azure Repos' webhooks, such as `http://broker.url.example:8000.`This URL is required to access features such as PR Fixes or PR Checks.
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Terraform, you can add an environment variable, `ACCEPT_IAC`, with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can add an environment variable, `ACCEPT_CODE=true`

**Use the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, and Code files (with the Code Agent).

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e AZURE_REPOS_TOKEN=<secret-azure-token> \
           -e AZURE_REPOS_ORG=<org-name> \
           -e AZURE_REPOS_HOST=<your.azure-server.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
       snyk/broker:azure-repos
```

**If necessary,** go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and **make any configuration changes** needed. For example, if the Azure Repos instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration or if you need to setup [proxy support](https://docs.snyk.io/integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client#proxy-support).

As an alternative to the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../../../snyk-admin/snyk-broker/snyk-broker-set-up-examples/derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Azure Repos integration.

### Custom allowlist through ACCEPT parameter

In addition, a fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for Azure Repos is attached in case you want to configure a custom allowlist. You cannot run the `ACCEPT_IAC` and `ACCEPT_CODE` arguments at the same time as the `ACCEPT` argument:

{% file src="../../../.gitbook/assets/accept (1) (1) (1) (3) (2).json" %}

**Paste the Broker Client configuration** to start the broker client container.

Once the container is up, the Azure Repos Integrations page shows the connection to Azure Repos and you can `Add Projects.`

## Basic troubleshooting for Broker with Azure Repos

* Run `docker logs <container id>` to look for any errors, where `container id` is the Azure Repos Broker container ID.
* Ensure relevant ports are exposed to Azure Repos.
* Make sure that file permissions for the local path to as well as the `accept.json` file itself are correct and accessible.
