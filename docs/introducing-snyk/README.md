# Introducing Snyk

Snyk tests for vulnerabilities in [your own code](https://snyk.io/product/snyk-code/), [open source dependencies](https://docs.snyk.io/snyk-open-source), [Container images](https://docs.snyk.io/snyk-container), and [Infrastructure as Code (IaC) configurations](https://snyk.io/product/infrastructure-as-code-security/), and offers context, prioritization, and fixes.

{% embed url="https://snyk.io/wp-content/uploads/Homepage-IDE-animation-Log4Shell-FINAL.mp4" %}

### Snyk products

* [Snyk Open Source](https://docs.snyk.io/snyk-open-source): Find and fix open source vulnerabilities. Snyk Open Source also includes [Snyk license compliance](https://docs.snyk.io/snyk-open-source) to help manage your open source license usage.
* [Snyk Code](https://snyk.io/product/snyk-code/): Find and fix vulnerabilities in your application code in real-time during the development process.
* [Snyk Container](https://docs.snyk.io/snyk-container): Find and fix vulnerabilities in container images and Kubernetes applications.
* [Snyk Infrastructure as Code (IaC)](https://docs.snyk.io/snyk-infrastructure-as-code): Find and fix insecure configurations in Terraform and Kubernetes code.

See [getting-started-snyk-products.md](../getting-started/getting-started-snyk-products.md "mention").

### Supported environments

Environments you can use to access your selected Snyk product include:

* **Snyk Web UI**: the core web-based UI ([app.snyk.io](https://app.snyk.io)).
* [**Snyk CLI**](https://docs.snyk.io/snyk-cli): the Command Line Interface.
* [**Snyk IDEs**](../ide-tools/): use IDE integrations to embed Snyk in your development environment.
* [**Snyk API**](https://support.snyk.io/hc/en-us/categories/360000665657-Snyk-API): allows you to programmatically integrate with Snyk.

### Supported languages

Snyk supports a range of languages, such as JavaScript, Java, .NET, and many others.

Languages / OSs supported depend on the relevant [Snyk product](./#snyk-products-and-platforms):

* Snyk Open Source: see [language-and-package-manager-support](../products/snyk-open-source/language-and-package-manager-support/ "mention")
* Snyk Code: see [snyk-code-language-and-framework-support.md](../products/snyk-code/snyk-code-language-and-framework-support.md "mention")
* Snyk Container: see [supported-operating-system-distributions.md](../products/snyk-container/snyk-container-security-basics/supported-operating-system-distributions.md "mention")
* Snyk IaC: we support configuration files for Terraform or Kubernetes environments

### Supported integrations

Taking a developer-first approach to security, Snyk integrates with leading IDE, repository, CI/CD, runtime, registry, and issue management tools.

[Snyk integrations](https://docs.snyk.io/integrations) for your software development process include:

* ****[**Source control**](../features/integrations/git-repository-scm-integrations/)**:** cloud and self-hosted SCMs such as Github.&#x20;
* ****[**CI / CD integrations**](../features/integrations/ci-cd-integrations/): such as Jenkins or TeamCity.
* **Artifact repositories:** such as Artifactory. See [Private registry gatekeeper plugins](https://docs.snyk.io/integrations/private-registry-gatekeeper-plugins) and [Private registry integrations](https://docs.snyk.io/integrations/private-registry-integrations).
* [**Serverless integrations**](https://docs.snyk.io/integrations/serverless-integrations): such as AWS Lambda.
* ****[**Platform as a Service (PaaS) integrations**](../features/integrations/platform-as-a-service-integrations/)**:** such as Heroku.
* [**Notification and ticketing system-integrations**](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations): such as Slack.
* [**Vulnerability Management Tools**](../features/integrations/vulnerability-management-tools/): such as RiskSense.

### Snyk CLI

You can use the Snyk CLI (Command Line Interface) to run scans on your local machine and integrate Snyk into your pipeline. You can use the Snyk CLI to scan your applications, containers, and Infrastructure as Code for security vulnerabilities. You can install the CLI via npm, Homebrew, Scoop, or manually.

For more information, see the [Snyk CLI docs](../snyk-cli/).

### Snyk API

Snyk’s [extensibility and API](https://snyk.io/blog/extensibility-and-the-snyk-api/) enable developers to tune Snyk’s security automation to their specific workflows, ensuring consistency in both developer experience and platform governance.

See how Snyk customers like [Twilio and Spotify](https://snyk.io/blog/snyk-watcher-keep-snyk-in-sync/) use the Snyk API in their workflows.

For more information, see the [Snyk API docs](../features/snyk-api-info/).
