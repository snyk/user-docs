# Best practices for JavaScript and Node.js



## The focus of this guide

This guide focuses on scanning your application code, which is a package manager and language-specific, [Snyk Container](../../../integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-container-specific-ci-cd-strategies.md) and [Snyk Infrastructure as Code (IaC)](../../../scan-cloud-configurations/snyk-infrastructure-as-code/) also support your container and infrastructure as code needs.&#x20;

Use this guide to understand the best way to apply Snyk in your workflow and to be aware of key considerations for your chosen technology stack.

## Developer-first approach

Snyk takes a developer-first approach to secure your application. This approach focuses less on overhead-heavy work, such as putting in hard QA gates, and more on building a secure application, providing visibility in a developer’s workflow and providing actionable insights. This approach's benefits are engaging developers in security practices as part of their development without seeing security as a costly overhead.

{% hint style="success" %}
**Securing your code and open source dependencies with Snyk**

* Consider using Snyk to focus on earlier enablement, not later enforcement.&#x20;
* Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a Project to minimize rework.
* Add and test a package before writing the code that interfaces with it. Similarly, after writing major sections of code, test it before moving on.
{% endhint %}

## Snyk overview

Some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in.&#x20;

{% hint style="info" %}
There is a monthly limit to the number of tests performed if a particular product is not purchased.&#x20;

:link: [Pricing plans](https://snyk.io/plans)&#x20;
{% endhint %}

### Snyk products

| Product                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Snyk Code**                                                                                                                                                                                    | Scan your code for security vulnerabilities using source code analysis.                                                                                                                                                                                                                                                                                                                                                       |
| <p><strong>Snyk Open Source</strong><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ</span> Some capabilities may be limited for some languages and package managers.</p> | <p></p><ul><li>Open Source vulnerability testing and monitoring (All plans).</li></ul><ul><li>Open Source dependency upgrade version bumping (All plans).</li></ul><ul><li>License Compliance (paid plans).</li></ul>                                                                                                                                                                                                         |
| **Snyk Infrastructure as Code**                                                                                                                                                                  | <p>Scan for configuration issues when you deploy your new applications using Kubernetes deployment files, Terraform, or Cloudformation templates. <br><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../scan-cloud-configurations/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/">Snyk CLI for Infrastructure as Code</a></p>                                      |
| **Snyk Container**                                                                                                                                                                               | Scan for issues with container images if you are building containers.                                                                                                                                                                                                                                                                                                                                                         |
| **Snyk Integrated IaC with cloud context**                                                                                                                                                       | <p></p><ul><li>Security from code to cloud and back.</li></ul><ul><li>Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.</li></ul><p><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../scan-cloud-configurations/integrated-iac-with-cloud-context/">Snyk Integrated IaC with cloud context</a></p> |

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

### Validating, monitoring, alerting, and gating

#### **With Git integrations**

<figure><img src="https://lh6.googleusercontent.com/EYPCKsyukOq5A9wNpYka8tUBa5FbzGQXrbmG2klrIigOxTNSInsA_Znj6P0jpGnBv7yRHAaiTsF_GX9Y9Zr1xdE35eZljg_1crKgqHBkhoZrEbvpTsdZstjXdVZ1hVF4jNyTgfLWbALbvqtDFbuI_ys" alt="PR Checks for Git integrations"><figcaption><p>PR Checks for Git integrations</p></figcaption></figure>

Snyk allows you to [run PR Checks](../../run-pr-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis and show results.

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
* Github Actions.
* The Snyk CLI can be used in most CI/CD tools (see [Snyk CI/CD integration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples)).
  * Fail the build based on criteria using options or the [snyk-filter tool](../../../snyk-api-info/other-tools/tool-snyk-filter.md).
  * [Containerized](https://hub.docker.com/r/snyk/snyk) versions of the CLI are available.
* Partner Platforms such as Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk.

### Production Monitoring

* Snyk can monitor container images and their open source/Linux-based packages being used in production via Kubernetes integration to notify customers of known vulnerabilities for applications in production.\
  :information\_source: **Snyk Enterprise plan only**\

* Where a production integration does not exist, use the [snyk monitor](../../../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.\
  :information\_source: **All Snyk plans**

### Package Registry Integrations (Artifactory/Nexus)

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

Snyk Open Source integrates with Artifactory and Nexus as local gatekeepers and interacts with the registry for security testing. Snyk uses this integration for fixes and also locks the Lockfile.

If your Projects reference private dependencies in Artifactory or Nexus but you are not a Snyk Enterprise user, you can use the Snyk CLI in a properly configured local environment (such as your build pipeline) so these dependencies can be resolved and included in the test.

For more information, see the following:

* Package registry integrations: [Artifactory Registry setup](../../../integrations/package-repository-integrations/artifactory-repository-setup/) and [Nexus Repository Manager setup](../../../integrations/package-repository-integrations/nexus-repo-manager-setup/).
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../../../integrations/gatekeeper-plugins/artifactory-gatekeeper-plugin-overview.md) and [Nexus Repository Manager Gatekeeper plugin](../../../integrations/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

### Gain recognition

You can place Snyk on your public repositories and have a badge indicating your package's security status.&#x20;

:link: [Instant security information with the Snyk security badge](https://snyk.io/blog/instant-security-information-with-the-snyk-security-badge/)

## Language and package manager considerations

### devDependencies analysis

devDependencies analysis is disabled by default as these are not typically elevated to production, often seen as “noise” by both security and development. To enable testing on dev-dependencies:

* Use the **--dev** parameter for CLI and CI/CD integrations.
* For Git integrations, set using **Settings > Languages** in the relevant configuration item.

### npm

Snyk can build a dependency tree with or without a lockfile. If a lockfile is present, this will be used.

* **Locally and CI/CD**: If a lockfile is not present and it is CLI/IDE, Snyk will look at `node_modules`  to determine what’s installed.
* **Git integration**: If a lockfile is not present, Snyk will approximate what the tree will look like at build time. This is highly valuable for getting insights into Projects in development or what the next build will look like when there is no lockfile present

As a user of npm, you may ask, “Why Snyk?” when npm-audit is at hand anytime you are working with your dependencies. You get the following capabilities:

* Snyk helps secure not only open source but also your first-party code. If you are using infrastructure as code and/or containers, Snyk can also provide visibility and remediation advice.
* It’s designed both for individuals and companies.
* In the context of Open Source:
  * You receive all the benefits of the curation, updates, and additional value that the Snyk Security Team adds, such as Known Exploit and Trending on Twitter.
  * You have Automated Remediation.
* Central reporting.
* Git Code repository integration, but not just there, Snyk has integrations across your pipeline and visibility into production.
* Broad support across programming languages and package managers.
* Ignore capabilities.

:link: [Snyk for npm](./#npm)

### Yarn

Requires yarn.lock and package.json

* Yarn 1 , 2, and 3 are fully supported in GIT and CLI.
* If a lock file is not present in CLI, the `node_modules` folder will be used to construct the dependency tree.
* `nohoist` is not supported for Yarn Workspaces..

:link: [Snyk for Yarn](./#yarn)

### Lerna/PNPM

Not officially supported, but if configured with Yarn workspaces, you may get Snyk IDE/CLI results.

### Unmanaged JavaScript

If you can get a full list of dependencies and their transitive dependencies, several API options are available.

{% hint style="info" %}
Snyk API access is required, available with the Snyk Enterprise plan.
{% endhint %}

To test for vulnerabilities, the following options are available:

* [Test package](https://snyk.docs.apiary.io/#reference/test/npm/test-for-issues-in-a-public-package-by-name-and-version)
* [Test Depgraph](https://snyk.docs.apiary.io/#reference/test/dep-graph)
* [Test Purl](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)

### Out of Sync Lockfiles

Control behavior when the lockfile and package file are in sync can be done using:

* CLI additional values: `--strict-out-of-sync`, `--fail-on`
* WebUI for Git Scans:
  * **Settings > Language > Javascript**

## Snyk integrations common usage patterns

npm and Yarn are well-designed package managers. The main considerations for npm and Yarn packages are if there are multiple package managers or Projects in the same repository or build and what criteria you want to apply, such as the threshold severity of critical/high for each.

For source code scanning, this ecosystem is straightforward, with no special options required. So testing runs easily under Git and Snyk CLI with the basic features and commands.

The following are some common usage patterns used for CLI:

## Deployment and rollout recommendations

Startups, small teams, individuals, open source maintainers typically onboard their applications via Git, getting results in minutes and starting to address issues almost immediately. Small teams have the benefit of being agile and determining what works best for their workflow.

With large organizations using hundreds of applications, a slower approach is recommended to get developer buy-in/adoption and to ensure a positive rollout experience.

1. Typically, large organizations start with daily monitoring of applications via Git integration, only initially turning on PR checks for a few key applications.
2. As developers become familiar with Snyk's capabilities, you can widen the scope of applications with PR checks for gating or blocking builds if checks fail.
3. Some customers use CI/CD to passively monitor and then turn on gating by using the `snyk [product] test` commands.
4. If you onboard a large number of legacy applications, you can use [Priority Score](../../../manage-issues/issue-management/priority-score.md) (typically 700 as a starting place) or criteria like “Known exploit” or “Fix available” to define a starting point to engage developers to start fixing vulnerabilities on key applications.

## Snyk CLI Tips and Tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

### What to test

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../snyk-cli/cli-reference.md)

#### Open Source libraries

\
The `snyk test` command tests the first manifest it can find and performs a test on that singular entry point. To have Snyk analyze all manifests in the directory, use the following options:

* `--all-projects`: This option detects and scans all Yarn and other Projects in this directory.
* `--yarn-workspaces`: For Yarn Workspaces use the `--all-projects` flag to test and monitor your packages with other package managers or Yarn workspaces or use `--yarn-workspaces` to specifically scan Yarn Workspaces Projects only.

{% hint style="info" %}
If you are using a package manager that requires options, it’s suggested to target individually with `--file=`
{% endhint %}

#### Codebase

* Framework support - see [Snyk Code - Supported languages and frameworks](../../snyk-code/snyk-code-language-and-framework-support.md).
* Use the `snyk code test` command from the root of the Project to perform source code analysis.

#### Containers

* Snyk will automatically look for application (open source) vulnerabilities as part of a container scan. Consider having Snyk integrated via CLI, earlier in the pipeline and utilize this as an additional signal/insight of what’s in production.
* If you ship your Node.JS application in a container, be aware that you might also be bundling insecure packages (Linux, open source), alongside your application in addition to what is brought in by the container base image. The Snyk Container CLI can help you identify a base image that minimizes the attack surface of your application.
* For more information on how you can filter to the layer you wish to work on,such as identifying a secure base image to build off of, the layers you are responsible for, or application (OS) vulnerabilities, see [Snyk CLI for container security](../../../scan-containers/snyk-cli-for-container-security/)

#### Infrastructure as code

:link: [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/)

### **Helpful Options/Plugins**

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../snyk-api-info/other-tools/tool-snyk-filter.md).

## Reporting

For Snyk Open Source, the `snyk [product] monitor` command is used to push results from the CLI back to Snyk for reporting in the Snyk UI. Don’t forget to use `--org=` to indicate what Organization to place the monitored results in or retrieve test settings from during the test.

{% hint style="info" %}
Snyk Enterprise plan customers can access the [Snyk API](../../../snyk-api-info/) for reporting and extracting data.
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

1. Being Proactive: [Snyk Advisor](https://snyk.io/advisor) for identifying better packages to begin designing. As well as the Snyk Advisor site, many IDE plugins use this data source to populate information about the package in the IDE.
2. Remediation advice: Snyk provides this across integrations on the scan results screens. For example, this advice can calculate the top level package requiring an update in package.json, or how to update a line of code to make it secure.
3. Automation:

* You can enable automatic fix pull requests, created when a new vulnerability is detected with a fix available.
* You can enable dependency upgrade related pull requests, created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.

{% hint style="info" %}
Snyk suggests enabling automated PRs on a key Project to start before enabling globally to familiarize development on how to interact with Snyk and security issues.
{% endhint %}

## Additional security topics that Impact Node.Js and JavaScript developers

The following is a collection of articles from the Snyk Security team and Developer Relations that impact this ecosystem. For more industry, security, and technology-related articles, visit the [Snyk Blog](https://snyk.io/blog/):

* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Best practices for creating Modern npm package](https://snyk.io/blog/best-practices-create-modern-npm-package/)
* [Detect and prevent dependency confusion attacks on npm to maintain supply chain security](https://snyk.io/blog/detect-prevent-dependency-confusion-attacks-npm-supply-chain-security/)
* [Switching between Node.Js versions](https://snyk.io/blog/mastering-node-js-version-management-and-npm-registry-sources-like-a-pro/)
* [DevSecOps tools for open source projects in JavaScript and Node.Js](https://snyk.io/blog/devsecops-tools-for-open-source-projects-in-javascript-and-node-js/)
