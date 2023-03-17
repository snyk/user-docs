# Managing notifications

### Introduction

Snyk automatically notifies you when new issues are found in the projects you are monitoring. These notifications can help make you aware of new risks in these projects.

Notifications can be sent:

* By email, if you have not disabled these in your [your notification settings](https://app.snyk.io/account/notifications).\
  **Note**: setting a Project to inactive will not stop Snyk from sending notifications, you must disable notifications separately for that Project.
* By Slack, if you have set up [Slack integration](https://docs.snyk.io/integrations/untitled-3/slack-integration).

Snyk also sends you a weekly update summarizing your security status across all of your organizations.

#### Notification types

Snyk offers notifications for:

* vulnerabilities
* license issues
* weekly summary report
* usage alerts
* report status

## **Managing notifications**

Snyk includes a range of controls to manage your own notifications.

Administrators can also manage the notification defaults for others in the [Group](notifications.md#define-group-notification-defaults) or [Organization](notifications.md#define-organization-notification-defaults).

You can also send notifications for an organization to a designated [Slack channel](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/slack-integration).

<figure><img src="../.gitbook/assets/image (48) (1) (1).png" alt="Notification example"><figcaption><p>Notification example</p></figcaption></figure>

### Define Group notification defaults

Define the notification settings at the [Group](managing-groups-and-organizations/whats-a-snyk-group.md) level to define the template for how Snyk sends issue alert emails, weekly report emails, usage alert emails, and report status emails for Organizations created in that Group. You can also access the default settings for existing individual Organizations at the Group level.

{% hint style="info" %}
When you change the Group default settings, it does not change the settings for existing Organizations or Projects. Individual users can override the default notification settings.
{% endhint %}

To navigate to the Group level notification settings:

1. Navigate to the Group overview for the group you want to change.
2. Select the Settings icon.
3. Select **Notifications**.

<figure><img src="../.gitbook/assets/image (349) (1).png" alt="Set Group-level notifications"><figcaption><p>Set Group-level notifications</p></figcaption></figure>

#### Issue alert emails

Issue alert emails are notifications Snyk sends the same day as it finds a new vulnerability, license issue, or remediation.

To set the defaults for issue alert emails:

1. Check the **Vulnerabilities** box when members of new organizations in this group should receive alert emails by default for new issues or remediations for all projects in that organization.
2. Check the **License Violations** box when members of new organizations in this group should receive alert emails for new license issues or remediations for all projects in that organization.
3. If either the **Vulnerabilities** or **License Violations** boxes are checked, indicate the severity of issues for which Snyk should send alert emails by selecting **All severities** or **Critical and high severity** from the drop down list.
4. To change the default for individual organizations, change the **Vulnerabilities**, **License Violations,** and **Severity** settings next to the organization name. These settings apply for any individual user who has not updated their personal notifications when you create new organizations in this group.

#### Weekly report emails

Weekly report emails are notifications Snyk sends to provide a summary of the vulnerability status across all projects and organizations to which a user belongs. For an organization with zero vulnerabilities across its projects, the notification lists the number of active projects, number of known vulnerabilities, and total dependencies.

<figure><img src="../.gitbook/assets/2022-06-27_13-45-21.png" alt="Weekly report emails"><figcaption><p>Weekly report emails</p></figcaption></figure>

To set the defaults for weekly report emails:

* Check the **Enabled by default** box when members of new organizations in this group should receive a weekly summary email.
* To change the defaults for individual organizations, clear or check the box next to the organization name. The defaults will apply for new organizations created in this group.

#### Usage alerts

Usage alert emails are notifications Snyk sends to warn when you are approaching usage limits.

<figure><img src="../.gitbook/assets/2022-06-27_13-47-51.png" alt="Usage alerts"><figcaption><p>Usage alerts</p></figcaption></figure>

To set the defaults for usage alerts:

* Check the Enabled by default box when members of new organizations in this group should receive usage alert emails.
* To change the defaults for individual organizations, clear or check the box next to the organization name. The defaults will apply for new organizations created in this group.

#### Report status emails

Report status emails are notifications Snyk sends to let you know when serverless project tests complete.

To set the defaults for usage alerts:

* Check the Enabled by default box when members of new organizations in this group should receive report status emails.
* To change the defaults for individual organizations, clear or check the box next to the organization name. The defaults will apply for new organizations created in this group.

### Define organization notification defaults

Define the notification settings for the organization to determine how Snyk sends emails to all individuals in the organization who have not updated their own personal notification preferences.

{% hint style="info" %}
When you change the organization default settings, it does not change the settings for existing projects or individual users. Individual users can override the default notification settings.
{% endhint %}

To navigate to the organization level notification settings:

1. Navigate to the organization you want to change.
2. Select the Settings icon.
3. Select **Notifications**.

<figure><img src="../.gitbook/assets/2022-06-27_13-56-10.png" alt="Organization-level notifications"><figcaption><p>Organization-level notifications</p></figcaption></figure>

#### Issue alert emails

Issue alert emails are notifications Snyk sends the same day as it finds a new vulnerability, license issue, or remediation.

To set the defaults for issue alert emails:

* To turn off issue alert emails for new projects loaded to this organization, clear the **On** box. Check the box when new projects in this organization should receive the defined alerts.
* Check the **Vulnerabilities** box to set the default for all projects in this organization to generate alert emails for new issues or remediations.
* Check the **License violations** box to set the default for all projects in this organization to generate alert emails for new license issues or remediations.
* If either the **Vulnerabilities** or **License violations** boxes are checked, indicate the severity of issues for which Snyk should send alert emails by selecting All severities or Critical and high severity from the drop down list.

#### Weekly report emails

Weekly report emails are notifications Snyk sends to provide a summary of the vulnerability status across all projects and organizations to which a user belongs.

To set the defaults for weekly report emails:

* Check the **Weekly report** box when members of this organization should receive a weekly summary email by default.

#### Usage alerts

Usage alert emails are notifications Snyk sends to warn when you are approaching usage limits.

To set the defaults for usage alerts:

* Check the **Approaching test limit** box when members of this organization should receive usage alert emails by default.

#### Report status emails

Report status emails are notifications Snyk sends to let you know when serverless project tests complete.

To set the defaults for report status emails:

* Check the Report ready to view box when members of this organization should receive report status emails by default.

### Change personal notification preferences

Organizations to which you belong have default settings for how Snyk sends emails.

To override these settings, change your notification preferences:

1. Select the dropdown next to your name in the top right corner of the Snyk Web UI.
2. Select **Notification preferences**.

The Account Settings page allows you to change the types of notifications for each organization to which you belong. You can also customize the notifications for individual projects.

#### Issue alert emails

Issue alert emails are notifications Snyk sends the same day as is finds a new vulnerability license issue, or remediation.

To customize your settings for issue alert emails for each organization for which you are a member:

* Check the **Vulnerabilities** box to receive alert emails for new issues or remediations for all projects in that organization.
* Check the **Licenses** box (if applicable) to receive alert emails for new license issues for all projects in that organization.
* If either the **Vulnerabilities** or **Licenses** boxes are checked, indicate the severity of issues for which Snyk should send alert emails by selecting **All severities** or **Critical and high severity** from the drop down list.

To customize your settings for individual projects:

1. Expand the organization to see a list of projects in that organization.
2. Next to each project name, clear or check the **Vulnerabilities** or **Licenses** boxes and set the severity for that project.

#### Weekly report emails

Weekly report emails are notifications Snyk sends to provide a summary of the vulnerability status for the projects in the organizations to which you belong.

To include organizations in your weekly report emails, check the box next to that organization.

#### Usage alerts

Usage alert emails are notifications Snyk sends to warn when you are approaching usage limits.

To receive usage alerts for an organization, check the box next to that organization.

#### Report status emails

Report status emails are notifications Snyk sends to let you know when serverless project tests are complete.

To receive report status emails for an organization, check the box next to that organization.
