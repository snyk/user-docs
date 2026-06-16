# Overview of asset discovery

Organizations often lack visibility into all their assets (web applications and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

## How asset discovery works

Snyk API & Web identifies assets through a multi-step process that discovers domains, analyzes infrastructure, and gathers detailed information about each asset.

### Step 1: Discover domains and subdomains

Snyk API & Web uses the following techniques to find domains and subdomains:

1. **Cloud provider connections**

   Snyk API & Web connects to your configured cloud providers to discover web assets. The type of assets accessed depends on the provider you have configured.

   * **Cloudflare or AWS**: Snyk API & Web connects to select the configured DNS zones and retrieves domains and subdomains from those zones.
   * **Akamai**: Snyk API & Web connects to obtain the configured APIs and domains.

1. **Certificate Transparency**

   Snyk API & Web searches Certificate Transparency, an Internet security standard for monitoring and auditing the issuance of digital certificates, to identify additional domains.

1. **Domain guessing**

   Snyk API & Web makes informed guesses about possible domains based on known domains. For example, if `www.example.com` exists, Snyk API & Web tries `admin.example.com` or `api.example.com`.

### Step 2: Analyze infrastructure

With the initial list of domains, Snyk API & Web performs the following analysis:

1. Identifies the IP address (or addresses) to which each domain resolves.
1. Performs a network and port scan for each IP address to identify open services and their type.
1. Attempts to determine if the service is a web application or an API.

### Step 3: Gather asset information

The discovered domains and subdomains for web applications and APIs become assets. Snyk API & Web gathers more detailed information about them:

1. Captures screenshots of web applications.
1. Runs [Security Headers](https://securityheaders.com/) to generate a security score.

Assets are listed in the **Discovery** page of Snyk API & Web. The following sections describe key concepts and common actions available on that page.

## Discovery sources

When you add a source to your account, Snyk API & Web starts performing regular discovery scans to identify assets in the source's attack surface.

There are four ways to add a source:

* By adding a domain. Visit [Scan a domain for asset discovery](scan-domain-asset-discovery.md).
* By connecting to Cloudflare. Visit [Scan a Cloudflare connection for asset discovery](scan-cloudflare-connection-asset-discovery.md).
* By connecting to AWS. Visit [Scan an AWS connection for asset discovery](scan-aws-connection-asset-discovery.md).
* By connecting to Akamai. Visit [Scan an Akamai connection for asset discovery](scan-akamai-connection-asset-discovery.md).

## Discovery assets

The assets resulting from discovery scans are listed in the **Discovery** page of Snyk API & Web.

At the top of the page, Snyk API & Web provides valuable information that you can use as quick filters to manage your assets and focus on the ones that matter most:

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-filters.png" alt="Asset discovery filters showing found, new, scanned, low score, and high risk metrics"></figure>

* **Found**: The total number of assets found so far. Click to show all assets.
* **New**: The total number of newly found assets. Click to filter the list to show only newly found assets.
* **Scanned**: The percentage of assets that were already added as targets to your Snyk API & Web account and were scanned, meaning they have a risk level associated. Click to filter the list to show only scanned assets.
* **Low score**: The percentage of assets with a Security Headers score of C or less. Click to filter the list to show assets with a score within this range.
* **High risk**: The percentage of assets already added as targets to your Snyk API & Web account that were scanned and identified as high risk. Click to filter the list to show only high-risk assets.

In addition to the quick filters, use the search box and generic filters to navigate the list.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-search-filters.png" alt="Asset discovery search box and filter options"></figure>

## Discovery asset details

To view the details of a specific asset, click its name in the list. This opens a side panel where you can see all associated vulnerabilities (findings) and begin fixing them.

For more detailed guidance, visit [Interpret target scan results](../review-and-fix/interpret-target-scan-results.md).

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-asset-details.png" alt="Asset details panel showing risk, insights, and images"></figure>

The panel shows the asset's name and URL and has three tabs with useful information to help you manage your assets:

* **Overview**
* **Redirect And IPs**
* **Log**

### Overview tab

The **Overview** tab shows three sections:

* **Risk**
  * The asset's risk classification (only set after the asset was added as a target and scanned at least once).
  * The Security Headers score with a link to the respective Security Headers details page.
* **Insights**
  * When the asset was last seen.
  * When the asset was first seen.
  * Labels and owners, if any have been set.
  * The list of technologies found on the asset.
  * How the asset was discovered.
* **Images**
  * A thumbnail of the asset's screenshot (click to see the full-size image).

### Redirect And IPs tab

The **Redirect And IPs** tab shows the following lists:

* **Redirect from**: The list of URLs that redirect to the asset, if any.
* **Redirect to**: The list of URLs to which the asset is redirected, if any.
* **IP**: The list of IP addresses of the asset.

### Log tab

The **Log** tab shows the list of events associated with that asset, including scans performed, risk or score updates, and detection of new technologies. You can also add notes for your team to see.

{% hint style="info" %}
You can switch between a small side panel and a full page by clicking the buttons at the top of the panel.
{% endhint %}

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-panel-expand.png" alt="Button to expand asset details panel to full page view"></figure>

## Actions on discovery assets

### Add target and scan

If you decide to, you can add assets as targets to your Snyk API & Web account to scan them for vulnerabilities. Use the **Add target** button in the asset's row.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-add-target.png" alt="Add target button for a discovery asset"></figure>

After adding an asset as a target, the **Add target** button changes to a **Scan** button, which you can click to start scanning for vulnerabilities.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-scan-button.png" alt="Scan button for a target that has been added"></figure>

After the target scan finishes, the **Risk** label of the corresponding asset is updated with the risk identified during the target scan.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-risk-label.png" alt="Risk label showing the severity level after scanning"></figure>

To access a target's details, click the three vertical dots next to the **Scan** button to display the overflow menu and choose **View target**. From there, you can analyze its vulnerabilities (findings) and begin fixing them. Visit [Interpret target scan results](../review-and-fix/interpret-target-scan-results.md) to learn more.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-view-target.png" alt="View target option in overflow menu"></figure>

### Mark as new or not new

Clicking the three vertical dots next to the main action button (**Add target**, **Scan**, or **Stop**) shows an overflow menu. The first two options, **Mark as new** and **Mark as not new**, let you identify the asset as new or not new, respectively. This action, along with the **State** filter, lets you organize and prioritize your list of assets.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-mark-new.png" alt="Mark as new or not new options in overflow menu"></figure>

### Hide or show

Clicking the three vertical dots next to the main action button shows an overflow menu. The **Hide** option lets you organize and prioritize your list of assets by hiding the ones you are not interested in at the moment. These assets do not disappear; they are filtered out of the default view. You can always filter them by choosing the **Hidden** option in the **State** filter. If you decide those assets are relevant again, click the **Show** option to restore visibility.

### Rename

Clicking the three vertical dots next to the main action button shows an overflow menu. The **Rename** option opens a modal that allows you to update the asset name.

Because assets and targets are synced together, updating an asset name also updates its matching target's name. The opposite is also true: updating the target name also updates the asset name automatically.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-rename.png" alt="Rename modal for updating asset name"></figure>

### Set labels

You can assign labels to assets the same way you would to targets. These are synced between assets and targets, so applying a label to an asset also automatically applies it to the respective target and vice versa.

Click the **Set labels** dropdown to assign one or multiple labels. Filter existing labels by typing in the search field, or create new ones and apply them in a single step.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-set-labels.png" alt="Set labels dropdown for assigning labels to assets"></figure>

### Set owners

You can assign user labels to assets to identify their owners. Click the **Set owners** dropdown to assign one or multiple user labels. Filter existing labels by typing in the search field, or create new ones and apply them in a single step.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-set-owners.png" alt="Set owners dropdown for assigning user labels to assets"></figure>

### Logs

Click the **Log** tab in the asset's details to open the logs of that asset since it was first discovered.

In the expanded view of the side panel, you can search for specific log messages or filter them with the **Event Type** dropdown. You can also add notes like you would for findings.

## Bulk actions on discovery assets

To improve your asset management, you can take certain actions in bulk. Check the checkboxes of the assets that interest you, and the bulk actions become available at the top of the list.

<figure><img src="../../../.gitbook/assets/overview-asset-discovery-bulk-actions.png" alt="Bulk actions available when multiple assets are selected"></figure>

You can choose to:

* **Set labels**: Apply labels to assets to help you filter and manage them. For example, you can set a **CRITICAL** label on assets that are most critical to protect in your organization, bringing attention to them. Labels assigned to assets are also synced to the respective targets.
* **Change state**: Change the state of a group of assets. For example, you can hide assets that are not important so you can focus on the ones that matter most. If you change your mind, you can show them again at any time. You can also set assets as new so you do not miss them, or as not new if you have completed all necessary actions.
