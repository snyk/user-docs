# Choose a service account type to use with Snyk APIs

## Group level or Organization level

Each Snyk action, whether for the Snyk API or the Snyk Web UI, requires a set of [user permissions](../../../snyk-platform-administration/user-roles/).

The permissions granted to a service account depend on the user role type it is assigned. For example, an Organization-level service account can use Organization-level roles, and a Group-level service account can use Group-level roles.

## Authentication types

#### Access tokens

Service accounts use Snyk [access tokens](../../../snyk-api/authentication-for-api/personal-access-tokens-pats.md) to secure your workflow. You control when the service account is created and can se the expiry date. The maximum expiry for an access token service account is one year, and on expiration you must create a new service account to continue using this authentication method.

Use an [OAuth 2.0 service account](service-accounts-using-oauth-2.0.md) instead of an access token service account, to prevent downtime when a token expires if you are using automation, for example, CI/CD pipeline.

#### OAuth 2.0

Service accounts have associated OAuth 2.0 client credentials that can fetch an OAuth 2.0 access token. These offer the highest level of security due to the short time-to-live of an OAuth access token and the standardized, automated refresh process.

{% hint style="info" %}
For more details on implementing this type of service account, visit [Service accounts using OAuth 2.0](service-accounts-using-oauth-2.0.md).
{% endhint %}

#### API key (legacy)

The service account has a Snyk API key associated with it. API keys are quick to implement in a workflow and do not expire. Due to the risk of a token which does not expire, this option is not recommended.
