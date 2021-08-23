# Container security with Nexus integration

{% hint style="info" %}
**Feature availability**  
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk integrates with Nexus Container Registry to enable you to import your container images and monitor them for vulnerabilities. 

Snyk tests the images you’ve imported \(referred to as \`projects\`\) for any known security vulnerabilities, testing them at a frequency you control and alerts you when new issues are detected.

To set up Nexus integration in Snyk and start managing image vulnerabilities: 

Prerequisites

* You must be an administrator for the organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with Nexus and does not support Nexus when configured for single sign-on \(SSO\).

**Configure integration**

1. In your Snyk account, navigate to **Integrations** from the menu bar at the top. ****Under the **Container Registries** section, find the **Nexus** option and click it.![mceclip0.png](../../../.gitbook/assets/mceclip0-9-.png) In the **Account credentials** section, enter your Nexus username and password login credentials. In the **container registry name** fill in the full URL to the registry you want to integrate with. To finish, click **Save**. ![mceclip1.png](https://support.snyk.io/hc/article_attachments/4403626381457/mceclip1.png)  In case you are using a self-hosted Nexus registry, contact us to provide you with a token. You can read more about setting up private registry integration [here](https://support.snyk.io/hc/en-us/articles/360017040957).  Snyk tests the connection values and the page reloads, now displaying Nexus integration information, and the **Add your Nexus images to Snyk** button becomes available. In case the connection to Nexus failed, notification appears under the **Connected to Nexus** section. Now you can use Snyk to scan your images from Nexus.

