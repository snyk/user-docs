# Slack setup to connect Snyk with AWS Lambda

To enable Snyk to communicate with Slack, start by setting up incoming webhooks through Slack Apps. These are provided by Slack to enable developers to communicate with Slack.

To set up Slack Apps with a webhook, follow these steps:

1. Go to https://api.slack.com/apps
2.  Select **Create New App**.\


    <figure><img src="https://lh5.googleusercontent.com/qw51g6soQ6IjBf95JM0hhIsON0RqAhwuDbd7p3FA_AoGatQWx_0VcefI7RhEoUuKkuDNmXQNSIxw9aD7T7uhG4YPxvRIsAhDnHtVCGT_PtGuAD3fZDO4Qlye45iz94j7xZb0Ze0g8h16xMNtE-3zhmsw8wmq-m_K6OI1UD8mN-CKbNZEJCynuOHEBg" alt="Your Apps in Slack"><figcaption><p>Your Apps in Slack</p></figcaption></figure>
3.  Select **From scratch**.\


    <figure><img src="https://lh4.googleusercontent.com/uDE4iWxfnHF0KvGGZlwZwAp39zNvG1Vav8yOSxak5DhOIdOl983GS8Xmr-YJ9WpfiS6WJD2b5yhgbUAxMpm7rDwpQkEH2W2zOSyNQZdDAqDvBFpBMP7uYZwDtPGE3OGt0-g-JW09Dx2RB2wcfghEpc8J47A-DH7fejMkupKPnhrspesfPt45duXivg" alt="Create a Slack app from scratch"><figcaption><p>Create a Slack app from scratch</p></figcaption></figure>
4. Give the app a name like SnykToSlack.
5. Select your workspace.
6. With the Slack App created, click **Add features and functionality**.
7.  Select **Incoming Webhooks**.\


    <figure><img src="https://lh3.googleusercontent.com/yc2jyH0npATioGnzPLv5WEmI762OIYoefYVztKfvfAS9iV6yHNudbralS8VfLE0NT2x9TqM7lDCVLfV_27cC6Z82P5qprCIu4FKnVco1FfzsDJb3t6_V5BowDpBYw8GrNEaW8TZGbb1hmXsQflr1eeCTNAhKNpbE-AbUJGnxT65Uu67niA_HdCklQg" alt="Add features and functionality to Slack App, Incoming webhooks"><figcaption><p>Add features and functionality to Slack App, Incoming webhooks</p></figcaption></figure>
8.  Activate incoming webhooks in that page.

    <figure><img src="../../../.gitbook/assets/image (1) (4) (1).png" alt=""><figcaption><p>Incoming webhooks activation</p></figcaption></figure>
9. Generate a webhook URL for the channel you want by clicking on **Add New Webhook to Workspace**&#x20;
10. Select the channel you want Snyk to post to. If you haven’t already done so, [create a channel](https://slack.com/intl/en-gb/help/articles/201402297-Create-a-channel).
11. When the webhook has been created, copy and save the webhook URL to use in the next step in AWS.

<figure><img src="https://lh3.googleusercontent.com/av55N4Y2DyLFYmbrhC2gEjU9CINSP4DWUYfkhJju65Q9mpI-MqkKKsf5H8af2TMVy8f-jP6m-6Y-TAaaFsgf6dJ6LbtgGxfYM-vqAkUU5zYVYSoV8u8jKbFeBI9wWWpi9CFrSYPTM-ee2m7DJYDo1p4uBIf-IxqZSLpkJ4kQhp34lT6-6RQ9QLqIEQ" alt="Generate Slack incoming webhook URL"><figcaption><p>Generated webhook with Webhook URL</p></figcaption></figure>
