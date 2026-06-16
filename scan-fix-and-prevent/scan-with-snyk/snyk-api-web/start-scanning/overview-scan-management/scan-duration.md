# Scan duration

The duration of a scan depends on a number of factors and it is not easy to estimate it beforehand. It depends on:

* **Number of pages of your site** - Scanning involves a certain number of tests per page, so the more you have, the more time it will take. However, this is not linear since the number of tests per page also depends on a number of factors.
* **Number of inputs (injection points) per page** - The more injection points you have on a page, the more tests will be performed. An input is, for instance, a form field. If you have a page with a large form, it will take more time to scan it.
* **Performance of the server that hosts your site** - If your server is slow (for example, if it takes several seconds to load a page), it will take more time to scan your site. If, on the other hand, your server's response time is fast, it will speed up the scan. The velocity of the scanner is adjusted automatically based on your site's response time. If the average response time increases, the scanner will slow down. If the average response time is stable, the scanner will increase the number of requests per second.
