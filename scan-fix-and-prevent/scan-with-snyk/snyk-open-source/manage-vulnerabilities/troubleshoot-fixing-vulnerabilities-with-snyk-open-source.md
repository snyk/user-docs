---
description: How to troubleshoot fixing vulnerabilities with Snyk Open Source
nav_context: agnostic
---

# Troubleshoot fixing vulnerabilities with Snyk Open Source

When you find a vulnerability, you have the opportunity to report that vulnerability to Snyk. For details, see [Disclosure of a vulnerability in an open-source package](https://docs.snyk.io/snyk-data-and-governance/disclosure-of-a-vulnerability-in-an-open-source-package).

## Unable to open a pull request or merge request for issues found by Snyk

When you import a Project, either through integration or by using the CLI, both CLI and SCM projects receive fix advice, while SCM projects additionally offer the option to open a Fix PR.

Snyk does not open PRs for transitive dependencies. For more information, see [Fixing transitive dependencies](vulnerability-fix-types.md#fixing-transitive-dependencies).

## Languages supported for Fix and Backlog Pull Requests or Merge Requests

Snyk can generate Fix Pull Requests (Fix PRs) or Merge Requests (MRs) for dependencies that may have a patch or an upgrade that will fix a vulnerability.

Snyk supports creating Fix PRs or MRs for the following languages:

* [Maven](https://docs.snyk.io/supported-languages/supported-languages-list/java-and-kotlin/git-repositories-with-maven-and-gradle#maven)
* [.NET (NuGet)](https://docs.snyk.io/supported-languages/supported-languages-list/.net)
* [npm](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#support-for-npm)
* [pnpm](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#support-for-pnpm)
* [Python](https://docs.snyk.io/supported-languages/supported-languages-list/python)
* [Ruby](https://docs.snyk.io/supported-languages/supported-languages-list/ruby)
* [Yarn](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#support-for-yarn)

## Languages supported for Upgrade Pull Requests or Merge Requests

Snyk can generate Upgrade Pull Requests (Upgrade PRs) or Merge Requests (MRs) for dependencies automatically when a new version is released.

Snyk supports generating Upgrade PRs or MRs for the following languages:

* [npm](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#support-for-npm)
* [pnpm](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#scanning-pnpm-workspaces)
* [Yarn](https://docs.snyk.io/supported-languages/supported-languages-list/javascript#support-for-yarn)
* [Maven](https://docs.snyk.io/supported-languages/supported-languages-list/java-and-kotlin/git-repositories-with-maven-and-gradle#maven)
