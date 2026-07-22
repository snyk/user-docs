---
description: How scan duration works for Snyk API and Web scans
nav_context: agnostic
---

# Scan duration

The duration of a scan depends on several factors and is not easy to estimate beforehand. It depends on:

* Number of pages of your site: Scanning involves a certain number of tests per page, so the more pages you have, the more time the scan takes. However, this is not linear, because the number of tests per page also depends on several factors.
* Number of inputs (injection points) per page: The more injection points you have on a page, the more tests the scanner performs. An input is, for example, a form field. A page with a large form takes more time to scan.
* Performance of the server that hosts your site: If your server is slow (for example, if it takes several seconds to load a page), scanning your site takes more time. If your server response time is fast, the scan speeds up. The scanner adjusts its velocity automatically based on your site response time. If the average response time increases, the scanner slows down. If the average response time is stable, the scanner increases the number of requests per second.
