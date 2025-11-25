# Google Workspace setup

This example shows setting up an Google Workspace SAML application and connecting it to Snyk to facilitate SSO.

For details in addition to the information provided on this page, see [Set up your own custom SAML app](https://support.google.com/a/answer/6087519).

Start by logging into the Google Workspace [admin area](https://admin.google.com).

1.  Go to **Apps** and then click **Web and mobile apps**.

    <figure><img src="../../../../.gitbook/assets/1 (2) (1).png" alt="Open Web and Mobile apps"><figcaption><p>Open Web and Mobile apps</p></figcaption></figure>
2.  Click on **Add app** and choose **Add custom SAML app**.

    <figure><img src="../../../../.gitbook/assets/2 (3).png" alt="Add new custom SAML app"><figcaption><p>Add new custom SAML app</p></figcaption></figure>
3.  Name your application appropriately and click **Continue**.

    <figure><img src="../../../../.gitbook/assets/3 (3).png" alt="Name the SAML app"><figcaption><p>Name the SAML app</p></figcaption></figure>
4.  Download the certificate and open it in your preferred text editor.

    <figure><img src="broken-reference" alt="Download signing certificate"><figcaption><p>Download signing certificate</p></figcaption></figure>
5.  Navigate to the Snyk portal, login and from the drop down at the top left select **GROUP OVERVIEW** and then the cog wheel (top right corner) to get to your group settings.

    <figure><img src="broken-reference" alt="Open group view in Snyk"><figcaption><p>Open group view in Snyk</p></figcaption></figure>
6.  Click on **SSO**, scroll down to step 2, and paste the Google SSO URL from step 4 into **Sign in URL** and the certificate in your text editor into **X509 signing certificate**.\
    Add the domain name(s) you are configuring this connection for in **Email domains and subdomains that need SSO access**.\
    Verify if an **IdP-initiated workflow** should be enabled and then save your modifications

    <figure><img src="broken-reference" alt="Enter details from Google Workspace"><figcaption><p>Enter details from Google Workspace</p></figcaption></figure>
7.  Scroll up to step 1 and copy the **Entity ID** and **ACS URL**.

    <figure><img src="../../../../.gitbook/assets/7.png" alt="Copy details from Snyk"><figcaption><p>Copy details from Snyk</p></figcaption></figure>
8.  Go back to the Google admin portal , click **Continue,** and paste those two values into their respective fields. Then tick **Signed response**.

    <figure><img src="../../../../.gitbook/assets/8.png" alt="Enter details from Snyk in Google"><figcaption><p>Enter details from Snyk in Google</p></figcaption></figure>
9.  Click **Continue**, add an app attribute named email tied to the **Primary Email**, and save the configuration.

    <figure><img src="broken-reference" alt="Add email attribute"><figcaption><p>Add email attribute</p></figcaption></figure>
10. Enable access to your app for your users by clicking **User Access**, tick **On for everyone**, and **Save**. Modify organizational access as needed.

    <figure><img src="../../../../.gitbook/assets/10 (1).png" alt="Enable SSO app for the organization"><figcaption><p>Enable SSO app for the organization</p></figcaption></figure>
11. Finalize the setup by going back to the Snyk portal and decide how new users should be processed when signing in. Choose the option you would like to use: **Group member**, **Org collaborator**, or **Org admin**.
12. Then add the previously created **email** app attribute to both **Email** and **Username** and save the configuration. If you wish to populate the full name you may configure a custom attribute in Google Workspace.

    <figure><img src="../../../../.gitbook/assets/11.png" alt="Tie together attributes from Google to Snyk"><figcaption><p>Tie together attributes from Google to Snyk</p></figcaption></figure>
