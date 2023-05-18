# Slack App

{% hint style="warning" %}
**Feature availability**\
This integration is available to customers as a closed beta. Contact your Snyk team if you would like access.
{% endhint %}

Snyk's new Slack App integration is built as an official [Slack App](https://api.slack.com/start/overview#apps), unlike the current Slack integration which uses Slack’s [Incoming Webhooks](https://api.slack.com/legacy/custom-integrations/messaging/webhooks); this is considered outdated.

When the Slack App integration is released, Snyk encourages all customers using the Slack Integration to adopt the Slack App integration and get the following benefits and more:

* Improved Support for Snyk Products (Open Source, Code, Container, and Iac)
* Support for filtering on severity levels

## Overview of Slack App integration

The Slack App integration lets you receive new issue notifications from your Snyk Projects in your Slack workspace. It is based on Snyk Apps that give you the benefit of more granular scopes and more ability to configure notifications at the Organization and Project level.

To enable the Slack integration, you must do the following:

1. Authorize the Slack App with Snyk to get new issues data that can be forwarded to your Slack workspace.
2. Authorize the Slack App with your Slack workspace to allow Slack App to send notifications to your channels in the workspace.
3. Configure the default notification settings for all Projects in your Organization and add any Project-level notification overrides if you would like

## Configure the integration in Snyk

Open the [Snyk integrations page](https://app.snyk.io/integrations), navigate to Notifications, and click the new Slack App tile:

<figure><img src="../../.gitbook/assets/slack-app1.png" alt="Snyk integrations Slack App tile"><figcaption><p>Snyk integrations Slack App tile</p></figcaption></figure>

This launches the authorization flow, allowing Snyk access to your Slack App:

<figure><img src="../../.gitbook/assets/slack-app2.png" alt="Add S;acl a[[ tp Snyk"><figcaption><p>Add Slack to Snyk</p></figcaption></figure>

After Snyk has been authorized, you will be asked to authorize Slack to connect to Snyk for Slack:

<figure><img src="../../.gitbook/assets/slack-app3.png" alt="Authorize SLack to connect to Snyk"><figcaption><p>Authorize Slack to connect to Snyk</p></figcaption></figure>

After this step is complete, you can configure the integration to choose which Slack channel the Organization notifications are sent to and also filter by severity level:

<figure><img src="../../.gitbook/assets/slack-app4.png" alt="Choose Slack channel and severity level for the notifications from the Organization"><figcaption><p>Choose Slack channel and severity level for the notifications from the Organization</p></figcaption></figure>

## Remove the Slack integration

To remove the Slack App integration, navigate to the Slack App settings page, locate **Remove Slack Snyk app** at the bottom of the page, and click the **Disconnect Slack** button:

<figure><img src="../../.gitbook/assets/slack-app5.png" alt="Remove Slack App integration"><figcaption><p>Remove Slack App integration</p></figcaption></figure>
