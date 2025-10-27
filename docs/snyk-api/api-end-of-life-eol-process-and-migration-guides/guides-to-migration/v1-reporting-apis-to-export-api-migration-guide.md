# V1 Reporting APIs to Export API migration guide

{% hint style="warning" %}
**End of life**

The v1 Reporting API is being deprecated, and future development and support are focused on the [Dataset Export API](../../using-specific-snyk-apis/export-api-specifications-columns-and-filters.md) and the [Issues REST API](../../reference/issues.md). Snyk recommends migrating to the new [Export API](../../using-specific-snyk-apis/export-api-specifications-columns-and-filters.md).
{% endhint %}

This guide outlines the migration path for all routes in the legacy v1 Reporting API.

### V1 Reporting legacy API to New Export API migration guide

The most significant change is the shift from a direct, synchronous request/response model (v1) to an asynchronous job-based model (Export API).

| Feature            | Legacy v1 Reporting API                          | New Export API                                                                                                                                                                                                                                                                                                                 |
| ------------------ | ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Endpoint Structure | Direct resource endpoints (/v1/reporting/issues) | Scoped job endpoints. The new API is focused on the Organization or Group scope, with Organization ID or Group ID as path parameters                                                                                                                                                                                           |
| Response Format    | JSON in the response body                        | CSV or JSON files using a signed URL                                                                                                                                                                                                                                                                                           |
| Data Retrieval     | Synchronous, immediate response                  | The Export API offers a robust, scalable solution for data export, providing results as CSV files stored securely and accessible using a self-signed link. It uses an asynchronous job pattern (Initiate, Check Status, Fetch Results) for large datasets, replacing direct, synchronous POST requests for large data volumes. |
| Rate Limit         | 70 requests per minute, per user                 | 20 export POST requests per hour (status/results checks are unlimited)                                                                                                                                                                                                                                                         |
| Required Headers   | Authorization                                    | Authorization, Version (for example `2024-10-15`)                                                                                                                                                                                                                                                                              |

#### New export workflow

To replace any v1 Reporting API call, follow these steps:

1. Initiate the Export:&#x20;
   1. Send a `POST` request with your filters to create an export job. You will receive an `export_id`.
   2. Endpoint: `POST /groups/{group_id}/export`
   3. Required Parameters: The body must include the filters (including at least one date filter: introduced or updated) and specify the dataset (issues or usage) in the request header.
2. Validate the Export Status:
   1. Poll the status endpoint using the `export_id` until the status is `FINISHED`.
   2. Endpoint: `GET /groups/{group_id}/jobs/export/{export_id}`
   3. Statuses: `PENDING`, `STARTED`, `FINISHED`, `ERROR`.
3. Fetch Results in a CSV:
   1. After `FINISHED`, fetch the results, which will return a self-signed URL to the CSV file.
   2. Endpoint: `GET /groups/{group_id}/export/{export_id}`

{% hint style="info" %}
The self-signed link is valid for 60 minutes by default, and can be configured between 0 to 60 minutes.
{% endhint %}

#### Route-specific migration

The v1 routes for issues and Project counts now primarily map to the Export API's issues dataset, while test counts map to the usage dataset.

<table><thead><tr><th>Legacy v1 Route</th><th width="147.451416015625">Purpose</th><th width="211.6319580078125">New Export API Mapping</th><th>Notes</th></tr></thead><tbody><tr><td><code>POST /v1/reporting/issues</code></td><td>Get list of issues in a timeframe</td><td>Dataset: issues</td><td>Use introduced (from/to) and/or updated (from/to) filters to replicate the timeframe logic.</td></tr><tr><td><code>POST /v1/reporting/issues/latest</code></td><td>Get list of latest issues</td><td>Dataset: issues</td><td>Use only the updated (from) filter with a recent date if you need issues updated since a specific point.</td></tr><tr><td><code>POST /v1/reporting/counts/issues</code></td><td>Get issue counts in a timeframe</td><td>Dataset: issues</td><td>Counting logic is now handled by processing the full exported CSV file. Use introduced and/or updated filters to replicate the timeframe.</td></tr><tr><td><code>POST /v1/reporting/counts/issues/latest</code></td><td>Get latest issue counts</td><td>Dataset: issues</td><td>Counting logic is now handled by processing the full exported CSV file.</td></tr><tr><td><code>POST /v1/reporting/counts/projects</code></td><td>Get project counts in a timeframe</td><td>Dataset: issues</td><td>Count logic is now handled by processing the full exported CSV file. Use project_public_id or project_name columns in the CSV for project identification.</td></tr><tr><td><code>POST /v1/reporting/counts/projects/latest</code></td><td>Get latest project counts</td><td>Dataset: issues</td><td>Count logic is now handled by processing the full exported CSV file.</td></tr><tr><td><code>POST /v1/reporting/counts/tests</code></td><td>Get test counts in a timeframe</td><td>Dataset: usage</td><td>Counting logic is now handled by processing the full exported CSV file, filtering by interaction_type: "Scan done". Use updated (from/to) to manage the timeframe.</td></tr></tbody></table>

#### Important considerations for count endpoints

The v1 issue count endpoints (`/v1/reporting/counts/issues` and `/latest`) relied on a legacy data model and do not accurately reflect the reality of all Snyk products.

Key issue:

* The legacy v1 issue count endpoints do not include issues from Snyk Code, and do not include issues generated by the most recent Snyk Infrastructure as Code (IaC) engines.
* Action: Migrating to the Export API with the issues dataset will provide a more complete and accurate count of all issues across all Snyk products, including IaC. Any count migration should be aware that the new number will likely be higher and more representative of your total issue landscape.

#### Filter parameter migration

The Export API uses filters in the initial `POST /export` request body.

{% hint style="info" %}
Filter values are case-sensitive.
{% endhint %}

| Filter                   | Applicable Datasets | Description                                                                                                    |
| ------------------------ | ------------------- | -------------------------------------------------------------------------------------------------------------- |
| updated (from and to)    | issues, usage       | The date and time of the last update that affected any attribute in the dataset. Required (or use introduced). |
| introduced (from and to) | issues              | The date and time when the issue was introduced. Required (or use updated).                                    |
| orgs                     | issues, usage       | Snyk Organization ID(s). Only available for the Group endpoints.                                               |
| environment              | issues              | The environment of the Project (for example, `external`).                                                      |
| lifecycle                | issues              | The lifecycle of the Project (for example, `production`).                                                      |
| product\_name            | issues              | Name of the Snyk product that produced the issue (for example, `Snyk IaC`).                                    |
| project\_type            | issues              | The scanning method used for the Project (for example, `npm`, `maven`). Case-sensitive value.                  |
| project\_tags            | issues              | Project tags as a key:value pair. Case-sensitive value.                                                        |

For any v1 filter not listed above (for example, `severity`, `fixable`, `languages`), the new process is to download the full CSV and apply the necessary filtering logic in your application or data warehouse using the corresponding column.

| v1 Filter Name              | Export API Column (for Post-Filtering) |
| --------------------------- | -------------------------------------- |
| severity                    | issue\_severity                        |
| languages                   | project\_type                          |
| ignored, isFixed, status    | issue\_status                          |
| fixable, isUpgradable, etc. | computed\_fixability                   |
| projects                    | project\_public\_id                    |

#### CSV to JSON Conversion Tool

Since the Export API returns data as CSV, and the legacy API returned JSON, you will likely need to convert the exported file for structured consumption by your application.

**Native JSON output**&#x20;

We are prioritizing adding the ability to export data directly in JSON format. It will be available early November. This feature will simplify data consumption for users migrating from the v1 API.

**Current workaround script**

Provided is an example of Python script [csv\_to\_json\_tool.py](https://gist.github.com/asabushaban/188ebc7516a0c28f2addec284bbc73fb) that automates downloading the CSV from the signed URL, converting it to JSON, and optionally formatting it using the command-line tool `jq`.

**Prerequisites for the script:**

* Python: The script requires Python 3.x.
* Libraries: requests (pip install requests).

{% hint style="info" %}
The command-line tool `jq` is recommended for pretty-printing the final JSON file.
{% endhint %}

### Alternative: Issues REST API for transactional data

While the Export API is the direct replacement for v1 reporting, the new [Issues REST API](../../reference/issues.md) is available for developers needing real-time, transactional access to issues (as JSON responses) for integration purposes, rather than bulk CSV reports.

| Endpoint                                | Method | Purpose                                                                                                              |
| --------------------------------------- | ------ | -------------------------------------------------------------------------------------------------------------------- |
| `/rest/orgs/{org_id}/issues`            | `GET`  | Get a paginated list of all issues for an Organization.                                                              |
| `/groups/{group_id}/issues`             | `GET`  | Get a paginated list of all issues for a Group.                                                                      |
| `/rest/orgs/{org_id}/issues/{issue_id}` | `GET`  | Get details for a single issue.                                                                                      |
| `/orgs/{org_id}/packages/{purl}/issues` | `GET`  | Query issues for a specific package version identified by Package URL (purl). (Returns direct vulnerabilities only.) |
| `/orgs/{org_id}/packages/issues`        | `POST` | Query issues for a batch of packages by purl (not available to all customers).                                       |

These endpoints support various query parameters for filtering (for example `effective_severity_level`, `type`, `status`, date ranges like `updated_after`), use cursor-based pagination (`starting_after`, `ending_before`), and return data in a structured JSON format.
