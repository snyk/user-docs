# Install the Snyk extension for your Azure pipelines

To start using the Snyk task as part of your pipeline build, from the [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan), install the extension into your Azure DevOps instance for your Organization.

## **Prerequisites for installing Snyk extension for Azure pipelines**

* Create a Snyk account at [https://snyk.io/](https://snyk.io)
* Ensure you are an owner or administrator of this account.

## **Step for installing the Snyk extension for Azure pipelines**

1. Access your Snyk account.
2. Token:
   1. For **free plans**, go to your **General Account Settings** and find, copy, and save your personal API authentication token on the side.
   2. For **paid plans**, navigate to the Organization where you want to integrate; then go to **Settings** to create a new service account token. Copy and save it on the side.
3. Access your Azure DevOps account and navigate to the **Extensions -> Browse marketplace.**
4. Search for the **Snyk Security Scan** extension and click **Get it free**.
5. Create a new **Service Connection** in your Project via **Project Settings** —> **Pipelines** —> **Service Connections**
6. Select the **Snyk Authentication** service connection:
   1. In the Snyk Authentication service connection form, enter the **Snyk API Token**.
   2. Click **Save**, ensuring the new service connection appears in your list of service connections.

<figure><img src="../../../.gitbook/assets/ap_-_search.jpg" alt="Create your first service connection"><figcaption><p>Create your first service connection</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/Azure-authentication-connection-setup.png" alt="New Snyk authentication service connection"><figcaption><p>New Snyk authentication service connection</p></figcaption></figure>



{% hint style="info" %}
If you are using the Snyk EU or AU data center, `you` must add another line above the -task lines as follows.

For EU: `- script: snyk config set endpoint=https://api.eu.snyk.io/v1/`

For AU: `- script: snyk config set endpoint=https://api.au.snyk.io/v1/`
{% endhint %}
