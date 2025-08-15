# Scala

## Applicability

Snyk supports [Scala for code analysis](scala-for-code-analysis.md) and [Scala for open source](scala-for-open-source.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

{% hint style="info" %}
**Supported Scala versions**

You can use versions 2.x.
{% endhint %}

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.&#x20;
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:maven`
* Test your app's packages using `pkg:maven`

## Package managers and supported file extensions

Snyk for Scala supports sbt as a package manager and [maven.org](https://maven.org/) as a package registry and supports the following file formats:

* Snyk Open Source: `build.sbt`
* Snyk Code: `.scala`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Scala:&#x20;

* Akka - Comprehensive&#x20;
* HTTP4S - Comprehensive&#x20;
* io.cequence.openaiscala - Comprehensive&#x20;
* Play Framework - Comprehensive&#x20;
* Scala standard library - Comprehensive
* Slick Framework - Comprehensive
* All [Java frameworks and libraries](../java-and-kotlin/#frameworks-and-libraries)

## Features

The following features are supported in Snyk for Scala:

| Snyk Open Source                                    | Snyk Code                                            |
| --------------------------------------------------- | ---------------------------------------------------- |
| <ul><li>License scanning </li><li>Reports</li></ul> | <ul><li>Reports</li><li>Interfile analysis</li></ul> |

PR Checks configured to “Only fail when the issues found have a fix available” rely on Snyk FixPR support and will not alert for Scala or other languages that do not support FixPRs.
