# Redesigned Analytics

## Snyk Analytics: Redesigned Tenant-Level Experience (Early Access)

{% hint style="info" %}
The redesigned Snyk Analytics experience in Early Access is available to all customers on an Enterprise plan with group-level reporting permissions.
{% endhint %}

Learn how to build custom dashboards, access enhanced metrics, and explore Tenant-level analytics. The redesigned Snyk Analytics includes:

* Customizable analytics landing page: A new tenant-level landing page featuring a widget inventory. You can select and arrange widgets to create a personalized dashboard focused on the metrics most relevant to your needs.
* Saved views: This new experience features ‌[Saved Views](../reporting/#saved-views).
* Report catalog: A centralized catalog of Snyk Group-level reports. Use this to discover and access pre-built reports relevant to specific Groups within your Organization.
* Issues/Assets drill-down: Enhanced drill-down capabilities allow you to explore issues and assets in greater detail directly from the dashboard.

<figure><img src="../../.gitbook/assets/image (433).png" alt=""><figcaption></figcaption></figure>

## User access permissions

Data access at the Tenant level is governed by your existing Group-level permissions. You can only see data for the Groups you are authorized to view. The following governs user permissions for Tenant-level analytics:

* Any Group member with reporting permissions can now see their Groups on the Tenant-level analytics page.
* Access is restricted to those Groups the user has permissions for, based on their Group-level permissions (in both predefined and custom roles).
*   Users who previously had Group reporting permissions but no direct access to Tenant-level analytics  can now access  Tenant-level analytics for the first time. For example, users missing a feature flag or without Tenant-level roles such as Tenant Admin or Tenant Viewer.

    Users with Group reporting permissions, who couldn't previously access Tenant-level analytics, will now be able to view Tenant analytics.

## Available metrics

You can select and remove metrics from each section to tailor your view.

#### Key Performance Indicators

* **Unique vulnerabilities**: This metric shows the distinct count of unresolved vulnerabilities, highlighting the breadth of different security problems.
* **Open issues (Featured zero-day)**: Focuses on open issues specifically related to featured zero-day vulnerabilities, highlighting exposure to critical emerging threats.
* **Tested repositories in CI/CD**: Tracks the number of repositories tested within your Continuous Integration/Continuous Delivery pipeline.

#### Coverage

* **Repository coverage by Snyk product**: Indicates the percentage of repositories covered by each Snyk product.
* **Repository coverage over time, by Snyk product**: Shows the trend of repository coverage by Snyk products over time.
* **Assets discovered over time, by type**: Tracks the trend of assets discovered, categorized by their type.

#### Exposure

* **Organizations with the most open issues**: Identifies high-risk areas, guiding where to focus remediation resources.
* **Open issues over time, by severity**: Shows the trend of open issues categorized by their severity.
* **Open issues by Snyk product**: Breaks down open issues by the Snyk product that identified them.
* **Open issues by severity**: Categorizes open issues by their severity level.
* **Open issues by repository freshness**: Reveals security risks in codebases based on their last commit (freshness), indicating potential blind spots. Active repositories had a last commit less than 3 months ago, Inactive 3-6 months ago, and Dormant over 6 months ago; these are configurable. N/A indicates no commits detected by the Group SCM integration.
* **Open issues (OWASP Top 10, 2021) by severity**: Categorizes open issues based on the OWASP Top 10 (2021) and their severity.
* **Open issues by asset class**: Enables prioritization of remediation efforts based on the business criticality of different assets.
* **SCA CVEs with the most open issues**: Pinpoints critical, high-impact open-source vulnerabilities (CVEs) that have the most open issues. These are identified via SCA (Software Composition Analysis).

### Remediation

* **Issues identified and resolved over time**: Assesses the ability to keep pace with new vulnerabilities and reduce the security backlog, guiding remediation prioritization.
* **Resolved issues by MTTR duration**: Displays the typical timeframes for issue resolution, indicating overall remediation efficiency. MTTR (Mean Time To Resolution) is the average time taken to fix an issue.
* **Organizations with the highest MTTR**: Identifies organizations with the longest average time to fix issues. High MTTR highlights teams where remediation processes or resources need improvement.
* **Snyk-generated PRs merged vs. open**: Compares the number of Snyk-generated pull requests that have been merged versus those that are still open.

### Prevention

* **Snyk IDE & CLI usage over time (developer count)**: Tracks the trend of developers testing using Snyk IDE plugins and the CLI for local development.
* **Organizations introducing most new SCA preventable issues**: Highlights organizations with the highest incidence of new SCA (Software Composition Analysis) preventable issues, informing where to focus prevention efforts.
* **New issues by introduction category**: Pinpoints the primary sources of new vulnerabilities, informing where to strengthen prevention strategies.
* **New SCA preventable issues introduced over time**: Tracks the volume of new open-source vulnerabilities that could have been avoided over time. Identified through SCA, these help assess prevention effectiveness for open-source risks.

### Report catalog

The comprehensive suite of reports helps you manage and improve your application security posture. The following reports are available at the Tenant level. For more details on each report, view [available Snyk reports](../reporting/available-snyk-reports.md).

<figure><img src="../../.gitbook/assets/Screenshot 2025-07-15 at 11.22.49 AM.png" alt=""><figcaption></figcaption></figure>

### Asset dashboard

The Asset Dashboard offers a comprehensive overview of security controls for your applications. It provides insights into scan coverage, a detailed breakdown of your inventory by asset class, source, and other relevant information, and allows for extensive global filtering. This report helps visualize all repository assets, track Snyk product control coverage, and prioritize mitigation based on business impact.

#### CWE Top 10 KEV (2023)

This report focuses on monitoring risk exposure and prioritizing vulnerabilities based on CISA's Known Exploited Vulnerabilities (KEV) catalog for 2023. Snyk helps identify and address these critical vulnerabilities within your Projects.

#### CWE Top 25 (2023)

The CWE Top 25 report helps Organizations monitor risk exposure and prioritize vulnerabilities based on the 25 most dangerous Common Weakness Enumerations (CWEs) identified by the SANS Institute and MITRE. Snyk assists in identifying and remediating these prevalent software weaknesses.

#### Developer IDE and CLI Usage

This report tracks and optimizes the adoption of Snyk IDE and CLI tools by developers. It provides insights into the total number of developers running scans, the volume of scans performed, and a breakdown by IDE plugin, fostering a more effective shift-left security culture.

#### Featured Zero-Day

The Featured Zero-Day report is designed to help identify affected assets within your environment and monitor ‌remediation progress for highlighted zero-day vulnerabilities. It enables you to quickly respond to newly discovered critical vulnerabilities and reduce their impact.

#### Issues Detail

The Issues Detail report displays all known security and license issues across the Projects monitored by Snyk. It provides granular details about each issue, identifies affected Projects, and offers information related to remediation, with options for quick aggregation and filtering.

#### Issues Summary

The Issues Summary report highlights the value Snyk delivers by enabling the identification and resolution of security issues. It provides an overview of how effectively teams are utilizing the Snyk platform, offering insights into the current state and trends of high-security risk items and the efficacy of remediation efforts.

#### OWASP Top 10 (2021)

This report helps monitor risk exposure and prioritize vulnerabilities based on the OWASP Top 10 most critical web application security risks from 2021. Snyk can scan for these vulnerabilities and generate corresponding reports, aiding in compliance and risk mitigation.

#### PCI-DSS v4.0.1

The PCI-DSS v4.0.1 report helps you estimate their readiness for meeting the Payment Card Industry Data Security Standard's (PCI DSS) Application Security (AppSec) requirements for Static Application Security Testing (SAST) and Software Composition Analysis (SCA) based on Snyk scan results. It also provides essential evidence for compliance.

#### Repositories Tested in CI/CD

This report tracks the adoption of Snyk tests within your Continuous Integration/Continuous Delivery (CI/CD) pipelines to ensure that critical vulnerabilities are detected early in the development lifecycle, before deployment. It provides insights into how Snyk is integrated and used for security scanning within your CI/CD environments.

#### SLA Management

The SLA Management report allows you to view the current Service Level Agreement (SLA) status of security issues based on your defined policies. It helps pinpoint gaps in adherence and identifies specific teams that require attention to improve remediation efforts and maintain compliance goals.

#### Vulnerabilities Detail

The Vulnerabilities Detail report helps prioritize remediation efforts, developer education, and eradication campaigns by focusing on the most prevalent security issues. Similar to the Issues Detail report, it groups issues by Snyk Problem ID, enabling a clearer understanding of the number of instances for a specific vulnerability and the Projects they affect.
