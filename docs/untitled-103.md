# Azure Functions overview

##  Azure Functions overview

Snyk's Azure Function Apps integration lets you monitor the deployed code of your Node.js, .Net and Java Azure Function Apps for any known vulnerabilities found in the application's dependencies, testing at a frequency you control.

For each test, Snyk will communicate directly with Azure to determine exactly what code is currently deployed and what dependencies are being used. Each dependency will in turn be tested against Snyk's vulnerability database to see if it contains any known vulnerabilities.

If vulnerabilities are found, you will be alerted \(via email or Slack\) so that you can take immediate action.

In order to turn on the Azure Function Apps integration you'll need to:

1. Connect to Azure from the integrations page
2. Add your Azure account credentials to Snyk
3. Select the projects you want to monitor and click "**Add to Snyk**"

###  **Connect Snyk to Azure Functions**

In order for Snyk to be able to monitor your deployed Azure Function apps, you'll first need to connect Snyk to your Azure account. You can do this by navigating to the [Integrations page](https://app.snyk.io/integrations), locating "Azure and clicking on "**Connect to Azure Functions**".

This will take you to a page where you'll be prompted to enter your Azure service principal credentials.

Instructions for how to generate and locate your Azure service principal credentials are below.

###  **Generate your Azure service principal**

To give Snyk access to your Azure account, you'll need a valid service principal.

To create a service principal for use by Snyk, you can either use the [Azure Portal](https://portal.azure.com/) or the [Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli).

After installing the Azure CLI 2.0, you should have the az command. Authenticate the CLI with your account using:

`az login`

Once authenticated, use it to create the service principal:

az ad sp create-for-rbac --name SpNameExample --role "Website Contributor"

This would result in JSON output similar to the following, which contains the service principal name, password and tenant that you'll need for setting up Snyk:

```text
{
"appId": "f5f1ce7d-c247-42e6-91a4-ad1e7example",
"displayName": "SpNameExample",
"name": "http://SpNameExample",
"password": "97adeba6-4178-4f2b-bf5f-782b3example",
"tenant": "874186fd-a7a8-4e98-9b9e-3df00example"
}
```

From there you can login to your Snyk account and paste in your name, password and tenant.

More information on creating a service principal for Azure can be found in Azure's documentation:

[CLI](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) , [Portal](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-create-service-principal-portal)

###  **Azure Functions: check your connection status**

At any time after you've entered your Azure service principal, you can check on the connection status in one of two places.

The first is on your integration settings page, where you'll see your current integrations listed as well as their connection status.

You can also check the status directly on the Azure Functions integration settings page \(found by clicking "**Edit settings**" on the integration settings page shown above\). If you've entered credentials, you'll see a box indicating whether or not Snyk is able to correctly connect to Azure.

If you are unable to connect, please re-enter your account credentials to verify that they are correct.

###  **Disable the Azure Functions integration**

###  **Add Azure Functions to Snyk**

Once you've successfully connected Snyk to your Azure account, you'll be able to select Azure functions that you would like Snyk to monitor. You can do this either using the **"Add projects"** button on the integrations page, or directly from the Azure Functions integration settings page.

In either case, you'll see a list of any available Function apps on the Azure account you connected. Select the ones you want to monitor and click the "**Add to Snyk**" button.

_Note: We currently support importing only v2 functions. v1 functions will be ignored._

As soon as you've added the projects to Snyk, Snyk will test them and begin to display a list of all monitored Azure functions in your [project dashboard](https://app.snyk.io/projects). You'll also see a snapshot of any current vulnerabilities, and be able to click through for a more detailed report including any steps to remediate.

Snyk will now continuously monitor each of those functions for known vulnerabilities. You can add more functions at any time.

_Note: for Node.js and .Net, dev dependencies will be ignored._

