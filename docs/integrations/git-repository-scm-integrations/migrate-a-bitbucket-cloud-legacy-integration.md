# Migrate a Bitbucket Cloud Personal Access Token

{% hint style="info" %}
Snyk recommends using the Bitbucket Cloud App integration for smoother integration.
{% endhint %}

This topic describes how to migrate your existing [Bitbucket Cloud Personal Access Token (PAT) integration](bitbucket-cloud-legacy-integration.md) —displayed in Snyk as **Bitbucket Cloud (Legacy)**—to the [**Bitbucket Cloud App**](bitbucket-cloud-app-integration.md) integration.

## Deciding which Bitbucket Cloud integration to use

In general, the recommended way of integrating to Bitbucket Cloud is the new app integration. However, the new integration doesn't fit all cases. This section should give you an idea of which Bitbucket Cloud integration best fits for you.

### Main capabilities unlocked by the new app integration

* It allows using Snyk with Bitbucket’s [allowlisting IP addresses](https://support.atlassian.com/bitbucket-cloud/docs/control-access-to-your-private-content/) premium tier feature.
* It helps handle rate limiting issues for companies who spread their repos across multiple workspaces in Bitbucket Cloud.
* The new integration supports the first party interface in Bitbucket Cloud (Snyk's Security tab) out-of-the-box, meaning you wouldn't be required to install and maintain the first-party extension app in order to consume Snyk's security insights from Bitbucket Cloud.

### Limitations of the new app integration

* In the new app integration, every Snyk Org can connect only to one workspace in Bitbucket Cloud. If you rather like to import projects from various different workspaces in Bitbucket into the same single Org in Snyk, you should use the PAT integration.
* The new app integration does not yet support Snyk Multi Tenant EU, Snyk Multi Tenant AUS and Snyk Single Tenant cloud deployments.
* For customers who are part of the custom branch closed beta, the new app integration's first party interface in Bitbucket Cloud does not allow importing projects from a non-default branches. It is still possible though to import a non-default branch - you'd just need to do it from within Snyk.io import modal.

### Are there any plans to End-of-Life the Personal Access Token (PAT) integration?

No - the Personal Access Token Bitbucket Cloud integration is fully supported and there aren’t any plans to stop supporting it.

However, there is the 1st-party-interface extension app - an app that serves as an extension layer to the PAT integration that allows developers to consume Snyk’s findings from within the Bitbucket interface. This extension app was developed and supported by an external contractor company. As this functionality is now an integral part of the new app integration, the extension app has now moved to no-support mode, meaning that customers who use the PAT integration alongside the first-party extension app would need to migrate to the new app integration in order to get support for the first party interface functionality.

## Migration instructions

### Things you should know before migrating

In order to migrate to the new app integration, you would need to remove all the previously imported projects from Snyk, delete the legacy PAT integration and its projects, set up the new app integration and reimport your projects to Snyk from the new integration.

{% hint style="info" %}
Before going through the migration process, it's important to note that the following project-level information would not presist:

* Historic project related data, including trend numbers of fixing vulnerabilities
* Project related metadata (ignores/tags)
{% endhint %}

### Migration process

The migration process includes:

1. [Deleting the existing projects](migrate-a-bitbucket-cloud-legacy-integration.md#1.-delete-existing-projects) that are connected to the Bitbucket Cloud PAT (Legacy) integration in Snyk.
2. [Disconnecting the Legacy integration in Snyk](migrate-a-bitbucket-cloud-legacy-integration.md#2.-disconnect-the-legacy-integration).
3. [Removing the first-party extension](migrate-a-bitbucket-cloud-legacy-integration.md#3.-remove-the-snyk-tab-for-the-legacy-integration-in-bitbucket-cloud-optional) for the Legacy integration in Bitbucket (optional)
4. [Connecting the Bitbucket Cloud App ](migrate-a-bitbucket-cloud-legacy-integration.md#set-up-the-new-bitbucket-cloud-app-integration.)and importing projects.

#### 1. Delete existing projects

Delete all the existing projects in Snyk that were previously imported from the Legacy integration. To use the bulk delete action for your projects, go to <img src="../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> (Organization settings) > **Usage** and scroll to the **Projects** section.

<figure><img src="../../.gitbook/assets/migrate_bulk actions_bbc-11oct2022.png" alt=""><figcaption><p>Delete existing projects from the Snyk Bitbucket Cloud PAT (Legacy) integration</p></figcaption></figure>

#### 2. Disconnect the Legacy integration

To disconnect the Bitbucket Cloud (Legacy) integration, go to the settings page of Bitbucket Cloud (Legacy) integration, scroll to the relevant section and click **Disconnect.**

<figure><img src="../../.gitbook/assets/image (461).png" alt=""><figcaption><p>Disconnect the Bitbucket Cloud PAT (Legacy) integration</p></figcaption></figure>

#### 3. Remove the Snyk tab for the Legacy integration in Bitbucket Cloud (optional)

The Bitbucket Cloud (Legacy) has an optional first-party interface app for Bitbucket Cloud.

This app can potentially be installed on your Bitbucket Cloud workspace to enrich the legacy integration with a first-party interface (a "_**Snyk**_" tab).

If you used this app, before setting up the Snyk Bitbucket Cloud App, remove the Legacy interface app in Bitbucket Cloud before continuing to the next step: this functionality is supported out-of-the-box in the Snyk App integration.\
\
Go to your **Workspace settings** page in **Bitbucket.org > Manage installed apps**, expand the **Snyk Security for Bitbucket Cloud** app, and click **Remove.**

<figure><img src="../../.gitbook/assets/remove_snyk-security-bbc_11oct2022.png" alt=""><figcaption><p>Remove the first-party Snyk Legacy interface app in Bitbucket.org</p></figcaption></figure>

#### 4. Set up the Bitbucket Cloud App integration

See the [Bitbucket Cloud App integration](bitbucket-cloud-app-integration.md) topic for more detailed information.

### Migration demo

#### **Watch this Snyk Bitbucket Cloud App migration demo to see how it's done!**

Marco Morales, a Partner Solutions Architect at Snyk, talks about the Snyk Bitbucket Cloud App and goes through the process to migrate an existing Legacy integration to the Snyk Bitbucket Cloud App - in under 5 minutes!

_Go to timestamp 2:34 to jump right into the demo:_

{% embed url="https://thoughtindustries-1.wistia.com/medias/32rgw3hkdk" %}
How to migrate to the new Snyk Bitbucket Cloud App integration
{% endembed %}
