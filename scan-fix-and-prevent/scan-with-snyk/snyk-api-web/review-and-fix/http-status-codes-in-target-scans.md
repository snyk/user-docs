# HTTP status codes in target scans

When analyzing your [target scan results](interpret-target-scan-results.md), you may find details about the crawler and the scanner, including a list of HTTP response status codes.

This information provides an overall view of how the application is responding to the crawler and scanner activities. When looking at them, bear in mind the following:

* **HTTP 2XX** - These are success status codes, that is, successful requests from Snyk API & Web.
* **HTTP 3XX** - These redirection status codes should be seen as a normal behavior of the application. They are most likely a result of the wide range of combinations in requests when Snyk API & Web crawls the application and searches for vulnerabilities.
* **HTTP 4XX** - These client error status codes should be seen as normal behavior of the application. They are most likely a result of the wide range of combinations in requests when Snyk API & Web crawls the application and searches for vulnerabilities.
* **HTTP 5XX** - These server error status codes indicate that the application is not responding well to Snyk API & Web requests due to server problems. Check the server to understand why this is happening.

## Troubleshooting

Although some of the reported HTTP status codes can be seen as normal behavior (because the applications or APIs are doing what they are supposed to), there are some cases that need troubleshooting and fixing.

### HTTP 5XX status codes

The server is not responding well.

Look into the server to find any misbehaviors or problems and fix them.

### HTTP 401 status code with API targets

Snyk API & Web is not able to log in to scan the API.

In the Snyk API & Web app, navigate to your target settings and review the authentication according to how the API was specified:

* **Swagger/OpenAPI** - The authentication configuration is in the target settings. Review and fix it.
* **Postman Collection** - The authentication is implemented in the Postman Collection. Review the implementation, fix it, export the collection, and upload it in the target settings or submit the Postman Collection schema URL.
