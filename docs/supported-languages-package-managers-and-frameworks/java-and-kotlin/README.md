# Java and Kotlin

## Applicability and integration

{% hint style="info" %}
Java and Kotlin are supported for Snyk Code and Snyk Open Source.&#x20;
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

{% hint style="info" %}
**Release status**

Improved Gradle SCM scanning is in Early Access. For more information, see [SCM integrations with Maven and Gradle](git-repositories-with-maven-and-gradle.md).
{% endhint %}

Available functions:

* Test your app's SBOM using `pkg:maven`
* Test your app's packages using `pkg:maven`

## Technical specifications

Snyk supports Java analysis for Java versions up to SE 21 and is designed to process code from newer Java versions where feasible.

### Supported frameworks and libraries

For Java and Kotlin,  the following frameworks and libraries are supported:

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

### Supported package managers and package registries <a href="#supported-package-managers-and-package-registries" id="supported-package-managers-and-package-registries"></a>

* Supported package managers: [Maven](https://maven.apache.org) and [Gradle](https://gradle.org), with the following supported versions:
  * Maven: `3.*` , `4.*`.  For more information, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
  * Gradle: `4.*`, `5.*`, `6.*`, `7.*`, `8.*`, `9*`. For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).
* Supported package registry: [maven.org](https://maven.org/) (Maven Central Repository)

## Java and Kotlin for Snyk Code

For Java and Kotlin with Snyk Code, the following file formats are supported:

* For Java: `.java`, `.jsp`, `.jspx`
* For Kotlin: `.kt`

Available features:

* Reports
* Custom rules
* Interfile analysis - Kotlin is fully supported
* Interfile analysis - Android is partially supported

## Java and Kotlin for Snyk Open Source

For Java and Kotlin with Snyk Open Source, the following file formats are supported:

* For Maven: `pom.xml`
* For Gradle: `build.gradle`, `build.gradle.kts`

Reports are available for Java and Kotlin with Open Source.&#x20;

Available features:

* Reports
* License scanning
* Fix PRs (for Gradle only Fix advice is available)

## Validating, monitoring, alerting, and gating for Java and Kotlin&#x20;

For SCM integrations, Snyk allows you to [run PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/) to validate submitted changes to code and open source packages before merging. Snyk can also retest and alert on the default branch on a scheduled basis. You can see the results on the **Projects** page.

For CI/CD integrations, Snyk can passively monitor and provide a QA gate by failing build checks during testing for policy violations.

Snyk provides flexible capabilities, including:

* [Gradle Plugins](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) (Community project)
* [Maven Plugins](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)
* Dedicated plugins for Jenkins, Circle CI, and others (see relevant marketplaces)
* Using [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/)
* The Snyk CLI can be used in most CI/CD systems (see [examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Fail the build based on criteria using options or the [snyk-filter](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* With Partner Platforms: Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk. For Java, Snyk suggests using the SCM integration with Bitbucket Cloud or using the CLI instead of the prepackaged Bitbucket Pipe.

Snyk can monitor container images and their open source or Linux based packages being used in production using Kubernetes integration, to notify customers of known vulnerabilities for applications in production. This feature is available for Enterprise plans only.

Where a production integration does not exist, use the [snyk monitor](../../developer-tools/snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what is being pushed to production (available for all plans).

## Java support for BOM

Maven supports [bill of materials (BOM) POM files](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html#bill-of-materials-bom-poms) to centralize dependency versions known to work together.

A BOM file includes:

* a `pom` packaging type: `<packaging>pom</packaging>`.
* a `dependencyManagement` section.

Third-party Projects can provide BOM files to make dependency management easier. Here are some common examples:

* [spring-data-bom](https://github.com/spring-projects/spring-data-bom) - The Spring team provides a BOM for their Spring Data Project.
* [jackson-bom](https://github.com/FasterXML/jackson-bom) - The Jackson Project provides a BOM for Jackson dependencies.

Example of a BOM file:

{% code title="Example 1 - BOM file" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <groupId>snyk</groupId>
    <artifactId>snyk-bom</artifactId>
    <version>1.0</version>
    <packaging>pom</packaging>
    <name>Snyk Bill Of Materials</name>
    
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>log4j</groupId>
                <artifactId>log4j</artifactId>
                <version>1.2.12</version>
            </dependency>
            <dependency>
                <groupId>commons-logging</groupId>
                <artifactId>commons-logging</artifactId>
                <version>1.1.1</version>
            </dependency>
        </dependencies>
    </dependencyManagement>
</project>
```
{% endcode %}

The `dependencyManagement` section contains dependency elements. Each dependency is a lookup reference for Maven to determine the version to select for transitive (and direct) dependencies.

Defining a dependency in the `dependencyManagement` section ia used only for lookup reference, it does not add it to the dependency tree of the Project.

You can run `mvn dependency:tree` on the previous BOM example to show that Maven does not treat the contents as dependencies of the file itself.

This BOM can be imported into a Project POM as a parent. You do not need to specify the `log4j` version, as it inherits it from the BOM:

{% code title="Example 2 - Project POM" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>snyk</groupId>
        <artifactId>snyk-bom</artifactId>
        <version>1.0</version>
    </parent>
    
    <groupId>snyk</groupId>
    <artifactId>snyk-project</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <packaging>jar</packaging>
    <name>Snyk Project</name>
    
    <dependency>
        <groupId>log4j</groupId>
        <artifactId>log4j</artifactId>
    </dependency>
</project>
```
{% endcode %}

Snyk applies the versions in the BOM `dependencyManagement` lookup to any dependencies declared in Project POMs that import it as a `parent`.

When Snyk scans the BOM files, the `dependencyManagement` contents are not considered dependencies of that file. These are only lookups.

For the previous examples, Snyk analyzes and treats the files as follows:

* BOM file - Snyk does not create a Snyk Project for this file because it has no dependencies.
* Project POM - Snyk creates a Project with a single dependency of `log4j,` with `v1.2.12`. Snyk applies the rules from the parent BOM to identify the correct version for `log4j`. The dependency `commons-logging` is not included, as it is not directly declared in the Project POM.

{% hint style="info" %}
If a BOM has direct dependencies outside `dependencyManagement`, then Snyk creates a Project for that BOM.
{% endhint %}

Snyk also offers fix advice, including recommendations for upgrading vulnerable packages with the [Fix PR feature](../../scan-with-snyk/pull-requests/#snyk-fix-prs).

Fix PRs can only be created for dependencies whose versions are managed in the POM file where the issue is reported.

If the version or dependency is managed in a parent BOM, then even though Snyk sees that it could fix the vulnerable path by changing the version, it cannot apply the fix.

See additional resources for Java developers on security topics and best practices:

* [Snyk Blog](https://snyk.io/blog/)
* [Securing your modern software supply chain](https://snyk.io/blog/software-supply-chain-security/)
* [Snyk for secure Java development](https://snyk.io/blog/snyk-for-secure-java-development/)
* [Advanced IntelliJ debugger features](https://snyk.io/blog/advanced-intellij-debugger-features/)
* [Spring4shell: the zero day RCE Spring Framework explained](https://snyk.io/blog/spring4shell-zero-day-rce-spring-framework-explained/)
* [Log4j vulnerability explained: Prevent Log4Shell RCE by updating to version 2.17.1](https://snyk.io/blog/log4j-rce-log4shell-vulnerability-cve-2021-44228/)
* [Best practices for managing Java dependencies](https://snyk.io/blog/best-practices-for-managing-java-dependencies/)
* [Exploring the Spring security authorization bypass (CVE-2022-31692)](https://snyk.io/blog/spring-security-authorization-bypass-cve-2022-31692/)
