---
nav_context: classic
description: The Snyk API and Web coverage report
---

# Coverage report

Coverage is a fundamental aspect of a scan. It can be the difference between a useful, successful scan and an uninformative one.

As soon as your scan starts, you can download a provisional coverage report to understand what is happening during your scan. This report may update until the scan is complete.

After the scan completes, you can export the detailed coverage report, which lists, for each URL discovered, the **HTTP request method** and **response codes**, whether the request was **authenticated**, whether the endpoint was sent to the scanner for **further tests** or just **crawled**, and why.

Use this report to check whether the scanner reaches every possible endpoint and filters them successfully.

## How coverage works

Before and during tests, the crawler navigates your website to find every possible endpoint while testing every input it finds. The crawler then sends those URLs to the scanner to test for vulnerabilities.

By default, the report shows only the accepted endpoints. To include rejected endpoints and additional information, navigate to **Target Settings** > **Reports** and change the **Coverage Detail** to **All URLs**.&#x20;

## Finding your report

Visit [Generate a CSV coverage report](generate-csv-coverage-report.md) to learn how to download your coverage report.

## Reading the coverage report

The coverage report is a CSV file, which means it is a plain text file that contains a list of data separated by commas. After downloading it, you can open it in your terminal, text editor, or spreadsheet application.

The `id` column shows the request ID. You can use the API to obtain more information about the respective request and response using the following endpoint:

`GET https://api.probely.com/targets/<target_id>/scans/<scan_id>/endpoints/<request_id>/`

The `request_method` shows the type of request the crawler made (HTTP requests such as GET, POST, PUT, DELETE, PATCH, and so on).

The `url` column shows the found or targeted URL. Check this column to verify that all possible endpoints of your website are being reached.

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
  * Words like "logout", "logoff", or "signout" are blocked to ensure that the crawler does not lose its session.
* `file extension ignored`
  * Meaning: Rejected because the URL file extension is on the internal reject list.
  * Extensions such as .exe, .zip, and .tgz get rejected by the crawler.
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

{% hint style="info" %}
By default, only "accepted" endpoints are present on the Coverage report. In your target **Settings > Reports > Coverage Detail**, select **All URLs** to enable "rejected" and the respective "reason".
{% endhint %}

The `info` column provides additional information, if available.

The `raw_request_size` and the `raw_response_size` are the request and response sizes in bytes, respectively.

The `request_parameters_count` is the number of parameters tested from that request.

The `raw_request_size_display` and the `raw_response_size_display` are the human-readable sizes of the request and response, respectively (e.g., instead of 1024, it displays 1 KB).

## Examples

Here are a few examples to help you understand your coverage report:

*   `"TBWoS3ogjjv","GET","http://example.com/product.php","200","false","accepted","","","773","6911","0","773 B","6.7 KB"`

    An unauthenticated GET request to http://example.com/product.php responded with 200 (OK) and has been accepted by the scanner (for example, it is tested for vulnerabilities).
*   `"ABXoS2ogzjv","POST","http://example.com/userinfo.php","302","false","accepted","","","853","4024","1","853 B","3.9 KB"`

    An unauthenticated POST request to http://example.com/userinfo.php responded with 302 - Found and has been accepted by the scanner (for example, it is tested for vulnerabilities).
*   `"QWLoP0ogcj3","GET","http://example.com/artists.php?artist=1","-","false","rejected","deduplicated (simhash)","","86","0","1","86 B","0 B"`

    A GET request to http://example.com/artists.php?artist=1 was rejected by the scanner because the endpoint's simhash was the same as another endpoint's.
*   `"XcKoc30bwjG","GET","http://example.com/showimage.php?file=./pictures/5.jpg","-","false","rejected","query string limit reached","","86","0","1","86 B","0 B"`

    A GET request to http://example.com/showimage.php?file=./pictures/5.jpg was rejected by the scanner because the base URL with the query string reached the visit limit.
*   `"Qw3bxFftFVsL","GET","http://example.com/#features","200","false","rejected","Crawled; not scanned (URL fragment)","","65","0","0","65 B","0 B"`

    A GET request to http://example.com/#features was rejected by the scanner because the URL includes a fragment and the fragmentless endpoint was already scanned.

With this knowledge, you can read your coverage feedback and identify any blind spots or misconfigurations of your target.
