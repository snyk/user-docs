# Authenticate the CLI with your account

To associate your Snyk account with the CLI, you must first authenticate your machine \(except for using the `snyk protect` command\). No repository permissions are needed at this stage, only your email address.

You can authenticate:

* Through your browser
* Using your API token

## Authenticate through your browser

1. Run `snyk auth` from the CLI.

   A browser tab opens, redirecting you to authenticate the CLI for use with your account.

2. Click **Authenticate**.

   The authentication ends and you can go back to your terminal to continue working.

![](../../.gitbook/assets/auth.gif/)

## Authenticate using your API token

1. Visit [your Snyk account](https://app.snyk.io/account/) \(**Account Settings &gt; API Token** section\).
2. In the **KEY** field, click **click to show**, then select and copy your API token:  
3. In the CLI, run `snyk config set api=XXXXXXXX`

![](../../.gitbook/assets/image%20%2811%29.png/)

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

