# Entra ID Enterprise application setup

This example shows setting up an Entra ID (formerly Azure AD) Enterprise Application and connecting this to Snyk to facilitate SSO. To configure your Azure Enterprise Application to use SSO with Snyk, first obtain an entity ID and a reply URL (Assertion Consumer Service URL) from Snyk.

1.  From the dropdown at the top left select **GROUP OVERVIEW** and then the cog icon (top right corner) to get to your group settings.

    <figure><img src="../../../../.gitbook/assets/5.png" alt="Select group overview"><figcaption><p>Select group overview</p></figcaption></figure>
2.  Click on **SSO** and copy the values under **Entity ID** and **ACS URL** or leave the browser tab open for easy access.

    <figure><img src="../../../../.gitbook/assets/2 (2).png" alt="Group Settings: SSO"><figcaption><p>Group Settings: SSO</p></figcaption></figure>
3.  Navigate to Azure and open Entra ID.

    <figure><img src="../../../../.gitbook/assets/3 (4).png" alt="Entra ID Default Directory"><figcaption><p>Entra ID Default Directory</p></figcaption></figure>
4.  Click **Add** then **Enterprise application**.

    <figure><img src="../../../../.gitbook/assets/4 (1).png" alt="Add Enterprise application"><figcaption><p>Add Enterprise application</p></figcaption></figure>
5.  Choose **Create your own application**.

    <figure><img src="../../../../.gitbook/assets/5 (5).png" alt="Create your own application"><figcaption><p>Create your own application</p></figcaption></figure>
6.  Name the application appropriately, for example, **Snyk-SSO**, making sure that **Integrate any other application you don't find in the gallery (Non-gallery)** is selected and then click **Create**.

    <figure><img src="../../../../.gitbook/assets/6 (1).png" alt="Application name and integration"><figcaption><p>Application name and integration</p></figcaption></figure>
7.  For the new app, select **Set up single sign on** and **Get started**.

    <figure><img src="../../../../.gitbook/assets/7 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Set up single sign-on, Get started"><figcaption><p>Set up single sign-on, Get started</p></figcaption></figure>
8.  Select **SAML** as the SSO method.

    <figure><img src="../../../../.gitbook/assets/8 (1) (1).png" alt="Select SAML"><figcaption><p>Select SAML</p></figcaption></figure>
9.  Click **Edit** under **Basic SAML configuration**.

    <figure><img src="../../../../.gitbook/assets/9 (1).png" alt="Edit basic SAML configuration"><figcaption><p>Edit basic SAML configuration</p></figcaption></figure>
10. Add the Identity (Entity ID) and reply URL (Assertion Consumer Service URL) you obtained from Snyk and click **Save**; then close the edit window.

    <figure><img src="../../../../.gitbook/assets/10.png" alt="Entity ID and Assertion Consumer Service URL"><figcaption><p>Entity ID and Assertion Consumer Service URL</p></figcaption></figure>
11. Scroll to find the login URL needed to finish the configuration in Snyk. Copy it and paste it into the SSO settings in the Snyk portal.

    <figure><img src="../../../../.gitbook/assets/11 (1).png" alt="Login URL"><figcaption><p>Login URL</p></figcaption></figure>

    <figure><img src="../../../../.gitbook/assets/1 (4).png" alt="Sign in URL in Snyk portal"><figcaption><p>Sign in URL in Snyk portal</p></figcaption></figure>
12. Return to Entra ID and click **Download** next to **Certificate (Base64)**.

    <figure><img src="../../../../.gitbook/assets/13.png" alt="Download SAML Certificate (Base 64)"><figcaption><p>Download SAML Certificate (Base 64)</p></figcaption></figure>
13. Open the downloaded certificate in your preferred text editor, copy the text and paste it into the Snyk **X509 signing certificate** field, and add the relevant domains that are supported by this SSO connection.\
    Finally, verify if an **IdP-initiated workflow** should be enabled and then click **Create Auth0 connection** if you are creating a completely new connection or **Save changes** if you are editing an existing connection.

    <figure><img src="../../../../.gitbook/assets/14.png" alt="Enter certificate and domains supported, set connection"><figcaption><p>Enter certificate and domains supported, set connection</p></figcaption></figure>
14. Decide how new users should be treated when signing in and choose the option you would like to use: **Group member**, **Org collaborator**, or **Org admin**. Finally, modify the **profile attributes** if your settings in Azure deviate from the default; then click **Save changes** and verify you can log in, either with the direct URL at the top of step 3 or by going to the [generic SSO login](https://app.snyk.io/login/sso).\
    \
    If you are not receiving profile values as expected, you may need to add email, name, and username as additional claims within Azure SSO settings and then map those accordingly in the Snyk SSO **Profile attributes** section.

    <figure><img src="../../../../.gitbook/assets/claim1.png" alt="Azure claim settings"><figcaption><p>Azure claim settings</p></figcaption></figure>

    <figure><img src="../../../../.gitbook/assets/image2.png" alt="Profile attributes section"><figcaption><p>Profile attributes section</p></figcaption></figure>

If you wish to add signature verification of the incoming Snyk request:

1. Download the **Signing certificate** at step 1 of the Snyk SSO settings.
2. Use the following openssl command to convert it to .cer-format `openssl x509 -outform DER -in snyk.pem -out snyk.cer`
3. At the bottom of the **SAML Certificates** settings of your SSO app in Active Directory, click **Edit** next to **Verification certificates.**
4. Check **Require verification certificates** and upload the certificate from the output of the above openssl command and click **Save**.
