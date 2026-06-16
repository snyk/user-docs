# Scan an Akamai connection for asset discovery

APIs often operate without the direct oversight of security teams, making them difficult to track and protect. To run a security scan on an API, you first need its specification file (schema), but finding the correct, up-to-date schema for every API in your organization can be a significant challenge.

This guide shows you how to connect your Akamai account to Snyk API & Web. This integration automatically discovers the APIs managed in your Akamai account and imports their schemas, allowing you to add them as targets in Snyk API & Web with a single click.

## Prerequisites

Before you start, you need administrator access to your organization's Akamai API Security portal to retrieve the required API credentials.

## Get your Akamai API credentials

To allow Snyk API & Web to access your assets, you need to obtain API credentials from your Akamai API Security tenant.

1. Log in to the Akamai API Security portal.
1. From your account's main navigation menu, select **Settings**.
1. From the **Settings** menu, select **User Management**.
1. Select the **Service Accounts** tab.
1. Click **Create Service Account** to open the configuration form.
1. Enter a name and a duration, then click **Save**.
1. Copy the following two values displayed in the **Service Account Credentials** window and save them in a secure location:
   * `Client ID`
   * `Client Secret`

{% hint style="warning" %}
The `Client Secret` is shown only once. Ensure you copy it before leaving the page. You will need these credentials for the next step.
{% endhint %}

## Add the Akamai connection in Snyk API & Web

Use the credentials you created to connect your Akamai account in Snyk API & Web. You can do this from either the global **Integrations** page or the **Discovery** page.

### From the Integrations page

1. In your Snyk API & Web account, navigate to **Settings** > **Integrations**.
1. Locate the **Akamai** integration module.
1. Enter your credentials obtained from Akamai:
   * Akamai Server URL
   * Akamai Client ID
   * Akamai Client Secret
1. Click **Save**.

### From the Discovery page

1. In your Snyk API & Web account, navigate to the **Discovery** page.
1. Click **Add source**.
1. From the list of sources, select **Akamai**.
1. Enter your credentials obtained from Akamai:
   * Akamai Server URL
   * Akamai Client ID
   * Akamai Client Secret
1. Click **Connect**.

## View your discovered assets

After a successful connection, Snyk API & Web immediately begins importing your domains from Akamai. This process runs periodically to keep your inventory up to date.

You can view these domains by navigating to **Targets** > **Domains**. Once your domains are imported, Snyk API & Web automatically runs a discovery scan on them to find the associated APIs. After this scan completes, your discovered assets are listed under the **Discovery** page. For each API, a **{...}** icon is displayed to indicate that its schema is available.

## Add a discovered API as a target

Once your API assets appear in the **Discovery** list, you can add them as targets.

1. Navigate to the **Discovery** page.
1. To find API assets, use the following filters:
   * Filter by **Type** > **API** to display only API assets.
   * Filter by **Source** > **Akamai** to display assets imported from this integration.
1. Click **Add Target** on the corresponding row.

Snyk API & Web automatically configures the API as a new target with its schema pre-populated, making it ready for security testing. Once you add the asset as a target, click **Scan** to get started.
