# Connect a third-party integration

{% hint style="warning" %}
The third-party integrations are available in a Closed Beta state and are applicable only to the Snyk AppRisk Pro version.  Please contact your salesperson if you are interested in Snyk AppRisk Pro.
{% endhint %}

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

You can customize your AppRisk integrations from the **Integrations Hub** where the following SAST and Secrets integrations are available:

SAST:

* Verocode SAST
* Checkmarx SAST&#x20;

Secrets:

* Nightfall

## Veracode setup guide

### Prerequisites

The Veracode application concept is matched into the Snyk AppRisk repository assets. You need to create and utilize the Veracode custom field by using the [Veracode API](https://app.swaggerhub.com/apis/Veracode/veracode-applications\_api\_specification/1.0#/Application%20information%20API/updateApplicationUsingPUT). More details about the Veracode custom field are available [here](https://docs.veracode.com/r/t\_create\_custom\_metadata).&#x20;

Ensure you are adding a custom field called repoURL:

```
{
"name": "repoURL", 
"value": <YOUR GITHUB URL>
}

```

### Required parameters

* API ID and API Key - API credentials associated with a user account. For more information, access this link: [https://help.veracode.com/r/c\_api\_credentials3](https://help.veracode.com/r/c\_api\_credentials3).

### &#x20;Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select the **SAST** tag and search for Veracode.&#x20;
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the **API ID** from your Veracode account.
6. Add the **API key** from your Veracode account.
7. Click the **Done** button.
8. When the connection is established, the status of the Veracode integration is changed to **Connected**.

<figure><img src="../../../.gitbook/assets/image (1).png" alt="Integration Hub - Veracode setup"><figcaption><p>Integration Hub - Veracode setup</p></figcaption></figure>

## Checkmarx setup guide

Use the following instructions to set up your Checkmarx SAST integration. Checkmarx SAST integration is only working for the Checkmarx SAST, we are not yet supporting Checkmarx One.

{% hint style="info" %}
Snyk AppRisk Pro does not currently support the Checkmarx One integration.
{% endhint %}

### Prerequisites

* Install and configure your [Snyk Broker](../../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md#checkmarx-sast-integration) connection for Snyk AppRisk.&#x20;

### Required parameters

1. API URL - The URL of Checkmarx API, for example, `checkmarx.customer.com`.
2. Username and Password - Credentials for a user account with Checkmarx SAST access.

### Integration Hub setup

After you have installed and configured Snyk Broker for AppRisk and you successfully established a connection for Checkmarx SAST, you need to also configure the integration from the Snyk AppRisk Integration Hub. &#x20;

1. Open the **Integration Hub** menu.&#x20;
2. Select the **SAST** tag and search for Checkmarx.&#x20;
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the Broker token for the Snyk AppRisk Checkmarx integration.
6. Add the Checkmarx host. E.g. `checkmarx.customer.com`
7. Click the **Done** button.
8. When the connection is established, the status of the Checkmarx integration is changed to **Connected**.

<figure><img src="../../../.gitbook/assets/image (370).png" alt="Integration Hub - Checkmarx setup"><figcaption><p>Integration Hub - Checkmarx setup</p></figcaption></figure>

## Nightfall setup guide

### Required parameters

* API Key. [Here](https://docs.nightfall.ai/docs/creating-an-api-key) you can find more details about how to create a Nightfall API key.&#x20;

### Integration Hub setup



1. Open the **Integration Hub** menu.&#x20;
2. Select the **Secrets** tag and search for Veracode.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** for this integration.
5. Add the **Base API URL** for this integration.
6. Add the **API Key** for this integration.
7. Click the **Done** button.
8. When the connection is established, the status of the Nightfall integration is changed to **Connected**.

<figure><img src="https://lh7-us.googleusercontent.com/FhueUeoeYL3j5ihKIc6lFSFuvoS43NU30PiSCIOpkfWkoPXI8XaLh_ABzR5DipTn-TVLAkX3pSmoBfdq6hNDL65JPV72E2lErTcTua9YTcUnFXzwVI_smOtlZqS4LzD2-XqGZHlygjVZHVlba4ju3pQ" alt="Integration Hub - Nightfall setup"><figcaption><p>Integration Hub - Nightfall setup</p></figcaption></figure>
