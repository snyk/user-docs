---
nav_context: agnostic
last_reviewed: 2026-09-07
last_reviewed_source: focused-edit
---

# What is Snyk?

Snyk is a platform you use to scan, prioritize, and fix security vulnerabilities in your code, open-source dependencies, container images, infrastructure as code (IaC) configurations, and live web applications and APIs. The Snyk platform uses a risk-based approach, focusing security efforts on issues that matter, and eliminating the noise of vulnerabilities that have no meaningful impact.

To manage and govern the security program, Snyk gives security teams immediate visibility into coverage and business context across all application assets, smart policies to automate and scale in large environments, and analytics and reporting to measure the performance of your security program.

Each Snyk product documents what it supports:

* Snyk Open Source and Snyk Code: visit [Supported languages, package managers, and frameworks](supported-languages/supported-languages-package-managers-and-frameworks.md).
* Snyk Secrets: visit [Snyk Secrets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-secrets).
* Snyk Container: visit [Supported operating system distributions](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container).
* Snyk IaC: visit [Supported IaC and cloud providers](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources).
* Snyk Essentials: visit [Snyk Essentials](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-essentials).
* Snyk API & Web: visit [Snyk API & Web](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-api-web).

## The Snyk developer-first approach

Snyk provides visibility in a developer's workflow and actionable insights. The benefit is engaging developers in security practices as part of their development work. Thus, the focus is on building a secure application rather than overhead-intensive work, such as putting in hard quality assurance gates.

Developers now assemble applications with a combination of proprietary and open-source code, run that code in containers, and then deploy with infrastructure as code configurations using technologies like Kubernetes and Terraform.

A robust security process secures each component where they are built and maintained. Snyk integrates into DevOps processes to work with developers using the methods each prefers, while following and supporting industry best practices. Snyk integrates directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit.

## Use Snyk in your workflow

* Secure your code: use [Snyk Open Source](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source) to fix vulnerabilities in your open-source dependencies and [Snyk Code](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code) to fix vulnerabilities in your source code.
* Secure your secrets: use [Snyk Secrets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-secrets) to detect hard-coded secrets, credentials, and API keys across the IDE, CLI, SCM, and pull request checks.
* Secure your containers: use [Snyk Container](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-container) to fix vulnerabilities in container images and Kubernetes applications.
* Secure your infrastructure: use [Snyk IaC](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-iac) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates.
* Secure your APIs and web applications: use [Snyk API & Web](https://snyk.io/product/dast-api-web/) to discover and test the security of all your APIs and web apps, including AI-generated ones.

## Choose how to run Snyk

You can run Snyk in the following ways:

* Web: the Snyk Web UI provides a browser-based experience with functions such as configuration settings, filtering and fixing discovered issues, and reports. Snyk hosts the Web UI in several regions, and your account exists in only one of them. Log in at the URL for your region: [app.snyk.io](https://app.snyk.io), [app.us.snyk.io](https://app.us.snyk.io), [app.eu.snyk.io](https://app.eu.snyk.io), or [app.au.snyk.io](https://app.au.snyk.io). For the full list, including Snyk for Government, visit [Regional hosting and data residency](https://docs.snyk.io/snyk-data-and-governance/regional-hosting-and-data-residency).
* [CLI](https://docs.snyk.io/developer-tools/snyk-cli/overview): the Snyk Command Line Interface lets you run vulnerability scans on your local machine and integrate Snyk into your pipeline.
* [IDEs](https://docs.snyk.io/developer-tools/integrations/snyk-ide-plugins-and-extensions): the Snyk IDE integrations let you embed Snyk in your development environment.
* [API](https://docs.snyk.io/developer-tools/snyk-api/snyk-api): the Snyk API lets you integrate with Snyk programmatically, tuning Snyk security automation to your specific workflows.

## What can Snyk integrate with?

Snyk integrations connect Snyk to your development and security processes, including source control, IDEs, continuous integration/continuous delivery (CI/CD), and many others.

For details, visit [Integrate with Snyk](https://docs.snyk.io/developer-tools/integrations/integrate-with-snyk).

## What does Snyk cost?

Snyk has several pricing plans available, from Free to Enterprise. Visit [Snyk Pricing Plans](https://snyk.io/plans/).

Snyk offers a trial of the platform with feature limitations. Visit [Trial limitations](implementation-guides/enterprise-implementation-guide/trial-limitations.md).

## What happens to my data?

For details, visit [How Snyk handles your data](https://docs.snyk.io/snyk-data-and-governance/how-snyk-handles-your-data).

## Next steps

Set up Snyk and run your first scan:

* [Getting started](getting-started-guides/getting-started.md): create or log in to an account, set up an integration, and scan your first Project.
* [Start scanning](getting-started-guides/start-scanning.md): choose where to scan from, and run a scan.
* [Navigate the Snyk Web UI](navigate-the-snyk-web-ui.md): find where features live in the interface.
* [Implement Snyk](implementation-guides/implement-snyk.md): roll Snyk out across a team or an enterprise.
