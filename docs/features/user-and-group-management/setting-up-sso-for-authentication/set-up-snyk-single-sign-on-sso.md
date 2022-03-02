# Set up Snyk Single Sign-On (SSO)

Set up single sign-on to allow your developers and teams easy access to Snyk through your existing SSO provider, so they can see the status of their projects, view reports, resolve vulnerabilities, and more.

The information you need to establish trust between Snyk and the identity provider depends on which type of SSO you are using.

Ensure you have at least one group and organization to indicate where new users will be assigned. See [Groups, Organizations and Users](https://docs.snyk.io/introducing-snyk/snyks-core-concepts/groups-organizations-and-users).

{% hint style="info" %}
Once you have gathered the requested information (detailed below), create a support ticket to request SSO set up.

For Business plan using SAML, Group Admins can configure Snyk Single Sign-On and [this guide](self-serve-single-sign-on-sso.md) details the steps.
{% endhint %}

## Overview

The process of establishing trust between your identity provider (IdP) and Snyk requires a few separate steps, coordinated between your SSO administrator and Snyk Support.

* In your own identity provider platform, enter details about the Snyk environment and user attributes.
* Provide Snyk with details from your IdP.
* Set up a user for testing and send Snyk the username and password for that user.
* Use the link provided by Snyk to generate a payload.
* After Snyk finalizes the connection, confirm the login process is working correctly.
* Users are provisioned to Snyk when they log in. If invitation is required, users may only see a list of your organizations until the admin adds them to the appropriate organizations.

After SSO is configured both from Snyk and your company's network, a trust relationship is established between Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.

Each type of SSO connection requires different details for establishing the trust between your identity provider and Snyk. The following sections explain those details. The details are also included in the worksheets in the Resources section at the end of this article.

## Use SAML for SSO

To establish trust with Snyk, add an Entity ID, an ACS URL, and a Signing certificate in your identity provider.

* The **Entity ID** is the URL that uniquely identifies Snyk as a SAML entity or service provider--note, **default Entity ID must be checked** manually as no default is set for this.
* The **Assertion Consumer Service (ACS)** is the endpoint on the Snyk network that listens for requests from your identity provider to enable communication between users on your network and Snyk. This URL is sometimes called a Reply URL.
* The **Signing certificate** is the Snyk certificate, stored on your server, that is needed to maintain the trust relationship. It contains the necessary encryption keys for authentication.

Use these details to set up the connection with your Identity provider (IdP):

| Details             | **Description**                                                                                                                                                                                          |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Entity ID           | **urn:auth0:snyk:saml-**_**{customer\_name}**_ (Replace _{customer\_name}_ with a name for your company.)                                                                                                |
| ACS URL             | [https://snyk.auth0.com/login/callback?connection=saml-](https://snyk.auth0.com/login/callback?connection=saml-)_**{customer\_name}**_ (Replace _{customer\_name}_ with the same name for your company.) |
| Signing certificate | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem)                                                                                                                                                 |

To map information from your Identity provider to Snyk, name your user attributes as follows (using the same capitalization and spelling):

| **Attribute** | **Description**                                     |
| ------------- | --------------------------------------------------- |
| **email**     | The user email address                              |
| **name**      | The name of the person to be authenticated          |
| **username**  | **The person’s username for the identity provider** |

If your user attributes do not match, note that the Snyk configuration for your SSO will take more time.

## SAML information to provide to Snyk

Obtain the following information from your identity provider and organization. Provide this information to Snyk to establish trust on the service-provider side.

| Information                   | Description                                                                                                                      |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Sign-in URL                   | The URL for your identity provider sign-in page                                                                                  |
| X509 Signing Certificate      | The identity provider public key, encoded in Base64 format                                                                       |
| Sign-out URL                  | <p>Optional, but recommended -</p><p>The URL for redirect whenever a user logs out of Snyk</p>                                   |
| User ID attribute             | Optional default is **http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier**                                     |
| Protocol binding              | HTTP-POST is recommended, HTTP-Redirect is also supported                                                                        |
| IdP initiated flow supported? | Idp-initiated flows carry a security risk and  are therefore not recommended. Make sure you understand the risks before enabling |
| Email domains and subdomains  | The email domains and subdomains that need access to the SSO                                                                     |

## Use OpenID Connect (OIDC) for SSO

When using OIDC for the connection between your Identity provider and Snyk, add the Callback/Redirect URIs and OAuth Grant Type in your identity provider to establish trust with Snyk.

| Information            | Description                                                                    |
| ---------------------- | ------------------------------------------------------------------------------ |
| Callback/Redirect URIs | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback) |
| OAuth Grant Type       | Implicit (or Authorization Code)                                               |

## OIDC information to provide to Snyk

Obtain the following information from your identity provider and organization. Provide this information to Snyk to establish trust on the service-provider side.

| Information                  | Description                                                                               |
| ---------------------------- | ----------------------------------------------------------------------------------------- |
| Issuer URL                   | The URL of the discovery document of the OpenID Connect provider you want to connect with |
| Client ID                    | The public identifier unique for your authorization server                                |
| Client Secret                | Only needed if the IdP does not allow the Implicit grant type                             |
| Email domains and subdomains | The email domains and subdomains that need access to the SSO                              |

## Use Azure AD as SSO (via App Registration/OIDC)

When using Azure AD for the connection between your Identity provider and Snyk, you'll need to add the Redirect URIs in your Identity provider to establish trust with Snyk.

{% hint style="info" %}
Use your Azure AD name when authenticating rather than the SCM user account name or a connection error can occur.
{% endhint %}

| Information   | Description                                                                    |
| ------------- | ------------------------------------------------------------------------------ |
| Redirect URIs | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback) |

## Azure AD information to provide to Snyk

Obtain the following information from your identity provider and organization. Provide this information to Snyk to establish trust on the service-provider side.

| Information               | Description                                                                                                                 |
| ------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| Client ID                 | The public identifier unique for your authorization server                                                                  |
| Client Secret             | The secret for your authorization that grants tokens to authorized requestors                                               |
| Microsoft Azure AD Domain | The numbers and letters shown in the Directory (tenant) ID, which can be found from the Snyk app you created under Overview |

## Use ADFS as SSO

When using ADFS for the connection between your Identity provider and Snyk, add the Realm Identifier, a Callback URL, and a Signing certificate in your Identity provider to establish trust with Snyk.

| Information      | **Description**                                                                                  |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Realm Identifier | **urn:auth0:snyk**                                                                               |
| Callback URL     | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                   |
| Signing cert     | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem) (add as a signature and not encryption) |

## ADFS information to provide to Snyk

Obtain the following information from your Identity provider and organization. You'll provide this information to Snyk in order to establish trust on the service-provider side.

* ADFS URL or Federation Metadata XML file

## Map Enterprise users

For Enterprise plans, Snyk can map new users to a specific organization and role when they first sign in using SSO. This option requires additional configuration, including specific naming conventions for organizations.

Work with your Customer Success Manager and Snyk Technical Services to prepare for implementing this SSO option.

## Complete SSO connection

After you set up the connection with your Identity provider and provide the necessary details to Snyk Support, Snyk sends you a link to generate a payload.

Ignore any error message you see after clicking this link the first time, as Snyk uses the generated payload to complete the configuration.

When Snyk finishes the configuration, the support agent asks you to access the log in page in incognito mode to prevent cookies from interfering with the log in process.

* Use [https://app.snyk.io/login/sso](https://app.snyk.io/login/sso) for logging into your production environment.
* Use [https://test.snyk.io](https://test.snyk.io) for logging into the test environment.

To complete your log in:

1. Enter your email address.
2. Select **Continue to provider**.
3. Log in with your identity provider as you would for other applications.
4. Let Snyk Support know which user to promote as the Group administrator.

## Resources

These worksheets include the information to enter in your Identity provider and the information you need to collect before submitting a ticket to Snyk Support to request single sign-on.

{% file src="../../../.gitbook/assets/sso-azure-worksheet.pdf" %}

{% file src="../../../.gitbook/assets/sso-saml-worksheet.pdf" %}

{% file src="../../../.gitbook/assets/sso-adfs-worksheet.pdf" %}

{% file src="../../../.gitbook/assets/sso-oidc-worksheet.pdf" %}
