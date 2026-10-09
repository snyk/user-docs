---
nav_context: classic
description: How to set up a OneLogin SAML application and connect it to Snyk for SSO
---

# OneLogin SAML Application setup

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

## OneLogin SAML Application setup

This example shows setting up a SAML application in OneLogin and connecting this to Snyk to facilitate SSO. To configure your OneLogin to use SSO with Snyk, start by obtaining an entity ID and a reply URL (Assertion Consumer Service URL) from Snyk.

1\. From the dropdown at the top left, select **GROUP OVERVIEW** and then the cog icon (top right corner) to get to your group settings.

2\. Click on **SSO** and copy the values under **Entity ID** and **ACS URL** or leave the browser tab open for easy access.

3\. Navigate to [OneLogin](https://www.onelogin.com/), open the **applications** menu, click on **Applications**, and then on the **Add App** button on the top right.

4\. Filter by “saml” and select the **SAML Custom Connector (Advanced)** app.

5\. Enter a **Display Name** for the app, for example, Snyk. Optionally, pick an icon. You can find the Snyk logo on Snyk’s [press kit](https://snyk.io/press-kit/) page.

6\. After the app is saved, go to the configuration page and enter the **Audience** (EntityID), **ACS URL,** and ACS URL Validator. Use the ACS URL value here as well. Copy the EntityID and ACS URL from step 2.

7\. Go to the parameters section and add a new parameter, as Snyk needs the email address as the `email` attribute in the SAML response. Click the **+** icon.

For the **Field name**, use **email**. Make sure the checkbox **Include in SAML assertion** is checked. **Save**.

On the next screen, select **Email** in the **Value** dropdown list. **Save**.

Repeat the same steps for the **name** attribute.

8\. Go to the **SSO** section, download the **certificate,** and make a note of the **SAML 2.0 Endpoint (HTTP)** value.

9\. To enable the app for users, go to the **Users** > **Roles** section and **create a new role** called Snyk.\
Be sure to select the Snyk SAML app as an enabled role app. **Save**.\
Then, assign this role to all the users who should be able to use Snyk.

10\. Go back to the Snyk portal, scroll to step 2, and make a note of entering the details from step 8, including the domain(s) you wish to use over the SSO connection.\
Verify if an **IdP-initiated workflow** should be enabled and then click **Create Auth0 connection**.

11\. Scroll to step 3 and determine how new users should be treated when signing in.\
Choose the option you would like to use: **Group member, Org collaborator,** or **Org admin**.\
Finally, enter the **Profile attributes** as you configured them in OneLogin (email, name, nickname), click **Save changes,** and verify you can log in, either with the direct URL at the top of step 3 or by going to the[ generic SSO login](https://app.snyk.io/login/sso).
