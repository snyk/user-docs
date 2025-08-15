# Guidance for Java and Kotlin

This guide is provided to help you apply Snyk effectively in your technology stack.

## Package Registry Integrations (Artifactory/Nexus) - Maven

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

* Snyk Open Source uses Artifactory or Nexus to resolve transitive dependencies through private packages.
* Snyk can be connected to a publicly available instance using username and password or a private server on your network using the Snyk Broker.
* Snyk Open Source provides integrations with Artifactory and Nexus, both as local gatekeepers and interacting with the registry for security testing. See [Nexus Repository Manager setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)

{% hint style="info" %}
If you are not a Snyk Enterprise user using Artifactory or Nexus, analysis is best performed using CLI, as the build system will retrieve the dependencies and be present locally.
{% endhint %}

For more information on package registry integrations, including Maven, see the following:

* Package registry integrations: [Nexus Repository Manager setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)
* [Artifactory Registry for Maven](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md)
* [Nexus Registry for Maven](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/nexus-repository-manager-for-maven.md)
* Nexus Container Registry: [Container security with Nexus integration](../../../scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-nexus-container-registry.md)
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/artifactory-gatekeeper-plugin.md) and [Nexus Repository Manager Gatekeeper plugin](../../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

## Language and package manager considerations

For open source, developers may have decided to use Maven or Gradle, which may impact how you best utilize Snyk to perform the analysis

*   Using Maven, or Gradle with a gradle.lockfile:

    The Git code repository integration is a great way to use Snyk and get visibility or you may decide to use CLI/IDE or CI/CD integrations to test/gate/monitor, or do both!
*   Using Gradle without a Gradle.lockfile:

    The full dependency tree may not be apparent or artifacts may be pulled in from external resources, so the CLI/IDE workflow (for local scans), and CI/CD is the recommended approach for analysis, otherwise you may not have a complete view of issues and dependencies.

### Maven

Snyk can generate a dependency tree from POM via the Git integration or CLI:

* Locally and using CI/CD: Snyk will interact with the package manager to produce a list of dependencies.
* Git integration: will approximate the build as if it were built at that time.

{% hint style="info" %}
Developer dependencies (`scope=test`) are ignored as they are not pushed to production and are generally considered noise. You can enable them in CLI by adding `--dev`
{% endhint %}

### **Gradle**

* Snyk will interact with the package manager to produce a list of dependencies.
* Typically Gradle will execute code and other actions during the build process that will impact the installed dependencies, so the CLI workflow is recommended if a gradle.lockfile is not present.

### Kotlin

The following manifest files are supported:

* build.gradle (Groovy DSL) for both SCM and CLI&#x20;
* build.gradle.kts (Kotlin DSL) for CLI only

See the [Java and Kotlin](./#features) page for more details about the supported features.

### **APIs**

Sometimes customers develop advanced dependency management strategies and may not necessarily use the standard and frequently used package managers. For that reason, Snyk has provided test APIs.

For on-time testing using the Snyk API, you can use the [Test](../../../snyk-api/reference/test-v1.md) endpoints. Examples include [Test for issues in a (Maven) public package by group id, artifact id and version](../../../snyk-api/reference/test-v1.md#test-maven-groupid-artifactid-version) and [List issues for a package](../../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues).

## Snyk Integrations and common usage patterns

Java provides powerful flexibility and configuration options for developers. There can be considerations for testing it, especially with Open Source.

### **Maven and Gradle Projects using gradle.lockfile**

Typically you can instrument testing as part of a build system or adopt a lockfile as part of their process.

* It is quite common for large organizations to monitor applications via Git integration, to begin with, daily monitoring, turning on PR checks for only key applications at the start.
* As developers become familiar with Snyk capabilities, they widen the scope of applications with PR checks for gating.
* Use CI/CD to passively monitor and then turn on gating by using the [snyk \[product\] test and monitor commands](../../../developer-tools/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).

### Gradle Projects without a lock file

* Use CI/CD to passively monitor and then turn on gating by using the [`snyk [product] test` and `monitor` commands](../../../developer-tools/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).
* Gating and failing the build typically will be turned on one Project to start so that everyone can become familiar with the process and use passive monitoring for the remainder of the portfolio.

## Snyk CLI Tips and Tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

### What to test

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md)

#### Testing your own code:

* Framework support - see [Snyk Code - Supported languages and frameworks](../../supported-languages-package-managers-and-frameworks.md).
* Use the `snyk code test` command from the root of the project to perform source code analysis.
* Use `--scan-all-unmanaged --all-projects` to recursively find all jars under the present working directory.

#### Open Source libraries

**Maven**

The `snyk test` command tests the first manifest it can find, and scans that singular entry point. To scan all manifests, follow these instructions:

* To scan aggregate projects, use the `--maven-aggregate-project` option\
  (for example, `snyk test --maven-aggregate-project`)
* To scan for all projects use `--all-projects` option:\
  (that is, `snyk test --all-projects`)

Snyk scans active profiles activated by default.

* Any additional Maven arguments can be passed, a common one is a non-standard settings.xml location. For example, `snyk test -- -s path/to/settings.xml`
* To scan a specific configuration, test a specific Maven profile using `-P [name]`. For example, use **`snyk test -- -P prod`** to scan the `prod` configuration.

**Gradle**

By default, Snyk CLI scans only the current project (the project in the root of the current folder), or the project that is specified by `--file=path/to/build.gradle`.

`--all-projects` can be used across all package managers, which also includes the behaviors of **--**`all-sub-projects`, mentioned below.

* To scan all projects at once (recommended), use the `--all-sub-projects` option:\
  (that is, `snyk test --all-sub-projects`). Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
* To scan a specific project (for example, myapp), use `--sub-project=` (that is, `snyk test --sub-project=myapp`).

To test specific configurations, see detailed examples here [Snyk for Java and Kotlin](./).

For Android Build variants, see [Snyk for Java and Kotlin](./).

**Unmanaged**

For more details on unmanaged Jars, navigate to the [Scan all unmanaged JAR files](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/scan-all-unmanaged-jar-files.md) page.

#### Testing containers

Snyk automatically looks for application (such as open source, maven, and npm) vulnerabilities as part of a container scan. We recommend you integrate via CLI or Registry earlier in the pipeline and use this as an additional signal/insight into what is in production.

See [Snyk CLI for container security](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) for more details.

#### Infrastructure as code

For more details, navigate to the [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/) page.

### Fixing vulnerabilities

For more details, navigate to the [Fixing vulnerabilities on Maven projects](https://snyk.io/blog/fixing-vulnerabilities-in-maven-projects/) page.

### Options and plugins

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).

Additionally, the Snyk team has built plugins to make it easy to integrate Snyk into your workflows:

* [Gradle Plugin](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) (Community project)
* [Maven Plugin](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)

## Additional security topics for Java developers

The following is a collection of articles from the Snyk Security team and Developer Relations that impact this ecosystem. For more industry, security, and technology-related articles

* [Snyk Blog](https://snyk.io/blog/)
* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Snyk for secure Java development](https://snyk.io/blog/snyk-for-secure-java-development/)
* [Advanced IntelliJ debugger features](https://snyk.io/blog/advanced-intellij-debugger-features/)
* [Spring4shell: the zero day RCE Spring Framework explained](https://snyk.io/blog/spring4shell-zero-day-rce-spring-framework-explained/)
* [Log4j vulnerability explained: Prevent Log4Shell RCE by updating to version 2.17.1](https://snyk.io/blog/log4j-rce-log4shell-vulnerability-cve-2021-44228/)
* [Best practices for managing Java dependencies](https://snyk.io/blog/best-practices-for-managing-java-dependencies/)
* [Exploring the Spring security authorization bypass (CVE-2022-31692)](https://snyk.io/blog/spring-security-authorization-bypass-cve-2022-31692/)\\
