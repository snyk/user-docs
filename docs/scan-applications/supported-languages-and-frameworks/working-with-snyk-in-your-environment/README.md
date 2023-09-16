# Working with Snyk in your environment

## The focus of this guide

This guide focuses on scanning your application code, which is a package manager and language-specific, [Snyk Container](../../../integrations/ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-container-specific-ci-cd-strategies.md) and [Snyk Infrastructure as Code (IaC)](../../../scan-infrastructure/snyk-infrastructure-as-code/) also support your container and infrastructure as code needs.&#x20;

Use this guide to understand the best way to apply Snyk in your workflow and to be aware of key considerations for your chosen technology stack.

## Developer-first approach

Snyk takes a developer-first approach to secure your application. This approach focuses less on overhead-heavy work, such as putting in hard QA gates, and more on building a secure application, providing visibility in a developer’s workflow, and providing actionable insights. This approach's benefits are engaging developers in security practices as part of their development without seeing security as a costly overhead.

{% hint style="success" %}
**Securing your code and open source dependencies with Snyk**

* Consider using Snyk to focus on earlier enablement, not later enforcement.&#x20;
* Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a Project to minimize rework.
* Add and test a package before writing the code that interfaces with it. Similarly, after writing major code sections, test it before moving on.
{% endhint %}

## Snyk overview

Some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in.&#x20;

{% hint style="info" %}
There is a monthly limit to the number of tests performed if a particular product is not purchased.&#x20;

:link: [Pricing plans](https://snyk.io/plans)&#x20;
{% endhint %}

### Snyk products

| Product                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Snyk Code**                                                                                                                                                                                    | Scan your code for security vulnerabilities using source code analysis.                                                                                                                                                                                                                                                                                                                         |
| <p><strong>Snyk Open Source</strong><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ</span> Some capabilities may be limited for some languages and package managers.</p> | <p></p><ul><li>Open Source vulnerability testing and monitoring (All plans).</li></ul><ul><li>Open Source dependency upgrade version bumping (All plans).</li></ul><ul><li>License Compliance (paid plans).</li></ul>                                                                                                                                                                           |
| **Snyk Infrastructure as Code**                                                                                                                                                                  | <p>Scan for configuration issues when you deploy your new applications using Kubernetes deployment files, Terraform, or Cloudformation templates. <br><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../scan-cloud-configurations/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/">Snyk CLI for Infrastructure as Code</a></p>        |
| **Snyk Container**                                                                                                                                                                               | Scan for issues with container images if you are building containers.                                                                                                                                                                                                                                                                                                                           |
| **Snyk Integrated IaC with cloud context**                                                                                                                                                       | <p></p><ul><li>Security from code to cloud and back.</li></ul><ul><li>Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.</li></ul><p><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../scan-infrastructure/snyk-iac+/">Snyk Integrated IaC with cloud context</a></p> |

### Product features to enhance your workflow

As you start planning and your code progresses through the pipeline, Snyk can provide different capabilities at each stage to help you find and fix security issues.

### Designing and planning

The following public resources are available for all users:

* [Snyk Advisor](https://snyk.io/advisor): Helps you pick healthy open source packages or base images to start developing with.
* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely.
* [Snyk Training](https://training.snyk.io/): Provides training on how to use Snyk.

### Coding

The following capabilities are available for all Snyk users.

* Snyk [IDE Plugins](../../../integrations/ide-tools/) for VS Code, IntelliJ, and others: Test your open source packages and first-party code as you develop. Test infrastructure as code (IaC) Kubernetes deployment files you create.
* [Snyk CLI](../../../snyk-cli/): A terminal program that allows you to test locally on your machine. Useful in testing containers and more complex IaC files that are templated with variables, such as Terraform plan files, as well as scanning open source and your code.

Snyk allows you to [run PR Checks](../../../scan-application-code/run-pr-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis and show results.

These functions can be used for:

* Code analysis with Snyk Code
* Open source and licensing with Snyk Open Source
  * Check for known vulnerabilities and create Fix Pull Requests (npm/Yarn 1,2,3)
  * Check for license compliance checks (Snyk Open Source)
  * Dependency upgrade - positioning updates to address technical debt (Snyk Open Source)

Additionally, with Git Integration, you can also monitor the following on a daily basis:

* Infrastructure as code (IaC) with Snyk Infrastructure as Code

#### **With CI/CD integrations**

Snyk can passively monitor and/or provide a QA gate by failing build checks during testing for policy violations.

Snyk provides the following capabilities:

* Dedicated plugins for Jenkins, Circle CI, and other CI/CD tools (see relevant marketplaces).
* [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/).
* The Snyk CLI can be used in most CI/CD tools (see [Snyk CI/CD integration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples)).
  * Fail the build based on criteria using options or the [snyk-filter tool](../../../snyk-api-info/other-tools/tool-snyk-filter.md).
  * [Containerized](https://hub.docker.com/r/snyk/snyk) versions of the CLI are available.
* Partner Platforms such as Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk.

### Production Monitoring

* Snyk can monitor container images and their open source/Linux-based packages being used in production via Kubernetes integration to notify customers of known vulnerabilities for applications in production.\
  :information\_source: **Snyk Enterprise plan only**\

* Where a production integration does not exist, use the [snyk monitor](../../../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.\
  :information\_source: **All Snyk plans**

### Gain recognition

You can place Snyk on your public repositories and have a badge indicating your package's security status.&#x20;

:link: [Instant security information with the Snyk security badge](https://snyk.io/blog/instant-security-information-with-the-snyk-security-badge/)

## Deployment and rollout recommendations

Startups, small teams, individuals, open source maintainers typically onboard their applications via Git, getting results in minutes and starting to address issues almost immediately. Small teams have the benefit of being agile and determining what works best for their workflow.

With large organizations using hundreds of applications, a slower approach is recommended to get developer buy-in/adoption and to ensure a positive rollout experience.

1. Typically, large organizations start with daily monitoring of applications via Git integration, only initially turning on PR checks for a few key applications.
2. As developers become familiar with Snyk's capabilities, you can widen the scope of applications with PR checks for gating or blocking builds if checks fail.
3. Some customers use CI/CD to passively monitor and then turn on gating by using the `snyk [product] test` commands.
4. If you onboard a large number of legacy applications, you can use [Priority Score](../../../manage-issues/priorities-for-fixing-issues/priority-score.md) (typically 700 as a starting place) or criteria like “Known exploit” or “Fix available” to define a starting point to engage developers to start fixing vulnerabilities on key applications.

## Reporting

For Snyk Open Source, the `snyk [product] monitor` command is used to push results from the CLI back to Snyk for reporting in the Snyk UI. Don’t forget to use `--org=` to indicate what Organization to place the monitored results in or retrieve test settings from during the test.

{% hint style="info" %}
Snyk Enterprise plan customers can access the [Snyk API](../../../snyk-api/) for reporting and extracting data.
{% endhint %}

## Proxies

:link: [CLI proxy guide](../../../snyk-cli/configure-the-snyk-cli/proxy-configuration-for-snyk-cli.md)&#x20;

## How to fix issues using Snyk

If you see hundreds or thousands of issues when first scanning your application, prioritization of issues becomes the most important.

### What to fix - prioritization factors

Some tools only use the single factor of severity to prioritize issues, but this can still result in thousands of results, with no clear starting point to fix these issues.

Snyk provides more factors to help you prioritize issues, such as having a known exploit, is it fixable, or social trending. This can be done both at the Project level when looking at a specific Project, or Enterprise customers can prioritize across all Projects.

To see prioritization in action, see the [Prioritize issues in the Snyk Web UI training course](https://training.snyk.io/learn/video/prioritize-ui).

### How to fix - addressing issues

Snyk offers capabilities in this ecosystem to help address issues, both reactively and proactively:

1. **Being proactive**\
   Use [Snyk Advisor](https://snyk.io/advisor) to identify better packages to begin designing.\
   Use the CLI and IDE plugins to test while developing.\
   Add a package, make sure it’s installed, and scan for security before writing your code.
2. **Remediation advice**\
   Snyk supplies this across integrations on the results screens that calculate the top-level package requiring an update in pom.xml or build.gradle or how to update the line of code to make it secure.
3. **Automation**
   * You can enable automatic fix pull requests created when a new vulnerability is detected with a fix available.
   * When first importing a Project into Snyk, it may already have vulnerabilities. Instead of overwhelming you with fixes, Snyk will present the highest priority known vulnerability for you to fix first. After fixing, the next highest priority known vulnerability is presented. In this way, you can catch up and fix existing issues. You can speed this process up by fixing all vulnerabilities in a single dependency at a time by enabling **Clear backlog faster** [Snyk Preview](../../../snyk-admin/manage-settings/snyk-preview.md).
   * You can enable dependency upgrade-related pull requests created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.

{% hint style="info" %}
We suggest enabling automated PRs on a key project to start before enabling globally to familiarize development on how to interact with Snyk and security issues.
{% endhint %}
