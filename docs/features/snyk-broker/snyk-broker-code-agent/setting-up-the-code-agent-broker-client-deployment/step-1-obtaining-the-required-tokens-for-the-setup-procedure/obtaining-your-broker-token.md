# Obtaining your Broker token

A Broker token is required for [the Broker Client component setup](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display), and it is used in the `-e BROKER_TOKEN` parameter. The Broker token is associated with a specific Organization by default and with a specific integrated SCM, and it enables for them the Snyk Broker deployment method. For each integrated SCM, a different Broker token is required.

* **Using an existing Broker token for the Code Agent setup** - if you already have a Broker token, which you used for running the Broker Client for another Snyk product in the same Organization and the same SCM, you can also use it for the setup of the Broker Client for the Code Agent.
* **Using the same Broker token for multiple Snyk Organizations** -\
  Although by default a Broker token is associated with only one Snyk Organization, by performing the following procedures you can use the same Broker token for multiple Organizations. To perform these procedures, you need Organizations Admin permissions: &#x20;
  * **A new Organization** - if you create a new Organization based on an existing Organization that has a Broker token, the existing Broker token will be cloned during the new Organization creation, and you can use it for the new Organization as well.
  * **An existing Organization** – if you want to use an existing Broker token for other existing Organizations, you can run the [Clone an integration](https://snyk.docs.apiary.io/#reference/integrations/integration-cloning/clone-an-integration-\(with-settings-and-credentials\)) API. This API clones existing Integration Settings, including the Broker token of the Integration.
* **Using a Broker token for redundancy** - if you set up two Broker Clients for the same Organization and the same SCM for redundancy purposes, you must use the same Broker token for both Broker Clients.You can obtain the Snyk Broker token in the following ways:

**Note**: It is recommended to ask your Snyk IC/TSM to generate for you a Broker token, and then to obtain it from the Web UI.

* Receive the Broker token from your Snyk IC/TSM.
* Generate the Broker token via Snyk APIs (see below).

In both ways, after a Broker token is generated, it is displayed on the Web UI, and you can [obtain it from there](obtaining-your-broker-token.md#obtaining-your-broker-token-via-the-web-ui).&#x20;

&#x20;&#x20;

### **Generating your Broker token via APIs**

You can generate the Broker token by using two APIs. After the Broker token will be generated, it will be displayed both in your API tool and on the Web UI.&#x20;

**To obtain the Broker token via Snyk APIs:**

1\.  Run the **Update Existing Integration** API, which enables the Snyk Broker for a specific Organization and a specific SCM:\
[https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration)

2\.  Run the **Provision new Broker token** API, which generates a Broker token:\
[https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-provisioning/provision-new-broker-token](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-provisioning/provision-new-broker-token)

The generated Broker token will appear in two places:

* On the Response body in your API tool.
* On the Web UI (see below).

3\.  Copy and save the Broker token in a secure location for future use, or obtain it later via the Web UI.

### **Obtaining your Broker token via the Web UI**

Your Broker token is displayed on the Web UI, only after your Snyk IC/TSM generated it for you, or you generated it by yourself via the Snyk APIs (see above).

&#x20;**To obtain your Broker token via the Web UI:**

1\.  After you received a Broker token from your IC/TSM, or generated it via the API, open the Snyk Web UI. Then, open the Organization for which you want to set up the Snyk Broker:

<figure><img src="../../../../../.gitbook/assets/Snyk Broker - Organization - Select (1).png" alt=""><figcaption></figcaption></figure>

**Note**: The Broker token is associated with a specific Organization and a specific Integration.

2\.  Once the required Organization is open, click the **Integration** tab on the top menu. Then, locate the Integration to which you want to connect the Snyk Broker, and click its **Settings** button <img src="../../../../../.gitbook/assets/Snyk Broker - Organization - Integrations - Settings Icon.png" alt="" data-size="line">  on the upper right corner:

<figure><img src="../../../../../.gitbook/assets/Snyk Broker - Organization - Integrations page.png" alt=""><figcaption></figcaption></figure>

3\.  On the **Settings** page of the selected Integration > **Broker Credentials** section, copy the Broker token from the **Token** box, and save it for future use:&#x20;

<figure><img src="../../../../../.gitbook/assets/Snyk Broker - Broker Token - box.png" alt=""><figcaption></figcaption></figure>

&#x20;
