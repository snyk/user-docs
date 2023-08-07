# Configure Azure Function environment variables

For more information about configuration of Azure Function environment variables, see the Microsoft Azure Functions documentation article [Manage your function app](https://learn.microsoft.com/en-us/azure/azure-functions/functions-how-to-use-azure-function-app-settings?tabs=portal).

Navigate to your Azure Function App Configuration and add the following new application settings:

* **NEW\_RELIC\_SECURITY\_URL**: URL for the New Relic Security API, that is, https://security-api.newrelic.com/security/v1
* **NEW\_RELIC\_LICENSE\_KEY**: New Relic License Key

If you are using Azure DevOps Repos in your Snyk account, you may also want to configure the following application setting:

* **AZURE\_DEVOPS\_ORG**: the name of your Azure DevOps organization

If you want to use the optional parameters for troubleshooting in a separate New Relic event, configure the following application settings:

* **NEW\_RELIC\_INSIGHTS\_URL**: URL for the New Relic account event API, that is, https://insights-collector.newrelic.com/v1/accounts/{NR-ACCOUNT-ID}/events

<figure><img src="../../../.gitbook/assets/azure-function-configuration.png" alt="Azure Function Configuration"><figcaption><p>Azure Function Configuration</p></figcaption></figure>

Next [Copy the Azure Function URL](copy-the-azure-function-url.md).
