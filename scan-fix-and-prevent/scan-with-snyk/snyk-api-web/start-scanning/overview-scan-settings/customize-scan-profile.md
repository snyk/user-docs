---
description: How to customize a Snyk API and Web scan profile
---

# Customize a scan profile

Snyk API & Web provides a variety of [built-in scan profiles](built-in-scan-profiles.md) to choose from and define how Snyk scans your targets. Each built-in scan profile is a group of scanning conditions that Snyk pre-configures to provide certain pre-defined scanning behaviors.

You can also create custom scan profiles to adjust and fine-tune scans for your targets.

## Create or edit a custom scan profile

To customize a scan profile in Snyk API & Web:

1. Open the **Settings** dropdown menu on the bottom-left corner of the navigation bar and click **Scan Profiles**.
1. On the Scan Profiles screen, you have three options to customize a scan profile:
   * **Add** - Click **ADD CUSTOM PROFILE** to create a new custom scan profile starting from a blank configuration.
   * **Clone** - Click **Clone** for a scan profile in the list to create a new custom scan profile based on an existing configuration and adjust it to your needs.
   * **Edit** - Click **Edit** for a scan profile in the list to adjust its configuration. This option is unavailable for [built-in scan profiles](built-in-scan-profiles.md), which can only be cloned.
1. In the form that follows, configure the custom scan profile:
   1. Type the name.
   1. Type a description (optional).
   1. Customize the scanning behavior, divided into three sections:

   ### Global

   * **Target type** - Choose the type of target for which this scan profile is available: Web applications or standalone APIs.
   * **HTTP methods** - Choose the type of HTTP methods to use in scanning requests. This lets you choose an ideal set of methods for targets with websites or applications that are in production.
   * **Scan Speed** - Choose the throughput of scanning requests relative to the target's response time. This avoids overloading the target with too many requests and optimizes the resources a scan consumes. Regardless of the scan speed, if Snyk detects that the target cannot handle the request throughput during a scan, the scanner automatically throttles down to attain optimal performance.
   * **Request Delay** - Set the time delay (in milliseconds) between requests for each scanning thread. This is an approximate value and is more accurate for slower Scan speed settings. The maximum delay allowed is 5000ms. If not defined, there is no delay between requests.
   * **Limit scan duration** - Set the maximum time the scan can run. If not set, there is no limit. Using this setting can cause the scan to miss vulnerabilities.

   ### Crawler

   * **Crawler deduplication** - Snyk uses a Simhash algorithm to detect similar pages and scan only a few of them. A page is considered similar if it shares the same HTML element structure. This feature is enabled by default. Clearing the check box turns it off and can increase the scan duration significantly.
   * **URL pattern detection** - Snyk detects patterns in URLs identifying similar pages and scans only a few of them. For example, pages such as `/2023-10-08-probely-scanner-finds-another` and `/2022-03-18-cibersecurity-is-important` share the pattern `/YYYY-MM-DD-` followed by several words separated by a hyphen. This feature is enabled by default. Clearing the check box turns it off and can increase the scan duration significantly.
   * **Maximum URLs crawled** - Set the maximum number of URLs the crawler visits. The maximum value available is 50,000. The default of 5,000 is a good compromise between coverage and scan time.


   ### Scanner

   * **Scanner Payloads** - Choose the diversity of payloads and headers used for testing vulnerabilities to fine-tune the number of scanning requests made to each URL of the target. Regardless of the scanner payloads, the vulnerabilities considered for testing are the same.
   * **Vulnerabilities** - Choose the vulnerabilities for the scanner to verify: all or a specific subset.

1. Click **SAVE** to finish customizing the scan profile.

After you create a custom scan profile, it is available in the list of profiles in the target settings. To use it, [switch the scan profile](switch-scan-profile.md) of the target to that custom scan profile.

## Delete a custom scan profile

To delete a custom scan profile, click **Delete** in the list of scan profiles. Snyk prompts you to confirm your action. If one or more targets still use the profile you are deleting, Snyk also indicates the replacement scan profile to set in those targets.
