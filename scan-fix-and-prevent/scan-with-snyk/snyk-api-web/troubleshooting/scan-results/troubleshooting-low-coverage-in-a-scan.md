# Troubleshoot low coverage in a scan

A scan must cover as much of the target scope as possible to identify the maximum number of vulnerabilities. If your scan shows low coverage, follow these troubleshooting steps to identify and resolve the issue.

For more information about coverage reports, visit [Coverage report](../../review-and-fix/overview-reports/coverage-report.md) and [Generate CSV coverage report](../../review-and-fix/overview-reports/generate-csv-coverage-report.md).

## The problem

When running a scan on a target, the coverage is low.

## Troubleshoot the problem

Work through the following steps to identify possible causes and solutions.

### Step 1: Check for target authentication

If the target has authentication, verify the scanner logged in.

1. Navigate to the **Targets** page.
2. Click the target name to see its details.
3. Click **Scan Activity** to see the list of scans.
4. Click **View** for the scan you want to check.
5. Click **Crawling Report** to download the spreadsheet with scanned URLs.
6. Check if URLs that are only available for authenticated users appear in the report.

If no URLs for authenticated users are listed, the scan failed to log in.

| Cause                                    | Solution                                                                                                                            |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| The scan failed to log in to the target. | Review the target authentication configuration. For common authentication issues, visit the Troubleshooting Authentication section. |

### Step 2: Check for missing SPA API

If the target is a Single-Page Application (SPA) with a backing API, check if the API uses a different URL.

For example:

* SPA URL: `https://example.com`
* SPA API URL: `https://api.example.com`

If the backing API has a URL different from the SPA, Snyk API & Web needs to know the API URL to scan the SPA properly.

| Cause                                                        | Solution                                                                                                                                                    |
| ------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The target is a SPA with its backing API on a different URL. | Add the API URL as an extra host in target settings. Visit [Configure extra hosts](../../configure-targets/configure-web-targets/configure-extra-hosts.md). |

### Step 3: Check for a blocking WAF

Check if a Web Application Firewall (WAF) is blocking scan requests after the scan has started.

1. Navigate to the **Targets** page.
2. Click the target name to see its details.
3. Click **Scan Activity** to see the list of scans.
4. Click **View** for the scan you want to check.
5. Click **Crawling Report** to download the spreadsheet with scanned URLs.
6. Check if URLs started receiving HTTP 403 error status at some point during the scan.
7. Open a browser in incognito mode and test those URLs to see if a WAF is blocking access.

If a WAF blocks access to URLs, Snyk API & Web cannot scan them.

| Cause                                             | Solution                                                                                                                                                           |
| ------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| A WAF is blocking access to URLs during the scan. | Add Snyk API & Web IPs to the WAF's allowlist. Visit [Configure IPs in WAFs](../../start-scanning/overview-scan-access-and-connectivity/configure-ips-in-wafs.md). |

### Step 4: Check for blocking WordPress plugin

If the target is WordPress, check if a WordPress security plugin (for example, WordFence) is blocking scan requests.

1. Navigate to the **Targets** page.
2. Click the target name to see its details.
3. Click **Scan Activity** to see the list of scans.
4. Click **View** for the scan you want to check.
5. Click **Crawling Report** to download the spreadsheet with scanned URLs.
6. Check if URLs have HTTP 403 error status.
7. Open a browser in incognito mode and test those URLs to see if a WordPress plugin is blocking access.

If a WordPress plugin is blocking access to URLs, Snyk API & Web cannot scan them.

| Cause                                                                   | Solution                                                                                                                                                                                                                      |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| A WordPress plugin (for example, WordFence) is blocking access to URLs. | Configure the WordPress plugin to allow requests from Snyk API & Web IPs. Visit [Scanner IP address](../../start-scanning/overview-scan-access-and-connectivity/scanner-ip-address.md) for the scanner's outgoing IP address. |

After following these steps and applying the solutions, scans achieve the expected coverage for your targets.
