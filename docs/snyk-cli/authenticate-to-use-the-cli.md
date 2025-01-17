# Authenticate to use the CLI

To scan your projects, you must authenticate with Snyk.&#x20;

{% hint style="info" %}
If you are not in the system default environment, `SNYK-US-01`, use the [`snyk config environment`](commands/config-environment.md) command to set your environment before you run [`snyk auth`](commands/auth.md).
{% endhint %}

## How to authenticate to use the CLI locally

### Steps to authenticate using OAuth 2.0 protocol

When you are using the CLI locally, **Snyk recommends that you use the OAuth 2.0 protocol.**  Follow these steps:

1. Run the `snyk auth` CLI command.
2. Log in if you are prompted to do so.
3. The next page asks for your authorization for the CLI to act on your behalf. Click **Grant app access**.
4. After you authenticate successfully, view the confirmation message; then close the browser window and return to the CLI in the terminal.&#x20;

After authentication is granted, a pair of access and refresh tokens is stored locally for future use. \
\
Multi-tenant users who do not belong to the default region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.

If you have problems, see [OAuth 2.0 authentication does not work](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md).

{% hint style="info" %}
OAuth 2.0 tokens are not static. You cannot copy these tokens from your Snyk account page.
{% endhint %}

### Steps to retrieve the Snyk API token and use it to authenticate

{% hint style="warning" %}
This method is inferior to the OAuth 2.0 method.
{% endhint %}

To authenticate using your Snyk API token, follow these steps:

1. Run the`snyk auth --auth-type=token` CLI command
2. Log in if you are prompted to do so.
3. The next page prompts you to authenticate your machine to associate the Snyk CLI or the IDE plugin with your account. Click **Authenticate**.
4. After you authenticate successfully, a confirmation message appears. Close the browser window and return to the CLI in the terminal.&#x20;

After you complete the dialog, the API token is stored locally for future use.&#x20;

All subsequent `test` commands will be authenticated automatically.&#x20;

### Steps to authenticate using a known Snyk API token

You can copy your personal API token from your **General Account settings** (under your username) in the Snyk Web UI, and then configure your CLI to use it locally.

All CLI `test` commands can automatically recognize the environment variable `SNYK_TOKEN` and use it for authentication. For details, see [Environment variables for Snyk CLI](configure-the-snyk-cli/environment-variables-for-snyk-cli.md).

To use API token-based authentication, set the `SNYK_TOKEN` environment variable and run the `test` command, for example:\
`SNYK_TOKEN=<SNYK_API_TOKEN> snyk test`

Alternatively, you can export the environment variable to make it available for future `test` commands:\
`export SNYK_TOKEN=<SNYK_API_TOKEN>`\
`snyk test`

This form of authentication is particularly useful for CI/CD pipelines. See [How to authenticate to use the CLI in CI/CD pipelines](authenticate-to-use-the-cli.md#how-to-authenticate-to-use-the-cli-in-ci-cd-pipelines).

You can also store the Snyk API token locally for later use by running the following CLI command:\
`snyk auth <SNYK_API_TOKEN>`

All subsequent test calls will be authenticated automatically.  For more information, see the [Auth command help](commands/auth.md).

## How to authenticate to use the CLI in CI/CD pipelines

**Free and Team plan users** are more likely to **use this method in a CI/CD pipeline** than to use OAuth 2.0. **Enterprise plan customers** are advised to use a [**service account**](../enterprise-setup/service-accounts/) in a CI/CD pipeline. For details about when to use an API token and when to use a service account token, see [Authentication for API](../snyk-api/rest-api/authentication-for-api/).

All CLI `test` commands can automatically recognize the environment variable `SNYK_TOKEN` and use it for authentication. For details, see [Environment variables for Snyk CLI](configure-the-snyk-cli/environment-variables-for-snyk-cli.md).

To use API token-based authentication, set the `SNYK_TOKEN` environment variable and run the `test` command, for example:\
`SNYK_TOKEN=<SNYK_API_TOKEN> snyk test`

Alternatively, you can export the environment variable to make it available for future `test` commands:\
`export SNYK_TOKEN=<SNYK_API_TOKEN>`\
`snyk test`

You can also store the Snyk API token locally for later use by running the following CLI command:\
`snyk auth <SNYK_API_TOKEN>`

All subsequent test calls will be authenticated automatically. For more information, see the [Auth command help](commands/auth.md).
