# Configure SSO

Configuring SSO may be done while onboarding projects, creating organization structure, or even while configuring integrations. Disable Notifications for the organizations if your intent is to add users prior to onboarding your projects into Snyk.&#x20;

Setting up [Single Sign On with Snyk](../../using-single-sign-on-sso-for-authentication/) gives your users a consistent way of logging in to your application, whilst increasing your ability to control who has access to your organizations.

In most cases, we recommend using the [Self-Serve option](../../using-single-sign-on-sso-for-authentication/self-serve-single-sign-on-sso/) that allows you to configure a SAML connection to your identity provider. This allows you to configure the email domains that are valid, and also the default permissions that new users will have in your organizations.

{% hint style="info" %}
**Getting Assistance**\
Depending on your license with Snyk, you may not see the ‘SSO’ option with your Group Settings page. If it is not displayed please [raise a support ticket](https://support.snyk.io/hc/en-us) and we will enable it. Also, if you want to use a different SSO configuration option (e.g. OIDC) then you will need to [raise a support ticket](https://support.snyk.io/hc/en-us) with your request so a member of the Snyk support team can assist.
{% endhint %}

{% hint style="info" %}
[**Custom Mapping**](../../using-single-sign-on-sso-for-authentication/custom-mapping-option/)

* This option requires paid professional services, and cannot be completed using the Self-Serve option. Please [contact the Snyk support team](https://support.snyk.io/hc/en-us) or your Snyk account team to purchase the required services.
{% endhint %}

{% hint style="info" %}
#### Provision Users via API

* The [Provision user endpoints](https://docs.snyk.io/snyk-admin/manage-users-in-organizations-and-groups/provision-users-to-orgs-using-the-snyk-api-v1) in the Snyk API v1 allows you to grant organization-level access and permissions to users before they log in to the Snyk platform.&#x20;
* API provisioning also allows you to limit organization access and assign custom member roles. This allows you to control:
  * Which _role_ is assigned to each user.
  * which _orgs_ they should have access to.&#x20;
{% endhint %}
