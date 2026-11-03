---
nav_context: agnostic
description: >-
  Snyk support for Groovy with Snyk Code, available in Early Access on
  Enterprise plans, including supported frameworks and libraries
last_reviewed: 2026-09-14
last_reviewed_source: focused-edit
---

# Groovy

{% hint style="info" %}
Snyk supports Groovy only for Snyk Code.

Code analysis support for Groovy is in Early Access and is available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://docs.snyk.io/platform-administration/snyk-hierarchy/snyk-preview).
{% endhint %}

For an overview of the supported security rules, visit [Groovy rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/groovy-rules).

## Supported frameworks and libraries

For Groovy, Snyk supports the following frameworks and libraries:

* Apache Camel
* groovy-cli-picocli
* Groovy standard library
* groovy-cli-commons
* Grails
* GDK
* Micronaut
* Play
* Spring

## Supported file formats

For Groovy, Snyk supports the following file formats:

* `.groovy`
* `.grt`
* `.gtpl`
* `.gvy`

## Available features

* Reports
* Interfile analysis

## Template file analysis

Snyk Code parses template files and follows data from your application code into the template, so it reports a cross-site scripting vulnerability that becomes exploitable only where the template renders its output.

For Groovy, Snyk Code supports the following template engines:

* FreeMarker
* Thymeleaf
* Velocity

For file extensions and the full set of supported language and template engine combinations, visit [Template file analysis](../technical-specifications-and-guidance.md#template-file-analysis).
