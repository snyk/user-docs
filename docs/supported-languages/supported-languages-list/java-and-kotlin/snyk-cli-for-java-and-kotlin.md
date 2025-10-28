# Snyk CLI for Java and Kotlin

To test Maven and Gradle Projects, use the `snyk test` command as follows:

* Snyk CLI with Gradle: to build the dependency graph, Snyk integrates with Gradle and inspects the dependencies reported by the tool. The following manifest files are supported: `build.gradle` (Groovy DSL) and `build.gradle.kts` (Kotlin DSL).
* Snyk CLI with Maven: to build the dependency tree, Snyk integrates with Maven and inspects the dependencies reported by the tool. The following manifest files are supported: `pom.xml`.

The following table lists options to start scanning your dependencies. It covers the `snyk test` and `snyk monitor` commands. For a list of all the options for these commands, see the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md).&#x20;

| Package manager-environment | Test help                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Monitor help                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Maven                       | <p><code>--maven-aggregate-project</code><br>See <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-maven-projects">Options for Maven Projects</a>.<br><br>Example for aggregate projects:<br><code>snyk test --maven-aggregate-project</code></p><p></p><p>Example for non-aggregate projects: <code>snyk test --all-projects</code> </p><p>Ensure you execute the options in the same directory as the root pom.xml file. </p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | <p><code>--maven-aggregate-project</code><br>See <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-maven-projects">Options for Maven Projects</a>.<br><br>Example for aggregate projects:<br><code>snyk monitor --maven-aggregate-project</code> </p><p></p><p>Example for non-aggregate projects: <code>snyk monitor --all-projects</code><br> </p><p>Ensure you execute the options in the same directory as the root pom.xml file. </p>                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Gradle                      | <p><code>--sub-project=&#x3C;NAME></code>, <code>--gradle-sub-project=&#x3C;NAME></code> - Tests a specific Gradle sub-project.</p><p></p><p><code>--all-sub-projects</code> - Tests all Gradle sub-projects.</p><p></p><p><code>--all-projects</code> - Tests all Gradle projects.</p><p></p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - Resolves dependencies using the first configuration that matches the specified Java regular expression.</p><p></p><p><code>--configuration-attributes=&#x3C;ATTRIBUTE>[,&#x3C;ATTRIBUTE>]...</code> - Selects certain values of configuration attributes to install and resolve dependencies.</p><p></p><p><code>--init-script=&#x3C;FILE></code> - Used for projects with a Gradle initialization script.</p><p></p><p>See <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-gradle-projects">Options for Gradle Projects</a>.</p><p></p> | <p><code>--sub-project=&#x3C;NAME></code>, <code>--gradle-sub-project=&#x3C;NAME></code> - Monitors a specific Gradle sub-project.</p><p></p><p><code>--all-sub-projects</code> - Monitors all Gradle sub-projects.</p><p></p><p><code>--all-projects</code> - Monitors all Gradle projects.</p><p></p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - Resolves dependencies using the first configuration that matches the specified Java regular expression.</p><p></p><p><code>--configuration-attributes=[,]...</code> - Selects certain values of configuration attributes to install dependencies and perform dependency resolution.</p><p></p><p><code>--init-script=&#x3C;FILE></code> - Used for projects with a Gradle initialization script.<br><br>See <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-gradle-projects">Options for Gradle Projects</a>.</p> |
| Build tools                 | <p><code>snyk test -- [&#x3C;context-specific_options>]</code></p><p>See <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-build-tools">Options for build tools</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | <p><code>snyk monitor -- [&#x3C;context-specific_options>]</code><br></p><p>See the <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-build-tools">Options for build tools</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| Unmanaged JAR files         | <p><code>--scan-unmanaged</code> - Tests unmanaged files <br></p><p><code>--scan-unmanaged --file=&#x3C;JAR_FILE_NAME></code> - Tests individual JAR, WAR, and AAR files<br></p><p><code>--scan-all-unmanaged</code> - Auto-detects Maven, JAR, WAR, and AAR files recursively from the current folder.<br><br>See <a href="../../../developer-tools/snyk-cli/commands/test.md#scan-all-unmanaged">Options for unmanaged JAR files</a>. </p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | <p><code>--scan-unmanaged</code> - Tests unmanagedfiles <br></p><p><code>--scan-unmanaged --file=&#x3C;JAR_FILE_NAME></code> - Testsindividual JAR, WAR, and AAR files<br></p><p><code>--scan-all-unmanaged</code> - Auto-detects Maven, JAR, WAR, and AAR files recursively from the current folder.</p><p></p><p>See <a href="../../../developer-tools/snyk-cli/commands/monitor.md#scan-all-unmanaged">Options for unmanaged JAR files</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

## CLI help for Maven Projects

A Maven aggregate Project is one that uses modules and inheritance. When scanning these types of Projects, Snyk performs a compile to ensure all modules are fixable by the Maven reactor.

To scan aggregate Projects, use the `--maven-aggregate-project` option:

```
snyk test --maven-aggregate-project
```

To scan non-aggregate Projects, use the `--all-projects` option:

```
snyk test --all-projects
```

You can use the same options with `snyk monitor`.

Ensure you execute the options in the same directory as the root pom.xml file. Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.

The following example outlines how Maven-specific options are used with the Snyk CLI.

1. Test a specific Maven profile called “prod”.

```
snyk test -- -prod
```

2. Add a system property from your pom.xml file, for example, the package version that appears in your pom.xml:

```
${pkg_version}
```

3. Define the system property:

```
snyk test -- -Dpkg_version=1.4
```

## CLI help for Gradle Projects

Gradle build can consist of several sub-projects, where each sub-project has its own build.gradle, while the root Project is the only one that also includes a `settings.gradle` file. Sub-projects depend on the root ProjectProjects but can be configured otherwise.

By default, Snyk CLI scans only the current Project, the Project in the root of the current folder, or the Project that is specified by `--file=path/to/build.gradle`).

To scan all Projects at once (recommended), use the `--all-sub-projects` option:

```
snyk test --all-sub-projects
```

Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.

To scan a specific Project (for example, "myapp"), use the following command:

```
snyk test --sub-project=myapp
```

### Gradle configurations

Gradle dependencies are declared for a particular scope. Each scope is represented by Gradle with the help of [Configurations](https://docs.gradle.org/current/userguide/declaring_dependencies.html#sec:what-are-dependency-configurations). For example:

* `implementation`: configuration for dependencies required at compile time and runtime, but not exposed to consumers.&#x20;
* `api`**:** configuration for dependencies required at compile time and runtime, and exposed to consumers.&#x20;
* `compileOnly`: configuration for dependencies required only at compile time.&#x20;
* `runtimeOnly`: configuration for dependencies required only at runtime.&#x20;
* `compileClasspath`: configuration for dependencies required at compile time.

In most cases, Snyk includes all the dependencies in the `compileClasspath` configuration, but this can vary in some circumstances.

To test a specific configuration:

* Use the `--configuration-matching` option with a [Java regular expression](https://docs.oracle.com/javase/tutorial/essential/regex/) (case-insensitive) as its parameter. Note that only the first configuration that matches is tested.
* If the different sub-projects include different configurations, scan each sub-project separately.

Examples of how you can use the `--configuration-matching` option:

* `--configuration-matching=compile`  to match `compile`, `testCompile,` `compileOnly`, and so on.
* `--configuration-matching=^compile$` to match only `compile`.
* `--configuration-matching='^(debug|release)compile$'` to match `debugCompile` and `releaseCompile`.
* `--configuration-matching='^(?!.*test).*$'`  to match all configurations except those containing the string "test".

### Android build variants

Android Gradle supports creating different versions of your app by configuring [build variants.](https://developer.android.com/studio/build/build-variants)

Because the Snyk default behavior is to merge all available configurations, the iterated variants cause a clash of configurations that can't be merged.

In these situations, the Snyk scan fails with an error from Gradle which may contain one of the following messages:

* Cannot choose between the following configurations of `project :mymodulewithvariants`
* Cannot choose between the following variants of `project :mymodulewithvariants`
* Could not select value from candidates

To avoid such conflicts:

*   Use a specific configuration(s): if you know of a build configuration that has all the required attributes and the configuration is identical across all sub-projects included in the test, specify that configuration.\
    For example:

    ```
    --configuration-matching=prodReleaseRuntimeClasspath
    ```
*   Explicitly specify the dependency configuration: modify intra-project dependencies in your build.gradle file(s) to use a specific configuration

    ```
      dependencies {
          implementation project(path: ':mymodulewithvariants', configuration: 'default')
      }
    ```
*   Suggest configuration attributes: if you receive an error when running the command, the error can indicate which attribute values are available, while the error details from Gradle also indicate which dependency variants match which attributes. Using these details, add the attribute filter option.\
    For example:

    ```
    snyk test --configuration-attributes=buildtype:release,usage:java-runtime,mode:demo
    ```

    matches the variants using `com.android.build.api.attributes.BuildTypeAttr=release` and `org.gradle.usage=java-runtime`

### Daemon

By default, Snyk passes `gradle build --no-daemon` in the background when running `snyk test` and `snyk monitor` on Windows.&#x20;

If you see `snyk test` or `snyk monitor` fail on other operating systems because of daemon-related issues, try adding the `--no-daemon` flag to the Snyk command or set `GRADLE_OPTS: '-Dorg.gradle.daemon=false'`.&#x20;

For tips on disabling the daemon, see the [Gradle documentation](https://docs.gradle.org/current/userguide/gradle_daemon.html#sec:disabling_the_daemon) .

### Lockfiles

If your Gradle Project uses a single `gradle.lockfile` or multiple `*.lockfile` per configuration, the following issue can appear:

{% code overflow="wrap" %}
```
Gradle Error (short): > Could not resolve all dependencies for configuration ':compileOnly'. > Locking strict mode: Configuration ':compileOnly' is locked but does not have lock state.
```
{% endcode %}

The **compileOnly** configuration has been deprecate&#x64;**,** and even if your Project successfully generates a lockfile, the `compileOnly` state is not included because this configuration cannot be resolved.&#x20;

Only resolvable configurations compute a dependency graph. To solve this issue, Snyk suggests you update your `build.gradle` containing `dependencyLocking` logic with the following instructio&#x6E;**:**

```
compileOnly {resolutionStrategy.deactivateDependencyLocking() }
```

This ignores the `compileOnly` and saves only the necessary information to analyze your Project.

If you have trouble testing your Gradle Projects with Snyk, contact Snyk support and provide the following details:

* `build.gradle`
* `settings.gradle` (especially if Snyk did not pick up a version of a package)
* The output from the following commands:
  * `$ snyk test -d`
  * `$ gradle dependencies -q`

## Workaround for `ant` and `ivy`

[Apache Ant](https://ant.apache.org/) is a Java build system focused solely on executing build tasks defined in XML. [Apache Ivy](https://ant.apache.org/ivy/) extends Ant by adding dependency management, handling library retrieval and transitive dependencies, which Ant alone does not manage.

Ivy dependencies are configured in an XML file, for example `ivy.xml`:

```xml
<ivy-module version="2.0">
    <info organisation="com.example" module="my-project" revision="1.0"/>

    <dependencies>
        <dependency org="junit" name="junit" rev="4.12" conf="default"/>
    </dependencies>
</ivy-module>
```

Such a dependency file is typically evaluated using an `ant` task defined in `build.xml`:

```xml
<target name="resolve-dependencies" depends="init">
    <ivy:retrieve pattern="${lib.dir}/[artifact]-[revision].[ext]"/>
</target>
```

Using the command `ant resolve-dependencies`, dependencies will be downloaded from Maven Central, just like regular Maven dependencies.

To let Snyk know about the dependency tree, you must first convert to the Maven POM format. Start by configuring a new `makepom` task in `build.xml`

```xml
<target name="makepom" depends="resolve-dependencies">
    <ivy:makepom ivyfile="${basedir}/ivy.xml" pomfile="${basedir}/pom.xml" conf="default,runtime">
        <mapping conf="default" scope="compile"/>
        <mapping conf="runtime" scope="runtime"/>
    </ivy:makepom>
</target>
```

With this, you can now run the following commands:

```
ant makepom
snyk test --file=pom.xml
```

The `pom.xml` file does not need to be checked in and can be deleted after a test is done using `snyk`. Additionally, the dependency tree can be monitored using:

```
snyk monitor --file=pom.xml
```

## Snyk CLI tips and tricks

See the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md) and the [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/). Use the `--help` option in the CLI for details of Snyk CLI commands.

### Testing your own code

* Use the `snyk code test` command from the root of the project to perform source code analysis.
* Use `--scan-all-unmanaged --all-projects` to recursively find all jars under the present working directory.

### Testing Open Source libraries

#### **Maven**

The `snyk test` command tests the first manifest it can find, and scans that singular entry point. To scan all manifests, follow these instructions:

* To scan aggregate projects, use the `--maven-aggregate-project` option\
  (for example, `snyk test --maven-aggregate-project`)
* To scan for all projects use `--all-projects` option:\
  (that is, `snyk test --all-projects`)

Snyk scans active profiles activated by default.

* Any additional Maven arguments can be passed, a common one is a non-standard settings.xml location. For example, `snyk test -- -s path/to/settings.xml`
* To scan a specific configuration, test a specific Maven profile using `-P [name]`. For example, use `snyk test -- -P prod` to scan the `prod` configuration.

#### **Gradle**

By default, Snyk CLI scans only the current project (the project in the root of the current folder), or the project that is specified by `--file=path/to/build.gradle`.

`--all-projects` can be used across all package managers, which also includes the behaviors of **--**`all-sub-projects`, mentioned below.

* To scan all projects at once (recommended), use the `--all-sub-projects` option:\
  (that is, `snyk test --all-sub-projects`). Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.
* To scan a specific project (for example, myapp), use `--sub-project=` (that is, `snyk test --sub-project=myapp`).

To test specific configurations, see detailed examples here [Snyk for Java and Kotlin](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/).

For Android Build variants, see [Snyk for Java and Kotlin](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/).

#### **Unmanaged**

For more details on unmanaged Jars, navigate to the [Scan all unmanaged JAR files](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/scan-all-unmanaged-jar-files.md) page.

### Testing containers

Snyk automatically looks for application (such as open source, maven, and npm) vulnerabilities as part of a container scan. Snyk recommends integrating via CLI or Registry earlier in the pipeline and use this as an additional signal or insight into what is in production. See [Snyk CLI for container security](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).

To test Infrastructure as Code, see [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/).

To fix vulnerabilities, see [Fixing vulnerabilities on Maven projects](https://snyk.io/blog/fixing-vulnerabilities-in-maven-projects/).

### Options and plugins

To help generate reports locally or at build time, see [snyk-to-html plugin](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).

See `--json` and `--sarif` options for generating output that can be programmatically accessed.

For advanced filtering options, see[ snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).
