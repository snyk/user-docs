# Bitbucket for Snyk Essentials

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

{% hint style="info" %}
The Bitbucket Cloud App is not supported at the Group level. The available options at the Group level are BitBucket Cloud and BitBucket Server.
{% endhint %}

### BitBucket setup guide

{% hint style="info" %}
Bitbucket Server and Bitbucket Cloud do not support automatic language detection. You can manually add language tags to a Bitbucket Cloud repository.\
After manually setting up the languages in your Bitbucket project, Snyk can automatically detect and ingest all those languages in your Snyk Essentials application.
{% endhint %}

#### Pulled entities <a href="#bitbucket-pulled-entities" id="bitbucket-pulled-entities"></a>

* Users
* Repositories

#### Integrate using Snyk Essentials <a href="#bitbucket-integrate-using-snyk-apprisk" id="bitbucket-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Access Token (`mandatory`) to create your BitBucket PAT from your BitBucket organization.
   * API URL (`mandatory`) - Input the API URL.
   * Username (`mandatory`): Input the BitBucket username of your organization.
   * App password (`mandatory`): Create an [API token](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#app-passwords) from your BitBucket account, with the following permissions:
     * **Account** - Read
     * **Projects** - Read

{% hint style="info" %}
Create a BitBucket app password by following these steps:

1. Open your BitBucket account&#x20;
2. Click the Settings option
3. Click the Personal BitBucket settings option&#x20;
4. Navigate to the App passwords sub-section from the ACCESS MANAGEMENT section.
{% endhint %}

3. Broker Token (`mandatory`) to create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) page.&#x20;
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
4. Service type (`mandatory`): Select the service type, Cloud, or On-premises.
5. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/) page.

#### API version <a href="#bitbucket-api-version" id="bitbucket-api-version"></a>

You can use the [BitBucket REST API V2](https://developer.atlassian.com/bitbucket/api/2/reference/resource/) repository to access information about the API.

