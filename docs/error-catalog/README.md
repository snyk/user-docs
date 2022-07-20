# Snyk Error Codes
  The error codes in the table below describe the codes that you may encounter while working with the [Snyk API](../snyk-api-info/README.md) or [CLI](../snyk-cli/README.md). When errors are encountered using the API, they will also have an appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). If you encounter errors without an error code, use the HTTP status code to determine the appropriate action.
  |Code|Title|Description|HTTP Status|Exit Code|Help|
  |---|---|---|---|---|---|
|[SNYK-OS-0001](#snyk-os-0001)|Unsupported language or package manager|The language or package manager is not supported|422|N/A|https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support|    
|[SNYK-OS-0002](#snyk-os-0002)|Unable to parse manifest file|The provided manifest file could not be parsed|422|N/A||    
|[SNYK-OSSI-0001](#snyk-ossi-0001)|Unauthorized|This is an example project error|401|N/A||    
|[SNYK-OSSI-1010](#snyk-ossi-1010)|Token not provided|Token has not been provided|400|N/A||    
|[SNYK-OSSI-1020](#snyk-ossi-1020)|Token is invalid|Token is invalid|401|N/A||    
|[SNYK-OSSI-1030](#snyk-ossi-1030)|Insufficient privileges|User has insufficient privileges|403|N/A||    
|[SNYK-OSSI-1040](#snyk-ossi-1040)|Organization not on allow list|Organization not on allow list|403|N/A||    
|[SNYK-OSSI-2010](#snyk-ossi-2010)|Invalid PURL|Invalid PURL has been provided|400|N/A||    
|[SNYK-OSSI-2020](#snyk-ossi-2020)|Unsupported Ecosystem|Ecosystem is not supported|422|N/A|https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support|    
|[SNYK-OSSI-2030](#snyk-ossi-2030)|Package not found|Requested package not found|404|N/A||    
|[SNYK-OSSI-2031](#snyk-ossi-2031)|Service down|Service is down|500|N/A||    
|[SNYK-OSSI-2032](#snyk-ossi-2032)|VulnDB error|Error querying VulnDB|500|N/A||    
|[SNYK-OSSI-2040](#snyk-ossi-2040)|Processing error|An error occurred processing the request|500|N/A||    
