# Interpret target scan results

After setting up your target at Snyk API & Web, you can start a scan and access the scan details to visualize the progress and results with real-time updates.

On this page, you will find three valuable sections to analyze and interpret a target scan:

* **Overview** tab, with the Risk, Status, Settings, and Details of the scan.
* **Findings** tab, with the list of vulnerabilities found by the scanner.
* **Reports** section, with a couple of options for reports about the scan.

This article provides all the details on these sections.

## Overview

The **Risk** section shows the counter of vulnerabilities found and the compliance tags. If a target does not meet the required checklist needed for being compliant, the tags will be red.

The **Status** section shows all the relevant information about the scan date (when it started and ended), its duration, status, whether the login was successful or not, and allows for the download of the crawling report (or provisory crawling report, if the scan is still running).

Visit [What happens during a scan](../start-scanning/what-happens-during-a-scan.md) for details on scan states.

The **Settings** section shows the settings that were used during the scan.

The **Details** section shows the evolution of the three major components at play, each one with a specific job:

* The **fingerprinter** detects the technologies used on the target.
* The **crawler** goes through the target's URLs and interacts with every element found, clicking on buttons and filling out forms, among other actions.
* The **scanner** finds vulnerabilities within the target's URLs obtained by the crawler.

Under the progress bars, you can have more details of the work performed by each component:

* **Fingerprinter** - Shows the list of technologies detected.
* **Crawler** - Shows the number of crawled URLs and how many were deduplicated. It also shows whether the login was successful or not, and the statistics of the HTTP response codes obtained during the crawling, which can be relevant to figuring out issues that may require attention. If the crawler is still working, you will also see which URLs are being crawled at the time. If you need to check the complete list of crawled URLs, issue a report as described in the Reports section.
* **Scanner** - Shows the number of scanned URLs, the average time it took, and, similarly to the crawler, the statistics of the HTTP response codes obtained, which can be relevant to figuring out issues that may require attention. If the scanner is still working, you will also see which URLs are being scanned at the time.

To know more about the HTTP response codes displayed in the crawler and scanner details, visit [HTTP status codes in target scans](http-status-codes-in-target-scans.md).

## Findings

In this section, you can see the list of all the vulnerabilities found by the scanner during the scan.

Depending on the type of vulnerability found, its exploitability, impact, and scope, a CVSS score and risk/severity classification are attributed to the finding, helping you prioritize the vulnerability fixes.

In general terms, vulnerabilities with a more significant impact that can be easily exploited have a higher risk. While vulnerabilities with a lower impact are more complex to replicate and require several specific conditions to be exploited, they likely represent a lower risk.

Here, you can analyze the findings and decide which actions to take.

## Reports

At the bottom of the page, the **Reports** dropdown button allows you to download the crawling and the scan reports.

The options available will depend on the scan status:

* **Preliminary crawling report** - While the scan is running, this option is available to issue and download a provisory coverage report. It allows you to check the endpoints the scanner has reached so far. Bear in mind, however, that this report might be subject to change until the scan finishes.
* **Crawling report** - Once the scan finishes, this option is available to issue and download the final version of the coverage report. It allows you to check every endpoint the scanner reached during the entire scan.
* **Scan report** - Once the scan finishes successfully, this button is available to issue and download a target scan report, which you can share within your company or with your auditors or customers. The report will list the vulnerabilities found, along with a detailed description and ways of fixing them, as well as a list of all the tests performed during the target scan. Depending on the type of report, it can have additional information on tests performed for specific compliances, such as OWASP Top 10 or PCI-DSS.
