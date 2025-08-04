# Explore the Snyk Web UI

In the Snyk Web UI, you can manage your Projects, view and address security vulnerabilities, monitor dependencies, and review the health of your code. You can also configure account settings, manage API and Auth tokens, authorize applications, set organizational preferences, and customize email notifications.

You can visualize information at the Group or Organization level by clicking the name of the Group or Organization. General information, like Reports, Issues, Dependencies, Members, Settings, Help, and Settings, is available for all level types. For more information, see [Groups and Organizations](../snyk-admin/groups-and-organizations/).

## Group level

The following Snyk functions are available with the Web UI, at the Group level, with Snyk Essentials:

* [Organizations available for the selected Group](snyk-web-ui.md#organizations-available-for-the-selected-group)
* [View the assets dashboard](snyk-web-ui.md#view-the-assets-dashboard)
* [View and manage your assets inventory](snyk-web-ui.md#view-and-manage-your-assets-inventory)
* [Manage and customize your policies](snyk-web-ui.md#manage-and-customize-your-policies)
* [Manage integrations for asset discovery, asset coverage, and issues from third-party vendors](snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors)

### Organizations available for the selected Group

If you navigate to the Group level and select the Organizations page, you will see a list with all the Organizations that you have access to from that Group and the Organization role for each available Organization.

### View the assets dashboard

The Snyk Essentials Asset Dashboard reporting page provides a comprehensive overview of the security controls associated with your application. It presents critical metrics and data regarding your assets, such as scan coverage and a detailed breakdown of your inventory categorized by asset class, source, and other relevant information. Furthermore, the dashboard includes an extensive global filtering option, enabling users to filter results based on specific applications and owners, using the context data related to the application.

Navigate to the [Asset Dashboard](../manage-issues/reporting/available-snyk-reports.md#asset-dashboard) documentation section for more details.

### View and manage your assets inventory

You can use the [Inventory](../manage-assets/overview.md) page to organize your repository assets, enabling you to visualize all repository assets from your SCM tools, track Snyk product control coverage, and prioritize coverage mitigation based on business impact.

The Overview tab will provide quick insights into the discovered repositories. Scanned artifacts are not supported through Policies.

### Manage and customize your policies

For information on how to automate the process of adding business context and receiving notifications, see [Policies](../manage-risk/policies/assets-policies/).

### Manage integrations for asset discovery, asset coverage, and issues from third-party vendors

{% hint style="info" %}
The Group-level Integrations view focuses on asset management and discovery and is available with Snyk Enterprise.
{% endhint %}

The **Integrations** page shows all active integrations, [SCM](../developer-tools/scms/group-level-integrations/) or, [third-party](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md), including any data automatically synced from your existing Snyk Organizations, and provides access to the Integration Hub. You can use the Integrations Hub button to add SCM integrations, connect a third-party integration, add App Context to an SCM integration, or use the Snyk Runtime Sensor.

You can find an overview of all your integrations on the Snyk Web UI **Integrations** page. You can enable or disable your integrations, edit them, or remove them from your configuration.

For more details about available integrations, see [Snyk Essentials SCM integrations](../developer-tools/scms/group-level-integrations/) and [Integrate with Snyk](../integrations/overview.md).

#### Enable or disable an integration

You can have an integration connected or paused. Click play or pause to enable or disable an integration.

#### Add a new profile for an integration

Each integration can be configured to run on more than one profile. This is helpful when retrieving data from multiple instances within the same source.

To add a new profile:

1. Click the **Settings** icon from an already available integration profile.
2. Click **Add profile**.
3. Fill in the configuration fields and click **Done**.

#### Edit an integration

You can edit an existing integration by clicking **Settings** on the integrations card and then clicking **Settings** again on the added Organization for that integration.

<figure><img src="../.gitbook/assets/image (472) (2).png" alt=""><figcaption><p>Edit an existing integration from the Integration Hub</p></figcaption></figure>

{% hint style="info" %}
For security reasons, all credentials are anonymized when you open the Settings of an already existing integration.
{% endhint %}

#### Remove an integration

To remove an existing integration from your environment, select the integration and click **Delete.**

{% hint style="info" %}
You cannot restore an integration that was already deleted. You must enable it again.
{% endhint %}

## Organization level

The following Snyk functions are available with the Web UI, at Organization level:

* [Explore the Dashboard](snyk-web-ui.md#explore-the-dashboard)
* [Manage your Projects](snyk-web-ui.md#manage-your-projects)
* [View reports](snyk-web-ui.md#view-reports)
* [View dependencies and licenses](snyk-web-ui.md#view-dependencies-and-licenses)
* [Manage your integrations](snyk-web-ui.md#manage-your-integrations)
* [View and prioritize issues](snyk-web-ui.md#view-and-prioritize-issues)
* [Manage Organization or Group members](snyk-web-ui.md#manage-organization-or-group-members)
* [Set Snyk Organization and Group settings](snyk-web-ui.md#snyk-organization-or-group-settings)
* [View helpful resources](snyk-web-ui.md#view-helpful-resources)
* [Manage account preferences and settings](snyk-web-ui.md#manage-account-preferences-and-settings)

{% hint style="info" %}
You can also use Snyk functions from the [Snyk CLI](../cli-ide-and-ci-cd-integrations/snyk-cli/), [in your IDE](../cli-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/), and with the [Snyk API](../snyk-api/overview.md).
{% endhint %}

### Explore the Dashboard

When you log in to an existing account and select an Organization, the Web UI opens the Dashboard for that Organization. You can see your top pending tasks and vulnerable Projects, and you can add new Projects.

#### Top pending tasks

The **Pending tasks** section shows the next chores to be handled for the Projects in a Snyk Organization.

#### View Projects

Use the links for Projects on the Dashboard to explore and manage the metadata, retest, and fix options for the Target files in your Projects. Each link opens a Project details page where you can view the Project **Overview**, or switch to the **History** and **Settings** tabs.

For more information, see [Snyk Projects](../snyk-admin/snyk-projects/).

#### Fix vulnerabilities

Snyk tracks and flags Pull Requests (PRs) in the top-most vulnerable Projects, including:

* PRs that can be raised to fix vulnerabilities in some of the most vulnerable Projects.
* PRs that have already been raised by or through Snyk and are open and awaiting review.

For Projects with the **Fix vulnerabilities** link, use this link to view Project details with an option to open a fix PR. For details, see [Snyk Fix Pull or Merge Requests](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/).

{% hint style="info" %}
Snyk tracks and flags PRs in GitHub, GitHub Enterprise, and Bitbucket Cloud only, and only for the top-most vulnerable Projects. If you use another SCM, the **Pending tasks** section shows PRs that can be raised but not PRs that have already been raised.
{% endhint %}

#### Top vulnerable projects

Similarly, the **top vulnerable projects** section shows the Snyk Projects assessed as the most vulnerable, with similar functions available as the **Pending tasks** section.

#### Add Project

To add a Snyk Project, use the **Add project** link on the Dashboard. Select how to add the Project from the dropdown. See [Import a Project](./#import-a-project-to-scan-and-identify-issues) for more details.

### Manage your Projects

To open the **Projects** listing page, select the **Projects** link in the side menu. On this page, you can perform several actions:

* Add a Project. Select how you want to add the Project from the **Add projects** dropdown.
* Filter, group, and sort your Projects.
* View tips and the latest import log for your Projects.
* Select the link for each Project to view the Project details page with a summary and Issue information.
* Use the plus icon and add a Target from a custom location when Projects are grouped by Target. This allows for grouping Projects in another Target in the list.
* Use the settings icon on the ungrouped **Projects** listing or the **Settings** tab on the Project detail page to configure General and GitHub integration settings for notifications, Project testing, and pull request (PR) frequency. On the Settings tab, you can also look up the unique Project ID and deactivate or delete a Project.
* View the Project history on the **History** tab.

### Manage your integrations

You can set up a range of integrations with Snyk from the [Integrations](../integrations/overview.md) page available on the Dashboard.

## General settings available for all level types

The following Snyk functions are available with the Web UI, for all level types:

* [View reports](snyk-web-ui.md#view-reports)
* [View and prioritize issues](snyk-web-ui.md#view-and-prioritize-issues)
* [View dependencies and licenses](snyk-web-ui.md#view-dependencies-and-licenses)
* [Manage Organization or Group members](snyk-web-ui.md#manage-organization-or-group-members)
* [View and manage Snyk Organization or Group settings](snyk-web-ui.md#snyk-organization-or-group-settings)
* [View helpful resources](snyk-web-ui.md#view-helpful-resources)
* [Manage account preferences and settings](snyk-web-ui.md#manage-account-preferences-and-settings)

### View reports

{% hint style="info" %}
**Feature availability**\
Reporting is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

You can view [reports](../manage-issues/reporting/) to gain visibility and insights into the state of all your Projects, vulnerabilities, and license issues. You will find detailed definitions for information on the reports in the tooltips in the Reporting user interface.

The Reports page has all report types available at both the Group and Organization levels. The exception to the rule is the Asset Dashboard report, which is available only at the Group level.

### View and prioritize issues

{% hint style="info" %}
Issues are available only for Snyk Enterprise users.
{% endhint %}

The [Issues page](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md) provides a centralized view of all the issues identified by Snyk with additional asset context. This helps you and your team to better triage and remediate issues in Snyk.

### View dependencies and licenses

You can [view dependencies](../manage-risk/reporting/dependencies-and-licenses/view-dependencies.md) and [license information](../manage-risk/reporting/dependencies-and-licenses/view-licenses.md) for all Projects in your Organization or Group, using the **Dependencies** option in your Organization or Group menu.

### Manage Organization or Group members

Select **Members** from the dashboard's navigation to view and manage users, roles, and how users authenticate in your Snyk [Organization](../snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations.md) or [Group](../snyk-admin/groups-and-organizations/groups/manage-users-in-a-group.md).

The Members page is available at both the Group and Organization levels.

{% hint style="info" %}
You must be assigned the [required Admin roles and permissions](../snyk-admin/user-roles/pre-defined-roles.md) to make changes to the **Members** tab.
{% endhint %}

### Snyk Organization or Group Settings

Use the **Settings** option to view and manage your Organization or Group settings. For details, see [Group and Organization settings](../snyk-admin/groups-and-organizations/group-and-organization-settings.md).

<figure><img src="../.gitbook/assets/org_settings.png" alt=""><figcaption><p>Group and Organization settings</p></figcaption></figure>

### View helpful resources

Select the **Help** in the navigation on the Dashboard and then select an option to view resources with information about Snyk.

Select the **Help** > **Product updates** menu option to visit [snyk.io updates](https://updates.snyk.io/).

### Manage account preferences and settings

Select your **name** in the navigation on the Dashboard and then **Account settings** to open your [account settings](https://app.snyk.io/account) page, where you can view and configure your user account settings, including:

* View and manage your API token or the Auth Token for free accounts. For details, see [How to obtain and authenticate with your Snyk API token](./#obtain-and-use-your-snyk-api-token).
* View the list of your **Authorized Applications**.
* Set your preferred Organization. See [Manage Organizations: Set your preferred Organization](../snyk-admin/groups-and-organizations/organizations/set-your-preferred-organization.md).
* **Delete** your account.
* Manage your Account Settings for email **Notifications** (link in the left navigation), including Issue email alerts, Weekly report emails, and Usage alerts, as well as email notifications when reports are available and preferences for sales and marketing communications.  For details, see [Manage notifications](../snyk-admin/manage-notifications.md).
* Get a referral link to **Share with a Friend**. The link is in the left navigation of your Account Settings.
