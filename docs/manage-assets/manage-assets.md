# Overview

Snyk defines an asset as an identifiable entity that is part of an application and relevant to security and developers. Snyk is generally focused on the development stages of application software, secures repository assets containing software package assets, and builds artifacts like container image assets.

## Inventory menu

Use the **Inventory** to organize your repository assets, visualize SCM tool assets, track Snyk coverage, and prioritize mitigation based on business impact.

The **Inventory** page in the Snyk Web UI is structured into multiple tabs and provides information about your assets.

Here is a list of all the available inventory tabs:

* **Overview:** Provides quick insights into discovered repositories, enabling AppSec teams to effectively operationalize their program using Snyk.&#x20;
* **All Assets**: All the discovered assets, grouped by their type.&#x20;
* **Asset Hierarchy**: All the assets in a hierarchical structure. The list of assets is sorted by issue counts, and where applicable, the package assets are listed underneath the repositories where they are located. The Asset hierarchy is visible only when there are no filters applied.
* **Teams**: SCM repository assets grouped by teams. Only SCM organizations with teams and repositories assigned to a team appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk Essentials.

Each inventory tab may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory layout.

The following video summarizes the main features of the Inventory view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1746127384/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-4c_-_v2_-_Inventory_Overview.mp4" %}
Reviewing asset inventory
{% endembed %}

{% hint style="info" %}
**Release status** \
Risk factors for Snyk AppRisk are in Early Access and are available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.

**Runtime discovered** and **Runtime last seen** filters take their release status from the [runtime integration](../integrations/connect-a-third-party-integration.md) for which they provide data and are available only for Snyk AppRisk.
{% endhint %}

## Assets overview

In the Snyk Enterprise context, you can find information about assets under the following menus: [Inventory](manage-assets.md), [Policies](../manage-risk/policies/assets-policies/), [Insights for Snyk AppRisk](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md), and [Analytics](../manage-risk/analytics/application-analytics.md).&#x20;

The main source of information is the **Inventory** menu. Here, you can view, filter, and manage your assets.&#x20;

The [**Inventory**](manage-assets.md#inventory-overview) menu provides a centralized view of all the issues identified by Snyk with additional asset context. The Inventory serves as the cornerstone of this interconnected system, providing detailed information on each asset to facilitate further processes.

In the [**Policies**](../manage-risk/policies/assets-policies/) menu, you can create policies for organizing the assets, classifying them, and always being up to date with the latest information about an asset. The policies are put in place to regulate these assets, making sure that each one follows security standards, organizational best practices, and regulatory requirements. These policies are enforced using the data from the Inventory, ensuring consistent compliance across all assets.

In the [**Analytics**](../manage-risk/analytics/application-analytics.md) menu, you can review and explore your application status and results, from a top-down approach. The Analytics page then uses this data to provide detailed insights into the performance and security metrics for the assets. This page helps identify trends, detect anomalies, and pinpoint potential vulnerabilities, providing a data-driven approach to asset management.

The [**Issues**](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md) menu uses information from both the Inventory and Policies menus to prioritize issues based on their severity and impact. This prioritization assists organizations in concentrating on addressing the most critical risks first, streamlining risk management, and enhancing overall security posture.

Through this interconnected framework, Snyk ensures robust asset protection and optimized performance, integrating inventory management, policy enforcement, analytical insights, and risk prioritization into a cohesive, efficient system.

<table data-full-width="false"><thead><tr><th></th><th>GitHub</th><th width="93.44140625">GitLab</th><th>BitBucket Cloud</th><th>BitBucket Server</th><th>Azure DevOps</th></tr></thead><tbody><tr><td>YXBi6VWXCKUG</td><td>n6geam4AQFbK</td><td>Wc6ibd1G8yaa</td><td>V9dfUtedUEmT</td><td>eB4OwAO2X6xx</td><td>DuK2FGF2fwgW</td></tr><tr><td>KIQORXVeKUlp</td><td>zilPP4HcuQjm</td><td>nCYibm74OqGE</td><td>Nou1J0ga9BcK</td><td>ekUWLlyAiRX9</td><td>L86XKPBUwU5X</td></tr><tr><td>VEyakITebTdY</td><td>gXTuErHqYSrN</td><td>Ng6fxyegq6AD</td><td>9e7DrWqhF6Lj</td><td>59WhRdqTxoSK</td><td>z2I1RSObplB7</td></tr><tr><td>VoLgUsYuuW43</td><td>Pwn2788JKjKg</td><td>UNrs8f0jThb4</td><td>ZxAZeSvLyUh3</td><td>cHHCoXG4Hy3J</td><td>RFUKTxvhTpBt</td></tr><tr><td>RykPXwKL5Kdv</td><td>cUbV7yfJudDn</td><td>LQiBAJu7N2xz</td><td>5sDQM2yhEylM</td><td>b1IQEDd7aYIc</td><td>GpZYselb3DvY</td></tr><tr><td>S03xzAX0lC50</td><td>ZfVGIoCMhVTS</td><td>iD8MYJkIYV2z</td><td>8F3K5p40mRFG</td><td>sVn3QCGGckiV</td><td>OES8aybw3K1v</td></tr><tr><td>HpgnwJF2fqDd</td><td>j7FUq9HwOj6p</td><td>rZKYxIn2cfWF</td><td>AlKaPZEdkTXb</td><td>1qg3pp9wrrYn</td><td>QIr9ZFdtmHMZ</td></tr><tr><td>Su35XAIFBmWR</td><td>qMi5bL6lCqi6</td><td>kUto1utr2CZs</td><td>9d4iRvje2lKN</td><td>PrJgi9bVpvYL</td><td>2bBYR2ixrEym</td></tr></tbody></table>

### Assets Enrichments by SCM Integration <a href="#assets-enrichments-by-scm-integration" id="assets-enrichments-by-scm-integration"></a>

The table below outlines the asset enrichments provided by each SCM Integration. It highlights which capabilities are available today that are not available due to the SCM provider.

| Organization/Workspace | ✅​                                        | ✅​             | ✅​                     | ✅​ | ✅​ |
| ---------------------- | ----------------------------------------- | -------------- | ---------------------- | -- | -- |
| SCM Projects           | ❌                                         | ❌              | ✅                      | ✅  | ✅  |
| Contributors           | ✅                                         | ❌              | ❌                      | ✅  | ✅  |
| Teams                  | ✅                                         | ✅              | ❌                      | ✅  | ✅  |
| Languages (tags)       | ✅                                         | ✅              | ✅ When manually set up | ❌  | ✅  |
| Tags                   | ✅GitHub topics / GitHub custom properties | ✅GitLab topics | ❌                      | ❌  | ❌  |
| Visibility             | ✅                                         | ✅              | ✅                      | ✅  | ✅  |
| Archive Repos          | ✅                                         | ✅              | ✅                      | ✅  | ✅  |
