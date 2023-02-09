# Configuring your JFrog Artifactory container registry integration

Enable integration between one Artifactory instance as a container registry and a Snyk organization to start managing your image security.

## Prerequisites

* You must be an administrator for the organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with Artifactory and does not support Artifactory when configured for single sign-on (SSO).
* If you're using a self-hosted Artifactory, refer to the Snyk documentation for [Broker configuration](../../../integrations/snyk-broker/snyk-broker-container-registry-agent/).

## Configure integration

1. Log in to [your Snyk account](https://app.snyk.io).
2.  Navigate to **Integrations** from the menu bar at the top; find and click the Artifactory option:

    <img src="../../../.gitbook/assets/image (57) (3).png" alt="Artifactory icon" data-size="original">
3. The configuration page in the **Settings** area loads.
4. Enter credentials as follows:
   1. **Username and Password**—use your Artifactory login credentials.
   2. **Container registry name -** the _full registry URL_ in format `<org>.jfrog.io/artifactory/api/docker/<repo-name>`.
5. Click **Save Changes**.

<figure><img src="https://user-images.githubusercontent.com/112600/144875482-078b715e-2834-469b-9983-7e88a65f175e.png" alt="Artifactory account credentials"><figcaption><p>Artifactory account credentials</p></figcaption></figure>

{% hint style="info" %}
**Note**\
\*\*\*\*To set up the integration, the Artifactory credentials should have read at least permissions to the relevant Artifactory repository.
{% endhint %}

![](../../../.gitbook/assets/uuid-3b329a90-394f-5ab3-af84-658b41a1edc0-en.png)

Snyk tests the connection values and the page reloads, now displaying integration details as you entered them. A confirmation message that the details were saved also appears in green at the top of the screen. In addition, if the connection failed, a notification appears.
