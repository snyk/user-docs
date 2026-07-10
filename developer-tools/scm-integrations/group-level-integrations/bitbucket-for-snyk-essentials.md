---
description: How to configure the Bitbucket integration for Snyk Essentials at the Group level
---

# Bitbucket for Snyk Essentials

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

{% hint style="info" %}
The Bitbucket Cloud App is not supported at the Group level. The available options at the Group level are BitBucket Cloud and BitBucket Server.

To improve security, the use of app passwords in Bitbucket Cloud is transitioning to API tokens. Existing integrations that use app passwords will continue to function temporarily until 9 June 2026. To ensure continued support and functionality, update your Bitbucket Cloud integration in Snyk to use an API token.
{% endhint %}

{% hint style="info" %}
Bitbucket Server and Bitbucket Cloud do not support automatic language detection. You can manually add language tags to a Bitbucket Cloud repository.\
After manually setting up the languages in your Bitbucket project, Snyk can automatically detect and ingest all those languages in your Snyk Essentials application.
{% endhint %}

## Pulled entities <a href="#bitbucket-pulled-entities" id="bitbucket-pulled-entities"></a>

* Users
* Repositories

## Prerequisites <a href="#azure-devops-integrate-using-snyk-apprisk" id="azure-devops-integrate-using-snyk-apprisk"></a>

To configure a Group-level integration, you must be a Group Admin or have a custom role that includes the `Edit Snyk Essentials` permissions under the [Group-level permissions](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/user-management/pre-defined-roles#group-level-permissions).

## Integrate using Snyk Essentials <a href="#bitbucket-integrate-using-snyk-apprisk" id="bitbucket-integrate-using-snyk-apprisk"></a>

1. **Profile name** (mandatory): Enter a name for this integration profile.
2. **Service type** (mandatory): Select **Cloud** or **Server**.
3. **Authentication**: Enter the credentials for your selected service type. \
   &#x20;     For **Bitbucket Cloud:**
   1. **User email**: Your Atlassian account email.
   2. **API Token**: Create a Bitbucket [API Token](https://support.atlassian.com/bitbucket-cloud/docs/create-an-api-token/) in your Bitbucket Organization with the following permissions: `read:user:bitbucket`, `read:workspace:bitbucket`, and `read:repository:bitbucket`.\
      For **Bitbucket Server**:
   3. **Username**: Your Bitbucket Server username.
   4. **App Password**: In your Bitbucket account, navigate to **Settings** > **Personal Bitbucket settings** > **App passwords** and create a password with these permissions: **Account: Read** and **Projects: Read**.
4. **Broker Token** (mandatory for Bitbucket Server not reachable through the Internet): [Obtain your Broker token for Snyk Broker](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker) and paste it here.
5. **Backstage Catalog** (optional): To sync your catalog, visit the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/)&#x20;

## API version <a href="#bitbucket-api-version" id="bitbucket-api-version"></a>

To access information about the API, use the [Bitbucket REST API V2](https://developer.atlassian.com/bitbucket/api/2/reference/resource/).
