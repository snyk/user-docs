# Introducing Snyk

## What is Snyk?

Snyk is a platform allowing you to scan, prioritize, and fix security vulnerabilities in your own code, open source dependencies, container images, and Infrastructure as Code (IaC) configurations.

## Snyk’s developer-first approach

Developers now assemble applications with a combination of proprietary and open source code, run that code in containers, and then deploy with infrastructure as code configurations with technologies like Kubernetes and Terraform.

A good security process secures each of these components where they are built and maintained. Snyk integrates into DevOps processes to work the with developers using the methods each prefers, while following and supporting industry best practices. Snyk integrates directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit.

<figure><img src="../.gitbook/assets/image (162) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Developer Security Platform"><figcaption><p>Developer Security Platform: Products and Developer experience</p></figcaption></figure>

## Snyk products: use Snyk in your workflow

* **Secure your code**: use [Snyk Open Source](../scan-application-code/snyk-open-source/) to fix vulnerabilities in your open source dependencies, and [Snyk Code](../scan-application-code/snyk-code/) to fix vulnerabilities in your source code.
* **Secure your containers**: use [Snyk Container](../scan-containers/) to fix vulnerabilities in container images and Kubernetes applications
* **Secure your deployment**: and [Snyk Infrastructure as Code (IaC)](../scan-cloud-deployment/snyk-infrastructure-as-code/) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates. Use [Snyk Cloud](../scan-cloud-deployment/snyk-cloud/) to fix misconfigurations in Amazon Web Services accounts, Microsoft Azure subscriptions, and Google Cloud projects.

## Snyk environments: choose how to run Snyk

You can run Snyk in the following ways:

* ****[**Web**](getting-started-with-the-snyk-web-ui.md): the Snyk Web UI ([app.snyk.io](https://app.snyk.io)) provides a browser-based experience, along with functions such as configuration settings, filtering and fixing discovered issues, and reports.
* ****[**CLI**](../snyk-cli/): the Snyk Command Line Interface enables you to run vulnerability scans on your local machine and integrate Snyk into your pipeline.
* [**IDEs**](../ide-tools/): the Snyk IDE integrations enable you to embed Snyk in your development environment.
* ****[**API**](../snyk-api-info/): the Snyk API enables you to programmatically integrate with Snyk, tuning Snyk’s security automation to your specific workflows.

This video shows using the CLI to scan for vulnerabilities.

{% embed url="https://snyk.io/wp-content/uploads/Homepage-IDE-animation-Log4Shell-FINAL.mp4" %}
Running Snyk from the command line.
{% endembed %}

## How can Snyk work in my environment?

Snyk products support multiple languages and formats:

* **Snyk Open Source**: see [Open Source - Supported languages and package managers](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/).
* **Snyk Code**: see [Snyk Code - Supported languages and frameworks](../scan-application-code/snyk-code/snyk-code-language-and-framework-support.md).
* **Snyk Container**: see [Supported operating system distributions.](../scan-containers/supported-operating-system-distributions.md)
* **Snyk Infrastructure as Code**: [Snyk IaC ](../scan-cloud-deployment/snyk-infrastructure-as-code/)supports configuration files for HashiCorp Terraform, AWS CloudFormation, Kubernetes, and Azure Resource Manager (ARM).
* **Snyk Cloud:** [Snyk Cloud](../scan-cloud-deployment/snyk-cloud/) supports scanning [Amazon Web Services resources](../scan-cloud-deployment/snyk-cloud/supported-resources-for-snyk-cloud/supported-aws-resources-for-snyk-cloud.md), [Azure resources](../scan-cloud-deployment/snyk-cloud/supported-resources-for-snyk-cloud/supported-azure-resources-for-snyk-cloud.md), and [Google Cloud resources](../scan-cloud-deployment/snyk-cloud/supported-resources-for-snyk-cloud/supported-google-resources-for-snyk-cloud.md).

## What can Snyk integrate with?

Snyk integrations for your software development process allow you to integrate Snyk into your development and security processes, including source control, CI/CD, and many others.

See [Snyk integrations](../integrations/) and [Snyk for IDEs](../ide-tools/) for details.

## **What does Snyk cost?**

Snyk has several [pricing plans](https://snyk.io/plans/) available, from free to Enterprise.

See [Snyk Pricing Plans](../more-info/plans.md) for more details.

## What happens to my data?

See [How Snyk handles your data](../more-info/how-snyk-handles-your-data.md) for details of Snyk data handling.
