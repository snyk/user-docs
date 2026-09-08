---
nav_context: agnostic
description: >-
  Run Snyk locally, in repositories, and in pipelines, and choose the developer
  tool integration that matches your workflow
---

# Overview

You can run Snyk locally, in repositories, and in pipelines to scan your code. Select the integration that matches your workflow and permissions.

## SCMs

Snyk integrates with SCMs like GitHub, GitLab, Bitbucket, and Azure Repos. Use these integrations to:

* import repositories
* run pull request checks
* open pull requests for fixes and upgrades

Visit the [SCMs ](scm-integrations/)page to learn more about workspaces, access scopes, and configuration at the Group and Organization levels.

## Snyk CI/CDs

[Snyk CI/CD](snyk-ci-cd-integrations/) integrations orchestrate development and operations in your system. Use these integrations to enforce policy gates, generate reports, and standardize scanning across pipelines.

Snyk provides platform-specific setup and parameters for:

* GitHub Actions
* Jenkins
* Azure Pipelines
* Bitbucket Pipelines

## Snyk CLI

Use the [Snyk CLI](snyk-cli/overview.md) for scripting and automation. It supports local testing, CI/CD pipelines, and advanced workflows like SBOM generation and custom configuration.

This section covers:

* installation
* authentication
* command reference
* configuration
* debugging
* analytics

## Snyk IDE plugins and extensions

[Snyk IDE plugins and extensions](snyk-ide-plugins-and-extensions/) provide scanning and remediation guidance in your Projects.

This section provides setup guidance, including authentication, proxy, workspace trust settings, and troubleshooting steps for each IDE.

{% hint style="warning" %}
Snyk IDE plugins and extensions rely on the [Snyk CLI](snyk-cli/overview.md) and the Snyk Vulnerability Database to perform many functions. Visit the individual IDE documentation for more information.
{% endhint %}

## Other integrations

Snyk also connects with:

* Package repositories and gatekeepers — [Artifactory Gatekeeper Plugin](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/manage-vulnerabilities/artifactory-gatekeeper-plugin), [package repository integrations](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-open-source/package-repository-integrations)
* [Container registries](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-container/container-registry-integrations)
* [Event forwarding](integrations/event-forwarding/)
* [Jira and Slack integrations](integrations/jira-and-slack-integrations/)
* [Partner integrations](integrations/partner-integrations.md)
* [Reporting and BI integrations](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/prevent/analytics/reports-tab/reporting-and-bi-integrations-snowflake-data-share)

For Snyk Essentials-specific integration availability, sync times, and connection statuses, see [Snyk Essentials integrations reference](integrations/snyk-essentials-integrations-reference.md).

## AI-assisted development

To secure AI coding assistants and agentic workflows, see the [Agent security](https://app.gitbook.com/s/N5N885PkllOWeBmgm3Bp/) section — it covers Snyk Studio and the Snyk MCP Server.
