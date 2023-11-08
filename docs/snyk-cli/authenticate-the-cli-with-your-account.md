# Authenticate the CLI with your account

To associate the Snyk CLI for use with your Snyk account, you must first authenticate your machine. No repository permissions are needed at this stage, only your email address. When you have authenticated, you can [get started](getting-started-with-the-snyk-cli.md) using the CLI.

{% hint style="info" %}
Free and Team users have access to personal tokens only. Personal tokens are recommended for use with IDEs and the local CLI.&#x20;

Snyk recommends that Enterprise customers use a service token to authenticate with a CI/CD. Avoid using a service token with an IDE.
{% endhint %}

You can authenticate by using the CLI `snyk auth` command to launch the authentication dialog in your browser.  See the [Auth command help](commands/auth.md). This method is the default and recommended.

You can also authenticate by using the `SNYK_TOKEN` environment variable. For details, see [Configure the Snyk CLI](configure-the-snyk-cli/). Use `SNYK_TOKEN` in a CI/CD environment.

You can specify either your personal API token or a service token, available for Enterprise customers only. See [Service accounts](../enterprise-setup/service-accounts/) for information on using service tokens.

Follow these steps to authenticate using your API token:

1. Go to [your Snyk account](https://app.snyk.io/account), **Account Settings > API Token** section.
2. In the **KEY** field, choose **click to show**; then select and copy your API token. A screenshot follows.
3. In the CLI, run `snyk auth [<API_TOKEN>]` or `snyk config set api=<token>`. The `<API_TOKEN>`is validated by the Snyk API.

![Snyk Account Settings, API Token](../.gitbook/assets/API-token-CLI-auth-details-22-01.png)

For more details, see [Obtaining your Snyk API token](../getting-started/obtaining-your-snyk-api-token.md) for instructions that apply to all applications and tools.
