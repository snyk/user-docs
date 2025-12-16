# Troubleshoot fixing vulnerabilities with Snyk Open Source

When you find a vulnerability, you have the opportunity to report that vulnerability to Snyk. For details, see [Disclosure of a vulnerability in an open-source package](../../../snyk-data-and-governance/disclosure-of-a-vulnerability-in-an-open-source-package.md).

## Unable to open a pull request or merge request for issues found by Snyk

When you import a Project, either through integration or by using the CLI, both CLI and SCM projects receive fix advice, while SCM projects additionally offer the option to open a Fix PR.

Snyk does not open PRs for transitive dependencies. For more information, see [Fixing transitive dependencies](vulnerability-fix-types.md#fixing-transitive-dependencies).

## Languages supported for Fix Pull Requests or Merge Requests

Snyk can generate Fix Pull Requests (Fix PRs) or Merge Requests (MRs) for dependencies that may have a patch or an upgrade that will fix a vulnerability.

Snyk supports creating Fix PRs or MRs for the following languages:

* [Maven](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#maven)
* [.NET](../../../supported-languages/supported-languages-list/.net/)
* [npm](/broken/pages/kJOeCbuqexRTuTQPx6aZ#npm)
* [Python](../../../supported-languages/supported-languages-list/python/)
* [Ruby](../../../supported-languages/supported-languages-list/ruby.md)
* [Yarn](/broken/pages/kJOeCbuqexRTuTQPx6aZ#yarn)
