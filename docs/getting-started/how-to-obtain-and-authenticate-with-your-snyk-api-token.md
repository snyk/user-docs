# How to obtain and authenticate with your Snyk API token

Your Snyk API token is a personal token available under your user profile. A token is required to authenticate with Snyk as an individual user and is also used in `SNYK_TOKEN` parameters. The Snyk API token is associated with your Snyk Account and not with a specific Organization.

Free, Team, and Trial plan users have access only to this personal token under the user profile. The **personal token can be used** to authenticate with:

* A CI/CD integration
* The Snyk CLI running on a local or a build machine
* An IDE, when setting a token manually

Enterprise users have access to a personal token under their profile and to service account tokens. For details, see [Service accounts](../getting-started-with-the-snyk-enterprise-plan/service-accounts/).

* **Enterprise users should use a service account** to authenticate for any kind of automation. This includes, but is not limited to, CI/CD scanning with the CLI or build system plugins and automations, including the API.
* **Enterprise users should use the personal token** under their user profile for:
  * Running the CLI locally on their machine
  * Authenticating with the IDE manually
  * Running API calls one time, for example, to test something

For more information on the personal Snyk API token, see the following pages: [Authenticate the CLI with your account](../snyk-cli/authenticate-the-cli-with-your-account.md) and [Authentication for API](../snyk-api-info/authentication-for-api.md).

{% hint style="info" %}
If you are using Snyk on the EU and AU tenants,  you must set your endpoints accordingly before authenticating. For more information, see [Regional hosting and data residency](../working-with-snyk/regional-hosting-and-data-residency.md).
{% endhint %}

Follow these steps to obtain your personal Snyk API token:

1. Log in to Snyk and navigate to your **Account settings**
2. In the **Account Settings,** select **General** > **Auth Token**
3. Click inside the **KEY** box to display your API token.
4. Copy the token and save it in a secure location for future use.

<figure><img src="../.gitbook/assets/Snyk Broker - API Token - Account settings - API Token box.png" alt="Settings page, display API token"><figcaption><p>Settings page, display API token</p></figcaption></figure>
