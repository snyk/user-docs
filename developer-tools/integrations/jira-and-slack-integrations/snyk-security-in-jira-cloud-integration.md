---
description: How the Snyk Security app for Jira Cloud surfaces vulnerabilities in your Jira projects
---

# Snyk Security in Jira Cloud integration

{% hint style="info" %}
Jira Server and Jira Data Center are not supported.
{% endhint %}

Snyk Security in Jira Cloud helps developers identify, prioritize, and triage security vulnerabilities related to their code repositories directly from the Jira interface.

The Snyk Security in Jira Cloud integration mirrors your Snyk scan results from the Snyk platform to Jira. You can view Snyk results in your native Jira environment and create Jira work items for your results as needed. You can see which Snyk Organizations are connected and which Snyk Targets are associated with your Jira spaces.

## Prerequisites for the installation of Snyk Security in Jira Cloud

* Snyk Security in Jira Cloud is a Jira app. To install and configure it, you must be a Jira Cloud administrator (in the site-admins, administrators, or jira-administrators group). Contact your IT team if you need help installing it.
* To connect the app to Snyk, you must be a Snyk Organization administrator.
* To add containers, you must have Space admin permissions in Jira.
* The Security feature must be enabled in Jira. In your Space, enable it under Space settings (Features > Development > Security). For the exact steps, see Atlassian's documentation.
* For the app to load, the base URL of your Jira instance must match the base URL configured in your Snyk Jira integration. If your Jira instance uses URL rewrites or redirects that change the accessible URL, the app may fail to load.
* Ensure you have the following permission scopes in Jira, which are required for the integration to operate:&#x20;

<table><thead><tr><th width="344.5">Required scope in Jira</th><th>Purpose</th></tr></thead><tbody><tr><td>Write data to the host application</td><td>Synchronize vulnerabilities in Snyk with Jira so they appear in the Security tab in Jira.</td></tr><tr><td>Read data from the host application</td><td>Read Jira vulnerabilities to optimize the synchronization process.</td></tr><tr><td>Delete data from the host application</td><td>Remove vulnerabilities from Jira when a Snyk Organization is removed from Jira.</td></tr></tbody></table>

## Install Snyk Security in Jira Cloud

Install the app from the Atlassian Marketplace: search for **Snyk Security in Jira Cloud** and follow Atlassian's [instructions for installing a Marketplace app](https://support.atlassian.com/organization-administration/docs/installing-and-managing-app-access/).

Install the app that matches your Snyk region:

* For the EU or AU regions, install the dedicated EU or AU app.
* If you use SNYK-US-02 ([https://app.us.snyk.io/](https://app.us.snyk.io/)), install the SNYK-US-02 app.
* Otherwise, install the standard Snyk Security in Jira Cloud app.

## Configure the Snyk Security in Jira Cloud app

After installing the app, connect it to Snyk:

1. In Jira, open the manage apps area and select **Snyk Security in Jira**.
2. Log in to your Snyk account, or sign up for a new Snyk account.
3. In Snyk, select **Grant access** to allow Snyk to read your Jira account information.
4. Select the Snyk Organizations to connect to your Jira site, then select **Grant app access**.

## Link code repositories to Jira spaces

After connecting Snyk to Jira, link your repositories so you can triage vulnerabilities in Jira. Typically, R\&D engineering managers do this because they own the Jira spaces and know their team's repositories.

1. In Jira, open your Space and select the Security tab.
2. Select **Connect security containers**.
3. Select the Snyk application, then select **Connect security containers** again.
4. Select your Snyk Organization, then choose the Snyk Targets to connect to Jira.

After linking, developers can view recent vulnerabilities in the linked repositories and create Jira work items from them, or link them to existing work items.

{% hint style="info" %}
Syncing between Snyk and Jira is not immediate, so work items may take a short time to appear in Jira. Only security vulnerabilities appear on the Security tab.
{% endhint %}

### Deleting a target or repository

To remove a target or repository from Snyk that is connected to Jira, first remove it in Jira, then remove it in Snyk:

1. In Jira, open your Space and select the Security tab.
2. Select **Connect security containers**, then select the Snyk application.
3. Select the security container to remove, then select **Remove connection**.

## Automate work item creation in Jira

Use a Jira automation rule to create work items automatically for Snyk vulnerabilities. For how to build and enable a rule, see Atlassian's documentation on [creating Jira automation rules](https://support.atlassian.com/cloud-automation/docs/create-and-edit-jira-automation-rules/). Configure the rule with these Snyk-specific values:

* **Trigger:** Vulnerability Found. Set a minimum vulnerability severity so the rule runs only for vulnerabilities at or above that level.
* **Action — create the work item:** set the Summary to `Fix {{vulnerability.displayName}}` and the Description to `{{vulnerability.description.wiki}}`.
* **Action — link the vulnerability:** add the Snyk **Link vulnerability to work item** action so the new work item is linked back to the vulnerability.

## Manage security vulnerabilities in Jira

After installing and configuring the app, view vulnerabilities on the Security tab of the Jira Space page. In the Vulnerabilities section, Snyk shows each vulnerability's severity, status, and identifiers. Select a vulnerability's title to see its details in the Snyk Web UI.

### Search, filter, and sort vulnerabilities

Use the search bar and filters in the Vulnerabilities section to narrow the list to the vulnerabilities relevant to your Organization. Ignored and closed vulnerabilities are hidden by default; to show them, use the **Vuln. status** filter. To sort the list, select a column title to order all vulnerabilities by that attribute.

### Create a Jira work item from a vulnerability

When triaging vulnerabilities, create a Jira work item so the remediation work is planned and tracked in your sprint or backlog. On the Security tab, find the vulnerability and select **Create work item**.

### Link an existing Jira work item to a vulnerability

If the vulnerability already has a Jira work item, you can link the vulnerability to the existing Jira work item by selecting the three dots in the Actions column and selecting **Link work item.**

### Auto-close resolved vulnerabilities in Jira

Use a scheduled Jira automation rule with a JQL search to transition work items whose vulnerabilities are now resolved. For how to build and schedule a rule, see Atlassian's documentation on [creating Jira automation rules](https://support.atlassian.com/cloud-automation/docs/create-and-edit-jira-automation-rules/). Configure the rule with these Snyk-specific values:

* **Schedule:** run the rule on a schedule that suits your workflow.
* **JQL search:** `status != Done AND vulnerability[status] = CLOSED`
* **Action:** transition the matching work items to **Done**, or another status that fits your workflow.

According to your schedule, Jira searches for open items with closed vulnerabilities and closes each item.

## Uninstall Snyk Security in Jira Cloud

Uninstalling Snyk Security in Jira Cloud disconnects Snyk vulnerabilities from their associated Jira work items. To uninstall the app, follow Atlassian's [instructions for managing and removing apps](https://support.atlassian.com/organization-administration/docs/installing-and-managing-app-access/).
