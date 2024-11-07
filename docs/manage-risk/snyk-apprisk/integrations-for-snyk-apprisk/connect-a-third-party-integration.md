# Third-party integrations for Snyk AppRisk

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that are automatically synced and provides access to the Integration Hub.

{% hint style="info" %}
The Loaded package risk factor is not supported by Snyk for operating system packages (such as Debian packages), only for packages which are hosted under package managers such as npm, Maven, or PyPI.
{% endhint %}

You can customize your AppRisk integrations from the **Integrations Hub** where the following integrations are available:

* [Veracode SAST](connect-a-third-party-integration.md#veracode-setup-guide)
* [Checkmarx SAST ](connect-a-third-party-integration.md#checkmarx-setup-guide)
* [SonarQube](connect-a-third-party-integration.md#sonarqube-setup-guide)
* [Nightfall](connect-a-third-party-integration.md#nightfall-setup-guide)
* [GitGuardian](connect-a-third-party-integration.md#gitguardian-setup-guide)
* [Jira](connect-a-third-party-integration.md#jira-setup-guide)
* [Dynatrace](connect-a-third-party-integration.md#dynatrace-setup-guide)
* [Sysdig](connect-a-third-party-integration.md#sysdig-setup-guide)
* [Orca Security](connect-a-third-party-integration.md#orca-security-setup-guide)
* [CrowdStrike](connect-a-third-party-integration.md#crowdstrike-setup-guide)
* [SentinelOne](connect-a-third-party-integration.md#sentinelone-setup-guide)
* [Datadog](connect-a-third-party-integration.md#datadog-setup-guide)
* [GCP](connect-a-third-party-integration.md#gcp-setup-guide)

{% hint style="info" %}
Data synchronization may take up to two hours after receiving the **Connected** status from a new integration setup.
{% endhint %}

## Veracode setup guide

{% hint style="info" %}
**Release status**

Veracode for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Prerequisites <a href="#verocode-prerequisites" id="verocode-prerequisites"></a>

The Veracode application concept is matched into the Snyk AppRisk repository assets. You need to create and utilize the Veracode custom field by using the [Veracode API](https://app.swaggerhub.com/apis/Veracode/veracode-applications\_api\_specification/1.0#/Application%20information%20API/updateApplicationUsingPUT). Access the [Veracode custom metadata field](https://docs.veracode.com/r/t\_create\_custom\_metadata) for more details.

Ensure you are adding a custom field called repoURL:

```
{
"name": "repoURL", 
"value": <YOUR GITHUB URL>
}

```

### Required parameters <a href="#veracode-required-parameters" id="veracode-required-parameters"></a>

* API ID and API Key - API credentials associated with a user account. For more information, access the [Veracode API credentials](https://help.veracode.com/r/c\_api\_credentials3) link.

### Integration Hub setup <a href="#veracode-integration-hub-setup" id="veracode-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.&#x20;
2. Select the **SAST** tag and search for Veracode.&#x20;
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the **API ID** from your Veracode account.
6. Add the **API key** from your Veracode account.
7. Click the **Done** button.
8. When the connection is established, the status of the Veracode integration is changed to **Connected**.

## Checkmarx setup guide

{% hint style="info" %}
**Release status**

Checkmarx for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

Use the following instructions to set up your Checkmarx SAST integration. Checkmarx SAST integration is only working for the Checkmarx SAST, we are not yet supporting Checkmarx One.

{% hint style="warning" %}
Snyk AppRisk Pro does not support the Checkmarx One integration.
{% endhint %}

### Prerequisites <a href="#checkmarx-prerequisites" id="checkmarx-prerequisites"></a>

* Install and configure your [Snyk Broker](../../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md#checkmarx-sast-integration) connection for Snyk AppRisk.&#x20;
* Ensure you have properly used Git Setting for your Checkmarx Project. Access the Checkmarx [Set project's remote source settings as GIT](https://checkmarx.stoplight.io/docs/checkmarx-sast-api-reference-guide/8312d35369b9b-set-project-s-remote-source-settings-as-git) documentation page for more details.&#x20;

### Required parameters <a href="#checkmarx-required-parameters" id="checkmarx-required-parameters"></a>

1. API URL - The URL of Checkmarx API, for example, `checkmarx.customer.com`.
2. Username and Password - Credentials for a user account with Checkmarx SAST access.

### Integration Hub setup <a href="#checkmarx-integration-hub-setup" id="checkmarx-integration-hub-setup"></a>

After you have installed and configured Snyk Broker for AppRisk and you successfully established a connection for Checkmarx SAST, you also need to configure the integration from the Snyk AppRisk Integration Hub. &#x20;

1. Open the **Integration Hub** menu.&#x20;
2. Select the **SAST** tag and search for Checkmarx.&#x20;
3. Click the **Add** button.
4. Add the profile name for this integration.
5. Add the Broker token for the Snyk AppRisk Checkmarx integration.
6. Add the Checkmarx host. For example `checkmarx.customer.com`
7. Click the **Done** button.
8. When the connection is established, the status of the Checkmarx integration is changed to **Connected**.

## SonarQube setup guide

{% hint style="info" %}
**Release status**

SonarQube for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#sonarqube-required-parameters" id="sonarqube-required-parameters"></a>

* API Key. [Here](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/) you can find more details about the SonarQube API Key.

### Integration Hub setup <a href="#sonarqube-integration-hub-setup" id="sonarqube-integration-hub-setup"></a>

* Open the **Integration Hub** menu.&#x20;
* Select the **SAST** tag and search for SonarQube.&#x20;
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Host URL** for this integration.
* Add the **API token**. Navigate to your SonarQube account, select User, select My Account, select Security, and then User Token. Access the SonarQube [generating and using tokens](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/) documentation page for more details about the SonarQube API Key.
* Click the **Done** button.
* When the connection is established, the status of the SonarQube integration is changed to **Connected**.

## Nightfall setup guide

{% hint style="info" %}
**Release status**

Nightfall for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#nightfall-required-parameters" id="nightfall-required-parameters"></a>

* API Key. Access the Nightfall [Creating an API Key](https://help.nightfall.ai/nightfall-firewall-for-ai/key-concepts/setting-up-nightfall/creating-api-key) documentation page for more details about how to create a Nightfall API key.&#x20;

### Integration Hub setup <a href="#nightfall-integration-hub-setup" id="nightfall-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.&#x20;
2. Select the **Secrets** tag and search for Nightfall.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** for this integration.
5. Add the **Base API URL** for this integration.
6. Add the **API Key** for this integration.
7. Click the **Done** button.
8. When the connection is established, the status of the Nightfall integration is changed to **Connected**.

The following video provides an overview of the Nightfall configuration from the Integration Hub:

{% embed url="https://www.youtube.com/watch?v=FJ5fAyMYSUs" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

After you set up your Nightfall integration using the Integration Hub, you can see the secrets detection coverage.

{% embed url="https://www.youtube.com/watch?v=o6TqPMSq1rk" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

## GitGuardian setup guide

{% hint style="info" %}
**Release status**

GitGuardian for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#gitguardian-required-parameters" id="gitguardian-required-parameters"></a>

* API Key. Access the GitGuardian [authentication](https://docs.gitguardian.com/api-docs/authentication) documentation page for more details about how to create a GitGuardian API Key.&#x20;

When you create a GitGuardian API Key, remember that it works for both service accounts and personal access token.

Ensure that the following permissions are set as READ:

* Incident (`mandatory`)
* Teams (`recommended` for GitGuardian paid accounts)

### Integration Hub setup <a href="#gitguardian-integration-hub-setup" id="gitguardian-integration-hub-setup"></a>

1. Open the **Integration Hub** menu.&#x20;
2. Select the **Secrets** tag and search for GitGuardian.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** for this integration.
5. Add the **API Token** for this integration.
6. Click the **Done** button.
7. When the connection is established, the status of the GitGuardian integration is changed to **Connected**.

The following video provides an overview of the GitGuardian configuration from the Integration Hub:

{% embed url="https://www.youtube.com/watch?v=4u4QrJBZTkI" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

After you set up your GitGuardian integration using the Integration Hub, you can see the secrets detection coverage:

{% embed url="https://www.youtube.com/watch?v=zh4c5f_vv1k" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

## Jira setup guide

{% hint style="info" %}
**Release status**

Jira for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Required parameters <a href="#jira-required-parameters" id="jira-required-parameters"></a>

* API Key - Access the Jira [API tokens](https://id.atlassian.com/manage-profile/security/api-tokens) documentation page for more details about how to generate a Jira API Token.&#x20;

{% hint style="info" %}
Ensure you have the correct user permissions before creating the API Token.
{% endhint %}

### Integration Hub setup <a href="#jira-integration-hub-setup" id="jira-integration-hub-setup"></a>

* Open the **Integration Hub** menu.&#x20;
* Select the **ITSM** tag and search for Jira.&#x20;
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **API Token** for this integration.
* Add the **User Email** used for this integration.
* Add the **Host URL** for this integration.
* Click the **Done** button.
* When the connection is established, the status of the Nightfall integration is changed to **Connected**.

{% hint style="warning" %}
You can add only one Jira profile to the Jira integration.
{% endhint %}

### Types of Jira integrations

Multiple Jira integrations are available when using Snyk, each designed to support specific needs.

* Jira - Manually create a ticket for issues from Snyk.
* [Jira Script](../../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md) - Automatically create tickets for new vulnerabilities.
* [Security in Jira](../../../integrate-with-snyk/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md) - View vulnerability information in Jira and create a ticket from Jira. The ticket is not visible in Snyk.
* [Jira for Snyk AppRisk](connect-a-third-party-integration.md#jira-setup-guide) - As part of the policy action, you can automatically create Jira tickets from Snyk AppRisk Assets.

The following table presents the functionality of all types of Jira integrations available in Snyk, specifies the supported Jira platform, the expected outcome, the authentication type, and the level of availability in Snyk.

| Jira integration type                                                                                                   | Functionality                                                                                                                                                                                                                                                                          | Authentication                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Jira                                                                                                                    | <p>Create a manual ticket for issues from Snyk.<br><br><strong>Supported for</strong>: <br>- Jira On-Cloud<br>- Jira Data Centre<br><br><strong>Outcome</strong>:<br>- Create Issue tickets </p>                                                                                       | <p><strong>Authentication type</strong>:  Personal Access Token</p><p><br><strong>Availability level</strong>: Snyk Organization</p>                                     |
| [Jira Script](../../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md)                             | <p>Automatically create tickets for new vulnerabilities.<br><br><strong>Supported for</strong>: <br>- Jira On-Cloud<br>- Jira Data Centre<br><br><strong>Outcome</strong>:<br>- Create Issue tickets </p>                                                                              | <p><strong>Authentication type</strong>:  Personal Access Token</p><p><br><strong>Availability level</strong>: Snyk Organization</p>                                     |
| [Security in Jira](../../../integrate-with-snyk/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md) | <p>View vulnerability information in Jira and create a ticket from Jira. <br>* Jira ticket is not visible in Snyk.<br><br><strong>Supported for</strong>:<br>- Jira On-Cloud<br><br><strong>Outcome</strong>:<br>- Create Issue tickets </p>                                           | <p><strong>Authentication type</strong>:  JWT(JSON Web Token) as part of the Connect App framework.</p><p><br><strong>Availability level</strong>: Snyk Organization</p> |
| [Jira for Snyk AppRisk](connect-a-third-party-integration.md#jira-setup-guide)                                          | <p>Use the "Create Jira ticket" action from a Snyk policy to create Jira tickets from Snyk AppRisk Assets automatically.<br><br><strong>Supported for</strong>:<br>- Jira On-Cloud<br>- Jira Data Centre (coming soon)<br><br><strong>Outcome</strong>:<br>- Create Asset tickets </p> | <p><strong>Authentication type</strong>:  Personal Access Token</p><p><br><strong>Availability level</strong>: Snyk Group</p>                                            |

## Dynatrace setup guide

{% hint style="info" %}
**Release status**

Dynatrace for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following risk factors are reported from the Dynatrace runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md), and [Loaded package](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

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
4. Create an OAuth client with the right permissions. Navigate to the [Dynatrace accounts](https://myaccount.dynatrace.com/accounts) page, then to `Identity & access management`. Select `OAuth clients` and click `Create client`. Fill in the details and check the following permissions; then click `Create client`:

<pre><code><strong>storage:entities:read
</strong></code></pre>

5. Save the Client ID and Client secret for later and click `Finish`.

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

## Sysdig setup guide

{% hint style="info" %}
**Release status**

Sysdig for Snyk AppRisk Pro is in Early Access and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following [risk factors](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk#risk-factors) are reported from the Sysdig runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md), and [Loaded package](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

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

* Open the **Integration Hub** menu.&#x20;
* Select the **Runtime** tag and search for Sysdig.&#x20;
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Account API Token**.
* Set the **Sysdig region**.
* Click the **Done** button.
* When the connection is established, the status of the Sysdig integration is changed to **Connected**.

{% hint style="info" %}
After the Sysdig runtime data becomes available from the runtime integration, it will appear in Snyk AppRisk within a few hours.
{% endhint %}

## Orca Security setup guide

{% hint style="info" %}
**Release status**

Orca Security for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following[ risk factors](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#risk-factors) are reported from the Orca runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md).

### Prerequisites <a href="#orca-prerequisites" id="orca-prerequisites"></a>

* Connect your vendor account to Orca. Navigate to the [Connecting Accounts](https://docs.orcasecurity.io/docs/onboarding) documentation for more details.
* Connect your cluster to Orca. Navigate to the [Connecting Clusters using Kubernetes Connector](https://docs.orcasecurity.io/docs/connecting-clusters-using-kubernetes-connector) documentation for more details.

### Required parameters <a href="#orca-required-parameters" id="orca-required-parameters"></a>

* **API Token** - generate an API token for Orca.

Create the Orca API Token by following these steps:

1. Open your Orca account, click **Settings**, then **Users and Permissions**, then **API**.
2. Click **Add API Token**.
3. Fill in the fields **Name** and **Description** for the API Token.
4. Select an **Expiration** date for the API Token.
5. Select **Integration Configuration** as a Role.
6. Check **Scope access to specific resources** and **Accounts**.
7. Select the account you want to integrate.

### Integration Hub setup <a href="#orca-integration-hub-setup" id="orca-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **Runtime** tag and search for Orca.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **API Token**.
* Set the **URL**.
* Click the **Done** button.
* When the connection is established, the status of the Orca integration is changed to **Connected**.

## CrowdStrike setup guide

{% hint style="info" %}
**Release status**

CrowdStrike for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following risk factor is reported from the CrowdStrike runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) risk factor.

### Prerequisites <a href="#crowdstrike-prerequisites" id="crowdstrike-prerequisites"></a>

* Deploy the [**Falcon Kubernetes Protection Agent (KPA)**](https://falcon.us-2.crowdstrike.com/documentation/page/q444c05c/deploy-falcon-kubernetes-protection-agent-with-a-helm-chart) to the cluster.
* Include the following scopes in the API Client:
  * **Falcon Container CLI** - write
  * **Falcon Container Image** - read/write
  * **Falcon Images Download** - read
  * **Kubernetes Protection Agent** - write
  * **Kubernetes Protection** - read/write
  * **Sensor Download** - read

### Required parameters <a href="#crowdstrike-required-parameters" id="crowdstrike-required-parameters"></a>

After generating the API Client following the [**Falcon Kubernetes Protection Agent (KPA)**](https://falcon.us-2.crowdstrike.com/documentation/page/q444c05c/deploy-falcon-kubernetes-protection-agent-with-a-helm-chart) deployment, the **Client ID** and the **Client Secret** are generated.

* **Client ID** - To retrieve the **Client ID**, follow these steps:
  * Left side navigation bar -> Support and resources -> API clients and keys.
  * Under **OAuth2 API Clients** press on **Create API Client**.
  * Select the API Client used in the **Falcon Kubernetes Protection Agent (KPA)**.
  * You can find the **Client ID** on the right side, under the details of the API Client.
* **Client Secret** - This value is available only at the first generation and cannot be accessed afterward.
* **URL** - Select the URL based on your network environment.

### Integration Hub setup <a href="#crowdstrike-integration-hub-setup" id="crowdstrike-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **CNAPP** tag and search for CrowdStrike.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Client ID**.
* Add the **Client Secret**.
* Set the **URL**.
* Click the **Done** button.
* When the connection is established, the status of the CrowdStrike integration is changed to **Connected**.

{% hint style="info" %}
After the CrowdStrike runtime data becomes available from the runtime integration, it will appear in Snyk AppRisk within a few hours.
{% endhint %}

## SentinelOne setup guide

{% hint style="info" %}
**Release status**

SentinelOne for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following risk factor is reported from the SentinelOne runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) risk factor.

### Prerequisites <a href="#sentinelone-prerequisites" id="sentinelone-prerequisites"></a>

* You must be a customer of Cloud Workload Protection Platform (CWPP).
* Your SentinelOne license needs to include the following SKUs: `CWS for Containers`, `CWS for Serverless Containers` or `CWS for Servers`

### Required parameters <a href="#sentinelone-required-parameters" id="sentinelone-required-parameters"></a>

Navigate to the [SentinelOne Community](https://community.sentinelone.com) page for more details about the required parameters.&#x20;

**Account IDs**

1. Navigate to **Settings,** then **Accounts.**
2. Select the account you want to integrate.
3. Navigate to **ACCOUNT INFO.**
4. Copy the **ACCOUNT ID.**

**Site IDs**

1. Navigate to **Settings,** then **Sites.**
2. Select the site you want to integrate.
3. Navigate to **SITE INFO.**
4. Copy the **SITE ID.**

**API token**

1. Navigate to **Settings, Users,** then **Service Users.**
2. Click on **Actions,** then **Create New Service User.**
3. Give the new service user a name and a description, and select an Expiration Date for it. When the expiration date is reached, the integration will need to be reset.
4. Select the desired Sites.
5. Click on **Create User.**
6. Enter your 2FA code.
7. Copy the API Token.

Alternatively, you can use an existing **API Token**.

### Integration Hub setup <a href="#sentinelone-integration-hub-setup" id="sentinelone-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **CWPP** tag and search for **SentinelOne**.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **SentinelOne account IDs**.
* Add the **SentinelOne site IDs**.
* Add the **API token**.
* Click the **Done** button.
* When the connection is established, the status of the SentinelOne integration is changed to **Connected**.

## Datadog setup guide

{% hint style="info" %}
**Release status**

Datadog for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The following risk factor is reported from the Datadog runtime integration: [Deployed](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) risk factor.

### Prerequisites <a href="#datadog-prerequisites" id="datadog-prerequisites"></a>

You must set the **Datadog Container Monitoring** for the desired cluster.

### Required parameters <a href="#datadog-required-parameters" id="datadog-required-parameters"></a>

* **API key** - Navigate to the [Datadog API and Application key](https://docs.datadoghq.com/account\_management/api-app-keys/#add-an-api-key-or-client-token) page for all the steps you need to follow to create a new API key in Datadog. You can also use an existing API key by selecting it and clicking the Copy button.
* **Application key** - Navigate to the [Datadog Add application keys](https://docs.datadoghq.com/account\_management/api-app-keys/#add-application-keys) page for all the steps you need to follow to create an Application key in Datadog.
* **Site** - If you are using a non-default Datadog site, you need to select it from the list.

### Integration Hub setup <a href="#datadog-integration-hub-setup" id="datadog-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **Observability** tag and search for Datadog.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **API key**.
* Add the **Application key**.
* Optional, select the Datadog **Site**.
* Click the **Done** button.
* When the connection is established, the status of the Datadog integration is changed to **Connected**.

## GCP setup guide

{% hint style="info" %}
**Release status**

The GCP integration for Snyk AppRisk Pro is in Closed Beta and available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

### Prerequisites <a href="#datadog-prerequisites" id="datadog-prerequisites"></a>

Set up a service account for your GCP project to use with the integration:

1. Navigate to your GCP project from the Google Cloud Console.
2. Navigate to **IAM & Admin** -> **Service Accounts**.
3. Click on **+ CREATE SERVICE ACCOUNT**.
4. Choose a name (for example `snyk-service-account`), add an optional description and click **CREATE AND CONTINUE**.
5. Under **Grant this service account access to project**, choose the following role to allow read-only access to all Kubernetes resources: `Cloud Asset Viewer`.
6. Click **Done**.
7. In the list of service accounts, click the service account you have just created.
8. Click the **KEYS** tab, click **ADD KEY** -> **Create new key** -> **JSON** and then **CREATE**. The JSON file will be downloaded to your computer, remember its location.

### Required parameters <a href="#datadog-required-parameters" id="datadog-required-parameters"></a>

* **Project ID** - The ID of the GCP project you want to integrate with.
* **JSON Key File** - The content of the JSON file you created and downloaded as part of the prerequisites.

### Integration Hub setup <a href="#datadog-integration-hub-setup" id="datadog-integration-hub-setup"></a>

* Open the **Integration Hub** menu.
* Select the **Cloud Providers** tag and search for **GCP**.
* Click the **Add** button.
* Add the **Profile name** for this integration.
* Add the **Project ID**.
* Add the **JSON Key File** - The content of the JSON file created in the prerequisites section.
* Click the **Done** button.
* When the connection is established, the status of the GCP integration is changed to **Connected**.
