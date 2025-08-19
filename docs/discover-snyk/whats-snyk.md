# What's Snyk?

Snyk is a platform that allows you to scan, prioritize, and fix security vulnerabilities in your code, open-source dependencies, container images, and infrastructure as code configurations. The Snyk platform uses a risk-based approach, focusing security efforts on issues that matter, and eliminating the noise of vulnerabilities that have no meaningful impact.

To manage and govern the security program, Snyk gives security teams immediate visibility into coverage and business context across all application assets, smart policies to automate and scale in large environments, and analytics and reporting to measure the performance of your security program.

* Snyk Open Source and Snyk Code: see [Supported languages and frameworks](../supported-languages/supported-languages-package-managers-and-frameworks.md).
* Snyk Container: see [Supported operating system distributions](../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md).
* Snyk Infrastructure as Code: see [Supported IaC and cloud providers](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/).
* Snyk Essentials: see [Snyk Essentials](../scan-with-snyk/snyk-essentials.md)&#x20;
* Snyk AppRisk: see [Snyk AppRisk](../scan-with-snyk/snyk-apprisk.md).
* Snyk API & Web: see [Snyk API & Web](https://snyk.io/product/dast-api-web/).

## The Snyk developer-first approach

Snyk provides visibility in a developer’s workflow and actionable insights. The benefit is engaging developers in security practices as part of their development work. Thus, the focus is on building a secure application rather than overhead-intensive work, such as putting in hard QA gates.

Developers now assemble applications with a combination of proprietary and open-source code, run that code in containers, and then deploy with infrastructure as code configurations using technologies like Kubernetes and Terraform.

A robust security process secures each component where they are built and maintained. Snyk integrates into DevOps processes to work with developers using the methods each prefers, while following and supporting industry best practices. Snyk integrates directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit.

<figure><img src="../.gitbook/assets/image (252).png" alt="Snyk Developer Security Platform: Products and Developer experience"><figcaption><p>Snyk Developer Security Platform: Products and Developer experience</p></figcaption></figure>

## Use Snyk in your workflow

* Secure your code: use [Snyk Open Source](../scan-with-snyk/snyk-open-source/) to fix vulnerabilities in your open source dependencies and [Snyk Code](../scan-with-snyk/snyk-code/) to fix vulnerabilities in your source code.
* Secure your containers: use [Snyk Container](../scan-with-snyk/snyk-container/) to fix vulnerabilities in container images and Kubernetes applications.
* Secure your infrastructure: use [Snyk Infrastructure as Code](../scan-with-snyk/snyk-iac/) (IaC) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates.

## Choose how to run Snyk

You can run Snyk in the following ways:

* Web: the Snyk Web UI ([app.snyk.io](https://app.snyk.io)) provides a browser-based experience with functions such as configuration settings, filtering and fixing discovered issues, and reports.
* [CLI](../developer-tools/snyk-cli/): the Snyk Command Line Interface enables you to run vulnerability scans on your local machine and integrate Snyk into your pipeline.
* [IDEs](../developer-tools/snyk-ide-plugins-and-extensions/): the Snyk IDE integrations enable you to embed Snyk in your development environment.
* [API](../snyk-api/snyk-api.md): the Snyk API enables you to integrate with Snyk programmatically, tuning Snyk security automation to your specific workflows.

## What can Snyk integrate with?

Snyk integrations for your software development process allow you to integrate Snyk into your development and security processes, including source control, IDE, CI/CD, and many others.

For details, see [Integrate with Snyk](../integrations/integrate-with-snyk.md).

## **What does Snyk cost?**

Snyk has several pricing plans available, from free to Enterprise. See [Snyk Pricing Plans](https://snyk.io/plans/).

Snyk offers a trial of the platform, but this has imposed feature limitations. See [Trial limitations](../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md).

## What happens to my data?

For details on Snyk handling, see [How Snyk handles your data](../snyk-data-and-governance/how-snyk-handles-your-data.md).
