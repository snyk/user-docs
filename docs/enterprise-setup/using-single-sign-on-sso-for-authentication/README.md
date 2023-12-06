# Using Single Sign-On (SSO) for authentication

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise Plans. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

## Overview of SSO

You can take advantage of your company's existing identity management system and have employees sign in to Snyk using their corporate identity. This makes provisioning Snyk to users easier. It also allows for deeper integration for Group and Organization membership, role-based access, and more.

<figure><img src="../../.gitbook/assets/image (1) (4).png" alt="Log in to Snyk with SSO"><figcaption><p>Log in to Snyk with SSO</p></figcaption></figure>

Snyk can integrate with any SAML-based and OpenID Connect (OIDC)-based SSO, as well as ADFS. You can also use your Enterprise Identity Provider for SSO, including [Azure AD](https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-whatis) and [Google G Suite](https://community.snowflake.com/s/article/configuring-g-suite-as-an-identity-provider). Read more about SAML in [the Auth0 documentation](https://auth0.com/docs/protocols/saml).

## User authentication and provisioning for SSO

With SSO configured, users are provisioned with a new Snyk account when they first sign on through SSO, even if they previously created their own account.

The sign-on process includes these steps:

1. When users select SSO from Snyk.io to log in, they are redirected to and authenticated by the identity provider you requested.
2. The identity provider communicates this authentication to Snyk servers, sending relevant data to Snyk to create each user.
3. Snyk checks the directory for that user.
4. If the user is already configured, Snyk enables the appropriate access. For a new user, Snyk adds the user to the directory and then redirects the user to Snyk.io with the appropriate access.

## Additional resource for SSO

Training: [SSO, authentication and user provisioning](https://learn.dev.snyk.io/lesson/sso-authentication-provisioning/)
