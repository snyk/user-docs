# Azure Repos - environment variables for Snyk Broker

The following environment variables are required to configure the Broker Client for Azure Repos:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Azure Repos integration settings view (app.snyk.io).
* `AZURE_REPOS_TOKEN` - an Azure Repos personal access token. Refer to this [Guide](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page) for how to get or create the token. Required scopes: ensure Custom defined is selected and under Code select _Read & write._
* `AZURE_REPOS_ORG` - Organization name, which can be found on your Organization Overview page in Azure. For Azure Repos on-prem, the typical organization name is `tfs/Main.` If you have more than one Azure organization, you must deploy a Broker for each one
* `AZURE_REPOS_HOST` - the hostname of your Azure Repos Server deployment, such as `your.azure-server.domain.com`.
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Azure Repos' webhooks, such as `http://broker.url.example:8000.`This URL is required to access features such as PR Fixes or PR Checks.
  * This must have http:// and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Terraform, you can add an environment variable, `ACCEPT_IAC`, with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can add an environment variable, `ACCEPT_CODE=true`
* `ACCEPT_APPRISK` - Enable Snyk AppRisk to identify your application assets, monitor them, and prioritize the risks. You can enable it by adding an environment variable `ACCEPT_APPRISK=true`.
