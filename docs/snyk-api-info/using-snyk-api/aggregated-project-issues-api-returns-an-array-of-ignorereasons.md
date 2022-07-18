# Aggregated project issues API returns an array of ignoreReasons

The Snyk API v1 endpoint [List all aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues) returns an array of `ignoreReasons` for each vulnerability.

This happens because ignores via CLI and API are path-based and thus potentially could have different `ignoreReasons` for different paths. Because the List all aggregated issues endpoint returns only one issue for all paths, the entire set of reasons is returned.

Snyk groups issues together by their identifier, so one response for the List all aggregated issues endpoint could correspond to the same issue across multiple paths. Thus the `ignoredReason` is across all issues that are aggregated and applies to that single grouped issue.
