# Notifying the Team

## Set email notifications

Snyk can automatically notify you when new issues are found in imported Projects. By default, email notifications are sent for any newly found vulnerabilities, but can customize these; for example, to only send messages for new High/Critical severity issues.&#x20;

#### Weekly updates

Snyk sends you a weekly update summarizing your security status across all of your Organizations.&#x20;

Also, Snyk generates usage alert emails warning when you are approaching usage limits, if you rely on the limited capacity of monthly free tests in one of your Snyk products. These are often also disabled during the initial setup period, but given the weekly frequency, it is less important.

### Configuring Notifications

There are several methods to control the behavior of email notifications, based on audience. and

* **Group Settings**: Configure the default settings for new Organizations, and get an overview of the current settings for each Organization.
* **Organization Settings**: Configure the default settings for new users in this Organization.
* **Personal Settings**: Each user can manually change their settings to have precise control over what emails they receive for each Organization they have access to.

{% hint style="info" %}
We suggest that you initially disable all email notifications so that users do not receive many notifications whilst Projects are imported. You would disable this at the Group level (for new organizations), and at the Organization level for all existing Organizations. Also see [Configure notifications](../phase-2-configure-account/set-visibility-and-configure-an-organization-template/configure-notifications.md).
{% endhint %}

Users (such as administrators) can opt-in to email notifications in their personal settings.

When you are ready to enable notifications more widely, you can see the overview for your Organizations and enable them in bulk on the **Group Settings** page.

See [Manage notifications](../../../snyk-admin/manage-notifications.md) for more details.

## Announce Snyk to your teams

Introducing Snyk to your developers is the first step in building awareness of Snyk functions, and how they can get the most out of Snyk. Typically we recommend disabling many of the more potentially intrusive features (for example, automatically raising PRs) until your teams are comfortable with Snyk use and results.

### Introducing Snyk to your Developers

This is your chance to get developers excited about better ways to introduce security into their work. If Snyk can make their lives easier, they’ll be more likely to get on board.

You can use the [Announcement templates](announcement-templates.md) to communicate the Snyk rollout to the rest of the developers.
