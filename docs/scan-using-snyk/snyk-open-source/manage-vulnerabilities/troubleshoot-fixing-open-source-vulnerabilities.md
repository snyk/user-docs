# Troubleshoot fixing Open Source vulnerabilities

## Unable to open a pull request or merge request for issues found by Snyk

When you import a Project, either through an integration or by using the CLI, Snyk provides the Fix PR button for direct dependencies only. Snyk does not open PRs for transitive dependencies. For more information, see [Fixing transitive dependencies](vulnerability-fix-types.md#fixing-transitive-dependencies).

## Languages supported for Fix Pull Requests or Merge Requests

Snyk can generate Fix Pull Requests (Fix PRs) or Merge Requests (MRs) for dependencies that may have a patch or an upgrade that will fix a vulnerability.

Snyk supports creating Fix PRs or MRs for the following languages:

* [Maven](../../supported-languages-and-frameworks/java-and-kotlin/best-practices-for-java-and-kotlin.md#maven)
* [.NET](../../../scan-applications/supported-languages-and-frameworks/.net/)
* [npm](../../supported-languages-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#npm)
* [Python](../../../scan-applications/supported-languages-and-frameworks/python.md)
* [Ruby](../../../scan-applications/supported-languages-and-frameworks/ruby.md)
* [Yarn](../../supported-languages-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#yarn)

