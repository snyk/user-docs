# Azure Functions overview

Snyk's Azure Function Apps integration lets you monitor the deployed code of your node.js, .Net and Java Azure Function Apps for any known vulnerabilities found in the application's dependencies, testing at a frequency you control.

For each test, Snyk will communicate directly with Azure to determine exactly what code is currently deployed and what dependencies are being used. Each dependency will in turn be tested against Snyk's vulnerability database to see if it contains any known vulnerabilities.

If vulnerabilities are found, you will be alerted (via email or Slack) so that you can take immediate action.

In order to turn on the Azure Function Apps integration you'll need to:

1. Connect to Azure from the integrations page
2. Add your Azure account credentials to Snyk
3. Select the projects you want to monitor and click "**Add to Snyk**"

## **Connect Snyk to Azure Functions**

For Snyk to monitor your deployed Azure Function apps, you must connect Snyk to your Azure account.

To do this, in [your Integrations page](https://app.snyk.io/integrations), navigate to **Serverless** and click **Azure Functions**:

![](<../../../.gitbook/assets/Screenshot 2021-10-27 at 09.36.33.png>)

This will take you to a page where you'll be prompted to enter your Azure service principal credentials:

![](<../../../.gitbook/assets/image (29).png>)

Instructions for how to generate and locate your Azure service principal credentials are below.

## **Generate your Azure service principal**

To give Snyk access to your Azure account, you'll need a valid service principal.

To create a service principal for use by Snyk, you can either use the [Azure Portal](https://portal.azure.com) or the [Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli).

After installing the Azure CLI 2.0, you have access to the `az` command. Authenticate the CLI with your account using:

```
az login
```

Once authenticated, create the service principal:

#### Input:

```
az ad sp create-for-rbac --name SpNameExample --role "Website Contributor"
```

#### Output:

```
{
"appId": "f5f1ce7d-c247-42e6-91a4-ad1e7example",
"displayName": "SpNameExample",
"name": "http://SpNameExample",
"password": "97adeba6-4178-4f2b-bf5f-782b3example",
"tenant": "874186fd-a7a8-4e98-9b9e-3df00example"
}
```

The results above contain JSON output of the service principal name, password, and tenant that you'll need for setting up Snyk. From there you can login to your Snyk account and paste in your name, password, and tenant. More information on creating a service principal for Azure can be found in Azure's documentation: [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) , [Portal](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-create-service-principal-portal)

## **Azure Functions: check your connection status**

At any time after you've entered your Azure service principal, you can check on the connection status in one of two places.

The first is on your integration settings page, where you'll see your current integrations listed as well as their connection status.

![](<../../../.gitbook/assets/image (25).png>)

You can also check the status directly on the Azure Functions integration settings page (found by clicking "**Edit settings**" on the integration settings page shown above). If you've entered credentials, you'll see a box indicating whether or not Snyk is able to correctly connect to Azure:

![](<../../../.gitbook/assets/image (31).png>)

If you are unable to connect, re-enter your account credentials to verify that they are correct:

![](<../../../.gitbook/assets/image (27).png>)

## **Disable the Azure Functions integration**

If you decide to disable this integration for any reason, you can accomplish this from the Integrations page in your Settings.

You need to find the specific integration you wish to deactivate in your list of integrations and click Edit settings. You are taken to a page that shows the current status of your integration, a place to update your credentials, specific to each integration (credentials, API key, Service Principal, or connection details), and a red box at the bottom to disconnect this integration, like in the example seen below:

![](<../../../.gitbook/assets/image (26).png>)

If you choose to disconnect, your credentials will be removed from Snyk and any integration-specific projects we had been monitoring will be deactivated on Snyk.

If you choose to re-enable this integration at any time, you need to re-enter your credentials and activate your projects.

## **Add Azure Functions to Snyk**

Once you've successfully connected Snyk to your Azure account, you'll be able to select Azure functions that you would like Snyk to monitor. You can do this either using the **"Add projects"** button on the integrations page, or directly from the Azure Functions integration settings page.

In either case, you'll see a list of any available Function apps on the Azure account you connected. Select the ones you want to monitor and click the "**Add to Snyk**" button.

{% hint style="warning" %}
**NOTE:** Current support exists for importing **v2 functions only--all functions will be ignored.**
{% endhint %}

![](<../../../.gitbook/assets/image (30).png>)

As soon as you've added the projects to Snyk, Snyk will test them and begin to display a list of all monitored Azure functions in your [project dashboard](https://app.snyk.io/projects). You'll also see a snapshot of any current vulnerabilities, and be able to click through for a more detailed report including any steps to fix:

![](<../../../.gitbook/assets/image (32).png>)

Snyk will now continuously monitor each of those functions for known vulnerabilities. You can add more functions at any time.

**NOTE**\
for Node.js and .Net, dev dependencies will be ignored.
