# Snyk CLI with Maven and Gradle

The Snyk CLI tests Maven and Gradle Projects as follows:

* **Snyk CLI with Gradle**: To build the dependency graph, Snyk integrates with Gradle and inspects the dependencies returned by the build. The following manifest files are supported: `build.gradle` (Groovy DSL) and `build.gradle.kts` (Kotlin DSL).
* **Snyk CLI with Maven**: To build the dependency tree, Snyk analyzes the output of the `pom.xml` files.

The following lists steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. To check the full list, see [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md).&#x20;



| Package manager     | Test help                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Monitor help                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Maven               | <p><code>--maven-aggregate-project</code><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-maven-projects">Options for Maven Projects</a> page for more details.<br><br>Example for aggregate projects:<br><code>snyk test --maven-aggregate-project</code><br><br>Example for non-aggregate projects: <code>snyk test --all-projects</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> Be sure to execute the options in the same directory as the root pom.xml file. </p>                                                                                                                                                                                                                                                                                                                                                                                                    | <p><code>--maven-aggregate-project</code><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-maven-projects">Options for Maven Projects</a> page for more details.<br><br>Example for aggregate projects:<br><code>snyk monitor --maven-aggregate-project</code><br><br>Example for non-aggregate projects: <code>snyk monitor --all-projects</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> Be sure to execute the options in the same directory as the root pom.xml file. </p>                                                                                                                                                                                                                                                                                                                                                                                 |
| Gradle              | <p><code>--sub-project=&#x3C;NAME></code>, <code>--gradle-sub-project=&#x3C;NAME></code> - Test a specific Gradle sub-project.</p><p></p><p><code>--all-sub-projects</code> - Test all Gradle sub-projects.</p><p></p><p><code>--all-projects</code> - Test all Gradle projects.</p><p></p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - Resolve dependencies using only configuration(s) that match the specified Java regular expression.</p><p></p><p>-<code>-configuration-attributes=&#x3C;ATTRIBUTE>[,&#x3C;ATTRIBUTE>]...</code>- Select certain values of configuration attributes to install and resolve dependencies.</p><p></p><p><code>--init-script=&#x3C;FILE></code> - Used for projects with a Gradle initialization script.</p><p></p><p><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-gradle-projects">Options for Gradle Projects </a>page for more details.</p><p></p> | <p><code>--sub-project=&#x3C;NAME></code>, <code>--gradle-sub-project=&#x3C;NAME></code> - Monitor a specific Gradle sub-project.</p><p></p><p><code>--all-sub-projects</code> - Monitor all Gradle sub-projects.</p><p></p><p><code>--all-projects</code> - Monitor all Gradle projects.</p><p></p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - Resolve dependencies using only configuration(s) that match the specified Java regular expression.</p><p></p><p><code>--configuration-attributes=[,]...</code> - Select certain values of configuration attributes to install dependencies and perform dependency resolution.</p><p></p><p><code>--init-script=&#x3C;FILE></code> - Used for projects with a Gradle initialization script.<br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-gradle-projects">Options for Gradle Projects</a> page for more details.</p> |
| Build tools         | <p><code>snyk test -- [&#x3C;context-specific_options>]</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-build-tools">Options for build tools</a> page for more details.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | <p><code>snyk monitor -- [&#x3C;context-specific_options>]</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-build-tools">Options for build tools</a> page for more details.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| Unmanaged JAR files | <p><code>--scan-unmanaged</code> - Test unmanager files <br></p><p><code>--scan-unmanaged --file=&#x3C;JAR_FILE_NAME></code> - Test individual JAR, WAR, and AAR files<br></p><p><code>--scan-all-unmanaged</code> - Auto-detect Maven, JAR, WAR, and AAR files recursively from the current folder.<br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="../../snyk-cli/commands/test.md#scan-all-unmanaged">Options for unmanaged JAR files</a> page for more details. </p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <p><code>--sub-project=&#x3C;NAME></code>, <code>--gradle-sub-project=&#x3C;NAME></code> - Monitor a specific Gradle sub-project.</p><p></p><p><code>--all-sub-projects</code> - Monitor all Gradle sub-projects.</p><p></p><p><code>--all-projects</code> - Monitor all Gradle projects.</p><p></p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - Resolve dependencies using only configuration(s) that match the specified Java regular expression.</p><p></p><p><code>--init-script=&#x3C;FILE></code> - Used for projects with a Gradle initialization script.<br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> See the <a href="../../snyk-cli/commands/monitor.md#scan-all-unmanaged">Options for unmanaged JAR files</a> page for more details.</p>                                                                                                                                                                                             |

## CLI help for Maven Projects

A **Maven aggregate Project** is one that uses modules and inheritance.

When scanning these types of Projects, Snyk performs a compile to ensure all modules are fixable by the Maven reactor.

*   To **scan aggregate Projects**, use the `--maven-aggregate-project` option:

    ```
    snyk test --maven-aggregate-project
    ```
*   To **scan non-aggregate Projects**, use the `--all-projects` option:

    ```
    snyk test --all-projects
    ```

The same options can be used with `snyk monitor`.

Be sure to execute the options in the same directory as the root pom.xml file.

Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.

## **Examples of how to use Maven-specific options with the Snyk CLI**

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

## CLI help for Gradle Projects

Gradle build can consist of several sub-projects, where each sub-project has its own build.gradle, while the root Project is the only one that also includes a `settings.gradle` file. Sub-projects depend on the root ProjectProjects but can be configured otherwise.

By default, Snyk CLI scans only the current Project, the Project in the root of the current folder, or the Project that is specified by `--file=path/to/build.gradle`).

*   To scan all Projects at once (recommended), use the `--all-sub-projects` option:

    ```
    snyk test --all-sub-projects
    ```

Each of the individual sub-projects appears as a separate Snyk Project in the Web UI.

*   To scan a specific Project (for example, myapp):

    ```
    snyk test --sub-project=myapp
    ```

## **Examples of how to use Gradle-specific options with the Snyk CLI**

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
