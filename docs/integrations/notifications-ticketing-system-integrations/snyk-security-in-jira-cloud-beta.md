---
description: Enabling developers to see and triage security issues from Jira
---

# Snyk Security in Jira Cloud

Snyk Security in Jira Cloud helps developers identify, prioritize and triage security vulnerabilities related to their code repositories directly from the Jira interface.

## How Snyk Security in Jira Cloud works

The Snyk Security in Jira Cloud integration mirrors the information in the Snyk platform to Jira, such as scan results. You can view these Snyk results directly in your native Jira environment and create Jira issues for them as needed.

## Available plans and compatibility

### Plans and pricing

Snyk Security in Jira integration is available for all Snyk and Jira plans, including Free versions.

### Supported Jira Software versions

Snyk Security in Jira is available for Jira Cloud only. Jira Server and Jira Data Center are not supported.

## Snyk Security in Jira Cloud installation

To install [**Snyk Security in Jira Cloud**](https://marketplace.atlassian.com/apps/1230482/snyk-security-in-jira-cloud) from the Atlassian Marketplace and start triaging security issues in Jira, you need to take the following actions:

1. [Install Snyk Security in Jira Cloud](snyk-security-in-jira-cloud-beta.md#install-snyk-security).
2. [Configure the Snyk Security app](snyk-security-in-jira-cloud-beta.md#configure-snyk-security).
3. [Link code repositories to Jira projects](snyk-security-in-jira-cloud-beta.md#link-code-repositories-to-jira-projects).

### Requirements

* To install and configure Jira apps, you need to be an administrator in the site-admins, administrators, or jira-administrators groups. Contact your IT team to support your effort in installing the Snyk Security app.&#x20;
* To connect the Jira app to Snyk, you need to be a [Snyk Organization administrator](../../snyk-admin/introduction-to-snyk-administration.md#member-user-types).
* Security in Jira Cloud turned on by toggling "Security" on under "Project Settings" > "Features" under Jira

### Permission scopes in Jira required for the integration

These are the required permission scopes in Jira needed for the integration to operate.

<table><thead><tr><th width="344.5">Required Scope</th><th>Purpose</th></tr></thead><tbody><tr><td>Write data to the host application</td><td>Sync vulnerabilities from Snyk to Jira so they appear in the Security tab in Jira.</td></tr><tr><td>Read data from the host application</td><td>Read vulnerabilities from Jira to optimize the issues sync process.</td></tr><tr><td>Delete data from the host application</td><td>Remove vulnerabilities from Jira when a Snyk Organization is removed from Jira. </td></tr></tbody></table>

### Install Snyk Security

1. In Jira, go to **Apps** > **Find new apps.**
2. Search for **Snyk Security in Jira Cloud**.
3. Click the app and then select **Get it now.**
4. Review the information about the app, and select **Get it now**.
5. Follow the instructions to install it.

### Configure Snyk Security

1. Go to **Apps** > **Manage apps.**
2. In the menu on the left-hand side, select **Snyk Security in Jira**.
3. [Log in to your Snyk account](../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md), or [sign up for a new Snyk account](../../getting-started/quickstart/create-a-snyk-account/).
4. Select **Grant access** to allow Snyk to read your Jira Software account information.
5. Select the specific Snyk organizations to connect to your Jira site, and select **Grant app access**.

### Link code repositories to Jira projects

{% hint style="info" %}
R\&D engineering managers who own Jira projects and know their team's code repositories can handle this task.
{% endhint %}

1. In Jira, go to **Project settings** > **Toolchain** and find Snyk in the list of tools (see [What is the project toolchain in Jira Software?](https://support.atlassian.com/jira-software-cloud/docs/what-is-the-project-toolchain-in-jira/)).
2. Select the Add connection button :heavy\_plus\_sign: for Snyk.
3. Choose the container (code repository) from the list and select **Add container**.

Developers can now use the security feature to view recent vulnerabilities found in the linked code repositories and start [creating Jira issues](snyk-security-in-jira-cloud-beta.md#create-a-jira-issue-out-of-a-vulnerability) from them or [linking them to existing Jira issues](snyk-security-in-jira-cloud-beta.md#link-an-existing-jira-issue-to-a-vulnerability).

## Manage security vulnerabilities in Jira

After the app installation, you can view vulnerabilities in the security tab on the Jira project page.

<figure><img src="../../.gitbook/assets/security tab.png" alt="Overview of the Snyk Security for Jira Cloud tab."><figcaption><p>Snyk Security for Jira Cloud tab</p></figcaption></figure>

To find vulnerabilities, go to the **Vulnerabilities** section. Snyk shows the severity, status, and identifiers. Click the title to get the details in Snyk Web UI.

### Search, filter and sort vulnerabilities

Use the search bar and filters in the **Vulnerabilities** section to customize the list of vulnerabilities to those relevant to your organization.

Ignored and closed vulnerabilities are not shown in the **Vulnerabilities** section by default, but you can view them using the **Vuln. status filter**.

Select the title of a column in the table to sort all vulnerabilities by that attribute.

### Create a Jira issue from a vulnerability&#x20;

When triaging issues, you can add a Jira issue to the sprint or backlog to ensure the required work for resolving the vulnerability is planned and tracked.

Snyk provides vulnerability information to Jira, enabling users to have comprehensive data for resolving issues.

To add a Jira issue, go to the Snyk Security tab, find a vulnerability, and then click **Create issue**.

<figure><img src="../../.gitbook/assets/create issue.png" alt="Jira issue created from a vulnerability found by Snyk."><figcaption><p>Jira issue created from a vulnerability</p></figcaption></figure>

### Link an existing Jira issue to a vulnerability

If the vulnerability already has a Jira issue, you can link the existing Jira issue by clicking the horizontal ellipsis icon and selecting **Link issue.**

## Uninstall Snyk Security

{% hint style="danger" %}
Uninstalling **Snyk Security in Jira will disconnect Snyk vulnerabilities from their associated Jira issues**.\
\
To uninstall a Jira app, you need to be an administrator in the site-admins, administrators, or jira-administrators groups.
{% endhint %}

1. In Jira, go to **Apps** in the main menu, then select **Manage your apps.**
2. Select **Snyk Security in Jira.**
3. Click the **Uninstall** button.

