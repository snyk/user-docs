# Configure integration for ACR

This page explains how to enable integration between a Microsoft Azure Container Registry (ACR) registry and a Snyk Organization and start managing your vulnerabilities. To integrate with multiple registries, create a unique Organization for each one.

Follow these steps:

1. Access your ACR account and retrieve unique service principal credentials for use by Snyk with the AcrPull role.
2. Log in to your Snyk account.
3. Navigate to **Integrations** from the menu bar at the top, find, and click the **ACR** option.\
   The ACR configuration page in the **Settings** area loads.
4. Enter the user name, password, and container registry name (myregistry.azurecr.io) that you received when you generated a service principal for this integration.
5.  Click **Save**.

    Snyk tests the connection values, and the page reloads, now displaying ACR integration information.\
    A confirmation message that the details were saved also appears in green at the top of the screen.\
    In addition, if the connection to Azure failed, a notification appears under the **Connected to ACR section**.
