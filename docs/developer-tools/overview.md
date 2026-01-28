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

Use the [Snyk CLI](snyk-cli/) for scripting and automation. It supports local testing, CI/CD pipelines, and advanced workflows like SBOM generation and custom configuration.

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
Snyk IDE plugins and extensions rely on the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli) and the Snyk Vulnerability Database to perform many functions. Visit the individual IDE documentation for more information.
{% endhint %}
