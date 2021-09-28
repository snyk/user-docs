# Connect to Snyk

At this stage of the configuration, you will be prompted to authenticate. Snyk supports the following identity providers for authentication:

* GitHub
* Bitbucket
* Google
* Azure
* Single sign-on \(SSO\)

Additional details on authentication may be found on our [documentation pages](https://support.snyk.io/hc/en-us/articles/360004008218-Authentication).

{% hint style="info" %}
It's important to call out that this integration is **FREE**! We take care of creating your free account as part of this configuration without any additional steps on your part. 
{% endhint %}

![](../../../.gitbook/assets/snyk-codepipeline-07.png)

Having selected your identity provider, you will receive a confirmation message to **Authorize** the app. Click this _once_ and proceed.

![](../../../.gitbook/assets/snyk-codepipeline-08.png)

The next step will be to configure the integration settings. Here you will select the **Snyk Organization** where scan results will be stored as well as **vulnerability handling** and threshold to **block deployment** by **severity**. These settings will differ based on your specific use case and organizational policies. 

![](../../../.gitbook/assets/snyk-codepipeline-09.png)

For the purpose of this lab, however, we will not block deployments and click **Continue**.

![](../../../.gitbook/assets/snyk-codepipeline-10.png)

Click **Confirm** to complete the authorization request.

![](../../../.gitbook/assets/snyk-codepipeline-11.png)

You will notice a confirmation that the connection was successfully established and at this point can click on **Done**.

If at any point in time, you need to revoke the authorization, you can do so by [logging into Snyk](https://app.snyk.io/) and clicking **Revoke** from the list of **Authorized Applications** under **Account Settings** as shown below.

![](../../../.gitbook/assets/snyk-codepipeline-12.png)

Finally, let's confirm our selections and click **Save** to complete our pipeline configuration.

![](../../../.gitbook/assets/snyk-codepipeline-13.png)

...and acknowledge the changes.

![](../../../.gitbook/assets/snyk-codepipeline-14.png)

Let's move on!

