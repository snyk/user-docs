# Remediation reports

The Remediation reports section includes the following reports:

* [Issues Summary](remediation-reports.md#issues-summary-report) report
* [SLA Management](remediation-reports.md#sla-management-report) report
* [Vulnerabilities Detail ](remediation-reports.md#vulnerabilities-detail-report)report
* [Zero-Day](remediation-reports.md#zero-day-report) report

## Issues Summary report

The Issues Summary report highlights the value that Snyk is providing by enabling both the identification and resolution of issues.

The report provides a glimpse into how well teams are optimizing the use of the Snyk platform for their workflow and provides a means to measure and improve security.

This report enables you to easily understand the current state and trends of the highest security risk items. This report also provides a quick view into where risk is coming from and where remediation efforts are most and least effective.

{% hint style="info" %}
Use the date filter in the upper right corner of the Issues Summary report to see key metrics and charts for a specified interval. The selected date range also impacts the compared period, which allows you to measure ‌progress across various key metrics.
{% endhint %}

At the top of the report, you can follow key metrics associated with security issues in the selected date range with a comparison to the previous sequential period's results. This allows you to get insights on trends. See the tooltips in Snyk Web UI for definitions of the metrics.

The **Issues Identified and Resolved** trend captures the accumulated security issues that were identified and resolved during the selected date range. The gap between the two lines indicates the open issues backlog.

This visual trend allows you to identify if too many issues are being introduced, meaning that prevention should become a higher priority. Conversely, if not enough issues are being resolved, it means that you need to further analyze metrics such as MTTR and SLA.

{% hint style="info" %}
The Total Open issues metric at the top completes the picture for this trend, by showing the total open issues at the end of the selected period compared with the total open issues at the beginning of the selected date range.
{% endhint %}

Reviewing the **Exposure Window** trend allows you to identify the capacity of security issues that are open within predefined periods. This is a relevant metric to follow when filtering by attributes such as severity, exploit maturity, or asset class. and ensuring that the most critical issues for sensitive assets are being remediated on time.

The **Time to Resolve by Week** trend provides visibility on the number of issues remediated within predefined periods, allowing you to measure remediation performance over time.

The **Risk breakdown** table helps you make data-driven decisions about where you need to focus. The tables allow you to review ‌performance metrics from several angles.

Use the dimension picker to browse:

* **Projects** - Available at the Organization level. Allows you to pinpoint Projects that require your attention.
* **Organizations** - Available at the Group level. Surface Snyk Organizations based on their performance.
* **Asset Classes** - Ensure that efforts are prioritized to secure the most sensitive assets first.
* **Introduction Categories** - Allows to determining if preventable issues are handled properly by looking at the percentage change of new preventable issues, as well as assessing the impact of new monitored assets on your AppSec Program. You can view this under the Baseline Issue category.

## SLA Management report

The report presents default SLA targets for each severity level based on common security standards, such as FedRAMP. These SLA targets can be modified to meet your own security requirements.

The SLA status of an issue can be:

* **Within SLA** - the age of the issue has not exceeded the SLA target, and it is expected to have sufficient lead time before breaching.
* **At Risk** - the issue is considered to be approaching an SLA breach and is flagged as “At Risk”.
* **Breached** - the age of the issue has exceeded the SLA target.

You can control the SLA targets and the transition of issues to the “At Risk” status by editing the **SLA target** and setting the **At risk duration before breach (days)** field.

<figure><img src="../../../.gitbook/assets/image (229).png" alt=""><figcaption><p>SLA Management Report - Edit SLA targets</p></figcaption></figure>

The SLA report includes additional filters under the SLA category, allowing for better identification of the age of issues in relation to the SLA target:

* **SLA status** - allows the filtering of the report according to a specific SLA status.
* **Issue age** - allows discovery of issues in a range of ages.
* **Time until breach** - identifies issues that will breach the SLA target in days.

{% hint style="info" %}
The report is, by default, showing only issues that are with high or critical severity. Update the severity filter if you want to view the SLA status for additional severities.
{% endhint %}

You can share the report with predefined SLA targets by sharing the report URL or return to a predefined SLA report by bookmarking the web page in your browser.

In the **Open issues** section, the **SLA severity breakdown** shows a distribution of severity levels by the SLA compliance status of the viewed Group or Organization.

The **SLA trend** shows the cumulative SLA status of issues over time.

The **SLA breakdown table** allows you to compare the SLA compliance results of Organizations in the Group view, or Targets in the Organization view. The table is sorted by default according to the quantity of breached issues.

The **Breached and at-risk open issues** table helps you prioritize issues based on their aging and SLA compliance status. You can use the **Modify Column** picker to add additional columns and learn more about the specific issues.

{% hint style="info" %}
You can download the **SLA Breakdown** and the **Breached and at risk open issues data** in a CSV format using the **Download CSV** option.
{% endhint %}

You can review the SLA results for resolved issues and perform a retrospective analysis by reviewing the **Resolved issues** section.

## Vulnerabilities Detail report

The Vulnerabilities Detail report is similar to the Issues Detail report but shows issues grouped by Snyk Problem ID ([see Snyk Vulnerability DB](https://security.snyk.io/vuln)).

You can easily see how many instances of a vulnerability exist and how many Projects are affected. Use this report to understand which vulnerabilities are most prevalent for both resolution and prevention use cases.

For a table of Total Issues, use Change Reports to switch to the Issues Detail report.

{% hint style="info" %}
**Dependencies and license information**

To view Dependencies and license information, select the **Dependencies** menu option. See [Dependencies and licenses](../../dependencies-and-licenses/) for details.
{% endhint %}

## Zero-Day report

This report addresses primary scenarios for managing and resolving emerging zero-day vulnerabilities, which carry significant consequences and attract substantial attention in the global AppSec community.

Use this report to discover your exposure to issues highlighted in a zero-day publication across various Targets and Projects. The report helps you prioritize zero-day issues and monitor the progress of remediation efforts against any remaining occurrences.

The [Security team at Snyk](https://snyk.io/platform/security-intelligence/) continuously updates the [Vulnerability Database](https://security.snyk.io/) with new vulnerabilities several times a day. When the team discovers a major new zero-day vulnerability—typically in a widely used package with high severity that affects many customers—it will be announced and addressed as a zero-day event.

Upon the announcement of a new zero-day event, begin by examining the **Impacted Targets** table to gain a deeper understanding of the exposure. Use filters such as Project Lifecycle, Environment, or Project Criticality to focus solely on Targets associated with Projects in production that are externally exposed or of high criticality. Gaining such insights depends on the [availability of Project attributes](../../../snyk-platform-administration/snyk-projects/project-attributes.md#available-attributes-and-their-values).

Next, proceed to the **All** **Issues** table and compile a prioritized list of issues requiring remediation. Typically, prioritization is determined by either the Snyk [Risk Score](../../prioritize-issues-for-fixing/risk-score.md) or NVD CVSS Score, with emphasis placed on addressing vulnerabilities within sensitive targets. Apply filters based on Project Lifecycle, Environment, or Project Criticality to identify and address these targets promptly.

For continuous monitoring of remediation progress and efficacy, refer to the trend diagrams.\
The **Accumulative Issues Backlog Trend** diagram shows the weekly changes in the zero-day backlog by accumulating the weekly delta between identified and resolved issues. Use this diagram to ensure that your R\&D teams are reducing the zero-day backlog consistently, which will be indicated by a negative trend line.

In parallel, review the **Issues Identified versus Resolved over Time** diagram to conclude whether additional emphasis should be placed on preventing the introduction of new issues or on accelerating the remediation efforts.
