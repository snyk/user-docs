# SCM integrations with Maven and Gradle

## Available SCM integrations

### Maven

When scanning Maven applications, Snyk creates a Project per `pom.xml` file. The Project includes all direct and indirect dependencies associated with that file.

The Project includes only the production dependencies in the `compile`, `provided`, and `runtime` scopes.

For Maven, Snyk can generate a dependency tree from POM through the SCM integration or the CLI:

* Locally and using CI/CD: Snyk interacts with the package manager to produce a list of dependencies.
* SCM integration: Snyk approximates the build as if it were built at that time.

{% hint style="info" %}
Developer dependencies (`scope=test`) are ignored as they are not pushed to production and are generally considered noise. You can enable them in CLI by adding `--dev`.

Because Maven resolves dependencies in the order it encounters them in the POM file, it is important that dev dependencies are listed last in the POM file. Not doing so can lead to dev dependencies being reported by Snyk.
{% endhint %}

### Gradle

For Gradle, Snyk interacts with the package manager to produce a list of dependencies. Typically Gradle executes code and other actions during the build process that impacts the installed dependencies, so Snyk recommends using the CLI if a gradle.lockfile is not present.

For Gradle, after you select a Project for import, Snyk builds the dependency tree based on the `build.gradle` file and (optional) `gradle.lockfile`.

Only production dependencies in the `api`, `compile`, `classpath`, `implementation`, `runtime` and `runtimeOnly` configurations are included.

If possible, enable [Gradle lockfiles](https://docs.gradle.org/current/userguide/dependency_locking.html) in your application. When present, Snyk can more accurately resolve the final version of dependencies used in the Project.

For Kotlin, the following manifest files are supported:

* build.gradle (Groovy DSL) for both SCM and CLI
* build.gradle.kts (Kotlin DSL) for CLI only

## **Maven and Gradle Projects using gradle.lockfile**

If you are using Maven or Gradle with a gradle.lockfile, the Git code repository integration is an efficient way to use Snyk and get visibility or you can use CLI/IDE or CI/CD integrations to test, gate, and, or monitor.

Typically you can instrument testing as part of a build system or adopt a lockfile as part of their process.

* It is quite common for large organizations to monitor applications via Git integration, to begin with, daily monitoring, turning on PR checks for only key applications at the start.
* As developers become familiar with Snyk capabilities, they widen the scope of applications with PR checks for gating.
* Use CI/CD to passively monitor and then turn on gating by using the [snyk \[product\] test and monitor commands](../../developer-tools/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).

## Gradle Projects without a lock file

If you're using Gradle without a Gradle.lockfile, it is possible that the full dependency tree is not apparent or artifacts are pulled in from external resources. Snyk recommends to:

* Use the CLI/IDE workflow for local scans
* Use CI/CD to passively monitor and then turn on gating by using the [`snyk [product] test` and `monitor` commands](../../developer-tools/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).
* Start with turning on gating and failing the build on one Project, so developers get familiar with the process and then use passive monitoring for the remainder of the portfolio.

## Improved Gradle SCM scanning

{% hint style="info" %}
**Release status**

Improved Gradle SCM scanning is in Early Access. You can enable the feature by using [Snyk Preview](../../snyk-platform-administration/snyk-preview.md).
{% endhint %}

### Supported Gradle features

* [Groovy](https://docs.gradle.org/current/userguide/groovy_build_script_primer.html) and [Kotlin](https://docs.gradle.org/current/userguide/kotlin_dsl.html) DSLs - `build.gradle(.kts)` and `settings.gradle(.kts)`
* [Built-in](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_public_repository) and [custom](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_custom_repository) package repositories, for example, Artifactory, Nexus
* Built-in objects [`ext`](https://docs.gradle.org/current/dsl/org.gradle.api.plugins.ExtraPropertiesExtension.html), [`project`](https://docs.gradle.org/current/dsl/org.gradle.api.Project.html), `rootProject`, and [`settings`](https://docs.gradle.org/current/dsl/org.gradle.api.initialization.Settings.html)
* Local and global variables, maps, and string interpolation
* `allprojects` and `subprojects` blocks
* Custom files referenced using `apply from`
* [Type-safe project accessors](https://docs.gradle.org/current/userguide/declaring_dependencies_basics.html#sec:type-safe-project-accessors)
* Gradle [lockfiles](https://docs.gradle.org/current/userguide/dependency_locking.html)
* [Gradle properties and system properties](https://docs.gradle.org/current/userguide/build_environment.html#sec:gradle_system_properties) - `gradle.properties`
* [Dependency exclusions](https://docs.gradle.org/current/userguide/dependency_downgrade_and_exclude.html#sec:excluding-transitive-deps)
* Version catalogs declared in [Gradle](https://docs.gradle.org/current/userguide/platforms.html#sub:version-catalog-declaration) and [TOML](https://docs.gradle.org/current/userguide/platforms.html#sub::toml-dependencies-format) files - `gradle/libs.versions.toml`
* [Multi-project builds](https://docs.gradle.org/current/userguide/declaring_dependencies_between_subprojects.html), project names, project references
* [Spring's `mavenBom`](https://docs.spring.io/dependency-management-plugin/docs/current/reference/html/#dependency-management-configuration-bom-import)
* [Spring Boot plugin BOMs](https://docs.spring.io/spring-boot/gradle-plugin/managing-dependencies.html)
* Maven BOMs as [`platform`](https://docs.gradle.org/current/userguide/platforms.html#sub:using-platform-to-control-transitive-deps) dependencies

The following Gradle features are not supported:

* Custom configuration in [buildSrc](https://docs.gradle.org/current/userguide/organizing_gradle_projects.html#sec:build_sources) directories
* Dependencies introduced via [plugins](https://docs.gradle.org/current/userguide/plugins.html).

### Enable improved Gradle SCM scanning

{% hint style="warning" %}
Improved Gradle SCM scanning supports importing a maximum limit of 5,000 `build.gradle(.kts)` files per SCM repository. Attempts to import repos with more than 5,000 Gradle build files will fail.
{% endhint %}

To enable this feature, follow these steps for your Snyk Organization:

1. Configure [package repository integrations](../../scan-with-snyk/snyk-open-source/package-repository-integrations/) (if you use Artifactory or Nexus, see [below](git-repositories-with-maven-and-gradle.md#package-repository-integrations)).
2. Enable [Workspaces for SCM integrations](../../developer-tools/scm-integrations/workspaces.md).
3. Enable **Improved Gradle scanning** in Snyk Preview.

After Improved Gradle SCM scanning is enabled:

* The SCM repositories that have been previously imported have existing Gradle Groovy DSL Projects automatically updated on the next manual or recurring test.
* Re-import the repository to start seeing results for Gradle Kotlin DSL Projects.

## Package repository integrations

If your application build uses private package repositories, you must configure the relevant Snyk integration to get the most accurate results.

To use package repository integrations with the Improved Gradle scanning feature, use the configuration instructions and settings for Maven. These will be detected and used in improved Gradle scans.

In the Java language settings, you can integrate Snyk with your private package repositories (for example, Artifactory or Nexus). This enables Snyk to build a complete dependency tree when scanning Maven or Gradle Projects that reference private packages.

## Package registry integrations (Artifactory/Nexus) - Maven

{% hint style="info" %}
Artifactory and Nexus package registry integrations are available only with Snyk Enterprise plans.
{% endhint %}

Snyk Open Source uses Artifactory or Nexus to resolve transitive dependencies through private packages.

Snyk can be connected to a publicly available instance using username and password or a private server on your network using the Snyk Broker.

Snyk Open Source provides integrations with Artifactory and Nexus, both as local gatekeepers and interacting with the registry for security testing. See [Nexus Repository Manager setup](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/).

{% hint style="info" %}
For users who do not have a Snyk Enterprise integration with Artifactory or Nexus, Snyk recommends using the CLI, as it works with the dependencies that your build system makes available locally.
{% endhint %}

For more information on package registry integrations, including Maven, see the following external resources:

* Package registry integrations: [Nexus Repository Manager setup](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)
* [Artifactory Registry for Maven](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md)
* [Nexus Registry for Maven](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/nexus-repository-manager-for-maven.md)
* Nexus Container Registry: [Container security with Nexus integration](../../scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-nexus-container-registry.md)
* Gatekeeper plugins: Artifactory Gatekeeper plugin

## Configure language settings for Snyk for Java

You can configure language settings for your open source libraries and licensing at the Organization level. The configuration settings apply to all Projects in that Organization. To configure:

1. In the Snyk Web UI, navigate to **Settings** > **Snyk Open Source** > **Languages** > **Java.**
2. Click **Edit settings.**
3. Configure the settings for **Maven**.
4. Click **Update Settings** to save changes.

## **APIs**

Customers develop advanced dependency management strategies and can choose not to use the standard and frequently used package managers.

For on-time testing using the Snyk API, you can use the [Test](../../snyk-api/reference/test-v1.md) endpoints. Examples include [Test for issues in a (Maven) public package by group id, artifact id and version](../../snyk-api/reference/test-v1.md#test-maven-groupid-artifactid-version) and [List issues for a package](../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues).
