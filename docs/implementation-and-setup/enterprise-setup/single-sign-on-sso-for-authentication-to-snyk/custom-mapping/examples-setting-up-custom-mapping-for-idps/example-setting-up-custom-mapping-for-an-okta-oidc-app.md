# Example: setting up custom mapping for an Okta OIDC app

Follow these steps configure an integration for OIDC Okta.

## Create an Okta OIDC app

1.  In Okta, select **Applications** > **Applications** > **Create App Integration** then choose **OICD OpenID Connect** and **Web Application.**

    <figure><img src="../../../../../.gitbook/assets/1 (2).png" alt="Create a new app integration in Okta"><figcaption><p>Create a new app integration in Okta</p></figcaption></figure>
2.  In the next step add an **App integration name** for your OIDC application, check the **Implicit** **Grant Type** and add the **Sign-in redirect URI** relevant to your [Snyk platform deployment](../../set-up-snyk-single-sign-on-sso.md). Remove the placeholder **Sign-out redirect URI** and choose your assignment access control before clicking **Save.**

    <figure><img src="../../../../../.gitbook/assets/2 (1) (1).png" alt="Provide details for new web app integration"><figcaption><p>Provide details for new web app integration</p></figcaption></figure>
3. On the application page that opens after saving, copy the details identified in  [OIDC information to provide to Snyk](../../set-up-snyk-single-sign-on-sso.md#oidc-information-to-provide-to-snyk) and provide to your Snyk contact:
   * Client ID
   * If you are not using the Implicit Grant type, the client secret
4. Also, share with Snyk the Issuer URL/domain. This is typically the URL you find in your browser address bar without "-admin", for example, https://customer.example.okta.com. It can also be found under the **Sign-On** tab of your application by editing the **OpenID Connect ID Token** from **Dynamic** Issuer to **Okta URL**.

If you wish to set up custom mapping, move on to the next section of this guide.

## Add custom mapping

[Custom mapping](../) for an OIDC application in Okta is easily managed through custom attributes on the Group level.

### Create a custom app user attribute to contain both the Snyk Organization name and role

1. In Okta, select your newly created OIDC application user profile under **Directory** > **Profile editor.**
2.  Select **+Add Attribute.**

    <figure><img src="../../../../../.gitbook/assets/3 (1).png" alt="Okta profile editor"><figcaption><p>Okta profile editor</p></figcaption></figure>
3.  In the corresponding fields, add the following details for this Attribute and click **Save**:\
    **Data type**: string array\
    **Display name**: Snyk roles\
    **Variable name:** roles\
    **Group Priority**: Combine values across groups

    <figure><img src="../../../../../.gitbook/assets/4 (2).png" alt="Attribute details"><figcaption><p>Attribute details</p></figcaption></figure>

### Assign the attribute to the relevant Okta groups

1. On the main page of Okta select **Directory** > **Groups**.
2.  Select a **Group**, navigate to the **Applications** tab, click **Assign** **application** if not already assigned, and choose your Snyk OIDC app,. Then click on the pencil next to the displayed Snyk OIDC app.

    <figure><img src="../../../../../.gitbook/assets/5 (4).png" alt="Group selected for modicification"><figcaption><p>Group selected for modicification</p></figcaption></figure>
3.  In the **Edit App Assignment** dialog, add the Snyk org name + role associated with your Okta group (no spaces or capital letter(s)), following the syntax explained in [custom mapping](../) (or [legacy custom mapping](../legacy-custom-mapping.md) if using the legacy mapping option). Example, `snyk:org:*:org_admin`.\


    <figure><img src="../../../../../.gitbook/assets/image (211).png" alt="Adding Snyk roles"><figcaption><p>Adding Snyk roles</p></figcaption></figure>
4. Repeat the preceding steps for all your applicable Okta groups to assign the org name and role combination to each user within each configured group.
