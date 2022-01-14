# Self Serve Single Sign-On(SSO)

Group Admins of Business plan who use SAML for SSO can configure Snyk Single Sign-on by themselves.

Ensure you have at least one group and organization to indicate where new users will be assigned. See [Groups, Organizations and Users](https://docs.snyk.io/introducing-snyk/snyks-core-concepts/groups-organizations-and-users).

## Use SAML for SSO - Overview

The process of establishing trust between your identity provider (IdP) and Snyk requires a few steps by the Group Admin.

1. Configure your identity provider platform by using the details about the Snyk environment that is displayed on-screen and user attributes.
2. Enter the details required to set up a successful Auth0 connection in the Group SSO Settings page.
3. Complete the setup by choosing how do you want your members to log in
4. After completing the SSO configuration, confirm the login process is working correctly.

Users are provisioned to Snyk when they log in. If new user role is selected as Group Member, they only see a list of your organizations until the admin adds them to the appropriate organizations.

After SSO is configured both from Snyk and your company's network, a trust relationship is established between Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.

## 1. Configure your Identity Provider

To establish trust with Snyk, add an Entity ID, an ACS URL, and a Signing certificate in your identity provider(IdP).

* The **Entity ID** is the URL that uniquely identifies Snyk as a SAML entity or service provider--please note, **default Entity ID must be checked** manually as no default is set for this.
* The **Assertion Consumer Service (ACS)** is the endpoint on the Snyk network that listens for requests from your identity provider to enable communication between users on your network and Snyk. This URL is sometimes called a Reply URL.
* The **Signing certificate** is the Snyk certificate, stored on your server, that is needed to maintain the trust relationship. It contains the necessary encryption keys for authentication.

**Group Admins** can find the Snyk details required to set up the connection with your Identity provider (IdP) here:

Select the Group from the drop down > click on Settings ![](<../../../.gitbook/assets/image (70).png>) > SSO

![](<../../../.gitbook/assets/Screenshot 2022-01-07 at 14.30.25.png>)

To map information from your Identity provider to Snyk, name your user attributes as follows (using the same capitalization and spelling). More details about these attributes can be found under Step 3.

| Attribute | Description                                     |
| --------- | ----------------------------------------------- |
| email     | The user’s email address                        |
| name      | The name of the person to be authenticated      |
| username  | The person’s username for the identity provider |

If your user attributes do not match, note that the Snyk configuration for your SSO will not work.

## 2. Set up Auth0 connection

Once the identity provider is set up to acknowledge Snyk, obtain the following information from your identity provider and organization. Group Admins can then proceed by clicking on ‘create a connection’ link to establish trust on the service-provider side.

![](<../../../.gitbook/assets/Screenshot 2022-01-07 at 16.11.10.png>)

After filling in the details click on ‘Create Auth0 connection’. Snyk will highlight if there are any errors, or you can reach out to the support team if you need assistance using the ‘Contact Support’ option.

## 3. Set up User Provisioning

Upon establishing a successful Auth0 connection, you can click on the ‘Edit Snyk SSOl details below’ link in the success banner to complete the setup.

![](<../../../.gitbook/assets/Screenshot 2022-01-14 at 14.56.56.png>)

Please choose the New User’s Role. More details about the options can be found [here](choose-a-provisioning-option.md)

![](<../../../.gitbook/assets/Screenshot 2022-01-14 at 14.58.19.png>)

The Profile Attributes are used to map the user's SAML payload data to ensure that Snyk receives the proper email, name, and userName. It should be the exact keys from the raw json from SAML payload. The values will be auto-populated by the system as nameIdAttributes.value. It is suggested that you consult your SSO administrator and determine how the email address, name (a display name, First name + Last Name), and username (a unique identifier) are represented when the Identity Provider was configured in Step 1.

Once the selections are made, click ‘Save Changes’ and that completes your SSO configuration.

## 4. Final Steps

Login to [**snyk**](http://snyk.io) in incognito mode to prevent cookies from interfering verify the SSO Login is working.

(Direct Link:[https://app.snyk.io/login/sso](https://app.snyk.io/login/sso)) or:

1. [https://snyk.io](https://snyk.io)
2. Log in
3. SSO
4. Enter your email address
5. Continue to your provider and log in
