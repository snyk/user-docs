---
description: How to generate a CSV coverage report in Snyk API and Web
nav_context: classic
---

# Generate CSV coverage report

To verify that your target was scanned in full and that the scan covered all URLs, Snyk API & Web lets you download a CSV coverage or crawling report. This report shows a list of all the URLs that Snyk visited while the scan was running.

There are a few different ways to generate a CSV coverage report:

* From a target's Scan Activity.
* From the global Scan Activity.
* From a target's scan results.

This article describes each way in detail.

## From a target's Scan Activity

Follow this procedure to generate a crawling report from the Scan Activity of a target.

1. In the **Targets** section, click the target name to show its details.
2. Click the **Scan Activity** tab. A **Reports** button appears next to each scan.
3. Click the button and choose the **Crawling report** to generate the report.

## From the global Scan Activity

Follow this procedure to generate the crawling report from the global Scan Activity in your account.

1. Access the **Scans** section. The **Scan Activity** tab displays a **Reports** button next to each scan.
2. Click the button and choose the **Crawling report** to generate the report.

## From a target's scan results

Follow this procedure to generate the scan report from the details of a target's scan results.

1. Navigate to the **Targets** section and click the target name to show its details.
2. Click the **Scan Activity** tab.
3. In the scans listed in the **Scan Activity** tab, click the row of the scan to view the scan details.
4. Download the **Crawling report** either from the **Status** section or from the **Reports** button.

If you open the full view of the scan details, you see the same two options.

After you request the report, Snyk downloads the CSV file automatically. The file lists all the endpoints Snyk accessed during that target scan.

Visit [Coverage report](coverage-report.md) to learn how to read and interpret the CSV coverage report.
