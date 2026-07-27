---
description: The Snyk API and Web coverage report
nav_context: classic
---

# Coverage report

Coverage is a fundamental aspect of a scan. It can be the difference between a useful, successful scan and an uninformative one.

As soon as your scan starts, you can download a provisional coverage report to understand what is happening during your scan. This report is subject to change until the scan is finished.

After the scan completes, you can export the detailed coverage report, which lists scanned URLs along with the ones that were not scanned.

Use this report to check whether the scanner is reaching every possible endpoint and filtering them successfully.

## How coverage works

Before and during tests, the crawler navigates your website to find every possible endpoint while testing every input it finds. The crawler then sends those URLs to the scanner to test for vulnerabilities.

## Finding your report

You can find your scan's coverage next to the scan's report and download it for further analysis.

By default, the report shows only the accepted endpoints. To include rejected endpoints, navigate to your **Target Settings** and change the **Coverage Detail** under **Scanner**.

Visit [Generate a CSV coverage report](generate-csv-coverage-report.md) to learn how to download your coverage report.

## Reading the CSV file

A CSV file is a plain text file that contains a list of data separated by commas.

After downloading the file, you can open it in your terminal, text editor, or spreadsheet application.

The first column shows the type of request the crawler made (HTTP requests such as GET, POST, PUT, DELETE, PATCH, and so on).

The second column shows the found or targeted URL. Check this column to verify that all possible endpoints for your website are being reached.

The third and fourth columns show the request's response and its meaning. The most frequent responses are:

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

These requests are then accepted or rejected by the engine's standards. If the engine rejects an endpoint, it provides a reason, such as:

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

## Examples

Here are a few examples to help you understand your coverage report:

*   `"GET","http://example.com/product.php","200","accepted",""`

    A GET request to http://example.com/product.php responded with 200 (OK) and has been accepted to be tested for vulnerabilities by the scanner.
*   `"POST","http://example.com/userinfo.php","302","accepted",""`

    A POST request to http://example.com/userinfo.php responded with 302 - Found and has been accepted to be tested for vulnerabilities by the scanner.
*   `"GET","http://example.com/artists.php?artist=1","-","rejected","deduplicated (simhash)"`

    A GET request to http://example.com/artists.php?artist=1 was rejected by the scanner because the endpoint's simhash was the same as another endpoint's.
*   `"GET","http://example.com/showimage.php?file=./pictures/5.jpg","-","rejected","query string limit reached"`

    A GET request to http://example.com/showimage.php?file=./pictures/5.jpg was rejected by the scanner because the base URL with the query string reached the visit limit.

With this knowledge, you can read your coverage feedback and identify any blind spots or misconfigurations of your target.
