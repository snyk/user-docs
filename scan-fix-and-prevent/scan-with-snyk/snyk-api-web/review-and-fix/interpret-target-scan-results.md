---
description: How to interpret target scan results in Snyk API and Web
---

# Interpret target scan results

After setting up your target at Snyk API & Web, you can start a scan and access the scan details to visualize the progress and results with real-time updates.

On this page, you find three valuable sections to analyze and interpret a target scan:

* **Overview** tab, with the Risk, Status, Settings, and Details of the scan.
* **Findings** tab, with the list of vulnerabilities found by the scanner.
* **Reports** section, with a couple of options for reports about the scan.

This article provides all the details on these sections.

## Overview

The **Risk** section shows the counter of vulnerabilities found and the compliance tags. If a target does not meet the required checklist for compliance, the tags are red.

The **Status** section shows all the relevant information about the scan date (when it started and ended), its duration, status, and whether the login was successful. This section also lets you download the crawling report, or the provisory crawling report if the scan is still running.

Visit [What happens during a scan](../start-scanning/what-happens-during-a-scan.md) for details on scan states.

The **Settings** section shows the settings used during the scan.

The **Details** section shows the evolution of the three major components at play, each one with a specific job:

* The **fingerprinter** detects the technologies used on the target.
* The **crawler** goes through the target's URLs and interacts with every element found, clicking on buttons and filling out forms, among other actions.
* The **scanner** finds vulnerabilities within the target's URLs obtained by the crawler.

Under the progress bars, you can find more details of the work performed by each component:

* **Fingerprinter** - Shows the list of technologies detected.
* **Crawler** - Shows the number of crawled URLs and how many were deduplicated. It also shows whether the login was successful, and the statistics of the HTTP response codes obtained during the crawling, which can be relevant to identifying problems that require attention. If the crawler is still working, you also see which URLs are being crawled at the time. To check the complete list of crawled URLs, issue a report as described in the Reports section.
* **Scanner** - Shows the number of scanned URLs, the average time it took, and, similar to the crawler, the statistics of the HTTP response codes obtained, which can be relevant to identifying problems that require attention. If the scanner is still working, you also see which URLs are being scanned at the time.

To learn more about the HTTP response codes displayed in the crawler and scanner details, visit [HTTP status codes in target scans](http-status-codes-in-target-scans.md).

## Findings

In this section, you can see the list of all the vulnerabilities found by the scanner during the scan.

Depending on the type of vulnerability found, its exploitability, impact, and scope, a CVSS score and risk/severity classification are attributed to the finding, helping you prioritize the vulnerability fixes.

In general terms, vulnerabilities with a more significant impact that can be exploited easily have a higher risk. Vulnerabilities with a lower impact are more complex to replicate and require several specific conditions to exploit, so they represent a lower risk.

Here, you can analyze the findings and decide which actions to take.

## Reports

At the bottom of the page, the **Reports** dropdown button lets you download the crawling and the scan reports.

The available options depend on the scan status:

* **Preliminary crawling report** - While the scan is running, this option is available to issue and download a provisory coverage report. Use it to check the endpoints the scanner has reached so far. This report is subject to change until the scan finishes.
* **Crawling report** - After the scan finishes, this option is available to issue and download the final version of the coverage report. Use it to check every endpoint the scanner reached during the entire scan.
* **Scan report** - After the scan finishes successfully, this button is available to issue and download a target scan report, which you can share within your company or with your auditors or customers. The report lists the vulnerabilities found, along with a detailed description and ways of fixing them, and a list of all the tests performed during the target scan. Depending on the type of report, it can have additional information on tests performed for specific compliances, such as OWASP Top 10 or PCI-DSS.
