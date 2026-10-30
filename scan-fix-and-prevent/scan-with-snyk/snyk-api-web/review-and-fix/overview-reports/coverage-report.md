---
nav_context: classic
description: The Snyk API & Web coverage report
---

{% include "../../../../.gitbook/includes/new-navigation-banner.md" %}

# Coverage report

Coverage is a fundamental aspect of a scan. It can be the difference between a useful, successful scan and an uninformative one.

As soon as your scan starts, you can download a provisional coverage report to understand what is happening during your scan. This report may update until the scan is complete.

After the scan completes, you can export the detailed coverage report, which lists, for each URL discovered, the **HTTP request method** and **response codes**, whether the request was **authenticated**, whether the endpoint was sent to the scanner for **further tests** or just **crawled**, and why.

Use this report to check whether the scanner reaches every possible endpoint and filters them successfully.

## How coverage works

Before and during tests, the crawler navigates your website to find every possible endpoint while testing every input it finds. The crawler then sends those URLs to the scanner to test for vulnerabilities.

By default, the report includes both accepted and rejected endpoints, with the reason for each rejection. Snyk keeps rejected endpoints for 90 days after the scan. To get only the accepted endpoints, navigate to **Target Settings** > **Reports** and change the **Coverage Detail** to **Scanned URLs**.

Rejected endpoints are available only if your plan includes them. If your plan does not, the report lists accepted endpoints only.

## Finding your report

Visit [Generate a CSV coverage report](generate-csv-coverage-report.md) to learn how to download your coverage report.

## Reading the coverage report

The coverage report is a CSV file, which means it is a plain text file that contains a list of data separated by commas. After downloading it, you can open it in your terminal, text editor, or spreadsheet application.

To explore the same data in the Snyk web interface, see [Interpret coverage results](../interpret-coverage-results.md).

The `id` column shows the request ID. You can use the API to obtain more information about the respective request and response using the following endpoint:

`GET https://api.probely.com/targets/<target_id>/scans/<scan_id>/endpoints/<request_id>/`

The `request_method` shows the type of request the crawler made (HTTP requests such as GET, POST, PUT, DELETE, PATCH, and so on).

The `host` column shows the origin of the URL: the scheme, host, and port, such as `https://example.com:8443`. The host is lowercase, and the port is left out when it is the default for the scheme (80 for `http`, 443 for `https`). The same host reached over `http://` and `https://` is listed as two different hosts.

The `url` column shows the path and query of the found or targeted URL, such as `/login?next=/`. Combine `host` and `url` to rebuild the full URL. Check this column to verify that all possible endpoints of your website are being reached.

The `status_code` shows the request's response. The most frequent responses are:

* 200 - OK.
* 301 - Moved Permanently.
* 302 - Found.
* 307 - Temporary Redirect.
* 308 - Permanent Redirect.
* 401 - Unauthorized.
* 403 - Forbidden.
* 404 - Not Found.
* 500 - Internal Server Error.
* 503 - Service Unavailable.

The `authenticated` column shows whether the request was authenticated. For the crawler to authenticate during the scan, configure your target **Settings > Authentication**.

The `result` column shows whether the engine **accepted** or **rejected** the requests. If the engine rejects an endpoint, it provides a **reason** (`reason` column), such as:

* `is on keyword reject list`
  * Meaning: Rejected because the URL contains a keyword that is on the internal keyword reject list.
  * The crawler blocks URLs that contain `logout`, `logoff`, or `signout`, so that it does not lose its session. It also blocks URLs that contain `delete`, `remove`, or `apagar` (Portuguese for "delete"), so that it does not trigger destructive actions.
    * Example: `GET /api/delete?id=1` is rejected with this reason.
* `file extension ignored`
  * Meaning: Rejected because the URL path ends in a file extension from the internal reject list.
  * The crawler skips URLs that end in any of the following extensions. The rule applies to the end of the URL path, not to the type of content the URL returns. For example, `https://example.com/report.pdf` is skipped, but the same document served at `https://example.com/download?doc=42` is crawled and scanned.
    * JavaScript files are an exception. When a page loads a `.js` file as a script, the scanner tests that file passively, as long as the **Include scripts** scope setting of your target is on (the default). The same `.js` file found only as a URL, for example in a link or a `src` attribute, is rejected with this reason.

<table><thead><tr><th width="204.6875">Type</th><th>Extensions</th></tr></thead><tbody><tr><td>Images</td><td><code>png</code>, <code>jpg</code>, <code>jpeg</code>, <code>gif</code>, <code>ico</code>, <code>bmp</code>, <code>webp</code>, <code>tiff</code>, <code>tif</code>, <code>svg</code>, <code>raw</code>, <code>psd</code></td></tr><tr><td>Media</td><td><code>mp4</code>, <code>m4v</code>, <code>mp3</code>, <code>avi</code>, <code>mov</code>, <code>mkv</code>, <code>mpg</code>, <code>mpeg</code>, <code>webm</code>, <code>ogg</code>, <code>wav</code>, <code>wma</code>, <code>wmv</code>, <code>flv</code>, <code>swf</code>, <code>vob</code>, <code>rm</code>, <code>3gp</code>, <code>3g2</code>, <code>m3u8</code></td></tr><tr><td>Fonts</td><td><code>woff</code>, <code>woff2</code>, <code>ttf</code>, <code>otf</code>, <code>eot</code></td></tr><tr><td>Archives and binaries</td><td><code>zip</code>, <code>rar</code>, <code>7z</code>, <code>gz</code>, <code>tar</code>, <code>tgz</code>, <code>bak</code>, <code>exe</code>, <code>iso</code>, <code>dmg</code>, <code>bin</code>, <code>apk</code>, <code>dll</code>, <code>msi</code>, <code>db</code>, <code>ini</code>, <code>icns</code></td></tr><tr><td>Documents</td><td><code>pdf</code>, <code>doc</code>, <code>odt</code>, <code>xls</code>, <code>xlsx</code>, <code>xlsm</code>, <code>ods</code>, <code>ppt</code>, <code>pptx</code>, <code>pps</code>, <code>odp</code></td></tr><tr><td>Web assets</td><td><code>css</code>, <code>scss</code>, <code>ts</code>, <code>js</code></td></tr></tbody></table>

* `is on path reject list`
  * Meaning: Rejected because the URL path matches a built-in list of infrastructure paths. This list contains the Cloudflare `/cdn-cgi/rum` and `/cdn-cgi/challenge-platform/` paths.
  * This list is built into the crawler. You do not configure it, and it is separate from your own reject list.
* `is on user reject list`
  * Meaning: Rejected because the URL matches an item that is on the user's reject list.
* `deduplicated (simhash)`
  * Meaning: The content structure of the endpoint's simhash was the same as another endpoint's, so it was rejected.
* `path limit reached`
  * Meaning: The base URL (without fragments and query strings) reached the visit limit.
* `query string limit reached`
  * Meaning: The base URL with the same query string parameters (values excluded) reached the visit limit. The default limit is 2.
* `fragment limit reached`
  * Meaning: The same base URL with fragments or hashes reached the visit limit.
* `auto pattern limit reached`
  * Meaning: This endpoint reached an automatic URL limiter limit. These limits detect IDs, hashes, slugs, localizations, UUIDs, and so on.
* `path pattern limit reached`
  * Meaning: This endpoint reached the user URL limiter limit.
* `URL out of scope`
  * Meaning: This URL is outside the scope of the scan.
* `Crawled; not scanned (URL fragment)`
  * Meaning: This URL is a fragment, so it was crawled but not scanned; the fragmentless URL is sent.
* `Crawled; not scanned (redirect document timing)`
  * Meaning: This page redirected back to its own URL, for example, with a 302 that sets a cookie and then a 200. The 302 is the endpoint sent to the scanner. The final page gets a second row with this reason so you can see its own status, response time, and page load. Snyk does not send this row to the scanner.
* `incremental already sent`
  * Meaning: This URL was not sent to the scanner again because it has not changed since the previous scan.
  * This reason appears only in incremental scans. To test all endpoints regardless of change status, run a full scan.

The `info` column provides additional information, if available.

The `raw_request_size` and the `raw_response_size` are the request and response sizes in bytes, respectively.

The `request_parameters_count` is the number of parameters tested from that request.

The `response_time_ms` is how long the target took to respond while the crawler visited the URL, in milliseconds. It is empty if the URL was never requested.

The `page_load_ms` is how long the page took to load in the browser, in milliseconds. It is measured for pages only, not for the requests a page makes. It is filled for web targets only, and is empty if the page was never loaded.

The `raw_request_size_display` and the `raw_response_size_display` are the human-readable sizes of the request and response, respectively (e.g., instead of 1024, it displays 1 KB).

## Examples

Here are a few examples to help you understand your coverage report:

*   `"TBWoS3ogjjv","GET","http://example.com","/product.php","200","FALSE","accepted","","","773","6911","0","","","773 B","6.7 KB"`

    An unauthenticated GET request to http://example.com/product.php responded with 200 (OK) and has been accepted by the scanner (for example, it is tested for vulnerabilities).
*   `"ABXoS2ogzjv","POST","http://example.com","/userinfo.php","302","FALSE","accepted","","","853","4024","1","","","853 B","3.9 KB"`

    An unauthenticated POST request to http://example.com/userinfo.php responded with 302 - Found and has been accepted by the scanner (for example, it is tested for vulnerabilities).
*   `"QWLoP0ogcj3","GET","http://example.com","/artists.php?artist=1","-","FALSE","rejected","deduplicated (simhash)","","86","0","1","","","86 B","0 B"`

    A GET request to http://example.com/artists.php?artist=1 was rejected by the scanner because the endpoint's simhash was the same as another endpoint's.
*   `"XcKoc30bwjG","GET","http://example.com","/showimage.php?file=./pictures/5.jpg","-","FALSE","rejected","query string limit reached","","86","0","1","","","86 B","0 B"`

    A GET request to http://example.com/showimage.php?file=./pictures/5.jpg was rejected by the scanner because the base URL with the query string reached the visit limit.
*   `"Qw3bxFftFVsL","GET","http://example.com","/#features","200","FALSE","rejected","Crawled; not scanned (URL fragment)","","65","0","0","","","65 B","0 B"`

    A GET request to http://example.com/#features was rejected by the scanner because the URL includes a fragment and the fragmentless endpoint was already scanned.

With this knowledge, you can read your coverage feedback and identify any blind spots or misconfigurations of your target.
