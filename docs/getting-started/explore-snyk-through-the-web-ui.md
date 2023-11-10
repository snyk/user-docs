# Explore Snyk through the Web UI

{% hint style="success" %}
If you have not done so already, you should [create an account](quickstart/create-or-log-in-to-a-snyk-account.md).
{% endhint %}

## Introduction to the Snyk Web UI and supported browsers

You can use the Snyk Web UI to run a full set of Snyk functions from your browser. See [supported browsers](../more-info/supported-browsers.md) for details.

<figure><img src="../.gitbook/assets/Screenshot 2023-07-13 at 11.06.29 AM.png" alt="Snyk Dashboard, the view when you log in"><figcaption><p>Snyk Dashboard, the view when you log in</p></figcaption></figure>

{% hint style="info" %}
You can also use Snyk functions from the [Snyk CLI](../snyk-cli/), from [within your IDE](../integrate-with-snyk/ide-tools/), and through the [Snyk API](../snyk-api/).
{% endhint %}

You can connect your code repositories, and then use Snyk to scan and secure your application code, open-source libraries, container registries, and configuration files.

## Web UI functions

This page explains the following functions of Snyk available through the Web UI:

* [Explore the Dashboard](explore-snyk-through-the-web-ui.md#dashboard)
* [View reports](explore-snyk-through-the-web-ui.md#view-reports)
* [Manage your Projects](explore-snyk-through-the-web-ui.md#manage-your-projects)
* [Manage your integrations](explore-snyk-through-the-web-ui.md#manage-your-integrations)
* [Manage Organization or Group members](explore-snyk-through-the-web-ui.md#manage-organization-or-group-members)
* [Snyk Organization and Group settings](explore-snyk-through-the-web-ui.md#snyk-organization-or-group-settings)
* [View Snyk updates](explore-snyk-through-the-web-ui.md#view-snyk-updates)
* [View helpful resources](explore-snyk-through-the-web-ui.md#view-helpful-resources)
* [Manage account preferences and settings](explore-snyk-through-the-web-ui.md#manage-account-preferences-and-settings)

### Explore the Dashboard

When you log in to an existing account, the Web UI opens to the Dashboard, where you can view your pending tasks and vulnerable Projects, invite team members, and add new Projects.

In the example that follows, the Snyk Dashboard for an Enterprise account shows pending tasks and vulnerable Projects.

<figure><img src="../.gitbook/assets/Screenshot 2023-07-13 at 11.09.02 AM.png" alt="Snyk dashboard for an Enterprise account"><figcaption><p>Snyk Dashboard for an Enterprise account</p></figcaption></figure>

#### Pending tasks

The **Pending tasks** section shows the next chores to be handled for the Projects in a Snyk Organization. This information includes:

* Pull Requests (PRs) that can be raised to fix vulnerabilities in some of the most vulnerable Projects.
* PRs that have already been raised by or through Snyk and are open and awaiting review.

Currently, Snyk tracks and flags PRs in GitHub, GitHub Enterprise, and Bitbucket Cloud only, and only for the top-most vulnerable Projects. If you use another SCM, the **Pending tasks** section shows only PRs that can be raised and not PRs that have already been raised.

<figure><img src="../.gitbook/assets/image (109) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt="Pending tasks and Vulnerable Projects listed on the Dashboard"><figcaption><p>Pending tasks and Vulnerable Projects listed on the Dashboard</p></figcaption></figure>

#### Add project

Use the **Add project** link on the Dashboard to add [Snyk Projects](../snyk-admin/introduction-to-snyk-projects/). Select how to add the Project from the dropdown.

Use the links for Projects on the Dashboard to explore and manage the metadata, retest, and fix options for the Target files in your Projects. Each link opens a Project details page where you can view the Project **Overview**, or switch to the **History** and **Settings** tabs.

* For Projects with the **Fix vulnerabilities** link, use the link to view Project details with an option to **Open a Fix PR.** Use this option to open a fix PR to implement the upgrades and patches in GitHub that address the issues.
* For Projects with a **View PR** link, use the link to open and view the Snyk-generated PR fixes in GitHub.

<figure><img src="../.gitbook/assets/demo-project-details-options (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (9).gif" alt="Demo, add project and Project details tabs"><figcaption><p>Demo, add project and Project details tabs</p></figcaption></figure>

### **View reports**

You can view [reports](../manage-issues/reporting/) to gain visibility and insights into the state of all your Projects, vulnerabilities, and license issues.

{% hint style="info" %}
**Feature availability**\
This feature is available for Enterprise plans.
{% endhint %}

The following example illustrates [Legacy reports](../manage-issues/reporting/legacy-reports/).

<figure><img src="../.gitbook/assets/reports.gif" alt="Demo of viewing reports, filters, and summary filters"><figcaption><p>Demo of viewing reports, filters, and summary filters</p></figcaption></figure>

### **Manage your** **Projects**

Select the **Projects** link in the navigation on the Dashboard to open the **Projects** listing page, where you can:

* Add a Project. Select the way you want to add the Project from the dropdown.
* Filter, group, and sort your Projects.
* View tips and the latest import log for your Projects.
* Select the link for each Project to view the Project details page with a summary and Issue information.
* Use the plus icon and add a Target from a custom location when Projects are grouped by Target. This allows for grouping Projects in another Target in the list.
* Use the settings icon on the ungrouped **Projects** listing or the **Settings** tab on the Project detail page to configure General and GitHub integration settings for notifications, Project testing, and pull request (PR) frequency. You can also look up the unique Project ID and deactivate or delete a Project on the Settings tab.
* View the Project history on the **History** tab.

<figure><img src="../.gitbook/assets/Project listing add projects.gif" alt="Options on the Projects listing page"><figcaption><p>Options on the Projects listing page</p></figcaption></figure>

### **Manage your** **Integrations**

Select **Integrations** from the navigation on the Dashboard to open the [Integrations](../integrate-with-snyk/) page, where you can do the following:

* Set up and view the [SCM integrations](../integrate-with-snyk/git-repository-scm-integrations/) that can connect to Snyk for vulnerability monitoring.
* Set up [Cloud platforms](../integrate-with-snyk/cloud-platforms-integrations/) and [Container integrations](../integrate-with-snyk/snyk-container-integrations/).
* Set up [CI/CD integrations](../integrate-with-snyk/snyk-ci-cd-integrations/), [IDE plugins](../integrate-with-snyk/ide-tools/), and [Gatekeeper plugins](../integrate-with-snyk/gatekeeper-plugins/).
* Set up [Notifications](../integrate-with-snyk/notifications-ticketing-system-integrations/) and [Cloud events](../integrate-with-snyk/event-forwarding/) integrations.

<figure><img src="../.gitbook/assets/image (123) (1) (2) (1).png" alt="Integrations page"><figcaption><p>Integrations page</p></figcaption></figure>

### Manage Organization or Group members

Select **Members** from the navigation on the Dashboard to view and manage users, roles, and how users authenticate in your Snyk [Organization](../snyk-admin/manage-users-in-organizations-and-groups/manage-users-in-organizations.md) or [Group](../snyk-admin/manage-users-in-organizations-and-groups/manage-users-in-a-group.md).

{% hint style="info" %}
You must be assigned the [required Admin roles and permissions](../snyk-admin/manage-permissions-and-roles/permissions-associated-with-each-pre-defined-role.md) to make changes on the **Members** tab.
{% endhint %}

### Snyk Organization or Group Settings

Use the **Settings** option to view and manage your Organization (team) or Group (company-wide) settings.

<figure><img src="../.gitbook/assets/Manage-settings-intro.png" alt="Group and Organization settings"><figcaption><p>Group and Organization settings</p></figcaption></figure>

See [Manage settings](../snyk-admin/manage-settings/) for more details.

### View Snyk updates

Select **Help** in the navigation on the Dashboard and then select **Product updates** to visit [snyk.io updates](https://updates.snyk.io/).

### View helpful resources

Select **Help** in the navigation on the Dashboard and then select an option to view resources with information about Snyk.

### Manage account preferences and settings

Select your **name** in the navigation on the Dashboard and then **Account settings** to open your [account settings](https://app.snyk.io/account) page, where you can view and configure your user account settings and your notification and sharing preferences.

You have access to the following information and options in the Account Settings:

* View and manage your API token or the Auth Token for free accounts. See [Obtaining your Snyk API token](how-to-obtain-and-authenticate-with-your-snyk-api-token.md) for instructions that apply to all applications and tools.
* View the list of your **Authorized Applications**.
* Set your preferred Organization. See [Manage Organizations: Set your preferred Organization](../snyk-admin/manage-groups-and-organizations/manage-organizations.md#set-your-preferred-organization).
* **Delete** your account.
* Manage your Account Settings for email **Notifications** (link in the left navigation), including Issue email alerts, Weekly report emails, and Usage alerts, as well as email notifications when reports are available and preferences for sales and marketing communications. See the [Manage notifications](../snyk-admin/manage-notifications.md) page for details about the email notifications as well as how to set personal preferences for notifications.
* Get a referral link to **Share with a Friend**. The link is in the left navigation of your Account Settings.
