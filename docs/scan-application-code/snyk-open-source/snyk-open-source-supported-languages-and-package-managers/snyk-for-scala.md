# Snyk for Scala

Snyk supports testing and monitoring Scala Projects that have their dependencies managed by [sbt.](https://www.scala-sbt.org)

## Features of Snyk for Scala

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

The following summarizes the features of Snyk for Scala by package manager. In addition to these features, Snyk may offer more functionality related to the specific integration configurations.

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs |
| --------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [sbt](https://www.scala-sbt.org/) | ✔︎          | ✔︎          | ✔︎               |         |

## How Snyk for Scala works

{% hint style="info" %}
To scan your dependencies, you must ensure you have first installed the relevant package manager and that your project contains the supported manifest files.
{% endhint %}

Snyk scans Scala projects by running `sbt` plugins or examining your `build.sbt` and compares the versions of every direct and transitive dependency in your project against Snyk's [Maven vulnerability database](https://snyk.io/vuln?type=maven).

## Snyk CLI for Scala Projects

{% hint style="warning" %}
The [Snyk CLI](../../../snyk-cli/) uses the [`sbt-dependency-graph`](https://github.com/sbt/sbt-dependency-graph) plugin, which has been [included](https://www.scala-sbt.org/1.x/docs/Combined+Pages.html#sbt-dependency-graph+is+in-sourced) in `sbt` as a built-in plugin since `sbt` 1.4.

However, the recommended method of calling the plugin in sbt 1.4+ is not currently compatible with Snyk. Use the legacy method, `addSbtPlugin()` instead (see the following).
{% endhint %}

You will likely want to install `sbt-dependency-graph` as a [global plugin](https://www.scala-sbt.org/1.x/docs/Using-Plugins.html#Global+plugins) so you can use it in any `sbt` project.

To do this, add the plugin dependency to `~/.sbt/0.13/plugins/plugins.sbt` for `sbt` 0.13 or `~/.sbt/1.0/plugins/plugins.sbt` for `sbt` 1.0+.

To add the plugin to a single Project only, update the `project/plugins.sbt` of your Project instead.

Regardless of which `sbt` version you are using, you _must_ use the following command in the relevant `plugins.sbt` file:

`addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1")`

{% hint style="warning" %}
Do not use the `addDependencyTreePlugin` command which the `sbt-dependency-graph` plugin docs recommend for `sbt` 1.4+. This is incompatible with the Snyk CLI. Use the `addSbtPlugin()` command as given above.
{% endhint %}

For more details on installing `sbt-dependency-graph` for use with the Snyk CLI, see the article [How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI.](https://support.snyk.io/hc/en-us/articles/360004167317)

## Git services for Scala Projects

Scala `sbt` Projects can be imported from any of the Git repositories that Snyk [supports](../../../integrations/git-repository-scm-integrations/).

To test your Scala Projects using `sbt` as a package manager, Snyk analyzes your `build.sbt` file.\
To ensure that this works properly, you must have this file in your repository before importing your projects.

{% hint style="info" %}
You can’t declare versions of dependencies in a file that is not accessible to Snyk via a Source Code Manager (SCM) integration, for example, `Dependencies.scala`.

* To ensure that your Scala dependencies are detected when you import your projects using an SCM integration, your `build.sbt` dependencies must be declared in a format that Snyk can detect, for example:\
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
