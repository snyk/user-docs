# Unable to test scala project

##  Unable to test scala project

Snyk currently supports vulnerability scanning for Scala / SBT projects using the **build.sbt** file. To read more and get up-to-date support information, please see this page on our Docs: [https://snyk.io/docs/snyk-for-scala/](https://snyk.io/docs/snyk-for-scala/)

To be able to test scala projects, you need to use the **sbt-dependency-graph** plugin for sbt. To read more about sbt-dependency-graph, go here: [https://github.com/jrudolph/sbt-dependency-graph](https://github.com/jrudolph/sbt-dependency-graph)

We recommend that you install it as a global plugin, although it can be installed by project as well.

Install sbt here: [https://www.scala-sbt.org/1.x/docs/Setup.html](https://www.scala-sbt.org/1.x/docs/Setup.html)

Using sbt plugins: [https://www.scala-sbt.org/0.13/docs/Using-Plugins.html](https://www.scala-sbt.org/0.13/docs/Using-Plugins.html)

**NOTE**: Upon initial install of sbt, the _~./sbt/&lt;version&gt;/plugins/plugins.sbt_ file and folder will not exist, you will have to create it.

The _~./sbt/&lt;version&gt;/plugins/plugins.sbt_ file should be updated with the following line:

```text
addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.9.2")
```

