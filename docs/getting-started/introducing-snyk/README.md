# Introducing Snyk

## What is Snyk?

Snyk is a platform allowing you to scan, prioritize, and fix security vulnerabilities in your own code, open source dependencies, container images, and Infrastructure as Code (IaC) configurations.

### Snyk’s developer-first approach

Developers now assemble applications with a combination of proprietary and open source code, run that code in containers, and then deploy with infrastructure as code configurations with technologies like Kubernetes and Terraform.

A good security process secures each of these components where they are built and maintained. Snyk integrates into DevOps processes to work the with developers using the methods each prefers, while following and supporting industry best practices. Snyk integrates directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit.

![](<../../.gitbook/assets/image (162) (1) (1) (1) (1) (1).png>)

#### Using Snyk in your workflow:

* **Secure your code**: use [Snyk Open Source](../../products/snyk-open-source/) to fix vulnerabilities in your open source dependencies, and [Snyk Code](../../products/snyk-code/) to fix vulnerabilities in your source code.
* **Secure your environments**: use [Snyk Container](../../products/snyk-container/) to fix vulnerabilities in container images and Kubernetes workloads, and [Snyk Infrastructure as Code (IaC)](../../products/snyk-infrastructure-as-code/) to fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates.

### Ways to run Snyk

Snyk runs in the following environments:

* In the browser as [Snyk Web UI](../../snyk-web-ui/).
* In your local Command Line Interface (or in your pipeline) with the [Snyk CLI](https://docs.snyk.io/snyk-cli).
* In various [IDE plugins](../../ide-tools/) that embed Snyk in your local dev environment.
* In the [Snyk API](https://support.snyk.io/hc/en-us/categories/360000665657-Snyk-API) (Application Programming Interface).

See [Snyk products and environments](snyk-products-and-environments.md) for more details.

{% embed url="https://snyk.io/wp-content/uploads/Homepage-IDE-animation-Log4Shell-FINAL.mp4" %}
Running Snyk from the command line.
{% endembed %}

### Supported languages and systems

[Snyk products](snyk-products-and-environments.md) support multiple languages and formats:

* **Snyk Open Source**: see [Open Source - Supported languages and package managers](../../products/snyk-open-source/language-and-package-manager-support/).
* **Snyk Code**: see [Snyk Code - Supported languages and frameworks](../../products/snyk-code/snyk-code-language-and-framework-support.md).
* **Snyk Container**: see [Supported operating system distributions.](../../products/snyk-container/supported-operating-system-distributions.md)
* **Snyk Infrastructure as Code**: [Snyk IaC ](../../products/snyk-infrastructure-as-code/)supports configuration files for HashiCorp Terraform, AWS CloudFormation, Kubernetes, and Azure Resource Manager (ARM).&#x20;
* &#x20;**Snyk Cloud:** [Snyk Cloud](../../products/snyk-cloud/) supports scanning [Amazon Web Services resources](../../products/snyk-cloud/supported-aws-resources-for-snyk-cloud.md) and [Google Cloud resources](../../products/snyk-cloud/getting-started-with-snyk-cloud-google/).

### Supported integrations

[Snyk integrations](https://docs.snyk.io/integrations) for your software development process include:

* [**Source control**](../../integrations/git-repository-scm-integrations/)**:** cloud and self-hosted Git-based code repositories such as Github.
* [**CI/CD integrations**](../../integrations/ci-cd-integrations/): such as Jenkins or TeamCity.
* **Artifact repositories:** such as Artifactory. See [Private registry gatekeeper plugins](https://docs.snyk.io/integrations/private-registry-gatekeeper-plugins) and [Private registry integrations](https://docs.snyk.io/integrations/private-registry-integrations).
* [**Serverless integrations**](https://docs.snyk.io/integrations/serverless-integrations): such as AWS Lambda.
* [**Platform as a Service (PaaS) integrations**](../../integrations/platform-as-a-service-integrations/)**:** such as Heroku.
* [**Notification and ticketing system-integrations**](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations): such as Slack.
* [**Vulnerability Management Tools**](../../integrations/vulnerability-management-tools/): such as RiskSense

### **What does it cost?**

Snyk has several [pricing plans](https://snyk.io/plans/) available, from free to Enterprise.

See [Plans](plans.md) for more details.

### **Who uses Snyk?**

Google, Salesforce, Atlassian, Twilio, Revolut, and many more are using Snyk to secure their code and monitor for vulnerabilities.

### **How do I get started?**

See [Getting started](../) for full information.

### What happens to my data?

See [How Snyk handles your data](https://docs.snyk.io/more-info/how-snyk-handles-your-data) for details of Snyk data handling.
