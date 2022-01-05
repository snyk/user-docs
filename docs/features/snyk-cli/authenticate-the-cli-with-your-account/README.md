# Authenticate the CLI with your account

To associate the Snyk CLI for use with your Snyk account, you must first authenticate your machine. No repository permissions are needed at this stage, only your email address.

You can authenticate:

* Through your browser
* Using your API token, required in a CI/CD environment

## Authenticate through your browser

1.  Run `snyk auth` from the CLI.

    A browser tab opens, redirecting you to authenticate your machine to associate the Snyk CLI for use with your account.
2.  Click **Authenticate**.

    The authentication ends and you can go back to your terminal to continue working.

## Authenticate using your API token

1. Visit [your Snyk account](https://app.snyk.io/account) (**Account Settings > API Token** section).
2. In the **KEY** field, click **click to show**, then select and copy your API token. A screenshot follows.
3. In the CLI, run `snyk config set api=<token>`

![Snyk Account Settings, API Token](../../../.gitbook/assets/API-token-CLI-auth-details-22-01.png)
