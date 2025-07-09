# Obtain the tokens required to set up Snyk Broker

{% hint style="info" %}
Snyk recommends rotating all API tokens and credentials used with Snyk Broker every 90 days.
{% endhint %}

By obtaining the tokens required to set up Snyk Broker, you generate the credentials for the Broker's target application. When this is complete for the Organizations where you want to use Broker, configure the environment variables for launching Snyk Broker. See [Enabling Broker across multiple Organizations](enabling-broker-across-multiple-organizations.md) if you want to deploy across Organizations.

To set up Snyk Broker, you must have the following tokens:

* **Broker token** - This token is **required for the Broker client setup**. It is used in the `-e BROKER_TOKEN` parameter. The Broker token is associated with a specific Organization by default and with a specific integrated SCM and enables Snyk Broker deployment for this Organization and SCM. A different Broker token is required for each SCM. Details are on this page.
* **Integrated SCM token** - This token is required for the Broker Client setup. It is used in the `-e <SCM>_TOKEN` parameter, for example, `-e GITHUB_TOKEN=xxx…`, to enable access to the SCM with certain permissions needed for the operation of the Broker and Snyk Code. For details, see [Integrated SCM tokens for Snyk Broker](../../../../scm-integrations/scm-integrations-and-snyk-broker.md#integrated-scm-tokens-for-classic-broker).

After you have obtained the required tokens, save them in a safe and accessible place for use in setting up the Broker Client.

## **Generate your Broker token using the Snyk API**

For code repository (SCM) integrations, you can generate a Broker token by using the API or by contacting [Snyk support](https://support.snyk.io).

1. Use the endpoint [Update Existing Integration](../../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-type) to enable Snyk Broker for a specific Organization and a specific SCM. Follow the example under "Set up a broker for an existing integration." This generates a Broker token in the UI.
2. To generate a Broker token programmatically after enabling Snyk Broker, use the endpoint [Provision new Broker token](../../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-authentication-provision-token) to generate a Broker token.\
   You can see the generated Broker token in the API response body and on the Web UI.
3. Verify the Broker token is generated in the Snyk Web UI under the specified SCM integration. by selecting **Settings** > **Integrations** for that specific integration to see the Broker token.
4. After generating the Broker token, copy and save it and store it in a secure location for future use, or obtain it later using the Web UI.

## **Generate a Broker token in the Web UI**

For [Artifactory Repository](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/) and [Nexus Repository Manager](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) brokered instances or [Jira](../install-and-configure-snyk-broker/jira-prerequisites-and-steps-to-install-and-configure-broker/jira-install-and-configure-using-docker.md) integration, you can obtain a Broker token in the Snyk UI or contact [Snyk support](https://support.snyk.io). The steps to generate a Broker token in the Web UI follow:

1. Select **Settings** > **Integrations** for that specific integration to generate the Broker token.
2. After the Broker token is generated, under the integration, the notification from this screen correctly displays **Could not connect to…**, as you have not yet installed and configured the client.
3. Copy and paste the Broker token from the UI to use it when you install the client.

## Obtain your Broker token from the Web UI

After your Broker token is generated, it is displayed on the Web UI. Follow these steps to obtain the token:

1. In the Snyk Web UI, select the **Organization** for which you want to set up the Snyk Broker.
2. In the selected Organization, select **Integration**. Find the Integration to which you want to connect Snyk Broker, and click the **Settings** icon.
3. On the **Settings** page of the selected Integration, in the **Broker Credentials** section, copy the Broker token from the **Token** box and save it for future use:

<figure><img src="../../../../.gitbook/assets/Snyk Broker - Broker Token - box.png" alt=""><figcaption><p>Copy the Broker token</p></figcaption></figure>

