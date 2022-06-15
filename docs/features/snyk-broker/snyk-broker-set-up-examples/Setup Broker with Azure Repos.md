# Setup Broker with Azure Repos

Configuring Azure Repos with broker is useful to ensure a secure connection with your on-premise or cloud Azure Repos deployment.

### To configure Broker to be used for Azure Repos

{% hint style="info" %}
Please ask for your CSM support to provide you with a Broker token
{% endhint %}

{% hint style="info" %}
You will require Docker or a way to run Docker containers
{% endhint %}

* Obtain your Azure Repos Broker token from your CSM or Support team
* To use the Broker client with [Azure](https://azure.microsoft.com/en-us/services/devops/), run `docker pull snyk/broker:azure-repos` tag. The following environment variables are mandatory to configure the Broker client:
  * `BROKER_TOKEN` - the Snyk Broker token, obtained from your Azure Repos integration settings view (app.snyk.io).
  * `AZURE_REPOS_TOKEN` - an Azure Repos personal access token. [Guide](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page) how to get/create the token. Required scopes: ensure Custom defined is selected and under Code select _Read & write_
  * `AZURE_REPOS_ORG` - organization name, which can be found in your Organization Overview page in Azure
  * `AZURE_REPOS_HOST` - the hostname of your Azure Repos Server deployment, such as `your.azure-server.domain.com`.
  * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Azure Repos' webhooks, such as `http://broker.url.example:8000`
* Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e AZURE_REPOS_TOKEN=<secret-azure-token> \
           -e AZURE_REPOS_ORG=<org-name> \
           -e AZURE_REPOS_<HOST=your.azure-server.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:azure-repos
```

* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../../integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes if required (For example, if the Azure Repos instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration). A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for Azure Repos is attached.

{% file src="../../../.gitbook/assets/accept (2).json" %}

* Paste the Broker Client configuration to start the broker client container
* Once the container is up, the Azure Repos Integrations page should show the connection to Azure Repos and you should be able to `Add Projects`

Basic Troubleshooting

* Run `docker logs <container id>` where container id is the Azure Repos Broker container ID to look for any errors
* Ensure relevant ports are exposed to Azure Repos
