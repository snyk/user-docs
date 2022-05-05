# Self Serve Single Sign-On (SSO)

Group Admins of Business plan who use SAML for SSO can configure Snyk Single Sign-on by themselves.

Ensure you have at least one group and organization to indicate where new users will be assigned. See [Groups, Organizations and Users](https://github.com/snyk/user-docs/blob/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/introducing-snyk/snyks-core-concepts/groups-organizations-and-users.md).

### Use SAML for SSO

#### Process overview

The process of establishing trust between your identity provider (IdP) and Snyk requires a few steps by the Group Admin:

1. **Configure your identity provider (IdP)** by using the details about the Snyk environment that is displayed on-screen and user attributes.
2. **Enter SAML attributes** from your identity provider(IdP) in the Group SSO Settings page.
3. **Configure Snyk SSO settings**, choosing how you want your members to log in.
4. Verify SSO login to confirm the login process is working correctly.

{% hint style="info" %}
After SSO is configured both from Snyk and your company's network, a trust relationship is established between Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.
{% endhint %}

**User log in**

Users are provisioned to Snyk when they log in (see [choose-a-provisioning-option.md](choose-a-provisioning-option.md "mention")). If the new user role is selected as Group Member, they only see a list of your organizations until the admin adds them to the appropriate organizations.

### Step 1. Configure your identity provider

To establish trust with Snyk, add an Entity ID, an ACS URL, and a Signing certificate in your identity provider(IdP).

* The **Entity ID** is the URL that uniquely identifies Snyk as a SAML entity or service provider--note, **default Entity ID must be checked** manually as no default is set for this.
* The **Assertion Consumer Service (ACS)** is the endpoint on the Snyk network that listens for requests from your identity provider to enable communication between users on your network and Snyk. This URL is sometimes called a Reply URL.
* The **Signing certificate** is the Snyk certificate, stored on your server, that is needed to maintain the trust relationship. It contains the necessary encryption keys for authentication.

**Group Admins** can find the Snyk details required to set up the connection with your Identity provider (IdP) here:

To access the Group Overview for your group, click on Settings [![](https://github.com/snyk/user-docs/raw/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/.gitbook/assets/image%20\(70\).png)](https://github.com/snyk/user-docs/blob/118bd8f19001bd64415f0ce63897f568c4b5327a/docs/.gitbook/assets/image%20\(70\).png) > **SSO**:

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 14.32.24.png>)

To map information from your Identity provider to Snyk, name your user attributes as follows (using the same capitalization and spelling).

| Attribute | Description                                     |
| --------- | ----------------------------------------------- |
| email     | The user’s email address                        |
| name      | The name of the person to be authenticated      |
| username  | The person’s username for the identity provider |

An example from Okta is given below:

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 14.19.18.png>)

{% hint style="warning" %}
If your user attributes do not match, the Snyk configuration for your SSO will not work.
{% endhint %}

For more details about these attributes, see [Step 3. Snyk SSO settings.](self-serve-single-sign-on-sso.md#step-3.-snyk-sso-setting)

### Step 2. Enter SAML Attributes

After the identity provider is set up to acknowledge Snyk, now obtain the following information from your identity provider and organization.

Click **create a connection** to establish trust on the service-provider side:

![](<../../../.gitbook/assets/image (66) (1) (1) (2).png>)

Provide SAML attributes in the below form.

* **Sign in URL** (mandatory): Identity Provider sign-in URL
* **Sign out URL**: Redirect URL when user signs out of Snyk, Recommended
* **X509 signing certificate** (mandatory): The identity provider public key. Download the certificate from your identity provider and paste it here. The system will encode it in _**Base64 format.**_
* **Email domains and sub-domains** that would need SSO access (mandatory)
* **Protocol binding**: HTTP-POST is recommended; HTTP-Redirect is also supported
* **IdP-Initiated workflow**: Enable this option to add Snyk tile to your Identity Provider. \
  **Note:** IdP-Initiated SSO behavior carries a [security risk](https://auth0.com/docs/authenticate/protocols/saml/saml-sso-integrations/identity-provider-initiated-single-sign-on#risks-and-considerations) and is therefore not recommended. The risk is explained on the IdP side, and should be understood before enabling this option.

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 14.40.24.png>)

After filling in the details click on **Save Changes**. Snyk will highlight if there are any errors. You can edit SAML attributes at any time, save your changes and it will reflect in the SSO connection immediately.

### Step 3. Snyk SSO settings

Now click **Configure Snyk SSO settings below** in the success banner to complete the setup.

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 15.37.44.png>)

Choose the new user’s role (see [choose-a-provisioning-option.md](choose-a-provisioning-option.md "mention")):

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 15.28.30.png>)

The **Profile attributes** values are used to map the user's SAML payload data, to ensure that Snyk receives the proper email, name, and userName. It should be the exact keys from the raw json from SAML payload.

The values will be auto-populated by the system as **nameIdAttributes.value**.

{% hint style="info" %}
We suggest that you consult your SSO administrator and determine how the email address, name (a display name, First name + Last Name), and username (a unique identifier) are represented when the Identity Provider was configured (as described in [#1.-configure-your-identity-provider](self-serve-single-sign-on-sso.md#1.-configure-your-identity-provider "mention"))
{% endhint %}

After the selections are made, click **Save Changes** to complete your SSO configuration.

### 4. Verify SSO login

Once the SSO settings are saved successfully, click on the direct login URL displayed on the screen for the SSO connection you set up. And you would be taken to your IDP through which you can login to your Snyk Group.

![](<../../../.gitbook/assets/Screenshot 2022-02-24 at 16.00.49.png>)

Alternatively, you can login to [**snyk**](http://snyk.io) in incognito mode to prevent cookies from interfering and verify the SSO Login is working.

(Direct SSO Link:[https://app.snyk.io/login/sso](https://app.snyk.io/login/sso)) or:

1. [https://snyk.io](https://snyk.io)
2. Log in
3. SSO
4. Enter your email address
5. Continue to your identity provider and log in.
