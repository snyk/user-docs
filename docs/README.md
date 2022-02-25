# Snyk User Documentation

{% hint style="success" %}
[Sign up to use Snyk for free!](https://snyk.io/login?cta=sign-up\&loc=nav\&page=support\_docs\_page)
{% endhint %}

### **What is Snyk?**

Snyk integrates directly into development tools and automation pipelines, making it easy to find, prioritize, and fix security vulnerabilities in code, dependencies, containers, and infrastructure as code. Backed by industry-leading vulnerability intelligence, and designed by developers for developers, Snyk fits into your development workflow to put security expertise in your toolkit.

![](<.gitbook/assets/Screen Shot 2022-02-22 at 8.18.28 AM.png>)

### **What products does Snyk offer?**

* [Snyk Open Source](https://snyk.io/product/open-source-security-management/): Find and automatically fix open source vulnerabilities
* [Snyk Code](https://snyk.io/product/snyk-code/): Find and fix vulnerabilities in your application code in real time
* [Snyk Container](https://snyk.io/product/container-vulnerability-management/): Find and fix vulnerabilities in container images and Kubernetes workloads
* [Snyk Infrastructure as Code (IaC)](https://snyk.io/product/infrastructure-as-code-security/): Find and fix misconfigurations in Terraform, CloudFormation, Kubernetes, and Azure templates.

{% tabs %}
{% tab title="Snyk CLI" %}
You can use the CLI for scanning and monitoring on your local machine and integrate it into your workflow pipeline. You can use the Snyk CLI to scan applications, containers, and infrastructure as code for security vulnerabilities. Install the CLI via npm, Homebrew, Scoop, or manually.

{% content-ref url="features/snyk-cli/" %}
[snyk-cli](features/snyk-cli/)
{% endcontent-ref %}

{% embed url="https://snyk.io/blog/snyk-cli-cheat-sheet" %}
{% endtab %}

{% tab title="Snyk API" %}
Snyk’s extensibility and API enables developers to tune Snyk’s security automation to their specific workflows--ensuring both a great developer experience and consistent platform governance. Learn more in our [Snyk API documentation](https://snyk.docs.apiary.io/#introduction/api-v3) and see how our customers like [Twilio and Spotify](https://snyk.io/blog/snyk-watcher-keep-snyk-in-sync/) use the Snyk API in their workflows.

#### V1

The Snyk v1 API is available to customers on [paid plans](https://snyk.io/plans) and allows you to programatically integrate with Snyk.

{% embed url="https://snyk.docs.apiary.io" %}

#### V3

The Snyk v3 API is a new, standards-driven API (OpenAPI, JSON API) available to all customers.

{% embed url="https://apidocs.snyk.io" %}
{% endtab %}
{% endtabs %}

### **What does it cost?**

Snyk has several [pricing plans](https://snyk.io/plans/) available:

* **Free**: For individual developers and small teams looking to secure while they build. Limited tests.
* **Team**: For dev teams looking to build security into their development process with shared visibility into projects. Unlimited tests.
* **Business**: Empower developers across an organization and provide reporting and advanced controls to manage teams and control to shift security left. Unlimited tests.
* **Enterprise**: Standardize dev-first security across the enterprise, with centralized policy governance. Unlimited tests.

### **Who uses Snyk?**

Google, Salesforce, Datadog, Atlassian, Twilio, Revolut and many more are using Snyk to secure their code and monitor for vulnerabilities.

### **How do I get started?**

1. [Sign up](https://snyk.io/login?cta=sign-up\&loc=nav\&page=support\_docs\_page) and connect to your project.
2. Run tests against your project.
3. Review your results to identify vulnerabilities.
4. Correct these vulnerabilities via Pull Requests.
5. Remain secure with monitoring.

See [getting-started](getting-started/ "mention") for more details.

{% hint style="info" %}
For details of Snyk data handling, see [How Snyk handles your data](more-info/how-snyk-handles-your-data.md).
{% endhint %}
