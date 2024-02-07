# Slack app

The Snyk app for Slack is built as an official [Slack App](https://api.slack.com/start/overview#apps).

{% hint style="warning" %}
Snyk recommends that all customers use the Snyk Slack app, as the [Slack integration](slack-integration.md) is outdated.
{% endhint %}

The Slack app provides the following benefits and more:

* Improved support for Snyk products: Open Source, Code, Container, and IaC
* Support for filtering on severity levels
* Support for overriding notifications at the Project level

{% hint style="info" %}
Vulnerabilities detected on initial import of Projects are not sent to Slack. Only newly disclosed vulnerabilities from imported Projects are sent to Slack.&#x20;
{% endhint %}

## Overview of the Snyk app for Slack

The Snyk app for Slack highlights vulnerability information across your software Projects and presents actionable context within your Slack workspace channels. Your teams get the issue notifications that matter the most and can act on them without leaving Slack.

This has the advantage of:

* Enhancing collaboration to discuss and resolve vulnerabilities quickly
* Maintaining awareness of relevant and emerging vulnerabilities
* Minimizing disruptions with issue notification thresholds

The Snyk app for Slack is based on [Snyk Apps](../../snyk-api-info/snyk-apps/), providing you with the benefit of more granular scopes and more ability to configure issue notifications at the Organization and Project levels.

## Steps to enable the Snyk app for Slack

To enable the Snyk app for Slack, you must do the following:

1. Authorize the app with Snyk to get new issues data that can be forwarded to your Slack workspace.
2. Authorize the app with your Slack workspace to allow Snyk to send notifications to your channels in the workspace.
3. Configure the default notification settings in Snyk for all Projects in your Organization and add [Project-level notification overrides](slack-app.md#manage-project-level-notification-overrides) if you would like.

## Configure the Snyk app for Slack

Ensure the user performing this installation has the permission **Snyk Apps Management - Install Apps** before continuing. See [documentation for member roles](../../snyk-admin/manage-permissions-and-roles/manage-member-roles.md).

Open the [Snyk integrations page](https://app.snyk.io/integrations), navigate to **Notifications**, and click the **Slack App** tile:

<figure><img src="../../.gitbook/assets/slack-app1.png" alt="Snyk integrations Slack App tile"><figcaption><p>Snyk integrations Slack App tile</p></figcaption></figure>

This launches the authorization flow, allowing Snyk access to your Snyk app for Slack:

<figure><img src="../../.gitbook/assets/slack-app2.png" alt="Add Slack to Snyk"><figcaption><p>Add Slack to Snyk</p></figcaption></figure>

After Snyk has been authorized, you will be asked to authorize Slack to connect to Snyk for Slack. Click **Allow.**

<figure><img src="../../.gitbook/assets/slack-app3.png" alt="Authorize Slack to connect to Snyk"><figcaption><p>Authorize Slack to connect to Snyk</p></figcaption></figure>

{% hint style="info" %}
If multiple Slack workspaces are available, a dropdown will be visible at the top right of the page. Select the desired Slack workspace.
{% endhint %}

After this step is complete, you can configure the integration to provide a Slack channel ID for the channel where issue notifications for the Organization are sent, and also filter by severity level.

<figure><img src="../../.gitbook/assets/slack-app-settings.png" alt="Choose Slack channel and severity level for the notifications for the Organization"><figcaption><p>Choose Slack channel and severity level for the notifications for the Organization</p></figcaption></figure>

{% hint style="info" %}
To find the channel ID of a Slack Channel, open Slack, right-click on the channel name, select **View channel details**, then scroll down to the bottom of the window where you will see the channel ID, for example, C2TB2222N.
{% endhint %}

{% hint style="info" %}
To add the Snyk for Slack app to a private channel, you must first add the app manually to the channel from within Slack and then select the channel within the Snyk integration.

In the Private channel, select **Channel settings - Integrations**, and then **Add an app**. Search for **Snyk for Slack** and select **add**. \
After you have done this, the channel is displayed on the **Settings** page for the integration.
{% endhint %}

{% hint style="warning" %}
If you are a Slack Admin, you can manually add the Snyk for Slack app to a private channel. To do this, type **@Snyk for Slack** in the chosen channel to summon the setup bot.
{% endhint %}

## Issue notifications

After the Slack app has been configured, new issue notifications are forwarded to the selected Slack channel according to the desired severity level threshold. \
New issue notifications may take up to an hour to start propagating to your Slack workspace after it is configured.&#x20;

<figure><img src="../../.gitbook/assets/image (113).png" alt="Example of a new critical vulnerability notification received in Slack"><figcaption><p>Example of a new critical vulnerability notification received in Slack</p></figcaption></figure>

## Manage Project-level notification overrides

To override Slack notification settings for a Project, use the customizations table.

<figure><img src="../../.gitbook/assets/image (118) (3).png" alt="Empty Project notification customizations table"><figcaption><p>Empty Project notification customizations table</p></figcaption></figure>

### Find your Project ID

To set a Project level override, find the ID of the Project for which you want to send notifications.

1. Navigate to the Projects page using the application sidebar.
2. Select the desired Project.
3. Copy the UUID value that follows the `project/` path in the URL bar.

<figure><img src="../../.gitbook/assets/image (315).png" alt="Location of the Project ID in the Project page URL"><figcaption><p>Location of the Project ID in the Project page URL</p></figcaption></figure>

### Add notification customizations

1. Select `Add customization` to open the creation dialog from the Slack App integration page.&#x20;
2. Paste your previously copied Project ID value.&#x20;
3. Select the desired issue severity level and the Slack channel to target.&#x20;
4. Save the customization.

<div data-full-width="true">

<figure><img src="../../.gitbook/assets/slack-app-project-override.png" alt="Project notification customization creation dialog" width="563"><figcaption><p>Project notification customization creation dialog</p></figcaption></figure>

</div>

### Edit or delete notification customizations

After a notification customization has been created, you can edit its configuration or delete it at any time. To do so, select the action in the ellipsis menu for the customization table entry.

<figure><img src="../../.gitbook/assets/image (302).png" alt="Edit and Delete options for notification customizations" width="443"><figcaption><p>Edit and Delete options for notification customizations</p></figcaption></figure>

## Manage notification customizations using the REST API

To override Slack notification settings on a per-Project basis, you can use a set of [Snyk REST API](https://apidocs.snyk.io/) endpoints.

Before attempting to use these endpoints, ensure that you have retrieved your authentication token and Organization ID (`org_id`) as outlined in the steps of the [guide for getting started using Snyk REST API.](../../snyk-api/try-a-simple-call-to-the-snyk-rest-api.md)

### P**rerequisites** for managing Project-level notification overrides

#### Find the Slack App Bot ID

To interact with the Project level notification customization endpoints, you must have a `bot_id`. You can obtain it with a request to the [Get a list of app bots](https://apidocs.snyk.io/?version=2023-08-04#get-/orgs/-org\_id-/app\_bots) endpoint.

GET [/orgs/{org\_id}/app\_bots](https://apidocs.snyk.io/?version=2023-08-04#get-/orgs/-org\_id-/app\_bots)

{% hint style="info" %}
Ensure you apply the `expand=app` query string on your request. This enables you to find the Bot for the related Snyk App named **Slack App**.
{% endhint %}

#### Find your Project ID

To target the desired Project, you must have its `project_id`. This can be obtained with a GET request to the [List all Projects for an Org](https://apidocs.snyk.io/?version=2023-08-04#get-/orgs/-org\_id-/projects) endpoint:

GET [/orgs/{org\_id}/projects](https://apidocs.snyk.io/?version=2023-08-04#get-/orgs/-org\_id-/projects)

### Relevant REST APIs

After retrieving the `org_id`, `bot_id`, and `project_id` values, you can use the following create, read, update, and delete API operations:

#### List all  Slack notification customizations for Projects

GET [/orgs/{org\_id}/slack\_app/{bot\_id}/projects](https://apidocs.snyk.io/?version=2023-08-04#get-/orgs/-org\_id-/slack\_app/-bot\_id-/projects)

#### Create a Slack notification customization for a Project

POST [/orgs/{org\_id}/slack\_app/{bot\_id}/projects/{project\_id}](https://apidocs.snyk.io/?version=2023-08-04#post-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

#### Update a Slack notification customization for a Project

PATCH [/orgs/{org\_id}/slack\_app/{bot\_id}/projects/{project\_id}](https://apidocs.snyk.io/?version=2023-08-04#patch-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

#### Delete a Slack notification customization for a Project

DELETE [/orgs/{org\_id}/slack\_app/{bot\_id}/projects/{project\_id}](https://apidocs.snyk.io/?version=2023-08-04#delete-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

## Remove the Snyk app for Slack

To remove the Snyk app for Slack, navigate to the settings page, locate **Remove Slack Snyk app** at the bottom of the page, and click the **Disconnect Slack** button:

<figure><img src="../../.gitbook/assets/slack-app5.png" alt="Remove Slack App integration"><figcaption><p>Remove Slack App integration</p></figcaption></figure>
