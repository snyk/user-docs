# Authentication for API

{% hint style="info" %}
**Feature Availability**

To use the Snyk API, you must be an Enterprise plan customer and have a token from Snyk.
{% endhint %}

{% hint style="warning" %}
Use the URL for your region when calling an API. See [API URLs](../rest-api/about-the-rest-api.md#api-urls).
{% endhint %}

Enterprise users have [access to a personal access token under their profile](./#how-to-obtain-your-personal-token) and to service account tokens. The PAT tokens are associated with your Snyk Account. Service accounts are associated with an Organization or a Group. For more information, see [Service accounts](../../implementation-and-setup/enterprise-setup/service-accounts/).

* **Enterprise users should use a service account** to authenticate for any kind of automation. This includes, but is not limited to, CI/CD scanning with the CLI or build system plugins and any automation, including automation with the API. Using a service account ensures continuity when users change roles or close their personal Snyk accounts.
* **Enterprise users should use the personal token** under their user profile for:
  * Running the CLI locally on their machine; for details, see [Authenticate to use the CLI](../../developer-tools/snyk-cli/authenticate-to-use-the-cli.md).
  * Authenticating with the IDE manually
  * Running API calls one time, for example, to test something

Note that for free and team plan users, the personal token does not have access to the API and may be used for authenticating to IDE, CLI, and CI/CD integrations only. For details, see [Obtain and use your PAT token](../../discover-snyk/getting-started/#obtain-and-use-your-snyk-api-token).

For additional information, see [Snyk PAT token permissions users can control](snyk-api-token-permissions-users-can-control.md).

## How to obtain your personal access token

You can create a personal access token through your [account settings](https://app.snyk.io/account/personal-access-tokens) after you register with Snyk and log in. For details, see [Personal Access Tokens](personal-access-tokens-pats.md).

### How to obtain an API token (legacy)

You can find your personal API token in your personal [account settings](https://app.snyk.io/account) after you register with Snyk and log in. In the **key** field, **Click to show**. Then highlight and copy the API key.

If you want a new API token, select **Revoke & Regenerate.** This will make the previous API token invalid. For details, see [Revoke and regenerate a Snyk API token](revoke-and-regenerate-a-snyk-api-token.md).

## Authenticating with a personal API token

When using the API directly, provide the API token in an `Authorization: token` header, as in the following example request, replacing `API_TOKEN` with your token

```bash
curl --request GET \
--url "https://api.snyk.io/rest/self?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

## Authenticating for Snyk Apps

If you are using the [Snyk Apps APIs](../using-specific-snyk-apis/snyk-apps-apis/), provide the `access_token` in an `Authorization: bearer` header as follows:

```bash
curl --request GET \
--url "https://api.snyk.io/rest/self?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: bearer API_TOKEN"
```

Otherwise, a `401 Unauthorized` error will be returned.
