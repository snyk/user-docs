# Scan with Snyk

You can use Snyk to scan and secure your codebase and cloud infrastructure configurations, taking advantage of the Snyk capabilities in Static Application Security Testing (SAST), Software Composition Analysis (SCA), and  Infrastructure as Code analysis.

For more information, see [Scanning overview](scanning-overview/) and [Start scanning using the CLI, Web UI, or API](start-scanning-using-the-cli-web-ui-or-api.md).

## Select scanning methods

Implement a workflow to secure your code and infrastructure in your environment using the scanning methods provided by Snyk products:

* [Snyk Open Source](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup.md): Find and fix known vulnerabilities and licensing issues in your open-source dependencies.  \
  More information: [Open Source Security Explained](https://snyk.io/series/open-source-security/).
* [Snyk Code](snyk-code/): Scan your codebase for known vulnerabilities and get remediation guidance either inline in your IDE or by importing your code repository to Snyk Web UI. \
  More information: [Exploring the advanced technologies behind Snyk Code](https://snyk.io/blog/advanced-technologies-behind-snyk-code/).
* [Snyk Container](snyk-container/): Find and automatically fix container and workload vulnerabilities.
* [Scan infrastructure](scan-infrastructure/): Secure cloud infrastructure configurations before and after deployment.

## Manage issues

Snyk has several features that help you determine which issues are the most important for you to fix and the sequence in which to fix the issues.

See [Find and manage priority issues](find-and-manage-priority-issues/) for details.

## Run pull request checks

Scan and automatically address potential vulnerabilities when you review pull requests (PRs), to prevent security issues in production.

For more information, see [Run PR checks](run-pr-checks/).
