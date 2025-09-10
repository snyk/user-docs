# Bitbucket for Snyk Essentials

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

{% hint style="info" %}
The Bitbucket Cloud App is not supported at the Group level. The available options at the Group level are BitBucket Cloud and BitBucket Server.

To improve security, the use of app passwords in Bitbucket Cloud is transitioning to API tokens. Existing integrations that use app passwords will continue to function temporarily until **9 June 2026**. To ensure continued support and functionality, update your Bitbucket Cloud integration in Snyk to use an API token.
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
2. Service type (`mandatory`): Select the service type, Cloud, or Server.
3. User email (`mandatory`): Atlassian account email.
4. API Token (`mandatory`) to create your BitBucket [API Token](https://support.atlassian.com/bitbucket-cloud/docs/create-an-api-token/) from your BitBucket organization,  with the following permissions:
   * `read:user:bitbucket`
   * `read:workspace:bitbucket`
   * `read:repository:bitbucket`
5.  Broker Token (`mandatory`) to create and add your Broker token if you use Snyk Broker.

    This step is only for BitBucket Server that are not reachable through the internet.

    Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) page. Copy and paste the Broker token on the integration setup menu from the Integration Hub.
6. Add Backstage Catalog (`optional`). If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/) page.

#### API version <a href="#bitbucket-api-version" id="bitbucket-api-version"></a>

You can use the [BitBucket REST API V2](https://developer.atlassian.com/bitbucket/api/2/reference/resource/) repository to access information about the API.

