---
description: Reference for the read-only tools and limits available through Snyk Remote MCP
---

# Snyk Remote MCP tools

Snyk Remote MCP advertises 28 read-only tools. Raw list and get tools help you discover and inspect Snyk resources. Reporting and workflow tools resolve readable scopes, combine related data, prioritize evidence, and return both Markdown and structured content.

## Identity and scope discovery

| Tool                    | Description                                                                                                         |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `snyk_get_current_user` | Returns the identity represented by the current OAuth token and identifies a Snyk App installation when applicable. |
| `snyk_list_orgs`        | Lists one bounded page of accessible Snyk Organizations.                                                            |
| `snyk_get_org`          | Resolves an Organization by name, slug, or UUID and returns its metadata.                                           |
| `snyk_list_groups`      | Lists accessible Snyk Groups. Availability depends on Group access.                                                 |
| `snyk_get_group`        | Returns one accessible Group by ID.                                                                                 |

## Projects, targets, and collections

| Tool                    | Description                                                                                   |
| ----------------------- | --------------------------------------------------------------------------------------------- |
| `snyk_list_projects`    | Lists one bounded page of Projects in an Organization.                                        |
| `snyk_get_project`      | Returns full metadata for one Project.                                                        |
| `snyk_list_targets`     | Lists one bounded page of repositories and other import targets onboarded to an Organization. |
| `snyk_get_target`       | Returns full metadata for one target.                                                         |
| `snyk_list_collections` | Lists Project collections in an Organization, or the Projects in one collection.              |

## Issues, reports, and remediation workflows

| Tool                               | Description                                                                                                                                                                                                                                             |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_list_issues`                 | Lists one cursor-paginated page of Organization issues. Supports optional Project, status, severity, type, and ignored-state filters.                                                                                                                   |
| `snyk_get_issue`                   | Returns full details for one issue instance, including source location, coordinates, and remediation data.                                                                                                                                              |
| `snyk_get_issue_counts`            | Aggregates all matching Organization issue instances by severity, type, status, and Project. A positive `limit` bounds analysis; `project_limit` controls the size of the Project breakdown.                                                             |
| `snyk_get_group_issues`            | Lists one cursor-paginated page of issues across Organizations in a Group. Requires Group-level access.                                                                                                                                                 |
| `snyk_get_issues_report`           | Produces a Project-grouped Markdown and structured report enriched with severity, CVSS, risk score, remediation guidance, affected packages, and code data flows.                                                                                       |
| `snyk_get_project_risk_report`     | Resolves an Organization and Project by name or ID, then returns prioritized Project risk and remediation evidence.                                                                                                                                     |
| `snyk_get_security_posture_report` | Summarizes Organization issue-instance risk by severity, issue type, status, Project, fixability, exploit signal, and Snyk risk score.                                                                                                                  |
| `snyk_get_remediation_backlog`     | Returns a prioritized Organization- or Project-level backlog containing only issues with a concrete fix signal or remediation guidance.                                                                                                                 |
| `snyk_get_group_risk_rollup`       | Resolves a Group by name or ID and compares issue-instance risk across its Organizations.                                                                                                                                                               |
| `snyk_get_remediation_handoff`     | Builds a read-only Snyk Studio handoff for one issue with stable identity, observed repository and source evidence, remediation guidance, missing context, and a recommended local scan category.                                                       |
| `snyk_verify_issue_remediation`    | Checks the current remote state of an issue after updated scan results reach Snyk, using the original issue ID and, when needed, fallback identity. Returns `verified_resolved`, `still_open`, or `not_observable`.                                      |

The reporting tools default to open, non-ignored issues. The four tools ending in `risk_report`, `posture_report`, `backlog`, and `risk_rollup` support status, severity, issue-type, and ignored-state filters. The `snyk_get_issues_report` tool supports status and severity filters and always excludes ignored issues. The four report tools use `limit: 0` by default to analyze all matching pages. Use a positive `limit` to bound analysis.

The `top` argument controls how much prioritized evidence is returned. The default is 100 for Project risk and remediation backlog reports and 20 for security posture and Group risk rollup reports. Set `top: 0` to return all analyzed evidence.

For `snyk_get_issue_counts`, `limit` defaults to 0 and analyzes all matching pages. A positive value bounds the analysis; when that bound is reached, `truncated` is true and a warning identifies the counts as incomplete.

The `project_limit` argument defaults to 10 for issue counts, security posture reports, and remediation backlog reports. It controls only the size of the `by_project` breakdown; set it to 0 to return every matching Project entry. Responses include `by_project_total` and `by_project_truncated` so clients can identify a bounded breakdown. For remediation backlogs, `action_group_limit` defaults to 20 and controls the number of returned remediation action groups; set it to 0 to return all action groups.

{% hint style="warning" %}
Use positive `limit` and `top` values for large scopes to keep tool calls and AI assistant context bounded. Review the returned `coverage` and truncation fields before drawing conclusions from a partial report.
{% endhint %}

The `snyk_get_issues_report` tool analyzes all matching pages and returns up to 100 issue instances by default. Set `limit: 0` to return all analyzed issues, or use any positive value to bound the returned evidence.

For `snyk_get_group_risk_rollup`, use `probe_only: true` to determine whether direct Group issue access is available or the per-Organization fallback is required. If a successful fallback result encounters more than five Organization failures, `coverage.org_failures` contains the first five examples. Failure responses apply the same bound to their top-level `org_failures` field. Use `org_failures_total` and `org_failures_truncated` to assess the complete failure count.

### Remediation handoff and verification

The `snyk_get_remediation_handoff` tool requires an Organization name, slug, or UUID and an issue UUID. It returns the primary issue identity and a fallback identity based on the Project, rule, source or manifest path, and affected package. Repository, branch, and source details are included only when Snyk Project, target, or issue data contains them. If the Project has a `prodsec-commit-hash` tag, the value is returned as `handoff.project.commit_hash`; otherwise, `commit_reference` is identified as missing in `handoff.context`.

The handoff recommends a local Snyk Studio scan category and remediation steps, but it does not run Snyk Studio or modify Snyk data.

After a local fix and Snyk Studio scan have updated Snyk, use `snyk_verify_issue_remediation` with the handoff identity. The tool requires an Organization and either the original `issue_id` or both the fallback `project_id` and `rule`. Optional file and package fields narrow fallback correlation. The fallback `limit` defaults to 100 and accepts values from 10 through 1000. Pass `baseline_updated_at` from the handoff to determine whether the observed issue snapshot advanced.

Verification reports only remote Snyk state. It does not prove that a local scan ran, and an issue that cannot be found or correlated is `not_observable`, not `verified_resolved`.

## Packages, SBOMs, and dependencies

| Tool                             | Description                                                                                                                                                                                                                              |
| -------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_review_project_sbom`       | Reviews an existing Project SBOM and dependency snapshot, correlates open package-vulnerability issue instances, and returns component and ecosystem counts, direct-dependency evidence, package risk, data-quality warnings, and coverage without running a scan. |
| `snyk_get_project_sbom`          | Retrieves an existing Project SBOM as CycloneDX 1.4, 1.5, or 1.6 JSON, or SPDX 2.3 JSON.                                                                                                                                                 |
| `snyk_list_project_dependencies` | Retrieves the dependency graph captured by the latest monitored Project snapshot. This tool reads the Snyk v1 API because there is no equivalent REST endpoint.                                                                          |
| `snyk_test_package`              | Looks up direct vulnerabilities for one Package URL without onboarding a Project.                                                                                                                                                        |
| `snyk_test_packages`             | Looks up direct vulnerabilities for a batch of up to 100 Package URLs. This capability is not enabled for every Snyk customer.                                                                                                           |

If batch package testing is unavailable, `snyk_test_packages` returns `snyk_test_package` as the fallback tool. Empty SBOM and dependency results include data-quality warnings.

The `snyk_review_project_sbom` tool accepts an Organization and Project by name or UUID. Use `project_type`, `project_origin`, and `target_file` to disambiguate duplicate Project names. The `format` argument supports CycloneDX 1.4, 1.5, or 1.6 JSON and SPDX 2.3 JSON; the default is `cyclonedx1.4+json`. Its evidence controls are:

* `component_limit`: defaults to 100 and accepts values from 0 through 1000. A value of 0 returns no component records.
* `issue_limit`: defaults to 500 and accepts values from 10 through 5000. It bounds the open package-vulnerability issue instances analyzed.
* `top`: defaults to 20 and accepts values from 0 through 100. A value of 0 returns no prioritized issue records.

The summary counts a package issue as fixable when it has a concrete fix flag, remedy, or remediation hint. High-risk package issue instances have a Snyk risk score of at least 700.

Review `coverage` and `warnings` before treating component or issue counts as complete. This workflow reads existing monitored data and does not run a new scan.

## Inventory and governance

| Tool                         | Description                                                                                                                              |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_list_container_images` | Lists one bounded page of container images known to an Organization.                                                                     |
| `snyk_search_audit_logs`     | Searches read-only Organization audit logs by date, event, user, Project, and sort order. Access also depends on the caller's Snyk role. |

Audit log pagination is time-bucketed by the upstream API. A page can contain no records while `has_more` is true. Continue with `next_cursor` until `has_more` is false instead of treating one empty page as a complete result.

## Pagination and views

Paginated tools return:

* `has_more` to indicate whether another page is available
* `next_cursor` to use for the next call
* `view` to identify whether summary or full resources were returned

The default page size is 50. A custom page size must be between 10 and 100 and a multiple of 10. This is a per-page API constraint, not a total result limit. Continue with `next_cursor` until `has_more` is false to retrieve the complete collection.

Use `view: "summary"` for normal agent workflows. Use `view: "full"` when you need raw Snyk API fields that are not included in the summary.

## Response format and errors

Every tool returns MCP `structuredContent` and a JSON text block. Reporting and workflow tools also include a `markdown` field designed for concise model context.

Errors include a category, a retryability indicator, and a recovery hint. Categories distinguish invalid input, authentication, forbidden access, not found, rate limiting, upstream failure, and internal failure. When available, results also include the upstream HTTP status and Snyk JSON:API errors.

Responses identify incomplete results with fields such as:

* `has_more` and `next_cursor` for paginated collections
* `truncated` and `candidate_truncated` for bounded analysis or evidence
* `coverage` for analyzed records, returned evidence, failures, and count semantics
* `warnings` for permission gaps, missing enrichment, or empty monitored data

Report and workflow results use `context.observed` and `context.missing` where relevant to distinguish evidence returned by Snyk from relationships or context that were not available. Missing owner, deployment, repository, commit, or reachability context is not inferred.

Reporting aggregates count issue instances, where one issue UUID represents one finding on one scan item. The same vulnerability or rule can produce multiple issue instances across Projects.

## Rate limits

Tool rate limits use a fixed one-minute window for each authenticated access token.

| Requests per minute | Tools                                                                                                                                                                                                                                                                                                                                 |
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 60                  | `snyk_test_package`                                                                                                                                                                                                                                                                                                                   |
| 30                  | `snyk_list_orgs`, `snyk_get_org`, `snyk_list_projects`, `snyk_get_project`, `snyk_list_targets`, `snyk_get_target`, `snyk_list_collections`, `snyk_get_issue`, `snyk_get_current_user`                                                                                                                                                |
| 20                  | `snyk_list_issues`, `snyk_get_project_sbom`, `snyk_list_project_dependencies`, `snyk_test_packages`, `snyk_list_groups`, `snyk_get_group`, `snyk_list_container_images`, `snyk_verify_issue_remediation`                                                                                                                               |
| 10                  | `snyk_get_issue_counts`, `snyk_get_group_issues`, `snyk_search_audit_logs`, `snyk_get_issues_report`, `snyk_get_project_risk_report`, `snyk_get_security_posture_report`, `snyk_get_remediation_backlog`, `snyk_get_group_risk_rollup`, `snyk_get_remediation_handoff`, `snyk_review_project_sbom`                                         |

When a tool reaches its limit, it returns a `rate_limit` error with `retry_after_seconds`.

{% hint style="info" %}
The legacy `snyk_get_project_issues` tool name remains callable for compatibility but is not advertised. Use `snyk_list_issues` for new workflows.
{% endhint %}
