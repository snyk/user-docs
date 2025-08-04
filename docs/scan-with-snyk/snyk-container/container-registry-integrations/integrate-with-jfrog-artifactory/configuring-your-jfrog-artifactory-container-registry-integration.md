# Configuring your JFrog Artifactory Container Registry integration

The instructions on this page explain how to enable integration between one Artifactory instance as a container registry and a Snyk Organization to start managing your image security.

## Prerequisites for Artifactory Container Registry integration

* You must be an administrator for the Organization you are configuring in Snyk.
* You must provide user credentials to integrate with Artifactory.&#x20;
* You must be running Docker. Snyk supports Docker repositories and the Docker package type for this integration.
* If you are using a self-hosted Artifactory instance, see [Snyk Broker - Container Registry Agent](../../../../implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/).

## Configure Artifactory Container Registry integration&#x20;

1. Log in to [your Snyk account](https://app.snyk.io).
2. Navigate to **Integrations**. From the list of integrations, select **Artifactory.** \
   The configuration page opens.
3. Enter your credentials:
   * **Username and Password  -** use your Artifactory login credentials. If you're using SSO configuration, you must generate an access token in Artifactory and use the token details to login.
   * **Container registry name -** the full registry URL: `<subdomain>.jfrog.io/artifactory/api/docker/<repo-name>`

<figure><img src="https://user-images.githubusercontent.com/112600/144875482-078b715e-2834-469b-9983-7e88a65f175e.png" alt="" width="375"><figcaption><p>Artifactory account credentials</p></figcaption></figure>

4. Click **Save Changes**.  A confirmation appears.

{% hint style="info" %}
To set up the integration, the Artifactory credentials need, at minimum, read permissions to the relevant Artifactory repository.
{% endhint %}

Snyk tests the connection values and the page reloads, now displaying integration details as you entered them. At the top of the screen, a confirmation message indicates that the details were saved. If the connection fails, a notification appears.
