# Snyk for Scala

Snyk supports testing Scala projects that have their dependencies managed by [sbt](https://www.scala-sbt.org/). Support is available via the Snyk UI and CLI.

{% hint style="info" %}
**NOTE**  
To use the Snyk CLI with versions of sbt 1.2 and older, you will need to first [install the sbt-dependency-graph plugin](https://support.snyk.io/hc/en-us/articles/360004167317)
{% endhint %}

## Testing Scala projects: how it works

We scan Scala projects by examining your build.sbt to compare the specific versions of every direct and deep dependency in your project against our [Maven vulnerability database](https://snyk.io/vuln?type=maven).

