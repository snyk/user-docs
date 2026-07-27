---
description: How to integrate Snyk API and Web with Slack
nav_context: classic
---

# Integrate with Slack

By connecting Snyk API & Web with Slack, you receive notifications about the activity of your targets in your Slack channels. For example, when target scans start or finish, when logins fail, or when Snyk finds or fixes vulnerabilities.

This integration involves two steps:

1. Configure the webhook in Slack.
2. Configure the Slack webhook in Snyk API & Web.

This article describes these steps in detail.

## Configure the webhook in Slack

Configure a webhook in Slack with a channel to receive notifications from Snyk:

1. Log in to your Slack account at `https://slack.com/signin` and select your Slack workspace.
2. Navigate to `https://api.slack.com/apps` and do the following:
   1. If this is your first app, click **Create an App**. Otherwise, click **Create New App**.
   2. In the displayed dialog, click **From Scratch**.
3. In the next dialog, fill out the form as follows:
   1. In the **App name**, enter a meaningful name. For example: "Snyk API & Web Integration".
   2. In the dropdown, select the Slack workspace you want to receive Snyk notifications.
   3. Click **Create App**. Slack redirects you to the **Basic Information** page.
4. In the sidebar menu, select **Incoming Webhooks**. If it is not turned on, click the toggle button of **Activate Incoming Webhooks** to turn it on.
5. Scroll to the bottom of the page and click **Add New Webhook to Workspace**.
6. On the displayed page, click the dropdown to see the list of Slack channels, select the one you want to receive Snyk notifications (for example, "Snyk API & Web notifications"), and click **Allow**.
7. Return to **Incoming Webhooks**, scroll down, and validate the newly created webhook.
8. Click **Copy** to store the webhook URL in the clipboard. You need it to configure the Slack integration in Snyk.

You have configured the webhook on the Slack side. Now move on to the Snyk side.

## Configure the Slack webhook in Snyk API & Web

Configure the Slack webhook on the target for which you want to receive notifications about its activity:

1. In the Snyk API & Web application, navigate to **Targets** and click the **gear icon** of the row of the target you want to configure.
2. Click the **Integrations** tab, navigate to the **Slack** section, and set the **Slack webhook URL**. If you copied the URL to the clipboard at the end of the previous step, paste it.
3. Select which events you want to be notified about. For example, target scan events such as when scans start and end, when they identify high-severity vulnerabilities, and when they fail to log in.
4. Click **Save** to conclude the Slack integration for the target.

After you integrate a target with Slack, you start receiving notifications from Snyk about the target's activity in your Slack channel. For example:

This integration lets you adjust notifications to the needs of your company. For example, you can use a single Slack webhook to configure your targets and receive all notifications in the same Slack channel. If necessary, you can create more webhooks and configure your targets so that you receive notifications of different targets, or groups of targets, in different Slack channels.
