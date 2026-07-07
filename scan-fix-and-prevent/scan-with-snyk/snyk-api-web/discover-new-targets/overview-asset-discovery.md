# Overview of asset discovery

Organizations often lack visibility into all their assets (web applications and APIs), which can lead to overlooked vulnerabilities and inadvertent security exposure. With Snyk API & Web asset discovery, you can identify your organization's assets and protect them before they become a liability.

## How asset discovery works

Snyk API & Web identifies assets through a multi-step process that discovers domains, analyzes infrastructure, and gathers detailed information about each asset.

### Step 1: Discover domains and subdomains

Snyk uses the following techniques to find domains and subdomains:

1.  **Cloud provider connections**

    Snyk connects to your configured cloud providers to discover web assets. The type of assets accessed depends on the provider you have configured.

    * **Cloudflare or AWS**: Snyk connects to select the configured DNS zones and retrieves domains and subdomains from those zones.
    * **Akamai**: Snyk connects to obtain the configured APIs and domains.
2.  **Certificate Transparency**

    Snyk searches Certificate Transparency, an Internet security standard for monitoring and auditing the issuance of digital certificates, to identify additional domains.
3.  **Domain guessing**

    Snyk makes informed guesses about possible domains based on known domains. For example, if `www.example.com` exists, Snyk tries `admin.example.com` or `api.example.com`.

### Step 2: Analyze infrastructure

With the initial list of domains, Snyk performs the following analysis:

1. Identifies the IP address (or addresses) to which each domain resolves.
2. Performs a network and port scan for each IP address to identify open services and their type.
3. Attempts to determine if the service is a web application or an API.

### Step 3: Gather asset information

The discovered domains and subdomains for web applications and APIs become assets. Snyk gathers more detailed information about them:

1. Captures screenshots of web applications.
2. Runs [Security Headers](https://securityheaders.com/) to generate a security score.

Snyk lists assets on the **Discovery** page. The following sections describe key concepts and common actions available on that page.

## Discovery sources

When you add a source to your account, Snyk starts performing regular discovery scans to identify assets in the source's attack surface.

There are four ways to add a source:

* By adding a domain. Visit [Scan a domain for asset discovery](scan-domain-asset-discovery.md).
* By connecting to Cloudflare. Visit [Scan a Cloudflare connection for asset discovery](scan-cloudflare-connection-asset-discovery.md).
* By connecting to AWS. Visit [Scan an AWS connection for asset discovery](scan-aws-connection-asset-discovery.md).
* By connecting to Akamai. Visit [Scan an Akamai connection for asset discovery](scan-akamai-connection-asset-discovery.md).

## Discovery assets

Snyk lists the assets resulting from discovery scans on the **Discovery** page.

At the top of the page, Snyk provides valuable information that you can use as quick filters to manage your assets and focus on the ones that matter most:

* **Found**: The total number of assets found so far. Click to show all assets.
* **New**: The total number of newly found assets. Click to filter the list to show only newly found assets.
* **Scanned**: The percentage of assets that were already added as targets to your Snyk account and were scanned, meaning they have a risk level associated. Click to filter the list to show only scanned assets.
* **Low score**: The percentage of assets with a Security Headers score of C or less. Click to filter the list to show assets with a score within this range.
* **High risk**: The percentage of assets already added as targets to your Snyk account that were scanned and identified as high risk. Click to filter the list to show only high-risk assets.

In addition to the quick filters, use the search box and generic filters to navigate the list.

## Discovery asset details

To view the details of a specific asset, click its name in the list. This opens a side panel where you can see all associated vulnerabilities (findings) and begin fixing them.

For more detailed guidance, visit [Interpret target scan results](../review-and-fix/interpret-target-scan-results.md).

The panel shows the asset's name and URL and has three tabs with information to help you manage your assets:

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

## Actions on discovery assets

### Add target and scan

You can add assets as targets to your Snyk account to scan them for vulnerabilities. Use the **Add target** button in the asset's row.

After you add an asset as a target, the **Add target** button changes to a **Scan** button, which you can click to start scanning for vulnerabilities.

After the target scan finishes, Snyk updates the **Risk** label of the corresponding asset with the risk identified during the target scan.

To access a target's details, click the three vertical dots next to the **Scan** button to display the overflow menu and select **View target**. From there, you can analyze its vulnerabilities (findings) and begin fixing them. Visit [Interpret target scan results](../review-and-fix/interpret-target-scan-results.md) to learn more.

### Mark as new or not new

Click the three vertical dots next to the main action button (**Add target**, **Scan**, or **Stop**) to show an overflow menu. The first two options, **Mark as new** and **Mark as not new**, identify the asset as new or not new, respectively. This action, along with the **State** filter, helps you organize and prioritize your list of assets.

### Hide or show

Click the three vertical dots next to the main action button to show an overflow menu. The **Hide** option helps you organize and prioritize your list of assets by hiding the ones you are not interested in at the moment. These assets do not disappear. Snyk filters them out of the default view. You can always filter them by selecting the **Hidden** option in the **State** filter. If you decide those assets are relevant again, click the **Show** option to restore visibility.

### Rename

Click the three vertical dots next to the main action button to show an overflow menu. The **Rename** option opens a modal where you update the asset name.

Because assets and targets are synced together, updating an asset name also updates its matching target's name. The opposite is also true: updating the target name also updates the asset name automatically.

### Set labels

You can assign labels to assets the same way you would to targets. These are synced between assets and targets, so applying a label to an asset also automatically applies it to the respective target and vice versa.

Click the **Set labels** dropdown to assign one or multiple labels. Filter existing labels by typing in the search field, or create new ones and apply them in a single step.

### Set owners

You can assign user labels to assets to identify their owners. Click the **Set owners** dropdown to assign one or multiple user labels. Filter existing labels by typing in the search field, or create new ones and apply them in a single step.

### Logs

Click the **Log** tab in the asset's details to open the logs of that asset since it was first discovered.

In the expanded view of the side panel, you can search for specific log messages or filter them with the **Event Type** dropdown. You can also add notes like you would for findings.

## Bulk actions on discovery assets

To improve your asset management, you can take certain actions in bulk. Select the check boxes of the assets that interest you, and the bulk actions become available at the top of the list.

You can:

* **Set labels**: Apply labels to assets to help you filter and manage them. For example, you can set a **CRITICAL** label on assets that are most critical to protect in your organization, bringing attention to them. Snyk also syncs labels assigned to assets to the respective targets.
* **Change state**: Change the state of a group of assets. For example, you can hide assets that are not important so you can focus on the ones that matter most. If you change your mind, you can show them again at any time. You can also set assets as new so you do not miss them, or as not new if you have completed all necessary actions.
