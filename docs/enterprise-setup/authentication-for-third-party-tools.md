# Authentication for third-party tools

When you work with Snyk from within any third-party tool, Snyk requires authentication in order to initiate its processes.

Snyk offers API tokens to enable integrations with third-party developer tools. You can authenticate through your personal account using your personal token or through a [service account](service-accounts/) using the token associated with that account. When you authenticate through a service account, you do not use any personal token.

{% hint style="info" %}
For authentication purposes, the third-party identity providers do not require access to your repositories, only your email address.
{% endhint %}

## Supported identity providers

You can use one of the following identity providers for authentication with Snyk::

* GitHub
* Bitbucket
* Google
* Entra ID (formerly Azure AD)
* Docker ID
* Single Sign-On (SSO): available with Enterprise plans.\
  See [Setting up Single Sign-On (SSO) for authentication](single-sign-on-sso-for-authentication-to-snyk/).

For additional instructions, see the integrations pages for [Git repositories (SCMs)](../developer-tools/scms/organization-level-integrations/).

{% hint style="info" %}
Logging in with a different provider from the one you registered with when you first created your Snyk account will create a separate new Snyk account.
{% endhint %}

## **How to authenticate for a third-party tool using your personal token**

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. In the token field, click to show and then select and copy your API token.
4. In the third-party interface, configure your integration by pasting your Snyk token when prompted.

<figure><img src="../.gitbook/assets/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3) (16).png" alt="API token screen; revoke; regenerate; click to show"><figcaption><p>API token screen; revoke; regenerate; click to show</p></figcaption></figure>
