# Snyk for JavaScript / Node.js developers

Use this guide to understand the best way to apply Snyk in your workflow, and to be aware of key considerations for your chosen technology stack.

### Developer-first approach

Snyk takes a developer-first approach to securing your application. This approach focuses less on overhead-heavy work (such as putting in hard QA gates), and more on building a secure application, providing visibility in a developer’s workflow, and providing actionable insights. The benefits of this approach are to engage developers in security practices as part of their development, without seeing security as a costly overhead.

So we recommend you use Snyk to focus on earlier enablement, not later enforcement. Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a project to minimize rework. Add a package: test it before writing the code that interfaces with it. Similarly after writing major sections of code, test it before moving on.

#### Focus of this guide

This guide focuses on scanning your application code, which are package manager and language specific, [Snyk Container](https://docs.snyk.io/scan-containers) and [Snyk Infrastructure as Code (IaC)](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code) also support your container and infrastructure as code needs.&#x20;

### Snyk overview

Before we get started, let's introduce Snyk, as some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in. Snyk [Pricing plans](https://snyk.io/plans) determine what features are available.

{% hint style="info" %}
Note that there is a monthly limit to the number of tests performed if a particular product is not purchased. You can find the limits outlined [here](https://snyk.io/plans/) in the first square.
{% endhint %}

#### Snyk products

* Snyk Code: scan your own code for security vulnerabilities using source code analysis.
* Snyk Open Source\*
  * Open Source vulnerability testing and monitoring (All plans)
  * Open Source dependency upgrade version bumping (All plans)
  * License Compliance (paid plans)
* Snyk Infrastructure as Code
  * Scan for configuration issues when you deploy your applications using Kubernetes deployment files, Terraform, or Cloudformation templates.&#x20;
* Snyk Container
  * Scan for issues with container images if you are building containers
* Snyk Cloud: Security from code to cloud and back
  * Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.

{% hint style="info" %}
\* Some capabilities may be limited for some languages and package managers.
{% endhint %}

### Product features to enhance your workflow

As you start planning and your code progresses through the pipeline, Snyk can provide different capabilities at each stage, to help you find and fix security issues.

#### Designing and planning

The following public resources are available for all users:

* [Snyk Advisor](https://snyk.io/advisor): Helps you pick healthy open source packages or base images to start developing with.
* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely.
* [Snyk Training](https://training.snyk.io/): Provides training on how to use Snyk.

#### Coding

{% hint style="info" %}
These capabilities focus on enablement, not enforcement. Use Snyk tools to test while working on a project, rather than testing after you commit the code and then discovering issues requiring code rework. Add a package and test it, before writing the code that interfaces with it. Similarly, after writing major sections of code, test the change before you continue.
{% endhint %}

The following capabilities are available for all Snyk users.

These capabilities focus on enablement, not enforcement. Use Snyk tools to test while working on a project, rather than testing after you commit the code and then discovering issues requiring code rework. Add a package and test it, before writing the code that interfaces with it. Similarly, after writing major sections of code, test the change before you continue.

* [IDE Plugins](../integrations/ide-tools/): for VS Code, IntelliJ, and others: Test your open source packages and first party code as you develop. Additionally test infrastructure as code (IaC) Kubernetes deployment files you create.
* [Snyk CLI](../snyk-cli/): A powerful terminal program that allows you to test locally on your machine. Very useful in testing containers and more complex IaC files that are templated with variables (such as Terraform plan files), as well as scanning open source and your own code.

#### Validating, monitoring, alerting and gating

The following capabilities are available for all Snyk users:

**With Git integrations**

<figure><img src="https://lh6.googleusercontent.com/EYPCKsyukOq5A9wNpYka8tUBa5FbzGQXrbmG2klrIigOxTNSInsA_Znj6P0jpGnBv7yRHAaiTsF_GX9Y9Zr1xdE35eZljg_1crKgqHBkhoZrEbvpTsdZstjXdVZ1hVF4jNyTgfLWbALbvqtDFbuI_ys" alt="PR Checks for Git integrations"><figcaption><p>PR Checks for Git integrations</p></figcaption></figure>

Snyk allows you to [run PR Checks](../scan-application-code/run-pr-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis, and show results.

These functions can be used for:

* Your code with Snyk Code
* Open Source with Snyk Open Source
  * Check for known vulnerabilities and create Fix Pull Requests (npm/Yarn 1,2,3)
  * Check for license compliance checks (Snyk Open Source)
  * Dependency upgrade - positioning updates to address technical debt (Snyk Open Source)

Additionally with the Git Integration, you can also monitor the following on a daily basis:

* Infrastructure as code (IaC) with Snyk Infrastructure as Code

**With CI/CD integrations**

Snyk can passively monitor, and provide a QA gate by failing build checks during testing for policy violations.

Snyk provides flexible capabilities, including:

* Dedicated plugins for Jenkins, Circle CI, and others (see relevant marketplaces)
* Using Github Actions
* The Snyk CLI can be used in most CI/CD systems (see [examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Fail the build based on criteria using options or the [snyk-filter](../snyk-api-info/other-tools/tool-snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* With Partner Platforms: Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk.

#### Production Monitoring

* (Snyk Enterprise plan only) Snyk can monitor container images and their open source/Linux based packages being used in production via Kubernetes integration, to notify customers of known vulnerabilities for applications in production.
* (All plans) Where a production integration does not exist, use the [snyk monitor](../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.

#### Package Registry Integrations (Artifactory/Nexus)

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

Snyk Open Source integrates with Artifactory and Nexus both as local gatekeeper, and interacting with the registry for security testing. Snyk uses this integration for fixes and also lock the Lockfile.

If your projects reference private dependencies in Artifactory or Nexus but you are not a Snyk Enterprise user, you can use the Snyk CLI in a properly configured local environment (such as your build pipeline) so these dependencies can be resolved and included in the test.

For more information, see:

* Package registry integrations: [Artifactory Registry setup](../integrations/private-registry-integrations/artifactory-repository-setup/) and [Nexus Repository Manager setup](../integrations/private-registry-integrations/nexus-repo-manager-setup.md).
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../integrations/gatekeeper-plugins/artifactory-gatekeeper-plugin-overview.md) and [Nexus Repository Manager Gatekeeper plugin](../integrations/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

For additional package registry integration capabilities see additional tools and security articles.

#### Gain recognition

* You can place Snyk on your public repositories and have a badge available indicating the security status of your package. See [Instant security information with the Snyk security badge](https://snyk.io/blog/instant-security-information-with-the-snyk-security-badge/).

### Language-specific and package manager-specific notes

#### devDependencies analysis

devDependencies analysis is disabled by default as these are not typically elevated to production, often seen as “noise” by both security and development. To enable testing on dev-dependencies:

* Use the **--dev** parameter for CLI and CI/CD integrations.
* For Git integrations, set using **Settings > Languages** in the relevant configuration item

#### Npm

Snyk can build a dependency tree with or without a lockfile

* If a lockfile is present this will be used.
  * Locally & CI/CD: If a lockfile is not present and it is CLI/IDE, Snyk will look at node\_modules to determine what’s installed.
  * Git integration: if a lockfile is not present , Snyk will approximate what the tree will look like at build time. This is highly valuable to get insights into projects in development or what the next build will look like when there is no lock file present
* As a user of npm you may ask “Why Snyk?” when npm-audit is at hand anytime you are working with your dependencies:
  * Snyk helps secure not only open source, but also your first party code. If you are using infrastructure as code and/or containers Snyk can also provide visibility and remediation advice.
  * It’s designed both for individuals and companies
  * In the context of Open Source
    * You receive all the benefits of the curation, updates and additional value that the Snyk Security Team adds, such as Known Exploit, Trending on Twitter, etc
    * Automated Remediation
  * Central reporting
  * Git Code repository integration, but not just there, Snyk has integrations across your pipeline and visibility into production
  * Broad support across programming languages and package managers
  * Ignore capabilities

Also see [Snyk for npm](broken-reference).

#### Yarn

Requires yarn.lock and package.json

* Yarn 1 , 2, and 3 are fully supported in GIT and CLI
* If a lock file is not present in CLI, then the node\_modules folder will be used to construct the dependency tree
* Note **nohoist** is not supported for Yarn Workspaces.

#### Lerna/PNPM

* Not officially supported, but if configured with Yarn workspaces you may get results with Snyk IDE/CLI

#### Unmanaged JavaScript

If you’re able to get a full list of dependencies and their transitives there are several API options available.

{% hint style="info" %}
Snyk API access is required, which is available with the Snyk Enterprise plan.
{% endhint %}

For example, to test for vulnerabilities; the following options are available:

* [Test package](https://snyk.docs.apiary.io/#reference/test/npm/test-for-issues-in-a-public-package-by-name-and-version)
* [Test Depgraph](https://snyk.docs.apiary.io/#reference/test/dep-graph)
* [Test Purl](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)

#### Out of Sync Lockfiles

Control behavior when the lockfile and package file are in sync, can be done using:

* CLI additional values: **--strict-out-of-sync, --fail-on**
* WebUI for Git Scans:
  * **Settings > Language > Javascript**

Also see [Snyk for Yarn](broken-reference).

### Snyk integrations common usage patterns

npm and Yarn are well designed package managers. The main considerations will be if there are multiple package managers or projects in the same repository or build, and what criteria you wish to apply (such as threshold of critical/high) for each.

For source code scanning, this ecosystem is straightforward, with no special options required. So testing runs easily under Git and Snyk CLI with the basic features and commands.

The following are some common usage patterns used for CLI:

#### Deployment and rollout recommendations

Startups, small teams, individuals, open source maintainers typically onboard their applications via Git getting results in minutes, and starting to address issues almost immediately. Small teams have the benefit of being nimble and determining what works best for their workflow.

With large organizations, using hundreds of applications, a slower approach is recommended, to get developer buy-in/adoption and to ensure a positive rollout experience.

1. Typically, large organizations start with daily monitoring of applications via Git integration, only initially turning on PR checks for a few key applications.
2. As developers become familiar with Snyk capabilities, you can widen the scope of applications with PR checks for gating or blocking builds if checks fail.
3. Some customers use CI/CD to passively monitor and then turn on gating by using the **snyk \[product] test** commands.
4. If you onboard a large number of legacy applications, you can use [Priority Score](../manage-issues/issue-management/priority-score.md) (typically 700 as a starting place) or criteria like “known exploit” or “fix available” to define a starting point to quickly engage developers to start fixing vulnerabilities on key applications.

### Snyk CLI Tips and Tricks

(Also see [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/).)

#### What to test

Use the **--help** option in the CLI or [docs](https://docs.snyk.io/snyk-cli/cli-reference) for details of Snyk CLI commands.

* Open Source libraries:\
  The “snyk test” command tests the first manifest it can find and performs a test on that singular entry point. To have Snyk analyze all manifests\* , use:
  * **--all-projects**: Use this option to detect and scan all Yarn and other projects in this directory.
  * **--yarn-workspaces**: For Yarn Workspaces use the **--all-projects** flag to test and monitor your packages with other package managers or Yarn workspaces or use **--yarn-workspaces** to specifically scan Yarn Workspaces Projects only.\
    \
    \*Note that if you are using a package manager that requires options, it’s suggested to target individually with **--file=**
* Your own code:
  * Framework support - see [Snyk Code - Supported languages and frameworks](../scan-application-code/snyk-code/snyk-code-language-and-framework-support.md).
  * Use the **snyk code test** command from the root of the project to perform source code analysis..
* Containers:
  * Note that Snyk will automatically look for application (open source) vulnerabilities as part of a container scan. It’s recommended to have Snyk integrated via CLI, earlier in the pipeline and utilize this as an additional signal/insight of what’s in production.\
    If you ship your Node.JS application in a container, be aware that you might also be bundling insecure packages (Linux, Open Source), alongside your application in addition to what is brought in by the container base image. The Snyk Container CLI can help you identify a base image that minimizes the attack surface of your application.
  * See [Snyk CLI for container security](../scan-containers/snyk-cli-for-container-security/) for more information on how you can filter to the layer you wish to work on, i.e. identifying a secure base image to build off of, the layers you are responsible for, or application (OS) vulnerabilities.
* Infrastructure as code:
  * See [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/)

**Helpful Options/Plugins**

* See [snyk-to-html](https://docs.snyk.io/scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature) plugin to help generate reports locally or at build time
* See **--json** and **--sarif** options for generating output that can be programmatically accessed
* See[ snyk-filter ](https://docs.snyk.io/other-tools/tool-snyk-filter)for advanced filtering options and [other tools](https://docs.snyk.io/other-tools) .

#### Reporting

For Snyk Open Source, the **snyk \[product] monitor** command is used to push results from the CLI back to Snyk for reporting in the Snyk UI. Don’t forget to use **--org=** to indicate what organization to place the monitored results in or retrieve test settings from during the test.

{% hint style="info" %}
Snyk Enterprise plan customers can access the [Snyk API](../snyk-api-info/) for reporting and extracting data.
{% endhint %}

#### Proxies

* See the [CLI proxy guide](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/proxy-configuration-for-snyk-cli) for more information

### How to fix issues using Snyk

There are two parts to this discussion: what to fix and how to fix.

When first scanning your application, you may see hundreds or thousands of issues. This is where prioritization of issues becomes most important.

#### What to fix - prioritization factors

Some tools only use the single factor of severity to prioritize issues, but this can still result in thousands of results, with no clear starting point to fix these issues.

Snyk provides more factors to help you prioritize issues, such as having a known exploit, is it fixable, social trending, and others. This can be done both at the project level when looking at a specific project, or Enterprise customers can prioritize across all your projects.

To see prioritization in action, see the [Prioritize issues in the Snyk Web UI](https://training.snyk.io/learn/video/prioritize-ui) training course.

#### How to fix - addressing issues

Snyk offers capabilities in this ecosystem to help address issues, both reactively and proactively:

1. Being Proactive: [Snyk Advisor](https://snyk.io/advisor) for identifying better packages to begin designing. As well as the Snyk Advisor site, many IDE plugins use this data source to populate information about the package in the IDE.
2. Remediation advice: Snyk provides this across integrations on the scan results screens. For example, this advice can calculate the top level package requiring an update in package.json, or how to update a line of code to make it secure.
3. Automation:

* You can enable automatic fix pull requests, created when a new vulnerability is detected with a fix available.
* You can enable dependency upgrade related pull requests, created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.\
  **Note**: Snyk suggests enabling automated PRs on a key project to start, before enabling globally to familiarize development on how to interact with Snyk and security issues.

## Additional security topics that Impact Node.Js and JavaScript developers

The following are a collection of articles from the Snyk Security team and Developer Relations that impact this ecosystem. For more industry, security, and technology related articles, visit the [Snyk Blog](https://snyk.io/blog/):

* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Best practices for creating Modern npm package](https://snyk.io/blog/best-practices-create-modern-npm-package/)
* [Detect and prevent dependency confusion attacks on npm to maintain supply chain security](https://snyk.io/blog/detect-prevent-dependency-confusion-attacks-npm-supply-chain-security/)
* [Switching between Node.Js versions](https://snyk.io/blog/mastering-node-js-version-management-and-npm-registry-sources-like-a-pro/)
* [DevSecOps tools for open source projects in JavaScript and Node.Js](https://snyk.io/blog/devsecops-tools-for-open-source-projects-in-javascript-and-node-js/)
