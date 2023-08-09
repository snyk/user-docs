# Snyk for .NET developers

## Introduction

Use this guide to understand the best way to apply Snyk in your workflow, and to be aware of key considerations for your chosen technology stack.&#x20;

This guide focuses on scanning your application code, which is specific to your package manager and language. [Snyk Container](https://docs.snyk.io/scan-containers) and [Snyk Infrastructure as Code (IaC)](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code) also support your container and Infrastructure as Code needs.&#x20;

### Developer-first approach

Snyk takes a developer-first approach to securing your application. This approach focuses less on overhead-heavy work (such as putting in hard QA gates), and more on building a secure application, providing visibility in a developer’s workflow, and providing actionable insights. The benefits of this approach are to engage developers in security practices as part of their development, without seeing security as a costly overhead.

So we recommend you use Snyk to focus on earlier enablement, not later enforcement. Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a project to minimize rework. Add a package: test it before writing the code that interfaces with it. Similarly after writing major sections of code, test it before moving on.

### Snyk overview

Before we get started, let's introduce Snyk, as some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in. Snyk [Pricing plans](https://snyk.io/plans) determine what features are available.

{% hint style="info" %}
Note that there is a monthly limit to the number of tests performed if a particular product is not purchased. You can find the limits outlined [here](https://snyk.io/plans/) in the first square.
{% endhint %}

**Snyk products**

* [Snyk Code](../scan-application-code/snyk-code/): scan your own code for security vulnerabilities using source code analysis.
* [Snyk Open Source](../scan-application-code/snyk-open-source/)\*
  * Open Source vulnerability testing and monitoring (All plans)
  * License Compliance (paid plans) (Nuget)
* [Snyk Container](../scan-containers/)
  * Scan for issues with container images if you are building containers
* [Snyk Infrastructure as Code](../scan-cloud-configurations/snyk-infrastructure-as-code/)
  * Scan for configuration issues when you deploy your applications using Azure Resource Manager, Kubernetes deployment files, Terraform, or AWS Cloudformation templates.
* [Integrated IaC with cloud context](../scan-cloud-configurations/snyk-iac+/): Security from code to cloud and back
  * Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.

{% hint style="info" %}
Some capabilities may be limited for some languages and package managers.
{% endhint %}

## Product features to enhance your workflow

As you start planning and your code progresses through the pipeline, Snyk can provide different capabilities at each stage, to help you find and fix security issues.

### Designing and planning

The following public resources are available for all users:

* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely.
* [Snyk Training](https://training.snyk.io/): Provides training on how to use Snyk.

### Coding

{% hint style="info" %}
These capabilities focus on enablement, not enforcement. Use Snyk tools to test while working on a project, rather than testing after you commit the code and then discovering issues requiring code rework. Add a package and test it, before writing the code that interfaces with it. Similarly, after writing major sections of code, test the change before you continue.
{% endhint %}

The following capabilities are available for all Snyk users.

* [Snyk CLI](../snyk-cli/):  A powerful terminal program that allows you to test locally on your machine. Very useful in testing containers and more complex IaC files that are templated with variables (such as Terraform plan files), as well as scanning open source and your own code.
* [IDE Plugins](../integrations/ide-tools/): for Visual Studio (Windows), VS Code, and others: Test your open source packages and first party code as you develop. Additionally test infrastructure as code (IaC) Kubernetes deployment files you create.

### Validating, monitoring, alerting, and gating

The following are available for all Snyk users.

#### With Git integrations

Snyk allows you to [run PR Checks](../scan-application-code/run-pr-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis, and show results.

These functions can be used for:

* Your code with Snyk Code
* [Azure Pipeline Integration](../integrations/ci-cd-integrations/azure-pipelines-integration/)
* Open Source with Snyk Open Source (for NuGet)
  * Check for known vulnerabilities
    * Create Fix Pull Requests to fix known vulnerabilities
  * License compliance checks (Snyk Open Source)

In addition to what is done with Code and Open source, with the Git Integration, you can monitor and alert on the following:

* Infrastructure as code (IaC) with Snyk Infrastructure as Code&#x20;

#### With CI/CD integrations&#x20;

Snyk can passively monitor, and/or provide a QA gate by failing build checks during testing for policy violations.&#x20;

Snyk provides flexible capabilities, including:

* CLI can be utilized in most CI/CD systems ([examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Fail the build based on criteria using options or the [snyk-filter](../snyk-api/other-tools/tool-snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* Partner Platforms - Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk or use the CLI.
* Using [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/)

### Production monitoring

* **Snyk Enterprise plan customers only**: Snyk can monitor container images and their open-source and Linux Linux-based packages being used in production via Kubernetes integration to notify customers of known vulnerabilities for applications in production.
* **All plans:** Where a production integration does not exist, use the [`snyk monitor`](../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.

## Language and package manager-specific facts for using Snyk

### C\#

Snyk Code can analyze your C# code using IDE, CLI, and Git integration.&#x20;

* Framework support - see [Snyk Code - Supported languages and frameworks](../scan-application-code/snyk-code/snyk-code-language-and-framework-support.md).

### Nuget

* **,Target Frameworks**: Snyk will identify the target frameworks and present results against each identified version when using the git integration.
* **Development dependencies**: By default Snyk generally does not scan developer dependencies, as they are not usually pushed to production and are seen as "noise". \
  Enable visibility in Nuget git import using the **Settings > Languages > .Net** settings. \
  Snyk will scan and fix build and `development Dependency` sections of your [`*.proj`](#user-content-fn-1)[^1], `packages.config` and `project.json` files
* **Lock files**: Currently **packages-lock.json** is not supported. Snyk interacts with the build system to determine the installed dependencies.
* **PackageReference:** Snyk currently requires a version attribute. If your Project lacks this, Snyk may fail to open a PR for your Project.
* **Git Analysis**
  * Dependency trees are created
    * For .NET Core, using the \*.proj files&#x20;
    * For .NET Framework, using the \*.proj file, and packages.config
  * Git integrations supports&#x20;
    * \*.csproj, \*.fsproj, \*.vbproj, packages.config
  * Fix Pull Requests
    * If you currently manage your Project dependencies with NuGet and leverage [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) Snyk can automatically update the dependency version in your manifest file, provided there is an actual fix for it. You can then review and merge your fixes.
*   **CLI Analysis**

    The CLI supports the following:

    * **project.assets.json**
      * Snyk can scan project.assets.json to determine dependencies, but the file must be generated. Similarly, if you point to the solution file (.sln), you must generate the file first
        * Run "**dotnet restore"** to generate the necessary `project.assets.json` prior to running the "**snyk test**" command.
        * The solution file contains pointers to the  files necessary to perform the analysis. Note that the projects themselves must have `project.assets.json` files to be able to be scanned. If you want Snyk to the solution file as an entry point for scanning, you can point the Snyk CLI to the solution file by using `--file=<filename>.sln`.
      * Where multiple target frameworks are used in the same project, the CLI scan returns results for the first target framework that is declared in the project.
    * packages.config
      * run "**nuget install -OutputDirectory packages**" prior to running the **snyk test** command

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment also known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

### Paket

* Snyk can analyze dependencies managed by Paket via the CLI.&#x20;
  * paket.dependencies and paket.lock must be present.
  * package.json is also supported but not recommended by Microsoft.
* For more information on Paket support, see [Snyk for .net](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net.md).

### Other

Sometimes customers develop advanced dependency management strategies and may not necessarily use the standard/well-traveled package managers. For that reason Snyk has provided test apis

* PURL See [Purl issues](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)&#x20;

### Build-time versus runtime dependencies

See [Snyk for .NET](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net.md) for more information

### Not Supported in .NET

* Private dependency scanning is not supported for the SCM integration. You can scan private dependencies using the Snyk CLI.
* `<ProjectReference>`elements are not currently supported.
* [`Directory.Build.props`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) and [`Directory.Build.targets`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) are not currently supported.

## Snyk Integrations and common usage patterns

### CLI Tips and Tricks

#### **Testing Your Code**

Run "snyk code test" from the root of your project

#### **Testing Open Source**

For open source analysis in the CLI, first install the dependencies. After installing the dependencies, use one of the following strategies when you go to run a test

* use **--file=** to target the solution file (.sln) or a specific file&#x20;
* use the **--all-projects** option to analyze your open source, especially if multiple languages/package managers/.sln files are involved.

#### **Containers**

See [Supported Operating System Distributions](../scan-containers/how-snyk-container-works/supported-operating-system-distributions.md).

See [Snyk CLI for container security](../scan-containers/snyk-cli-for-container-security/) for more details.

#### Infrastructure as code

See [Supported Providers](../scan-cloud-configurations/supported-iac-and-cloud-providers.md).

See [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/) for more details.

#### Helpful options and plugins

* See [snyk-to-html](../scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/) plugin to help generate reports locally or at build time
* See --json and --sarif options for generating output that can be programmatically accessed
* See [snyk-filter](../snyk-api/other-tools/tool-snyk-filter.md) for advanced filtering options and [other tools](../snyk-api/other-tools/)&#x20;

#### Reporting

The `snyk monitor` CLI command is used to push results from the CLI back to Snyk for reporting in the Snyk UI. Use **--org=** to indicate what organization to place the monitored results in or retrieve test settings from during the test.

{% hint style="info" %}
Snyk Enterprise customers can also access the [Snyk API](../snyk-api/) for reporting and extracting data.
{% endhint %}

#### Proxies

See the [CLI proxy guide](../snyk-cli/configure-the-snyk-cli/proxy-configuration-for-snyk-cli.md) for more information.

## Deployment and rollout recommendations

Startups, small teams, individuals, open-source maintainers typically onboard their applications via Git getting results in minutes, and starting to address issues almost immediately. Small teams have the benefit of being nimble and determining what works best for their workflow.

With large organizations and potentially hundreds of applications, a slower approach is recommended to get developer buy-in/adoption and to ensure a positive experience. Passive monitoring would be performed to start, with perhaps one or two key projects having gating enabled early on to familiarize everyone with the process. Gating using PR checks or blocking builds on a wider scope typically starts after the first 30 days.

If you onboard a large number of legacy applications, we suggest you use priority score (typically 700 as a starting place) or criteria such as “known exploit” or “fix available”, to define a starting point to quickly engage developers to start fixing vulnerabilities on key applications.

## How to fix issues using Snyk

This discussion has two sides: what to fix and how to fix it.

### What to fix - prioritization factors

Some tools only use the single factor of severity to prioritize issues, but this can still result in thousands of results, with no clear starting point to fix these issues.

Snyk provides more factors to help you prioritize issues, such as having a known exploit, is it fixable, social trending, and others. This can be done both at the project level when looking at a specific project, or Enterprise customers can prioritize across all your projects.

To see prioritization in action, see the [Prioritize issues in the Snyk Web UI](https://training.snyk.io/learn/video/prioritize-ui) training course.

### How to fix - addressing issues

Snyk offers capabilities in this ecosystem to help address issues, both reactively and proactively:

1. **Being proactive**:\
   Use the CLI and IDE plugins to test while developing.\
   Add a package, make sure it’s installed and scan for security, before writing your code.
2. **Remediation advice**:\
   Snyk supplies advice across integrations on the results screens that calculates the top level package requiring an update or how to update the line of code to make it secure.
3.  **Automation**: Under Git configuration options

    (Nuget) You can enable automatic fix pull requests that will be created when new known vulnerabilities are detected and there is a fix available

{% hint style="info" %}
Snyk suggests enabling automated PRs on a key Project to start before enabling globally, to familiarize development with how to interact with Snyk and security issues.
{% endhint %}

## Additional security topics for .Net developers

The following is a collection of articles from the Snyk Security team and Developer Relations that relate to this ecosystem.&#x20;

* [Snyk Blog](https://snyk.io/blog/)
* [Snyk for .Net](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net.md)
* [Best Practices for Containerizing .NET applications](https://snyk.io/blog/best-practices-for-containerizing-net-applications/)

[^1]: 
