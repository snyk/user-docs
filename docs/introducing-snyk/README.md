# Introducing Snyk

Snyk tests for vulnerabilities in [your own code](https://snyk.io/product/snyk-code/), [open source dependencies](https://docs.snyk.io/snyk-open-source), [Container images](https://docs.snyk.io/snyk-container), and [Infrastructure as Code (IaC) configurations](https://snyk.io/product/infrastructure-as-code-security/), and offers context, prioritization, and fixes.

![](<../.gitbook/assets/Screen Shot 2022-02-22 at 8.18.50 AM.png>)

### Snyk products&#x20;

* [Snyk Open Source](https://docs.snyk.io/snyk-open-source): Find and fix open source vulnerabilities. Snyk Open Source also includes [Snyk license compliance](https://docs.snyk.io/snyk-open-source) to help manage your open source license usage.
* [Snyk Code](https://snyk.io/product/snyk-code/): Find and fix vulnerabilities in your application code in real-time during the development process.
* [Snyk Container](https://docs.snyk.io/snyk-container): Find and fix vulnerabilities in container images and Kubernetes applications.
* [Snyk Infrastructure as Code (IaC)](https://docs.snyk.io/snyk-infrastructure-as-code): Find and fix insecure configurations in Terraform and Kubernetes code.

See [getting-started-snyk-products.md](../getting-started/getting-started-snyk-products.md "mention").

### Supported languages&#x20;

Snyk supports a range of languages, such as JavaScript, Java, .NET, and many others.

Languages / OSs supported depend on the relevant [Snyk product](./#snyk-products-and-platforms):

* Snyk Open Source: see [language-and-package-manager-support](../products/snyk-open-source/language-and-package-manager-support/ "mention")
* &#x20;Snyk Code: see [snyk-code-language-and-framework-support.md](../products/snyk-code/snyk-code-language-and-framework-support.md "mention")
* Snyk Container: see [supported-operating-system-distributions.md](../products/snyk-container/snyk-container-security-basics/supported-operating-system-distributions.md "mention")
* Snyk IaC: we support configuration files for Terraform or Kubernetes environments&#x20;

### Supported integrations

Taking a developer-first approach to security, Snyk integrates with leading IDE, repository, CI/CD, runtime, registry, and issue management tools.

[Snyk integrations](https://docs.snyk.io/integrations) for your software development process include:

* **Source control:** cloud and self-hosted SCMs such as Github. See [Git repository (SCM) integrations](../integrations/git-repository-scm-integrations/).
* **Container registries & Kubernetes:** container registries such as Docker Hub, and Kubernetes. See [Snyk Container](https://docs.snyk.io/snyk-container).
* **Continuous integration**: such as Jenkins or TeamCity. See [CI/CD Integrations](https://docs.snyk.io/integrations/ci-cd-integrations)
* **IDE plugins:** using IDE tools such as Eclipse. See [IDE tools](https://docs.snyk.io/integrations/ide-tools).
* **Artifact repositories:** such as Artifactory. See [Private registry gatekeeper plugins](https://docs.snyk.io/integrations/private-registry-gatekeeper-plugins) and [Private registry integrations](https://docs.snyk.io/integrations/private-registry-integrations).
* **Serverless**: such as AWS Lambda. See [Serverless integrations](https://docs.snyk.io/integrations/serverless-integrations).
* **Platform as a service:** such as Heroku. See [Platform as a service integrations](https://docs.snyk.io/integrations/platform-as-a-service-integrations).
* **Notifications**: such as Slack. See [Notification and ticketing system-integrations](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations).
* **Vulnerability tools**: such as RiskSense. See [Vulnerability Management Tools](../integrations/vulnerability-management-tools/).

### Snyk CLI

You can use the Snyk CLI (Command Line Interface) to scan and monitor on your local machine, and integrate it into your pipeline. You can use the Snyk CLI to scan your applications, containers, and infrastructure as code for security vulnerabilities. You can install the CLI via npm, Homebrew, Scoop, or manually.

See [snyk-cli](../snyk-cli/ "mention") for more details.

### Snyk API

Snyk’s extensibility and API enable developers to tune Snyk’s security automation to their specific workflows, ensuring both developer experience and consistent platform governance.&#x20;

Learn more in our [Snyk API documentation](https://support.snyk.io/hc/en-us/articles/360000914857-Does-Snyk-have-an-API-) and see how our customers like [Twilio and Spotify](https://snyk.io/blog/snyk-watcher-keep-snyk-in-sync/) use the Snyk API in their workflows.

