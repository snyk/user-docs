# Slack integration

{% hint style="info" %}
Snyk recommends that all customers use the [Slack App](slack-app.md), as the Slack integration is outdated.
{% endhint %}

You can set up Slack to receive Snyk’s alerts about new vulnerabilities that affect your Projects and new upgrades or patches that have become available.

{% hint style="info" %}
Vulnerabilities detected on initial import of projects are not sent to Slack immediately.
{% endhint %}

You will get the following alerts in Slack.

A newly disclosed vulnerability affects you:

<figure><img src="../../.gitbook/assets/image (23) (1) (1).png" alt="Newly disclosed vulnerability notification"><figcaption><p>Newly disclosed vulnerability notification</p></figcaption></figure>

A new upgrade or patch is available for a vulnerability that you previously ignored or patched:

<figure><img src="../../.gitbook/assets/image (22) (1).png" alt="New upgrade available"><figcaption><p>New upgrade available</p></figcaption></figure>

To set the integration up, you must generate a Slack webhook. You can do this either via the [Incoming WebHooks app](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) or by [creating your own Slack app](https://api.slack.com/incoming-webhooks).

Once you have generated your Slack Webhook URL, go to your **Manage Organization** settings and enter the URL.

<figure><img src="../../.gitbook/assets/image (24) (1) (1) (1) (1) (1) (1).png" alt="Enter URL of the Slack webhook"><figcaption><p>Enter URL of the Slack webhook</p></figcaption></figure>

{% hint style="info" %}
Currently, only webhooks created with Slack Apps are supported; webhooks created with Slack Workflows are not currently supported.
{% endhint %}
