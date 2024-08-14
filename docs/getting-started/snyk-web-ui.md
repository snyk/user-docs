# Snyk Web UI

{% hint style="info" %}
If you have not done so already, [create an account](quickstart/create-or-log-in-to-a-snyk-account.md) with Snyk.
{% endhint %}

Use the Snyk Web UI to run Snyk from any [supported browser](quickstart/#supported-browsers).

In the Snyk Web UI, you can visualize information at the [Group](snyk-web-ui.md#group-level) or [Organization](snyk-web-ui.md#organization-level) level by clicking the name of the Group or Organization. [General information](snyk-web-ui.md#general-settings-available-for-all-level-types), like the Reports, Issues, Dependencies, Members, Settings, Help, and Settings, is available for all level types.

{% hint style="info" %}
Snyk AppRisk Essentials is available with Snyk Enterprise.&#x20;
{% endhint %}

## Group level

The following Snyk functions are available with the Web UI:

* [Organizations available for the selected Group](snyk-web-ui.md#organizations-available-for-the-selected-group)
* [View the assets dashboard](snyk-web-ui.md#view-the-assets-dashboard)
* [View and manage your assets inventory](snyk-web-ui.md#view-and-manage-your-assets-inventory)
* [Manage and customize your policies](snyk-web-ui.md#manage-and-customize-your-policies)
* [Manage integrations for asset discovery, asset coverage, and issues from third-party vendors](snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors)

### Organizations available for the selected Group

If you navigate to the Group level and select the Organizations page, you will see a list with all the Organizations that you have access to from that Group and the Organization role for each available Organization.

### View the assets dashboard

{% hint style="warning" %}
The Asset Dashboard menu option is available only for Snyk AppRisk Essentials users.&#x20;

If you use Snyk AppRisk Pro, navigate to the [Application Analytics ](../manage-risk/enterprise-analytics/application-analytics.md)page.
{% endhint %}

On the Snyk AppRisk dashboard page, you can add widgets that display an overview of your application and security controls.

You can customize the dashboard widgets as desired. Choose to rename or modify display configurations, or add multiple widget instances.

#### Configure a widget

Customize your dashboard with the available widgets. You can change the settings of an existing widget or the way it is displayed. You have the ability to move a widget around the dashboard, rename it, display or hide the legend, view it in full screen, export or download it. See the [Repositories assets](../manage-assets/assets-inventory-layouts.md#repository-assets) section for additional details regarding archived or deleted repositories and the information presented in the dashboard widgets.

#### Settings menu

You can make several changes to a widget. All widgets allow you to change the name. Other particular settings are available for each widget. You can access the Settings menu by following these steps:

1. Select a widget and click the **Setting** menu.
2. Customize the widget by changing its name or other specific details.&#x20;
3. After all changes are done, click **Apply**.

#### Widget menu

Access the full list of general options from the widget menu. You can access the widget menu by following these steps:

1. Select a widget and click the menu available in the top right corner of the widget.
2. Select one of the following actions:
   * View in full-screen
   * CSV export
   * XLS export
   * Download PNG
   * Download PDF

### View and manage your assets inventory

{% hint style="warning" %}
The Inventory menu option is available only for Snyk AppRisk users.&#x20;
{% endhint %}

The [Inventory](../manage-assets/) is available if you are using Snyk AppRisk. You can use the Inventory page to organize your repository assets, enabling you to visualize all repository assets from your SCM tools, track Snyk product control coverage, and prioritize coverage mitigation based on business impact.

Each line in the inventory represents either a repository asset or a scanned artifact from Snyk that is likely a repository but lacks some identifying information. Scanned artifacts are not supported through Policies.

### Manage and customize your policies

{% hint style="warning" %}
The Policies menu option is available only for Snyk AppRisk users.&#x20;
{% endhint %}

You can use the [Policies](../manage-risk/policies/assets-policies/) page to automate the process of adding business context and receiving notifications.

### Manage integrations for asset discovery, asset coverage, and issues from third-party vendors

{% hint style="warning" %}
The Integrations menu option for the Group level is available only for Snyk AppRisk users.&#x20;
{% endhint %}

The Integrations page shows all active integrations, [SCM](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) or, [third-party](../integrate-with-snyk/#integrations-for-snyk-apprisk), including any data automatically synced from your existing Snyk Organizations, and provides access to the Integration Hub. You can use the Integrations Hub button to add SCM integrations, connect a third-party integration, add App Context to an SCM integration, or use the Snyk Runtime Sensor.&#x20;

You can find the overview of all your integrations on the Snyk Web UI Integrations page and more details about the available integrations on the [Snyk AppRisk SCM integrations](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) page and on the [Integrate with Snyk](../integrate-with-snyk/#integrations-for-snyk-apprisk) page.

You can enable or disable your integrations, edit them, or remove them from your configuration.

#### Enable or disable an integration

You can have an integration connected or paused. Click play or pause to enable or disable an integration.

#### Add a new profile for an integration

Each integration can be configured to run on more than one profile. This is helpful when retrieving data from multiple instances within the same source.&#x20;

You can add a new profile by following these steps:

1. Click the **Settings** icon from an already available integration profile.
2. Click **Add profile**.
3. Fill in the configuration fields and click **Done**.

#### Edit an integration

You can edit an existing integration by clicking Settings on the integrations card and then clicking Settings again on the added organization for that integration.

<figure><img src="../.gitbook/assets/image (472).png" alt="Edit an existing integration from the Integration Hub"><figcaption><p>Edit an existing integration from the Integration Hub</p></figcaption></figure>

{% hint style="info" %}
For security reasons, all credentials are anonymized when you open the Settings of an already existing integration.
{% endhint %}

#### Remove an integration

Select an existing integration and click **Delete** to remove it from your environment.

You cannot restore an integration that was already deleted. To add it again, you need to follow the steps from the [Using the Integration Hub](snyk-web-ui.md#using-the-integration-hub) section.

## Organization level

The following Snyk functions are available with the Web UI:

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
You can also use Snyk functions from the [Snyk CLI](../snyk-cli/), [in your IDE](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/), and with the [Snyk API](../snyk-api/).
{% endhint %}

### Explore the Dashboard

When you log in to an existing account and select an Organization, the Web UI opens the Dashboard for that Organization:

<figure><img src="../.gitbook/assets/image (503).png" alt="Snyk Dashboard, the view when you log in and select an Organization"><figcaption><p>Snyk Dashboard, the view when you log in and select an Organization</p></figcaption></figure>

Use the Dashboard to see your top pending tasks and vulnerable Projects and add new Projects.

#### Top pending tasks

The **Pending tasks** section shows the next chores to be handled for the Projects in a Snyk Organization.&#x20;

#### View Projects

Use the links for Projects on the Dashboard to explore and manage the metadata, retest, and fix options for the Target files in your Projects. Each link opens a Project details page where you can view the Project **Overview**, or switch to the **History** and **Settings** tabs.

See [Snyk Projects](../snyk-admin/snyk-projects/) for more information.

#### Fix vulnerabilities

Snyk tracks and flags Pull Requests (PRs) in the top-most vulnerable Projects, including:

* PRs that can be raised to fix vulnerabilities in some of the most vulnerable Projects.
* PRs that have already been raised by or through Snyk and are open and awaiting review.

For Projects with the **Fix vulnerabilities** link, use this link to view Project details with an option to open a fix PR. See [Snyk Fix Pull or Merge Requests](../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/) for more details.

{% hint style="warning" %}
Currently, Snyk tracks and flags PRs in GitHub, GitHub Enterprise, and Bitbucket Cloud only, and only for the top-most vulnerable Projects. If you use another SCM, the **Pending tasks** section shows PRs that can be raised but not PRs that have already been raised.
{% endhint %}

#### Top vulnerable projects

Similarly, the **Top vulnerable projects** section shows the [Snyk Projects](../snyk-admin/snyk-projects/) assessed as the most vulnerable, with similar functions available as the **Pending tasks** section.&#x20;

#### Add Project

Use the **Add project** link on the Dashboard to add [Snyk Projects](../snyk-admin/snyk-projects/). Select how to add the Project from the dropdown. See [Import a Project](quickstart/import-a-project.md) for more details.

### **Manage your** **Projects**

Select the **Projects** link in the side menu to open the **Projects** listing page:

* Add a Project. Select how you want to add the Project from the **Add projects** dropdown.
* Filter, group, and sort your Projects.
* View tips and the latest import log for your Projects.
* Select the link for each Project to view the Project details page with a summary and Issue information.
* Use the plus icon and add a Target from a custom location when Projects are grouped by Target. This allows for grouping Projects in another Target in the list.
* Use the settings icon on the ungrouped **Projects** listing or the **Settings** tab on the Project detail page to configure General and GitHub integration settings for notifications, Project testing, and pull request (PR) frequency. On the Settings tab, you can also look up the unique Project ID and deactivate or delete a Project.
* View the Project history on the **History** tab.

### **Manage your** **Integrations**

Select **Integrations** from the navigation on the Dashboard to open the [Integrations](../integrate-with-snyk/) page, where you can set up a range of integrations with Snyk:

<figure><img src="../.gitbook/assets/image (506).png" alt="Organization level - Integrations page"><figcaption><p>Organization level - Integrations page</p></figcaption></figure>

## General settings available for all level types

The following Snyk functions are available with the Web UI:

* [View reports ](snyk-web-ui.md#view-reports)
* [View and prioritize issues ](snyk-web-ui.md#view-and-prioritize-issues)
* [View dependencies and licenses  ](snyk-web-ui.md#view-dependencies-and-licenses)
* [Manage Organization or Group members](snyk-web-ui.md#manage-organization-or-group-members) &#x20;
* [View and manage Snyk Organization or Group settings ](snyk-web-ui.md#snyk-organization-or-group-settings)
* [View helpful resources  ](snyk-web-ui.md#view-helpful-resources)
* [Manage account preferences and settings](snyk-web-ui.md#manage-account-preferences-and-settings)

### **View reports**

You can view [reports](../manage-issues/reporting/) to gain visibility and insights into the state of all your Projects, vulnerabilities, and license issues. You will find detailed definitions for information on the reports in the tooltips in the Reporting user interface.

The Reports page is available at both the Group and Organization levels.

{% hint style="info" %}
**Feature availability**\
Reporting is available for Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

### View and prioritize issues

{% hint style="warning" %}
The Issues menu option is available only for Snyk AppRisk Pro users.&#x20;
{% endhint %}

The [Issues page](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-apprisk.md) provides a centralized view of all the issues identified by Snyk with additional asset context. This helps you and your team to better triage and remediate issues in Snyk.

The Reports page is available at both the Group and Organization levels.

### View dependencies and licenses

You can [view dependencies](https://docs.snyk.io/manage-risk/reporting/dependencies-and-licenses/view-dependencies) and [license information](https://docs.snyk.io/manage-risk/reporting/dependencies-and-licenses/view-licenses) for all Projects in your Organization or Group, using the **Dependencies** option in your Organization or Group menu:

### Manage Organization or Group members

Select **Members** from the navigation on the Dashboard to view and manage users, roles, and how users authenticate in your Snyk [Organization](../snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations.md) or [Group](../snyk-admin/groups-and-organizations/groups/manage-users-in-a-group.md).&#x20;

The Members page is available at both Group and Organization level.

{% hint style="warning" %}
You must be assigned the [required Admin roles and permissions](../snyk-admin/manage-permissions-and-roles/pre-defined-roles.md) to make changes to the **Members** tab.
{% endhint %}

### Snyk Organization or Group Settings

Use the **Settings** option to view and manage your Organization or Group settings:

<figure><img src="../.gitbook/assets/image (507).png" alt="Group and Organization settings"><figcaption><p>Group and Organization settings</p></figcaption></figure>

See [Manage settings](../snyk-admin/groups-and-organizations/group-and-organization-settings.md) for more details.

### View helpful resources

Select the **Help** in the navigation on the Dashboard and then select an option to view resources with information about Snyk.

Select the **Help** > **Product updates** menu option to visit [snyk.io updates](https://updates.snyk.io/).

### Manage account preferences and settings

Select your **name** in the navigation on the Dashboard and then **Account settings** to open your [account settings](https://app.snyk.io/account) page, where you can view and configure your user account settings, including:

* View and manage your API token or the Auth Token for free accounts. See [How to obtain and authenticate with your Snyk API token](how-to-obtain-and-authenticate-with-your-snyk-api-token.md) for more detalis.
* View the list of your **Authorized Applications**.
* Set your preferred Organization. See [Manage Organizations: Set your preferred Organization](../snyk-admin/groups-and-organizations/organizations/create-and-delete-organizations.md#set-your-preferred-organization).
* **Delete** your account.
* Manage your Account Settings for email **Notifications** (link in the left navigation), including Issue email alerts, Weekly report emails, and Usage alerts, as well as email notifications when reports are available and preferences for sales and marketing communications. See the [Manage notifications](../snyk-admin/manage-notifications.md) page for more details.
* Get a referral link to **Share with a Friend**. The link is in the left navigation of your Account Settings.
