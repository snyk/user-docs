# Generate a CSV coverage report

Sometimes, you might want to make sure that your target was scanned in full and that all URLs were covered by the scan. For that reason, Snyk API & Web allows you to download a CSV coverage or crawling report, which shows a list of all the URLs that Snyk API & Web visited while the scan was running.

There are a few different ways to generate a CSV coverage report:

* From a target's Scan Activity.
* From the global Scan Activity.
* From a target's scan results.

This article describes each way in detail.

## From a target's Scan Activity

Follow this procedure to generate a crawling report from the Scan Activity of a target.

1. In the **Targets** section, click the target name to show its details.
1. Click the **Scan Activity** tab. You will see a **Reports** button next to each scan.
1. Generate the report by clicking the button and choosing the Crawling report.

<figure><img src="../../../../.gitbook/assets/generate-csv-coverage-report-target-scan-activity.png" alt="Scan Activity tab showing Reports button with Crawling report option"></figure>

## From the global Scan Activity

Follow this procedure to generate the crawling report from the global Scan Activity in your account.

1. Once you access the **Scans** section, the **Scan activity** tab displays a **Reports** button next to each scan.
1. Generate the report by clicking the button and choosing the Crawling report.

<figure><img src="../../../../.gitbook/assets/generate-csv-coverage-report-global-scan-activity.png" alt="Global Scan Activity tab showing Reports button with Crawling report option"></figure>

## From a target's scan results

Follow this procedure to generate the scan report from the details of a target's scan results.

1. Go to the **Targets** section and click the target name to show its details.
1. Click the **Scan Activity** tab.
1. In the scans listed in the **Scan Activity** tab, click the row of the scan to view the scan details.
1. At this point, you can download the **Crawling report** either from the **Status** section or from the **Reports** button.

<figure><img src="../../../../.gitbook/assets/generate-csv-coverage-report-scan-results-status.png" alt="Scan results page showing Crawling report download link in Status section"></figure>

<figure><img src="../../../../.gitbook/assets/generate-csv-coverage-report-scan-results-reports.png" alt="Scan results page showing Reports button with Crawling report option"></figure>

If you open the full view of the scan details, you will see the same two options:

<figure><img src="../../../../.gitbook/assets/generate-csv-coverage-report-scan-details.png" alt="Full scan details view showing both Crawling report download options"></figure>

Once you request the report, the CSV file is automatically downloaded, and you have a list of all the endpoints Snyk API & Web accessed during that particular target scan.

Visit [Coverage report](coverage-report.md) to learn how to read and interpret the CSV coverage report.
