# IDE and CLI usage telemetry

Snyk Language Server collects usage telemetry after each successful test, through the Snyk proprietary Analytics service. Data does not pass through any third-party services.

The request body contains several categories of data:

* **Versions and types of integration**, for example, Language Server version, IDE type and version, plugin type and version
* **Local environment**: OS, architecture
* **Project data**, for example, repository git URL, branch name
* **Performance data**: duration of the scan in milliseconds
* **Event details**, for example, event type, timestamp, status, number of findings

The API path contains a unique Snyk Organization ID. Thus, data is annotated with Organization information.

Personally identifiable information (PII) for users is not part of the request. The Snyk internal analytics service API requires an authentication token to be present. Incoming event data is merged with user data inside the backend service, to allow Snyk to build a special report page: [Developer IDE and CLI usage](../../../manage-risk/reporting/available-snyk-reports.md#developer-ide-and-cli-usage).

You can see the open-source logic and examine it directly in the GitHub repository: [snyk/go-application-framework](https://github.com/snyk/go-application-framework/blob/main/internal/api/clients/analytics_v20240307.go). This is useful when a security audit is required.

