# Snyk Error Codes
  The error codes in the table below describe the codes that you may encounter while working with the [Snyk API](../snyk-api-info/README.md) or [CLI](../snyk-cli/README.md). When errors are encountered using the API, they will also have an appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). If you encounter errors without an error code, use the HTTP status code to determine the appropriate action.
  ||Code|Title|Description|HTTP Status|Exit Code|Help|
  |---|---|---|---|---|---|---|
|Ecosystems|||||||
||[SNYK-OS-0001](#snyk-os-0001)|Unsupported Ecosystem|The language or package manager is not supported.|[422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)|N/A|https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support|    
||[SNYK-OS-0002](#snyk-os-0002)|Unable to parse manifest file|The provided manifest file could not be parsed|[422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)|N/A||    
|IsolatedBuilds|||||||
||[SNYK-OS-8001](#snyk-os-8001)|Invalid request|The provided request payload is not valid|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A||    
||[SNYK-OS-8002](#snyk-os-8002)|Too many requests|The service has received too many requests and will be throttled|[429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)|N/A||    
||[SNYK-OS-8003](#snyk-os-8003)|Unsupported ecosystem|The ecosystem that has been provided is not supported. Please review the documentation in the links provided.|[403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)|N/A|https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support|    
||[SNYK-OS-8004](#snyk-os-8004)|Build environment not found|The build environment for the provided context could not be found. Please ensure you have created the build environment first.|[404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)|N/A||    
||[SNYK-OS-8999](#snyk-os-8999)|Server error|An unexpected server error was encountered.|[500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)|N/A||    
|Snapshots|||||||
||[SNYK-OSSI-0001](#snyk-ossi-0001)|Unauthorized|This is an example project error|[401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)|N/A||    
|PurlVulnerabilityFetching|||||||
||[SNYK-OSSI-1040](#snyk-ossi-1040)|Your organisation is not authorised to perform this action|Your organisation is not authorised to perform this action|[403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)|N/A||    
||[SNYK-OSSI-1050](#snyk-ossi-1050)|Authorization request failure|Unexpected error when authenticating. Please try again, and if you continue to experience issues please contact support|[500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)|N/A||    
||[SNYK-OSSI-2010](#snyk-ossi-2010)|Invalid PURL|Invalid PURL has been provided|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A||    
||[SNYK-OSSI-2011](#snyk-ossi-2011)|Package requested without namespace|You have requested a package type which requires a namespace (eg. maven group id). Please supply the namespace in order to retrieve the package correctly|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A||    
||[SNYK-OSSI-2020](#snyk-ossi-2020)|Unsupported Ecosystem|Ecosystem is not supported|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A|https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support|    
||[SNYK-OSSI-2021](#snyk-ossi-2021)|Purl component(s) required|Currently we require a list of components of the package url specification. The purl supplied by the user did not specify all the components which we currently require|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A||    
||[SNYK-OSSI-2022](#snyk-ossi-2022)|Purl component not supported|You have submitted a purl with components which are not supported|[400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)|N/A||    
||[SNYK-OSSI-2030](#snyk-ossi-2030)|Package not found|Requested package not found|[404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)|N/A||    
||[SNYK-OSSI-2031](#snyk-ossi-2031)|Vulnerability service unavailable|Vulnerability service is currently not available. Please try again in a few minutes|[503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)|N/A||    
||[SNYK-OSSI-2032](#snyk-ossi-2032)|Vulnerability response invalid|An unexpected error occurred. Please try again, and if you continue to experience issues please contact support|[500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)|N/A||    
||[SNYK-OSSI-2033](#snyk-ossi-2033)|Vulnerability service error|An unexpected error occured with the vulnerability service. Please try again, and if you continue to experience issues please contact support|[500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)|N/A||    
||[SNYK-OSSI-2040](#snyk-ossi-2040)|Internal server error|An error was experienced by the service when processing the request.|[500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)|N/A||    
Genererated at 2022-08-22T10:14:05.405Z