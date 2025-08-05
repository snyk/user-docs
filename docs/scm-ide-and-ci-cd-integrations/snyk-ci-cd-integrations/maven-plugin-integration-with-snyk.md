# Maven plugin integration with Snyk

Snyk offers a [Maven plugin](https://github.com/snyk/snyk-maven-plugin) based on the [Snyk CLI](../../cli-ide-and-ci-cd-integrations/snyk-cli/). This plugin allows you to scan and monitor your Maven dependencies for vulnerabilities.

See all releases in the [Maven Central Repository](https://search.maven.org/artifact/io.snyk/snyk-maven-plugin).

## Installation of Maven plugin

1. Get your [Snyk API token or Snyk PAT](../../snyk-api/authentication-for-api/).
2. Add the Snyk Maven Plugin to your `pom.xml` and configure it as needed.

```xml
<!-- Example Plugin Configuration -->
<build>
  <plugins>
    <plugin>
      <groupId>io.snyk</groupId>
      <artifactId>snyk-maven-plugin</artifactId>
      <version>2.2.0</version>
      <inherited>false</inherited>
      <executions>
        <execution>
          <id>snyk-test</id>
          <goals>
            <goal>test</goal>
          </goals>
        </execution>
        <execution>
          <id>snyk-monitor</id>
          <goals>
            <goal>monitor</goal>
          </goals>
        </execution>
      </executions>
      <configuration>
        <apiToken>${env.SNYK_TOKEN}</apiToken>
        <args>
          <arg>--all-projects</arg>
        </args>
      </configuration>
    </plugin>
  </plugins>
</build>
```

## Supported versions

* Java 8 and above
* Maven 3.2.5 and above

## Goals

### `code-test` (experimental)

Default phase: `test`

Perform a static-analysis of your Project's source code and provide a list of vulnerabilities if any are found.

### `container-test` (experimental)

Default phase: `install`

Perform analysis of the layers of a container image.

Provide the tag of the image to be scanned as an argument:

```xml
<!-- Example of specifying the tag of the image to scan -->
<configuration>
  <args>
    <arg>--print-deps</arg>
    <arg>nginx:1.9.5</arg>
  </args>
</configuration>
```

### `test`

Default Phase: `test`

Scan your Project's dependencies and provide a list of vulnerabilities if any are found.

### `monitor`

Default Phase: `install`

Take a snapshot of your Project's dependency tree and monitor it on [snyk.io](https://snyk.io). You wil be alerted when new relevant vulnerabilities, updates, or patches are disclosed.

## Configuration for the Maven plugin

You can configure the following parameters inside the `<configuration>` section. All parameters are optional.

### `apiToken` \[string]

Do NOT include your API token directly in your `pom.xml`. Use a variable instead.

You must provide a Snyk API token to access Snyk services. You can do so by:

* Providing `apiToken` in your configuration using a variable.
* Providing a `SNYK_TOKEN` environment variable.
* Authenticating using the CLI  `snyk auth` command before using this plugin.

### `skip` \[boolean]

Default: `false`

Skip this execution entirely.

When you are running `mvn`, you can also use `-Dsnyk.skip` to enable this behavior.

### `failOnIssues` \[boolean]

Default: `true`

When this variable is set to `true`, if the Snyk CLI tool indicates that action is required to remedy a security issue, the Maven build will be considered failed. When this variable is set to `false` , the build will continue even if action is required.

### `args` \[array\<string>]

This plugin uses the [Snyk CLI](../../cli-ide-and-ci-cd-integrations/snyk-cli/), so you can pass any supported arguments using `<args>`. See the example that follows.

For a list of supported CLI options, see the [CLI commands and options summary](../../cli-ide-and-ci-cd-integrations/snyk-cli/cli-commands-and-options-summary.md).

```xml
<!-- Example Arguments Configuration -->
<configuration>
  <args>
    <arg>--severity-threshold=high</arg>
    <arg>--scan-all-unmanaged</arg>
    <arg>--json</arg>
  </args>
</configuration>
```

### `cli` \[object]

Lets you configure the Snyk CLI used by this plugin.

By default, the CLI is automatically downloaded and updated for you.

See the CLI configuration section that follows.

## CLI configuration

For most use cases, you need not set any `<cli>` options.

You can configure the CLI in three different modes:

* Auto-Download and Update (default)
* Custom CLI Executable
* Specific CLI Version

Follow the link for each mode to see which parameters are available.

```xml
<!-- Example CLI Configuration -->
<configuration>
  <cli>
    <updatePolicy>daily</updatePolicy>
  </cli>
</configuration>
```

### Auto-Download and Update

#### **`updatePolicy` \[string]**

Default: `daily`

How often to download the latest CLI release. Snyk recommends always keeping your CLI installation updated to the latest version. Can be one of the following:

* `daily` - On the first execution of the day
* `always` - On every execution
* `never` - Never update after the initial download
* `interval:<minutes>` - On the execution after more than a specific number of `<minutes>` has passed since the last update. For example, `interval:60` to update after an hour

#### **`downloadDestination` \[string]**

Default: OS-specific

Where to place the downloaded executable. By default, this is OS-specific as follows:

* Linux - `$XDG_DATA_HOME/snyk/snyk-linux` or `~/.local/share/snyk/snyk-linux`
* macOS - `~/Library/Application Support/Snyk/snyk-macos`
* Windows - `%APPDATA%\Snyk\snyk-win.exe`

### Custom CLI Executable

#### **`executable` \[string]**

Example: `~/.local/share/snyk/snyk-linux`

Path to a pre-installed Snyk CLI executable. You can find executables on the [Snyk CLI releases page](https://github.com/snyk/snyk/releases).

### Specific CLI Version

#### **`version` \[string]**

Example: `1.542.0`

Specify if you want to use a specific version. You can find versions on the [Snyk CLI releases page](https://github.com/snyk/snyk/releases).

Setting this option triggers a download of the CLI on every execution.

## Demonstration

To try out this plugin, see [the demo project](https://github.com/snyk/demo-snyk-maven-plugin).

## Migrating from Snyk Maven Plugin v1 to v2

Move all plugin parameters from v1 to the `<args>` object, to keep them in line with CLI usage. For example:

* `org` => `<arg>--org=my-org-name</arg>`
* `failOnSeverity` => `<arg>--severity-threshold=low|medium|high</arg>`
* `failOnAuthError` => Use `<skip>true</skip>` to skip plugin execution.
* `includeProvidedDependencies` => `provided` to include dependencies always.

***
