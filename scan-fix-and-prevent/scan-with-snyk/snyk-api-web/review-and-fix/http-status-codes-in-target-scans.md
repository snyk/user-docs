# HTTP status codes in target scans

When analyzing your [target scan results](interpret-target-scan-results.md), you can find details about the crawler and the scanner, including a list of HTTP response status codes.

This information provides an overall view of how the application responds to the crawler and scanner activities. When looking at them, keep the following in mind:

* **HTTP 2XX** - These are success status codes, that is, successful requests from Snyk API & Web.
* **HTTP 3XX** - These redirection status codes indicate normal behavior of the application. They most likely result from the wide range of request combinations when Snyk crawls the application and searches for vulnerabilities.
* **HTTP 4XX** - These client error status codes indicate normal behavior of the application. They most likely result from the wide range of request combinations when Snyk crawls the application and searches for vulnerabilities.
* **HTTP 5XX** - These server error status codes indicate that the application is not responding well to Snyk requests because of server problems. Check the server to understand why this is happening.

## Troubleshooting

Although some reported HTTP status codes indicate normal behavior (because the applications or APIs are doing what they are supposed to), some cases need troubleshooting and fixing.

### HTTP 5XX status codes

The server is not responding well.

Look into the server to find any misbehaviors or problems and fix them.

### HTTP 401 status code with API targets

Snyk cannot log in to scan the API.

In the Snyk app, navigate to your target settings and review the authentication according to how the API was specified:

* **Swagger/OpenAPI** - The authentication configuration is in the target settings. Review and fix it.
* **Postman Collection** - The authentication is implemented in the Postman Collection. Review the implementation, fix it, export the collection, and upload it in the target settings or submit the Postman Collection schema URL.
