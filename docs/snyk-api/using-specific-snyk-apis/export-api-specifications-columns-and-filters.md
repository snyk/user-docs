# Export API: Specifications, columns, and filters

The Export API, which Snyk Analytics supports, makes it easier to export data by allowing users to create and manage CSV files. These files are safely stored by Snyk. Designed for efficiency and security, the Export API helps users organize and scale the export of large datasets, which is useful for reporting and analytics tasks.

You can use the Export API to export the Snyk **issues** and **usage** **events** datasets in the scope of `Snyk Organization` or `Snyk Group`. Navigate to the [available columns and filters](export-api-specifications-columns-and-filters.md#available-columns-and-filters) section to see the full lists.

{% hint style="info" %}
Before running the first export, ensure that all API requests include:

* The API version parameter. The latest version is `2024-10-15`. You can also include the date of the current day for the version if you want to auto-upgrade when you use the API.
* The authorization header. Use a user or a service account Snyk API Token.
* The `dataset` parameter. The only valid values are `issues` or `usage`. This parameter is required to specify which dataset you want to export.
* At least one date filter (`introduced` for the `issues` dataset or `updated` for either)
{% endhint %}

## Data consumption process

The Export API includes three endpoints for each scope, where the scope can be `Snyk Organization` or `Snyk Group`. Use the following workflow to successfully run an export using this API.

1. **Initiate the export**\
   Start by initiating an export process. The response to that request returns the `export_id`.\
   Set the [filters and columns](export-api-specifications-columns-and-filters.md#available-columns-and-filters) based on your preferences.

```javascript
POST /groups/{group_id}/export
```

2. **Validate the export status**\
   Validate the status using the export status endpoint and the `export_id` returned in the previous step.\
   Available statuses:
   * `PENDING` - the export process is preparing to start
   * `STARTED` - the export process has started
   * `FINISHED` - the export process has completed successfully. If the process finishes successfully when this request arrives, the results data will be included in the response.
   * `ERROR` - the export process has failed

```javascript
GET /groups/{group_id}/jobs/export/{export_id}
```

3. **Fetch results**\
   After the export process returns the `FINISHED` status, fetch the exported files using the export result endpoint. Use the `export_id` returned in the first step.

```javascript
GET /groups/{group_id}/export/{export_id}
```

## Export API specifications

### Data freshness

The data provided by the Export API service updates approximately every two hours. Given the data freshness, cyclic exports should not be scheduled more frequently than once every two hours.

### Rate limits

The API is limited by:

* The export `POST` endpoint allows up to 20 export requests per hour, while the status checks and results retrieval are unlimited.

{% hint style="info" %}
Given that the data is typically refreshed every two hours, Snyk anticipates that the applied rate limits will allow comfortable consumption. Snyk recommends requesting an export per relevant Group once in a few hours or on a daily basis.
{% endhint %}

### Data retention

The exported CSV files will remain available in the designated S3 bucket for a period of three days.

{% hint style="danger" %}
While the files are accessible for three days, the self-signed link to retrieve the export results is available only for 60 minutes after its creation by default. Users can limit the link expiration by passing a value between 0 and 3600 to the `url_expiration_seconds` attribute.
{% endhint %}

## Available columns and filters

### Default columns

If the Export API call does not define the specific columns, the returned data includes all the available columns by default.

### Available filters

Although the requested filters are not case-sensitive, the values for those filters are case-sensitive.

Use the exact filter value as it appears in the Snyk Web UI. To clarify this requirement, case-sensitive filters are indicated in the table of available filters.

{% hint style="info" %}
At least one date filter (`introduced` for the `issues` dataset or `updated` for either) must be included in your request.
{% endhint %}

<table><thead><tr><th width="205.140625">Filter</th><th>Applicable Datasets</th><th width="366.3046875">Description</th></tr></thead><tbody><tr><td>updated (from and to)</td><td>issues, usage</td><td><p>The date and time of the last update that affected any attribute in the dataset.</p><p>Use this filter during cyclic exports to export only data that was updated since the last export.</p><p><br>Acceptable format: <code>YYYY-MM-DDTHH:MM:SSZ</code><br>(example: <code>2024-11-28T09:10:00Z</code>)</p></td></tr><tr><td>introduced (from and to)</td><td>issues</td><td>Date when the issue was introduced.<br>Acceptable format: <code>YYYY-MM-DDTHH:MM:SSZ</code><br>(example: <code>2024-11-28T09:10:00Z</code>)</td></tr><tr><td>orgs</td><td>issues, usage</td><td>Snyk <code>Organization ID</code> (available only for the Group endpoints).</td></tr><tr><td>environment</td><td>issues</td><td>The environment of the Project (case insensitive).</td></tr><tr><td>lifecycle</td><td>issues</td><td>The lifecycle of the Project (case insensitive).</td></tr><tr><td>product_name</td><td>issues</td><td>Name of the Snyk product that produced the issue (case sensitive).</td></tr><tr><td>project_type</td><td>issues</td><td>The scanning method to use for a particular Project (case sensitive).</td></tr><tr><td>project_tags</td><td>issues</td><td>All tags (as key:value pair) which have been assigned to this Project (case sensitive).</td></tr><tr><td>empty_project_tags</td><td>issues</td><td><p>Takes one of three values:</p><ul><li><code>include</code> - Includes issues with null/empty project tags when using the <code>project_tags</code> filter.</li><li><code>exclude</code> - Excludes issues with null/empty project tags. (Not necessary when using the <code>project_tags</code> filter, which implicitly filters these out.)</li><li><code>only</code> - Returns only issues with null/empty project tags. Request will return 0 issues if the <code>project_tags</code> filter is also populated.</li></ul></td></tr></tbody></table>

### Issues dataset columns

#### Available columns

<details>

<summary>A list of all columns that can be easily copied to the request body</summary>

```json
"PROBLEM_ID",
"GROUP_PUBLIC_ID",
"GROUP_DISPLAY_NAME",
"ORG_PUBLIC_ID",
"ORG_SLUG",
"EPSS_SCORE",
"EPSS_PERCENTILE",
"ISSUE_SEVERITY",
"SCORE",
"NVD_SCORE",
"PROBLEM_TITLE",
"CVE",
"CWE",
"VULN_DB_URL",
"NVD_SEVERITY",
"PROJECT_NAME",
"PROJECT_URL",
"PROJECT_PUBLIC_ID",
"PROJECT_IS_MONITORED",
"PROJECT_OWNER_USERNAME",
"PROJECT_ORIGIN",
"PROJECT_CRITICALITIES",
"PROJECT_ENVIRONMENTS",
"PROJECT_LIFECYCLES",
"PROJECT_TAGS",
"PROJECT_TARGET_DISPLAY_NAME",
"PROJECT_TARGET_REF",
"PROJECT_COLLECTIONS",
"EXPLOIT_MATURITY",
"COMPUTED_FIXABILITY",
"EXISTS_IN_DIRECT_DEPENDENCY",
"FIRST_INTRODUCED",
"LAST_INTRODUCED",
"LAST_IGNORED",
"LAST_RESOLVED",
"PRODUCT_NAME",
"ISSUE_URL",
"ISSUE_STATUS",
"ISSUE_TYPE",
"ISSUE_SUB_TYPE",
"FIXED_IN_AVAILABLE",
"FIXED_IN_VERSION",
"SEMVER_VULNERABLE_RANGE",
"INTRODUCTION_CATEGORY",
"VULNERABILITY_PUBLICATION_DATE",
"PACKAGE_NAME_AND_VERSION",
"SNYK_CVSS_SCORE",
"SNYK_CVSS_VECTOR",
"REACHABILITY",
"GROUP_SLUG",
"ORG_DISPLAY_NAME",
"PROJECT_TYPE",
"PROJECT_TYPE_DISPLAY_NAME",
"PROJECT_TEST_FREQUENCY",
"PROJECT_TARGET_RUNTIME",
"PROJECT_IS_PRIVATE_TARGET",
"PROJECT_TARGET_SOURCE_TYPE",
"PROJECT_TARGET_SOURCE_TYPE_DISPLAY_VALUE",
"PROJECT_TARGET_UPSTREAM_URL",
"PROJECT_TARGET_FILE",
"PROJECT_OWNER_EMAIL",
"ISSUE_DELETED_AT",
"PROJECT_DELETED_AT",
"GROUP_DELETED_AT",
"COMMIT_ID",
"FILE_PATH",
"CODE_REGION",
"CODE_REGION_DISPLAY_VALUE",
"ASSET_FINDING_ID",
"ASSET_ID",
"PARENT_ASSET_ID",
"ASSET_NAME",
"PARENT_ASSET_NAME",
"ASSET_CLASS",
"ASSET_TYPE",
"ASSET_TAGS",
"REPOSITORY_FRESHNESS",
"ASSET_APPLICATION",
"ASSET_OWNER",
"ASSET_CATEGORY",
"ASSET_CATALOG_NAME",
"ASSET_LIFECYCLE",
"JIRA_ISSUES",
"HAS_JIRA_ISSUE_ASSIGNED",
"LATEST_JIRA_ISSUE",
"UPDATED_AT"
```

</details>

#### **Severity**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>score</code></td><td>A score based on an analysis model. Priority score is released in General Availability, while Risk Score is in Early Access.</td></tr><tr><td><code>issue_severity</code></td><td>Indicates the assessed level of risk, <code>critical</code>, <code>high</code>, <code>medium</code>, or <code>low</code>.</td></tr><tr><td><code>snyk_cvss_score</code></td><td>The Snyk recommended Common Vulnerability Scoring System (CVSS) score.</td></tr><tr><td><code>nvd_severity</code></td><td>The severity of a vulnerability as rated by NVD.</td></tr><tr><td><code>nvd_score</code></td><td>The score of a vulnerability as calculated by NVD.</td></tr></tbody></table>

#### **Likelihood**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>exploit_maturity</code></td><td>Represents the legacy existence and maturity, as defined by Snyk, of public exploits validated by Snyk, for example, Mature or Proof of Concept.</td></tr><tr><td><code>exploit_maturity_cvss_v4</code></td><td>Represents the existence and maturity of public exploits validated by Snyk, using the CVSS v4 values: Not Defined, POC, Attacked.</td></tr><tr><td><code>snyk_cvss_vector</code></td><td>The vector string of the metric values used to determine the CVSS score.</td></tr><tr><td><code>epss_score</code></td><td>The probability of exploitation in the wild in the next 30 days.</td></tr><tr><td><code>epss_percentile</code></td><td>The proportion of all vulnerabilities with the same or lower EPSS score.</td></tr><tr><td><code>reachability</code></td><td>Indicates whether the issue is related to functions that are being called by the application and thus has a greater risk of exploitability.</td></tr><tr><td><code>project_is_private_target</code></td><td>Indicates whether the Target's source is private or publicly reachable.</td></tr></tbody></table>

#### **Impact**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>asset_class</code></td><td>The customer configured business criticality of the asset (A, most critical to D, least critical).</td></tr><tr><td><code>project_target_runtime</code></td><td>The environment in which the Target is executed and run.</td></tr><tr><td><code>project_criticalities</code></td><td>A Project attribute that indicates business criticality. For example, <code>low</code>, <code>medium</code>, <code>high</code>, <code>critical</code>.</td></tr><tr><td><code>project_lifecycles</code></td><td>A Project attribute, for example, <code>production</code>, <code>development</code>, <code>sandbox</code>.</td></tr><tr><td><code>asset_lifecycle</code></td><td>The lifecycle state of the asset.</td></tr><tr><td><code>project_environments</code></td><td>A Project attribute, for example, <code>frontend</code>, <code>backend</code>, <code>internal</code>, <code>external</code>, <code>mobile</code>, <code>saas</code>, <code>onprem</code>, <code>hosted</code>, <code>distributed</code>.</td></tr></tbody></table>

#### **SCA fixability**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>computed_fixability</code></td><td>Indicates whether the issue can be fixed based on the vulnerability remediation paths.</td></tr><tr><td><code>fixed_in_available</code></td><td>Indicates whether Is the given vulnerability fixed in a different version of the responsible source.</td></tr><tr><td><code>fixed_in_version</code></td><td>The first version in which a given vulnerability was fixed.</td></tr><tr><td><code>exists_in_direct_dependency</code></td><td>Indicates if the vulnerability exists in a direct dependency. If false, the vulnerability only exists in transitive dependencies.</td></tr></tbody></table>

#### **Snyk hierarchy**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>group_public_id</code></td><td>A universally unique identifier for a Group, assigned in the source database of the record.</td></tr><tr><td><code>org_public_id</code></td><td>A universally unique identifier for an Organization, assigned in the source database of the record.</td></tr><tr><td><code>group_display_name</code></td><td>The display name set for this Group.</td></tr><tr><td><code>group_slug</code></td><td>The name of the Group within Snyk.</td></tr><tr><td><code>org_display_name</code></td><td>The display name set for this Organization.</td></tr><tr><td><code>org_slug</code></td><td>The name for the Organization within Snyk.</td></tr></tbody></table>

**Issue context**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>problem_id</code></td><td>Snyk Vulnerability Database ID that uniquely identifies the vulnerability.</td></tr><tr><td><code>product_name</code></td><td>The Snyk product which initially identified the issue.</td></tr><tr><td><code>problem_title</code></td><td>Name of the Snyk discovered vulnerability.</td></tr><tr><td><code>vuln_db_url</code></td><td>URL which directs to the Snyk Vulnerability Database.</td></tr><tr><td><code>issue_type</code></td><td>Indicates whether the issue is related to a vulnerability, license, or configuration.</td></tr><tr><td><code>issue_sub_type</code></td><td>A more granular variation of issue type.</td></tr><tr><td><code>issue_url</code></td><td>URL that directs to the given Project instance of this vulnerability on the Snyk website.</td></tr><tr><td><code>issue_status</code></td><td>Indicates whether the issue is open, resolved, or ignored.</td></tr><tr><td><code>issue_severity</code></td><td>Indicates the assessed level of risk, <code>critical</code>, <code>high</code>, <code>medium</code>, or <code>low</code>.</td></tr><tr><td><code>commit_id</code></td><td>they can be uniquely identified. Snyk provides Commit ID only for Snyk Code issues.</td></tr><tr><td><code>file_path</code></td><td>The path to the file where Snyk Code identified the specific issue.</td></tr><tr><td><code>code_region</code></td><td>The line numbers and columns range where the issues were found within a file.</td></tr><tr><td><code>code_region_display_value</code></td><td>The display representation of the line numbers and columns range where the issues was found within a file.</td></tr><tr><td><code>asset_finding_id</code></td><td>A unique issue ID in the level of repository, only applicable for Snyk Code issue</td></tr><tr><td><code>cve</code></td><td>The CVE ID(s).</td></tr><tr><td><code>cwe</code></td><td>The CWE ID(s).</td></tr><tr><td><code>introduction_category</code></td><td>A classification generated by Snyk describing how an issue was introduced in the context of using Snyk products, such as <code>Baseline Issue</code>, <code>Non-Preventable Issue</code>, and <code>Preventable Issue</code>.</td></tr><tr><td><code>package_name_and_version</code></td><td>The associated package name and version of the vulnerability.</td></tr><tr><td><code>semver_vulnerable_range</code></td><td>The vulnerable range of package versions (based on semantic versioning).</td></tr><tr><td><code>vulnerability_publication_date</code></td><td>The date a given vulnerability was first published by Snyk.</td></tr><tr><td><code>has_jira_issue_assigned</code></td><td>Displays <code>true</code> when at least one Jira issue is assigned, otherwise displays <code>false</code>.</td></tr><tr><td><code>latest_jira_issue</code></td><td>The most recently created Jira Issue for this issue.</td></tr><tr><td><code>jira_issues</code></td><td>All Jira Issues ever created for this issue.</td></tr><tr><td><code>first_introduced</code></td><td>The timestamp of the first scan that identified the issue.</td></tr><tr><td><code>last_introduced</code></td><td>The most recent instance of an issue having been introduced (or reintroduced).</td></tr><tr><td><code>last_ignored</code></td><td>The most recent instance of an issue has been ignored within the Snyk product.</td></tr><tr><td><code>last_resolved</code></td><td>The most recent instance of an issue having been resolved.</td></tr><tr><td><code>issue_deleted_at</code></td><td>When the issue record was deleted from Snyk.</td></tr><tr><td><code>updated_at</code></td><td>When the issue or any related context was last updated.</td></tr></tbody></table>

**Project and Target context**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>project_public_id</code></td><td>A universally unique identifier for a Project, assigned in the source database or the record.</td></tr><tr><td><code>project_name</code></td><td>The name given to this Project, when added to Snyk.</td></tr><tr><td><code>project_url</code></td><td>The project URL in Snyk platform.</td></tr><tr><td><code>project_is_monitored</code></td><td>The Project is set to be actively monitored. By default, the API returns only monitored issues of the Project. To fetch issues of deactivated Projects, check the API parameters.</td></tr><tr><td><code>project_type</code></td><td>The scanning method to use for a particular Project, such as Static Application Security Testing (SAST) for scanning using Snyk Code, or Maven for a Maven Project using Snyk Open Source. This is part of the configuration for scanning.</td></tr><tr><td><code>project_type_display_name</code></td><td>A display name Snyk assigned to internal Project type values.</td></tr><tr><td><code>project_test_frequency</code></td><td>The frequency of testing for a given Project, for example, Daily, Weekly, and so on.</td></tr><tr><td><code>project_origin</code></td><td>The Origin defines the Target ecosystem, such as CLI, GitHub, or Kubernetes. Origins are a property of Targets.</td></tr><tr><td><code>project_target_ref</code></td><td>A reference that differentiates this Project, for example, a branch name or version. Projects having the same reference can be grouped based on that reference.</td></tr><tr><td><code>project_target_runtime</code></td><td>The environment in which the Target is executed and run.</td></tr><tr><td><code>project_target_display_name</code></td><td>A display name for the Project's Target.</td></tr><tr><td><code>project_is_private_target</code></td><td>Indicates whether the Target's source is private or publicly reachable</td></tr><tr><td><code>project_target_source_type</code></td><td>The hosting provider of a givenTarget, for example, <code>docker-hub</code>, <code>github</code>, and so on.</td></tr><tr><td><code>project_target_source_type_display_value</code></td><td>A display value that represents the grouping forTarget sources, for example, Source Control, Container Registry, and so on.</td></tr><tr><td><code>project_target_upstream_url</code></td><td>The URL that points to a Target's upstream source, such as a URL for a GitHub repository.</td></tr><tr><td><code>project_target_file</code></td><td>The full file path within a project that Snyk is targeting for security scanning, such as <code>/var/www/composer.lock</code>, <code>/app/package.json</code>, or other dependency manifest files.</td></tr><tr><td><code>project_criticalities</code></td><td>A Project attribute that indicates business criticality. For example, <code>low</code>, <code>medium</code>, <code>high</code>, <code>critical</code>.</td></tr><tr><td><code>project_lifecycles</code></td><td>A Project attribute, for example, <code>production</code>, <code>development</code>, <code>sandbox</code>.</td></tr><tr><td><code>project_environments</code></td><td>A Poject attribute, for example, <code>frontend</code>, <code>backend</code>, <code>internal</code>, <code>external</code>, <code>mobile</code>, <code>saas</code>, <code>onprem</code>, <code>hosted</code>, <code>distributed</code>.</td></tr><tr><td><code>project_collections</code></td><td>All Project collections to which this Project has been added.</td></tr><tr><td><code>project_tags</code></td><td>All tags which have been assigned to this Project.</td></tr><tr><td><code>project_owner_email</code></td><td>The email of the user assigned as the owner of this Project.</td></tr><tr><td><code>project_owner_username</code></td><td>The username of the user assigned as the owner of this Project.</td></tr></tbody></table>

**Assets and application context**

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>asset_id</code></td><td>Asset ID.</td></tr><tr><td><code>parent_asset_id</code></td><td>Parent Asset ID.</td></tr><tr><td><code>asset_name</code></td><td>The display name of the asset.</td></tr><tr><td><code>parent_asset_name</code></td><td>The display name of the parent asset.</td></tr><tr><td><code>asset_class</code></td><td>The customer configured business criticality of the asset (A, most critical to D, least critical).</td></tr><tr><td><code>asset_type</code></td><td>Specific type of the asset (Repository, Package, Container Image, Image Package, or Scanned Artifact).</td></tr><tr><td><code>asset_tags</code></td><td>Array of the tags that were assigned to the asset based on imported data or user input.</td></tr><tr><td><code>repository_freshness</code></td><td>The repository activity status based on the last commit date.</td></tr><tr><td><code>asset_application</code></td><td>The application or service that the asset is associated with.</td></tr><tr><td><code>asset_owner</code></td><td>The code owner of the asset, usually a development team.</td></tr><tr><td><code>asset_category</code></td><td>Category from integrated development platforms, such as Backstage and Roadie.</td></tr><tr><td><code>asset_catalog_name</code></td><td>The catalog name as mentioned in the application context (ServiceNow, DataDog, and so on).</td></tr><tr><td><code>asset_lifecycle</code></td><td>The lifecycle state of the asset.</td></tr></tbody></table>

### Usage events dataset columns

#### Available columns

<details>

<summary>A list of all columns that can be easily copied to the request body</summary>

```yaml
"ID",
"ORG_PUBLIC_ID",
"GROUP_PUBLIC_ID",
"PRODUCT_DISPLAY_NAME",
"INTERACTION_TYPE",
"INTERACTION_CATEGORIES",
"INTERACTION_TIMESTAMP",
"INTERACTION_STATUS",
"INTERACTION_STAGE",
"INTERACTION_EXIT_CODE",
"INTERACTION_TARGET_ID",
"RUNTIME_APPLICATION_NAME",
"RUNTIME_APPLICATION_VERSION",
"RUNTIME_APPLICATION_DATA_SCHEMA_VERSION",
"RUNTIME_PLATFORM_OS",
"RUNTIME_PLATFORM_ARCH",
"RUNTIME_ENVIRONMENT_NAME",
"ENVIRONMENT_DISPLAY_NAME",
"RUNTIME_ENVIRONMENT_VERSION",
"RUNTIME_INTEGRATION_NAME",
"RUNTIME_INTEGRATION_VERSION",
"RUNTIME_PERFORMANCE_DURATION_MS",
"USER_EMAIL",
"USER_NAME",
"ORG_DISPLAY_NAME",
"ORG_SLUG",
"GROUP_DISPLAY_NAME",
"GROUP_SLUG",
"ORG_DELETED_AT",
"GROUP_DELETED_AT",
"UPDATED_AT"
```

</details>

#### Interaction context

<table><thead><tr><th width="318.18359375">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>id</code></td><td>A unique identifier for an interaction event</td></tr><tr><td><code>product_display_name</code></td><td>The Snyk product used during this interaction, for example, Snyk Open Source, Snyk IaC, Snyk Code, Snyk Container.</td></tr><tr><td><code>interaction_type</code></td><td>The type of interaction, could be "Scan done". Scan Done indicates that a test was run no matter if the CLI or IDE ran it, other types can be freely chosen types.</td></tr><tr><td><code>interaction_categories</code></td><td>The category vector used to describe the interaction in detail, "oss", "test".</td></tr><tr><td><code>interaction_timestamp</code></td><td>When the interaction was started in UTC.</td></tr><tr><td><code>interaction_status</code></td><td>Status would be "success" or "failure", where success means the action was executed, while failure means it didn't run.</td></tr><tr><td><code>interaction_stage</code></td><td>The stage of the SDLC where the Interaction occurred, such as "dev"|"cicd"|"prchecks"|"unknown".</td></tr><tr><td><code>interaction_exit_code</code></td><td>The interaction's exit code as returned by the running process. More info about the exit codes and their meaning is available in Snyk Docs per a given interaction (test, monitor, and so on).</td></tr><tr><td><code>interaction_target_id</code></td><td>A purl is a URL composed of seven components. scheme:type/namespace/name@version?qualifiers#subpath The purl specification is available here: <code>https://github.com/package-url/purl-spec</code> Some purl examples <code>pkg:github/package-url/purl-spec@244fd47e07d1004f0aed9c</code> <code>pkg:npm/%40angular/animation@12.3.1</code> <code>pkg:pypi/django@1.11.1</code></td></tr><tr><td><code>updated_at</code></td><td>When the interaction event or any related context was last updated.</td></tr></tbody></table>

#### Runtime context

<table><thead><tr><th width="330.3359375">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>runtime_application_name</code></td><td>The application used to execute a snyk interaction, for example, PyCharm, Visual Studio, snyk-ls, snyk-cli.</td></tr><tr><td><code>runtime_application_version</code></td><td>The version of the integration.</td></tr><tr><td><code>runtime_application_data_schema_version</code></td><td>The data schema version of Snyk's runtime interactions. The current version (v2) was released in Q2 2024. Prior versions' data may behave differently.</td></tr><tr><td><code>runtime_platform_os</code></td><td>The operating system for the integration (darwin, windows, linux, etc).</td></tr><tr><td><code>runtime_platform_arch</code></td><td>The architecture for the integration (AMD64, ARM64, 386, ALPINE).</td></tr><tr><td><code>runtime_environment_name</code></td><td>The environment for the integration (e.g., IntelliJ Ultimate, Pycharm).</td></tr><tr><td><code>environment_display_name</code></td><td>The Environment used during this interaction, for example: CLI, Eclipse, Jetbrains IDE, Visual Studio, Visual Studio Code, or Other</td></tr><tr><td><code>runtime_environment_version</code></td><td>The version of the integration environment (e.g. 2023.3)</td></tr><tr><td><code>runtime_integration_name</code></td><td>The name of the integration, could be a plugin or extension.</td></tr><tr><td><code>runtime_integration_version</code></td><td>The version of the integration, for example: 2.3.4.</td></tr><tr><td><code>runtime_performance_duration_ms</code></td><td>The duration in milliseconds of the interaction</td></tr></tbody></table>

#### Snyk hierarchy

<table><thead><tr><th width="318">Column name</th><th>Description</th></tr></thead><tbody><tr><td><code>group_public_id</code></td><td>A universally unique identifier for a Group, assigned in the source database of the record.</td></tr><tr><td><code>org_public_id</code></td><td>A universally unique identifier for an Organization, assigned in the source database of the record.</td></tr><tr><td><code>group_display_name</code></td><td>The display name set for this Group.</td></tr><tr><td><code>group_slug</code></td><td>The name of the Group within Snyk.</td></tr><tr><td><code>org_display_name</code></td><td>The display name set for this Organization.</td></tr><tr><td><code>org_slug</code></td><td>The name for the Organization within Snyk.</td></tr><tr><td><code>user_email</code></td><td>The email of the user who was authenticated during the interaction.</td></tr><tr><td><code>user_name</code></td><td>The name of the user who was authenticated during the interaction.</td></tr></tbody></table>
