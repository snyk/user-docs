# Set up Snyk Single Sign-On (SSO)

Set up Single Sign-On (SSO) to allow your developers and teams easy access to Snyk through your existing SSO provider.

The information you need to establish trust between Snyk and the identity provider depends on which type of SSO you are using.

Ensure you have at least one Group and Organization to indicate where new users will be assigned. For details, see [Manage Groups and Organizations](../../../snyk-platform-administration/groups-and-organizations/).

{% hint style="info" %}
After you gather the needed information identified in the following sections, create a support ticket to request SSO setup.

Alternatively, Group Admins can configure Snyk Single Sign-On. See [Self-Serve Single Sign-On (SSO)](configure-self-serve-single-sign-on-sso/) for the steps.
{% endhint %}

## Overview of SSO

The process of establishing trust between your identity provider (IdP) and Snyk requires a few steps coordinated between your SSO administrator and Snyk Support.

* In your identity provider platform, enter details about the Snyk environment and user attributes.
* Provide Snyk with details from your IdP.
* Set up a user for testing and send Snyk the username and password for that user.
* Use the link provided by Snyk to generate a payload.
* After Snyk finalizes the connection, confirm the login process is working correctly.

Users are provisioned to Snyk when they log in. If an invitation is required, users may only see a list of your Organizations until the admin adds them to the appropriate Organizations.

After SSO is configured both from Snyk and your company's network, a trust relationship is established with Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.

Each type of SSO connection requires different details for establishing trust between your identity provider and Snyk. The following sections explain those details. The details are also included in the worksheets in the Resources section at the end of this article.

## Set up SAML for SSO

To establish trust with Snyk, add an Entity ID, an Assertion Consumer Service (ACS) URL, and a Signing certificate in your identity provider.

* The **Entity ID** is the URL that uniquely identifies Snyk as a SAML entity or service provider. The **d**efault Entity ID must be checked manually as no default is set for this.
* The **Assertion Consumer Service (ACS)** is the endpoint on the Snyk network that listens for requests from your identity provider to enable communication between users on your network and Snyk. This URL is sometimes called a Reply URL.
* The **Signing certificate** is the Snyk certificate, stored on your server that is needed to maintain the trust relationship. It contains the necessary encryption keys for authentication.

Use these details to set up the connection with your Identity provider (IdP):

| Entity ID (SNYK-US-01 and SNYK-US-02)           | urn:auth0:snyk:saml-{group-slug}                                                                                                                             |
| ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Entity ID (SNYK-EU-01)                          | urn:auth0:snyk-mt-eu-prod-1:saml-{group-slug}                                                                                                                |
| Entity ID (SNYK-AU-01)                          | urn:auth0:snyk-mt-au-prod-1:saml-{group-slug}                                                                                                                |
| ACS URL (SNYK-US-01 and SNYK-US-02)             | [https://snyk.auth0.com/login/callback?connection=saml-](https://snyk.auth0.com/login/callback?connection=saml-){group-slug}                                 |
| ACS URL (SNYK-EU-01)                            | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback?connection=saml-](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback?connection=saml-){group-slug} |
| ACS URL (SNYK-AU-01)                            | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback?connection=saml](https://snyk-mt-au-prod-1.au.auth0.com/login/callback?connection=saml)-{group-slug}  |
| Signing certificate (SNYK-US-01 and SNYK-US-02) | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem)                                                                                                     |
| Signing certificate (SNYK-EU-01)                | [https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection)                                     |
| Signing certificate (SNYK-AU-01)                | [https://snyk-mt-au-prod-1.au.auth0.com/pem?cert=connection)](https://snyk-mt-au-prod-1.au.auth0.com/pem?cert=connection\))                                  |

{% hint style="info" %}
Replace {group-slug} with the name of your Snyk Group. If your Group name includes spaces, replace them with dashes. For example, if your Group name is `Your Company Group`, then the {group-slug} value is your-company-group.
{% endhint %}

To map information from your Identity provider to Snyk, name your user attributes as follows, using the same capitalization and spelling:

| email    | The user email address                          |
| -------- | ----------------------------------------------- |
| name     | The name of the person to be authenticated      |
| username | The person’s username for the identity provider |

If your user attributes do not match, note that the Snyk configuration for your SSO will take more time.

## SAML information to provide to Snyk

Obtain the following information from your identity provider. Provide this information to Snyk to establish trust on the service-provider side.

| Information                   | Description                                                                                                                                                                                                                                                                                                                                                                            |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Sign-in URL                   | The URL for your identity provider sign-in page                                                                                                                                                                                                                                                                                                                                        |
| X509 Signing Certificate      | The identity provider public key, encoded in Base64 format                                                                                                                                                                                                                                                                                                                             |
| Sign-out URL                  | <p>Optional, but recommended -</p><p>The URL for redirect whenever a user logs out of Snyk</p>                                                                                                                                                                                                                                                                                         |
| User ID attribute             | <p>Optional default is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier<br><br>This value uniquely identifies Snyk users and if changed will result in a duplicate user being created. If you see a duplicate user after changing identity provider <a href="https://support.snyk.io">submit a request </a>to Snyk support to have the duplicate user removed.</p> |
| Protocol binding              | HTTP-POST is recommended, HTTP-Redirect is also supported                                                                                                                                                                                                                                                                                                                              |
| IdP initiated flow supported? | Idp-initiated flows carry a security risk and are therefore not recommended. Make sure you understand the risks before enabling                                                                                                                                                                                                                                                        |
| Email domains and subdomains  | The email domains and subdomains that need access to the SSO                                                                                                                                                                                                                                                                                                                           |

## Set up OpenID Connect (OIDC) for SSO

{% hint style="info" %}
The IdP (or issuer URL) must be publicly reachable. If these cannot be made public then SAML should be used rather than OIDC
{% endhint %}

When using OIDC for the connection between your Identity provider and Snyk, add the Callback/Redirect URIs and OAuth Grant Type in your identity provider to establish trust with Snyk.

| Information                                        | Description                                                                                                    |
| -------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Callback/Redirect URIs (SNYK-US-01 and SNYK-US-02) | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                 |
| Callback/Redirect URIs (SNYK-EU-01)                | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback) |
| Callback/Redirect URIs (SNYK-AU-01)                | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback) |
| OAuth Grant Type                                   | Implicit (or Authorization Code)                                                                               |

## OIDC information to provide to Snyk

Obtain the following information from your identity provider. Provide this information to Snyk to establish trust on the service-provider side.

| Information                  | Description                                                                                                                 |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| Issuer URL                   | The URL of the discovery document of the OpenID Connect provider you want to connect with. This must be publicly reachable. |
| Client ID                    | The public identifier unique for your authorization server                                                                  |
| Client Secret                | Needed only if the IdP does not allow the Implicit grant type                                                               |
| Email domains and subdomains | The email domains and subdomains that need access to the SSO                                                                |

## Set up Entra ID as SSO (via App Registration/OIDC)

When using Entra ID (formerly Azure AD) for the connection between your Identity provider and Snyk, you must add the Redirect URIs in your Identity provider to establish trust with Snyk.

{% hint style="info" %}
Use your Entra ID name when authenticating rather than the SCM user account name, or a connection error can occur.
{% endhint %}

| Information                               | Description                                                                                                    |
| ----------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Redirect URIs (SNYK-US-01 and SNYK-US-02) | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                 |
| Redirect URIs (SNYK-EU-01)                | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback) |
| Redirect URIs (SNYK-AU-01)                | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback) |

## Entra ID information to provide to Snyk

Obtain the following information from your identity provider. Provide this information to Snyk to establish trust on the service-provider side.

| Information               | Description                                                                                                                 |
| ------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| Client ID                 | The public identifier unique for your authorization server                                                                  |
| Client Secret             | The secret for your authorization that grants tokens to authorized requestors                                               |
| Microsoft Entra ID Domain | The numbers and letters shown in the Directory (tenant) ID, which can be found from the Snyk app you created under Overview |

## Set up ADFS as SSO

When using Active Directory Federation Service (ADFS) for the connection between your Identity provider and Snyk, add the Realm Identifier, a Callback URL, and a Signing certificate in your Identity provider to establish trust with Snyk. For more information, see [Connecting Auth0 to an ADFS server (video)](https://www.youtube.com/watch?v=ICW6sGP9ht8).

| Realm Identifier (SNYK-US-01 and SNYK-US-02) | urn:auth0:snyk                                                                                                                                                   |
| -------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Realm Identifier (SNYK-EU-01)                | urn:auth0:snyk-mt-eu-prod-1                                                                                                                                      |
| Realm Identifier (SNYK-AU-01)                | urn:auth0:snyk-mt-au-prod-1                                                                                                                                      |
| Callback URL (SNYK-US-01 and SNYK-US-02)     | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                                                                   |
| Callback URL (SNYK-EU-01)                    | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback)                                                   |
| Callback URL (SNYK-AU-01)                    | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback)                                                   |
| Signing cert (SNYK-US-01 and SNYK-US-02)     | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem) (add as a signature and not encryption)                                                                 |
| Signing cert (SNYK-EU-01)                    | [https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection) (add as a signature and not encryption) |
| Signing cert (SNYK-AU-01)                    | [https://snyk-mt-eu-prod-1.au.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.au.auth0.com/pem?cert=connection) (add as a signature and not encryption) |

## ADFS information to provide to Snyk

Obtain the following information from your Identity provider. Provide this information to Snyk in order to establish trust on the service-provider side: ADFS URL or Federation Metadata XML file

For Enterprise plans, Snyk can map new users to a specific Organization and role when they first sign in using SSO. This option requires additional configuration, including specific naming conventions for organizations.

{% hint style="info" %}
Work with your Snyk account team to prepare for implementing this SSO option.
{% endhint %}

## Complete SSO connection

After you set up the connection with your Identity provider and provide the necessary details to Snyk Support, Snyk sends you a link to generate a payload.

{% hint style="info" %}
Ignore any error message you see after clicking this link the first time, as Snyk uses the generated payload to complete the configuration.
{% endhint %}

When Snyk finishes the configuration, the support agent asks you to navigate to the login page in incognito mode to prevent cookies from interfering with the login process.

Use [https://app.snyk.io/login/sso](https://app.snyk.io/login/sso) (or the regional equivalent) for logging into your production environment.

To complete your login:

1. Enter your email address.
2. Select **Continue to provider**.
3. Log in with your identity provider as you would for other applications.
4. Let Snyk Support know which user to promote as the Group administrator.

## Resources for SSO setup

These worksheets include the information to enter in your Identity provider and the information you need to collect before submitting a ticket to Snyk Support to request single sign-on. Please make sure you replace any Snyk URL with the correct regional environment per the tables above.

{% file src="../../../.gitbook/assets/SSO Azure Worksheet (3).pdf" %}

{% file src="../../../.gitbook/assets/SSO SAML Worksheet (1).pdf" %}

{% file src="../../../.gitbook/assets/SSO ADFS Worksheet (1).pdf" %}

{% file src="../../../.gitbook/assets/SSO OIDC Worksheet (1).pdf" %}
