# Install the Snyk extension for your Azure pipelines

To start using the Snyk task as part of your pipeline build, first install the extension into your Azure DevOps instance per organization, from the [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan).

## **Prerequisites**

* Create a Snyk account at [https://snyk.io/](https://snyk.io)
* Ensure you are an owner of or an administrator for this account.

## **Steps**

1. Access your Snyk account.
2. For free plans, go to your **General Account Settings** and find, copy and save your personal API authentication token on the side.
3. For paid plans, navigate to the organization where you want to integrate; then go to **Settings** to create a new service account token. Copy and save it on the side.
4. Access your Azure DevOps account and navigate to **Extensions -> Browse marketplace.**
5. Search for the **Snyk Security Scan** extension, click **Get it free**.
6. Create a new **Service Connection** in your project via **Project Settings** —> **Pipelines** —> **Service Connections**
7. Select the **Snyk Authentication** service connection:
8. In the Snyk Authentication service connection form, enter the **Server URL** and the **Snyk API Token** along with a **Service connection name**.
9. Click **Save**, ensuring the new service connection appears in your list of service connections.

![Create your first service connection](../../../.gitbook/assets/ap\_-\_search.jpg)

![New Snyk authentication service connection](../../../.gitbook/assets/ap\_-\_config.jpg)
