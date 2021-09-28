# Configure the Azure Repos integration

### Enable the integration

Let's begin by familiarizing ourselves with the [integration documentation](https://support.snyk.io/hc/en-us/articles/360004002198-Azure-Repos-integration). Then, from the [Snyk](https://snyk.io) web console, navigate to `Integrations`. Search and select `Azure Repos`. 

![](../../../.gitbook/assets/snyk_integrations_09.png)

From the configuration menu, you will need to perform three steps:

1. Enter your `organization`
2. Generate a personal access token
3. Copy and paste that token and click `save`

![](../../../.gitbook/assets/snyk_integrations_10.png)

#### Generate a Personal Access Token

When you click the `Create a personal access token` button illustrated in the above diagram, you will be redirected to your Azure DevOps organization. You will need to perform the following four steps:

1. Provide a name for your token \(i.e. `snyk`\)
2. Select `Custom defined` for Scopes
3. Select `Read & write` for Code
4. Click `Create`

![](../../../.gitbook/assets/azure_tokens_01.png)

Make sure that you copy the token and paste this into the Snyk integrations page for Azure Repos.

![](../../../.gitbook/assets/azure_tokens_02.png)

You are now ready to add your Azure Repos repositories for scanning!

