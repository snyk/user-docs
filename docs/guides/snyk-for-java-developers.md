# Snyk for Java developers

Use this guide to understand the best way to apply Snyk in your workflow, and to be aware of key considerations for your chosen technology stack.

## Developer-first approach

Snyk takes a developer-first approach to securing your application. This approach focuses less on overhead-heavy work (such as putting in hard QA gates), and more on building a secure application, providing visibility in a developer’s workflow, and providing actionable insights. The benefits of this approach are to engage developers in security practices as part of their development, without seeing security as a costly overhead.

So we recommend you use Snyk to focus on earlier enablement, not later enforcement. Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a project to minimize rework. Add a package: test it before writing the code that interfaces with it. Similarly after writing major sections of code, test it before moving on.

### Focus of this guide

This guide focuses on scanning your application code, which are package manager and language specific. [Snyk Container](https://docs.snyk.io/scan-containers) and [Snyk Infrastructure as Code (IaC)](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code) also support your container and infrastructure as code needs.

### Snyk overview

Before we get started, let's introduce Snyk, as some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in. Snyk [Pricing plans](https://snyk.io/plans) determine what features are available.

#### Snyk Plans and Products

Snyk has many integrations, designed to satisfy key workflow needs and provide visibility at various stages of development. Additionally each Snyk product provides key capabilities for the ecosystems you are working in.

The following product list indicates what each Snyk product can review with respect to what you are creating, and the [plan](https://snyk.io/plans) (Free, Team, legacy Business, and Enterprise) determines what features are available.

{% hint style="info" %}
Note that there is a monthly limit to the number of tests performed if a particular product is not purchased. You can find the limits outlined [here](https://snyk.io/plans/) in the first square.
{% endhint %}

* Snyk Code - scan your own code for security vulnerabilities using source code analysis.
* Snyk Open Source\*
  * Open Source vulnerability testing and monitoring (All plans)
  * Open Source dependency upgrade version bumping (All plans)(Maven)
  * License Compliance (Snyk Team, Legacy Business and Enterprise plans)
* Snyk Infrastructure as Code
  * If you deploy your applications using Kubernetes deployment files, Terraform, or AWS Cloudformation templates, Snyk can scan for configuration issues that may lead to security vulnerabilities. [Supported Environments](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code/snyk-iac-supported-environments).
* Snyk Container
  * Analyze your container image to make suggestions on utilizing better base images, and identifying known vulnerabilities against container or open source packages in the layers you are creating
* Snyk Cloud: Security from code to cloud and back
  * Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source

{% hint style="info" %}
\*Some capabilities may be limited for some languages and package managers
{% endhint %}

### Language and package manager considerations

Analyzing your Java source code with Snyk Code is an easy capability to implement, requiring little configuration. So you can use standard [Snyk IDE](../integrations/ide-tools/) functions, standard [Snyk CLI](../snyk-cli/) (testing and monitoring) commands, and you can scan repositories using standard Snyk [Git repository integration](../integrations/git-repository-scm-integrations/).

{% hint style="warning" %}
Kotlin is currently in beta support for Snyk Code. Contact Snyk for more information.
{% endhint %}

### Maven and Gradle

For open source, developers may have decided to use Maven or Gradle, which may impact how you best utilize Snyk to perform the analysis

*   **Using** **Maven, or Gradle with a gradle.lockfile**:

    The Git code repository integration is a great way to use Snyk and get visibility or you may decide to use CLI/IDE or CI/CD integrations to test/gate/monitor, or do both!
*   **Using Gradle** **without a Gradle.lockfile**:

    The full dependency tree may not be apparent or artifacts may be pulled in from external resources, so the CLI/IDE workflow (for local scans), and CI/CD is the recommended approach for analysis, otherwise you may not have a complete view of issues and dependencies.

### Product features to enhance your workflow

As you start planning and your code progresses through the pipeline, Snyk can provide different capabilities at each stage, to help you find and fix security issues.

#### Designing and planning

The following public resources are available for all users:

* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely.
* [Snyk Training](https://training.snyk.io/): Provides training on how to use Snyk.
* [Snyk Advisor](https://snyk.io/advisor): Helps you pick healthy base images to start developing with.

#### Coding

{% hint style="info" %}
These capabilities focus on enablement, not enforcement. Use Snyk tools to test while working on a project, rather than testing after you commit the code and then discovering issues requiring code rework. Add a package and test it, before writing the code that interfaces with it. Similarly, after writing major sections of code, test the change before you continue.
{% endhint %}

The following capabilities are available for all Snyk users.

* [Snyk CLI](../snyk-cli/): A powerful terminal program that allows you to test locally on your machine. Very useful in testing containers and more complex IaC files that are templated with variables (such as Terraform plan files), as well as scanning open source and your own code.
* [IDE Plugins](../integrations/ide-tools/): for IntelliJ, Eclipse, VS Code, and others: Test your open source packages and first party code as you develop. Additionally test infrastructure as code (IaC) Kubernetes deployment files you create.
* Additionally the Snyk team has built plugins to make it easy to integrate Snyk into your workflows:
  * [**Gradle Plugin**](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) **(Community project)**
  * [**Maven Plugin**](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)

#### Validating, Monitoring, Alerting and Gating

The following capabilities are available for all Snyk users:

**With Git integrations**

Snyk allows you to [run PR Checks](../scan-application-code/run-pr-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis, and show results.

These results are viewable on the Snyk projects screen, for:

* Your code with Snyk Code
* Open Source with Snyk Open Source
  * Check for known vulnerabilities (Snyk Open Source)
    * Create Fix Pull Requests to fix known vulnerabilities (Maven)
  * License compliance checks (Snyk Open Source)(Maven)
  * Dependency upgrade - positioning updates to address technical debt (Snyk Open Source) (Maven)

Additionally with the Git Integration, you can also monitor the following on a daily basis:

* Infrastructure as code (IaC) with Snyk Infrastructure as Code

**With CI/CD integrations**

Snyk can passively monitor, and/or provide a QA gate by failing build checks during testing for policy violations.

Snyk provides flexible capabilities, including:

* [Gradle Plugins](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) **(Community project)**
* [Maven Plugins](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)
* Dedicated plugins for Jenkins, Circle CI, and others (see relevant marketplaces)
* Using [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/)
* The Snyk CLI can be used in most CI/CD systems (see [examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Fail the build based on criteria using options or the [snyk-filter](../snyk-api-info/other-tools/tool-snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* With Partner Platforms: Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk.
  * Note for Java: using the Git integration with Bitbucket Cloud or using the CLI instead of the prepackaged Bitbucket Pipe is suggested

#### Production monitoring

* (Snyk Enterprise plan only) Snyk can monitor container images and their open source/Linux based packages being used in production via Kubernetes integration, to notify customers of known vulnerabilities for applications in production.
* (All plans) Where a production integration does not exist, use the [snyk monitor](../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production.

### Package Registry Integrations (Artifactory/Nexus) - Maven

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

* Snyk Open Source uses Artifactory or Nexus to resolve transitive dependencies through private packages.
* Snyk can be connected to a publicly available instance using username and password or a private server on your network using the Snyk Broker..
* Snyk Open Source provides integrations with Artifactory and Nexus both as local gatekeeper, and interacting with the registry for security testing. See [Nexus Repository Manager setup](../integrations/package-repository-integrations/nexus-repo-manager-setup/) and [Artifactory Registry setup](../integrations/package-repository-integrations/artifactory-repository-setup/)

{% hint style="info" %}
If you are not a Snyk Enterprise user and you use Artifactory or Nexus, analysis is best performed via CLI as the build system will retrieve the dependencies and be present locally.
{% endhint %}

#### For more information

* Package registry integrations: [Nexus Repository Manager setup](../integrations/package-repository-integrations/nexus-repo-manager-setup/) and [Artifactory Registry setup](../integrations/package-repository-integrations/artifactory-repository-setup/)
* [Artifactory Registry for Maven](../integrations/package-repository-integrations/artifactory-repository-setup/artifactory-registry-for-maven.md)
* [Nexus Registry for Maven](../integrations/package-repository-integrations/nexus-repo-manager-setup/artifactory-registry-for-maven-1.md)
* Nexus Container Registry: [Container security with Nexus integration](../scan-containers/image-scanning-library/nexus-image-scanning/container-security-with-nexus-integration.md)
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../integrations/gatekeeper-plugins/artifactory-gatekeeper-plugin-overview.md) and [Nexus Repository Manager Gatekeeper plugin](../integrations/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

#### Recognition

You can place Snyk on your public repositories and have a badge available indicating the security status of your package. See [Instant security information with the Snyk security badge](https://snyk.io/blog/instant-security-information-with-the-snyk-security-badge/).

### Language-specific and package manager-specific notes

**Maven**

Snyk can generate a dependency tree from POM via the Git integration or CLI:

* Locally and using CI/CD: Snyk will interact with the package manager to produce a list of dependencies.
* Git integration: will approximate the build as if it were built at that time.

{% hint style="info" %}
Developer dependencies (that is, scope=test) are ignored as they are not pushed to production and are generally considered noise. You can enable them in CLI by adding “--dev”
{% endhint %}

**Gradle**

* Snyk will interact with the package manager to produce a list of dependencies
* Typically Gradle will execute code and other actions during the build process that will impact the installed dependencies, so the CLI workflow is recommended if a gradle.lockfile is not present.

[**Kotlin**](#user-content-fn-1)[^1]

The following manifest files are supported:

* build.gradle (Groovy DSL)
* build.gradle.kts (Kotlin DSL)

**Other**

Sometimes customers develop advanced dependency management strategies and may not necessarily use the standard/well-traveled package managers. For that reason Snyk has provided test APIs

* One-time testing API
  * Test - will generate a dependency tree. See [Test](https://snyk.docs.apiary.io/#reference/test) docs.
    * [Maven testing endpoint](https://snyk.docs.apiary.io/#reference/test/maven/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)
  * [PURL for maven](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)

### Snyk Integrations and common usage patterns

Java provides powerful flexibility and configuration options for developers. There can be considerations on how you test it, especially with Open Source.

The following provides some guidance on general rollout recommendations, in addition to commands that are relevant for the Maven/Gradle ecosystems.

### Snyk CLI Tips and Tricks

(Also see [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/).)

#### What to test

Use the **--help** option in the CLI or [docs](https://docs.snyk.io/snyk-cli/cli-reference) for details of Snyk CLI commands.

#### Testing your own code:

* Framework support - see [Snyk Code - Supported languages and frameworks](../scan-application-code/snyk-code/snyk-code-language-and-framework-support.md)
* Use the **snyk code test** command from the root of the project to perform source code analysis.

#### Testing open source dependencies

* Maven
  * The **snyk test** command tests the first manifest it can find, and scans that singular entry point. To scan all manifests:
    * To scan aggregate projects, use the **--maven-aggregate-project** option\
      (for example, **snyk test --maven-aggregate-project**)
    * To scan for all projects use **--all-projects** option:\
      (that is, **snyk test --all-projects**)
  * Snyk scans active profiles activated by default.
    * Any additional Maven args can be passed, a common one is a non-standard settings.xml location. For example, **snyk test -- -s path/to/settings.xml**
  * To scan a specific configuration, test a specific Maven profile using **-P \[name]**. For example, use **snyk test -- -P prod** to scan the **prod** configuration.
* Gradle
  * By default, Snyk CLI scans only the current project (the project in the root of the current folder), or the project that is specified by **--file=path/to/build.gradle**.
  * As above, **--all-projects** can be used across all package managers, which also includes the behaviors of **--all-sub-projects**, mentioned below.
    * To scan all projects at once (recommended), use the **--all-sub-projects** option:\
      (that is, **snyk test --all-sub-projects**). Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
    * To scan a specific project (for example, myapp), use **--sub-project=** (that is, **snyk test --sub-project=myapp**)
  * To test specific configurations, see detailed examples here [Snyk for Java and Kotlin](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-java-and-kotlin.md#configurations-for-gradle-project)
  * Android Build variants, see [Snyk for Java and Kotlin](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-java-and-kotlin.md#configurations-for-gradle-project)
  * Unmanaged Jars - see [Scan all unmanaged JAR files](../snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files.md).
    * Use **--scan-all-unmanaged --all-projects** to recursively find all jars under the present working directory.

#### Testing containers

Snyk automatically looks for application (such as open source, maven and npm) vulnerabilities as part of a container scan. We recommend you integrate via CLI or Registry earlier in the pipeline and use this as an additional signal/insight of what is in production.

See [Snyk CLI for container security](../scan-containers/snyk-cli-for-container-security/) for more details.

#### Testing Infrastructure as code:

See [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/) for more details.

#### Helpful Options/Plugins

* See [snyk-to-html](https://docs.snyk.io/scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature) plugin to help generate reports locally or at build time
* See --json and --sarif options for generating output that can be programmatically accessed
* See[ snyk-filter ](https://docs.snyk.io/other-tools/tool-snyk-filter)for advanced filtering options and [other tools](https://docs.snyk.io/other-tools).

#### Reporting

For Snyk Open Source, the **snyk \[product] monitor** command is used to push results from the CLI back to Snyk for reporting in the Snyk UI. Don’t forget to use **--org=** to indicate what Organization to place the monitored results in or retrieve test settings from during the test.

{% hint style="info" %}
Snyk Enterprise plan customers can access the [Snyk API](../snyk-api-info/) for reporting and extracting data.
{% endhint %}

#### Proxies

See the [CLI proxy guide](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/proxy-configuration-for-snyk-cli) for more information.

### Deployment and rollout recommendations

Startups, small teams, individuals, open source maintainers typically onboard their applications via Git getting results in minutes, and starting to address issues almost immediately. Small teams have the benefit of being nimble and determining what works best for their workflow.

With large organizations, and potentially hundreds of applications, a slower approach is recommended to get developer buy-in/adoption and to ensure a positive experience. Passive monitoring would be performed to start, with perhaps one or two key projects having gating enabled early on to familiarize everyone with the process. Gating using PR checks or blocking builds on a wider scope typically starts after the first 30 days.

If you onboard a large number of legacy applications, we suggest you use priority score (typically 700 as a starting place) or criteria such as “known exploit” or “fix available”, to define a starting point to quickly engage developers to start fixing vulnerabilities on key applications.

### **Maven and Gradle projects using gradle.lockfile**

Typically customers will instrument testing as part of a build system, or adopt a lockfile as part of their process.

* It is quite common for large organizations to monitor applications via Git integration to begin with using the daily monitoring, turning on PR checks for only key applications at the start.
* As developers become familiar with Snyk capabilities, widen the scope of applications with PR checks for gating.
* Some customers use CI/CD to passively monitor and then turn on gating by using the **snyk \[product] test** command.

### Gradle projects without a lock file

* Use CI/CD to passively monitor and then turn on gating by using the the **snyk \[product] test** command.
* Gating/failing the build typically will be turned on one project to start so that everyone can become familiar with the process, and use passive monitoring for the remainder of the portfolio.

### How to fix issues using Snyk

This discussion has two sides: what to fix and how to fix it.

#### What to fix - prioritization factors

Some tools only use the single factor of severity to prioritize issues, but this can still result in thousands of results, with no clear starting point to fix these issues.

Snyk provides more factors to help you prioritize issues, such as having a known exploit, is it fixable, social trending, and others. This can be done both at the project level when looking at a specific project, or Enterprise customers can prioritize across all your projects.

To see prioritization in action, see the [Prioritize issues in the Snyk Web UI](https://training.snyk.io/learn/video/prioritize-ui) training course.

#### How to fix - addressing issues

Snyk offers capabilities in this ecosystem to help address issues, both reactively and proactively:

1. **Being proactive**:\
   Use [Snyk Advisor](https://snyk.io/advisor) for identifying better packages to begin designing.\
   Use the CLI and IDE plugins to test while developing.\
   Add a package, make sure it’s installed and scan for security, before writing your code.
2. **Remediation advice**:\
   Snyk supplies this across integrations on the results screens that calculate the top-level package requiring an update in pom.xml or build.gradle or how to update the line of code to make it secure.
3. **Automation (Maven)**:
   * You can enable automatic fix pull requests that will be created when new known vulnerabilities are detected and there is a fix available.
   * When first importing a project into Snyk it may already have vulnerabilities. Instead of overwhelming you with fixes, Snyk will present the highest priority known vulnerability for you to fix first. Once fixed the next highest priority known vulnerability is presented, in this way you can catch up and fix existing issues. You can speed this process up by fixing all vulnerabilities in a single dependency at a time by enabling the setting **Clear backlog faster**.
   * You can enable dependency upgrade related pull requests, created when new versions of a package are available. This helps with technical debt by providing PR nudges to update dependencies.

{% hint style="info" %}
We suggest enabling automated PRs on a key project to start before enabling globally, to familiarize development on how to interact with Snyk and security issues.
{% endhint %}

Additional information: [Helpful guidance on fixing vulnerabilities on Maven projects](https://snyk.io/blog/fixing-vulnerabilities-in-maven-projects/)

### Additional security topics for Java developers

The following is a collection of articles from the Snyk Security team and Developer Relations that impact this ecosystem. For more industry, security, and technology-related articles

* [Snyk Blog](https://snyk.io/blog/)
* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Snyk for secure Java development](https://snyk.io/blog/snyk-for-secure-java-development/)
* [Advanced IntelliJ debugger features](https://snyk.io/blog/advanced-intellij-debugger-features/)
* [Spring4shell: the zero day RCE Spring Framework explained](https://snyk.io/blog/spring4shell-zero-day-rce-spring-framework-explained/)
* [Log4j vulnerability explained: Prevent Log4Shell RCE by updating to version 2.17.1](https://snyk.io/blog/log4j-rce-log4shell-vulnerability-cve-2021-44228/)
* [Best practices for managing Java dependencies](https://snyk.io/blog/best-practices-for-managing-java-dependencies/)
* [Exploring the Spring security authorization bypass (CVE-2022-31692)](https://snyk.io/blog/spring-security-authorization-bypass-cve-2022-31692/)\\

[^1]: this makes it look like Kotlin is a separate tool like Maven or Gradle. I'd remove the heading, the supported build files are part of Gradle
