# Snyk for Java and Kotlin

Snyk offers security scanning to test your projects for vulnerabilities, both from the [Snyk CLI ](../../../snyk-cli/)and the [Snyk Web UI](../../../snyk-web-ui/), using different [Snyk Integrations](../../../integrations/).

## Features

The following tables provide an outline of the general features Snyk offers by language. In addition to these features, we offer additional functionality related to the specific integration configurations.

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
Gradle projects imported via Git are tested by parsing `build.gradle` files. As the only truly reliable way to resolve Gradle dependencies is to execute the tool itself, this method can sometimes provide incomplete results.

If possible, enable [lockfiles](snyk-for-java-gradle-maven.md#git-services-for-gradle-projects) in your Gradle project to improve the accuracy for Git imports.

However, for the most accurate results, we recommend using the [Snyk CLI](../../../snyk-cli/) to test Gradle projects.
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

### Maven

* CLI - Maven `3.*` ([more details](https://github.com/snyk/snyk-mvn-plugin#support))
* Git - Maven `3.*`

### Gradle

* CLI - Gradle `4.*`, `5.*`, `6.*`,  `7.*` ([more details](https://github.com/snyk/snyk-gradle-plugin#support))
* Git - Gradle `4.*`, `5.*`, `6.*`

## Snyk CLI tool for Java and Kotlin projects (CI/CD)

The way Snyk analyzes and builds the dependencies varies depending on the language and package manager of the project.

Learn how to use the tool for your Java projects as follows:

* Snyk CLI with Gradle: To build the dependency graph, Snyk integrates with Gradle and inspects the dependencies returned by the build. The following manifest files are supported: `build.gradle` (Groovy DSL) and `build.gradle.kts` (Kotlin DSL).
* Snyk CLI with Maven: To build the dependency tree, Snyk analyzes the output of the `pom.xml` files.

## CLI parameters for Java and Kotlin

This section describes the unique CLI commands available when working with Java-based projects as follows:

### Prerequisites

* Install the relevant package manager before you use the Snyk CLI tool.
* Include the relevant manifest files supported by Snyk before testing.
* Install and authenticate the Snyk CLI to start analyzing projects from your local environment. See [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

### Snyk CLI parameters

When working with Gradle projects from our CLI, you can add any of the following options to further refine the way the scan works:

| **Option**                    | **Description**                                                                                                                                                                                                                                                                                     |
| ----------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--sub-project=`              | For Gradle "multi-project" configurations, test a specific sub-project.                                                                                                                                                                                                                             |
| `--all-sub-projects`          | For "multi-project" configurations, test all sub-projects.                                                                                                                                                                                                                                          |
| `--configuration-matching=`   | Resolve dependencies using only configuration(s) that match the provided Java regular expression. For example: `'^releaseRuntimeClasspath$'`                                                                                                                                                        |
| `--configuration-attributes=` | Select certain values of configuration attributes to resolve the dependencies. For example: `'buildtype:release,usage:java-runtime'`                                                                                                                                                                |
| `--all-projects`              | Use for monorepos. This will detect all supported manifests. For Gradle monorepos Snyk will only look for root level **build.gradle / build.gradle.kts** files and apply the same logic as `--all-sub-projects` behind the scenes. This command is designed to be run in the root of your monorepo. |
| `--maven-aggregate-project`   | Use this argument instead of `--all-projects` when scanning Maven aggregate projects.                                                                                                                                                                                                               |

### Pass extra arguments directly to Gradle or Maven via Snyk CLI

You can pass any extra Gradle or Maven arguments directly to **gradle** or **mvn** by providing them after a Snyk command like so:

```
snyk test -- --build-cache
```

**Examples of how you can use Maven arguments with the Snyk CLI**

Test a specific Maven profile called “prod”.

```
snyk test -- -prod
```

Add a system property from your pom.xml file.

For example:

The package version appears in your pom.xml

```
${pkg_version}
```

Define the system property like this:

```
snyk test -- -Dpkg_version=1.4
```

## CLI help for Maven projects

### Aggregate projects

A Maven aggregate project is one that uses modules and inheritance.&#x20;

When scanning these types of projects Snyk will perform a compile to ensure all modules are resolvable by the Maven reactor.&#x20;

*   To scan aggregate projects, use the `--maven-aggregate-project` param:

    ```
    snyk test --maven-aggregate-project
    ```
*   To scan non-aggregate projects use `--all-projects` param:

    ```
    snyk test --all-projects
    ```

The same param can be used with `snyk monitor`.

Make sure to execute this in the same directory as the root pom.xml file.

{% hint style="info" %}
Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
{% endhint %}

## CLI help for Gradle projects

### Sub-projects

Gradle build can consist of several sub-projects, where each sub-project has its own build.gradle, while the root project is the only one that also includes a `settings.gradle` file. Sub-projects depend on the root project, but can be configured otherwise.

By default, Snyk CLI scans only the current project (the project in the root of the current folder), or the project that is specified by `--file=path/to/build.gradle`).

*   To scan all projects at once (recommended), use the `--all-sub-projects` flag:

    ```
    snyk test --all-sub-projects
    ```

{% hint style="info" %}
Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
{% endhint %}

*   To scan a specific project (for example, _myapp_):

    ```
    snyk test --sub-project=myapp
    ```

### Configurations

Gradle dependencies are declared for a particular scope, each scope is represented by Gradle with the help of [Configurations](https://docs.gradle.org/current/userguide/declaring\_dependencies.html#sec:what-are-dependency-configurations). For example:

* **compileOnly** configuration for development dependencies
* **compile** configuration that includes compile and runtime dependencies

By default Snyk merges all configurations returned by Gradle to dependency graph based on the sum total of the dependencies across all configurations in the project/projects.

To test a specific configuration:

* Use the `--configuration-matching` option with a [Java regular expression](https://docs.oracle.com/javase/tutorial/essential/regex/) (case-insensitive) as its parameter. The configuration that matches is then tested. If several configurations match, they are all merged and resolved together. For example: `'^releaseRuntimeClasspath$|^compile$'`
* If the different sub-projects include different configurations, scan each sub-project separately.

**Examples of how you can use the --configuration-matching:**

* `--configuration-matching=compile` will match compile, testCompile, compileOnly etc;
* `--configuration-matching=^compile$` will match only compile;
* `--configuration-matching='^(debug|release)compile$'` will match debugCompile and releaseCompile
* `--configuration-matching='^(?!.*test).*$'` will match all configurations _except_ those containing the string "test"

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

By default, Snyk passes `gradle build --no-daemon` in the background when running `snyk test` and `snyk monitor`. If for any reason, you run into trouble, try this:

1. Start the Gradle daemon.
2. Add `--daemon` to your `snyk test` or `snyk monitor`.

### Lockfiles

If your Gradle project makes use of a single **gradle.lockfile** or multiple **\*.lockfile** per configuration and you are having the following issue

**Gradle Error (short): > Could not resolve all dependencies for configuration ':compileOnly'. > Locking strict mode: Configuration ':compileOnly' is locked but does not have lock state.**

Bear in mind that **compileOnly configuration** **has been deprecated** and even if your project successfully generates a lockfile, it will not contain \`compileOnly\` state because this configuration cannot be resolved. Only resolvable configurations compute a dependency graph. In order to solve this issue we suggest you **update your build.gradle containing dependencyLocking logic with the following instruction**

```
compileOnly {resolutionStrategy.deactivateDependencyLocking() }
```

This will **ignore compileOnly** and save only the necessary information to analyse your project/projects.

### Support

If you are having any trouble testing your Gradle projects with Snyk, collect the following details and send them to us at `<`[`support@snyk.io`](mailto:support@snyk.io)`>` so we can help you out:

* `build.gradle`
* `settings.gradle` (especially if we did not pick up a version of a package)
* The output from the following commands:
  * `$ snyk test -d`
  * `$ gradle dependencies -q`

## Git services for Maven projects

After you select a project for import, we build the dependency tree based on the `pom.xml` file.

## Git services for Gradle projects

After you select a project for import, we build the dependency tree based on the `build.gradle` file and (optional) `gradle.lockfile`.

{% hint style="info" %}
Only dependencies in the following configurations are included - `api`, `compile`, `classpath`, `implementation`, `runtime`, `runtimeOnly` .
{% endhint %}

If a lockfile is present, Snyk will use it to accurately resolve the final version of dependencies used in the project.

Gradle lockfiles are an opt-in feature that, among other benefits, enable reproducible builds.Read more about Gradle dependency locking at [https://docs.gradle.org/current/userguide/dependency\_locking.html](https://docs.gradle.org/current/userguide/dependency\_locking.html)

{% hint style="warning" %}
**Kotlin**: `build.gradle.kts` files are not currently supported in Git.
{% endhint %}

## Git settings for Java

From the Snyk UI you can specify mirrors or repositories from which you’d like to resolve packages in Artifactory for Maven.

See the page below for more details on configuring the Artifactory integration.

{% content-ref url="../../../integrations/private-registry-integrations/artifactory-registry-for-maven.md" %}
[artifactory-registry-for-maven.md](../../../integrations/private-registry-integrations/artifactory-registry-for-maven.md)
{% endcontent-ref %}

## Additional Snyk support for Java

In addition to the CLI and Snyk UI features, you can also check your Java projects with these integrations.

{% content-ref url="../../../integrations/ci-cd-integrations/maven-plugin-integration.md" %}
[maven-plugin-integration.md](../../../integrations/ci-cd-integrations/maven-plugin-integration.md)
{% endcontent-ref %}
