# Configure Self-Serve Single Sign-On (SSO)

Group Admins on a Snyk Enterprise plan who use SAML for SSO can configure Snyk Single Sign-on themselves. Ensure you have at least one Group and Organization where you can assign new users. See [Manage Groups and Organizations](../../../../snyk-platform-administration/groups-and-organizations/).

{% hint style="info" %}
To enable the self-serve SSO option, contact your Snyk account team or [Snyk support](https://support.snyk.io).\
\
This option does not accommodate [custom role mapping](../custom-mapping/). To set up custom role mapping with SSO for your Snyk Group, contact your Snyk account team.
{% endhint %}

The following video demonstrates the process and steps for setting up single sign-on when using SAML.

{% embed url="https://thoughtindustries-1.wistia.com/medias/dyg9opxlv8" %}

## Use SAML for SSO: process overview

The process of establishing trust between your identity provider (IdP) and Snyk requires that the Group Admin do the following:

1. Configure your identity provider (IdP) by using the details about the Snyk environment displayed on-screen and user attributes.
2. Enter SAML attributes from your identity provider(IdP) on the Group SSO Settings page.
3. Configure Snyk SSO settings, choosing how you want your members to log in.
4. Verify SSO login to confirm the login process is working correctly.

{% hint style="info" %}
After SSO is configured both from Snyk and your company's network, a trust relationship is established with Snyk, Auth0 (on behalf of Snyk), and your network. Any sensitive data is encrypted and stored in Auth0 only for the purposes of enabling user logins.

Although not all the examples following this page cover verifying the Snyk signature, it is recommended that you improve the trust relationship and ensure integrity even further. Follow your respective IdP's documentation to add SP signature verification where possible.
{% endhint %}

## **User login**

Users are [provisioned ](../choose-a-provisioning-option.md)to Snyk when they log in. If the new user role selected is Group Member, the new user sees only a list of your Organizations until the admin adds them to the appropriate Organizations.

## High-level configuration steps

### **Configuring the IdP**

1. In the Snyk web UI, navigate to **Group** > **Settings.**
2. Select **SSO** and copy the needed information in step 1, namely **Entity ID**, **ACS URL,** and the Snyk **signing certificate URL**.
3. Enter these details in the IdP where appropriate and upload the Snyk signing certificate after downloading it locally in case the IdP does not accept only the certificate URL.
4. Before moving back to the Snyk UI, copy the IdP provided sign in URL and copy or download the IdP-provided **X509 signing certificate** details.

### **Configuring Snyk**

1. In Step 2 of the SSO settings page in the Snyk web UI, enter the details collected from the IdP by providing the sign in URL, sign out URL if available and desired, the IdP signing certificate and domains and subdomains that will be served over the SSO connection.
2. In case the connection requires HTTP-Redirect protocol binding, change that option from the default HTTP-POST.
3. Finally, verify if an IdP-initiated workflow should be enabled and then select **Create Connection** or **Save changes** if you are modifying an existing connectio&#x6E;**.**

### **Setting up user provisioning**

1. To make sure users are assigned the correct role when logging in for the first time, choose either **Group member**, **Org Collaborator** or **Org Admin**. Refer to [choosing a provisioning option](../choose-a-provisioning-option.md) for details on the options in this step.
2. In the section **Profile Attributes**, the fields are auto-populated but verify that **Email**, **Name** and **Username**, if known, exactly match the corresponding keys in the SAML payload raw JSON sent by the IdP to Snyk. Select **Save Changes** and move to the final step where you verify the setup.

### **Testing and verifying the configuration**

1. Provided all details have been entered correctly, the direct URL from the top of **Step 3** in the Snyk web UI can now be used to verify the configuration works as intended by logging in as a user in the directory of the IdP.
2. When all stored details are verified to be accurate (name, e-mail, permissions), Snyk generally recommends existing users that previously were logged in through Social login methods be removed from the Snyk platform. This can be accomplished under the Group menu **Members.**
