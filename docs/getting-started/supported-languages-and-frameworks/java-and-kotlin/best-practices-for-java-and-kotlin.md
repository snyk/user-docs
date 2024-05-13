# Best practices for Java and Kotlin

Use this guide to apply Snyk effectively in your technology stack.

## Package Registry Integrations (Artifactory/Nexus) - Maven

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

* Snyk Open Source uses Artifactory or Nexus to resolve transitive dependencies through private packages.
* Snyk can be connected to a publicly available instance using username and password or a private server on your network using the Snyk Broker.
* Snyk Open Source provides integrations with Artifactory and Nexus, both as local gatekeepers and interacting with the registry for security testing. See [Nexus Repository Manager setup](../../../integrate-with-snyk/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/)

{% hint style="info" %}
If you are not a Snyk Enterprise user using Artifactory or Nexus, analysis is best performed via CLI, as the build system will retrieve the dependencies and be present locally.
{% endhint %}

For more information on package registry integrations, including Maven, see the following:

* Package registry integrations: [Nexus Repository Manager setup](../../../integrate-with-snyk/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/)
* [Artifactory Registry for Maven](../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md)
* [Nexus Registry for Maven](../../../integrate-with-snyk/package-repository-integrations/nexus-repository-manager-connection-setup/nexus-repository-manager-for-maven.md)
* Nexus Container Registry: [Container security with Nexus integration](../../../integrate-with-snyk/snyk-container-integrations/container-security-with-nexus-integration.md)
* Gatekeeper plugins: [Artifactory Gatekeeper plugin](../../../integrate-with-snyk/gatekeeper-plugins/artifactory-gatekeeper-plugin.md) and [Nexus Repository Manager Gatekeeper plugin](../../../integrate-with-snyk/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

## Language and package manager considerations

For open source, developers may have decided to use Maven or Gradle, which may impact how you best utilize Snyk to perform the analysis

*   **Using** **Maven, or Gradle with a gradle.lockfile**:

    The Git code repository integration is a great way to use Snyk and get visibility or you may decide to use CLI/IDE or CI/CD integrations to test/gate/monitor, or do both!
*   **Using Gradle** **without a Gradle.lockfile**:

    The full dependency tree may not be apparent or artifacts may be pulled in from external resources, so the CLI/IDE workflow (for local scans), and CI/CD is the recommended approach for analysis, otherwise you may not have a complete view of issues and dependencies.

### **Maven**

Snyk can generate a dependency tree from POM via the Git integration or CLI:

* Locally and using CI/CD: Snyk will interact with the package manager to produce a list of dependencies.
* Git integration: will approximate the build as if it were built at that time.

{% hint style="info" %}
Developer dependencies (`scope=test`) are ignored as they are not pushed to production and are generally considered noise. You can enable them in CLI by adding `--dev`
{% endhint %}

### **Gradle**

* Snyk will interact with the package manager to produce a list of dependencies.
* Typically Gradle will execute code and other actions during the build process that will impact the installed dependencies, so the CLI workflow is recommended if a gradle.lockfile is not present.

### [**Kotlin**](#user-content-fn-1)[^1]

The following manifest files are supported:

* build.gradle (Groovy DSL) for both SCM and CLI&#x20;
* build.gradle.kts (Kotlin DSL) for CLI only

See the [Java and Kotlin](./#open-source-and-licensing) page for more details about the supported features.

### **Other**

Sometimes customers develop advanced dependency management strategies and may not necessarily use the standard/well-traveled package managers. For that reason, Snyk has provided test APIs

* One-time testing API
  * Test - will generate a dependency tree. See [Test](https://snyk.docs.apiary.io/#reference/test) docs.
    * [Maven testing endpoint](https://snyk.docs.apiary.io/#reference/test/maven/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)
  * [PURL for maven](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)

## Snyk Integrations and common usage patterns

Java provides powerful flexibility and configuration options for developers. There can be considerations for testing it, especially with Open Source.

### **Maven and Gradle Projects using gradle.lockfile**

Typically you can instrument testing as part of a build system or adopt a lockfile as part of their process.

* It is quite common for large organizations to monitor applications via Git integration, to begin with, daily monitoring, turning on PR checks for only key applications at the start.
* As developers become familiar with Snyk capabilities, they widen the scope of applications with PR checks for gating.
* Use CI/CD to passively monitor and then turn on gating by using the [snyk \[product\] test and monitor commandsd](../../../integrate-with-snyk/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).

### Gradle Projects without a lock file

* Use CI/CD to passively monitor and then turn on gating by using the [snyk \[product\] test and monitor commandsd](../../../integrate-with-snyk/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).
* Gating and failing the build typically will be turned on one Project to start so that everyone can become familiar with the process and use passive monitoring for the remainder of the portfolio.

## Snyk CLI Tips and Tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

### What to test

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../snyk-cli/cli-commands-and-options-summary.md)

#### Testing your own code:

* Framework support - see [Snyk Code - Supported languages and frameworks](../../../scan-using-snyk/supported-languages-frameworks-and-feature-availability-overview/).
* Use the **snyk code test** command from the root of the project to perform source code analysis.
* Use **--scan-all-unmanaged --all-projects** to recursively find all jars under the present working directory.

#### Open Source libraries

<table data-full-width="true"><thead><tr><th>Maven</th><th>Gradle</th><th>Unmanaged</th></tr></thead><tbody><tr><td><ul><li><p>The <strong>snyk test</strong> command tests the first manifest it can find, and scans that singular entry point. To scan all manifests:</p><ul><li>To scan aggregate projects, use the <strong>--maven-aggregate-project</strong> option<br>(for example, <strong>snyk test --maven-aggregate-project</strong>)</li><li>To scan for all projects use <strong>--all-projects</strong> option:<br>(that is, <strong>snyk test --all-projects</strong>)</li></ul></li></ul><ul><li><p>Snyk scans active profiles activated by default.</p><ul><li>Any additional Maven args can be passed, a common one is a non-standard settings.xml location. For example, <strong>snyk test -- -s path/to/settings.xml</strong></li></ul></li></ul><ul><li>To scan a specific configuration, test a specific Maven profile using <strong>-P [name]</strong>. For example, use <strong>snyk test -- -P prod</strong> to scan the <strong>prod</strong> configuration.</li></ul></td><td><p></p><ul><li>By default, Snyk CLI scans only the current project (the project in the root of the current folder), or the project that is specified by <strong>--file=path/to/build.gradle</strong>.</li></ul><ul><li><p>As above, <strong>--all-projects</strong> can be used across all package managers, which also includes the behaviors of <strong>--all-sub-projects</strong>, mentioned below.</p><ul><li>To scan all projects at once (recommended), use the <strong>--all-sub-projects</strong> option:<br>(that is, <strong>snyk test --all-sub-projects</strong>). Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.</li><li>To scan a specific project (for example, myapp), use <strong>--sub-project=</strong> (that is, <strong>snyk test --sub-project=myapp</strong>).</li></ul></li></ul><ul><li>To test specific configurations, see detailed examples here <a href="./">Snyk for Java and Kotlin</a>.</li></ul><ul><li>For Android Build variants, see <a href="./">Snyk for Java and Kotlin</a>.</li></ul></td><td>Unmanaged Jars - see <a href="../../../snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files.md">Scan all unmanaged JAR files</a>.</td></tr></tbody></table>

#### Testing containers

Snyk automatically looks for application (such as open source, maven, and npm) vulnerabilities as part of a container scan. We recommend you integrate via CLI or Registry earlier in the pipeline and use this as an additional signal/insight into what is in production.

See [Snyk CLI for container security](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) for more details.

#### Infrastructure as code

:link: [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/)

### Fixing vulnerabilities

:link: [Fixing vulnerabilities on Maven projects](https://snyk.io/blog/fixing-vulnerabilities-in-maven-projects/)

### Options and plugins

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).

Additionally, the Snyk team has built plugins to make it easy to integrate Snyk into your workflows:

* [**Gradle Plugin**](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) **(Community project)**
* [**Maven Plugin**](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)

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

[^1]: this makes it look like Kotlin is a separate tool like Maven or Gradle. I'd remove the heading, the supported build files are part of Gradle
