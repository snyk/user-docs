# Configuring your JFrog Artifactory Container Registry integration

The instructions on this page explain how to enable integration between one Artifactory instance as a container registry and a Snyk organization to start managing your image security.

## Prerequisites

* You must be an administrator for the organization you are configuring in Snyk.
* Snyk needs user credentials to integrate with Artifactory and does not support Artifactory when configured for single sign-on (SSO).
* If you are using a self-hosted Artifactory instance, refer to Snyk Broker - Container Registr Agent.

## Configure integration

1. Log in to [your Snyk account](https://app.snyk.io).
2.  Navigate to **Integrations**; select the **Artifactory** option:

    <img src="../../../.gitbook/assets/image (57) (1).png" alt="Artifactory integration" data-size="original">
3. The configuration page in the **Settings** area loads.
4. Enter credentials as follows:
   * **Username and Password**—use your Artifactory login credentials.
   * **Container registry name -** the _full registry URL_ in format `<org>.jfrog.io/artifactory/api/docker/<repo-name>`.
5. Click **Save Changes**; a confirmation appears.

{% hint style="info" %}
To set up the integration, the Artifactory credentials need at minimum read permissions to the relevant Artifactory repository.
{% endhint %}

<figure><img src="https://user-images.githubusercontent.com/112600/144875482-078b715e-2834-469b-9983-7e88a65f175e.png" alt="Artifactory account credentials"><figcaption><p>Artifactory account credentials</p></figcaption></figure>

Snyk tests the connection values and the page reloads, now displaying integration details as you entered them. A confirmation message that the details were saved also appears in green at the top of the screen. If the connection failed, a notification appears.

<figure><img src="../../../.gitbook/assets/uuid-3b329a90-394f-5ab3-af84-658b41a1edc0-en.png" alt="Artifactory setup confirmation"><figcaption><p>Artifactory setup confirmation</p></figcaption></figure>
