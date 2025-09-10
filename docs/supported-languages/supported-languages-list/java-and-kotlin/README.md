# Java and Kotlin

## Applicability

Snyk supports [Java and Kotlin for code analysis](java-and-kotlin-for-code-analysis.md) and [Java and Kotlin for open source](java-and-kotlin-for-open-source.md).

{% hint style="info" %}
**Supported Java versions**

Snyk supports Java analysis for Java versions up to SE 21 and is designed to process code from newer Java versions where feasible.
{% endhint %}

There are special considerations for [Snyk CLI for Java and Kotlin](snyk-cli-for-java-and-kotlin.md) and [SCM integrations with Maven and Gradle](git-repositories-with-maven-and-gradle.md).

[Guidance for Java and Kotlin](guidance-for-java-and-kotlin.md) is available, along with information about the [Snyk workflow with Java and Kotlin](snyk-workflow-with-java-and-kotlin.md) and [More information about Java support](more-information-about-java-support.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:maven`
* Test your app's packages using `pkg:maven`

## Package managers and supported file extensions

Snyk for Java and Kotlin supports Maven and Gradle as package managers and [maven.org](https://maven.org/) (Maven Central Repository) as a package registry.

Use any of the following versions:

* Maven: `3.*.` For more information, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Gradle: `4.*`, `5.*`, `6.*`, `7.*`, `8.*`, `9.*` For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).

Snyk for Java and Kotlin supports the following file formats:

* Snyk Open Source:
  * Maven: `pom.xml`
  * Gradle: `build.gradle`, `build.gradle.kts`
* Snyk Code:
  * Java: `.java`, `.jsp`, `.jspx`
  * Kotlin: `.kt`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Java and Kotlin:

* Amazon AWS SDK - Comprehensive
* Android Standard Library - Partial
* Apache Commons - Comprehensive
* Apache Tomcat - Partial
* Apache XML - Comprehensive
* apache.mahou - Comprehensive
* bouncycastle - Comprehensive
* com.azure.ai.openai - Comprehensive
* com.google.ai.client.generativeai - Comprehensive
* com.google.cloud.vertexai.generativeai - Comprehensive
* com.google.re2j - Comprehensive
* com.google.gwt - Partial
* Dropwizard - Comprehensive
* elasticsearch - Partial
* FasterXML Jackson - Comprehensive
* Google Guava - Comprehensive
* grpc-java - Comprehensive
* hibernate - Comprehensive
* http4k - Comprehensive
* io.jsonwebtoken - Comprehensive
* Jakarta EE - Partial
* Jakarta XML Services - Partial
* Java EE - Partial
* Java Servlet - Comprehensive
* Java Servlet (javax) - Comprehensive
* Java Server Pages - Partial
* Java Standard Edition - Comprehensive
* javalin - Partial
* Jax-RS - Comprehensive
* jooq - Comprehensive
* Kyro - Comprehensive
* Micronaut - Comprehensive
* mongo-java-driver - Comprehensive
* Netty - Comprehensive
* okhttp3 - Comprehensive
* org.apache.hc.client5 - Comprehensive
* org.apache.http.client - Comprehensive
* org.apache.sling - Partial
* org.apache.tools.zip - Comprehensive
* org.codehaus.plexus - Comprehensive
* org.dom4j.io - Comprehensive
* Playframework - Comprehensive
* rxhttp - Comprehensive
* Seam logger - Comprehensive
* SnakeYaml - Comprehensive
* Spongycastle - Comprehensive
* Spring AI - Partial
* Spring boot - Partial
* Spring Web, MVC and JDBC - Comprehensive
* Spring WebFlux - Comprehensive
* Struts - Partial
* Vaadin - Comprehensive
* XStream - Comprehensive

Kotlin only:

* Android Standard Library - Partial
* com.aallam.openai - Comprehensive
* com.expediagroup.graphql.server - Comprehensive
* Javalin - Partial
* Ktor - Comprehensive
* Kotlin Standard Library - Comprehensive
* khttp - Comprehensive

## Features

The following features are supported in Snyk for Java and Kotlin:

| Snyk Open Source                                                           | Snyk Code                                                                                                                                                   |
| -------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <ul><li>Reports</li><li>Fix PRs (Maven)</li><li>License scanning</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis - Kotlin fully supported</li><li>Interfile analysis - Android partially supported</li></ul> |
