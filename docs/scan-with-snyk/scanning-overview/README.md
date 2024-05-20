# Scanning overview

{% hint style="info" %}
Scans may be limited on your account, depending on your[ Pricing Plan](../../implement-snyk/enterprise-implementation-guide/trial-limitations.md). See [What counts as a scan?](../../working-with-snyk/what-counts-as-a-test.md) for more information.
{% endhint %}

Snyk takes a developer-first approach to secure your development work by integrating directly into your IDEs, workflows, and automation pipelines to add security expertise to your toolkit. This approach allows you to:

* Use Snyk to focus on early enablement, not later enforcement.&#x20;
* Run scans while working on a Project, before you commit any code. This minimizes rework by finding issues that require changes early on.
* Add and test packages before writing the code that interfaces with each package.
* &#x20;After writing a major section of code, scan it to find issues before continuing work.

## Learn how to design secure code

The following resources are available for all users:

* [Snyk Advisor](https://snyk.io/advisor): Helps you pick healthy open-source packages or base images to develop code.
* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely, and provides training on how to use Snyk.

## Write and deploy your code

* Using the [Snyk CLI](../../snyk-cli/), you can scan locally on your machine. This is useful in scanning open-source and static code as well as containers and infrastructure as code configurations, including complex files that are templated with variables, such as Terraform plan files.
* Using [Snyk IDE Plugins](../../integrate-with-snyk/use-snyk-in-your-ide/), you can test your open-source packages, first-party code, and infrastructure as code (IaC) Kubernetes deployment files in your development environment as you create your Project.
* Using [Git integrations](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/), you can improve security in your Git repositories for both your code and deployed applications.
* Using [CI/CD integrations](../../integrate-with-snyk/snyk-ci-cd-integrations/), you can fail the build in your integration and deployment pipeline to keep vulnerabilities out of your code.

## Monitor your code in production

Before integrating your code into production, use the [`snyk monitor`](../../snyk-cli/commands/monitor.md) or [`snyk container monitor`](../../snyk-cli/commands/container-monitor.md) CLI command to identify issues introduced into open-source and container Projects, monitoring these Projects for vulnerabilities before pushing them into production.

See [Monitor your projects at regular intervals](../../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md) for more details.

## Manage and fix issues using Snyk

If you see hundreds or thousands of issues when first scanning your application, prioritization of issues becomes important. For more details, see [Prioritize your issues](../../manage-risk/prioritize-issues-for-fixing/).

Snyk offers capabilities to address issues both reactively and proactively:

* **Being proactive**
  * Use [Snyk Advisor](https://snyk.io/advisor) to identify better packages to begin designing.
  * Use the CLI and IDE plugins to test while developing.
  * Add a package, ensure it is installed, and scan for security before writing your code.
* **Fix advice**\
  Snyk provides this advice across integrations, calculating the top-level package requiring an update in the package manifest or how to update the line of code to make it secure and displaying the advice on results screens.
* **Automation**
  * Snyk can create automatic fix pull requests when a new vulnerability is detected with a fix available.
  * You can enable dependency upgrade-related pull requests created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.

## Deployment and rollout recommendations

#### Smaller businesses

Startups, small teams, individuals, and open-source maintainers typically onboard their applications using Git, getting results in minutes and starting to address issues almost immediately. Small teams benefit from being agile and determining what works best for their workflow. &#x20;

For more details, see the [Team implementation guide](../../implement-snyk/team-implementation-guide/).

#### Larger businesses

With large organizations developing hundreds of applications, a slower approach is recommended to get developer buy-in and adoption and to ensure a positive rollout experience.

For more details, see the [Enterprise implementation guide](../../implement-snyk/enterprise-implementation-guide/).
