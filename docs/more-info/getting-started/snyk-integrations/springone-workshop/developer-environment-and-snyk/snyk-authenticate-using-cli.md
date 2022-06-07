# Authenticate with Snyk

## Authenticate the CLI

For this lab, use the instructions to **authenticate with your token** detailed below. **Authenticate with your account** via a browser is provided as an alternative.

### Authenticate the CLI with your account

When you are working with the Snyk CLI tool, Snyk first requires authentication.

To authenticate:

1.  Run `snyk auth` from the CLI.

    A browser tab opens, redirecting you to authenticate the CLI for use with your account.
2.  Click Authenticate.

    The authentication ends and you can go back to your terminal to continue working.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/auth\_image\_1.gif)

### Authenticate with your token

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. From the token field, click **click to show** and then select and copy your API token.
4. In the CLI, run `snyk config set api=XXXXXXXX`

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/auth\_image\_2.png)

{% hint style="info" %}
We will use your personal API token during the maven plugin.
{% endhint %}
