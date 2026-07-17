---
description: Reference for the read-only tools and limits available through Snyk Remote MCP
---

# Snyk Remote MCP tools

Snyk Remote MCP advertises 25 read-only tools. Raw list and get tools help you discover and inspect Snyk resources. Reporting tools resolve readable scopes, combine related data, prioritize evidence, and return both Markdown and structured content.

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

## Issues and reports

| Tool                               | Description                                                                                                                                                       |
| ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_list_issues`                 | Lists one cursor-paginated page of Organization issues. Supports optional Project, status, severity, type, and ignored-state filters.                             |
| `snyk_get_issue`                   | Returns full details for one issue instance, including source location, coordinates, and remediation data.                                                        |
| `snyk_get_issue_counts`            | Aggregates all matching Organization issue instances by severity, type, status, and Project. A positive `limit` can optionally bound the analysis.                 |
| `snyk_get_group_issues`            | Lists one cursor-paginated page of issues across Organizations in a Group. Requires Group-level access.                                                           |
| `snyk_get_issues_report`           | Produces a Project-grouped Markdown and structured report enriched with severity, CVSS, risk score, remediation guidance, affected packages, and code data flows. |
| `snyk_get_project_risk_report`     | Resolves an Organization and Project by name or ID, then returns prioritized Project risk and remediation evidence.                                               |
| `snyk_get_security_posture_report` | Summarizes Organization issue-instance risk by severity, issue type, status, Project, fixability, exploit signal, and Snyk risk score.                            |
| `snyk_get_remediation_backlog`     | Returns a prioritized Organization- or Project-level backlog containing only issues with a concrete fix signal or remediation guidance.                           |
| `snyk_get_group_risk_rollup`       | Resolves a Group by name or ID and compares issue-instance risk across its Organizations.                                                                         |

The reporting tools default to open, non-ignored issues. The four tools ending in `risk_report`, `posture_report`, `backlog`, and `risk_rollup` support status, severity, issue-type, and ignored-state filters. They use `limit: 0` by default to analyze all matching pages. Use a positive `limit` to bound analysis.

The `top` argument controls how much prioritized evidence is returned. The default is 100 for Project risk and remediation backlog reports and 20 for security posture and Group risk rollup reports. Set `top: 0` to return all analyzed evidence.

{% hint style="warning" %}
Use positive `limit` and `top` values for large scopes to keep tool calls and AI assistant context bounded. Review the returned `coverage` and truncation fields before drawing conclusions from a partial report.
{% endhint %}

The `snyk_get_issues_report` tool analyzes all matching pages and returns up to 100 issue instances by default. Set `limit: 0` to return all analyzed issues, or use any positive value to bound the returned evidence.

## Packages, SBOMs, and dependencies

| Tool                             | Description                                                                                                                                                     |
| -------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_get_project_sbom`          | Retrieves an existing Project SBOM as CycloneDX 1.4, 1.5, or 1.6 JSON, or SPDX 2.3 JSON.                                                                        |
| `snyk_list_project_dependencies` | Retrieves the dependency graph captured by the latest monitored Project snapshot. This tool reads the Snyk v1 API because there is no equivalent REST endpoint. |
| `snyk_test_package`              | Looks up direct vulnerabilities for one Package URL without onboarding a Project.                                                                               |
| `snyk_test_packages`             | Looks up direct vulnerabilities for a batch of up to 100 Package URLs. This capability is not enabled for every Snyk customer.                                  |

If batch package testing is unavailable, `snyk_test_packages` returns `snyk_test_package` as the fallback tool. Empty SBOM and dependency results include data-quality warnings.

## Inventory and governance

| Tool                         | Description                                                                                                                              |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `snyk_list_container_images` | Lists one bounded page of container images known to an Organization.                                                                     |
| `snyk_search_audit_logs`     | Searches read-only Organization audit logs by date, event, user, Project, and sort order. Access also depends on the caller's Snyk role. |

## Pagination and views

List tools return:

* `has_more` to indicate whether another page is available
* `next_cursor` to use for the next call
* `view` to identify whether summary or full resources were returned

The default page size is 50. A custom page size must be between 10 and 100 and a multiple of 10. This is a per-page API constraint, not a total result limit. Continue with `next_cursor` until `has_more` is false to retrieve the complete collection.

Use `view: "summary"` for normal agent workflows. Use `view: "full"` when you need raw Snyk API fields that are not included in the summary.

## Response format and errors

Every tool returns MCP `structuredContent` and a JSON text block. Reporting tools also include a `markdown` field designed for concise model context.

Errors include a category, a retryability indicator, and a recovery hint. Categories distinguish invalid input, authentication, forbidden access, not found, rate limiting, upstream failure, and internal failure. When available, results also include the upstream HTTP status and Snyk JSON:API errors.

Responses identify incomplete results with fields such as:

* `has_more` and `next_cursor` for paginated collections
* `truncated` and `candidate_truncated` for bounded analysis or evidence
* `coverage` for analyzed records, returned evidence, failures, and count semantics
* `warnings` for permission gaps, missing enrichment, or empty monitored data

Reporting aggregates count issue instances, where one issue UUID represents one finding on one scan item. The same vulnerability or rule can produce multiple issue instances across Projects.

## Rate limits

Tool rate limits use a fixed one-minute window for each authenticated access token.

| Requests per minute | Tools                                                                                                                                                                                                                                  |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 60                  | `snyk_test_package`                                                                                                                                                                                                                    |
| 30                  | `snyk_list_orgs`, `snyk_get_org`, `snyk_list_projects`, `snyk_get_project`, `snyk_list_targets`, `snyk_get_target`, `snyk_list_collections`, `snyk_get_issue`, `snyk_get_current_user`                                                 |
| 20                  | `snyk_list_issues`, `snyk_get_project_sbom`, `snyk_list_project_dependencies`, `snyk_test_packages`, `snyk_list_groups`, `snyk_get_group`, `snyk_list_container_images`                                                                |
| 10                  | `snyk_get_issue_counts`, `snyk_get_group_issues`, `snyk_search_audit_logs`, `snyk_get_issues_report`, `snyk_get_project_risk_report`, `snyk_get_security_posture_report`, `snyk_get_remediation_backlog`, `snyk_get_group_risk_rollup` |

When a tool reaches its limit, it returns a `rate_limit` error with `retry_after_seconds`.

{% hint style="info" %}
The legacy `snyk_get_project_issues` tool name remains callable for compatibility but is not advertised. Use `snyk_list_issues` for new workflows.
{% endhint %}
