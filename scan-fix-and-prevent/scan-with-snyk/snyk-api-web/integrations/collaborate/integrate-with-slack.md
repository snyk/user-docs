# Integrate with Slack

By connecting Snyk API & Web with Slack, you receive notifications about the activity of your targets in your Slack channels. For example, when target scans start or finish, when logins fail, or when vulnerabilities are found or fixed.

This integration involves two steps:

1. Configure the webhook in Slack.
2. Configure the Slack webhook in Snyk API & Web.

This article describes these steps in detail.

## Configure the webhook in Slack

The first step is to configure a webhook in Slack with a channel to receive notifications from Snyk API & Web:

1. Sign in to your Slack account at `https://slack.com/signin` and select your Slack workspace.
2. Go to `https://api.slack.com/apps` and do the following:
   1. If this is your first app, click **Create an App**. Otherwise, click **Create New App**.
   2. In the displayed dialog, click **From Scratch**.
3. In the next dialog, fill out the form as follows:
   1. In the **App name**, type in a meaningful name. For example: "Snyk API & Web Integration".
   2. In the dropdown, pick the Slack workspace you want to receive Snyk API & Web notifications.
   3. Click **Create App**, and you are redirected to the **Basic Information** page.
4. On the sidebar menu, select **Incoming Webhooks**, and if not turned on, click the toggle button of **Activate Incoming Webhooks** to turn it on.
5. Scroll to the bottom of the page and click **Add New Webhook to Workspace**.
6. On the displayed page, click the dropdown to see the list of Slack channels, select the one you want to receive Snyk API & Web notifications to (for example, "Snyk API & Web notifications"), and click **Allow**.
7. Back to **Incoming Webhooks**, scroll down and validate the newly created webhook.
8. Click **Copy** to store the webhook URL in the clipboard. You need it to configure the Slack integration in Snyk API & Web.

The webhook has been configured on the Slack side. Now move on to the Snyk API & Web side.

## Configure the Slack webhook in Snyk API & Web

In this step, configure the Slack webhook on the target for which you want to receive notifications about its activity:

1. In the Snyk API & Web application, go to **Targets** and click the **gear icon** of the row of the target you want to configure.
2. Click the **Integrations** tab, go to the **Slack** section, and set the **Slack webhook URL**. If you copied the URL to the clipboard at the end of the previous step, paste it.
3. Now choose which events you want to be notified about. For example, events of target scans such as when they start and end, when they identify high-severity vulnerabilities, and if they fail to log in.
4. Click **Save** to conclude the Slack integration for the target.

Once a target is integrated with Slack, you start receiving notifications from Snyk API & Web about the target's activity in your Slack channel. For example:

This integration allows you to adjust notifications to your organization's needs. For example, you can have a single Slack webhook to configure your targets, and you receive all notifications in the same Slack channel. However, if necessary, you can create more webhooks and configure your targets accordingly so that you receive notifications of different targets (or groups of targets) in different Slack channels.
