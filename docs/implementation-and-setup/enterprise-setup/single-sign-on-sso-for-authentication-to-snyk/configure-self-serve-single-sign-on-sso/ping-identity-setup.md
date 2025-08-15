# Ping Identity setup

This page explains how to set up a Ping Identity Application and connect it to Snyk to facilitate SSO.

Before configuring your Ping Identity Application to use SSO with Snyk, obtain an entity ID and a reply URL (Assertion Consumer Service URL) from Snyk. Then follow these steps:

1.  In the left menu, select your **Group**, then **Settings**.

    <figure><img src="../../../../.gitbook/assets/Screenshot 2023-09-05 at 10.54.23 AM.png" alt="" width="375"><figcaption></figcaption></figure>
2.  Select **SSO** and copy the values under **Entity ID** and **ACS URL** or leave the browser tab open for easy access.

    <figure><img src="../../../../.gitbook/assets/2 (2).png" alt="Group Settings: SSO"><figcaption><p>Group Settings: SSO</p></figcaption></figure>
3.  Navigate to Ping Identity and select **Applications** in the **Connections** menu. Click on the plus sign to create a new application.&#x20;

    <figure><img src="../../../../.gitbook/assets/1 (6).png" alt="Create a new application"><figcaption><p>Create a new Application</p></figcaption></figure>
4.  Name your application appropriately, select **SAML Application**, and click **Configure.**

    <figure><img src="../../../../.gitbook/assets/2 (5).png" alt="Configure as SAML Application" width="563"><figcaption><p>Configure as SAML Application</p></figcaption></figure>
5.  Enter the details you copied from Snyk, the **ACS URL** and **Entity ID,** and select **Save**.

    <figure><img src="../../../../.gitbook/assets/3 (5).png" alt="Add Snyk configuration details" width="563"><figcaption><p>Add Snyk configuration details</p></figcaption></figure>
6.  Select **Configuration** and download the signing certificate in PEM format.

    <figure><img src="../../../../.gitbook/assets/4 (4).png" alt="Download signing certificate"><figcaption><p>Download signing certificate</p></figcaption></figure>
7.  Scroll further down and copy the **Single Signon Service** details.

    <figure><img src="../../../../.gitbook/assets/5 (7).png" alt="Copy the Single Signon Service details"><figcaption><p>Copy the Single Signon Service details</p></figcaption></figure>
8.  Return to the the Snyk portal and paste the single sign-in URL copied at step 2 into the **Sign in URL** field. \


    <figure><img src="../../../../.gitbook/assets/single-sign-on-URL-field.png" alt="Paste Sign in URL"><figcaption><p>Paste Sign in URL</p></figcaption></figure>
9.  Open the downloaded certificate in your preferred text editor, copy the text and paste it into the Snyk **X509 signing certificate** field, and add the relevant domains that are supported by this SSO connection.\
    Finally, verify if an IdP-initiated workflow should be enabled and then click **Create Auth0 connection** if you are creating a completely new connection or **Save changes** if you are editing an existing connection.&#x20;

    <figure><img src="../../../../.gitbook/assets/Screenshot 2023-09-05 at 11.01.53 AM.png" alt="Enter the Ping Identity details"><figcaption><p>Enter the Ping Identity details</p></figcaption></figure>
10. In Ping Identity, select **Attribute mappings** and click the pencil to edit.

    <figure><img src="../../../../.gitbook/assets/6 (5).png" alt="Edit attribue mappings"><figcaption><p>Edit attribue mappings</p></figcaption></figure>
11. Click the cog icon and add the following attributes:

    **email**: Email Address\
    **username**: Username\
    **name**: the expression `user.name.given + ' ' + user.name.famil`y; click the cog icon to enter an advanced description.&#x20;

    <figure><img src="../../../../.gitbook/assets/7 (5).png" alt="Add attribute mappings"><figcaption><p>Add attribute mappings</p></figcaption></figure>

    <figure><img src="../../../../.gitbook/assets/8 (5).png" alt="Add an advanced expression for the name attribute"><figcaption><p>Add an advanced expression for the name attribute</p></figcaption></figure>
12. In the Snyk portal, decide how new users should be treated when signing in and choose the option you would like to use: **Group member**, **Org collaborator**, or **Org admin**.
13. Change the profile attributes to the attribute names you entered in Ping Identity then click **Save changes.**\


    <figure><img src="../../../../.gitbook/assets/Screenshot 2023-09-05 at 11.07.37 AM.png" alt="Step 3 Snyk SSO settings"><figcaption><p>Step 3 Snyk SSO settings</p></figcaption></figure>
14. Verify you can log in, either with the direct URL at the top of **Step 3 Snyk SSO settings** (not shown in the image) or by going to the [generic SSO login](https://app.snyk.io/login/sso).
15. As a final step, enable the application and assign it to users.

    <figure><img src="../../../../.gitbook/assets/10 (2).png" alt="Enable and assign the application to users"><figcaption><p>Enable and assign the application to users</p></figcaption></figure>
