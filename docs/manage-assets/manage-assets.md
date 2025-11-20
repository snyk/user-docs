# Overview

Snyk defines an asset as an identifiable entity that is part of an application and relevant to security and developers. Snyk is generally focused on the development stages of application software, secures repository assets containing software package assets, and builds artifacts like container image assets.

## Inventory menu

Use the **Inventory** to organize your repository assets, visualize SCM tool assets, track Snyk coverage, and prioritize mitigation based on business impact.

The **Inventory** page in the Snyk Web UI is structured into multiple tabs and provides information about your assets.

Here is a list of all the available inventory tabs:

* **Overview:** Provides quick insights into discovered repositories, enabling AppSec teams to effectively operationalize their program using Snyk.
* **All Assets**: All the discovered assets, grouped by their type.
* **Asset Hierarchy**: All the assets in a hierarchical structure. The list of assets is sorted by issue counts, and where applicable, the package assets are listed underneath the repositories where they are located. The Asset hierarchy is visible only when there are no filters applied.
* **Teams**: SCM repository assets grouped by teams. Only SCM organizations with teams and repositories assigned to a team appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk Essentials.

Each inventory tab may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory layout.

The following video summarizes the main features of the Inventory view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1746127384/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-4c_-_v2_-_Inventory_Overview.mp4" %}
Reviewing asset inventory
{% endembed %}

{% hint style="info" %}
**Release status**\
Risk factors for Snyk AppRisk are in Early Access and are available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.

**Runtime discovered** and **Runtime last seen** filters take their release status from the [runtime integration](../integrations/connect-a-third-party-integration.md) for which they provide data and are available only for Snyk AppRisk.
{% endhint %}

## Assets overview

In the Snyk Enterprise context, you can find information about assets under the following menus: [Inventory](manage-assets.md), [Policies](../manage-risk/policies/assets-policies/), [Insights for Snyk AppRisk](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md), and [Analytics](broken-reference).

The main source of information is the **Inventory** menu. Here, you can view, filter, and manage your assets.

The [**Inventory**](manage-assets.md#inventory-overview) menu provides a centralized view of all the issues identified by Snyk with additional asset context. The Inventory serves as the cornerstone of this interconnected system, providing detailed information on each asset to facilitate further processes.

In the [**Policies**](../manage-risk/policies/assets-policies/) menu, you can create policies for organizing the assets, classifying them, and always being up to date with the latest information about an asset. The policies are put in place to regulate these assets, making sure that each one follows security standards, organizational best practices, and regulatory requirements. These policies are enforced using the data from the Inventory, ensuring consistent compliance across all assets.

In the [**Analytics**](broken-reference) menu, you can review and explore your application status and results, from a top-down approach. The Analytics page then uses this data to provide detailed insights into the performance and security metrics for the assets. This page helps identify trends, detect anomalies, and pinpoint potential vulnerabilities, providing a data-driven approach to asset management.

The [**Issues**](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md) menu uses information from both the Inventory and Policies menus to prioritize issues based on their severity and impact. This prioritization assists organizations in concentrating on addressing the most critical risks first, streamlining risk management, and enhancing overall security posture.

Through this interconnected framework, Snyk ensures robust asset protection and optimized performance, integrating inventory management, policy enforcement, analytical insights, and risk prioritization into a cohesive, efficient system.

### Assets Enrichments by SCM Integration <a href="#assets-enrichments-by-scm-integration" id="assets-enrichments-by-scm-integration"></a>

The table below outlines the asset enrichments provided by each SCM Integration. It highlights which capabilities are available today that are not available due to the SCM provider.

<table data-full-width="false"><thead><tr><th width="212.5364990234375">Capability</th><th>GitHub</th><th>GitLab</th><th>BitBucket Cloud</th><th>BitBucket Server</th><th>Azure DevOps</th></tr></thead><tbody><tr><td>Organization/Workspace</td><td>✅​</td><td>✅​</td><td>✅​</td><td>✅​</td><td>✅​</td></tr><tr><td>SCM Projects</td><td>❌</td><td>❌</td><td>✅</td><td>✅</td><td>✅</td></tr><tr><td>Contributors</td><td>✅</td><td>❌</td><td>❌</td><td>✅</td><td>✅</td></tr><tr><td>Teams</td><td>✅</td><td>✅</td><td>❌</td><td>✅</td><td>✅</td></tr><tr><td>Languages (tags)</td><td>✅</td><td>✅</td><td>✅ When manually set up</td><td>❌</td><td>✅</td></tr><tr><td>Tags</td><td>✅ GitHub topics / GitHub custom properties</td><td>✅ GitLab topics</td><td>❌</td><td>❌</td><td>❌</td></tr><tr><td>Visibility</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td></tr><tr><td>Archive Repos</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td></tr></tbody></table>
