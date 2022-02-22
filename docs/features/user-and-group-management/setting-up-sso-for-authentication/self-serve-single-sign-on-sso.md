# Self Serve Single Sign-On (SSO)

Group Admins of Business plan who use SAML for SSO can configure Snyk Single Sign-on by themselves.

Ensure you have at least one group and organization to indicate where new users will be assigned. See [Groups, Organizations and Users](https://github.com/snyk/user-docs/blob/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/introducing-snyk/snyks-core-concepts/groups-organizations-and-users.md).

### Use SAML for SSO

#### Process overview

The process of establishing trust between your identity provider (IdP) and Snyk requires a few steps by the Group Admin:

1. Configure your identity provider (IdP) by using the details about the Snyk environment that is displayed on-screen and user attributes.
2. Set up your Auth0 connection in the Group SSO Settings page.
3. Set up user provisioning, choosing how you want your members to log in.
4. Verify SSO login to confirm the login process is working correctly.

{% hint style="info" %}
After SSO is configured both from Snyk and your company's network, a trust relationship is established between Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.
{% endhint %}

**User log in**

Users are provisioned to Snyk when they log in (see [choose-a-provisioning-option.md](choose-a-provisioning-option.md "mention")). If the new user role is selected as Group Member, they only see a list of your organizations until the admin adds them to the appropriate organizations.

### 1. Configure your identity provider

To establish trust with Snyk, add an Entity ID, an ACS URL, and a Signing certificate in your identity provider(IdP).

* The **Entity ID** is the URL that uniquely identifies Snyk as a SAML entity or service provider--please note, **default Entity ID must be checked** manually as no default is set for this.
* The **Assertion Consumer Service (ACS)** is the endpoint on the Snyk network that listens for requests from your identity provider to enable communication between users on your network and Snyk. This URL is sometimes called a Reply URL.
* The **Signing certificate** is the Snyk certificate, stored on your server, that is needed to maintain the trust relationship. It contains the necessary encryption keys for authentication.

**Group Admins** can find the Snyk details required to set up the connection with your Identity provider (IdP) here:

To access the Group Overview for your group, click on Settings [![](https://github.com/snyk/user-docs/raw/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/.gitbook/assets/image%20\(70\).png)](https://github.com/snyk/user-docs/blob/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/.gitbook/assets/image%20\(70\).png) > **SSO**:

![](<../../../.gitbook/assets/image (76) (1).png>)

To map information from your Identity provider to Snyk, name your user attributes as follows (using the same capitalization and spelling).

| Attribute | Description                                     |
| --------- | ----------------------------------------------- |
| email     | The user’s email address                        |
| name      | The name of the person to be authenticated      |
| username  | The person’s username for the identity provider |

An example from Okta is given below:

![](<../../../.gitbook/assets/image (95).png>)

{% hint style="warning" %}
If your user attributes do not match, the Snyk configuration for your SSO will not work.
{% endhint %}

For more details about these attributes, see [#3.-set-up-user-provisioning](self-serve-single-sign-on-sso.md#3.-set-up-user-provisioning "mention").

### 2. Set up your Auth0 connection

After the identity provider is set up to acknowledge Snyk, now obtain the following information from your identity provider and organization.

Click **create a connection** to establish trust on the service-provider side:

![](<../../../.gitbook/assets/image (73).png>)

Provide Auth0 connection details in the below form.

* **Sign in URL** (mandatory): Identity Provider sign-in URL
* **Sign out URL**: Redirect URL when user signs out of Snyk, Recommended
* **X509 signing certificate** (mandatory): The identity provider public key, encoded in _**Base64 format**_
* **Email domains and sub domains** that would need SSO access (mandatory)
* **User ID attribute**: The attribute in SAML token that will be mapped to the user\_id property in Auth0
* **Protocol binding**: HTTP-POST is recommended, HTTP-Redirect is also supported
* **IdP initiated workflow**: Idp-initiated SSO behaviour carry a security risk and therefore not recommended. Make sure you understand the risks before enabling this option.

![](<../../../.gitbook/assets/image (81) (1).png>)

After filling in the details click on **Create Auth0 connection**. Snyk will highlight if there are any errors. In this version you cannot edit Auth0 attributes once set up. If you need to change Auth0 attributes or if you need assistance, you can reach out to the support team using the ‘Contact Support’ option.

### 3. Set up User Provisioning

Now click **Edit Snyk SSO details below** in the success banner to complete the setup.

![](<../../../.gitbook/assets/image (77) (1).png>)

Choose the new user’s role (see [choose-a-provisioning-option.md](choose-a-provisioning-option.md "mention")):

![](<../../../.gitbook/assets/image (79).png>)

The **Profile attributes** values are used to map the user's SAML payload data, to ensure that Snyk receives the proper email, name, and userName. It should be the exact keys from the raw json from SAML payload.

The values will be auto-populated by the system as **nameIdAttributes.value**.

{% hint style="info" %}
We suggest that you consult your SSO administrator and determine how the email address, name (a display name, First name + Last Name), and username (a unique identifier) are represented when the Identity Provider was configured (as described in [#1.-configure-your-identity-provider](self-serve-single-sign-on-sso.md#1.-configure-your-identity-provider "mention"))
{% endhint %}

After the selections are made, click **Save Changes** to complete your SSO configuration.

### 4. Verify SSO login

Once the SSO settings are saved successfully, click on the direct login URL displayed on the screen for the SSO connection you set up. And you would be taken to your IDP through which you can login to your Snyk Group.

![](<../../../.gitbook/assets/Screenshot 2022-02-08 at 12.54.37.png>)



Alternatively,  you can login to [**snyk**](http://snyk.io) in incognito mode to prevent cookies from interfering and verify the SSO Login is working.

(Direct SSO Link:[https://app.snyk.io/login/sso](https://app.snyk.io/login/sso)) or:

1. [https://snyk.io](https://snyk.io)
2. Log in
3. SSO
4. Enter your email address
5. Continue to your provider and log in.
