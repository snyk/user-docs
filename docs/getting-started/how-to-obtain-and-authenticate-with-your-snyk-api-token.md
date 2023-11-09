# How to obtain and authenticate with your Snyk API token

Your Snyk API token is a personal token, required to authenticate with Snyk as an individual user and used in `SNYK_TOKEN` parameters. The Snyk API token is associated with your Snyk Account and not with a specific Organization.

Free and Team plan users have access to this personal token only. Snyk recommends that all users authenticate with the CLI and IDEs using the personal token.

Enterprise users have access to service accounts and should authenticate with a service account token when the documentation recommends doing so. Snyk recommends that Enterprise users authenticate with CI/CDs using a service account, All users should authenticate with the CLI and IDEs using the personal token.

For more information on the Snyk API token, see the following pages: [Authenticate the CLI with your account](../snyk-cli/authenticate-the-cli-with-your-account.md) and [Authentication for API](../snyk-api-info/authentication-for-api.md).

Follow these steps to obtain your Snyk API token:

1. Log in to Snyk and navigate to your **Account settings**
2. In the **Account Settings,** select **General** > **Auth Token**
3. Click inside the **KEY** box to display your API token.
4. Copy the token and save it in a secure location for future use.

<figure><img src="../.gitbook/assets/Snyk Broker - API Token - Account settings - API Token box.png" alt="Settings page, display API token"><figcaption><p>Settings page, display API token</p></figcaption></figure>
