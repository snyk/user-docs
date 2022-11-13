# Migrate a Bitbucket Cloud Legacy integration

{% hint style="info" %}
Snyk recommends using the new Bitbucket Cloud App integration for smoother integration and to ensure long-term support.
{% endhint %}

This topic describes how to migrate your existing [Bitbucket Cloud Personal Access Token (PAT) integration](bitbucket-cloud-integration.md) —displayed in Snyk as **Bitbucket Cloud (Legacy)**—to the new **** [**Bitbucket Cloud App**](bitbucket-cloud-app-integration.md) **** integration.

#### **Watch this Snyk Bitbucket Cloud App migration demo to see how it's done!**&#x20;

Marco Morales, a Partner Solutions Architect at Snyk, talks about the Snyk Bitbucket Cloud App and goes through the process to migrate an existing Legacy integration to the new Snyk Bitbucket Cloud App - in under 5 minutes! &#x20;

_Go to timestamp 2:34 to jump right into the demo:_ &#x20;

{% embed url="https://thoughtindustries-1.wistia.com/medias/32rgw3hkdk" %}
How to migrate to the new Snyk Bitbucket Cloud App integration
{% endembed %}

## Migration Process

The migration process includes:&#x20;

1. [Deleting the existing projects](migrate-a-bitbucket-cloud-legacy-integration.md#1.-delete-existing-projects) that are connected to the Bitbucket Cloud PAT (Legacy) integration in Snyk.
2. [Disconnecting the Legacy integration in Snyk](migrate-a-bitbucket-cloud-legacy-integration.md#2.-disconnect-the-legacy-integration).
3. [Removing the first-party extension](migrate-a-bitbucket-cloud-legacy-integration.md#3.-remove-the-snyk-tab-for-the-legacy-integration-in-bitbucket-cloud-optional) for the Legacy integration in Bitbucket (optional)
4. [Connecting the Bitbucket Cloud App ](migrate-a-bitbucket-cloud-legacy-integration.md#set-up-the-new-bitbucket-cloud-app-integration.)and importing projects.

### 1. Delete existing projects

Delete all the existing projects in Snyk that were previously imported from the Legacy integration. To use the bulk delete action for your projects, go to <img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line">  (Organization settings) > **Usage** and scroll to the **Projects** section.

<figure><img src="../../.gitbook/assets/migrate_bulk actions_bbc-11oct2022.png" alt=""><figcaption><p>Delete existing projects from the Snyk Bitbucket Cloud PAT (Legacy) integration</p></figcaption></figure>

### 2. Disconnect the Legacy integration

To disconnect the Bitbucket Cloud (Legacy) integration, go to the settings page of Bitbucket Cloud (Legacy) integration, scroll to the relevant section and click **Disconnect.**

<figure><img src="../../.gitbook/assets/image (3) (1).png" alt=""><figcaption><p>Disconnect the Bitbucket Cloud PAT (Legacy) integration</p></figcaption></figure>

### 3.  Remove the Snyk tab for the Legacy integration in Bitbucket Cloud (optional)&#x20;

The Bitbucket Cloud (Legacy) has an optional first-party interface app for Bitbucket Cloud.&#x20;

This app can potentially be installed on your Bitbucket Cloud workspace to enrich the legacy integration with a first-party interface (a "_**Snyk**_" tab).&#x20;

If you used this app, before setting up the Snyk Bitbucket Cloud App, remove the Legacy interface app in Bitbucket Cloud before continuing to the next step: this functionality is supported out-of-the-box in the Snyk App integration.\
\
Go to your **Workspace settings** page in **Bitbucket.org > Manage installed apps**, expand the **Snyk Security for Bitbucket Cloud** app, and click **Remove.**

<figure><img src="../../.gitbook/assets/remove_snyk-security-bbc_11oct2022.png" alt=""><figcaption><p>Remove the first-party Snyk Legacy interface app in Bitbucket.org</p></figcaption></figure>

### 4. Set up the Bitbucket Cloud App integration&#x20;

See the [Bitbucket Cloud App integration](bitbucket-cloud-app-integration.md) topic for more detailed information.
