# Scala for open source

## Snyk support for Scala for open source

Refer to the [Scala details](./) for supported package managers and features.

If you need help, [contact Snyk Support](https://support.snyk.io).

## Open source and licensing

The following summarizes Snyk support for sbt.

| Package managers / Features       | CLI support | Git support | License scanning | Fix PRs |
| --------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [sbt](https://www.scala-sbt.org/) | ✔︎          | ✔︎          | ✔︎               |         |

### CLI support for Scala for open source

The [Snyk CLI](../../../developer-tools/snyk-cli/) uses the [`sbt-dependency-graph`](https://github.com/sbt/sbt-dependency-graph) plugin, which has been [included](https://www.scala-sbt.org/1.x/docs/Combined+Pages.html#sbt-dependency-graph+is+in-sourced) in `sbt` as a built-in plugin since `sbt` 1.4.

However, the recommended method of calling the plugin in sbt 1.4+ is not compatible with Snyk. Use the legacy method, `addSbtPlugin()` instead. Snyk recommends installing the `sbt-dependency-graph` as a [global plugin](https://www.scala-sbt.org/1.x/docs/Using-Plugins.html#Global+plugins) so you can use it in any `sbt` project.

To do this, add the plugin dependency to `~/.sbt/0.13/plugins/plugins.sbt` for `sbt` 0.13 or `~/.sbt/1.0/plugins/plugins.sbt` for `sbt` 1.0+.

To add the plugin to a single Project only, update the `project/plugins.sbt` of your Project instead.

Regardless of which `sbt` version you are using, you must use the following command in the relevant `plugins.sbt` file:

`addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1")`

{% hint style="warning" %}
Do not use the `addDependencyTreePlugin` command which the `sbt-dependency-graph` plugin docs recommend for `sbt` 1.4+. This is incompatible with the Snyk CLI.\
Use the `addSbtPlugin()` command as given above.
{% endhint %}

For more information on installing `sbt-dependency-graph` for use with the Snyk CLI, see the article [How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI](https://support.snyk.io/s/article/How-to-install-the-SBT-dependency-graph-plugin-to-test-Scala-projects-with-Snyk-CLI).

### Git repository integration support for Scala for open source

Scala `sbt` Projects can be imported from any of the Git repositories that Snyk [supports](../../../developer-tools/scm-integrations/organization-level-integrations/).

To test your Scala Projects using `sbt` as a package manager, Snyk analyzes your `build.sbt` file.\
To ensure that this works properly, you must have this file in your repository before importing your projects.

You cannot declare versions of dependencies in a file that is not accessible to Snyk using a Source Code Manager (SCM) integration, for example, `Dependencies.scala`.

* To ensure that your Scala dependencies are detected when you import your Projects using an SCM integration, your `build.sbt` dependencies must be declared in a format that Snyk can detect, for example:\
  `"commons-io" % "commons-io" % "2.11.0"`.
* You can use a version declared in a variable if the variable is in the `build.sbt` file, for example:

```scala
  lazy val derbyVersion = "10.4.1.3"
  libraryDependencies ++= Seq(
    "org.apache.derby" % "derby" % derbyVersion
  ) 
```

For more information, see [Differences in Open Source vulnerability counts across environments](../../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/differences-in-open-source-vulnerability-counts-across-environments.md).
