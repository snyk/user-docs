# Snyk Security in Jira Cloud Integration

{% hint style="warning" %}
**Release status** \
Snyk Security in Jira Cloud is available for all Snyk and Jira plans, including Free versions.&#x20;

Jira Server and Jira Data Center are not supported.
{% endhint %}

Snyk Security in Jira Cloud helps developers identify, prioritize, and triage security vulnerabilities related to their code repositories directly from the Jira interface.

The Snyk Security in Jira Cloud integration mirrors your Snyk scan results from the Snyk platform to Jira. You can view Snyk results in your native Jira environment and create Jira issues for your results as needed. You can see which Snyk Organizations are connected and which Snyk Targets are associated with your Jira projects.

## Prerequisites for installation of Snyk Security in Jira Cloud

Snyk Security in Jira Cloud is a Jira app.

To install and configure the Jira app, you must be a Jira Cloud administrator in the site-admins, administrators, or jira-administrators group. Contact your IT team to support your effort in installing the Snyk Security in Jira Cloud app.

To connect the Jira app to Snyk, you must be a [Snyk Organization administrator](../../snyk-admin/introduction-to-snyk-administration.md#member-user-types).

To activate Security in Jira Cloud in Jira, navigate to **Project Settings > Features > Development > Security** and toggle **Security** **ON**.&#x20;

{% hint style="warning" %}
Ensure that the project is _**not**_ of type **Classic**. If you are unable to find **Features** in your project settings, this may be because the project type is **Classic**. In this case, the  project must be rebuilt or migrated to a newer version.
{% endhint %}

Ensure you have the following permission scopes in Jira, which are required for the integration to operate.

<table><thead><tr><th width="344.5">Required scope in Jira</th><th>Purpose</th></tr></thead><tbody><tr><td>Write data to the host application</td><td>Synchronize vulnerabilities in Snyk with Jira so they appear in the Security tab in Jira.</td></tr><tr><td>Read data from the host application</td><td>Read vulnerabilities from Jira to optimize the issues synchronization process.</td></tr><tr><td>Delete data from the host application</td><td>Remove vulnerabilities from Jira when a Snyk Organization is removed from Jira. </td></tr></tbody></table>

## Install Snyk Security in Jira Cloud

Follow these steps to install [**Snyk Security in Jira Cloud**](https://marketplace.atlassian.com/apps/1230482/snyk-security-in-jira-cloud) from the Atlassian Marketplace.

1. In Jira, navigate to **Apps** > **Find new apps.**
2. Search for **Snyk Security in Jira Cloud**.
3. Click the app and then select **Get it now.**
4. Review the information about the app, and select **Get it now**.
5. Follow the instructions to install the app.

## Configure the Snyk Security in Jira Cloud app

1. Go to **Apps** > **Manage apps.**
2. In the left menu, select **Snyk Security in Jira**.
3. [Log in to your Snyk account](broken-reference), or [sign up for a new Snyk account](../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md).
4. In Snyk, select **Grant access** to allow Snyk to read your Jira Software account information.
5. Select the specific Snyk Organizations to connect to your Jira site, and select **Grant app access**.

## Link code repositories to Jira projects

When you have completed the following steps to connect Snyk to Jira, you can start triaging security issues in Jira.

{% hint style="info" %}
Typically research and development engineering managers do his task because they own the Jira projects and know their team's code repositories.
{% endhint %}

1. In Jira, navigate to **Project settings** > **Toolchain** and find Snyk in the list of tools.\
   See [What is the project toolchain in Jira Software?](https://support.atlassian.com/jira-software-cloud/docs/what-is-the-project-toolchain-in-jira/)
2. Select the **Add connection** plus sign button for Snyk.
3. Choose the container code repository from the list and select **Add container**. This is a Snyk Target.\
   You can connect more than one code repository to Jira.

Developers can now use the security feature to view recent vulnerabilities found in the linked code repositories and start [creating Jira issues](snyk-security-in-jira-cloud-integration.md#create-a-jira-issue-from-a-vulnerability) from those vulnerabilities or [linking them to existing Jira issues](snyk-security-in-jira-cloud-integration.md#link-an-existing-jira-issue-to-a-vulnerability).

{% hint style="info" %}
Only security vulnerabilities will be shown on the Jira Security tab.
{% endhint %}

## Manage security vulnerabilities in Jira

After installing and configuring the Snyk Security in Jira Cloud app, you can view vulnerabilities on the security tab on the Jira project page.

<figure><img src="../../.gitbook/assets/Secuirity-Jira-Cloud-tab.png" alt="Snyk Security in Jira Cloud tab."><figcaption><p>Snyk Security in Jira Cloud tab</p></figcaption></figure>

To find vulnerabilities, navigate to the **Vulnerabilities** section. Snyk shows the severity, status, and identifiers. Click the title to see the details in Snyk Web UI.

### Search, filter, and sort vulnerabilities

Use the search bar and filters in the **Vulnerabilities** section to customize the list of vulnerabilities to show those relevant to your Organization.

Ignored and closed vulnerabilities are not shown in the **Vulnerabilities** section by default, but you can view them using the **Vuln. status filter**.

Select the title of a column in the table to sort all vulnerabilities by that attribute.

### Create a Jira issue from a vulnerability&#x20;

When triaging issues, you can add a Jira issue to the sprint or backlog to ensure the required work for resolving the vulnerability is planned and tracked.

Snyk provides vulnerability information to Jira, enabling users to have comprehensive data for resolving issues.

To add a Jira issue, navigate to the Snyk Security tab, find a vulnerability, and click **Create issue**.

<figure><img src="../../.gitbook/assets/Security-Jira-Cloud-create-issue.png" alt="Jira issue created from a vulnerability found by Snyk"><figcaption><p>Jira issue created from a vulnerability found by Snyk</p></figcaption></figure>

### Link an existing Jira issue to a vulnerability

If the vulnerability already has a Jira issue, you can link the vulnerability to the existing Jira issue by clicking the three dots in the Actions column and selecting **Link issue.**

## Uninstall Snyk Security in Jira Cloud

{% hint style="danger" %}
Uninstalling Snyk Security in Jira Cloud will disconnect Snyk vulnerabilities from their associated Jira issues.\
\
To uninstall a Jira app, you must be an administrator in the site-admins, administrators, or jira-administrators groups.
{% endhint %}

1. In Jira, navigate to **Apps** in the main menu and select **Manage your apps.**
2. Select **Snyk Security in Jira.**
3. Click the **Uninstall** button.

