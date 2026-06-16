# Customize a scan profile

Snyk API & Web provides a variety of [built-in scan profiles](built-in-scan-profiles.md) to choose from and define how your targets are scanned. Each of these built-in scan profiles is a group of scanning conditions that are pre-configured by Snyk API & Web to provide certain pre-defined scanning behaviors.

You can also create custom scan profiles to adjust and fine-tune scans for your targets.

## Create or edit a custom scan profile

In Snyk API & Web, customize a scan profile as follows:

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
   * **HTTP methods** - Choose the type of HTTP methods to be used in scanning requests. This allows the choice of an ideal set of methods for targets with websites or applications that are in production.
   * **Scan Speed** - Choose the throughput of scanning requests regarding the target's response time to avoid overloading the target with too many requests and optimizing the resources consumed by a scan. Regardless of the scan speed, if Snyk API & Web detects that the target is not able to handle the requests throughput during a scan, the scanner will automatically throttle down to attain optimal performance.
   * **Request Delay** - Set the time delay (in milliseconds) between requests for each scanning thread. This is an approximate value and is more accurate for slower Scan speed settings. The maximum delay allowed is 5000ms. If not defined, there is no delay between requests.
   * **Limit scan duration** - Set the maximum time the scan is allowed to run. If not set, there is no limit. The usage of this setting might cause the scan to miss vulnerabilities.

   ### Crawler

   * **Crawler deduplication** - Snyk API & Web uses a Simhash algorithm to detect similar pages and scan only a few of them. A page is considered similar if it shares the same HTML element structure. This feature is enabled by default. Disabling the checkbox turns it off and can increase the scan duration significantly.
   * **URL pattern detection** - Snyk API & Web detects patterns in URLs identifying similar pages and scans only a few of them. For instance, pages like `/2023-10-08-probely-scanner-finds-another` and `/2022-03-18-cibersecurity-is-important` share the pattern `/YYYY-MM-DD-` followed by several words separated by a hyphen. This feature is enabled by default. Disabling the checkbox turns it off and can increase the scan duration significantly.
   * **Maximum URLs crawled** - Set the maximum number of URLs the crawler visits. The maximum value available is 50,000. The default of 5,000 is a good compromise between coverage and scan time.


   ### Scanner

   * **Scanner Payloads** - Choose the diversity of payloads and headers used for testing vulnerabilities to fine-tune the number of scanning requests made to each URL of the target. Regardless of the scanner payloads, the vulnerabilities considered for testing are the same.
   * **Vulnerabilities** - Choose the vulnerabilities to be verified by the scanner: all or a specific subset.

1. Click **SAVE** to finish the customization of the scan profile.

Once created, custom scan profiles are available in the list of profiles in the target settings. You only have to [switch the scan profile](switch-scan-profile.md) of the target to the desired custom scan profile.

## Delete a custom scan profile

You can delete a custom scan profile by clicking **Delete** on the list of scan profiles. Snyk API & Web will prompt you to confirm your action. If one or more targets still use the profile to be deleted, Snyk API & Web also indicates which will be the replacement scan profile to set in those targets.
