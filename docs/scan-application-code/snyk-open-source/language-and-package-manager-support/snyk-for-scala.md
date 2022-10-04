# Snyk for Scala

Snyk supports testing and monitoring Scala projects that have their dependencies managed by [sbt](https://www.scala-sbt.org)

## Features of Snyk support for Scala

The following tables provide an outline of the general features Snyk offers by package manager. In addition to these features, Snyk may offer extra functionality related to the specific integration configurations.

{% hint style="info" %}
Some features might not be available, depending on your pricing plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs |
| --------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [sbt](https://www.scala-sbt.org/) | ✔︎          | ✔︎          | ✔︎               |         |

## How scanning Scala projects works

{% hint style="info" %}
To scan your dependencies, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

Snyk scans Scala projects by running `sbt` plugins or examining your `build.sbt`, and compares the versions of every direct and transitive dependency in your project against Snyk's [Maven vulnerability database](https://snyk.io/vuln?type=maven).

## Snyk CLI tool for Scala projects

{% hint style="warning" %}
The [Snyk CLI ](../../../snyk-cli/)uses the [`sbt-dependency-graph`](https://github.com/sbt/sbt-dependency-graph) plugin which has been [included](https://www.scala-sbt.org/1.x/docs/Combined+Pages.html#sbt-dependency-graph+is+in-sourced) in `sbt` as a built-in plugin since `sbt` 1.4.

However, the recommended method of calling the plugin in sbt 1.4+ is not currently compatible with Snyk. Use the legacy method, `addSbtPlugin()`, instead (see below).
{% endhint %}

You will most likely want to install `sbt-dependency-graph` as a [global plugin](https://www.scala-sbt.org/1.x/docs/Using-Plugins.html#Global+plugins) so you can use it in any `sbt` project.

To do this, add the plugin dependency to `~/.sbt/0.13/plugins/plugins.sbt` for `sbt` 0.13 or `~/.sbt/1.0/plugins/plugins.sbt` for `sbt` 1.0+.

To add the plugin to a single project only, update the `project/plugins.sbt` of your project instead.

Regardless of which `sbt` version you are using, you _must_ use the following command in the relevant `plugins.sbt` file:

`addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1")`

{% hint style="warning" %}
Do not use the `addDependencyTreePlugin` command which the `sbt-dependency-graph` plugin docs recommend for `sbt` 1.4+. This is incompatible with the Snyk CLI. use the `addSbtPlugin()` command as given above.
{% endhint %}

For more details on installing `sbt-dependency-graph` for use with the Snyk CLI, see the [article](https://support.snyk.io/hc/en-us/articles/360004167317) "How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI."

## Git services for Scala projects

Scala `sbt` projects can be imported from any of the Git repositories Snyk [supports](../../../integrations/git-repository-scm-integrations/).

To test your Scala projects using `sbt` as a package manager, Snyk analyzes your `build.sbt` file, and so you must have this file in your repository before importing.
