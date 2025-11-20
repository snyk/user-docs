# Third-party integrations for Snyk AppRisk

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that are automatically synced, and provides access to the Integration Hub.

{% hint style="info" %}
The Loaded package risk factor is not supported by Snyk for operating system packages (such as Debian packages), only for packages hosted under package managers such as npm, Maven, or PyPI.
{% endhint %}

You can customize your AppRisk integrations from the **Integrations Hub** where the following integrations are available:

* [Veracode SAST](connect-a-third-party-integration.md#veracode-setup-guide)
* [Checkmarx SAST](connect-a-third-party-integration.md#checkmarx-setup-guide)
* [SonarQube](connect-a-third-party-integration.md#sonarqube-setup-guide)
* [Nightfall](connect-a-third-party-integration.md#nightfall-setup-guide)
* [GitGuardian](connect-a-third-party-integration.md#gitguardian-setup-guide)
* [Dynatrace](connect-a-third-party-integration.md#dynatrace-setup-guide)
* [Sysdig](connect-a-third-party-integration.md#sysdig-setup-guide)
* [Crowdstrike](connect-a-third-party-integration.md#crowdstrike-setup-guide)

{% hint style="info" %}
Data synchronization may take up to two hours after receiving the **Connected** status from a new integration setup. See the [Integration connected statuses](integrate-with-snyk.md#integration-connection-statuses) page for more details.
{% endhint %}

After you finish setting up an integration, you can see it listed with a Connected status.

<figure><img src="../.gitbook/assets/image (223).png" alt=""><figcaption><p>Third-party integration card - Connected status</p></figcaption></figure>

## Veracode setup guide

{% hint style="info" %}
**Release status**

Veracode for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Prerequisites <a href="#verocode-prerequisites" id="verocode-prerequisites"></a>

The Veracode application concept is matched to the Snyk AppRisk repository assets. You need to create and utilize the Veracode custom field by using the [Veracode API](https://app.swaggerhub.com/apis/Veracode/veracode-applications_api_specification/1.0#/Application%20information%20API/updateApplicationUsingPUT). Access the [Veracode custom metadata field](https://docs.veracode.com/r/t_create_custom_metadata) for more details.

Ensure you are adding a custom field called repoURL:

```
{
"name": "repoURL", 
"value": <YOUR GITHUB URL>
}

```

### Required parameters <a href="#veracode-required-parameters" id="veracode-required-parameters"></a>

* API ID and API Key - API credentials associated with a user account. For more information, access the [Veracode API credentials](https://help.veracode.com/r/c_api_credentials3) link.

### Integration Hub setup <a href="#veracode-integration-hub-setup" id="veracode-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.
2. Select the **SAST** tag and search for Veracode.
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the **API ID** from your Veracode account.
6. Add the **API key** from your Veracode account.
7. Click the **Done** button.
8. When the connection is established, the status of the Veracode integration is changed to **Connected**.

<figure><img src="../.gitbook/assets/image (206).png" alt=""><figcaption><p>Veracode - Setup screen</p></figcaption></figure>

## Checkmarx setup guide

{% hint style="info" %}
**Release status**

Checkmarx for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

Use the following instructions to set up your Checkmarx SAST integration. Checkmarx SAST integration is only working for the Checkmarx SAST, we are not yet supporting Checkmarx One.

{% hint style="warning" %}
Snyk AppRisk does not support the Checkmarx One integration.
{% endhint %}

### Prerequisites <a href="#checkmarx-prerequisites" id="checkmarx-prerequisites"></a>

* Install and configure your [Snyk Broker](broken-reference) connection for Snyk AppRisk.
* Ensure you have properly used Git Setting for your Checkmarx Project. For more details, see the Checkmarx [Set project's remote source settings as GIT](https://checkmarx.stoplight.io/docs/checkmarx-sast-api-reference-guide/8312d35369b9b-set-project-s-remote-source-settings-as-git) documentation page.

### Required parameters <a href="#checkmarx-required-parameters" id="checkmarx-required-parameters"></a>

1. API URL - The URL of Checkmarx API, for example, `checkmarx.customer.com`.
2. Username and Password - Credentials for a user account with Checkmarx SAST access.

### Integration Hub setup <a href="#checkmarx-integration-hub-setup" id="checkmarx-integration-hub-setup"></a>

After you have installed and configured Snyk Broker for AppRisk and have successfully established a connection for Checkmarx SAST, you also need to configure the integration from the Snyk AppRisk Integration Hub.

1. Open the **Integration Hub** menu.
2. Select the **SAST** tag and search for Checkmarx.
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the Broker token for the Snyk AppRisk Checkmarx integration.
6. Add the Checkmarx host. For example `checkmarx.customer.com`
7. Click the **Done** button.
8. When the connection is established, the status of the Checkmarx integration is changed to **Connected**.

<figure><img src="../.gitbook/assets/image (207).png" alt=""><figcaption><p>Checkmarx - Setup screen</p></figcaption></figure>

## SonarQube setup guide

{% hint style="info" %}
**Release status**

SonarQube for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#sonarqube-required-parameters" id="sonarqube-required-parameters"></a>

* API Key. [Here](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/) you can find more details about the SonarQube API Key.

### Integration Hub setup <a href="#sonarqube-integration-hub-setup" id="sonarqube-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **SAST** tag and search for SonarQube.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Host URL** for this integration.
* Add the **API token**. Navigate to your SonarQube account, select User, select My Account, select Security, and then User Token. Access the SonarQube [generating and using tokens](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/) documentation page for more details about the SonarQube API Key.
* Click the **Done** button.
* When the connection is established, the status of the SonarQube integration is changed to **Connected**.

<figure><img src="../.gitbook/assets/image (208).png" alt=""><figcaption><p>SonarQube - Setup screen</p></figcaption></figure>

## Nightfall setup guide

{% hint style="info" %}
**Release status**

Nightfall for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#nightfall-required-parameters" id="nightfall-required-parameters"></a>

* API Key. Access the Nightfall [Creating an API Key](https://help.nightfall.ai/nightfall-firewall-for-ai/key-concepts/setting-up-nightfall/creating-api-key) documentation page for more details about how to create a Nightfall API key.

### Integration Hub setup <a href="#nightfall-integration-hub-setup" id="nightfall-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.
2. Select the **Secrets** tag and search for Nightfall.
3. Click the **Add** button.
4. Add the **Profile name** for this integration.
5. Add the **Base API URL** for this integration.
6. Add the **API Key** for this integration.
7. Click the **Done** button.
8. When the connection is established, the status of the Nightfall integration is changed to **Connected**.

After you set up your Nightfall integration using the Integration Hub, you can see the secrets detection coverage.

The following video explains how to configure Nightfall.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1740510692/snyk-learn/product-training-videos/Snyk_AppRisk_-11a_-_v1_-_Secrets_detection_coverage_with_Nightfall_AI.mp4" %}
Configure Nightfall
{% endembed %}

The following video explains how Nightfall uses data in Snyk.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1740510717/snyk-learn/product-training-videos/Snyk_AppRisk_-11b_-_v1_-_Secrets_detection_coverage_with_Nightfall_AI_-_Snyk_Interface.mp4" %}
How Nightfall uses data in Snyk
{% endembed %}

## GitGuardian setup guide

{% hint style="info" %}
**Release status**

GitGuardian for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#gitguardian-required-parameters" id="gitguardian-required-parameters"></a>

* API Key. Access the GitGuardian [authentication](https://docs.gitguardian.com/api-docs/authentication) documentation page for more details about how to create a GitGuardian API Key.

When you create a GitGuardian API Key, remember that it works for both service accounts and personal access token.

Ensure that the following permissions are set as READ:

* Incident (`mandatory`)
* Teams (`recommended` for GitGuardian paid accounts)
* Sources (`mandatory`)

### Integration Hub setup <a href="#gitguardian-integration-hub-setup" id="gitguardian-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.
2. Select the **Secrets** tag and search for GitGuardian.
3. Click the **Add** button.
4. Add the **Profile name** for this integration.
5. Add the **API Token** for this integration.
6. Click the **Done** button.
7. When the connection is established, the status of the GitGuardian integration is changed to **Connected**.

After you set up your GitGuardian integration using the Integration Hub, you can see the secrets detection coverage.

The following video explains how to configure GitGuardian.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1740510716/snyk-learn/product-training-videos/Snyk_AppRisk_-12a_-_v1_-_Secrets_detection_coverage_with_GitGuardian.mp4" %}
Configure GitGuardian
{% endembed %}

The following video explains how GitGuardian uses data in Snyk.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1740510727/snyk-learn/product-training-videos/Snyk_AppRisk_-12b_-_v1_-_Secrets_detection_coverage_with_GitGuardian_Part_2.mp4" %}
How GitGuardian uses data in Snyk
{% endembed %}

## Dynatrace setup guide

{% hint style="info" %}
**Release status**

Dynatrace for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following risk factors are reported from the Dynatrace runtime integration: [Deployed](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) and [Loaded package](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

{% hint style="info" %}
The supported languages for the Loaded package risk factor reported from the Dynatrace integration are detailed on the Dynatrace [Supported technologies](https://docs.dynatrace.com/docs/platform-modules/application-security/getting-started/get-started-with-application-security#tech) page.
{% endhint %}

### Prerequisites <a href="#dynatrace-prerequisites" id="dynatrace-prerequisites"></a>

* Use Dynatrace SaaS on the DPS licensing model.
* The Dynatrace [Kubernetes app](https://docs.dynatrace.com/docs/platform-modules/infrastructure-monitoring/container-platform-monitoring/kubernetes-app/overview) is configured to monitor at least one cluster.
* The user is associated with a group that has permissions (through policies) to query the entity model. In the Dynayrace policy, set the following permission: `storage:entities:read`.

Comply with the following steps before integrating Dynatrace with Snyk AppRisk:

1. Retrieve the `account-uuid` from your Dynatrace account. Navigate to the [Dynatrace accounts](https://myaccount.dynatrace.com/accounts) page and select the account whose environment you want to integrate into Snyk. Identify the `account-uuid` in the URL and save it for later use.
2. Ensure you have OneAgent deployed in your Kubernetes environment. Navigate to `Settings` then `Environments` and select the environment you want to integrate into Snyk. Save the environment ID for later use (available in the URL of the new window as well). Click `Deploy OneAgent` then `Kubernetes` and follow the instructions. Ensure OneAgent is running in full-stack mode.
3. Ensure your deployment is activated. On your environment's page, click `Kubernetes` , then `Recommendations` and activate the cluster where you deployed OneAgent.
4. Under **account management**:
   * Create a new Permissions Group. Add the following permissions to the group:
     * Enable `View and manage account and billing information` under Account Management permissions.
     * Add both `All Grail data read access` and `Read Events` under general permissions.
   * Create a new, dedicated user to be used by the integration and assign to it the newly created Permissions Group mentioned above.
   * Create a new OAuth client:
     * Set the Service User Email as the email address of the newly generated user mentioned above.
     * Set the Permissions for the client - Under `Grail data ingest and query` select `View entities (storage:entities:read)`.
     * Ensure to hold a copy of the `Client ID` and `Client Secret` to be used in a later phase when configuring the integration in Snyk.

### Required parameters <a href="#dynatrace-required-parameters" id="dynatrace-required-parameters"></a>

1. **Account UUID** - the `account-uuid` of your Dynatrace account.
2. **Environment ID** - the ID of the environment monitored in Dynatrace.
3. **OAuth client ID** - the ID of the OAuth client created in the prerequisites.
4. **OAuth client secret** - the secret of the OAuth client created in the prerequisites.

### Integration Hub setup <a href="#dynatrace-integration-hub-setup" id="dynatrace-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **Runtime** tag and search for **Dynatrace**.
* Click the **Add** button.
* Edit the **Profile name** of your integration.
* Enter the **Account UUID**.
* Enter the **Environment ID**.
* Enter the **OAuth client ID**.
* Enter the **OAuth client secret**.
* Click the **Done** button.
* When the connection is established, the **Dynatrace** integration status changes to **Connected**.

{% hint style="info" %}
After the Dynatrace runtime data becomes available from the runtime integration, it will appear in Snyk AppRisk within a few hours.
{% endhint %}

<figure><img src="../.gitbook/assets/image (209).png" alt=""><figcaption><p>Dynatrace - Setup screen</p></figcaption></figure>

## Sysdig setup guide

{% hint style="info" %}
**Release status**

Sysdig for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following [risk factors](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#risk-factors) are reported from the Sysdig runtime integration: [Deployed](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md), and [Loaded package](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

{% hint style="info" %}
The supported languages for the Loaded package risk factor are: Go, Java, JavaScript/TypeScript and Python.
{% endhint %}

### Prerequisites <a href="#sysdig-prerequisites" id="sysdig-prerequisites"></a>

* The account must have access to Sysdig Secure product.
* Contact your Sysdig representative to activate the in-use packages feature flag.

### Required parameters <a href="#sysdig-required-parameters" id="sysdig-required-parameters"></a>

* **Service Account API Token** - Navigate to the [Service Account setup instructions page](https://docs.sysdig.com/en/docs/administration/administration-settings/access-and-secrets/user-and-team-administration/manage-teams-and-roles/#service-accounts) for details on how to create a Sysdig Service Account in order to obtain an API Token.
  * Set **View Only** as the Role for this Service Account.
  * Set an **Expiration Date** for the Service Account. After the Service Account expires, the Sysdig integration will no longer be able to pull information until updated with a new Service Account.

{% hint style="info" %}
The created **Service Account** must be under **Sysdig Secure**, not **Sysdig Monitor**.
{% endhint %}

* **Region -** Navigate to the Sysdig [SAAS regions and IP ranges](https://docs.sysdig.com/en/docs/administration/saas-regions-and-ip-ranges) page for details about the Sysdig region URLs.

### Known limitations

* If the Sysdig Agent is not deployed on every node of a cluster, runtime data available from this integration may be incomplete.
* Various Sysdig scans run at different intervals, which may cause a delay between applying changes to a resource within a cluster and reporting this information through the integration.

### Integration Hub setup <a href="#sysdig-integration-hub-setup" id="sysdig-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **Runtime** tag and search for Sysdig.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Account API Token**.
* Set the **Sysdig region**.
* Click the **Done** button.
* When the connection is established, the status of the Sysdig integration is changed to **Connected**.

{% hint style="info" %}
After the Sysdig runtime data becomes available from the runtime integration, it will appear in Snyk AppRisk within a few hours.
{% endhint %}

<figure><img src="../.gitbook/assets/image (210).png" alt=""><figcaption><p>Sysdig - Setup screen</p></figcaption></figure>

## Crowdstrike setup guide

{% hint style="info" %}
**Release status**

Crowdstrike for Snyk AppRisk is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following [risk factors](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#risk-factors) are reported from the Crowdstrike runtime integration: [Deployed](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md).

### Prerequisites <a href="#sysdig-prerequisites" id="sysdig-prerequisites"></a>

* Deploy the [**Kubernetes Admission Controller Agent (KAC)**](https://falcon.us-2.crowdstrike.com/documentation/page/d0a3095c/deploy-falcon-kubernetes-admission-controller-with-a-helm-chart) to the cluster.
* Create a [new API Client](https://falcon.us-2.crowdstrike.com/api-clients-and-keys/clients) for the integration:
  * API Client should include the **Falcon Container Image** scope.
  * Make sure you keep hold of the **Client Secret** being generated, as it's only visible once - during the generation of a new API Client.

### Required Parameters

The following parameters are required for setting the integration in Snyk UI. All values should be available from the newly generated API Client mentioned above.

* **Client ID**
* **Client Secret**
* **Base URL**

### Integration Hub Setup <a href="#crowdstrike-integration-hub-setup" id="crowdstrike-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **CNAPP** tag and search for CrowdStrike.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Set the **Client ID, Client Secret** and **URL** from the API Client as mentioned above.
* Click the **Done** button.
* When the connection is established, the status of the CrowdStrike integration is changed to **Connected**.

{% hint style="info" %}
After the Crowdstrike runtime data becomes available from the runtime integration, it will appear in Snyk AppRisk within a few hours.
{% endhint %}

<figure><img src="../.gitbook/assets/Screenshot 2025-05-29 at 17.01.33.png" alt=""><figcaption><p>Crowdstrike - Setup screen</p></figcaption></figure>
