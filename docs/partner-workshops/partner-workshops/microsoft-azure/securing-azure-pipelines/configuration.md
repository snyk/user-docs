# Configuration

Once **Snyk Security Scan** extension for Azure DevOps is installed, we will need to perform a few simple steps to configure it so that it can communicate with your Snyk account. To do this, we will create a **Snyk Service Account** \(preferred\) or use your **Snyk** **Organization ID** if you are on a free-tier plan.

{% hint style="info" %}
You can learn more about Snyk Service Accounts by visiting our [documentation page](https://support.snyk.io/hc/en-us/articles/360004037597-Service-accounts).
{% endhint %}

You can obtain the token by logging into [**Snyk**](https://app.snyk.io) and clicking on the _gear_ icon in the upper-right navigation bar as shown below.

![](../../../.gitbook/assets/azure-devops-11.png)

Under the **Settings** menu, you will select the **General** tab and then either **Click** the copy button if using the Org ID or **Manage service accounts** \(recommended\) if you are on a paid plan.

Now, we will move back to our Azure DevOps browser tab to complete the configuration. You can add or change your token by navigating to your Azure Pipelines **Project Settings** and selecting **Service Extensions** as shown below:

![](../../../.gitbook/assets/azure-devops-05.png)

From there you can configure the **Service URL**, input the **Snyk API Token**, and name your connection. Please refer to the image below as an example:

![](../../../.gitbook/assets/azure-devops-06.png)

When complete, click **Save**. Now, let's move onto the next section!

