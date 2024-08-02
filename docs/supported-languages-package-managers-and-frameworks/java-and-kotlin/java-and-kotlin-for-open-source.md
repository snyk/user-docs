# Java and Kotlin for open source

## Snyk for Java and Kotlin for open source support

**Build tools**: Maven, Gradle

**Build tool versions**:&#x20;

* Maven: `3.*.` For more information, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Gradle: `4.*`, `5.*`, `6.*`, `7.*.` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).

**Package registry**: [maven.org](https://maven.org/)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:maven`

**Test your app's packages**: Available, `pkg:maven`

**Features**:&#x20;

* Fix PRs (Maven)
* License scanning
* Reports

## Open source and licensing

The following summarizes Snyk support for Java and Kotin,

### **Java**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          | ✔︎          | ✔︎               | Fix advice only |

### **Kotlin**

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs         |
| --------------------------------- | ----------- | ----------- | ---------------- | --------------- |
| [Maven](https://maven.apache.org) | ✔︎          | ✔︎          | ✔︎               | ✔︎              |
| [Gradle](https://gradle.org)      | ✔︎          |             | ✔︎               | Fix advice only |

## Snyk support for Maven and Gradle

| Maven                                                                                                               | Gradle                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CLI - Maven `3.*` For details, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support). | CLI - Gradle `4.*`, `5.*`, `6.*`, `7.*`, `8.*` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support). |
| Git - Maven `3.*`                                                                                                   | Git - Gradle `4.*`, `5.*`, `6.*`, `7.*`, `8.*`                                                                                                                |

Gradle Projects imported using a Git integration are tested by parsing `build.gradle` files. You can resolve Gradle dependencies only by executing the tool itself, but even this method can sometimes provide incomplete results.

If possible, enable [lockfiles](java-and-kotlin-for-open-source.md#git-services-for-maven-and-gradle) in your Gradle Project to improve the accuracy of Git imports.

For Gradle projects without lockfiles, Snyk recommends using the Snyk CLI for the most accurate results.

### Support for Gradle with Snyk

If you are having any trouble testing your Gradle Projects with Snyk, collect the following details and send them to Snyk at [support@snyk.io](mailto:support@snyk.io):

* `build.gradle`
* `settings.gradle` (especially if Snyk did not pick up a version of a package)
* The output from the following commands:
  * `$ snyk test -d`
  * `$ gradle dependencies -q`

### Configurations for Gradle Project

Gradle dependencies are declared for a particular scope; each scope is represented by Gradle with the help of [Configurations](https://docs.gradle.org/current/userguide/declaring\_dependencies.html#sec:what-are-dependency-configurations). For example:

* **compileOnly** configuration for development dependencies
* **compileClasspath** configuration that includes compile and runtime dependencies

In most cases, Snyk will include all the dependencies in the **compileClasspath** configuration, but this can vary in some circumstances.

To test a specific configuration:

* Use the `--configuration-matching` option with a [Java regular expression](https://docs.oracle.com/javase/tutorial/essential/regex/) (case-insensitive) as its parameter. The first configuration that matches is tested.
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

* Cannot choose between the following configurations of `project :mymodulewithvariants`
* Cannot choose between the following variants of `project :mymodulewithvariants`
* Could not select value from candidates

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

If your Gradle Project makes use of a single `gradle.lockfile` or multiple `*.lockfile` per configuration, you may see the following issue:

{% code overflow="wrap" %}
```
Gradle Error (short): > Could not resolve all dependencies for configuration ':compileOnly'. > Locking strict mode: Configuration ':compileOnly' is locked but does not have lock state.
```
{% endcode %}

The **compileOnly configuration** **has been deprecated,** and even if your Project successfully generates a lockfile, the `compileOnly` state is not included because this configuration cannot be resolved.&#x20;

Only resolvable configurations compute a dependency graph. To solve this issue, Snyk suggests you update your `build.gradle` containing `dependencyLocking` logic with the following instruction**:**

```
compileOnly {resolutionStrategy.deactivateDependencyLocking() }
```

This ignores the compileOnly and saves only the necessary information to analyze your Project.

If you are having any trouble testing your Gradle Projects with Snyk, collect the following details and send them to Snyk at [support@snyk.io](mailto:support@snyk.io):

* `build.gradle`
* `settings.gradle` (especially if Snyk did not pick up a version of a package)
* The output from the following commands:
  * `$ snyk test -d`
  * `$ gradle dependencies -q`

