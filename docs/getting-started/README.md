# Getting started

## What is Snyk?

Snyk is a platform that allows you to scan, prioritize, and fix security vulnerabilities in your code, open-source dependencies, container images, and infrastructure as code configurations.

## The Snyk developer-first approach

Snyk provides visibility in a developer’s workflow and actionable insights. The benefit is engaging developers in security practices as part of their development work. Thus, the focus is on building a secure application rather than overhead-intensive work, such as putting in hard QA gates.&#x20;

Developers now assemble applications with a combination of proprietary and open-source code, run that code in containers, and then deploy with infrastructure as code configurations using technologies like Kubernetes and Terraform.

A robust security process secures each of these components where they are built and maintained. Snyk integrates into DevOps processes to work with developers using the methods each prefers, while following and supporting industry best practices. Snyk integrates directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit.

<figure><img src="../.gitbook/assets/image (162) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Developer Security Platform: Products and Developer experience"><figcaption><p>Developer Security Platform: Products and Developer experience</p></figcaption></figure>

## Use Snyk in your workflow

* **Secure your code**: use [Snyk Open Source](../scan-with-snyk/snyk-open-source/) to fix vulnerabilities in your open source dependencies and [Snyk Code](../scan-with-snyk/snyk-code/) to fix vulnerabilities in your source code.
* **Secure your containers**: use [Snyk Container](../scan-with-snyk/snyk-container/) to fix vulnerabilities in container images and Kubernetes applications.
* **Secure your infrastructure**: use [Snyk Infrastructure as Code (IaC)](../scan-with-snyk/snyk-iac/scan-your-iac-source-code/) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates. Use [IaC+](../scan-with-snyk/snyk-iac/iac+-code-to-cloud-capabilities/) to fix misconfigurations in Amazon Web Services accounts, Microsoft Azure subscriptions, and Google Cloud Projects.

## Choose how to run Snyk

You can run Snyk in the following ways:

* [**Web**](snyk-web-ui.md): the Snyk Web UI ([app.snyk.io](https://app.snyk.io)) provides a browser-based experience with functions such as configuration settings, filtering and fixing discovered issues, and reports.
* [**CLI**](../snyk-cli/): the Snyk Command Line Interface enables you to run vulnerability scans on your local machine and integrate Snyk into your pipeline.
* [**IDEs**](../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/): the Snyk IDE integrations enable you to embed Snyk in your development environment.
* [**API**](../snyk-api/): the Snyk API enables you to integrate with Snyk programmatically, tuning Snyk security automation to your specific workflows.

This video shows using the CLI to scan for vulnerabilities.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Running Snyk from the command line.
{% endembed %}

## How can Snyk work in my environment?

The Snyk tech stacks that are supported depend on the Snyk product you use:

* **Snyk Open Source** and **Snyk Code**: see [Supported languages and frameworks](supported-languages-and-frameworks/).
* **Snyk Container**: see [Supported operating system distributions](../scan-with-snyk/snyk-container/how-snyk-container-works/supported-operating-system-distributions.md).
* **Snyk Infrastructure as Code**: see [Supported IaC and cloud providers](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/).

## What can Snyk integrate with?

Snyk integrations for your software development process allow you to integrate Snyk into your development and security processes, including source control, IDE, CI/CD, and many others.

See [Integrate with Snyk](../integrate-with-snyk/) for details.

## **What does Snyk cost?**

Snyk has several pricing plans available, from free to Enterprise. See [Snyk Pricing Plans](../implement-snyk/enterprise-implementation-guide/trial-limitations.md).

## What happens to my data?

See [How Snyk handles your data](../working-with-snyk/how-snyk-handles-your-data.md) for details of Snyk data handling.&#x20;

