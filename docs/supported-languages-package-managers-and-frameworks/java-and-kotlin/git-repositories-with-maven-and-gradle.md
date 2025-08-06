# SCM integrations with Maven and Gradle

## SCM integrations available for Maven and Gradle Projects

### Maven Projects

Snyk creates a Project per `pom.xml` file when it scans Maven applications. The Project includes all direct and indirect dependencies associated with that file.

The Project includes only the production dependencies in the `compile`, `provided`, and `runtime` scopes.

### Gradle Projects

After you select a Project for import, Snyk builds the dependency tree based on the `build.gradle` file and (optional) `gradle.lockfile`.

Improved scanning for Gradle Projects (including Groovy and Kotlin DSLs) is now in Early Access as explained [on this page](git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning).

Only production dependencies in the `api`, `compile`, `classpath`, `implementation`, `runtime` and `runtimeOnly` configurations are included.

If possible, enable [Gradle lockfiles](https://docs.gradle.org/current/userguide/dependency_locking.html) in your application. When present, Snyk can more accurately resolve the final version of dependencies used in the Project.

For Gradle projects without lockfiles, Snyk recommends using the Snyk CLI for the most accurate results.

## Improved Gradle SCM scanning

{% hint style="info" %}
**Release status**

Improved Gradle SCM scanning is in Early Access. You can enable the feature by using [Snyk Preview](../../snyk-admin/snyk-preview.md).
{% endhint %}

You can now obtain more accurate results for your Gradle Projects imported through Git integrations by using Improved Gradle SCM scanning.

### Supported Gradle features

The following lists some of the main supported Gradle features:

* [Groovy](https://docs.gradle.org/current/userguide/groovy_build_script_primer.html) and [Kotlin](https://docs.gradle.org/current/userguide/kotlin_dsl.html) DSLs - `build.gradle(.kts)` and `settings.gradle(.kts)`
* [Built-in](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_public_repository) and [custom](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_custom_repository) package repositories, or example, Artifactory, Nexus
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
* [Spring's `mavenBom`](https://docs.spring.io/dependency-management-plugin/docs/current/reference/html/#dependency-management-configuration-bom-import)&#x20;
* [Spring Boot plugin BOMs](https://docs.spring.io/spring-boot/gradle-plugin/managing-dependencies.html)
* Maven BOMs as [`platform`](https://docs.gradle.org/current/userguide/platforms.html#sub:using-platform-to-control-transitive-deps) dependencies

Some Gradle features are not supported, and this may influence the scan results. These Gradle features include:

* Custom configuration in [buildSrc](https://docs.gradle.org/current/userguide/organizing_gradle_projects.html#sec:build_sources) directories
* Dependencies introduced via [plugins](https://docs.gradle.org/current/userguide/plugins.html).

### How to enable improved Gradle SCM scanning

{% hint style="warning" %}
Improved Gradle scanning supports importing a maximum limit of 5,000 `build.gradle(.kts)` files per Git repository. Attempts to import repos with more than 5,000 Gradle build files will fail.
{% endhint %}

To enable this feature, follow these steps for your Snyk Organisation:

1. Configure [package repository integrations](../../scan-with-snyk/snyk-open-source/package-repository-integrations/) (if you use Artifactory or Nexus, see [below](git-repositories-with-maven-and-gradle.md#package-repository-integrations)).
2. Enable [Workspaces for SCM integrations](../../scm-integrations/workspaces.md).
3. Enable **Improved Gradle scanning** in Snyk Preview.

After Improved Gradle SCM scanning is enabled:

* Previously imported Git repositories will have existing Gradle Groovy DSL Projects automatically updated on the next manual or recurring test.
* Re-import the repository to start seeing results for Gradle Kotlin DSL Projects.

## Configure language settings for Snyk for Java

Configure language settings for your open source and licensing at the Organization level. The configuration settings apply to all Projects in that Organization.

1. Open Snyk Web UI and navigate to **Settings** > **Languages**.
2. Under **Languages**, navigate to **Java** and select **Edit settings.**
3. Configure the settings for **Maven**.
4. **Update Settings** to save changes.

## Package repository integrations

{% hint style="info" %}
If your application build uses private package repositories, you must configure the relevant Snyk integration to get the most accurate results.
{% endhint %}

{% hint style="info" %}
To use package repository integrations with the Improved Gradle scanning Early Access feature, use the configuration instructions and settings for Maven.

These will be detected and used in improved Gradle scans.
{% endhint %}

In the Java language settings, you can integrate Snyk with your private package repositories (for example, Artifactory or Nexus).

This enables Snyk to build a complete dependency tree when scanning Maven or Gradle (Early Access) projects that reference private packages.

For more information, see [Artifactory Registry for Maven](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md) in the [Package repository integrations](../../scan-with-snyk/snyk-open-source/package-repository-integrations/).
