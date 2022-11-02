# Self Serve Single Sign-On (SSO)

Group Admins of a Business or Enterprise plan who use SAML for SSO can configure Snyk Single Sign-on by themselves. Ensure you have at least one Group and Organization to indicate where new users will be assigned. See [Groups, Organizations, and Users](../../../../introducing-snyk/groups-organizations-and-users.md).

{% hint style="warning" %}
Self serve SSO does not accommodate [custom mapping](../custom-mapping-option/). To set up custom SSO for your Snyk Group, contact your Snyk account team.
{% endhint %}

The following video demonstrates the process and steps for setting up single sign-on when using SAML.

{% embed url="https://thoughtindustries-1.wistia.com/medias/dyg9opxlv8" %}

## Use SAML for SSO: process overview

The process of establishing trust between your identity provider (IdP) and Snyk requires that the Group Admin do the following:

1. **Configure your identity provider (IdP)** by using the details about the Snyk environment displayed on-screen and user attributes.
2. **Enter SAML attributes** from your identity provider(IdP) in the Group SSO Settings page.
3. **Configure Snyk SSO settings**, choosing how you want your members to log in.
4. Verify SSO login to confirm the login process is working correctly.

{% hint style="info" %}
After SSO is configured both from Snyk and your company's network, a trust relationship is established with Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.&#x20;

Although not all the examples following this page cover verifying the Snyk signature, it is recommended that you improve the trust relationship and ensure integrity even further. Follow your respective IdP's documentation to add SP signature verification where possible.
{% endhint %}

## **User login**

Users are provisioned to Snyk when they log in (see [Choose a provisioning option](../choose-a-provisioning-option.md)). If the new user role selected is Group Member, the new user sees only a list of your Organizations until the admin adds them to the appropriate Organizations.
