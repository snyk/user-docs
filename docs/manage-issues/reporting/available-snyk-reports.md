# Available Snyk Reports

The following reports are available:

* [Issues Detail report](available-snyk-reports.md#issues-detail-report)
* [Issues Summary report](available-snyk-reports.md#issues-summary-report)
* [Vulnerabilities Detail report](available-snyk-reports.md#vulnerabilities-detail-report)
* [Featured Zero-Day report](available-snyk-reports.md#featured-zero-day-report)
* [SLA Management report](available-snyk-reports.md#sla-management-report)
* [OWASP TOP 10 report](available-snyk-reports.md#owasp-top-10-report)
* [CWE TOP 25 report](available-snyk-reports.md#cwe-top-25-report)
* [CWE TOP 10 KEV report](available-snyk-reports.md#cwe-top-10-kev-report)
* [Developer IDE and CLI usage report](available-snyk-reports.md#developer-ide-and-cli-usage)
* [Cloud Compliance Issues report](available-snyk-reports.md#cloud-compliance-issues-report)
* [Snyk Generated Pull Requests](available-snyk-reports.md#snyk-generated-pull-requests)
* [Asset Dashboard](available-snyk-reports.md#asset-dashboard)

Select **Change Report** to change the report displayed:

<div align="left"><figure><img src="../../.gitbook/assets/select-report.png" alt="Select Change Report to display different reports" width="563"><figcaption><p>Select Change Report to display different reports</p></figcaption></figure></div>

## Issues Detail report

The Issues Detail report displays all known issues in all of your Projects that are being monitored by Snyk. The report gives details about each issue and which of your Projects are affected and provides links to fix information.

The Issues Detail report displays the number of issues as well as the number of unique vulnerabilities that make up the issues.

Quick aggregations are available by categories including **Severity**, **Product Name**, and **Issue Type.**

Individual issues are displayed in a table according to the selected category. You can modify columns as needed.

For a table of only the unique vulnerabilities, use Change Report to switch to the Vulnerabilities Detail report.

## Issues Summary report

The Issues Summary report highlights the value that Snyk is providing by enabling both the identification and resolution of issues.&#x20;

The report provides a glimpse into how well teams are optimizing the use of the Snyk platform for their workflow and provides a means to measure and improve security.

This report enables you to easily understand the current state and trends of the highest security risk items. This report also provides a quick view into where risk is coming from and where remediation efforts are most and least effective.

Use the date filter in the upper right corner of the Issues Summary report to see key metrics and charts for a specified interval.

This report shows a number of key metrics associated with issues in that interval with a comparison to the same metrics in the previous period so you can get a quick understanding of trends. See tooltips in the app for definitions of the metrics.

Scroll down for additional charts that show trend information in greater detail.

Key metrics are then broken down to point out information at the Organization or Project level. You can drill down to see what new and resolved issues were introduced during the date range selected.

## Vulnerabilities Detail report

The Vulnerabilities Detail report is similar to the Issues Detail report but shows issues grouped by Snyk Problem ID ([see Snyk Vulnerability DB](https://security.snyk.io/vuln)).&#x20;

You can easily see how many instances of a vulnerability exist and how many Projects are affected. Use this report to understand which vulnerabilities are most prevalent for both resolution and prevention use cases.

For a table of Total Issues, use Change Reports to switch to the Issues Detail report.

{% hint style="info" %}
#### Dependencies and license information

To view Dependencies and license information, select the **Dependencies** menu option. See [Dependencies and licenses](../../manage-risk/reporting/dependencies-and-licenses/) for details.
{% endhint %}

## Featured Zero-Day report

This report addresses primary scenarios for managing and resolving emerging zero-day vulnerabilities, which carry significant consequences and attract substantial attention in the global AppSec community.

Use this report to discover your exposure to issues highlighted in a zero-day publication across various Targets and Projects. The report helps you prioritize zero-day issues and monitor the progress of remediation efforts against any remaining occurrences.

The [Security team at Snyk](https://snyk.io/platform/security-intelligence/) continuously updates the [Vulnerability Database](https://security.snyk.io/) with new vulnerabilities several times a day. When the team discovers a major new zero-day vulnerability—typically in a widely used package with high severity that affects many customers—it will be announced and addressed as a zero-day event.

Upon the announcement of a new zero-day event, begin by examining the **Impacted Targets** table to gain a deeper understanding of the exposure. Use filters such as Project Lifecycle, Environment, or Project Criticality to focus solely on Targets associated with Projects in production that are externally exposed or of high criticality. Gaining such insights depends on the [availability of Project attributes](../../snyk-admin/snyk-projects/project-attributes.md#available-attributes-and-their-values).

Next, proceed to the **All** **Issues** table and compile a prioritized list of issues requiring remediation. Typically, prioritization is determined by either the Snyk [Risk Score](../../manage-risk/prioritize-issues-for-fixing/risk-score.md) or NVD CVSS Score, with emphasis placed on addressing vulnerabilities within sensitive targets. Apply filters based on Project Lifecycle, Environment, or Project Criticality to identify and address these targets promptly.

For continuous monitoring of remediation progress and efficacy, refer to the trend diagrams.\
The **Accumulative Issues Backlog Trend** diagram shows the weekly changes in the zero-day backlog by accumulating the weekly delta between identified and resolved issues. Use this diagram to ensure that your R\&D teams are reducing the zero-day backlog consistently, which will be indicated by a negative trend line.

In parallel, review the **Issues Identified versus Resolved over Time** diagram to conclude whether additional emphasis should be placed on preventing the introduction of new issues or on accelerating the remediation efforts.

## SLA Management report&#x20;

The report presents a set of default SLA targets per severity based on common security standards, such as FedRAMP. These SLA targets can be modified to meet your own security requirements.

The SLA status of an issue can be:

* **Within SLA** - the age of the issue has not exceeded the SLA target, and it is expected to have sufficient lead time before breaching.
* **At Risk** - the issue is considered to be approaching an SLA breach and is flagged as “At Risk”.
* **Breached** - the age of the issue has exceeded the SLA target.

You can control the SLA targets and the transition of issues to the “At Risk” status by editing the **SLA target** and setting the **At risk duration before breach (days)** field.

<figure><img src="../../.gitbook/assets/image (639).png" alt=""><figcaption><p>SLA Management Report - Edit SLA targets</p></figcaption></figure>

The SLA report includes additional filters under the SLA category, allowing for better identification of the age of issues in relation to the SLA target:

* **SLA status** - allows the filtering of the report according to a specific SLA status.
* **Issue age** - allows to discover issues within a range of age.
* **Time until breach** - identifies issues that will breach the SLA target within days.

{% hint style="info" %}
The report is, by default, showing only issues that are with high or critical severity. Update the severity filter if you want to view the SLA status for additional severities.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (634).png" alt="" width="300"><figcaption><p>SLA Filters within the filters picker </p></figcaption></figure>

You can share the report with predefined SLA targets by sharing the report URL or return to a predefined SLA report by bookmarking the web page in your browser.

In the **Open issues** section, the **SLA severity breakdown** shows a distribution of severity levels by the SLA compliance status of the viewed Group or Organization.&#x20;

The **SLA trend** shows the cumulative SLA status of issues over time.

<figure><img src="../../.gitbook/assets/image (635).png" alt=""><figcaption><p>SLA Management Report - Open issues section</p></figcaption></figure>

The **SLA breakdown table** allows you to compare the SLA compliance results of Organizations in the Group view, or Targets in the Organization view. The table is sorted by default according to the quantity of breached issues.

<figure><img src="../../.gitbook/assets/image (636).png" alt=""><figcaption><p>SLA Management Report - SLA Breakdown</p></figcaption></figure>

The **Breached and at-risk open issues** table helps you prioritize issues based on their aging and SLA compliance status. You can use the **Modify Column** picker to add additional columns and learn more about the specific issues.

<figure><img src="../../.gitbook/assets/image (637).png" alt=""><figcaption><p>SLA Management Report - Breached and at risk open issues</p></figcaption></figure>

{% hint style="info" %}
You can download the **SLA Breakdown** and the **Breached and at risk open issues data** in a CSV format using the **Download CSV** option.
{% endhint %}

You can review the SLA results for resolved issues and perform a retrospective analysis by reviewing the **Resolved issues** section.

<figure><img src="../../.gitbook/assets/image (638).png" alt=""><figcaption><p>Resolved issues section</p></figcaption></figure>

## OWASP Top 10 report

The [OWASP Top 10](https://owasp.org/www-project-top-ten/) is a standard awareness document for developers and web application security. It represents a broad consensus about the most critical security risks for web applications and is globally recognized by developers as the first step towards more secure coding.&#x20;

Each control in the list (A1, A2, and so on) is based on a list of Common Weakness Enumerations (CWEs). For example, [A01:2021 – Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/) is based on a list of 34 CWEs.&#x20;

The CWEs are mapped to Snyk-IDs (), which are mapped to issues.&#x20;

For example, the critical vulnerability [SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720](https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720) is classified as [CWE-94](https://cwe.mitre.org/data/definitions/94.html), which is part of the OWASP TOP 10 [A03:2021 - Injection](https://owasp.org/Top10/A03_2021-Injection/). All the issues related to this vulnerability will be under the A03 category.

Learn more by using the [OWASP TOP 10 Learning path](https://learn.snyk.io/learning-paths/owasp-top-10/) on Snyk Learn.

The report is based on the latest mapping released in 2021. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## CWE Top 25 report

The [CWE Top 25](https://cwe.mitre.org/top25/) Most Dangerous Software Weaknesses is a list that demonstrates the current most common and impactful software weaknesses based on Common Vulnerabilities and Exposures (CVEs) severity and their exploitation potential.&#x20;

The report is based on the latest version released in 2023 by Mitre. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## CWE Top 10 KEV report&#x20;

The [CWE Top 10 KEV Weaknesses](https://cwe.mitre.org/top25/archive/2023/2023_kev_list.html) list identifies the top ten CWEs in the Cybersecurity and Infrastructure Security Agency’s (CISA) [Known Exploited Vulnerabilities](https://www.cisa.gov/known-exploited-vulnerabilities-catalog) (KEV) Catalog, a database of security flaws in software applications and weaknesses that have been exposed and leveraged by attackers.&#x20;

The report is based on the version released in 2023 by Mitre. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## Developer IDE and CLI usage

To use this report, you must ensure you have installed the following prerequisites:&#x20;

* Snyk CLI version 1.1292.1 or newer
* VS Code 1.86.0 or newer and Snyk Security plugin 2.3.3 or newer
* IntelliJ IDEs 2023.3 or newer and Snyk Security plugin 2.7.3 or newer
* Visual Studio 2019, 2022 and Snyk Security Plugin 1.1.47 or newer
* Eclipse 2023.12 or newer and Snyk Security plugin 2.1.0 or newer

This report shows the adoption of Snyk testing in local development through the IDE plugins and using the CLI locally. The report is available under the Change Report dropdown at the Group and Organization levels.

{% hint style="info" %}
This report focuses on the local developer experience and thus does not include the use of CI/CD. In addition, the report does not show organizations or developers that have never used the CLI or IDE.
{% endhint %}

Security teams can use this report to demonstrate strong shift left behavior as model behavior to bring to other teams. This report also shows where teams or individual developers are not adopting Snyk locally. Companies can use this report to encourage more shift left behavior.

This report shows the test usage in the IDE and CLI by developers:

* Total number of developers running scans and the number of scans in IDE and CLI
* Charts and summary tables breaking down this data by different dimensions, including IDE plugin
* List of organizations and developers adopting Snyk locally

Teams can filter by date and Organization.

## Cloud Compliance Issues report

{% hint style="info" %}
This report is available only if you have the legacy Snyk Cloud enabled.
{% endhint %}

The Cloud Compliance Issues report shows cloud issues for an entire Organization, organized by [compliance standard](../../scan-with-snyk/snyk-iac/getting-started-with-cloud-scans/key-concepts-for-cloud-scans.md#docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e).

You can view a report for a single version of a compliance standard at a time, for example, CIS AWS Foundations Benchmark v1.4.0, by selecting the desired standard from the drop-down menu. Each report includes a list of compliance controls organized by control category, with corresponding issue counts.

Selecting an issue count lets you view the list of issues associated with that control in the [Cloud Issues UI](../../scan-with-snyk/snyk-iac/getting-started-with-cloud-scans/manage-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md), where you can view each issue in detail.

Use the information in the Cloud Compliance Issues report to investigate, triage, and fix cloud compliance issues.

## Snyk Generated Pull Requests

{% hint style="info" %}
**Release status**

Snyk Generated Pull Requests report is now in Early Access and available only for Enterprise plan customers on the following SCM integrations: GitHub, GitHub Enterprise, and GitHub Cloud App. For more information, see [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

### Access the report

The Generated Pull Requests report can be accessed at both Group and Organization level from the **Change Report** drop down in the Reports menu.

<figure><img src="../../.gitbook/assets/2025-03-05_11-19-35.png" alt="Snyk generated pull requests report"><figcaption><p>Snyk generated pull requests report</p></figcaption></figure>

This report type provides an overview of how [Fix](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md), [Backlog](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md), and [Upgrade PRs](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/) are used and highlights the efficiency of PR merges.

The analytics report covers the following:

* Overview of PRs status by type and the PR merge ratio.
* Visibility of issues.
* Breakdown by repository for PR status.

The report summary enables you to check the total number of Snyk PRs created, the total pull requests merged, and the mean time to merge for those pull requests.

### Report features

Use the date filter in the upper right corner of the report to display data based on a specific interval.

Add various filters to narrow down results to specific configurations. The filter options are Organization, SCM, Project, and Repository.

#### Snyk Generated Pull Requests usage

<figure><img src="../../.gitbook/assets/2025-03-05_11-20-10.png" alt="Pull Request usage graph and table"><figcaption><p>Pull Request usage graph and table</p></figcaption></figure>

Pull Request usage is visualized in a **Pull requests by type** graph and a **Pull requests by status** table, displaying the same data in different formats. These distinguish the number of PRs into Fix, Backlog, and Dependency upgrade categories, segmented by Open, Merged, and Closed status types. Merge rate is presented as a percentage for each row.

#### Open vs Fixed issues

<figure><img src="../../.gitbook/assets/open-vs-fixed-issues-pr-report.png" alt="Open vs Fixed issues graph and table"><figcaption><p>Open vs Fixed issues graph and table </p></figcaption></figure>

The Open vs Fixed issues in Snyk PRs graph and table displays the number of open and fixed issues based on severity.

#### Snyk Generated Pull Requests by repository

<figure><img src="../../.gitbook/assets/2025-03-05_10-02-13.png" alt="Projects/Orgs/Repository table for PRs of different status"><figcaption><p>Projects/Orgs/Repository table for PRs of different status</p></figcaption></figure>

The **Projects/Org/Repository** table displays the number of Total, Open, Merged, and Closed PRs for each Organization and repository relationship. Merge rate is presented as a percentage for each row.

Select a repository name to open a modal containing additional metrics for that specific repository.

<figure><img src="../../.gitbook/assets/repo-breakdown-report-blur.png" alt="Repository breakdown by PR type and PR status"><figcaption><p>Repository breakdown by PR type and PR status</p></figcaption></figure>

The repository breakdown details the number of PRs segmented by PR type and PR status. Merge rate is presented as a percentage for each row. It also lists the Projects within that repository, with the number of issues categorised by severity.

## Asset Dashboard

The Asset Dashboard provides a comprehensive overview of your application and security controls. It displays essential data such as the status and trends of open issues, control coverage, and repository metadata.&#x20;

The Asset Dashboard is a central hub for managing and reviewing assets, making tracking inventory size easier over time and understanding the interaction between different asset types.

While Snyk Inventory enables the discovery and management of your assets that should be secured,  the Snyk Asset Dashboard allows you to go beyond the details and better understand the main building blocks of your inventory. \
\
The Asset Dashboard brings all the asset data that is available in your inventory and helps to answer various questions, such as:

* Does my AppSec program meet the coverage requirements for business-critical assets and strategic applications?
* Are the assets being classified properly according to their criticality?&#x20;
* Do you know which repositories belong to which application or code owners? Are newly introduced repositories being updated with that data?&#x20;
* What are the main programming languages and package managers that are used in repositories that have been worked on recently?&#x20;

### Filters

The filters are located at the top left of the page, with the following filtering options: **Asset Class**, **Asset type,** **Add filter**. The filter selection applies to all available data widgets.

Here are the available filters:

| Filter               | Description                                                                                                                                                                                                                                                                                                                                                            |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Asset Class          | The business criticality of an asset (A - most critical to D - least critical).                                                                                                                                                                                                                                                                                        |
| Asset type           | The type of an asset (Container image, Package, Repository). Most data widgets already present certain asset types by default.                                                                                                                                                                                                                                         |
| \*Application        | The list of the applications for which you have configured the application context catalog in Snyk Essentials.                                                                                                                                                                                                                                                         |
| \*Catalog name       | The name of your application context catalog.                                                                                                                                                                                                                                                                                                                          |
| \*Category           | The category of a repository asset. For example, `service` or `library`.                                                                                                                                                                                                                                                                                               |
| Discovered           | The period when the asset was discovered.                                                                                                                                                                                                                                                                                                                              |
| Last Seen            | The period when the asset was last imported from the integration.                                                                                                                                                                                                                                                                                                      |
| \*Lifecycle          | The lifecycle state of the application context catalog component. For example `production`, `experimental`, `deprecated`.                                                                                                                                                                                                                                              |
| \*Owner              | The team that owns the repository for which the application context catalog was configured.                                                                                                                                                                                                                                                                            |
| Repository Freshness | <p>The last commit date in the repository: </p><ul><li><strong>Active</strong>: Had commits in the last 3 months.</li><li><strong>Inactive</strong>: The last commits were made in the last 3 - 6 months.</li><li><strong>Dormant</strong>: No commits in the last 6 months.</li><li><strong>N/A</strong>: There are no commits detected by Snyk Essentials.</li></ul> |
| Source               | The integration that imported the asset.                                                                                                                                                                                                                                                                                                                               |
| Tags                 | The asset tags. For more details about tagging assets using a policy, see the [Tagging policy](../../manage-risk/policies/assets-policies/use-cases-for-policies/tagging-policy-use-case.md) page.                                                                                                                                                                     |
| \*Title              | The name of the component for which the application context catalog was configured.                                                                                                                                                                                                                                                                                    |

**\***&#x41;ll filters marked with `*` are visible only if you configured the [application context](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/) catalog for your SCM integrations.

### Repository coverage widget

The repository coverage widget provides an overview of the percentage of scanned repositories compared to the total number of available repositories, using integrated Snyk or third-party security products.

Hover over any column to see how the coverage percentage is calculated.&#x20;

<figure><img src="../../.gitbook/assets/image (585).png" alt=""><figcaption><p>Repository Coverage</p></figcaption></figure>

### Asset class breakdown

The asset class breakdown widget surfaces the distribution of repositories and container images by [asset class](../../manage-assets/assets-inventory-components.md#class). Reviewing this widget allows you to determine the percentage of business-critical assets in your inventory and drill down to see the actual assets.

{% hint style="info" %}
**Tips**

* Having the context of the asset class is crucial for prioritizing assets. It is recommended to categorize your inventory by implementing [classification policies](../../manage-risk/policies/assets-policies/use-cases-for-policies/classification-policy-use-case.md) to proactively classify existing and newly introduced assets.
* Using the filters enables narrowing down the asset class distribution within specific applications or code owners, as well as focusing on active repositories or a set of assets based on the asset tags.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (584).png" alt=""><figcaption><p>Asset Class Breakdown</p></figcaption></figure>

### Top 10 technologies breakdown

The top 10 technologies widget identifies the leading programming languages and frameworks used in repositories. Using the available filters enables you to determine the most commonly used technologies in active or business-critical repositories. Moreover, you can investigate specific applications or code owners.

{% hint style="info" %}
**Tips**

* The technology data is available in the [asset tags](../../manage-assets/assets-inventory-components.md#tags).&#x20;
* Click a presented technology to open the inventory page in a new browser tab. This will allow you to review the related repositories in detail.
{% endhint %}

### Top 10 package managers breakdown

The top 10 package managers widget allows you to identify the leading package managers in your inventory. The quantities represent assets of package type. A [package asset](../../manage-assets/assets-inventory-layouts.md#packages) is defined as software or library that is managed by package management systems.&#x20;

### Repository freshness

The repository freshness widget displays the distribution of repositories according to the last commit date:

* **Active**: Had commits in the last 3 months.
* **Inactive**: The last commits were made in the last 3 - 6 months.
* **Dormant**: No commits in the last 6 months.
* **N/A**: Commits data is unavailable.

You can use this widget to surface the quantity of repositories that are more or less maintained in various contexts, such as specific applications.&#x20;

{% hint style="info" %}
**Tips**

You can use the asset class filter to identify business-critical assets that are not being maintained. Click a specific slice to open the inventory page in a new browser tab where you can browse and learn more about those assets.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (583).png" alt=""><figcaption><p>Repository freshness</p></figcaption></figure>

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

<figure><img src="../../.gitbook/assets/image (581).png" alt=""><figcaption><p>Application Context Availability</p></figcaption></figure>

### Asset source breakdown

The asset source breakdown widget visualizes the quantities of detected assets from various sources. A source can be a platform where the asset is being managed directly (such as an SCM, container registry, and so on) or a platform that enriches the assets (such as security products and ASTs).

{% hint style="info" %}
**Tips**

* The widget displays the net quantities of detected assets for each source. If an asset is detected in more than one source, it will be counted once for each detected source.
* When asset inventory quantities seem incomplete or exceed expectations, this widget will help you discover which integrations should be examined and potentially configured differently.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (582).png" alt=""><figcaption><p>Asset source breakdown</p></figcaption></figure>
