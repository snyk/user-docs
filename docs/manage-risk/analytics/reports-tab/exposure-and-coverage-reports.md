# Exposure and coverage reports

The Exposure and coverage reports section includes the following reports:

* [Asset Dashboard](exposure-and-coverage-reports.md#asset-dashboard) report
* [Issues Detail](exposure-and-coverage-reports.md#issues-detail-report) report
* [Risk Exposure](exposure-and-coverage-reports.md#risk-exposure-report) report

## Asset Dashboard

The Asset Dashboard provides a comprehensive overview of your application and security controls. It displays essential data such as the status and trends of open issues, control coverage, and repository metadata.

The Asset Dashboard is a central hub for managing and reviewing assets, making tracking inventory size easier over time and understanding the interaction between different asset types.

While Snyk Inventory enables the discovery and management of your assets that should be secured, the Snyk Asset Dashboard allows you to go beyond the details and better understand the main building blocks of your inventory.\
\
The Asset Dashboard shows all the asset data that is available in your inventory and helps you answer various questions, such as:

* Does my AppSec program meet the coverage requirements for business-critical assets and strategic applications?
* Are the assets being classified properly according to their criticality?
* Do you know which repositories belong to which application or code owners? Are newly introduced repositories being updated with that data?
* What are the main programming languages and package managers that are used in repositories that have been worked on recently?

### Filters

The filters are located at the top left of the page, with the following filtering options: **Asset Class**, **Asset type,** **Add filter**. The filter selection applies to all available data widgets.

The available filters are:

| Filter               | Description                                                                                                                                                                                                                                                                                                                                                           |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Asset Class          | The business criticality of an asset (A - most critical to D - least critical).                                                                                                                                                                                                                                                                                       |
| Asset type           | The type of an asset (Container image, Package, Repository). Most data widgets already present certain asset types by default.                                                                                                                                                                                                                                        |
| \*Application        | The list of the applications for which you have configured the application context catalog in Snyk Essentials.                                                                                                                                                                                                                                                        |
| \*Catalog name       | The name of your application context catalog.                                                                                                                                                                                                                                                                                                                         |
| \*Category           | The category of a repository asset. For example, `service` or `library`.                                                                                                                                                                                                                                                                                              |
| Discovered           | The period when the asset was discovered.                                                                                                                                                                                                                                                                                                                             |
| Last Seen            | The period when the asset was last imported from the integration.                                                                                                                                                                                                                                                                                                     |
| \*Lifecycle          | The lifecycle state of the application context catalog component. For example `production`, `experimental`, `deprecated`.                                                                                                                                                                                                                                             |
| \*Owner              | The team that owns the repository for which the application context catalog was configured.                                                                                                                                                                                                                                                                           |
| Repository Freshness | <p>The last commit date in the repository:</p><ul><li><strong>Active</strong>: Had commits in the last 3 months.</li><li><strong>Inactive</strong>: The last commits were made in the last 3 - 6 months.</li><li><strong>Dormant</strong>: No commits in the last 6 months.</li><li><strong>N/A</strong>: There are no commits detected by Snyk Essentials.</li></ul> |
| Source               | The integration that imported the asset.                                                                                                                                                                                                                                                                                                                              |
| Tags                 | The asset tags. For more details about tagging assets using a policy, see the [Tagging policy](../../policies/assets-policies/use-cases-for-policies/tagging-policy.md) page.                                                                                                                                                                                         |
| \*Title              | The name of the component for which the application context catalog was configured.                                                                                                                                                                                                                                                                                   |

**\***&#x41;ll filters marked with `*` are visible only if you configured the [application context](../../../developer-tools/scm-integrations/application-context-for-scm-integrations/) catalog for your SCM integrations.

### Repository coverage widget

The repository coverage widget provides an overview of the percentage of scanned repositories compared to the total number of available repositories, using integrated Snyk or third-party security products.

Hover over any column to see how the coverage percentage is calculated.

<figure><img src="../../../.gitbook/assets/image (204).png" alt=""><figcaption><p>Repository Coverage</p></figcaption></figure>

### Asset class breakdown

The asset class breakdown widget surfaces the distribution of repositories and container images by [asset class](../../../manage-assets/assets-inventory-components.md#class). Reviewing this widget allows you to determine the percentage of business-critical assets in your inventory and drill down to see the actual assets.

{% hint style="info" %}
**Tips**

* Having the context of the asset class is crucial for prioritizing assets. It is recommended to categorize your inventory by implementing [classification policies](../../policies/assets-policies/use-cases-for-policies/classification-policy.md) to proactively classify existing and newly introduced assets.
* Using the filters enables narrowing down the asset class distribution within specific applications or code owners, as well as focusing on active repositories or a set of assets based on the asset tags.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (203).png" alt=""><figcaption><p>Asset Class Breakdown</p></figcaption></figure>

### Top 10 technologies breakdown

The top 10 technologies widget identifies the leading programming languages and frameworks used in repositories. Using the available filters enables you to determine the most commonly used technologies in active or business-critical repositories. Moreover, you can investigate specific applications or code owners.

{% hint style="info" %}
**Tips**

* The technology data is available in the [asset tags](../../../manage-assets/assets-inventory-components.md#tags).
* Click a presented technology to open the inventory page in a new browser tab. This will allow you to review the related repositories in detail.
{% endhint %}

### Top 10 package managers breakdown

The top 10 package managers widget allows you to identify the leading package managers in your inventory. The quantities represent assets of package type. A [package asset](../../../manage-assets/assets-inventory-layouts.md#packages) is defined as software or library that is managed by package management systems.

### Repository freshness

The repository freshness widget displays the distribution of repositories according to the last commit date:

* **Active**: Had commits in the last 3 months.
* **Inactive**: The last commits were made in the last 3 - 6 months.
* **Dormant**: No commits in the last 6 months.
* **N/A**: Commits data is unavailable.

You can use this widget to surface the quantity of repositories that are more or less maintained in various contexts, such as specific applications.

{% hint style="info" %}
**Tips**

You can use the asset class filter to identify business-critical assets that are not being maintained. Click a specific slice to open the inventory page in a new browser tab where you can browse and learn more about those assets.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (202).png" alt=""><figcaption><p>Repository freshness</p></figcaption></figure>

### Application context availability

The application context availability widget allows you to discover gaps in the context of assets. The available columns include:

* **Application Context** - displays the analyzed context attribute.
* **Unique Values** - shows how many unique instances exist for an attribute. For example, you can check how many unique applications or code owners are available for any of the listed attributes.
* **Availability in Repos** - indicates the completeness of a certain attribute across the repositories.

{% hint style="info" %}
**Tips**

* Before reviewing this widget, ensure that the results are cleaned up by filtering out the "dummy" attribute values, such as "unknown", "-", and so on.\
  You can clean up the values by selecting only the relevant values.
* Filtering by asset class allows you to identify business-critical repositories without a known code owner or associated application.
* Filtering by the "active" value of the repository freshness filter allows you to discover context gaps in repositories that are actively being developed.
* Reviewing the unique values allows you to spot gaps in context. For example, you may realize that the number of unique code owners does not match the number of teams.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (199).png" alt=""><figcaption><p>Application Context Availability</p></figcaption></figure>

### Asset source breakdown

The asset source breakdown widget visualizes the quantities of detected assets from various sources. A source can be a platform where the asset is being managed directly (such as an SCM, container registry, and so on) or a platform that enriches the assets (such as security products and ASTs).

{% hint style="info" %}
**Tips**

* The widget displays the net quantities of detected assets for each source. If an asset is detected in more than one source, it will be counted once for each detected source.
* When asset inventory quantities seem incomplete or exceed expectations, this widget will help you discover which integrations should be examined and potentially configured differently.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (200).png" alt=""><figcaption><p>Asset source breakdown</p></figcaption></figure>

## Issues Detail report

The Issues Detail report displays all known issues in all of your Projects that are being monitored by Snyk. The report provides issue details and which of your Projects are affected, providing links to fix information.

The Issues Detail report displays the number of issues as well as the number of unique vulnerabilities that make up the issues.

Quick aggregations are available by categories including **Severity**, **Product Name**, and **Issue Type.**

Individual issues are displayed in a table according to the selected category. You can modify columns as needed.

For a table of only the unique vulnerabilities, use **Change Report** to switch to the Vulnerabilities Detail report.

## Risk exposure report

This report gives you a single, consolidated view of your security risks. It allows you to quickly understand your risk exposure, track your progress in reducing it, and pinpoint high-risk areas.

The Risk Exposure Report helps AppSec teams make quicker, more informed decisions. Rather than reviewing multiple reports, it provides a clear overview of the security landscape, allowing you to:

* Make faster decisions by quickly identifying your biggest security challenges and where to focus your attention.
* Prioritize effectively by using data to guide your mitigation efforts toward the areas that contribute the most risk.
* Show progress by tracking the impact of your team on reducing risk over time with easy-to-understand visualizations.

### Severity source

![](https://lh7-rt.googleusercontent.com/docsz/AD_4nXdufjGtE0kED7zHIl_L4jGrLbWkgeFfzbNwzEISsiINoEyWo2mQSnJxEBrzRca5bD1QCz-u60m-CQvDHVC-lx4gYd4LvsDrtTUkMcl6ff8V2q4uc5lUi1S8zAieM5s36JNVFbLU-Q?key=Dqdjzf6y3TJS6QA9IfBneg)

Choose your preferred severity source and automatically update selected severity throughout the report:

* **Snyk**: utilizing Snyk proprietary CVSS calculations and other factors, including the relative importance of the Linux distributor.
* **NVD CVSS**: leveraging severity scores from the National Vulnerability Database (NVD).
* **Non-SCA Severities:** For non-SCA issues (for example, Code, IaC), Snyk severity calculates High, Medium, and Low levels for specific code vulnerabilities and makes use of the Common Configuration Scoring System (CCSS) for IaC severity determinations

The report includes two main sections to provide a comprehensive view of your risk landscape:

### **Risk exposure trends**

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXd5HXXMoOzL2GsdBqF8tdO17PhaHx-1GdOdjVLAKpg46xqSMl1ooJB_KoaOkZb61O6Pu44KVI5hYkrn08aLiYfvKbIu0vZIraPlI1t44JcZP49KGbyYczwgn-jbXObBOmx-b_XF?key=Dqdjzf6y3TJS6QA9IfBneg" alt=""><figcaption></figcaption></figure>

This section provides a visual overview of your issues over time. You can view these trends by:

* **Severity**: See the distribution of issues across different severity levels.
* **Introduction Category**: Understand how issues are being introduced into your Projects.
* **Asset Class**: Group issues by the type of asset they affect.
* **Snyk Product**: Group issues by scanner.
* **Exploit Maturity:** Understand which issues have mature exploits.
* **Top 10 Orgs:** View risk by the Top 10 Organizations with the most open issues.

### **Risk exposure breakdown**

This detailed table breaks down issues and impacted assets. You can dynamically group the data to fit your needs by selecting from the following dimensions: Group, Organization, Project, introduction category, and asset class.

The table is sorted by default to surface the total number of critical and high-severity issues, helping you focus on the most urgent risks first.

You can also export data to PDF or CSV and drill down into issues for more detail.
