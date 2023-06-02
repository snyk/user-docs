# Snyk for Java and Kotlin

ProjectSnyk offers security scanning to test your Projects for vulnerabilities, both from the [Snyk CLI ](../../../snyk-cli/)and the [Snyk Web UI](../../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md) using different [Snyk Integrations](../../../integrations/).

## Features of Snyk for Java and Kotlin

The following tables provide an outline of the general features Snyk offers by language. In addition to these features, Snyk offers additional functionality related to the specific integration configurations.

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
Gradle Projects imported via Git are tested by parsing `build.gradle` files. As the only truly reliable way to resolve Gradle dependencies is to execute the tool itself, this method can sometimes provide incomplete results.

If possible, enable [lockfiles](snyk-for-java-and-kotlin.md#git-services-for-gradle-projects) in your Gradle Project to improve the accuracy for Git imports.

However, for the most accurate results, Snyk recommends using the [Snyk CLI](../../../snyk-cli/) to test Gradle Projects.
{% endhint %}

{% tabs %}
{% tab title="Java" %}
**Java**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | ✔︎          | ✔︎               | Fix advice only |
{% endtab %}

{% tab title="Kotlin" %}
**Kotlin**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          |             | ✔︎               | Fix advice only |
{% endtab %}
{% endtabs %}

## Supported versions

### Maven versions supported

* CLI - Maven `3.*` For details, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Git - Maven `3.*`

### Gradle versions supported

{% hint style="info" %}
Gradle 8 is not yet supported in the CLI.&#x20;

However, if your app does not use Gradle 8 specific features, it is generally possible to install Gradle 7 instead before running Snyk CLI scans
{% endhint %}

* CLI - Gradle `4.*`, `5.*`, `6.*`, `7.*` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).
* Git - Gradle `4.*`, `5.*`, `6.*`, `7.*`

## Snyk CLI for Java and Kotlin Projects (CI/CD)

The way Snyk analyzes and builds the dependencies varies depending on the language and package manager of the Project.

The Snyk CLI tests Maven and Gradle Projects as follows:

* Snyk CLI with Gradle: To build the dependency graph, Snyk integrates with Gradle and inspects the dependencies returned by the build. The following manifest files are supported: `build.gradle` (Groovy DSL) and `build.gradle.kts` (Kotlin DSL).
* Snyk CLI with Maven: To build the dependency tree, Snyk analyzes the output of the `pom.xml` files.

## CLI options for Java and Kotlin

This section describes the unique CLI commands available when working with Java-based Projects.

### Prerequisites for Snyk CLI with Java and Kotlin

* Install the relevant package manager before you use the Snyk CLI.
* Include the relevant manifest files supported by Snyk before testing.
* Install and authenticate the Snyk CLI to start analyzing Projects from your local environment. See [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

### Snyk CLI options for Maven and Gradle

For information about the `snyk test` and `snyk monitor` options available for use with Maven and Gradle, see the following pages:

* [Options for Maven Projects](https://docs.snyk.io/snyk-cli/commands/test#options-for-maven-projects) in the Test help
* [Options for Gradle ProjectsProjects ](https://docs.snyk.io/snyk-cli/commands/test#options-for-gradle-projects)in the Test help
* [Options for Maven Projects](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-maven-projects) in the Monitor help
* [Options for Gradle Projects](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-gradle-projects) in the Monitor help
* [Options for build tools](https://docs.snyk.io/snyk-cli/commands/test#options-for-build-tools) in the Test help
* [Options for build tools](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-build-tools) in the Monitor help
* [Options for unmanaged JAR files](../../../snyk-cli/commands/test.md#scan-all-unmanaged) in the Test help
* [Options for unmanaged JAR files](../../../snyk-cli/commands/monitor.md#scan-all-unmanaged) in the Monitor help

### Scan unmanaged JAR files&#x20;

If you are not using Maven or Gradle but have JAR files, you can identify the open-source packages these JAR files relate to using Snyk CLI.  For more information, see [Scan all unmanaged JAR files](../../../snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files.md).

### **Examples of how you can use Maven arguments with the Snyk CLI**

Test a specific Maven profile called “prod”.

```
snyk test -- -prod
```

Add a system property from your pom.xml file.

Example:

The package version appears in your pom.xml

```
${pkg_version}
```

Define the system property like this:

```
snyk test -- -Dpkg_version=1.4
```

## CLI help for Maven Projects: Aggregate Projects

A Maven aggregate Project is one that uses modules and inheritance.

When scanning these types of Projects, Snyk performs a compile to ensure all modules are resolvable by the Maven reactor.

*   To scan aggregate Projects, use the `--maven-aggregate-project` option:

    ```
    snyk test --maven-aggregate-project
    ```
*   To scan non-aggregate Projects, use `--all-projects` option:

    ```
    snyk test --all-projects
    ```

The same options can be used with `snyk monitor`.

Make sure to execute the options in the same directory as the root pom.xml file.

{% hint style="info" %}
Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
{% endhint %}

## CLI help for Gradle Projects

### Sub-projects

Gradle build can consist of several sub-projects, where each sub-project has its own build.gradle, while the root Project is the only one that also includes a `settings.gradle` file. Sub-projects depend on the root ProjectProjects but can be configured otherwise.

By default, Snyk CLI scans only the current Project, the Project in the root of the current folder, or the Project that is specified by `--file=path/to/build.gradle`).

*   To scan all Projects at once (recommended), use the `--all-sub-projects` option:

    ```
    snyk test --all-sub-projects
    ```

{% hint style="info" %}
Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
{% endhint %}

*   To scan a specific Project (for example, _myapp_):

    ```
    snyk test --sub-project=myapp
    ```

### Configurations for Gradle Project

Gradle dependencies are declared for a particular scope; each scope is represented by Gradle with the help of [Configurations](https://docs.gradle.org/current/userguide/declaring\_dependencies.html#sec:what-are-dependency-configurations). For example:

* **compileOnly** configuration for development dependencies
* **compile** configuration that includes compile and runtime dependencies

By default, Snyk merges all configurations returned by Gradle to a dependency graph based on the sum total of the dependencies across all configurations in the Project or Projects.

To test a specific configuration:

* Use the `--configuration-matching` option with a [Java regular expression](https://docs.oracle.com/javase/tutorial/essential/regex/) (case-insensitive) as its parameter. The configuration that matches is then tested. If several configurations match, they are all merged and resolved together. For example: `'^releaseRuntimeClasspath$|^compile$'`
* If the different sub-projects include different configurations, scan each sub-project separately.

### **Examples of how you can use the --configuration-matching option**

* `--configuration-matching=compile` will match compile, testCompile, compileOnly, and so on.
* `--configuration-matching=^compile$` will match only compile.
* `--configuration-matching='^(debug|release)compile$'` will match debugCompile and releaseCompile.
* `--configuration-matching='^(?!.*test).*$'` will match all configurations _except_ those containing the string "test".

### Android build variants

Android Gradle supports creating different versions of your app by configuring [build variants.](https://developer.android.com/studio/build/build-variants)

Because the Snyk default behavior is to merge all available configurations, the iterated variants cause a clash of configurations that can't be merged.

In these situations, the Snyk scan fails with an error from Gradle which may contain one of the following messages:

* _Cannot choose between the following configurations of `project :mymodulewithvariants`_
* _Cannot choose between the following variants of `project :mymodulewithvariants`_
* _Could not select value from candidates_

To avoid such conflicts:

*   **Use a specific configuration(s):** if you know of a build configuration that has all the required attributes and the configuration is identical across all sub-projects included in the test, specify that configuration.\
    For example:

    ```
    --configuration-matching=prodReleaseRuntimeClasspath
    ```
*   **Explicitly specify the dependency configuration:** modify intra-project dependencies in your build.gradle file(s) to use a specific configuration

    ```
      dependencies {
          implementation project(path: ':mymodulewithvariants', configuration: 'default')
      }
    ```
*   **Suggest configuration attributes:** if you receive an error when running the command, the error may indicate which attribute values are available, while the error details from Gradle also indicate which dependency variants match which attributes. Using these details, add the attribute filter option.\
    For example:

    ```
    snyk test --configuration-attributes=buildtype:release,usage:java-runtime,mode:demo
    ```

    matches the variants using `com.android.build.api.attributes.BuildTypeAttr=release` and `org.gradle.usage=java-runtime`

### Daemon

By default, Snyk passes `gradle build --no-daemon` in the background when running `snyk test` and `snyk monitor` on Windows. If, for some reason, you see `snyk test` or `snyk monitor` fail on other operating systems because of daemon-related issues, try adding the `--no-daemon` flag to the Snyk command or set `GRADLE_OPTS: '-Dorg.gradle.daemon=false'`. See the [Gradle documentation](https://docs.gradle.org/current/userguide/gradle\_daemon.html#sec:disabling\_the\_daemon) for tips on disabling the daemon.

### Lockfiles

If your Gradle Project makes use of a single **gradle.lockfile** or multiple **\*.lockfile** per configuration and you are having the following issue

**Gradle Error (short): > Could not resolve all dependencies for configuration ':compileOnly'. > Locking strict mode: Configuration ':compileOnly' is locked but does not have lock state.**

Note that that **compileOnly configuration** **has been deprecated,** and even if your Project successfully generates a lockfile, it will not contain the \`compileOnly\` state because this configuration cannot be resolved. Only resolvable configurations compute a dependency graph. In order to solve this issue, Snyk suggests you **update your build.gradle containing dependencyLocking logic with the following instruction:**

```
compileOnly {resolutionStrategy.deactivateDependencyLocking() }
```

This will **ignore compileOnly** and save only the necessary information to analyze your Project of Projects.

### Support for Gradle with Snyk

If you are having any trouble testing your Gradle Projects with Snyk, collect the following details and send them to Snyk at [support@snyk.io](mailto:support@snyk.io):

* `build.gradle`
* `settings.gradle` (especially if Snyk did not pick up a version of a package)
* The output from the following commands:
  * `$ snyk test -d`
  * `$ gradle dependencies -q`

## Git services for Maven Projects

After you select a Project for import, Snyk builds the dependency tree based on the `pom.xml` file.

## Git services for Gradle Projects

After you select a Project for import, Snyk builds the dependency tree based on the `build.gradle` file and (optional) `gradle.lockfile`.

{% hint style="info" %}
Only dependencies in the following configurations are included - `api`, `compile`, `classpath`, `implementation`, `runtime`, `runtimeOnly`.
{% endhint %}

If a lockfile is present, Snyk uses the lockfile to more accurately resolve the final version of dependencies used in the Project.

Gradle lockfiles are an opt-in feature that, among other benefits, enable reproducible builds. For more information, see the [Gradle docs on dependency locking](https://docs.gradle.org/current/userguide/dependency\_locking.html).

{% hint style="warning" %}
**Kotlin**: `build.gradle.kts` files are not currently supported in Git.
{% endhint %}

## Git settings for Java

From the Snyk UI, you can specify mirrors or repositories from which you’d like to resolve packages in Artifactory for Maven. For more information, see [Artifactory Registry for Maven](../../../integrations/private-registry-integrations/artifactory-repository-setup/artifactory-registry-for-maven.md).

## Additional Snyk support for Java

In addition to the CLI and Snyk UI features, you can also check your Java Projects using [Maven integration](../../../integrations/ci-cd-integrations/maven-plugin-integration.md).
