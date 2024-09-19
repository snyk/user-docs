# Getting started

## What is Snyk?

Snyk begins with developer-focused security, providing security tools designed for developers to create secure code from the start. Snyk integrates into all stages of development - IDEs, source code managers, CI/CD pipelines, and repositories, to detect high-risk code, open source packages, containers, and cloud configurations, and give developers the precise information they need to fix each issue - or even to automate the fix, where desired.

The Snyk platform uses a risk-based approach, focusing security efforts on issues that matter, and eliminating the noise of vulnerabilities that have no meaningful impact. To manage and govern the security program, Snyk gives security teams immediate visibility into coverage and business context across all application assets, smart policies to automate and scale in large environments, and analytics and reporting to measure the performance of your security program.

## The Snyk developer-first approach

Developers now assemble applications with a combination of proprietary and open-source code, run them in containers, and deploy them using infrastructure as code configurations with technologies like Kubernetes and Terraform. Snyk provides visibility in a developer's workflow and actionable insights, engaging developers in security practices as part of their development work. This approach focuses on building secure applications rather than relying on overhead-intensive processes like stringent

Snyk builds trust between developers and AppSec teams by embracing DevOps, Cloud, and AI without limits. The path to trusted software begins with Developer Security, scales with DevSec Governance, and is navigated through Application Risk Management. Snyk integrates into DevOps processes, IDEs, workflows, and automation pipelines to add security expertise, aligning with industry best practices and enhancing each development method's security.

Implementing the Snyk platform allows you to:

* Secure code-to-cloud developer experience to find and fix hand-written and AI-generated code
* Unite developers and security through comprehensive and scalable governance capabilities&#x20;
* Uncover and manage actual business risk through holistic application risk management

<figure><img src="../.gitbook/assets/image (565).png" alt="Snyk Developer Security Platform: Products and Developer experience"><figcaption><p>Snyk Developer Security Platform: Products and Developer experience</p></figcaption></figure>

## Use Snyk in your workflow

* **Discover assets** from your application to ensure Snyk is securing your business-critical assets and identifying gaps:
  * Use the [Manage assets](../manage-assets/) functionality to discover and group assets from an application.
* **Scan your code** to fix vulnerabilities:&#x20;
  * Use [Snyk Open Source](../scan-with-snyk/snyk-open-source/) to fix vulnerabilities in your open source dependencies. &#x20;
  * Use [Snyk Code](../scan-with-snyk/snyk-code/) to fix vulnerabilities in your source code.&#x20;
  * Use [Snyk Container](../scan-with-snyk/snyk-container/) to fix vulnerabilities in container images and Kubernetes applications.
  * Use [Snyk Infrastructure as Code (IaC)](../scan-with-snyk/snyk-iac/scan-your-iac-source-code/) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates. Use [IaC+](../scan-with-snyk/snyk-iac/iac+-code-to-cloud-capabilities/) to fix misconfigurations in Amazon Web Services accounts, Microsoft Azure subscriptions, and Google Cloud Projects.
* **Prioritize issues** based on actual risk by using comprehensive visibility into your application and vulnerabilities:&#x20;
  * Use Snyk AppRisk to prioritize and manage risk more holistically with visibility into the runtime.&#x20;

## Choose how to run Snyk

You can run Snyk in the following ways:

* [**Web**](snyk-web-ui.md): the Snyk Web UI ([app.snyk.io](https://app.snyk.io)) provides a browser-based experience with functions such as configuration settings, filtering and fixing discovered issues, and reports.
* [**CLI**](../snyk-cli/): the Snyk Command Line Interface enables you to run vulnerability scans on your local machine and integrate Snyk into your pipeline.
* [**IDEs**](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/): the Snyk IDE integrations enable you to embed Snyk in your development environment.
* [**API**](../snyk-api/): the Snyk API enables you to integrate with Snyk programmatically, tuning Snyk security automation to your specific workflows.

This video shows using the CLI to scan for vulnerabilities.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Running Snyk from the command line.
{% endembed %}

## How can Snyk work in my environment?

The Snyk tech stacks that are supported depend on the Snyk product you use:

* **Snyk Open Source** and **Snyk Code**: see [Supported languages and frameworks](../supported-languages-package-managers-and-frameworks/).
* **Snyk Container**: see [Supported operating system distributions](../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md).
* **Snyk Infrastructure as Code**: see [Supported IaC and cloud providers](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/).
* **Snyk AppRisk**: see [Using Snyk AppRisk](../scan-with-snyk/snyk-apprisk/using-snyk-apprisk.md).

## What can Snyk integrate with?

Snyk integrations for your software development process allow you to integrate Snyk into your development and security processes, including source control, IDE, CI/CD, and many others.

See [Integrate with Snyk](../integrate-with-snyk/) at both the Group and Organization levels for details.

## **What does Snyk cost?**

Snyk has several pricing plans available, from free to Enterprise. See [Snyk Pricing Plans](https://snyk.io/plans/).

Snyk offers a trial of the platform, but this has imposed feature limitations. See [Trial limitations](https://docs.snyk.io/implement-snyk/enterprise-implementation-guide/trial-limitations).

## What happens to my data?

See [How Snyk handles your data](../working-with-snyk/how-snyk-handles-your-data.md) for details of Snyk data handling.

