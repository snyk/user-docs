# Best practices for JavaScript and Node.js

| Product                                                                                                                                                                                           | Description                                                                                                                                                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Snyk Code**                                                                                                                                                                                     | Scan your code for security vulnerabilities using source code analysis.                                                                                                                                                                                                                                                                                                                                                   |
| <p><strong>Snyk Open Source</strong><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> Some capabilities may be limited for some languages and package managers.</p> | <p></p><ul><li>Open Source vulnerability testing and monitoring (All plans).</li></ul><ul><li>Open Source dependency upgrade version bumping (All plans).</li></ul><ul><li>License Compliance (paid plans).</li></ul>                                                                                                                                                                                                     |
| **Snyk Infrastructure as Code**                                                                                                                                                                   | <p>Scan for configuration issues when you deploy your new applications using Kubernetes deployment files, Terraform, or Cloudformation templates. <br><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">Snyk CLI for Infrastructure as Code</a></p>                                                         |
| **Snyk Container**                                                                                                                                                                                | Scan for issues with container images if you are building containers.                                                                                                                                                                                                                                                                                                                                                     |
| **Snyk Integrated IaC with cloud context**                                                                                                                                                        | <p></p><ul><li>Security from code to cloud and back.</li></ul><ul><li>Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.</li></ul><p><span data-gb-custom-inline data-tag="emoji" data-code="1f517">🔗</span> <a href="../../../scan-with-snyk/snyk-iac/iac+-code-to-cloud-capabilities/">Snyk Integrated IaC with cloud context</a></p> |

Use this guide to apply Snyk effectively in your technology stack.

## Validating, monitoring, alerting, and gating

### **With Git integrations**

<figure><img src="https://lh6.googleusercontent.com/EYPCKsyukOq5A9wNpYka8tUBa5FbzGQXrbmG2klrIigOxTNSInsA_Znj6P0jpGnBv7yRHAaiTsF_GX9Y9Zr1xdE35eZljg_1crKgqHBkhoZrEbvpTsdZstjXdVZ1hVF4jNyTgfLWbALbvqtDFbuI_ys" alt="PR Checks for Git integrations"><figcaption><p>PR Checks for Git integrations</p></figcaption></figure>

* Snyk can monitor container images and their open source/Linux-based packages being used in production via Kubernetes integration to notify customers of known vulnerabilities for applications in production.\
  :information\_source: **Snyk Enterprise plan only**\

* Where a production integration does not exist, use the [snyk monitor](../../../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.\
  :information\_source: **All Snyk plans**

## Package Registry Integrations (Artifactory/Nexus)

Artifactory, Nexus, npm Teams and npm Enterprise Package Registry integrations are available to Snyk Enterprise plan users.

Snyk Open Source Gatekeeper plugins integrate with Artifactory and Nexus to block builds from downloading packages with vulnerability and license issues.

Snyk Open Source can also integrate with Artifactory, Nexus, npm Teams and npm Enterprise to assist in security testing of your applications. Snyk uses this integration for dependency resolution, fix calculation, and re-locking lock files.

If your Projects reference private dependencies in these repositories but you are not a Snyk Enterprise user, you can use the Snyk CLI in a properly configured local environment (such as your build pipeline) so these dependencies can be resolved and included in the test.

For more information, see the following:

* Package registry integrations: [npm Teams and npm Enterprise](../../../integrate-with-snyk/package-repository-integrations/npm-teams-and-npm-enterprise-integration.md), [Artifactory Registry setup](../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/) and [Nexus Repository Manager setup](../../../integrate-with-snyk/package-repository-integrations/nexus-repository-manager-connection-setup/).
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../../../integrate-with-snyk/gatekeeper-plugins/artifactory-gatekeeper-plugin.md) and [Nexus Repository Manager Gatekeeper plugin](../../../integrate-with-snyk/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

## Language and package manager considerations

### devDependencies analysis

**`devDependencies`** analysis is disabled by default as these are not typically elevated to production, often seen as “noise” by both security and development. To enable testing on dev-dependencies:

* Use the **--dev** parameter for CLI and CI/CD integrations.
* For Git integrations, set using **Settings > Languages** in the relevant configuration item.

### optionalDependencies analysis

optionalDependencies are included by default for CLI and CI/CD, as well as Git integrations.

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

For source code scanning, this ecosystem is straightforward, with no special options required. Testing runs under Git and Snyk CLI with the basic features and commands.

## Snyk CLI tips and tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

### What to test

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../snyk-cli/cli-commands-and-options-summary.md)

#### Open Source libraries

\
The `snyk test` command tests the first manifest it can find and performs a test on that singular entry point. To have Snyk analyze all manifests in the directory, use the following options:

* `--all-projects`: This option detects and scans all Yarn and other Projects in this directory.
* `--yarn-workspaces`: For Yarn Workspaces use the `--all-projects` flag to test and monitor your packages with other package managers or Yarn workspaces or use `--yarn-workspaces` to specifically scan Yarn Workspaces Projects only.

{% hint style="info" %}
If you are using a package manager that requires options, it’s suggested to target individually with `--file=`
{% endhint %}

#### Codebase

* Framework support - see [Supported languages, frameworks, and feature availability overview](../).
* Use the `snyk code test` command from the root of the Project to perform source code analysis.

#### Containers

* Snyk will automatically look for application (open source) vulnerabilities as part of a container scan. Consider having Snyk integrated via CLI earlier in the pipeline and utilize this for an additional signal of and insight into what is in production.
* If you ship your Node.JS application in a container, be aware that you might also be bundling insecure packages (Linux, open source), alongside your application in addition to what is brought in by the container base image. The Snyk Container CLI can help you identify a base image that minimizes the attack surface of your application.
* For more information on how you can filter to the layer you wish to work on,such as identifying a secure base image to build off of, the layers you are responsible for, or application (OS) vulnerabilities, see [Snyk CLI for container security](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)

#### Infrastructure as code

:link: [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/)

### **Options and plugins**

* To help generate reports locally or at build time, see the [snyk-to-html plugin](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).

## Additional security topics that Impact Node.Js and JavaScript developers

The following is a collection of articles from the Snyk Security team and Developer Relations that impact this ecosystem. For more industry, security, and technology-related articles, visit the [Snyk Blog](https://snyk.io/blog/):

* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Best practices for creating Modern npm package](https://snyk.io/blog/best-practices-create-modern-npm-package/)
* [Detect and prevent dependency confusion attacks on npm to maintain supply chain security](https://snyk.io/blog/detect-prevent-dependency-confusion-attacks-npm-supply-chain-security/)
* [Switching between Node.Js versions](https://snyk.io/blog/mastering-node-js-version-management-and-npm-registry-sources-like-a-pro/)
* [DevSecOps tools for open source projects in JavaScript and Node.Js](https://snyk.io/blog/devsecops-tools-for-open-source-projects-in-javascript-and-node-js/)
