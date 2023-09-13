# Scala

{% hint style="warning" %}
You might encounter false positives or false negatives for partially supported frameworks and package managers.
{% endhint %}

## Supported frameworks and package managers

### Code analysis

{% hint style="info" %}
Snyk Code for Scala is currently in Open Beta, and you can [enable it through Snyk Preview](../../snyk-admin/manage-settings/snyk-preview.md).
{% endhint %}

Snyk Code supports the following frameworks:

* Play Framework
* Akka
* HTTP4S

### Open source and licensing

Below is a summary of the features offered by Snyk for Scala, organized by package manager. In addition to these features, Snyk may offer additional functionality related to specific integration configurations.

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs |
| --------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [sbt](https://www.scala-sbt.org/) | ✔︎          | ✔︎          | ✔︎               |         |

{% hint style="info" %}
To scan your dependencies, you must ensure you have first installed the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

Snyk scans Scala projects by running `sbt` plugins or examining your `build.sbt` and compares the versions of every direct and transitive dependency in your Project against Snyk's [Maven vulnerability database](https://snyk.io/vuln?type=maven).

## Getting started with Snyk for Scala across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/)
2. [Install Snyk CLI and authenticate your machine](../../snyk-cli/start-using-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](../../scan-application-code/snyk-code/cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests/setting-the-default-organization-for-all-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../scan-application-code/snyk-code/cli-for-snyk-code/excluding-directories-and-files-from-the-snyk-code-cli-test.md)
* [Explore test results in a JSON or SARIF format in the terminal](../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md)
* [Exporting the test results to a JSON or SARIF file](../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md)

#### Open source and licensing

{% hint style="warning" %}
The [Snyk CLI](../../snyk-cli/) uses the [`sbt-dependency-graph`](https://github.com/sbt/sbt-dependency-graph) plugin, which has been [included](https://www.scala-sbt.org/1.x/docs/Combined+Pages.html#sbt-dependency-graph+is+in-sourced) in `sbt` as a built-in plugin since `sbt` 1.4.

However, the recommended method of calling the plugin in sbt 1.4+ is not currently compatible with Snyk. Use the legacy method, `addSbtPlugin()` instead (see the following).
{% endhint %}

Snyk recommends installing the `sbt-dependency-graph` as a [global plugin](https://www.scala-sbt.org/1.x/docs/Using-Plugins.html#Global+plugins) so you can use it in any `sbt` project.

To do this, add the plugin dependency to `~/.sbt/0.13/plugins/plugins.sbt` for `sbt` 0.13 or `~/.sbt/1.0/plugins/plugins.sbt` for `sbt` 1.0+.

To add the plugin to a single Project only, update the `project/plugins.sbt` of your Project instead.

Regardless of which `sbt` version you are using, you must use the following command in the relevant `plugins.sbt` file:

`addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1")`

{% hint style="warning" %}
Do not use the`addDependencyTreePlugin` command which the`sbt-dependency-graph`plugin docs recommend for`sbt` 1.4+. This is incompatible with the Snyk CLI. \
Use the`addSbtPlugin()` command as given above.
{% endhint %}

For more details on installing `sbt-dependency-graph` for use with the Snyk CLI, see the article [How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI.](https://support.snyk.io/hc/en-us/articles/360004167317)

Scala `sbt` Projects can be imported from any of the Git repositories that Snyk [supports](../../integrations/git-repository-scm-integrations/).

To test your Scala Projects using `sbt` as a package manager, Snyk analyzes your `build.sbt` file.\
To ensure that this works properly, you must have this file in your repository before importing your projects.

{% hint style="info" %}
You can’t declare versions of dependencies in a file that is not accessible to Snyk using a Source Code Manager (SCM) integration, for example, `Dependencies.scala`.

* To ensure that your Scala dependencies are detected when you import your Projects using an SCM integration, your `build.sbt` dependencies must be declared in a format that Snyk can detect, for example:\
  `"commons-io" % "commons-io" % "2.11.0"`.
* You can use a version declared in a variable if the variable is in the `build.sbt` file, for example:

```scala
  lazy val derbyVersion = "10.4.1.3"
  libraryDependencies ++= Seq(
    "org.apache.derby" % "derby" % derbyVersion
  ) 
```

For more details, see [Scanning a remote repository using the Snyk Web UI.](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/differences-in-vulnerability-counts-across-environments#scanning-a-remote-repository-using-the-web-ui)
{% endhint %}

### Snyk Web UI (Git repository integration)

Snyk for Scala is not available at the moment for Snyk Web UI. You can use it with the [CLI,](scala.md#snyk-cli) or with the available [Snyk integrations](scala.md#snyk-integrations).

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../integrations/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
