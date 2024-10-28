# Manage assets

Snyk defines an asset as an identifiable entity that is part of an application and relevant to security and developers. Snyk is generally focused on the development stages of application software, secures repository assets containing software package assets, and builds artifacts like container image assets.

## Inventory overview

The Inventory page in the Snyk Web UI is structured into Inventory layouts and provides information about your assets. Here is a list with all the available inventory layouts:

* **All Assets**: All the discovered assets, grouped by their type.&#x20;
* **Asset Hierarchy**: Asset Hierarchy layout shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. Assets hierarchy is visible only when there are no filters applied.
* **Teams**: SCM repository assets grouped by teams. Only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk AppRisk.

Each inventory layout may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory layout.

The following video presents an overview of the Inventory view from the Snyk Web UI.

{% embed url="https://youtu.be/BIvGd_TCwHw" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)!
{% endembed %}

{% hint style="warning" %}
**Release status** \
**Risk factors** on assets are in Early Access and available only for Snyk AppRisk Pro.

**Runtime discovered** and **Runtime last seen** filters take their release status from the [runtime integration](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md) for which they provide data and are available only for Snyk AppRisk Pro.
{% endhint %}

## Assets overview

In the Snyk AppRisk context, you can find information about assets under the following menus: [Inventory](./), [Policies](../manage-risk/policies/assets-policies/), [Insights for Snyk AppRisk](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-apprisk.md), and [Analytics](../manage-risk/enterprise-analytics/application-analytics.md).&#x20;

The main source of information is the **Inventory** menu. Here, you can view, filter, and manage your assets.&#x20;

The [**Inventory**](./) menu provides a centralized view of all the issues identified by Snyk with additional asset context. The Inventory serves as the cornerstone of this interconnected system, providing detailed information on each asset to facilitate further processes.

In the [**Policies**](../manage-risk/policies/assets-policies/) menu, you can create policies for organizing the assets, classifying them, and always being up to date with the latest information about an asset. The policies are put in place to regulate these assets, making sure that each one follows security standards, organizational best practices, and regulatory requirements. These policies are enforced using the data from the Inventory, ensuring consistent compliance across all assets.

In the [**Analytics**](../manage-risk/enterprise-analytics/application-analytics.md) menu, you can review and explore your application status and results, from a top-down approach. The Analytics page then uses this data to provide detailed insights into the performance and security metrics of the assets. This page helps identify trends, detect anomalies, and pinpoint potential vulnerabilities, providing a data-driven approach to asset management.

The [**Issues**](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-apprisk.md) menu utilizes information from both the Inventory and Policies menus to prioritize issues based on their severity and impact. This prioritization assists organizations in concentrating on addressing the most critical risks first, streamlining risk management, and enhancing overall security posture.

Through this interconnected framework, Snyk ensures robust asset protection and optimized performance, integrating inventory management, policy enforcement, analytical insights, and risk prioritization into a cohesive, efficient system.



