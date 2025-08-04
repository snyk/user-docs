# Configure SSO

Configuring SSO may be done while onboarding Projects, creating Organization structure, or even while configuring integrations. Disable notifications for the Organizations if your intent is to add users prior to onboarding your Projects into Snyk.

Setting up [Single Sign On with Snyk](../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/) gives your users a consistent way of logging in to your application while increasing your ability to control who has access to your Organizations.

In most cases, Snyk recommends using the [Self-Serve Single Sign-On option](../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/configure-self-serve-single-sign-on-sso/) that allows you to configure a SAML connection to your identity provider. This allows you to configure the email domains that are valid, and also the default permissions that new users will have in your Organizations.

## How to get assistance

Depending on your license with Snyk, you may not see the SSO option on your Group Settings page. If it is not displayed, [submit a request to Support](https://support.snyk.io) to enable it. This option lets you configure SSO with SAML, but if you want to use a different SSO configuration option, for example, OIDC, [submit a request to Snyk Support](https://support.snyk.io), and a member of the Snyk support team will assist.

## Assistance with SSO Custom Mapping

Using this option requires paid professional services and cannot be completed using the Self-Serve SSO option. [Contact the Snyk support team](https://support.snyk.io) or your Snyk account team to purchase the required services.

## How to provision users through the API

The [Provision user endpoints in the Snyk API](../../../snyk-platform-administration/user-management-with-the-api/provision-users-to-organizations-using-the-api.md) allow you to grant Organization-level access and permissions to users before they log in to the Snyk platform.

API provisioning also allows you to limit Organization access and assign custom member roles. This allows you to control:

* Which role is assigned to each user
* Which Organizations each user should have access to
