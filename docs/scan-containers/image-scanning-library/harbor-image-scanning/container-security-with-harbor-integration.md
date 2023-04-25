# Container security with Harbor integration

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk integrates with Harbor Container Registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you have imported (referred to as Projects) for any known security vulnerabilities, testing them at a frequency you control, and alerts you when new issues are detected.

Follow these instructions to set up Harbor integration in Snyk and start managing image vulnerabilities.:

## Prerequisites for Harbor integration

* You must be an administrator for the organization you are configuring in Snyk.
* Snyk needs user credentials to integrate with Harbor and does not support Harbor when configured for single sign-on (SSO).

## **Configure integration**

1. In your Snyk account, navigate to **Integrations** from the menu bar at the top. Under the **Container Registries** section, select the Harbor option.
2. In the **Account credentials** section, enter your Harbor username and password login credentials.
3. In the **Container registry name** fill in the full URL to the registry you want to integrate with.
4. To finish, select **Save**.

If you are using a self-hosted Harbor registry, contact Snyk to provide you with a token. For more information about setting up private registry integration see [Snyk Container for self-hosted container registries (with broker)](../../integrate-self-hosted-container-registries.md).

<figure><img src="../../../.gitbook/assets/mceclip2-1-.png" alt="Harbor Account credentials and Container registry name"><figcaption><p>Harbor Account credentials and Container registry name</p></figcaption></figure>

{% hint style="info" %}
**Note**\
To set up the integration, the Harbor user should be an admin user. It currently uses /v2/\_catalog endpoint for listing repos.
{% endhint %}

<figure><img src="../../../.gitbook/assets/mceclip1-8-.png" alt="Harbor Account credentials with Broker token"><figcaption><p>Harbor Account credentials with Broker token</p></figcaption></figure>

Snyk tests the connection values and the page reloads, now displaying Harbor integration information, and the **Add your Harbor images to Snyk** button becomes available. In case the connection to Harbor failed, notification appears under the **Connected to Harbor** section.\
Now you can use Snyk to scan your images from Harbor.
