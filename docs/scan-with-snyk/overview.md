# Overview

You can use Snyk to scan and secure your codebase and cloud infrastructure configurations, taking advantage of the Snyk capabilities in Static Application Security Testing (SAST), Software Composition Analysis (SCA), and  Infrastructure as Code analysis.

## Select scanning methods

Snyk supports scanning methods that correspond to Snyk products. Choose the right scanning method for the job you want to do, to find and fix issues not only early in the Software Development Life Cycle, but also after your web application is live.

* [Snyk Open Source](snyk-open-source/): scan your open-source libraries for vulnerabilities and license issues.\
  For more information, see [Open Source Security Explained](https://snyk.io/series/open-source-security/).
* [Snyk Code](snyk-code/): scan your code for security vulnerabilities using source code analysis.\
  For more information, see [Exploring the advanced technologies behind Snyk Code](https://snyk.io/blog/advanced-technologies-behind-snyk-code/).
* [Snyk Container](snyk-container/): scan for container image and workload vulnerabilities.
* [Snyk Infrastructure as Code: ](snyk-iac/)scan for issues in your cloud infrastructure configurations, before and after deployment.
* [Snyk API & Web](snyk-api-web/): discover and test the security of all your APIs and web apps, including those AI-generated. See the [product page](https://snyk.io/product/dast-api-web/) and the [Developers documentation](https://developers.probely.com/).

## Run pull request checks

Scan and automatically address potential vulnerabilities when you review pull requests (PRs), to prevent security issues in production, for your open-source libraries and your own code.

Snyk can also retest and alert on the default branch on a scheduled basis and show results.&#x20;

For more information, see [Run PR checks](pull-requests/pull-request-checks/).

{% hint style="info" %}
Scans may be limited on your account, depending on your[ Pricing Plan](../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md). For more information, see [What counts as a test?](../snyk-data-and-governance/what-counts-as-a-test.md)
{% endhint %}

## The scanning process

Snyk takes a developer-first approach to secure your development work by integrating directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit. This approach allows you to:

* Use Snyk to focus on early enablement, not later enforcement.
* Run scans while working on a Project, before you commit any code. This minimizes rework by finding issues that require changes early on.
* Add and test packages before writing the code that interfaces with each package.
* After writing a major section of code, scan it to find issues before continuing work.

### Learn how to design secure code

The following resources are available for all users:

* [Snyk Advisor](https://snyk.io/advisor): Helps you pick healthy open-source packages or base images to develop code.
* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely, and provides training on how to use Snyk.

### Write and deploy your code

* Using the [Snyk CLI](../developer-tools/snyk-cli/), you can scan locally on your machine. This is useful in scanning open-source and static code as well as containers and infrastructure as code configurations, including complex files that are templated with variables, such as Terraform plan files.
* Using [Snyk IDE Plugins](../developer-tools/snyk-ide-plugins-and-extensions/), you can test your open-source packages, first-party code, and infrastructure as code (IaC) Kubernetes deployment files in your development environment as you create your Project.
* Using [Git integrations](../developer-tools/scm-integrations/organization-level-integrations/), you can improve security in your Git repositories for both your code and deployed applications.
* Using [CI/CD integrations](../developer-tools/snyk-ci-cd-integrations/), you can fail the build in your integration and deployment pipeline to keep vulnerabilities out of your code.

### Monitor your code in production

Before integrating your code into production, use the [`snyk monitor`](../developer-tools/snyk-cli/commands/monitor.md) or [`snyk container monitor`](../developer-tools/snyk-cli/commands/container-monitor.md) CLI command to identify issues introduced into open-source and container Projects, monitoring these Projects for vulnerabilities before pushing them into production.

For more information, see [Monitor your projects at regular intervals](../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md).

### Manage and fix issues using Snyk

If you see hundreds or thousands of issues when first scanning your application, prioritization of issues becomes important. For more information, see [Prioritize your issues](../manage-risk/prioritize-issues-for-fixing/).

Snyk offers capabilities to address issues both reactively and proactively:

* Being proactive
  * Use [Snyk Advisor](https://snyk.io/advisor) to identify better packages to begin designing.
  * Use the CLI and IDE plugins to test while developing.
  * Add a package, ensure it is installed, and scan for security before writing your code.
* Fix advice\
  Snyk provides this advice across integrations, calculating the top-level package requiring an update in the package manifest or how to update the line of code to make it secure and displaying the advice on results screens.
* Automation
  * Snyk can create automatic fix pull requests when a new vulnerability is detected with a fix available.
  * You can enable dependency upgrade-related pull requests created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.

### Deployment and rollout recommendations

Startups, small teams, individuals, and open-source maintainers typically onboard their applications using Git, getting results in minutes and starting to address issues almost immediately. Small teams benefit from being agile and determining what works best for their workflow. For more details, see the [Team implementation guide](../implementation-and-setup/team-implementation-guide/).

With large organizations developing hundreds of applications, a slower approach is recommended to get developer buy-in and adoption and to ensure a positive rollout experience. For more details, see the [Enterprise implementation guide](../implementation-and-setup/enterprise-implementation-guide/).

## Schedule dynamic scans (Snyk API & Web)

Automate recurring dynamic scans of your web apps and APIs to track drift and regressions. See how to configure schedules and policies:

- [Schedule a scan](snyk-api-web/scanning/schedule-a-scan.md)

## Troubleshooting Snyk API & Web scanning

Resolve common issues when scanning dynamic targets:

- [Improve coverage (common issues)](snyk-api-web/troubleshooting/low-coverage.md)
- [Verify site ownership for scanning (DNS)](snyk-api-web/troubleshooting/domain-verification-dns.md)
- [Authenticate login forms for crawlers](snyk-api-web/troubleshooting/login-failed-target-authentication-login-form.md)
