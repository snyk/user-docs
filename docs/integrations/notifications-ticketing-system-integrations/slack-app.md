# Slack app

Snyk's new Snyk app for Slack is built as an official [Slack App](https://api.slack.com/start/overview#apps), unlike the current Slack integration which uses Slack’s [Incoming Webhooks](https://api.slack.com/legacy/custom-integrations/messaging/webhooks); this is considered outdated.

When the Snyk app for Slack is released, Snyk encourages all customers using the current Slack integration to adopt the Snyk app for Slack and get the following benefits and more:

* Improved Support for Snyk Products (Open Source, Code, Container, and IaC)
* Support for filtering on severity levels

{% hint style="info" %}
Vulnerabilities detected on initial import of projects are not sent to Slack immediately.
{% endhint %}

## Overview of the Snyk app for Slack

The Snyk app for Slack lets you receive new issue notifications from your Snyk Projects in your Slack workspace. It is based on [Snyk Apps](../../snyk-api-info/snyk-apps/) that give you the benefit of more granular scopes and more ability to configure notifications at the Organization and Project level.

To enable the Snyk app for Slack, you must do the following:

1. Authorize the Snyk app for Slack with Snyk to get new issues data that can be forwarded to your Slack workspace.
2. Authorize the Snyk app for Slack with your Slack workspace to allow us to send notifications to your channels in the workspace.
3. Configure the default notification settings for all Projects in your Organization and add any Project-level notification overrides if you would like

## Configure the Snyk app for Slack

Open the [Snyk integrations page](https://app.snyk.io/integrations), navigate to Notifications, and click the new Slack App tile:

<figure><img src="../../.gitbook/assets/slack-app1.png" alt="Snyk integrations Slack App tile"><figcaption><p>Snyk integrations Slack App tile</p></figcaption></figure>

This launches the authorization flow, allowing Snyk access to your Snyk app for Slack:

<figure><img src="../../.gitbook/assets/slack-app2.png" alt="Add S;acl a[[ tp Snyk"><figcaption><p>Add Slack to Snyk</p></figcaption></figure>

After Snyk has been authorized, you will be asked to authorize Slack to connect to Snyk for Slack:

<figure><img src="../../.gitbook/assets/slack-app3.png" alt="Authorize SLack to connect to Snyk"><figcaption><p>Authorize Slack to connect to Snyk</p></figcaption></figure>

After this step is complete, you can configure the integration to choose which Slack channel the Organization notifications are sent to and also filter by severity level:

<figure><img src="../../.gitbook/assets/slack-app4.png" alt="Choose Slack channel and severity level for the notifications from the Organization"><figcaption><p>Choose Slack channel and severity level for the notifications from the Organization</p></figcaption></figure>

{% hint style="info" %}
To add the Snyk for Slack app to a private channel, you must first add the app manually to the channel from within Slack and then select the channel within the Snyk integration.

In the Private channel, select **Channel settings - Integrations**, and then **Add an app**. Search for **Snyk for Slack** and select **add**. \
After you have done this, the channel is displayed on the Settings page for the integration.
{% endhint %}

## Remove the Snyk app for Slack

To remove the Snyk app for Slack, navigate to the settings page, locate **Remove Slack Snyk app** at the bottom of the page, and click the **Disconnect Slack** button:

<figure><img src="../../.gitbook/assets/slack-app5.png" alt="Remove Slack App integration"><figcaption><p>Remove Slack App integration</p></figcaption></figure>
