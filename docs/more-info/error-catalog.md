# Snyk Error Codes
  The error codes in the table below describe the codes that you may encounter while working with the [Snyk API](../snyk-api-info/README.md) or [CLI](../snyk-cli/README.md). When errors are encountered using the API, they will also have an appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). If you encounter errors without an error code, use the HTTP status code to determine the appropriate action.
  
  
---
# [Snyk](https://docs.snyk.io/introducing-snyk)
### [SNYK-0001](#snyk-0001-too-many-requests) - Too many requests
  The service has received too many requests and will be throttled

  **HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-9999](#snyk-9999-server-error) - Server error
  An unexpected server error was encountered.

  **HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

  **Exit Code:** N/A

  **Help Links:**

  
---
# [Open Source Languages & Package Managers](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support)
### [SNYK-OS-0001](#snyk-os-0001-unsupported-ecosystem) - Unsupported Ecosystem
  The language or package manager is not supported.

  **HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

  **Exit Code:** N/A

  **Help Links:**

  - [https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support)
### [SNYK-OS-0002](#snyk-os-0002-unable-to-parse-manifest-file) - Unable to parse manifest file
  The provided manifest file could not be parsed

  **HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

  **Exit Code:** N/A

  **Help Links:**

  
---
# Builds
### [SNYK-OS-8001](#snyk-os-8001-invalid-request) - Invalid request
  The provided request payload is not valid for the selected ecosystem. Please review the API documentation

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  - [https://apidocs.snyk.io/](https://apidocs.snyk.io/)
### [SNYK-OS-8002](#snyk-os-8002-build-environment-not-found) - Build environment not found
  The build environment for the provided context could not be found. Please ensure you have created the build environment first.

  **HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

  **Exit Code:** N/A

  **Help Links:**

  
---
# [PURL Vulnerabilities](https://docs.snyk.io/introducing-snyk/getting-started-snyk-intel-vuln-db-access#about-the-snyk-vulnerability-database)
### [SNYK-OSSI-1040](#snyk-ossi-1040-your-organisation-is-not-authorised-to-perform-this-action) - Your organisation is not authorised to perform this action
  Your organisation is not authorised to perform this action

  **HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-1050](#snyk-ossi-1050-authorization-request-failure) - Authorization request failure
  Unexpected error when authenticating. Please try again, and if you continue to experience issues please contact support

  **HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2010](#snyk-ossi-2010-invalid-purl) - Invalid PURL
  Invalid PURL has been provided

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2011](#snyk-ossi-2011-package-requested-without-namespace) - Package requested without namespace
  You have requested a package type which requires a namespace (eg. maven group id). Please supply the namespace in order to retrieve the package correctly

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2020](#snyk-ossi-2020-unsupported-ecosystem) - Unsupported Ecosystem
  Ecosystem is not supported

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  - [https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support)
### [SNYK-OSSI-2021](#snyk-ossi-2021-purl-component(s)-required) - Purl component(s) required
  Currently we require a list of components of the package url specification. The purl supplied by the user did not specify all the components which we currently require

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2022](#snyk-ossi-2022-purl-component-not-supported) - Purl component not supported
  You have submitted a purl with components which are not supported

  **HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2030](#snyk-ossi-2030-package-not-found) - Package not found
  Requested package not found

  **HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2031](#snyk-ossi-2031-vulnerability-service-unavailable) - Vulnerability service unavailable
  Vulnerability service is currently not available. Please try again in a few minutes

  **HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2032](#snyk-ossi-2032-vulnerability-response-invalid) - Vulnerability response invalid
  An unexpected error occurred. Please try again, and if you continue to experience issues please contact support

  **HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2033](#snyk-ossi-2033-vulnerability-service-error) - Vulnerability service error
  An unexpected error occured with the vulnerability service. Please try again, and if you continue to experience issues please contact support

  **HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

  **Exit Code:** N/A

  **Help Links:**

  
### [SNYK-OSSI-2040](#snyk-ossi-2040-internal-server-error) - Internal server error
  An error was experienced by the service when processing the request.

  **HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

  **Exit Code:** N/A

  **Help Links:**

  
---

Genererated at 2022-09-02T10:15:10.026Z